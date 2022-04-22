package fieldlabs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/pkg/errors"
	"github.com/replicatedhq/replicated/cli/print"
	"github.com/replicatedhq/replicated/pkg/kotsclient"
	"github.com/replicatedhq/replicated/pkg/types"
)

const (
	OneWeek = 7 * 24 * time.Hour
)

type Environment struct {
	// name of the environment
	Name string `json:"name,omitempty" csv:"name"`
	// slug of the environment
	Slug string `json:"slug,omitempty" csv:"slug"`
	// public key of the user who will access this environment
	PubKey string `json:"pub_key,omitempty" csv:"pub_key"`
	// email to invite to vendor.replicated.com if Params.InviteUsers is set
	Email string `json:"email,omitempty" csv:"email"`
	// password to be set on kotsadm instances
	KotsadmPassword string `json:"password,omitempty" csv:"password"`

	App types.App `json:"-" csv:"-"`
}

type ExtraReleaseSpec struct {
	// Dir of YAML sources to promote to channel
	YAMLDir string `json:"yaml_dir"`
	// If set, promote this release to a channel
	PromoteChannel string `json:"promote_channel"`
}

type ExtraReleaseStatus struct {
	Spec    ExtraReleaseSpec
	YAML    string
	Release *types.ReleaseInfo
}

type LabSpec struct {
	// Name of the lab
	Name string `json:"name"`
	// Slug of the lab
	Slug string `json:"slug"`
	// Channel name to create
	Channel string `json:"channel"`
	// Slug of channel name
	ChannelSlug string `json:"channel_slug"`
	// Customer Name to Create
	Customer string `json:"customer"`
	// Dir of YAML sources to promote to channel
	YAMLDir string `json:"yaml_dir"`
	// Path to Installer YAML to promote to channel
	K8sInstallerYAMLPath string `json:"k8s_installer_yaml_path"`
	// List of additional releases to promote
	ExtraReleases []ExtraReleaseSpec `json:"extra_releases"`
	// whether to include a license file and install KOTS
	SkipInstallKots bool `json:"skip_install_kots"`
	// whether to include a license file and install the app
	SkipInstallApp bool `json:"skip_install_app"`
	// kots config values to pass during install
	ConfigValues string `json:"config_values"`
	// bash source to run before installing KOTS
	PreInstallSH string `json:"pre_install_sh"`
	// bash source to run after installing KOTS
	PostInstallSH string `json:"post_install_sh"`

	// add a public ip?
	UsePublicIP bool `json:"use_public_ip"`
	// add a squid proxy?
	UseProxy bool `json:"use_proxy"`
	// add a jump box?
	UseJumpBox bool `json:"use_jump_box"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type Instance struct {
	Name          string `json:"name"`
	Prefix        string `json:"prefix"`
	InstallScript string `json:"provision_sh"`
	MachineType   string `json:"machine_type"`
	BookDiskGB    string `json:"boot_disk_gb"`

	// used to define if a proxy server should be used.
	UseProxy bool `json:"use_proxy"`

	// used in a tf for_each, just put nils in both, the keys and values are ignored
	PublicIps map[string]interface{} `json:"public_ips"`
}

type Lab struct {
	Spec   LabSpec
	Status LabStatus
}

type LabStatus struct {
	InstanceToMake Instance
	Env            Environment

	App           types.App
	Channel       *types.Channel
	Customer      *types.Customer
	Release       *types.ReleaseInfo
	ExtraReleases []ExtraReleaseStatus
	Installer     *types.InstallerSpec
}

type EnvironmentManager struct {
	Log    *print.Logger
	Writer io.Writer
	Params *Params
	Client *kotsclient.VendorV3Client
}

func (e *EnvironmentManager) Validate(envs []Environment, labs []LabSpec) error {
	slug.CustomSub = map[string]string{"_": "-"}

	for _, env := range envs {
		if env.Name == "" {
			return errors.Errorf("no name set for env %v", env)
		}

		if env.Slug == "" {
			return errors.Errorf("no slug set for env %s", env.Name)
		}

		slugified := slug.Make(env.Slug)
		if env.Slug != slugified {
			return errors.Errorf("slugified form of env.Slug %q didn't match provided slug %q", slugified, env.Slug)
		}

		// slug is used as the password to all instances in the environment, so make sure it meets the kots minimum reqs
		if env.KotsadmPassword == "" {
			return errors.Errorf("no kotsadm password set for env %s", env.Name)
		}
	}

	for _, lab := range labs {
		slugifiedChannel := slug.Make(lab.Channel)

		if lab.ChannelSlug == "" {
			lab.ChannelSlug = slugifiedChannel
		}

		if slugifiedChannel != lab.ChannelSlug {
			return errors.Errorf("slugified form of Channel name %q was %q, did not match specified slug %q", lab.Channel, slugifiedChannel, lab.ChannelSlug)
		}
	}

	return nil
}
func (e *EnvironmentManager) Ensure(envs []Environment, labSpecs []LabSpec) error {
	envs, err := e.createApps(envs)
	if err != nil {
		return errors.Wrap(err, "create apps")
	}

	if e.Params.InviteUsers {
		policies, err := e.getPolicies()
		if err != nil {
			return errors.Wrap(err, "get policies")
		}

		err = e.createRBAC(envs, policies)
		if err != nil {
			return errors.Wrap(err, "invite rbac")
		}

		policies, err = e.getPolicies()
		if err != nil {
			return errors.Wrap(err, "get policies")
		}

		err = e.inviteUsers(envs, policies)
		if err != nil {
			return errors.Wrap(err, "invite users")
		}
	}

	labStatuses, err := e.createVendorLabs(envs, labSpecs)
	if err != nil {
		return errors.Wrap(err, "create vendor labs")
	}

	err = e.mergeWriteTFInstancesJSON(labStatuses)
	if err != nil {
		return errors.Wrap(err, "write tf instances json")
	}

	e.Log.ActionWithoutSpinner("Preparing terraform command:")
	fmt.Printf("make instances provisioner_json_out=%q\n", e.Params.InstanceJSONOutput)

	return nil
}

// write TF json for the instances
// merging the new instances with any already present in the
// json file
func (e *EnvironmentManager) mergeWriteTFInstancesJSON(labStatuses []Lab) error {
	bs, err := ioutil.ReadFile(e.Params.InstanceJSONOutput)
	if err != nil && !os.IsNotExist(err) {
		return errors.Wrapf(err, "read file %q", e.Params.InstanceJSONOutput)
	}

	gcpInstances := map[string]Instance{}
	if len(bs) > 0 {
		err := json.Unmarshal(bs, &gcpInstances)
		if err != nil {
			return errors.Wrapf(err, "unmarshal existing instances from %q", e.Params.InstanceJSONOutput)
		}
	}

	for _, labInstance := range labStatuses {
		firstname := strings.Split(labInstance.Status.Env.Name, " ")
		if _, ok := gcpInstances[labInstance.Status.InstanceToMake.Name]; ok {
			e.Log.Error(errors.Errorf("WARNING -- instance %q already present in %q, refusing to overwrite", labInstance.Status.InstanceToMake.Name, e.Params.InstanceJSONOutput))
		}
		gcpInstances[labInstance.Status.InstanceToMake.Name] = labInstance.Status.InstanceToMake
		if labInstance.Spec.UseJumpBox {
			jumpBoxName := fmt.Sprintf("%s-jump", labInstance.Status.InstanceToMake.Name)
			gcpInstances[jumpBoxName] = Instance{
				Name:   jumpBoxName,
				Prefix: e.Params.NamePrefix,
				InstallScript: fmt.Sprintf(`
#!/bin/bash 
# add new user
sudo useradd -s /bin/bash -d /home/%[1]v -m -p safWNrcAGYqm2 -G sudo %[1]v
sudo groups %[1]v
echo '%[1]v ALL=(ALL)        NOPASSWD: ALL' | sudo EDITOR='tee -a' visudo
# update ssh to allow password login
sudo sed -i 's/no/yes/g' /etc/ssh/sshd_config
sudo service ssh restart
# add %[1]v to google-sudoers
sudo usermod -aG google-sudoers,%[1]v %[1]v
# user must change password on first login
#sudo chage --lastday 0 %[1]v

set -euo pipefail

mkdir -p ~/.ssh
cat <<EOF >>~/.ssh/authorized_keys
%s
EOF
`, strings.ToLower(firstname[0]), labInstance.Status.Env.PubKey),
				MachineType: "n1-standard-1",
				BookDiskGB:  "10",
				PublicIps: map[string]interface{}{
					"_": nil,
				},
			}
		}
	}

	serialized, err := json.MarshalIndent(gcpInstances, "", "  ")
	if err != nil {
		return errors.Wrap(err, "serialize instance json")
	}

	err = ioutil.WriteFile(e.Params.InstanceJSONOutput, serialized, 0644)
	if err != nil {
		return errors.Wrapf(err, "write file %q", e.Params.InstanceJSONOutput)
	}

	return nil
}

func (e *EnvironmentManager) createVendorLabs(envs []Environment, labSpecs []LabSpec) ([]Lab, error) {
	var labs []Lab

	for _, env := range envs {
		app := env.App
		firstname := strings.Split(env.Name, " ")
		for _, labSpec := range labSpecs {
			var lab Lab
			lab.Spec = labSpec
			lab.Status.App = app
			lab.Status.Env = env
			appLabSlug := fmt.Sprintf("%s-%s", lab.Status.App.Slug, lab.Spec.Slug)
			e.Log.ActionWithSpinner("Provision lab %s", appLabSlug)

			// load yaml for releases first to ensure directories exist
			kotsYAML, err := readYAMLDir(labSpec.YAMLDir)
			if err != nil {
				return nil, errors.Wrapf(err, "read yaml dir %q", labSpec.YAMLDir)
			}

			for _, extraRelease := range lab.Spec.ExtraReleases {
				kotsYAML, err := readYAMLDir(extraRelease.YAMLDir)
				if err != nil {
					return nil, errors.Wrapf(err, "read yaml dir %q", labSpec.YAMLDir)
				}
				lab.Status.ExtraReleases = append(lab.Status.ExtraReleases, ExtraReleaseStatus{
					Spec: extraRelease,
					YAML: kotsYAML,
				})

			}

			kurlYAML, err := ioutil.ReadFile(labSpec.K8sInstallerYAMLPath)
			if err != nil {
				return nil, errors.Wrapf(err, "read installer yaml %q", labSpec.K8sInstallerYAMLPath)
			}

			channel, err := e.getOrCreateChannel(lab)
			if err != nil {
				return nil, errors.Wrapf(err, "get or create channel %q", lab.Spec.Channel)
			}
			lab.Status.Channel = channel

			customer, err := e.getOrCreateCustomer(lab)
			if err != nil {
				return nil, errors.Wrapf(err, "create customer for lab %q app %q", labSpec.Slug, app.Slug)
			}
			lab.Status.Customer = customer

			release, err := e.Client.CreateRelease(app.ID, kotsYAML)
			if err != nil {
				return nil, errors.Wrapf(err, "create release for %q", labSpec.YAMLDir)
			}

			lab.Status.Release = release

			err = e.Client.PromoteRelease(app.ID, labSpec.Name, labSpec.Slug, release.Sequence, channel.ID)
			if err != nil {
				return nil, errors.Wrapf(err, "promote release %d to channel %q", release.Sequence, channel.Slug)
			}

			for _, extraRelease := range lab.Status.ExtraReleases {
				releaseInfo, err := e.Client.CreateRelease(app.ID, extraRelease.YAML)
				if err != nil {
					return nil, errors.Wrapf(err, "create release for %q", extraRelease.Spec.YAMLDir)
				}
				extraRelease.Release = releaseInfo

				if extraRelease.Spec.PromoteChannel != "" {

					continue
				}
			}

			installer, err := e.Client.CreateInstaller(app.ID, string(kurlYAML))
			if err != nil {
				return nil, errors.Wrapf(err, "create installer from %q", labSpec.K8sInstallerYAMLPath)
			}
			lab.Status.Installer = installer

			err = e.Client.PromoteInstaller(app.ID, installer.Sequence, channel.ID, labSpec.Slug)
			if err != nil {
				return nil, errors.Wrapf(err, "promote installer %d to channel %q", installer.Sequence, channel.Slug)
			}

			licenseContents, err := e.Client.DownloadLicense(app.ID, customer.ID)
			if err != nil {
				return nil, errors.Wrap(err, "download license")
			}

			kotsProvisionScript := fmt.Sprintf(`
curl -fSsL https://k8s.kurl.sh/%s-%s | sudo bash &> kURL.output
`, lab.Status.App.Slug, lab.Spec.ChannelSlug)

			appProvisioningScript := fmt.Sprintf(`
KUBECONFIG=/etc/kubernetes/admin.conf kubectl patch secret kotsadm-tls -p '{"metadata": {"annotations": {"acceptAnonymousUploads": "0"}}}'

KUBECONFIG=/etc/kubernetes/admin.conf kubectl kots install %s-%s \
	--license-file ./license.yaml \
	--namespace default \
	--shared-password %s			
`, lab.Status.App.Slug, lab.Spec.ChannelSlug, env.KotsadmPassword)

			if lab.Spec.SkipInstallKots && lab.Spec.SkipInstallApp {
				kotsProvisionScript = ""
			}

			if lab.Spec.SkipInstallApp {
				appProvisioningScript = ""
			}

			publicIPs := map[string]interface{}{}
			if lab.Spec.UsePublicIP {
				// hack: used in a tf for_each loop, just need something here
				publicIPs["_"] = nil
			}

			lab.Status.InstanceToMake = Instance{
				Name:        appLabSlug,
				Prefix:      e.Params.NamePrefix,
				MachineType: "n1-standard-4",
				BookDiskGB:  "200",
				UseProxy:    lab.Spec.UseProxy,
				PublicIps:   publicIPs,
				InstallScript: fmt.Sprintf(`
#!/bin/bash 
# add new user
sudo useradd -s /bin/bash -d /home/%[1]v -m -p safWNrcAGYqm2 -G sudo %[1]v
sudo groups %[1]v
echo '%[1]v ALL=(ALL)        NOPASSWD: ALL' | sudo EDITOR='tee -a' visudo
# update ssh to allow password login
sudo sed -i 's/no/yes/g' /etc/ssh/sshd_config
sudo service ssh restart
# add %[1]v to google-sudoers
sudo usermod -aG google-sudoers,%[1]v %[1]v
# user must change password on first login
#sudo chage --lastday 0 %[1]v

set -euo pipefail

%s

cat <<EOF > ./license.yaml
%s
EOF

mkdir -p ~/.ssh
cat <<EOF >>~/.ssh/authorized_keys
# added by kots-field-labs
%s
EOF

%s

%s

%s
`, strings.ToLower(firstname[0]), lab.Spec.PreInstallSH, licenseContents, env.PubKey, kotsProvisionScript, appProvisioningScript, lab.Spec.PostInstallSH),
			}
			e.Log.FinishSpinner()
			labs = append(labs, lab)
		}
	}

	return labs, nil
}

func (e *EnvironmentManager) getOrCreateChannel(lab Lab) (*types.Channel, error) {
	channels, err := e.Client.ListChannels(lab.Status.App.ID, lab.Status.App.Slug, lab.Spec.Channel)
	if err != nil {
		return nil, errors.Wrapf(err, "list channel %q for app %q", lab.Spec.Channel, lab.Status.App.Slug)
	}

	var matchedChannels []types.Channel
	for _, channel := range channels {
		if channel.Name == lab.Spec.Channel || channel.Slug == lab.Spec.Channel {
			matchedChannels = append(matchedChannels, channel)
		}
	}

	if len(matchedChannels) == 1 {
		return &matchedChannels[0], nil
	}

	if len(matchedChannels) > 1 {
		return nil, errors.Errorf("expected at most one channel to match %q, found %d", lab.Spec.Channel, len(matchedChannels))
	}

	channel, err := e.Client.CreateChannel(lab.Status.App.ID, lab.Spec.Slug, lab.Spec.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "create channel for lab %q app %q", lab.Spec.Slug, lab.Status.App.Slug)
	}
	return channel, nil
}

func (e *EnvironmentManager) getOrCreateCustomer(lab Lab) (*types.Customer, error) {
	customers, err := e.Client.ListCustomers(lab.Status.App.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "list customer %q for app %q", lab.Spec.Channel, lab.Status.App.Slug)
	}

	for _, customer := range customers {
		if customer.Name == lab.Spec.Customer {
			return &customer, nil
		}
	}

	customer, err := e.Client.CreateCustomer(lab.Spec.Customer, lab.Status.App.ID, lab.Status.Channel.ID, OneWeek)
	if err != nil {
		return nil, errors.Wrapf(err, "create customer for lab %q app %q", lab.Spec.Slug, lab.Status.App.Slug)
	}
	return customer, nil
}

func (e *EnvironmentManager) getAppName(env Environment) string {
	return fmt.Sprintf("%s-%s", e.Params.NamePrefix, env.Slug)
}

func (e *EnvironmentManager) createApps(envs []Environment) ([]Environment, error) {
	var outEnvs []Environment
	var appsCreated []types.App
	for _, env := range envs {
		appName := e.getAppName(env)
		app, err := e.getOrCreateApp(appName)
		if err != nil {
			return nil, errors.Wrapf(err, "get or create app %q", appName)
		}
		env.App = *app
		outEnvs = append(outEnvs, env)
		appsCreated = append(appsCreated, *app)

	}
	_ = e.PrintApps(appsCreated)
	return outEnvs, nil
}

func (e *EnvironmentManager) getOrCreateApp(appName string) (*types.App, error) {
	existingApp, err := e.Client.GetApp(appName)
	if err != nil && !e.isNotFound(err) {
		return nil, errors.Wrapf(err, "check for existing app")
	}
	if existingApp != nil {
		e.Log.ActionWithoutSpinner(fmt.Sprintf("Found Existing app %s", appName))
		return &types.App{
			ID:        existingApp.ID,
			Name:      existingApp.Name,
			Scheduler: "kots",
			Slug:      existingApp.Slug,
		}, nil
	}

	e.Log.ActionWithSpinner(fmt.Sprintf("Creating App %s", appName))
	app, err := e.Client.CreateKOTSApp(appName)
	if err != nil {
		e.Log.FinishSpinnerWithError()
		return nil, errors.Wrapf(err, "create app %s", appName)
	}
	e.Log.FinishSpinner()

	return &types.App{
		ID:        app.Id,
		Name:      app.Name,
		Scheduler: "kots",
		Slug:      app.Slug,
	}, nil
}

func (e *EnvironmentManager) isNotFound(err error) bool {
	return strings.Contains(err.Error(), "App not found")
}

func (e *EnvironmentManager) getPolicies() (map[string]string, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/v1/policies", e.Params.IDOrigin),
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "build policies request")
	}
	req.Header.Set("Authorization", e.Params.SessionToken)
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send policies request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GET /v1/policies %d: %s", resp.StatusCode, body)
	}
	var policies []PolicyListItem
	err = json.Unmarshal([]byte(body), &policies)
	if err != nil {
		return nil, errors.Wrap(err, "list policies unmarshal")
	}

	policiesMap := make(map[string]string)
	for i := 0; i < len(policies); i += 1 {
		policiesMap[policies[i].Name] = policies[i].Id
	}
	return policiesMap, nil
}

func (e *EnvironmentManager) createRBAC(envs []Environment, policies map[string]string) error {
	for _, env := range envs {
		if _, policyExists := policies[e.getAppName(env)]; policyExists {
			// Policy already exists, not recreating
			continue
		}
		policyDefinition := &PolicyDefinition{
			V1: PolicyDefinitionV1{
				Name: "Policy Name",
				Resources: PolicyResourcesV1{
					Allowed: []string{fmt.Sprintf("kots/app/%s/**", env.App.ID), "kots/license/**", "user/token/**"},
					Denied:  []string{},
				},
			},
		}
		policyDefinitionBytes, err := json.Marshal(policyDefinition)
		if err != nil {
			return errors.Wrap(err, "marshal definition body")
		}
		rbacBody := &Policy{
			Name:        e.getAppName(env),
			Description: e.getAppName(env),
			Definition:  string(policyDefinitionBytes),
		}

		rbacBodyBytes, err := json.Marshal(rbacBody)
		if err != nil {
			return errors.Wrap(err, "marshal rbac body")
		}
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/v1/policy", e.Params.IDOrigin),
			bytes.NewReader(rbacBodyBytes),
		)
		if err != nil {
			return errors.Wrap(err, "build rbac request")
		}
		req.Header.Set("Authorization", e.Params.SessionToken)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return errors.Wrap(err, "send rbac request")
		}
		defer resp.Body.Close()
		if resp.StatusCode != 201 {
			body, _ := ioutil.ReadAll(resp.Body)
			return fmt.Errorf("POST /v1/policy %d: %s", resp.StatusCode, body)
		}
	}
	return nil

}

func (e *EnvironmentManager) inviteUsers(envs []Environment, policies map[string]string) error {
	for _, env := range envs {
		if env.Email == "" {
			continue
		}
		inviteBody := map[string]string{
			"email":     env.Email,
			"policy_id": policies[e.getAppName(env)],
		}
		inviteBodyBytes, err := json.Marshal(inviteBody)
		if err != nil {
			return errors.Wrap(err, "marshal invite body")
		}
		req, err := http.NewRequest(
			"POST",
			fmt.Sprintf("%s/v1/team/invite", e.Params.IDOrigin),
			bytes.NewReader(inviteBodyBytes),
		)
		if err != nil {
			return errors.Wrap(err, "build invite request")
		}
		req.Header.Set("Authorization", e.Params.SessionToken)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return errors.Wrap(err, "send invite request")
		}
		defer resp.Body.Close()
		// rate limit returned when already invited
		if resp.StatusCode == 429 {
			e.Log.ActionWithoutSpinner("Skipping invite %q due to 429 error", env.Email)
			continue
		}
		if resp.StatusCode != 204 {
			body, _ := ioutil.ReadAll(resp.Body)
			return fmt.Errorf("POST /team/invite %d: %s", resp.StatusCode, body)
		}
	}
	return nil
}

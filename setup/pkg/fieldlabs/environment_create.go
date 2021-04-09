package fieldlabs

import (
	"encoding/json"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/pkg/errors"
	"github.com/replicatedhq/replicated/cli/print"
	"github.com/replicatedhq/replicated/pkg/kotsclient"
	"github.com/replicatedhq/replicated/pkg/types"
	"io"
	"io/ioutil"
	"time"
)

const (
	OneWeek = 7 * 24 * time.Hour
)

type Environment struct {
	// name of the environment
	Name string
	// slug of the environment
	Slug string
	// public key of the user who will access this environment
	PubKey string
	// email to invite to vendor.replicated.com if Params.InviteUsers is set
	Email string
	// password to be set on kotsadm instances
	KotsadmPassword string

	App types.App
}

type LabSpec struct {
	// Name of the lab
	Name string
	// Slug of the lab
	Slug string
	// Channel name to create
	Channel string
	// Slug of channel name
	ChannelSlug string
	// Customer Name to Create
	Customer string
	// Dir of YAML sources to promote to channel
	YAMLDir string
	// Path to Installer YAML to promote to channel
	K8sInstallerYAMLPath string

	// whether to include a license file and install KOTS
	SkipInstallKots bool
	// kots config values to pass during install
	ConfigValues string
	// bash source to run before installing KOTS
	PreInstallSH string
	// bash source to run after installing KOTS
	PostInstallSH string

	// add a public ip?
	PublicIP bool
	// add a jump box?
	JumpBox bool
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type Instance struct {
	Name          string `json:"name"`
	InstallScript string `json:"provision_sh"`
	MachineType   string `json:"machine_type"`
	BookDiskGB    string `json:"boot_disk_gb"`

	// used in a tf for_each, just put nils in both, the keys and values are ignored
	PublicIps map[string]interface{} `json:"public_ips"`
}

type Lab struct {
	Spec   LabSpec
	Status LabStatus
}

type LabStatus struct {
	InstanceToMake Instance

	App       types.App
	Channel   *types.Channel
	Customer  *types.Customer
	Release   *types.ReleaseInfo
	Installer *types.InstallerSpec
}

type EnvironmentManager struct {
	Log     *print.Logger
	Writer  io.Writer
	Params  *Params
	Client  *kotsclient.VendorV3Client
	GClient *kotsclient.GraphQLClient
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

	labStatuses, err := e.createVendorLabs(envs, labSpecs)
	if err != nil {
		return errors.Wrap(err, "create vendor labs")
	}

	err = e.writeTFInstancesJSON(labStatuses)
	if err != nil {
		return errors.Wrap(err, "write tf instances json")
	}

	return nil
}

// write TF json for the instances
func (e *EnvironmentManager) writeTFInstancesJSON(labStatuses []Lab) error {
	gcpInstances := map[string]Instance{}
	for _, labInstance := range labStatuses {
		gcpInstances[labInstance.Status.InstanceToMake.Name] = labInstance.Status.InstanceToMake
		if labInstance.Spec.JumpBox {
			jumpBoxName := fmt.Sprintf("jump-%s", labInstance.Status.InstanceToMake.Name)
			gcpInstances[jumpBoxName] = Instance{
				Name:          jumpBoxName,
				InstallScript: "",
				MachineType:   "n1-standard-1",
				BookDiskGB:    "10",
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
		for _, labSpec := range labSpecs {
			var lab Lab
			lab.Spec = labSpec
			lab.Status.App = app

			kotsYAML, err := readYAMLDir(labSpec.YAMLDir)
			if err != nil {
				return nil, errors.Wrapf(err, "read yaml dir %q", labSpec.YAMLDir)
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

			customer, err := e.Client.CreateCustomer(labSpec.Customer, app.ID, channel.ID, OneWeek)
			if err != nil {
				return nil, errors.Wrapf(err, "create customer for lab %q app %q", labSpec.Slug, app.Slug)
			}
			lab.Status.Customer = customer

			release, err := e.GClient.CreateRelease(app.ID, kotsYAML)
			if err != nil {
				return nil, errors.Wrapf(err, "create release for %q", labSpec.YAMLDir)
			}

			lab.Status.Release = release

			err = e.GClient.PromoteRelease(app.ID, release.Sequence, labSpec.Slug, labSpec.Name, channel.ID)
			if err != nil {
				return nil, errors.Wrapf(err, "promote release %d to channel %q", release.Sequence, channel.Slug)
			}

			installer, err := e.GClient.CreateInstaller(app.ID, string(kurlYAML))
			if err != nil {
				return nil, errors.Wrapf(err, "create installer from %q", labSpec.K8sInstallerYAMLPath)
			}
			lab.Status.Installer = installer

			err = e.GClient.PromoteInstaller(app.ID, installer.Sequence, channel.ID, labSpec.Slug)
			if err != nil {
				return nil, errors.Wrapf(err, "promote installer %d to channel %q", installer.Sequence, channel.Slug)
			}

			licenseContents, err := e.Client.DownloadLicense(app.ID, customer.ID)
			if err != nil {
				return nil, errors.Wrap(err, "download license")
			}

			kotsProvisionScript := fmt.Sprintf(`
curl -fSsL https://k8s.kurl.sh/%s-%s | sudo bash 

KUBECONFIG=/etc/kubernetes/admin.conf kubectl kots install %s-%s \
  --license-file ./license.yaml \
  --namespace default \
  --shared-password %s
`, lab.Status.App.Slug, lab.Spec.ChannelSlug, lab.Status.App.Slug, lab.Spec.ChannelSlug, env.KotsadmPassword)

			if lab.Spec.SkipInstallKots {
				kotsProvisionScript = ""
			}

			publicIPs := map[string]interface{}{}
			if lab.Spec.PublicIP {
				// hack: used in a tf for_each loop, just need something here
				publicIPs["_"] = nil
			}

			lab.Status.InstanceToMake = Instance{
				Name:        fmt.Sprintf("%s-%s", lab.Status.App.Slug, lab.Spec.Slug),
				MachineType: "n1-standard-4",
				BookDiskGB: "200",
				PublicIps: publicIPs,
				InstallScript: fmt.Sprintf(`
#!/bin/bash 

set -euo pipefail

%s

cat <<EOF > ./license.yaml
%s
EOF

mkdir -p ~/.ssh
cat <<EOF >>~/.ssh/authorized_keys
%s
EOF

%s

%s
`, lab.Spec.PreInstallSH, licenseContents, env.PubKey, kotsProvisionScript, lab.Spec.PostInstallSH),
			}
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
		return nil, errors.New("expected at most one channel to match %q, found %d")
	}

	channel, err := e.GClient.CreateChannel(lab.Status.App.ID, lab.Spec.Slug, lab.Spec.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "create channel for lab %q app %q", lab.Spec.Slug, lab.Status.App.Slug)
	}
	return channel, nil
}

func (e *EnvironmentManager) createApps(envs []Environment) ([]Environment, error) {
	var outEnvs []Environment
	var appsCreated []types.App
	for _, env := range envs {
		appName := fmt.Sprintf("%s-%s", e.Params.NamePrefix, env.Slug)
		e.Log.ActionWithSpinner(fmt.Sprintf("Creating App %s", appName))
		app, err := e.Client.CreateKOTSApp(appName)
		if err != nil {
			e.Log.FinishSpinnerWithError()
			return nil, errors.Wrapf(err, "create app %s", appName)
		}
		e.Log.FinishSpinner()

		outApp := types.App{
			ID:        app.ID,
			Name:      app.Name,
			Scheduler: "kots",
			Slug:      app.Slug,
		}
		env.App = outApp

		outEnvs = append(outEnvs, env)
		appsCreated = append(appsCreated, outApp)

	}
	_ = e.PrintApps(appsCreated)
	return outEnvs, nil
}

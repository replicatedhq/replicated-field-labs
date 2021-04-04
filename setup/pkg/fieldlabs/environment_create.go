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
	Name                 string
	Slug                 string
	Channel              string
	ChannelSlug          string
	Customer             string
	YAMLDir              string
	K8sInstallerYAMLPath string
	ConfigValues         string
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type Instance struct {
	Name          string
	InstallScript string
}

type LabStatus struct {
	Spec           LabSpec
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

	return nil
}
func (e *EnvironmentManager) Ensure(envs []Environment, labs []LabSpec) error {
	envs, err := e.createApps(envs)
	if err != nil {
		return errors.Wrap(err, "create apps")
	}

	labStatuses, err := e.createVendorLabs(envs, labs)
	if err != nil {
		return errors.Wrap(err, "create vendor labs")
	}

	gcpInstances := map[string]string{}
	for _, labInstance := range labStatuses {
		gcpInstances[labInstance.InstanceToMake.Name] = labInstance.InstanceToMake.InstallScript
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

func (e *EnvironmentManager) createVendorLabs(envs []Environment, labs []LabSpec) ([]LabStatus, error) {
	var labStatuses []LabStatus

	for _, env := range envs {
		app := env.App
		for _, labSpec := range labs {
			var lab LabStatus
			lab.Spec = labSpec
			lab.App = app

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
			lab.Channel = channel

			customer, err := e.Client.CreateCustomer(labSpec.Customer, app.ID, channel.ID, OneWeek)
			if err != nil {
				return nil, errors.Wrapf(err, "create customer for lab %q app %q", labSpec.Slug, app.Slug)
			}
			lab.Customer = customer

			release, err := e.GClient.CreateRelease(app.ID, kotsYAML)
			if err != nil {
				return nil, errors.Wrapf(err, "create release for %q", labSpec.YAMLDir)
			}

			lab.Release = release

			err = e.GClient.PromoteRelease(app.ID, release.Sequence, labSpec.Slug, labSpec.Name, channel.ID)
			if err != nil {
				return nil, errors.Wrapf(err, "promote release %d to channel %q", release.Sequence, channel.Slug)
			}

			installer, err := e.GClient.CreateInstaller(app.ID, string(kurlYAML))
			if err != nil {
				return nil, errors.Wrapf(err, "create installer from %q", labSpec.K8sInstallerYAMLPath)
			}
			lab.Installer = installer

			err = e.GClient.PromoteInstaller(app.ID, installer.Sequence, channel.ID, labSpec.Slug)
			if err != nil {
				return nil, errors.Wrapf(err, "promote installer %d to channel %q", installer.Sequence, channel.Slug)
			}

			licenseContents, err := e.Client.DownloadLicense(app.ID, customer.ID)
			if err != nil {
				return nil, errors.Wrap(err, "download license")
			}

			lab.InstanceToMake = Instance{
				Name: fmt.Sprintf("%s-%s", lab.App.Slug, lab.Spec.Slug),
				InstallScript: fmt.Sprintf(`
cat <<EOF > ./license.yaml
%s
EOF

mkdir -p ~/.ssh
cat <<EOF >>~/.ssh/authorized_keys
%s
EOF

curl -fSsL https://k8s.kurl.sh/%s-%s | sudo bash 

KUBECONFIG=/etc/kubernetes/admin.conf kubectl kots install %s-%s \
  --license-file ./license.yaml \
  --namespace default \
  --shared-password %s
`, licenseContents, env.PubKey, lab.App.Slug, lab.Spec.ChannelSlug, lab.App.Slug, lab.Spec.ChannelSlug, env.KotsadmPassword),
			}
			labStatuses = append(labStatuses, lab)
		}
	}

	return labStatuses, nil
}

func (e *EnvironmentManager) getOrCreateChannel(lab LabStatus) (*types.Channel, error) {
	channels, err := e.Client.ListChannels(lab.App.ID, lab.App.Slug, lab.Spec.Channel)
	if err != nil {
		return nil, errors.Wrapf(err, "list channel %q for app %q", lab.Spec.Channel, lab.App.Slug)
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

	channel, err := e.GClient.CreateChannel(lab.App.ID, lab.Spec.Slug, lab.Spec.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "create channel for lab %q app %q", lab.Spec.Slug, lab.App.Slug)
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

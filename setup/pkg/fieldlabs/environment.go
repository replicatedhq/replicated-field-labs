package fieldlabs

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/pkg/errors"
	"github.com/replicatedhq/replicated/cli/print"
	"github.com/replicatedhq/replicated/pkg/kotsclient"
	"github.com/replicatedhq/replicated/pkg/types"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	OneWeek = 7 * 24 * time.Hour
)

type Environment struct {
	Name string
	Slug string
}

type LabSpec struct {
	Name                 string
	Slug                 string
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
	Spec           *LabSpec
	Environment    *Environment
	InstanceToMake *Instance

	App      *types.App
	Channel  *types.Channel
	Customer *types.Customer
	Release  *types.ReleaseInfo
	Installer  *types.InstallerSpec
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
	}

	return nil
}
func (e *EnvironmentManager) Ensure(envs []Environment, labs []LabSpec) error {
	apps, err := e.createApps(envs)
	if err != nil {
		return errors.Wrap(err, "create apps")
	}

	labStatuses, err := e.createVendorLabs(apps, labs)
	if err != nil {
		return errors.Wrap(err, "create vendor labs")
	}

	fmt.Printf(mustJSON(labStatuses))
	return nil
}

func (e *EnvironmentManager) createVendorLabs(apps []types.AppAndChannels, labs []LabSpec) ([]LabStatus, error) {
	var labStatuses []LabStatus

	for _, appChan := range apps {
		app := appChan.App
		for _, labSpec := range labs {

			kotsYAML, err := readYAMLDir(labSpec.YAMLDir)
			if err != nil {
				return nil, errors.Wrapf(err, "read yaml dir %q", labSpec.YAMLDir)
			}
			kurlYAML, err := ioutil.ReadFile(labSpec.K8sInstallerYAMLPath)
			if err != nil {
				return nil, errors.Wrapf(err, "read installer yaml %q", labSpec.K8sInstallerYAMLPath)
			}

			var lab LabStatus
			lab.Spec = &labSpec
			lab.App = app
			channel, err := e.GClient.CreateChannel(app.ID, labSpec.Slug, labSpec.Name)
			if err != nil {
				return nil, errors.Wrapf(err, "create channel for lab %q app %q", labSpec.Slug, app.Slug)
			}
			lab.Channel = channel

			customer, err := e.Client.CreateCustomer(labSpec.Slug, app.ID, channel.ID, OneWeek)
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

			lab.InstanceToMake = &Instance{
				Name: fmt.Sprintf("%s-%s", lab.App.Slug, lab.Spec.Slug),
				InstallScript: fmt.Sprintf(`
cat <<EOF > ./license.yaml
%s
EOF

curl -fSsL https://k8s.kurl.sh/%s-%s | sudo bash

KUBECONFIG=/etc/kubernetes/admin.conf kubectl kots install %s-%s \
  --license-file ./license.yaml \
  --namespace default \
  --shared-password %s
`, licenseContents, lab.App.Slug, lab.Spec.Slug, lab.App.Slug, lab.Spec.Slug, lab.App.Slug),
			}
			labStatuses = append(labStatuses, lab)
		}
	}

	return labStatuses, nil
}

func (e *EnvironmentManager) createApps(envs []Environment) ([]types.AppAndChannels, error) {
	var appsCreated []types.AppAndChannels
	for _, env := range envs {
		appName := fmt.Sprintf("%s-%s", e.Params.NamePrefix, env.Slug)
		e.Log.ActionWithSpinner(fmt.Sprintf("Creating App %s", appName))
		app, err := e.Client.CreateKOTSApp(appName)
		if err != nil {
			e.Log.FinishSpinnerWithError()
			return nil, errors.Wrapf(err, "create app %s", appName)
		}
		e.Log.FinishSpinner()

		appsCreated = append(appsCreated, types.AppAndChannels{
			App: &types.App{
				ID:        app.ID,
				Name:      app.Name,
				Scheduler: "kots",
				Slug:      app.Slug,
			},
			Channels: nil,
		})
	}
	_ = e.PrintApps(appsCreated)
	return appsCreated, nil
}

func (e *EnvironmentManager) Destroy(envs []Environment) error {
	var appsToDelete []types.AppAndChannels

	// find exactly one matching app per env slug
	for _, env := range envs {
		app, err := e.findExactlyOneApp(env)
		if err != nil {
			return errors.Wrapf(err, "find app")
		}
		appsToDelete = append(appsToDelete, *app)
	}

	// print.go apps to delete
	err := e.PrintApps(appsToDelete)
	if err != nil {
		return errors.Wrap(err, "write apps")
	}
	// confirm delete
	answer, err := PromptConfirmDelete()
	if err != nil {
		return errors.Wrap(err, "confirm delete")
	}

	if answer != "yes" {
		return errors.New("prompt declined")
	}

	for _, app := range appsToDelete {
		e.Log.ActionWithSpinner(fmt.Sprintf("Deleting App %s", app.App.Slug))
		err := e.GClient.DeleteKOTSApp(app.App.ID)
		if err != nil {
			e.Log.FinishSpinnerWithError()
			return errors.Wrapf(err, "delete app %q %q", app.App.Slug, app.App.ID)
		}
		e.Log.FinishSpinner()
	}

	return nil
}

func (e *EnvironmentManager) findExactlyOneApp(env Environment) (*types.AppAndChannels, error) {
	var envApps []types.AppAndChannels
	testString := fmt.Sprintf("%s-%s", e.Params.NamePrefix, env.Slug)
	apps, err := e.GClient.ListApps()
	if err != nil {
		return nil, errors.Wrapf(err, "list apps")
	}
	for _, app := range apps {
		if strings.Contains(app.App.Slug, testString) {
			envApps = append(envApps, app)
		}
	}
	if len(envApps) != 1 {
		_ = e.PrintApps(envApps)
		return nil, errors.Errorf("expected exactly one app to match %s, found %d", testString, len(envApps))

	}

	return &envApps[0], nil
}

func readYAMLDir(yamlDir string) (string, error) {

	var allKotsReleaseSpecs []kotsSingleSpec
	err := filepath.Walk(yamlDir, func(path string, info os.FileInfo, err error) error {
		spec, err := encodeKotsFile(yamlDir, path, info, err)
		if err != nil {
			return err
		} else if spec == nil {
			return nil
		}
		allKotsReleaseSpecs = append(allKotsReleaseSpecs, *spec)
		return nil
	})
	if err != nil {
		return "", errors.Wrapf(err, "walk %s", yamlDir)
	}

	jsonAllYamls, err := json.Marshal(allKotsReleaseSpecs)
	if err != nil {
		return "", errors.Wrap(err, "marshal spec")
	}
	return string(jsonAllYamls), nil
}

type kotsSingleSpec struct {
	Name     string   `json:"name"`
	Path     string   `json:"path"`
	Content  string   `json:"content"`
	Children []string `json:"children"`
}

func encodeKotsFile(prefix, path string, info os.FileInfo, err error) (*kotsSingleSpec, error) {
	if err != nil {
		return nil, err
	}

	singlefile := strings.TrimPrefix(filepath.Clean(path), filepath.Clean(prefix)+"/")

	if info.IsDir() {
		return nil, nil
	}
	if strings.HasPrefix(info.Name(), ".") {
		return nil, nil
	}
	ext := filepath.Ext(info.Name())
	switch ext {
	case ".tgz", ".gz", ".yaml", ".yml":
		// continue
	default:
		return nil, nil
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "read file %s", path)
	}

	var str string
	switch ext {
	case ".tgz", ".gz":
		str = base64.StdEncoding.EncodeToString(bytes)
	default:
		str = string(bytes)
	}

	return &kotsSingleSpec{
		Name:     info.Name(),
		Path:     singlefile,
		Content:  str,
		Children: []string{},
	}, nil
}
func mustJSON(v interface{}) string {
	bs, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(bs)
}

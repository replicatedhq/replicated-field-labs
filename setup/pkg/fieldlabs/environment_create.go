package fieldlabs

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"
  "encoding/json"

	"github.com/pkg/errors"
	"github.com/replicatedhq/replicated/pkg/kotsclient"
	"github.com/replicatedhq/replicated/pkg/logger"
	"github.com/replicatedhq/replicated/pkg/types"
)

const (
	OneWeek = 7 * 24 * time.Hour
)

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

type TrackSpec struct {
	// Name of the track
	Name string `json:"name"`
	// Slug of the track
	Slug string `json:"slug"`
	// Channel: if not defined or empty use Stable
	Channel string `json:"channel"`
	// Customer Name to Create
	Customer string `json:"customer"`
	// Dir of YAML sources to promote to channel
	YAMLDir string `json:"yaml_dir"`
	// Path to Installer YAML to promote to channel
	K8sInstallerYAMLPath string `json:"k8s_installer_yaml_path"`
	// List of additional releases to promote
	ExtraReleases []ExtraReleaseSpec `json:"extra_releases"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type Track struct {
	Spec   TrackSpec
	Status TrackStatus
}

type TrackStatus struct {
	App           types.App
	Channel       *types.Channel
	Customer      *types.Customer
	Release       *types.ReleaseInfo
	ExtraReleases []ExtraReleaseStatus
	Installer     *types.InstallerSpec
}

type EnvironmentManager struct {
	Log       *logger.Logger
	Writer    io.Writer
	Params    *Params
	Client    *kotsclient.VendorV3Client
	VendorLoc string
}

func (e *EnvironmentManager) Ensure(track *TrackSpec) error {
	app, err := e.createApp(track.Name)
	if err != nil {
		return errors.Wrap(err, "create apps")
	}

	policies, err := e.getPolicies()
	if err != nil {
		return errors.Wrap(err, "get policies")
	}

	err = e.createRBAC(*app, policies)
	if err != nil {
		return errors.Wrap(err, "invite rbac")
	}

	policies, err = e.getPolicies()
	if err != nil {
		return errors.Wrap(err, "get policies")
	}

	members, err := e.getMembersMap()
	if err != nil {
		return errors.Wrap(err, "get members")
	}

	err = e.addMember(members, policies)
	if err != nil {
		return errors.Wrap(err, "add member")
	}
  memberJson, _ := json.Marshal(members)
  policyJson, _ := json.Marshal(policies)
  e.Log.ActionWithSpinner("Added member: %s %s", memberJson, policyJson)

	err = e.createVendorTrack(*app, *track)
	if err != nil {
		return errors.Wrap(err, "create vendor track")
	}

	return nil
}

func (e *EnvironmentManager) createVendorTrack(app types.App, trackSpec TrackSpec) error {
	var track Track
	track.Spec = trackSpec
	track.Status.App = app
	appTrackSlug := fmt.Sprintf("%s-%s", app.Slug, track.Spec.Slug)
	e.Log.ActionWithSpinner("Provision track %s", appTrackSlug)

  // get the stable channel to assign for customer
	e.Log.ActionWithSpinner("Get stable channel")
  channel, err := e.getChannel(track)
  if err != nil {
    return errors.Wrapf(err, "get Stable channel")
  }
  track.Status.Channel = channel
  e.Log.ActionWithSpinner("Got channel: %s", channel)

  // Create customer
  e.Log.ActionWithSpinner("Create customer")
  if trackSpec.Customer != "" {
    customer, err := e.getOrCreateCustomer(track)
    if err != nil {
      return errors.Wrapf(err, "create customer for track %q app %q", trackSpec.Slug, app.Slug)
    }
    track.Status.Customer = customer
  }

	// load yaml for releases first to ensure directories exist
  if trackSpec.YAMLDir != "" { 
    kotsYAML, err := readYAMLDir(fmt.Sprintf("%s/%s", e.VendorLoc, trackSpec.YAMLDir))

    if err != nil {
      return errors.Wrapf(err, "read yaml dir %q", fmt.Sprintf("%s/%s", e.VendorLoc, trackSpec.YAMLDir))
    }

    for _, extraRelease := range track.Spec.ExtraReleases {
      kotsYAML, err := readYAMLDir(fmt.Sprintf("%s/%s", e.VendorLoc, extraRelease.YAMLDir))
      if err != nil {
        return errors.Wrapf(err, "read yaml dir %q", fmt.Sprintf("%s/%s", e.VendorLoc, trackSpec.YAMLDir))
      }
      track.Status.ExtraReleases = append(track.Status.ExtraReleases, ExtraReleaseStatus{
        Spec: extraRelease,
        YAML: kotsYAML,
      })

    }

    release, err := e.Client.CreateRelease(app.ID, kotsYAML)
    if err != nil {
      return errors.Wrapf(err, "create release for %q", fmt.Sprintf("%s/%s", e.VendorLoc, trackSpec.YAMLDir))
    }

    track.Status.Release = release

    e.Log.ActionWithSpinner("Promote release")
    err = e.Client.PromoteRelease(app.ID, release.Sequence, "0.1.0", trackSpec.Slug, false, channel.ID)
    if err != nil {
      return errors.Wrapf(err, "promote release %d to channel %q", release.Sequence, channel.Slug)
    }

    for _, extraRelease := range track.Status.ExtraReleases {
      releaseInfo, err := e.Client.CreateRelease(app.ID, extraRelease.YAML)
      if err != nil {
        return errors.Wrapf(err, "create release for %q", fmt.Sprintf("%s/%s", e.VendorLoc, extraRelease.Spec.YAMLDir))
      }
      extraRelease.Release = releaseInfo

      if extraRelease.Spec.PromoteChannel != "" {

        continue
      }
    }

    if trackSpec.K8sInstallerYAMLPath != "" {
      kurlYAML, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", e.VendorLoc, trackSpec.K8sInstallerYAMLPath))
      if err != nil {
        return errors.Wrapf(err, "read installer yaml %q", fmt.Sprintf("%s/%s", e.VendorLoc, trackSpec.K8sInstallerYAMLPath))
      }

      installer, err := e.Client.CreateInstaller(app.ID, string(kurlYAML))
      if err != nil {
        return errors.Wrapf(err, "create installer from %q", fmt.Sprintf("%s/%s", e.VendorLoc, trackSpec.K8sInstallerYAMLPath))
      }
      track.Status.Installer = installer

      err = e.Client.PromoteInstaller(app.ID, installer.Sequence, channel.ID, trackSpec.Slug)
      if err != nil {
        return errors.Wrapf(err, "promote installer %d to channel %q", installer.Sequence, channel.Slug)
      }
    }
  }
  e.Log.ActionWithSpinner("Provisioned")
	return nil
}

func (e *EnvironmentManager) getChannel(track Track) (*types.Channel, error) {
	channels, err := e.Client.ListChannels(track.Status.App.ID, track.Spec.Channel)
	if err != nil {
		return nil, errors.Wrapf(err, "list channel %q for app %q", track.Spec.Channel, track.Status.App.Slug)
	}

	var matchedChannels []types.Channel
	for _, channel := range channels {
		if channel.Name == track.Spec.Channel {
			matchedChannels = append(matchedChannels, *channel)
		}
	}

	if len(matchedChannels) == 1 {
		return &matchedChannels[0], nil
	}

	if len(matchedChannels) > 1 {
		return nil, errors.Errorf("expected at most one channel to match %q, found %d", track.Spec.Channel, len(matchedChannels))
	}

	channel, err := e.Client.CreateChannel(track.Status.App.ID, track.Spec.Channel, track.Spec.Channel)
	if err != nil {
		return nil, errors.Wrapf(err, "create channel for track %q app %q", track.Spec.Slug, track.Status.App.Slug)
	}
	return channel, nil

}

func (e *EnvironmentManager) getOrCreateCustomer(track Track) (*types.Customer, error) {
	customers, err := e.Client.ListCustomers(track.Status.App.ID, false)
	if err != nil {
		return nil, errors.Wrapf(err, "list customer for app %q", track.Status.App.Slug)
	}

	for _, customer := range customers {
		if customer.Name == track.Spec.Customer {
			return &customer, nil
		}
	}

  var createOpts = kotsclient.CreateCustomerOpts{ Name: track.Spec.Customer, AppID: track.Status.App.ID, Email: track.Spec.Customer + "@replicated-labs.com", ChannelID: track.Status.Channel.ID, ExpiresAt: OneWeek }
	customer, err := e.Client.CreateCustomer(createOpts) 
	if err != nil {
		return nil, errors.Wrapf(err, "create customer for track %q app %q", track.Spec.Slug, track.Status.App.Slug)
	}
	return customer, nil
}

func (e *EnvironmentManager) createApp(trackName string) (*types.App, error) {
	appName := trackName + " " + e.Params.ParticipantId
	app, err := e.getOrCreateApp(appName)
	if err != nil {
		return nil, errors.Wrapf(err, "get or create app %q", appName)
	}
	return app, nil
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

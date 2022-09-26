package fieldlabs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gosimple/slug"
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

func (e *EnvironmentManager) Validate(track *TrackSpec) error {
	slug.CustomSub = map[string]string{"_": "-"}

	slugifiedChannel := slug.Make(track.Channel)

	if track.ChannelSlug == "" {
		track.ChannelSlug = slugifiedChannel
	}

	if slugifiedChannel != track.ChannelSlug {
		return errors.Errorf("slugified form of Channel name %q was %q, did not match specified slug %q", track.Channel, slugifiedChannel, track.ChannelSlug)
	}

	return nil
}
func (e *EnvironmentManager) Ensure(track *TrackSpec) error {
	app, err := e.createApp()
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

	err = e.inviteUser(members, policies)
	if err != nil {
		return errors.Wrap(err, "invite users")
	}

	err = e.createVendorTrack(*app, *track)
	if err != nil {
		return errors.Wrap(err, "create vendor track")
	}

	return nil
}

func (e *EnvironmentManager) createVendorTrack(app types.App, trackSpec TrackSpec) error {
	if trackSpec.Slug != e.Params.TrackSlug {
		return errors.Errorf("Track with slug %q not found", e.Params.TrackSlug)
	}
	var track Track
	track.Spec = trackSpec
	track.Status.App = app
	appTrackSlug := fmt.Sprintf("%s-%s", app.Slug, track.Spec.Slug)
	e.Log.ActionWithSpinner("Provision track %s", appTrackSlug)

	// load yaml for releases first to ensure directories exist
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

	channel, err := e.getOrCreateChannel(track)
	if err != nil {
		return errors.Wrapf(err, "get or create channel %q", track.Spec.Channel)
	}
	track.Status.Channel = channel

	customer, err := e.getOrCreateCustomer(track)
	if err != nil {
		return errors.Wrapf(err, "create customer for track %q app %q", trackSpec.Slug, app.Slug)
	}
	track.Status.Customer = customer

	release, err := e.Client.CreateRelease(app.ID, kotsYAML)
	if err != nil {
		return errors.Wrapf(err, "create release for %q", fmt.Sprintf("%s/%s", e.VendorLoc, trackSpec.YAMLDir))
	}

	track.Status.Release = release

	err = e.Client.PromoteRelease(app.ID, release.Sequence, trackSpec.Name, trackSpec.Slug, false, channel.ID)
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

	return nil
}

func (e *EnvironmentManager) getOrCreateChannel(track Track) (*types.Channel, error) {
	channels, err := e.Client.ListChannels(track.Status.App.ID, track.Status.App.Slug, track.Spec.ChannelSlug)
	if err != nil {
		return nil, errors.Wrapf(err, "list channel %q for app %q", track.Spec.ChannelSlug, track.Status.App.Slug)
	}

	var matchedChannels []types.Channel
	for _, channel := range channels {
		if channel.Name == track.Spec.ChannelSlug || channel.Slug == track.Spec.ChannelSlug {
			matchedChannels = append(matchedChannels, channel)
		}
	}

	if len(matchedChannels) == 1 {
		return &matchedChannels[0], nil
	}

	if len(matchedChannels) > 1 {
		return nil, errors.Errorf("expected at most one channel to match %q, found %d", track.Spec.Channel, len(matchedChannels))
	}

	channel, err := e.Client.CreateChannel(track.Status.App.ID, track.Spec.ChannelSlug, track.Spec.Channel)
	if err != nil {
		return nil, errors.Wrapf(err, "create channel for track %q app %q", track.Spec.Slug, track.Status.App.Slug)
	}
	return channel, nil
}

func (e *EnvironmentManager) getOrCreateCustomer(track Track) (*types.Customer, error) {
	customers, err := e.Client.ListCustomers(track.Status.App.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "list customer %q for app %q", track.Spec.Channel, track.Status.App.Slug)
	}

	for _, customer := range customers {
		if customer.Name == track.Spec.Customer {
			return &customer, nil
		}
	}

	customer, err := e.Client.CreateCustomer(track.Spec.Customer, track.Status.App.ID, track.Status.Channel.ID, OneWeek)
	if err != nil {
		return nil, errors.Wrapf(err, "create customer for track %q app %q", track.Spec.Slug, track.Status.App.Slug)
	}
	return customer, nil
}

func (e *EnvironmentManager) getAppName() string {
	appName := strings.Replace(e.Params.ParticipantEmail, "@", "-", 1)
	appName = strings.Replace(appName, "+", "-", -1)
	return strings.Replace(appName, ".", "-", -1)
}

func (e *EnvironmentManager) createApp() (*types.App, error) {
	appName := e.getAppName()
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

func (e *EnvironmentManager) getMembersMap() (map[string]MemberList, error) {
	members, err := e.GetMembers()
	if err != nil {
		return nil, errors.Wrap(err, "get members")
	}

	membersMap := make(map[string]MemberList)
	for i := 0; i < len(members); i += 1 {
		membersMap[members[i].Email] = members[i]
	}
	return membersMap, nil
}

func (e *EnvironmentManager) createRBAC(app types.App, policies map[string]string) error {
	if _, policyExists := policies[e.getAppName()]; policyExists {
		// Policy already exists, not recreating
		return nil
	}
	//read + write policy
	policyDefinition := &PolicyDefinition{
		V1: PolicyDefinitionV1{
			Name: "Policy Name",
			Resources: PolicyResourcesV1{
				Allowed: []string{fmt.Sprintf("kots/app/%s/**", app.ID), "kots/license/**", "user/token/**"},
				Denied:  []string{},
			},
		},
	}
	policyDefinitionBytes, err := json.Marshal(policyDefinition)
	if err != nil {
		return errors.Wrap(err, "marshal definition body")
	}
	rbacBody := &Policy{
		Name:        e.getAppName(),
		Description: e.getAppName(),
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

	//readonly policy
	policyDefinition = &PolicyDefinition{
		V1: PolicyDefinitionV1{
			Name: "Policy Name",
			Resources: PolicyResourcesV1{
				Allowed: []string{fmt.Sprintf("kots/app/%s/read", app.ID), "kots/license/**", "user/token/**"},
				Denied:  []string{},
			},
		},
	}
	policyDefinitionBytes, err = json.Marshal(policyDefinition)
	if err != nil {
		return errors.Wrap(err, "marshal definition body")
	}
	rbacBody = &Policy{
		Name:        fmt.Sprintf("%s-readonly", e.getAppName()),
		Description: fmt.Sprintf("%s-readonly", e.getAppName()),
		Definition:  string(policyDefinitionBytes),
	}

	rbacBodyBytes, err = json.Marshal(rbacBody)
	if err != nil {
		return errors.Wrap(err, "marshal rbac body")
	}
	req, err = http.NewRequest(
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

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "send rbac request")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("POST /v1/policy %d: %s", resp.StatusCode, body)
	}
	return nil

}

func (e *EnvironmentManager) updateRBAC(member MemberList, policyId string) error {
	policyUpdateBody := &PolicyUpdate{
		Id: policyId,
	}

	policyUpdateBytes, err := json.Marshal(policyUpdateBody)
	if err != nil {
		return errors.Wrap(err, "marshal policy update body")
	}
	url := fmt.Sprintf("%s/v1/team/member/%s", e.Params.APIOrigin, member.Id)
	if member.Is_Pending_Invite {
		url = fmt.Sprintf("%s/v1/team/invite/%s", e.Params.APIOrigin, member.Email)
	}
	req, err := http.NewRequest(
		"PUT",
		url,
		bytes.NewReader(policyUpdateBytes),
	)
	if err != nil {
		return errors.Wrap(err, "build update policy request")
	}
	req.Header.Set("Authorization", e.Params.SessionToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "send update policy request")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("PUT /v1/team/member %d: %s", resp.StatusCode, body)
	}
	return nil
}

func (e *EnvironmentManager) inviteUser(members map[string]MemberList, policies map[string]string) error {
	inviteEmail := strings.Replace(e.Params.ParticipantEmail, "@", "+replabs@", 1)
	if _, memberExists := members[inviteEmail]; memberExists {
		// Update RBAC policy TODO
		err := e.updateRBAC(members[inviteEmail], policies[e.getAppName()])
		if err != nil {
			return errors.Wrap(err, "update rbac policy")
		}
		return nil
	}
	inviteBody := map[string]string{
		"email":     inviteEmail,
		"policy_id": policies[e.getAppName()],
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
		e.Log.ActionWithoutSpinner("Skipping invite %q due to 429 error", inviteEmail)
		return nil
	}
	if resp.StatusCode != 204 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("POST /team/invite %d: %s", resp.StatusCode, body)
	}
	return nil
}

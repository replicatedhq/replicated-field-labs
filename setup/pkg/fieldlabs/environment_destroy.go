package fieldlabs

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func (e *EnvironmentManager) Destroy(track *TrackSpec) error {
	e.Log.Verbose()
	e.Log.Debug("Destroying environment %s", e.Params.ParticipantId)
	e.Log.Debug("Get policies")
	policies, err := e.getPolicies()
	if err != nil {
		return errors.Wrap(err, "get policies")
	}

	e.Log.Debug("Get member map")
	members, err := e.getMembersMap()
	if err != nil {
		return errors.Wrap(err, "get members")
	}
	inviteEmail := e.Params.ParticipantId + "@replicated-labs.com"

	memberJson, _ := json.Marshal(members)
	e.Log.Debug("Need to find %s in: %s", inviteEmail, memberJson)
	e.Log.Debug("Delete member %s", inviteEmail)
	err = e.DeleteMember(members[inviteEmail].Id)
	if err != nil {
		return err
	}
	e.Log.Debug("Delete policy id %s", policies[e.Params.ParticipantId])
	err = e.DeletePolicyId(policies[e.Params.ParticipantId])
	if err != nil {
		return err
	}

	// Delete the app
	e.Log.Debug("List apps")
	apps, err := e.Client.ListApps()
	if err != nil {
		return errors.Wrapf(err, "list apps")
	}
	for _, app := range apps {
		if app.App.Name == track.Name+" "+e.Params.ParticipantId {
      e.Log.Debug("Delete app %s", app.App.Name)
			err = e.Client.DeleteKOTSApp(app.App.ID)
			if err != nil {
				return errors.Wrapf(err, "delete app")
			}
			break
		}
	}
	return nil
}

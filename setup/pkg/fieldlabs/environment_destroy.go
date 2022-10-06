package fieldlabs

import (
	"github.com/pkg/errors"
)

func (e *EnvironmentManager) Destroy() error {
	policies, err := e.getPolicies()
	if err != nil {
		return errors.Wrap(err, "get policies")
	}
	members, err := e.getMembersMap()
	if err != nil {
		return errors.Wrap(err, "get members")
	}
	inviteEmail := e.Params.ParticipantId + "@replicated-labs.com"

	err = e.DeleteMember(members[inviteEmail].Id)
	if err != nil {
		return err
	}
	err = e.DeletePolicyId(policies[e.Params.ParticipantId])
	if err != nil {
		return err
	}

	// Delete the app
	apps, err := e.Client.ListApps()
	if err != nil {
		return errors.Wrapf(err, "list apps")
	}
	for _, app := range apps {
		if app.App.Name == e.Params.ParticipantId {
			err = e.Client.DeleteKOTSApp(app.App.ID)
			if err != nil {
				return errors.Wrapf(err, "delete app")
			}
			break
		}
	}
	return nil
}

package fieldlabs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Get list of team members created with multi-player mode
func (e *EnvironmentManager) GetMembers() ([]MemberList, error) {
	url := fmt.Sprintf("%s/v1/team/members", e.Params.IDOrigin)
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)

	if err != nil {
		return []MemberList{}, err
	}
	req.Header.Set("Authorization", e.Params.SessionToken)
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return []MemberList{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	if resp.StatusCode != 200 {
		return []MemberList{}, fmt.Errorf("GET /v1/team/members %d: %s", resp.StatusCode, body)
	}

	var members []MemberList
	err = json.Unmarshal([]byte(body), &members)
	if err != nil {
		return []MemberList{}, err
	}

	return members, nil
}

// Delete team members created with multi-player mode
func (e *EnvironmentManager) DeleteMember(id string) error {
	url := fmt.Sprintf("%s/v1/team/member?user_id=%s", e.Params.IDOrigin, id)
	req, err := http.NewRequest(
		"DELETE",
		url,
		nil,
	)

	if err != nil {
		return err
	}
	req.Header.Set("Authorization", e.Params.SessionToken)
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("GET /v1/team/member %d: %s", resp.StatusCode, body)
	}
	return nil
}

// Delete policies create through multi-player mode
func (e *EnvironmentManager) DeletePolicyId(id string) error {
	url := fmt.Sprintf("%s/v1/policy/%s", e.Params.IDOrigin, id)
	req, err := http.NewRequest(
		"DELETE",
		url,
		nil,
	)

	if err != nil {
		return err
	}
	req.Header.Set("Authorization", e.Params.SessionToken)
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("GET /v1/policy %d: %s", resp.StatusCode, body)
	}
	return nil
}

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

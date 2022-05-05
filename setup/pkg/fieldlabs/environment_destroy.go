package fieldlabs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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

func (e *EnvironmentManager) Destroy() error {
	policies, err := e.getPolicies()
	if err != nil {
		return errors.Wrap(err, "get policies")
	}
	members, err := e.getMembersMap()
	if err != nil {
		return errors.Wrap(err, "get members")
	}
	inviteEmail := strings.Replace(e.Params.ParticipantEmail, "@", "+labs@", 1)
	if err != nil {
		return errors.Wrap(err, "get policies")
	}
	err = e.updateRBAC(members[inviteEmail], policies[fmt.Sprintf("%s-readonly", e.getAppName())])
	if err != nil {
		return errors.Wrap(err, "update rbac policy")
	}

	return nil
}

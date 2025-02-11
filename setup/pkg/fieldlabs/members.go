package fieldlabs

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type MemberList struct {
	Id                string `json:"id"`
	Email             string `json:"email"`
	Is_Pending_Invite bool   `json:"is_pending_invite"`
}

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

func (e *EnvironmentManager) getMembersMap() (map[string]MemberList, error) {
	members, err := e.GetMembers()
	if err != nil {
		return nil, errors.Wrap(err, "get members")
	}

	membersJson, _ := json.Marshal(members)
	fmt.Sprintf("members: %s", membersJson)

	membersMap := make(map[string]MemberList)
	for i := 0; i < len(members); i += 1 {
		memberJson, _ := json.Marshal(members)
		fmt.Sprintf("member: %s", memberJson)
		fmt.Sprintf("member: %s", members[i].Email)
		membersMap[members[i].Email] = members[i]
	}
	return membersMap, nil
}

// Delete team members created with multi-player mode
func (e *EnvironmentManager) DeleteMember(id string) error {
	requestUrl := fmt.Sprintf("%s/v1/team/member?id=%s", e.Params.IDOrigin, id)
	req, err := http.NewRequest(
		"DELETE",
		requestUrl,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("DELETE %s", requestUrl))
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
    return fmt.Errorf("DELETE %s %d: %s", requestUrl, resp.StatusCode, body)
	}
	return nil
}

func (e *EnvironmentManager) addMember(members map[string]MemberList, policies map[string]string) error {
	inviteEmail := e.Params.ParticipantId + "@replicated-labs.com"
	err := e.inviteMember(inviteEmail, members, policies)
	if err != nil {
		return err
	}

	// Signup
	sr, err := e.signupMember(inviteEmail)
	if err != nil {
		return err
	}

	// Verify
	vr, err := e.verifyMember(sr)
	if err != nil {
		return err
	}

	// Capture Invite Id
	invite, err := e.captureInvite(vr)
	if err != nil {
		return err
	}

	// Accept Invite
	err = e.acceptInvite(invite, e.Params.ParticipantId, vr)
	if err != nil {
		return err
	}
	return nil
}


type VerifyResponse struct {
	Token string `json:"token"`
}

type SignupResponse struct {
	Token string `json:"token"`
}

func (e *EnvironmentManager) signupMember(inviteEmail string) (*SignupResponse, error) {
	signupBody := map[string]string{
		"email": inviteEmail,
	}
	signupBodyBytes, err := json.Marshal(signupBody)
	if err != nil {
		return nil, errors.Wrap(err, "marshal signup body")
	}
  requestUrl := fmt.Sprintf("%s/vendor/v1/signup", e.Params.IDOrigin)
	req, err := http.NewRequest(
		"POST",
    requestUrl,
		bytes.NewReader(signupBodyBytes),
	)
	if err != nil {
		return nil, errors.Wrap(err, "build signup request")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send signup request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("POST %s %d: %s", requestUrl, resp.StatusCode, body)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body")
	}
	var body SignupResponse
	if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&body); err != nil {
		return nil, errors.Wrap(err, "decode body")
	}
	return &body, nil

}

func (e *EnvironmentManager) inviteMember(inviteEmail string, members map[string]MemberList, policies map[string]string) error {
	if _, memberExists := members[inviteEmail]; memberExists {
		// This should never happen?
		return nil
	}
	inviteBody := map[string]string{
		"email":     inviteEmail,
		"policy_id": policies[e.Params.ParticipantId],
	}
	inviteBodyBytes, err := json.Marshal(inviteBody)
	if err != nil {
		return errors.Wrap(err, "marshal invite body")
	}
  requestUrl := fmt.Sprintf("%s/vendor/v1/team/invite", e.Params.IDOrigin)
	req, err := http.NewRequest(
		"POST",
    requestUrl,
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
    return errors.Wrap(err, fmt.Sprintf("send invite request: %s", requestUrl))
	}
	defer resp.Body.Close()
	// rate limit returned when already invited
	if resp.StatusCode == 429 {
		e.Log.ActionWithoutSpinner("Skipping invite %q due to 429 error", inviteEmail)
		return nil
	}
	if resp.StatusCode != 204 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("POST %s %d: %s", requestUrl, resp.StatusCode, body)
	}
	return nil
}

type InvitedTeams struct {
	Teams []struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		InviteId string `json:"invite_id"`
	} `json:"invited_teams"`
}

func (e *EnvironmentManager) verifyMember(sr *SignupResponse) (*VerifyResponse, error) {
	verifyBody := map[string]string{
		"token": sr.Token,
	}
	verifyBodyBytes, err := json.Marshal(verifyBody)
	if err != nil {
		return nil, errors.Wrap(err, "marshal verify body")
	}
  requestUrl := fmt.Sprintf("%s/vendor/v1/signup/verify", e.Params.IDOrigin)
	req, err := http.NewRequest(
		"POST",
    requestUrl,
		bytes.NewReader(verifyBodyBytes),
	)
	if err != nil {
    return nil, errors.Wrap(err, fmt.Sprintf("build verify request: %s", requestUrl))
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send verify request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		body, _ := ioutil.ReadAll(resp.Body)
    return nil, fmt.Errorf("POST %s %d: %s", requestUrl, resp.StatusCode, body)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body")
	}
	var body VerifyResponse
	if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&body); err != nil {
		return nil, errors.Wrap(err, "decode body")
	}
	return &body, nil
}

func (e *EnvironmentManager) captureInvite(vr *VerifyResponse) (*InvitedTeams, error) {
	e.Log.Verbose()
  requestUrl := fmt.Sprintf("%s/vendor/v1/signup/teams", e.Params.IDOrigin)
	req, err := http.NewRequest(
		"GET",
    requestUrl,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "build signup teams request")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", vr.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "getting the invite")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
    return nil, fmt.Errorf("GET %s %d: %s", requestUrl, resp.StatusCode, body)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body")
	}
	var body InvitedTeams
	if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&body); err != nil {
		return nil, errors.Wrap(err, "decode body")
	}
	return &body, nil
}

type AcceptBody struct {
	InviteId          string `json:"invite_id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Password          string `json:"password"`
	ReplaceAccount    bool   `json:"replace_account"`
	FromTeamSelection bool   `json:"from_team_selection"`
}

func (e *EnvironmentManager) acceptInvite(invite *InvitedTeams, participantId string, vr *VerifyResponse) error {
	h := sha256.Sum256([]byte(participantId))
	sum := fmt.Sprintf("%x", h)
	ab := AcceptBody{InviteId: (*invite).Teams[0].InviteId, FirstName: "Instruqt", LastName: "Participant", Password: string(sum[0:20]), ReplaceAccount: false, FromTeamSelection: true}
	acceptBodyBytes, err := json.Marshal(ab)
	if err != nil {
		return errors.Wrap(err, "marshal accept body")
	}

  requestUrl := fmt.Sprintf("%s/vendor/v1/signup/accept-invite", e.Params.IDOrigin)
	req, err := http.NewRequest(
		"POST",
    requestUrl,
		bytes.NewReader(acceptBodyBytes),
	)
	if err != nil {
		return errors.Wrap(err, "build accept request")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
    return errors.Wrap(err, fmt.Sprintf("send accept request: %s", requestUrl))
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		body, _ := ioutil.ReadAll(resp.Body)
    return fmt.Errorf("POST %s %d: %s", requestUrl, resp.StatusCode, body)
	}
	return nil
}

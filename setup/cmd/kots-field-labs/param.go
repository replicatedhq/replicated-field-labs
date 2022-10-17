package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/replicatedhq/kots-field-labs/setup/pkg/fieldlabs"
)

var actions = map[string]fieldlabs.Action{
	"create":  fieldlabs.ActionCreate,
	"destroy": fieldlabs.ActionDestroy,
}

func missingParam(s string) error {
	return errors.New(fmt.Sprintf("Missing or invalid parameters: %s", s))
}

func GetParams() (*fieldlabs.Params, error) {
	params := &fieldlabs.Params{
		ParticipantId:   os.Getenv("PARTICIPANT_ID"),
		Branch:          os.Getenv("REPLICATED_BRANCH"),
		TrackSlug:       os.Getenv("REPLICATED_TRACK_SLUG"),
		InviterEmail:    os.Getenv("REPLICATED_INVITER_EMAIL"),
		InviterPassword: os.Getenv("REPLICATED_INVITER_PASSWORD"),
		APIToken:        os.Getenv("REPLICATED_API_TOKEN"),
		APIOrigin:       os.Getenv("REPLICATED_API_ORIGIN"),
		GraphQLOrigin:   os.Getenv("REPLICATED_GRAPHQL_ORIGIN"),
		KURLSHOrigin:    os.Getenv("REPLICATED_KURLSH_ORIGIN"),
		IDOrigin:        os.Getenv("REPLICATED_ID_ORIGIN"),
	}

	if params.ParticipantId == "" {
		return nil, missingParam("PARTICIPANT_ID")
	}

	if params.Branch == "" {
		params.Branch = "main"
	}

	if params.TrackSlug == "" {
		return nil, missingParam("REPLICATED_TRACK_SLUG")
	}

	if params.APIToken == "" {
		return nil, missingParam("REPLICATED_API_TOKEN")
	}
	if params.APIOrigin == "" {
		params.APIOrigin = "https://api.replicated.com/vendor"
	}
	if params.GraphQLOrigin == "" {
		params.GraphQLOrigin = "https://g.replicated.com/graphql"
	}
	if params.KURLSHOrigin == "" {
		params.KURLSHOrigin = "https://kurl.sh"
	}
	if params.IDOrigin == "" {
		params.IDOrigin = "https://id.replicated.com"
	}

	actionString := os.Getenv("REPLICATED_ACTION")
	if actionString == "" {
		actionString = "create"
	}

	action, ok := actions[actionString]
	if !ok {
		return nil, errors.Errorf("unkown action %s", actionString)
	}
	params.Action = action

	err := getSessionTokenAndCheckInviteParams(params)
	if err != nil {
		return nil, errors.Wrap(err, "validate invite user params")
	}

	return params, nil
}

func getSessionTokenAndCheckInviteParams(params *fieldlabs.Params) error {
	err2 := validateInviteParams(params)
	if err2 != nil {
		return err2
	}

	sessionToken, err := getLoginResponse(params)
	if err != nil {
		return errors.Wrap(err, "get session token")
	}

	params.SessionToken = *sessionToken

	return nil
}

func validateInviteParams(params *fieldlabs.Params) error {
	if params.InviterEmail == "" {
		return errors.Errorf("REPLICATED_INVITER_EMAIL must be set if action is create")
	}
	if params.InviterPassword == "" {
		return errors.Errorf("REPLICATED_INVITER_PASSWORD must be set if action is create")
	}
	return nil
}

func getLoginResponse(params *fieldlabs.Params) (*string, error) {
	loginParams := map[string]string{
		"email":    params.InviterEmail,
		"password": params.InviterPassword,
	}

	loginBody, err := json.Marshal(loginParams)
	if err != nil {
		return nil, errors.Wrap(err, "marshal login params")
	}

	loginReq, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/login", params.IDOrigin), bytes.NewBuffer(loginBody))
	if err != nil {
		return nil, errors.Wrap(err, "build login request")
	}
	loginReq.Header.Set("Accept", "application/json")
	loginReq.Header.Set("Content-Type", "application/json")
	loginResp, err := http.DefaultClient.Do(loginReq)
	if err != nil {
		return nil, errors.Wrap(err, "send login request")
	}

	defer loginResp.Body.Close()
	if loginResp.StatusCode != 201 {
		body, _ := ioutil.ReadAll(loginResp.Body)
		return nil, fmt.Errorf("GET /policies %d: %s", loginResp.StatusCode, body)
	}
	bodyBytes, err := ioutil.ReadAll(loginResp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body")
	}
	var body SessionResponse
	if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&body); err != nil {
		return nil, errors.Wrap(err, "decode body")
	}
	return body.Token, nil
}

type SessionResponse struct {
	Token *string `json:"token"`
}

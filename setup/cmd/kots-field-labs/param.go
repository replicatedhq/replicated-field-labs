package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/replicatedhq/kots-field-labs/setup/pkg/fieldlabs"
	"io/ioutil"
	"net/http"
	"os"
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
		NamePrefix:         os.Getenv("REPLICATED_NAME_PREFIX"),
		EnvironmentsJSON:   os.Getenv("REPLICATED_ENVIRONMENTS_JSON"),
		LabsJSON:           os.Getenv("REPLICATED_LABS_JSON"),
		InstanceJSONOutput: os.Getenv("REPLICATED_INSTANCE_JSON_OUT"),
		InviteUsers:        os.Getenv("REPLICATED_INVITE_USERS") != "",
		RBACPolicyID:       os.Getenv("REPLICATED_INVITE_RBAC_POLICY_ID"),
		APIToken:           os.Getenv("REPLICATED_API_TOKEN"),
		APIOrigin:          os.Getenv("REPLICATED_API_ORIGIN"),
		GraphQLOrigin:      os.Getenv("REPLICATED_GRAPHQL_ORIGIN"),
		KURLSHOrigin:       os.Getenv("REPLICATED_KURLSH_ORIGIN"),
	}

	if params.NamePrefix == "" {
		return nil, missingParam("REPLICATED_NAME_PREFIX")
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

	if params.InstanceJSONOutput == "" {
		params.InstanceJSONOutput = "./terraform/provisioner_pairs.json"
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

	if params.InviteUsers {
		if params.RBACPolicyID == "" {
			return nil, errors.Errorf("REPLICATED_INVITE_RBAC_POLICY_ID must be set if REPLICATED_INVITE_USERS is set")
		}
		req, err := http.NewRequest("GET", "https://id.replicated.com/v1/policies", nil)
		if err != nil {
			return nil, errors.Wrap(err, "build policy list request")
		}
		req.Header.Set("Authorization", params.APIToken)
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, errors.Wrap(err, "send request")
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, fmt.Errorf("GET /policies %d: %s", resp.StatusCode, body)
		}
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "read body")
		}
		body := []map[string]interface{}{}
		if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&body); err != nil {
			return nil, errors.Wrap(err, "decode body")
		}
		foundPolicy := false
		for _, policy := range body {
			if policyId, ok := policy["id"]; ok {
				if policyId == params.RBACPolicyID {
					foundPolicy = true
					break
				}
			}
		}

		if !foundPolicy {
			return nil, errors.Errorf("Could not find policy %q, found %d total", params.RBACPolicyID, len(body))
		}
	}

	return params, nil
}

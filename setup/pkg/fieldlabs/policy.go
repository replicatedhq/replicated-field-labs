package fieldlabs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/replicatedhq/replicated/pkg/types"
)

type Policy struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Definition  string `json:"definition"`
}

// Copied from vendor-api/policy
// PolicyDefinition implements the JSON schema a user can write to define a policy
type PolicyDefinition struct {
	V1 PolicyDefinitionV1 `json:"v1"`
}

// Copied from vendor-api/policy
// PolicyDefinitionV1 implements the V1 JSON schema for a policy definition
type PolicyDefinitionV1 struct {
	Name      string            `json:"name"`
	Resources PolicyResourcesV1 `json:"resources"`
}

// Copied from vendor-api/policy
// PolicyResourcesV1 implements the resources list in a V1 JSON policy definition
type PolicyResourcesV1 struct {
	Allowed []string `json:"allowed"`
	Denied  []string `json:"denied"`
}

type PolicyListResponse struct {
	Policies []PolicyListItem `json:"policies"`
}

type PolicyListItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PolicyUpdate struct {
	Id string `json:"policy_id"`
}

func (e *EnvironmentManager) getPolicies() (map[string]string, error) {
	requestUrl := fmt.Sprintf("%s/vendor/v1/policies", e.Params.IDOrigin)
	req, err := http.NewRequest(
		"GET",
		requestUrl,
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
		return nil, fmt.Errorf("GET %s %d: %s", requestUrl, resp.StatusCode, body)
	}
	var policies PolicyListResponse
	err = json.Unmarshal([]byte(body), &policies)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("list policies unmarshal %s", body))
	}

	policiesMap := make(map[string]string)
	for i := 0; i < len(policies.Policies); i += 1 {
		policiesMap[policies.Policies[i].Name] = policies.Policies[i].Id
	}
	return policiesMap, nil
}

func (e *EnvironmentManager) createRBAC(app types.App, policies map[string]string) error {
	if _, policyExists := policies[e.Params.ParticipantId]; policyExists {
		// Policy already exists, not recreating
		return nil
	}
	//read + write policy
	policyDefinition := &PolicyDefinition{
		V1: PolicyDefinitionV1{
			Name: "Policy Name",
			Resources: PolicyResourcesV1{
				Allowed: []string{fmt.Sprintf("kots/app/%s/**", app.ID), "kots/license/**", "user/token/**", "kots/externalregistry/list"},
				Denied:  []string{},
			},
		},
	}
	policyDefinitionBytes, err := json.Marshal(policyDefinition)
	if err != nil {
		return errors.Wrap(err, "marshal definition body")
	}
	rbacBody := &Policy{
		Name:        e.Params.ParticipantId,
		Description: e.Params.ParticipantId,
		Definition:  string(policyDefinitionBytes),
	}

	rbacBodyBytes, err := json.Marshal(rbacBody)
	if err != nil {
		return errors.Wrap(err, "marshal rbac body")
	}
	requestUrl := fmt.Sprintf("%s/vendor/v1/policy", e.Params.IDOrigin)
	req, err := http.NewRequest(
		"POST",
		requestUrl,
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
	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("POST %s %d: %s", requestUrl, resp.StatusCode, body)
	}
	return nil

}

// Get a single policy by id
func (e *EnvironmentManager) GetPolicyId(id string) (*Policy, error) {
	requestUrl := fmt.Sprintf("%s/vendor/v1/policy/%s", e.Params.IDOrigin, id)
  e.Log.Debug(fmt.Sprintf("GET %s", requestUrl))
	req, err := http.NewRequest(
		"GET",
		requestUrl,
		nil,
	)

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("GET %s", requestUrl))
	}
	req.Header.Set("Authorization", e.Params.SessionToken)
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("send get request: %s", requestUrl))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GET %s %d %s", requestUrl, resp.StatusCode, body)
  }
  var policy Policy
  err = json.Unmarshal([]byte(body), &policy)
  if err != nil {
    return nil, errors.Wrap(err, fmt.Sprintf("unmarshal %s", body))
  }
  return &policy, nil
} 

// Delete policies create through multi-player mode
func (e *EnvironmentManager) DeletePolicyId(id string) error {
	requestUrl := fmt.Sprintf("%s/vendor/v1/policy/%s", e.Params.IDOrigin, id)
  e.Log.Debug(fmt.Sprintf("DELETE %s", requestUrl))
	req, err := http.NewRequest(
		"DELETE",
		requestUrl,
		nil,
	)

  policy, _ := e.GetPolicyId(id)
  policyJson, _ := json.Marshal(policy)
  e.Log.Debug(fmt.Sprintf("policy %s contains %s", policy.Name, policyJson))
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
	if resp.StatusCode != 204 && resp.StatusCode != 200 {
		return fmt.Errorf("DELETE %s %d: %s", requestUrl, resp.StatusCode, body)
	}
	return nil
}

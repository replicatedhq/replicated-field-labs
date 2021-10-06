package fieldlabs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/replicatedhq/replicated/pkg/types"
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

func (e *EnvironmentManager) Destroy(envs []Environment) error {
	// Check if using single player mode and skip
	if len(e.Params.EnvironmentsJSON) == 0 {
		members, err := e.GetMembers()
		if err != nil {
			return err
		}
		for _, env := range envs {
			for _, member := range members {
				if env.Email == member.Email && member.Is_Pending_Invite {
					continue
				}
				if env.Email == member.Email && !member.Is_Pending_Invite {
					err := e.DeleteMember(member.Id)
					if err != nil {
						return err
					}
				}
			}

			policies, err := e.getPolicies()
			if err != nil {
				return err
			}
			appSlug := fmt.Sprintf("%s-%s", e.Params.NamePrefix, env.Slug)
			for policyName, Id := range policies {
				if policyName == appSlug {
					err := e.DeletePolicyId(Id)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	var appsToDelete []types.App
	apps, err := e.Client.ListApps()
	if err != nil {
		return errors.Wrapf(err, "list apps")
	}

	for _, env := range envs {
		testString := fmt.Sprintf("%s-%s", e.Params.NamePrefix, env.Slug)
		// find all apps matching the prefixed slug for this env
		for _, app := range apps {
			if strings.Contains(app.App.Slug, testString) {
				appsToDelete = append(appsToDelete, *app.App)
			}
		}
	}

	err = e.PrintApps(appsToDelete)
	if err != nil {
		return errors.Wrap(err, "print apps to delete")
	}
	// confirm delete
	answer, err := PromptConfirmDelete()
	if err != nil {
		return errors.Wrap(err, "confirm delete")
	}

	if answer != "yes" {
		return errors.New("prompt declined")
	}

	for _, app := range appsToDelete {
		e.Log.ActionWithSpinner(fmt.Sprintf("Deleting App %s", app.Slug))
		err := e.Client.DeleteKOTSApp(app.ID)
		if err != nil {
			e.Log.FinishSpinnerWithError()
			return errors.Wrapf(err, "delete app %q %q", app.Slug, app.ID)
		}
		e.Log.FinishSpinner()
	}

	return nil
}

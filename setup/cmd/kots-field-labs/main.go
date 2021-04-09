package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/replicatedhq/kots-field-labs/setup/pkg/fieldlabs"
	"github.com/replicatedhq/replicated/cli/print"
	"github.com/replicatedhq/replicated/pkg/kotsclient"
	"github.com/replicatedhq/replicated/pkg/platformclient"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {

	params, err := GetParams()
	if err != nil {
		return errors.Wrap(err, "get params")
	}

	environments, labs, err := loadConfig(params)
	if err != nil {
		return errors.Wrap(err, "load config")
	}

	platformClient := *platformclient.NewHTTPClient(params.APIOrigin, params.APIToken)
	envManager := &fieldlabs.EnvironmentManager{
		Log:     print.NewLogger(os.Stdout),
		Writer:  os.Stdout,
		Params:  params,
		Client:  &kotsclient.VendorV3Client{HTTPClient: platformClient},
		GClient: kotsclient.NewGraphQLClient(params.GraphQLOrigin, params.APIToken, params.KURLSHOrigin),
	}

	if err := envManager.Validate(environments, labs); err != nil {
		return errors.Wrap(err, "validate environments")
	}

	switch params.Action {
	case fieldlabs.ActionCreate:
		return envManager.Ensure(environments, labs)
	case fieldlabs.ActionDestroy:
		return envManager.Destroy(environments)
	}

	return nil
}

func loadConfig(params *fieldlabs.Params) ([]fieldlabs.Environment, []fieldlabs.LabSpec, error) {
	environments := []fieldlabs.Environment{}
	envJSON, err := ioutil.ReadFile(params.EnvironmentsJSON)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "read environments json from %q", params.EnvironmentsJSON)
	}
	err = json.Unmarshal(envJSON, &environments)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "read labs json from %q", params.EnvironmentsJSON)
	}

	labs := []fieldlabs.LabSpec{}
	labJSON, err := ioutil.ReadFile(params.LabsJSON)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "read labs json from %q", params.LabsJSON)
	}
	err = json.Unmarshal(labJSON, &labs)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal labs json from %q", params.LabsJSON)
	}

	return environments, labs, nil
}

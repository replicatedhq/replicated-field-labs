package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/replicatedhq/kots-field-labs/setup/pkg/fieldlabs"
	"os"
)

var actions = map[string]fieldlabs.Action{
	"create":   fieldlabs.ActionCreate,
	"destroy": fieldlabs.ActionDestroy,
}

func missingParam(s string) error {
	return errors.New(fmt.Sprintf("Missing or invalid parameters: %s", s))
}

func GetParams() (*fieldlabs.Params, error) {
	params := &fieldlabs.Params{
		NamePrefix:       os.Getenv("REPLICATED_NAME_PREFIX"),
		EnvironmentsJSON: os.Getenv("REPLICATED_ENVIRONMENTS_JSON"),
		APIToken:         os.Getenv("REPLICATED_API_TOKEN"),
		APIOrigin:        os.Getenv("REPLICATED_API_ORIGIN"),
		GraphQLOrigin:    os.Getenv("REPLICATED_GRAPHQL_ORIGIN"),
		KURLSHOrigin:     os.Getenv("REPLICATED_KURLSH_ORIGIN"),
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

	actionString := os.Getenv("REPLICATED_ACTION")
	if actionString == "" {
		actionString = "create"
	}

	action, ok := actions[actionString]
	if !ok {
		return nil, errors.Errorf("unkown action %s", actionString)
	}
	params.Action = action

	return params, nil
}

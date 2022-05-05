package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/replicatedhq/kots-field-labs/setup/pkg/fieldlabs"
	"github.com/replicatedhq/replicated/cli/print"
	"github.com/replicatedhq/replicated/pkg/kotsclient"
	"github.com/replicatedhq/replicated/pkg/platformclient"
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

	labs, err := loadConfig(params)
	if err != nil {
		return errors.Wrap(err, "load config")
	}

	platformClient := *platformclient.NewHTTPClient(params.APIOrigin, params.APIToken)
	envManager := &fieldlabs.EnvironmentManager{
		Log:    print.NewLogger(os.Stdout),
		Writer: os.Stdout,
		Params: params,
		Client: &kotsclient.VendorV3Client{HTTPClient: platformClient},
	}

	if err := envManager.Validate(labs); err != nil {
		return errors.Wrap(err, "validate labs")
	}

	switch params.Action {
	case fieldlabs.ActionCreate:
		return envManager.Ensure(labs)
	case fieldlabs.ActionDestroy:
		return envManager.Destroy()
	}

	return nil
}

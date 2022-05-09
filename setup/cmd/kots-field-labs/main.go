package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	"github.com/replicatedhq/kots-field-labs/setup/pkg/fieldlabs"
	"github.com/replicatedhq/replicated/cli/print"
	"github.com/replicatedhq/replicated/pkg/kotsclient"
	"github.com/replicatedhq/replicated/pkg/platformclient"
)

func main() {
	runInLambda := os.Getenv("RUN_IN_LAMBDA")
	if runInLambda != "" {
		lambda.Start(HandleRequest)
	} else {
		params, err := GetParams()
		if err != nil {
			log.Fatal(err, "get params")
		}
		if err := Run(params); err != nil {
			log.Fatal(err)
		}
	}
}

func HandleRequest(event fieldlabs.LambdaEvent) error {
	params := &fieldlabs.Params{
		ParticipantEmail: event.ParticipantEmail,
		LabsJSON:         "./labs/labs_all.json",
		LabSlug:          event.LabSlug,
		InviterEmail:     event.InviterEmail,
		InviterPassword:  event.InviterPassword,
		APIToken:         event.APIToken,
		APIOrigin:        "https://api.replicated.com/vendor",
		GraphQLOrigin:    "https://g.replicated.com/graphql",
		KURLSHOrigin:     "https://kurl.sh",
		IDOrigin:         "https://id.replicated.com",
	}

	action, ok := actions[event.Action]
	if !ok {
		return errors.Errorf("unkown action %s", event.Action)
	}
	params.Action = action

	err := getSessionTokenAndCheckInviteParams(params)
	if err != nil {
		return errors.Wrap(err, "validate invite user params")
	}

	return Run(params)
}

func Run(params *fieldlabs.Params) error {
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

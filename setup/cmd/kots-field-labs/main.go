package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	"github.com/replicatedhq/kots-field-labs/setup/pkg/fieldlabs"
	"github.com/replicatedhq/replicated/pkg/kotsclient"
	"github.com/replicatedhq/replicated/pkg/logger"
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
		ParticipantId:   event.ParticipantId,
		Branch:          event.Branch,
		TrackSlug:       event.TrackSlug,
		InviterEmail:    event.InviterEmail,
		InviterPassword: event.InviterPassword,
		APIToken:        event.APIToken,
		APIOrigin:       "https://api.replicated.com/vendor",
		GraphQLOrigin:   "https://g.replicated.com/graphql",
		KURLSHOrigin:    "https://kurl.sh",
		IDOrigin:        "https://api.replicated.com/vendor",
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
	// clone git repo / subfolder
	tempDir, err := gitSparseCheckout(params.TrackSlug, params.Branch)
	defer os.RemoveAll(tempDir)
	if err != nil {
		return errors.Wrap(err, "git sparse checkout")
	}
	vendorLoc := fmt.Sprintf("%s/kots-field-labs/instruqt/%s/vendor", tempDir, params.TrackSlug)
	track, err := loadConfig(vendorLoc)
	if err != nil {
		return errors.Wrap(err, "load config")
	}

	platformClient := *platformclient.NewHTTPClient(params.APIOrigin, params.APIToken)
	envManager := &fieldlabs.EnvironmentManager{
		Log:       logger.NewLogger(os.Stdout),
		Writer:    os.Stdout,
		Params:    params,
		Client:    &kotsclient.VendorV3Client{HTTPClient: platformClient},
		VendorLoc: vendorLoc,
	}

	switch params.Action {
	case fieldlabs.ActionCreate:
		return envManager.Ensure(track)
	case fieldlabs.ActionDestroy:
		return envManager.Destroy(track)
	}

	return nil
}

func gitSparseCheckout(trackSlug string, branch string) (string, error) {
	// the caller of this function is repsonsible for deleting this folder
	tempDir, err := os.MkdirTemp("", "instruqt")
	if err != nil {
		return "", errors.Wrap(err, "failed to create temp dir")
	}

	command := exec.Command("git", "clone", "--depth", "1", "--filter=blob:none", fmt.Sprintf("--branch=%s", branch), "--sparse", "https://github.com/replicatedhq/kots-field-labs.git")
	command.Dir = tempDir
	out, err := command.CombinedOutput()
	if err != nil {
		return tempDir, errors.Wrap(err, fmt.Sprintf("git clone in %s: %s", tempDir, command))
	}
	log.Println("Git clone", string(out))

	command = exec.Command("git", "sparse-checkout", "set", fmt.Sprintf("instruqt/%s/vendor", trackSlug))
	command.Dir = fmt.Sprintf("%s/%s", tempDir, "kots-field-labs")

	out, err = command.CombinedOutput()
	if err != nil {
		return tempDir, errors.Wrap(err, "Git sparse checkout")
	}
	log.Println("Git sparse checkout", string(out))

	return tempDir, nil
}

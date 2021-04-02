package main

import (
	"github.com/pkg/errors"
	"github.com/replicatedhq/kots-field-labs/setup/pkg/fieldlabs"
	"github.com/replicatedhq/replicated/cli/print"
	"github.com/replicatedhq/replicated/pkg/kotsclient"
	"github.com/replicatedhq/replicated/pkg/platformclient"
	"log"
	"os"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

var testEnvs = []fieldlabs.Environment{
	{Name: "Dex", Slug: "dex"},
}

var testLabs = []fieldlabs.LabSpec{
	{
		Name:                 "Lab 1.1: Using Support Analyzers",
		Slug:                 "lab1-e1-ui-analyzers",
		YAMLDir:              "lab1-kots/lab1-e1-ui-analyzers/manifests",
		K8sInstallerYAMLPath: "lab1-kots/lab1-e1-ui-analyzers/kurl-installer.yaml",
		ConfigValues: `---
apiVersion: kots.io/v1beta1
kind: ConfigValues
metadata:
  name: config
spec: {}
`,
	},
}

func Run() error {
	params, err := GetParams()
	if err != nil {
		return errors.Wrap(err, "get params")
	}

	environments := testEnvs
	if params.EnvironmentsJSON != "" {
		// load from file
		//
	}
	labs := testLabs
	if params.LabsJSON != "" {
		// load from file
		//
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

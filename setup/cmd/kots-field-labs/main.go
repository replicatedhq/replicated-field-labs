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

var testEnvs = []fieldlabs.Environment{
	{
		Name:            "Dex",
		Slug:            "dextest",
		KotsadmPassword: "password",
		PubKey:          "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDFK52oX9POpHodnsx0XT4ltw37VYUfulr4z62ZDLOFXl0wztjuo+19DHnVuD70tY8fB1UcyCBkKRy09vQZwOSmV5U4kpVIC9fH9toAZte4Rb7a8wWXyNujBrtKhSMpdNxiNKouf6OjZvRWmoIOXfiEo7oekaERt4dilIkefSK4AT3ccWMWs/pt0GbhyNbCorWW7HHKfeJ+gMkOMV70uQO76Lwhu/7e/Ll72aALpq9RPt7xaOllBTq4iIz7x/E7k9/w2h9D5/xHiKIOBhJJw8Vd9yS0Tj+u1jg1a68CF2YQhdakTpqDhISsKKVtkb31MPqrdZpqNKu37J29Q6fxNN3KpaZkt19BMG+L28uOXon9+782AIUJqTGnqKcJhziyCOZpKaBiu2S1cbDSRJpyaqHZi3vMy5eleblWgQn/tbUQMtWh1UR5KANGhvBVS84hxFWkPuCwWORnewQCpz8jPXMpaOnLK2n7ZZSBmSXOYOozQh/MfNamtRajiUhBfHxuh5jD3FcXsAVy2yYmCZVAXJB/XzJMeNKGz6mmWH+9xBufa8oFYedQAUiyyVgW6QODNO5uu3YVQtySjuwsenxp2guBfiteSUtMJeclQjSbglCLtvrDXkF6AKiYkx/+5Bz2RpoitgXvL92EAEPiAOLxOVKRtbkMjG4xLM8gYQXkncpy+Q== dex",
	},
}

var testLabs = []fieldlabs.LabSpec{
	{
		Name:                 "Lab 1.0: Hello World",
		Slug:                 "lab1-e0-hello-world",
		Channel:              "Unstable",
		ChannelSlug:          "unstable",
		Customer:             "Dev Customer",
		YAMLDir:              "lab1-kots/lab1-e0-hello-world/manifests",
		K8sInstallerYAMLPath: "lab1-kots/lab1-e0-hello-world/kurl-installer.yaml",
		SkipInstallKots:      true,
		PublicIP:             true,
	},
}

func writeLabs() error {
	labs := testLabs

	jsonLabs, err := json.MarshalIndent(labs, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal labs")
	}

	_ = ioutil.WriteFile("labs.json", jsonLabs, 0644)
	return errors.New("thats all folks")
}

func Run() error {
	if err := writeLabs(); err != nil {
		return err
	}

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
	environments := testEnvs
	if params.EnvironmentsJSON != "" {
		environments = []fieldlabs.Environment{}
		contents, err := ioutil.ReadFile(params.EnvironmentsJSON)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "read environments json from %q", params.EnvironmentsJSON)
		}
		err = json.Unmarshal(contents, &environments)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "read labs json from %q", params.EnvironmentsJSON)
		}
	}

	labs := testLabs
	if params.LabsJSON != "" {
		labs = []fieldlabs.LabSpec{}
		contents, err := ioutil.ReadFile(params.LabsJSON)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "read labs json from %q", params.LabsJSON)
		}
		err = json.Unmarshal(contents, &environments)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "unmarshal labs json from %q", params.LabsJSON)
		}
	}

	return environments, labs, nil
}

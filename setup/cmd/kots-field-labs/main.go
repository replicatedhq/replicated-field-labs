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
	{
		Name:            "Dex",
		Slug:            "dex",
		KotsadmPassword: "password",
		PubKey:          "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDFK52oX9POpHodnsx0XT4ltw37VYUfulr4z62ZDLOFXl0wztjuo+19DHnVuD70tY8fB1UcyCBkKRy09vQZwOSmV5U4kpVIC9fH9toAZte4Rb7a8wWXyNujBrtKhSMpdNxiNKouf6OjZvRWmoIOXfiEo7oekaERt4dilIkefSK4AT3ccWMWs/pt0GbhyNbCorWW7HHKfeJ+gMkOMV70uQO76Lwhu/7e/Ll72aALpq9RPt7xaOllBTq4iIz7x/E7k9/w2h9D5/xHiKIOBhJJw8Vd9yS0Tj+u1jg1a68CF2YQhdakTpqDhISsKKVtkb31MPqrdZpqNKu37J29Q6fxNN3KpaZkt19BMG+L28uOXon9+782AIUJqTGnqKcJhziyCOZpKaBiu2S1cbDSRJpyaqHZi3vMy5eleblWgQn/tbUQMtWh1UR5KANGhvBVS84hxFWkPuCwWORnewQCpz8jPXMpaOnLK2n7ZZSBmSXOYOozQh/MfNamtRajiUhBfHxuh5jD3FcXsAVy2yYmCZVAXJB/XzJMeNKGz6mmWH+9xBufa8oFYedQAUiyyVgW6QODNO5uu3YVQtySjuwsenxp2guBfiteSUtMJeclQjSbglCLtvrDXkF6AKiYkx/+5Bz2RpoitgXvL92EAEPiAOLxOVKRtbkMjG4xLM8gYQXkncpy+Q== dex",
	},
}

var InternalEnvs =[]fieldlabs.Environment{
	{
		Name:            "Dex",
		Slug:            "dex",
		KotsadmPassword: "password",
		PubKey:          "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDFK52oX9POpHodnsx0XT4ltw37VYUfulr4z62ZDLOFXl0wztjuo+19DHnVuD70tY8fB1UcyCBkKRy09vQZwOSmV5U4kpVIC9fH9toAZte4Rb7a8wWXyNujBrtKhSMpdNxiNKouf6OjZvRWmoIOXfiEo7oekaERt4dilIkefSK4AT3ccWMWs/pt0GbhyNbCorWW7HHKfeJ+gMkOMV70uQO76Lwhu/7e/Ll72aALpq9RPt7xaOllBTq4iIz7x/E7k9/w2h9D5/xHiKIOBhJJw8Vd9yS0Tj+u1jg1a68CF2YQhdakTpqDhISsKKVtkb31MPqrdZpqNKu37J29Q6fxNN3KpaZkt19BMG+L28uOXon9+782AIUJqTGnqKcJhziyCOZpKaBiu2S1cbDSRJpyaqHZi3vMy5eleblWgQn/tbUQMtWh1UR5KANGhvBVS84hxFWkPuCwWORnewQCpz8jPXMpaOnLK2n7ZZSBmSXOYOozQh/MfNamtRajiUhBfHxuh5jD3FcXsAVy2yYmCZVAXJB/XzJMeNKGz6mmWH+9xBufa8oFYedQAUiyyVgW6QODNO5uu3YVQtySjuwsenxp2guBfiteSUtMJeclQjSbglCLtvrDXkF6AKiYkx/+5Bz2RpoitgXvL92EAEPiAOLxOVKRtbkMjG4xLM8gYQXkncpy+Q== dex",
	},
	{
		Name:            "Todd",
		Slug:            "todd",
		KotsadmPassword: "password",
		PubKey:          "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIFF+vXc0McDwC9y5qTBlTH3Z8dnwbV28w72M5FhKA2xG todd@replicated.com",
	},
	{
		Name:            "Dan",
		Slug:            "flux-capacitor",
		KotsadmPassword: "password",
		PubKey:          "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAINzndPU93H5NerufeEKuh2QuRJXdiRK6cr0ulLRLZNzJ dans@replicated.com",
	},
	{
		Name:            "Jalaja",
		Slug:            "jala",
		KotsadmPassword: "password",
		PubKey:          "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICfp3rCP6O6B+LVm7KmKB0Hp1IU7Q8EuqfbnW2pCJ9AA jalaja@replicated.com",
	},
	{
		Name:            "Ethan",
		Slug:            "ethan",
		KotsadmPassword: "password",
		PubKey:          "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCZvDX1RQtu/Kxxh+oHXMBiTp8kyUOnmd1SE3gr50VlJRqR0R43EZcAVQ+BOFtXNeOs0h5FeXK0CBXzQvU9WeF97SGJStY+KVpxZkRzILBSCppPHrVtvLm+X6QQbzsMcwdkmHJz1Aucrn6izMP1+b4DCTMnL/yE82YQ40GxL0fKiTXGIn3K5JxM6wB+HQGi4V4QJchDW2lpOU5rUuOg7X6/2GtR+Jg1yDc8gqk1QrI4N2ctndUKehd4JVFscOAguHBlYoaSxsn71rTKw0Z4XLzcvldoy0N6yG4tyfbrnhO0aS/WiMuo0sAUY63/NJPx18Zvw6WBmf4rUV0GXO54HC8SGrNu7OAJAM9boELddifbrn4FWazQKLG0jwYETRbeGiSbuI40ae1Iq1VEQc1tnzMghZvmSkKIS+eNAgrat/KCVbJ6oP6/9XqtOeACumpwPjpQ/sGWEg2KyPjT5OjkveqIIJOq8BvW+ZbAA3HmWjzG1FPCiHkd+byDnjq0IC86A82hyUJpRaKRItO7s1dLnW3VIrcbwCWknfz5XgIqmt4Iv7CF0BI444MDugOxFFItvaje6cNO3jTEPwvEs6Np63blBPSMeAn2dEoDj3bf2xZSl7wuvJGd9WLNrT57SEccTGuvoC5Zye4/QoRAAtLJh1iPUNPl6MHn+HuWI+vmgv19IQ== ethan@replicated.com",
	},
	{
		Name:            "Dmitriy",
		Slug:            "dmitriy",
		KotsadmPassword: "password",
		PubKey:          "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFuOjxawceuaDpy9LYcOzNtKEHJcS5q0u9+xrLKynpfUa71FGeSDeeebyynHZxy1eW1walbPxj4DxPpMdZco6MdsN2iUE5PpQNMgH51P4s3/KZrQyZVWxOXQkUirvQeEgjuDD1pUVGM8HB/cW8iGLUwXPq92cAzv+3Kd7VjSeywHwzf8b3UwdeliCWpyhKd0pR0MRYFl9OLvR6vNCenJtpH7fu0ncVJHC7p+XO6RmtHHABZ9l9no/Nr81i5JiyjgNdD7r3vrcfSF+urxDpgX33ey1tylTBBOdysoDUbvL4LhJetK8F7bNWAgeAABpFFr2awq2cCiDaX413pCanFuKb divolgin@Ds-MacBook-Pro.local",
	},
}

var Stragglers = []fieldlabs.Environment{
	{
		Name:            "Max",
		Slug:            "max",
		KotsadmPassword: "password",
		PubKey:          "",
	},
	{
		Name:            "Barry",
		Slug:            "barry",
		KotsadmPassword: "password",
		PubKey:          "",
	},
	{
		Name:            "Josh",
		Slug:            "josh",
		KotsadmPassword: "password",
		PubKey:          "",
	},
	{
		Name:            "Javier",
		Slug:            "javier",
		KotsadmPassword: "password",
		PubKey:          "",
	},
	{
		Name:            "Fernando",
		Slug:            "fernando",
		KotsadmPassword: "password",
		PubKey:          "",
	},
	{
		Name:            "AJ",
		Slug:            "aj",
		KotsadmPassword: "password",
		PubKey:          "",
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
		JumpBox:              false,
	},
	{
		Name:                 "Lab 1.1: Using Support Analyzers",
		Slug:                 "lab1-e1-ui-analyzers",
		Channel:              "lab1-e1-ui-analyzers",
		ChannelSlug:          "lab1-e1-ui-analyzers",
		Customer:             "lab1-e1-ui-analyzers",
		YAMLDir:              "lab1-kots/lab1-e1-ui-analyzers/manifests",
		K8sInstallerYAMLPath: "lab1-kots/lab1-e1-ui-analyzers/kurl-installer.yaml",
		PublicIP:             true,
		JumpBox:              false,
	},
	{
		Name:                 "Lab 1.2: Adding Analyzers",
		Slug:                 "lab1-e2-adding-analyzers",
		Channel:              "lab1-e2-adding-analyzers",
		ChannelSlug:          "lab1-e2-adding-analyzers",
		Customer:             "lab1-e2-adding-analyzers",
		YAMLDir:              "lab1-kots/lab1-e2-adding-analyzers/manifests",
		K8sInstallerYAMLPath: "lab1-kots/lab1-e2-adding-analyzers/kurl-installer.yaml",
		PublicIP:             true,
		JumpBox:              false,
		PreInstallSH: `
sudo mkdir -p /etc/lab1-e2/
sudo touch /etc/lab1-e2/config.txt
sudo chmod 400 /etc/lab1-e2/config.txt
`,
	},
	//{
	//	Name:                 "Lab 1.5: Airgapped Install",
	//	Slug:                 "lab1-e5-airgap",
	//	Channel:              "lab1-e5-airgap",
	//	ChannelSlug:          "lab1-e5-airgap",
	//	Customer:             "lab1-e5-airgap",
	//	YAMLDir:              "lab1-kots/lab1-e5-airgap/manifests",
	//	K8sInstallerYAMLPath: "lab1-kots/lab1-e5-airgap/kurl-installer.yaml",
	//	SkipInstallKots:      true,
	//	PublicIP:             false,
	//	JumpBox:              true,
	//},
}

func Run() error {
	params, err := GetParams()
	if err != nil {
		return errors.Wrap(err, "get params")
	}

	environments := InternalEnvs
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

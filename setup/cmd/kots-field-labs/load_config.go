package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/replicatedhq/kots-field-labs/setup/pkg/fieldlabs"
)

func loadConfig(params *fieldlabs.Params) ([]fieldlabs.LabSpec, error) {
	labs := []fieldlabs.LabSpec{}
	labJSON, err := ioutil.ReadFile(params.LabsJSON)
	if err != nil {
		return nil, errors.Wrapf(err, "read labs json from %q", params.LabsJSON)
	}
	err = json.Unmarshal(labJSON, &labs)
	if err != nil {
		return nil, errors.Wrapf(err, "unmarshal labs json from %q", params.LabsJSON)
	}

	return labs, nil
}

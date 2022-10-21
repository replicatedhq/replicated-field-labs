package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/replicatedhq/kots-field-labs/setup/pkg/fieldlabs"
)

func loadConfig(vendorLoc string) (*fieldlabs.TrackSpec, error) {
	track := fieldlabs.TrackSpec{}
	vendorJSON, err := ioutil.ReadFile(fmt.Sprintf("%s/vendor.json", vendorLoc))
	if err != nil {
		return nil, errors.Wrapf(err, "read vendor json from %q", vendorLoc)
	}
	err = json.Unmarshal(vendorJSON, &track)
	if err != nil {
		return nil, errors.Wrapf(err, "unmarshal vendor json from %q", vendorLoc)
	}

	if track.Channel == "" {
		track.Channel = "Stable"
	}

	return &track, nil
}

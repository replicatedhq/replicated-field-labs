package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/replicatedhq/kots-field-labs/setup/pkg/fieldlabs"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name    string
		params  fieldlabs.Params
		files   map[string]string
		wantErr string
	}{
		{
			name: "csv",
			files: map[string]string{
				"testing-envs.csv": `Timestamp,name,email,slug,pub_key,password
4/9/2021 23:07:06,Dex,dex@replicated.com,dex,ssh-rsa public-key-bytes dex@replicated.com,password`,
				"testing-labs.json": `[]`,
			},
			params: fieldlabs.Params{
				LabsJSON: "testing-labs.json",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := require.New(t)
			for name, contents := range test.files {
				err := ioutil.WriteFile(name, []byte(contents), 0644)
				req.NoError(err)
			}

			defer func() {
				for name := range test.files {
					_ = os.RemoveAll(name)
				}
			}()
			_, err := loadConfig(&test.params)
			if test.wantErr != "" {
				req.EqualError(err, test.wantErr)
			}

		})
	}
}

func mustJSON(v interface{}) string {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(bytes)

}

Field Lab Setup Tools
========================


### Parameters

There are no CLI flags because folks were in a hurry -- review the supported env vars in [setup/cmd/kots-field-labs/param.go](./setup/cmd/kots-field-labs/param.go)


    # Token for provisioning new apps and users
    export REPLICATED_API_TOKEN=...
    # JSON containing environment (session attendee) specs
    export REPLICATED_ENVIRONMENTS_JSON=./environments_test.json
    # JSON containing lab exercise specs
    export REPLICATED_LABS_JSON=./labs_test.json
    # unique prefix to avoid app slug collisions
    export REPLICATED_NAME_PREFIX=abcd


### Usage

    cd setup
    go install ./cmd/kots-field-labs
    cd ../
    REPLICATED_ACTION=create  kots-field-labs

this will create apps, save the list and share w/ participants

Review ./terraform/provisioner_pairs.json

    cd terraform
    terraform init 
    terraform apply -var user=$GCP_NAME

### Cleanup

    cd terraform
    terraform destroy
    cd ..
    REPLICATED_ACTION=destroy kots-field-labs

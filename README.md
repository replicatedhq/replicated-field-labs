Field Lab Setup Tools
========================


### Parameters

There are no CLI flags on the binary itself, Makefile wraps env vars:

    $ head Makefile
    user := ${USER}
    labs_json := "labs_e0.json"
    env_json := ""
    env_csv := ""
    action := "create"
    prefix := ""
    provisioner_json_out := "terraform/provisioner_pairs.json"
    
You'll also need to set an API token

    # Token for provisioning new apps and users
    export REPLICATED_API_TOKEN=...


### Usage

Prefix should be set to a short unique string to prevent name collisions

	make apps prefix=abcd env_json=... labs_json=...
	make tf

### Cleanup

    make -C terraform destroy
    make apps action=destroy prefix=abcd env_json=... labs_json=...

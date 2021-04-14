user := ${USER}
labs_json := "labs_e0.json"
env_json := ""
env_csv := ""
action := "create"
prefix := ""
provisioner_json_out := "terraform/provisioner_pairs.json"

.PHONY: install
install:
	@$(MAKE) -C setup install

.PHONY: apps
apps: install
	REPLICATED_LABS_JSON=$(labs_json) \
	REPLICATED_ENVIRONMENTS_JSON=$(env_json) \
	REPLICATED_ENVIRONMENTS_CSV=$(env_csv) \
	REPLICATED_ACTION=$(action) \
	REPLICATED_NAME_PREFIX=$(prefix) \
	REPLICATED_INSTANCE_JSON_OUT=$(provisioner_json_out) \
	kots-field-labs

.PHONY: instances
instances:
	TF_VAR_user=$(user) \
	TF_VAR_provisioner_pairs_json=$(provisioner_json_out) \
	$(MAKE) -C terraform apply
outputs:
	$(MAKE) -C terraform output

.PHONY: both
both: apps instances

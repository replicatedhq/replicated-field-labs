user := ${USER}
labs_json := "labs_e0.json"
env_json := ""
env_csv := ""
action := "create"
prefix := ""
provisioner_json_out := ""

.PHONY: install
install:
	@$(MAKE) -C setup install

.PHONY: launch-one
launch-one: install
	REPLICATED_LABS_JSON=$(labs_json) \
	REPLICATED_ENVIRONMENTS_JSON=$(env_json) \
	REPLICATED_ENVIRONMENTS_CSV=$(env_csv) \
	REPLICATED_ACTION=$(action) \
	REPLICATED_NAME_PREFIX=$(prefix) \
	REPLICATED_INSTANCE_JSON_OUTPUT=$(provisioner_json_out) \
	kots-field-labs
	terraform apply ./terraform -var user=$(user) -var provisioner_pairs_json=$(provisioner_json_out)
	
	
	


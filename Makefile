user := $(or ${REPLICATED_GCP_USER},${USER})
labs_json := "labs_e0.json"
env_json := ""
env_csv := ""
action := "create"
prefix := ""
invite_users := ""
inviter_email := "dexter+training415@replicated.com"
inviter_password := ""
provisioner_json_out := "terraform/provisioner_pairs.json"
REPLICATED_GCP_PROJECT ?= "kots-field-labs"
REPLICATED_GCP_ZONE ?= "us-central1-b"
OWNER ?= ${USER}

.PHONY: install
install:
	@$(MAKE) -C setup install

.PHONY: test
test:
	@$(MAKE) -C setup test

.PHONY: apps
apps: install
	REPLICATED_LABS_JSON=$(labs_json) \
	REPLICATED_ENVIRONMENTS_JSON=$(env_json) \
	REPLICATED_ENVIRONMENTS_CSV=$(env_csv) \
	REPLICATED_ACTION=$(action) \
	REPLICATED_NAME_PREFIX=$(prefix) \
	REPLICATED_INVITE_USERS=$(invite_users) \
	REPLICATED_INVITER_EMAIL=$(inviter_email) \
	REPLICATED_INVITER_PASSWORD=$(inviter_password) \
	REPLICATED_INSTANCE_JSON_OUT=$(provisioner_json_out) \
	kots-field-labs

.PHONY: instances
instances:
	TF_VAR_user=$(user) \
	TF_VAR_owner=$(OWNER) \
	TF_VAR_gcp_project=$(REPLICATED_GCP_PROJECT) \
	TF_VAR_gcp_zone=$(REPLICATED_GCP_ZONE) \
	TF_VAR_provisioner_pairs_json=$(provisioner_json_out) \
	$(MAKE) -C terraform apply

.PHONY: tfdestroy
tfdestroy:
	TF_VAR_user=$(user) \
	TF_VAR_owner=$(OWNER) \
	TF_VAR_gcp_project=$(REPLICATED_GCP_PROJECT) \
	TF_VAR_gcp_zone=$(REPLICATED_GCP_ZONE) \
	TF_VAR_provisioner_pairs_json=$(provisioner_json_out) \
	$(MAKE) -C terraform destroy

outputs:
	$(MAKE) -C terraform output

pretty_ips:
	tail -n 1000 terraform/etchosts/*

.PHONY: both
both: apps instances

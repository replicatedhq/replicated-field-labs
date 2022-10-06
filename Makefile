participant_id := ${USER}
branch := "main"
track_slug := "sample"
inviter_email := "dexter+training415@replicated.com"
inviter_password := ""

.PHONY: install
install:
	@$(MAKE) -C setup install

.PHONY: test
test:
	@$(MAKE) -C setup test

.PHONY: create
create: install
	REPLICATED_ACTION=create \
    PARTICIPANT_ID=$(participant_id) \
	REPLICATED_BRANCH=$(branch) \
	REPLICATED_TRACK_SLUG=$(track_slug) \
	REPLICATED_INVITER_EMAIL=$(inviter_email) \
	REPLICATED_INVITER_PASSWORD=$(inviter_password) \
	kots-field-labs

.PHONY: destroy
destroy: install
	REPLICATED_ACTION=destroy \
	PARTICIPANT_ID=$(participant_id) \
	REPLICATED_BRANCH=$(branch) \
	REPLICATED_TRACK_SLUG=$(track_slug) \
	REPLICATED_INVITER_EMAIL=$(inviter_email) \
	REPLICATED_INVITER_PASSWORD=$(inviter_password) \
	kots-field-labs



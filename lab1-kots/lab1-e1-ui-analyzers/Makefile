CHANNEL=lab1-e1-ui-analyzers

release: 
	@replicated release create \
		--app ${REPLICATED_APP} \
		--token ${REPLICATED_API_TOKEN} \
		--auto \
		--promote $(CHANNEL)

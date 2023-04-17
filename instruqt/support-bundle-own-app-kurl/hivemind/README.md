# Hivemind

A really bad command-and-control service for instruqt sandboxes

## usage:

`hivemind super_secret_passphrase`

## TODO:

- maybe use pubsub instead of ntfy.sh

## local testing

due to cors issues testing the site with file:// scheme doesn't work. running `caddy.sh` with podman installed spins up a local webserver for testing.
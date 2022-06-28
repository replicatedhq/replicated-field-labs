Single Player Mode
======================


This guide will walk you through running all the labs on your own.

You should probably skim through [the architecture outline](./01-architecture.md) first.

## 1. Prerequisites

* Create a vendor.replicated.com account
* Create and copy an API token, set it in your shell with `export REPLICATED_API_TOKEN=...`
* Choose a unique name for your session, e.g. `dh-test`
* Install terraform
* Install a Go toolchain (this has been tested w/ 1.16), ensure `$GOPATH` is set and `$GOPATH/bin` is added to your `$PATH` 
  ```
  export GOPATH=$HOME/go
  export GOBIN=$GOPATH/bin
  export PATH="$PATH:$GOBIN"
  ```
* Install the `gcloud` CLI and log in with application default credentials: `gcloud auth application-default login`
* Add your `google_compute_engine` ssh key to the `ssh-agent` or configure `~/.ssh/config` correctly. (ssh-add ~./ssh/google_compute_engine)
* You should probably skim through [the architecture outline](./01-architecture.md) first.

## 2. Create an environment JSON

Copy the example file and edit it with your name and slug.
You can ignore the `pub_key` and `password`. Those are currently not used, and a default password is used to ssh into the box. 
You'll already have access to the vendor account, so email will be not be required either.

```
cp environments_test.json env-dh-test.json
vim env-dh-test.json
```

Whatever file you create, you should git-ignore it. 
I usually drop mine in a `.dex` folder which is globally git-ignored.

## 3. Choose your labs

You can use the `labs_all.json` file, or you can copy it to another file and edit it to remove labs you'd like to skip.
If you'd like to set up all the labs, you can skip the below commands and just pass `labs_json=labs_all.json` in step 4.

```
cp labs_all.json labs-dh-test.json
vim labs-dh-test.json
```

## 4. Provision your apps

The Makefile provides a handy wrapper around the golang code in `setup`. 
A simple invocation of the above two files might look like:

```shell
make apps \
  prefix=dh-test \
  labs_json=labs-dh-test.json \
  env_json=env-dh-test.json \
  inviter_password='YOUR_VENDORPORTAL_PASSWORD' \
  inviter_email=you@vendorportal.com 
```

The output of this will be a file at `terraform/provisioner_pairs.json`.
You can review this file to understand which GCP instances will be created.

You should log into your vendor.replicated.com account to briefly review the app, channels, releases, and customers that were created.

## 5. Terraform Init

* By default this procedure will deploy instances to the GCP Project `kots-field-labs` in zone `us-central1-b`. To set an alternate project `export REPLICATED_GCP_PROJECT=...`. To set an alternate zone `export REPLICATED_GCP_ZONE=...`
* By default this procedure will provision a user account on the GCP instances that matches your currently logged in local user. To override this in cases where your GCP username differs from your workstation, set `export REPLICATED_GCP_USER=...`
* The provisioned instances will have the following labels set
    * `expires-on`: Set for 14 days from the moment of creation.
    * `owner`: Defaults to $USER. If you want to override, use `export OWNER=...`.
* The terraform provisioners' connection settings leverage ssh-agent. If terraform errors with ssh timeouts, consider adding your local private key to ssh agent with `ssh-add -K ...`.

If you haven't already, initialize the terraform dir. You can cd in and run `terraform init`, or there's a `make` wrapper for this.

```shell
make -C terraform init
```

## 6. Terraform apply

There's a handy `make` wrapper for this.

```shell
make instances
```

## 7. Host file

Grab the host file from the `terraform/etchosts/` directory and follow the instructions in the comment to add host entries for the lab servers:

```text
$ cat terraform/etchosts/aj
# copy the below and add it to your hosts file with
#
#     echo '
#     <PASTE>
#     ' | sudo tee -a /etc/hosts

34.121.47.43	lab05-airgap-jump	# dppt-aj-lab05-airgap-jump
104.198.254.92	lab06-proxy-jump	# dppt-aj-lab06-proxy-jump
```

## 8. You're ready!

For each lab you provisioned, head to the README for that lab. 
For example, to walk through Lab 0, navigate to the [labs/lab00-hello-world README](../labs/lab00-hello-world).

## 9. Iterating

If you make changes to Lab YAMLs as part of a lab exercise, use the `make release` command for that lab.

If you are developing changes to a lab spec itself (e.g. changing how the instance is provisioned), you can re-run the
same `make apps` and `make instances` commands. `make apps` will merge any new instances into `provisioner_pairs.json`, although
it **will not** overwrite existing instance specs -- you'll have to remove these manually before running `make apps`.

`make apps` is mostly idempotent. It includes some rudimentary get-or-create logic for apps, channels, and customers. 
The exception is releases -- new releases for app YAML and kURL installer will always created and promoted to the lab's channel.

## 10. Cleaning up

```shell
make tfdestroy
make apps \
  action=destroy \
  prefix=dh-test \
  labs_json=labs-dh-test.json \
  env_json=env-dh-test.json
  inviter_password='YOUR_VENDORPORTAL_PASSWORD' \
  inviter_email=you@vendorportal.com 
```

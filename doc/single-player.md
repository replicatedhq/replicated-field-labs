Single Player Mode
======================


This guide will walk you through running all the labs on your own.

You should probably skim through [the architecture outline](./architecture.md) first.

## 1. Prerequisites

* Create a vendor.replicated.com account
* Create and copy an API token, set it in your shell with `export REPLICATED_API_TOKEN=...`
* Choose a unique name for your session, e.g. `dh-test`
* Install terraform
* Install a Go toolchain (this has been tested w/ 1.16)
* Install the `gcloud` CLI and log in with application default credentials: `gcloud auth application-default login`
* You should probably skim through [the architecture outline](./architecture.md) first.

## 2. Create an environment JSON

Copy the example file and edit it with your name and slug.
Since single-player mode requires you to have server access to provision the boxes in the first place,
so you can omit the SSH public key if you'd like. 
You'll already have access to the vendor account, so email will be not be required either.

```
cp environments_test.json env-dh-test.json
vim env-dh-test.json
```

Whatever file you create, you should git-ignore it. 
I usually drop mine in a `.dex` folder which is globally git-ignored.

## 3. Choose your labs

You can use the `labs_all.json` file, or you can copy it to another file and edit it to remove labs you'd like to skip.
If you'd like to set up all the labs, you can skip this next step.

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
  env_json=env-dh-test.json
```

The output of this will be a file at `terraform/provisioner_pairs.json`.
You can review this file to understand which GCP instances will be created.

You should log into your vendor.replicated.com account to briefly review the app, channels, releases, and customers that were created.

## 5. Terraform Init

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

34.121.47.43	lab5-airgap-jump	# dppt-aj-lab5-airgap-jump
104.198.254.92	lab6-proxy-jump	# dppt-aj-lab6-proxy-jump
```

## 8. You're ready!

For each lab you provisioned, head to the README for that lab. 
For example, to walk through Lab 0, navigate to the [labs/lab0-hello-world README](../labs/lab0-hello-world).

## 9. Iterating

If you make changes to Lab YAMLs as part of a lab exercise, use the `make release` command for that lab.

If you are developing changes to a lab spec itself (e.g. changing how the instance is provisioned), you can re-run the
same `make apps` and `make instances` commands. `make apps` will merge any new instances into `provisioner_pairs.json`, although
it **will not** overwrite existing instance specs -- you'll have to remove these manually before running `make apps`.

`make apps` is mostly idempotent. It includes some rudimentary get-or-create logic for apps, channels, and customers. 
The exception is releases -- new releases for app YAML and kURL installer will always created and promoted to the lab's channel.

## 10. Cleaning up

```shell
make -C terraform destroy
make apps \
  action=destroy \
  prefix=dh-test \
  labs_json=labs-dh-test.json \
  env_json=env-dh-test.json
```
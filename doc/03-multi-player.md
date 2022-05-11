Multi Player Mode
======================


This guide will walk you through running all the labs for a team of people.

## 1. Prerequisites


+ You should probably skim through [the architecture outline](./01-architecture.md) first.
+ You should be able to setup the labs in [single player mode](./02-single-player.md).

## 2. Create the files

Create a CSV file with a row for each participant.

```csv
name,password,email,slug
Participant One,password,participant.one+[PREFIX]@somecompany.com,participant
```

In above example, the following fields are being mapped:
+ `name`: full name of the participant
+ `password`: password to be set on kotsadm instance UI
+ `email`: participant email address to send invite to vendor.replicated.com if Params.InviteUsers is set
+ `slug`: slug of the environment, match the participant name, all lowercase no spaces

Create a JSON file with the labs you'd like to run. You can use `labs_all.json` as a starter, and remove the labs you don't want.

```json
[
  {
    "name": "Lab 17: Persistent Datastores",
    "slug": "lab17-persistent-storage",
    "channel": "lab17-persistent-storage",
    "channel_slug": "lab17-persistent-storage",
    "customer": "lab17-persistent-storage",
    "yaml_dir": "labs/lab17-persistent-storage/manifests",
    "k8s_installer_yaml_path": "labs/lab17-persistent-storage/kurl-installer.yaml",
    "skip_install_kots": false,
    "skip_install_app": false,
    "config_values": "",
    "pre_install_sh": "",
    "post_install_sh": "",
    "use_public_ip": true,
    "use_jump_box": false
  }
]
```

## 3. Initialization

When provisioning for an actual training, the following is needed (to invite users):

```shell
make apps invite_users=1 \
  env_csv=[CSV File with ssh public keys] \
  labs_json=labs_all.json \
  prefix=[PREFIX] \
  inviter_password=[Your Vendor web password] \
  inviter_email=[your Vendor web email]
```
> NOTE: 
When creating the field labs the prefix cannot begin with a number or contain `lab`. Additionally the inviter_password cannot contain a '!'.


## 4. VM Instances

See [Terraform apply](./02-single-player.md#6-terraform-apply)
```shell
make instances
```

## 5. Notifying particpants

The below can be used as a starter for notifying participants that their labs have been provisioned.

```
Hi folks,

As promised, lab environments have been provisioned for $USERS -- to get started, you'll need to do a few things

1. Check your email and accept the invite to create an account (subject will be "Invitation to join team on replicated" from "contact@replicated.com")
2. Navigate to the [first lab](https://github.com/replicatedhq/kots-field-labs/tree/main/labs/lab00-hello-world) and start working on the readme
3. You will see prompts to "insert your IP address here" -- those IPs for each participant are found below. The password for each SSH login will be $SOMETHING and the UI password for logging into the browser view will be $SOMETHING_ELSE.


$IP_ADDRESS_INFO

If if you get stuck, feel free to reach out via $CHANGEME

Thanks and good luck! 
```

## 6. Tearing it down

After the training is done, we usually give them a couple days to finish the labs.
1. End of training day 2: Ask who is finished with the labs and remove their instances
    ```shell
    gcloud compute instances delete $(gcloud compute instances list --filter="name:[PREFIX]" | awk '{ print $1 }' | grep [USERNAME]) --zone us-central1-b
    ```
1. Day after the training: Ask again who is finished with the labs and remove their instances.
1. 3 days after the training was done: Let them know all instances will be removed EOD, unless someone lets us know they will need more time. Delete all resources created by running the following two commands:

The second `make apps` command will be identical to your creation command, but with `action=destroy` added.

```shell
make tfdestroy

make apps action=destroy \
  invite_users=1 \
  env_csv=[CSV File with ssh public keys]\
  labs_json=labs_all.json prefix=[PREFIX] \
  inviter_password=[Your Vendor web password] \
  inviter_email=[your Vendor web email]
```
> NOTE: Use the same PREFIX and `labs_all.json` that was used during creation.

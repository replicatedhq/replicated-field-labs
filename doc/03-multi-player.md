Multi Player Mode
======================


This guide will walk you through running all the labs for a team of people.

## 1. Prerequisites


+ You should probably skim through [the architecture outline](./01-architecture.md) first.
+ You should be able to setup the labs in [single player mode](./02-single-player.md).

## 2. Initialization

When provisioning for an actual training, the following is needed (to invite users):

```shell
make apps invite_users=1 env_csv=[CSV File with ssh public keys] labs_json=labs_all.json prefix=[PREFIX] inviter_password=[Your Vendor web password] inviter_email=[your Vendor web email]
```
> NOTE: 
When creating the field labs the prefix cannot begin a number or contain `lab`. Additionally the inviter_password cannot contain a '!'.

A csv example:
```csv
name,pub_key,password,email,slug
Participant One,,password,participant.one+[PREFIX]@somecompany.com,participant
```

In above example, the following fields are being mapped:
+ `name`: full name of the participant
+ `pub_key`: leave empty
+ `password`: password to be set on kotsadm instances
+ `email`: participant email address to send invite to vendor.replicated.com if Params.InviteUsers is set
+ `slug`: slug of the environment




## 3. VM Instances

See [Terraform apply](./02-single-player.md#6-terraform-apply)
```shell
make instances
```

## 3. Notifying particpants

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

## 4. Tearing it down

After the training is done, we usually give them a couple days to finish the labs.
1. End of training day 2: Ask who is finished with the labs and remove their instances
    ```shell
    gcloud compute instances delete $(gcloud compute instances list --filter="name:[PREFIX]" | awk '{ print $1 }' | grep [USERNAME]) --zone us-central1-b
    ```
1. Day after the training: Ask again who is finished with the labs and remove their instances.
1. 3 days after the training was done: Let them know all instances will be removed EOD, unless someone lets us know they will need more time. Delete all resources created by running the following two commands:
```shell
make tfdestroy

make apps action=destroy invite_users=1 env_csv=[CSV File with ssh public keys] labs_json=labs_all.json prefix=[PREFIX] inviter_password=[Your Vendor web password] inviter_email=[your Vendor web email]
```
> NOTE: Use the same PREFIX and `labs_all.json` that was used during creation.

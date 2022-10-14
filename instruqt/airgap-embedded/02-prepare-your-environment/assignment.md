---
slug: prepare-your-environment
id: oxmjwa2nm04e
type: challenge
title: Prepare Your Environment
teaser: Prepare your jumpbox environment for interacting with Replicated
notes:
- type: text
  contents: Making sure you have access to vendor.replicated.com from your jumpbox
    command line
tabs:
- title: Shell
  type: terminal
  hostname: shell
  workdir: /home/replicant
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 300
---

🚀 Let's start
==============


### Configure environment

Log into the Replicated Vendor Portal in the "Vendor Portal" tab using the username
and password printed to your screen in the "Shell" tab.

```
username: [PARTICIPANT_ID]@replicated-labs.com
password: [PARTICIPANT_ID]
```

Once you have the credentials, you can login into the Vendor tab and you should land on the Channels tab.

![Vendor Portal Login](../assets/vendor-portal-login.png)

After logging in, you're going to identify your application and create an API token to use with the
Replicated command-line, then set up some environment variables in your shell to store them.

When you log in, you'll be on our release channels page.

![Release Channels on the Vendor Portal](../assets/vendor-portal-login.png)

Go from your

![Finding Your Application Slug](../assets/cli-setup-quickstart-settings.png)

When you go back to the "Shell" tab you'll set the variable `REPLICATED_APP` to the app slug from the Settings page.


Next, create a `read/write` User API token from your [Account Settings](https://vendor.replicated.com/account-settings)
page:

![Creating an API token](../assets/create-api-token.png)

Note: Ensure the token has "Write" access or you'll be unable create new releases.

Once you have the values, set them in your environment.

```
export REPLICATED_APP=...
export REPLICATED_API_TOKEN=...
```

You can ensure this is working with

```
replicated release ls
```

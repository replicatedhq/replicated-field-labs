---
slug: prepare-for-access
id: oxmjwa2nm04e
type: challenge
title: Prepare for Access
teaser: A short description of the challenge.
notes:
- type: text
  contents: Making sure you have access to vendor.replicated.com from the command
    line
tabs:
- title: Shell
  type: terminal
  hostname: shell
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 600
---

🚀 Let's start
==============

You should have received an invite via email to log into https://vendor.replicated.com -- you'll want to accept this invite and set your password.

### 1. Configure environment

**Important Note:** It is important to logout of any existing session in the Replicated vendor portal so that when clicking on the Labs Account invitation email link it takes you to a specific new registration page where you enter your name and password details.  If you get a login screen then this is probably the issue.

Once registered, your application will be automatically selected:

<p><img width="600" alt="Application Selected" src="../assets/application-selected.png"></p>

Now, you'll need to set up environment variables to interact with vendor.replicated.com and instance.

`REPLICATED_APP` should be set to the app slug from the Settings page.

<p align="center"><img src="../assets/cli-setup-quickstart-settings.png" width=600 alt="Finding your application slug"></img></p>

Next, create a `read/write` User API token from your [Account Settings](https://vendor.replicated.com/account-settings) page:
> Note: Ensure the token has "Write" access or you'll be unable create new releases.

<p align="center"><img src="../assets/create-api-token.png" width="1368" alt="Create and API token"></img></p>

Once you have the values,
set them in your environment.

```
export REPLICATED_APP=...
export REPLICATED_API_TOKEN=...
```

You can ensure this is working with

```
replicated release ls
```



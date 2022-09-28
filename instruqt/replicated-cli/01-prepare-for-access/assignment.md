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
  workdir: /home/replicant
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 300
---

👋 Introduction
===============

This exercise is designed to give you a sandbox to ensure you have a basic understanding how to work with the Replicated CLI.

* **What you will do**:
  * Use the Replicated CLI to vaidate and release your application
  * Find your installation command and install to an existing cluster
  * Iterate on your work and deploy an updated release
* **Who this is for**: This lab is for anyone who works with app code, docker images, k8s yamls, or does field support for multi-prem applications
  * Full Stack / DevOps / Product Engineers
  * Support Engineers
  * Implementation / Field Engineers
* **Prerequisites**:
  * Basic working knowledge of Kubernetes
* **Outcomes**:
  * You will build a working understanding of how to release and deploy a kubernetes application packaged with Replicated using kots.

* * *

🚀 Let's start
==============

### 1. Check Your Email!

If you previously already done any tracks and accepted the invite for the Vendor Portal, you can skip this and go to section 2.

You should have received an invite via email to log into https://vendor.replicated.com -- you'll want to accept this invite and set your password.

**Important Note:** It is important to logout of any existing session in the Replicated vendor portal so that when clicking on the Labs Account invitation email link it takes you to a specific new registration page where you enter your name and password details.  If you get a login screen then this is probably the issue.

The email should look like this:

<p align="center"><img src="../assets/email-invite.png" width=600></img></p>

Once you click on the button, it should open a browser to a page similar to this:

<p align="center"><img src="../assets/create-account.png" width=600></img></p>

Fill in the rest of the form and click on the **Create Account** button to get started.

Once you have created your account you should land on the Channels. Channels allow you to manage who has access to which releases of your application.


### 2. Configure environment

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

### 2. Saving your lab setup

The lab environment you are working in is a little bit different
than your own workstation. Because of the lab structure, we need
to take an extra step to support the next challenges. You would
not need to do this in a real-world environment.

```
save_lab_setup
```

---
slug: prepare-your-jumpbox
id: ztordyohaj5l
type: challenge
title: Prepare Your Environment
teaser: Prepare your jumpbox environment for interacting with Replicated
notes:
- type: text
  contents: "\U0001F44B Introduction\n===============\n\nIn this exercise you will
    learn how to perform installations in air gap environments, and\nhow to collect
    support bundles in air gap environments.\n\n* **What you will do**:\n    * Access
    and verify a single-node air gap setup via a bastion server\n    * Learn to use
    KOTS to install in an air gap environment\n    * Create an SSH tunnel to configure
    an air gap instance\n    * Perform an upgrade of an application runnning in an
    air gap\n    * Collect a support bundle in an air-gapped environment\n* **Who
    this is for**: This lab is for anyone who builds/maintains KOTS applications (see\n
    \   note below)\n    * Full Stack / DevOps / Product Engineers\n* **Outcomes**:\n
    \   * You will be ready to deliver a KOTS application into an air gap environment\n
    \   * You will build confidence in performing upgrades and troubleshooting in\n
    \   air gap environments\n"
- type: text
  contents: "\U0001F512 Packaging\n============\nFirst, we'll push a release -- in
    the background, Replicated's air gap builder will prepare an air gap bundle.\n![Air
    Gap Deployment Packaging](../assets/airgap-slide-1.png)"
- type: text
  contents: "\U0001F512 Delivery\n===========\n\nNext, we'll collect a license file,\n
    \   a download link, and a public kURL bundle.\n\n![Air Gap Deployment Delivery](../assets/airgap-slide-2.png)\n"
- type: text
  contents: "\U0001F512 Deployment\n============\n\nFrom there, we'll move all three
    artifacts into the datacenter via a jump box.\n\n![Air Gap Delivery](../assets/airgap-slide-3.png)\n\nThe
    above diagram shows a three node cluster, but we'll use only a single node.\nWhile
    the KOTS bundle will be moved onto the server via SCP as in the diagram,\nthe
    app bundle and license file will be uploaded via a browser UI through an SSH tunnel.\n"
tabs:
- title: Shell
  type: terminal
  hostname: jumpbox
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 300
---

🚀 Let's start
==============

## Connecting to the Replicated Vendor Portal

Log into the Replicated Vendor Portal in the "Vendor Portal" tab using the following credentials

Username: `[[ Instruqt-Var key="USERNAME" hostname="jumpbox" ]]`<br/>
Password: `[[ Instruqt-Var key="PASSWORD" hostname="jumpbox" ]]`

![Vendor Portal Login](../assets/vendor-portal-login.png)

### Configure environment

When you log in, you'll be on our release channels page.

![Release Channels on the Vendor Portal](../assets/release-channels.png)

Go from the channels page to the settings page to see the application slug. The slug is how the
Replicated CLI and API uniquely identify applications. We'll need to know the slug to use the
CLI later in the lab.

![Finding Your Application Slug](../assets/application-slug-in-settings.png)

When you go back to the "Shell" tab you'll set the variable `REPLICATED_APP` to the app slug. This tells
the `replicated` command which application you are working on without you having to passing it as
an argument to every command.

The other thing you need to work with your application is an API token. You
would usually create `read/write` User API token from the [Account
Settings](https://vendor.replicated.com/account-settings) page. For this lab,
we've created one for you. You set the variable `REPLICATED_API_TOKEN` to the
value of the token.

Set the environment variables in your shell.

```
export REPLICATED_APP=[[ Instruqt-Var key="REPLICATED_APP" hostname="jumpbox"]]
export REPLICATED_API_TOKEN=[[Instruqt-Var key="REPLICATED_API_TOKEN" hostname="jumpbox"]]
```

You can ensure you are authenticated and using the right application by running

```
replicated release ls
```

If you see a list of releases everything is configured correctly.

```
SEQUENCE    CREATED                 EDITED                  ACTIVE_CHANNELS
1           2023-06-23T00:45:18Z    0001-01-01T00:00:00Z    Unstable
```

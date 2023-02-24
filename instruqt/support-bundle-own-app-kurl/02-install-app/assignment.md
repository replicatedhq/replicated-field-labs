---
slug: install-app
id: jtzlouxvykhg
type: challenge
title: Install app
teaser: Install Your Application using Replicated
notes:
- type: text
  contents: Let's install your Application
tabs:
- title: Workstation
  type: terminal
  hostname: cloud-client
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: Cluster Node 1
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-1
- title: Cluster Node 2
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-2
- title: Cluster Node 3
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-3
difficulty: intermediate
timelimit: 3600
---

🚀 Let's begin!
=================

# Vendor Portal login

Log into the Vendor Portal with your existing account, and note your application *app slug*, *release channel*, and the "embedded cluster install command".  It should look something like `curl -sSL https://kurl.sh/<appslug>-<channel> | sudo bash`.  You don't need to install it yet! We have a little bit of setup to complete, first.

# Download your test license

Navigate to the Vendor Portal tab and download the license that you've provisioned for your development work.

  ![Support Bundle Customer](../assets/support-bundle-customer.png)

# Configure the VM environment

## Set up the Workstation
Next, configure the VM environment for automation by exporting the name of your app slug and the release channel.  In the Workstation shell, type:
- `export APP_SLUG=app && export CHANNEL=stable`
replacing `app` and `stable` with your app slug and release channel, and hit Enter.

Then run the following snippet to add these variables to your shell environment - they'll be used to set up the challenges.  You can paste the following snippet entirely:

```shell
echo "export APP_SLUG=${APP_SLUG}" >> ~/.bashrc
echo "export CHANNEL=${CHANNEL}" >> ~/.bashrc
```

# Install the Replicated embedded cluster

## Setup Cluster Node 1
**In the ***Cluster Node 1*** tab** begin your embedded cluster installation.  You're already `root` so you don't need to use `sudo`:

Example: for an HA installation (3 primary nodes)
- `curl -sSL https://kurl.sh/<installer-name> | bash -s ha`

Example: for an installation with a single primary node
- `curl -sSL https://kurl.sh/<installer-name> | bash`

Your embedded installer command may have additional [advanced installation options](https://kurl.sh/docs/install-with-kurl/advanced-options).  Double check with your team for the expected options to use.

*When prompted for the loadbalancer IP address, leave it blank to use the internal LB*

When the install script completes, copy the join command and run it in the *Cluster Node 2* tab and the *Cluster Node 3* tab if you're adding more nodes.  You can also provision additional GCP resources, if needed, from the *Workstation* tab.  Any resources you create will be destroyed at the completion of this exercise.

# Upload your license and install the app

After installation succeeds, navigate to the [App Installer Admin Console](http://loadbalancer.[[ Instruqt-Var key="SANDBOX_ID" hostname="cloud-client" ]].instruqt.io:8800), login and upload your license.

  ![Application installer](../assets/deploy.png)

In the admin console, continue to configure your application, run preflight checks, and deploy your application.

Once your application is deployed and the admin console reports it is ready to use, we can move on to the interactive troubleshooting exercises.

Click "Check" to continue.

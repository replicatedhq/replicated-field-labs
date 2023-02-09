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
- title: Shell
  type: terminal
  hostname: kurl-1
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: intermediate
timelimit: 3600
---

🚀 Let's begin!
=================

# Vendor Portal login

Log into the Vendor Portal with your existing account, and note your application *app slug*, *release channel*, and the "existing cluster install command".  It should look something like `kubectl kots install <app-slug/channel>`.  You don't need to install it yet! We have a little bit of setup to complete, first.

  ![Existing Cluster Install Command](../assets/release-channel.png)

# Download the license

Navigate to the Vendor Portal tab and download the license that you've provisioned for your development work.

  ![Support Bundle Customer](../assets/support-bundle-customer.png)

# Install the Replicated admin console

Configure the VM environment for automation by exporting the name of your app slug and the release channel.  Type the following into your shell, replacing `your-app` and `stable` with your app slug and release channel, and hit Enter.

```shell
export APP_SLUG=your-app
export CHANNEL=stable
```

Then run the following snippet to add these variables to your shell environment - they'll be useful later for automating the challenges.  You can paste the following snippet entirely:

```shell
echo "export APP_SLUG=${APP_SLUG}" >> ~/.bashrc
echo "export CHANNEL=${CHANNEL}" >> ~/.bashrc
```

Then, install your application by executing `kots install` with your app slug.  You can copy/paste this snippet:

```shell
kubectl kots install ${APP_SLUG}/${CHANNEL} \
  --no-port-forward=true
```

# Expose the admin console

To reach the admin console through the VM's firewall, expose the Kubernetes Service for `kotsadm`.  Paste this whole snippet:

```shell
kubectl expose deployment kotsadm \
  -n $(kubectl get pods -A -l app=kotsadm --no-headers | awk '{ print $1 }' ) \
  --type=LoadBalancer \
  --name=kotsadm2 \
  --port=8800 \
  --target-port=3000
```

# Upload your license and install the app

After installation succeeds, navigate to the [Application Installer admin console](http://[[ Instruqt-Var key="HOSTNAME" hostname="kubernetes-vm" ]].[[ Instruqt-Var key="SANDBOX_ID" hostname="kubernetes-vm" ]].instruqt.io:8800), login and upload your license.

  ![Application installer](../assets/deploy.png)

In the admin console, continue to configure your application, run preflight checks, and deploy your application.

Once your application is deployed and the admin console reports it is ready to use, we can move on to the interactive troubleshooting exercises.

Click "Check" to continue.
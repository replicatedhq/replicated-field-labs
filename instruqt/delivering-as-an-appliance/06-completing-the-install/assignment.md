---
slug: completing-the-install
id: am2hqdznfpff
type: challenge
title: Completing the Slackernews Installation
teaser: |
  Complete installing Slackernews with the Admin Console
notes:
- type: text
  contents: |
    Let's walk through the rest of the installation
tabs:
- id: i1b6db1hruhx
  title: Admin Console
  type: website
  url: http://node.${_SANDBOX_ID}.instruqt.io:30000
  new_window: true
- id: omi9odffx8es
  title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 1200
---

We're part way through the installation of Slackernews as a Kubernetes
Appliance. We've run the Slackernews installer which has create a single-node
Kubernetes cluster running all of the infrastructure components of the
Replicated Embedded Cluster. The next step is to complete the install using
one of those components, the Admin Console.

About the Admin Console
=======================

The Replicated Admin Console is a web-based interface that plays two major
roles:

1. It wraps the Helm command with a user-friendly interface that guides the
   user through the installation process.
2. It offers a GUI for "Day Two" operations like checking for (and installing)
   updates, collecting support bundles, and managing cluster nodes.

We're going to focus primarily on the first role for now.

Completing the Install
=====================

To help us complete our installation, the Admin Console will guide us a
through a few steps, starting with configuring its certificate. Since it's
not possible to provide a "safe" certificate to the Admin Console out of the
box, the process starts with a warning and instructions for the user to
validate the self-signed certificate the installer created.

![Explaining the Initial Certificate to Your Customer](../assets/certificate-warning.png)

The next step will ask your customer to configure the permanent certificate
for the Admin Console. They can choose between a self-signed certificate or
uploading one signed by a certificate authority. I recommend that they use a
signed certificate, this is something you may also want to suggest in your
installation documentation.

For the lab, though, we'll stick with the self-signed certificate. When you
click the "Continue to Setup" button, you'll get a warning about the initial
self-signed certificate. Assuming you accept the risk, you'll get a form that
allows you to set up the Slackernews Admin Console certificate. Keep it set on
"self-signed" and click "Continue".

![Configuring the Admin Console Certificate](../assets/certificate-configuration.png)

You'll be asked to log into the Admin Console on the next page. The password
for the Admin Console is the password you specified in the last step. It will
be `[[ Instruqt-Var key="ADMIN_CONSOLE_PASSWORD" hostname="node" ]]` if you use the suggestion
from the lab text or skipped ahead to this step.

### Adding Nodes to the Cluster

The first thing you'll see when you log in to the Admin Console is the Node
Management screen. This is where you can add additional nodes to the
Slackernews Embedded Cluster. Node management is the first step in the install
process because not all applications can run on a single node.

![Admin Console Node Managemet Interface](../assets/node-management.png)

To simplify this lab, we're going to stick with a single node cluster. Click
"Continue" to continue the installation on one node. Slackernews is a fairly
lightweight application and can run fairly easily on a single node.

### Configuring Your Instance

Next is the configuratiion screen that we defined earlier. Fill in the values
for each entry.

![Entering the Slackernews Configuration in the Admin Console](../assets/providing-configuration.png)

The first two sections are fairly straightforward:

* **Slackernews Domain**: `[[ Instruqt-Var key="SLACKERNEWS_DOMAIN" hostname="node" ]]`
* **Admin Users**: `[[ Instruqt-Var key="ADMIN_EMAIL" hostname="node" ]]`
* **Service Type**: NodePort
* **Node Port**: `30443`
* **Certificate Source**: Generate

The next section presents us with a bit of a challenge. We don't have a Slack
team to connect to, so any values we enter will be invalid. The values are
required, though, so we need to enter something. Let's use the following dummy
values:

* **Slack Client ID**: `notavalidid`
* **Slack Client Secret**: `notavalidsecret`
* **User OAuth Token**: `xoxp-notavalidtoken`
* **Bot User Auth Token**: `xoxb-notavalidtoken`

Note that we prefixed the two tokens. Those are the prefixes used by actual Slack
tokens. We needed to do this to pass the validation on those fields.

Click "Continue" to move to the next steps and run our preflight checks.
Preflight checks run after your customer specifies their configuration so that
you can use the configuration values as part of your checks.

### Running Preflight Checks

The Admin Console will run preflight checks it finds in the release. Preflight
checks are resource processed by the Admin Console to ensure that your
customer has met the prerequisites to successfully install your application.
They can be included as standard manifests within the Embeddded Cluster
release or as secrets within your application's Helm chart.

We strongly recommend including them in your Helm chart so that they can be
used regardless of installation method. See [Avoiding Installation
Pitfalls](https://play.instruqt.com/manage/replicated/tracks/avoiding-installation-pitfalls)
for more details on writing preflight checks and including them in your
application.

We didn't add any preflight checks in this lab because the Slackernews Helm
chart already includes them. They check for some basic prerequisites:

1. The version of Kubernetes is one that the Slackernews team is prepared to
   support. This is a less relevant check for an Embedded Cluster installation
   since the version of Kubernetes is tightly coupled to the application
   release, but it's critical when installing into an existing cluster.
2. Access to the Slack API. This is probably the most critical check for the
   Slackernews application. The application provides no value if it cannot
   connect to Slack, and a lot of production servers are configured to limit
   access to the internet. This check will fail if the server cannot connect
   and will provide guidance on how to resolve the issue.
3. The amount of memory available in the cluster. Slackernews is not terribly
   resource intensive, but this is a good check for most applications and the
   Slackernews team wants to make sure they have space to run.

All of your checks should pass, and you should see a screen like the one
below. Click "Deploy" to continue.

![Successful Slackernews Preflight Checks](../assets/successful-preflight-checks.png)

If one or more check had failed or provided a warning, you may still have had
the option to continue the deployment. You can decide when writing your
preflight checks whether a failure stops a deployment or allows it to
continue.

### Validate the Installation is Running

The Admin Console played a little trick on you. Since all of your preflight
checks passed it continued the installation without waiting for you to click
"Deploy". I note this because the instance it may be running successfully
when you land on the main Admin Console screen. Otherwise, just wait until you
see the status change to "Ready" underneath the title "Slackernews".

![Slackernews Admin Console with "Ready" Status](../assets/slackernews-admin-console-ready.png)

While you're connected to the Admin Console, notice some of it's other
capabilities:

* You can manage the cluster nodes using the "Cluster Management" tab.
* You can use the "Troubleshoot" tab to collect and analyze support bundles
  defined as part of the application.
* You can reconfigure the application using the "Config" tab.

Each of these is beyond the scope of the current lab, but feel free to poke
around at them before you move on.

Your Perspective on the Install
===============================

You've just run through the customer experience of installing Slackernews with
the Replicated Embedded Cluster. Before we wrap up the lab, let's take a
moment to see the how this new instance looks on the Replicated Vendor portal.
Knowing what happened when your customer installed is an important piece of
information for you as as software vendor.

Use the "Vendor Portal" tab to connect to the Replicated Vendor Portal. You may
still be logged in from the previous step. If not, log in again
with the following credentials:

Username: `[[ Instruqt-Var key="USERNAME" hostname="node" ]]`<br/>
Password: `[[ Instruqt-Var key="PASSWORD" hostname="node" ]]`

Navigate to the "Customers" page in the left hand menu and you'll see that
Nitflex has a new instance of the application running.

![Customer Page Showing New Running Instance for Nitflex](../assets/new-nitflex-instance.png)

Click in on the instance name to see details about the instance. You not only
see information about what version of your application they are running, but
also details about their preflight checks, the details of the cluster, even
the cloud provider. This information can be useful to your support and product
teams as the support the customer journey and evolve your product.

![Instance Details for the New Nitflex Instance](../assets/new-nitflex-instance.png)

🏁 Finish
=========

You've now successfully build and installed a Kubernetes appliance with the
Replicated Embedded Cluster. You've seen how the Admin Console guides the
installation and had a look what features it supports. Your current
installation of Slackernews runs on a single-node cluster, but you've seen how
your customer could expand it to have additional nodes.


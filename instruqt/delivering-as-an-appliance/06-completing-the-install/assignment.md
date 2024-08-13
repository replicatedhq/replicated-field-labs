---
slug: completing-the-install
id: f0ae2ipcuup0
type: challenge
title: Completing the Slackernews Installation
teaser: |
  Complete installing Slackernews with the Admin Console
notes:
- type: text
  contents: |
    How a customer installs with a Kubernetes appliance
tabs:
- title: Admin Console
  type: website
  url: http://node.${_SANDBOX_ID}.instruqt.io:30000
  new_window: true
- title: Slackernews
  type: service
  hostname: node
  port: 30443
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

We're going to focus primarily on the first role for now. In the next step of
the lab we'll also look at managing cluster nodes.

Completing the Install
=====================

To help us complete our installation, the Admin Console will guide us a
through a few steps, starting with configuring its certificates. Since it's
not possible to provide a "safe" certificate to the Admin Console out of the
box, the process starts with a warning and instructions for the user to
validate the self-signed certificate the installer created.

![Explaining the Initial Certificate to Your Customer](../assets/certificate-warning.png)

The next step will ask your customer to configure the permanent certificate
for the Admin Console. They can choose between a self-signed certificate or
uploading one signed by a certificate authority. I recommend that they use a
sign certificate, this is something you may also want to suggest in your
installation documentation.

For the lab, though, we'll stick with the self-signed certificate. When you click the
"Continue to Setup" button, you'll get a warning about the initial
certificate. Assuming you accept the risk, you'll get a form that allows you
to set up the Slackernews Admin Console certificate. Keep it set on
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
process because not all applications can run on a single node. Slackernews is
a fairly lightweight application and can run fairly easily on a single node.
To simplify this lab, we're going to stick with a single node cluster.

### Configuring Your Instance

Once you've logged in, the Admin Console will guide you through a few more
steps, starting with configuring Slackernews. You'll see the configuration
form we defined earlier. Fill in the values for each entry.

The first two sections are fairly straightforward:

* **Slackernews Domain**: `[[ Instruqt-Var key="SLACKERNEWS_DOMAIN" hostname="cluster" ]]`
* **Admin Users**: `[[ Instruqt-Var key="ADMIN_EMAIL" hostname="node" ]]`
* **Service Type**: NodePort
* **Node Port**: `30443`
* **Certificate Source**: Generate

The next section presents us with a bit of a challenge. We don't have a Slack
team to connect to, so any values we enter will be invalid when we run our
pre-flight checks. The values are required, though, so we need to enter
something. Let's use the following dummy values:

* **Slack Client ID**: `notavalidid`
* **Slack Client Secret**: `notavalidsecret`
* **User OAuth Token**: `xoxp-notavalidtoken`
* **Bot User Auth Token**: `xoxb-notavalidtoken`

Note that we prefixed the two tokens with the prefix used by actual Slack
tokens. We needed to do this to pass the validation set set up for the form.

Click "" to move to the next steps and run our preflight checks.

### Running Preflight Checks

### Validate the Installation is Running

### Observing the Instance in Vendor Portal

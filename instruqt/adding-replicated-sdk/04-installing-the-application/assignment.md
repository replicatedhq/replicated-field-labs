---
slug: installing-the-application
id: lxlwklpzidy0
type: challenge
title: Installing the Application
teaser: Let's install the application as your customer
notes:
- type: text
  contents: Let's see how your customer installs an application
tabs:
- title: Shell
  type: terminal
  hostname: shell
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 600
---

Now that we have a release in the Replicated Platform, you can
distribute it's Helm chart to you customers using entitlements
that we manage for you. In this step, we're going to install the
Mastodon Helm chart the same way a customer would install your
application.

Getting the Install Instructions
================================

We're going to use the Replicated Vendor Portal to look up the
installation instructions of the customer Omozan. The Vendor
Portal is a core interface into the platform. We'll use it again
later in this lab to review the telemetry information we receive
from the SDK.

Click on the Vendor Portal tab to open up a new browser window and
access the portal. Log in with these credentials

Username: `[[ Instruqt-Var key="USERNAME" hostname="shell" ]]`<br/>
Password: `[[ Instruqt-Var key="PASSWORD" hostname="shell" ]]`



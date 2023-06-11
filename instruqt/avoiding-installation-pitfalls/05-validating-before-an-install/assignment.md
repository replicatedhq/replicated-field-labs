---
slug: validating-before-an-install
id: aray1djn6poy
type: challenge
title: Validating Before an Install
teaser: A short description of the challenge.
notes:
- type: text
  contents: Replace this text with your own text
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

Having a release with your preflights included means you can
take advantage of the Replicated Platform to distribute those
preflights. The entitlements that provide access to your
application control access to the preflight checks as well.
To run the preflights, your customer templates out your
Helm chart and pipes the output to the `preflight` plugin
to `kubectl`, just like you did to test your changes.

Logging Into the Vendor Portal
==============================

To run the preflights as a customer, we need to have their
login credentials to the Replicated registry. The lab setup
process configured a customer for the Harbor application,
but in this step we're going to add a new customer. We'll
do this in the Replicated Vendor Portal,.

Click on the Vendor Portal tab to open up a new browser window and
access the portal. Log in with these credentials

Username: `[[ Instruqt-Var key="USERNAME" hostname="shell" ]]`<br/>
Password: `[[ Instruqt-Var key="PASSWORD" hostname="shell" ]]`

You'll land on the "Channels" page for your app, which will show
the release channels we discussed in the previous step. Notice that
each of the default channels shows the current version `16.8.0`,
while the channel LTS, which we haven't released to, reflects
that.

![Vendor Portal Release Channels](../assets/vendor-portal-landing.png)

Creating a Customer
===================


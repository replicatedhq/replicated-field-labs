---
slug: validating-before-an-install
id: esalpc0nrb1q
type: challenge
title: Validating Before an Install
teaser: How your customer uses preflights to validate their environment
notes:
- type: text
  contents: Experience how your customer executes your preflight checks
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
login credentials to the Replicated registry. The lab set up
process configured a customer for the Slackernews application,
but in this step we're going to add a new customer. We'll
do this in the Replicated Vendor Portal.

Click on the Vendor Portal tab to open up a new browser window and
access the portal. Log in with these credentials:

Username: `[[ Instruqt-Var key="USERNAME" hostname="shell" ]]`<br/>
Password: `[[ Instruqt-Var key="PASSWORD" hostname="shell" ]]`

You'll land on the "Channels" page for your app, which will show
the release channels we discussed in the previous step. Notice that
each of the default channels shows the current version `0.3.0`,
while the channel LTS, which we haven't released to, reflects
that.

![Vendor Portal Release Channels](../assets/vendor-portal-landing.png)

Creating a Customer
===================

To create a customer, select "Customers" from the menu on the left,
then click the "+ Create Customer" button.

![Creating a Customer](../assets/create-customer-button.png)

Call your customer "Geeglo" and assign them to the `Stable` channel.
Let's assume this customer is not yet a customer, but a prospect who
is evaluating our application. Let's allow our customer to
evaluate our software on a trial license for 30 days.  Set them up
as a "Trial" customer and set up the expiration date and customer
email as follows:

Expiration Date: `[[ Instruqt-Var key="LICENSE_EXPIRY" hostname="shell" ]]`<br/>
Customer Email: `[[ Instruqt-Var key="CUSTOMER_EMAIL" hostname="shell" ]]`

![Customer Details](../assets/new-customer-details.png)

Running Preflight Checks
========================

Your customer will prepare for their installation by running
the preflights checks. Since the preflights are embedded in
the Slackernews Helm chart, they first need to log into the Replicated
registry with the `helm` command. This gives them access to
your Helm chart via the Replicated Platform.

If you completed the [Distributing Your Application with Replicated](https://play.instruqt.com/replicated/tracks/distributing-your-application-with-replicated)
lab, you'll remember that each customer has their own login
credentials that are specific to their entitlements to your
application. You can get those instructions by clicking the
"Helm Install Instructions" button on the customer page.

![Install instruction](../assets/install-instructions.png)

In this lab, we only want to execute the first step in the process,
which is to log into the registry. Because you just created
the customer, you'll have to copy the login command from the
customer page.

From there, they can use the `helm template` command to extract the
preflight checks to run with `kubectl preflight`.

```
helm template oci://registry.replicated.com/[[ Instruqt-Var key="REPLICATED_APP" hostname="shell" ]]/slackernews \
  | kubectl preflight -
```

Running the checks this way shows them the same screen you saw when
you tested the preflights earlier, with the Kubernetes version check
passing, the memory check failing, and a warning for the CPU available.

![Customer Preflight Checks](../assets/customer-preflight-checks-before.png)

To have a successful installation, the customer has to make a decision
about whether they should install to a different cluster or add
more resources to this cluster.

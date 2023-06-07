---
slug: installing-the-application
id: f9ve4gcr0dzn
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
Harbor Helm chart the same way a customer would install your
application.

Logging Into the Vendor Portal
==============================

We're going to use the Replicated Vendor Portal to look up the
installation instructions of the customer Omozan. The Vendor
Portal is a core interface into the platform. We'll use it again
later in this lab to review the telemetry information we receive
from the SDK.

Click on the Vendor Portal tab to open up a new browser window and
access the portal. Log in with these credentials

Username: `[[ Instruqt-Var key="USERNAME" hostname="shell" ]]`<br/>
Password: `[[ Instruqt-Var key="PASSWORD" hostname="shell" ]]`

You'll land on the "Channels" page for your app, which will show
the release channels we discussed in the previous step. Notice that
each channel shows the current version `16.7.0`.

![Vendor Portal Release Channels](../assets/vendor-portal-landing.png)


Getting the Install Instructions
================================

Installation instructions are specific to each customer, since they
require unique login credentials for the Replicated registry. We're
going to install as the customer "Omozan" that has been set up as
part of the lab.

Select the "Customers" link from the left navigation. You'll be on
the customer landing page and see the adoption graph for the application.
The graph is currently empty because the application hasn't been
installed anywhere yet.

![Customers Landing Page](../assets/customers-page.png)

Below the graph you'll see the list of customers, with the customer
"Omozon" as the only one in the list. Click on their name and you'll
be brought to their customer page. In the top right corner you'll
see a link to their install instructions.

![Customers Landing Page](../assets/single-customer-page.png)

Click on the link and you'll be prompted to enter a customer
email. You can use any address you want, but the rest of the
instructions assume you used the username above
(`[[ Instruqt-Var key="USERNAME" hostname="shell" ]]`).

![Install Instructions](../assets/helm-install-instructions.png)

You're going to use these instructions to complete your install.
We'll skip the preflight checks for this lab since we haven't
added any to our chart.

Installing the Application
==========================

Your customer starts their installation by logging into our 
registry with the `helm` command. This gives them access to 
your Helm chart via the Replicated Platform.

```
helm registry login registry.replicated.com \
  --username [[ Instruqt-Var key="CUSTOMER_EMAIL" hostname="shell" ]] \
  --password [[ Instruqt-Var key="REGISTRY_PASSWORD" hostname="shell" ]]
```

From there, they do a simple Helm install. In our case, we're going to
tack some additional values that Harbor needs to come up correctly. 
This helps us make sure the installation is complete before we move
onto the next step in the lab.

```
helm install harbor \
  oci://registry.replicated.com/[[ Instruqt-Var key="REPLICATED_APP" hostname="shell" ]]/harbor \
  --set service.type=NodePort --set nodePort.https=443 \
  --set externalURL=https://[[  Instruqt-Var key="EXTERNAL_URL" hostname="cluster" ]]
```

---
slug: create-a-customer
id: rwexuqt6fnmv
type: challenge
title: Creating a Customer License
teaser: Time to add a customer and configure their license
notes:
- type: text
  contents: |
    Add a customer to your application and configure their license
tabs:
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 600
---


Creating a Customer License
===========================

Let's connect to the Replicated Vendor Portal and create a new customer. This
will also create their license. Click on the "Open External Window" button to
open a new browser window and access the portal. Log in with these credentials:

Username: `[[ Instruqt-Var key="USERNAME" hostname="shell" ]]`<br/>
Password: `[[ Instruqt-Var key="PASSWORD" hostname="shell" ]]`

You'll land on the "Channels" page showing the default release channels.

![Vendor Portal Release Channels](../assets/vendor-portal-landing.png)

To create a customer, select "Customers" from the menu on the left, and you'll
see your two existing customers "Omozan" and "Geeglo".

![Your Existing Customers](../assets/customer-landing-page.png)

We're going to assume you've just landed a new customer named "Nitflex" and
create them. The platform also includes an
[API](https://replicated-vendor-api.readme.io/v3/reference/createapp) you can
use to automate customer creation as part of your existing onboarding workflow.
For the purpose of the lab, click the "+ Create Customer" button to create
"Nitflex" manually.

![Creating a Customer](../assets/create-customer-button.png)

Enter the name "NitFlex" and assign them to the `Stable` channel. They've
subscribed to your software for a year, so let's make sure we capture the
expiration date. We also need an email for them to use as a login to install
your Helm chart. Note that we never use that email, it's your customer, not
ours.

Expiration Date: `[[ Instruqt-Var key="LICENSE_EXPIRY" hostname="shell" ]]`<br/>
Customer Email: `[[ Instruqt-Var key="CUSTOMER_EMAIL" hostname="shell" ]]`

![Customer Details](../assets/new-customer-details.png)


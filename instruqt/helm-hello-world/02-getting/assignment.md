---
slug: getting
id: juwinvh4ipce
type: challenge
title: Getting Started
teaser: Let's make sure you have access to vendor portal
notes:
- type: text
  contents: |-
    Let's start by making sure you have access to the Vendor Portal, so we'll cover:
    - Accepting the Invite to Vendor Portal
    - Walkthrough of the Application
    - Download the customer license
tabs:
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 600
---

🚀 Let's start
==============

### 1. Vendor Portal login

On the **Vendor** tab, login to the Vendor Portal using the credentials below.

Username: `[[ Instruqt-Var key="USERNAME" hostname="kubernetes-vm" ]]`<br/>
Password: `[[ Instruqt-Var key="PASSWORD" hostname="kubernetes-vm" ]]`

You will land on the Release Channels page showing the release channels for the Wordpress application. Channels allow you to manage who has access to which releases of your application.

### 2. Review the Application

The default channels are `Stable`, `Beta` and `Unstable`.

On the Stable channel card, click on **Release history** to get the list of releases. Here you can see all of the releases that have been promoted to this channel.

<p align="center"><img src="../assets/hellohelmchannel.png" width=300></img></p>

We want to view the contents of the latest release, so to do that click on the **View Release YAML** icon as shown below:

<p align="center"><img src="../assets/releases-channel.png" width=600></img></p>

You will see a file navigator similar to the one shown below. This view shows you the content of the current release. As you can see there are some files above the line and files below it. The files above are files used to configure some of the Replicated features. The files below are the ones needed to deploy the application, which in our case is Wordpress.

<p align="center"><img src="../assets/release-contents.png" width=600></img></p>

As you can see we are using the Wordpress Helm Chart, and in this view, the top level **Chart.yaml** and **Values.yaml** file are exposed. The **wordpress.yaml** is a file that declares how Replicated will manage the Chart. For example, you can override the default values, set up rules for optional charts and more.

**Managing Values**

When installing from a Helm Chart, there are scenarios where the default values need to be overridden or preset for a given customer. With Replicated, you can map the values in the **Values.yaml** file with values that an end user can enter in a config UI or from a Replicated License file. Below is a screenshot of the **wordpress.yaml** file with some value overrides:

<p align="center"><img src="../assets/values-overide.png" width=600></img></p>

Note that for some values, the value is not a hard coded value, rather it has something like `repl{{ ConfigOption ... }}`. Replicated supports [templating](https://docs.replicated.com/vendor/packaging-template-functions) which is how you can dynamically assign values.

## 3. Copy Install Command

Go back to **Channels** and go to the `Stable` channel. On the bottom of the channel card, select to copy the install command for `Existing Cluster`

<p align="center"><img src="../assets/install-command.png" width=600></img></p>

## 4. Download Customer File

A customer license (downloadable as a `.yaml` file) is required to install any KOTS application.
To create a customer license, go to `Customers > Wordpress Customer` by selecting the "Customers" link on the left in the Vendor Portal.

<p align="center"><img src="../assets/helm-customer-list.png" width=600></img></p>

You can view the customer details by clicking the row.
For this Hello World exercise we'll use `Wordpress Customer`.
You'll notice that the customer is assigned to the the "Stable" channel on the right hand side, and the Customer Type is set to "Development".
When you've reviewed these, you can click the "Download License" link in the top right corner.

<p align="center"><img src="../assets/helm-cust-details.png" width=600></img></p>

This will download the file with your customer name and a `.yaml` extension.
This is the license file a customer would need to install your application.

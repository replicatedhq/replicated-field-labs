---
slug: enable-feature
id: 0ql5zjsxwew6
type: challenge
title: Enable the Feature
teaser: Next, we'll change the feature flag and update the application!
notes:
- type: text
  contents: |- 
    In this challenge we'll cover:
    - Updating the Custom License Field
    - updating the deployed application
tabs:
- title: Shell
  type: terminal
  hostname: shell
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: Admin Console
  type: website
  url: http://kubernetes-vm.${_SANDBOX_ID}.instruqt.io:8800
  new_window: true
difficulty: basic
timelimit: 600
---

## Update the Customer License

Let's now enable the feature for this customer.

Navigate back to **Customers** and open the customer created earler and set the **Enable Fetaure** custom field to `true`:

<p align="center"><img src="../assets/lic-updated-customer.png" width=600></img></p>

Save the changes. The feature is now available for the customer. If the customer has an online installation you do not need to download the license again.


## Update Deployed App

Click on the **Admin Console** tab to access the Admin Console.

Click on **sync licenses** highlighted below

<p align="center"><img src="../assets/lic-sync-licenses.png" width=600></img></p>

A dialog is displayed explaining that there is a new version of the application. This is because the underlying manifests have been modified by the license change.

<p align="center"><img src="../assets/lic-new-version.png" width=600></img></p>

Click on **Version History** which will take you to **Version history**

<p align="center"><img src="../assets/lic-version-tab.png" width=600></img></p>

Click on **Deploy** to deploy the update that should now include the Super Duper feature.

Once the update is deployed, click on the Dashboard tab and once the application is **Ready** click on the **Open Custom Fields App** link, which should now look like this:

<p align="center"><img src="../assets/lic-dashboard-updated.png" width=600></img></p>

You may need to clear your cache or do a hard refresh to see the change

<p align="center"><img src="../assets/lic-updated-app.png" width=600></img></p>

If you receive the above result, pat yourself on the back for a great job done!

This track is now complete.

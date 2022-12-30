---
slug: creating-custom-fields
id: 0efpwhe3qy3q
type: challenge
title: Creating Custom Fields
teaser: Learn how to create Custom License Fields
notes:
- type: text
  contents: Let's get you logged in to Vendor Portal and create a Custom License Field!
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 600
---
### Vendor Portal login

In this lab, we'll use your partipant ID for your credentials in the Vendor Portal. The credentials are included in the **Shell** tab and are displayed in the following format:

```
username: [PARTICIPANT_ID]@replicated-labs.com
password: [PASSWORD]
```

To access the Vendor Portal, click on the **Vendor** tab and login to the Vendor Portal. Once you log in, you should land on the Channels, which allow you to manage who has access to which releases of your application.

### Create Custom Field

Let's create the [Custom License Field](https://docs.replicated.com/vendor/licenses-adding-custom-fields) that we will use as a switch for our Super Duper Feature:

Navigate to **License Fields** on the left hand side navigator.

<p align="center"><img src="../assets/nav-lic-fields.png" width=600></img></p>

Click on **Create custom field** and use the following values:

* **Field:** enable-feature
* **Title:** Enable Feature
* **Type:** Boolean
* **Default:** false

Leave the **Required** and **Hidden** boxes unchecked. Below is a screenshot of what it should look like:

<p align="center"><img src="../assets/create-field.png" width=450></img></p>

Click on the **Create** button to create the field. Stay logged in to the Vendor Portal as we will need it in the next challege.

Congratulations! You have completed this challenge.

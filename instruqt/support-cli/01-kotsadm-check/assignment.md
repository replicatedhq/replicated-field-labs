---
slug: kotsadm-check
id: rxsy3qldj4b2
type: challenge
title: kotsadm check
teaser: Initial application check via kotsadm
notes:
- type: text
  contents: Please wait while your environment is provisioned..
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 1200
---

👋 Initial Application Check
============================

**In this exercise you will:**

 * Check the application status via the kotsadm web console


### 1. Access Vendor portal to download license

Click on the Vendor tab to launch the Vendor Portal login in new browser tab.

To access the Vendor Portal, you will need your participant id. If you go to the Shell tab, it will show you the username and password to be used for the Vendor tab. It will be of the following format:
```
username: [PARTICIPANT_ID]@replicated-labs.com
password: [PARTICIPANT_ID]
```

Once you have the credentials, you can login into the Vendor tab and you should land on the Channels tab.

### 2. Download Application License

A sample end customer has been pre-created and associated with the Stable release channel where the test application release has been promoted to.  View this customer by navigating to the Customers tab on the left hand side of the UI, the customer name is *Hola SupportCLI Customer*.

Click on the download license icon on the right of the customer entry as you'll use that in the next challenge.

![license-dlicon](../assets/license-download-icon.png)



To complete this challenge, press **Check**.

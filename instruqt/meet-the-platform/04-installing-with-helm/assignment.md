---
slug: installing-with-helm
type: challenge
title: Installing Your Application with Helm
teaser: A short description of the challenge.
notes:
- type: text
  contents: Replace this text with your own text
tabs:
- title: Enterprise Portal
  type: website
  url: https://get.replicated.com
  new_window: true
- title: Customer Terminal
  type: terminall
  hostname: shell
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 600
enhanced_loading: null
---

Once the preflight checks have passed, your customer continues on to install
your application using your Helm chart. The chart is pulled from the
Replicated Proxy Registry using credentials that are tied to their license.

> Note: This step will take them through the Helm install. After the install
> is finished, we'll guide them to see the instance exists in both the
> Enterprise Portal and the Vendor Portal, we'll look into the instance
> details in the next step.

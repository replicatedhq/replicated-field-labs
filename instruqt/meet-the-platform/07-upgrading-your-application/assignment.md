---
slug: upgrading-your-application
type: challenge
title: Upgrading Your Application
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

After you diagnose your customer's issue, you may find you need to release a
patch to solve the issue. We've released a patch for you so that you can focus
on the experience you'll have as your customer deploys the fix.

> Note: We'll have the customer review the update in the Enterprise Portal
> (including the Securiy Center) and then run `helm upgrade` in their terminal
> to deploy the patched version. From their, we'll see that they've upgraded
> in the Vendor Portal to know that all is well.

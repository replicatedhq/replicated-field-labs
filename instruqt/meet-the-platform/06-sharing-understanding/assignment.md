---
slug: sharing-understanding
type: challenge
title: Sharinng More Details for Troubleshooting
teaser: A short description of the challenge.
notes:
- type: text
  contents: Replace this text with your own text
tabs:
- title: Customer Terminal
  type: terminal
  hostname: shell
- title: Enterprise Portal
  type: website
  url: https://get.replicated.com
  new_window: true
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 600
enhanced_loading: null
---

As much as we hate it, sometimes something goes wrong with our applications.
This has been one of the main arguments against self-hosted software in the
past. It's just so much easier to help when you have access to everything.
The Replicated Platform appreciates this and provides a powerful
troubleshooting component to close the gap between you and your customers.

> Note: This step will take them trough colelcting a support bundle as a
> customer and uploading it to the Enterprise Portal. We'll then switch over
> to the Vendor Portal and review the support bundle. It will show a fairly
> obvious error that is caught by an analyzer, but also discuss how they can
> review all the files in the bundle in the Vendor Portal or download it for
> use with `sbctl`. We won't have them use `sbctl` but instead instruct them
> to the support bundle lab.

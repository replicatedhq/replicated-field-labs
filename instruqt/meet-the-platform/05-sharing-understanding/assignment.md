---
slug: sharing-understanding
id: yod1u3o4xn68
type: challenge
title: Sharinng More Details for Troubleshooting
teaser: A short description of the challenge.
notes:
- type: text
  contents: Replace this text with your own text
tabs:
- title: SlackerNews
  type: website
  url: https://slackernews.io   # replace with generated URI
- id: 54o8pdlzguw3
  title: Customer Terminal
  type: terminal
  hostname: shell
- id: 1yjw8bibrupy
  title: Enterprise Portal
  type: website
  url: https://get.replicated.com
  new_window: true
- id: gnagj9cpljs9
  title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 600
enhanced_loading: null
---

> Note: Before we start this step, we'll deliberately break something in the
> running instance so we can show something real going wrong in the support
> bundle and use that as a way to elaborate on what the support bundle is
> doing.

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

When Something Goes Wrong
=========================

> Note: Have them try to access the application 

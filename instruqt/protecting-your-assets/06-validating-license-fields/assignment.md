---
slug: validating-license-fields
id: 7zuxim1ohfcx
type: challenge
title: Validating License Fields with the Replicated SDK
teaser: Calling Replicated SDK to assure customer entitlements
notes:
- type: text
  contents: Learn how to validate the entitlements we just created
tabs:
- title: Shell
  type: terminal
  hostname: shell
  workdir: /home/replicant
difficulty: basic
timelimit: 600
---

The Replicated SDK provides an in-cluster API for the Replicated platform. One
of it's core features is access to the complete customer license. We're going
to run some simple shell commands against the SDK to show how your team can
validating entitlements as part of your application.

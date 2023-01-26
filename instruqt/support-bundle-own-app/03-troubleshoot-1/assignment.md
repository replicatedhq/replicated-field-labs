---
slug: troubleshoot-1
id: wmm7b9ophrjc
type: challenge
title: Correcting the broken application
teaser: Time to fix the problem
notes:
- type: text
  contents: Time to fix the problem...
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
- title: Application Installer
  type: website
  url: http://kubernetes-vm.${_SANDBOX_ID}.instruqt.io:8800
  new_window: true
difficulty: basic
timelimit: 3600
---

🚀 Let's start
=================

Now we will explore solving an application problem in *[[ Instruqt-Var key="APP_SLUG" hostname="kubernetes-vm" ]]/[[ Instruqt-Var key="CHANNEL" hostname="kubernetes-vm" ]]*.  Imagine: you are supporting a customer and they report to you that one of their application pods is crashing.  How do you begin to solve the problem?

💡 Hints
=================
- How do you list pods?
- How do you describe pods?
  - What if you wanted to see events from multiple pods at once?
- How do you get logs from a pod?
  - What if you wanted to see a previous version of the pod's logs?
- When would you look at `describe` output vs. gathering pod logs?

✔️ Solution
=================
A random deployment has been selected and the memory limit reduced to 10Mi.  This will cause the application to crash.

Remediation
=================
Patch or edit the affected deployment to increase the memory request and limit to a reasonable amount.

- How can you edit or patch a resource in-place?
- How can you edit or patch a resource from a file?
- How can we make sure that this doesn't happen again?

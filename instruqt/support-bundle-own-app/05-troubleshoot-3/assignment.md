---
slug: troubleshoot-3
id: vbniffgw5lxc
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

iptables/coredns issue

TODO: find the iptables rules for the coredns pod, and drop them

should simulate DNS failures across the cluster

💡 Hints
=================

✔️ Solution
=================

Remediation
=================

---
slug: troubleshoot-2
id: fvqrxka6fxu8
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

You get another report from a customer saying that the application isn't working, as if two components are unable to communicate.  How would you begin to solve the problem?

💡 Hints
=================

- Think about the traffic flow to your application
- How does traffic get to workloads inside kubernetes

✔️ Solution
=================

A random service has been selected and the port number changed to a random port number.

Remediation
=================

Patch or edit the affected service to correct the port number. you may have to refer to the other resources in the cluster to identify the correct port number.

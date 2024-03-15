---
slug: troubleshoot-1
id: araxpgiqal1r
type: challenge
title: Correcting the broken application
teaser: Where are my pods?
notes:
- type: text
  contents: Time to fix the problem...
tabs:
- title: Workstation
  type: terminal
  hostname: cloud-client
difficulty: basic
timelimit: 3600
---
# [App Installer Admin Console](http://loadbalancer.[[ Instruqt-Var key="SANDBOX_ID" hostname="cloud-client" ]].instruqt.io:8800)

🚀 Let's start
================

Let's imagine that our environment belongs to a customer, who are now experiencing an issue with their install.

They've raised a rather vague issue to your support team suggesting that the application "doesn't work" after one of their admins was playing with the setup.

Use the `sbctl` tool to inspect a support bundle they've shared with you, and try to determine what's amiss.

💡 Hints
=================

- use `sbctl shell -s /path/to/bundle.tar.gz` to get kubectl access to the bundle

- How are applications deployed in kubernetes?

- What controls a pod's lifecycle?

💡 More Hints
=================

- How do I see deployments?

Troubleshooting Procedure
=================

Identify the problematic deployment from `kubectl get deployments -n <namespace>`.  Notice any pods that have 0 replicas.

✔️ Solution
=================

A deployment has been scaled to 0

Remediation
=================

Patch or edit the affected deployment to increase the memory request and raise replicas to desired ammount.

```bash
kubectl scale deployment <deployment-name> --replicas=1
```

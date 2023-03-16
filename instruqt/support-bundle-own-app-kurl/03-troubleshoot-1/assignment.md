---
slug: troubleshoot-1
id: sgroektituzf
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
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: Cluster Node 1
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-1
- title: Cluster Node 2
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-2
- title: Cluster Node 3
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-3
difficulty: intermediate
timelimit: 3600
---
[App Installer Admin Console](http://loadbalancer.[[ Instruqt-Var key="SANDBOX_ID" hostname="cloud-client" ]].instruqt.io:8800)

🚀 Let's start
================
Let's imagine that the embedded cluster and app we just installed was for a customer, who are now experiencing an issue with their install.

They've raised a rather vague issue to your support team suggesting that the application "doesn't work" after one of their admins was playing with the setup.

We'll start by exploring how to solve an application problem in *[[ Instruqt-Var key="APP_SLUG" hostname="cloud-client" ]]/[[ Instruqt-Var key="CHANNEL" hostname="cloud-client" ]]*.


💡 Hints
=================
- How are applications deployed in kubernetes?

- What controlls a pod's lifecycle?

💡 More Hints
=================
- How do I see deployments?

Troubleshooting Procedure
=================

Identify the problematic deployment from `kubectl get deployments -n <namespace>`.  Notice any pods that have 0 replicas.

✔️ Solution
=================
A random deployment has been selected and scaled to 0

Remediation
=================
Patch or edit the affected deployment to increase the memory request and raise replicas to desired ammount.

```bash
kubectl scale deployment <deployment-name> --replicas=1
```

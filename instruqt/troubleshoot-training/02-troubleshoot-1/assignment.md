---
slug: troubleshoot-1
id: araxpgiqal1r
type: challenge
title: Correcting the broken application
teaser: Where are my pods?
notes:
- type: text
  contents: The website is down
tabs:
- title: Workstation
  type: terminal
  hostname: cloud-client
difficulty: basic
timelimit: 3600
---
# [App Installer Admin Console](http://loadbalancer.[[ Instruqt-Var key="SANDBOX_ID" hostname="cloud-client" ]].instruqt.io:8800)

password: password

🚀 Let's start
================

Let's imagine that our environment belongs to a customer, who are now experiencing an issue with their install.

They've raised a rather vague issue to your support team suggesting that the application "doesn't work" after one of their admins was playing with the setup.

Use the `sbctl` tool to inspect a support bundle they've shared with you, and try to determine what's amiss.

When you've identified the problem, type out the commmand you would use to resolve the problem in `/root/solution.txt`

The answer should be one line, on the first line of the file.

(The file does not exist, you will have to create it with your preferred text editor.)

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

✔️  Solution
==================

A deployment has been scaled to 0

🛠️ Remediation
=================

```bash
kubectl scale deployment <deployment-name> --replicas=1
```

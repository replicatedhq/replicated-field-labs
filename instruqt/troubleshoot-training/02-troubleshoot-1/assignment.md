---
slug: troubleshoot-1
id: araxpgiqal1r
type: challenge
title: Where are my pods?
teaser: "\U0001F914"
notes:
- type: text
  contents: The website is down
tabs:
- id: veuyak3z63de
  title: Workstation
  type: terminal
  hostname: cloud-client
difficulty: basic
timelimit: 3600
enhanced_loading: null
---
Let's imagine that our environment belongs to a customer, who are now experiencing an issue with their install.

They've raised a rather unclear issue to your support team suggesting that the application "doesn't work" after one of their users accientally made a change from the command line.

They've shared a support bundle with you, and you've been asked to help investigate.

Let's use the `sbctl` tool to inspect the support bundle and try to determine what's amiss.  `sbctl` should already be installed and the customer's support bundle should be in your home folder.  `sbctl` simulates having access to the customer's environment, but all of the data is taken from the support bundle.  It lets us use the familiar `kubectl` tool to explore the customer's environment, even without direct access.

When you've identified the problem, write out the commmand you would use to resolve the problem into a file at `/root/solution.txt`

The answer should be one line, on the first line of the file.

(The file does not exist, you will have to create it with your preferred text editor.)

💡 Using `sbctl`
=================

- Try `sbctl help` to see what commands are available

💡 Hints
=================

- Try the interactive shell prompt using `sbctl` and make sure to provide the path to the support bundle in your home folder

- How are applications deployed in kubernetes?

- What controls a pod's lifecycle?

💡 More Hints
=================

- How do I see deployments?

Troubleshooting Procedure
=================

Identify the problematic deployment from `kubectl get deployments -n <namespace>`.  Notice any pods that have 0 replicas, but should have 1 or more.

✔️  Solution
==================

A deployment has been scaled to 0

🛠️ Remediation
=================

```bash
kubectl scale deployment <deployment-name> --replicas=1
```

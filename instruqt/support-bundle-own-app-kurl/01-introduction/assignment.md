---
slug: introduction
id: tbq9chthjagw
type: challenge
title: Introduction
teaser: Practical Application of Support Bundles and Analyzers
notes:
- type: text
  contents: In this track, we'll work together to apply some practical methods for
    troubleshooting problems on Kubernetes clusters, using your own application.
tabs:
- title: Workstation
  type: terminal
  hostname: cloud-client
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
timelimit: 600
---

👋 Introduction
===============

* **What you will do**:
  * Learn to troubleshoot application & cluster problems using Support Bundles and `sbctl`
* **Who this is for**:
  * This track is for anyone who will build KOTS applications **plus** anyone who will be user-facing:
    * Full Stack / DevOps / Product Engineers
    * Support Engineers
    * Implementation / Field Engineers
    * Success / Sales Engineers
* **Prerequisites**:
  * Basic working knowledge of Linux (Bash)
  * An Embedded Cluster release of your application is available in Replicated Vendor Portal
  * A trial or dev license for your application so you can install it for yourself
* **Outcomes**:
  * You will be able to determine if the problem is in your application, in Kubernetes, or in the infrastructure environment
  * You will be ready to use Replicated's Support Bundle features to diagnose first-line issues in end-user environments
  * You will reduce escalations and expedite time to remediate for such issues

The environment is prepped for an *embedded cluster* installation.

Before we begin, select the text editor you're most comfortable with by running

```bash
update-alternatives --config editor
```

Press **Next** when you're ready to begin.

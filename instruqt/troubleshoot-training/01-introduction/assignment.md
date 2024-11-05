---
slug: introduction
id: b5dftki3524w
type: challenge
title: Introduction
teaser: Practical Application of Support Bundles and Analyzers
notes:
- type: text
  contents: In this track, we'll work together to apply some practical methods for
    troubleshooting some Kubernetes problems using Replicated tooling.
tabs:
- id: vy80bd4zzksw
  title: Workstation
  type: terminal
  hostname: cloud-client
difficulty: intermediate
timelimit: 600
enhanced_loading: null
---

👋 Introduction
===============

* **What you will do**:
  * Learn to troubleshoot application & cluster problems
* **Who this is for**:
  * This track is for anyone who will build KOTS applications **plus** anyone user-facing who support these applications:
    * Full Stack / DevOps / Product Engineers
    * Support Engineers
    * Implementation / Field Engineers
    * Success / Sales Engineers
* **Prerequisites**:
  * Basic working knowledge of Linux and the `bash` shell
* **Outcomes**:
  * You will be able to determine if the problem is in your application, in Kubernetes, or in the infrastructure environment
  * You will reduce escalations and expedite time to remediation for such issues

# Configure the VM environment

## Set up the Workstation

The environment is prepped for an *embedded cluster* installation.

### Configure your editor

Before we begin, let's choose an editor.  The default editor is `nano`, but if you'd like to use `vim` instead, you can switch to it by running the following command and selecting option `2`:

```bash
update-alternatives --config editor
```

Press **Check** when you're ready to begin.

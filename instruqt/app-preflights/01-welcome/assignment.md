---
slug: welcome
id: kklmmc12yvx8
type: challenge
title: Welcome
teaser: Welcome, track overview and check kubernetes cluster running
notes:
- type: text
  contents: |-
    This track uses a single node Kubernetes cluster on a sandbox virtual machine.
    Please wait while we boot the VM for you and start Kubernetes.

    ## Objectives

    In this track, this is what you'll cover:
    - Configure various Application Preflight Checks
    - View preflight checks in Kots Admin UI
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 600
---

👋 Introduction
===============

* **What you will do**:
    * This exercise will guide you through the core capabilities of the Replicated application preflight capabilities
    * Configure various Application Preflight tests and view and view resulant output on application deployment
* **Who this is for**: This track is for anyone who will build KOTS applications **plus** anyone who will be user-facing
    * Full Stack / DevOps / Product Engineers
    * Support Engineers
    * Implementation / Field Engineers
    * Success / Sales Engineers
* **Prerequisites**:
    * Basic working knowledge of Linux (Bash)
* **Outcomes**:
    * You will be ready to use KOTS's support bundle feature to diagnose first-line issues in end-user environments
    * You will reduce escalations and expedite time to remediate for such issues


## Get Started

Use the terminal to check that kubernetes is running:

```
kubectl get nodes
```

To complete this challenge, press **Check**.

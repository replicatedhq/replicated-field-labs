---
slug: intro
id: x3nb9trf3eet
type: challenge
title: Introduction
teaser: Introduction to the Replicated CLI Challenge
notes:
- type: text
  contents: |-
    This track uses a single node Kubernetes cluster on a sandbox virtual machine.
    Please wait while we boot the VM for you and start Kubernetes.

    ## Objectives

    **This track demonstrates how to:**
    - Validate and release your application from the command line
    - Find your download URL without leaving your shell
    - Install your application to an existing Kubernetes cluster
    - Put it together to iterate on your work
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 600
---

👋 Introduction
===============

This exercise is designed to give you a sandbox to ensure you have a basic understanding how to work with the Replicated CLI.

* **What you will do**:
  * Use the Replicated CLI to vaidate and release your application
  * Find your installation command and install to an existing cluster
  * Iterate on your work and deploy an updated release
* **Who this is for**: This lab is for anyone who works with app code, docker images, k8s yamls, or does field support for multi-prem applications
  * Full Stack / DevOps / Product Engineers
  * Support Engineers
  * Implementation / Field Engineers
* **Prerequisites**:
  * Basic working knowledge of Kubernetes
* **Outcomes**:
  * You will build a working understanding of how to deploy a kubernetes application packaged with Replicated using kots.

* * *

## Get started
Use the terminal to check if kubernetes is running:

```
kubectl get nodes
```

🏁 Finish
=========

To complete this challenge, press **Check**.

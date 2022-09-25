---
slug: intro
id: fcmnc0sptag8
type: challenge
title: Intro
teaser: Introduction to deploy a hello world app
notes:
- type: text
  contents: |-
    This track uses a single node Kubernetes cluster on a sandbox virtual machine.
    Please wait while we boot the VM for you and start Kubernetes.

    ## Objectives

    **This track demonstrates how to:**
    - Install Replicated's application installer
    - Deploy a webserver (NGINX) on Kubernetes using Replicated
    - Make a change in the webserver configuration and update the application
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 300
---

👋 Introduction
===============

This exercise is designed to give you a sandbox to ensure you have a basic understanding how to deploy a kubernetes application using Replicated.

The README and the YAML sources draw from https://github.com/replicatedhq/replicated-starter-kots

* **What you will do**:
  * Complete the simplest possible "Hello World" setup with a minimal KOTS application designed for demos
* **Who this is for**: This lab is for anyone who works with app code, docker images, k8s yamls, or does field support for multi-prem applications
  * Full Stack / DevOps / Product Engineers
  * Support Engineers
  * Implementation / Field Engineers
* **Prerequisites**:
  * Basic working knowledge of Kubernetes
* **Outcomes**:
  * You will build a working understanding of how to deploy a kubernetes application packaged with Replicated using kots.

* * *

🐚 Get started
===============

Use the terminal to check if kubernetes is running:

```
kubectl get nodes
```

🏁 Finish
=========

To complete this challenge, press **Check**.

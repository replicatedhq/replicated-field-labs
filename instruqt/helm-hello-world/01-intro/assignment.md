---
slug: intro
id: z273ux2hwcu2
type: challenge
title: Intro
teaser: Introduction to Deploying a Helm Based Kubernetes Application
notes:
- type: text
  contents: |-
    This track uses a single node Kubernetes cluster on a sandbox virtual machine.
    Please wait while we boot the VM for you and start Kubernetes.

    ## Objectives

    **This track demonstrates how to:**
    - Install Replicated's application installer
    - Deploy Wordpress on Kubernetes using Replicated
    - Provide the Initial Blog's name during the installation
    - Open Wordpress from the Admin Console
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 300
---
👋 Introduction
===============

This exercise is designed to give you a sandbox to ensure you have a basic understanding how to deploy a kubernetes application packaged as a Helm Chart using Replicated.

The README and the YAML sources draw from https://github.com/replicatedhq/replicated-starter-kots

* **What you will do**:
  * Review the Wordpress Helm Chart packaged in Replicated and deploy it to a sandbox environment
* **Who this is for**: This lab is for anyone who works with app code, docker images, k8s yamls, or does field support for multi-prem applications
  * Full Stack / DevOps / Product Engineers
  * Support Engineers
  * Implementation / Field Engineers
* **Prerequisites**:
  * Basic working knowledge of Kubernetes
* **Outcomes**:
  * You will build a working understanding of how to deploy a kubernetes application packaged as a Helm Chart using Replicated.

* * *

## Get started
Use the terminal to check if kubernetes is running:

```
kubectl get nodes
```

To complete this challenge, press **Check**.

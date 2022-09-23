---
slug: intro
id: xml6e2vefn99
type: challenge
title: Intro
teaser: Introduction to using Support Bundles and Analyzers
notes:
- type: text
  contents: In this track, we'll use the Support Bundle analyzers feature to debug
    an application, modifying the host in order to create the correct conditions for
    the application to start.
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 300
---

👋 Introduction
===============

* **What you will do**:
    * Learn to query, read, and understand support bundle analyzers
    * Use the analyzers to fix a problem on the server and get the application up and running
* **Who this is for**: This lab is for anyone who will build KOTS applications **plus** anyone who will be user-facing
    * Full Stack / DevOps / Product Engineers
    * Support Engineers
    * Implementation / Field Engineers
    * Success / Sales Engineers
* **Prerequisites**:
    * Basic working knowledge of Linux (SSH, Bash)
* **Outcomes**:
    * You will be ready to use KOTS's support bundle feature to diagnose first-line issues in end-user environments
    * You will reduce escalations and expedite time to remediate for such issues

🐚 Get started
===============

Use the terminal to check if kubernetes is running:

```
kubectl get nodes
```
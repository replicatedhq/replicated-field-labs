---
slug: intro
type: challenge
title: intro
teaser: Create your own collectors and analyzers
notes:
- type: text
  contents: In this track we will learn how to capture the logs of our application
    and add automated analyzers to them.
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
    * Learn to download and manage support bundles when analyzers cannot immediately resolve the issue
    * Learn to write new analyzers to detect more complex issues as they arise
* **Who this is for**: This lab is for anyone who will own continuously evolving the reliability and supportability of a KOTS application.
    * Full Stack / DevOps / Product Engineers
    * Support Engineers
    * Implementation / Field Engineers
* **Prerequisites**:
    * You should be familiar with deploying applications using Replicated's Application Installer and use the Replicated CLI.
    * Basic working knowledge of Kubernetes
* **Outcomes**:
    * You will be ready to use support bundles to collaborate with your team when escalating issues from the field
    * You will be confident in the process for adding new analyzers, evolving your support tooling over time to continuously
      reduce unnecessary escalations
    * You will understand how to escalate issues to the Replicated support team  


🐚 Get started
===============

Use the terminal to check if kubernetes is running:

```
kubectl get nodes
```

To complete this challenge, press **Check**.
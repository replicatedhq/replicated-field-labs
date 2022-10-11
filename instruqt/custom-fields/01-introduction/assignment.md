---
slug: introduction
type: challenge
title: Introduction
teaser: Learn what this lab is about!
notes:
- type: text
  contents: Get ready to learn about Custom License Fields
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 300
---

👋 Introduction
===============

This exercise is designed to give you a sandbox to ensure you have a basic understanding how Custom License Fields work.

* **What you will do**:
  * Learn how to create a Custom License Field
  * Learn how to use the field in a sample application
  * Deploy the sample app and then update the field to see the change
* **Who this is for**: This lab is for anyone who works with app code, docker images, k8s yamls, or does field support for multi-prem applications
  * Full Stack / DevOps / Product Engineers
  * Support Engineers
  * Implementation / Field Engineers
* **Prerequisites**:
  * Hello World Lab
  * Replicated CLI lab
  * Basic working knowledge of Kubernetes
* **Outcomes**:
  * You will build a working understanding of how you can leverage Custom License Fields in your application.


🐚 Get started
===============

Use the terminal to check if kubernetes is running:

```
kubectl get nodes
```

You should see a Kubernetes node listed. If you do, **Congratulations!** you are ready to move to the next challenge.

To complete this challenge, press **Next**.
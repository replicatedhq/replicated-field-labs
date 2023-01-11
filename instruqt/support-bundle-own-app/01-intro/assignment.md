---
slug: supportbundle-practical-1
type: challenge
title: Support Bundles in Practice 1
teaser: Practical Application of Support Bundles and Analyzers
notes:
- type: text
  contents: In this track, we'll work together with your own application to practice troubleshooting Kubernetes applications using Support Bundles
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: intermediate
timelimit: 300
---

👋 Introduction
===============

* **What you will do**:
    * Learn to troubleshoot application & cluster problems using Support Bundles and `sbctl`
* **Who this is for**: This track is for anyone who will build KOTS applications **plus** anyone who will be user-facing
    * Full Stack / DevOps / Product Engineers
    * Support Engineers
    * Implementation / Field Engineers
    * Success / Sales Engineers
* **Prerequisites**:
    * Basic working knowledge of Linux (Bash)
    * A release of your application is available in Replicated Vendor Portal
    * A trial or dev license for your application so you can install it for yourself
* **Outcomes**:
    * You will be able to determine if the problem is in your application, in Kubernetes, or in the infrastructure environment
    * You will be ready to use KOTS's support bundle feature to diagnose first-line issues in end-user environments
    * You will reduce escalations and expedite time to remediate for such issues

🐚 Get started
===============

Use the terminal to check if kubernetes is running:

```
kubectl get nodes
```

To complete this challenge, press **Check**.

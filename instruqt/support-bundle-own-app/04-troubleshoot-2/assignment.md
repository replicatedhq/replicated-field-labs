---
slug: troubleshoot-2
id: fvqrxka6fxu8
type: challenge
title: Correcting the broken application
teaser: Time to fix another problem
notes:
- type: text
  contents: Time to fix another problem...
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
- title: Application Installer
  type: website
  url: http://kubernetes-vm.${_SANDBOX_ID}.instruqt.io:8800
  new_window: true
difficulty: basic
timelimit: 3600
---

🚀 Let's start
=================

You get another report from a customer saying that the application isn't working, as if two components are unable to communicate.  How would you begin to solve the problem?

- You may want to start by verifying that all the expected pods are running, then move on to checking the communication between the pods and the client.

💡 Hints
=================

- The Kubernetes documentation has a [great manual on debugging Services](https://kubernetes.io/docs/tasks/debug/debug-application/debug-service/)

- Think about the traffic flow to your application
  - There are multiple *hops* in the network path, and any of them _could_ be a potential break in the path.  - Which hops can you identify?

- How does traffic get to workloads inside kubernetes
- How does Kubernetes handle DNS resolution and load balancing for Pods?

✔️ Solution
=================

A random Service's `targetPort` has been patched to be something in the 30k range.  Any pod in the cluster that tries to connect to this pod's Service name (which is what gets programmed in to DNS in the cluster) will fail to connect, because the pod is not listening on the the same port that the Service is trying to connect to.

Remediation
=================

Patch or edit the affected service to correct the port number. you may have to refer to the other resources in the cluster to identify the correct port number.

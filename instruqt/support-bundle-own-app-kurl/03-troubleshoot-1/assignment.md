---
slug: troubleshoot-1
id: sgroektituzf
type: challenge
title: Correcting the broken application
teaser: A Pod is crashing...
notes:
- type: text
  contents: Time to fix the problem...
tabs:
- title: Workstation
  type: terminal
  hostname: cloud-client
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: Cluster Node 1
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-1
- title: Cluster Node 2
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-2
- title: Cluster Node 3
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-3
difficulty: intermediate
timelimit: 3600
---
[App Installer Admin Console](http://loadbalancer.[[ Instruqt-Var key="SANDBOX_ID" hostname="cloud-client" ]].instruqt.io:8800)

🚀 Let's start
=================

Now we will explore solving an application problem in *[[ Instruqt-Var key="APP_SLUG" hostname="kubernetes-vm" ]]/[[ Instruqt-Var key="CHANNEL" hostname="kubernetes-vm" ]]*.  Imagine: you are supporting a customer and they report to you that one of their application pods is crashing.  How do you begin to solve the problem?

We think it's a good habit to use Support Bundles whenever possible - and we'll use another tool we've developed called `sbctl` to help with that.

First let's take a support bundle from the cluster:

```
kubectl support-bundle --load-cluster-specs --output support-bundle.tar.gz
```

Right now we have access to the cluster, but what if we did not?  `sbctl` to the rescue - let's invoke it:

```
sbctl shell --support-bundle-location ./support-bundle.tar.gz
```

This will give us a shell in the context of the support bundle.  `sbctl` mocks a Kubernetes API so we can use `kubectl` to inspect the data in the bundle - just like a regular cluster.


💡 Hints
=================
- How do you list pods?

- How do you describe pods?
  - What if you wanted to see events from multiple pods at once?

- How do you get logs from a pod?
  - What if you wanted to see a previous version of the pod's logs?

- When would you look at `describe` output vs. gathering pod logs?

- Review the [Kubernetes documentation on debugging Pods](https://kubernetes.io/docs/tasks/debug/debug-application/debug-running-pod/)

💡 More Hints
=================
- How do you find the exit code of a Pod?

- What could it mean if a Pod is exiting before it has a chance to emit any logs?

✔️ Solution
=================
A random deployment has been selected and the memory limit reduced to 10Mi.  This will cause the application to crash.

Remediation
=================
Patch or edit the affected deployment to increase the memory request and limit to a reasonable amount.

- How can you edit or patch a resource in-place?

- How can you edit or patch a resource from a file?

- How can we make sure that this doesn't happen again?

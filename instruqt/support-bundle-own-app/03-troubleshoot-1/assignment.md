---
slug: troubleshoot-1
id: wmm7b9ophrjc
type: challenge
title: Correcting the broken application
teaser: Time to fix the problem
notes:
- type: text
  contents: Time to fix the problem...
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
Our chaos script randomly selected a Deployment in your application and reduced the memory limit to 10Mi, which should cause those pods to start crashing.

Remediation
=================
Patch or edit the affected deployment to increase the memory request and limit to a reasonable amount.

- How can you edit or patch a resource in-place?

- How can you edit or patch a resource from a file?

- How can we make sure that this doesn't happen again?
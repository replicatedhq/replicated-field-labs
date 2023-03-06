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

Now we will explore solving an application problem in *[[ Instruqt-Var key="APP_SLUG" hostname="cloud-client" ]]/[[ Instruqt-Var key="CHANNEL" hostname="cloud-client" ]]*.  Imagine: you are supporting a customer and they report to you that one of their application pods is crashing.  How do you begin to solve the problem?


💡 Hints
=================
- How do you list pods?

- How do you describe pods?
  - What if you wanted to see data from multiple pods at once?

- How do you get logs from a pod?
  - What if you wanted to see a previous version of the pod's logs?

- When would you look at `describe` output vs. gathering pod logs?

- Review the [Kubernetes documentation on debugging Pods](https://kubernetes.io/docs/tasks/debug/debug-application/debug-running-pod/)

💡 More Hints
=================
- How do you find the exit code of a Pod?

- What could it mean if a Pod is exiting before it has a chance to emit any logs?

Troubleshooting Procedure
=================

Identify the problematic Pod from `kubectl get pods -n <namespace>`.  Notice any pods that are not in the Running state.

Describe the current state of the Pod with `kubectl describe pod -n <namespace> <pod-name>`.  Here are some things to look out for:
  - each Container's current State and Reason
  - each Container's Last State and Reason
    - the Last State's Exit Code
  - each Container's Ready status
  - the Events table

For a Pod that is crashing, expect that the current state will be `Waiting` or `Terminated`, and the last state will probably also be `Terminated`.  Notice the reason for the termination, and especially notice the exit code.  There are standards for the exit code originally set by the `chroot` standards, but they are not strictly enforced since applications can always set their own exit codes.

In short, if the exit code is >128, then the application exited as a result of Kubernetes killing the Pod.  If that's the case, you'll commonly see code 137 or 143, which is 128 + the value of the `kill` signal sent to the container.

If the exit code is <128, then the application crashed or exited abnormally.  If the exit code is 0, then the application exited normally (most commonly seen in init containers or Jobs/CronJobs)

Look for any Events that may indicate a problem.  Look for the Pod's State,

First, let's

✔️ Solution
=================
A random deployment has been selected and the memory limit reduced to 10Mi.  This will cause the application to crash.

Remediation
=================
Patch or edit the affected deployment to increase the memory request and limit to a reasonable amount.

- How can you edit or patch a resource in-place?

- How can you edit or patch a resource from a file?

- How can we make sure that this doesn't happen again?

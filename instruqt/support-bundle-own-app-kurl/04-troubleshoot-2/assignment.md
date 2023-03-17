---
slug: troubleshoot-2
id: bbn7etmrxhet
type: challenge
title: Correcting the broken application
teaser: A Pod is crashing...
notes:
- type: text
  contents: Time to fix another problem...
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
difficulty: basic
timelimit: 3600
---
[App Installer Admin Console](http://loadbalancer.[[ Instruqt-Var key="SANDBOX_ID" hostname="cloud-client" ]].instruqt.io:8800)

🚀 Let's start
================
The customer opens another issue about pods, but this time pods seem to be crashing.

Let's investigate our app and see if we can identify the issue.


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

For a Pod that is crashing, expect that the current state will be `Waiting`, `Terminated` or `Error`, and the last state will probably also be `Terminated`.  Notice the reason for the termination, and especially notice the exit code.  There are standards for the exit code originally set by the `chroot` standards, but they are not strictly enforced since applications can always set their own exit codes.

In short, if the exit code is >128, then the application exited as a result of Kubernetes killing the Pod.  If that's the case, you'll commonly see code 137 or 143, which is 128 + the value of the `kill` signal sent to the container.

If the exit code is <128, then the application crashed or exited abnormally.  If the exit code is 0, then the application exited normally (most commonly seen in init containers or Jobs/CronJobs)

Look for any Events that may indicate a problem.  Events by default last 1 hour, unless they occur repeatedly.  Events in a repetition loop are especially noteworthy:

```
Events:
  Type     Reason                  Age                      From     Message
  ----     ------                  ----                     ----     -------
  Warning  BackOff                 2d19h (x9075 over 4d4h)  kubelet  Back-off restarting failed container sentry-workers in pod sentry-worker-696456b57c-twpj7_default(82eb1dde-2987-4f58-af64-883470ffcb58)
```


✔️ Solution
=================
A random deployment has been selected and the memory limit reduced to 10Mi.  This will cause the application to crash.

Remediation
=================
Patch or edit the affected deployment to increase the memory request and limit to a reasonable amount.

- How can you edit or patch a resource in-place?

- How can you edit or patch a resource from a file?

- How can we make sure that this doesn't happen again?

---
slug: troubleshoot-4
id: vccuaq9got6x
type: challenge
title: it's all made of stars
teaser: The final frontier...
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
## [App Installer Admin Console](http://loadbalancer.[[ Instruqt-Var key="SANDBOX_ID" hostname="cloud-client" ]].instruqt.io:8800)

🚀 Let's start
=================

You get a new report from a customer saying that many pods are failing; some may display Errors, while others may be Evicted or Pending, or even in an Unknown state.

How do you begin to debug this problem?

💡 Hints
=================

- Remember from our first challenge how to examine the state of a Pod in the cluster.

- What patterns can be observed from the state of the Pods that are failing?
  - Are they all in the same namespace?
  - Are they all using the same container image?
  - Are they all scheduled to the same Node?

💡 More Hints
=================

- Get more information from `get pods` by using the `-o wide` option.

- Remember that Nodes are also a resource managed by Kubernetes - how do you view the state of a Node?

- Nodes have Events and States/Conditions just like other resources.

- Make sure to examine the underlying Linux filesystem
  - `lsblk`
  - `df -h`

✔️ Solution
=================

One of the disks in the cluster is full, which causes problems in the operating system as well as with Kubernetes.  Pod eviction thresholds have been exceeded, causing pods to be evicted and image garbage collection to remove images from the node.

Remediation
=================

Find the offending file and remove it, and the cluster should recover.  Command line tools like `find` and `du` can be helpful here.

Wait a few minutes for the Kubelet to recognize that disk space is available, and then clean up any Errored or Evicted pods:

```
kubectl delete pods --field-selector status.phase=Failed -A
```

We also have a Troubleshoot spec that looks for particularly large files at https://github.com/replicatedhq/troubleshoot-specs/blob/main/host/resource-contention.yaml#L36

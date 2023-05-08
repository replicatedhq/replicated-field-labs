---
slug: troubleshoot-7
id: nuxfa0vjx8yc
type: challenge
title: Correcting the broken application
teaser: Rook-Ceph is unhealthy...
notes:
- type: text
  contents: Replace this text with your own text
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
timelimit: 600
---
[App Installer Admin Console](http://loadbalancer.[[ Instruqt-Var key="SANDBOX_ID" hostname="cloud-client" ]].instruqt.io:8800)

🚀 Let's start
=================

Time for a new challenge! Now, a customer has reported that they've recently upgraded their kURL cluster, and afterward Rook-Ceph is unhealthy. Let's see if we can figure out what's going on.

💡 Hints
=================

- Ceph status is available from `ceph status` inside the rook-ceph-tools pod.
  - `kubectl -n rook-ceph exec -it rook-ceph-tools-xxxxx-yyyyyy -- ceph status`

- Generally, Ceph expects a `mon` deployment to be running on each node in the cluster.

- The `rook-ceph-operator` is responsible for reconciling the Ceph cluster.

💡 More Hints
=================

- Remember that Deployments manage ReplicaSets, which in turn manage Pods.

✔️ Solution
=================

One of the Ceph monitor (`mon`) deployments has a PriorityClass that is not available on the cluster. This causes the deployment to be stuck in a `Pending` state.

Remediation
=================

Edit the `rook-ceph-mon` deployment and edit the `priorityClassName` field to `system-cluster-critical`.

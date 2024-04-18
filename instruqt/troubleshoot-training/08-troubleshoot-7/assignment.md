---
slug: troubleshoot-7
id: gfetpub34n8i
type: challenge
title: Storage
teaser: "\U0001F9FA"
notes:
- type: text
  contents: Replace this text with your own text
tabs:
- title: Workstation
  type: terminal
  hostname: cloud-client
difficulty: basic
timelimit: 600
---

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

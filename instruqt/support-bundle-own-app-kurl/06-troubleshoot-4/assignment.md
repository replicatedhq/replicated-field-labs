---
slug: troubleshoot-4
id: ksonxafcrrhe
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
difficulty: advanced
timelimit: 3600
---
[App Installer Admin Console](http://loadbalancer.[[ Instruqt-Var key="SANDBOX_ID" hostname="cloud-client" ]].instruqt.io:8800)

🚀 Let's start
=================

You get a new report from a customer saying Rook Ceph is reporting an Unhealthy status.  How do you begin to troubleshoot a problem with Ceph?

💡 Hints
=================

- Use `kubectl get pods -n rook-ceph` to see the status of the Rook Ceph pods

- Make a "toolbox" Pod to run `ceph` commands
    - `kubectl --namespace rook-ceph create -f https://raw.githubusercontent.com/rook/rook/master/cluster/examples/kubernetes/ceph/toolbox.yaml`
    - `kubectl --namespace rook-ceph exec -it deploy/rook-ceph-tools -- bash`

- Review the [Rook Ceph troubleshooting documentation](https://rook.io/docs/rook/v1.11/Troubleshooting/common-issues/)

- Note some common commands for troubleshooting a Ceph cluster
  - `ceph status`
  - `ceph osd status`
  - `ceph osd df`
  - `ceph osd utilization`
  - `ceph osd pool stats`
  - `ceph osd tree`
  - `ceph pg stat`

- Make sure to examine the underlying Linux filesystem and block devices being used by Ceph
  - `lsblk`
  - `df -h`

✔️ Solution
=================

Remediation
=================

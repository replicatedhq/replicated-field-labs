---
slug: troubleshoot-4
id: ksonxafcrrhe
type: challenge
title: Correcting the broken application
teaser: Rook is unhealthy...
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

rook/ceph problem - `dd` into a rook block device?

💡 Hints
=================

✔️ Solution
=================

Remediation
=================

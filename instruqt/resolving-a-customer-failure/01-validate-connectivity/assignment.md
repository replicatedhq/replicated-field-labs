---
slug: validate-connectivity
id: yggrsxmjgfqy
type: challenge
title: validate-connectivity
teaser: Validates connectivity between the shell and the cluster
notes:
- type: text
  contents: Let's check that we can connect
tabs:
- title: Shell
  type: terminal
  hostname: shell
  workdir: /home/replicant
- title: Cluster
  type: terminal
  hostname: cluster
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 3000
---

#### Let's check our connectivity

```
kubectl get nodes
```

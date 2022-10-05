---
slug: validate-connectivity
id: zwgjxayvrlqk
type: challenge
title: validate-connectivity
teaser: Validates connectivity between the jumpbox and the airgap
notes:
- type: text
  contents: Let's check that we can connect
tabs:
- title: Jumpbox
  type: terminal
  hostname: jumpbox
  workdir: /home/replicant
- title: Cluster
  type: terminal
  hostname: airgap
difficulty: basic
timelimit: 3000
---

#### Let's check our connectivity

```
kubectl get nodes
```

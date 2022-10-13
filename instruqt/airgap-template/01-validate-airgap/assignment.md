---
slug: validate-airgap
id: zwgjxayvrlqk
type: challenge
title: validate-airgap
teaser: Validates the cluster is air-gapped and we can connect over SSH
notes:
- type: text
  contents: Let's check that we can't connect
tabs:
- title: Jumpbox
  type: terminal
  hostname: jumpbox
  workdir: /home/replicant
difficulty: basic
timelimit: 300
---

#### Let's check our air gap

```
ssh cluster curl --connect-timeout 30 https://google.com
```

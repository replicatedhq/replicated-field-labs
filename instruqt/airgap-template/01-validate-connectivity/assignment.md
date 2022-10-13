---
slug: validate-connectivity
id: zwgjxayvrlqk
type: challenge
title: validate-connectivity
teaser: Validates connectivity between the jumpbox and the airgap
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
ssh curl https://google.com
```

---
slug: connecting
id: x5fau7sif8sa
type: challenge
title: Connecting
teaser: Connecting with the airgapped server
notes:
- type: text
  contents: Connecting with the airgapped server
tabs:
- title: Jumpstation
  type: terminal
  hostname: jumpstation
difficulty: basic
timelimit: 600
---

### Connecting

From the Jumpstation terminal, you can SSH into the Air Gap server using the following command:

```shell
ssh airgap
```

Accept the fingerprint generation by accepting the defaults.

Once you're on the Air Gap server, you can verify that the server indeed does not have internet access. Once you're convinced, you
can ctrl+C the command and proceed to the next section

```shell
curl https://kubernetes.io
```

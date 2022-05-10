---
slug: connecting
id: 8bczmdyolixx
type: challenge
title: Connecting
teaser: Connecting with the airgapped server
notes:
- type: text
  contents: Replace this text with your own text
tabs:
- title: CLI
  type: terminal
  hostname: cli
- title: Jumpstation
  type: terminal
  hostname: jumpstation
difficulty: basic
timelimit: 600
---

### Connecting

First set your application slug, the public IP of your jump box and your first name:

```shell
export JUMP_BOX_IP=...
export REPLICATED_APP=... # your app slug
export FIRST_NAME=... # your first name
```

Next, you can SSH into the Air Gap server using the following command:

```shell
ssh -J ${FIRST_NAME}@${JUMP_BOX_IP} ${FIRST_NAME}@${REPLICATED_APP}-lab05-airgap
```

The `-J` option, allows to connect to the target host by first making a ssh connection to the jump host (`${JUMP_BOX_IP}`) described by destination and then establishing a TCP forwarding to the ultimate destination (`${REPLICATED_APP}-lab05-airgap`) from there.

You can also do it in multiple steps and achieve the same:

```shell
local> ssh ${FIRST_NAME}@${JUMP_BOX_IP}
jump> export REPLICATED_APP=...
jump> ssh ${REPLICATED_APP}-lab05-airgap
```

Once you're on the Air Gap server, you can verify that the server indeed does not have internet access. Once you're convinced, you
can ctrl+C the command and proceed to the next section

```shell
curl -v https://kubernetes.io
```

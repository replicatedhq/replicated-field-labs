---
slug: speciying-the-cluster
id: bvpn5mxammoz
type: challenge
title: Configuring the Embedded Cluster
teaser: Enable and configure the embedded cluster
notes:
- type: text
  contents: Let's enable the embedded cluster
tabs:
- title: Shell
  type: terminal
  hostname: shell
- title: Release Editor
  type: code
  hostname: shell
  path: /home/replicant
difficulty: basic
timelimit: 300
---

You can release you application as a Kubernetes applicance by
specifying a few extra configuration files as part of your Replicated
release. The first of these files is the configuration for the
Embedded Cluster itself. This configuration can be very simple. All
it's required to provide is the version of the cluster to use.

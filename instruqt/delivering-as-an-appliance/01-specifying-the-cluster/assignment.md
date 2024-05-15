---
slug: specifying-the-cluster
id: rsxl9h79cen1
type: challenge
title: Configuring the Embedded Cluster
teaser: Enable and configure the embedded cluster
notes:
- type: text
  contents: Let's enable the embedded cluster
tabs:
- title: Shell
  type: terminal
  hostname: node
- title: Release Editor
  type: code
  hostname: node
  path: /home/replicant
difficulty: basic
timelimit: 300
---

A virtual Kubernetes appliance consists of your application, an embedded
Kubernetes cluster, and a console that you customer uses to install your
application and manage the appliance. You release your application as a
Kubernetes applicance by specifying a few configuration files. These files are
included as part of your Replicated release.

The first file is the configuration for the Embedded Cluster itself, which can
very simple. All it requires is the version of the cluster to use.

```
apiVersion: embeddedcluster.replicated.com/v1beta1
kind: Config
spec:
  version: [[ Instruqt-Var key="EMBEDDED_CLUSTER_VERSION" hostname="node" ]]
```

The Replicated Release
======================

The Replicated Platform distributes your software to your customers. To do
this, it needs to know about your application, it's customers, and the files
you're shipping to them. Each release is built around a Helm chart, and that's
all it needs to include. If you've completed the [Distributing Your Software
with Replicated](https://play.instruqt.com/replicated/tracks/distributing-with-replicated)
lab, you built a release around a Helm chart and installed it using Helm tools.

For the Emnbe

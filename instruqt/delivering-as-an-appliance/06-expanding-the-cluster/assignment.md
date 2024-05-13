---
slug: expanding-the-cluster
id: eiyviocvnrbs
type: challenge
title: Expanding the Embedded Cluster Instance
teaser: |
  Adding nodes to a running cluster to provide sufficient capacity
notes:
- type: text
  contents: Moving from a single node cluster to multiple nodes
tabs:
- title: Shell
  type: terminal
  hostname: node
difficulty: basic
timelimit: 300
---

The initial install of a Kubernets appliance converts a single (often
virtual) machine into a one-node cluster running your application.
Your customer can expand the cluster by adding additional nodes
either as need or prior to their initial install. Expanding a cluster
is a "clickops" experience that any adminstrator will be able to
complete.

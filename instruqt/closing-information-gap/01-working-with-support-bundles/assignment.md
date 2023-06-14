---
slug: working-with-support-bundles
type: challenge
title: Working with Support Bundkes
teaser: |-
  Learn how support bundles help you understand what is going
  on with a customer instance
notes:
- type: text
  contents: Let's learn about Support Bundles
tabs:
- title: Shell
  type: terminal
  hostname: shell
- title: Manifest Editor
  type: code
  hostname: shell
  path: /home/replicant
difficulty: basic
timelimit: 300
---

👋 Introduction
===============

When your software is running in a customer cluster, you no longer
have direct access to troubleshoot when things go wrong. You won't
be able to see what's running, read logs, or even confirm that the
ever started up. Your cutsomer can do these things, but they may
need your guidance to do them correctly and coordinating that 
information sharing can be challenging.

The Replicated Platform let's you define a support bundles that
your customer can send to you to bring you the visbility you need.
Support bundles can also surface specific issues and provide 
guidance to your customer in order to repair issues on their own.
They are part of the [Troubleshoot](https://troubleshoot.sh) open 
source project.

What is a Support Bundle?
=========================

Support Bundles collect the information you need to understand their
cluster and how your application is running in it. The Replicated 
Platform allows you to do this without installing anything else
into the cluster

You define your bundle in a YAML file that follows the same format
as a Kubernetes object. The simplest support bundle object looks
like this, and it's in the file `simplest-support-bundle.yaml`:

```
apiVersion: troubleshoot.sh/v1beta2
kind: SupportBundle
metadata:
  name: simplest-support-bundle
spec:
  collectors: []
  analyzers: []
```

As the name `empty-` suggests, this is an empty set
of checks and will not execute. Let's try it out.

```
kubectl support-bundle ./simplest-support-bundle.yaml
```

It will take a few seconds to generate a support bundle in a
file named `support-bundle-$TIMESTAMP.tar.gz` that contains
some simplest information about the cluster. You can get a
flavor for what's in the bundle by running

```
tar -tzf support-bundle-*.tar.gz | less
```

You'll see the files that were collected cataloging all of the
resources in the cluster and some information about the cluster
itself.

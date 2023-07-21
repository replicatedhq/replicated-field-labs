---
slug: basics
id: fytbmwjxpo78
type: challenge
title: Compatibility Matrix Basics
teaser: Learn the basic commands for working with the Compatibility Matrix CLI.
notes:
- type: text
  contents: Let's explore the Compatibility Matrix
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---
Introduction
===============
In this exercise we will examine the different matrix commands and their uses as well as build and tear down a few environments for practice.

Getting Started
===
Let’s start with how to create an environment using the Compatibility Matrix. 

We'll start by looking at the new CLI commands that the `replicated` command offers:

<link to docs>


Using the replicated cluster commands
===

To see the list of available `replicated cluster` subcommands, run:

``` replicated cluster --help ```

The output provides a list of possible subcommands. We can learn more about any subcommand by using the `--help` flag. Let's learn a little more about `replicated cluster create`:

``` replicated cluster create --help ```

As you can see, the `create` subcommand takes flags to customize the cluster creation. 

Let's go ahead and find out what versions we have access to:

``` replicated cluster versions  ```

Now that we know what we can use, let's build a cluster. We'll start with a simple k3s 1.24 cluster:

 # 3. Create a cluster from scratch. Minimum flags: --version, --distribution. Connect to it, then delete it.

``` replicated cluster create --distribution k3s --version 1.24 ```

When the cluster is created, cluster information will be displayed. You can also see available clusters with the `ls` command:

``` replicated cluster ls ```

Watch this cluster build with `watch replicated cluster ls` as it goes through the build stages: assigned, preparing, provisioning, running.

We can create a cluster with some additional flags to further customize it for our needs. We're going to start with a cluster with a that only lasts ten minutes. Note that the flag for cluster duration accepts anywhere from `10m0s` through `48h0m0s`:

``` replicated cluster create --distribution k3s --version 1.24 --ttl 2s ```

Let's delete these clusters to keep our workspace clean. The `rm` subcommand can take multiple IDs, separated by spaces, so we will just need to grab multiple IDs from an `ls`

``` replicated cluster rm {{ID}} {{ID2}} ```


4. Create another cluster with a very short (2 minute) duration. Verify it exists
5. Create 3-node cluster with multiple OS.  (Keep this)
6. Install application in all three nodes.
7. Verify that the cluster from step 4 is gone, as the TTL has expired


To complete this track, click the **Check** button.


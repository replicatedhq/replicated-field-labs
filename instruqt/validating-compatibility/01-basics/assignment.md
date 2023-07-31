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
The Compatibility Matrix creates ephemeral lab-like environments which can be used for testing upgrades and deployments, troubleshooting customer issues, and exploring potential changes to systems. Because it is CLI-driven, it can be easily automated and integrated into existing (or planned) CI/CD processes without requiring a permanent lab environment. In this set of labs, we'll explore how to use the Replicated Compatibility Matrix and some of the ways in which it can be incorporated into existing pipelines. Let's dive in!

Getting Started
===
In this exercise we will examine the different matrix commands and their uses as well as build and tear down a few environments. The goal of this exercise is to familiarize us with the basics of the compatibility matrix and get us comfortable working with it.

First we'll learn how to create an environment using the Compatibility Matrix. Once we understand the basics of creating an environment, we can use that to automate building environments to meet our testing, troubleshooting, and deployment goals. Because the environments we'll be building are customizable to match our customers' specs, they can mimic the actual environments our customers are using, and because they're ephemeral, they won't require any maintenance. 

Let's start by looking at the new CLI commands that the `replicated` command offers:

<link to docs>

Using the replicated cluster commands
===

To see the list of available `replicated cluster` subcommands, run:

``` replicated cluster --help ```

The output provides a list of possible subcommands. The `--help` flag works here, too. Let's learn a little more about `replicated cluster create`:

``` replicated cluster create --help ```

The `create` subcommand takes flags to customize cluster creation, which allows us to tailor our clusters to match our cuatomer environments.

One of the customizable options is the type of cluster. Let's go ahead and find out what versions we have access to:

``` replicated cluster versions  ```

Now that we know what we can use and what a creation command looks like, let's build a cluster. We'll start with a simple k3s 1.24 cluster:

``` replicated cluster create --distribution k3s --version 1.24 ```

When the cluster is created, cluster information will be displayed. You can also see available clusters with the `ls` command:

``` replicated cluster ls ```

Watch this cluster build with `watch replicated cluster ls` as it goes through the build stages: assigned, preparing, provisioning, running. This is the simplest form of cluster we can build, but as `replicated cluster create --help` showed, there are a lot of options for customizing it to suit our needs.

Let's create a more complex cluster. Our BigBank customer is running a standard setup, with 3 nodes in a kind cluster with the following specs, so to match that environment we will do the same.

Specs:

Version: kind 1.27.0 
Memory: 8GB
vCPU: 8
Nodes: 3

This is an unusual config, so let's verify we've assembled the correct command before building. The`--dry-run` flag will allow us to verify that our specs are supported:

``` replicated cluster create --distribution kind --version 1.27.0 --name bigbank-testing --memory 8 --vcpu 8 --node-count
 3 --dry-run ```

As we can see, the dry run was a success so we can go ahead and build this new testing cluster:

``` replicated cluster create --distribution kind --version 1.27.0 --name bigbank-testing --memory 8 --vcpu 8 --node-count
 3 ```

As it builds, we can see both clusters are up and running/provisioning  with another `replicated cluster ls`.

We don't need the first cluster we created, since now we have a cluster that represents our customer accurately. Let's delete that cluster by using its ID:

``` replicated cluster rm {{cluster-ID}} ```

This cluster drops off our cluster list immediately, though in the background it is still cleaning up.

Now we have one running cluster with the same configurations as Big Bank. Let's go ahead and install our app into that cluster so we can begin using it to do testing and troubleshooting. 

First, we need to connect to our cluster. If we were to run a `kubectl get namespaces` right now, we wouldn't be interacting with the testing cluster. The `kubeconfig` subcommand allows us to interact with the cluster directly:

``` replicated cluster kubeconfig --name {{name}} ```

Now, when we run `kubectl get namespaces` we can see a list of the current namespaces. 

Let's go ahead and get our app installed to this cluster. This uses the same kots installer as a regular installation:

``` kubectl kots install {{app-slug}} ```

Now a `kubeadm get namespaces` shows that our app has installed in the namespace we chose during installation.

Congratulations! You've created your first few clusters using the Replicated compatibility matrix and installed an app to an ephemeral CM cluster! Next, let's look at how we can use a customer's support dump to troubleshoot their errors.

To complete this track, click the **Check** button.


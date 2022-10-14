---
slug: airgap-overview
id: mq22qgecjrh9
type: challenge
title: airgap-overview
teaser: An overview of the Replicated airgap overview
notes:
- type: text
  contents: Our workflow for air-gapped deployment
tabs:
- title: Jumpbox
  type: terminal
  hostname: jumpbox
  workdir: /home/replicant
difficulty: basic
timelimit: 300
---

👋 Introduction
===============

In this exercise you will learn how to perform installations in Air Gap environments, and
how to collect support bundles in Air Gap environments.

* **What you will do**:
    * Access and verify a single-node Air Gap setup via a bastion server
    * Learn to use KOTS to install in an Air Gap environment
    * Create an SSH tunnel to configure an Air Gap instance
    * Perform an upgrade of an application in an Air Gap environment
    * Use the `kubectl support-bundle` CLI in an Air Gap environment
* **Who this is for**: This lab is for anyone who builds/maintains KOTS applications (see note below)
    * Full Stack / DevOps / Product Engineers
* **Prerequisites**:
    * [Development environment setup from Lab 0](../lab00-hello-world)
    * Basic working knowledge of Kubernetes
* **Outcomes**:
    * You will be ready to deliver a KOTS application into an Air Gap environment
    * You will build confidence in performing upgrades and troubleshooting in Air Gap environments

* **Note** -- a more minimal Air Gap lab is in the works for non-dev teams to learn just the user-side installation
    workflow without needing to understand the building/packaging of new Air Gap versions. 
    Until that is made available, this lab is also appropriate for
    * Implementation / Field Engineers
    * Support Engineers


🔒 Airgap Workflow Overview
===========================

First, we'll push a release -- in the background, Replicated's Air Gap builder will prepare an Air Gap bundle.

![airgap-slide-1](img/airgap-slide-1.png)

Next, we'll collect a license file, a download link, and a public kURL bundle.

![airgap-slide-2](img/airgap-slide-2.png)

From there, we'll move all three artifacts into the datacenter via a jump box.

![airgap-slide-3](img/airgap-slide-3.png)

The above diagram shows a three node cluster, but we'll use only a single node.
While the KOTS bundle will be moved onto the server via SCP as in the diagram,
the app bundle and license file will be uploaded via a browser UI through an SSH tunnel.


🚀 Let's start
==============



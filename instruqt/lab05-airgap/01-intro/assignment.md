---
slug: intro
id: yf4mxxu9mlhv
type: challenge
title: Intro
teaser: An overview of the Airgap Lab
notes:
- type: text
  contents: Lab 5 - Airgap
tabs:
- title: CLI
  type: terminal
  hostname: cli
difficulty: basic
timelimit: 600
---

Lab 1.5: Airgap
=========================================

In this lab, we'll review how to perform installations in Air Gap environments, and
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


## Overview

In this case we'll start with a bare Air Gap server with no KOTS installation, so you can
get practice performing an Air Gap install from scratch.

Once that's done, we'll explore how some of the support techniques differ between online and Air Gap environments.

***
## Airgap Workflow Overview

First, we'll push a release -- in the background, Replicated's Air Gap builder will prepare an Air Gap bundle.

![airgap-slide-1](../assets/airgap-slide-1.png)

Next, we'll collect a license file, a download link, and a public kURL bundle.

![airgap-slide-2](../assets/airgap-slide-2.png)

From there, we'll move all three artifacts into the datacenter via a jump box.

![airgap-slide-3](../assets/airgap-slide-3.png)

The above diagram shows a three node cluster, but we'll use only a single node.
While the KOTS bundle will be moved onto the server via SCP as in the diagram,
the app bundle and license file will be uploaded via a browser UI through an SSH tunnel.

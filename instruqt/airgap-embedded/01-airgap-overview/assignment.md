---
slug: airgap-overview
id: mq22qgecjrh9
type: challenge
title: The Replicated Air Gap Worfklow
teaser: An overview of the Replicated air gap workflow
notes:
- type: text
  contents: Our workflow for air-gapped deployment
tabs:
- title: Jumpbox
  type: terminal
  hostname: jumpbox
  workdir: /home/replicant
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 300
---

👋 Introduction
===============

In this exercise you will learn how to perform installations in air gap environments, and
how to collect support bundles in air gap environments.

* **What you will do**:
    * Access and verify a single-node air gap setup via a bastion server
    * Learn to use KOTS to install in an air gap environment
    * Create an SSH tunnel to configure an air gap instance
    * Perform an upgrade of an application runnning in an air gap
    * Collect a support bundle in an air-gapped environment
* **Who this is for**: This lab is for anyone who builds/maintains KOTS applications (see note below)
    * Full Stack / DevOps / Product Engineers
* **Outcomes**:
    * You will be ready to deliver a KOTS application into an air gap environment
    * You will build confidence in performing upgrades and troubleshooting in air gap environments

🔒 Airgap Workflow Overview
===========================

First, we'll push a release -- in the background, Replicated's air gap builder will prepare an air gap bundle.

![Air Gap Deployment Packaging](../assets/airgap-slide-1.png)

Next, we'll collect a license file, a download link, and a public kURL bundle.

![Air Gap Deployment Delivery](../assets/airgap-slide-2.png)

From there, we'll move all three artifacts into the datacenter via a jump box.

![Air Gap Delivery](../assets/airgap-slide-3.png)

The above diagram shows a three node cluster, but we'll use only a single node.
While the KOTS bundle will be moved onto the server via SCP as in the diagram,
the app bundle and license file will be uploaded via a browser UI through an SSH tunnel.

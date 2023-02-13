---
slug: intro
id: kuutyxh4oiv4
type: challenge
title: Intro
teaser: Introduction to Rapid Development with Replicated
notes:
- type: text
  contents: |-
    This lab is composed of several steps. Do not skip any step as these build on each other.
    In this step, we introduce the lab and check our sandbox environment to make sure we are ready to go.

    Have fun!
tabs:
- title: Dev
  type: terminal
  hostname: shell
  workdir: /home/replicant
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: Cluster
  type: terminal
  hostname: cluster
difficulty: basic
timelimit: 300
---

👋 Introduction
===============

* **What you will do**:
    * Learn to use the `kots download` & `kots upload` command to rapidly iterate on a deployed instance
    * Learn how to take the modified manifests and then create a release
* **Who this is for**: This lab is for anyone who will build KOTS applications **plus** anyone who will be working with your users
    * Full Stack / DevOps / Product Engineers
    * Support Engineers
    * Implementation / Field Engineers
    * Success / Sales Engineers
* **Prerequisites**:
    * Basic working knowledge of Linux (Bash)
    * This is an advanced topic so make sure you have completed the following labs or have relevant hands-on experience
      * Deploy Hello World Application
      * Replicated CLI
* **Outcomes**:
    * You will be ready to rapidly develop on the Replicated platform!


🐚 Get started
===============

In this lab we are going to use several tabs so let's review how we will use each one:

* **Dev**: This tab provides access to our dev environment. Here we already installed the `Replicated` and `KOTS` command lines for your convenience.
* **Vendor**: This tab launches a browser to https://vendor.replicated.com which is the Vendor Portal. We will use this to get our Application Slug, API token and license file.
* **Cluster**: This tab contains access to the VM hosting our K3s cluster, which we will use as our dev cluster.
* **Admin Console**: This tab is not visible in this step but will be later in this lab. The Admin Console will be used to install our sample application that we will iterate on.
* **Code Editor**: This tab is not visible in this step but will be later in this lab. The Code Editor will be used to create a new deployment manifest later in the lab.

## Checking your Environment

Let's make sure we have our environment ready to go before we get started.

Let's start with the **Dev** tab. Not only does it already have our command line tools installed, it also has `kubectl` access to our K3 cluster.
Select the **Dev** tab and try the following commands to ensure we have everything set up correctly:

```bash
replicated version
```

The above should return the version of the `replicated` command line. Anything else points to an issue with the dev environment. Next, let's see if we have access to the cluster:

```bash
kubectl get nodes
```

The above should return a single node, which we can access over on the **Cluster** tab. Before we go there, let's make sure we also have the `kots` command line as well:

```bash
kubectl kots version
```

The above should return the version of the `kots` cli.

Let's head over now to the **Cluster** tab, where you should see the username and password we will use to login to the Vendor Portal. The username and password are based on your Participant ID, which is generated when you start a lab.

We won't use this tab much in this lab, but will be available in case we want to troubleshoot any issues with the cluster or application or need to retrieve the Vendor Portal credentials.


## Vendor Portal login

To access the Vendor Portal, you will need your participant id.  These are
```
username: [[ Instruqt-Var key="USERNAME" host="cluster" ]] 
password: [[ Instruqt-Var key="PASSWORD" host="cluster" ]]
```

Once you have the credentials, you can login into the Vendor tab and you should land on the Channels. Channels allow you to manage who has access to which releases of your application.

**Tip:** Stay logged in to Vendor Portal as you we will be using it on the next step.



---
slug: troubleshooting
id: a0cvc186bru1
type: challenge
title: Troubleshooting
teaser: Using the compatibility-matrix for Troubleshooting
notes:
- type: text
  contents: contents go here
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---

As we saw in the previous exercise, the CM can create a variety of environments. Now let’s look at how we might use it to troubleshoot existing customer installations without connecting to the live/production/airgap installations our customer is running. To start off, we will download a support bundle from an example customer and build a support environment that matches our customer’s. In this case, we have a customer who is not able to upgrade their environment from Kubernetes X to X (dependency issue)

1. The support bundle can be downloaded from the “Bundle” tab at the top of this lab. Go ahead and download it now.

Let's take a look at the support bundle our customer provided and see what kind of a cluster we need to build. First, let's untar the bundle, and use `sbctl` to update our kubeconfig to use it:

``` tar zxvf {{supportbundle-date.tgz}} ```
``` sbctl shell -s . ```
``` kubectl get nodes -n {{namespace}} ```
``` kubectl describe node {{nodename}} | grep -A 3 Resource ```
``` grep k8s {{bundle-dir}}/kots/admin_console/app-info.json ```

With this set of information, we have all we need to mimic our customer's environment for examination. Let's go ahead and build that cluster now so we can take a look:

``` replicated cluster create --memory `




2. Use the bundle to determine requisite versions
3. Build a cluster from the bundle manually
4. Build a cluster from the bundle automatically (possible?)
5. Connect to the cluster and look at the errors the customer is seeing
6. Correct the error and update Kubernetes.

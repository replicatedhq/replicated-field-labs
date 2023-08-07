---
slug: troubleshooting
id: rgkxp5f7zamn
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

As we saw in the previous exercise, the CM can create a variety of environments. Now let’s look at how we might use it to troubleshoot existing customer installations without connecting to the live/production/airgap installations our customer is running. In this case, we'll be looking at a voting app that our customer BigBank is having trouble with. They are having trouble with their voting app, which they need up and running before an important presentation next week where they will use the voting app to interact with the audience. In their testing, whenever more than a few people at a time access the app, it crashes.

Because BigBank is strictly controlled, we aren't able to connect directly to their environment. The support dumps we've gotten so far aren't giving us any indications as to why they might be crashing. We also haven't been able to reproduce the issue on our testing environment. Let's see if we can use the Compatibility Matrix to figure something out.

The lastest support bundle can be downloaded from the “Bundle” tab at the top of this lab. Go ahead and download it now.

Let's take a look at the support bundle and see what kind of a cluster we need to build. First, let's untar the bundle, and use `sbctl` to update our kubeconfig to use it:

``` tar zxvf {{supportbundle-date.tgz}} ```
``` sbctl shell -s . ```
``` kubectl get nodes -n {{namespace}} ```
``` kubectl describe node {{nodename}} | grep -A 3 Resource ```
``` grep k8s {{bundle-dir}}/kots/admin_console/app-info.json ```

With this set of information, we have all we need to mimic our customer's environment for examination. Let's go ahead and build that cluster now so we can take a look:

``` replicated cluster create --memory 8 --vcpu 8 --distribution kind --version 1.27.0 --name big-bank-voting ```

Now we have a cluster, `big-bank-voting`, that matches our customer's system. Let's set our context to use that system, then we can install our app and see what happens:

``` replicated cluster kubeconfig --name big-bank-voting ```
``` kubectl kots install {{app-slug}} ```

During the installation to this cluster, we see that there are preflight failures. This system only has 8GB of memory, as we specified to match our customer's environment. However, preflights are showing that we will need considerably more:

'
Every node in the cluster must have at least 8 GB of memory, with 32 GB recommended

All nodes must have at least 20 GB of memory.
'

After a little bit of testing we are able to duplicate the error. Using the CM again, we can try increasing memory to see if that resolves the issue:

``` replicated cluster create --memory 32 --vcpu 8 --distribution kind --version 1.27.0 --name big-bank-voting-more-mem ```
``` replicated cluster kubeconfig --name big-bank-voting-more-mem ```
``` kube kods install {{app-slug}}}

With additional testing, we are no longer able to reproduce the problem. In other words, BigBank is having trouble due to resource constraints with memory. We can recommend that they increase the memory on their cluster node(s) to 32GB and see if that fixes the issues they're seeing.

Now that we've figured it out, we can clean up our testing clusters. They will disappear automatically in 8 hours, but since we're done with them we can clean up now.

``` replicated cluster rm $ID $ID ```

Now that we've gotten familiar with what CM can do, and also explored a little bit about using it for troubleshooting, let's get to the really good stuff: Using the compatibility matrix to do testing and integrate it into our existing CI/CD pipeline.

To complete this track, click the **Check** button.

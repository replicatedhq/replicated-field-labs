---
slug: creating-a-cluster
id: fav4s8wyibpy
type: challenge
title: Creating a Cluster
teaser: Use the Replicated CLI to create a test cluster
notes:
- type: text
  contents: |
    Let's create a cluster on the command-line using the options we selected in
    the Vendor Portal. The Replicated CLI can help you integrate the
    Compatibility Matrix into your continuous delivery pipeline to test your
    application in the types of clusters your customers are using.
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---

Creating a Cluster from the CLI
===============================

The Vendor Portal is a great way to create an _ad hoc_ cluster, but the
Compatibility Matrix really shines in automaated scenarios. The most common use
is as part of your continous delivery pipeline to create cluster to execute
tests. Your pipeline will likely use the Replicated CLI to create a cluster.

Running the Cluster Create Command
==================================

In the last step of the lab you defined a cluster using the Vendor Portal and
copied the command to create it to the clipboard. Don't worry, if you forgot to
copy it or overwrote it, you can use the command below.

```
replicated cluster create --name [[ Instruqt-Var key="USERNAME" hostname="shell" ]] --distribution eks --instance-type m6i.large --nodes 3 --version 1.27 --disk 100
```

Paste the command into the shell (or type it as above) to create a new cluster.

Watching Your Cluster Come Up
=============================

The cluster will take a few minutes to provision and validate, so you may want
to watch for it to be ready. Run the following command and watch for the
cluster to reach the `running` status. You may see multiple cluster running
since this is a shared lab environment.

```
watch replicated cluster ls
```

Once you see the cluster named `[[ Instruqt-Var key="USERNAME" hostname="shell" ]]`
reach status `running`, it is ready to use. Hit `<Ctrl>-C` to exit the
`watch` command.

Accessing Your Cluster
======================

Once your cluster is running, you can use the Replicated CLI to get access to
it. The following command will update the Kubernetes client configuration to
allow access to the cluster.

```
replicated cluster kubeconfig --name [[ Instruqt-Var key="USERNAME" hostname="shell" ]]
```

You can run any `kubectl` command at this point. Let's run `kubectl get nodes`
to see the nodes your specified for the cluster.

```
kubectl get nodes
```

This should show the nodes of your cluster. If you followed my lead, you'll see
three nodes similar to the ones below.

```
NAME                          STATUS   ROLES    AGE     VERSION
ip-10-0-150-35.ec2.internal   Ready    <none>   2m48s   v1.27.6-eks-a5df82a
ip-10-0-185-13.ec2.internal   Ready    <none>   2m46s   v1.27.6-eks-a5df82a
ip-10-0-63-26.ec2.internal    Ready    <none>   2m48s   v1.27.6-eks-a5df82a
```

🏁 Finish
=========

After running these commands in your pipeline, you could set up and run your
test suite and validate your application running with the distribution and
version you selected.

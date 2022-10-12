---
slug: vendor-portal
id: pxcgks564wye
type: challenge
title: vendor-portal
teaser: Work with support bundles received from a customer
notes:
- type: text
  contents: How can a customer share a support bundle with you.
tabs:
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: Application Installer
  type: website
  url: http://kubernetes-vm.${_SANDBOX_ID}.instruqt.io:8800
  new_window: true
- title: Upload
  type: code
  hostname: shell
  path: /home/replicant
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---

Download the support bundle
===========================

The key lesson of this track is around what to do if analyzers don't provide any useful information. So far we've been mainly assuming we always have access to the Application Installer. But that will not be the case for many customers. So let's walk through downloading the bundle, collaborating with your team to diagnose the error, and preventing it in the future.

<div align="center"><blockquote><h3>If support analyzers don't surface any issues, the bundle should be downloaded and sent to the vendor team</h3></blockquote></div>

In this case, even though we're wearing a "customer" hat for parts of these exercises, "the vendor" is also you in this case :slightly_smiling_face:

Let's grab the bundle using the "download bundle" option so we can upload it to https://vendor.replicated.com to collaborate with the team:

![download-bundle](../assets/download-bundle.png)


Upload in Vendor Portal
=======================

Next, we'll navigate to https://vendor.replicated.com/troubleshoot and choose "upload bundle", then drag and drop the bundle you downloaded from the instance.
Having access to the instance to download the bundle simplifies things for us in this case.
In the real world, you can receive the bundle from a customer by whatever means is most secure/convenient for them.

![upload-bundle](../assets/upload-bundle.png)

Once it's uploaded, you should see a similar view of the analyzers.

![vendor-analyzers](../assets/vendor-analyzers.png)

This view is identical to what your customer will see!

Next we'll dig into the logs by navigating to the file inspector tab.
We can use the line-level deep linking in the UI to highlight and share specific errors and log lines.

**Exercise** navigate to `cluster-info/cluster_version.json`, and determine the version of Kubernetes that is running on the server side.
Test the line-level linking by copying the URL into a new tab.

![bundle-line-links](../assets/bundle-line-links.png)

Note that as you click line numbers in the gutter, the URL changes.

<div align="center"><blockquote><h3>Once a bundle has been uploaded to the vendor console, use deep-linking to collaborate on the diagnosis with your team.</h3></blockquote></div>

These links can be attached to a support ticket, shared in Slack, etc.

If you choose the "Share with Replicated" option in the https://vendor.replicated.com, the same links will become accessible to Replicated's support team.


Support Bundle ctl
==================

As you might have noticed, besides the application logs, the support bundle also contains a lot more information. Each support bundle does contain some "default" collectors that always are included. One of them is the [cluster-resources](https://troubleshoot.sh/docs/collect/cluster-resources/) which will capture all kubernetes resources from the cluster.

In the `Vendor Portal` tab, you can easily browse all the collected output if you go to `File Inspector` and open one of the files under `cluster-resources`

![Cluster resources](../assets/cluster-resources.png)

Although the collected information is all `json`, there is a lot of information in it that can be very helpfull! And there is even an easy way to interact with it called `sbctl`: a command line tool for examining K8s resources in Troubleshoot's support bundles.

The `Shell` tab already has the `sbctl` cli installed. So all we have to do is upload the previously downloaded support bundle into it. Go the the `Editor` tab, and upload the `.tar.gz` file into it.

Next, you can go to the `Shell` tab and run:

```bash
sbctl shell -s /home/replicant/[YOUR_TAR_GZ_FILE]
```

This will allow you to interact with the support bundle, using `kubectl`. Try some of the following commands, and see what info you get back:

```bash
kubectl get ns
```

```bash
kubectl get pods -n kube-system
```

```bash
kubectl describe deployment nginx -n support-bundle
```

Congratulations! You have completed this track!
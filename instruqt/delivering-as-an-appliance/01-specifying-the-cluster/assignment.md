---
slug: specifying-the-cluster
id: rsxl9h79cen1
type: challenge
title: Starting Your Appliance Configuration
teaser: Enable and configure the embedded cluster
notes:
- type: text
  contents: Let's build a Kubernetes appliance
tabs:
- title: Shell
  type: terminal
  hostname: node
- title: Release Editor
  type: code
  hostname: node
  path: /home/replicant
difficulty: basic
timelimit: 300
---

A virtual Kubernetes appliance consists of your application, a Kubernetes
cluster, and a console that you customer uses to install your application and
manage the appliance. We refer to the Kubernetes cluster that's included as the
Embedded Cluster since it's "embedded" with your application. The console to
isntall and manage your application is called the Admin Console. Under the
hood, the Admin Console uses Helm to install and upgrade your application.

You release your application as a Kubernetes applicance by releasing a Helm
chart and some additional configuration on the Replicated Vendor Portal. The
moost important file is the configuration for the Embedded Cluster, which can
very simple. All it requires is the version of the cluster to use.

```
apiVersion: embeddedcluster.replicated.com/v1beta1
kind: Config
spec:
  version: [[ Instruqt-Var key="EMBEDDED_CLUSTER_VERSION" hostname="node" ]]
```

That specification lets the Replicated Vendor Portal know which version of
Kubernetes to embed. There are [more options for the configuration]
(https://docs.replicated.com/reference/embedded-config), but that's all you
need to get started.

The Replicated Release
======================

The Replicated Platform distributes your software to your customers. To do
this, it needs to know about your application, it's customers, and the files
you're shipping to them. We talk about those files as a release.

Every release is built around a Helm chart, and that's all it needs. If you've
completed the [Distributing Your Software with
Replicated](https://play.instruqt.com/replicated/tracks/distributing-with-replicated)
lab, you built a release around a Helm chart and installed it using Helm tools.
In this lab, we'll add the cluster configuration and a few other files to the
release to enable the appliance experience.

Your Initial Appliance
======================

Let's create a simple Kubernetes appliance for Slackernews and release it with
the Platform. We're going to add our Helm chart and two configuration files to
the `release` directory. These are the bare minimum set of files we need to
create the appliance.

### Adding a Helm Chart to a Release

To prepare the release, we first need to make sure our Helm chart is configured
as part of it. This lab uses the application
[Slackernews](https://slackernews.io) and the Helm sources for it are in the
`slackernews` directory. Let's package the chart and include it in the
`release` directory, where we'll also add the additional files we need.

```
helm package -u slackernews -d release --version 0.6.0
```

We're then going to add a file that lets the Admin Console know about the Helm
chart. It uses this file to identify which chart to install and pass the
appropriate values. We'll look at the passing values in a later section of the
lab when we set up the configuration screen for our application.

Go to the "Release Editor" tab to add a file to your release.

![Creating a manifest file describing your Helm
chart](../assets/creating-the-helmchart-object.png)

The editor may not open your new file automatically. If it doesn't, click on it
to open it. Add the following content to the file. Note that it looks like a
Kubernetes custom resource, but it's really not. Instead, it's processed by the
Admin Console to avoid the complexity of creating a CRD and the relevant
controllers.

```
apiVersion: kots.io/v1beta2
kind: HelmChart
metadata:
  name: slackernews
spec:
  # chart identifies a matching chart from a .tgz
  chart:
    name: slackernews
    chartVersion: 0.6.0

  # values are used in the customer environment, as a pre-render step
  # these values will be supplied to helm template
  values: {}
```

The name and version in this file need to match the metadata for our
Slackernews Helm chart to identify it correctly. You're not limited to only one
Helm chart as part of your application. Including multiple `HelmChart` objects
let's the Admin Console know it has to install multiple Helm charts. For
Slackernews, we have only a single chart.

### Including the Embedded Cluster

We showed a simple Embedded Cluster configuration earlier in the lab. We're
going to use that basic configuration for Slackernews. Create another file in
the `release` folder named `embedded-cluster.yaml` and copy the contents of the
configuration into it.

```
apiVersion: embeddedcluster.replicated.com/v1beta1
kind: Config
spec:
  version: [[ Instruqt-Var key="EMBEDDED_CLUSTER_VERSION" hostname="node" ]]
```

Releasing Your Appliance
========================

---
slug: airgap-deployment-assets
id: cqt3rfjrlyp8
type: challenge
title: Air Gap Deployment Assets
teaser: The files you need to deploy your air-gapped instance
notes:
- type: text
  contents: The assets that make up your air-gapped deployment
tabs:
- title: Jumpbox
  type: terminal
  hostname: jumpbox
  workdir: /home/replicant
difficulty: basic
timelimit: 800
---

Air Gap Deployment Assets
=========================

Our next step is to collect the assets we need for an Air Gap installation:

1. A license with the Air Gap entitlement enabled
2. An Air Gap bundle containing the kURL cluster components
3. An Air Gap bundle containing the application components

Items (2) and (3) are separate artifacts to cut down on bundle size during
upgrade scenarios where only the application version is changing and the
underlying cluster does not need to change.

Starting the kURL Bundle Download
=================================

We're going to start with downloading the bundle for the kURL cluster. This will
turn our air-gapped instance into a single-node Kubernetes cluster which in
turn will run our application. We start with the cluster download since it's
the largest of the three assets and we can download the others while its
download is running.

From the "Jumpbox" tab, run the command below:

```
replicated channel inspect Unstable
```

This command shows you the details of the `Unstable` release channel,
which we'll use for the air-gap install.

The channel details include the information you need to install the
application in one of three ways: into an existing (connected) cluster,
onto a connected machine without a pre-existing clsuter available (the
install includes its own "embedded" cluster"), and onto an
air-gapped machine. The air-gapped install method also includes it's
own cluster.

```text
ID:             2GSIQnqsl6ybfWVacXrutlLQa28
NAME:           Unstable
DESCRIPTION:
RELEASE:        1
VERSION:        Installing in an Air-Gapped Environment
EXISTING:

    curl -fsSL https://kots.io/install | bash
    kubectl kots install uws24vkeurcz-replicated-labs-com-caiman/unstable

EMBEDDED:

    curl -fsSL https://k8s.kurl.sh/uws24vkeurcz-replicated-labs-com-caiman-unstable | sudo bash

AIRGAP:

    curl -fSL -o uws24vkeurcz-replicated-labs-com-caiman-unstable.tar.gz https://k8s.kurl.sh/bundle/uws24vkeurcz-replicated-labs-com-caiman-unstable.tar.gz
    # ... scp or sneakernet uws24vkeurcz-replicated-labs-com-caiman-unstable.tar.gz to airgapped machine, then
    tar xvf uws24vkeurcz-replicated-labs-com-caiman-unstable.tar.gz
    sudo bash ./install.sh airgap
```

The file download we're interested in is in the `AIRGAP` section of the
output. We're going to run the first command in that list to get the bundle
onto our jumpbox.

In my case:

```bash
curl -fSL -o uws24vkeurcz-replicated-labs-com-caiman-unstable.tar.gz https://k8s.kurl.sh/bundle/uws24vkeurcz-replicated-labs-com-caiman-unstable.tar.gz
```

This will take several minutes, leave this running and proceed to the next step, we'll come back in a few minutes.


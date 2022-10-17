---
slug: moving-assets-into-place
id: cqt3rfjrlyp8
type: challenge
title: Moving Assets into Place
teaser: Preparing the files you need to deploy your air-gapped instance
notes:
- type: text
  contents: Let's prepare the files we need for our deployment
tabs:
- title: Jumpbox
  type: terminal
  hostname: jumpbox
  workdir: /home/replicant
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 800
---

## Moving Assets into place

Our next step is to collect the assets we need for an Air Gap installation:

1. A license with the Air Gap entitlement enabled
2. An Air Gap bundle containing the kURL cluster components
3. An Air Gap bundle containing the application components

(2) and (3) are separate artifacts to cut down on bundle size during upgrade
scenarios where only the application version
is changing and no changes are needed to the underlying cluster.

#### Starting the kURL Bundle Download

We're going to start with downloading the bundle for kURL cluster. This will
turn our air-gapped instance into as single-node Kubernetes cluster which in
turn will run our application. We start with the cluster download since it's
the largest of the thres assets and we can download the others while its
download is running.

From the "Jumpbox" tab, run the command below:

```
replicated channel inspect development
```

This command shows you the details of the `development` release channel, 
which we'll use for the air-gap install.

The channel details include the information you need to install the
application in one of three ways: into an existing (connected) cluster, 
onto a connected machine without a pre-existing clsuter available (the
install includes its own "embedded" cluster"), and onto an
air-gapped machine. The air-gapped install method also includes it's  
own cluster.

```text
ID:             2G8bwopWjbBhyGqKut11tpMcTz6
NAME:           development
DESCRIPTION:    
RELEASE:        2
VERSION:        Installing in an Air-Gapped Environment
EXISTING:

    curl -fsSL https://kots.io/install | bash
    kubectl kots install uws24vkeurcz-replicated-labs-com/development

EMBEDDED:

    curl -fsSL https://k8s.kurl.sh/uws24vkeurcz-replicated-labs-com-development | sudo bash

AIRGAP:

    curl -fSL -o uws24vkeurcz-replicated-labs-com-development.tar.gz https://k8s.kurl.sh/bundle/uws24vkeurcz-replicated-labs-com-development.tar.gz
    # ... scp or sneakernet uws24vkeurcz-replicated-labs-com-development.tar.gz to airgapped machine, then
    tar xvf uws24vkeurcz-replicated-labs-com-development.tar.gz
    sudo bash ./install.sh airgap
```

The file download we're interested in is in the `AIRGAP` section of the 
output. We're going to run the first command in that list to get the bundle 
onto our jumpbox.

In my case:

```text
curl -fSL -o uws24vkeurcz-replicated-labs-com-development.tar.gz https://k8s.kurl.sh/bundle/uws24vkeurcz-replicated-labs-com-development.tar.gz
```

This will take several minutes, leave this running and proceed to the next step, we'll come back in a few minutes.


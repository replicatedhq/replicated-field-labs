---
slug: installing-your-release
id: 2hcel4ej9cx3
type: challenge
title: Installing your release
teaser: Let's install your release
notes:
- type: text
  contents: Installing the new release
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
- title: Kuard
  type: website
  url: http://kubernetes-vm.${_SANDBOX_ID}.instruqt.io:8080
  new_window: true
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 1200
---

💡 Installing your release
==========================

This challenge may seem familiar if you've already run through
the "Hello World" track.  We're going to go through the same
steps to install using Replicated's command line tooling.

### 0. Configuring API access

We're running this install from a different shell than our other
challenges, so we'll need to make sure we can connect to the
Replicated API. Let's set that up like we did in the early
challenge.

You'll need the API token that you set up earlier for this
step. Hopefully you kept track of it, otherwise you can log
into the vendor portal and create another one. Use the tab
"Vendor" do to that if you need to (see below for a resfresher).

```
export REPLICATED_API_TOKEN=...
```

Next we'll set up our application slug so the command line
knows which application we're working with. You can uyse
Now, you'll need to set up environment variables to interact
with. If you don't recall it from the earlier challenges,
you can list your applications with

```
replicated app ls
```

which will show something like

```
ID                             NAME                                                     SLUG                                                     SCHEDULER
2FK67XXGtleLCE8bm3KJnLdezww    chuck-instruqt-replicon-2022q3-replabs-replicated-com    chuck-instruqt-replicon-2022q3-replabs-replicated-com    kots
```

Set `REPLICATED_APP` to the application slug.

```
export REPLICATED_APP=...
```

### 1. Getting the install command

Each channel for  your application has a custom install command.
You can view the install command either in the vendor portal or
using the `replicated` command line tool. Let's get it using the
CLI.

```
replicated channel inspect replicated-cli
```

We're going to install into an existing Kubernetes cluster.
You'll find that one listed first, with the label `EXISTING`.


```
ID:             2FKLzwElQOFuHs6YlYEvZ6ncNEo
NAME:           replicated-cli
DESCRIPTION:
RELEASE:        4
VERSION:        0.0.1
EXISTING:

    curl -fsSL https://kots.io/install | bash
    kubectl kots install chuck-instruqt-replicon-2022q3-replabs-replicated-com/replicated-cli

EMBEDDED:

    curl -fsSL https://k8s.kurl.sh/chuck-instruqt-replicon-2022q3-replabs-replicated-com-replicated-cli | sudo bash

AIRGAP:

    curl -fSL -o chuck-instruqt-replicon-2022q3-replabs-replicated-com-replicated-cli.tar.gz https://k8s.kurl.sh/bundle/chuck-instruqt-replicon-2022q3-replabs-replicated-com-replicated-cli.tar.gz
    # ... scp or sneakernet chuck-instruqt-replicon-2022q3-replabs-replicated-com-replicated-cli.tar.gz to airgapped machine, then
    tar xvf chuck-instruqt-replicon-2022q3-replabs-replicated-com-replicated-cli.tar.gz
    sudo bash ./install.sh airgap
```

Your outputs will be in the same format but a bit different, since you'll
be installing your own application.

The command you'll use will look like this:

```
curl -fsSL https://kots.io/install | bash
kubectl kots install [YOUR APP NAME]/replicated-cli
```

We'll come back to that in a later step.

### 2. Download a Customer License

A customer license (downloadable as a `.yaml` file) is required
to install any KOTS application. We're going to use the command-line
to both create a customer and to download their license file.

```
replicated customer create --name "Replicant" --channel replicated-cli
```

once the customer is created, we'll download their license file to use
as part of the install.

```
replicated customer download-license --customer "Replicant" > license.yaml
```

### 3. Run your install

When you looked up your install command, you saw something
like this.

```
curl -fsSL https://kots.io/install | bash
kubectl kots install [YOUR APP NAME]/replicated-cli
```

This command installs the `kots` plugin to `kubectl` and then
starts the install of your application. It then sets up a port
forward to the admin console, where you finish the install. We
are going to skip the console element and work entirely from
the command line.

Our shell already has the `kots` plugin installed, so we can
skip the first line. We are also going to embelish the second
line a little bit to fill in the values that your customer
would typically enter into the admin console.

We're also setting the password for the admin console that
Replicated provides for managing the application. Since we're
setting it on the command-line, let's use `this-is-unsafe` as
a reminder not to leak secrets in the real world.

```
kubectl kots install ${REPLICATED_APP}/replicated-cli --namespace kuard \
  --shared-password this-is-unsafe --license-file ~/license.yaml
```

#### 4. Check your application

Click on the "Kuard" tab to see you application running.

🏁 Finish
=========

Congratulations! You've finished the "Replicated CLI" track.

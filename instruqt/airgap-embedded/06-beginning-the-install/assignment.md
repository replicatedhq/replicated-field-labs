---
slug: beginning-the-install
id: bsgw1cyzyjgd
type: challenge
title: Beginning the Install
teaser: Kicking of your air-gapped install`
notes:
- type: text
  contents: Starting the air-gapped install
tabs:
- title: Jumpbox
  type: terminal
  hostname: jumpbox
  workdir: /home/replicant
difficulty: basic
timelimit: 600
---

In this lab we're using kURL to turn our air-gapped system into a single-node
Kuberenetes cluster. Our installation begins by connecting to the air-gapped
instane over SSH. From there we unpack the kURL bundle and run the included
install script. When these steps are complete we'll be able to install our
application into the cluster.

Let's Connect
=============

Since we're installing directly on the air-gapped node, let's connect to it.
We're using our jumpbox's SSH connection into the air-gap network.

```shell
ssh cluster
```

Once you're on the node, untar the bundle and run the install script
with the `airgap` flag. kURL install flags are documented
[in the kurl.sh docs](https://kurl.sh/docs/install-with-kurl/advanced-options).
Your bundle name will vary, but end with the `-replicated-labs-com-development.tar.gz`
suffix like mine does.

```shell
tar xvzf uws24vkeurcz-replicated-labs-com-development.tar.gz
sudo bash install.sh airgap
```

At the end, you should see a `Installation Complete` message as shown below.
Since the instance is Air Gap, we'll need to use a port forward and proxy for
the UI in the next step. Note that the IP address and password you see will
differ from mine.

```text
configmap/kurl-config created


		Installation
		  Complete ✔


Kotsadm: http://10.128.1.47:30880
Login with password (will not be shown again): iunIEfPyc
This password has been set for you by default. It is recommended that you change this password; this can be done with the following command: kubectl kots reset-password default
```

You'll need the password in next step, so be sure to copy it and put it
in your notes.

Disconnect
==========

Make sure you disconnect from the cluster node before moving on to the
next step. Press "control-D" to log out.

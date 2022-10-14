---
slug: building-an-air-gap-release
id: uppk7yks93gg
type: challenge
title: Building an Air Gap Release
teaser: A short description of the challenge.
notes:
- type: text
  contents: Replace this text with your own text
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
timelimit: 600
---

#### Building an Airgap Release

By default, only the Stable and Beta channels will automatically build Air Gap bundles

- manually build
- set channel to auto build

For a production application, Air Gap releases will be built automatically on the Stable channel, so this won't
be necessary.

In this case, since we're working off the `lab05-airgap` channel, you'll want to enable Air Gap builds on that channel.

You can check the build status by navigating to the "Release History" for the channel.

![release-history](../assets/channel-release-history.png)

You can build invividual bundles on the Release History page, but you'll likely want to edit the channel and enable "build all releases for this channel".

![edit-channel](../assets/channel-edit-info-btn.png)

![auto-build](../assets/channel-enable-airgap.png)

Now you should see all the bundles building or built on the release history page. If you do not see "Airgap Built" for the release, click **Build**.

![airgap-built](../assets/airgap-builds.png)

#### Enabling Airgap for a customer

The first step will be to enable Air Gap for the `lab5` customer:

![enable-airgap](../assets/airgap-customer-enable.png)


#### Download Airgap Assets
After saving the customer, scroll to the bottom of the page to the `Download Portal` section.

![download-portal](../assets/airgap-customer-portal.png)

Generate a new password and save it somewhere in your notes.
Next, click the link to open the download portal.
This is a link you would usually send to your customer, so from here on we'll be wearing our "end user" hat.


Navigate to the "embedded cluster" option and review the three downloadable assets.

![download-portal-view](../assets/download-portal-view.png)

Download the license file, but **don't download the kURL bundle** -- this is the download we already started on the server.

You'll also want to download the other bundle `Latest Lab 1.5: Airgap Bundle` to your workstation.

From your jumpbox, check that the download has finished, so you can copy it to the Air Gap server. If you have not started the download, see the [Starting the kURL Bundle Download](#starting-the-kurl-bundle-download) instructions above.

You can use the DNS name in this case, as described in [Instance Overview](#instance-overview).

```bash
export REPLICATED_APP=... # your app slug
export FIRST_NAME=... # your first name
scp kurlbundle.tar.gz ${REPLICATED_APP}-lab05-airgap:/home/${FIRST_NAME}
```

> **Note**: -- we use SCP via an SSH tunnel in this case, but the Air Gap methods in this lab also extend to
more locked down environments where e.g. physical media is required to move assets into the datacenter.

Now we'll SSH all the way to Air Gap node. If you still have a shell on your jump box, you can use the instance name.

```bash
ssh ${REPLICATED_APP}-lab05-airgap
```

Otherwise, from your local system you can use the one below

```shell
ssh -J ${FIRST_NAME}@${JUMP_BOX_IP} ${FIRST_NAME}@${REPLICATED_APP}-lab05-airgap
```

Once you're on the Air Gap node, untar the bundle and run the install script with the `airgap` flag.
kURL install flags are documented [in the kurl.sh docs](https://kurl.sh/docs/install-with-kurl/advanced-options).

```shell
tar xvf kurlbundle.tar.gz
sudo bash install.sh airgap
```

At the end, you should see a `Installation Complete` message as shown below. Since the instance is Air Gap, we'll need to create a port forward to access the UI from your workstation in the next step.

```text
configmap/kurl-config created


		Installation
		  Complete ✔


Kotsadm: http://10.128.1.47:30880
Login with password (will not be shown again): iunIEfPyc
This password has been set for you by default. It is recommended that you change this password; this can be done with the following command: kubectl kots reset-password default
```

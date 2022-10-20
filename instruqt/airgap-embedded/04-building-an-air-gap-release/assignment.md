---
slug: building-an-air-gap-release
id: uppk7yks93gg
type: challenge
title: Building an Air Gap Release
teaser: Creating the air-gap deployment bundle for your releaese
notes:
- type: text
  contents: Let's build your airgap release
tabs:
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 600
---

Building an Airgap Release
==========================

While the kURL download continues, let's build the air-gap bundle for
our current application release. You can see your download status in
the "Jumpbox" tab. We'll mostly be working in the "Vendor Portal" for
this lab.

Since air-gapped bundles can be quite large, not all release channels
build them automatically. Two of the default channels, `Stable` and
`Beta` are configured to do it. Other channels can be set to auto
build, or you can manually build a bundle when needed.

You can build air-gap bundles for any of your releases by looking at
the release history page for a channel. If you have a channel that
you'll want regular airgap bundles on, you'll likely want to edit
that channel to enable auto builds.

![Release History for the development Channel](../assets/channel-release-history.png)

Let's go through setting that up for our `development` channel we're
using for this lab. Start by editing the channel info

![Editing Channel Details](../assets/channel-edit-info-btn.png)

then enable auto-builds by flipping the toggle labeled "Automatically
create airgap builds for newly promoted release in this channel"

![Enabling Automatic Airgap Builds](../assets/channel-enable-airgap.png)

Enabling Airgap for a customer
==============================

Customers are automatically enabled for air gap deployment. This
gives you flexibility in terms of product packaging and deployment.
Let's enable air-gap downloads for the example customer we're using
for the lab. Go to "Customers" in the Vendor portal and select the
"Replicant" customer to enable the airgap.

![Enabling Airgap Downloads for a Custtomer](../assets/airgap-customer-enable.png)

Click the checkbox next to "Airgap Download Enabled" and make sure
you "Save Changes" with the bottom on the bottom right.

Download Airgap Assets
======================

After saving the customer, scroll to the bottom of the page to the
`Download Portal` section.

![Customer Download Portal Section](../assets/airgap-customer-portal.png)

Generate a new password and save it somewhere in your notes. Next,
click the link to open the download portal. This is a link you would
usually send to your customer. The rest of the lab we'll be back on
our Jumpbox and working wearing our "end user" hat.

Navigate to the "embedded cluster" option and review the three
downloadable assets.

![Viewing the Customer Download Portal](../assets/download-portal-view.png)

This is where your customer downloads the assets they neeed for an
air gap install: the kURL bundle, their license file, and the airgap
bundle for your application. They will usually download all three at
once for the initial install.

We're only going to download the license file and the application
bundle right now. We started downloading the kURL bundle in the previous
step, so we don't need to do that again.


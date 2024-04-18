---
slug: proxying-private-images
id: wlrdo1mqkezo
type: challenge
title: Proxying Private Images
teaser: Protect your private container images with the Replicated proxy registry
notes:
- type: text
  contents: |
    Share your private images without exposting your private registry to your customers
tabs:
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: Shell
  type: terminal
  hostname: shell
  workdir: /home/replicant
difficulty: basic
timelimit: 600
---

One of the core features of the Replicated Platform is it's proxy registry. The
proxy registry controls acccess to your images using the Replicated license
This relieves you of the burden of managing authentication and authorization
for the private images your application depends on. You provide Replicated with
access and we manage the rest.

Using the Proxy Registry
------------------------

The first step in using the Replicated proxy registry is to provide access to
your private registry for the Replicated Vendor Portal. This has already been
done in the lab environmnet, so we're just going to review how it was set.

Log into the Vendor Portal with the following credentials:

Username: `[[ Instruqt-Var key="USERNAME" hostname="shell" ]]`<br/>
Password: `[[ Instruqt-Var key="PASSWORD" hostname="shell" ]]`

From the "Channels" page you landed on, select "Images" in the left-hand menu.
You'll land 

![Managing Images on the Vendor Portal](../assets/managing-images.png)

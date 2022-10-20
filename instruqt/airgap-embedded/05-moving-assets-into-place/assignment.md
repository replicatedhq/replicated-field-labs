---
slug: moving-assets-into-place
id: 097iz3ahw0xv
type: challenge
title: Moving Assets Into Place
teaser: Getting your airgap assets ready for deployment
notes:
- type: text
  contents: Getting our airgap assets ready to deploy
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

Recall the three assets we need for an Air Gap installation:

1. A license with the Air Gap entitlement enabled
2. An Air Gap bundle containing the kURL cluster components
3. An Air Gap bundle containing the application components

We've already begun the download of item (2), since it's the largest
one and we needed some time for it to complete. We also saw how your
customer gets access to all three assets from the Replicated download
portal, then grapped their license file. Now we're going to grab the
application bundle using the command line.

At this point, your download of the kURL bundle should be completed. If
not, wait for it to complete before you continue.

Copying the Airgap Bundle to the Air-Gapped Host
================================================

In the wild, this step could take many forms. Customers may have a "soft"
airgap like the one we us in this lab, which means you can copy the
file directly from a jumpbox or through a bastion host. They may have
a much more robust airgap requiring a cross-domain solution or burning
the bundles to read-only media and using a sneakernet to get move that
media to the other side. Regardless of the requirements, as long as the
three assets are available on the isolated network they'll be able to
install your application.

In our simple case, we're going to use the SSH connection between our
Jumpbox and the airgapped instance to manage our install. First, we'll
copy the bundle files to the instance using `scp`, then we'll use
an SSH tunnel to configure the application and complete the install.

Let's copy the kURL bundle first. Take a look at the files in your
current directory. There should be one (the name of your kURL bundle
will different but have the same ending after the random identifier
at the beginning).

```
replicant@jumpbox:~$ ls
uws24vkeurcz-replicated-labs-com-development.tar.gz
```

Since there's only one file, let's save some typing and use a
wildcard. We're copying the to the machine named `cluster` which
will become a single-node Kubernetes cluster running our application.

```
scp *.tar.gz cluster:
```

Ready to Install
================

Now that the bundle is in place, we're ready to install.

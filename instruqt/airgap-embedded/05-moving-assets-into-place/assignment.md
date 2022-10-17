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
one and we needed some time for it to completed. We also saw how your
customer gets access to all three assets from the Replicated download
portal, then grapped their license file. Now we're going to grab the
application bundle using the command line.

### Downloading the Application Bundle

At this point, your download the kURL bundle should be completed. If
not, wait for it to complete before you continue.

The next step is to install the application download bundle. You copied
at the end of the last step, lets paste it as the argument to the
`curl` command in the Jumpbox shell. The URL carries a lot of detail in
it's parameters, so don't be surprised by it's length. Also make sure
it's in quotes.

```
curl -fSL -o application-bundle.tar.gz '[PASTE URL HERE]'
```

Wait for this download to finish, then we'll move our two bundles onto
the air-gapped host.

### Copying the Airgap Bundles to the Air-Gapped Host

In the wild, this step could take many forms. Customers may have a "soft"
airgap like the one we us in this lab, which means you can copy the
file directly from a jumpbox or throuhg a bastion host. They may have
a much more robust airgap requiring a cross-domain solution or burning
the bundles to read-only media and using a sneakernet to get move that
media to the other side. Regardless of the requirements, as long as the
three assets are available on the isolated network they'll be able to
install your application.

In our simple case, we're going to use the SSH connection between our
Jumpbox and the airgapped instance to manage our install. First, we'll
copy the two bundle files to the instance using `scp`, then we'll use
an SSH tunnel to configure the application and complete the install.

Let's copy the kURL bundle first. Take a look at the files in your
current directory. There should be two:

```
replicant@jumpbox:~$ ls
application-bundle.tar.gz  uws24vkeurcz-replicated-labs-com-development.tar.gz

```

We're going to move both files, so let's be efficient and use a
wildcard.

```
scp *.tar.gz cluster:
```

### Ready to Install

Now that the assets are in place, we're ready to start the install.

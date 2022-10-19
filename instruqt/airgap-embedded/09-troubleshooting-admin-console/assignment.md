---
slug: collecting-a-support-bundle
id: kryjgrtoqw2l
type: challenge
title: What if the Admin Console fails?
teaser: Addressing a failed admin console in an air-gapped environment
notes:
- type: text
  contents: Troubelshooting the air-gapped Admin Console
tabs:
- title: Jumpbox
  type: terminal
  hostname: jumpbox
  workdir: /home/replicant
difficulty: basic
timelimit: 600
---

If your application is at least partially installed in the customer
enviornmnet, they can use the support bundle you defined to troubelshoot
application failures. But what happens if it fails before that?


The Admin Console Support Bundle
================================

In a connected network, there is an easy way to connect a support
bundle from a failing Admin Console

```shell
kubectl support-bundle https://kots.io
```

Let's connect to the air-gapped cluster and give this a try:

```
TO DO
```

As expected, this doesn't work. Looks like you'll need to a 
local support bundle specification that does the same thing.

Moving the KOTS Support Bundle Spec
===================================

We're going to pull the spec from the KOTS project, then move
it onto the cluster node. The steps will vary depending on how your
airgap is configured.

In our case, we can grab it from GitHub onto our Jumpbox, then
move the file up onto the air-gapped instance and use it from
there.

```
curl -o support-bundle.yaml https://github.com/replicatedhq/kots/blob/master/pkg/supportbundle/defaultspec/spec.yaml
scp support-bundle.yaml cluster:
```

Collecting the Support Bundle
=============================

Now we can connect to our cluster node and collect the support
bundle.

```shell
kubectl support-bundle ./support-bundle.yaml
```

There's an in depth post with some other options at [How Can I Generate a Support Bundle If 
I Cannot Access the Admin Console?](https://help.replicated.com/community/t/kots-how-can-i-generate-a-support-bundle-if-i-cannot-access-the-admin-console/455).

🏁 Finish
=========

Congratulations! You've now explored the air-gapped workflow
for applications deployed with Replicated.

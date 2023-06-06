---
slug: preparing-to-use-the-sdk
id: kkhcf1j4eaym
type: challenge
title: Preparing to Use the SDK
teaser: Getting ready to use the Replicated SDK
notes:
- type: text
  contents: Let's get ready to use the Replicated SDK
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 300
---

👋 Introduction
===============

The Replicated SDK is implemented as a small service that runs
alongside your application and enables access to the Replicated
Platform. The SDK allows you to enforce your entitlements and
take advantage of the telemetry that Replicated provides to help
you better understand customer instances.

✅ Preparing to use the SDK
===========================

To make use of the Replicated SDK, you'll need a couple of
things. In the lab environment. They've been set up for you
in this lab environment.

1. A Helm chart for your application. We're going to use the
   Open Source Bitnami Helm chat for the Harbor container registry
   for this lab, but it could be the Helm chart for your product
   or any other Helm chart you wanted to distribute this way.
2. Access to the Replicated Vendor Portal. You've been given
   access for the duration of this lab with the username
   `[[ Instruqt-Var key="USERNAME" hostname="shell" ]]` and
   the password `[[ Instruqt-Var key="PASSWORD" hostname="shell" ]]`
3. An application on the Replicated Vendor Portal. We've also
   created that for you as part of the lab setup.
4. A customer for that application. We've created the customer as
   well.

🔤 Getting Started
==================

To use the SDK, we need to add a dependency to the Harbor
Helm chart. Let's pull down the chart so that we can get
started.

```bash
helm pull oci://registry-1.docker.io/bitnamicharts/harbor --untar
```

Let's also set up our shell for interacting with the Replicated
platform.

```
export REPLICATED_API_TOKEN="[[ Instruqt-Var key="REPLICATED_API_TOKEN" hostname="shell" ]]"
```

And lastly make sure we are working with the Harbor app

```
export REPLICATED_APP="[[ Instruqt-Var key="REPLICATED_APP" hostname="shell" ]]"
```


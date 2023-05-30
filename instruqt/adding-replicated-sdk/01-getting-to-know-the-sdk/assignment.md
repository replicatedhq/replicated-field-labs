---
slug: getting-to-know-the-sdk
id: svl7xy6jgkzb
type: challenge
title: Getting to Know the SDK
teaser: Learn about the Replicated SDK
notes:
- type: text
  contents: Let's learn about the Replicated SDK
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
   Open Source Bitnami Helm chat for the Mastadoon social network.
2. Access to the Replicated Vendor Portal. You've been given
   access for the duration of this lab with the username
   `[[ Instruqt-Var key="USERNAME" hostname="shell" ]]` and
   the password `[[ Instruqt-Var key="PASSWORD" hostname="shell" ]]`

Everything else you need we'll do as part of the lab.

🔤 Getting Started
==================

To use the SDK, we need to add a dependency to the Mastodon
Helm chart. Let's pull down the chart so that we can get
started.

```bash
helm pull my-release oci://registry-1.docker.io/bitnamicharts/mastodon
```

Let's also set up our shell for interacting with the Replicated
platform.

```
export REPLICATED_API_TOKEN="[[ Instruqt-Var key="API_TOKEN" hostname="shell" ]]"
```

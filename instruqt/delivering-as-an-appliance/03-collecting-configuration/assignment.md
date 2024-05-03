---
slug: collecting-configuration
id: gx0mrlxjywou
type: challenge
title: Collecting Application Configuration
teaser: |
  Help your users configuration your application with a custom
  configuration screen.
notes:
- type: text
  contents: |
    Help your users configure your application with a simple form
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 300
---

Configuring your application for their own environment can create
a challenge for your customers. Helm values are great, but
documenting a giant values file and distringuishing which values to
change is a challenge. The configuration capabilities of the
Replicated Embedded Cluster let you offer a simple web form to assure
your customers set critical values correctly.

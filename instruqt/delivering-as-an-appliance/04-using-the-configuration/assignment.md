---
slug: using-the-configuration
id: n8juiythxhxt
type: challenge
title: Using the Configuration to Install and Upgrade
teaser: Now we can customize the installation with the configuration
notes:
- type: text
  contents: |
    Provide the user configuration to your Helm chart
tabs:
- title: Shell
  type: terminal
  hostname: node
difficulty: basic
timelimit: 300
---

The configuration screen we built looks great, guides the customer through
their configuration, and helps make sure they set their configuration is set
up correctly. The next step is to configure the application using the options
they provide. Values from the configuration are mapped to the Helm chart(s)
that make up your application using the `HelmChart` resource. The Admin
Console uses this resource to prepare the values passed to Helm when
installing or upgrading the chart.

Passing Configurations to Helm
==============================

We saw the `HelmChart` object when we initially prepared the cluster. In that
section we specified a chart and version to install and explicitly provided no
values to the Helm command (`values: {}`). We're going to fix that now and
provide values based on the configuration the Admin Console collected.

### A Word on Templating

[Templating in the Admin Console](https://docs.replicated.com/reference/template-functions-about)
is a big topic and an in-depth treatment would take it's own lab (or three).
I'm going to explain the basics that we need to use here, building on the
glimpse we had when setting up conditional fields on the configuration screen.
To use the configuration we're going to look at a few more template functions
and how to use them.

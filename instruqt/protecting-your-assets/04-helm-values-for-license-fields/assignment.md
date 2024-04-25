---
slug: helm-values-for-license-fields
id: egqdclosyyl5
type: challenge
title: Checking License Fields in Your Helm Chart
teaser: Using the fields of the Replicated license as values in your Helm chart
notes:
- type: text
  contents: Let's assure a valid license before deploying your Helm chart
tabs:
- title: Shell
  type: terminal
  hostname: shell
  workdir: /home/replicant
- title: Cluster
  type: terminal
  hostname: cluster
difficulty: basic
timelimit: 600
---

When you distribute your software with Replicated, Replicated injects the
license into your Helm chart in two ways:

1. As a value provided to the Replicated SDK to access via an in-cluster API
2. As global values that you can use in other components, including directly in your Helm templates.

We're going to take advantage of the second option to update the Slackernews
chart to only install when the license has not expired.

### A Word of Caution

The approach we're demonstrating here is easily defeated by overriding a Helm
value on the command-line or in a values file. It's meant to remind an honest
customer their license is expired rather than prevent them from tampering with
the license to install anyway. The additional protections provided by the proxy
registry and the Replicated SDK should also be used to assure compliance and
prevent tampering.

License Field Values
--------------------



```
helm registry login [[ Instruqt-Var key="REGISTRY_HOST" hostname="shell" ]]  --username [[ Instruqt-Var key="REGISTRY_USERNAME" hostname="shell" ]]  --password [[ Instruqt-Var key="REGISTRY_PASSWORD" hostname="shell" ]]
helm show values oci://[[ Instruqt-Var key="REGISTRY_HOST" hostname="shell"]]/[[ Instruqt-Var key="REPLICATED_APP" host="shell" ]]/slackernews
```

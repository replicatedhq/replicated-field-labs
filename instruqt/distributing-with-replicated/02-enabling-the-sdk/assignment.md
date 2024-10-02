---
slug: enabling-the-sdk
id: 9weportqrwy6
type: challenge
title: Enabling the Replicated SDK
teaser: Incorporate the SDK into your application
notes:
- type: text
  contents: Introducing the Replicated SDK into your application
tabs:
- id: lep5xpebtpkw
  title: Shell
  type: terminal
  hostname: shell
- id: cnhzcdcy5ots
  title: Manifest Editor
  type: code
  hostname: shell
  path: /home/replicant
difficulty: basic
timelimit: 300
---

Now that we've got our environment set up, let's incorporate the
SDK into our application. Replicated makes it easy for you to do
this by providing a Helm chart you can drop into your chart as a
dependency. When you deliver your Helm chart from the Replicated
registry, we'll embed your customer's license into the final
chart.

This injection serves a few purposes:

1. The license is available to your application logic through a
   call to an API provided by an in-cluster service.
2. Access to your container images and other registry
   assets is secured using customer-specific credentials
3. The in-cluster service can connect securely to the Replicated
   vendor portal for telemetry, upgrade checks, etc.

Adding the Dependency
======================

Go to the the "Manifest Editor" tab and edit the file `Chart.yaml` in
the source directory `slackernews`. You're going to make two changes to
this file.

First, you're going to add a dependency on the Replicated SDK Helm
chart.

```
- name: replicated
  repository: oci://registry.replicated.com/library
  version: [[ Instruqt-Var key="REPLICATED_SDK_VERSION" hostname="shell" ]]
```

You should put the dependency into the array with any other chart dependencies.
In this case, we don't have any dependencies so we need to add the key
`dependencies` to the YAML file. Use the version shown above, since it may be
newer than the one in the screenshot.

![Adding the Dependency](../assets/adding-the-dependency.png)

You should also bump the version number of your chart. Adding
telemetry and preparing to distribute with Replicated feels like
a fairly large change. It's not a breaking change, though, so
let's just bump the minor version number.

```
version: 0.2.0
```

![Bumping the Chart Version](../assets/bumping-the-version.png)

After you make each change, the editor will automatically save
them. Be sure they are saved before you continue.

![Assuring Your Changes are Saved](../assets/saving-your-changes.png)

After saving, drop back in to the "Shell" tab and update your
dependencies.

```shell
helm dependency update slackernews
```

Repackaging Your Chart
=====================

After updating dependencies, you should repackage your Helm
chart into a new tarball including the changes.

```
helm package slackernews --destination ./release
```

You should now have a tarball in directory `release` in your
home directory.

```
ls release
```

which shows

```
slackernews-0.2.0.tgz
```

---
slug: create-release
id: c1pyibvplfog
type: challenge
title: Create Release
teaser: A short description of the challenge.
notes:
- type: text
  contents: Replace this text with your own text
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---
## Set the Application Slug

We will use the Application Slug to set the `REPLICATED_APP` environment varilable used by the `replicated` cli to know which application to update. To access the Application Slug, navigate to **Settings** and copy the value highlighted in red below:


To set the environment variable run the command `export REPLICATED_APP=<your-app-slug>`.

## Set the Api Token

We will use the API Token to set the `REPLICATED_API_TOKEN` environment variable used by the `replicated` cli to authenticate access to the vendor portal. To access the API Token, click on the user profile and click on **Settings**. Scroll down until you see **API Tokens**. Create a new Read/Write token with whatever name you choose.

## Set the Manifests directory

The **Upstream** directory includes the **userdata** sub directory which we don't want to include in our release. Let's create a new directory called **Manifests**

```bash
mk dir Manifests
```
Let' copy the yamls

```bash

cp upstream/*.yaml manifests

```

Create the release by running the following command

```bash
 replicated release create --yaml-dir...

```
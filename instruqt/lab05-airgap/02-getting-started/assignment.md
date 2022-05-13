---
slug: getting-started
id: dwkmxoxunztv
type: challenge
title: Getting Started
teaser: Time to setup the Replicated CLI
notes:
- type: text
  contents: Time to setup the Replicated CLI
tabs:
- title: CLI
  type: terminal
  hostname: cli
difficulty: basic
timelimit: 600
---

## Getting Started

> **Note:** If you've already completed [Lab 0](../lab00-hello-world), you can skip to [Instance Overview](#instance-overview).

You should have received an invite to log into https://vendor.replicated.com -- you'll want to accept this invite and set your password.

Now, you'll need to install the **Replicated CLI** and set up two environment variables to interact with vendor.replicated.com. See [Get Started -> Steps 1 and 2](https://github.com/replicatedhq/kots-field-labs/blob/main/labs/lab00-hello-world/README.md)


`REPLICATED_APP` should be set to the app slug from the Settings page. You should have received your App Name
ahead of time.

![kots-app-slug](../assets/application-slug.png)

`REPLICATED_API_TOKEN` should be set to the previously created user api token. See [Get Started -> Steps 1 and 2](https://github.com/replicatedhq/kots-field-labs/blob/main/labs/lab00-hello-world/README.md)

Once you have the values,
set them in your environment.

```
echo export REPLICATED_APP=... >> ~/.bashrc
echo export REPLICATED_API_TOKEN=... >> ~/.bashrc
```

Lastly before continuing make sure this repo is cloned locally as we will be modifying `lab05` later during the workshop. (should already be done for you)
```bash
git clone https://github.com/replicatedhq/kots-field-labs
cd kots-field-labs/labs/lab05-airgap
```
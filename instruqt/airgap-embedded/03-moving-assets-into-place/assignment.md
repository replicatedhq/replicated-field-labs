---
slug: moving-assets-into-place
id: cqt3rfjrlyp8
type: challenge
title: Moving Assets into Place
teaser: A short description of the challenge.
notes:
- type: text
  contents: Replace this text with your own text
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

***
## Moving Assets into place

If you haven't already, log out of the Air Gap instance with `exit` or ctrl+D.
Our next step is to collect the assets we need for an Air Gap installation:

1. A license with the Air Gap entitlement enabled
2. An Air Gap bundle containing the kURL cluster components
3. An Air Gap bundle containing the application components

(2) and (3) are separate artifacts to cut down on bundle size during upgrade scenarios where only the application version
is changing and no changes are needed to the underlying cluster.


#### Starting the kURL Bundle Download
From your local system run the command below and record the `AIRGAP` section output.

```
replicated channel inspect lab05-airgap
```
<details>
  <summary>Example Output:</summary>

```bash
❯ replicated channel inspect lab05-airgap
ID:             1wyFvAQANNcga1zkRoMIPpQpb9q
NAME:           lab05-airgap
DESCRIPTION:
RELEASE:        1
VERSION:        0.0.1
EXISTING:

    curl -fsSL https://kots.io/install | bash
    kubectl kots install lab05-airgap

EMBEDDED:

    curl -fsSL https://k8s.kurl.sh/lab05-airgap | sudo bash

AIRGAP:

    curl -fSL -o lab05-airgap.tar.gz https://k8s.kurl.sh/bundle/lab05-airgap.tar.gz
    # ... scp or sneakernet lab05-airgap.tar.gz to airgapped machine, then
    tar xvf lab05-airgap.tar.gz
    sudo bash ./install.sh airgap
```

</details>
<br>

Now, let's SSH to our jump box (the one with the public IP) `ssh ${FIRST_NAME}@${JUMP_BOX_IP}` and download the kurl bundle. Replace <URL> with the URL from the `AIRGAP` output that you recorded in the previous step.

```text
curl -o kurlbundle.tar.gz <URL>
```

This will take several minutes, leave this running and proceed to the next step, we'll come back in a few minutes.


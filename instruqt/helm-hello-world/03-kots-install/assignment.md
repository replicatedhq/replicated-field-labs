---
slug: kots-install
id: qxdapcfsfg3o
type: challenge
title: kots-install
teaser: A short description of the challenge.
notes:
- type: text
  contents: Now that we have reviewed how the application is packaged, let's install
    it! In this challenge we will install the Admin Console which will help us manage
    the deployment of the application.
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 600
---
 Install kots
================

In the previous challenge you already copied the installation command for an existing cluster:

```bash
curl https://kots.io/install | bash
kubectl kots install [YOUR-APP-NAME]
```

Run this command in the `Shell` tab. When asked for the `namespace`, you can just press enter and use the one suggested.


The installation will take a couple minutes, and ask to provide a secure password. Remember it as it will be needed in the next Challenge.


Once finished you will see the following output:

<p align="center"><img src="../assets/helm-vm-output.png" width=600></img></p>


Feel free to press `Ctrl+C` and press **Check** to continue to the next challenge.

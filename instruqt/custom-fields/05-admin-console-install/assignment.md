---
slug: admin-console-install
id: beo1zan4ozx3
type: challenge
title: Install Admin Console
teaser: Install the Admin Console, which we'll use to deploy our sample app
notes:
- type: text
  contents: In this challenge we will use a terminal to install the Admin Console
    powered by KOTS
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true  
difficulty: basic
timelimit: 600
---
💡 Install kots
================

In the previous challenge you already copied the installation command for an existing cluster. It looked like:
```bash
curl https://kots.io/install | bash
kubectl kots install [YOUR-APP-NAME]
```

If you don't have the command anymore, you can always go back to the `Vendor` tab and copy it from `Channels > Stable` (the existing cluster install command).

Run this command in the `Shell` tab. If propmted for an installation path, leave it blank to use the default. When asked for the `namespace`, you can just press enter and use the one suggested:

<p align="center"><img src="../assets/lic-namespace.png" width=600></img></p>

The installation will take a couple minutes, and ask to provide a secure password. Remember it as it will be needed in the next Challenge.

<p align="center"><img src="../assets/lic-password.png" width=600></img></p>

Once finished you will see the following output:

<p align="center"><img src="../assets/lic-install-complete.png" width=600></img></p>

Congratulations for finishing this challenge! Click on **Next** to proceed to the next challenge!

---
slug: kots-install
id: 9yyncefqzmrx
type: challenge
title: Kots install
teaser: Start the application installer
notes:
- type: text
  contents: Let's install the application installer
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 300
---

💡 Install kots
================

In the previous challenge you already copied the installation command for an existing cluster:
```bash
curl https://kots.io/install | bash
kubectl kots install [YOUR-APP-NAME]/helloworld
```

Run this command in the `Shell` tab. When asked for the `namespace`, you can just press enter and use the one suggested:

![Namespace](../assets/namespace.png)

The installation will take a couple minutes, and ask to provide a secure password. Remember it as it will be needed in the next Challenge.

![Password](../assets/password.png)

Once finished you will see the following output:

![Kots finished](../assets/kots-install-finished.png)

🏁 Finish
=========

Feel free to press `Ctrl+C` and press **Check** to continue to the next challenge.

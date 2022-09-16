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

Run this command in the `Shell` tab and give it a `namespace` name of your choice:

```
$> kubectl kots install [YOUR-APP-NAME]/helloworld
Enter the namespace to deploy to:
```

When asked for a password, provide a secure password and remember it as it will be needed in the next Challenge.

Once finsihed, press **Check** to continue to the next challenge.

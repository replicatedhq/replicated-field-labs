---
slug: exposing-nginx
id: mrlgykq8y8ed
type: challenge
title: Expose the NGINX service
teaser: Why can't I access nginx
notes:
- type: text
  contents: You've just deployed NGINX. Now let's try browsing it
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
- title: NodePort
  type: service
  hostname: kubernetes-vm
  path: /
  port: 30001
difficulty: basic
timelimit: 1800
---

👋 Introduction
===============

Use the terminal, the browser tab and `kubectl` to understand why you can't access nginx.

🏁 Solution
===========

Use the terminal and `kubectl` to fix the issue.

<details>
  <summary>Hint</summary>

```
k create svc nodeport my-nginx --node-port=30001 --tcp=80 --dry-run=client -o yaml > nginxsvc.yaml
```

And change the `selecter` to use `app: nginx`.

</details>

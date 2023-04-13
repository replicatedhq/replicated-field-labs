---
slug: awkward-scaling
id: rkvhyttuiahw
type: challenge
title: Awkward scaling
teaser: Someone did some weird way of scaling
notes:
- type: text
  contents: Something is wrong with the pods. Try to understand why, and find one
    or multiple ways of fixing it.
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

Use the terminal and `kubectl` to understand why you sometimes can't access nginx.
Use the `check-nginx.sh` script to verify if nginx is running.

<details>
  <summary>The script</summary>

```
#!/bin/bash

echo "Checking nginx"
for i in {1..10}
do
curl -s -o /dev/null -w "Response: %{http_code} %{errormsg}" localhost:30001
echo ""
sleep 2
done
```
</details>

🏁 Solution
===========

Use the terminal and `kubectl` to fix the issue.

<details>
  <summary>Hint</summary>

Someone created a second deployment instead of changing the number of replicas. Use `k scale`? Or edit the deployment?
Also this is a great opportunity to talk about `readinessProbes` and `livenessProbes`.

</details>

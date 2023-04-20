---
slug: demo
id: kjotiq7dpx8l
type: challenge
title: Embedded k8s with kURL
teaser: Embedded k8s
notes:
- type: text
  contents: Get ready to install k8s with kURL
tabs:
- title: Slides
  type: website
  url: https://docs.google.com/presentation/d/e/2PACX-1vS99mmZQQ8IGbTkArn29bAttXMWbabgdu6E3VbVtFKMXQU6TlcDh3hayR7uyd_WLj7Q2yQtLas2YAiO/embed?start=false&loop=false
- title: Vanilla
  type: terminal
  hostname: vanilla
- title: kURL installed
  type: terminal
  hostname: kurl-installed
- title: kURL Multi-Node
  type: terminal
  hostname: kurl-ha-1
difficulty: basic
timelimit: 3600
---

## Notes

### Infrastructure

* The environment consist of the following setup
  * Vanilla: A single node VM allowing to show the single command install
  * kURL Installed: A single node VM with kURL Pre-installed
    * The script is started when the Track is started and runs in the background.
    * Tail the log with: `tail -f kurl.log`
  * kURL Multi-Node: A 3 node HA cluster with kURL installed

### Commands:

* Show all the nodes: `kubectl get nodes`
* Show all the pods: `kubectl get pods -A`
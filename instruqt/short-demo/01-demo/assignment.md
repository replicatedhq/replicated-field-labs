---
slug: demo
id: s0jzecmlrymt
type: challenge
title: Demo
teaser: The demo environment
notes:
- type: video
  url: ../assets/demo.mp4
- type: image
  url: ../assets/connected.png
tabs:
- title: Shell
  type: terminal
  hostname: shell
- title: Code Editor
  type: code
  hostname: shell
  path: /home/replicant/demo-app
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: Application Installer Init
  type: website
  url: http://kotsadm.${_SANDBOX_ID}.instruqt.io:8800
  new_window: true
- title: Application Installed
  type: website
  url: http://application.${_SANDBOX_ID}.instruqt.io:8800
  new_window: true
difficulty: basic
timelimit: 1800
---

Instructions
============

kubectl for kotsadm (default)
```
export KUBECONFIG=~/.kube/config-kotsadm
```

kubectl for application
```
export KUBECONFIG=~/.kube/config-application
```
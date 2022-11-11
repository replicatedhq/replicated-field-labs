---
slug: demo
id: s0jzecmlrymt
type: challenge
title: Demo
teaser: The demo environment
notes:
- type: image
  url: ../assets/slide1.png
- type: image
  url: ../assets/slide2.png
- type: image
  url: ../assets/slide3.png
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
- title: kURL
  type: website
  url: http://kurl.${_SANDBOX_ID}.instruqt.io:8800
  new_window: true
difficulty: basic
timelimit: 3600
---

Kubeconfig
==========

kubectl for kotsadm
```
export KUBECONFIG=~/.kube/config-kotsadm
```

kubectl for application
```
export KUBECONFIG=~/.kube/config-application
```


Remove app
==========

```shell
export KUBECONFIG=~/.kube/config-kotsadm
kubectl kots remove short-demo-${INSTRUQT_PARTICIPANT_ID} -n default --force
```

kURL Embedded Install
=====================

You can show the kURL install output from the embedded installation in the shell tab using
```
tail -100f ~/kurl.log
```

Or if you want the full output from the beginning
```
cat ~/kurl.log | more
```

If you want to install the application on the embedded kURL instance, go to the `kURL` tab, upload the license and enjoy.

Create new release
==================

```shell
replicated release create --yaml-dir ./demo-app/manifests --promote Stable --version 0.2.0
```


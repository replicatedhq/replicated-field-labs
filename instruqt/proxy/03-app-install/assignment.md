---
slug: app-install
id: zfxhfqtljsrb
type: challenge
title: Kots application install with proxy config
teaser: Install kots application via proxy server
notes:
- type: text
  contents: |
    # Install kots application using a proxy server
tabs:
- title: Host
  type: terminal
  hostname: isolated-host
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: KotsAdm
  type: website
  url: http://isolated-host.${_SANDBOX_ID}.instruqt.io:8800
  new_window: true
difficulty: basic
timelimit: 600
---

💡 Shell
=========

Install sample application with proxy config in place.

Access the vendor portal using the credentials output in the Host tab and download the application license
A customer entry has been pre-created called "Hola ProxyCLI Customer", download the license associated with this customer.
Paste license contents into a file called license.yaml

Install using kots cli including license file to upload automatically vs uploading via kotsadm web UI
```
kubectl kots install proxy --skip-preflights --license-file ~/license.yaml --shared-password mytestapp --namespace default --no-port-forward
```

Text access to kotsapp via the KotsAdm tab..


🏁 Finish
==========
Once the nginx application is running and you have reviewed the application, you can Complete the track.

To Finish this track, press **Check**.

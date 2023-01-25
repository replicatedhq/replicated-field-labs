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
timelimit: 4200
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

Log output
==========

Showing the log output for an existing cluster install
```
cat kotsadm.log
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

Support Bundle
==============

The `Application Installed` instance also has a support-bundle pre generated. So you can browse to the application installer and use that one in case you can't wait for the results.

FYI: It takes a couple minutes before it is generated. If you see the following message in `trigger.out`, it should be loaded: `A copy of this support bundle was written to the current directory, named "support-bundle-....tar.gz"`

Demo Examples
==============

Below are different demo examples that can be used as needed:

<details>
  <summary>Default</summary>

1. Vendor Portal: Applications / Channels
2. Vendor Portal: What is a Replicated Application / Release
3. Application Installer: Day 1 Operations - Login -> Vendor Portal Customers to get License
   * Value: License Management / Expiration / Enterprise Feature Flags
4. Application Installer: Day 1 Operations - Customer Config
   * Value: Ease of use
5. Application Installer: Day 1 Operations - Preflights
6. Application Installer: Day 2 Operations - Update
   * Value: Increased adoption rates
   * Additional: Show Reporting in Vendor Portal
   * Additional: Make a change in the editor tab, and create a new release.
7. Application Installer: Day 2 Operations - Support bundles
   * Value: Faster Time To Resolution
   * Additional: Show `sbctl`

</details>

<details>
  <summary>Conferences / 5 minute demo</summary>

You will need to use the "Hot Load" functionality from Instruqt to make sure you always have an environment available. This can allow you to reset your demo environment every hour, without having to wait 5 minutes to create a new environment. Also use the `kubectl kots remove` functionality to reset and be able to redo a clean install.
1. Discovery: What is the main interest for the prospect? Initial Deploy? Updates? Support Bundles? Airgap?
2. Address one or two main interest points.
3. Try scheduling a more in depth demo (20-30 minutes)

</details>

<details>
  <summary>Airgap</summary>
TBD: This environment is not fully airgapped. It is possible to do an airgap deploy with the embedded kURL instance. However, the instance is not airgapped. Nor is there a sandbox that has everything already installed airgapped.

</details>

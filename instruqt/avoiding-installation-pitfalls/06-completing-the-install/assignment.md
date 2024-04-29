---
slug: completing-the-install
id: qo2vdg2dfuq4
type: challenge
title: Completing the Install
teaser: Finishing the install once the cluster passes its preflights
notes:
- type: text
  contents: Completing the installation when cluster capacity is increased
tabs:
- title: Shell
  type: terminal
  hostname: shell
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: Slackernews
  type: service
  hostname: cluster
  port: 30443
  new_window: true
difficulty: basic
timelimit: 600
---

We're going to continue playing the role of the customer
who had a failing preflight check for Slackernews.
Let's also assume that the customer decided to increase
the capacity of the cluster and is ready to perform
the installation.

Validating the Added Capacity
=============================

The first step to resuming the installation process is
to make sure that the cluster has been upgraded with the
necessary resources. We can do that by re-running the
preflight checks against the now expanded cluster.

```
helm template oci://registry.replicated.com/[[ Instruqt-Var key="REPLICATED_APP" hostname="shell" ]]/slackernews \
  | kubectl preflight -
```

In the background, the lab setup process added two
additional nodes to the cluster so that the memory
and CPU capacity has increased. The results of the
checks confirm this change.

![Customer Cluster is Ready for Install](../assets/customer-preflight-checks-after.png)

Installing the Application
==========================

Now that the preflight checks have passed, it is safe to
install the application. You can find the installation
command for the customer "Geeglo" in the Replicated
Vendor Portal. Since we've already run the first two
steps of logging into the registry and running our
preflight checks, we have only the installation
with the Helm command to complete.

You can use the Vendor Portal tab to look up the
install instruction, just like you did in the
previous step before running the preflights.
Click "Customers" in the left navigation, then
click on "Geeglo". You access the instructions
using the "Helm Install Instructions" button on
the top right.

![Customer Installation Commands](../assets/install-instructions.png)

We need to add some additional values that Slackernews
needs to come up correctly. This helps us make sure
the installation will complete.

```
helm install --namespace slackernews --create-namespace  \
    slackernews \
    oci://registry.replicated.com/[[Instruqt-Var key="REPLICATED_APP" hostname="shell" ]]/slackernews \
    --set slackernews.domain=[[ Instruqt-Var key="SLACKERNEWS_DOMAIN" hostname="cluster" ]] \
    --set service.type=NodePort
```

Note that the cluster we're using is fairly limited, so we're using `NodePort`
to simplify access.

Verifying the Installation
==========================

From the customer perspective, the installation is complete when they can log
into the application and see that it was complete. Once their install is
complete, the tab "Slackernews" should show the login page for Slackernews.

![Slackernews Landing Page](../assets/slackernews.png)

🏁 Finish
=========

You've now successfully seen how you can provide preflight
checks to your customer to help them avoid potential pitfalls
installing your application. You've also seen how those
preflights can help your customer understand what they need
to do to prepare for an install, and how they can complete
a successful install once the preflight checks pass.


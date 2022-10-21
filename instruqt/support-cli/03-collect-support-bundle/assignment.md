---
slug: collect-support-bundle
id: em5ownshn4bh
type: challenge
title: Collect Support Bundle
teaser: Collect an application support bundles and diagnose
notes:
- type: text
  contents: Your next challenge is being initialised..
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
- title: KotsAdm
  type: website
  url: http://kubernetes-vm.${_SANDBOX_ID}.instruqt.io:8800
  new_window: true
difficulty: basic
timelimit: 2700
---

👋 Load the application license in kotadm UI
============================================

**In this exercise you will:**

 * Collect a support bundle
 * Investigate application health issues using the cli

***

### 1. Application Status CLI

While the goal of this lab is to show you how to get rich diagnostic information without using granular kubectl CLI commands, we'll pause for a second here to do some very basic inspection of what's happening using kubectl get pod.

```
kubectl get pod
```

Notice the status of the nginx pod is in Init:CrashLoopBackoff status

![supportcli-kubectl-broken1](../assets/supportcli-kubectl-broken1.png)

Rather than diagnosing using kubectl, follow the steps to analyse using the troubleshoot.io tools.


### 2. Support Bundle Collection

In the KotsAdm UI, navigate to the Troubleshoot tab and select *Analyze support-cli track*

![supportcli-kotsadm-supbundle1](../assets/supportcli-kotsadm-supbundle1.png)

This will trigger a support bundle collection via the kotsadm UI

![supportcli-kotsadm-supbundle2](../assets/supportcli-kotsadm-supbundle2.png)

Note the text at the bottom of the screen:

*If you'd prefer, click here to get a command to manually generate a support bundle.*

click on the *click here* link to show the commands to collect a support bundle via the CLI.

![supportcli-kotsadm-supbundle3](../assets/supportcli-kotsadm-supbundle3.png)

This is useful as sometimes the kotsadm UI may be unavailable and diagnosis needs to be performed from the CLI only.

Note that the commands are in two parts.  Firstly install the support-bundle kubernetes cli plugin using krew:
```
curl https://krew.sh/support-bundle | bash
```

Note: If the kubernetes cluster was kURL then the support-bundle plugin would have been pre-installed.

Then the bundle generation command:
```
kubectl support-bundle secret/default/kotsadm-support-cli-${INSTRUQT_PARTICIPANT_ID}-supportbundle --redactors=configmap/default/kotsadm-redact-spec/redact-spec,configmap/default/kotsadm-support-cli-${INSTRUQT_PARTICIPANT_ID}-redact-spec/redact-spec
```

Run both of the above commands now to generate the CLI support bundle.
Note that the support-bundle cli command launches a console ui to view the bundle contents


### 3. CLI Support Bundle Analysis

Note the help output from the support-bundle:

```
kubectl support-bundle --help
```

![supportcli-supbundle-help](../assets/supportcli-supbundle-help.png)

In the support bundle console UI, step through the issues and address them one by one.

Here is some sample output similar to what you should see in your environment:
![supportcli-supbundle1](../assets/supportcli-supbundle1.png)
![supportcli-supbundle4](../assets/supportcli-supbundle4.png)


### 4. Diagnose issues and fix application

Based on the advice presented in the support bundle content, perform corrective action on the cluster.

<details>
  <summary>Open for a hint on config file issue</summary>

Scrolling to the failing check, we can review the error message:
![supportcli-supbundle2](../assets/supportcli-supbundle2.png)

Specifically, you'll see the error message:
```shell
Could not find a file at /etc/support/config.txt with 400 permissions
```

To fix this, run:
```shell
chmod 400 /etc/support/config.txt
```

</details>

<details>
  <summary>Open for a hint on the restraining-bolt issue</summary>

Scrolling to the failing check, we can review the error message:
![supportcli-supbundle3](../assets/supportcli-supbundle3.png)

Specifically, you'll see the error message:
```shell
Restraining bolt in /etc/support has short circuited the startup process. If you remove it, we might be able to launch the application.
```

We can remove this file with

```shell
rm /etc/support/restraining-bolt.txt
```

</details>


### 5. App Status Check - CLI

Check the status of the application services via the cli

```
kubectl get all
```

The nginx pod should be in *Running* state, it can take some time for it to loop after it has backed off so it can be deleted and the replicaset will replace with a new one.
```
kubectl get pod/$(kubectl get pod | grep nginx | awk '{print $1}')
kubectl delete pod/$(kubectl get pod | grep nginx | awk '{print $1}')
watch "kubectl get pod/$(kubectl get pod | grep nginx | awk '{print $1}')"
```

Note: the new pod does take a bit of time to step through the initialisation steps.



### 6. App Status Check - UI

Once the application is confirmed to be running on the CLI, check in the UI

The Application StatusInformers should be Green now:

![supportcli-kotsadm-status-running](../assets/supportcli-kotsadm-status-running.png)

The application launcher should also now be functional, from the Dashboard click on *Open nginx app* link


***

🏁 Finish
=========
Once the nginx application is running and you have reviewed the application, you can Complete the track.

To Finish this track, press **Check**.

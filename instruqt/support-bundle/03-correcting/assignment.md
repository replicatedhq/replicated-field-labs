---
slug: correcting
id: inxjzxmjgxfs
type: challenge
title: Correcting the broken application
teaser: Time to fix the problem
notes:
- type: text
  contents: Time to fix the problem
tabs:
- title: Shell
  type: terminal
  hostname: kubernetes-vm
- title: Application Installer
  type: website
  url: http://kubernetes-vm.${_SANDBOX_ID}.instruqt.io:8800
  new_window: true
difficulty: basic
timelimit: 600
---

🐚 Correcting
===============

In order to correct this issue you'll need to add the missing file.

<details>
  <summary>Expand for shell commands</summary>

```
sudo touch /etc/support/config.txt
sudo chmod 400 /etc/support/config.txt
```
</details>

🏆  Validating
===============

If we run another support bundle, we should now see this check passes:

![check-passes](../assets/check-passes.png)

Once the fix is done, we can wait for the nginx pod to recover from CrashLoopBackoff, or we can give the pod a nudge to get it to retry immediately:

```text
kubectl delete pod -l app=nginx
```

Furthermore, we should now see that the application shows ready in the admin console, and we can open it via the link:

![app-ready](../assets/app-ready.png)

![congrats-page](../assets/congrats-page.png)

Congrats! You've completed the Support Bundle Track!

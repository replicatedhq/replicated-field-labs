Lab 1.1: Using Support Analyzers
=========================================

In this lab, we'll use the Support Bundle analyzers feature to debug an application, modifying the host in order to create the correct conditions for the application to start. In this case, the app is already deployed, but something is not quite right. The Status Informers show "Unavailable".

![lab1-e1-kots-ui-unavailable](img/lab1-e1-kots-ui-unavailable.png)

### Investigating

As you know, the first step when the application isn't starting will be to collect a support bundle. Doing so will run a series of diagnostic checks to help diagnose problems with the application. In the case that a problem cannot be diagnosed automatically, a bundle will be ready for download so you can share with your broader team to help diagnose the issue. Start by navigating to the "Troubleshoot" tab and clicking the "Analzye" button.

<div align="center"><blockquote><h3>If an application isn't starting, always collect a support bundle</h3></blockquote></div>


![click-analzyer](img/click-analyze.png)


Once the bundle is collected, you should see an informative error message in the analyzers:

![failing-check](img/failing-check.png)


> Could not find a file at /etc/lab1-e1/config.txt with 400 permissions -- please ensure this file exists with any content

### Correcting

In order to correct this issue, you'll need to SSH your node and add this file. Assuming your environment is called `train-gmiller`, you'll want to ssh the IP for `train-gmiller-lab1-e1-ui-analyzers`.

<details>
  <summary>Expand for shell commands</summary>

```
sudo touch /etc/lab1-e1/config.txt
sudo chmod 400 /etc/lab1-e1/config.txt
```
</details>

### Validating

Running another support bundle, we should now see this check passes:


![check-passes](img/check-passes.png)

Once the fix is done, we can wait for the nginx pod to recover from CrashLoopBackoff, or we can give the pod a nudge to get it to retry immediately:

```text
kubectl delete pod -l app=nginx
```

Furthermore, we should now see that the application shows ready in the admin console, and we can open it via the link:

![app-ready](img/app-ready.png)

![congrats-page](img/congrats-page.png)

Congrats! You've completed Exercise 1! [Back To Exercise List](../../)

---
slug: troubleshoot-2
id: gzv8orjeqdcg
type: challenge
title: CrashLoopBackOff
teaser: "\U0001F648"
notes:
- type: text
  contents: Time to fix another problem...
tabs:
- title: Workstation
  type: terminal
  hostname: cloud-client
difficulty: intermediate
timelimit: 3600
---
The customer opens another issue, but this time pods seem to be crashing.

Let's investigate our app and see if we can identify the issue. Again, we'll use `sbctl` to explore the support bundle.

To pass this challenge, find the faulty resource, save the YAML spec for that resource to `~/solution.yaml`, correct the problem in the resource, then click "Next" to check your work.

💡 Using `sbctl`
=================

- Remember that you can use the interactive shell prompt with `sbctl shell -s <path-to-support-bundle>`

💡 Using `kubectl`
=================

- How do you make `kubectl` print output in YAML format?
  -- What if you wanted to save that output to a file?

💡 Hints
=================

- How do you list pods?

- How do you describe pods?
  - What if you wanted to see data from multiple pods at once?

- How do you get logs from a pod?
  - What if you wanted to see a previous version of the pod's logs?

- When would you look at `describe` output vs. gathering pod logs?

- Review the [Kubernetes documentation on debugging Pods](https://kubernetes.io/docs/tasks/debug/debug-application/debug-running-pod/)

💡 More Hints
=================

- How do you find the exit code of a Pod?

- What could it mean if a Pod is exiting before it has a chance to emit any logs?

- Review the Linux exit code conventions: `0` means the process exited normally, `1`-`127` generally mean that the process exited because of a crash or error, and >`128` generally means that the process was killed by a signal (think Ctrl-C or the `kill` command).

Troubleshooting Procedure
=================

Identify the problematic Pod from `kubectl get pods -n <namespace>`.  Notice any pods that are not in the Running state.

Describe the current state of the Pod with `kubectl describe pod -n <namespace> <pod-name>`.  Here are some things to look out for:

- each Container's current **State** and **Reason**
- each Container's **Last State** and **Reason**
- the Last State's **Exit Code**
- each Container's **Ready** status
- the **Events** table

For a Pod that is crashing, expect that the current state will be `Waiting`, `Terminated` or `Error`, and the last state will probably also be `Terminated`.  Notice the reason for the termination, and especially notice the exit code.  There are standards for the exit code originally set by the `chroot` standards, but they are not strictly enforced since applications can always set their own exit codes.

In short, if the exit code is >128, then the application exited as a result of Kubernetes killing the Pod.  If that's the case, you'll commonly see code 137 or 143, which is 128 + the value of the `kill` signal sent to the container.

If the exit code is <128, then the application crashed or exited abnormally.  If the exit code is 0, then the application exited normally (most commonly seen in init containers or Jobs/CronJobs)

Look for any Events that may indicate a problem.  Events by default last 1 hour, unless they occur repeatedly.  Events in a repetition loop are especially noteworthy:

```plaintext
Events:
  Type     Reason                  Age                      From     Message
  ----     ------                  ----                     ----     -------
  Warning  BackOff                 2d19h (x9075 over 4d4h)  kubelet  Back-off restarting failed container sentry-workers in pod sentry-worker-696456b57c-twpj7_default(82eb1dde-2987-4f58-af64-883470ffcb58)
```

Another way to get even more information about a pod is to use the `-o yaml` option with `kubectl get pods`.  This will output the entire pod definition in YAML format.  This is useful for debugging issues with the pod definition itself.  Here you will see some info that isn't present in `describe pods`, such as Annotations, Tolerations, restart policy, ports, and volumes.

✔️  Solution
=================

One of the deployments has a memory limit that is too low for the Pod to run successfully.

🛠️ Remediation
=================

Write the YAML spec for the affected deployment into a file at `~/solution.yaml`, then increase the memory limit for the Pod to a reasonable amount.  You may have to make an educated guess about what the correct memory limit should be.

To think about:

- How can we make sure that this doesn't happen again?

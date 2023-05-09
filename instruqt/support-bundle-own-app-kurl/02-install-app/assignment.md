---
slug: install-app
id: jtzlouxvykhg
type: challenge
title: Install app
teaser: Install Your Application using Replicated
notes:
- type: text
  contents: Let's install your Application
tabs:
- title: Cluster Node 1
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-1
- title: Cluster Node 2
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-2
- title: Cluster Node 3
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-3
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: intermediate
timelimit: 3600
---

🚀 Let's begin!

1. Install kURL
=================

# Install the Replicated embedded cluster

**In the ***Cluster Node 1*** tab** begin your embedded cluster installation.  You're already `root` so you don't need to use `sudo`:

We recommend doing this inside a tmux session, so we don't lose the script output if we get disconnected

Example: **for an HA installation (3 primary nodes)**

- `curl -sSL https://kurl.sh/<installer-name> | bash -s ha`

Example: **for an installation with a single primary node**

- `curl -sSL https://kurl.sh/<installer-name> | bash`

Your embedded installer command may have additional [advanced installation options](https://kurl.sh/docs/install-with-kurl/advanced-options).  Double check with your team for the expected options to use.

*When prompted for the loadbalancer IP address, **leave it blank** to use the internal LB*

2. (Optional) Add More Nodes
=================

# If you need to add more nodes to the cluster, do the following, otherwise skip to the "Upload your license" step

**When the install script completes,** copy the primary or secondary node join command printed in green at the end of the installation and run it in the *Cluster Node 2* tab and the *Cluster Node 3* tab.

3. Upload License and Deploy
=================

# Upload your license and install your application

## Navigate to the [Admin Console](http://loadbalancer.[[ Instruqt-Var key="SANDBOX_ID" hostname="cloud-client" ]].instruqt.io:8800)

After installation succeeds, navigate to the **[App Installer Admin Console](http://loadbalancer.[[ Instruqt-Var key="SANDBOX_ID" hostname="cloud-client" ]].instruqt.io:8800)**, login and upload your license.

  ![Application installer](../assets/deploy.png)

In the admin console, continue to configure your application, run preflight checks, and deploy your application.

Once your application is deployed and the admin console reports it is ready to use, check that your application pods are all "Running" before we move on to the interactive troubleshooting exercises.

On "Cluster Node 1": `kubectl get pods -n <your application namespace>`

Click "Check" to continue.

Common Problems
=================

### The connection to the server localhost:8080 was refused - did you specify the right host or port?

Run `bash -l` to reload the shell and try again.

### Reset the Admin Console password

Use the `kots` cli to reset the password in a namespace:

```
kubectl kots reset-password <namespace>
kubectl kots reset-password default
```

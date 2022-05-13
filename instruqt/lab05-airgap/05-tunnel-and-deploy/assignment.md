---
slug: tunnel-and-deploy
id: lpezue9darej
type: challenge
title: Tunnel and Deploy
teaser: A short description of the challenge.
notes:
- type: text
  contents: Setup ssh tunnel and deploy
tabs:
- title: Jumpstation
  type: terminal
  hostname: jumpstation
difficulty: basic
timelimit: 600
---

## Accessing the UI via SSH tunnel, Configuring the instance

You'll want to create a port forward from your jumpbox in order to access to UI locally.

On your jumpbox terminal, run the following. You'll need the `AIRGAP_IP`. You can see it using a `ping airgap`.

```shell
echo $HOSTNAME.$_SANDBOX_ID.instruqt.io # You'll need this output to browse to kotsadm

export AIRGAP_IP=... # your airgap IP

ssh -L 0.0.0.0:8800:$AIRGAP_IP:8800 airgap
```

This will run in the foreground, and you wont see any output. At this point, Kubernetes and the Admin Console are running inside the air gapped server, but the application isn't deployed yet.
To complete the installation, visit http://jumpstation.[SANDBOX_ID].instruqt.io:8800 in your browser.

Click "**Continue and Setup**" in the browser to continue to the secure Admin Console.

![kots-tls-wanring](../assets/kots-tls-warning.png)

Click the "**Skip and continue**" to Accept the insecure certificate in the admin console.
> **Note**: For production installations we recommend uploading a trusted cert and key, but for this tutorial we will proceed with the self-signed cert.

![Console TLS](../assets/admin-console-tls.png)

At the login screen paste in the password noted previously on the `Installation Complete` screen. The password is shown in the output from the installation script.

![Log In](../assets/admin-console-login.png)

Until this point, this server is just running Docker, Kubernetes, and the kotsadm containers.
The next step is to upload a license file so KOTS can validate which application is authorized to be deployed. Use the license file we downloaded earlier.

Click the Upload button and select your `.yaml` file to continue, or drag and drop the license file from a file browser.

![Upload License](../assets/upload-license.png)

After you upload your license, you'll be greeted with an Airgap Upload screen. Select **choose a bundle to upload** and use the "application bundle" that you
downloaded to your workstation using the customer portal here. Click **Upload Air Gap bundle** to continue the upload process.

![airgap-upload](../assets/airgap-upload.png)

You'll see the bundle uploaded and images being pushed to kURL's internal registry. This will take a few minutes to complete.

![airgap-push](../assets/airgap-push.png)

Once uploaded `Preflight Checks` will run. These are designed to ensure this server has the minimum system and software requirements to run the application.
Depending on your YAML in `preflight.yaml`, you may see some of the example preflight checks fail.
If you have failing checks, you can click continue -- the UI will show a warning that will need to be dismissed before you can continue.

![Preflight Checks](https://kots.io/images/guides/kots/preflight.png)


We'll find that the application is unavailable.

![app-down](../assets/app-down.png)

While we'll explore [support techniques for airgapped environments](#collecting-a-cli-support-bundle)
below, in this case you should observe that our deployment is simply not valid, specifically, the
standard nginx entrypoint has been overriden:

```yaml
      containers:
        - name: nginx
          image: nginx
          command:
            - exit
            - "1"
```

So we'll need to create a new release in order to fix this. Continue to the next step.
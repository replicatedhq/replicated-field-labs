---
slug: configuring-the-application
id: igbszqdl9ylb
type: challenge
title: Configuring the Application
teaser: A short description of the challenge.
notes:
- type: text
  contents: Replace this text with your own text
tabs:
- title: Jumpbox
  type: terminal
  hostname: jumpbox
  workdir: /home/replicant
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 600
---

## Accessing the UI via SSH tunnel, Configuring the instance

You'll want to create a port forward from your local workstation in order to access to UI locally.

Run `exit` or click ctrl+D to log out of the air gap instance. Run `exit` or click ctrl+D again to log out of the jump box.

On your local workstation, run the following (again, we'll use `REPLICATED_APP` to construct the DNS name but you can input it manually as well):

```shell
export JUMP_BOX_IP=... # your jumpbox IP
export REPLICATED_APP=... # your app slug
export FIRST_NAME=... # your first name

ssh -NL 30880:${REPLICATED_APP}-lab05-airgap:30880 -L 30888:${REPLICATED_APP}-lab05-airgap:30888 ${FIRST_NAME}@${JUMP_BOX_IP}
```

This will run in the foreground, and you wont see any output. At this point, Kubernetes and the Admin Console are running inside the air gapped server, but the application isn't deployed yet.
To complete the installation, visit http://localhost:30880 in your browser.

Click "**Continue and Setup**" in the browser to continue to the secure Admin Console.

![kots-tls-wanring](assets/kots-tls-warning.png)

Click the "**Skip and continue**" to Accept the insecure certificate in the admin console.
> **Note**: For production installations we recommend uploading a trusted cert and key, but for this tutorial we will proceed with the self-signed cert.

![Console TLS](assets/admin-console-tls.png)

At the login screen paste in the password noted previously on the `Installation Complete` screen. The password is shown in the output from the installation script.

![Log In](assets/admin-console-login.png)

Until this point, this server is just running Docker, Kubernetes, and the kotsadm containers.
The next step is to upload a license file so KOTS can validate which application is authorized to be deployed. Use the license file we downloaded earlier.

Click the Upload button and select your `.yaml` file to continue, or drag and drop the license file from a file browser.

![Upload License](assets/upload-license.png)

After you upload your license, you'll be greeted with an Airgap Upload screen. Select **choose a bundle to upload** and use the "application bundle" that you
downloaded to your workstation using the customer portal here. Click **Upload Air Gap bundle** to continue the upload process.

![airgap-upload](assets/airgap-upload.png)

You'll see the bundle uploaded and images being pushed to kURL's internal registry. This will take a few minutes to complete.

![airgap-push](assets/airgap-push.png)

Once uploaded `Preflight Checks` will run. These are designed to ensure this server has the minimum system and software requirements to run the application.
Depending on your YAML in `preflight.yaml`, you may see some of the example preflight checks fail.
If you have failing checks, you can click continue -- the UI will show a warning that will need to be dismissed before you can continue.

![Preflight Checks](assets/airgap-preflight.png)


We'll find that the application is unavailable.

![app-down](assets/app-down.png)

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

So we'll need to create a new release in order to fix this.


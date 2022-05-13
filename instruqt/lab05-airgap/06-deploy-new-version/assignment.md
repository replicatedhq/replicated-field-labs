---
slug: deploy-new-version
id: zll9xufs3mbq
type: challenge
title: Deploy new version
teaser: A short description of the challenge.
notes:
- type: text
  contents: Deploy a new version
tabs:
- title: CLI
  type: terminal
  hostname: cli
- title: Jumpstation
  type: terminal
  hostname: jumpstation
- title: Editor
  type: code
  hostname: cli
  path: /root/kots-field-labs/labs/lab05-airgap/manifests
difficulty: basic
timelimit: 600
---

## Deploying a new version

From the `labs/lab05-airgap` directory, update the `manifests/nginx-deployment.yaml` file to remove the command override as shown below. (You can do this in the Editor tab)


```diff
diff --git a/labs/lab05-airgap/manifests/nginx-deployment.yaml b/labs/lab05-airgap/manifests/nginx-deployment.yaml
index fa29e8d..3a66405 100644
--- a/labs/lab05-airgap/manifests/nginx-deployment.yaml
+++ b/labs/lab05-airgap/manifests/nginx-deployment.yaml
@@ -16,9 +16,6 @@ spec:
       containers:
         - name: nginx
           image: nginx:latest
-          command:
-            - exit
-            - "1"
           volumeMounts:
             - mountPath: /usr/share/nginx/html
               name: html
```

Once you're satisfied with your `nginx-deployment.yaml` create a new release with `make release`. (Using the CLI terminal tab)
You'll need to ensure you have your `REPLICATED_APP` and `REPLICATED_API_TOKEN` set. See the **Getting Started** section for information on how to obtain and set these.

```shell
make release
```

Once the release is made, you should be able to navigate back to the customer download portal we accessed from the customer page.
Scrolling to the bottom, you can click "show older bundles" to see the history of releases on the lab05-airgap channel.
The new release may take a minute or two to build, so you're want to refresh the make until you see one
with a timestamp that matches when you ran `make release`.

![download-portal-more](../assets/download-portal-more.png)

Once you've downloaded the new version, in the KOTS Admin Console select **Version History** and click "**Upload a new version**" and select your bundle.

![airgap-new-upload](../assets/airgap-new-upload.png)

You'll see the bundle upload as before and you'll have the option to deploy it once the
preflight checks complete. Click **Deploy** to perform the upgrade.

Click the **Application** button to navigate back to the main landing page. The app should now show as **Ready** status on the main dashboard.

In order to access the application select **Open Lab 5**.
> **Note**: For this work successfully you'll need to ensure the SSH tunnel for the app's port (8888) was initialized.

Congrats! You've installed and then upgraded an Air Gap instance!
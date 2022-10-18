---
slug: deploying-a-new-version
id: 3a8ilcdrsyed
type: challenge
title: Deploying a New Version
teaser: A short description of the challenge.
notes:
- type: text
  contents: Replace this text with your own text
tabs:
- title: Jumpbox
  type: terminal
  hostname: jumpbox
  workdir: /home/replicant
- title: Code Editor
  type: code
  hostname: shell
  path: /home/replicant/
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: Application Installer
  type: website
  url: http://jumpbox.${_SANDBOX_ID}.instruqt.io:8800
difficulty: basic
timelimit: 600
---

Now that we've installed into the airgapped environment, let's explore
what your customer will do when you release a new version of your
application. The KUARD application has multiple versions published, so
we're going to change our configuration to deploy the `green` version
rather than the `blue` version we just installed.

### Updating and Releasing

Go to the the "Code Editor" tab and edit the file `deployment.yaml` in
your application manifests.` Where is references the `blue` image for
the application, switch it instead to read `green`.

![Screenshot TO DO]()

After making the change, create a new Replicated release.

```shell
replicated release create ... TO DO
```

### Downloading the Updated Bundle

Once the release is made, you should be able to navigate back to the
customer download portal we accessed from the customer page. Scrolling to 
the bottom, you can click "show older bundles" to see the history of 
releases on the `development` channel. The new release may take a minute 
or two to build, so you're want to refresh the make until you see one
with a timestamp that matches when you created the release. 

![download-portal-more](assets/download-portal-more.png)

### Installing the Update

Once you've downloaded the new version, in the KOTS Admin Console 
select **Version History** and click "**Upload a new version**" and 
select your bundle.

![airgap-new-upload](assets/airgap-new-upload.png)

You'll see the bundle upload as before and you'll have the option to deploy 
it once the preflight checks complete. Click **Deploy** to perform the upgrade.

Click the **Application** button to navigate back to the main landing page. 
The app should now show as **Ready** status on the main dashboard.

Congrats! You've installed and then upgraded an Air Gap instance!


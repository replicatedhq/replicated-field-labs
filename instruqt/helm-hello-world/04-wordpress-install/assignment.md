---
slug: wordpress-install
id: kxqg90q9wucw
type: challenge
title: wordpress-install
teaser: A short description of the challenge.
notes:
- type: text
  contents: |-
    We have installed the Admin Console, now we are ready to deploy Wordpress.

    This challenge will walk through deploying the Wordpress application using the Admin Console.
tabs:
- title: Application Installer
  type: service
  hostname: kubernetes-vm
  port: 8800
  new_window: true
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 600
---
## Step 01

Go to the Application Installer tab, and login using the password you used in the previous step.

<p align="center"><img src="../assets/helm-login.png" width=600></img></p>

## Step 02

Upload the license for the `Helm Customer` you downloaded in Challenge #2

<p align="center"><img src="../assets/helm-license.png" width=600></im></p>

## Step 03

Set the initial blog name in Wordpress in the text field as shown below.

<p align="center"><img src="../assets/helm-config.png" width=600></img></p>

## Step 04


Once you click on `Continue`, the Application Installer will deploy the Wordpress Application.

If you want to check the Wordpress App, wait for the Status Informers to show `Ready` and then click on the `Open Wordpress` link.

<p align="center"><img src="../assets/helm-admin-console.png" width=600></img></p>


It will open a new tab and you should see something similar like the screenshot below:

<p align="center"><img src="../assets/wordpress.png" width=600></img></p>


🏁 Finish
=========

If you've viewed the initial blog in Wordpress, congratulations! You've deployed your first application using the Replicated Application Installer. You can click **Check** to finish this track.

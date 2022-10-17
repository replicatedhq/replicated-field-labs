---
slug: hello-world-install
id: 6qxrucass5to
type: challenge
title: Hello World install
teaser: Install Hello World using Replicated Application Installer
notes:
- type: text
  contents: Let's install the Hello World Application
tabs:
- title: Application Installer
  type: website
  url: http://kubernetes-vm.${_SANDBOX_ID}.instruqt.io:8800
  new_window: true
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 900
---

👋 Install Hello World
===============

## Step 01

Go to the Application Installer tab, and login using the password you used in the previous step.

![Login](../assets/login.png)

If you forgot the password, enter the shell and run `kubectl kots reset-password -n <yournamespace>`

## Step 02

Upload the license for the `Hola Customer` you downloaded in Challenge #2

![Upload License](../assets/upload-license.png)

## Step 03

Customize the hello world application by adding some text examples like below:

![Configuration](../assets/config.png)

For now, you can just ignore the `API token` and `Readonly text` fields.

## Step 04

Click `Continue` and watch the Preflights run. These preflights will validate the application environment.

![Preflights Run](../assets/preflights-run.png)

Once the preflights are finished, you can check the results which will look like below

![Preflights Results](../assets/preflights-results.png)

For now we will ignore the warnings and click `Continue`. As there is an issue with one of the preflights, you will have to confirm that you want to `Deploy and continue`.

![Preflights Deploy](../assets/preflights-deploy.png)

## Step 05

Once you clicked on `Continue`, the Application Installer will deploy the Hello World Application.

![Dashboard](../assets/dashboard.png)

If you want to check the "Hello World" application, you can do so by clicking on `Open App`

![Open App](../assets/open-app.png)

It will open a new tab and you should see something similar like the screenshot below:

![Hola Result](../assets/hola-result.png)

🏁 Finish
=========

If you've viewed the Hello World app, congratulations! You've deployed your first application using the Replicated Application Installer. You can click **Check** to finish this track.

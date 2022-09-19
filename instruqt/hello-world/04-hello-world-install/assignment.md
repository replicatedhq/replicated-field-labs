---
slug: hello-world-install
id: 6qxrucass5to
type: challenge
title: Hello World install
teaser: Install Hello World using Replicated Application Installer
notes:
- type: text
  contents: Let's install the application installer
tabs:
- title: Application Installer
  type: website
  url: http://kubernetes-vm.${_SANDBOX_ID}.instruqt.io:8800
  new_window: true
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 600
---

👋 Install Hello World
===============

## Step 01

Go to the Application Installer tab, and login using the password you used in the previous step.

![Login](../assets/login.png)

## Step 02

Upload the license for the `Hola Customer` you downloaded in Challenge #2

![Upload License](../assets/upload-license.png)

## Step 03

Customize the hello world application by adding some text examples like below:

![Configuration](../assets/config.png)

## Step 04

Click `Continue` and watch the Preflights run. These preflights will validate the application environment.

![Preflights Run](../assets/preflights-run.png)

Once the preflights are finished, you can check the results which will look like below

![Preflights Results](../assets/preflights-results.png)

For now we will ignore the warnings and click `Continue`.

## Step 05

Once you clicked on `Continue`, the Application Installer will deploy the Hello World Application.

![Dashboard](../assets/dashboard.png)

If you want to check the "Hello World", you can do so by clicking on `Open App`


🏁 Finish
=========

If you've viewed the Hello World app, click **Check** to finish this track.

---
slug: wordpress-install
type: challenge
title: wordpress-install
teaser: A short description of the challenge.
notes:
- type: text
  contents: Replace this text with your own text
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---
## Step 01

Go to the Application Installer tab, and login using the password you used in the previous step.

![Login](../assets/login.png)

## Step 02

Upload the license for the `Helm Customer` you downloaded in Challenge #2

![Upload License](../assets/upload-license.png)

## Step 03

Set the initial blog name in Wordpress

![Configuration](../assets/config.png)

## Step 04

Click `Continue` and watch the Preflights run. These preflights will validate the application environment.

![Preflights Run](../assets/preflights-run.png)

Once the preflights are finished, you can check the results which will look like below

![Preflights Results](../assets/preflights-results.png)

For now we will ignore the warnings and click `Continue`.

## Step 05

Once you clicked on `Continue`, the Application Installer will deploy the Wordpress Application.

![Dashboard](../assets/dashboard.png)

If you want to check the Wordpress App, you can do so by clicking on `Open Wordpress`

![Open App](../assets/open-app.png)

It will open a new tab and you should see something similar like the screenshot below:

![Hola Result](../assets/hola-result.png)

🏁 Finish
=========

If you've viewed the initial blog in Wordpress, congratulations! You've deployed your first application using the Replicated Application Installer. You can click **Check** to finish this track.

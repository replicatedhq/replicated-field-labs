---
slug: intro
type: challenge
title: Intro
teaser: Learn how to use Config Options in your Application
notes:
- type: text
  contents: In this track, our application has already been deployed. Before we get
    much further, let's make sure you have access to everything.
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---

## About the Config Custom Resource

This track covers the [KOTS Config Custom Resource](https://docs.replicated.com/reference/custom-resource-config) which allows you to present end users with a Configuration UI. This custom resource provides several [item types](https://docs.replicated.com/reference/custom-resource-config#item-types) that allow you to manage the end user to provide configurations in various formats.

## About this track

In this track we have a simple flask application that displays the databases available on a Postgres instance. This application currently ships with its own `postgres` instance. We are going to use Config Items to provide the end user with the option to connect to their own instance of `postgres`.

We are going to use several tabs throughout, so let's review them and ensure you have access to them:

* **Vendor Portal** We will use this to retrieve our app slug and API token.
* **K3s VM** We have deployed the app to this VM running K3s. This tab also displays the username and password needed to login to the **Vendor Portal** tab.
* **Dev** This is our dev environment where we have already installed all the command line tools you will need. We will set our `REPLICATED_APP` and `REPLICATED_API_TOKEN` here. We will also create our releases from here.
* **Code** This is a text editor provided by instruqt. We will use this editor to add and edit our manifests.
* **Admin Console** This is the admin console that is deployed on the **K3s VM** that we will use to manage updates and use the fields provided by the Config Custom Resources to provide the connection string to our external database.
* **Postgres VM** This is our "external" database to connect our application to. Once we have all our config fields updated, we then point our application to this instance to validate that our application can connect to the external database.


## Checking your environment

Before we get going, let's make sure your sandbox environment is ready to go.

Navigate to the **Admin Console** tab, where you should see the Admin Console login screen. The password is your participant id. To retrieve it, go to the tab and run `echo ...` . Once logged in you should see the Admin Console. We have already deployed a simple Stateful App that uses `postgres`. 

Click on the **Open StatefulApp** link to open the application. It should show you the output of a SQL command to get the available databases of the instance. 






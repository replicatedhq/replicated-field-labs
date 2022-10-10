---
slug: updating-the-application
id: fpxcohlfy6y3
type: challenge
title: Updating the Application
teaser: Update the application to consume the value set for the Custom License Field
notes:
- type: text
  contents: |-
    In this next challenge we are going to get our hands dirty with some YAML!
    We are going to update the application manifests to support turning on/off the Super Duper Feature.
tabs:
- title: Dev
  type: terminal
  hostname: shell
  workdir: /home/replicant
difficulty: basic
timelimit: 600
---

## Setting Up Dev Environment ##

Before we update the application we need to make sure we have a dev environment set up. We have provided a dev environment for you to use but you could also run the `replicated` command line from other Linux or Mac machine like your laptop. To use the dev environment we provided for you click on the **Dev** tab.

Make sure to set the `REPLICATED_APP` and `REPLICATED_API_TOKEN` environment variables. This is covered in the Replicated CLI track in more detail.

We also need to download the manifests we want to update. To download the latest version of the applicaiton, we'll employ the [release download](https://docs.replicated.com/reference/replicated-cli-release-download) command.

First, we need to get the latest sequence number. To do this run the following command

```
replicated release ls

```

You should see results similar to this:

```
SEQUENCE    CREATED                 EDITED                  ACTIVE_CHANNELS
10          2022-10-05T15:32:40Z    0001-01-01T00:00:00Z    CustomFields
9           2022-10-05T13:26:25Z    0001-01-01T00:00:00Z    hellohelm,License100
8           2022-10-04T20:46:30Z    0001-01-01T00:00:00Z
7           2022-10-04T20:33:16Z    0001-01-01T00:00:00Z    support-bundle
6           2022-10-04T20:06:59Z    0001-01-01T00:00:00Z
5           2022-10-04T20:01:59Z    0001-01-01T00:00:00Z
4           2022-10-04T18:57:23Z    0001-01-01T00:00:00Z
3           2022-10-03T18:32:48Z    0001-01-01T00:00:00Z
2           2022-09-28T13:54:29Z    0001-01-01T00:00:00Z
1           2022-09-20T19:53:57Z    0001-01-01T00:00:00Z    hello-world
```

We want to update the release currently on the **CustomFields** channel. Note the **SEQUENCE** associated to that channel as that is what we are going to use in the next command.

Let's create a directory structure before we start dowloading files. Create a directory for this lab in your environment and a `manifests` sub directory to store the manifests.

```
mkdir custom-fields-track

cd custom-fields-track

mkdir manifests

```
To download the contents of the release run the following command

```
replicated release download [The SEQUENCE number from above] -d ./manifests

```

As an exmaple, if I wanted to download the manifests associated to the release in the **CustomFields** example above, I would run:

```
replicated release download 10 -d ./manifests
```

We are going to add a second `ConfigMap` that will be used when the Super Duper Feature is enabled.

Create a new file in the **./manifests** directory called `cp-feature-on.yaml` with the following content

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: nginx-feature-on
data:
  index.html: |
    <!doctype html>
    <html lang="en">
    <head>
      <meta charset="utf-8">
      <title>Custom Fields Track</title>
    </head>
    <body>
      <h2>Congrats!</h2>
      <h2>You have turned on the Super Duper Feature!</h2>
      This is the default NGINX app
    </body>
    </html>
```
Now we are going to update the `nginx-deployment.yaml` file to choose which ConfigMap to use. To do this, we will use sprig in the `volumes` section to determine the `ConfigMap` at run time.

```diff
      volumes:
        - name: html
          configMap:
-           name: nginx-feature-off
+           name: '{{repl if (eq (LicenseFieldValue "enable-feature") "true") }}nginx-feature-on{{repl else}}nginx-feature-off{{repl end}}'
```

The above basically states that if the Custom License Field is set to true, the value of the `name` key is `nginx-feature-on`, otherwise the value will be `nginx-feature-off`

Save changes and create a new release:

```
replicated release create --version [NEW VERSION] --release-notes "Update for Super Duper feature" \
  --promote stable --yaml-dir manifests
```

Let's verify our release we indeed created and promoted to the channel by running `replicated release ls`. There should ne a new SEQUENCE associated to the **CustomFields** channel. If that is the case, you have completed this challenge and ready to move to the next one!
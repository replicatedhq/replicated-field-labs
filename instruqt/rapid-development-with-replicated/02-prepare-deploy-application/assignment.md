---
slug: prepare-deploy-application
id: ewyefebg62cu
type: challenge
title: Prepare to Deploy Sample Application
teaser: Let's retrieve the installation command and license to deploy our sample app
notes:
- type: text
  contents: |-
    In this challenge we'll retrieve from the Vendor Portal our install command and our customer license file.
    we'll also install the Admin Console which we'll use in the next challenge to deploy the sample app.
tabs:
- title: K3S-VM
  type: terminal
  hostname: cluster
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
difficulty: basic
timelimit: 600
---

## Log in to Vendor Portal

If you are still logged in to the Vendor Portal from the previous challenge, continue to the next section. If you are not logged in to Vendor Portal, click on the **Vendor** tab to lauch a new window and log in using the credentials provided in the **Shell** tab.

## Retrieve the Install Command

Navigate to **Channels** if not there already. The release we are going to use has been promoted to the **Stable** channel so let's use that install command. Since we already have a cluster setup, we are going to use the **existing cluster** command highlighted below:

<p align="center"><img src="../assets/rdk-channels.png" width=300></img></p>

## Download the License File

Navigate to **Customers** where we should have our `Dev Customer` there. Click on the download icon highlighted in red below to download the license file.

<p align="center"><img src="../assets/rdk-customers.png" width=600></img></p>

## Install the Admin Console

Head over to the **K3S-VM** tab and paste the install command which should look something like:

```shell
curl https://kots.io/install | bash
kubectl kots sample-app-{$INSTRUQT_PARTICIPANT_ID}
```

Use the default namespace and set the password, keeping in mind that you will need it throughout the lab.

<p align="center"><img src="../assets/rdk-output.png" width=600></img></p>

Once you see the prompt to browse to the admin console as shown above, you have completed this challenge.


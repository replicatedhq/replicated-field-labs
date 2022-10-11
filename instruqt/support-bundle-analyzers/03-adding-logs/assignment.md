---
slug: adding-logs
type: challenge
title: adding-logs
teaser: Create a Support Bundle
notes:
- type: text
  contents: Create a Support Bundle and add Application logs
tabs:
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
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


### Create a Support Bundle

Create a support bundle in kotsadm

Everything looks ok. Also not much to see in the files.

### Add application logs

The application consists of an nginx deployment. Let's add a logs collector to th `kots-support-bundle.yaml`, create a new release and promote it to the `Stable` channel. In the Application installer, deploy the new version, and generate another support bundle.

```
- logs:
    name: nginx
    selector:
      - app=nginx
    containerNames:
      - nginx
      - k8slove
```

Still nothing to see in the analyzers, but let's go to the file inspector to see if our logs collector can tell us something. The `nginx.log` is not so special. But the `k8slove.log` is a bit more special and contains `Artist yes`. That doesn't sound correct, as the Artist for the song "California love" should be `2Pac`. (Told you it is a bit of a contrived example).


### Add Analyzer

Let's add an Analyzer to the support bundle that shows an error if the `k8slove.log` does not contain `Artist 2Pac`, and helps the end user to understand what they should to fix the application.

In the Vendor Portal, create a new release and update the `kots-support-bundle.yaml` to contain the following:
```
apiVersion: troubleshoot.sh/v1beta2
kind: SupportBundle
metadata:
  name: support
spec:
  collectors:
    - clusterInfo: {}
    - clusterResources: {}
    - logs:
        name: nginx
        selector:
          - app=nginx
        containerNames:
          - nginx
          - k8slove
  analyzers:
    - textAnalyze:
        checkName: k8slove
        fileName: nginx/**/k8slove.log
        regex: 'Artist 2Pac'
        outcomes:
          - pass:
              when: "true"
              message: "Congratulations, the artist behind California love is 2Pac! Or not?"
          - fail:
              when: "false"
              message: "You should update the Config, and put `2Pac` in `k8slove Artist`."
```

The above analyzer, will search for any text `Artist 2Pac` and fail if not found.

Save and promote the release to the Stable channel. Update the app in the Application Installer, and generate a new Support Bundle.

![Analyzer error](../assets/analyzer-error.png)

Great! So now the end-user will also know they have to update the config option. Go ahead, update the config, and deploy the new version. Once deployed, you can refresh the deployed app and should see the following output:

![Working nginx](../assets/working-nginx.png)

If you go back to the application installer, and generate another support bundle, the Analyzer view should now show an additional test that passses.

![Analyzer pass](../assets/analyzer-pass.png)
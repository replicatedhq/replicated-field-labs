---
slug: kotsadm-ui
id: bisphrzpeuon
type: challenge
title: Admin Console UI
teaser: Access the Admin Console UI and view pre-flight check output
notes:
- type: text
  contents: Navigate the Admin Console UI and complete app install
tabs:
- title: Admin Console
  type: service
  hostname: kubernetes-vm
  port: 8800
  new_window: true
- title: Vendor
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: Shell
  type: terminal
  hostname: kubernetes-vm
difficulty: basic
timelimit: 600
---

👋 Kotsadm UI
=============

**In this exercise you will:**

 * Access Admin Console, login and upload license
 * View the initial run of the preflights
 * After the app is running view the preflight log for current deployment


### 1. Finish App Install

In the previous challenge the kots services were installed for the application on the existing kubernetes cluster.

In this challenge the installation is completed and as part of that the application install pre-flight checks will be run.

Launch the Admin Console console UI using the **Open external window** launch button in the KotsAdm tab


### 2. Admin Console authentication and License upload

![preflight-login](../assets/preflight-login.png)

Login to the Admin Console UI using the password you set in the previous challenge.
After authenticating you will be prompted to upload an application license file, select the file that you downloaded earlier for the **Hola Preflight Customer**.

![preflight-login](../assets/preflight-license-upload1.png)
![preflight-login](../assets/preflight-license-upload2.png)


### 3. Application Configuration Values

Once the license is uploaded you will be presented with the application configuration set in (kots-config.yaml), this data entry is made before the application is deployed for the first time.

![preflight-login](../assets/preflight-config.png)

You can add values for the config as you like, or accept details and click **Continue**


### 4. Pre-Flight Checks!!

Next the Application Pre-Flights checks will be run for the first time:

![preflight-login](../assets/preflight-preflight-checks1.png)
![preflight-login](../assets/preflight-preflight-results1.png)

Notice that the Total CPU Cores check hit the **warn** outcome

Validating the environment before doing the actual installation is a very important step in the installation and upgrade process. In this case, we're trying to install the application on a single node cluster that has 2 cpu cores. In Challenge 2 we already showed the `kots-preflight.yaml` that defines all the different preflight checks. Below is the specific analyzer that checks the number of cpu cores for this demo application:


```yaml
   - nodeResources:
     checkName: Total CPU Cores
     outcomes:
      - fail:
          when: "sum(cpuCapacity) < 2"
          message: The cluster must contain at least 2 cores, and should contain at least 4 cores.
      - warn:
          when: "sum(cpuCapacity) < 4"
          message: The cluster should contain at least 4 cores.
      - pass:
          message: There are at least 4 cores in the cluster.
```

In the above snippet, we make use of the `nodeResources` analyzer that gives us access to the `cpuCapacity`. In practice, you should always make sure you have the right Collectors and Analyzers defined for your Application preflight checks. At least they should check your Application minimum system requirements.

### 5. Application Deployment

When satisfied with the Preflight Check results, click **Continue** to carry on with the deployment.

![preflight-login](../assets/preflight-deploy.png)

Click **Deploy and continue** and the application resource status informer will turn to <font color="Green">Green</font> after a few moments

![preflight-login](../assets/preflight-app-deployed.png)

If you click on Details, you'll see the detailed status informer result.

![preflight-login](../assets/preflight-app-running.png)


### 6. View PreFlights post deployment

The Application Preflight results can be viewed post application deployment by selecting the PreFlights icon

![preflight-login](../assets/preflight-preflight-icon.png)
![preflight-login](../assets/preflight-preflight-checks2.png)

Note that the preflights can be run at any time from this screen


🏁 Finish
=========

Congratulations you have finished the Applications Pre-Flights Lab!!  You can click **Check** to finish this track.


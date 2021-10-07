Lab 1.1: Using Support Analyzers
=========================================

In this lab, we'll use the Support Bundle analyzers feature to debug an application, modifying the host in order to create the correct conditions for the application to start.

* **What you will do**:
    * Learn to query, read, and understand support bundle analyzers
    * Use the analyzers to fix a problem on the server and get the application up and running
* **Who this is for**: This lab is for anyone who will build KOTS applications **plus** anyone who will be user-facing
    * Full Stack / DevOps / Product Engineers
    * Support Engineers
    * Implementation / Field Engineers
    * Success / Sales Engineers
* **Prerequisites**:
    * Basic working knowledge of Linux (SSH, Bash)
* **Outcomes**:
    * You will be ready to use KOTS's support bundle feature to diagnose first-line issues in end-user environments
    * You will reduce escalations and expedite time to remediate for such issues

### Ground Rules

In this lab and most of those that follow it, some of the failure scenarios are not entirely realistic, but are used to highlight how to troubleshoot issues.
It is very possible to reverse-engineer the solution by reading the Kubernetes YAML instead of following the lab steps.
If you want to get the most of out these labs, use the presented debugging steps to get experience with the toolset.

### Accessing the UI

1. Navigate to the KOTS admin console. Use your instance's IP address.

    ```
    https://$IP_ADDRESS:8800
    ```

    > Note: Ensure you have [configured /etc/hosts](../../doc/01-architecture.md#terraform), so that you can access it.

1. Enter the password to access the admin console. The password to your instance UI will be provided as part of the lab, or you can reset by SSHing the node and running

    ```shell
    kubectl kots reset-password -n default
    ```

### The Issue

In this case, the app is already deployed, but something is not quite right.
The Status Informers show "Unavailable".


![lab01-kots-ui-unavailable](img/lab1-kots-ui-unavailable.png)

### Investigation

The first step in troubleshooting is to collect a support bundle. Doing so will run a series of diagnostic checks to help diagnose problems with the application. In the case that a problem cannot be diagnosed automatically, a bundle will be ready for download so you can share with your broader team to help diagnose the issue.

<div align="center"><blockquote><h3>If an application isn't starting, always collect a support bundle</h3></blockquote></div>

1. Navigate to the "Troubleshoot" tab

1. Click the "Analzye" button.

    ![click-analzyer](img/click-analyze.png)

    Once the bundle is collected, you should see an informative error message in the analyzers:

    ![failing-check](img/failing-check.png)

### Issue Correction

1. SSH into your `lab01-ui-analyzers` node

1. Create the following file as `/etc/lab1/config.txt`

    ```bash
    export FIRST_NAME="<your first name>"

    ssh ${FIRST_NAME}@<server ip address>
    ```

    <details>
      <summary>Expand for shell commands</summary>

    ```
    sudo touch /etc/lab1/config.txt
    sudo chmod 400 /etc/lab1/config.txt
    ```
    </details>

### Validation

1. Check the Nginx pod's status.
    ```bash
    kubectl get pod -l app=nginx
    ```

    > Note: We can either wait for the pod to recover from `CrashLoopBackoff`, or we can move things along by deleing the pod using `kubectl delete pod -l app=nginx`

1. Generate another support bundle (see instructions above). We should now see that the issue now passes:

    ![check-passes](img/check-passes.png)

1. Click the "Application" tab in the admin console to view the application's status. The application should now show as ready in the admin console.

1. Click the "Open Lab 1 Exercise 1" button to view the running application.

    ![app-ready](img/app-ready.png)

The application should present a page like this.

![congrats-page](img/congrats-page.png)

Congrats! You've completed Exercise 1!

[Back To Exercise List](https://github.com/replicatedhq/kots-field-labs/tree/main/labs)

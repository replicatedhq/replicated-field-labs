Lab 1.8: Airgap
=========================================

In this lab, you will review how to perform installations and collect support bundles in air gap environments.

* **What you will do**:
    * Access and verify a single node air gap setup using a bastion server
    * Learn to use app manager to install in an air gap environment
    * Create an SSH tunnel to configure an air gap instance
    * Perform an upgrade of an application in an air gap environment
    * Use the `kubectl support-bundle` CLI in an air gap environment
* **Who this lab is for**:
    * Anyone who builds or maintains app manager applications (see note below)
    * Full Stack / DevOps / Product Engineers
* **Prerequisites**:
    * Basic working knowledge of Kubernetes
* **Outcomes**:
    * You will be ready to deliver a app manager application into an air gap environment
    * You will build confidence in performing upgrades and troubleshooting in air gap environments

## Overview

Starting with a bare air gap server, so you will practice performing an air gap install from scratch.

Then you will explore how some of the support techniques differ between online and air gap environments.

***
## Airgap Workflow Overview

First, you will push a release -- in the background, the Replicated air gap builder prepares an air gap bundle.

![airgap-slide-1](img/airgap-slide-1.png)

Next, you will collect a license file, a download link, and a public kURL bundle.

![airgap-slide-2](img/airgap-slide-2.png)

Then you will move all three artifacts into the datacenter using a jump box.

![airgap-slide-3](img/airgap-slide-3.png)

The previous diagram shows a three node cluster, but this lab uses only a single node. While the app manager bundle is moved onto the server using SCP, as shown in the diagram, the app bundle and license file are uploaded using a browser UI through an SSH tunnel.

***
## Getting Started


1. Check your email for an invitation to log into https://vendor.replicated.com. Accept this invitation and set your password.

2. From the **Settings** page, copy the Application Slug. You will need that later to set the `REPLICATED_APP`.

  ![kots-app-slug](img/application-slug.png)

***
## Instance Overview

A separate email was sent to you containing the IP of a jump box and the name of an air gap server that is assigned specifically to you. For example:

```text
dx411-dex-lab08-airgap-lite-jump = 104.155.131.205
dx411-dex-lab08-airgap-lite
```

In general, the name of the private server will be the same as the jump box, with the characters `-jump` removed from the suffix. Put another way, you could construct both instance names programatically as:

```shell
${REPLICATED_APP}-lab08-airgap-lite-jump
```

and

```shell
${REPLICATED_APP}-lab08-airgap-lite
```

### Connecting to the air gap server

1. Set your application slug, the public IP of your jump box, and your first name:

  ```shell
  $ export JUMP_BOX_IP=<ip_address>
  $ export REPLICATED_APP=<application_slug>
  $ export FIRST_NAME=<your_first_name> # use lower case
  ```
  These will run in the foreground, and you will not see any output.

2. Perform an SSH command into the air gap server:

  ```shell
  $ ssh -J ${FIRST_NAME}@${JUMP_BOX_IP} ${FIRST_NAME}@${REPLICATED_APP}-lab08-airgap-lite
  ```

  The `-J` option, allows connections to the target host by first making a ssh connection to the jump host (`${JUMP_BOX_IP}`) described by destination and then establishing a TCP forwarding to the ultimate destination (`${REPLICATED_APP}-lab08-airgap-lite`) from there.

  Alternatively, you can also do the same thing using multiple steps:

  ```shell
  local> ssh ${FIRST_NAME}@${JUMP_BOX_IP}
  jump> export REPLICATED_APP=<application_slug>
  jump> ssh ${REPLICATED_APP}-lab08-airgap-lite
  ```

3. After you are on the air gap server, verify that the server does not have internet access using the following command:

  ```shell
  curl -v https://kubernetes.io
  ```

  This command will fail if there is no internet connection, which means that your environment is safe.

4. Press **ctrl+C**.

5. Log out of the air gap instance with `exit` or press **ctrl+D**.

***
## Moving assets into place

Our next step is to collect the assets we need for an air gap installation:

* An air gap bundle containing the kURL cluster components
* An air gap bundle containing the application components
* A license with the air gap entitlement enabled

The air gap bundles are separate artifacts used to cut down on bundle size during upgrade scenarios where only the application version is changing and no changes are needed to the underlying cluster.


#### Starting the kURL bundle download

1. From your home directory, run the SSH command to your jump box that uses the public IP:

  ```text
  ssh ${FIRST_NAME}@${JUMP_BOX_IP}
  ```
2. Enter the password.
3. Download the kURL bundle.

##### Enabling air gap for a customer

1. Log in to the vendor web at link:http://endor.replicated.com.

2. From the **Customers** page, select the **Airgap Download Enabled** checkbox:

  ![enable-airgap](./img/airgap-customer-enable.png)

3. Click **Save changes**.

  Air gap is enabled for the `lab8` customer.

##### Download air gap assets
1. From the **Customers** page, scroll down to the Download Portal section.

  ![download-portal](img/airgap-customer-portal.png)

2. Generate a new password and save it in your notes.
3. Right-click on **View download portal** and copy the URL to your notes.

  This is a link you would usually send to your customer, so from here on we'll be wearing our "end user" hat.

4. Click the **View download portal**.
  The **Download portal** window opens.  

5. Select **Embedded cluster**.

6. Under Latest kURL embedded install, select **Download bundle**.

  ![download-portal](img/download-portal-kurl.png)

6. In your terminal, from your jump box, run the following curl command using the URL that you copied from the download portal:

  Example:

  ```text
  kots@dx411-dex-lab08-airgap-lite-jump ~$ curl -o kurlbundle.tar.gz <URL>
  ```

  The download takes several minutes. While this runs, proceed to building an air gap release. You will return to this downlad later.

#### Building an Airgap Release

By default, only the Stable and Beta channels will automatically build air gap bundles. For other channels, you have two options to create an air gap release:

- Manually build from the Release History view for a channel
- Set a channel to automatically build all promoted releases

> **Note:** For a production application, air gap releases are built automatically on the Stable channel, but that is not
be necessary for this lab.

In this lab, we are working off the `lab08-airgap-lite` channel, so you will enable air gap builds on that channel.

1. From the vendor web UI, check the build status by selecting **Channels** > **lab08-airgap-lite** > **Release History**.

  ![release-history](img/channel-release-history.png)

2. Although you can build individual bundles on the Release History page, for this lab we recommend that you select the edit icon for the channel and enable **build all releases for this channel**.

  ![edit-channel](img/channel-edit-info-btn.png)

  ![auto-build](img/channel-enable-airgap.png)

3. To see all of the bundles building or built, select **Release history** from the Latest release pane.

  ![airgap-built](img/airgap-builds.png)

  The Release history page displays the build history and status.

#### Download Airgap Assets

Air gap assets are downloaded from the same download portal UI where you downloaded the kURL bundle.

1. From the download portal, select **Embedded cluster**.

2. Download the license.

3. Download the `Latest Lab 1.8: Airgap Bundle` to your workstation.

  > **Note:** Do not download the kURL bundle because that is the download you already started on the server.

  ![download-portal-view](img/download-portal-view.png)

4. To check the download of the kURL bundle, run the SSH command to your jump box that has the public IP. Replace the URL with the one you copied above.

  Example:

  ```text
  ssh kots@<jump box IP address>
  ```

  At the beginning of the lab, you downloaded the bundle with this command from your jump box:

  Example:

  ```text
  kots@dx411-dex-lab08-airgap-lite-jump ~$ curl -o kurlbundle.tar.gz <URL>
  ```

5. When the download is finished, copy the download to the air gap serverusing the following scp command. You can use the DNS name in this case, as described in [Instance Overview](#instance-overview).

  Example:

  ```text
  kots@dx411-dex-lab08-airgap-lite-jump ~$ scp kurlbundle.tar.gz ${REPLICATED_APP}-lab08-airgap-lite:~

  ```

  > **Note**: -- we use SCP via an SSH tunnel in this case, but the air gap methods in this lab also extend to
  more locked down environments where physical media is required to move assets into the datacenter.

6. Run the SSH command to the air gap node. If you still have a shell on your jump box, you can use the jump instance name.

  Example:

  ```text
  kots@dx411-dex-lab08-airgap-lite-jump ~$ ssh ${REPLICATED_APP}-lab08-airgap-lite
  ```

  Alternatively, you can use the following SSH command:

  ```shell
  ssh -J ${FIRST_NAME}@${JUMP_BOX_IP} ${FIRST_NAME}@${REPLICATED_APP}-lab08-airgap-lite
  ```

7. From the air gap node, untar the bundle and run the installation script using the `airgap` flag. kURL install flags are documented [in the kurl.sh docs](https://kurl.sh/docs/install-with-kurl/advanced-options).

  ```shell
  tar xvf kurlbundle.tar.gz
  sudo bash install.sh airgap
```

  An `Installation Complete` message appears when the action is complete.

  ![kurl-password](img/kurl-password.png)

8. Use `exit` or press **ctl+D** to exit the air gap server.

  Since the instance is air gapped, in the next steps you will create a port forward to access the UI from your workstation.
***
## Accessing the UI and configuring the instance

1. Create a port forward using SSH tunnel from your workstation to access the UI locally.

  > **Note:** In this lab, we will use `REPLICATED_APP` to construct the DNS name, but you can input it manually as well.

  ```shell
  $ export FIRST_NAME=... # your firstname (lowercase)
  $ export JUMP_BOX_IP=... # your jumpbox IP
  $ export REPLICATED_APP=... # your app slug
  $ ssh -NL 8800:${REPLICATED_APP}-lab08-airgap-lite:8800 -L 8888:${REPLICATED_APP}-lab08-airgap-lite:8888 ${FIRST_NAME}@${JUMP_BOX_IP}
  ```

  These commands run in the foreground, and you will not see any output. At this point, Kubernetes and the admin console are running inside the air gapped server, but the application is not deployed yet.

2. To complete the installation, visit http://localhost:8800 in your browser.

3. Click "**Continue and Setup**" in the browser to continue to the secure admin console.

  ![kots-tls-wanring](img/kots-tls-warning.png)

4. Click **Skip and continue** to accept the insecure certificate in the admin console.

  > **Note**: For production installations, we recommend uploading a trusted certificate and key. For this tutorial, we will proceed with the self-signed cert.

  ![Console TLS](img/admin-console-tls.png)

5. At the login screen, paste in the password noted previously on the `Installation Complete` screen. The password is shown in the output from the installation script.

  ![Log In](img/admin-console-login.png)

  Until this point, this server is running Kubernetes and the kotsadm containers only.

6. Upload the license file that you downloaded earlier, so that app manager can validate which application is authorized to be deployed:

  1. Click **Upload** and select your `.yaml` file, or drag and drop the license file from a file browser.

    ![Upload License](img/upload-license.png)

    The Airgap Upload page opens.

  2. Select **choose a bundle to upload** and use the application bundle that you previously downloaded to your workstation from the customer portal. Click **Upload Air Gap bundle** to continue the upload process.

    ![airgap-upload](img/airgap-upload.png)

    You will see the bundle uploaded and images being pushed to kURL's internal registry. This process takes a few minutes to complete.

    ![airgap-push](img/airgap-push.png)

7. After the license is uploaded, `Preflight Checks` run automatically. These are designed to ensure this server has the minimum system and software requirements to run the application.

  Depending on your YAML in `preflight.yaml`, you may see some of the example preflight checks fail. If you have failing checks, you can click **Continue**. and then dismiss the warning that appears.

  ![Preflight Checks](https://kots.io/images/guides/kots/preflight.png)

  You see that the application is unavailable.

  ![app-down](img/app-down.png)

  In this case, the deployment is not valid, specifically, the
  standard `nginx` entry point has been overridden:

  ```yaml
      containers:
        - name: nginx
          image: nginx
          command:
            - exit
            - "1"
    ```

    To fix this, we will deploy a new release in the next steps.

    > **Note:** We will explore [support techniques for airgapped environments](#collecting-a-cli-support-bundle) later in this lab.
***
## Deploying a new version

As part of the lab setup, a new release has been created in the vendor portal with the fix.

1. To make the release available, go to **Releases** > **Sequence 2**, and click **Promote**. Select the `lab08-airgap-lite` channel, to which the new release will be promoted.

  ![app-down](img/promote-sequence-2.png)

2. (Optional) You can review the difference between the two releases in the vendor portal. It is also shown below:

  ```diff
  diff --git a/labs/lab05-airgap/manifests/nginx-deployment.yaml b/labs/lab05-airgap/manifests/nginx-deployment.yaml
  index fa29e8d..3a66405 100644
  --- a/labs/lab05-airgap/manifests/nginx-deployment.yaml
  +++ b/labs/lab05-airgap/manifests/nginx-deployment.yaml
  @@ -16,9 +16,6 @@ spec:
         containers:
           - name: nginx
             image: nginx:latest
  -          command:
  -            - exit
  -            - "1"
             volumeMounts:
               - mountPath: /usr/share/nginx/html
                 name: html
  ```

3. When the promotion is done, navigate back to the customer download portal that you accessed from the customer page.
Scroll down and click **show older bundles** to see the history of releases on the lab08-airgap channel.

  The new release may take a minute or two to build, so refresh the page until you see a release with a timestamp that matches the time that you promoted the release.

  ![download-portal-more](img/download-portal-more.png)

4. From the admin console, select **Version History** and click "**Upload a new version**" and select your bundle.

  ![airgap-new-upload](img/airgap-new-upload.png)

  The new bundle uploads and the preflight checks run.

5. Click **Deploy** to perform the upgrade.

6. Click **Application** to return to the main landing page. The app will show with the **Ready** status on the main dashboard.

7. To access the application, select **Open Lab 8**.

  > **Note**: For this to work successfully, ensure that the SSH tunnel for the app's port (8888) was initialized.

Congratulations! You have installed and upgraded an air gap instance!

***
## Collecting a CLI support bundle

As a final step, let's review how to collect support bundles.

What would you do in the case that the app installation itself was failing?

1. Try loading the `kots.io` support bundle from the air gap server.

  ```shell
  kubectl support-bundle https://kots.io
  ```

  As you might expect, this fails because we cannot fetch the spec from the internet. You receive an error message similar to the following message:

  ```text
  Error: failed to load collector spec: failed to get spec from URL: execute request: Get "https://kots.io": dial tcp 104.21.18.220:443: i/o timeout
  ```

2. Pull in the spec from https://github.com/replicatedhq/kots/blob/master/support-bundle.yaml.

  How you get this file onto the server is up to you -- expand the following details for an option that uses `cat` with a heredoc.


<details>

```shell
cat <<EOF > support-bundle.yaml
apiVersion: troubleshoot.sh/v1beta2
kind: SupportBundle
metadata:
  name: collector-sample
spec:
  collectors:
    - clusterInfo: {}
    - clusterResources: {}
    - ceph: {}
    - exec:
        args:
          - "-U"
          - kotsadm
        collectorName: kotsadm-postgres-db
        command:
          - pg_dump
        containerName: kotsadm-postgres
        name: kots/admin_console
        selector:
          - app=kotsadm-postgres
        timeout: 10s
    - exec:
        args:
          - "http://localhost:3030/goroutines"
        collectorName: kotsadm-goroutines
        command:
          - curl
        containerName: kotsadm
        name: kots/admin_console
        selector:
          - app=kotsadm
        timeout: 10s
    - exec:
        args:
          - "http://localhost:3030/goroutines"
        collectorName: kotsadm-operator-goroutines
        command:
          - curl
        containerName: kotsadm-operator
        name: kots/admin_console
        selector:
          - app=kotsadm-operator
        timeout: 10s
    - logs:
        collectorName: kotsadm-postgres-db
        name: kots/admin_console
        selector:
          - app=kotsadm-postgres
    - logs:
        collectorName: kotsadm-api
        name: kots/admin_console
        selector:
          - app=kotsadm-api
    - logs:
        collectorName: kotsadm-operator
        name: kots/admin_console
        selector:
          - app=kotsadm-operator
    - logs:
        collectorName: kotsadm
        name: kots/admin_console
        selector:
          - app=kotsadm
    - logs:
        collectorName: kurl-proxy-kotsadm
        name: kots/admin_console
        selector:
          - app=kurl-proxy-kotsadm
    - logs:
        collectorName: kotsadm-dex
        name: kots/admin_console
        selector:
          - app=kotsadm-dex
    - logs:
        collectorName: kotsadm-fs-minio
        name: kots/admin_console
        selector:
          - app=kotsadm-fs-minio
    - logs:
        collectorName: kotsadm-s3-ops
        name: kots/admin_console
        selector:
          - app=kotsadm-s3-ops
    - secret:
        collectorName: kotsadm-replicated-registry
        includeValue: false
        key: .dockerconfigjson
        name: kotsadm-replicated-registry
    - logs:
        collectorName: rook-ceph-agent
        selector:
          - app=rook-ceph-agent
        namespace: rook-ceph
        name: kots/rook
    - logs:
        collectorName: rook-ceph-mgr
        selector:
          - app=rook-ceph-mgr
        namespace: rook-ceph
        name: kots/rook
    - logs:
        collectorName: rook-ceph-mon
        selector:
          - app=rook-ceph-mon
        namespace: rook-ceph
        name: kots/rook
    - logs:
        collectorName: rook-ceph-operator
        selector:
          - app=rook-ceph-operator
        namespace: rook-ceph
        name: kots/rook
    - logs:
        collectorName: rook-ceph-osd
        selector:
          - app=rook-ceph-osd
        namespace: rook-ceph
        name: kots/rook
    - logs:
        collectorName: rook-ceph-osd-prepare
        selector:
          - app=rook-ceph-osd-prepare
        namespace: rook-ceph
        name: kots/rook
    - logs:
        collectorName: rook-ceph-rgw
        selector:
          - app=rook-ceph-rgw
        namespace: rook-ceph
        name: kots/rook
    - logs:
        collectorName: rook-discover
        selector:
          - app=rook-discover
        namespace: rook-ceph
        name: kots/rook
    - exec:
        collectorName: weave-status
        command:
        - /home/weave/weave
        args:
        - --local
        - status
        containerName: weave
        exclude: ""
        name: kots/kurl/weave
        namespace: kube-system
        selector:
        - name=weave-net
        timeout: 10s
    - exec:
        collectorName: weave-report
        command:
        - /home/weave/weave
        args:
        - --local
        - report
        containerName: weave
        exclude: ""
        name: kots/kurl/weave
        namespace: kube-system
        selector:
        - name=weave-net
        timeout: 10s

  analyzers:
    - textAnalyze:
        checkName: Weave Status
        exclude: ""
        fileName: kots/kurl/weave/kube-system/weave-net-*/weave-status-stdout.txt
        outcomes:
        - fail:
            message: Weave is not ready
        - pass:
            message: Weave is ready
        regex: 'Status: ready'
    - textAnalyze:
        checkName: Weave Report
        exclude: ""
        fileName: kots/kurl/weave/kube-system/weave-net-*/weave-report-stdout.txt
        outcomes:
        - fail:
            message: Weave is not ready
        - pass:
            message: Weave is ready
        regex: '"Ready": true'
EOF
```
</details>

3. Run the following command to collect a bundle as usual.

  ```shell
  kubectl support-bundle ./support-bundle.yaml
  ```

Congratulations! You ave completed Exercise 8! [Back To Exercise List](https://github.com/replicatedhq/kots-field-labs/tree/main/labs)


**Additional resources:**

There is an in-depth post with some other options at [How can i generate a support bundle if i cannot access the admin console?](https://help.replicated.com/community/t/kots-how-can-i-generate-a-support-bundle-if-i-cannot-access-the-admin-console/455).



***
## Extra exercises

If you finish the lab early, you can:

1. Experiment with copying the CLI-generated bundle off the server and uploading it to https://vendor.replicated.com

2. Experiment with expanding or building your own `support-bundle.yaml` and using it to collect other information about the host

Lab 1.20: GitOps & Sealed Secrets
=========================================

Kubernetes makes it easy to describe the state of your system in manifests, and it's
very convenient to be able to check those into a version control system or share them
with your community, except for Secrets.

Discover how KOTS can utilize the "Sealed Secrets" features of the [bitnami project](https://github.com/bitnami-labs/sealed-secrets#installation) to ensure that Secrets
are encrypted and safe to store, even in a public repository!

* **What you will do**:
    * Learn how to deploy your application in GitOps mode for a `git`-driven deployment
    * Learn to configure KOTS to automatically encrypt Secrets to make them safe to share, even publicly
    * Learn how `kubeseal` can be used from your workstation to achieve the same
* **Who this is for**: This lab is for anyone who will deploy KOTS applications, in particular if you may also be using the GitOps features of KOTS
    * Full Stack / DevOps / Product Engineers
    * Support Engineers
    * Implementation / Field Engineers
    * Success / Sales Engineers
* **Prerequisites**:
    * Basic working knowledge of Linux (SSH, bash)
    * A cluster that has an application installed with KOTS
    * A workstation (which may be the same as the cluster control plane node)
    * some Kubernetes Secrets you want to protect
    * a `git` repo hosted on a supported provider, such as GitHub, GitLab, or Bitbucket, and an associated Deploy key
* **Outcomes**:
    * You will be able to effectively utilize Sealed Secrets to prevent passwords, certificates, etc. from leaking out of your Kubernetes manifests, and build confidence in your automation workflow using `git`

### Deploy an application and enable GitOps mode

1.



### Prepare the workstation to use Sealed Secrets

#### 1. Install `helm` and, optionally, `kubeseal`

- If using a Mac as your workstation, install with `homebrew`

    ```bash
    brew install helm
    brew install kubeseal
    ```

- If using a Linux workstation, install from binaries

    ```bash
    curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

    # navigate to https://github.com/bitnami-labs/sealed-secrets/releases and grab the latest
    # release for your architecture, unpack it, and put it in your $PATH
    wget https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.17.5/kubeseal-0.17.5-linux-amd64.tar.gz
    tar xvzf kubeseal-0.17.5-linux-arm64.tar.gz
    chmod +x kubeseal
    sudo mv kubeseal /usr/local/bin/kubeseal
    ```
#### 2.

#### 3. Prepare the cluster

1. Install the `sealed-secrets` CRDs, controller, etc. from the [bitnami project](https://github.com/bitnami-labs/sealed-secrets#installation) in the `kube-system` namespace

    ```bash
    helm repo add sealed-secrets https://bitnami-labs.github.io/sealed-secrets
    helm repo update
    helm install -n kube-system sealed-secrets sealed-secrets/sealed-secrets
    ```

1. Obtain the public key certificate from the `sealed-secrets` controller

    - with `kubeseal`
    ```bash
    kubeseal \
      --controller-name=sealed-secrets \
      --controller-namespace=kube-system \
      --fetch-cert > sealed-secrets-cert.pem
    ```

    - or copy the certificate from the `sealed-secrets` pod logs
    and create a text file called `sealed-secrets-cert.yaml`
    ```bash
    kubectl logs -n kube-system sealed-secrets-7684c7b86c-6bhhw
    # 2022/04/20 15:49:49 Starting sealed-secrets controller version: 0.17.5
    # controller version: 0.17.5
    # 2022/04/20 15:49:49 Searching for existing private keys
    # 2022/04/20 15:49:58 New key written to kube-system/sealed-secrets-keyxmwv2
    # 2022/04/20 15:49:58 Certificate is
    # -----BEGIN CERTIFICATE-----
    # MIIEzDCCArSgAwIBAgIQIkCjUuODpQV7zK44IB3O9TANBgkqhkiG9w0BAQsFADAA
    # ...
    # jCwIzOs3BKuiotGAWACaURFiKhyY+WiEpsIN1H6hswAwY0lcV1rrOeQgg9rfYvoN
    # 0tXH/eHuyzyHdWt0BX6LLY4cqP2rP5QyP117Vt2i1jY=
    # -----END CERTIFICATE-----
    ```

1. Create a Kubernetes Secret in the same namespace as KOTS that will hold the public key from `sealed-secrets` controller, and KOTS will use it to transform Secrets into SealedSecrets
    
    ```bash
    kubectl create secret generic sealed-secrets-cert \
      --dry-run=client \
      -o yaml \
      --from-file=cert.pem=sealed-secrets-cert.pem \
      -n sentry-pro > sealed-secrets-cert.yaml
    ```
1. Edit the `sealed-secrets-cert.yaml` and add the following labels to the `.metadata.labels` field so KOTS knows it must use this certificate
    - `kots.io/buildphase: secret`
    - `kots.io/secrettype: sealedsecrets`
    
    Example:
    ```yaml
    ---
    apiVersion: v1
    kind: Secret
    metadata:
      name: sealed-secrets-cert
      namespace: sentry-pro
      labels:
        kots.io/buildphase: secret
        kots.io/secrettype: sealedsecrets
    data:
      cert.pem: "..."
    ```

1. Apply the manifest to the cluster in the same namespace as KOTS

    ```bash
    kubectl apply -n sentry-pro -f sealed-secrets-cert.yaml
    # secret/sealed-secrets-cert created
    ```

1. Enable GitOps and allow KOTS to push a release to the repo.  Observe that `Secrets` have been replaced with `SealedSecrets`

1. Apply the finished manifest to the cluster and observe that the `SealedSecrets` have been decrypted by the SealedSecrets controller

### Support Bundle

#### The Issue

In this case, the app is already deployed, but something is not quite right.
The Status Informers show "Unavailable".


![lab10-demo-status-unavailable](img/lab10-demo-status-unavailable.png)

#### Investigating

As you know, the first step when the application isn't starting will be to collect a support bundle. Doing so will run a series of diagnostic checks to help diagnose problems with the application. In the case that a problem cannot be diagnosed automatically, a bundle will be ready for download so you can share with your broader team to help diagnose the issue. Start by navigating to the "Troubleshoot" tab and clicking the "Analyze" button.

<div align="center"><blockquote><h3>If an application isn't starting, always collect a support bundle</h3></blockquote></div>

![click-analzyer](img/click-analyze.png)


Once the bundle is collected, you should see an informative error message in the analyzers:

![failing-check](img/failing-check.png)


> Could not find a file at /etc/lab10/config.txt with 400 permissions -- please ensure this file exists with any content

#### Correcting

In order to correct this issue, you'll need to SSH into your `lab10-demo` node and add this file. 

```bash
export FIRST_NAME=... # your first name

ssh ${FIRST_NAME}@<server ip address>
```

<details>
  <summary>Expand for shell commands</summary>

```
sudo touch /etc/lab10/config.txt
sudo chmod 400 /etc/lab10/config.txt
```
</details>

#### Validating

Running another support bundle, we should now see this check passes:


![check-passes](img/check-passes.png)

Once the fix is done, we can wait for the nginx pod to recover from CrashLoopBackoff, or we can give the pod a nudge to get it to retry immediately:

```text
kubectl delete pod -l app=nginx
```

Furthermore, we should now see that the application shows ready in the admin console, and we can open it via the link:

![app-ready](img/app-ready.png)

![congrats-page](img/congrats-page.png)

Congrats! You've completed Exercise 10! [Back To Exercise List](https://github.com/replicatedhq/kots-field-labs/tree/main/labs)

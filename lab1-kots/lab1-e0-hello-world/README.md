Lab 1.0: Hello World
=========================================

This exercise is designed to give you a sandbox to ensure you have the basic CLI tools set up and are prepared to proceed 
with the exercises in Lab 1. 

The README and the YAML sources draw from https://github.com/replicatedhq/replicated-starter-kots

### Get started

To start, you'll want to clone this repo somewhere. Optionally, you can fork it first (or you can do this later).

```shell script
git clone git@github.com:replicatedhq/kots-field-labs
cd kots-field-labs/lab1-kots/lab1-e0-hello-world
```

#### Install CLI

### 1. Install CLI

To start, you'll want to install the `replicated` CLI.
You can install with [homebrew](https://brew.sh) or grab the latest Linux or macOS version from [the replicatedhq/replicated releases page](https://github.com/replicatedhq/replicated/releases).

##### Brew

```shell script
brew install replicatedhq/replicated/cli
```

##### Manual

```shell script
curl -s https://api.github.com/repos/replicatedhq/replicated/releases/latest \
           | grep "browser_download_url.*$(uname | tr '[:upper:]' '[:lower:]')_amd64.tar.gz" \
           | cut -d : -f 2,3 \
           | tr -d \" \
           | cat <( echo -n "url") - \
           | curl -fsSL -K- \
           | tar xvz replicated
```
Then move `./replicated` to somewhere in your `PATH`:


```shell script
mv replicated /usr/local/bin/
```

##### Verifying

You can verify it's installed with `replicated version`:

```text
$ replicated version
```
```json
{
  "version": "0.31.0",
  "git": "c67210a",
  "buildTime": "2020-09-03T18:31:11Z",
  "go": {
      "version": "go1.14.7",
      "compiler": "gc",
      "os": "darwin",
      "arch": "amd64"
  }
}
```


#### Configure environment

You should have received an invite to log into https://vendor.replicated.com -- you'll want to accept this invite and set your password.

Once you're logged in, you'll need to set up two environment variables to interact with vendor.replicated.com:

```
export REPLICATED_APP=...
export REPLICATED_API_TOKEN=...
```

`REPLICATED_APP` should be set to the app slug from the Settings page. You should have received your App Name
ahead of time.

<p align="center"><img src="https://kots.io/images/guides/kots/cli-setup-quickstart-settings.png" width=600></img></p>

Next, create an API token from the [Teams and Tokens](https://vendor.replicated.com/team/tokens) page:

<p align="center"><img src="https://kots.io/images/guides/kots/cli-setup-api-token.png" width=600></img></p>

Ensure the token has "Write" access or you'll be unable create new releases. Once you have the values,
set them in your environment.

```
export REPLICATED_APP=...
export REPLICATED_API_TOKEN=...
```

You can ensure this is working with

```
replicated release ls
```

#### Verifying manifests

You should have a few YAML files in `manifests`:


```text
$ ls -la manifests
.rw-r--r--  114 dex 16 Aug  7:59 config-map.yaml
.rw-r--r--  906 dex 16 Aug  7:59 config.yaml
.rw-r--r--  608 dex 16 Aug  7:59 deployment.yaml
.rw-r--r--  296 dex 16 Aug  7:59 ingress.yaml
.rw-r--r-- 2.5k dex 16 Aug  7:59 preflight.yaml
.rw-r--r--  399 dex 16 Aug  7:59 replicated-app.yaml
.rw-r--r--  205 dex 16 Aug  7:59 service.yaml
.rw-r--r--  355 dex 16 Aug  7:59 support-bundle.yaml
```

You can verify this yaml with `replicated release lint`:

```shell script
replicated release lint --yaml-dir=manifests
```

If there are no errors, you'll an empty list and a zero exit code.

```text
RULE    TYPE    FILENAME    LINE    MESSAGE
```

* * *

### Creating our first release


Now that we have some YAML, let's create a release and promote it to the `Unstable` channel so we can test it internally.
You can inspect the `Makefile` to get a sense of what is happening under the hood, but for now, for simplicity we'll use the Makefile command,
for this and all future labs in this program.


```shell script
make release
```

You can verify the release was created with `release ls`:

```text
$ replicated release ls
SEQUENCE    CREATED                      EDITED                  ACTIVE_CHANNELS
1           2020-09-03T11:48:45-07:00    0001-01-01T00:00:00Z    Unstable
```

* * *

### Creating a Customer License

Now that we've created a release, we can create a "customer" object.
A customer represents a single licensed end user of your application.

In this example, we'll create a customer named `Some-Big-Bank` with an expiration in 10 days.
Since we created our release on the `Unstable` channel, we'll assign the customer to this channel.

```shell script
replicated customer create \
  --name "Some-Big-Bank" \
  --expires-in "240h" \
  --channel "Unstable"
```

Your output should look something like this:

```text
ID                             NAME             CHANNELS     EXPIRES                          TYPE
1h0yojS7MmpAUcZk8ekt7gn0M4q    Some-Big-Bank     Unstable    2020-09-13 19:48:00 +0000 UTC    dev
```

You can also verify this with `replicated customer ls`.

```text
replicated customer ls
```

Now that we have a customer, we can download a license file

```shell script
replicated customer download-license \
  --customer "Some-Big-Bank"
```

You'll notice this just dumps the license to stdout, so you'll probably want to redirect the output to a file:

```shell script
export LICENSE_FILE=~/Some-Big-Bank-${REPLICATED_APP}-license.yaml
replicated customer download-license --customer "Some-Big-Bank" > "${LICENSE_FILE}"
```

You can verify the license was written properly with `cat` or `head`:

```text
$ head ${LICENSE_FILE}

apiVersion: kots.io/v1beta1
kind: License
metadata:
  name: some-big-bank
spec:
  appSlug: kots-dex
  channelName: Unstable
  customerName: Some-Big-Bank
  endpoint: https://replicated.app
```
 * * *

### 6. Getting an install command

Next, let's get the install commands for the Unstable channel with `channel inspect`:

```text
$ replicated channel inspect Unstable
ID:             VEr0nhJBBUdaWpPaOIK-SOryKZEwa3Mg
NAME:           Unstable
DESCRIPTION:
RELEASE:        1
VERSION:        Unstable-ba710e5
EXISTING:

    curl -fsSL https://kots.io/install | bash
    kubectl kots install cli-quickstart-puma/unstable

EMBEDDED:

    curl -fsSL https://k8s.kurl.sh/cli-quickstart-puma-unstable | sudo bash

AIRGAP:

    curl -fSL -o cli-quickstart-puma-unstable.tar.gz https://k8s.kurl.sh/bundle/cli-quickstart-puma-unstable.tar.gz
    # ... scp or sneakernet cli-quickstart-puma-unstable.tar.gz to airgapped machine, then
    tar xvf cli-quickstart-puma-unstable.tar.gz
    sudo bash ./install.sh airgap
```

* * *

### 7. Installing KOTS

A server has already been provisioned for this exercise, you'll want to find the one with the name matching `lab1-e0-hello-world`. 
KOTS has not yet been installed on this server, to give you an opportunity to experiment with the install process.

###### On the Server

Next, ssh into the server we just created, and run the install script from above, using the `EMBEDDED` version:

```shell
curl -sSL https://k8s.kurl.sh/<your-app-name-and-channel> | sudo bash
```

This script will install Docker, Kubernetes, and the KOTS admin console containers (kotsadm).

Installation should take about 5-10 minutes.

Once the installation script is completed, it will show the URL you can connect to in order to continue the installation:

```text

Kotsadm: http://[ip-address]:8800
Login with password (will not be shown again): [password]


To access the cluster with kubectl, reload your shell:

    bash -l

The UIs of Prometheus, Grafana and Alertmanager have been exposed on NodePorts 30900, 30902 and 30903 respectively.

To access Grafana use the generated user:password of admin:[password] .

To add worker nodes to this installation, run the following script on your other nodes
    curl -sSL https://kurl.sh/cli-quickstart-puma-unstable/join.sh | sudo bash -s kubernetes-master-address=[ip-address]:6443 kubeadm-token=[token] kubeadm-token-ca-hash=sha256:[sha] kubernetes-version=1.16.4 docker-registry-ip=[ip-address]

```

Following the instructions on the screen, you can reload the shell and `kubectl` will now work:

```bash
user@kots-guide:~$ kubectl get pods
NAME                                  READY   STATUS      RESTARTS   AGE
kotsadm-585579b884-v4s8m              1/1     Running     0          4m47s
kotsadm-migrations                    0/1     Completed   2          4m47s
kotsadm-operator-fd9d5d5d7-8rrqg      1/1     Running     0          4m47s
kotsadm-postgres-0                    1/1     Running     0          4m47s
kurl-proxy-kotsadm-77c59cddc5-qs5bm   1/1     Running     0          4m46s
user@kots-guide:~$
```

* * *

### 8. Install the Application

At this point, Kubernetes and the Admin Console are running, but the application isn't deployed yet.
To complete the installation, visit the URL that the installation script displays when completed.

Once you've bypassed the insecure certificate warning, you have the option of uploading a trusted cert and key.
For production installations we recommend using a trusted cert, but for this tutorial we'll click the "skip this step" button to proceed with the self-signed cert.

![Console TLS](https://kots.io/images/guides/kots/admin-console-tls.png)

Next, you'll be asked for a password -- you'll want to grab the password from the CLI output above and use it to log in to the console.

![Log In](https://kots.io/images/guides/kots/admin-console-login.png)

Until this point, this server is just running Docker, Kubernetes, and the kotsadm containers.
The next step is to upload a license file so KOTS can pull containers and run your application.
Click the Upload button and select your `.yaml` file to continue, or drag and drop the license file from a file browser. 

![Upload License](https://kots.io/images/guides/kots/upload-license.png)

Preflight checks are designed to ensure this server has the minimum system and software requirements to run the application.
Depending on your YAML in `preflight.yaml`, you may see some of the example preflight checks fail.
If you have failing checks, you can click continue -- the UI will show a warning that will need to be dismissed before you can continue.

![Preflight Checks](https://kots.io/images/guides/kots/preflight.png)


You should see the app Deployed and Ready.
If you are still connected to this server over ssh, `kubectl get pods` will now show the example nginx service we just deployed.

![Cluster](https://kots.io/images/guides/kots/application.png)

### View the application

Since we used the default nginx application and enabled the ingress object, we can view the application at `http://${INSTANCE_IP}/` with no port, and you should see a basic (perhaps familiar) nginx server running:

![Cluster](https://kots.io/images/guides/kots/example-nginx.png)

Next, we'll walk through creating and delivering an update to the application we just installed.

* * *

### 9. Iterating

From our local repo, we can update the nginx deployment to test a simple update to the application.
We'll add a line to `deployment.yaml`, right after `spec:`. The line to add is

```yaml
  replicas: 2
```

Using `head` to view the first 10 lines of the file should give the output below

```shell script
head manifests/deployment.yaml
```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-nginx
  labels:
    app: example
    component: nginx
spec:
  replicas: 2
  selector:
```

Once you've added the `replicas` line, you can create a new release:

```shell script
make release
```

### Update the Test Server

To install and test this new release, we need to connect to the Admin Console dashboard on port :8800 using a web browser.
At this point, it will likely show that our test application is "Up To Date" and that "No Updates" Are Available.
The Admin Console can be configured to check for new updates at regular intervals but for now we'll trigger a check manually by clicking "Check for Updates".
You should see a new release in the history now.
You can click the +/- diff numbers to review the diff, but for now let's click "Deploy" to roll out this new version.

![View Update](https://kots.io/images/guides/kots/view-update.png)

Clicking the Deploy button will apply the new YAML which will change the number of nginx replicas, this should only take a few seconds.
You can verify this on the server by running

```shell script
kubectl get pod -l app=nginx
```

You should see two pods running.

* * *

### Next Steps

From here, it's time to start iterating on your application.
Continue making changes and using `make release` to publish them.


If you want to learn more about KOTS features, you can explore some of the [intermediate and advanced guides](/vendor/guides), some good next steps might be

- [Integrating your release workflow with CI](/vendor/guides/ci-cd-integration)
- [Integrating a Helm Chart](/vendor/guides/helm-chart)

If you already have a release published in https://vendor.replicated.com you'd like to use as a starting point, check out the help docs for `replicated release download`:

```shell script
replicated release download --help
```

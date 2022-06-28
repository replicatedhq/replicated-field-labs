Lab 18: Using Helm Charts
=========================================

In this lab, you will learn how to include a Helm Chart in your Replicated Application. For this lab we wil use a simpler HelmChart, but the concepts covered here should help you understand how to handle more complex and larger charts. 

In this lab, we'll cover the following:

* Packaging a Helm Chart from source
* Adding a Helm Chart to a Replicated Application Release
* Map config fields to fields in the Values file
* Deploy the Helm Chart using the Replicated Application Manager
* Cover how to manage private images
* Cover how to manage multiple HelmCharts in a single Replicated Application Release

## Prerequisites

While not an absolute requirement, this lab is written with *nix systems in mind, not Windows. If you are following this lab on a Windows machine and find gaps, please let us know.

You will need to have the following installed:

* Lastest version of Helm - In the writing and testing of the lab, we used version 3.8.2
* Git

You also should have received an invite to join Replicated. Make sure you accepted and activated your account.

## Packaging Helm Chart

For this lab we are going to use the WordPress [Helm Chart](https://github.com/bitnami/charts/tree/master/bitnami/wordpress) available from Bitnami. The Helm Chart includes a couple of sub charts (common and MariaDB) as dependencies.

### 0. Get Dev Environment ready

We will be running some commands on the terminal and will download some files. Let's create a directory where we'll do our work

```bash
   $ mkdir repl-helm-lab
   $ cd repl-helm-lab

```
Next, we are going to download the nescessary files to package our Helm Chart.

### 1. Clone the Bitnami Repo

The repository we are about to clone contains multiple Helm Charts. While this lab will focus on one chart, you could follow the same steps covered in this lab with other charts.

```bash
git clone https://github.com/bitnami/charts.git
```

This command will now download all the helm chart files. In order to include a Helm Chart in Replicated, we have to package it first.

### 2. Change Directory to the Wordpress Chart

As mentioned above, this repo contains several Helm Charts, and we want to only package WordPress, so we need to change directory:

```bash
cd charts/bitnami/wordpress/
```

We are now at the root of the Wordpress Helm Chart. If you list the contents of this directory, you will see that it contains a `Chart.yaml` and `Values.yaml` file. The `Chart.yaml` contains details about the chart, and the `Values.yaml` contains fields that the end user can override with their own values. Fields in this file can also be passed as parameters when you run `helm install`.

### 3. Package the Helm Chart

To package the chart, we are going to run the `helm package` command.

As mentioned above, this Helm Chart includes some dependencies. To ensure that these are included in this Chart, include the `-u` option. Also note that we are passing a `.` for the location of the chart. This is where the top level ```Chart.yaml``` file is located. For more details on this command, please refer to the [Helm Documentation](https://helm.sh/docs/helm/helm_package/#helm).

```bash
helm package -u .
```
The command may display some warnings, but ultimately should display an output similar to

```bash
Successfully packaged chart and saved it to: /.../charts/bitnami/wordpress/wordpress-14.0.7.tgz

```

We will include this file in our Replicated Application Release.

## Adding the Helm Chart to a Release

You should have received an invite to log into https://vendor.replicated.com -- you'll want to accept this invite and set your password. 

In this lab, we have already created an application for you, but there are no releases created yet. We are going to cover creating the release with the UI and through the CLI. Going through the UI is easier and better if you just want to understand the basic concepts as it will automatically create resources for you. With the CLI, we'll take a more 'starting from scratch' approach. Later in the lab we'll be making changes to the application, and will only cover the CLI.

### Using the UI

* Log in to the Vendor Portal using your email and password.
* On the navigation menu, select **Releases**

    <p align="center"><img src="img/releases.png" alt="vendor-ui" width="1000" margin=auto/></img></p>

* Click on **Create Release**

    <p align="center"><img src="img/create-release.png" alt="create-release" width="1000" margin=auto/></img></p>

* We will not use the default nginx app, so let's delete the following files:
    * example-configmap.yaml
    * example-deployment.yaml
    * example-ingress.yaml
    * example-service.yaml

    <p align="center"><img src="img/delete-files.png" alt="files-to-delete" width="1000" margin=auto/></img></p>

* Now we are ready to add the Chart! Open a window and browse to the directory where the Helm Chart we created is located.
* Drag and Drop it where the Replicated Release Files navigator as shown below
    <p align="center"><img src="img/drag-drop-chart.png" alt="drag and drop chart" width="1000" margin=auto/></img></p>

* You will be prompted to select which method to use when deploying the Helm Chart:
   
    <p align="center"><img src="img/helm-opts.png" alt="helm options" width="600" margin=auto/></img></p>

   * **Native Helm** - This should be the default option and should be the one selected. With this option, App Manager will use Helm to deploy the Helm Chart in the customer's environment.
   * **Replicated** - This option will eventually be depecrated and should not be selected. With this option, App Manager will first generate Kubernetes Manifests from the Helm Chart (i.e., run `helm template`) before they are deployed to the cluster as manifests.

* Note how a new file is created by Replicated that includes details about the chart we just added.

<p align="center"><img src="img/helm-file.png" alt="helm file" width="600" margin=auto/></img></p>

### Using the CLI

For the CLI, you should have already completed the [Hello World](https://github.com/replicatedhq/kots-field-labs/blob/main/labs/lab00-hello-world/README.md) Lab, as it covers how to set up your local Dev environment with the Replicated CLI.

Make sure to update your environment variables to interact with this application. See [Get Started -> Steps 1 and 2](https://github.com/replicatedhq/kots-field-labs/blob/main/labs/lab00-hello-world/README.md)


`REPLICATED_APP` should be set to the app slug from the Settings page. You should have received your App Name
ahead of time.

<p align="center"><img src="img/application-slug.png" alt="Application Slug" width="600" margin=auto/></img></p>


`REPLICATED_API_TOKEN` should be set to the previously created user api token. See [Get Started -> Steps 1 and 2](https://github.com/replicatedhq/kots-field-labs/blob/main/labs/lab00-hello-world/README.md)

Once you have the values, set them in your environment. If you are still in the `charts` directory, let's get back to the root of the directory we created earlier

```bash
    $ cd ../../..
```
Now to set your environment variables:

```bash
    $ export REPLICATED_APP=...
    $ export REPLICATED_API_TOKEN=...
```

We need a directory that will contain the Helm Chart and the Replicated manifests. Let's create a directory for our application, and in there a `manifests` directory that will contain the manifests and helm chart.

```bash
   $ mkdir helm-wordpress
   $ cd helm-wordpress
   $ mkdir manifests
```

We need to move the helm chart package we created previously to the manifests directory we just created. Depending on where your helm chart was created, the path in the example command below may be different.

```bash
    $ mv ../charts/bitnami/wordpress/wordpress-14.2.2.tgz manifests
```

Next, let's create a couple of manifests to finish our first release. Both of these manifests are Replicated [Custom Resources](https://docs.replicated.com/reference/custom-resource-about).

#### Adding kots-app.yaml

The first Custom Resource we'll create is [Application](https://docs.replicated.com/reference/custom-resource-application). This manifest includes details about the application we created.

``` yaml
#kots-app.yaml
apiVersion: kots.io/v1beta1
kind: Application
metadata:
  name: helm-wordpress
spec:
  title: Helm Wordpress
  icon: https://upload.wikimedia.org/wikipedia/commons/9/93/Wordpress_Blue_logo.png
  statusInformers:
    - deployment/wordpress

```

#### Adding wordpress.yaml

The second Custom Resource we'll create is [Helm Chart](https://docs.replicated.com/reference/custom-resource-helmchart). This file has information about the Wordpress Helm Chart. This file is specific to the Helm Chart, so if your application contains multiple, separate charts, you will need a manifest for each Helm Chart. You do not need to define a `Helm Chart` manifest for any chart included as a sub-chart.

```yaml
#wordpress.yaml
apiVersion: kots.io/v1beta1
kind: HelmChart
metadata:
  name: wordpress
spec:
  # chart identifies a matching chart from a .tgz
  chart:
    name: wordpress
    chartVersion: "14.0.6"
  # helmVersion identifies the Helm Version used to render the chart. Default is v2.
  helmVersion: v3

  # useHelmInstall identifies whether this Helm chart will use the
  # Replicated Helm installation (false) or native Helm installation (true). Default is false.
  # Native Helm installations are only available for Helm v3 charts.
  useHelmInstall: true

  # weight determines the order that charts with "useHelmInstall: true" are applied, with lower weights first.
  weight: 0

  # values are used in the customer environment, as a pre-render step
  # these values will be supplied to helm template
  values: {}
```

Before we create our release, let's initialize git. We will use a command line option that uses Git metadata to determine the version label, channel to promote to and more.

```bash
    $ git init
    $ git add .
    $ git commit -m "this is my first commit"
```
Now we are ready to create our first release:

```bash
    $ replicated release create --auto -y
```

## Deploy the Application

We now have something we can deploy! As part of this lab, you were provided with the IP address of a Virtual Machine. Log in using `ssh kots@{IP-Address}` using the IP address provided. The password will be provided during this lab.

Once you are on the terminal, copy the install command from the channel. 

## Mapping Field Values

Let's take some of the field values and map them to the ui!

We are going to make two changes:
* Map the `wordpressBlogName` field in the Values.yaml file to a UI field.
* Override the default port of 80 to 8080.

### Copy the field name from the Values file

For this lab we are going to use the `wordpressBlogName` field in the values file and expose it in the UI. For reference, we are going to be referencing fields in the ../charts/bitnami/wordpress/values.yaml file.

![Values File](./img/values-file.png)

### Create a field in the Config.yaml file

Select the `kots-config.yaml` file. Replace the contents of this file with the following:

```yaml
apiVersion: kots.io/v1beta1
kind: Config
metadata:
  name: wordpress-config
spec:
  groups:
    - name: wordpress
      title: Wordpress
      description: Wordpress Defaults
      items:
        - name: wordpressBlogName
          title: Wordpress Blog Name
          type: text

```

As you can see, we are using the values field name as the config field name and used for the title, which is what the end user will see.


### Map the Config file field with the field in the Values file

Select the `wordpress.yaml` file and scroll down to the values section and make the following changes:

```diff
# values are used in the customer environment, as a pre-render step
# these values will be supplied to helm template
-- values: {}
++ values:
++   wordpressBlogName: '{{repl configOption "wordpressBlogName"}}'
```

## Update the Application

Let's see our changes on the Admin Console. Click on Check for Update as shown below.

Before you click on Deploy, note that now you have a Config tab with our field in it.

Deploy the new version of our app

### Random Password Generation

If you were paying close attention to the update we did in the previous step, you would have noticed that the credentials that are generated for MariaDB get generated again. When you peform an intial install or a subsequent upgrade, App Manager will use `helm upgrade -i` and will let Helm determine if there are any changes in the Chart.

Obviously we don't want this password to be reset each time we do an upgrade, so to solve this we are going to create a hidden config field which will have a random string value. We will use this value for the password that is used with MariaDB.

#### Create the field

Create a new release and go to the Config.yaml file we created earlier. We are now going to create a field that will generate a random string.

```yaml
# config.yaml
apiVersion: kots.io/v1beta1
kind: Config
metadata:
  name: wordpress-config
spec:
  groups:
    - name: wordpress
      title: Wordpress
      description: Wordpress Defaults
      items:
        - name: wordpressBlogName
          title: Wordpress Blog Name
          type: text
        - name: wordpress-db-secret
          hidden: true
          type: password
          value: "{{repl RandomString 16}}"
```

Next, we need to map this value to the field in the Values file. Note that if this field was not exposed in the Values file, we would likely need to change the chart so the field would be exposed in the values file.

```yaml
# wordpress.yaml

apiVersion: kots.io/v1beta1
kind: HelmChart
metadata:
  name: wordpress
spec:
  # chart identifies a matching chart from a .tgz
  chart:
    name: wordpress
    chartVersion: "14.0.6"
  # helmVersion identifies the Helm Version used to render the chart. Default is v2.
  helmVersion: v3

  # useHelmInstall identifies whether this Helm chart will use the
  # Replicated Helm installation (false) or native Helm installation (true). Default is false.
  # Native Helm installations are only available for Helm v3 charts.
  useHelmInstall: true

  # weight determines the order that charts with "useHelmInstall: true" are applied, with lower weights first.
  weight: 0

  # values are used in the customer environment, as a pre-render step
  # these values will be supplied to helm template
  values:
    wordpressBlogName: '{{repl ConfigOption "wordpressBlogName"}}'
    mariadb:
      auth:
        rootPassword: '{{ repl ConfigOption "wordpress-db-secret"}}'

```

## Make Update to Access Wordpress UI

By default, Wordpress runs on port 80 but given that it is a pretty popular port it may already be in use. In the `values.yaml` file, the port is exposed like this:



```diff
    values:
      wordpressBlogName: '{{repl ConfigOption "wordpressBlogName"}}'
++    service:
++      port: 8080
      mariadb:
        auth:
          rootPassword: '{{ repl ConfigOption "wordpress-db-secret"}}'
```

## Update and Test Change

Now that we have a new release available, let's update our deployed application to test the change.

Click on **Check for Update** and click on **Deploy** once available.


## Managing Private Images

In this lab, we used publicly available images so we didn't have to worry about accessing images that may be hosted in a private registry. When it comes to Helm Charts, we need to look at if and how the images are avaialble to be modified via the Values file.

When using private images, you must first configure the Vendor Portal to have access to those images. There are two options:

* Use Replicated as a pull-through proxy - This is an option if you are using a private cloud registry like Elastic Container Registry (ECR). To link the registry you would need to provide the credentials of a user with `pull` permissions.

* Push your images to the Replicated proxy - This option is for when your private container registry is behind a firewall or otherwise inaccessible to us or it's easier for you to push to our registry than to give Replicated access to your registry.

Depending on the approach taken, we may need modify image tags. If you opted to push your private images to our registry, then you would need to override your image tags to point to the Replicated registry. Usually these are exposed as fields in the ```Values.yaml``` file.

If you opt to use our registry, you should not have to change your image tags, as long as they are exposed in the `Values.yaml` as an `image` field. At deploy time, the App Manager will try to pull the images using your tags. When it receives an unauthorized response, it will append the image with our proxy path (i.e.,  `proxy.replicated.com.proxy/your-app/`)

If you have a field with a different name, Replicated will not know that it references a registry, repository path or image tag and will not modify it. In those instances, you may need to override the value of that field and add the proxy prefix.

## Preparing for Airgap

### Managing Private Images

When you go to build an airgap bundle, the Air Gap Builder will go through the `Values.yaml` and look for any `image` fields. If we find any, they are pulled and included in the bundle.

### Manging Public Images

Depending on the complexity of your application, you may rely on one or more containers that are pulled from Public Registries, like Docker hub. If these images are exposed in the `Values.yaml`, then the Air Gap Builder should include them in the bundle. If they are not, however, you should then add them to the `additionalImages` section of the `kots-app.yaml` file.

### Optional Images

Another thing to consider about images is making sure that all of them will be available in an airgap install. For example,  you may have a service like database that they customer has the option to deploy or connect to their own instance. In an airgap installation, you would need to make sure that the image is included in the bundle in case a user selects to deploy the service.

### Base64 Icon

By default, the icon used in the Admin Console is a publicly available image via the internet. In an airgap scenario, you will need to download the logo you are using and convert it to base64. There are several utilities out there to do this. Once you have the Base64 encoding (or in your clipboard), we recommend that you add this field to the bottom of the `kots-app.yaml` file for readability purposes


## Using Multiple Charts

In this lab we have used a single chart that contained two sub-charts. Note that a Replicated Application Release can contain more than one Helm chart. If you do need to include more than one Helm chart, it is important to understand the sequence by which these charts are deployed. 

Starting with version 1.69* of App Manager, charts are deployed in alphanumeric order. If this is not the sequence that you desire, you can use the `weight` field in the corresponding Replicated `kind:helmchart` file. The default value is 0 for any charts that do not have this field specified. Charts will be deployed in ascending numeric order based on this field. *For charts that have the same weight specified, then alpha numeric order takes precendence.
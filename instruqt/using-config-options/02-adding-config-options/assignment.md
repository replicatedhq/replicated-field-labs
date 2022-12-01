---
slug: adding-config-options
type: challenge
title: Adding Config Options
teaser: Let's add Config Options to our Application
notes:
- type: text
  contents: In this challenge we are going to add config options to manage an external
    datastore
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---

## Getting our Dev Environment Ready

We want to get the latest manifests from the application. To do this, we will use the `replicated` command line.

Go to the **Vendor Portal** tab and login using the credentials displayed on the **K3sVM** tab. Under **Settings** retrieve your Application slug and set it as your `REPLICATED_APP` environment variable on the **Dev** tab by running a command similar to

```bash

export REPLICATED_APP=...

```

Under your user settings, scroll all the way to the bottom to generate an API token. Use this to set as your `REPLICATED_API_TOKEN` environment variable on the **Dev** tab by running a command similar to

```bash

export REPLICATED_API_TOKEN=...

```

To ensure you have your environment variables set up correctly, run the following command:

```bash

replicated release ls

```

You should see a list of releases containing the initial release we have created for you. Next, we'll run the [replicated release download](https://docs.replicated.com/reference/replicated-cli-release-download) to download the manifests for us to build upon.

Let's create a directory to work on and then change to that directory

```bash
mkdir sts-app
cd sts-app
```
The `replicated` cli by default expects a `./manifests` directory so let's create one so we don't have to override that option later

```bash
mkdir manifests
```

```bash
replicated release download 1 -d ./manifests
```


## Managing Secrets

Our sample application does not follow security best pratices as it is using a hard coded value for the password in the initial admin account. Let's fix that. Let's have Replicated generate a random password for us that we can display to the end user.

Navigate to the **Code** tab and make sure you have selected the **manifests** directory we created earlier. Click on the icon to create a new file as shown below


Add the following content to the file:

```yaml
#kots-config.yaml
apiVersion: kots.io/v1beta1
kind: Config
metadata:
  name: sts-app
spec:
  groups:
    - name: database
      title: Database Options
      items:        
        - name: embedded_postgres_password
          type: password
          hidden: true
          value: "{{repl RandomString 32}}"
```

We will build on the Conig file in a moment, but for now all it does is generate a random string that we'll use as the password for Postgres. You'll notice that the value is in the format of `"{{repl ...}}"` which is an example of how to use [template functions](https://docs.replicated.com/reference/template-functions-about). In this specific case, we are using the [RandomString](https://docs.replicated.com/reference/template-functions-static-context#randomstring) [static context](https://docs.replicated.com/reference/template-functions-static-context) to generate the value.

Let's create another file, this time a `Secrets` definition file:

```yaml
# postgres-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: postgres
data:
  DB_PASSWORD: '{{repl ConfigOption "embedded_postgres_password" | Base64Encode }}'

```

In the secret definition file we are using the [ConfigOption](https://docs.replicated.com/reference/template-functions-config-context#configoption), which allows us to get the value from the specified field, in our case the **embedded_postgress_password** field defined in the `kots-config.yaml` file. Since this is a secret we need to encode it, so note how we use the [Base64Encode](https://docs.replicated.com/reference/template-functions-static-context#base64encode) Static Context.

Next we need to update both the `postgres.yaml` and `sts-app.yaml` to use the secret.  


In both files, the password is set as such:

```yaml

            - name: POSTGRES_PASSWORD
              value: "postgres"

```
Update both files to have this instead:

```yaml
            - name: POSTGRES_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: postgres
                  key: DB_PASSWORD
```

Save your changes and let's create a new release by heading over to the **Dev** tab and run the following command:

```bash

replicated release create --version 2.0 ...

```

### Validating the change

Since the current app deployment already deployed `postgres`, our update only works for the initial password and will not update an existing instance. To make sure this change works, delete the postgres StatefulSet. 

Navigate BACK to the **Dev** tab which has `kubectl` access to the K3s cluster and run the following command, keeping in mind that the `namespace` is likely the application slug we used to set the `REPLICATED_APP` environment. 

```bash

kubectl get sts -n $REPLICATED_APP

```

If it can't find the namespace, get the list of namespaces by running `kubectl get namespaces` and try again with the correct namespace.

Run the following command to delete the postgres StatefulSet:

```bash

kubectl delete sts postgres -n $REPLICATED_APP

```

Once deleted, head over to the **Admin Console** tab to consume the update.

Once deployed, you can verify that the new password is being used by first running:

```bash
kubectl exec -it pod postgres-0 bash
```

Once inside the `postgres` container, run the following command to get the password:

```bash
echo $POSTGRES_PASSWORD
```

You should see a random string as the password. 


## Adding External Database Options

To provide our end users with the option to specify an external database, we'll need the following:

* The option to indicate to use an external database
* Text fields for the following
  * Hostname
  * Port
  * username
  * database
* Password field for the database password

The fields that will be used for the database connection string will only become visible if the user chooses to use an external database.

To start we'll add on to our `kots-config.yaml` file, first by adding the option to choose which database to use.

To do this, we'll add a new [select one](https://docs.replicated.com/reference/custom-resource-config#select_one) item type to the file like so:

```yaml
        - name: postgres_type
          help_text: This app needs Postgres. Would you like us to deploy it, or would you rather connect to an external instance that you manage?
          type: select_one
          title: Postgres
          default: embedded_postgres
          items:
            - name: embedded_postgres
              title: Yes deploy postgres!
            - name: external_postgres
              title: Let's connect to my own instance!
```
The updated file should look like this:

```yaml
#kots-config.yaml
apiVersion: kots.io/v1beta1
kind: Config
metadata:
  name: sts-app
spec:
  groups:
    - name: database
      title: Database Options
      items:        
        - name: embedded_postgres_password
          type: password
          hidden: true
          value: "{{repl RandomString 32}}"
        - name: postgres_type
          help_text: This app needs Postgres. Would you like us to deploy it, or would you rather connect to an external instance that you manage?
          type: select_one
          title: Postgres
          default: embedded_postgres
          items:
            - name: embedded_postgres
              title: Yes deploy postgres!
            - name: external_postgres
              title: Let's connect to my own instance!
```

At this point you can save changes and create a new release using the Replicated CLI. Updating the deployed app should show you the radio buttons for they are not wired to anything.

## Making Fields Visible Based on Condition

Next, we are going to add fields to capture the information we need to connect to the external database. However, we only want to display them when the end user selects **Let's connect to my own instance!**. To accomplish this we'll take advantage of the [when](https://docs.replicated.com/reference/custom-resource-config#when-1) item property. When you add this property and set to `true` the field is displayed, otherwise the field is not rendered.

However, we can't hard code this property value, as it will be based on the user's selection. Thankfully, we can use a template function to set the property value based on the selection of the user:

```yaml
when: '{{repl ConfigOptionEquals "postgres_type" "external_postgres"}}'

```

Let's go back to the code editor and add the following items to our `Group`:

```yaml
        - name: external_postgres_host
          title: Postgres Host
          when: '{{repl ConfigOptionEquals "postgres_type" "external_postgres"}}'
          type: text
          default: postgres
        - name: external_postgres_port
          title: Postgres Port
          when: '{{repl ConfigOptionEquals "postgres_type" "external_postgres"}}'
          type: text
          default: "5432"
        - name: external_postgres_user
          title: Postgres Username
          when: '{{repl ConfigOptionEquals "postgres_type" "external_postgres"}}'
          type: text
          required: true
        - name: external_postgres_password
          title: Postgres Password
          when: '{{repl ConfigOptionEquals "postgres_type" "external_postgres"}}'
          type: password
          required: true
        - name: external_postgres_db
          title: Postgres Database
          when: '{{repl ConfigOptionEquals "postgres_type" "external_postgres"}}'
          type: text
          default: postgres    
```
Your entire config file should look like this:

```yaml
#kots-config.yaml
apiVersion: kots.io/v1beta1
kind: Config
metadata:
  name: sts-app
spec:
  groups:
    - name: database
      title: Database Options
      items:        
        - name: embedded_postgres_password
          type: password
          hidden: true
          value: "{{repl RandomString 32}}"
        - name: postgres_type
          help_text: This app needs Postgres. Would you like us to deploy it, or would you rather connect to an external instance that you manage?
          type: select_one
          title: Postgres
          default: embedded_postgres
          items:
            - name: embedded_postgres
              title: Yes deploy postgres!
            - name: external_postgres
              title: Let's connect to my own instance!
        - name: external_postgres_host
          title: Postgres Host
          when: '{{repl ConfigOptionEquals "postgres_type" "external_postgres"}}'
          type: text
          default: postgres
        - name: external_postgres_port
          title: Postgres Port
          when: '{{repl ConfigOptionEquals "postgres_type" "external_postgres"}}'
          type: text
          default: "5432"
        - name: external_postgres_user
          title: Postgres Username
          when: '{{repl ConfigOptionEquals "postgres_type" "external_postgres"}}'
          type: text
          required: true
        - name: external_postgres_password
          title: Postgres Password
          when: '{{repl ConfigOptionEquals "postgres_type" "external_postgres"}}'
          type: password
          required: true
        - name: external_postgres_db
          title: Postgres Database
          when: '{{repl ConfigOptionEquals "postgres_type" "external_postgres"}}'
          type: text
          default: postgres
```
Save your changes and create a new release. If you update the deployed application, you should see the fields we have added get rendered when you select the **Let's connect to my own instance!** option. In the next track, we are going to map these fields to our application to connect to the right `postgres` instance.
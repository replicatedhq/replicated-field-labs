---
slug: using-config-values
type: challenge
title: Using Config Values
teaser: A short description of the challenge.
notes:
- type: text
  contents: Replace this text with your own text
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---

In the previous challenge we added the config options so our end users have the option to use their database instance. We also got a brief introduction into using templating to map a value to a definition file.

## Ensuring Postgres is Only Deployed When Needed

Let's start by making sure that if the end user wants to use their own database, we don't waste any resources deploying an embedded postgres. In our sample app, there are three definition files related to Postgres:

* **postgres.yaml:** this defines the Postgres StatefulSet
* **postgres-service.yaml** Defines the service used to connect to postgres
* **postgres-pvs.yaml** Defines the PVC used by postgres

For each of these we want to add the following line to the label annotation section:

```yaml
  kots.io/when: '{{repl ConfigOptionEquals "postgres_type" "embedded_postgres"}}'
```

Then above line tells Replicated to only send this to the Kubernetes API when the condition resolves to `true`. When the user selects to connect to their own database, that will not be the case and any definition that includes the above line will not be sent to the Kubernetes API.


We've included the modified files below for reference, but collapsed them for readability
<details>
  <summary>Click to see the modified postgres.yaml</summary>

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: postgres
  labels:
    app: postgres
  annotations:
    kots.io/when: '{{repl ConfigOptionEquals "postgres_type" "embedded_postgres"}}'
spec:
  replicas: 1
  selector:
    matchLabels:
      app: postgres
  serviceName: postgres-service
  template:
    metadata:
      labels:
        app: postgres
    spec:
      containers:
      - env:
        - name: PGDATA
          value: /var/lib/postgresql/data/pgdata
        # create a db called "postgres"
        - name: POSTGRES_DB
          value: "postgres"
        # create admin user with name "postgres"
        - name: POSTGRES_USER
          value: "postgres"
        - name: PGUSER
          value: "postgres"
        # use admin password from secret
        - name: POSTGRES_PASSWORD
          valueFrom:
            secretKeyRef:
              name: postgres
              key: DB_PASSWORD
        - name: POD_IP
          valueFrom: { fieldRef: { fieldPath: status.podIP } }
        ports:
        - name: postgresql
          containerPort: 5432  
        image: postgres:9.6
        name: postgres
        volumeMounts:
        - mountPath: /var/lib/postgresql/data
          name: postgres-pvc
      volumes:
      - name: postgres-pvc
        persistentVolumeClaim:
          claimName: postgres-pvc
  volumeClaimTemplates:
  - metadata:
      name: postgres-pvc
    spec:
      accessModes:
      - ReadWriteOnce
      resources:
        requests:
          storage: 1Gi
```          
</details>  


<details>
  <summary>Click to see the modified postgres-serice.yaml</summary>

```yaml
apiVersion: v1
kind: Service
metadata:
  name: repl-db-provider-service
  labels:
    app: repl-db-provider-service
  annotations:
    kots.io/when: '{{repl ConfigOptionEquals "postgres_type" "embedded_postgres"}}'
spec:
  type: ClusterIP
  ports:
  - name: postgresql
    port: 5432
    protocol: TCP
  selector:
    app: repl-db-provider
```

</details>

<details>
  <summary>Click to see the modified postgres-pvc.yaml</summary>

```yaml

kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: postgres-pvc
  labels:
    app: postgres
  annotations:
    kots.io/when: '{{repl ConfigOptionEquals "postgres_type" "embedded_postgres"}}'  
spec:
  accessModes:
    - "ReadWriteOnce"
  resources:
    requests:
      storage: "100Gi"
```

</details>

<details>
  <summary>Click to see the modified postgres-secret.yaml</summary>

```yaml

apiVersion: v1
kind: Secret
metadata:
  name: postgres
  annotations:
    kots.io/when: '{{repl ConfigOptionEquals "postgres_type" "embedded_postgres"}}' 
data:
  DB_PASSWORD: '{{repl ConfigOption "embedded_postgres_password" | Base64Encode }}'

```
</details>


Save your changes and create a new release. With the changes made, when a user selects to connect to their own database, the postgres resources (StatefulSet, Service, Secret & PVC) will not be created.


## Getting the STS App to Connect to the Correct Database

So how do we tell our app to connect to the right database? Well we have a couple of choices. 

The first option is to create two separate `sts-app.yaml` file, one which is the file in its current form (using hard coded values & secret) and the second one that get values from the config options. You can the use the same `label annotations` used in the above section to determine which file to deploy.

The second option is to have a single `sts-app.yaml` file, and use an `if` statement to determine which value to use. In this lab we are going to go with this option.

Here is the basic syntax of the command:

```yaml

 '{{repl if <condition that returns true or false>}}<value to use if true>{{repl else}}{{ <value to use if false> }}{{repl end}}'

```

Here is what we'll use for the `POSTGRES_USER` environment variable:

```yaml
  '{{repl if ConfigOptionEquals "postgres_type" "embedded_postgres"}}postgres{{repl else}}{{repl ConfigOption "external_postgres_user" }}{{repl end}}'

```

We'll repeat the same pattern for all of the fields:

```yaml
          env:
          # Define the environment variable
            - name: POSTGRES_HOST
              value: '{{repl if ConfigOptionEquals "postgres_type" "embedded_postgres"}}postgres{{repl else}}{{repl ConfigOption "external_postgres_host" }}{{repl end}}'
            - name: POSTGRES_USER
              value: '{{repl if ConfigOptionEquals "postgres_type" "embedded_postgres"}}postgres{{repl else}}{{repl ConfigOption "external_postgres_user" }}{{repl end}}'
            - name: POSTGRES_DB
              value: '{{repl if ConfigOptionEquals "postgres_type" "embedded_postgres"}}postgres{{repl else}}{{repl ConfigOption "external_postgres_db" }}{{repl end}}'
            - name: POSTGRES_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: postgres
                  key: DB_PASSWORD
            - name: POSTGRES_PORT
              value: '{{repl if ConfigOptionEquals "postgres_type" "embedded_postgres"}}5432{{repl else}}{{repl ConfigOption "external_postgres_port" }}{{repl end}}'

```

For reference, the entire file is included below in the collapsed section

<details>
  <summary>Click here to see entire sts-app.yaml</summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sts-app
  labels:
    app: sts-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: sts-app
  template:
    metadata:
      labels:
        app: sts-app
    spec:
      containers:
        - name: repl-db-consumer
          image: cremerfc/repl-db-consumer:latest
          ports:
            - containerPort: 5000
          env:
            - name: POSTGRES_HOST
              value: '{{repl if ConfigOptionEquals "postgres_type" "embedded_postgres"}}postgres{{repl else}}{{repl ConfigOption "external_postgres_host" }}{{repl end}}'
            - name: POSTGRES_USER
              value: '{{repl if ConfigOptionEquals "postgres_type" "embedded_postgres"}}postgres{{repl else}}{{repl ConfigOption "external_postgres_user" }}{{repl end}}'
            - name: POSTGRES_DB
              value: '{{repl if ConfigOptionEquals "postgres_type" "embedded_postgres"}}postgres{{repl else}}{{repl ConfigOption "external_postgres_db" }}{{repl end}}'
            - name: POSTGRES_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: postgres
                  key: DB_PASSWORD
            - name: POSTGRES_PORT
              value: '{{repl if ConfigOptionEquals "postgres_type" "embedded_postgres"}}5432{{repl else}}{{repl ConfigOption "external_postgres_port" }}{{repl end}}'
```
</details>  
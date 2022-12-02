---
slug: testing-the-external-database-connection
type: challenge
title: Testing the External Database Connection
teaser: Let's validate the connection and add preflights and support bundle
notes:
- type: text
  contents: -|
    In this challenge, we have provisioned a postgres instance for you to validate the changes made to the application in the previous challenge.
    We will also add a preflight check and support bundle to get the app ready for production.
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---

## Validating the External Database Connection

We have provisioned for you a postgres instance to connecto to.

Here is the connection information:

#todo

## Adding a preflight check

Ideally we would want to validate the connection to the database before the application attempts to. That can prevent the application from starting and getting into a broken state due it being unable to connect to a backend.

Let's add a preflight check to test the database:

```yaml
## needs verified and tested
apiVersion: troubleshoot.sh/v1beta2
kind: Preflight
metadata:
  name: helm-wordpress-preflights
spec:
  collectors:
    - postgres:
        exclude: repl{{ ConfigOptionEquals "option_select" "embedded_db" }}
        collectorName: pg
        uri: postgresql://'{{repl ConfigOption "db_username"}}':'{{repl ConfigOption "db_password"}}'@'{{repl "db_hostname"}}':5432/'{{repl ConfigOption "db_name"}}'?sslmode=required
  analyzers:
    - postgres:
        exclude: repl{{ ConfigOptionEquals "option_select" "embedded_db" }}
        strict: 'repl{{  eq (LicenseFieldValue "licenseType") "prod" }}'
        checkName: pg
        collectorName: pg
        outcomes:
          - fail:
              when: connected == false
              message: Cannot connect to PostgreSQL server
          - pass:
              message: The PostgreSQL server is ready
```

We won't get too deep into the actual `collector` or `analyzer` used other than to note how we are using templating to make these dynamic. 

First we are using the `exclude` shared property to only run the collector and analyzer when the user selects to use an external postgres.

Next, we are using the `strict` shared property in the analyzer to prevent the user from installing the application if the database connection test fails. In the above example this is only enforced if the license being used is **paid** (internal name is **prod**). If you use a different type of license, you'll be able to continue.

## Adding a Support Bundle

Our app does not have a support bundle so we should add a couple of things to help us troubleshoot. We can take the above `collector` and `analyzer` and make some minor changes:


```yaml

# a support bundle spec that can check for postgres internal or external.

```
---
slug: using-the-configuration
id: n8juiythxhxt
type: challenge
title: Using the Configuration to Install and Upgrade
teaser: Now we can customize the installation with the configuration
notes:
- type: text
  contents: |
    Provide the user configuration to your Helm chart
tabs:
- title: Shell
  type: terminal
  hostname: node
difficulty: basic
timelimit: 300
---

The configuration screen we built looks great, guides the customer through
their configuration, and helps make sure they set their configuration is set
up correctly. The next step is to configure the application using the options
they provide. Values from the configuration are mapped to the Helm chart(s)
that make up your application using the `HelmChart` resource. The Admin
Console uses this resource to prepare the values passed to Helm when
installing or upgrading the chart.

Passing Configurations to Helm
==============================

We saw the `HelmChart` object when we initially prepared the cluster. In that
section we specified a chart and version to install and explicitly provided no
values to the Helm command (`values: {}`). We're going to fix that now and
provide values based on the configuration the Admin Console collected.

### A Word on Templating

[Templating in the Admin Console](https://docs.replicated.com/reference/template-functions-about)
is a big topic and an in-depth treatment would take it's own lab (or three).
I'm going to explain the basics that we need to use here, building on the
glimpse we had when setting up conditional fields on the configuration screen.
To use the configuration we're going to look at a few more template functions
and how to use them.

### Setting Helm Chart Values

We're going to elaborate on our `HelmChart` manifest to include all the values
we collected from the user. Let's expand the update the file `slackernews-chart.yaml`
incrementally to set all the values we've asked for.

The `HelmChart` object has two ways of supplying values during installations
and updates. The `values` key supplies values that are used for every
operation, while the `optionalValues` key allows you to conditionally set
values. The combined set of values will be passed to the Helm command when the
Admin Console installs or updates the application.


### Values

The `values` key is for values that will always be set the same way during an
install or upgrade operation. These value can be templated, so they can change
based on context like user configuration or license details. But they will
always be passed to the `helm` command and will always be set using the same
expression. The authentication information for Slackernews is a good example.
It is always set using the four values the user provides during configuration.

```yaml
  values:
    slack:
      botToken: repl{{ ConfigOption "slack_bot_token" | quote }}
      userToken: repl{{ ConfigOption "slack_user_token" | quote }}
      clientId: repl{{ ConfigOption "slack_clientid" | quote }}
      clientSecret: repl{{ ConfigOption "slack_clientsecret" | quote }}
```

We're using KOTS template functions to pull in the configuration options. For
example, `repl{{ ConfigOption "slack_bot_token" | quote }}` will retrieve
the value of the `slack_bot_token` configuration option, quote it, and make
sure it's provided as `slack.botToken` to the `helm` command.

### Optional Values

Optional values are used when you want to conditionally set values based on
some other context that the Admin Console can provide. Generally that context
comes from the user's configuration, but it could also come from the license
or the cluster. You can use as many conditions as you need for optional
values. Each condition is part of the `optionalValues` list, specified with a
`when` key.

You also specify how to merge the values from `optionalValues` and `values`
together using `recursiveMerge`. If `recurisveMerge` is false, then the
top-level key from the `optionalValues` clause overwrites the top-level key
from `values`.

In other words, if you have `recursiveMerge: false` (the default), then:

```
values:
  slack:
    botToken: repl{{ ConfigOption "slack_bot_token" | quote }}
    userToken: repl{{ ConfigOption "slack_user_token" | quote }}
optionalValues:
  - when: `{{repl eq (LicenseFieldValue "licenseType") "trial" }}
    values:
      slack:
        mock: true
```

then the result will be:

```
slack:
  mock: true
```

whereas if you have `recursiveMerge: true`, then the result will be

```
  slack:
    botToken: xoxb-pretendthisisyourbottoken
    userToken: xoxp-pretendthisisyourusertoken
    mock: true
```

For Slackernews, we use optional values to determine how to configure the
database. Since there are some database configurations that are always set in
the same way, we specify a recursive merge.

```yaml
  optionalValues:
    - when: '{{repl ConfigOptionEquals "deploy_postgres" "1"}}'
      recursiveMerge: true
      values:
        postgres:
          password: '{{repl ConfigOption "postgres_password" }}'

    - when: '{{repl ConfigOptionEquals "deploy_postgres" "0"}}'
      recursiveMerge: true
      values:
        postgres:
          uri: '{{repl ConfigOption "postgres_external_uri" }}'
```

Specifying Image Values
=======================

If you completed the [Protecting Your
Assets](https://play.instruqt.com/manage/replicated/tracks/protecting-your-assets)
lab, you worked through the configuration for the Replicated Proxy Service to
secure access to your private images. You should specify how to access the
images through the proxy service as part of your Replicated Embedded Cluster
configuration as well. There are a few template functions to facilitate this.
These functions are part of the [config
context](https://docs.replicated.com/reference/template-functions-config-context)
within the Admin Console.

```yaml
    images:
      slackernews:
        pullSecret: repl{{ ImagePullSecretName }}
        repository: '{{repl HasLocalRegistry | ternary LocalRegistryHost "proxy.replicated.com" }}/{{repl HasLocalRegistry | ternary LocalRegistryNamespace (print "proxy/" (LicenseFieldValue "appSlug") "/ghcr.io/slackernews" ) }}/slackernews-web:1.0.17'
```

You are probably already referring to the Replicated Proxy Service in your
default Helm chart values, so it's natural to ask why you should also specify
it in the `HelmChart` object. Let's break down these values to understand why.

First, we set the `images.slackernewws.pullSecret` value to `repl{{ ImagePullSecretName }}`.
This grabs the image pull secret created by the Admin Console and let's the
chart know to use it when pulling the Slackernews image. In [Protecting You
Assets](https://play.instruqt.com/manage/replicated/tracks/protecting-your-assets),
you configured your own secret and that secret will work in an online
installation---the secrets will be identical. But in an airgap scenario, the
secret created by the Admin Console will contain credentials to a different
registry: the one configured inside the airgap.

Likewise, the `LocalRegistryHost` and `LocalRegistryNamespace` function refer
to the registry inside the airgapped environment. Using the `HasLocalRegistry`
function, you can distinguish between the online scenario (where you provide
the proxy service URI) and the airgapped scenario (where the Admin Console
fills in the details).

Providing a Complete Set of Values
==================================

Like with the last step of the lab, there's a more robust set of values needed
than just he ones we went through. To include the complete set of values in
your release, move the file `complete-helmchart.yaml` in your home directory
into the `release` directory.

```shell
mv complete-helmchart.yaml release/slackernews-chart.yaml
```

Releasing an Update
===================

We now have a complete release of the Slackernews application that a customer
can install. Let's release our update and move it through the release process.

Like our last release, we're going to bump the Helm chart version to keep the
versions aligned across all our install methods. Remember this is optional, so
if you feel funny bumping the chart version when you didn't change the chart
itself you can skip that part. In practice you'll usually be making changes to
both the chart and the Replicated Embedded Cluste configuration in parallel so
this probably won't be an issue.

Run the following commands to bump the chart version and add it to your
release.

```
yq -i '.version = "0.6.2"' slackernews/Chart.yaml
helm package -u slackernews -d release
rm release/slackernews-0.6.1.tgz
```

The `HelmChart` object needs to refer to the new chart version, so we need to
line up the version as well.

```
yq -i '.spec.chart.chartVersion = "0.6.2"' release/slackernews-chart.yaml
```

Now we can create our release and simulate a full release process by promoting
across the `Unstable`, `Beta`, and `Stable` channels. First build the release
and promote it directly to `Unstable`.

```
replicated release create --promote Unstable --chart ${HOME_DIR}/release/slackernews-0.6.2.tgz --version 0.6.2 \
  --release-notes "Collects configuration from the user and provides it to Helm" \
  --app ${REPLICATED_APP} --token ${REPLICATED_API_TOKEN}
```

Then you can promote to `Beta` using the release sequence from the output.

```
replicated release promote 7 Beta --version 0.6.2 \
  --release-notes "Collects configuration from the user and provides it to Helm" \
  --app ${REPLICATED_APP} --token ${REPLICATED_API_TOKEN}
```

And on to `Stable`

```
replicated release promote 7 Stable --version 0.6.2 \
```

Next, we'll see what your customer experiences when they install the
application.

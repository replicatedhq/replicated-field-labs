---
slug: helm-values-for-license-fields
id: egqdclosyyl5
type: challenge
title: Using License Fields in Your Helm Chart
teaser: Working with the Replicated license values in your Helm chart
notes:
- type: text
  contents: Let's assure a valid license before deploying your Helm chart
tabs:
- title: Shell
  type: terminal
  hostname: shell
  workdir: /home/replicant
- title: Slackernews Chart
  type: code
  hostname: shell
  path: /home/replicant
difficulty: basic
timelimit: 600
---

When you distribute your software with Replicated, Replicated injects the
license into your Helm chart in two ways:

1. As a value provided to the Replicated SDK to access via an in-cluster API
2. As global values that you can use in other components, including directly in
   your Helm templates.

We're going to take advantage of the second option to update the Slackernews
chart to only install when the license has not expired.

### A Word of Caution

The approach we're demonstrating here is easily defeated by overriding a Helm
value on the command-line or in a values file. It's meant to remind an honest
customer their license is expired rather than prevent them from tampering in
order to install anyway. Additional features provided by the proxy registry and
the Replicated SDK should be used to assure compliance and prevent tampering.

Available License Values
========================

The global values provided by the Replicated registry are the most valuable to
use within your Helm chart. Let's take a look at our Slackernews license to see
what other information about the license is available. Log in to the Replicated
registry as your Geeglo customer.

```
helm registry login [[ Instruqt-Var key="REGISTRY_HOST" hostname="shell" ]]  --username [[ Instruqt-Var key="REGISTRY_USERNAME" hostname="shell" ]]  --password [[ Instruqt-Var key="REGISTRY_PASSWORD" hostname="shell" ]]
```

You can then view the values using the `helm show values` command. The command
below will get the values and isolate the globals. For Slackernews, this
will only include the values injected by the Registry.

```
helm show values oci://[[ Instruqt-Var key="REGISTRY_HOST" hostname="shell"]]/[[ Instruqt-Var key="REPLICATED_APP" hostname="shell" ]]/slackernews | yq -P .global
```

You'll see something like (the encoded fields and expiration will probably be
different).

```
replicated:
  channelName: Unstable
  customerEmail: [[ Instruqt-Var key="CUSTOMER_EMAIL" hostname="shell" ]]
  customerName: Geeglo
  dockerconfigjson: eyJhdXRocyI6eyJpbWFnZXMuc2hvcnRyaWIuaW8iOnsiYXV0aCI6Ik1tWlhSMVEzZDBGck0yOVVTbGhWY25GVWFrWlFUV0k1ZVV4dE9qSm1WMGRVTjNkQmF6TnZWRXBZVlhKeFZHcEdVRTFpT1hsTWJRPT0ifSwicmVnaXN0cnkuc2hvcnRyaWIuaW8iOnsiYXV0aCI6Ik1tWlhSMVEzZDBGck0yOVVTbGhWY25GVWFrWlFUV0k1ZVV4dE9qSm1WMGRVTjNkQmF6TnZWRXBZVlhKeFZHcEdVRTFpT1hsTWJRPT0ifX19
  licenseFields:
    expires_at:
      name: expires_at
      title: Expiration
      description: License Expiration
      value: "2025-05-31T00:00:00Z"
      valueType: String
      signature:
        v1: UwI4/IL6JR4K5Tw7gapvfW6+zkirfMulbAxaQqkVIAZazip+pehegNRVhHEhbM9V9EibONGBbOazipb8aQeWO2hYoN0mcQOelUxVmK7U2GFP862tyorwAPwxMg+ZbAunUsoKP4/GT+Up5bhC8UN+NgyfZFzmCo3G6TK+2tbtI/tHXN0IwFacY3TyvryAfB+6qRGpsb0efb/Wl4DNmzuo/z9qE/1HbdG/TGUdq3SEmaH4iGBSeUZMHrHmW7/fHM8DDAABi7NW8v7HGEKI467yufKPcohzim8roLl2mdLsnLq2o3J5ovsAImgXUK3ac7ymHBLdT9WRg5cuBtAWNI5oaQ==
  licenseID: 2fWGT7wAk3oTJXUrqTjFPMb9yLm
  licenseType: prod
```

There are a few useful things in there:

* Which customer this license is for, including name, a unique license ID, and
  their email.
* The customer's credentials for the Replicated registry and proxy (we used
  this earlier).
* The type of license: `prod` (representing a paid license), `dev`, `trial`, or
  `community`.
* A set of fields representing license entitlements. We'll talk about those in
  the next section.

You can use this information to customize the user experience, guard components
(e.g. don't allow persistence with trials), and tailor the install to the
license entitlements.

License Fields
==============

The license has several fields that represent specific entitlements of this
license. In Helm, they're available as an array
`global.replicated.licenseFields`. Each is fields is represented as a YAML
dictionary. You can add your own fields to represent your unique entitlements.
We'll do that in the next part of the lab.

Every application will have at least one field in this array for the license
expiration date. If the license does not expire, the value will be blank.

```
expires_at:
  name: expires_at
  title: Expiration
  description: License Expiration
  value: "2025-05-31T00:00:00Z"
  valueType: String
  signature:
    v1: UwI4/IL6JR4K5Tw7gapvfW6+zkirfMulbAxaQqkVIAZazip+pehegNRVhHEhbM9V9EibONGBbOazipb8aQeWO2hYoN0mcQOelUxVmK7U2GFP862tyorwAPwxMg+ZbAunUsoKP4/GT+Up5bhC8UN+NgyfZFzmCo3G6TK+2tbtI/tHXN0IwFacY3TyvryAfB+6qRGpsb0efb/Wl4DNmzuo/z9qE/1HbdG/TGUdq3SEmaH4iGBSeUZMHrHmW7/fHM8DDAABi7NW8v7HGEKI467yufKPcohzim8roLl2mdLsnLq2o3J5ovsAImgXUK3ac7ymHBLdT9WRg5cuBtAWNI5oaQ==
```

Other license fields will be displayed in the same format. The name,
title, description, and type are set when you set up the field. Entitlements
can be integers, strings, text (multi-line strings), or boolean. All fields are
signed to prevent tampering. Unfortunately the built-in functions in Helm do
not include the capabilities we need to verify the signature. It's also trivial
to override both the value and the signature.

In spite of those limitations, there's still value to using the global values
in your Helm templates. You can use them to determine which features to
install, how to configure a service, or even whether to install at all. Just be
sure to also use the Replicated SDK to validate any critical license conditions
in your application code.

Checking for a Valid License
============================

Since our license has an expiration date, let's take advantage of that. We're
going to update the Slackernews deployment to prevent installation when the
expiration date has passed.

Using the editor in the "Slackernews Chart" tab, open the file
`slackernews/templates/slackernews-deployment`.

![Editing the Slackernews Deployment](../assets/editing-the-deployment.png)

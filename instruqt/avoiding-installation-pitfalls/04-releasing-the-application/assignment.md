---
slug: releasing-the-application
id: q5wbycoymxhk
type: challenge
title: Releasing the Application
teaser: Releasing with preflights on the Replicated Platform
notes:
- type: text
  contents: Time to release an update with our preflight checks
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 300
---

Since we're distributing our application with the Replicated
Platform, we need to let the platform know about the changes
we've made to the application by creating a new release. This
process will be familiar to you if you have completed
the [Distributing Your Application with Replicated](https://play.instruqt.com/replicated/tracks/distributing-your-application-with-replicated)
lab. If not, you may want to go through that lab to learn a bit
more about how releases and release channels work.

A Quick Look at Release Channels
================================

The Replicated Platform provides a way to connect
each customer to the right release(s) for them. It does this
by organizing releases into _channels_, and assigning each
customer license to the appropriate channel. Release channels
help you account for these different release cadences for
your software.

By default, Replicated creates three release channels for
each application: `Unstable`, `Beta`, and `Stable`. We're
going to release our updates to Harbor across all three of
those channels.

Preparing to Release
====================

Before we release, we need to make sure we're authenticated
to the Replicated Platform. We're going to use an API
token to do that. The lab setup created one for you. Let's
set it in an environment variable.


```
export REPLICATED_API_TOKEN="[[ Instruqt-Var key="REPLICATED_API_TOKEN" hostname="shell" ]]"
```

We also need to tell the `replicated` command which
application to work with. We can do this with every command,
but it's easier to just set an environment variable.

```
export REPLICATED_APP="[[ Instruqt-Var key="REPLICATED_APP" hostname="shell" ]]"
```

Creating a New Release
======================

There are two releases already available for the Harbor
application. We're going to release a third that includes
the preflight checks.

```
replicated release ls
```

You'll see that the the same release is current across
all three channels, and it has the sequence number `1`.
All releases are assigned a sequence number based on the
order in which they were created.

```
SEQUENCE    CREATED                 EDITED                  ACTIVE_CHANNELS
1           2023-06-08T00:23:40Z    0001-01-01T00:00:00Z    Stable,Beta,Unstable
```

To release our new version, we create a new release and
(optionally) assign it to a channel. It's a good practice
to make new releases on either the `Unstable` channel or
a channel specific to the feature you are working on. Let's
use the `Unstable` channel for this lab, since the latter
approach is best for teams working with feature branches.

```
replicated release create --promote Unstable --chart ./release/harbor-19.3.0.tgz --version 19.3.0  \
  --release-notes "Adds preflight checks to enable customers to validate cluster prerequisites before installing"
```

This creates a release for version `19.3.0` of the Harbor Helm
chart, and promotes it to the `Unstable` channel. The `create`
command output a sequence number that you'll need for `promote` (it
will be `2` if you haven't explored releasing a bit more).

```
  _ Reading manifests from ./release _
  _ Creating Release _
    _ SEQUENCE: 2
  _ Promoting _
    _ Channel 2Qa7rGeBiT3DaDK85s6FVKRC7Mn successfully set to release 2
```

For the lab, we're going to assume this release can be directly
shared on the `Beta` and `Stable` channels. Your actual release
process may have many more activities before releasing to either
of those channels---your teams review and approval processes,
steps in a continuous delivery pipeline, or both. Run the following command to promote our release to the `Beta` channel:

```
replicated release promote 2 Beta --version 19.3.0 \
  --release-notes "Adds preflight checks to enable customers to validate cluster prerequisites before installing"
```

Then promote to the `Stable` channel:

```
replicated release promote 2 Stable --version 19.3.0 \
  --release-notes "Adds preflight checks to enable customers to validate cluster prerequisites before installing"
```

List your releases again to see that the release has been
promoted.

```
replicated release ls
```

Your list of releases will now show three releases with the third
release available on the `Unstable`, `Beta`, and `Unstable` channels.

```
SEQUENCE    CREATED                 EDITED                  ACTIVE_CHANNELS
2           2023-06-10T20:22:14Z    0001-01-01T00:00:00Z    Stable,Beta,Unstable
1           2023-06-10T20:20:02Z    0001-01-01T00:00:00Z
```

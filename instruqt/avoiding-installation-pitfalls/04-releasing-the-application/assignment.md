---
slug: releasing-the-application
id: 14xfg2xrtgef
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
lab. If not, you may want to go throght lab to learn a bit
more about how releases and release channels work.

A Quick Look at Release Channels
================================

The Replicated platform provides a way to connect
each customer to the right release(s) for them. It does this
by organizing release into _channels_, and assigning each
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
set it into an environment variable.


```
export REPLICATED_API_TOKEN="[[ Instruqt-Var key="REPLICATED_API_TOKEN" hostname="shell" ]]"
```

We also need to tell the `replicated` commmand which
application to work with. We can do this with every command,
but it's easier to just set an environment variable.

```
export REPLICATED_APP="[[ Instruqt-Var key="REPLICATED_APP" hostname="shell" ]]"
```

Creating a New Release
======================

There are two releases already available for the Harbor
applicaiton. We're going to releaea a third that includes
the preflight checks.

```
replicated release ls
```

You'll see that the the same release is current across
all three channels, and it has the sequence number `2`.
All releases are assigned a sequence number base on the
order in which they are created.

```
SEQUENCE    CREATED                 EDITED                  ACTIVE_CHANNELS
2           2023-06-08T00:23:40Z    0001-01-01T00:00:00Z    Stable,Beta,Unstable
1           2023-06-08T00:19:43Z    0001-01-01T00:00:00Z
```

To release our new version, we create a new releaes and
(optionally) assign it to a channel. It's a good practice
to make new releases on either the `Unstable` channel or
a channel specific to the feature you are working on. Let's
use the `Unstable` channel for this lab, since the latter
approach is best for teams working with feature branches.

```
replicated release create --promote Unstable --yaml-dir ./release --version 16.8.0  \
  --release-notes "Adds preflight checks to enable customers to validate cluster prerequisites before installing"
```

This creates a release for version `16.8.0` of the Harbor Helm
chart, and promotes it to the `Unstable` channel. The `create`
command output sequence number that you'll need for `promote` (it
will be `3` if you haven't explored releasing a bit more).

```
  _ Reading manifests from ./release _
  _ Creating Release _
    _ SEQUENCE: 3
  _ Promoting _
    _ Channel 2Qa7rGeBiT3DaDK85s6FVKRC7Mn successfully set to release 2
```

For the lab, we're going to assume this release can be directly
shared on the `Beta` and `Stable` channels. Your actual release
process may have many more activities before releasing to either
of those channels---ther team processes, steps in a continuous
delivery pipeline, or both.

```
replicated release promote 3 Beta --version 16.8.0 \
  --release-notes "Adds preflight checks to enable customers to validate cluster prerequisites before installing"
```

and then

```
replicated release promote 3 Stable --version 16.8.0 \
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
3           2023-06-10T20:22:14Z    0001-01-01T00:00:00Z    Stable,Beta,Unstable
2           2023-06-10T20:21:13Z    0001-01-01T00:00:00Z    
1           2023-06-10T20:20:02Z    0001-01-01T00:00:00Z
```
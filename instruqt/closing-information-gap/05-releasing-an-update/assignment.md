---
slug: releasing-an-update
id: bb2ckqmiideb
type: challenge
title: Releasing an Update with the Support Bundle
teaser: Releasing a new version with the support bundle included
notes:
- type: text
  contents: Let's create a release to distribute the support bundle
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 300
---

To make the support bundle available to the customer, we need to distribute it
to them. Since the Replicated Platform handles distribution for us, we need to
let the platform know a new release is available. This process will be familiar
to you if you have completed the [Distributing Your Application with
Replicated](https://play.instruqt.com/replicated/tracks/distributing-your-application-with-replicated)
lab. You'll learn all you need to know to relasze your update here, but that
lab can help you to get a more complete picture if you'd like one.

Preparing to Release
====================

You'll need to be authenticated to the Replicated Platform to release your
update. The simplest way to do that is with the API token the lab setup create
for you. The `replicated` CLI will read that from an environment variable, so
let's set it.


```
export REPLICATED_API_TOKEN="[[ Instruqt-Var key="REPLICATED_API_TOKEN" hostname="shell" ]]"
```

The `replicated` command also needs to know which application to work with. You
can set it with a command flag or with an environment variable. Let's set the
variable to save some typing.

```
export REPLICATED_APP="[[ Instruqt-Var key="REPLICATED_APP" hostname="shell" ]]"
```

Creating an Upgraded Release
============================

There will be three releases already available for the Slackernews application. This
shows the eovlution of the application over time.

```
replicated release ls
```

All releases are assigned a sequence number based on the order in which they
are created. They may also be assigned to zero or more release channels that
provide specific streams of releases to specific customers. You'll see that the
the same release is current across all three channels, and it has the sequence
number `2`

```
SEQUENCE    CREATED                 EDITED                  ACTIVE_CHANNELS
2           2023-06-20T14:52:16Z    0001-01-01T00:00:00Z    Stable,Beta,Unstable
1           2023-06-20T14:50:52Z    0001-01-01T00:00:00Z
```

To release our update, we're going to create a new release including our
changes and assign it to the release channel `Unstable`, which is one of the
default release channels that is generally used internally for releases that
may or not be ready for customers.

```
replicated release create --promote Unstable --chart ./release/slackernews-0.4.0.tgz --version 0.4.0  \
  --release-notes "Adds a support bundle spec to facilitate troubleshooting"
```

This creates a release for version `0.4.0` of the Slackernews Helm chart, and
promotes it to the `Unstable` channel. The `create` command output sequence
number that you'll need for `promote` (it will be `3` if you haven't explored
releasing a bit more).

```
  _ Reading manifests from ./release _
  _ Creating Release _
    _ SEQUENCE: 3
  _ Promoting _
    _ Channel 2Qa7rGeBiT3DaDK85s6FVKRC7Mn successfully set to release 4
```


A Note About Release Channels
================================

The previous section mentioned release channels a few times. The [Distributing
Your Application with
Replicated](https://play.instruqt.com/replicated/tracks/distributing-your-application-with-replicated)
lab goes into detail about release channels but it's worth explaining them a
pbit more here in case you haven't had a chance to go through that lab yet.
Release channels on the Replicated Platform allow you to provide difference
streams or release to different types of customers to make sure you get the
right releaes to each one. Each customer license is assigned to a channel, and
you assign each of your release to one or more channels.

By default, Replicated creates three release channels for each application:
`Unstable`, `Beta`, and `Stable`. You can probably guess from the names what
types of releases each one is intended for. You will likely also add your own
release channels to model how you distribute your software. Replicated
recommends, for example, that teams that use feature branches and/or a pull
request workflow create release channels for each branch. This allow teams to
create and distribuite releases from the branch for testing and validation.
Other team  creates release for customers who want either more or less frequent
releases that the standard release cadence.

Promoting the Release
=====================

For the lab, we're going to assume this release can be directly shared on the
`Beta` and `Stable` channels. You'll no doubt have a much more thorough process
to determine whether a release should be promoted to each of those
channels---hopefully automated as part of your continuous delivery pipelines.

```
replicated release promote 3 Beta --version 0.4.0 \
  --release-notes "Adds a support bundle spec to facilitate troubleshooting"
```

and then

```
replicated release promote 3 Stable --version 0.4.0 \
  --release-notes "Adds a support bundle spec to facilitate troubleshooting"
```

List your releases again to see that the release has been promoted.

```
replicated release ls
```

Your list of releases will now show four releases with the most recent release
available on the `Unstable`, `Beta`, and `Unstable` channels.

```
SEQUENCE    CREATED                 EDITED                  ACTIVE_CHANNELS
3           2023-06-20T14:55:32Z    0001-01-01T00:00:00Z    Stable,Beta,Unstable
2           2023-06-20T14:52:08Z    0001-01-01T00:00:00Z
1           2023-06-20T14:50:52Z    0001-01-01T00:00:00Z
```

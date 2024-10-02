---
slug: creating-a-release
id: vxnhm8nlkifh
type: challenge
title: Releasing an Application
teaser: Creating a release on the Replicated Platform
notes:
- type: text
  contents: It's time to distribute our application with Replicated
tabs:
- id: y93opy8anppl
  title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 300
---

To take advantage of the Replicated Platform to distribute
an application, we need to let the platform know about the
application, its releases, and the customers who are entitled
to access it. The lab environment has created the
Slackernews application for us, and created a customer, "Omozan",
that has access to it. All we need to do is create a
release and we'll be ready to go.

Replicated Release Channels
===========================

There's one more part of the equation that we didn't mention
above. The Replicated Platform provides a way to connect
each customer to the right release(s) for them. It does this
by organizing releases into _channels_, and assigning each
customer license to the appropriate channel.

We encounter this concept in our day-to-day use of software
all the time. For some applications, you sign up to
receive beta releases, while others you may receive updates only
when they're GA. You may even have some software, for example
your Linux distribution, where you use only releases that
have long term support.

Release channels help you account for the different release
cadences of your software. By default, Replicated creates
three release channels for each application.

```
replicated channel ls
```

You should see three channels in the output:

* `Unstable` is, as it sounds, releases that may be unstable
   and subject to defects and/or constant change. You may have
   every merge PR hit this channel, for example.
* `Beta` represents releases that have beta quality, for
  the customers you have as part of a beta program for new
  releases.
* `Stable` is for GA releases that you want to be broadly available
  to your customer base. This would be the default
  channel for any customer who did not opt-in to an alternative.

You may consider a few other uses for release channels in your
release process. For examples, let's add a channel called `LTS`
for those customers who want longer term guarantees of
support and fitness that you provide for your standard releases.

```
replicated channel create --name LTS \
  --description "Releases with long-term support available"
```

Other examples:

* `Edge` for customers who want continuous delivery of your
   software to their environments.
* Channels named after the feature branches in your source
  code. These can help product teams validate releases before
  they are merged for release on your primary channels.
  Replicated recommends all teams follow this approach.

Creating Your Release
=====================

To create a release, run the following command. We're using the
`Unstable` channel since we're releasing our most recent change.

```
replicated release create --promote Unstable --chart ./release/slackernews-0.2.0.tgz --version 0.2.0  \
  --release-notes "Prepares for distribution with Replicated by incorporating the Replicated SDK"
```

This creates a release for version `0.2.0` of your Slackernews Helm
chart, and promotes it to the `Unstable` channel.  The `create`
command output a sequence number that you'll need for `promote`
(it will be `1` if you haven't explored releasing a bit more).

```
  _ Reading manifests from ./release _
  _ Creating Release _
    _ SEQUENCE: 1
  _ Promoting _
    _ Channel 2Qa7rGeBiT3DaDK85s6FVKRC7Mn successfully set to release 2
```

The sequence number uniquely identifies a release among all the
releases you've made for your application. You can list your
releases using

```
replicated release ls
```

which should show the initial release created during lab
set up, as well as the release you just created. It will
also show the channel each release is currently available on,
if any.

```
SEQUENCE    CREATED                 EDITED                  ACTIVE_CHANNELS
1           2023-06-08T00:23:40Z    0001-01-01T00:00:00Z    Unstable
```

To make an existing release available on another channel, use
`replicated release promote`. In your actual release process,
there may be a lot of activity between releasing to `Unstable`,
promoting to `Beta`, and ultimately releasing on `Stable`.
For the purposes of the lab, let's just promote the release straight through.

```
replicated release promote 1 Beta --version 0.2.0 \
  --release-notes "Prepares for distribution with Replicated by incorporating the Replicated SDK"
```

and then

```
replicated release promote 1 Stable --version 0.2.0 \
  --release-notes "Prepares for distribution with Replicated by incorporating the Replicated SDK"
```

You can see they were promoted by listing your releases again. You should see
similar output to the following:

```
SEQUENCE    CREATED                 EDITED                  ACTIVE_CHANNELS
1           2023-06-08T00:23:40Z    0001-01-01T00:00:00Z    Stable,Beta,Unstable
```

---
slug: creating-a-release
id: zf43x5d6newr
type: challenge
title: Releasing an Application
teaser: Creating a release on the Replicated Platform
notes:
- type: text
  contents: It's time to distribute our application with Replicated
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 300
---

To take advantage of the Replicated Platform to distribute
an application we need to let the platform know about the
application, its releases, and the customers who are entitled
to access it. The lab environment has created the application
Mastodon application for us, and created a customer "Omozan"
that has access to it. All we need to do is to create a
release and we'll be ready to go.

Replicated Release Channels
===========================

There's one more part of the equation that we didn't mention
above. The Replicated platform provides a way to connect
each customer to the right release(s) for them. It does this
by organizing release into _channels_, and assigning each
customer license to the appropriate channel.

We encounter this concept in our day-to-day use of software
all the time. For some applications, you're signed up to
receive beta releases, while others your receive updates only
when they're GA. You may even have some software, for example
your Linux distribution, where you use only releases that
have long term support.

Release channels help you account for these different release
cadences for your software. By default, Replicated creates
three release channels for each application:

* `Unstable` is, as it sounds, releases that may be unstable
   and subject to defects and/or constant change. You may have
   every merge PR hit this channel, for example.
* `Beta` represents release that have release beta quality, for
  those customer you have as part of a beta program for new
  releases.
* `Stable` is for GA release that you want to be available
  broadly to your customer base. This would be the default
  channel for any customer who did not opt-in to an alternative.

You may consider a few other uses for release channels in your
release process. Some examples:

* `LTS` for those customer who want longer term gaurantees
  of support and fitness that you provide for your standard
  releases.
* `Edge` for customer who want continuous delivery of your
   software to their environmnets.
* Channels named after the feature branches in your source
  code. These can help product teams validate release before
  they are merged for release on your primary channels.
  Replicated recomments all teams follow this approach.

Creating Your Release
=====================

To create a release, run the following command:

```
replicated release create --promote Unstable --yaml-dir ./release --version 1.6.0  \
  --release-notes "Prepares for distribution with Replicated by incorporating the Replicated SDK"
```

This creates a release for version `1.6.0` of your Mastodon Helm
Chart, and promotes it to the `Unstable` channel. To release it
to another channel, use `replicated release promote`. The `create`
command output sequence number that you'll need for `promote` (it
will be `2` if you haven't explored releasing a bit more).

```
  _ Reading manifests from ./release _
  _ Creating Release _
    _ SEQUENCE: 2
  _ Promoting _
    _ Channel 2Qa7rGeBiT3DaDK85s6FVKRC7Mn successfully set to release 2
```

In your actual release process, there may be a lot of activity
between releasing to `Unstable`, promoting to `Beta`, and
ultimately releasing on `Stable`. For the purposes of the lab,
let's just promote the release straight through.

```
replicated release promote 2 Beta --version 1.6.0 \
  --release-notes "Prepares for distribution with Replicated by incorporating the Replicated SDK"
```

and then

```
replicated release promote 2 Stable --version 1.6.0 \
  --release-notes "Prepares for distribution with Replicated by incorporating the Replicated SDK"
```


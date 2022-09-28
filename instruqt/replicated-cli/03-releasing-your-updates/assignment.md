---
slug: releasing-your-updates
id: vtwciot9j8bb
type: challenge
title: Releasing Your Updates
teaser: Cut a new release for your customers
notes:
- type: text
  contents: Releasing your changes to your customer
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---

🚀 Let's start
==============

In the last challenge, we updated our release manifests to
follow the adviced provided by the Replicated linter. We
need to release our changes in order for our customers to
take advantage of them.


### 1. Release channels

Releases are shared with customers user channels. We have set
up the channels for your application for you. Let's take a
look at what channels are available.

```shell script
replicated channel ls
```

You should get a list that has four channels in it. Our
application has the default release channels as well as one
channel we created for the lab. The names of your channels
should be the same as you see here. The other fields may
vary.

```
ID                             NAME              RELEASE    VERSION
2FK67d5b0y2ilwbAkouIm5Ly98U    Stable            1          0.0.1
2FK67cR5l41w4FrJ17oQz1f0pDX    Beta              1          0.0.1
2FK67bqiSUUpl0PpEHe8fZYuXXK    Unstable          2          Sample Track
2FKLzwElQOFuHs6YlYEvZ6ncNEo    replicated-cli    3          The Replicated CLI
```

We're going to release our changes on the `replicated-cli`
channel in the later steps.

### 2. Creating your first release

Let's create a release with our improved YAML files. We're going
to create our release and make it available on the `Unstable` channel
for internal user. After we review  the release, we'll promote
it to our `replicated-cli` channel to simulate releasing to the customer.
Make sure to change the version number (in my case, I'll go from
`0.0.1` to `0.0.2`).

```
replicated release create --version [NEW VERSION] --release-notes "Adds resource requests to our deployment" \
  --promote Unstable --yaml-dir manifests
```

You can view your new release (along with previous releases) using

```
replicated release ls
```

Your new release is at the top of the list with a unique sequence
number.

```
SEQUENCE    CREATED                 EDITED                  ACTIVE_CHANNELS
4           2022-09-27T20:39:40Z    0001-01-01T00:00:00Z    Unstable
3           2022-09-26T23:48:15Z    0001-01-01T00:00:00Z    replicated-cli
2           2022-09-26T21:39:47Z    0001-01-01T00:00:00Z
1           2022-09-26T21:37:47Z    0001-01-01T00:00:00Z
```

Your date and times will be different, and it's OK if the sequence numbers
and the active channels for older releases differ.

### 3. Promoting the release for customers

To make the release available for customers, let's use our customer
release channel `replicated-cli`. In a real use case, you might promote
through other channles like `Beta` before going to customers, but we'll
skip ahead this time.

Look back at your list of releases and take note of the sequence number
for your latest release (sequence `4` in my case). You're going to
promote that release in the next command.

```
replicated release promote [SEQUENCE] replicated-cli \
  --version [NEW VERSION] --release-notes "Adds resource requests to our deployment"
```

After promoting the release, take a look at your releases again

```
replicated release ls
```

Your new release is at the top of the list is now active on two
channels, `Unstable` and `replicated-cli`.

```
SEQUENCE    CREATED                 EDITED                  ACTIVE_CHANNELS
4           2022-09-27T20:39:40Z    0001-01-01T00:00:00Z    Unstable,replicated-cli
3           2022-09-26T23:48:15Z    0001-01-01T00:00:00Z
2           2022-09-26T21:39:47Z    0001-01-01T00:00:00Z
1           2022-09-26T21:37:47Z    0001-01-01T00:00:00Z
```

---
slug: getting-to-know-the-template
type: challenge
title: Getting to Know the Template
teaser: Some tips and tricks for using this template
notes:
  - type: text
    contents: Let's learn about this template
tabs:
  - title: Shell
    type: terminal
    hostname: node
  - title: Cluster
    type: terminal
    hostname: cluster
difficulty: basic
timelimit: 300
---

👋 Introduction
===============

This template is a baseline for labs that need to persist their shell
environment across challenges. This may be because you as the learner
to set some environment variables, or because they've started a long
running process, or just to make it feel more like the real world
where they're doing everything in the same shell session.

As a cool side-effect, you can also use this template if you want
to interact with the contents of the learner's shell session. The
track uses `tmux` to persist the shell, and with that comes the
opportunity to read the content of the entire `tmux` pane. That
content includes the commands the learner types and the output of
those commands. This can come in very handy in lifecycle scripts, as
can `tmux`'s ability to send keystrokes into the session.


🔤 Basics
=========

You don't really have to do anything special to use this template.
It's configured to start a shell container and a single node Kubernetes
cluster. The shell uses our Instruqt shell image, and runs a `tmux`
session named `shell`. In that sesion it starts a login shell as the
user `replicant` using `su`.

The first challenge will create the session, and additional challenges
will connect to the existing session. This is enabled by following
command which is the `shell` specified in `config.yml` for the `Shell`
sandbox.

```yaml
- name: shell
  image: gcr.io/kots-field-labs/shell:instruqt-feature-tmux-template
  shell: tmux new-session -A -s shell su - replicant
```

This one command will either create a new session named `shell` running
`su - replicant`, or join an existing session named `shell`. The existing
session will continue with whatever command it was running in the prior
challenge which may just be `replicant`'s shell.

🧪 Try It
=========

Let's set an environment variable in this challenge so we can take
advantage of it in the next one.

```shell
export THIS="the way"
```

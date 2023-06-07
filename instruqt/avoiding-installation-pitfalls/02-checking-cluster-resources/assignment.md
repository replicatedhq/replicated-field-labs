---
slug: checking-cluster-resources
id: iy5zaqt0tx9a
type: challenge
title: Checking Cluster Resources
teaser: Some guidance on using the template
notes:
- type: text
  contents: How and Why to Use This Template
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 300
---

✨ Uses
=======

## Environment Variables

Only one shell runs across all challenges. This means the values of
environment variables persist from challenge to challenge without the
user setting them into their `.bashrc`.

## Long Running Commands

If a challenge needs to end with long running command (for example
downloading an airgap bundle or starting a kURL install), tell the user
they can click **Check** and leave the comamnd running. When the next
challenge starts then their shell will look the same and the command
will still be running (unless it happened to finish during the Cleanup
and Check scripts).

🔄 Lifecycle Scripts
====================

Lifecycle scripts can take advantage of `tmux` to read and write from
the learner's session. This is useful in Check scripts, for example,
to read what the user has typed and what the output was from those
commands. It also means that Setup, Cleanup, and Solve scripts can
type into the users shell to run commands.

Here are a couple of `tmux` commands to be aware of to interact with
the session:

`tmux capture-pane`
: This command the history of what's been done in the learner's
shell so you can intertact with it, for example to test whether
they typed the commands you expected

`tmux save-buffer`
: After you've captured what the learner has done, you can use the
`save-buffer` command to access it. The combination of the two is
useful in Check scripts

`tmux send-keys`
: Allows you to send keystrokes to the learner's shell. You have to
be explicit about charaters like `SPACE` and `ENTER` so that they
are sent to. This can be great for Solve scripts.

`tmux clear-histry`
: Clears the scrollback history (not the shell history) to keep
what's captured by `capture-pane` nice and fresh.

🧪 Did It Work?
===============

Remember that variable we set in the last step? Let's make sure it
stuck around like we expected.

```
echo $THIS
```

It should have, so you should see

```text
replicant@shell:~$ echo $THIS
the way
replicant@shell:~$
```

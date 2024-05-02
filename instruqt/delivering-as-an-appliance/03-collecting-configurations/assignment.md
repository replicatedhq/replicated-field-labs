---
slug: tips-and-tricks
id: nk3jmszcvy6l
type: challenge
title: Tips and Tricks
teaser: Some tips and recomended defaults
notes:
- type: text
  contents: Making the most of this template
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 300
---

💡 Tips and Tricks
==================

## Checking What Happened

As mentioned above, you can capture what went on during a challenge
using a pair of `tmux` commands. Here's an example of how you might
capture the entire history of a challenge:

```shell
# save the entire session to check user inputs and outputs
tmux capture-pane -t shell -S -
SESSION=$(tmux save-buffer -)
```

You may want to capture a subset of what was done. In this example,
we get the last ten lines.

```shell
# save the last _LINES_ lines to check inputs and outputs
LINES=10
HEIGHT=$(tmux list-panes -F "#{pane_height}")
SESSION=$(tmux capture-pane -t shell -S $(expr $HEIGHT - $LINES) -p)
```

In either case, you can use `grep` or either shell commands to examine
`SESSION` and see if they user followed you instructions.

## Entering Commands in Solve

Your solve command can run as usual and not interact with the learner's
shell. That's the recommended approach for most challenges. Some solve
scripts can be done that way, for example setting environment variables
you'll need in another challenge. In that case, you can use `send-keys`
in your solve script to change the environment.

```shell
tmux send-keys -t shell export SPACE REPLICATED_APP=wordpress-civet ENTER
```

Note that since the shell isn't restarted with each challenge, you'll
need to do this for any variable you want to persist in the environment
after the first challenge.

## Default Cleanup Script

The following is a good default cleanup script. Use this to make the
shell look clean and new just like it does in tracks that don't use
`tmux`.

```shell
#!/usr/bin/env bash

# This set line ensures that all failures will cause the script to error and exit
set -euxo pipefail

# clear the tmux pane and scrollback to look like a fresh shell
tmux clear-history -t shell
tmux send-keys -t shell clear ENTER
```

This is the cleanup script for this track, when you click **Check** to
move on to the next challenge you can see it's results.

## Testing

If you use any `tmux` commands in your lifecycle scripts, you will need
to make sure that the session is created if you want to run tests with
`instruqt track test`. This is necessary since the test lifecycle does not
run the shell commands in `config.yml`.

Put the following early in your `solve-shell` script for your first
challenge to make sure testing behaves.

```shell
### Assure the tmux session exists
#
# In a test scenario Instuqt does not run the user shell for the
# challenge, which means the tmux session is never established. We
# need to session for the solve scripts for other challenges to
# succeed, so let's create it here.
#

if ! tmux has-session -t shell ; then
  tmux new-session -d -s shell su - replicant
fi
```

🏁 Finish
=========

You've now had a bit of a tour through this template. You're ready to
base a lab on it. Feel free to browse through the source code to see
examples of these tips in action.

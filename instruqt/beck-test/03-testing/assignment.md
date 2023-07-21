---
slug: testing
id: un6bxcajkw47
type: challenge
title: Testing
teaser: A short description of the challenge.
notes:
- type: text
  contents: Replace this text with your own text
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---
Now that we’ve practiced some of the basics of working with the CM, let’s look at how it can be used for testing application changes prior to releasing them. To start this next exercise, we will be using our environment from the first exercise to test a possible upgrade.

1. Connect to GitHub
2. Make a simple change to the git project
3. Push it, and watch how CM picks up the change and spins up testing envs

We’ve made a very simple change that works on all of our nodes. What happens if one of the nodes fails testing? Let’s find out. In this exercise, we’re dropping support for OpenShift and deprecating EKS.

1. Add a preflight requiring one not OpenShift
2. Push to git
3. See failure at incorrect OS
4. See warning at EKS
5. See success at GKE


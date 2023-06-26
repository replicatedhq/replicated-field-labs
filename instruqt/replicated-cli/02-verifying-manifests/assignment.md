---
slug: verifying-manifests
id: zajj4jldzsgt
type: challenge
title: Verifying Manifests
teaser: Getting started with the Replicated CLI
notes:
- type: text
  contents: Let's see what the Replicated CLI can do
tabs:
- title: Shell
  type: terminal
  hostname: shell
difficulty: basic
timelimit: 600
---


🚀 Let's start
==============

We've created a set of manifests for you that configures
the "Kubernetes Up and Running" sample application to
deploy with Replicated. In the next few challenges we'll
validate those manifests and create a new application
you can deploy using Replicated.

### 1. Verifying manifests

You should have a few YAML files in `manifests`.

```text
$ ls -la manifests
total 28
drwxr-xr-x. 2 root root  161 Apr 12 18:27 .
drwxr-xr-x. 4 root root   94 Apr 12 18:27 ..
-rw-r--r--. 1 root root  179 Apr 12 18:27 k8s-app.yaml
-rw-r--r--. 1 root root 4186 Apr 12 18:27 kots-app.yaml
-rw-r--r--. 1 root root  990 Apr 12 18:27 kots-preflight.yaml
-rw-r--r--. 1 root root  347 Apr 12 18:27 kots-support-bundle.yaml
-rw-r--r--. 1 root root  447 Apr 12 18:27 nginx-deployment.yaml
-rw-r--r--. 1 root root  438 Apr 12 18:27 nginx-service.yaml
```
Since you're shipping them to your customers, you want to make
sure they'll install cleanly. Replicated includes a linter
to check that your installation files are correct. It also
checks for best practices for installing into customer clusters.

Let's verify these manifests with the linter, which is part of
the `replicated release` subcommand.

```shell script
replicated release lint --yaml-dir=manifests
```

You should get a list that returns no errors. It should have at least
one info message that looks something like the output. It's OK if you have more
info or warning messages, since we often update our linting rules to
capture more best practices.

```text
RULE                           TYPE    FILENAME                     LINE    MESSAGE
container-resource-requests    info    manifests/deployment.yaml    20      Missing resource requests
```

The command will also exit with a `0` status since there were no
errors. This helps you include linting in your CI/CD process or
local build commands and fail on an error. You can verify this if
you want.

```shell script
echo $?
```

### 2. Following the linter's advice

Let's make a quick change to the manifests to follow the advice
that the linter gave us. We're going to specify the resource
requests for our deployment so that it's clear what we expect
to have as available resources in the cluster. We have a pretty
simple application, so we won't ask for much.

Most times, you'll do this interactively in your favorite editor.
To simplify for the lab, we're going to fix it from the CLI to
keep things moving.

```
yq -i '.spec.template.spec.containers[0].resources.requests.cpu = "100m"' manifests/deployment.yaml
yq -i '.spec.template.spec.containers[0].resources.requests.memory = "64Mi"' manifests/deployment.yaml
```

### 3. Check your release again

Let's check the release again with these changes.

```shell script
replicated release lint --yaml-dir=manifests
```

You should get a list that returns no errors. It should no longer
show the information message we saw before. Like before, it's OK if
you have additional info or warning messages.

```
RULE    TYPE    FILENAME    LINE    MESSAGE
```



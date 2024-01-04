---
slug: proxying-private-images
type: challenge
title: Proxying Private Images
teaser: Protect your private container images with the Replicated proxy registry
notes:
- type: text
  contents: |
    Share your private images without exposting your private registry to your customers
tabs:
- title: Shell
  type: terminal
  hostname: shell
  workdir: /home/replicant
difficulty: basic
timelimit: 600
---

One of the core features of the Replicated Platform is it's proxy registry. The
proxy registry uses the Replicated license to control access to images that you
store in any other container registry. This relieves you of the burden of
managing authentication and authorization for the private images your
application depends on.

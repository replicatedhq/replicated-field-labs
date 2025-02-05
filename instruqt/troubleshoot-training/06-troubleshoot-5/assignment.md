---
slug: troubleshoot-5
id: blomhvawugyh
type: challenge
title: It can't be DNS...
teaser: It's always DNS
notes:
- type: text
  contents: Time to fix the problem...
tabs:
- id: g7eab4ov6vwd
  title: Workstation
  type: terminal
  hostname: cloud-client
difficulty: advanced
timelimit: 3600
enhanced_loading: null
---
A new issue has been reported saying that there are DNS resolution failures in some Pod logs.

The customer has provided a support bundle

How do you begin to troubleshoot the problem?

once you think you know the answer, run:

```run
quiz
```

💡 Hints
=================

- How do Pods resolve DNS names?

- Start by checking for any pods that may be failing or in a crash loop, and have a look at the pod logs.  You may want to use the `--previous` flag to see the logs from the previous instance of the Pod.

- Keep on the lookout for `tcp: lookup <hostname>: no such host`,  `getaddrinfo failed` or `Name or service not found` to confirm DNS resolution failures.

- Try to determine any patterns that may be present.  Does the problem affect a single Pod, multiple Pods, or all Pods?
  - Is there a pattern that affects only Pods on a specific Node or Namespace?

- If the behaviour affects only a single Pod, it might be a good idea to delete the Pod and let Kubernetes recreate it.  But, if the problem affects multiple Pods, it's more likely a problem in `coredns` or `kube-dns` itself.

- Review the [Debugging DNS Resolution](https://kubernetes.io/docs/tasks/administer-cluster/dns-debugging-resolution/) article from the Kubernetes project.


💡 More Hints
=================

- The DNS service in the cluster can be user-configured.  How would a cluster admin customize the DNS configs?

- The DNS zone for a Kubernetes cluster is expected to be `cluster.local`.  The fully-qualified domain name for a Service would be `<namespace>.svc.cluster.local`.

- You can verify if queries are being received by `coredns` by [configuring logging](https://kubernetes.io/docs/tasks/administer-cluster/dns-debugging-resolution/#are-dns-queries-being-received-processed).  Enable logging for `coredns` and then send some test queries.  What responses are logged on the server side?

💡 Even More Hints
=================

- an NXDOMAIN response is returned when a DNS query is made for a name that does not exist in the DNS zone.  This is a valid response, so DNS is _working_; if we expect the Kubernetes zone to be `cluster.local`, why are we getting `NXDOMAIN` responses for `cluster.local`?

✔️ Solution
=================

The `coredns` deployment has been reconfigured to only answer for a DNS zone of `cluster.nonlocal`.  This is causing DNS resolution failures for Pods that are expecting to resolve names in the `cluster.local` zone.

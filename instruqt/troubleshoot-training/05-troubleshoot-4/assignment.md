---
slug: troubleshoot-4
id: vccuaq9got6x
type: challenge
title: Everything is broke...
teaser: "\U0001F527"
notes:
- type: text
  contents: Time to fix the problem...
tabs:
- title: Workstation
  type: terminal
  hostname: cloud-client
difficulty: intermediate
timelimit: 3600
---
You get a new report from a customer saying that many pods are failing; some may display Errors, while others may be Evicted or Pending, or even in an Unknown state.

They have provided a cluster-down support bundle since they can't get one from the kots admin panel.

This is a type of support-bundle that collects data from the node itself rather than from inside the cluster.

Extract the support bundle with

```run
tar -xvf support-bundle.tar.gz
```

Explore it to determine the issue.

Once you think you have your answer, run:

```
quiz
```

💡 Hints
=================

Host support bundles have some key files:
- `analysis.json` contains the analysis results of the support bundle
- `host-collectors/` contains the raw output from the collectors
- `host-collectors/run-host/crictl-logs*` contain logs from important containers

💡 More Hints
=================

You can pull warnings from the analysis.json with a simple jq filter:

```run
cd /root/support-bundle-2024-04-17T09_50_09
jq '.[] | select(.insight.severity == "warn")' analysis.json
```

`host-collectors/run-host/df.txt` shows the output of running `df` on the host

✔️ Solution
=================

The node's disk is full, which causes problems in the operating system as well as with Kubernetes.  Pod eviction thresholds have been exceeded, causing pods to be evicted and image garbage collection to remove images from the node.

We have a Troubleshoot spec that looks for particularly large files at https://github.com/replicatedhq/troubleshoot-specs/blob/main/host/resource-contention.yaml#L36

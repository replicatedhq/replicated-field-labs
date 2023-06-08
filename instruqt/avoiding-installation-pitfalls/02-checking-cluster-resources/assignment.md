---
slug: checking-cluster-resources
id: iy5zaqt0tx9a
type: challenge
title: Checking Cluster Resources
teaser: Use preflight checks to validate minimum cluster requirements
notes:
- type: text
  contents: Making sure your cluster has sufficient resources to run your application
tabs:
- title: Shell
  type: terminal
  hostname: shell
- title: Manifest Editor
  type: code
  hostname: shell
  path: /home/replicant
difficulty: basic
timelimit: 300
---

Now that we know we're installing to a supported version of
Kubernetes, let's see if that cluster has the resources to
support running Harbor.

Cluster Resources
=================

The default `clusterResources` collector collects information
about all of the nodes in the cluster. This allows us to
write analyzers that check whether the cluster has sufficient
resources to run our cluster: most often we write checks to
determine whether CPU, memmory, and storage meet the base
requirements of the application.

When analyzing resources in the cluster, we can write expressions
baesd on whether the node has the
[capacity required and whether that capacity isallocatable](https://kubernetes.io/docs/concepts/architecture/nodes/#capacity).
Alloctable has a very specific meaning to Kubernetes, and is not
the same as "free" or "available". It means only that he capacity
is not being reserved by Kubernetes or the underlying system. This
distinction often trips up developer who are new to Kubernetes.

Verify Resource Requirements
============================

The best way to define your preflight checks for cluster resources
is to make sure they align with your documentation for the minimum
and recommended values. The preflight check makes those prerequisites
executable and lets your customer know whether there install will
succeed. Let's look at [Harbor's documentation](https://goharbor.io/docs/2.8.0/install-config/installation-prereqs/)
for guidance on our preflights.

<table>
<thead>
<tr>
<th>Resource</th>
<th>Minimum</th>
<th>Recommended</th>
</tr>
</thead>
<tbody>
<tr>
<td>CPU</td>
<td>2 CPU</td>
<td>4 CPU</td>
</tr>
<tr>
<td>Mem</td>
<td>4 GB</td>
<td>8 GB</td>
</tr>
<tr>
<td>Disk</td>
<td>40 GB</td>
<td>160 GB</td>
</tr>
</tbody>
</table>

Since we have both recommended and minimum values, we have thresholds
for both warning and failure. Let's add the CPU check into our
`harbor-preflights.yaml` manifest. Open up the "Manifest Editor" tab
and paste this new analyzer after the one checking the Kubernetes version.

```
    - nodeResources:
        checkName: Cluster CPU resources are sufficient to install Harbor
        outcomes:
          - fail:
              when: "sum(cpuAllocatable) < 2"
              message: |-
                Harbor requires a minimum of 2 CPU cores in order to run, and runs best with
                at 4 cores. Your current cluster has less than 2 CPU cores available to Kubernetes
                workloads. Please increase cluster capacity or install into a different cluster.
              uri: https://goharbor.io/docs/2.8.0/install-config/installation-prereqs/
          - warn:
              when: "sum(cpuAllocatable) < 4"
              message: |-
                Harbor runs best with a minimum of 4 CPU cores. Your current cluster has less
                than 4 CPU cores available to run Kubernetes workloads. You may want to consider
                increasing cluster capacity or installing into a different cluster for the best
                experience
              uri: https://goharbor.io/docs/2.8.0/install-config/installation-prereqs/
          - pass:
              message: Your cluster has sufficient CPU resources available to run Harbor
```

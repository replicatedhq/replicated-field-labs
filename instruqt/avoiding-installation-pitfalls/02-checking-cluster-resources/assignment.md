---
slug: checking-cluster-resources
id: iy5zaqt0tx9a
type: challenge
title: Checking Cluster Resources
teaser: Use preflight checks to validate minimum cluster requirements
notes:
- type: text
  contents: Making sure a cluster has sufficient resources to run your application
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
determine whether CPU, memory, and storage meet the base
requirements of the application.

When analyzing resources in the cluster, we can write expressions
based on whether the node has the
[capacity required and whether that capacity is allocatable](https://kubernetes.io/docs/concepts/architecture/nodes/#capacity).
Allocatable has a very specific meaning to Kubernetes, and is not
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
</tbody>
</table>

Since we have both recommended and minimum values, we have thresholds
for both warning and failure. Let's add the CPU check into our
`harbor-preflights.yaml` manifest. Open up the "Manifest Editor" tab
and paste this new analyzer after the one checking the Kubernetes version.

```
    - nodeResources:
        checkName: Cluster CPU resources are sufficient to install and run Harbor
        outcomes:
          - fail:
              when: "sum(cpuAllocatable) < 2"
              message: |-
                Harbor requires a minimum of 2 CPU cores in order to run, and runs best with
                at least 4 cores. Your current cluster has less than 2 CPU cores available to Kubernetes
                workloads. Please increase cluster capacity or install into a different cluster.
              uri: https://goharbor.io/docs/2.8.0/install-config/installation-prereqs/
          - warn:
              when: "sum(cpuAllocatable) < 4"
              message: |-
                Harbor runs best with a minimum of 4 CPU cores. Your current cluster has less
                than 4 CPU cores available to run workloads. For the best experience, consider
                increasing cluster capacity or installing into a different cluster.
              uri: https://goharbor.io/docs/2.8.0/install-config/installation-prereqs/
          - pass:
              message: Your cluster has sufficient CPU resources available to run Harbor
```

After saving your changes run the update preflight checks to see the outcome.

```
kubectl preflight ./harbor-preflights.yaml
```

You'll see that our cluster generates a warning since it has only two CPU
cores available. This should be fine for our lab environment, so we can
ignore the warning for now.

![CPU Preflight Warning](../assets/cpu-preflight-warning.png)

To round out the resource checks, add a similar check for memory

```
    - nodeResources:
        checkName: Cluster memory is sufficient to install and run Harbor
        outcomes:
          - fail:
              when: "sum(memoryAllocatable) < 4G"
              message: |-
                Harbor requires a minimum of 4 GB of memory in order to run, and runs best with
                at least 8 GB. Your current cluster has less than 4 GB available to Kubernetes
                workloads. Please increase cluster capacity or install into a different cluster.
              uri: https://goharbor.io/docs/2.8.0/install-config/installation-prereqs/
          - warn:
              when: "sum(memoryAllocatable) < 8Gi"
              message: |-
                Harbor runs best with a minimum of 8 GB of memory. Your current cluster has less
                than 8 GB of memory available to run workloads. For the best experience, consider
                increasing cluster capacity or installing into a different cluster.
              uri: https://goharbor.io/docs/2.8.0/install-config/installation-prereqs/
          - pass:
              message: Your cluster has sufficient memory available to run Harbor
```

Running the Revised Preflights
==============================

Now that we have a thorough set of preflights for cluster resources, let's run
them:

```
kubeclt preflight ./harbor-preflights.yaml
```

You'll see that all three preflights are run, and that the memory
preflight has failed. This is an expected failure, since we have
single node cluster that does not have enough memory to run Harbor.

![Failing Memory Preflight](../assets/memory-preflight-failure.png)


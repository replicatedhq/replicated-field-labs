# Support your own application with Replicated Support Bundles and Analyzers

## Introduction
The intent of this course is for customers to bring their own app into a controlled environment
where they can practice troubleshooting and support techniques.

## Deployment Architecture

In order to utilize Instruqt.io's DNS for sending back the user's admin console URL, we needed a machine


Hosts:
managed by Instruqt.io  | managed by sandbox instanced GCP
----------------------- | ---------------------------------
cloud-client            | cluster-node1
loadbalancer            | cluster-node2
                        | cluster-node3

The track is configured with 5 hosts:
- a `cloud-client` container that bootstraps a GCP project and resources to support kURL, provides a working environment for the user, and provides a jump host for SSH access to the kURL nodes.
- a `loadbalancer` VM that runs haproxy to route traffic to the kURL nodes
- 3 `cluster-node` VMs, provisioned in an instanced GCP project.  This GCP project doesn't share any data/integration with the managed sandbox provided by instruqt.io.
  - one additional GCP disk per node, attached to each VM in order to support Rook/Ceph

```
               [ ssh ]
cloud-client ----------> cluster-node{1,2,3}
( container )

              [ haproxy ]
loadbalancer -------------> cluster-node{1,2,3}:80
(   vm    )  |------------> cluster-node{1,2,3}:443
             |------------> cluster-node{1,2,3}:6443
             |------------> cluster-node{1,2,3}:8800
```

## Challenge lifecycle

https://docs.instruqt.com/concepts/lifecycle-scripts#instruqt-scripts-and-the-lifecycle

- user clicks "start" on a challenge
- Instruqt sandbox gets created in GCP, additional GCP projects created for the kURL nodes
- track_scripts fire for each machine, and the `cloud-client` machine bootstraps the kURL nodes over ssh.  the loadbalancer configures itself to route to the kURL nodes.
- challenge loop starts
  - setup runs first, and should set up the challenge
  - take backups of Kubernetes resources you might modify and save to /opt/backups - that makes it easy to undo your changes in the `solve` script if a user wants to skip the challenge
  - `check` scripts


## Developing new challenges

- add a challenge folder within this track folder `support-bundle-own-app-kurl`, in increasing number order.  the name of the folder after the first 2 digits gets used as the track title.
- create `assignment.md` and write the challenge.
  - top-level metadata provides challenge title, teaser, difficulty, and timelimit
- create challenge lifecycle scripts
  - https://docs.instruqt.com/concepts/lifecycle-scripts/helper-scripts
  - each track needs a setup, solve, and check script
  - instruqt manages DNS and provides environment variables for VMs in the `config.yaml` but *not* for things created in the sandbox GCP account
  - https://docs.instruqt.com/reference/instruqt-platform/networking#resolving-the-external-ip-of-a-sandbox-vm

### Lifecycle scripts tips
- `agent get variable` and `agent set variable` to store state between steps
  - only lifecycle scripts can use this feature, inaccessible to the user
  - shells are ephemeral and get destroyed when moving through challenge

- scripts should exit 0 on success, 1 on failure
  - `check` scripts `fail-message` is displayed to the user on a failed check
  - `debug-message` can send debug logs to the instruqt track logger

- instruqt provides some useful [environment variables](https://docs.instruqt.com/concepts/runtime-variables)

## Testing challenges

`instruqt track open` in the root of the track to open it in a browser

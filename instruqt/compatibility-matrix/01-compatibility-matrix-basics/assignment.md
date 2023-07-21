---
slug: compatibility-matrix 
id: 
type: challenge
title: compatibility-matrix
teaser: 
notes:
- type: text
  contents: Examines using the Replicated CLI to work with Compatibility Matrix
tabs:
- title: Shell
  type: terminal
  hostname: shell
  workdir: /home/replicant
- title: Cluster
  type: terminal
  hostname: cluster
difficulty: basic
timelimit: 3000
---

#### Let's look at using the Compatibility Matrix
The Replicated Compatibility Matrix is a service for quickly and easily spinning up ephemeral clusters. In this lab, we will learn some basics of how to work with the Compatibility Matrix: using it to troubleshooting existing environments, testing application changes in a comprehensive manner, and integrating it with existing CI/CD pipelines for automated testing. The Compatibility Matrix is part of our Builders Plan and can be added to your account at any time. 
	
	
Let’s start with how to create an environment using the Compatibility Matrix. In this exercise we will examine the different matrix commands and their uses as well as build and tear down a few environments for practice.	
	
1. Look at the new cli commands and run a few to see what’s up	
2. Get list of supported images (wc)
3. Create a cluster from scratch. Connect to it, then delete it.
4. Create another cluster with a very short (2 minute) duration. Verify it exists	
5. Create 3-node cluster with multiple OS.  (Keep this)
6. Install application in all three nodes.
7. Verify that the cluster from step 4 is gone, as the TTL has expired	
	

As we saw in the previous exercise, the CM can create a variety of environments. Now let’s look at how we might use it to troubleshoot existing customer installations without connecting to the live/production/airgap installations our customer is running. To start off, we will download a support bundle from an example customer and build a support environment that matches our customer’s. In this case, we have a customer who is not able to upgrade their environment from Kubernetes X to X (dependency issue)

1. The support bundle can be downloaded from the “Bundle” tab at the top of this lab. Go ahead and download it now.
2. Use the bundle to determine requisite versions 
3. Build a cluster from the bundle manually
4. Build a cluster from the bundle automatically (possible?)
5. Connect to the cluster and look at the errors the customer is seeing
6. Correct the error and update Kubernetes.

Now that we’ve practiced some of the basics of working with the CM, let’s look at how it can be used for testing application changes prior to releasing them. To start this next exercise, we will be using our environment from the first exercise to test a possible upgrade.

1. Connect to GitLab project *note: would it be better to use GitHub?
2. Make a simple change to the git project
3. Push it, and watch how CM picks up the change and spins up testing envs

We’ve made a very simple change that works on all of our nodes. What happens if one of the nodes fails testing? Let’s find out. In this exercise, we’re dropping support for OpenShift and deprecating EKS. 

1. Add a preflight requiring one not OpenShift
2. Push to git
3. See failure at incorrect OS
4. See warning at EKS
5. See success at GKE


```
kubectl get nodes
```

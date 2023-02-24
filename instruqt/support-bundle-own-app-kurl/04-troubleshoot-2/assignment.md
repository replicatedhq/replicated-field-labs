---
slug: troubleshoot-2
id: bbn7etmrxhet
type: challenge
title: Correcting the broken application
teaser: A Pod is not responding...
notes:
- type: text
  contents: Time to fix another problem...
tabs:
- title: Workstation
  type: terminal
  hostname: cloud-client
- title: Vendor Portal
  type: website
  url: https://vendor.replicated.com
  new_window: true
- title: Cluster Node 1
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-1
- title: Cluster Node 2
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-2
- title: Cluster Node 3
  type: terminal
  hostname: cloud-client
  cmd: ssh -oStrictHostKeyChecking=no kurl-node-3
difficulty: basic
timelimit: 3600
---
[App Installer Admin Console](http://loadbalancer.[[ Instruqt-Var key="SANDBOX_ID" hostname="cloud-client" ]].instruqt.io:8800)

🚀 Let's start
=================

You get another report from a customer saying that the application isn't working, as if **a Pod is not responding and connections time out**.  How would you begin to solve the problem?

- What component handles communication between Kubernetes workloads?
- What component handles communication between workloads and the outside world?

💡 Hints
=================

- The Kubernetes documentation has a [great manual on debugging Services](https://kubernetes.io/docs/tasks/debug/debug-application/debug-service/)

- Think about the traffic flow to your application
  - There are multiple *hops* in the network path, and any of them _could_ be a potential break in the path.  - Which hops can you identify?

- How does traffic get to workloads inside kubernetes
- How does Kubernetes handle DNS resolution and load balancing for Pods?

💡 More Hints
=================

- Kubernetes Services act as load balancers to Pods
  - Pods advertise a `containerPort` that they are listening on, but we don't want to keep track of their IP addresses since they change all the time.  Services are a way to abstract away the IP addresses of the Pods, and instead use a DNS name to connect to the Pods.
  - Services advertise a listening `port` and forward connections to a `targetPort`.

 Troubleshooting Procedure
=================

#### Understand the limits of the problem
Let's enumerate all the hops in our system.  We know there is a Pod, and a Service, and potentially an IngressController pod, and perhaps an outgoing proxy; so begin by figuring out the network path from the client to the Pod.  At each step, we can check our connection with something like `netcat` or `curl` depending on what kind of protocol we are using.  Perhaps the web frontend is not working, or perhaps a backend API is not working; start to understand the limits of the problem, and work backwards from Pods that are no longer responding to the client.

List all the pods and check for any that are not `Running` with `kubectl get pods -n <namespace> -o wide`.  If there are any Pods that are not `Running`, check the logs with `kubectl logs -n <namespace> <pod-name>`.  If the Pod is `Running`, check the logs with `kubectl logs -n <namespace> <pod-name>`.  Note the IP address of any affected Pods.

Check to make sure that Services exist with `kubectl get svc -n <namespace>`.  We can see the Type of the service which tells us how we can expect to connect to this service; for example, a `ClusterIP` service is only accessible from within the cluster, and a `NodePort` service is accessible from outside the cluster too.  Describe the Service with `kubectl describe svc`, and check that the `targetPort` is correctly set to the `containerPort` in your Pod spec.

#### Peel the onion

We will start to debug the problem from the inside out.  # TODO: insert onion model diagram here , Container -> Pod -> Service -> Ingress -> Load Balancer -> Client

First, we want to ensure the the application in the Pod itself is running and responding.  This may not be necessary if your pod is configured with HealthChecks that will restart the pod if it is not responding correctly.  If you have a pod that is not responding, but is not restarting, you may need to troubleshoot the problem.  Try `kubectl exec` to enter the Pod and run a shell.  From here we can use tools like `ps`, `netstat` or `ss`, `curl` to make sure the application is up and responding correctly.

```
kubectl exec -n <namespace> -it <pod-name> -- /bin/sh
```

If the image you're using doesn't have a shell or package manager to install command line tools, that's OK - versions of Kubernetes above 1.25 support ephemeral debug containers - you can use `kubectl debug` to create a debug container in the Pod based on a more permissive image like `alpine` or `busybox` or [`netshoot`](https://github.com/nicolaka/netshoot)

```
kubectl debug -n <namespace> -it <pod-name> --image=nicolaka/netshoot --share-processes
```

Once you can confirm that the application is responding, let's move on to the next hop in the network path, which would be across the CNI.

If you have another working Pod in the cluster with tools like `curl` or `netcat` in the base image, or has a package manager that can install these tools, we can use `kubectl exec` to exec into another Pod and test connections the affected Pod.  If not, we can use an image that has these tools installed, like `busybox` `ubuntu` or `netshoot`.  We can use `kubectl run` to create a throwaway Pod:

```
kubectl run tmp-shell --rm -i --tty --image nicolaka/netshoot
kubectl run tmp-shell --rm -i --tty --image busybox
kubectl run tmp-shell --rm -i --tty --image ubuntu
```

From the Pod shell, use `curl` or `netcat` to test the connection to the affected Pod IP.  If the connection fails here, we have a problem with the CNI.  If the connection succeeds, we can move on to the next hop in the network path, the Service.

```
# For applications that speak HTTP
#
curl -vvvv -sSL <pod-ip>:<container-port>
#
# -vvvv is verbose output
# -sSL is silent, show errors, follow redirects

# For applications that expect a TCP connection (like MySQL)
#
nc -vz <pod-ip> <container-port>
#
# -v is verbose output
# -z is zero I/O mode, just check if the port is open
```

From the same Pod shell, first let's check that cluster DNS is working by resolving the Service name:

```
dig <service-name>
dig <service-name>.<namespace>.svc.cluster.local

# or
nslookup <service-name>
nslookup <service-name>.<namespace>.svc.cluster.local
```

Expect both the short name and the FQDN of the service to resolve to the same IP address.  If not, there is a problem with DNS resolution.

Next, let's check that the Service is working by connecting to the Service name:

```
# For applications that speak HTTP
#
curl -vvvv -sSL <service-name>:<service-port>
curl -vvvv -sSL <service-name>.<namespace>.svc.cluster.local:<service-port>
#

# For applications that expect a TCP connection (like MySQL)
#
nc -vz <service-name> <service-port>
nc -vz <service-name>.<namespace>.svc.cluster.local <service-port>
#
```

Expect both connections to the short name and the FQDN of the service to succeed.  If not, there is a problem with the Service configuration or the `kube-proxy` component of the cluster.

I expect that the problem in this challenge is visible now, but the same technique can be applied to any network path in the cluster.  As long as you are able to make a connection to the next hop in the network path, you can move on to testing access through an IngressController Pod, an infrastructure loadbalancer, etc.

Some `curl` examples that may be useful, particularly for testing IngressControllers and other layer 7 loadbalancers:

```
# connect to a server and send a "Host:" header
curl -vvvv -SsL -H "Host: service.domain.com" http://<ingress-ip>:80



# connect to a server that expects a TLS connection and override DNS resolution for SNI
# useful to fake a connection through an IngressController or other layer 7 loadbalancer since the "Host:" header cannot be inspected
curl -vvvv -kSsL --resolve <service-fqdn>:<service-port>:<ingress-ip> https://<service-fqdn>:<service-port>
```

Reference:
[Debugging Kubernetes Services](https://kubernetes.io/docs/tasks/debug/debug-application/debug-service/)
[curl name resolution tricks](https://everything.curl.dev/usingcurl/connections/name)



✔️ Solution
=================

A random Service's `targetPort` has been patched to be something in the 30k range.  Any pod in the cluster that tries to connect to this pod's Service name (which is what gets programmed in to DNS in the cluster) will fail to connect, because the pod is not listening on the the same port that the Service is trying to connect to.

Remediation
=================

Patch or edit the affected service to correct the port number. you may have to refer to the other resources in the cluster to identify the correct port number.


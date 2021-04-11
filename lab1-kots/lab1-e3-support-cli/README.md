Lab 1.3: Support CLI
=========================================

In this lab, we'll learn how to debug and diagnose support problems when 

You can open the KOTS admin console your your node by navigating to https://$IP_ADDRESS:8800 in a browser. The password to your instance will be provided as part of the lab, or you can reset by SSHing the node and running

```shell
kubectl kots reset-password -n default
```

### Ground Rules

In this lab and most of those that follow it, some of the failure scenarios are quite contrived.
It is very possible to reverse-engineer the solution by reading the Kubernetes YAML rather instead of following the lab steps.
If you want to get the most of out these labs, use the presented debugging steps to get experience with the toolset.

### Investigating

- app not running
- collect bundle, 502
- collect bundle, 401






When the KOTS admin console is 

NAME                                  READY   STATUS        RESTARTS   AGE
file-check-pod-5fb558b75b-djltv       1/1     Running       0          5m25s
kotsadm-589555b5c7-6c96r              0/1     Init:0/4      0          2s
kotsadm-589555b5c7-t78q8              1/1     Terminating   0          53s
kotsadm-operator-674545cbb6-66xfp     1/1     Running       0          6m50s
kotsadm-postgres-0                    1/1     Running       0          6m49s
kurl-proxy-kotsadm-5bd9b6956d-c8xpn   1/1     Running       0          6m48s
nginx-8b679cd99-zmv2w                 0/1     Init:2/3      0          5m25s


CLI Bundle

```shell
kubectl support-bundle secret/default/kotsadm-dx411-dex-supportbundle --redactors=configmap/default/kotsadm-redact-spec/redact-spec,configmap/default/kotsadm-dx411-dex-redact-spec/redact-spec
```

- normal install
- fix normal config.txt problem (again)
- kotsadm crashing, collect CLI support bundle 
- understand new file needed, fix issue
- add preflight check and analyzer  
  
- 



```shell
$ kubectl get pod
NAME                                  READY   STATUS        RESTARTS   AGE
file-check-pod-6799b757fb-gf2gn       1/1     Running       0          6m26s
kotsadm-794468644d-7rhml              0/1     Init:0/4      0          7s
kotsadm-794468644d-hdrgz              1/1     Terminating   0          42s
kotsadm-operator-674545cbb6-4zdfr     1/1     Running       0          9m59s
kotsadm-postgres-0                    1/1     Running       0          9m59s
kurl-proxy-kotsadm-5bd9b6956d-qpmhk   1/1     Running       0          9m58s
nginx-5d6b75bd99-mj7nb                0/1     Init:2/3      0          47s
nginx-67c9547d89-wv5pv                1/1     Running       0          4m
```










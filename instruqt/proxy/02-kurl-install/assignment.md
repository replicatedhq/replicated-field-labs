---
slug: kurl-install
id: c0g5jayumshi
type: challenge
title: kURL install with proxy config
teaser: Install kURL cluster via proxy server
notes:
- type: text
  contents: |
    # Install kURL using a proxy server
tabs:
- title: Host
  type: terminal
  hostname: isolated-host
difficulty: basic
timelimit: 3600
---

💡 Shell
=========

Check proxy env variables present
```
env | grep proxy
```

Download kurl installer
```
curl --proxy http://${LOCAL_PROXY_IP}:3128 https://kurl.sh/proxy-$INSTRUQT_PARTICIPANT_ID -o /root/kurl-install.sh
```

Replace instruqt local hostname for IP address in installer file (needed as not an fqdn for kubernetes DNS to work later)
```
sed -i 's/proxy-host/${LOCAL_PROXY_IP}/' /root/kurl-install.sh
```

Install kurl kubernetes cluster
note the below command uses 'screen' to run the install in the background which protects from session timeouts on laptops etc.
```
screen -d -m bash -c 'time bash /root/kurl-install.sh 2>&1 | tee -a /root/kurl-install.log'
```

note: this will take around 12 minutes to complete.

View the deployment log while installing..
```
tail -f /root/kurl-install.log
```

Note the proxy settings picked up by the kurl installer:
```
cat /root/kurl-install.log | grep proxy
```

When the installation has completed, test kubernetes is running
```
kubectl get nodes
```

and kotsadm services are deployed and running
```
kubectl get all
```

Change the default randomly generated kotsadm password to a known value
```
echo mytestapp | /usr/bin/kubectl kots --kubeconfig=/etc/kubernetes/admin.conf reset-password -n default
```

🏁 Finish
==========

To complete the challenge, press **Next**.

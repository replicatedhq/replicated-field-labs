---
slug: proxy-config
id: rkd0axhzc3me
type: challenge
title: Proxy Config
teaser: Test connectivity and configure proxy
notes:
- type: text
  contents: |
    # Test external network access and configure proxy

    This track environment has an isolated deployment host that can access external sites using a proxy server.
tabs:
- title: Host
  type: terminal
  hostname: isolated-host
difficulty: basic
timelimit: 1200
---

💡 Shell
=========

Test that the host has outbound traffic blocked.

Test blocked website:
```bash
curl -v -m 5 https://lwn.net 2>&1 | egrep 'Trying|HTTP/|Failed|error|unreachable'
```

Test website with proxy config:
```
curl -v -m 5 --proxy http://proxy-host:3128 https://lwn.net 2>&1 | egrep 'Trying|HTTP/|Failed|error'
```

Ping blocked test:
```
ping -c 3 google.com
```

Package manager already configured with proxy as usually packages available in internal repos.
```
apt-get install -y tree
```

Set the proxy environment variables and retest site access"
```
export http_proxy=http://proxy-host:3128
export https_proxy=http://proxy-host:3128
```

Test website access now works without the additional --proxy parameter::
```bash
curl -v -m 5 https://lwn.net 2>&1 | egrep 'Trying|HTTP/|Failed|error|unreachable'
```


🏁 Next
=======

To complete the challenge, press **Next**.

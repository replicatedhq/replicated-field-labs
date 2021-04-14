Lab 1.6: Proxy
=========================================


- Exact same as airgap but with a authenticated squid proxy




SSH via jump box


Validate no internet

```shell
curl https://kubernetes.io
```

```shell
curl curl -x kots-field-labs-squid-proxy:3128 https://kubernetes.io
curl -x kots-field-labs-squid-proxy:3128 https://api.replicated.com/market/v1/echo/ip
```


```shell
export HTTP_PROXY=kots-field-labs-squid-proxy:3128
curl https://api.replicated.com/market/v1/echo/ip
# doesn't work, need HTTPS
export HTTPS_PROXY=kots-field-labs-squid-proxy:3128
curl https://api.replicated.com/market/v1/echo/ip
```
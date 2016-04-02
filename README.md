# Remote Proxy

This is just purpurin sprinkled on top of golang's [httputil.NewSingleHostReverseProxy](https://golang.org/pkg/net/http/httputil/#NewSingleHostReverseProxy).

It takes a list (map) of domains and ports from a yaml and reverse proxies them (port 80). No tls support. 


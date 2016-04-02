package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"gopkg.in/yaml.v2"
)

var proxies map[string]string

func init() {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, &proxies); err != nil {
		panic(err)
	}
}

func main() {

	for path, port := range proxies {
		vhost, err := url.Parse("http://127.0.0.1:" + port)
		if err != nil {
			panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(vhost)
		http.HandleFunc(path+"/", handler(proxy))
	}

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}

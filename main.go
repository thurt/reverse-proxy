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
		http.Handle(path+"/", &ProxyHandler{proxy})
	}

	err := http.ListenAndServe("", nil)
	if err != nil {
		panic(err)
	}
}

type ProxyHandler struct {
	p *httputil.ReverseProxy
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// add customization here (if desired)

	ph.p.ServeHTTP(w, r)
}

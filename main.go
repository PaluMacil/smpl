package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	dwn, err := url.Parse("http://127.0.0.1:3035")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(dwn)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}

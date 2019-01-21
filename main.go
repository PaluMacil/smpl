package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/mholt/certmagic"
)

func main() {
	dwn, err := url.Parse("http://127.0.0.1:3035")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	proxy := httputil.NewSingleHostReverseProxy(dwn)
	mux.HandleFunc("/", handler(proxy))

	certmagic.Agreed = true
	certmagic.Email = "dcwolf@gmail.com"
	certmagic.CA = certmagic.LetsEncryptProductionCA
	err = certmagic.HTTPS([]string{"danwolf.net"}, mux)
	if err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		p.ServeHTTP(w, r)
	}
}

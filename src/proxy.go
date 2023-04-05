package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func NewProxy(targetUrl string) (*httputil.ReverseProxy, error) {
	target, err := url.Parse(targetUrl)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	return proxy, nil
}

func NewHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = mux.Vars(r)["target"]
		proxy.ServeHTTP(w, r)
	}
}

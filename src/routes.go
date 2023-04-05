package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/msw-x/moon/ulog"
)

func routes(routesFile string) *mux.Router {
	log := ulog.New("routes")
	r := mux.NewRouter()
	r.HandleFunc("/a", func(http.ResponseWriter, *http.Request) {

	})
	rules, err := loadRules(routesFile)
	if err == nil {
		for _, rule := range rules {
			log.Info(rule)
			proxy, err := NewProxy(rule.Target)
			if err == nil {
				r.HandleFunc(rule.Source+"{target:.*}", NewHandler(proxy))
			} else {
				log.Error(err)
			}
		}
	} else {
		log.Error(err)
	}
	return r
}

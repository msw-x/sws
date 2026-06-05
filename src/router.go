package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/msw-x/moon/ulog"
)

func router(uiDir, routesFile string) *mux.Router {
	log := ulog.New("router")
	r := mux.NewRouter()
	if routesFile != "" {
		log.Info("routes:", routesFile)
		rules, err := loadRules(routesFile)
		if err == nil {
			for _, rule := range rules {
				log.Info(rule)
				proxy, err := NewProxy(rule.Target, log)
				if err == nil {
					r.HandleFunc(rule.Source+"{target:.*}", NewHandler(proxy))
				} else {
					log.Error(err)
				}
			}
		} else {
			log.Error(err)
		}
	}
	if uiDir != "" {
		log.Info("static:", uiDir)
		r.PathPrefix("/").Handler(http.FileServer(http.FS(os.DirFS(uiDir)))).Methods(http.MethodGet)
	}
	return r
}

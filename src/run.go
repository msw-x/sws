package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/msw-x/moon/app"
)

func run(conf Conf) {
	if certHost, ok := os.LookupEnv("CERT_HOST"); ok {
		conf.CertHost = certHost
	}
	log.SetOutput(ioutil.Discard)
	srv := New(conf)
	defer srv.Close()
	app.WaitInterrupt()
}

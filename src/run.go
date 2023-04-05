package main

import (
	"io/ioutil"
	"log"

	"github.com/msw-x/moon/app"
)

func run(conf Conf) {
	log.SetOutput(ioutil.Discard)
	srv := New(conf)
	defer srv.Close()
	app.WaitInterrupt()
}

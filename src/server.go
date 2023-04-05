package main

import "github.com/msw-x/moon/webs"

type Server struct {
	s *webs.DualServer
}

func New(conf Conf) *Server {
	o := new(Server)
	o.s = webs.NewDual().WithSecretDir(conf.CertDir)
	if conf.CertHost != "" {
		o.s.WithRedirectToTls(conf.CertHost)
		o.s.WithAutoSecret(conf.CertDir, conf.CertHost)
	}
	o.s.Run(conf.Http, conf.Https, routes(conf.RoutesFile))
	return o
}

func (o *Server) Close() {
	o.s.Shutdown()
}

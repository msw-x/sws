package main

type Conf struct {
	Http        string
	Https       string
	CertDir     string
	CertHost    string
	RoutesFile  string
	LogRequests bool
	LogErrors   string
}

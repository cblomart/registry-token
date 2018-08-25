package main

import (
	"flag"

	"net/http"

	"github.com/golang/glog"
)

func init() {
	flag.Parse()
}

func main() {
	glog.Infof("Starting registry-token-ldap server")
	http.HandleFunc("/"+folder, HandleAuth)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		glog.Errorf("error starting server: %s", err)
	}
}

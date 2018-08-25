package main

import (
	"flag"
	"fmt"

	"net/http"

	"github.com/golang/glog"
	"github.com/jinzhu/configor"
)

// Config holds the global configuration
var Config Config

func init() {
	flag.Parse()
	// read config file
	configor.Load(&Config, "/etc/registry-token-ldap.yml")
}

func main() {
	glog.Infof("Starting registry-token-ldap server")
	http.HandleFunc("/"+folder, HandleAuth)
	err := http.ListenAndServe(fmt.Sprintf(":%d", Config.Port), nil)
	if err != nil {
		glog.Errorf("error starting server: %s", err)
	}
}

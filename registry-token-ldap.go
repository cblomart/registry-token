package main

import (
	"flag"
	"fmt"
	"os"

	"net/http"

	"github.com/golang/glog"
	"github.com/jinzhu/configor"
)

// AuthConfig holds the global configuration
var AuthConfig Config

func init() {
	flag.Parse()
	// read config file
	err := configor.Load(&AuthConfig, "/etc/registry-token-ldap/config.yml")
	if err != nil {
		glog.Errorf("Error loading config: %s", err)
		return
	}
	// check file presence
	if _, err := os.Stat(AuthConfig.JWSCert); os.IsNotExist(err) {
		glog.Errorf("JWS certificate does not exist: %s", AuthConfig.JWSCert)
		return
	}
	if _, err := os.Stat(AuthConfig.JWSKey); os.IsNotExist(err) {
		glog.Errorf("JWS certificate does not exist: %s", AuthConfig.JWSKey)
		return
	}
}

func main() {
	glog.Infof("Starting registry-token-ldap server")
	http.HandleFunc("/"+AuthConfig.Path, HandleAuth)
	glog.Infof("Listening on port %d for auth on /%s", AuthConfig.Port, AuthConfig.Path)
	err := http.ListenAndServe(fmt.Sprintf(":%d", AuthConfig.Port), nil)
	if err != nil {
		glog.Errorf("error starting server: %s", err)
	}
}

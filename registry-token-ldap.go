package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"net/http"

	"github.com/docker/libtrust"
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
	// check CA presence
	if len(AuthConfig.LDAPCa) > 0 {
		if _, err := os.Stat(AuthConfig.LDAPCa); err != nil {
			glog.Errorf("Could not generate private key: %s", err)
			panic(err)
		}
	}
	// check file presence
	if _, err := os.Stat(AuthConfig.JWSKey); os.IsNotExist(err) {
		glog.Infof("JWS key does not exist: %s", AuthConfig.JWSKey)
		GenerateKey("RSA4096")
	}
	if _, err := os.Stat(AuthConfig.JWSCert); os.IsNotExist(err) {
		glog.Infof("JWS Certificate does not exist: %s", AuthConfig.JWSCert)

	}
	_, err = libtrust.LoadKeyFile(AuthConfig.JWSKey)
	if err != nil {
		glog.Errorf("Could not read key file: %s", err)
		panic(err)
	}
	certs, err := libtrust.LoadCertificateBundle(AuthConfig.JWSCert)
	if err != nil {
		glog.Errorf("Could not read certificate file: %s", err)
		panic(err)
	}
	if len(certs) == 0 {
		glog.Errorf("No certificates read from certificate")
		panic("No certificates read from certificate")
	}
	if certs[0].NotAfter.Unix() < time.Now().UTC().Unix() {
		glog.Errorf("Certificate is not valid anymore")
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

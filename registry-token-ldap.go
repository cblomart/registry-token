package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"encoding/pem"
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
	// check file presence
	if _, err := os.Stat(AuthConfig.JWSKey); os.IsNotExist(err) {
		glog.Infof("JWS key does not exist: %s", AuthConfig.JWSKey)
		glog.Infof("Generating a new private key.")
		privkey, err := libtrust.GenerateRSA4096PrivateKey()
		if err != nil {
			glog.Errorf("Could not generate private key: %s", err)
			panic(err)
		}
		err = libtrust.SaveKey(AuthConfig.JWSKey, privkey)
		if err != nil {
			glog.Errorf("Could not save private key: %s", err)
			panic(err)
		}
		glog.Infof("Generating private key saved to %s", AuthConfig.JWSKey)
	}
	if _, err := os.Stat(AuthConfig.JWSCert); os.IsNotExist(err) {
		glog.Infof("JWS Certificate does not exist: %s", AuthConfig.JWSCert)
		glog.Infof("Generating new certificate")
		privkey, err := libtrust.LoadKeyFile(AuthConfig.JWSKey)
		if err != nil {
			glog.Errorf("Could not load private key: %s", err)
			panic(err)
		}
		cert, err := libtrust.GenerateSelfSignedClientCert(privkey)
		if err != nil {
			glog.Errorf("Could gnerate certificate: %s", err)
			panic(err)
		}
		certout, err := os.Create(AuthConfig.JWSCert)
		if err != nil {
			glog.Errorf("Could not create certificate file: %s", err)
			panic(err)
		}
		err = pem.Encode(certout, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})
		if err != nil {
			glog.Errorf("Could not encode certificate to pem file: %s", err)
			panic(err)
		}
		err = certout.Close()
		if err != nil {
			glog.Errorf("Could not close certificate file: %s", err)
			panic(err)
		}
		glog.Infof("Certificate saved to %s", AuthConfig.JWSCert)
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
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", AuthConfig.Port), AuthConfig.JWSCert, AuthConfig.JWSKey, nil)
	if err != nil {
		glog.Errorf("error starting server: %s", err)
	}
}

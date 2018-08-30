package main

import (
	"encoding/pem"
	"fmt"
	"os"

	"github.com/docker/libtrust"
	"github.com/golang/glog"
)

// GenerateKey generates a SSL key
func GenerateKey(alg string) {
	glog.Infof("Generating a new private key.")
	var privkey libtrust.PrivateKey
	switch alg {
	case "ECP256":
		key, err := libtrust.GenerateECP256PrivateKey()
		if err != nil {
			glog.Errorf("Could not generate %s private key: %s", alg, err)
			panic(err)
		}
		privkey = key
	case "ECP384":
		key, err := libtrust.GenerateECP384PrivateKey()
		if err != nil {
			glog.Errorf("Could not generate %s private key: %s", alg, err)
			panic(err)
		}
		privkey = key
	case "ECP521":
		key, err := libtrust.GenerateECP521PrivateKey()
		if err != nil {
			glog.Errorf("Could not generate %s private key: %s", alg, err)
			panic(err)
		}
		privkey = key
	case "RSA2048":
		key, err := libtrust.GenerateRSA2048PrivateKey()
		if err != nil {
			glog.Errorf("Could not generate %s private key: %s", alg, err)
			panic(err)
		}
		privkey = key
	case "RSA3072":
		key, err := libtrust.GenerateRSA3072PrivateKey()
		if err != nil {
			glog.Errorf("Could not generate %s private key: %s", alg, err)
			panic(err)
		}
		privkey = key
	case "RSA4096":
		key, err := libtrust.GenerateRSA4096PrivateKey()
		if err != nil {
			glog.Errorf("Could not generate %s private key: %s", alg, err)
			panic(err)
		}
		privkey = key
	default:
		glog.Errorf("Alogrithm %s not supported", alg)
		panic(fmt.Errorf("Alogrithm %s not supported", alg))
	}
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

// GenerateCert generates a SSL certificate
func GenerateCert() {
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

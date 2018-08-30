package main

import (
	"testing"
	"os"
	"io/ioutil"

	"github.com/docker/libtrust"
)

func TestGenerateKey(t *testing.T) {
	type args struct {
		alg string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ECP256",
			args: args{
				alg: "ECP256",
			},
		},
		{
			name: "ECP384",
			args: args{
				alg: "ECP384",
			},
		},
		{
			name: "ECP521",
			args: args{
				alg: "ECP521",
			},
		},
		{
			name: "RSA2048",
			args: args{
				alg: "RSA2048",
			},
		},
		{
			name: "RSA3072",
			args: args{
				alg: "RSA3072",
			},
		},
		{
			name: "RSA4096",
			args: args{
				alg: "RSA4096",
			},
		},
	}
	AuthConfig = Config{
		JWSKey: "/tmp/cert.key",
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateKey(tt.args.alg)
			_, err := libtrust.LoadKeyFile(AuthConfig.JWSKey)
			if err != nil {
				t.Errorf("GenerateKey() error = %v",err)
			}
		})
	}
	os.Remove(AuthConfig.JWSKey)
	AuthConfig = Config{}
	
}

func TestGenerateCert(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Cert",
		},
	}
	AuthConfig = Config{
		JWSKey: "/tmp/cert.key",
		JWSCert: "/tmp/cert.crt",
	}
	ioutil.WriteFile("/tmp/cert.key", []byte(key), 0600)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateCert()
			certs, err := libtrust.LoadCertificateBundle(AuthConfig.JWSCert)
			if err != nil {
				t.Errorf("GenerateCert() error = %v",err)
			}
			if len(certs) == 0 {
				t.Errorf("GenerateCert() = empty")
			}
		})
	}
	os.Remove(AuthConfig.JWSKey)
	os.Remove(AuthConfig.JWSCert)
	AuthConfig = Config{}
}
package main

import (
	"flag"

	"github.com/golang/glog"
)

func init() {
	flag.Parse()
}

func main() {
	glog.Infof("Starting registry-token-ldap server")
}

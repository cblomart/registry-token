package main

import (
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

// HandleAuth Authenticates and resturns a token.
func HandleAuth(w http.ResponseWriter, r *http.Request) {
	glog.Infof("Call to authentication endpoint")
	fmt.Fprintf(w, "Test page.")
}

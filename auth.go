package main

import (
	"fmt"
	"net/http"
)

// HandleAuth Authenticates and resturns a token.
func HandleAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test page.")
}

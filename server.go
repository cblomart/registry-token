package main

import (
	"net/http"
	"time"

	"github.com/golang/glog"
)

// TokenResponse represents the response structure
type TokenResponse struct {
	Token       string `json:"token"`
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	IssuedAt    string `json:"issued_at,omitempty"`
}

// PasswordString hides passwords on output
type PasswordString string

func (s PasswordString) String() string {
	if len(s) == 0 {
		return ""
	}
	return "*"
}

// HandleAuth Authenticates and resturns a token.
func HandleAuth(w http.ResponseWriter, r *http.Request) {
	glog.Infof("Call to authentication endpoint")
	// check parameters
	anr := GetAuthRequest(r)
	if anr == nil {
		glog.Error("No Authentication request returned")
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}
	glog.Infof("Authentication Request: %s", anr.String())
	azr, ok := Authenticate(anr.UserName, string(anr.Password))
	if !ok {
		glog.Infof("User %s not authenticated", anr.UserName)
		http.Error(w, "Not authorized", 401)
		return
	}
	glog.Infof("User %s authenticated check authorizations", anr.UserName)
	jti, err := GenerateJTI()
	if err != nil {
		glog.Infof("Failed to generate unique token id")
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}
	accesses := Scopes{}
	if len(anr.Scopes) == 0 {
		glog.Infof("Authenticating user %s with no scopes: returning empty token", anr.UserName)
		accesses = Authorize(azr, anr.Scopes)
	}
	iat := time.Now().UTC()
	token, err := GenerateToken(accesses, anr.Service, anr.UserName, iat, jti)
	if err != nil {
		glog.Errorf("Could not generate token: %s", err)
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}
	glog.Infof("token: %s", token)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(GenerateTokenResponse(token, iat))
	if err != nil {
		glog.Errorf("Could write repsonse: %s", err)
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}
}

// GenerateTokenResponse generate the json repsonse for a token request
func GenerateTokenResponse(token string, iat time.Time) []byte {
	response := TokenResponse{
		Token:       token,
		AccessToken: token,
		ExpiresIn:   TokenValidity,
		IssuedAt:    iat.Format(time.RFC3339),
	}
	return MustMarshal(response)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

// AuthnRequest respresent an authentication request
type AuthnRequest struct {
	RemoteAddr string
	UserName   string
	Password   PasswordString
	ClientID   string
	Service    string
	Scopes     Scopes
}

// GetAuthRequest gets an authentication request from the http request
func GetAuthRequest(r *http.Request) *AuthnRequest {
	ar := AuthnRequest{}
	ar.RemoteAddr = r.RemoteAddr
	if username, password, ok := r.BasicAuth(); ok {
		ar.UserName = username
		ar.Password = PasswordString(password)
	} else {
		glog.Errorf("Authentication request didn't have basic authentication")
		return nil
	}
	service := r.FormValue("service")
	if len(service) == 0 {
		glog.Errorf("No service provided")
		return nil
	}
	ar.Service = service
	clientid := r.FormValue("client_id")
	if len(clientid) == 0 {
		glog.Error("No ClientId provided")
	} else {
		ar.ClientID = clientid
	}
	scopeString := r.FormValue("scope")
	if len(scopeString) == 0 {
		glog.Infof("No scopes provided")
	}
	ar.Scopes = *GetScopes(scopeString)
	return &ar
}

func (ar *AuthnRequest) String() string {
	return fmt.Sprintf("%s:%s - ip='%s' client_id='%s' service='%s' scopes='%s'", ar.UserName, ar.Password, ar.RemoteAddr, ar.ClientID, ar.Service, ar.Scopes)
}

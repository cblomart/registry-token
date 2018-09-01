package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestGetAuthRequest(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want *AuthnRequest
	}{
		{
			name: "full request",
			args: args{
				r: &http.Request{
					RemoteAddr: "127.0.0.1:54657",
					Header: http.Header{
						"Authorization": []string{
							fmt.Sprintf("Basic %s", base64.URLEncoding.EncodeToString([]byte("username:password"))),
						},
					},
					Form: url.Values{
						"service":   []string{"registry.docker.io"},
						"client_id": []string{"Docker"},
						"scope":     []string{"repository:cblomart/foo:latest:pull,push"},
					},
				},
			},
			want: &AuthnRequest{
				RemoteAddr: "127.0.0.1:54657",
				UserName:   "username",
				Password:   PasswordString("password"),
				ClientID:   "Docker",
				Service:    "registry.docker.io",
				Scopes: Scopes{
					Scope{
						Type:    "repository",
						Name:    "cblomart/foo:latest",
						Actions: []string{"pull", "push"},
					},
				},
			},
		},
		{
			name: "request without scope",
			args: args{
				r: &http.Request{
					RemoteAddr: "127.0.0.1:54657",
					Header: http.Header{
						"Authorization": []string{
							fmt.Sprintf("Basic %s", base64.URLEncoding.EncodeToString([]byte("username:password"))),
						},
					},
					Form: url.Values{
						"service": []string{"registry.docker.io"},
					},
				},
			},
			want: &AuthnRequest{
				RemoteAddr: "127.0.0.1:54657",
				UserName:   "username",
				Password:   PasswordString("password"),
				Service:    "registry.docker.io",
				Scopes:     Scopes{},
			},
		},
		{
			name: "request without client id",
			args: args{
				r: &http.Request{
					RemoteAddr: "127.0.0.1:54657",
					Header: http.Header{
						"Authorization": []string{
							fmt.Sprintf("Basic %s", base64.URLEncoding.EncodeToString([]byte("username:password"))),
						},
					},
					Form: url.Values{
						"service": []string{"registry.docker.io"},
						"scope":   []string{"repository:cblomart/foo:latest:pull,push"},
					},
				},
			},
			want: &AuthnRequest{
				RemoteAddr: "127.0.0.1:54657",
				UserName:   "username",
				Password:   PasswordString("password"),
				Service:    "registry.docker.io",
				Scopes: Scopes{
					Scope{
						Type:    "repository",
						Name:    "cblomart/foo:latest",
						Actions: []string{"pull", "push"},
					},
				},
			},
		},
		{
			name: "request without service",
			args: args{
				r: &http.Request{
					RemoteAddr: "127.0.0.1:54657",
					Header: http.Header{
						"Authorization": []string{
							fmt.Sprintf("Basic %s", base64.URLEncoding.EncodeToString([]byte("username:password"))),
						},
					},
					Form: url.Values{
						"client_id": []string{"Docker"},
						"scope":     []string{"repository:cblomart/foo:latest:pull,push"},
					},
				},
			},
			want: nil,
		},
		{
			name: "empty request",
			args: args{
				r: &http.Request{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAuthRequest(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthnRequest_String(t *testing.T) {
	tests := []struct {
		name string
		ar   *AuthnRequest
		want string
	}{
		{
			name: "auth request to string",
			ar: &AuthnRequest{
				RemoteAddr: "127.0.0.1:54657",
				UserName:   "username",
				Password:   PasswordString("password"),
				ClientID:   "Docker",
				Service:    "registry.docker.io",
				Scopes: Scopes{
					Scope{
						Type:    "repository",
						Name:    "cblomart/foo:latest",
						Actions: []string{"pull", "push"},
					},
				},
			},
			want: "username:* - ip='127.0.0.1:54657' client_id='Docker' service='registry.docker.io' scopes='repository:cblomart/foo:latest:pull,push'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ar.String(); got != tt.want {
				t.Errorf("AuthnRequest.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

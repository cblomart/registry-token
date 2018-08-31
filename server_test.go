package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestScope_String(t *testing.T) {
	tests := []struct {
		name string
		s    Scope
		want string
	}{
		{
			name: "full access to simple repository",
			s: Scope{
				Type:    "repository",
				Name:    "sample",
				Actions: []string{"push", "pull"},
			},
			want: "repository:sample:push,pull",
		},
		{
			name: "pull access to simple repository",
			s: Scope{
				Type:    "repository",
				Name:    "sample",
				Actions: []string{"pull"},
			},
			want: "repository:sample:pull",
		},
		{
			name: "push access to simple repository",
			s: Scope{
				Type:    "repository",
				Name:    "sample",
				Actions: []string{"push"},
			},
			want: "repository:sample:push",
		},
		{
			name: "no access to simple repository",
			s: Scope{
				Type:    "repository",
				Name:    "sample",
				Actions: []string{},
			},
			want: "repository:sample:",
		},
		{
			name: "full access to tagged repository",
			s: Scope{
				Type:    "repository",
				Name:    "sample:latest",
				Actions: []string{"push", "pull"},
			},
			want: "repository:sample:latest:push,pull",
		},
		{
			name: "pull access to tagged repository",
			s: Scope{
				Type:    "repository",
				Name:    "sample:latest",
				Actions: []string{"pull"},
			},
			want: "repository:sample:latest:pull",
		},
		{
			name: "push access to tagged repository",
			s: Scope{
				Type:    "repository",
				Name:    "sample:latest",
				Actions: []string{"push"},
			},
			want: "repository:sample:latest:push",
		},
		{
			name: "no access to tagged repository",
			s: Scope{
				Type:    "repository",
				Name:    "sample:latest",
				Actions: []string{},
			},
			want: "repository:sample:latest:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Scope.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScopes_String(t *testing.T) {
	tests := []struct {
		name string
		ss   Scopes
		want string
	}{
		{
			name: "one repo",
			ss: []Scope{
				Scope{
					Type:    "repository",
					Name:    "sample:latest",
					Actions: []string{"pull,push"},
				},
			},
			want: "repository:sample:latest:pull,push",
		},
		{
			name: "two repos",
			ss: []Scope{
				Scope{
					Type:    "repository",
					Name:    "sample:latest",
					Actions: []string{"pull,push"},
				},
				Scope{
					Type:    "repository",
					Name:    "foo",
					Actions: []string{"push"},
				},
			},
			want: "repository:sample:latest:pull,push repository:foo:push",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ss.String(); got != tt.want {
				t.Errorf("Scopes.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPasswordString_String(t *testing.T) {
	tests := []struct {
		name string
		s    PasswordString
		want string
	}{
		{
			name: "a password",
			s:    PasswordString("@P@ssw0rd"),
			want: "*",
		},
		{
			name: "no password",
			s:    PasswordString(""),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("PasswordString.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestGetScope(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Scope
	}{
		{
			name: "full access to simple repository",
			args: args{
				s: "repository:sample:pull,push",
			},
			want: &Scope{
				Type:    "repository",
				Name:    "sample",
				Actions: []string{"pull", "push"},
			},
		},
		{
			name: "pull access to simple repository",
			args: args{
				s: "repository:sample:pull",
			},
			want: &Scope{
				Type:    "repository",
				Name:    "sample",
				Actions: []string{"pull"},
			},
		},
		{
			name: "push access to simple repository",
			args: args{
				s: "repository:sample:push",
			},
			want: &Scope{
				Type:    "repository",
				Name:    "sample",
				Actions: []string{"push"},
			},
		},
		// BUG: doesn't work for go access but string seems the same...
		{
			name: "no access to simple repository",
			args: args{
				s: "repository:sample:",
			},
			want: &Scope{
				Type: "repository",
				Name: "sample",
			},
		},
		{
			name: "full access to tagged repository",
			args: args{
				s: "repository:cblomart/foo:latest:pull,push",
			},
			want: &Scope{
				Type:    "repository",
				Name:    "cblomart/foo:latest",
				Actions: []string{"pull", "push"},
			},
		},
		{
			name: "pull access to tagged repository",
			args: args{
				s: "repository:cblomart/foo:latest:pull",
			},
			want: &Scope{
				Type:    "repository",
				Name:    "cblomart/foo:latest",
				Actions: []string{"pull"},
			},
		},
		{
			name: "push access to tagged repository",
			args: args{
				s: "repository:cblomart/foo:latest:push",
			},
			want: &Scope{
				Type:    "repository",
				Name:    "cblomart/foo:latest",
				Actions: []string{"push"},
			},
		},
		// BUG: doesn't work for go access but string seems the same...
		{
			name: "no access to simple repository",
			args: args{
				s: "repository:cblomart/foo:latest:",
			},
			want: &Scope{
				Type: "repository",
				Name: "cblomart/foo:latest",
			},
		},
		{
			name: "too few fields",
			args: args{
				s: "cblomart/foo:pull,push",
			},
			want: nil,
		},
		{
			name: "too many fields",
			args: args{
				s: "repository:cblomart/foo:latest:pull,push:hey",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetScope(tt.args.s)
			if got != nil && tt.want == nil {
				t.Errorf("GetScope() = %v, want nil", got)
			} else if got != nil && tt.want != nil {
				if !reflect.DeepEqual(got.String(), tt.want.String()) {
					t.Errorf("GetScope() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestHandleAuth(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleAuth(tt.args.w, tt.args.r)
		})
	}
}

func TestGetScopes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Scopes
	}{
		// BUG: doesn't work for go access but string seems the same...
		{
			name: "one scope",
			args: args{
				s: "repository:sample:latest:pull,push",
			},
			want: &Scopes{
				Scope{
					Type:    "repository",
					Name:    "sample:latest",
					Actions: []string{"pull,push"},
				},
			},
		},
		// BUG: doesn't work for go access but string seems the same...
		{
			name: "two scope",
			args: args{
				s: "repository:sample:latest:pull,push repository:foo:push",
			},
			want: &Scopes{
				Scope{
					Type:    "repository",
					Name:    "sample:latest",
					Actions: []string{"pull,push"},
				},
				Scope{
					Type:    "repository",
					Name:    "foo",
					Actions: []string{"push"},
				},
			},
		},
		{
			name: "two scope with one not parseable",
			args: args{
				s: "repository:sample:latest:pull,push cblomart/foo:push",
			},
			want: &Scopes{
				Scope{
					Type:    "repository",
					Name:    "sample:latest",
					Actions: []string{"pull,push"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetScopes(tt.args.s); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("GetScopes() = %v, want %v", got, tt.want)
			}
		})
	}
}

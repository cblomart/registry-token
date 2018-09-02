package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

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

func TestHandleAuth(t *testing.T) {
	ioutil.WriteFile("/tmp/cert.key", []byte(key), 0600)
	ioutil.WriteFile("/tmp/cert.crt", []byte(cert), 0600)
	type args struct {
		r *http.Request
		c Config
	}
	tests := []struct {
		name         string
		args         args
		wantCode     int
		wantResponse string
	}{
		/*{
			name: "request all authorized all",
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
				c: Config{
					JWSCert: "/tmp/cert.crt",
					JWSKey:  "/tmp/cert.key",
					Issuer:  "auth.dokcer.io",
					Rules: []Rule{
						Rule{Match: "${user}/.*", Actions: []string{"pull", "push"}},
					},
				},
			},
			wantCode:     200,
			wantResponse: "",
		},
		{
			name: "request all authorized pull",
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
				c: Config{
					JWSCert: "/tmp/cert.crt",
					JWSKey:  "/tmp/cert.key",
					Issuer:  "auth.dokcer.io",
					Rules: []Rule{
						Rule{Match: "${user}/.*", Actions: []string{"pull"}},
					},
				},
			},
			wantCode:     200,
			wantResponse: "",
		},
		{
			name: "request all authorized push",
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
				c: Config{
					JWSCert: "/tmp/cert.crt",
					JWSKey:  "/tmp/cert.key",
					Issuer:  "auth.dokcer.io",
					Rules: []Rule{
						Rule{Match: "${user}/.*", Actions: []string{"push"}},
					},
				},
			},
			wantCode:     200,
			wantResponse: "",
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
				c: Config{
					JWSCert: "/tmp/cert.crt",
					JWSKey:  "/tmp/cert.key",
					Issuer:  "auth.dokcer.io",
					Rules: []Rule{
						Rule{Match: "${user}/.*", Actions: []string{"pull", "push"}},
					},
				},
			},
			wantCode:     200,
			wantResponse: "",
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
				c: Config{
					JWSCert: "/tmp/cert.crt",
					JWSKey:  "/tmp/cert.key",
					Issuer:  "auth.dokcer.io",
					Rules: []Rule{
						Rule{Match: "${user}/.*", Actions: []string{"pull", "push"}},
					},
				},
			},
			wantCode:     200,
			wantResponse: "",
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
				c: Config{
					JWSCert: "/tmp/cert.crt",
					JWSKey:  "/tmp/cert.key",
					Issuer:  "auth.dokcer.io",
					Rules: []Rule{
						Rule{Match: "${user}/.*", Actions: []string{"pull", "push"}},
					},
				},
			},
			wantCode:     200,
			wantResponse: "",
		},
		{
			name: "empty request",
			args: args{
				r: &http.Request{},
				c: Config{
					JWSCert: "/tmp/cert.crt",
					JWSKey:  "/tmp/cert.key",
					Issuer:  "auth.dokcer.io",
					Rules: []Rule{
						Rule{Match: "${user}/.*", Actions: []string{"pull", "push"}},
					},
				},
			},
			wantCode:     200,
			wantResponse: "",
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			AuthConfig = tt.args.c
			HandleAuth(rec, tt.args.r)
			if rec.Code != tt.wantCode {
				t.Errorf("HandleAuth() Code = %v, want %v", rec.Code, tt.wantCode)
			}
			if rec.Body.String() != tt.wantResponse {
				t.Errorf("HandleAuth() Response = %v, want %v", rec.Body.String(), tt.wantResponse)
			}
			AuthConfig = Config{}
		})
	}
	os.Remove(AuthConfig.JWSKey)
	os.Remove(AuthConfig.JWSCert)
}

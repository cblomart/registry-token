package main

import (
	"reflect"
	"testing"
)

func TestJosePart(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  string
		panic bool
	}{
		{
			name: "test jose header",
			args: args{
				v: &Header{
					Type:       "JWT",
					SigningAlg: "RS256",
					KeyID:      "Random ID Key",
				},
			},
			want:  "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IlJhbmRvbSBJRCBLZXkifQ",
			panic: false,
		},
		{
			name: "test jose claims",
			args: args{
				v: &ClaimSet{
					Issuer:     "issuer",
					Subject:    "subject",
					Audience:   "audience",
					Expiration: 1535721792,
					NotBefore:  1535720892,
					IssuedAt:   1535720892,
					JWTID:      "random bytes",
					Access:     Scopes{},
				},
			},
			want:  "eyJpc3MiOiJpc3N1ZXIiLCJzdWIiOiJzdWJqZWN0IiwiYXVkIjoiYXVkaWVuY2UiLCJleHAiOjE1MzU3MjE3OTIsIm5iZiI6MTUzNTcyMDg5MiwiaWF0IjoxNTM1NzIwODkyLCJqdGkiOiJyYW5kb20gYnl0ZXMifQ",
			panic: false,
		},
		{
			name: "test byte array",
			args: args{
				v: []byte("random bytes"),
			},
			want:  "cmFuZG9tIGJ5dGVz",
			panic: false,
		},
		{
			name: "test an unsupported object",
			args: args{
				v: Scope{
					Type:    "repository",
					Name:    "sample",
					Actions: []string{"pull", "push"},
				},
			},
			panic: true,
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.panic {
				if got := JosePart(tt.args.v); got != tt.want {
					t.Errorf("JosePart() = %v, want %v", got, tt.want)
				}
			} else {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("JosePart() wants panic")
					}
				}()
				JosePart(tt.args.v)
			}
		})
	}
}

func TestJoseBase64UrlEncode(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoseBase64UrlEncode(tt.args.b); got != tt.want {
				t.Errorf("JoseBase64UrlEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustMarshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		panic bool
	}{

		{
			name: "test jose header",
			args: args{
				v: &Header{
					Type:       "JWT",
					SigningAlg: "RS256",
					KeyID:      "Random ID Key",
				},
			},
			want:  []byte("{\"typ\":\"JWT\",\"alg\":\"RS256\",\"kid\":\"Random ID Key\"}"),
			panic: false,
		},
		{
			name: "test jose claims",
			args: args{
				v: &ClaimSet{
					Issuer:     "issuer",
					Subject:    "subject",
					Audience:   "audience",
					Expiration: 1535721623,
					NotBefore:  1535720723,
					IssuedAt:   1535720723,
					JWTID:      "random bytes",
					Access:     Scopes{},
				},
			},
			want:  []byte("{\"iss\":\"issuer\",\"sub\":\"subject\",\"aud\":\"audience\",\"exp\":1535721623,\"nbf\":1535720723,\"iat\":1535720723,\"jti\":\"random bytes\"}"),
			panic: false,
		},
		{
			name: "test byte array",
			args: args{
				v: []byte("random bytes"),
			},
			want:  []byte("\"cmFuZG9tIGJ5dGVz\""),
			panic: false,
		},
		{
			name: "test an unsupported object",
			args: args{
				v: make(chan int),
			},
			panic: true,
			want:  []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.panic {
				if got := MustMarshal(tt.args.v); !reflect.DeepEqual(string(got), string(tt.want)) {
					t.Errorf("MustMarshal() = %v, want %v", string(got), string(tt.want))
				}
			} else {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("JosePart() wants panic")
					}
				}()
				MustMarshal(tt.args.v)
			}
		})
	}
}

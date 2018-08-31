package main

import (
	"reflect"
	"testing"
	"time"
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
			want: "",
		},
		{
			name: "test jose claims",
			args: args{
				v: &ClaimSet{
					Issuer:     "issuer",
					Subject:    "subject",
					Audience:   "audience",
					Expiration: time.Now().UTC().Unix() + TokenValidity,
					NotBefore:  time.Now().UTC().Unix(),
					IssuedAt:   time.Now().UTC().Unix(),
					JWTID:      "random bytes",
					Access:     Accesses{},
				},
			},
			want: "",
		},
		{
			name: "test byte array",
			args: args{
				v: []byte("random bytes"),
			},
			want: "",
		},
		{
			name: "test an unsupported object",
			args: args{
				v: Access{
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
			}
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("JosePart() wants panic")
				}
			}()
			JosePart(tt.args.v)
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
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustMarshal(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

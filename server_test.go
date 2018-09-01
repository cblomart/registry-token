package main

import (
	"net/http"
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

package main

import (
	"reflect"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	type args struct {
		user     string
		password string
	}
	tests := []struct {
		name  string
		args  args
		want  AuthzRequest
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Authenticate(tt.args.user, tt.args.password)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Authenticate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

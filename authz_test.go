package main

import (
	"reflect"
	"testing"
)

func TestRule_Eval(t *testing.T) {
	type args struct {
		user   string
		group  string
		scope  Scope
		access *Scope
	}
	tests := []struct {
		name string
		r    *Rule
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Eval(tt.args.user, tt.args.group, tt.args.scope, tt.args.access)
		})
	}
}

func TestAuthorize(t *testing.T) {
	type args struct {
		request AuthzRequest
		scopes  []Scope
	}
	tests := []struct {
		name string
		args args
		want Scopes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Authorize(tt.args.request, tt.args.scopes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authorize() = %v, want %v", got, tt.want)
			}
		})
	}
}

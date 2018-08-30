package main

import (
	"reflect"
	"testing"
)

func TestAccess_String(t *testing.T) {
	tests := []struct {
		name string
		a    Access
		want string
	}{
		{
			name: "full access to simple repository",
			a: Access{
				Type:    "repository",
				Name:    "sample",
				Actions: []string{"push", "pull"},
			},
			want: "repository:sample:push,pull",
		},
		{
			name: "pull access to simple repository",
			a: Access{
				Type:    "repository",
				Name:    "sample",
				Actions: []string{"pull"},
			},
			want: "repository:sample:pull",
		},
		{
			name: "push access to simple repository",
			a: Access{
				Type:    "repository",
				Name:    "sample",
				Actions: []string{"push"},
			},
			want: "repository:sample:push",
		},
		{
			name: "no access to simple repository",
			a: Access{
				Type:    "repository",
				Name:    "sample",
				Actions: []string{},
			},
			want: "repository:sample:",
		},
		{
			name: "full access to tagged repository",
			a: Access{
				Type:    "repository",
				Name:    "sample:latest",
				Actions: []string{"push", "pull"},
			},
			want: "repository:sample:latest:push,pull",
		},
		{
			name: "pull access to tagged repository",
			a: Access{
				Type:    "repository",
				Name:    "sample:latest",
				Actions: []string{"pull"},
			},
			want: "repository:sample:latest:pull",
		},
		{
			name: "push access to tagged repository",
			a: Access{
				Type:    "repository",
				Name:    "sample:latest",
				Actions: []string{"push"},
			},
			want: "repository:sample:latest:push",
		},
		{
			name: "no access to tagged repository",
			a: Access{
				Type:    "repository",
				Name:    "sample:latest",
				Actions: []string{},
			},
			want: "repository:sample:latest:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("Access.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccesses_String(t *testing.T) {
	tests := []struct {
		name string
		as   Accesses
		want string
	}{
		{
			name: "one repo",
			as: []Access{
				Access{
					Type:    "repository",
					Name:    "sample:latest",
					Actions: []string{"pull,push"},
				},
			},
			want: "repository:sample:latest:pull,push",
		},
		{
			name: "two repos",
			as: []Access{
				Access{
					Type:    "repository",
					Name:    "sample:latest",
					Actions: []string{"pull,push"},
				},
				Access{
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
			if got := tt.as.String(); got != tt.want {
				t.Errorf("Accesses.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRule_Eval(t *testing.T) {
	type args struct {
		user   string
		group  string
		scope  Scope
		access *Access
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
		want Accesses
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

func TestGetAccess(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Access
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAccess(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAccesses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Accesses
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAccesses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccesses() = %v, want %v", got, tt.want)
			}
		})
	}
}

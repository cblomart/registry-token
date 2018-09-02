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
		want string
	}{
		{
			name: "full user access to specific reposistory",
			r:    &Rule{User: "john", Group: "", Match: "", Actions: []string{"push", "pull"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "people/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:people/foo:pull,push",
		},
		{
			name: "push user access to specific reposistory",
			r:    &Rule{User: "john", Group: "", Match: "", Actions: []string{"push"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "people/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:people/foo:push",
		},
		{
			name: "pull user access to specific reposistory",
			r:    &Rule{User: "john", Group: "", Match: "", Actions: []string{"pull"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "people/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:people/foo:pull",
		},
		{
			name: "deny user access to specific reposistory",
			r:    &Rule{User: "john", Group: "", Match: "", Actions: []string{}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "people/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:people/foo:",
		},
		{
			name: "full user access to group repositories",
			r:    &Rule{User: "john", Group: "admin", Match: "admins/.*", Actions: []string{"push", "pull"}},
			args: args{
				user:   "john",
				group:  "admin",
				scope:  Scope{Type: "repository", Name: "admins/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:admins/foo:pull,push",
		},
		{
			name: "push user access to group repositories",
			r:    &Rule{User: "john", Group: "admin", Match: "admins/.*", Actions: []string{"push"}},
			args: args{
				user:   "john",
				group:  "admin",
				scope:  Scope{Type: "repository", Name: "admins/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:admins/foo:push",
		},
		{
			name: "pull user access to group reposistories",
			r:    &Rule{User: "john", Group: "admin", Match: "admins/.*", Actions: []string{"pull"}},
			args: args{
				user:   "john",
				group:  "admin",
				scope:  Scope{Type: "repository", Name: "admins/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:admins/foo:pull",
		},
		{
			name: "deny user access to group reposistories",
			r:    &Rule{User: "john", Group: "admin", Match: "admins/.*", Actions: []string{}},
			args: args{
				user:   "john",
				group:  "admin",
				scope:  Scope{Type: "repository", Name: "admnins/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:admins/foo:",
		},
		{
			name: "full user access to dynamic reposistory",
			r:    &Rule{User: "john", Group: "", Match: "${user}/.*", Actions: []string{"push", "pull"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "john/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:people/foo:pull,push",
		},
		{
			name: "push user access to dynamic reposistory",
			r:    &Rule{User: "john", Group: "", Match: "${user}/.*", Actions: []string{"push"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "john/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:people/foo:push",
		},
		{
			name: "pull user access to dynamic reposistory",
			r:    &Rule{User: "john", Group: "", Match: "${user}/.*", Actions: []string{"pull"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "john/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:people/foo:pull",
		},
		{
			name: "deny user access to dynamic reposistory",
			r:    &Rule{User: "john", Group: "", Match: "${user}/.*", Actions: []string{}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "john/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:people/foo:",
		},
		{
			name: "full group access to dynamic repositories",
			r:    &Rule{User: "john", Group: "project", Match: "${group}/.*", Actions: []string{"push", "pull"}},
			args: args{
				user:   "john",
				group:  "project",
				scope:  Scope{Type: "repository", Name: "project/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:project/foo:pull,push",
		},
		{
			name: "push group access to dynamic repositories",
			r:    &Rule{User: "john", Group: "project", Match: "${group}/.*", Actions: []string{"push"}},
			args: args{
				user:   "john",
				group:  "project",
				scope:  Scope{Type: "repository", Name: "project/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:project/foo:push",
		},
		{
			name: "pull group access to dynamic reposistories",
			r:    &Rule{User: "john", Group: "project", Match: "${group}/.*", Actions: []string{"pull"}},
			args: args{
				user:   "john",
				group:  "project",
				scope:  Scope{Type: "repository", Name: "project/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:project/foo:pull",
		},
		{
			name: "deny group access to dynamic reposistories",
			r:    &Rule{User: "john", Group: "project", Match: "${group}/.*", Actions: []string{}},
			args: args{
				user:   "john",
				group:  "project",
				scope:  Scope{Type: "repository", Name: "project/foo", Actions: []string{"pull", "push"}},
				access: &Scope{},
			},
			want: "repository:project/foo:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Eval(tt.args.user, tt.args.group, tt.args.scope, tt.args.access)
			if tt.args.access.String() == tt.want {
				t.Errorf("Eval() = %v, want %v", tt.args.access.String(), tt.want)
			}
		})
	}
}

func TestAuthorize(t *testing.T) {
	type args struct {
		request AuthzRequest
		scopes  string
		rules   []Rule
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "general pull access and user access tp",
			args: args{
				request: AuthzRequest{
					User:   "cblomart",
					Groups: []string{},
				},
				rules: []Rule{
					Rule{User: "", Group: "", Match: "${user}/.*", Actions: []string{"push"}},
					Rule{User: "", Group: "", Match: ".*", Actions: []string{"pull"}},
				},
				scopes: "repository:cblomart/foo:pull,push",
			},
			want: "repository:cblomart/foo:pull,push",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Authorize(tt.args.request, tt.args.scopes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authorize() = %v, want %v", got, tt.want)
			}
		})
	}
}

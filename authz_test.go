package main

import (
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
				access: &Scope{Type: "repository", Name: "people/foo"},
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
				access: &Scope{Type: "repository", Name: "people/foo"},
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
				access: &Scope{Type: "repository", Name: "people/foo"},
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
				access: &Scope{Type: "repository", Name: "people/foo"},
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
				access: &Scope{Type: "repository", Name: "admins/foo"},
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
				access: &Scope{Type: "repository", Name: "admins/foo"},
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
				access: &Scope{Type: "repository", Name: "admins/foo"},
			},
			want: "repository:admins/foo:pull",
		},
		{
			name: "deny user access to group reposistories",
			r:    &Rule{User: "john", Group: "admin", Match: "admins/.*", Actions: []string{}},
			args: args{
				user:   "john",
				group:  "admin",
				scope:  Scope{Type: "repository", Name: "admins/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "admins/foo"},
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
				access: &Scope{Type: "repository", Name: "john/foo"},
			},
			want: "repository:john/foo:pull,push",
		},
		{
			name: "full user access to dynamic reposistory denied as wrong user",
			r:    &Rule{User: "john", Group: "", Match: "${user}/.*", Actions: []string{"push", "pull"}},
			args: args{
				user:   "jane",
				group:  "",
				scope:  Scope{Type: "repository", Name: "john/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "john/foo"},
			},
			want: "repository:john/foo:",
		},
		{
			name: "push user access to dynamic reposistory",
			r:    &Rule{User: "john", Group: "", Match: "${user}/.*", Actions: []string{"push"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "john/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "john/foo"},
			},
			want: "repository:john/foo:push",
		},
		{
			name: "pull user access to dynamic reposistory",
			r:    &Rule{User: "john", Group: "", Match: "${user}/.*", Actions: []string{"pull"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "john/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "john/foo"},
			},
			want: "repository:john/foo:pull",
		},
		{
			name: "deny user access to dynamic reposistory",
			r:    &Rule{User: "john", Group: "", Match: "${user}/.*", Actions: []string{}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "john/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "john/foo"},
			},
			want: "repository:john/foo:",
		},
		{
			name: "full group access to dynamic repositories",
			r:    &Rule{User: "john", Group: "project", Match: "${group}/.*", Actions: []string{"push", "pull"}},
			args: args{
				user:   "john",
				group:  "project",
				scope:  Scope{Type: "repository", Name: "project/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "project/foo"},
			},
			want: "repository:project/foo:pull,push",
		},
		{
			name: "full group access to dynamic repositories denied as wrong group",
			r:    &Rule{User: "john", Group: "project", Match: "${group}/.*", Actions: []string{"push", "pull"}},
			args: args{
				user:   "john",
				group:  "otherproject",
				scope:  Scope{Type: "repository", Name: "project/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "project/foo"},
			},
			want: "repository:project/foo:",
		},
		{
			name: "push group access to dynamic repositories",
			r:    &Rule{User: "john", Group: "project", Match: "${group}/.*", Actions: []string{"push"}},
			args: args{
				user:   "john",
				group:  "project",
				scope:  Scope{Type: "repository", Name: "project/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "project/foo"},
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
				access: &Scope{Type: "repository", Name: "project/foo"},
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
				access: &Scope{Type: "repository", Name: "project/foo"},
			},
			want: "repository:project/foo:",
		},
		{
			name: "full user access to specific reposistory",
			r:    &Rule{User: "john", Group: "", Match: "", Actions: []string{"push", "pull"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "namespace", Name: "people/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "namespace", Name: "people/foo"},
			},
			want: "namespace:people/foo:",
		},
		{
			name: "full user access to specific reposistory with pull access already granted",
			r:    &Rule{User: "john", Group: "", Match: "", Actions: []string{"push", "pull"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "people/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "people/foo", Actions: []string{"pull"}},
			},
			want: "repository:people/foo:pull,push",
		},
		{
			name: "full user access to specific reposistory with push access already granted",
			r:    &Rule{User: "john", Group: "", Match: "", Actions: []string{"push", "pull"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "people/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "people/foo", Actions: []string{"push"}},
			},
			want: "repository:people/foo:pull,push",
		},
		{
			name: "full user access to specific reposistory but scope diffrent from access",
			r:    &Rule{User: "john", Group: "", Match: "", Actions: []string{"push", "pull"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "people/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "john/foo", Actions: []string{}},
			},
			want: "repository:john/foo:",
		},
		{
			name: "specific user access",
			r:    &Rule{User: "john", Group: "", Match: ".*", Actions: []string{"push", "pull"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "john/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "john/foo", Actions: []string{}},
			},
			want: "repository:john/foo:pull,push",
		},
		{
			name: "specific group access",
			r:    &Rule{User: "", Group: "all", Match: ".*", Actions: []string{"push", "pull"}},
			args: args{
				user:   "cedric",
				group:  "all",
				scope:  Scope{Type: "repository", Name: "john/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "john/foo", Actions: []string{}},
			},
			want: "repository:john/foo:pull,push",
		},
		{
			name: "specific user access not matching",
			r:    &Rule{User: "john", Group: "", Match: ".*", Actions: []string{"push", "pull"}},
			args: args{
				user:   "cedric",
				group:  "",
				scope:  Scope{Type: "repository", Name: "john/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "john/foo", Actions: []string{}},
			},
			want: "repository:john/foo:",
		},
		{
			name: "specific group access not matching",
			r:    &Rule{User: "", Group: "all", Match: ".*", Actions: []string{"push", "pull"}},
			args: args{
				user:   "cedric",
				group:  "project",
				scope:  Scope{Type: "repository", Name: "john/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "john/foo", Actions: []string{}},
			},
			want: "repository:john/foo:",
		},
		{
			name: "invalid regex",
			r:    &Rule{User: "", Group: "", Match: "${user}/[0-", Actions: []string{"push", "pull"}},
			args: args{
				user:   "john",
				group:  "",
				scope:  Scope{Type: "repository", Name: "people/foo", Actions: []string{"pull", "push"}},
				access: &Scope{Type: "repository", Name: "people/foo"},
			},
			want: "repository:people/foo:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Eval(tt.args.user, tt.args.group, tt.args.scope, tt.args.access)
			if tt.args.access.String() != tt.want {
				t.Errorf("Eval() = '%v', want '%v'", tt.args.access.String(), tt.want)
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
		{
			name: "general pull access and user access tp",
			args: args{
				request: AuthzRequest{
					User:   "cblomart",
					Groups: []string{"project", "all"},
				},
				rules: []Rule{
					Rule{User: "", Group: "", Match: "${user}/.*", Actions: []string{"push"}},
					Rule{User: "", Group: "", Match: "${group}/.*", Actions: []string{"push"}},
					Rule{User: "", Group: "", Match: ".*", Actions: []string{"pull"}},
				},
				scopes: "repository:cblomart/foo:pull,push repository:project/foo:pull,push repository:john/test:pull,push",
			},
			want: "repository:cblomart/foo:pull,push repository:project/foo:pull,push repository:john/test:pull",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AuthConfig = Config{
				Rules: tt.args.rules,
			}
			if got := Authorize(tt.args.request, *GetScopes(tt.args.scopes)); got.String() != tt.want {
				t.Errorf("Authorize() = %v, want %v", got, tt.want)
			}
			AuthConfig = Config{}
		})
	}
}

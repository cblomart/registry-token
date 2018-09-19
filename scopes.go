package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/golang/glog"
)

// Scope defined the required resources and actions
type Scope struct {
	Type    string   `json:"type"`
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}

func (s Scope) String() string {
	sort.Strings(s.Actions)
	return fmt.Sprintf("%s:%s:%s", s.Type, s.Name, strings.Join(s.Actions, ","))
}

// Scopes represents a set of scopes
type Scopes []Scope

func (ss Scopes) String() string {
	scopes := ""
	for _, s := range ss {
		scopes = scopes + " " + s.String()
	}
	return strings.Trim(scopes, " ")
}

// GetScope gets the scope from a string
func GetScope(s string) *Scope {
	scope := Scope{}
	parts := strings.Split(s, ":")
	switch len(parts) {
	case 3:
		scope.Type = parts[0]
		scope.Name = parts[1]
		scope.Actions = strings.Split(parts[2], ",")
	case 4:
		scope.Type = parts[0]
		scope.Name = fmt.Sprintf("%s:%s", parts[1], parts[2])
		scope.Actions = strings.Split(parts[3], ",")
	default:
		return nil
	}
	return &scope
}

// GetScopes gets scopes from a string
func GetScopes(s string) *Scopes {
	// trim input
	s = strings.Trim(s, " ")
	// split scopes
	ss := strings.Split(s, " ")
	scopes := Scopes{}
	for _, v := range ss {
		scope := GetScope(v)
		if scope != nil {
			scopes = append(scopes, *scope)
		} else {
			glog.Errorf("Could not parse scope %s", v)
		}
	}
	return &scopes
}

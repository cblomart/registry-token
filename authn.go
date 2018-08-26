package main

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/golang/glog"
	"gopkg.in/ldap.v2"
)

func ldapDial(server string, usetls bool) (*ldap.Conn, error) {
	if usetls {
		return ldap.DialTLS("tcp", server, &tls.Config{InsecureSkipVerify: true})
	}
	return ldap.Dial("tcp", server)
}

func ldapSearch(limit uint, search string, attributes []string) *ldap.SearchRequest {
	return ldap.NewSearchRequest(
		AuthConfig.LDAPBase,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		int(limit),
		10,
		false,
		search,
		attributes,
		nil,
	)
}

//Authenticate authenticates the user and returns to groups he is member of and wether or not the user was found.
func Authenticate(user, password string) ([]string, bool) {
	binduser := user
	groups := []string{}
	// complete the login
	if !strings.Contains(user, "@") && !strings.Contains(user, "\\") {
		binduser = fmt.Sprintf("%s\\%s", AuthConfig.DefaultDomain, user)
	}
	// parse server
	port := "389"
	if AuthConfig.LDAPTls {
		port = "636"
	}
	server := AuthConfig.LDAPServer
	parts := strings.Split(AuthConfig.LDAPServer, ":")
	switch len(parts) {
	case 1:
		server = fmt.Sprintf("%s:%s", server, port)
	case 2:
		// do nothing
	default:
		glog.Errorf("Server provided is incorrect (%s)", server)
		return groups, false
	}
	glog.Infof("Authenticate to ldap server: %s", server)
	// Connect to LDAP
	l, err := ldapDial(server, AuthConfig.LDAPTls)
	if err != nil {
		glog.Errorf("Could not connect to LDAP %s: %s", server, err)
		return groups, false
	}
	defer l.Close()
	err = l.Bind(binduser, password)
	if err != nil {
		glog.Errorf("Could not bind to LDAP %s: %s", server, err)
		return groups, false
	}
	// search user dn
	dnsearch := ldapSearch(
		1,
		fmt.Sprintf("(&(%s=%s)(objectCategory=person))", AuthConfig.LDAPAttribute, user),
		[]string{"dn"},
	)
	dnresult, err := l.Search(dnsearch)
	if err != nil {
		glog.Errorf("Could not search for %s DN: %s", user, err)
		return groups, true
	}
	if len(dnresult.Entries) != 1 {
		glog.Errorf("Unexpected amount of DN returned for %s.", user)
		return groups, true
	}
	dn := dnresult.Entries[0].DN
	// search the groups
	groupsearch := ldapSearch(
		0,
		fmt.Sprintf("(&(member=%s)(objectCategory=group))", dn),
		[]string{AuthConfig.LDAPAttribute},
	)
	groupresult, err := l.Search(groupsearch)
	if err != nil {
		glog.Errorf("Could not search for %s groups: %s", user, err)
		return groups, true
	}
	for _, v := range groupresult.Entries {
		groups = append(groups, v.GetAttributeValue(AuthConfig.LDAPAttribute))
	}
	return groups, true
}

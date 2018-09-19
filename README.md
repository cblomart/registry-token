# registry-token-ldap

[![Go Report Card](https://goreportcard.com/badge/github.com/cblomart/registry-token-ldap)](https://goreportcard.com/report/github.com/cblomart/registry-token-ldap) [![Maintainability](https://api.codeclimate.com/v1/badges/1b846ff830e068ea7658/maintainability)](https://codeclimate.com/github/cblomart/registry-token-ldap/maintainability) [![codecov](https://codecov.io/gh/cblomart/registry-token-ldap/branch/master/graph/badge.svg)](https://codecov.io/gh/cblomart/registry-token-ldap) [![Drone Build Status](https://drone.blomart.net/api/badges/cblomart/registry-token-ldap/status.svg)](https://drone.blomart.net/cblomart/registry-token-ldap) [![](https://images.microbadger.com/badges/image/cblomart/registry-token-ldap.svg)](https://microbadger.com/images/cblomart/registry-token-ldap "Get your own image badge on microbadger.com") [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Authentication Server for Registry v2 with ldap

Will provide tokens on basis of LDAP authentification.
LDAP authentification will be done by binding to ldap with the username and password provided.
This plugin is oriented to AD so the username will be matched to SamAccountName and a default domain is required.

The set of rules will be evaluated and the resultant actions for the scope will be returned.
Rules can be set on users or on groups.

## configuration file

```yaml
# cert and key will be generated if file are not present
jwscert: /etc/registry-token-ldap/cert.crt
jwskey: /etc/registry-token-ldap/cert.key
# issuer must match registry config
issuer: "auth.registry.local"
# ldap server to use
ldapserver: ad.contoso.com
# base to search for users
ldapbase: "DC=contoso,DC=com"
# domain to automaticaly add to auth request
defaultdomain: CONTOSO
# rules to provide access (cumulative)
rules:
  # Admin can do all
  - group: "AdminGroup"
    match: ".+"
    actions: [ "push", "pull" ]
  # Users can do all on their repo
  - match: "${user}/.+"
    actions: [ "push", "pull" ]
  # Everybody can pull
  - match: ".+"
    actions: [ "pull" ]
```
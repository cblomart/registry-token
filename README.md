# registry-token-ldap

[![Go Report Card](https://goreportcard.com/badge/github.com/cblomart/registry-token-ldap)](https://goreportcard.com/report/github.com/cblomart/registry-token-ldap) [![Maintainability](https://api.codeclimate.com/v1/badges/1b846ff830e068ea7658/maintainability)](https://codeclimate.com/github/cblomart/registry-token-ldap/maintainability) [![Drone Build Status](https://drone.blomart.net/api/badges/cblomart/registry-token-ldap/status.svg)](https://drone.blomart.net/cblomart/registry-token-ldap) [![](https://images.microbadger.com/badges/image/cblomart/registry-token-ldap.svg)](https://microbadger.com/images/cblomart/registry-token-ldap "Get your own image badge on microbadger.com") [![Known Vulnerabilities](https://snyk.io/test/github/cblomart/registry-token-ldap/badge.svg)](https://snyk.io/test/github/cblomart/registry-token-ldap) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Authentication Server for Registry v2 with ldap

Will provide tokens on basis of LDAP authentification.
LDAP authentification will be done by binding to ldap with the username and password provided.
This plugin is oriented to AD so the username will be matched to SamAccountName and a default domain is required.

The set of rules will be evaluated and the resultant actions for the scope will be returned.
Rules can be set on users or on groups.
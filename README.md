# registry-token-ldap

Authentication Server for Registry v2 with ldap

Will provide tokens on basis of LDAP authentification.
LDAP authentification will be done by binding to ldap with the username and password provided.
This plugin is oriented to AD so the username will be matched to SamAccountName and a default domain is required.

The set of rules will be evaluated and the resultant actions for the scope will be returned.
Rules can be set on users or on groups.

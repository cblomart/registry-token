package main

// Config defines the basic config of the server
type Config struct {
	Port          uint   `default:"5001"`
	Path          string `default:"token"`
	JWSCert       string `required:"true"`
	JWSKey        string `required:"true"`
	Issuer        string `required:"true"`
	LDAPServer    string `required:"true"`
	LDAPTls       string `default:"insecure"`
	LDAPAttribute string `default:"sAMAccountName"`
	LDAPBase      string `required:"true"`
	LDAPCa        string
	DefaultDomain string `required:"true"`
	Rules         []Rule `required:"true"`
}

// Rule define a rule pattern
type Rule struct {
	User    string
	Group   string
	Match   string   `required:"true"`
	Actions []string `required:"true"`
}

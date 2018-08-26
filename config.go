package main

const (
	folder = "token"
)

// Config defines the basic config of the server
type Config struct {
	Port          uint    `default:"5001"`
	Path          string  `default:"token"`
	JWSCert       string  `required:"true"`
	JWSKey        string  `required:"true"`
	LDAPServer    string  `required:"true"`
	LDAPTls       bool    `default:"true"`
	LDAPAttribute string  `default:"SamAccountName"`
	LDAPBase      string  `required:"true"`
	DefaultDomain string  `required:"true"`
	Rules         []Rules `required:"true"`
}

// Rules define a rule pattern
type Rules struct {
	User    string
	Group   string
	Match   string   `required:"true"`
	Actions []string `required:"true"`
}

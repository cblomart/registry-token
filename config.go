package main

const (
	folder = "token"
)

// Config defines the basic config of the server
type Config struct {
	Port          uint    `default:"5001"`
	JWSCert       string  `required:"true"`
	LDAPServer    string  `required:"true"`
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

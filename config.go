package main

const (
	folder = "token"
)

// Config defines the basic config of the server
type Config struct {
	Port       int
	JWSCert    string
	LDAPServer string
	Rules      []Rules
}

// Rules define a rule pattern
type Rules struct {
	User    string
	Group   string
	Match   string
	Actions []string
}

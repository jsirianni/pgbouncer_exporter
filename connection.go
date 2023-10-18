package main

import (
	"fmt"
	"net/url"
)

// Connection represents a connection string for accessing pgBouncer.
// postgres://postgres:@localhost:6543/pgbouncer?sslmode=disable
type Connection struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	SSLMode  string `json:"sslmode"`
}

func (c *Connection) Build() (string, error) {
	u := url.URL{
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", c.Hostname, c.Port),
		Path:     c.Database,
		User:     url.UserPassword(c.Username, c.Password),
		RawQuery: fmt.Sprintf("sslmode=%s", c.SSLMode),
	}
	return u.String(), nil
}

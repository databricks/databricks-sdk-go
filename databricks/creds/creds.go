package creds

import "net/http"

type Provider interface {
	Refresh() error
	Credentials() Credentials
}

type Credentials interface {
	Apply(r *http.Request) error
	IsExpired() bool
}

package cache

import (
	"errors"

	"golang.org/x/oauth2"
)

type TokenCache interface {
	Store(key string, t *oauth2.Token) error
	Lookup(key string) (*oauth2.Token, error)
}

var ErrNotConfigured = errors.New("databricks OAuth is not configured for this host")

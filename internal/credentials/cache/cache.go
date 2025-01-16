package cache

import (
	"errors"

	"golang.org/x/oauth2"
)

// TokenCache is an interface for storing and looking up OAuth tokens.
type TokenCache interface {
	// Store stores the token with the given key, replacing any existing token.
	Store(key string, t *oauth2.Token) error

	// Lookup looks up the token with the given key. If the token is not found, it
	// returns ErrNotConfigured.
	Lookup(key string) (*oauth2.Token, error)
}

var ErrNotConfigured = errors.New("databricks OAuth is not configured for this host")

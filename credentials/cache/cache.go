package cache

import (
	"golang.org/x/oauth2"
)

type TokenCache interface {
	Store(key string, t *oauth2.Token) error
	Lookup(key string) (*oauth2.Token, error)
}

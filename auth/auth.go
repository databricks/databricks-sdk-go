// Package auth is an internal package that provides authentication utilities.
package auth

import (
	"context"

	"golang.org/x/oauth2"
)

// A TokenSource is anything that can return a token.
type TokenSource interface {
	// Token returns a token or an error. Token must be safe for concurrent use
	// by multiple goroutines. The returned Token must not be modified.
	Token(context.Context) (*oauth2.Token, error)
}

// TokenSourceFn is an adapter to allow the use of ordinary functions as
// TokenSource.
//
// Example:
//
//	   ts := TokenSourceFn(func(ctx context.Context) (*oauth2.Token, error) {
//			return &oauth2.Token{}, nil
//	   })
type TokenSourceFn func(context.Context) (*oauth2.Token, error)

func (fn TokenSourceFn) Token(ctx context.Context) (*oauth2.Token, error) {
	return fn(ctx)
}

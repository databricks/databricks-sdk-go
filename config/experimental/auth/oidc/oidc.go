// Package oidc provides utilities for working with OIDC ID tokens.
//
// This package is experimental and subject to change.
package oidc

import (
	"context"
	"fmt"
	"os"
)

// IDToken is a token that can be exchanged for a an access token.
// Value is the token string.
type IDToken struct {
	Value string
}

// IDTokenSource is anything that returns an IDToken given an audience.
type IDTokenSource interface {
	// Function to get the token
	IDToken(ctx context.Context, audience string) (*IDToken, error)
}

// IDTokenSourceFn is an adapter to allow the use of ordinary functions as
// IDTokenSource.
//
// Example:
//
//	   ts := IDTokenSourceFn(func(ctx context.Context) (*IDToken, error) {
//			return &IDToken{}, nil
//	   })
type IDTokenSourceFn func(ctx context.Context, audience string) (*IDToken, error)

func (fn IDTokenSourceFn) IDToken(ctx context.Context, audience string) (*IDToken, error) {
	return fn(ctx, audience)
}

// NewEnvIDTokenSource returns an IDTokenSource that reads the token from
// environment variable v.
//
// Note that the IDTokenSource does not cache the token and will read the token
// from environment variable v each time.
func NewEnvIDTokenSource(v string) IDTokenSource {
	return IDTokenSourceFn(func(ctx context.Context, _ string) (*IDToken, error) {
		t := os.Getenv(v)
		if t == "" {
			return nil, fmt.Errorf("missing env var %q", v)
		}
		return &IDToken{Value: t}, nil
	})
}

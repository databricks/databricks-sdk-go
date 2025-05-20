package config

import (
	"context"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
)

// Creates a CredentialsStrategy from a TokenSource.
func NewTokenSourceStrategy(name string, ts auth.TokenSource) CredentialsStrategy {
	return &tokenSourceStrategy{name: name, ts: ts}
}

// tokenSourceStrategy is wrapper on a auth.TokenSource which converts it into
// a CredentialsStrategy.
type tokenSourceStrategy struct {
	name string
	ts   auth.TokenSource
}

// Configure implements [CredentialsStrategy.Configure].
func (tss *tokenSourceStrategy) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	cp := credentials.NewOAuthCredentialsProviderFromTokenSource(auth.NewCachedTokenSource(tss.ts))

	// Sanity check that a token can be obtained.
	if _, err := cp.Token(); err != nil {
		return nil, err
	}

	return cp, nil
}

// Name implements [CredentialsStrategy.Name].
func (t *tokenSourceStrategy) Name() string {
	return t.name
}

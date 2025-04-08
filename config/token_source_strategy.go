package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/authconv"
	"github.com/databricks/databricks-sdk-go/logger"
)

type IDToken struct {
	Value string
}

// TokenProvider is an interface for a Token Provider with an audience.
type TokenProvider interface {
	// Function to get the token
	IDToken(ctx context.Context, audience string) (*IDToken, error)
}

// TokenSourceStrategy is wrapper on a auth.TokenSource which converts it into a CredentialsStrategy
type TokenSourceStrategy struct {
	tokenSource auth.TokenSource
	name        string
}

// Configure implements [CredentialsStrategy.Configure].
func (t *TokenSourceStrategy) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {

	// If we cannot get a token, skip this CredentialsStrategy.
	// We don't want to fail here because it's possible that the supplier is enabled
	// without the user action. For instance, jobs running in GitHub will have
	// OIDC environment variables added automatically
	cached := auth.NewCachedTokenSource(t.tokenSource)
	if _, err := cached.Token(ctx); err != nil {
		logger.Debugf(ctx, fmt.Sprintf("Skipping %s due to error: %v", t.name, err))
		return nil, nil
	}

	visitor := refreshableAuthVisitor(cached, ctx)
	return credentials.NewOAuthCredentialsProvider(visitor, authconv.OAuth2TokenSource(cached).Token), nil
}

// Name implements [CredentialsStrategy.Name].
func (t *TokenSourceStrategy) Name() string {
	return t.name
}

// Creates a CredentialsStrategy from a TokenSource.
func NewTokenSourceStrategy(
	name string,
	tokenSource auth.TokenSource,
) CredentialsStrategy {
	return &TokenSourceStrategy{
		name:        name,
		tokenSource: tokenSource,
	}
}

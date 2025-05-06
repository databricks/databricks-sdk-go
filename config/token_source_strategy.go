package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/authconv"
	"github.com/databricks/databricks-sdk-go/logger"
)

// Creates a CredentialsStrategy from a TokenSource.
func NewTokenSourceStrategy(
	name string,
	tokenSource auth.TokenSource,
) CredentialsStrategy {
	return &tokenSourceStrategy{
		name:        name,
		tokenSource: tokenSource,
	}
}

// tokenSourceStrategy is wrapper on a auth.TokenSource which converts it into a CredentialsStrategy
type tokenSourceStrategy struct {
	tokenSource auth.TokenSource
	name        string
}

// Configure implements [CredentialsStrategy.Configure].
func (t *tokenSourceStrategy) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {

	// If we cannot get a token, skip this CredentialsStrategy.
	// We don't want to fail here because it's possible that the supplier is enabled
	// without the user action. For instance, jobs running in GitHub will have
	// OIDC environment variables added automatically
	cached := auth.NewCachedTokenSource(t.tokenSource)
	if _, err := cached.Token(ctx); err != nil {
		logger.Debugf(ctx, fmt.Sprintf("Skipping %s due to error: %v", t.name, err))
		return nil, nil
	}

	visitor := refreshableAuthVisitor(cached)
	return credentials.NewOAuthCredentialsProvider(visitor, authconv.OAuth2TokenSource(cached).Token), nil
}

// Name implements [CredentialsStrategy.Name].
func (t *tokenSourceStrategy) Name() string {
	return t.name
}

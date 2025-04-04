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

type TokenProvider interface {
	// Function to get the token
	IDToken(ctx context.Context, audience string) (*IDToken, error)
}

type TokenSourceStrategy struct {
	tokenSource auth.TokenSource
	name        string
}

func (t *TokenSourceStrategy) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {

	// If we cannot get a token, skip this CredentialsStrategy.
	// We don't want to fail here because it's possible that the supplier is enabled
	// without the user action. For instance, jobs running in GitHub will have
	// OIDC environment variables added automatically
	if _, err := t.tokenSource.Token(ctx); err != nil {
		logger.Debugf(ctx, fmt.Sprintf("Skipping %s due to error: %v", t.name, err))
		return nil, nil
	}

	visitor := refreshableVisitor(authconv.OAuth2TokenSource(t.tokenSource))
	return credentials.NewOAuthCredentialsProvider(visitor, authconv.OAuth2TokenSource(t.tokenSource).Token), nil
}

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

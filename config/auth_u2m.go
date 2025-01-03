package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/credentials"
	"github.com/databricks/databricks-sdk-go/credentials/oauth"
)

type U2MCredentials struct {
	Auth oauth.PersistentAuth
}

// Name implements CredentialsStrategy.
func (u U2MCredentials) Name() string {
	return "oauth-u2m"
}

// Configure implements CredentialsStrategy.
func (u U2MCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	f := func(r *http.Request) error {
		arg := oauth.BasicOAuthArgument{
			Host:      cfg.Host,
			AccountID: cfg.AccountID,
		}
		token, err := u.Auth.Load(r.Context(), arg)
		if err != nil {
			return fmt.Errorf("oidc: %w", err)
		}
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
		return nil
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, "", nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	// Try to load the credential from the token cache. If absent, fall back
	// to the next credentials strategy.
	if err := f(r); err != nil {
		return nil, nil
	}

	return credentials.NewCredentialsProvider(f), nil
}

var _ CredentialsStrategy = U2MCredentials{}

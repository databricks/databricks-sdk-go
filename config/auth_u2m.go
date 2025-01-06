package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/credentials"
	"github.com/databricks/databricks-sdk-go/credentials/cache"
	"github.com/databricks/databricks-sdk-go/credentials/oauth"
	"github.com/databricks/databricks-sdk-go/logger"
)

type U2MCredentials struct {
	// Auth is the persistent auth object to use. If not specified, a new one will
	// be created, using the default cache and locker.
	Auth *oauth.PersistentAuth

	// ErrorHandler controls the behavior of Configure() when loading the OAuth
	// token fails. If not specified, any error will cause Configure() to return
	// said error.
	ErrorHandler func(context.Context, error) error

	name string
}

// Name implements CredentialsStrategy.
func (u U2MCredentials) Name() string {
	if u.name != "" {
		return "oauth-u2m"
	}
	return u.name
}

// Configure implements CredentialsStrategy.
func (u U2MCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	a := u.Auth
	if a == nil {
		var err error
		a, err = oauth.NewPersistentAuth(ctx)
		if err != nil {
			logger.Debugf(ctx, "failed to create persistent auth: %v, continuing", err)
			return nil, nil
		}
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, "", nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}

	f := func(r *http.Request) error {
		arg := oauth.BasicOAuthArgument{
			Host:      cfg.Host,
			AccountID: cfg.AccountID,
		}
		token, err := a.Load(r.Context(), arg)
		if err != nil {
			return fmt.Errorf("oidc: %w", err)
		}
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
		return nil
	}

	// Try to load the credential from the token cache. If absent, fall back to
	// the next credentials strategy. If a token is present but cannot be loaded
	// (e.g. expired), return an error. Otherwise, fall back to the next
	// credentials strategy.
	if err := f(r); err != nil {
		if u.ErrorHandler != nil {
			return nil, u.ErrorHandler(ctx, err)
		}
		return nil, err
	}

	return credentials.NewCredentialsProvider(f), nil
}

var _ CredentialsStrategy = U2MCredentials{}

var databricksCliCredentials = U2MCredentials{
	ErrorHandler: func(ctx context.Context, err error) error {
		if errors.Is(err, cache.ErrNotConfigured) {
			return nil
		}
		if _, ok := err.(*oauth.InvalidRefreshTokenError); ok {
			return err
		}
		logger.Debugf(ctx, "failed to load token: %v, continuing", err)
		return nil
	},
	name: "databricks-cli",
}

package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/credentials/u2m/cache"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

// u2mCredentials is a credentials strategy that uses the U2M OAuth flow to
// authenticate with Databricks. It loads a token from the token cache for the
// given workspace or account, refreshing it if needed. If the user has not
// authenticated with OAuth U2M, it falls back to the next credentials strategy.
// If they have but their access and refresh tokens are both invalid, it returns
// a special error message that instructs the user how to reauthenticate.
type u2mCredentials struct {
	// ts supplies the token source for the U2M OAuth flow.
	ts oauth2.TokenSource
}

// Name implements CredentialsStrategy.
func (u u2mCredentials) Name() string {
	// When we support allowing users to configure a custom U2M strategy, we
	// should use a different name here.
	return "databricks-cli"
}

// Configure implements CredentialsStrategy.
func (u u2mCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Host == "" {
		return nil, nil
	}

	arg, err := cfg.getOAuthArgument()
	if err != nil {
		return nil, fmt.Errorf("oidc: %w", err)
	}

	if u.ts == nil {
		auth, err := u2m.NewPersistentAuth(ctx, u2m.WithOAuthArgument(arg), u2m.WithPort(cfg.OAuthCallbackPort))
		if err != nil {
			logger.Debugf(ctx, "failed to create persistent auth: %v, continuing", err)
			return nil, nil
		}
		u.ts = auth
	}

	// Construct the visitor, and try to load the credential from the token
	// cache. If absent, fall back to the next credentials strategy. If a token
	// is present but cannot be loaded (e.g. expired), return an error.
	// Otherwise, fall back to the next credentials strategy.
	visitor := u.makeVisitor(u.ts)
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, "", nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	if err := visitor(r); err != nil {
		return nil, u.errorHandler(ctx, cfg, arg, err)
	}

	return credentials.NewOAuthCredentialsProvider(visitor, u.ts.Token), nil
}

func (u u2mCredentials) makeVisitor(auth oauth2.TokenSource) func(*http.Request) error {
	return func(r *http.Request) error {
		token, err := auth.Token()
		if err != nil {
			return fmt.Errorf("oidc: %w", err)
		}
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
		return nil
	}
}

func (u u2mCredentials) errorHandler(ctx context.Context, cfg *Config, arg u2m.OAuthArgument, err error) error {
	// If the current OAuth argument doesn't have a corresponding session
	// token, fall back to the next credentials strategy.
	if errors.Is(err, cache.ErrNotFound) {
		return nil
	}
	// If there is an existing token but the refresh token is invalid,
	// return a special error message for invalid refresh tokens. To help
	// users easily reauthenticate, include a command that the user can
	// run, prepopulating the profile, host and/or account ID.
	target := &u2m.InvalidRefreshTokenError{}
	if errors.As(err, &target) {
		return &CliInvalidRefreshTokenError{
			loginCommand: buildLoginCommand(cfg.Profile, arg),
			err:          err,
		}
	}
	// Otherwise, log the error and continue to the next credentials strategy.
	logger.Debugf(ctx, "failed to load token: %v, continuing", err)
	return nil
}

var _ CredentialsStrategy = u2mCredentials{}

// CliInvalidRefreshTokenError is a special error type that is returned when a
// new access token could not be retrieved because the refresh token is invalid.
// It is returned only by the `databricks-cli` credentials strategy.
type CliInvalidRefreshTokenError struct {
	loginCommand string
	err          error
}

func (e *CliInvalidRefreshTokenError) Error() string {
	return fmt.Sprintf(`a new access token could not be retrieved because the refresh token is invalid. If using the CLI, run the following command to reauthenticate:

  $ %s`, e.loginCommand)
}

func (e *CliInvalidRefreshTokenError) Unwrap() error {
	return e.err
}

// buildLoginCommand returns the `databricks auth login` command that the user
// can run to reauthenticate. The command is prepopulated with the profile, host
// and/or account ID.
func buildLoginCommand(profile string, arg u2m.OAuthArgument) string {
	cmd := []string{
		"databricks",
		"auth",
		"login",
	}
	if profile != "" {
		cmd = append(cmd, "--profile", profile)
	} else {
		switch arg := arg.(type) {
		case u2m.AccountOAuthArgument:
			cmd = append(cmd, "--host", arg.GetAccountHost(), "--account-id", arg.GetAccountId())
		case u2m.WorkspaceOAuthArgument:
			cmd = append(cmd, "--host", arg.GetWorkspaceHost())
		}
	}
	return strings.Join(cmd, " ")
}

var DatabricksCliCredentials = u2mCredentials{}

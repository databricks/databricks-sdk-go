package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/credentials/cache"
	"github.com/databricks/databricks-sdk-go/credentials/oauth"
	"github.com/databricks/databricks-sdk-go/logger"
)

// U2MCredentials is a credentials strategy that uses the U2M OAuth flow to
// authenticate with Databricks.
//
// To authenticate with U2M OAuth, the user must already have an existing OAuth
// session. The specific OAuth session is indicated by the OAuth argument
// provided by GetOAuthArg. By default, the OAuth argument is determined by the
// account host and account ID or workspace host in the Config.
//
// Error handling for this strategy is controlled by the ErrorHandler field. If
// ErrorHandler is not specified, any error will cause Configure() to return said
// error.
type U2MCredentials struct {
	// Auth is the persistent auth object to use. If not specified, a new one will
	// be created, using the default cache and locker.
	Auth *oauth.PersistentAuth

	// GetOAuthArg is a function that returns the OAuth argument to use for
	// loading the OAuth session token. If not specified, the OAuth argument is
	// determined by the account host and account ID or workspace host in the
	// Config.
	GetOAuthArg func(context.Context, *Config) (oauth.OAuthArgument, error)

	// ErrorHandler controls the behavior of Configure() when loading the OAuth
	// token fails. If not specified, any error will cause Configure() to return
	// said error.
	ErrorHandler func(context.Context, *Config, oauth.OAuthArgument, error) error

	name string
}

// Name implements CredentialsStrategy.
func (u U2MCredentials) Name() string {
	if u.name != "" {
		return u.name
	}
	return "oauth-u2m"
}

// Configure implements CredentialsStrategy.
func (u U2MCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Host == "" {
		return nil, nil
	}
	a := u.Auth
	if a == nil {
		var err error
		a, err = oauth.NewPersistentAuth(ctx)
		if err != nil {
			logger.Debugf(ctx, "failed to create persistent auth: %v, continuing", err)
			return nil, nil
		}
	}

	var arg oauth.OAuthArgument
	var err error
	if u.GetOAuthArg != nil {
		arg, err = u.GetOAuthArg(ctx, cfg)
	} else {
		arg, err = defaultGetOAuthArg(ctx, cfg)
	}
	if err != nil {
		return nil, fmt.Errorf("oidc: %w", err)
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, "", nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}

	f := func(r *http.Request) error {
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
			return nil, u.ErrorHandler(ctx, cfg, arg, err)
		}
		return nil, err
	}

	return credentials.NewCredentialsProvider(f), nil
}

func defaultGetOAuthArg(_ context.Context, cfg *Config) (oauth.OAuthArgument, error) {
	if cfg.IsAccountClient() {
		return oauth.NewBasicAccountOAuthArgument(cfg.Host, cfg.AccountID)
	}
	return oauth.NewBasicWorkspaceOAuthArgument(cfg.Host)
}

var _ CredentialsStrategy = U2MCredentials{}

// CliInvalidRefreshTokenError is a special error type that is returned when a
// new access token could not be retrieved because the refresh token is invalid.
// It is returned only by the `databricks-cli` credentials strategy.
type CliInvalidRefreshTokenError struct {
	loginCommand string
	err          error
}

func (e *CliInvalidRefreshTokenError) Error() string {
	return fmt.Sprintf("a new access token could not be retrieved because the refresh token is invalid. If using the CLI, run `%s` to reauthenticate", e.loginCommand)
}

func (e *CliInvalidRefreshTokenError) Unwrap() error {
	return e.err
}

// buildLoginCommand returns the `databricks auth login` command that the user
// can run to reauthenticate. The command is prepopulated with the profile, host
// and/or account ID.
func buildLoginCommand(ctx context.Context, profile string, arg oauth.OAuthArgument) string {
	cmd := []string{
		"databricks",
		"auth",
		"login",
	}
	if profile != "" {
		cmd = append(cmd, "--profile", profile)
	} else {
		switch arg := arg.(type) {
		case oauth.AccountOAuthArgument:
			cmd = append(cmd, "--host", arg.GetAccountHost(ctx), "--account-id", arg.GetAccountId(ctx))
		case oauth.WorkspaceOAuthArgument:
			cmd = append(cmd, "--host", arg.GetWorkspaceHost(ctx))
		}
	}
	return strings.Join(cmd, " ")
}

// databricksCliCredentials is a credentials strategy that emulates the behavior
// of the earlier `databricks-cli` credentials strategy which invoked the
// `databricks auth token` command.
var databricksCliCredentials = U2MCredentials{
	ErrorHandler: func(ctx context.Context, cfg *Config, arg oauth.OAuthArgument, err error) error {
		// If the current OAuth argument doesn't have a corresponding session
		// token, fall back to the next credentials strategy.
		if errors.Is(err, cache.ErrNotConfigured) {
			return nil
		}
		// If there is an existing token but the refresh token is invalid,
		// return a special error message for invalid refresh tokens. To help
		// users easily reauthenticate, include a command that the user can
		// run, prepopulating the profile, host and/or account ID.
		target := &oauth.InvalidRefreshTokenError{}
		if errors.As(err, &target) {
			return &CliInvalidRefreshTokenError{
				loginCommand: buildLoginCommand(ctx, cfg.Profile, arg),
				err:          err,
			}
		}
		// Otherwise, log the error and continue to the next credentials strategy.
		logger.Debugf(ctx, "failed to load token: %v, continuing", err)
		return nil
	},
	GetOAuthArg: defaultGetOAuthArg,
	name:        "databricks-cli",
}

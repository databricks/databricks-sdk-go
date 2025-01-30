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
type u2mCredentials struct {
	// auth is the persistent auth object to use. If not specified, a new one will
	// be created, using the default cache and locker.
	auth *u2m.PersistentAuth

	// getOAuthArg is a function that returns the OAuth argument to use for
	// loading the OAuth session token. If not specified, the OAuth argument is
	// determined by the account host and account ID or workspace host in the
	// Config.
	getOAuthArg func(context.Context, *Config) (u2m.OAuthArgument, error)

	// errorHandler controls the behavior of Configure() when loading the OAuth
	// token fails. If not specified, any error will cause Configure() to return
	// said error.
	errorHandler func(context.Context, *Config, u2m.OAuthArgument, error) error

	name string
}

// Name implements CredentialsStrategy.
func (u u2mCredentials) Name() string {
	if u.name != "" {
		return u.name
	}
	return "oauth-u2m"
}

// Configure implements CredentialsStrategy.
func (u u2mCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Host == "" {
		return nil, nil
	}

	var arg u2m.OAuthArgument
	var err error
	if u.getOAuthArg != nil {
		arg, err = u.getOAuthArg(ctx, cfg)
	} else {
		arg, err = defaultGetOAuthArg(ctx, cfg)
	}
	if err != nil {
		return nil, fmt.Errorf("oidc: %w", err)
	}

	if u.auth == nil {
		var err error
		u.auth, err = u2m.NewPersistentAuth(ctx, u2m.WithOAuthArgument(arg))
		if err != nil {
			logger.Debugf(ctx, "failed to create persistent auth: %v, continuing", err)
			return nil, nil
		}
	}

	// Construct the visitor, and try to load the credential from the token
	// cache. If absent, fall back to the next credentials strategy. If a token
	// is present but cannot be loaded (e.g. expired), return an error.
	// Otherwise, fall back to the next credentials strategy.
	visitor := u.makeVisitor()
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, "", nil)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	if err := visitor(r); err != nil {
		if u.errorHandler != nil {
			return nil, u.errorHandler(ctx, cfg, arg, err)
		}
		return nil, err
	}

	return credentials.NewOAuthCredentialsProvider(visitor, func() (*oauth2.Token, error) { return u.auth.Token() }), nil
}

func (u u2mCredentials) makeVisitor() func(*http.Request) error {
	return func(r *http.Request) error {
		token, err := u.auth.Token()
		if err != nil {
			return fmt.Errorf("oidc: %w", err)
		}
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
		return nil
	}
}

func defaultGetOAuthArg(_ context.Context, cfg *Config) (u2m.OAuthArgument, error) {
	if cfg.IsAccountClient() {
		return u2m.NewBasicAccountOAuthArgument(cfg.Host, cfg.AccountID)
	}
	return u2m.NewBasicWorkspaceOAuthArgument(cfg.Host)
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

// DatabricksCliCredentials is a credentials strategy that emulates the behavior
// of the earlier `databricks-cli` credentials strategy which invoked the
// `databricks auth token` command.
var DatabricksCliCredentials = u2mCredentials{
	errorHandler: func(ctx context.Context, cfg *Config, arg u2m.OAuthArgument, err error) error {
		// If the current OAuth argument doesn't have a corresponding session
		// token, fall back to the next credentials strategy.
		if errors.Is(err, cache.ErrNotConfigured) {
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
	},
	getOAuthArg: defaultGetOAuthArg,
	name:        "databricks-cli",
}

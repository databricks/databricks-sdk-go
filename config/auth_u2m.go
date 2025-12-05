package config

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"golang.org/x/oauth2"
)

// u2mCredentials is a credentials strategy that uses the U2M OAuth flow to
// authenticate with Databricks. It loads a token from the token cache for the
// given workspace or account, refreshing it using the associated refresh token
// if needed.
type u2mCredentials struct{}

// Name implements CredentialsStrategy.
func (u u2mCredentials) Name() string {
	// When we support allowing users to configure a custom U2M strategy, we
	// should use a different name here.
	return "databricks-cli"
}

// Configure implements CredentialsStrategy.
func (u u2mCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Host == "" {
		return nil, fmt.Errorf("host is required")
	}

	arg, err := cfg.getOAuthArgument()
	if err != nil {
		return nil, fmt.Errorf("oidc: %w", err)
	}

	ts, err := u2m.NewPersistentAuth(ctx, u2m.WithOAuthArgument(arg), u2m.WithPort(cfg.OAuthCallbackPort))
	if err != nil {
		return nil, err
	}

	// TODO: Having to handle the CLI error here is not ideal as it couples the
	// SDK logic with the CLI's. Remove this wrapping logic as soon as the CLI
	// is able to handle it on its own.
	wts := wrapWithCLIErrorHandling(cfg, arg, ts)

	cts := auth.NewCachedTokenSource(wts) // important
	return credentials.NewOAuthCredentialsProviderFromTokenSource(cts), nil
}

func wrapWithCLIErrorHandling(cfg *Config, arg u2m.OAuthArgument, ts oauth2.TokenSource) auth.TokenSource {
	cliCmd := buildLoginCommand(cfg.Profile, arg)
	return auth.TokenSourceFn(func(ctx context.Context) (*oauth2.Token, error) {
		t, err := ts.Token()
		if err != nil {
			// If there is an existing token but the refresh token is invalid,
			// return a special error message for invalid refresh tokens.
			// To help users easily reauthenticate, include a command that the
			// user can run, prepopulating the profile, host and/or account ID.
			target := &u2m.InvalidRefreshTokenError{}
			if !errors.As(err, &target) {
				return nil, err
			}
			return nil, &CliInvalidRefreshTokenError{
				loginCommand: cliCmd,
				err:          err,
			}
		}
		return t, nil
	})
}

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

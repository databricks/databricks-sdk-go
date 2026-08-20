package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
)

type u2mCredentials struct{}

func (u u2mCredentials) Name() string {
	return "databricks-cli"
}

func (u u2mCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Host == "" {
		return nil, fmt.Errorf("host is required")
	}
	if err := validateCliGroupID(cfg); err != nil {
		return nil, err
	}

	// We only partially support custom scopes with databricks-cli auth.
	// Users can specify their scopes when logging in with `databricks auth login`,
	// but not when using `databricks-cli` auth in the SDKs.
	// The token store is keyed by profile name (when an explicit profile is set)
	// or by host (legacy/implicit default). In either case, scopes are not part
	// of the cache key, so custom scopes would be silently ignored.
	// TODO: remove this validation once the token store can identify scopes based on their permissions. This will
	// allow users to specify scopes explicitly in the SDKs.
	if err := validateCliScopes(cfg); err != nil {
		return nil, err
	}

	ts, err := NewCliTokenSource(ctx, cfg)
	if err != nil {
		return nil, err
	}
	_, err = ts.Token(ctx)
	if err != nil {
		return nil, err
	}
	return credentials.NewOAuthCredentialsProviderFromTokenSource(auth.NewCachedTokenSource(ts, cacheOptions(cfg)...)), nil
}

// ErrCustomGroupIDNotSupported is returned when a group ID is specified in
// code or the environment with databricks-cli auth. Role selection belongs to
// the CLI profile because `databricks auth token` has no group argument.
const ErrCustomGroupIDNotSupported = "cannot set group_id when using the CLI authentication method. " +
	"Use the CLI to manage the group for the selected profile"

// validateCliGroupID ensures that databricks-cli role selection comes from the
// selected configuration profile rather than code or environment overrides.
func validateCliGroupID(cfg *Config) error {
	if cfg.GroupID == "" {
		return nil
	}
	for _, attr := range ConfigAttributes {
		if attr.Name != "group_id" {
			continue
		}
		if cfg.getSource(&attr).Type == SourceFile {
			return nil
		}
		return errors.New(ErrCustomGroupIDNotSupported)
	}
	return nil
}

// ErrCustomScopesNotSupported is returned when custom scopes are specified
// with databricks-cli auth, which is not supported because the CLI's token
// cache is keyed by host, not by scopes.
const ErrCustomScopesNotSupported = "custom scopes are not supported with databricks-cli auth; " +
	"scopes are determined by what was last used when logging in with `databricks auth login`"

// validateCliScopes returns an error if custom scopes are
// specified with databricks-cli auth. The CLI's token cache is keyed by host,
// not by scopes, so custom scopes would be silently ignored otherwise. Custom scopes
// from config files are allowed since `databricks auth login` writes them there.
func validateCliScopes(cfg *Config) error {
	if len(cfg.Scopes) == 0 {
		return nil
	}
	for _, attr := range ConfigAttributes {
		if attr.Name != "scopes" {
			continue
		}
		if cfg.getSource(&attr).Type == SourceFile {
			return nil
		}
		return errors.New(ErrCustomScopesNotSupported)
	}
	return nil
}

var DatabricksCliCredentials = u2mCredentials{}

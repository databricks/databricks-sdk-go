package config

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/logger"
)

type u2mCredentials struct{}

func (u u2mCredentials) Name() string {
	return "databricks-cli"
}

func (u u2mCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Host == "" {
		return nil, fmt.Errorf("host is required")
	}
	ts, err := NewCliTokenSource(cfg)
	if err != nil {
		// If the user explicitly requested this auth type, return the error.
		// Otherwise, skip to the next auth method in the chain.
		if cfg.AuthType == "databricks-cli" {
			return nil, err
		}
		if errors.Is(err, ErrCliNotFound) || errors.Is(err, ErrLegacyCliDetected) {
			logger.Debugf(ctx, "databricks-cli auth: %v", err)
			return nil, nil
		}
		return nil, err
	}
	_, err = ts.Token(ctx)
	if err != nil {
		// If user hasn't logged in for this host (no token in cache) and didn't
		// explicitly request this auth type, skip to the next auth method.
		if strings.Contains(err.Error(), "databricks OAuth is not") && cfg.AuthType != "databricks-cli" {
			logger.Debugf(ctx, "databricks-cli auth: %v", err)
			return nil, nil
		}
		// Pass the CLI error through unchanged. The CLI already provides
		// helpful error messages with login commands for common cases.
		return nil, err
	}
	logger.Infof(ctx, "Using Databricks CLI authentication")
	return credentials.NewOAuthCredentialsProviderFromTokenSource(
		auth.NewCachedTokenSource(ts, auth.WithAsyncRefresh(!cfg.DisableAsyncTokenRefresh))), nil
}

var DatabricksCliCredentials = u2mCredentials{}

package config

import (
	"context"
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
	ts, err := NewCliTokenSource(cfg)
	if err != nil {
		return nil, err
	}
	_, err = ts.Token(ctx)
	if err != nil {
		return nil, err
	}
	return credentials.NewOAuthCredentialsProviderFromTokenSource(auth.NewCachedTokenSource(ts)), nil
}

var DatabricksCliCredentials = u2mCredentials{}

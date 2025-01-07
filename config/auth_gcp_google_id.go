package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
)

type GoogleDefaultCredentials struct {
	// options used to enable unit testing mode for OIDC
	opts []option.ClientOption
}

func (c GoogleDefaultCredentials) Name() string {
	return "google-id"
}

func (c GoogleDefaultCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.GoogleServiceAccount == "" || !cfg.IsGcp() {
		return nil, nil
	}
	inner, err := c.idTokenSource(ctx, cfg.Host, cfg.GoogleServiceAccount, c.opts...)
	if err != nil {
		return nil, err
	}
	if !cfg.IsAccountClient() {
		logger.Infof(ctx, "Using Google Default Application Credentials for Workspace")
		visitor := refreshableVisitor(inner)
		return credentials.NewCredentialsProvider(visitor), nil
	}
	// source for generateAccessToken
	platform, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
		TargetPrincipal: cfg.GoogleServiceAccount,
		Scopes: []string{
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/compute",
		},
	}, c.opts...)
	if err != nil {
		return nil, err
	}
	logger.Infof(ctx, "Using Google Default Application Credentials for Accounts API")
	visitor := serviceToServiceVisitor(inner, platform, "X-Databricks-GCP-SA-Access-Token")
	return credentials.NewOAuthCredentialsProvider(visitor, inner.Token), nil
}

func (c GoogleDefaultCredentials) idTokenSource(ctx context.Context, host, serviceAccount string,
	opts ...option.ClientOption) (oauth2.TokenSource, error) {
	ts, err := impersonate.IDTokenSource(ctx, impersonate.IDTokenConfig{
		Audience:        host,
		TargetPrincipal: serviceAccount,
		IncludeEmail:    true,
	}, opts...)
	if err != nil {
		err = fmt.Errorf("could not obtain OIDC token. %w Running 'gcloud auth application-default login' may help", err)
		return nil, err
	}
	return ts, err
}

package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/oidc"
	"github.com/databricks/databricks-sdk-go/logger"
)

const authDocURL = "https://docs.databricks.com/en/dev-tools/auth.html#databricks-client-unified-authentication"

type DefaultCredentials struct {
	name string
}

func (c *DefaultCredentials) Name() string {
	if c.name == "" {
		return "default"
	}
	return c.name
}

func (c *DefaultCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	// Order in which strategies are tested. Iteration proceeds from most
	// specific to most generic, and the first strategy to return a non-nil
	// credentials provider is selected.
	//
	// Modifying this order could break authentication for users whose
	// environments are compatible with multiple strategies and who rely on the
	// current priority for tie-breaking. While arguably an anti-pattern, this
	// order is maintained for backward compatibility.
	strategies := []CredentialsStrategy{
		PatCredentials{},
		BasicCredentials{},
		M2mCredentials{},
		u2mCredentials{},
		MetadataServiceCredentials{},
		// OIDC Strategies from most specific to most generic.
		oidcStrategy(cfg, "github-oidc", githubOIDCTokenSource(cfg)),
		oidcStrategy(cfg, "env-oidc", envOIDCTokenSource(cfg)),
		oidcStrategy(cfg, "file-oidc", fileOIDCTokenSource(cfg)),
		// Azure strategies from most specific to most generic.
		AzureGithubOIDCCredentials{},
		AzureMsiCredentials{},
		AzureClientSecretCredentials{},
		AzureCliCredentials{},
		// Google strategies from most specific to most generic.
		GoogleCredentials{},
		GoogleDefaultCredentials{},
	}

	// If an auth type is specified, try to configure the credentials for that
	// specific auth type. If an error is encountered, return it.
	if cfg.AuthType != "" {
		for _, s := range strategies {
			if s.Name() == cfg.AuthType {
				logger.Tracef(ctx, "Attempting to configure auth: %q", s.Name())
				c.name = s.Name()
				return s.Configure(ctx, cfg)
			}
		}
		return nil, fmt.Errorf("auth type %q not found, please check %s for a list of supported auth types", cfg.AuthType, authDocURL)
	}

	// If no auth type is specified, try the strategies in order. If a strategy
	// succeeds, returns the credentials provider. If a strategy fails, swallow
	// the error and try the next strategy.
	for _, s := range strategies {
		logger.Tracef(ctx, "Attempting to configure auth: %q", s.Name())
		cp, err := s.Configure(ctx, cfg)
		if err != nil || cp == nil {
			logger.Tracef(ctx, "Failed to configure auth: %q", s.Name())
			continue
		}
		c.name = s.Name()
		return cp, nil
	}

	return nil, fmt.Errorf("cannot configure default credentials, please check %s to configure credentials for your preferred authentication method", authDocURL)
}

// oidcStrategy returns a new CredentialsStrategy to authenticate with
// Databricks using the given OIDC IDTokenSource.
func oidcStrategy(cfg *Config, name string, ts oidc.IDTokenSource) CredentialsStrategy {
	oidcConfig := oidc.DatabricksOIDCTokenSourceConfig{
		ClientID:              cfg.ClientID,
		Host:                  cfg.CanonicalHostName(),
		TokenEndpointProvider: cfg.getOidcEndpoints,
		Audience:              cfg.TokenAudience,
		IDTokenSource:         ts,
	}
	if cfg.IsAccountClient() {
		oidcConfig.AccountID = cfg.AccountID
	}
	tokenSource := oidc.NewDatabricksOIDCTokenSource(oidcConfig)
	return NewTokenSourceStrategy(name, tokenSource)
}

func githubOIDCTokenSource(cfg *Config) oidc.IDTokenSource {
	return oidc.NewGithubIDTokenSource(
		cfg.refreshClient,
		cfg.ActionsIDTokenRequestURL,
		cfg.ActionsIDTokenRequestToken,
	)
}

func envOIDCTokenSource(cfg *Config) oidc.IDTokenSource {
	v := cfg.OIDCTokenEnv
	if v == "" {
		v = "DATABRICKS_OIDC_TOKEN"
	}
	return oidc.NewEnvIDTokenSource(v)
}

func fileOIDCTokenSource(cfg *Config) oidc.IDTokenSource {
	return oidc.NewFileTokenSource(cfg.OIDCTokenFilepath)
}

package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	authcred "github.com/databricks/databricks-sdk-go/config/experimental/auth/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/oidc"
	"github.com/databricks/databricks-sdk-go/logger"
)

const authDocURL = "https://docs.databricks.com/en/dev-tools/auth.html#databricks-client-unified-authentication"

// ErrCannotConfigureDefault indicates that the DefaultCredentials strategy was
// unable to configure the default credentials.
var ErrCannotConfigureDefault = fmt.Errorf("cannot configure default credentials, please check %s to configure credentials for your preferred authentication method", authDocURL)

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
		pat(),
		BasicCredentials{},
		M2mCredentials{},
		u2mCredentials{},
		MetadataServiceCredentials{},
		// OIDC Strategies.
		githubOIDC(cfg),
		envOIDC(cfg),
		fileOIDC(cfg),
		// Azure strategies.
		AzureGithubOIDCCredentials{},
		AzureMsiCredentials{},
		AzureClientSecretCredentials{},
		AzureCliCredentials{},
		// Google strategies.
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
			logger.Debugf(ctx, "Failed to configure auth: %q", s.Name())
			continue
		}
		c.name = s.Name()
		return cp, nil
	}

	return nil, ErrCannotConfigureDefault
}

type credStrategy struct {
	name      string
	configure func(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error)
}

func (c *credStrategy) Name() string {
	return c.name
}

func (c *credStrategy) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	return c.configure(ctx, cfg)
}

func fromCredentials(c auth.Credentials, err error) (credentials.CredentialsProvider, error) {
	if err != nil {
		return nil, err
	}
	return credentials.CredentialsProviderFn(func(r *http.Request) error {
		headers, err := c.AuthHeaders(context.Background())
		if err != nil {
			return err
		}
		for k, v := range headers {
			r.Header.Set(k, v)
		}
		return nil
	}), nil
}

func pat() CredentialsStrategy {
	return &credStrategy{
		name: "pat",
		configure: func(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
			return fromCredentials(authcred.NewPATCredentials(cfg.Token))
		},
	}
}

func githubOIDC(cfg *Config) CredentialsStrategy {
	return oidcStrategy(cfg, "github-oidc", oidc.NewGithubIDTokenSource(
		cfg.refreshClient,
		cfg.ActionsIDTokenRequestURL,
		cfg.ActionsIDTokenRequestToken,
	))
}

func envOIDC(cfg *Config) CredentialsStrategy {
	v := cfg.OIDCTokenEnv
	if v == "" {
		v = "DATABRICKS_OIDC_TOKEN"
	}
	return oidcStrategy(cfg, "env-oidc", oidc.NewEnvIDTokenSource(v))
}

func fileOIDC(cfg *Config) CredentialsStrategy {
	return oidcStrategy(cfg, "file-oidc", oidc.NewFileTokenSource(cfg.OIDCTokenFilepath))
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

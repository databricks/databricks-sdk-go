package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/logger"
)

var authProviders = []CredentialsStrategy{
	PatCredentials{},
	BasicCredentials{},
	M2mCredentials{},
	DatabricksCliCredentials{},
	MetadataServiceCredentials{},

	// Attempt to configure auth from most specific to most generic (the Azure CLI).
	AzureGithubOIDCCredentials{},
	AzureMsiCredentials{},
	AzureClientSecretCredentials{},
	AzureCliCredentials{},

	// Attempt to configure auth from most specific to most generic (Google Application Default Credentials).
	GoogleCredentials{},
	GoogleDefaultCredentials{},
}

type DefaultCredentials struct {
	name string
}

func (c *DefaultCredentials) Name() string {
	if c.name == "" {
		return "default"
	}
	return c.name
}

var authFlowUrl = "https://docs.databricks.com/en/dev-tools/auth.html#databricks-client-unified-authentication"
var errorMessage = fmt.Sprintf("cannot configure default credentials, please check %s to configure credentials for your preferred authentication method", authFlowUrl)

// ErrCannotConfigureAuth (experimental) is returned when no auth is configured
var ErrCannotConfigureAuth = errors.New(errorMessage)

func (c *DefaultCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	for _, p := range authProviders {
		if cfg.AuthType != "" && p.Name() != cfg.AuthType {
			// ignore other auth types if one is explicitly enforced
			logger.Infof(ctx, "Ignoring %s auth, because %s is preferred", p.Name(), cfg.AuthType)
			continue
		}
		logger.Tracef(ctx, "Attempting to configure auth: %s", p.Name())
		credentialsProvider, err := p.Configure(ctx, cfg)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", p.Name(), err)
		}
		if credentialsProvider == nil {
			continue
		}
		c.name = p.Name()
		return credentialsProvider, nil
	}
	return nil, ErrCannotConfigureAuth
}

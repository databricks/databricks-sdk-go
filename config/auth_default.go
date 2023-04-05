package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/logger"
)

var (
	authProviders = []CredentialsProvider{
		PatCredentials{},
		BasicCredentials{},
		M2mCredentials{},
		BricksCliCredentials{},
		MetadataServiceCredentials{},

		// Attempt to configure auth from most specific to most generic (the Azure CLI).
		AzureMsiCredentials{},
		AzureClientSecretCredentials{},
		AzureCliCredentials{},

		// Attempt to configure auth from most specific to most generic (Google Application Default Credentials).
		GoogleCredentials{},
		GoogleDefaultCredentials{},
	}
)

type DefaultCredentials struct {
	name string
}

func (c *DefaultCredentials) Name() string {
	if c.name == "" {
		return "default"
	}
	return c.name
}

// ErrCannotConfigureAuth (experimental) is returned when no auth is configured
var ErrCannotConfigureAuth = errors.New("cannot configure default credentials")

func (c *DefaultCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	for _, p := range authProviders {
		if cfg.AuthType != "" && p.Name() != cfg.AuthType {
			// ignore other auth types if one is explicitly enforced
			logger.Infof(ctx, "Ignoring %s auth, because %s is preferred", p.Name(), cfg.AuthType)
			continue
		}
		logger.Tracef(ctx, "Attempting to configure auth: %s", p.Name())
		visitor, err := p.Configure(ctx, cfg)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", p.Name(), err)
		}
		if visitor == nil {
			continue
		}
		c.name = p.Name()
		return visitor, nil
	}
	return nil, ErrCannotConfigureAuth
}

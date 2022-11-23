package databricks

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/logger"
)

var (
	authProviders = []CredentialsProvider{
		PatCredentials{},
		BasicCredentials{},
		AzureClientSecretCredentials{},
		AzureCliCredentials{},
		GoogleDefaultCredentials{},
		GoogleCredentials{},
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

func (c *DefaultCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	for _, p := range authProviders {
		if cfg.AuthType != "" && p.Name() != cfg.AuthType {
			// ignore other auth types if one is explicitly enforced
			logger.Infof("Ignoring %s auth, because %s is preferred", p.Name(), cfg.AuthType)
			continue
		}
		logger.Tracef("Attempting to configure auth: %s", p.Name())
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
	return nil, fmt.Errorf("cannot configure default credentials")
}

type PatCredentials struct {
}

func (c PatCredentials) Name() string {
	return "pat"
}

func (c PatCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if cfg.Token == "" || cfg.Host == "" {
		return nil, nil
	}
	return func(r *http.Request) error {
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.Token))
		return nil
	}, nil
}

type BasicCredentials struct {
}

func (c BasicCredentials) Name() string {
	return "basic"
}

func (c BasicCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if cfg.Username == "" || cfg.Password == "" || cfg.Host == "" {
		return nil, nil
	}
	tokenUnB64 := fmt.Sprintf("%s:%s", cfg.Username, cfg.Password)
	b64 := base64.StdEncoding.EncodeToString([]byte(tokenUnB64))
	return func(r *http.Request) error {
		r.Header.Set("Authorization", fmt.Sprintf("Basic %s", b64))
		return nil
	}, nil
}

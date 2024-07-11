package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
)

type oidcTokenSupplier interface {
	Name() string

	// GetOIDCToken returns an OIDC token for the given audience.
	GetOIDCToken(ctx context.Context, audience string) (string, error)
}

type githubOIDCTokenSupplier struct {
	idTokenRequestURL   string
	idTokenRequestToken string
	client              *httpclient.ApiClient
}

func githubOIDCTokenSupplierFromConfig(cfg *Config) githubOIDCTokenSupplier {
	return githubOIDCTokenSupplier{
		idTokenRequestURL:   cfg.ActionsIDTokenRequestURL,
		idTokenRequestToken: cfg.ActionsIDTokenRequestToken,
		client:              cfg.refreshClient,
	}
}

func (g githubOIDCTokenSupplier) Name() string {
	return "github"
}

// requestIDToken requests an ID token from the Github Action.
func (g githubOIDCTokenSupplier) GetOIDCToken(ctx context.Context, audience string) (string, error) {
	if g.idTokenRequestURL == "" {
		logger.Debugf(ctx, "Missing cfg.ActionsIDTokenRequestURL, likely not calling from a Github action")
		return "", nil
	}
	if g.idTokenRequestToken == "" {
		logger.Debugf(ctx, "Missing cfg.ActionsIDTokenRequestToken, likely not calling from a Github action")
		return "", nil
	}
	url := g.idTokenRequestURL
	if audience != "" {
		url = fmt.Sprintf("%s&audience=%s", url, audience)
	}
	resp := struct { // anonymous struct to parse the response
		Value string `json:"value"`
	}{}
	err := g.client.Do(ctx, "GET", url,
		httpclient.WithRequestHeader("Authorization", fmt.Sprintf("Bearer %s", g.idTokenRequestToken)),
		httpclient.WithResponseUnmarshal(&resp),
	)
	if err != nil {
		return "", fmt.Errorf("failed to request ID token from %s: %w", g.idTokenRequestURL, err)
	}

	return resp.Value, nil
}

var _ oidcTokenSupplier = githubOIDCTokenSupplier{}

type oidcTokenSuppliers []oidcTokenSupplier

func (c *Config) getAllOIDCSuppliers() oidcTokenSuppliers {
	return []oidcTokenSupplier{
		githubOIDCTokenSupplierFromConfig(c),
	}
}

func (o oidcTokenSuppliers) GetOIDCToken(ctx context.Context, audience string) (string, error) {
	for _, s := range o {
		token, err := s.GetOIDCToken(ctx, audience)
		if err != nil {
			return "", err
		}
		if token != "" {
			return token, nil
		}
		logger.Debugf(ctx, "No OIDC token found from %s", s.Name())
	}
	return "", nil
}

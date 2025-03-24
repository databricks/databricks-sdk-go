package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
)

type GithubOIDCTokenSupplier struct {
	cfg *Config
}

// requestIDToken requests an ID token from the Github Action.
func (g *GithubOIDCTokenSupplier) GetOIDCToken(ctx context.Context, audience string) (string, error) {
	if g.cfg.ActionsIDTokenRequestURL == "" {
		logger.Debugf(ctx, "Missing cfg.ActionsIDTokenRequestURL, likely not calling from a Github action")
		return "", nil
	}
	if g.cfg.ActionsIDTokenRequestToken == "" {
		logger.Debugf(ctx, "Missing cfg.ActionsIDTokenRequestToken, likely not calling from a Github action")
		return "", nil
	}

	resp := struct { // anonymous struct to parse the response
		Value string `json:"value"`
	}{}
	requestUrl := g.cfg.ActionsIDTokenRequestURL
	if audience != "" {
		requestUrl = fmt.Sprintf("%s&audience=%s", requestUrl, audience)
	}
	err := g.cfg.refreshClient.Do(ctx, "GET", requestUrl,
		httpclient.WithRequestHeader("Authorization", fmt.Sprintf("Bearer %s", g.cfg.ActionsIDTokenRequestToken)),
		httpclient.WithResponseUnmarshal(&resp),
	)
	if err != nil {
		return "", fmt.Errorf("failed to request ID token from %s: %w", g.cfg.ActionsIDTokenRequestURL, err)
	}

	return resp.Value, nil
}

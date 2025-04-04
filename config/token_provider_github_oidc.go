package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
)

type GithubProvider struct {
	cfg *Config
}

func (g *GithubProvider) IDToken(ctx context.Context, audience string) (*IDToken, error) {
	if g.cfg.ActionsIDTokenRequestURL == "" {
		logger.Debugf(ctx, "Missing cfg.ActionsIDTokenRequestURL, likely not calling from a Github action")
		return nil, errors.New("missing cfg.ActionsIDTokenRequestURL")
	}
	if g.cfg.ActionsIDTokenRequestToken == "" {
		logger.Debugf(ctx, "Missing cfg.ActionsIDTokenRequestToken, likely not calling from a Github action")
		return nil, errors.New("missing cfg.ActionsIDTokenRequestToken")
	}

	resp := &IDToken{}
	requestUrl := g.cfg.ActionsIDTokenRequestURL
	if audience != "" {
		requestUrl = fmt.Sprintf("%s&audience=%s", requestUrl, audience)
	}
	err := g.cfg.refreshClient.Do(ctx, "GET", requestUrl,
		httpclient.WithRequestHeader("Authorization", fmt.Sprintf("Bearer %s", g.cfg.ActionsIDTokenRequestToken)),
		httpclient.WithResponseUnmarshal(resp),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to request ID token from %s: %w", g.cfg.ActionsIDTokenRequestURL, err)
	}

	return resp, nil
}

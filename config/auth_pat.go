package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/config/credentials"
)

type PatCredentials struct {
}

func (c PatCredentials) Name() string {
	return "pat"
}

func (c PatCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Token == "" || cfg.Host == "" {
		return nil, nil
	}
	visitor := func(r *http.Request) error {
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.Token))
		return nil
	}
	return credentials.NewCredentialsProvider(visitor), nil
}

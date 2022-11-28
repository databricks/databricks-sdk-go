package config

import (
	"context"
	"fmt"
	"net/http"
)

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

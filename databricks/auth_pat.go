package databricks

import (
	"context"
	"fmt"
	"net/http"
)

type PatCredentials struct {
	Token string `name:"token" env:"DATABRICKS_TOKEN" auth:"token,sensitive"`
}

func (c PatCredentials) Name() string {
	return "pat"
}

func (c PatCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if c.Token == "" || cfg.Host == "" {
		return nil, nil
	}
	return func(r *http.Request) error {
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
		return nil
	}, nil
}

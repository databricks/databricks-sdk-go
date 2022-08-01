package databricks

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
)

type BasicCredentials struct {
	Username string `name:"username" env:"DATABRICKS_USERNAME" auth:"password"`
	Password string `name:"password" env:"DATABRICKS_PASSWORD" auth:"password,sensitive"`
}

func (c BasicCredentials) Name() string {
	return "basic"
}

func (c BasicCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if c.Username == "" || c.Password == "" || cfg.Host == "" {
		return nil, nil
	}
	tokenUnB64 := fmt.Sprintf("%s:%s", c.Username, c.Password)
	b64 := base64.StdEncoding.EncodeToString([]byte(tokenUnB64))
	return func(r *http.Request) error {
		r.Header.Set("Authorization", fmt.Sprintf("Basic %s", b64))
		return nil
	}, nil
}

package config

import "net/http"

func NewMockConfig(auth func(r *http.Request) error) *Config {
	if auth == nil {
		auth = func(r *http.Request) error {
			return nil
		}
	}
	return &Config{
		AuthType: "mock",
		auth:     auth,
		resolved: true,
	}
}

package databricks

import "net/http"

func NewMockConfig(auth func(r *http.Request) error) *Config {
	return &Config{
		AuthType: "mock",
		auth:     auth,
		resolved: true,
	}
}

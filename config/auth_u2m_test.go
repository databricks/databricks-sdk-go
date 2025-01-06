package config

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/credentials/oauth"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestU2MCredentials(t *testing.T) {
	tests := []struct {
		name       string
		cfg        *Config
		auth       func() (*oauth.PersistentAuth, error)
		expectErr  bool
		expectAuth string
	}{
		{
			name: "happy path",
			cfg: &Config{
				Host: "https://myworkspace.cloud.databricks.com",
			},
			auth: func() (*oauth.PersistentAuth, error) {
				return oauth.NewPersistentAuth(
					context.Background(),
					oauth.WithTokenCache(&InMemoryTokenCache{
						Tokens: map[string]*oauth2.Token{
							"https://myworkspace.cloud.databricks.com": {
								AccessToken: "dummy_access_token",
								Expiry:      time.Now().Add(1 * time.Hour),
							},
						},
					}))
			},
			expectErr:  false,
			expectAuth: "Bearer dummy_access_token",
		},
		{
			name: "expired token with invalid refresh token",
			cfg: &Config{
				Host: "https://databricks.com",
			},
			auth: func() (*oauth.PersistentAuth, error) {
				return oauth.NewPersistentAuth(
					context.Background(),
					oauth.WithTokenCache(&InMemoryTokenCache{
						Tokens: map[string]*oauth2.Token{
							"https://myworkspace.cloud.databricks.com": {
								AccessToken:  "dummy_access_token",
								RefreshToken: "dummy_refresh_token",
								Expiry:       time.Now().Add(-1 * time.Hour),
							},
						},
					}),
					oauth.WithApiClient(httpclient.NewApiClient(httpclient.ClientConfig{
						Transport: fixtures.SliceTransport{
							{
								MatchAny: true,
								Status:   401,
								Response: `{"error":"invalid_refresh_token","error_description":"Refresh token is invalid"}`,
							},
						},
					})),
				)
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			auth, err := tt.auth()
			require.NoError(t, err)
			strat := U2MCredentials{
				Auth: auth,
			}
			provider, err := strat.Configure(ctx, tt.cfg)
			if tt.expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://databricks.com", nil)
			require.NoError(t, err)

			err = provider.SetHeaders(req)
			require.NoError(t, err)
			require.Equal(t, tt.expectAuth, req.Header.Get("Authorization"))
		})
	}
}

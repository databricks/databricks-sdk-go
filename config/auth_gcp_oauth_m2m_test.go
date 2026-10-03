package config

import (
	"context"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func staticAuthTokenSource(token string) auth.TokenSource {
	return auth.TokenSourceFn(func(ctx context.Context) (*oauth2.Token, error) {
		return &oauth2.Token{TokenType: "Bearer", AccessToken: token}, nil
	})
}

// The core behavior: identity token in Authorization, Google access token in
// the X-Databricks-GCP-SA-Access-Token passthrough header. The Google source is
// injected so the test is hermetic; the Databricks token is minted against a
// mocked account token endpoint.
func TestGcpM2m_SetsBothHeaders(t *testing.T) {
	cfg := &Config{
		Host:                 "https://accounts.gcp.databricks.com",
		AccountID:            "abc",
		Cloud:                environment.CloudGCP,
		ClientID:             "b",
		ClientSecret:         "c",
		GoogleServiceAccount: "sa@proj.iam.gserviceaccount.com",
		AuthType:             "oauth-m2m-gcp",
		ConfigFile:           "/dev/null",
		HTTPTransport: fixtures.MappingTransport{
			"POST /oidc/accounts/abc/v1/token": {
				Response: oauth2.Token{
					TokenType:   "Bearer",
					AccessToken: "db-oauth",
				},
			},
		},
	}
	require.NoError(t, cfg.EnsureResolved())

	strat := GcpM2mCredentials{googleTokenSource: staticAuthTokenSource("gcp-access")}
	provider, err := strat.Configure(context.Background(), cfg)
	require.NoError(t, err)
	require.NotNil(t, provider)

	req, err := http.NewRequest("GET", "http://localhost", nil)
	require.NoError(t, err)
	require.NoError(t, provider.SetHeaders(req))

	require.Equal(t, "Bearer db-oauth", req.Header.Get("Authorization"))
	require.Equal(t, "gcp-access", req.Header.Get("X-Databricks-GCP-SA-Access-Token"))
}

// Combining oauth (client_id/client_secret) and google
// (google_service_account) config groups is rejected by the single-auth-method
// conflict check unless an explicit auth_type is set.
func TestGcpM2m_RequiresExplicitAuthType(t *testing.T) {
	newCfg := func(authType string) *Config {
		return &Config{
			Host:                 "https://accounts.gcp.databricks.com",
			AccountID:            "abc",
			Cloud:                environment.CloudGCP,
			ClientID:             "b",
			ClientSecret:         "c",
			GoogleServiceAccount: "sa@proj.iam.gserviceaccount.com",
			AuthType:             authType,
			ConfigFile:           "/dev/null",
		}
	}

	err := newCfg("").EnsureResolved()
	require.Error(t, err)
	require.Contains(t, err.Error(), "more than one authorization method")

	require.NoError(t, newCfg("oauth-m2m-gcp").EnsureResolved())
}

// The strategy is inert (returns a nil provider so the chain moves on) unless it
// is on GCP with Databricks OAuth SP credentials and a Google source.
func TestGcpM2m_ConfigureSkips(t *testing.T) {
	cases := []struct {
		name string
		cfg  *Config
	}{
		{
			name: "not gcp",
			cfg: &Config{
				Host: "https://foo.cloud.databricks.com", ClientID: "b", ClientSecret: "c",
				GoogleServiceAccount: "sa@proj.iam.gserviceaccount.com", ConfigFile: "/dev/null",
			},
		},
		{
			name: "missing client secret",
			cfg: &Config{
				Cloud: environment.CloudGCP, Host: "https://accounts.gcp.databricks.com",
				ClientID: "b", GoogleServiceAccount: "sa@proj.iam.gserviceaccount.com", ConfigFile: "/dev/null",
			},
		},
		{
			name: "no google source",
			cfg: &Config{
				Cloud: environment.CloudGCP, Host: "https://accounts.gcp.databricks.com",
				ClientID: "b", ClientSecret: "c", ConfigFile: "/dev/null",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			provider, err := GcpM2mCredentials{}.Configure(context.Background(), tc.cfg)
			require.NoError(t, err)
			require.Nil(t, provider)
		})
	}
}

func TestGoogleAccessTokenSource_RequiresASource(t *testing.T) {
	_, err := googleAccessTokenSource(context.Background(), &Config{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "requires google_credentials or google_service_account")
}

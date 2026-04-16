package config

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/stretchr/testify/require"
)

func TestHTTPClientConfigSetsOrgIdHeader(t *testing.T) {
	cfg := &Config{
		Host:        "https://myhost.databricks.com",
		Token:       "dapi1234",
		WorkspaceID: "7474644166319138",
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/api/2.0/test",
				ExpectedHeaders: map[string]string{
					"X-Databricks-Org-Id": "7474644166319138",
					"Authorization":       "Bearer dapi1234",
				},
				Response: `{"ok": true}`,
			},
		},
	}

	clientCfg, err := HTTPClientConfigFromConfig(cfg)
	require.NoError(t, err)

	apiClient := httpclient.NewApiClient(clientCfg)
	err = apiClient.Do(t.Context(), "GET", "/api/2.0/test")
	require.NoError(t, err)
}

func TestHTTPClientConfigOmitsOrgIdHeaderWhenEmpty(t *testing.T) {
	cfg := &Config{
		Host:  "https://myhost.databricks.com",
		Token: "dapi1234",
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:   "GET",
				Resource: "/api/2.0/test",
				ExpectedHeaders: map[string]string{
					"Authorization": "Bearer dapi1234",
				},
				Response: `{"ok": true}`,
			},
		},
	}

	clientCfg, err := HTTPClientConfigFromConfig(cfg)
	require.NoError(t, err)

	apiClient := httpclient.NewApiClient(clientCfg)
	err = apiClient.Do(t.Context(), "GET", "/api/2.0/test")
	require.NoError(t, err)
}

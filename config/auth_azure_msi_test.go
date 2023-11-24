package config

import (
	"net/http"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/stretchr/testify/require"
)

func init() {
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelDebug,
	}
}

func TestMsiHappyFlow(t *testing.T) {
	cfg := Config{
		AzureUseMSI:     true,
		AzureResourceID: "/a/b/c",
		HTTPTransport: fixtures.HTTPFixtures{
			{
				Method:   "GET",
				Resource: "/metadata/identity/oauth2/token?api-version=2018-02-01&resource=https%3A%2F%2Fmanagement.azure.com%2F",
				Response: map[string]any{
					"token_type":   "Bearer",
					"access_token": "bcd",
					"expires_on":   time.Now().Add(5 * time.Minute).Unix(),
				},
			},
			{
				Method:   "GET",
				Resource: "/a/b/c?api-version=2018-04-01",
				Response: `{"properties": {
					"workspaceUrl": "https://abc"
				}}`,
			},
			{
				Method:   "GET",
				Resource: "/metadata/identity/oauth2/token?api-version=2018-02-01&resource=62a912ac-b58e-4c1d-89ea-b2dbfc7358fc",
				Response: map[string]any{
					"token_type":   "Bearer",
					"access_token": "cde",
					"expires_on":   time.Now().Add(5 * time.Minute).Unix(),
				},
			},
			{
				Method:   "GET",
				Resource: "/metadata/identity/oauth2/token?api-version=2018-02-01&resource=https%3A%2F%2Fmanagement.core.windows.net%2F",
				Response: map[string]any{
					"token_type":   "Bearer",
					"access_token": "def",
					"expires_on":   time.Now().Add(5 * time.Minute).Unix(),
				},
			},
		},
	}
	req, _ := http.NewRequest("GET", "http://localhost", nil)
	err := cfg.Authenticate(req)
	require.NoError(t, err)
	require.Equal(t, "/a/b/c", req.Header.Get("X-Databricks-Azure-Workspace-Resource-Id"))
	require.Equal(t, "def", req.Header.Get("X-Databricks-Azure-SP-Management-Token"))
	require.Equal(t, "Bearer cde", req.Header.Get("Authorization"))
}

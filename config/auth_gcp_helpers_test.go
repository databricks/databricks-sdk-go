package config

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestRequiresGcpSaAccessToken_WorkspaceFromMetadata(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `", "workspace_id": "` + testHMWorkspaceID + `"}`,
			},
		},
	}
	result := requiresGcpSaAccessToken(context.Background(), cfg)
	assert.False(t, result, "workspace host should not require GCP SA token")
}

func TestRequiresGcpSaAccessToken_AccountFromMetadata(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       200,
				Response:     `{"oidc_endpoint": "` + testHMHost + `/oidc", "account_id": "` + testHMAccountID + `"}`,
			},
		},
	}
	result := requiresGcpSaAccessToken(context.Background(), cfg)
	assert.True(t, result, "account host (no workspace_id) should require GCP SA token")
}

func TestRequiresGcpSaAccessToken_MetadataError_FallsBackToAccountID(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:      testHMHost,
		AccountID: testHMAccountID,
		Loaders:   []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       500,
				Response:     `{"error": "internal error"}`,
			},
		},
	}
	result := requiresGcpSaAccessToken(context.Background(), cfg)
	assert.True(t, result, "with account_id set and metadata error, should require SA token")
}

func TestRequiresGcpSaAccessToken_MetadataError_NoAccountID(t *testing.T) {
	noopLoader := mockLoader(func(cfg *Config) error { return nil })
	cfg := &Config{
		Host:    testHMHost,
		Loaders: []Loader{noopLoader},
		HTTPTransport: fixtures.SliceTransport{
			{
				Method:       "GET",
				Resource:     "/.well-known/databricks-config",
				ReuseRequest: true,
				Status:       500,
				Response:     `{"error": "internal error"}`,
			},
		},
	}
	result := requiresGcpSaAccessToken(context.Background(), cfg)
	assert.False(t, result, "without account_id and metadata error, should not require SA token")
}

func TestRequiresGcpSaAccessToken_NoHost_WithAccountID(t *testing.T) {
	cfg := &Config{
		AccountID: testHMAccountID,
	}
	result := requiresGcpSaAccessToken(context.Background(), cfg)
	assert.True(t, result, "no host with account_id should require SA token")
}

func TestRequiresGcpSaAccessToken_NoHost_NoAccountID(t *testing.T) {
	cfg := &Config{}
	result := requiresGcpSaAccessToken(context.Background(), cfg)
	assert.False(t, result, "no host and no account_id should not require SA token")
}

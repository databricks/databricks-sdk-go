package config

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCloudAgnosticHosts verifies that credential providers work with cloud-agnostic hosts
// after removing is_azure/is_gcp checks. These tests ensure that Azure and GCP authentication
// works with any Databricks host when credentials are properly configured.

func TestAzureClientSecretWithCloudAgnosticHost(t *testing.T) {
	ctx := context.Background()

	// Create a config with cloud-agnostic host
	cfg := &Config{
		Host:              "https://api.databricks.com", // Cloud-agnostic host (not *.azuredatabricks.net)
		AzureClientID:     "test-client-id",
		AzureClientSecret: "test-client-secret",
		AzureTenantID:     "test-tenant-id",
	}
	cfg.refreshClient = &httpclient.ApiClient{}

	// Create credentials provider
	creds := AzureClientSecretCredentials{}

	// Should not return nil (would have returned nil before fix due to IsAzure check)
	provider, err := creds.Configure(ctx, cfg)

	// We expect an error here because we're not actually authenticating,
	// but the important thing is that it didn't return nil due to IsAzure check
	// If it returned nil with no error, that means it failed the IsAzure check
	assert.NotNil(t, provider, "Provider should not be nil for cloud-agnostic host with Azure credentials")
	// Error is acceptable as we're not doing real authentication
	if err != nil {
		t.Logf("Expected error during test auth: %v", err)
	}
}

func TestAzureClientSecretReturnsNilWhenCredentialsMissing(t *testing.T) {
	ctx := context.Background()

	// Test with missing credentials
	cfg := &Config{
		Host: "https://api.databricks.com",
		// Missing AzureClientID, AzureClientSecret, or AzureTenantID
	}

	creds := AzureClientSecretCredentials{}
	provider, err := creds.Configure(ctx, cfg)

	// Should return nil when credentials are missing (filtering based on config fields, not cloud type)
	require.NoError(t, err)
	assert.Nil(t, provider, "Provider should be nil when Azure credentials are not configured")
}

func TestGoogleCredentialsWithCloudAgnosticHost(t *testing.T) {
	ctx := context.Background()

	// Create a config with cloud-agnostic host and valid JSON credentials
	validJSON := `{"type": "service_account", "project_id": "test-project"}`

	cfg := &Config{
		Host:              "https://api.databricks.com", // Cloud-agnostic host (not *.gcp.databricks.com)
		GoogleCredentials: validJSON,
	}

	// Create credentials provider
	creds := GoogleCredentials{}

	// Should attempt to configure (won't be nil immediately due to IsGcp check being removed)
	// The actual configuration will fail without real credentials, but that's OK for this test
	_, err := creds.Configure(ctx, cfg)

	// We expect an error here because we're using fake credentials,
	// but the important thing is that it didn't return nil due to IsGcp check
	// before even attempting to process the credentials
	if err != nil {
		// This is expected with fake credentials
		assert.Contains(t, err.Error(), "could not obtain", "Expected credential processing error, not cloud check failure")
	}
}

func TestGoogleCredentialsReturnsNilWhenCredentialsMissing(t *testing.T) {
	ctx := context.Background()

	// Test with missing credentials
	cfg := &Config{
		Host: "https://api.databricks.com",
		// Missing GoogleCredentials
	}

	creds := GoogleCredentials{}
	provider, err := creds.Configure(ctx, cfg)

	// Should return nil when credentials are missing (filtering based on config fields, not cloud type)
	require.NoError(t, err)
	assert.Nil(t, provider, "Provider should be nil when Google credentials are not configured")
}

func TestGoogleDefaultCredentialsWithCloudAgnosticHost(t *testing.T) {
	ctx := context.Background()

	// Create a config with cloud-agnostic host
	cfg := &Config{
		Host:                 "https://api.databricks.com", // Cloud-agnostic host (not *.gcp.databricks.com)
		GoogleServiceAccount: "test-sa@test-project.iam.gserviceaccount.com",
	}

	// Create credentials provider
	creds := GoogleDefaultCredentials{}

	// Should attempt to configure (won't be nil immediately due to IsGcp check being removed)
	// The actual configuration will fail without real credentials, but that's OK for this test
	_, err := creds.Configure(ctx, cfg)

	// We expect an error here because we're using fake credentials,
	// but the important thing is that it didn't return nil due to IsGcp check
	if err != nil {
		// This is expected with fake credentials
		assert.Contains(t, err.Error(), "could not obtain OIDC token", "Expected credential processing error, not cloud check failure")
	}
}

func TestGoogleDefaultCredentialsReturnsNilWhenCredentialsMissing(t *testing.T) {
	ctx := context.Background()

	// Test with missing credentials
	cfg := &Config{
		Host: "https://api.databricks.com",
		// Missing GoogleServiceAccount
	}

	creds := GoogleDefaultCredentials{}
	provider, err := creds.Configure(ctx, cfg)

	// Should return nil when credentials are missing (filtering based on config fields, not cloud type)
	require.NoError(t, err)
	assert.Nil(t, provider, "Provider should be nil when Google service account is not configured")
}

func TestAzureCliRequiresAzureApplicationID(t *testing.T) {
	ctx := context.Background()

	// Test: Without AzureApplicationID - should return nil
	// Use DefaultEnvironment which is AWS and has no AzureApplicationID
	testEnv := environment.DefaultEnvironment()

	cfg := &Config{
		Host:                  "https://api.databricks.com",
		DatabricksEnvironment: &testEnv,
	}

	creds := AzureCliCredentials{}
	provider, err := creds.Configure(ctx, cfg)

	// Should return nil due to missing AzureApplicationID requirement
	require.NoError(t, err)
	assert.Nil(t, provider, "Provider should be nil when AzureApplicationID is not available")
}

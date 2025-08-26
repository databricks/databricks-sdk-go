package config

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/stretchr/testify/assert"
)

func TestAzureCliCredentials_FederatedTokenServicePrincipal(t *testing.T) {
	// This test verifies the fix where service principals authenticated via
	// federated token (like in AKS with workload identity) should be detected via GUID pattern
	// and skip tenant ID usage entirely (no fallback needed with the new approach).

	env.CleanupEnvironment(t)
	t.Setenv("PATH", testdataPath())

	// Simulate a service principal with a client ID (not systemAssignedIdentity/userAssignedIdentity)
	// This represents the federated token scenario from the JIRA issue
	t.Setenv("AZ_USER_NAME", "5817e630-86b3-4f67-a38e-a63e6a1a401c")
	t.Setenv("AZ_USER_TYPE", "servicePrincipal")

	// This would make the mock az command fail if --tenant were passed, but with the GUID
	// detection fix, --tenant should never be passed for this GUID-like service principal name
	t.Setenv("FAIL_IF_TENANT_ID_SET", "true")

	aa := AzureCliCredentials{}
	cfg := &Config{
		Host:          "https://adb-1891644720860465.5.azuredatabricks.net/",
		AzureTenantID: "e6a2f6d5-ece9-4c0d-9464-9c493497cb8f",
	}

	// With the GUID detection fix, this should work because the service principal name
	// matches the GUID pattern and is treated like MSI (no --tenant parameter used)
	visitor, err := aa.Configure(context.Background(), cfg)

	assert.NoError(t, err, "Authentication should work with federated token service principals via GUID detection")
	assert.NotNil(t, visitor, "Should return a valid credentials provider")
}

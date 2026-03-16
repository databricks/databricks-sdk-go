package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHostMetadataResolution(t *testing.T) {
	loadDebugEnvIfRunsFromIDE(t, "workspace")
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	envType := testEnvironmentType()
	if envType != "" {
		if envType != "WORKSPACE" && envType != "UC_WORKSPACE" {
			skipf(t)("Skipping workspace test: TEST_ENVIRONMENT_TYPE=%s", envType)
		}
	}
	t.Parallel()

	cfg := &config.Config{}
	err := cfg.EnsureResolved()
	require.NoError(t, err)

	assert.NotEmpty(t, cfg.AccountID, "expected account_id to be resolved from host metadata")
	// Workspace hosts should always have a workspace_id in metadata.
	assert.NotEmpty(t, cfg.WorkspaceID, "expected workspace_id to be resolved from host metadata")
}

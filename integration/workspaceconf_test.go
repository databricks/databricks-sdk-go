package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/settings/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccWorkspaceConf(t *testing.T) {
	ctx := workspaceTest(t)
	w, err := settings.NewWorkspaceConfClient(nil)
	assert.NoError(t, err)
	conf, err := w.GetStatus(ctx, settings.GetStatusRequest{
		Keys: "maxTokenLifetimeDays,enableIpAccessLists,enableWorkspaceFilesystem",
	})
	require.NoError(t, err)
	assert.Equal(t, 3, len(*conf))

	err = w.SetStatus(ctx, settings.WorkspaceConf{
		"enableWebTerminal": "true",
	})
	require.NoError(t, err)
}

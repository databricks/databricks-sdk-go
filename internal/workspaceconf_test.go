package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/workspaceconf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccWorkspaceConf(t *testing.T) {
	ctx, w := workspaceTest(t)
	conf, err := w.WorkspaceConf.GetStatus(ctx, workspaceconf.GetStatus{
		Keys: "maxTokenLifetimeDays,enableIpAccessLists,enableWorkspaceFilesystem",
	})
	require.NoError(t, err)
	assert.Equal(t, 3, len(*conf))

	err = w.WorkspaceConf.SetStatus(ctx, workspaceconf.WorkspaceConf{
		"enableWebTerminal": "true",
	})
	require.NoError(t, err)
}

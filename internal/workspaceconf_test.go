package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuxiaoshuo/databricks-sdk-go/service/settings"
)

func TestAccWorkspaceConf(t *testing.T) {
	ctx, w := workspaceTest(t)
	conf, err := w.WorkspaceConf.GetStatus(ctx, settings.GetStatusRequest{
		Keys: "maxTokenLifetimeDays,enableIpAccessLists,enableWorkspaceFilesystem",
	})
	require.NoError(t, err)
	assert.Equal(t, 3, len(*conf))

	err = w.WorkspaceConf.SetStatus(ctx, settings.WorkspaceConf{
		"enableWebTerminal": "true",
	})
	require.NoError(t, err)
}

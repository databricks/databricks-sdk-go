package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/stretchr/testify/assert"
)

func TestAccWorkspaceConf(t *testing.T) {
	ctx, w := workspaceTest(t)
	conf, err := w.WorkspaceConf.GetStatus(ctx, settings.GetStatusRequest{
		Keys: "maxTokenLifetimeDays,enableIpAccessLists,enableWorkspaceFilesystem",
	})
	assert.NoError(t, err)
	assert.Equal(t, 3, len(*conf))

	err = w.WorkspaceConf.SetStatus(ctx, settings.WorkspaceConf{
		"enableWebTerminal": "true",
	})
	assert.NoError(t, err)
}

package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccCommands(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	wsc := workspaces.Must(workspaces.NewClient())

	clusterId := GetEnvOrSkipTest(t, "TEST_GOSDK_CLUSTER_ID")

	info, err := wsc.Clusters.GetByClusterId(ctx, clusterId)
	require.NoError(t, err)
	if !info.IsRunningOrResizing() {
		_, err = wsc.Clusters.StartByClusterIdAndWait(ctx, clusterId)
		require.NoError(t, err)
	}

	res := wsc.CommandExecutor.Execute(ctx, clusterId, "python", "print(1)")
	require.NoError(t, res.Err())
	assert.Equal(t, "1", res.Text())

	res = wsc.CommandExecutor.Execute(ctx, clusterId, "python", "1/0")
	assert.EqualError(t, res.Err(), "ZeroDivisionError: division by zero")
}

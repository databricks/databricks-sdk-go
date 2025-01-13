package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccCommands(t *testing.T) {
	ctx, w := workspaceTest(t)

	clusterId := GetEnvOrSkipTest(t, "TEST_DEFAULT_CLUSTER_ID")
	commandContext, err := w.CommandExecution.Start(ctx, clusterId, compute.LanguagePython)
	require.NoError(t, err)

	t.Cleanup(func() {
		// TODO: add "destroys commandContext" examples converter hint
		err = commandContext.Destroy(ctx)
		require.NoError(t, err)
	})

	// TODO: figure out how to propagate this to examples converter
	res, err := commandContext.Execute(ctx, "print(1)")
	require.NoError(t, res.Err())
	assert.Equal(t, "1", res.Text())

	res, err = commandContext.Execute(ctx, "1/0")
	require.NoError(t, err)
	assert.EqualError(t, res.Err(), "ZeroDivisionError: division by zero")
}

func TestAccCommandsDirectUsage(t *testing.T) {
	ctx, w := workspaceTest(t)

	clusterId := GetEnvOrSkipTest(t, "TEST_DEFAULT_CLUSTER_ID")
	err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
	require.NoError(t, err)

	context, err := w.CommandExecution.CreateAndWait(ctx, compute.CreateContext{
		ClusterId: clusterId,
		Language:  compute.LanguagePython,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err = w.CommandExecution.Destroy(ctx, compute.DestroyContext{
			ClusterId: clusterId,
			ContextId: context.Id,
		})
		require.NoError(t, err)
	})

	textResults, err := w.CommandExecution.ExecuteAndWait(ctx, compute.Command{
		ClusterId: clusterId,
		ContextId: context.Id,
		Language:  compute.LanguagePython,
		Command:   "print(1)",
	})
	require.NoError(t, err)

	assert.Equal(t, "1", textResults.Results.Text())

	failingCommand, err := w.CommandExecution.ExecuteAndWait(ctx, compute.Command{
		ClusterId: clusterId,
		ContextId: context.Id,
		Language:  compute.LanguagePython,
		Command:   "1/0",
	})
	require.NoError(t, err)
	assert.EqualError(t, failingCommand.Results.Err(), "ZeroDivisionError: division by zero")
}

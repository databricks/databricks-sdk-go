package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/compute/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO:
// Enable or Delete this test when we have a decision on Start mixin in CommandExecutionAPI
// As this mixin have a cross dependency on the ClustersAPI.
// func TestAccCommands(t *testing.T) {
// 	ctx := workspaceTest(t)

// 	clusterId := GetEnvOrSkipTest(t, "TEST_DEFAULT_CLUSTER_ID")
// 	CommandExecutionAPI, err := compute.NewCommandExecutionClient(nil)
// 	require.NoError(t, err)
// 	commandContext, err := CommandExecutionAPI.Start(ctx, clusterId, compute.LanguagePython)
// 	require.NoError(t, err)

// 	t.Cleanup(func() {
// 		// TODO: add "destroys commandContext" examples converter hint
// 		err = commandContext.Destroy(ctx)
// 		require.NoError(t, err)
// 	})

// 	// TODO: figure out how to propagate this to examples converter
// 	res, err := commandContext.Execute(ctx, "print(1)")
// 	require.NoError(t, res.Err())
// 	assert.Equal(t, "1", res.Text())

// 	res, err = commandContext.Execute(ctx, "1/0")
// 	require.NoError(t, err)
// 	assert.EqualError(t, res.Err(), "ZeroDivisionError: division by zero")
// }

func TestAccCommandsDirectUsage(t *testing.T) {
	ctx := workspaceTest(t)

	clusterId := GetEnvOrSkipTest(t, "TEST_DEFAULT_CLUSTER_ID")
	ClustersAPI, err := compute.NewClustersClient(nil)
	require.NoError(t, err)
	err = ClustersAPI.EnsureClusterIsRunning(ctx, clusterId)
	require.NoError(t, err)

	CommandExecutionAPI, err := compute.NewCommandExecutionClient(nil)
	require.NoError(t, err)
	createWaiter, err := CommandExecutionAPI.Create(ctx, compute.CreateContext{
		ClusterId: clusterId,
		Language:  compute.LanguagePython,
	})
	require.NoError(t, err)
	context, err := createWaiter.WaitUntilDone(ctx, nil)
	require.NoError(t, err)

	t.Cleanup(func() {
		_, err = CommandExecutionAPI.Destroy(ctx, compute.DestroyContext{
			ClusterId: clusterId,
			ContextId: context.Id,
		})
		require.NoError(t, err)
	})

	executeWaiter, err := CommandExecutionAPI.Execute(ctx, compute.Command{
		ClusterId: clusterId,
		ContextId: context.Id,
		Language:  compute.LanguagePython,
		Command:   "print(1)",
	})
	require.NoError(t, err)
	textResults, err := executeWaiter.WaitUntilDone(ctx, nil)
	require.NoError(t, err)
	assert.Equal(t, "1", textResults.Results.Text())

	executeWaiter, err = CommandExecutionAPI.Execute(ctx, compute.Command{
		ClusterId: clusterId,
		ContextId: context.Id,
		Language:  compute.LanguagePython,
		Command:   "1/0",
	})
	require.NoError(t, err)
	failingCommand, err := executeWaiter.WaitUntilDone(ctx, nil)
	require.NoError(t, err)
	assert.EqualError(t, failingCommand.Results.Err(), "ZeroDivisionError: division by zero")
}

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
		err = commandContext.Destroy(ctx)
		require.NoError(t, err)
	})

	res, err := commandContext.Execute(ctx, "print(1)")
	require.NoError(t, res.Err())
	assert.Equal(t, "1", res.Text())
	
	res, err = commandContext.Execute(ctx, "1/0")
	require.NoError(t, err)
	assert.EqualError(t, res.Err(), "ZeroDivisionError: division by zero")
}

package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccCommands(t *testing.T) {
	ctx, w := workspaceTest(t)
	clusterId := sharedRunningClusterNoTranspile(t, ctx, w)

	res := w.CommandExecutor.Execute(ctx, clusterId, "python", "print(1)")
	require.NoError(t, res.Err())
	assert.Equal(t, "1", res.Text())

	res = w.CommandExecutor.Execute(ctx, clusterId, "python", "1/0")
	assert.EqualError(t, res.Err(), "ZeroDivisionError: division by zero")
}

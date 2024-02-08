package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUcAccWorkspaceClient_CurrentWorkspaceId(t *testing.T) {
	ctx, w := ucwsTest(t)
	expectedWorkspaceId := MustParseInt64(GetEnvOrSkipTest(t, "THIS_WORKSPACE_ID"))
	workspaceId, err := w.CurrentWorkspaceID(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedWorkspaceId, workspaceId)
}

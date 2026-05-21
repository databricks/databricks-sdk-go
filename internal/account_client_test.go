package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUcAccAccountClient_GetWorkspaceClient_NoTranspile(t *testing.T) {
	ctx, a := ucacctTest(t)
	workspaceId := MustParseInt64(GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID"))
	ws, err := a.Workspaces.GetByWorkspaceId(ctx, int64(workspaceId))
	require.NoError(t, err)
	w, err := a.GetWorkspaceClient(*ws)
	assert.NoError(t, err)
	me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
	assert.NoError(t, err)
	assert.True(t, me.Active)
}

func TestAccUnifiedHost_GetWorkspaceClient_NoTranspile(t *testing.T) {
	ctx, a, _ := unifiedHostAccountTest(t)
	workspaceId := MustParseInt64(GetEnvOrSkipTest(t, "TEST_WORKSPACE_ID"))
	ws, err := a.Workspaces.GetByWorkspaceId(ctx, int64(workspaceId))
	require.NoError(t, err)
	w, err := a.GetWorkspaceClient(*ws)
	assert.NoError(t, err)
	me, err := w.CurrentUser.Me(ctx, iam.MeRequest{})
	assert.NoError(t, err)
	assert.True(t, me.Active)
}

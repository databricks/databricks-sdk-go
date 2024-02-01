package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMwsAccAccountClient_GetWorkspaceClient_NoTranspile(t *testing.T) {
	ctx, a := accountTest(t)
	wss, err := a.Workspaces.List(ctx)
	require.NoError(t, err)

	if len(wss) == 0 {
		t.Skip("No workspaces found")
	}
	w, err := a.GetWorkspaceClient(wss[0])
	assert.NoError(t, err)
	me, err := w.CurrentUser.Me(ctx)
	assert.NoError(t, err)
	assert.True(t, me.Active)
}

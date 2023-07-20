package internal

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuxiaoshuo/databricks-sdk-go/service/compute"
)

func TestAccGlobalInitScripts(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.GlobalInitScripts.Create(ctx, compute.GlobalInitScriptCreateRequest{
		Name:     RandomName("go-sdk-"),
		Script:   base64.StdEncoding.EncodeToString([]byte("echo 1")),
		Enabled:  true,
		Position: 10,
	})
	require.NoError(t, err)

	defer w.GlobalInitScripts.DeleteByScriptId(ctx, created.ScriptId)

	err = w.GlobalInitScripts.Update(ctx, compute.GlobalInitScriptUpdateRequest{
		ScriptId: created.ScriptId,
		Name:     RandomName("go-sdk-updated-"),
		Script:   base64.StdEncoding.EncodeToString([]byte("echo 2")),
	})
	require.NoError(t, err)

	byId, err := w.GlobalInitScripts.GetByScriptId(ctx, created.ScriptId)
	require.NoError(t, err)

	all, err := w.GlobalInitScripts.ListAll(ctx)
	require.NoError(t, err)

	names, err := w.GlobalInitScripts.GlobalInitScriptDetailsNameToScriptIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, byId.ScriptId, names[byId.Name])

	byName, err := w.GlobalInitScripts.GetByName(ctx, byId.Name)
	require.NoError(t, err)
	assert.Equal(t, byName.ScriptId, byId.ScriptId)
}

package integration

import (
	"encoding/base64"
	"testing"

	"github.com/databricks/databricks-sdk-go/compute/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccGlobalInitScripts(t *testing.T) {
	ctx := workspaceTest(t)

	GlobalInitScriptsAPI, err := compute.NewGlobalInitScriptsClient(nil)
	created, err := GlobalInitScriptsAPI.Create(ctx, compute.GlobalInitScriptCreateRequest{
		Name:     RandomName("go-sdk-"),
		Script:   base64.StdEncoding.EncodeToString([]byte("echo 1")),
		Enabled:  true,
		Position: 10,
	})
	require.NoError(t, err)

	defer GlobalInitScriptsAPI.DeleteByScriptId(ctx, created.ScriptId)

	err = GlobalInitScriptsAPI.Update(ctx, compute.GlobalInitScriptUpdateRequest{
		ScriptId: created.ScriptId,
		Name:     RandomName("go-sdk-updated-"),
		Script:   base64.StdEncoding.EncodeToString([]byte("echo 2")),
	})
	require.NoError(t, err)

	byId, err := GlobalInitScriptsAPI.GetByScriptId(ctx, created.ScriptId)
	require.NoError(t, err)

	all, err := GlobalInitScriptsAPI.ListAll(ctx)
	require.NoError(t, err)

	names, err := GlobalInitScriptsAPI.GlobalInitScriptDetailsNameToScriptIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, byId.ScriptId, names[byId.Name])

	byName, err := GlobalInitScriptsAPI.GetByName(ctx, byId.Name)
	require.NoError(t, err)
	assert.Equal(t, byName.ScriptId, byId.ScriptId)
}

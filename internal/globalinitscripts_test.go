package internal

import (
	"encoding/base64"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/globalinitscripts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccGlobalInitScripts(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.GlobalInitScripts.CreateScript(ctx, globalinitscripts.GlobalInitScriptCreateRequest{
		Name:     RandomName("go-sdk-"),
		Script:   base64.StdEncoding.EncodeToString([]byte("echo 1")),
		Enabled:  true,
		Position: 10,
	})
	require.NoError(t, err)

	defer w.GlobalInitScripts.DeleteScriptByScriptId(ctx, created.ScriptId)

	err = w.GlobalInitScripts.UpdateScript(ctx, globalinitscripts.GlobalInitScriptUpdateRequest{
		ScriptId: created.ScriptId,
		Name:     RandomName("go-sdk-updated-"),
		Script:   base64.StdEncoding.EncodeToString([]byte("echo 2")),
	})
	require.NoError(t, err)

	script, err := w.GlobalInitScripts.GetScriptByScriptId(ctx, created.ScriptId)
	require.NoError(t, err)

	all, err := w.GlobalInitScripts.ListScripts(ctx)
	require.NoError(t, err)

	names, err := w.GlobalInitScripts.GlobalInitScriptDetailsNameToScriptIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, script.ScriptId, names[script.Name])

	// TODO: content on load?...
	byName, err := w.GlobalInitScripts.GetGlobalInitScriptDetailsByName(ctx, script.Name)
	require.NoError(t, err)
	assert.Equal(t, byName.ScriptId, script.ScriptId)
}

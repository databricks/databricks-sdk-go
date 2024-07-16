package internal

import (
	"encoding/base64"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/assert"
)

func TestAccGlobalInitScripts(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.GlobalInitScripts.Create(ctx, compute.GlobalInitScriptCreateRequest{
		Name:     RandomName("go-sdk-"),
		Script:   base64.StdEncoding.EncodeToString([]byte("echo 1")),
		Enabled:  true,
		Position: 10,
	})
	assert.NoError(t, err)

	defer w.GlobalInitScripts.DeleteByScriptId(ctx, created.ScriptId)

	err = w.GlobalInitScripts.Update(ctx, compute.GlobalInitScriptUpdateRequest{
		ScriptId: created.ScriptId,
		Name:     RandomName("go-sdk-updated-"),
		Script:   base64.StdEncoding.EncodeToString([]byte("echo 2")),
	})
	assert.NoError(t, err)

	byId, err := w.GlobalInitScripts.GetByScriptId(ctx, created.ScriptId)
	assert.NoError(t, err)

	all, err := w.GlobalInitScripts.ListAll(ctx)
	assert.NoError(t, err)

	names, err := w.GlobalInitScripts.GlobalInitScriptDetailsNameToScriptIdMap(ctx)
	assert.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, byId.ScriptId, names[byId.Name])

	byName, err := w.GlobalInitScripts.GetByName(ctx, byId.Name)
	assert.NoError(t, err)
	assert.Equal(t, byName.ScriptId, byId.ScriptId)
}

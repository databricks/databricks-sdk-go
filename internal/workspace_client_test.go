package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUcAccWorkspaceClient_CurrentWorkspaceId_NoTranspile(t *testing.T) {
	ctx, w := ucwsTest(t)
	expectedWorkspaceId := MustParseInt64(GetEnvOrSkipTest(t, "THIS_WORKSPACE_ID"))
	workspaceId, err := w.CurrentWorkspaceID(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedWorkspaceId, workspaceId)
}

func TestAccWorkspaceClient_AccountIDResolvedFromMetadata(t *testing.T) {
	_, w := workspaceTest(t)
	assert.NotEmpty(t, w.Config.AccountID, "account_id must be populated after config resolution")
}

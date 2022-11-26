package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/tokens"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccTokens(t *testing.T) {
	ctx, w := workspaceTest(t)

	token, err := w.Tokens.Create(ctx, tokens.CreateTokenRequest{
		Comment:         "xyz",
		LifetimeSeconds: 300,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Tokens.DeleteByTokenId(ctx, token.TokenInfo.TokenId)
		require.NoError(t, err)
	})

	names, err := w.Tokens.PublicTokenInfoCommentToTokenIdMap(ctx)
	require.NoError(t, err)
	assert.True(t, len(names) >= 1)

	wscInner := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Host:     w.Config.Host,
		Token:    token.TokenValue,
		AuthType: "pat",
	}))

	me, err := w.CurrentUser.Me(ctx)
	require.NoError(t, err)
	me2, err := wscInner.CurrentUser.Me(ctx)
	require.NoError(t, err)
	assert.Equal(t, me2.UserName, me.UserName)
}

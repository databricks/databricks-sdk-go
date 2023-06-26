package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccTokens(t *testing.T) {
	ctx, w := qa.WorkspaceTest(t)

	token, err := w.Tokens.Create(ctx, settings.CreateTokenRequest{
		Comment:         qa.RandomName("go-sdk-"),
		LifetimeSeconds: 300,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Tokens.DeleteByTokenId(ctx, token.TokenInfo.TokenId)
		require.NoError(t, err)
	})

	all, err := w.Tokens.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)

	byName, err := w.Tokens.GetByComment(ctx, token.TokenInfo.Comment)
	require.NoError(t, err)
	assert.Equal(t, token.TokenInfo.TokenId, byName.TokenId)

	wscInner := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Host:     w.Config.Host,
		Token:    token.TokenValue,
		AuthType: "pat",
	}))

	me2, err := wscInner.CurrentUser.Me(ctx)
	require.NoError(t, err)
	assert.Equal(t, me2.UserName, me(t, w).UserName)
}

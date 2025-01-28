package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/iam/v2"
	"github.com/databricks/databricks-sdk-go/settings/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccTokens(t *testing.T) {
	ctx := workspaceTest(t)

	TokensAPI, err := settings.NewTokensClient(nil)
	require.NoError(t, err)
	token, err := TokensAPI.Create(ctx, settings.CreateTokenRequest{
		Comment:         RandomName("go-sdk-"),
		LifetimeSeconds: 300,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = TokensAPI.DeleteByTokenId(ctx, token.TokenInfo.TokenId)
		require.NoError(t, err)
	})

	all, err := TokensAPI.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)

	byName, err := TokensAPI.GetByComment(ctx, token.TokenInfo.Comment)
	require.NoError(t, err)
	assert.Equal(t, token.TokenInfo.TokenId, byName.TokenId)

	CurrentUserAPIInner, err := iam.NewCurrentUserClient(&config.Config{
		Host:     TokensAPI.Config.Host,
		Token:    token.TokenValue,
		AuthType: "pat",
	})
	require.NoError(t, err)
	me2, err := CurrentUserAPIInner.Me(ctx)
	require.NoError(t, err)
	assert.Equal(t, me2.UserName, me(t, TokensAPI.Config).UserName)
}

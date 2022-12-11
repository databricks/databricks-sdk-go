package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/scim"
	"github.com/databricks/databricks-sdk-go/service/tokenmanagement"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOboToken(t *testing.T) {
	ctx, w := workspaceTest(t)
	if !w.Config.IsAws() {
		t.Skip("works only on aws")
	}
	groups, err := w.Groups.GroupDisplayNameToIdMap(ctx, scim.ListGroupsRequest{})
	require.NoError(t, err)

	spn, err := w.ServicePrincipals.Create(ctx, scim.ServicePrincipal{
		DisplayName: RandomName(t.Name()),
		Groups: []scim.ComplexValue{
			{
				Value: groups["admins"],
			},
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.ServicePrincipals.DeleteById(ctx, spn.Id)
		require.NoError(t, err)
	})

	obo, err := w.TokenManagement.CreateOboToken(ctx, tokenmanagement.CreateOboTokenRequest{
		ApplicationId:   spn.ApplicationId,
		LifetimeSeconds: 60,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.TokenManagement.DeleteByTokenId(ctx, obo.TokenInfo.TokenId)
		require.NoError(t, err)
	})

	byId, err := w.TokenManagement.GetByTokenId(ctx, obo.TokenInfo.TokenId)
	require.NoError(t, err)
	t.Log(byId)

	all, err := w.TokenManagement.ListAll(ctx, tokenmanagement.List{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

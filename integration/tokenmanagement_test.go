package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/iam/v2"
	"github.com/databricks/databricks-sdk-go/settings/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccCreateOboTokenOnAws(t *testing.T) {
	ctx := workspaceTest(t)
	cfg := &config.Config{}
	GroupsAPI, err := iam.NewGroupsClient(cfg)
	require.NoError(t, err)
	if !cfg.IsAws() {
		t.Skip("works only on aws")
	}
	groups, err := GroupsAPI.GroupDisplayNameToIdMap(ctx, iam.ListGroupsRequest{})
	require.NoError(t, err)

	ServicePrincipalsAPI, err := iam.NewServicePrincipalsClient(nil)
	require.NoError(t, err)
	spn, err := ServicePrincipalsAPI.Create(ctx, iam.ServicePrincipal{
		DisplayName: RandomName(t.Name()),
		Groups: []iam.ComplexValue{
			{
				Value: groups["admins"],
			},
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err = ServicePrincipalsAPI.DeleteById(ctx, spn.Id)
		require.NoError(t, err)
	})

	TokenManagementAPI, err := settings.NewTokenManagementClient(nil)
	require.NoError(t, err)
	obo, err := TokenManagementAPI.CreateOboToken(ctx, settings.CreateOboTokenRequest{
		ApplicationId:   spn.ApplicationId,
		LifetimeSeconds: 60,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err = TokenManagementAPI.DeleteByTokenId(ctx, obo.TokenInfo.TokenId)
		require.NoError(t, err)
	})

	byId, err := TokenManagementAPI.GetByTokenId(ctx, obo.TokenInfo.TokenId)
	require.NoError(t, err)
	t.Log(byId)

	all, err := TokenManagementAPI.ListAll(ctx, settings.ListTokenManagementRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/stretchr/testify/assert"
)

func TestAccCreateOboTokenOnAws(t *testing.T) {
	ctx, w := workspaceTest(t)
	if !w.Config.IsAws() {
		t.Skip("works only on aws")
	}
	groups, err := w.Groups.GroupDisplayNameToIdMap(ctx, iam.ListGroupsRequest{})
	assert.NoError(t, err)

	spn, err := w.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		DisplayName: RandomName(t.Name()),
		Groups: []iam.ComplexValue{
			{
				Value: groups["admins"],
			},
		},
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = w.ServicePrincipals.DeleteById(ctx, spn.Id)
		assert.NoError(t, err)
	})

	obo, err := w.TokenManagement.CreateOboToken(ctx, settings.CreateOboTokenRequest{
		ApplicationId:   spn.ApplicationId,
		LifetimeSeconds: 60,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = w.TokenManagement.DeleteByTokenId(ctx, obo.TokenInfo.TokenId)
		assert.NoError(t, err)
	})

	byId, err := w.TokenManagement.GetByTokenId(ctx, obo.TokenInfo.TokenId)
	assert.NoError(t, err)
	t.Log(byId)

	all, err := w.TokenManagement.ListAll(ctx, settings.ListTokenManagementRequest{})
	assert.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

package integration

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/apierr"
	"github.com/databricks/databricks-sdk-go/iam/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccAccountServicePrincipal(t *testing.T) {
	ctx, cfg := accountTest(t)
	ServicePrincipalsAPI, err := iam.NewAccountServicePrincipalsClient(cfg)
	require.NoError(t, err)
	spCreate, err := ServicePrincipalsAPI.Create(ctx, iam.ServicePrincipal{
		Active:      true,
		DisplayName: RandomName("go-sdk-sp-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err := ServicePrincipalsAPI.Delete(ctx, iam.DeleteAccountServicePrincipalRequest{Id: spCreate.Id})
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	sp, err := ServicePrincipalsAPI.GetById(ctx, spCreate.Id)
	require.NoError(t, err)
	assert.Equal(t, spCreate.Id, sp.Id)

	spGet, err := ServicePrincipalsAPI.Get(ctx, iam.GetAccountServicePrincipalRequest{Id: sp.Id})
	require.NoError(t, err)
	assert.Equal(t, spGet.DisplayName, sp.DisplayName)

	spGetByDisplayName, err := ServicePrincipalsAPI.GetByDisplayName(ctx, sp.DisplayName)
	require.NoError(t, err)
	assert.Equal(t, spGetByDisplayName.Id, sp.Id)

	spList, err := ServicePrincipalsAPI.ListAll(ctx, iam.ListAccountServicePrincipalsRequest{
		Filter: fmt.Sprintf("displayName eq %v", sp.DisplayName),
	})
	require.NoError(t, err)
	assert.Equal(t, spList[0].Id, sp.Id)

	_, err = ServicePrincipalsAPI.Patch(ctx, iam.PartialUpdate{
		Id: sp.Id,
		Operations: []iam.Patch{
			{
				Op:    iam.PatchOpReplace,
				Path:  "active",
				Value: "false",
			},
		},
		Schemas: []iam.PatchSchema{
			iam.PatchSchemaUrnIetfParamsScimApiMessages20PatchOp,
		},
	})
	require.NoError(t, err)

	_, err = ServicePrincipalsAPI.Update(ctx, iam.ServicePrincipal{
		Active:      true,
		DisplayName: sp.DisplayName,
		Id:          sp.Id,
	})
	require.NoError(t, err)
}

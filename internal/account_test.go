package internal

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccAccountServicePrincipal(t *testing.T) {
	ctx, a := accountTest(t)

	spCreate, err := a.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		Active:      true,
		DisplayName: RandomName("go-sdk-sp-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.ServicePrincipals.Delete(ctx, iam.DeleteAccountServicePrincipalRequest{Id: spCreate.Id})
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	sp, err := a.ServicePrincipals.GetById(ctx, spCreate.Id)
	require.NoError(t, err)
	assert.Equal(t, spCreate.Id, sp.Id)

	spGet, err := a.ServicePrincipals.Get(ctx, iam.GetAccountServicePrincipalRequest{Id: sp.Id})
	require.NoError(t, err)
	assert.Equal(t, spGet.DisplayName, sp.DisplayName)

	spGetByDisplayName, err := a.ServicePrincipals.GetByDisplayName(ctx, sp.DisplayName)
	require.NoError(t, err)
	assert.Equal(t, spGetByDisplayName.Id, sp.Id)

	spList, err := a.ServicePrincipals.ListAll(ctx, iam.ListAccountServicePrincipalsRequest{
		Filter: fmt.Sprintf("displayName eq %v", sp.DisplayName),
	})
	require.NoError(t, err)
	assert.Equal(t, spList[0].Id, sp.Id)

	err = a.ServicePrincipals.Patch(ctx, iam.PartialUpdate{
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

	err = a.ServicePrincipals.Update(ctx, iam.ServicePrincipal{
		Active:      true,
		DisplayName: sp.DisplayName,
		Id:          sp.Id,
	})
	require.NoError(t, err)
}

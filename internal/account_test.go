package internal

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
)

func TestAccAccountServicePrincipal(t *testing.T) {
	ctx, a := accountTest(t)

	spCreate, err := a.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		Active:      true,
		DisplayName: RandomName("go-sdk-sp-"),
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := a.ServicePrincipals.Delete(ctx, iam.DeleteAccountServicePrincipalRequest{Id: spCreate.Id})
		assert.True(t, err == nil || apierr.IsMissing(err))
	})

	sp, err := a.ServicePrincipals.GetById(ctx, spCreate.Id)
	assert.NoError(t, err)
	assert.Equal(t, spCreate.Id, sp.Id)

	spGet, err := a.ServicePrincipals.Get(ctx, iam.GetAccountServicePrincipalRequest{Id: sp.Id})
	assert.NoError(t, err)
	assert.Equal(t, spGet.DisplayName, sp.DisplayName)

	spGetByDisplayName, err := a.ServicePrincipals.GetByDisplayName(ctx, sp.DisplayName)
	assert.NoError(t, err)
	assert.Equal(t, spGetByDisplayName.Id, sp.Id)

	spList, err := a.ServicePrincipals.ListAll(ctx, iam.ListAccountServicePrincipalsRequest{
		Filter: fmt.Sprintf("displayName eq %v", sp.DisplayName),
	})
	assert.NoError(t, err)
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
	assert.NoError(t, err)

	err = a.ServicePrincipals.Update(ctx, iam.ServicePrincipal{
		Active:      true,
		DisplayName: sp.DisplayName,
		Id:          sp.Id,
	})
	assert.NoError(t, err)
}

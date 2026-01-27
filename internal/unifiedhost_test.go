package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccUnifiedHost(t *testing.T) {
	ctx, a := unifiedHostAccountTest(t)
	user, err := a.Users.Create(ctx, iam.User{
		DisplayName: RandomName("Me "),
		UserName:    RandomEmail(),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.Users.DeleteById(ctx, user.Id)
		require.True(t, err == nil || apierr.IsMissing(err))
	})

	assert.Equal(t, 0, len(user.Roles))
	err = a.Users.Patch(ctx, iam.PartialUpdate{
		Id:      user.Id,
		Schemas: []iam.PatchSchema{iam.PatchSchemaUrnIetfParamsScimApiMessages20PatchOp},
		Operations: []iam.Patch{
			{Op: iam.PatchOpAdd, Value: iam.User{
				Roles: []iam.ComplexValue{
					{Value: "account_admin"},
				},
			}},
		},
	})
	require.NoError(t, err)

	byId, err := a.Users.GetById(ctx, user.Id)
	require.NoError(t, err)
	assert.Equal(t, user.Id, byId.Id)
	assert.Equal(t, 1, len(byId.Roles))
}

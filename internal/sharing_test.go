package internal

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUcAccProviders(t *testing.T) {
	ctx, w := ucwsTest(t)

	// this is a publicly available test token
	publicShareRecipient := `{
		"shareCredentialsVersion":1,
		"bearerToken":"dapiabcdefghijklmonpqrstuvwxyz",
		"endpoint":"https://sharing.delta.io/delta-sharing/"
	}`
	created, err := w.Providers.Create(ctx, sharing.CreateProvider{
		Name:                RandomName("go-sdk-"),
		RecipientProfileStr: publicShareRecipient,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.Providers.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})
	_, err = w.Providers.Update(ctx, sharing.UpdateProvider{
		Name:    created.Name,
		Comment: "Comment for update",
	})
	require.NoError(t, err)

	_, err = w.Providers.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := w.Providers.ListAll(ctx, sharing.ListProvidersRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)

	shares, err := w.Providers.ListSharesAll(ctx, sharing.ListSharesRequest{
		Name: created.Name,
	})
	require.NoError(t, err)
	assert.Equal(t, "delta_sharing", shares[0].Name)
}

// TODO: remove NoTranspile
func TestUcAccRecipientActivationNoTranspile(t *testing.T) {
	ctx, w := ucwsTest(t)

	created, err := w.Recipients.Create(ctx, sharing.CreateRecipient{
		Name:               RandomName("go-sdk-"),
		AuthenticationType: sharing.AuthenticationTypeToken,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Recipients.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})

	split := strings.Split(created.Tokens[0].ActivationUrl, "?")
	recipientProfile, err := w.RecipientActivation.RetrieveTokenByActivationUrl(ctx, split[1])
	require.NoError(t, err)

	recipientProfileStr, err := json.Marshal(recipientProfile)
	require.NoError(t, err)

	newProvider, err := w.Providers.Create(ctx, sharing.CreateProvider{
		Name:                RandomName("go-sdk-"),
		RecipientProfileStr: string(recipientProfileStr),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Providers.DeleteByName(ctx, newProvider.Name)
		require.NoError(t, err)
	})

	_, err = w.Providers.ListSharesAll(ctx, sharing.ListSharesRequest{
		Name: newProvider.Name,
	})
	require.NoError(t, err)
}

func TestUcAccRecipients(t *testing.T) {
	ctx, w := ucwsTest(t)

	created, err := w.Recipients.Create(ctx, sharing.CreateRecipient{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.Recipients.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})
	err = w.Recipients.Update(ctx, sharing.UpdateRecipient{
		Name:    created.Name,
		Comment: RandomName("comment "),
	})
	require.NoError(t, err)

	recipientInfo, err := w.Recipients.RotateToken(ctx, sharing.RotateRecipientToken{
		Name:                         created.Name,
		ExistingTokenExpireInSeconds: 0,
	})
	require.NoError(t, err)
	assert.NotEqual(t, created.UpdatedAt, recipientInfo.UpdatedAt)

	sharePermissions, err := w.Recipients.SharePermissionsByName(ctx, created.Name)
	require.NoError(t, err)
	assert.Equal(t, 0, len(sharePermissions.PermissionsOut))

	_, err = w.Recipients.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := w.Recipients.ListAll(ctx, sharing.ListRecipientsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccShares(t *testing.T) {
	ctx, w := ucwsTest(t)

	createdCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: RandomName("catalog_"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
			Name:  createdCatalog.Name,
			Force: true,
		})
		require.NoError(t, err)
	})

	createdSchema, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        RandomName("schema_"),
		CatalogName: createdCatalog.Name,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Schemas.DeleteByFullName(ctx, createdSchema.FullName)
		require.NoError(t, err)
	})

	tableName := RandomName("foo_")
	tableFullName := fmt.Sprintf("%s.%s.%s", createdCatalog.Name, createdSchema.Name, tableName)

	// creates tableName
	_, err = w.StatementExecution.ExecuteAndWait(ctx, sql.ExecuteStatementRequest{
		WarehouseId: GetEnvOrSkipTest(t, "TEST_DEFAULT_WAREHOUSE_ID"),
		Catalog:     createdCatalog.Name,
		Schema:      createdSchema.Name,
		Statement:   fmt.Sprintf("CREATE TABLE %s AS SELECT 2+2 as four", tableName),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Tables.DeleteByFullName(ctx, tableFullName)
		require.NoError(t, err)
	})

	createdShare, err := w.Shares.Create(ctx, sharing.CreateShare{
		Name: RandomName("sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Shares.DeleteByName(ctx, createdShare.Name)
		require.NoError(t, err)
	})

	_, err = w.Shares.Update(ctx, sharing.UpdateShare{
		Name: createdShare.Name,
		Updates: []sharing.SharedDataObjectUpdate{
			{
				Action: sharing.SharedDataObjectUpdateActionAdd,
				DataObject: &sharing.SharedDataObject{
					Name:           tableFullName,
					DataObjectType: "TABLE",
				},
			},
		},
	})
	require.NoError(t, err)

	_, err = w.Shares.GetByName(ctx, createdShare.Name)
	require.NoError(t, err)

	all, err := w.Shares.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

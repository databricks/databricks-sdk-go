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
	assert.NoError(t, err)

	t.Cleanup(func() {
		err := w.Providers.DeleteByName(ctx, created.Name)
		assert.NoError(t, err)
	})
	_, err = w.Providers.Update(ctx, sharing.UpdateProvider{
		Name:    created.Name,
		Comment: "Comment for update",
	})
	assert.NoError(t, err)

	_, err = w.Providers.GetByName(ctx, created.Name)
	assert.NoError(t, err)

	all, err := w.Providers.ListAll(ctx, sharing.ListProvidersRequest{})
	assert.NoError(t, err)
	assert.True(t, len(all) >= 1)

	shares, err := w.Providers.ListSharesAll(ctx, sharing.ListSharesRequest{
		Name: created.Name,
	})
	assert.NoError(t, err)
	assert.Equal(t, "delta_sharing", shares[0].Name)
}

// TODO: remove NoTranspile
func TestUcAccRecipientActivationNoTranspile(t *testing.T) {
	ctx, w := ucwsTest(t)
	if w.Config.IsAzure() {
		skipf(t)("temporarily skipping this test on Azure until RetrieveToken uses the same host as specified in the activation URL")
	}

	created, err := w.Recipients.Create(ctx, sharing.CreateRecipient{
		Name:               RandomName("go-sdk-"),
		AuthenticationType: sharing.AuthenticationTypeToken,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := w.Recipients.DeleteByName(ctx, created.Name)
		assert.NoError(t, err)
	})

	split := strings.Split(created.Tokens[0].ActivationUrl, "?")
	recipientProfile, err := w.RecipientActivation.RetrieveTokenByActivationUrl(ctx, split[1])
	assert.NoError(t, err)

	recipientProfileStr, err := json.Marshal(recipientProfile)
	assert.NoError(t, err)

	newProvider, err := w.Providers.Create(ctx, sharing.CreateProvider{
		Name:                RandomName("go-sdk-"),
		RecipientProfileStr: string(recipientProfileStr),
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := w.Providers.DeleteByName(ctx, newProvider.Name)
		assert.NoError(t, err)
	})

	_, err = w.Providers.ListSharesAll(ctx, sharing.ListSharesRequest{
		Name: newProvider.Name,
	})
	assert.NoError(t, err)
}

func TestUcAccRecipients(t *testing.T) {
	ctx, w := ucwsTest(t)

	created, err := w.Recipients.Create(ctx, sharing.CreateRecipient{
		Name: RandomName("go-sdk-"),
	})
	assert.NoError(t, err)

	t.Cleanup(func() {
		err := w.Recipients.DeleteByName(ctx, created.Name)
		assert.NoError(t, err)
	})
	err = w.Recipients.Update(ctx, sharing.UpdateRecipient{
		Name:    created.Name,
		Comment: RandomName("comment "),
	})
	assert.NoError(t, err)

	recipientInfo, err := w.Recipients.RotateToken(ctx, sharing.RotateRecipientToken{
		Name:                         created.Name,
		ExistingTokenExpireInSeconds: 0,
	})
	assert.NoError(t, err)
	assert.NotEqual(t, created.UpdatedAt, recipientInfo.UpdatedAt)

	sharePermissions, err := w.Recipients.SharePermissionsByName(ctx, created.Name)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(sharePermissions.PermissionsOut))

	_, err = w.Recipients.GetByName(ctx, created.Name)
	assert.NoError(t, err)

	all, err := w.Recipients.ListAll(ctx, sharing.ListRecipientsRequest{})
	assert.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccShares(t *testing.T) {
	ctx, w := ucwsTest(t)
	if w.Config.IsGcp() {
		skipf(t)("Statement Execution API not available on GCP, skipping")
	}

	createdCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: RandomName("catalog_"),
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
			Name:  createdCatalog.Name,
			Force: true,
		})
		assert.NoError(t, err)
	})

	createdSchema, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        RandomName("schema_"),
		CatalogName: createdCatalog.Name,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := w.Schemas.DeleteByFullName(ctx, createdSchema.FullName)
		assert.NoError(t, err)
	})

	tableName := RandomName("foo_")
	tableFullName := fmt.Sprintf("%s.%s.%s", createdCatalog.Name, createdSchema.Name, tableName)

	// creates tableName
	_, err = w.StatementExecution.ExecuteAndWait(ctx, sql.ExecuteStatementRequest{
		WarehouseId: GetEnvOrSkipTest(t, "TEST_DEFAULT_WAREHOUSE_ID"),
		Catalog:     createdCatalog.Name,
		Schema:      createdSchema.Name,
		Statement:   fmt.Sprintf("CREATE TABLE %s TBLPROPERTIES (delta.enableDeletionVectors=false) AS SELECT 2+2 as four", tableName),
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = w.Tables.DeleteByFullName(ctx, tableFullName)
		assert.NoError(t, err)
	})

	createdShare, err := w.Shares.Create(ctx, sharing.CreateShare{
		Name: RandomName("sdk-"),
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := w.Shares.DeleteByName(ctx, createdShare.Name)
		assert.NoError(t, err)
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
	assert.NoError(t, err)

	_, err = w.Shares.GetByName(ctx, createdShare.Name)
	assert.NoError(t, err)

	all, err := w.Shares.ListAll(ctx)
	assert.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

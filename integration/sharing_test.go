package integration

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/catalog/v2"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/sharing/v2"
	"github.com/databricks/databricks-sdk-go/sql/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUcAccProviders(t *testing.T) {
	ctx := ucwsTest(t)

	// this is a publicly available test token
	publicShareRecipient := `{
		"shareCredentialsVersion":1,
		"bearerToken":"dapiabcdefghijklmonpqrstuvwxyz",
		"endpoint":"https://sharing.delta.io/delta-sharing/"
	}`
	ProvidersAPI, err := sharing.NewProvidersClient(nil)
	require.NoError(t, err)
	created, err := ProvidersAPI.Create(ctx, sharing.CreateProvider{
		Name:                RandomName("go-sdk-"),
		RecipientProfileStr: publicShareRecipient,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := ProvidersAPI.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})
	_, err = ProvidersAPI.Update(ctx, sharing.UpdateProvider{
		Name:    created.Name,
		Comment: "Comment for update",
	})
	require.NoError(t, err)

	_, err = ProvidersAPI.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := ProvidersAPI.ListAll(ctx, sharing.ListProvidersRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)

	shares, err := ProvidersAPI.ListSharesAll(ctx, sharing.ListSharesRequest{
		Name: created.Name,
	})
	require.NoError(t, err)
	assert.Equal(t, "delta_sharing", shares[0].Name)
}

// TODO: remove NoTranspile
func TestUcAccRecipientActivationNoTranspile(t *testing.T) {
	ctx := ucwsTest(t)
	cfg := &config.Config{}
	RecipientsAPI, err := sharing.NewRecipientsClient(cfg)
	require.NoError(t, err)
	if cfg.IsAzure() {
		skipf(t)("temporarily skipping this test on Azure until RetrieveToken uses the same host as specified in the activation URL")
	}

	created, err := RecipientsAPI.Create(ctx, sharing.CreateRecipient{
		Name:               RandomName("go-sdk-"),
		AuthenticationType: sharing.AuthenticationTypeToken,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := RecipientsAPI.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})

	split := strings.Split(created.Tokens[0].ActivationUrl, "?")
	RecipientActivationAPI, err := sharing.NewRecipientActivationClient(nil)
	require.NoError(t, err)
	recipientProfile, err := RecipientActivationAPI.RetrieveTokenByActivationUrl(ctx, split[1])
	require.NoError(t, err)

	recipientProfileStr, err := json.Marshal(recipientProfile)
	require.NoError(t, err)

	ProvidersAPI, err := sharing.NewProvidersClient(nil)
	require.NoError(t, err)
	newProvider, err := ProvidersAPI.Create(ctx, sharing.CreateProvider{
		Name:                RandomName("go-sdk-"),
		RecipientProfileStr: string(recipientProfileStr),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := ProvidersAPI.DeleteByName(ctx, newProvider.Name)
		require.NoError(t, err)
	})

	_, err = ProvidersAPI.ListSharesAll(ctx, sharing.ListSharesRequest{
		Name: newProvider.Name,
	})
	require.NoError(t, err)
}

func TestUcAccRecipients(t *testing.T) {
	ctx := ucwsTest(t)

	RecipientsAPI, err := sharing.NewRecipientsClient(nil)
	require.NoError(t, err)
	created, err := RecipientsAPI.Create(ctx, sharing.CreateRecipient{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := RecipientsAPI.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})
	_, err = RecipientsAPI.Update(ctx, sharing.UpdateRecipient{
		Name:    created.Name,
		Comment: RandomName("comment "),
	})
	require.NoError(t, err)

	recipientInfo, err := RecipientsAPI.RotateToken(ctx, sharing.RotateRecipientToken{
		Name:                         created.Name,
		ExistingTokenExpireInSeconds: 0,
	})
	require.NoError(t, err)
	assert.NotEqual(t, created.UpdatedAt, recipientInfo.UpdatedAt)

	sharePermissions, err := RecipientsAPI.SharePermissionsByName(ctx, created.Name)
	require.NoError(t, err)
	assert.Equal(t, 0, len(sharePermissions.PermissionsOut))

	_, err = RecipientsAPI.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := RecipientsAPI.ListAll(ctx, sharing.ListRecipientsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccShares(t *testing.T) {
	ctx := ucwsTest(t)
	cfg := &config.Config{}
	CatalogsAPI, err := catalog.NewCatalogsClient(cfg)
	require.NoError(t, err)
	if cfg.IsGcp() {
		skipf(t)("Statement Execution API not available on GCP, skipping")
	}

	createdCatalog, err := CatalogsAPI.Create(ctx, catalog.CreateCatalog{
		Name: RandomName("catalog_"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := CatalogsAPI.Delete(ctx, catalog.DeleteCatalogRequest{
			Name:  createdCatalog.Name,
			Force: true,
		})
		require.NoError(t, err)
	})

	SchemasAPI, err := catalog.NewSchemasClient(nil)
	require.NoError(t, err)
	createdSchema, err := SchemasAPI.Create(ctx, catalog.CreateSchema{
		Name:        RandomName("schema_"),
		CatalogName: createdCatalog.Name,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := SchemasAPI.DeleteByFullName(ctx, createdSchema.FullName)
		require.NoError(t, err)
	})

	tableName := RandomName("foo_")
	tableFullName := fmt.Sprintf("%s.%s.%s", createdCatalog.Name, createdSchema.Name, tableName)

	// creates tableName
	StatementExecutionAPI, err := sql.NewStatementExecutionClient(nil)
	require.NoError(t, err)
	_, err = StatementExecutionAPI.ExecuteAndWait(ctx, sql.ExecuteStatementRequest{
		WarehouseId: GetEnvOrSkipTest(t, "TEST_DEFAULT_WAREHOUSE_ID"),
		Catalog:     createdCatalog.Name,
		Schema:      createdSchema.Name,
		Statement:   fmt.Sprintf("CREATE TABLE %s TBLPROPERTIES (delta.enableDeletionVectors=false) AS SELECT 2+2 as four", tableName),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		TablesAPI, err := catalog.NewTablesClient(nil)
		require.NoError(t, err)
		err = TablesAPI.DeleteByFullName(ctx, tableFullName)
		require.NoError(t, err)
	})

	SharesAPI, err := sharing.NewSharesClient(nil)
	require.NoError(t, err)
	createdShare, err := SharesAPI.Create(ctx, sharing.CreateShare{
		Name: RandomName("sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := SharesAPI.DeleteByName(ctx, createdShare.Name)
		require.NoError(t, err)
	})

	_, err = SharesAPI.Update(ctx, sharing.UpdateShare{
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

	_, err = SharesAPI.GetByName(ctx, createdShare.Name)
	require.NoError(t, err)

	all, err := SharesAPI.ListAll(ctx, sharing.ListSharesRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

package internal

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/unitycatalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUcAccStorageCredentials(t *testing.T) {
	ctx, w := ucwsTest(t)
	if !w.Config.IsAws() {
		skipf(t)("not not aws")
	}

	// TODO: OpenAPI: retry protocol on late validation for storage
	// See https://github.com/databricks/terraform-provider-databricks/issues/1424
	created, err := w.StorageCredentials.Create(ctx, unitycatalog.CreateStorageCredential{
		Name: RandomName("go-sdk-"),
		AwsIamRole: &unitycatalog.AwsIamRole{
			RoleArn: GetEnvOrSkipTest(t, "TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.StorageCredentials.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})

	_, err = w.StorageCredentials.Update(ctx, unitycatalog.UpdateStorageCredential{
		Name:    created.Name,
		Comment: RandomName("comment "),
		AwsIamRole: &unitycatalog.AwsIamRole{
			RoleArn: GetEnvOrSkipTest(t, "TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	require.NoError(t, err)

	byName, err := w.StorageCredentials.GetByName(ctx, created.Name)
	require.NoError(t, err)
	assert.NotEqual(t, "", byName.Id)

	all, err := w.StorageCredentials.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccExternalLocations(t *testing.T) {
	ctx, w := ucwsTest(t)
	if !w.Config.IsAws() {
		skipf(t)("not not aws")
	}

	credential, err := w.StorageCredentials.Create(ctx, unitycatalog.CreateStorageCredential{
		Name: RandomName("go-sdk-"),
		AwsIamRole: &unitycatalog.AwsIamRole{
			RoleArn: GetEnvOrSkipTest(t, "TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.StorageCredentials.DeleteByName(ctx, credential.Name)
		require.NoError(t, err)
	})

	created, err := w.ExternalLocations.Create(ctx, unitycatalog.CreateExternalLocation{
		Name:           RandomName("go-sdk-"),
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", GetEnvOrSkipTest(t, "TEST_BUCKET"), RandomName("l-")),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.ExternalLocations.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})

	_, err = w.ExternalLocations.Update(ctx, unitycatalog.UpdateExternalLocation{
		Name:           created.Name,
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", GetEnvOrSkipTest(t, "TEST_BUCKET"), RandomName("l-")),
	})
	require.NoError(t, err)

	_, err = w.ExternalLocations.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := w.ExternalLocations.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccMetastores(t *testing.T) {
	t.Skip("metastore force delete doesn't work yet")
	ctx, w := ucwsTest(t)
	created, err := w.Metastores.Create(ctx, unitycatalog.CreateMetastore{
		Name:        RandomName("go-sdk-"),
		StorageRoot: fmt.Sprintf("s3://%s/%s", GetEnvOrSkipTest(t, "TEST_BUCKET"), RandomName("t=")),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Metastores.Delete(ctx, unitycatalog.DeleteMetastoreRequest{
			Id:    created.MetastoreId,
			Force: true,
		})
		require.NoError(t, err)
	})

	_, err = w.Metastores.Update(ctx, unitycatalog.UpdateMetastore{
		Id:   created.MetastoreId,
		Name: RandomName("go-sdk-updated"),
	})
	require.NoError(t, err)

	_, err = w.Metastores.GetById(ctx, created.MetastoreId)
	require.NoError(t, err)

	err = w.Metastores.Assign(ctx, unitycatalog.CreateMetastoreAssignment{
		MetastoreId: created.MetastoreId,
		WorkspaceId: GetEnvInt64OrSkipTest(t, "TEST_WORKSPACE_ID"),
	})
	require.NoError(t, err)

	// TODO: Fix once APP-1331 is implemented
	err = w.Metastores.Unassign(ctx, unitycatalog.UnassignRequest{
		MetastoreId: created.MetastoreId,
		WorkspaceId: GetEnvInt64OrSkipTest(t, "TEST_WORKSPACE_ID"),
	})
	require.NoError(t, err)

	_, err = w.Metastores.Summary(ctx)
	require.NoError(t, err)

	all, err := w.Metastores.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccCatalogs(t *testing.T) {
	t.Skip("needs force delete")
	ctx, w := ucwsTest(t)

	created, err := w.Catalogs.Create(ctx, unitycatalog.CreateCatalog{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Catalogs.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})

	_, err = w.Catalogs.Update(ctx, unitycatalog.UpdateCatalog{})
	require.NoError(t, err)

	_, err = w.Catalogs.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := w.Catalogs.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccSchemas(t *testing.T) {
	t.Skip("needs force delete")
	ctx, w := ucwsTest(t)

	catalog, err := w.Catalogs.Create(ctx, unitycatalog.CreateCatalog{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		// TODO: force delete
		err := w.Catalogs.DeleteByName(ctx, catalog.Name)
		require.NoError(t, err)
	})

	created, err := w.Schemas.Create(ctx, unitycatalog.CreateSchema{
		Name:        RandomName("go-sdk-"),
		CatalogName: catalog.Name,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.Schemas.DeleteByFullName(ctx, created.FullName)
		require.NoError(t, err)
	})
	_, err = w.Schemas.Update(ctx, unitycatalog.UpdateSchema{
		FullName: created.FullName,
		Comment:  RandomName("comment "),
	})
	require.NoError(t, err)

	_, err = w.Schemas.GetByFullName(ctx, created.FullName)
	require.NoError(t, err)

	all, err := w.Schemas.ListAll(ctx, unitycatalog.ListSchemasRequest{
		CatalogName: catalog.Name,
	})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccProviders(t *testing.T) {
	t.Skip("fixme")
	ctx, w := ucwsTest(t)

	created, err := w.Providers.Create(ctx, unitycatalog.CreateProvider{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.Providers.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})
	_, err = w.Providers.Update(ctx, unitycatalog.UpdateProvider{})
	require.NoError(t, err)

	_, err = w.Providers.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := w.Providers.ListAll(ctx, unitycatalog.ListProvidersRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccRecipients(t *testing.T) {
	ctx, w := ucwsTest(t)

	created, err := w.Recipients.Create(ctx, unitycatalog.CreateRecipient{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.Recipients.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})
	err = w.Recipients.Update(ctx, unitycatalog.UpdateRecipient{
		Name:    created.Name,
		Comment: RandomName("comment "),
	})
	require.NoError(t, err)

	_, err = w.Recipients.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := w.Recipients.ListAll(ctx, unitycatalog.ListRecipientsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccShares(t *testing.T) {
	t.Skip("expand with catalogs")
	ctx, w := ucwsTest(t)

	created, err := w.Shares.Create(ctx, unitycatalog.CreateShare{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.Shares.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})
	_, err = w.Shares.Update(ctx, unitycatalog.UpdateShare{
		Name: RandomName("go-sdk-"),
		Updates: []unitycatalog.SharedDataObjectUpdate{
			{
				Action: unitycatalog.SharedDataObjectUpdateActionAdd,
				DataObject: &unitycatalog.SharedDataObject{
					Name: RandomName("go-sdk-"),
				},
			},
		},
	})
	require.NoError(t, err)

	_, err = w.Shares.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := w.Shares.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

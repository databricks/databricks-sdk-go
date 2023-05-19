package internal

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
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
	created, err := w.StorageCredentials.Create(ctx, catalog.CreateStorageCredential{
		Name: RandomName("go-sdk-"),
		AwsIamRole: &catalog.AwsIamRole{
			RoleArn: GetEnvOrSkipTest(t, "TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.StorageCredentials.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})

	_, err = w.StorageCredentials.Update(ctx, catalog.UpdateStorageCredential{
		Name:    created.Name,
		Comment: RandomName("comment "),
		AwsIamRole: &catalog.AwsIamRole{
			RoleArn: GetEnvOrSkipTest(t, "TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	require.NoError(t, err)

	byName, err := w.StorageCredentials.GetByName(ctx, created.Name)
	require.NoError(t, err)
	assert.NotEqual(t, "", byName.Id)

	all, err := w.StorageCredentials.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccExternalLocations(t *testing.T) {
	ctx, w := ucwsTest(t)
	if !w.Config.IsAws() {
		skipf(t)("not not aws")
	}

	credential, err := w.StorageCredentials.Create(ctx, catalog.CreateStorageCredential{
		Name: RandomName("go-sdk-"),
		AwsIamRole: &catalog.AwsIamRole{
			RoleArn: GetEnvOrSkipTest(t, "TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.StorageCredentials.DeleteByName(ctx, credential.Name)
		require.NoError(t, err)
	})

	created, err := w.ExternalLocations.Create(ctx, catalog.CreateExternalLocation{
		Name:           RandomName("go-sdk-"),
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", GetEnvOrSkipTest(t, "TEST_BUCKET"), RandomName("l-")),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.ExternalLocations.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})

	_, err = w.ExternalLocations.Update(ctx, catalog.UpdateExternalLocation{
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
	ctx, w := ucwsTest(t)
	created, err := w.Metastores.Create(ctx, catalog.CreateMetastore{
		Name:        RandomName("go-sdk-"),
		StorageRoot: fmt.Sprintf("s3://%s/%s", GetEnvOrSkipTest(t, "TEST_BUCKET"), RandomName("t=")),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Metastores.Delete(ctx, catalog.DeleteMetastoreRequest{
			Id:    created.MetastoreId,
			Force: true,
		})
		require.NoError(t, err)
	})

	_, err = w.Metastores.Update(ctx, catalog.UpdateMetastore{
		Id:   created.MetastoreId,
		Name: RandomName("go-sdk-updated"),
	})
	require.NoError(t, err)

	_, err = w.Metastores.GetById(ctx, created.MetastoreId)
	require.NoError(t, err)

	err = w.Metastores.Assign(ctx, catalog.CreateMetastoreAssignment{
		MetastoreId: created.MetastoreId,
		WorkspaceId: GetEnvInt64OrSkipTest(t, "TEST_WORKSPACE_ID"),
	})
	require.NoError(t, err)

	err = w.Metastores.Unassign(ctx, catalog.UnassignRequest{
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
	ctx, w := ucwsTest(t)

	created, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
			Name:  created.Name,
			Force: true,
		})
		require.NoError(t, err)
	})

	_, err = w.Catalogs.Update(ctx, catalog.UpdateCatalog{
		Name:    created.Name,
		Comment: "updated",
	})
	require.NoError(t, err)

	_, err = w.Catalogs.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := w.Catalogs.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccSchemas(t *testing.T) {
	ctx, w := ucwsTest(t)

	newCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
			Name:  newCatalog.Name,
			Force: true,
		})
		require.NoError(t, err)
	})

	created, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        RandomName("go-sdk-"),
		CatalogName: newCatalog.Name,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := w.Schemas.DeleteByFullName(ctx, created.FullName)
		require.NoError(t, err)
	})
	_, err = w.Schemas.Update(ctx, catalog.UpdateSchema{
		FullName: created.FullName,
		Comment:  RandomName("comment "),
	})
	require.NoError(t, err)

	_, err = w.Schemas.GetByFullName(ctx, created.FullName)
	require.NoError(t, err)

	all, err := w.Schemas.ListAll(ctx, catalog.ListSchemasRequest{
		CatalogName: newCatalog.Name,
	})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

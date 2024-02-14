package internal

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUcAccVolumes(t *testing.T) {
	ctx, w := ucwsTest(t)
	if !w.Config.IsAws() {
		skipf(t)("not on aws")
	}

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

	storageCredential, err := w.StorageCredentials.Create(ctx, catalog.CreateStorageCredential{
		Name: RandomName("creds-"),
		AwsIamRole: &catalog.AwsIamRole{
			RoleArn: GetEnvOrSkipTest(t, "TEST_METASTORE_DATA_ACCESS_ARN"),
		},
		Comment: "created via SDK",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.StorageCredentials.DeleteByName(ctx, storageCredential.Name)
		require.NoError(t, err)
	})

	externalLocation, err := w.ExternalLocations.Create(ctx, catalog.CreateExternalLocation{
		Name:           RandomName("location-"),
		CredentialName: storageCredential.Name,
		Comment:        "created via SDK",
		Url:            "s3://" + GetEnvOrSkipTest(t, "TEST_BUCKET") + "/" + RandomName("somepath-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.ExternalLocations.DeleteByName(ctx, externalLocation.Name)
		require.NoError(t, err)
	})

	createdVolume, err := w.Volumes.Create(ctx, catalog.CreateVolumeRequestContent{
		CatalogName:     createdCatalog.Name,
		SchemaName:      createdSchema.Name,
		Name:            RandomName("volume_"),
		StorageLocation: externalLocation.Url,
		VolumeType:      catalog.VolumeTypeExternal,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Volumes.DeleteByName(ctx, createdVolume.FullName)
		require.NoError(t, err)
	})

	loadedVolume, err := w.Volumes.ReadByName(ctx, createdVolume.FullName)
	require.NoError(t, err)
	assert.Equal(t, createdVolume.Name, loadedVolume.Name)

	_, err = w.Volumes.Update(ctx, catalog.UpdateVolumeRequestContent{
		Name:    loadedVolume.FullName,
		Comment: "Updated volume comment",
	})
	require.NoError(t, err)

	allVolumes, err := w.Volumes.ListAll(ctx, catalog.ListVolumesRequest{
		CatalogName: createdCatalog.Name,
		SchemaName:  createdSchema.Name,
	})
	require.NoError(t, err)

	assert.Equal(t, 1, len(allVolumes))
}

func TestUcAccTables(t *testing.T) {
	ctx, w := ucwsTest(t)
	if w.Config.IsGcp() {
		skipf(t)("Statement Execution API not available on GCP, skipping")
	}

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

	// creates tableFullName
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

	allTables, err := w.Tables.ListAll(ctx, catalog.ListTablesRequest{
		CatalogName: createdCatalog.Name,
		SchemaName:  createdSchema.Name,
	})
	require.NoError(t, err)
	assert.Equal(t, 1, len(allTables))

	createdTable, err := w.Tables.GetByFullName(ctx, tableFullName)
	require.NoError(t, err)
	assert.Equal(t, createdCatalog.Name, createdTable.CatalogName)

	accountLevelGroupName := GetEnvOrSkipTest(t, "TEST_DATA_ENG_GROUP")
	x, err := w.Grants.Update(ctx, catalog.UpdatePermissions{
		FullName:      createdTable.FullName,
		SecurableType: catalog.SecurableTypeTable,
		Changes: []catalog.PermissionsChange{
			{
				Add: []catalog.Privilege{
					catalog.PrivilegeModify,
					catalog.PrivilegeSelect,
				},
				Principal: accountLevelGroupName,
			},
		},
	})
	require.NoError(t, err)
	assert.True(t, len(x.PrivilegeAssignments) > 0)

	grants, err := w.Grants.GetEffectiveBySecurableTypeAndFullName(ctx, catalog.SecurableTypeTable, createdTable.FullName)
	require.NoError(t, err)
	assert.True(t, len(grants.PrivilegeAssignments) > 0)

	summaries, err := w.Tables.ListSummariesAll(ctx, catalog.ListSummariesRequest{
		CatalogName:       createdCatalog.Name,
		SchemaNamePattern: createdSchema.Name,
	})
	require.NoError(t, err)
	assert.Equal(t, len(allTables), len(summaries))
}

func TestUcAccStorageCredentialsOnAws(t *testing.T) {
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

	all, err := w.StorageCredentials.ListAll(ctx, catalog.ListStorageCredentialsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccExternalLocationsOnAws(t *testing.T) {
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

	all, err := w.ExternalLocations.ListAll(ctx, catalog.ListExternalLocationsRequest{})
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

	currentMetastore, err := w.Metastores.Current(ctx)
	assert.NoError(t, err)
	assert.NotEqual(t, created.MetastoreId, currentMetastore.MetastoreId)

	_, err = w.Metastores.Update(ctx, catalog.UpdateMetastore{
		Id:      created.MetastoreId,
		NewName: RandomName("go-sdk-updated"),
	})
	require.NoError(t, err)

	_, err = w.Metastores.GetById(ctx, created.MetastoreId)
	require.NoError(t, err)

	workspaceId := MustParseInt64(GetEnvOrSkipTest(t, "DUMMY_WORKSPACE_ID"))

	err = w.Metastores.Assign(ctx, catalog.CreateMetastoreAssignment{
		MetastoreId: created.MetastoreId,
		WorkspaceId: workspaceId,
	})
	require.NoError(t, err)

	// w.Metastores.UpdateAssignment can only be done for the current WS

	err = w.Metastores.Unassign(ctx, catalog.UnassignRequest{
		MetastoreId: created.MetastoreId,
		WorkspaceId: workspaceId,
	})
	require.NoError(t, err)

	summary, err := w.Metastores.Summary(ctx)
	require.NoError(t, err)
	assert.Equal(t, currentMetastore.MetastoreId, summary.MetastoreId)

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

func TestUcAccCatalogWorkspaceBindings(t *testing.T) {
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
		Name:          created.Name,
		IsolationMode: catalog.IsolationModeIsolated,
	})
	require.NoError(t, err)

	thisWorkspaceID := MustParseInt64(GetEnvOrSkipTest(t, "THIS_WORKSPACE_ID"))
	_, err = retries.New[catalog.CurrentWorkspaceBindings](
		retries.OnErrors(apierr.ErrDeadlineExceeded),
		retries.WithTimeout(1*time.Minute),
	).Run(ctx, func(ctx context.Context) (*catalog.CurrentWorkspaceBindings, error) {
		return w.WorkspaceBindings.Update(ctx, catalog.UpdateWorkspaceBindings{
			Name:             created.Name,
			AssignWorkspaces: []int64{thisWorkspaceID},
		})
	})
	require.NoError(t, err)

	bindings, err := w.WorkspaceBindings.GetByName(ctx, created.Name)
	require.NoError(t, err)

	assert.Equal(t, thisWorkspaceID, bindings.Workspaces[0])
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

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/catalog/v2"
	"github.com/databricks/databricks-sdk-go/databricks/apierr"
	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/sql/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUcAccVolumes(t *testing.T) {
	ctx := ucwsTest(t)
	CatalogsAPI, err := catalog.NewCatalogsClient(nil)
	require.NoError(t, err)
	if !CatalogsAPI.Config.IsAws() {
		skipf(t)("not on aws")
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

	StorageCredentialsAPI, err := catalog.NewStorageCredentialsClient(nil)
	storageCredential, err := StorageCredentialsAPI.Create(ctx, catalog.CreateStorageCredential{
		Name: RandomName("creds-"),
		AwsIamRole: &catalog.AwsIamRoleRequest{
			RoleArn: GetEnvOrSkipTest(t, "TEST_METASTORE_DATA_ACCESS_ARN"),
		},
		Comment: "created via SDK",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := StorageCredentialsAPI.DeleteByName(ctx, storageCredential.Name)
		require.NoError(t, err)
	})

	ExternalLocationsAPI, err := catalog.NewExternalLocationsClient(nil)
	externalLocation, err := ExternalLocationsAPI.Create(ctx, catalog.CreateExternalLocation{
		Name:           RandomName("location-"),
		CredentialName: storageCredential.Name,
		Comment:        "created via SDK",
		Url:            "s3://" + GetEnvOrSkipTest(t, "TEST_BUCKET") + "/" + RandomName("somepath-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := ExternalLocationsAPI.DeleteByName(ctx, externalLocation.Name)
		require.NoError(t, err)
	})

	VolumesAPI, err := catalog.NewVolumesClient(nil)
	createdVolume, err := VolumesAPI.Create(ctx, catalog.CreateVolumeRequestContent{
		CatalogName:     createdCatalog.Name,
		SchemaName:      createdSchema.Name,
		Name:            RandomName("volume_"),
		StorageLocation: externalLocation.Url,
		VolumeType:      catalog.VolumeTypeExternal,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = VolumesAPI.DeleteByName(ctx, createdVolume.FullName)
		require.NoError(t, err)
	})

	loadedVolume, err := VolumesAPI.ReadByName(ctx, createdVolume.FullName)
	require.NoError(t, err)
	assert.Equal(t, createdVolume.Name, loadedVolume.Name)

	_, err = VolumesAPI.Update(ctx, catalog.UpdateVolumeRequestContent{
		Name:    loadedVolume.FullName,
		Comment: "Updated volume comment",
	})
	require.NoError(t, err)

	allVolumes, err := VolumesAPI.ListAll(ctx, catalog.ListVolumesRequest{
		CatalogName: createdCatalog.Name,
		SchemaName:  createdSchema.Name,
	})
	require.NoError(t, err)

	assert.Equal(t, 1, len(allVolumes))
}

func TestUcAccTables(t *testing.T) {
	ctx := ucwsTest(t)
	CatalogsAPI, err := catalog.NewCatalogsClient(nil)
	if CatalogsAPI.Config.IsGcp() {
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

	// creates tableFullName
	StatementExecutionAPI, err := sql.NewStatementExecutionClient(nil)
	require.NoError(t, err)
	_, err = StatementExecutionAPI.ExecuteAndWait(ctx, sql.ExecuteStatementRequest{
		WarehouseId: GetEnvOrSkipTest(t, "TEST_DEFAULT_WAREHOUSE_ID"),
		Catalog:     createdCatalog.Name,
		Schema:      createdSchema.Name,
		Statement:   fmt.Sprintf("CREATE TABLE %s AS SELECT 2+2 as four", tableName),
	})
	require.NoError(t, err)

	TablesAPI, err := catalog.NewTablesClient(nil)
	require.NoError(t, err)
	t.Cleanup(func() {
		err = TablesAPI.DeleteByFullName(ctx, tableFullName)
		require.NoError(t, err)
	})

	allTables, err := TablesAPI.ListAll(ctx, catalog.ListTablesRequest{
		CatalogName: createdCatalog.Name,
		SchemaName:  createdSchema.Name,
	})
	require.NoError(t, err)
	assert.Equal(t, 1, len(allTables))

	createdTable, err := TablesAPI.GetByFullName(ctx, tableFullName)
	require.NoError(t, err)
	assert.Equal(t, createdCatalog.Name, createdTable.CatalogName)

	accountLevelGroupName := GetEnvOrSkipTest(t, "TEST_DATA_ENG_GROUP")
	GrantsAPI, err := catalog.NewGrantsClient(nil)
	require.NoError(t, err)
	x, err := GrantsAPI.Update(ctx, catalog.UpdatePermissions{
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

	grants, err := GrantsAPI.GetEffectiveBySecurableTypeAndFullName(ctx, catalog.SecurableTypeTable, createdTable.FullName)
	require.NoError(t, err)
	assert.True(t, len(grants.PrivilegeAssignments) > 0)

	summaries, err := TablesAPI.ListSummariesAll(ctx, catalog.ListSummariesRequest{
		CatalogName:       createdCatalog.Name,
		SchemaNamePattern: createdSchema.Name,
	})
	require.NoError(t, err)
	assert.Equal(t, len(allTables), len(summaries))
}

func TestUcAccStorageCredentialsOnAws(t *testing.T) {
	ctx := ucwsTest(t)
	StorageCredentialsAPI, err := catalog.NewStorageCredentialsClient(nil)
	require.NoError(t, err)
	if !StorageCredentialsAPI.Config.IsAws() {
		skipf(t)("not not aws")
	}

	// TODO: OpenAPI: retry protocol on late validation for storage
	// See https://github.com/databricks/terraform-provider-databricks/issues/1424
	created, err := StorageCredentialsAPI.Create(ctx, catalog.CreateStorageCredential{
		Name: RandomName("go-sdk-"),
		AwsIamRole: &catalog.AwsIamRoleRequest{
			RoleArn: GetEnvOrSkipTest(t, "TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := StorageCredentialsAPI.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})

	_, err = StorageCredentialsAPI.Update(ctx, catalog.UpdateStorageCredential{
		Name:    created.Name,
		Comment: RandomName("comment "),
		AwsIamRole: &catalog.AwsIamRoleRequest{
			RoleArn: GetEnvOrSkipTest(t, "TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	require.NoError(t, err)

	byName, err := StorageCredentialsAPI.GetByName(ctx, created.Name)
	require.NoError(t, err)
	assert.NotEqual(t, "", byName.Id)

	all, err := StorageCredentialsAPI.ListAll(ctx, catalog.ListStorageCredentialsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccExternalLocationsOnAws(t *testing.T) {
	ctx := ucwsTest(t)
	StorageCredentialsAPI, err := catalog.NewStorageCredentialsClient(nil)
	require.NoError(t, err)
	if !StorageCredentialsAPI.Config.IsAws() {
		skipf(t)("not not aws")
	}

	credential, err := StorageCredentialsAPI.Create(ctx, catalog.CreateStorageCredential{
		Name: RandomName("go-sdk-"),
		AwsIamRole: &catalog.AwsIamRoleRequest{
			RoleArn: GetEnvOrSkipTest(t, "TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := StorageCredentialsAPI.DeleteByName(ctx, credential.Name)
		require.NoError(t, err)
	})

	ExternalLocationsAPI, err := catalog.NewExternalLocationsClient(nil)
	require.NoError(t, err)
	created, err := ExternalLocationsAPI.Create(ctx, catalog.CreateExternalLocation{
		Name:           RandomName("go-sdk-"),
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", GetEnvOrSkipTest(t, "TEST_BUCKET"), RandomName("l-")),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := ExternalLocationsAPI.DeleteByName(ctx, created.Name)
		require.NoError(t, err)
	})

	_, err = ExternalLocationsAPI.Update(ctx, catalog.UpdateExternalLocation{
		Name:           created.Name,
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", GetEnvOrSkipTest(t, "TEST_BUCKET"), RandomName("l-")),
	})
	require.NoError(t, err)

	_, err = ExternalLocationsAPI.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := ExternalLocationsAPI.ListAll(ctx, catalog.ListExternalLocationsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccMetastores(t *testing.T) {
	ctx := ucwsTest(t)
	MetastoresAPI, err := catalog.NewMetastoresClient(nil)
	created, err := MetastoresAPI.Create(ctx, catalog.CreateMetastore{
		Name:        RandomName("go-sdk-"),
		StorageRoot: fmt.Sprintf("s3://%s/%s", GetEnvOrSkipTest(t, "TEST_BUCKET"), RandomName("t=")),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := MetastoresAPI.Delete(ctx, catalog.DeleteMetastoreRequest{
			Id:    created.MetastoreId,
			Force: true,
		})
		require.NoError(t, err)
	})

	currentMetastore, err := MetastoresAPI.Current(ctx)
	assert.NoError(t, err)
	assert.NotEqual(t, created.MetastoreId, currentMetastore.MetastoreId)

	_, err = MetastoresAPI.Update(ctx, catalog.UpdateMetastore{
		Id:      created.MetastoreId,
		NewName: RandomName("go-sdk-updated"),
	})
	require.NoError(t, err)

	_, err = MetastoresAPI.GetById(ctx, created.MetastoreId)
	require.NoError(t, err)

	workspaceId := MustParseInt64(GetEnvOrSkipTest(t, "DUMMY_WORKSPACE_ID"))

	err = MetastoresAPI.Assign(ctx, catalog.CreateMetastoreAssignment{
		MetastoreId: created.MetastoreId,
		WorkspaceId: workspaceId,
	})
	require.NoError(t, err)

	// MetastoresAPI.UpdateAssignment can only be done for the current WS

	err = MetastoresAPI.Unassign(ctx, catalog.UnassignRequest{
		MetastoreId: created.MetastoreId,
		WorkspaceId: workspaceId,
	})
	require.NoError(t, err)

	summary, err := MetastoresAPI.Summary(ctx)
	require.NoError(t, err)
	assert.Equal(t, currentMetastore.MetastoreId, summary.MetastoreId)

	all, err := MetastoresAPI.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccCatalogs(t *testing.T) {
	ctx := ucwsTest(t)
	CatalogsAPI, err := catalog.NewCatalogsClient(nil)
	require.NoError(t, err)
	created, err := CatalogsAPI.Create(ctx, catalog.CreateCatalog{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := CatalogsAPI.Delete(ctx, catalog.DeleteCatalogRequest{
			Name:  created.Name,
			Force: true,
		})
		require.NoError(t, err)
	})

	_, err = CatalogsAPI.Update(ctx, catalog.UpdateCatalog{
		Name:    created.Name,
		Comment: "updated",
	})
	require.NoError(t, err)

	_, err = CatalogsAPI.GetByName(ctx, created.Name)
	require.NoError(t, err)

	all, err := CatalogsAPI.ListAll(ctx, catalog.ListCatalogsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestUcAccCatalogWorkspaceBindings(t *testing.T) {
	ctx := ucwsTest(t)
	CatalogsAPI, err := catalog.NewCatalogsClient(nil)
	require.NoError(t, err)
	created, err := CatalogsAPI.Create(ctx, catalog.CreateCatalog{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := CatalogsAPI.Delete(ctx, catalog.DeleteCatalogRequest{
			Name:  created.Name,
			Force: true,
		})
		require.NoError(t, err)
	})

	_, err = CatalogsAPI.Update(ctx, catalog.UpdateCatalog{
		Name:          created.Name,
		IsolationMode: catalog.CatalogIsolationModeIsolated,
	})
	require.NoError(t, err)

	thisWorkspaceID := MustParseInt64(GetEnvOrSkipTest(t, "THIS_WORKSPACE_ID"))

	WorkspaceBindingsAPI, err := catalog.NewWorkspaceBindingsClient(nil)
	require.NoError(t, err)
	_, err = retries.New[catalog.CurrentWorkspaceBindings](
		retries.OnErrors(apierr.ErrDeadlineExceeded),
		retries.WithTimeout(1*time.Minute),
	).Run(ctx, func(ctx context.Context) (*catalog.CurrentWorkspaceBindings, error) {
		return WorkspaceBindingsAPI.Update(ctx, catalog.UpdateWorkspaceBindings{
			Name:             created.Name,
			AssignWorkspaces: []int64{thisWorkspaceID},
		})
	})
	require.NoError(t, err)

	bindings, err := WorkspaceBindingsAPI.GetByName(ctx, created.Name)
	require.NoError(t, err)

	assert.Equal(t, thisWorkspaceID, bindings.Workspaces[0])
}

func TestUcAccSchemas(t *testing.T) {
	ctx := ucwsTest(t)

	CatalogsAPI, err := catalog.NewCatalogsClient(nil)
	require.NoError(t, err)
	newCatalog, err := CatalogsAPI.Create(ctx, catalog.CreateCatalog{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := CatalogsAPI.Delete(ctx, catalog.DeleteCatalogRequest{
			Name:  newCatalog.Name,
			Force: true,
		})
		require.NoError(t, err)
	})

	SchemasAPI, err := catalog.NewSchemasClient(nil)
	require.NoError(t, err)
	created, err := SchemasAPI.Create(ctx, catalog.CreateSchema{
		Name:        RandomName("go-sdk-"),
		CatalogName: newCatalog.Name,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := SchemasAPI.DeleteByFullName(ctx, created.FullName)
		require.NoError(t, err)
	})
	_, err = SchemasAPI.Update(ctx, catalog.UpdateSchema{
		FullName: created.FullName,
		Comment:  RandomName("comment "),
	})
	require.NoError(t, err)

	_, err = SchemasAPI.GetByFullName(ctx, created.FullName)
	require.NoError(t, err)

	all, err := SchemasAPI.ListAll(ctx, catalog.ListSchemasRequest{
		CatalogName: newCatalog.Name,
	})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

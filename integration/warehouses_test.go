package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/sql/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO: Enable this test when LRO is implemented
func TestAccSqlWarehouses(t *testing.T) {
	ctx := workspaceTest(t)

	WarehousesAPI, err := sql.NewWarehousesClient(nil)
	require.NoError(t, err)
	resp, err := WarehousesAPI.Create(ctx, sql.CreateWarehouseRequest{
		Name:           RandomName("go-sdk-"),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
		Tags: &sql.EndpointTags{
			CustomTags: []sql.EndpointTagPair{
				{
					Key:   "Owner",
					Value: "eng-dev-ecosystem-team_at_databricks.com",
				},
			},
		},
	})
	require.NoError(t, err)
	created, err := resp.WaitUntilDone(ctx, nil)
	require.NoError(t, err)

	t.Cleanup(func() {
		// err = WarehousesAPI.DeleteById(ctx, created.Id)
		_, err = WarehousesAPI.Delete(ctx, sql.DeleteWarehouseRequest{
			Id: created.Id,
		})
		require.NoError(t, err)
	})

	_, err = WarehousesAPI.Edit(ctx, sql.EditWarehouseRequest{
		Id:             created.Id,
		Name:           RandomName("go-sdk-updated-"),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
	})
	require.NoError(t, err)

	wh, err := WarehousesAPI.GetById(ctx, created.Id)
	require.NoError(t, err)

	all, err := WarehousesAPI.ListAll(ctx, sql.ListWarehousesRequest{})
	require.NoError(t, err)

	names, err := WarehousesAPI.EndpointInfoNameToIdMap(ctx, sql.ListWarehousesRequest{})
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, wh.Id, names[wh.Name])

	byName, err := WarehousesAPI.GetByName(ctx, wh.Name)
	require.NoError(t, err)
	assert.Equal(t, wh.Id, byName.Id)
}

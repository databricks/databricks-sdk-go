package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/warehouses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSqlWarehouses(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.Warehouses.CreateWarehouseAndWait(ctx, warehouses.CreateWarehouseRequest{
		Name:           RandomName("go-sdk-"),
		ClusterSize:    "2X-Small", // TODO: add enum
		MaxNumClusters: 1,
		AutoStopMins:   10,
	})
	require.NoError(t, err)

	defer w.Warehouses.DeleteWarehouseByIdAndWait(ctx, created.Id)

	err = w.Warehouses.EditWarehouse(ctx, warehouses.EditWarehouseRequest{
		Id:             created.Id,
		Name:           RandomName("go-sdk-updated-"),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
	})
	require.NoError(t, err)

	wh, err := w.Warehouses.GetWarehouseById(ctx, created.Id)
	require.NoError(t, err)

	all, err := w.Warehouses.ListWarehousesAll(ctx, warehouses.ListWarehousesRequest{})
	require.NoError(t, err)

	names, err := w.Warehouses.EndpointInfoNameToIdMap(ctx, warehouses.ListWarehousesRequest{})
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, wh.Id, names[wh.Name])

	byName, err := w.Warehouses.GetEndpointInfoByName(ctx, wh.Name)
	require.NoError(t, err)
	assert.Equal(t, wh.Id, byName.Id)
}

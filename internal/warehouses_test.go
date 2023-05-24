package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccSqlWarehouses(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.Warehouses.CreateAndWait(ctx, sql.CreateWarehouseRequest{
		Name:           RandomName("go-sdk-"),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err = w.Warehouses.DeleteById(ctx, created.Id)
		require.NoError(t, err)
	})

	err = w.Warehouses.Edit(ctx, sql.EditWarehouseRequest{
		Id:             created.Id,
		Name:           RandomName("go-sdk-updated-"),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
	})
	require.NoError(t, err)

	wh, err := w.Warehouses.GetById(ctx, created.Id)
	require.NoError(t, err)

	all, err := w.Warehouses.ListAll(ctx, sql.ListWarehousesRequest{})
	require.NoError(t, err)

	names, err := w.Warehouses.EndpointInfoNameToIdMap(ctx, sql.ListWarehousesRequest{})
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, wh.Id, names[wh.Name])

	byName, err := w.Warehouses.GetByName(ctx, wh.Name)
	require.NoError(t, err)
	assert.Equal(t, wh.Id, byName.Id)
}

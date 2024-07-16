package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/stretchr/testify/assert"
)

func TestAccSqlWarehouses(t *testing.T) {
	ctx, w := workspaceTest(t)

	created, err := w.Warehouses.CreateAndWait(ctx, sql.CreateWarehouseRequest{
		Name:           RandomName("go-sdk-"),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
	})
	assert.NoError(t, err)

	t.Cleanup(func() {
		err = w.Warehouses.DeleteById(ctx, created.Id)
		assert.NoError(t, err)
	})

	_, err = w.Warehouses.Edit(ctx, sql.EditWarehouseRequest{
		Id:             created.Id,
		Name:           RandomName("go-sdk-updated-"),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
	})
	assert.NoError(t, err)

	wh, err := w.Warehouses.GetById(ctx, created.Id)
	assert.NoError(t, err)

	all, err := w.Warehouses.ListAll(ctx, sql.ListWarehousesRequest{})
	assert.NoError(t, err)

	names, err := w.Warehouses.EndpointInfoNameToIdMap(ctx, sql.ListWarehousesRequest{})
	assert.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, wh.Id, names[wh.Name])

	byName, err := w.Warehouses.GetByName(ctx, wh.Name)
	assert.NoError(t, err)
	assert.Equal(t, wh.Id, byName.Id)
}

package sql_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/sql"
)

func ExampleWarehousesAPI_Create_sqlWarehouses() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Warehouses.CreateAndWait(ctx, sql.CreateWarehouseRequest{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	_, err = w.Warehouses.DeleteByIdAndWait(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleWarehousesAPI_Edit_sqlWarehouses() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Warehouses.CreateAndWait(ctx, sql.CreateWarehouseRequest{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = w.Warehouses.Edit(ctx, sql.EditWarehouseRequest{
		Id:             created.Id,
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	_, err = w.Warehouses.DeleteByIdAndWait(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleWarehousesAPI_EndpointInfoNameToIdMap_sqlWarehouses() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	names, err := w.Warehouses.EndpointInfoNameToIdMap(ctx, sql.ListWarehousesRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", names)

}

func ExampleWarehousesAPI_Get_sqlWarehouses() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Warehouses.CreateAndWait(ctx, sql.CreateWarehouseRequest{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		ClusterSize:    "2X-Small",
		MaxNumClusters: 1,
		AutoStopMins:   10,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	wh, err := w.Warehouses.GetById(ctx, created.Id)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", wh)

	// cleanup

	_, err = w.Warehouses.DeleteByIdAndWait(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleWarehousesAPI_ListAll_sqlWarehouses() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Warehouses.ListAll(ctx, sql.ListWarehousesRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

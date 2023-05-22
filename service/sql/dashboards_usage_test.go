// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package sql_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/sql"
)

func ExampleDashboardsAPI_Create_dashboards() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Dashboards.Create(ctx, sql.CreateDashboardRequest{
		Name:                    fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		DashboardFiltersEnabled: false,
		IsDraft:                 true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.Dashboards.DeleteByDashboardId(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleDashboardsAPI_Delete_dashboards() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Dashboards.Create(ctx, sql.CreateDashboardRequest{
		Name:                    fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		DashboardFiltersEnabled: false,
		IsDraft:                 true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = w.Dashboards.DeleteByDashboardId(ctx, created.Id)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Dashboards.DeleteByDashboardId(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleDashboardsAPI_Get_dashboards() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Dashboards.Create(ctx, sql.CreateDashboardRequest{
		Name:                    fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		DashboardFiltersEnabled: false,
		IsDraft:                 true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := w.Dashboards.GetByDashboardId(ctx, created.Id)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.Dashboards.DeleteByDashboardId(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleDashboardsAPI_ListAll_dashboards() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Dashboards.ListAll(ctx, sql.ListDashboardsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleDashboardsAPI_Restore_dashboards() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Dashboards.Create(ctx, sql.CreateDashboardRequest{
		Name:                    fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		DashboardFiltersEnabled: false,
		IsDraft:                 true,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = w.Dashboards.Restore(ctx, sql.RestoreDashboardRequest{
		DashboardId: created.Id,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Dashboards.DeleteByDashboardId(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}

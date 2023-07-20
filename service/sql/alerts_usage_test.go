// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package sql_test

import (
	"context"
	"fmt"
	"time"

	"github.com/xuxiaoshuo/databricks-sdk-go"
	"github.com/xuxiaoshuo/databricks-sdk-go/logger"

	"github.com/xuxiaoshuo/databricks-sdk-go/service/sql"
)

func ExampleAlertsAPI_Create_alerts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	srcs, err := w.DataSources.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", srcs)

	query, err := w.Queries.Create(ctx, sql.QueryPostContent{
		Name:         fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		DataSourceId: srcs[0].Id,
		Description:  "test query from Go SDK",
		Query:        "SELECT 1",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	alert, err := w.Alerts.Create(ctx, sql.CreateAlert{
		Options: sql.AlertOptions{
			Column: "1",
			Op:     "==",
			Value:  "1",
		},
		Name:    fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		QueryId: query.Id,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", alert)

	// cleanup

	err = w.Queries.DeleteByQueryId(ctx, query.Id)
	if err != nil {
		panic(err)
	}
	err = w.Alerts.DeleteByAlertId(ctx, alert.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleAlertsAPI_Get_alerts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	srcs, err := w.DataSources.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", srcs)

	query, err := w.Queries.Create(ctx, sql.QueryPostContent{
		Name:         fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		DataSourceId: srcs[0].Id,
		Description:  "test query from Go SDK",
		Query:        "SELECT 1",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	alert, err := w.Alerts.Create(ctx, sql.CreateAlert{
		Options: sql.AlertOptions{
			Column: "1",
			Op:     "==",
			Value:  "1",
		},
		Name:    fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		QueryId: query.Id,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", alert)

	byId, err := w.Alerts.GetByAlertId(ctx, alert.Id)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.Queries.DeleteByQueryId(ctx, query.Id)
	if err != nil {
		panic(err)
	}
	err = w.Alerts.DeleteByAlertId(ctx, alert.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleAlertsAPI_ListAll_alerts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Alerts.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleAlertsAPI_Update_alerts() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	srcs, err := w.DataSources.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", srcs)

	query, err := w.Queries.Create(ctx, sql.QueryPostContent{
		Name:         fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		DataSourceId: srcs[0].Id,
		Description:  "test query from Go SDK",
		Query:        "SELECT 1",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	alert, err := w.Alerts.Create(ctx, sql.CreateAlert{
		Options: sql.AlertOptions{
			Column: "1",
			Op:     "==",
			Value:  "1",
		},
		Name:    fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		QueryId: query.Id,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", alert)

	err = w.Alerts.Update(ctx, sql.EditAlert{
		Options: sql.AlertOptions{
			Column: "1",
			Op:     "==",
			Value:  "1",
		},
		AlertId: alert.Id,
		Name:    fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		QueryId: query.Id,
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Queries.DeleteByQueryId(ctx, query.Id)
	if err != nil {
		panic(err)
	}
	err = w.Alerts.DeleteByAlertId(ctx, alert.Id)
	if err != nil {
		panic(err)
	}

}

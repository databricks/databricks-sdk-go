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

func ExampleQueriesAPI_Create_queries() {
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
		Query:        "SHOW TABLES",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	// cleanup

	err = w.Queries.DeleteByQueryId(ctx, query.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleQueriesAPI_Create_alerts() {
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

	// cleanup

	err = w.Queries.DeleteByQueryId(ctx, query.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleQueriesAPI_Get_queries() {
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
		Query:        "SHOW TABLES",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	byId, err := w.Queries.GetByQueryId(ctx, query.Id)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.Queries.DeleteByQueryId(ctx, query.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleQueriesAPI_Update_queries() {
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
		Query:        "SHOW TABLES",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", query)

	updated, err := w.Queries.Update(ctx, sql.QueryEditContent{
		QueryId:      query.Id,
		Name:         fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		DataSourceId: srcs[0].Id,
		Description:  "UPDATED: test query from Go SDK",
		Query:        "SELECT 2+2",
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", updated)

	// cleanup

	err = w.Queries.DeleteByQueryId(ctx, query.Id)
	if err != nil {
		panic(err)
	}

}

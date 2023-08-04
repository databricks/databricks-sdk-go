package main

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/sql"
)

func main() {
	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{Profile: "aws-prod-ucws"}))
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelDebug,
	}
	ctx := context.Background()
	dashes, err := w.Dashboards.ListAll(ctx, sql.ListDashboardsRequest{})
	if err != nil {
		panic(err)
	}
	for _, dash := range dashes {
		w.Dashboards.Get(ctx, sql.GetDashboardRequest{DashboardId: dash.Id})
	}
}

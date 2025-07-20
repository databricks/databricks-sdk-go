package main

import (
	"context"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/apps"
)

func main() {

	cfg := &databricks.Config{
		Profile: "dbc-1232e87d-9384",
	}

	w := databricks.Must(databricks.NewWorkspaceClient(cfg))
	ctx := context.Background()

	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelDebug,
	}

	response, err := w.Apps.Create(ctx, apps.CreateAppRequest{
		App: apps.App{
			Name: "test-app",
		},
		NoCompute: true,
	})

	if err != nil {
		panic(err)
	}

	app, err := w.Apps.WaitGetAppActive(ctx, response.Name, 10*time.Second, func(app *apps.App) {
		logger.Infof(ctx, "app created: %s", app.Name)
	})
	if err != nil {
		panic(err)
	}

	_, err = w.Apps.Delete(ctx, apps.DeleteAppRequest{
		Name: app.Name,
	})
	if err != nil {
		panic(err)
	}

	logger.Infof(ctx, "app deleted: %s", response.Name)

}

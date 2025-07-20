package main

import (
	"context"
	"io"
	"os"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/files"
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

	response, err := w.Files.Download(ctx, files.DownloadRequest{
		FilePath: "/Volumes/parth-testing/default/parth_files_api/large_random_file.bin",
	})

	if err != nil {
		panic(err)
	}

	written, err := io.Copy(os.Stdout, response.Contents)
	logger.Infof(ctx, "written %d bytes", written)

	if err != nil {
		panic(err)
	}

	response.Contents.Close()
}

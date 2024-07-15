package main

import (
	"context"
	"log"
	"os"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

func main() {
	if os.Getenv("HTTPS_PROXY") != "https://localhost:8443" {
		log.Printf(`To run this example, see the instructions in examples/http-proxy/README.md.`)
		os.Exit(1)
	}
	log.Printf("Constructing client...")
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelDebug,
	}
	w := databricks.Must(databricks.NewWorkspaceClient())
	log.Printf("Listing clusters...")
	all, err := w.Clusters.ListAll(context.Background(), compute.ListClustersRequest{})
	if err != nil {
		panic(err)
	}
	for _, c := range all {
		println(c.ClusterName)
	}
}

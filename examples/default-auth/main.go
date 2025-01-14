package main

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

func main() {
	w := databricks.Must(databricks.NewWorkspaceClient())
	all, err := w.Clusters.ListAll(context.Background(), compute.ListClustersRequest{})
	if err != nil {
		panic(err)
	}
	println("Found %d clusters.", len(all))
}

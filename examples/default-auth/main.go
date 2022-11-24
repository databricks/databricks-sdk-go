package main

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/service/clusters"
)

func main() {
	w := databricks.Must(databricks.NewWorkspaceClient())
	all, err := w.Clusters.ListAll(context.Background(), clusters.ListRequest{})
	if err != nil {
		panic(err)
	}
	for _, c := range all {
		println(c.ClusterName)
	}
}

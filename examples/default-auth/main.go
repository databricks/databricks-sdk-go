package main

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/workspaces"
)

func main() {
	w := workspaces.New()
	all, err := w.Clusters.ListAll(context.Background(), clusters.ListRequest{})
	if err != nil {
		panic(err)
	}
	for _, c := range all {
		println(c.ClusterName)
	}
}

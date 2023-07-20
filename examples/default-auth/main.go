package main

import (
	"context"

	"github.com/xuxiaoshuo/databricks-sdk-go/service/compute"
)

func main() {
	w := databricks.Must(databricks.NewWorkspaceClient())
	all, err := w.Clusters.ListAll(context.Background(), compute.ListClustersRequest{})
	if err != nil {
		panic(err)
	}
	for _, c := range all {
		println(c.ClusterName)
	}
}

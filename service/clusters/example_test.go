package clusters_test

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/workspaces"
)

func ExampleClustersAPI() {
	w := workspaces.New()
	ctx := context.Background()

	// Fetch list of spark runtime versions
	sparkVersions, err := w.Clusters.SparkVersions(ctx)
	if err != nil {
		panic(err)
	}

	// Select the latest LTS version
	latestLTS, err := sparkVersions.Select(clusters.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	if err != nil {
		panic(err)
	}

	// Fetch list of available node types
	nodeTypes, err := w.Clusters.ListNodeTypes(ctx)
	if err != nil {
		panic(err)
	}

	// Select the smallest node type id
	smallestWithDisk, err := nodeTypes.Smallest(clusters.NodeTypeRequest{
		LocalDisk: true,
	})
	if err != nil {
		panic(err)
	}

	// Create cluster and wait until it's ready to use
	runningCluster, err := w.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
		ClusterName:            "Test cluster from SDK",
		SparkVersion:           latestLTS,
		NodeTypeId:             smallestWithDisk,
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Cluster is ready: %s#setting/clusters/%s/configuration\n",
		w.Config.Host, runningCluster.ClusterId)
}
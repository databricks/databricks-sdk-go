package compute_test

import (
	"context"
	"os"

	"github.com/databricks/databricks-sdk-go"

	"github.com/databricks/databricks-sdk-go/service/compute"
)

func ExampleLibrariesAPI_Update_libraries() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	clusterId := func() string {
		clusterId := os.Getenv("DATABRICKS_CLUSTER_ID")
		err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
		if err != nil {
			panic(err)
		}
		return clusterId
	}()

	err = w.Libraries.UpdateAndWait(ctx, compute.Update{
		ClusterId: clusterId,
		Install: []compute.Library{compute.Library{
			Pypi: &compute.PythonPyPiLibrary{
				Package: "dbl-tempo",
			},
		}},
	})
	if err != nil {
		panic(err)
	}

}

package integration

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/compute/v2"
	"github.com/stretchr/testify/require"
)

func sharedRunningCluster(t *testing.T, ctx context.Context, c *compute.ClustersClient) string {
	clusterId := GetEnvOrSkipTest(t, "TEST_GO_SDK_CLUSTER_ID")
	err := c.EnsureClusterIsRunning(ctx, clusterId)
	require.NoError(t, err)
	return clusterId
}

func TestAccLibraries(t *testing.T) {
	ctx := workspaceTest(t)
	lc, err := compute.NewLibrariesClient(nil)
	require.NoError(t, err)
	cc, err := compute.NewClustersClient(nil)
	require.NoError(t, err)
	clusterId := sharedRunningCluster(t, ctx, cc)

	err = lc.UpdateAndWait(ctx, compute.Update{
		ClusterId: clusterId,
		Install: []compute.Library{
			{
				Pypi: &compute.PythonPyPiLibrary{
					Package: "dbl-tempo",
				},
			},
		},
	})
	require.NoError(t, err)

	err = lc.UpdateAndWait(ctx, compute.Update{
		ClusterId: clusterId,
		Uninstall: []compute.Library{
			{
				Pypi: &compute.PythonPyPiLibrary{
					Package: "dbl-tempo",
				},
			},
		},
	})
	require.NoError(t, err)
}

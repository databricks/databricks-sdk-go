package internal

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccClustersCreateFailsWithTimeout(t *testing.T) {
	os.Setenv("DATABRICKS_CONFIG_PROFILE", "azure-deco")
	// t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()

	ctx := context.Background()
	w := workspaces.New()

	// Fetch list of spark runtime versions
	sparkVersions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	// Select the latest LTS version
	latestLTS, err := sparkVersions.Select(clusters.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	require.NoError(t, err)

	// Fetch list of available node types
	nodeTypes, err := w.Clusters.ListNodeTypes(ctx)
	require.NoError(t, err)

	// Select the smallest node type id
	smallestWithDisk, err := nodeTypes.Smallest(clusters.NodeTypeRequest{
		LocalDisk: true,
	})
	require.NoError(t, err)

	var clusterId string

	// Create a cluster with unreasonably low timeout
	_, err = w.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
		ClusterName:            RandomName(t.Name()),
		SparkVersion:           latestLTS,
		NodeTypeId:             smallestWithDisk,
		AutoterminationMinutes: 10,
		NumWorkers:             1,
	}, clusters.CreateTimeout(15*time.Second),
		func(i *retries.Info[clusters.ClusterInfo]) {
			clusterId = i.Info.ClusterId
		})
	assert.EqualError(t, err, "timed out: Finding instances for new nodes, acquiring more instances if necessary")
	_, err = w.Clusters.DeleteByClusterIdAndWait(ctx, clusterId)
	assert.NoError(t, err)
}

func TestAccClustersApiIntegration(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()

	ctx := context.Background()
	wsc := workspaces.New()

	clusterName := RandomName("sdk-go-cluster-")

	// Fetch list of spark runtime versions
	sparkVersions, err := wsc.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	// Select the latest LTS version
	latestLTS, err := sparkVersions.Select(clusters.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	require.NoError(t, err)

	// Fetch list of available node types
	nodeTypes, err := wsc.Clusters.ListNodeTypes(ctx)
	require.NoError(t, err)

	// Select the smallest node type id
	smallestWithDisk, err := nodeTypes.Smallest(clusters.NodeTypeRequest{
		LocalDisk: true,
	})
	require.NoError(t, err)

	// Create cluster and wait for it to start properly
	clstr, err := wsc.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latestLTS,
		NodeTypeId:             smallestWithDisk,
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	}, retries.Timeout[clusters.ClusterInfo](20*time.Minute))
	require.NoError(t, err)

	t.Cleanup(func() {
		// Permanently delete the cluster
		err := wsc.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
		require.NoError(t, err)
	})

	byId, err := wsc.Clusters.GetByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)
	assert.Equal(t, clusterName, byId.ClusterName)
	assert.Equal(t, clusters.ClusterInfoStateRunning, byId.State)

	// Pin the cluster in the list
	err = wsc.Clusters.PinByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// Unpin the cluster
	err = wsc.Clusters.UnpinByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// Edit the cluster: change auto-termination and number of workers
	err = wsc.Clusters.Edit(ctx, clusters.EditCluster{
		ClusterId:    clstr.ClusterId,
		SparkVersion: latestLTS,
		NodeTypeId:   smallestWithDisk,
		ClusterName:  clusterName,

		// change auto-termination and number of workers
		AutoterminationMinutes: 10,
		NumWorkers:             2,
	})
	require.NoError(t, err)

	// Assert edit changes are reflected in the cluster
	byId, err = wsc.Clusters.GetByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)
	assert.Equal(t, 10, byId.AutoterminationMinutes)
	assert.Equal(t, 2, byId.NumWorkers)

	// Terminate the cluster
	_, err = wsc.Clusters.DeleteByClusterIdAndWait(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// Assert that the cluster we've just deleted has Terminated state
	byId, err = wsc.Clusters.GetByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)
	assert.Equal(t, byId.State, clusters.ClusterInfoStateTerminated)

	// Start cluster and wait until it's running again
	_, err = wsc.Clusters.StartByClusterIdAndWait(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// Resize the cluster back to 1 worker and wait till completion
	byId, err = wsc.Clusters.ResizeAndWait(ctx, clusters.ResizeCluster{
		ClusterId:  clstr.ClusterId,
		NumWorkers: 1,
	})
	require.NoError(t, err)
	assert.Equal(t, 1, byId.NumWorkers)

	// Restart the cluster and wait for it to run again
	err = wsc.Clusters.Restart(ctx, clusters.RestartCluster{
		ClusterId: clstr.ClusterId,
	})
	require.NoError(t, err)

	// Get events for the cluster and assert its non empty
	getEventsResponse, err := wsc.Clusters.Events(ctx, clusters.GetEvents{
		ClusterId: clstr.ClusterId,
	})
	require.NoError(t, err)
	assert.True(t, len(getEventsResponse.Events) > 0)

	// List clusters in workspace
	listClustersResponse, err := wsc.Clusters.List(ctx, clusters.ListRequest{})
	require.NoError(t, err)

	var seen int
	for _, clusterInfo := range listClustersResponse.Clusters {
		if clusterInfo.ClusterId == clstr.ClusterId {
			seen++
		}
	}
	// The test clusters should only occur once in the list clusters response
	assert.True(t, seen == 1)
}

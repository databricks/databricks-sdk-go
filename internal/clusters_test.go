package internal

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccClustersCreateFailsWithTimeout(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()

	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

	// Fetch list of spark runtime versions
	sparkVersions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	// Select the latest LTS version
	latest, err := sparkVersions.Select(clusters.SparkVersionRequest{
		Latest: true,
	})
	require.NoError(t, err)

	var clusterId string

	// Create a cluster with unreasonably low timeout
	_, err = w.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
		ClusterName:            RandomName(t.Name()),
		SparkVersion:           latest,
		InstancePoolId:         GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 10,
		NumWorkers:             1,
	}, retries.Timeout[clusters.ClusterInfo](15*time.Second),
		func(i *retries.Info[clusters.ClusterInfo]) {
			clusterId = i.Info.ClusterId
		})
	assert.True(t, strings.HasPrefix(err.Error(), "timed out: "))
	_, err = w.Clusters.DeleteByClusterIdAndWait(ctx, clusterId)
	assert.NoError(t, err)
}

func TestAccClustersApiIntegration(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()

	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

	clusterName := RandomName("sdk-go-cluster-")

	// Fetch list of spark runtime versions
	sparkVersions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	// Select the latest LTS version
	latest, err := sparkVersions.Select(clusters.SparkVersionRequest{
		Latest: true,
	})
	require.NoError(t, err)

	// Create cluster and wait for it to start properly
	clstr, err := w.Clusters.CreateAndWait(ctx, clusters.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	}, retries.Timeout[clusters.ClusterInfo](20*time.Minute))
	require.NoError(t, err)

	t.Cleanup(func() {
		// Permanently delete the cluster
		err := w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
		require.NoError(t, err)
	})

	byId, err := w.Clusters.GetByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)
	assert.Equal(t, clusterName, byId.ClusterName)
	assert.Equal(t, clusters.StateRunning, byId.State)

	// Pin the cluster in the list
	err = w.Clusters.PinByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// Unpin the cluster
	err = w.Clusters.UnpinByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// Edit the cluster: change auto-termination and number of workers
	err = w.Clusters.Edit(ctx, clusters.EditCluster{
		ClusterId:      clstr.ClusterId,
		SparkVersion:   latest,
		ClusterName:    clusterName,
		InstancePoolId: GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),

		// change auto-termination and number of workers
		AutoterminationMinutes: 10,
		NumWorkers:             2,
	})
	require.NoError(t, err)

	// Assert edit changes are reflected in the cluster
	byId, err = w.Clusters.GetByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)
	assert.Equal(t, 10, byId.AutoterminationMinutes)
	assert.Equal(t, 2, byId.NumWorkers)

	// Terminate the cluster
	_, err = w.Clusters.DeleteByClusterIdAndWait(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// Assert that the cluster we've just deleted has Terminated state
	byId, err = w.Clusters.GetByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)
	assert.Equal(t, byId.State, clusters.StateTerminated)

	// Start cluster and wait until it's running again
	_, err = w.Clusters.StartByClusterIdAndWait(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// Resize the cluster back to 1 worker and wait till completion
	byId, err = w.Clusters.ResizeAndWait(ctx, clusters.ResizeCluster{
		ClusterId:  clstr.ClusterId,
		NumWorkers: 1,
	})
	require.NoError(t, err)
	assert.Equal(t, 1, byId.NumWorkers)

	// Restart the cluster and wait for it to run again
	err = w.Clusters.Restart(ctx, clusters.RestartCluster{
		ClusterId: clstr.ClusterId,
	})
	require.NoError(t, err)

	// Get events for the cluster and assert its non empty
	events, err := w.Clusters.EventsAll(ctx, clusters.GetEvents{
		ClusterId: clstr.ClusterId,
	})
	require.NoError(t, err)
	assert.True(t, len(events) > 0)

	// List clusters in workspace
	clusters, err := w.Clusters.ListAll(ctx, clusters.List{})
	require.NoError(t, err)

	var seen int
	for _, clusterInfo := range clusters {
		if clusterInfo.ClusterId == clstr.ClusterId {
			seen++
		}
	}
	// The test clusters should only occur once in the list clusters response
	assert.True(t, seen == 1)
}

package internal

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/scim"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// use mutex around starting TEST_GO_SDK_CLUSTER_ID
var mu sync.Mutex

func sharedRunningCluster(t *testing.T, ctx context.Context,
	w *databricks.WorkspaceClient) string {
	mu.Lock()
	defer mu.Unlock()

	clusterId := GetEnvOrSkipTest(t, "TEST_GO_SDK_CLUSTER_ID")
	info, err := w.Clusters.GetByClusterId(ctx, clusterId)
	require.NoError(t, err)

	switch info.State {
	case clusters.StateRunning:
		// noop
	case clusters.StateTerminated:
		_, err = w.Clusters.StartByClusterIdAndWait(ctx, clusterId)
		require.NoError(t, err)
	case clusters.StatePending,
		clusters.StateResizing,
		clusters.StateRestarting:
		_, err = w.Clusters.GetByClusterIdAndWait(ctx, clusterId)
		require.NoError(t, err)
	default:
		t.Errorf("Cluster is in %s state", info.State)
	}
	return clusterId
}

func TestAccClustersCreateFailsWithTimeout(t *testing.T) {
	ctx, w := workspaceTest(t)

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
			if i.Info == nil {
				return
			}
			clusterId = i.Info.ClusterId
		})
	assert.True(t, strings.HasPrefix(err.Error(), "timed out: "))
	_, err = w.Clusters.DeleteByClusterIdAndWait(ctx, clusterId)
	require.NoError(t, err)
}

func TestAccAwsInstanceProfiles(t *testing.T) {
	ctx, w := workspaceTest(t)
	if !w.Config.IsAws() {
		t.Skipf("runs only on AWS")
	}

	arn := "arn:aws:iam::000000000000:instance-profile/abc"
	err := w.InstanceProfiles.Add(ctx, clusters.AddInstanceProfile{
		InstanceProfileArn: arn,
		SkipValidation:     true,
		IamRoleArn:         "arn:aws:iam::000000000000:role/bcd",
	})
	require.NoError(t, err)

	defer w.InstanceProfiles.RemoveByInstanceProfileArn(ctx, arn)

	err = w.InstanceProfiles.Edit(ctx, clusters.InstanceProfile{
		InstanceProfileArn: arn,
		IamRoleArn:         "arn:aws:iam::000000000000:role/bcdf",
	})
	require.NoError(t, err)

	all, err := w.InstanceProfiles.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestAccClustersApiIntegration(t *testing.T) {
	ctx, w := workspaceTest(t)

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
	}, retries.Timeout[clusters.ClusterInfo](20*time.Minute),
		retries.OnPoll(func(i *clusters.ClusterInfo) {
			t.Logf("cluster is %s", i.State)
		}))
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

	byName, err := w.Clusters.GetByClusterName(ctx, byId.ClusterName)
	require.NoError(t, err)
	assert.Equal(t, byId.ClusterId, byName.ClusterId)

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

	all, err := w.Clusters.ListAll(ctx, clusters.List{})
	require.NoError(t, err)

	// List clusters in workspace
	names, err := w.Clusters.ClusterInfoClusterNameToClusterIdMap(ctx,
		clusters.List{})
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Contains(t, names, clusterName)

	otherOwner, err := w.Users.Create(ctx, scim.User{
		UserName: RandomEmail("sdk-go-"),
	})
	require.NoError(t, err)
	defer w.Users.DeleteById(ctx, otherOwner.Id)

	// terminate the cluster
	_, err = w.Clusters.DeleteByClusterIdAndWait(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// cluster must be terminated to change the owner
	err = w.Clusters.ChangeOwner(ctx, clusters.ChangeClusterOwner{
		ClusterId:     clstr.ClusterId,
		OwnerUsername: otherOwner.UserName,
	})
	require.NoError(t, err)

	nodes, err := w.Clusters.ListNodeTypes(ctx)
	require.NoError(t, err)
	assert.True(t, len(nodes.NodeTypes) > 1)

	if w.Config.IsAws() {
		zones, err := w.Clusters.ListZones(ctx)
		require.NoError(t, err)
		assert.True(t, len(zones.Zones) > 1)
	}
}

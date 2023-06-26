package internal

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func sharedRunningCluster(t *testing.T, ctx context.Context,
	w *databricks.WorkspaceClient) string {
	clusterId := qa.GetEnvOrSkipTest(t, "TEST_GO_SDK_CLUSTER_ID")
	err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
	require.NoError(t, err)
	return clusterId
}

func TestAccClustersCreateFailsWithTimeoutNoTranspile(t *testing.T) {
	ctx, w := qa.WorkspaceTest(t)

	// Fetch list of spark runtime versions
	sparkVersions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	// Select the latest LTS version
	latest, err := sparkVersions.Select(compute.SparkVersionRequest{
		Latest: true,
	})
	require.NoError(t, err)

	var clusterId string

	// Create a cluster with unreasonably low timeout
	_, err = w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            qa.RandomName(t.Name()),
		SparkVersion:           latest,
		InstancePoolId:         qa.GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 10,
		NumWorkers:             1,
	}, retries.Timeout[compute.ClusterInfo](15*time.Second),
		func(i *retries.Info[compute.ClusterInfo]) {
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
	ctx, w := qa.WorkspaceTest(t)
	if !w.Config.IsAws() {
		t.Skipf("runs only on AWS")
	}

	arn := "arn:aws:iam::000000000000:instance-profile/abc"
	err := w.InstanceProfiles.Add(ctx, compute.AddInstanceProfile{
		InstanceProfileArn: arn,
		SkipValidation:     true,
		IamRoleArn:         "arn:aws:iam::000000000000:role/bcd",
	})
	require.NoError(t, err)

	defer w.InstanceProfiles.RemoveByInstanceProfileArn(ctx, arn)

	err = w.InstanceProfiles.Edit(ctx, compute.InstanceProfile{
		InstanceProfileArn: arn,
		IamRoleArn:         "arn:aws:iam::000000000000:role/bcdf",
	})
	require.NoError(t, err)

	all, err := w.InstanceProfiles.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestAccClustersApiIntegration(t *testing.T) {
	ctx, w := qa.WorkspaceTest(t)

	clusterName := qa.RandomName("sdk-go-cluster-")

	// Select the latest LTS version
	latest, err := w.Clusters.SelectSparkVersion(ctx, compute.SparkVersionRequest{
		Latest: true,
	})
	require.NoError(t, err)

	// Create cluster and wait for it to start properly
	clstr, err := w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            clusterName,
		SparkVersion:           latest,
		InstancePoolId:         qa.GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 15,
		NumWorkers:             1,
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		// Permanently delete the cluster
		err := w.Clusters.PermanentDeleteByClusterId(ctx, clstr.ClusterId)
		require.NoError(t, err)
	})

	byId, err := w.Clusters.GetByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)
	assert.Equal(t, clusterName, byId.ClusterName)
	assert.Equal(t, compute.StateRunning, byId.State)

	// Pin the cluster in the list
	err = w.Clusters.PinByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// Unpin the cluster
	err = w.Clusters.UnpinByClusterId(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// Edit the cluster: change auto-termination and number of workers
	_, err = w.Clusters.EditAndWait(ctx, compute.EditCluster{
		ClusterId:      clstr.ClusterId,
		SparkVersion:   latest,
		ClusterName:    clusterName,
		InstancePoolId: qa.GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),

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
	assert.Equal(t, byId.State, compute.StateTerminated)

	byName, err := w.Clusters.GetByClusterName(ctx, byId.ClusterName)
	require.NoError(t, err)
	assert.Equal(t, byId.ClusterId, byName.ClusterId)

	// Start cluster and wait until it's running again
	_, err = w.Clusters.StartByClusterIdAndWait(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// Resize the cluster back to 1 worker and wait till completion
	byId, err = w.Clusters.ResizeAndWait(ctx, compute.ResizeCluster{
		ClusterId:  clstr.ClusterId,
		NumWorkers: 1,
	})
	require.NoError(t, err)
	assert.Equal(t, 1, byId.NumWorkers)

	// Restart the cluster and wait for it to run again
	_, err = w.Clusters.RestartAndWait(ctx, compute.RestartCluster{
		ClusterId: clstr.ClusterId,
	})
	require.NoError(t, err)

	// Get events for the cluster and assert its non empty
	events, err := w.Clusters.EventsAll(ctx, compute.GetEvents{
		ClusterId: clstr.ClusterId,
	})
	require.NoError(t, err)
	assert.True(t, len(events) > 0)

	// List clusters in workspace
	all, err := w.Clusters.ListAll(ctx, compute.ListClustersRequest{})
	require.NoError(t, err)
	foundCluster := false
	for _, info := range all {
		if info.ClusterName == clusterName {
			foundCluster = true
		}
	}
	assert.True(t, foundCluster)

	// Get cluster by name and assert it still exists
	clusterInfo, err := w.Clusters.GetByClusterName(ctx, clusterName)
	require.NoError(t, err)
	assert.Equal(t, clusterInfo.ClusterName, clusterName)

	otherOwner, err := w.Users.Create(ctx, iam.User{
		UserName: qa.RandomEmail(),
	})
	require.NoError(t, err)
	defer w.Users.DeleteById(ctx, otherOwner.Id)

	// terminate the cluster
	_, err = w.Clusters.DeleteByClusterIdAndWait(ctx, clstr.ClusterId)
	require.NoError(t, err)

	// cluster must be terminated to change the owner
	err = w.Clusters.ChangeOwner(ctx, compute.ChangeClusterOwner{
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

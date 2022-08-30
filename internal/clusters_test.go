package internal

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getDefaultNodeTypeForCloudEnv() string {
	if IsCloudEnvAzure() {
		return "Standard_D3_v2"
	} else if IsCloudEnvGcp() {
		return "n1-standard-4"
	}
	return "i3.xlarge"
}

func waitForTerminatingClusterToTerminate(ctx context.Context, workspacesClient workspaces.WorkspacesClient, t *testing.T, clusterId string) error {
	return retries.Wait(ctx, 30*time.Minute, func() *retries.Err {
		clusterInfo, err := workspacesClient.Clusters.GetCluster(
			ctx,
			clusters.GetClusterRequest{ClusterId: clusterId},
		)
		if err != nil {
			return retries.Halt(err)
		}
		if clusterInfo.State == clusters.GetClusterResponseStateTerminated {
			return nil
		}
		if clusterInfo.State != clusters.GetClusterResponseStateTerminating {
			// Maybe manually test that this path
			return retries.Halt(fmt.Errorf("Expected cluster %s to be in a TERMINATING state, instead its in %s state", clusterId, clusterInfo.State))
		}
		return retries.Continue(fmt.Errorf("%s is %s, but has to be TERMINATING", clusterId, clusterInfo.State))
	})
}

func waitForClusterRunning(ctx context.Context, workspacesClient workspaces.WorkspacesClient, t *testing.T, clusterId string) error {
	return retries.Wait(ctx, 30*time.Minute, func() *retries.Err {
		clusterInfo, err := workspacesClient.Clusters.GetCluster(
			ctx,
			clusters.GetClusterRequest{ClusterId: clusterId},
		)
		if err != nil {
			return retries.Halt(err)
		}
		if clusterInfo.State == clusters.GetClusterResponseStateRunning {
			return nil
		}
		if clusterInfo.State == clusters.GetClusterResponseStateTerminated ||
			clusterInfo.State == clusters.GetClusterResponseStateTerminating {
			// Maybe manually test that this path
			return retries.Halt(fmt.Errorf("cluster %s can't transition into a RUNNING state because its in %s state", clusterId, clusterInfo.State))
		}
		return retries.Continue(fmt.Errorf("%s is %s, but has to be RUNNING", clusterId, clusterInfo.State))
	})
}

func TestAccListClustersIntegration(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	ctx := context.TODO()
	workspacesClient := workspaces.New()
	const DEFAULT_RUNTIME_VERSION = "7.3.x-scala2.12"
	cluster1_name := RandomName("acceptance-test-cluster-1-")
	cluster2_name := RandomName("acceptance-test-cluster-2-")

	// Assert default node type used in this test (from getDefaultNodeTypeForCloudEnv)
	// is present in list node types
	listNodeTypesResponse, err := workspacesClient.Clusters.ListNodeTypes(ctx)
	require.NoError(t, err)
	isDefaultNodeTypePresent := false
	for _, nodeType := range listNodeTypesResponse.NodeTypes {
		if nodeType.NodeTypeId == getDefaultNodeTypeForCloudEnv() {
			isDefaultNodeTypePresent = true
			break
		}
	}
	assert.True(t, isDefaultNodeTypePresent)

	// Assert default runtime version used in this test (7.3.x-scala2.12)
	// is present spark versions api response
	getSparkVersionsResponse, err := workspacesClient.Clusters.GetSparkVersions(ctx)
	require.NoError(t, err)
	isDefaultRuntimeVersionPresent := false
	for _, runtimeVersion := range getSparkVersionsResponse.Versions {
		if runtimeVersion.Key == DEFAULT_RUNTIME_VERSION {
			isDefaultRuntimeVersionPresent = true
			break
		}
	}
	assert.True(t, isDefaultRuntimeVersionPresent)

	createClusterRequest := clusters.CreateClusterRequest{
		NumWorkers:             1,
		ClusterName:            cluster1_name,
		SparkVersion:           DEFAULT_RUNTIME_VERSION,
		AutoterminationMinutes: 15,
		NodeTypeId:             getDefaultNodeTypeForCloudEnv(),
	}

	// Create cluster
	clusterCreateResponse, err := workspacesClient.Clusters.CreateCluster(ctx, createClusterRequest)
	require.NoError(t, err)
	assert.NotEmpty(t, clusterCreateResponse.ClusterId)

	clusterId1 := clusterCreateResponse.ClusterId

	// Get cluster information. Assert cluster metadata is as expected
	getClusterResponse, err := workspacesClient.Clusters.GetCluster(
		ctx,
		clusters.GetClusterRequest{ClusterId: clusterId1},
	)
	require.NoError(t, err)
	assert.True(t, getClusterResponse.NumWorkers == createClusterRequest.NumWorkers)
	assert.True(t, getClusterResponse.ClusterName == createClusterRequest.ClusterName)
	assert.True(t, getClusterResponse.AutoterminationMinutes == createClusterRequest.AutoterminationMinutes)
	assert.True(t, getClusterResponse.NodeTypeId == createClusterRequest.NodeTypeId)
	assert.True(t, getClusterResponse.SparkVersion == createClusterRequest.SparkVersion)

	// Pin the cluster
	err = workspacesClient.Clusters.PinCluster(
		ctx,
		clusters.PinClusterRequest{
			ClusterId: clusterId1,
		},
	)
	require.NoError(t, err)

	// Unpin the cluster
	err = workspacesClient.Clusters.UnpinCluster(
		ctx,
		clusters.UnpinClusterRequest{
			ClusterId: clusterId1,
		},
	)
	require.NoError(t, err)

	// Terminate the cluster
	err = workspacesClient.Clusters.DeleteCluster(
		ctx,
		clusters.DeleteClusterRequest{
			ClusterId: clusterId1,
		},
	)
	require.NoError(t, err)

	// Assert the cluster state is TERMINATING or TERMINATED
	getClusterResponse, err = workspacesClient.Clusters.GetCluster(
		ctx,
		clusters.GetClusterRequest{ClusterId: clusterId1},
	)
	require.NoError(t, err)
	assert.True(
		t,
		getClusterResponse.State == clusters.GetClusterResponseStateTerminated ||
			getClusterResponse.State == clusters.GetClusterResponseStateTerminating,
	)

	// Wait until cluster reaches TERMINATED state
	err = waitForTerminatingClusterToTerminate(ctx, *workspacesClient, t, clusterId1)
	require.NoError(t, err)

	// Assert cluster state is terminated
	getClusterResponse, err = workspacesClient.Clusters.GetCluster(
		ctx,
		clusters.GetClusterRequest{ClusterId: clusterId1},
	)
	require.NoError(t, err)
	assert.True(t, getClusterResponse.State == clusters.GetClusterResponseStateTerminated)

	// Edit the cluster
	err = workspacesClient.Clusters.EditCluster(
		ctx,
		clusters.EditClusterRequest{
			ClusterId:              clusterId1,
			NumWorkers:             2,
			AutoterminationMinutes: 30,
			SparkVersion:           DEFAULT_RUNTIME_VERSION,
			NodeTypeId:             getDefaultNodeTypeForCloudEnv(),
			ClusterName:            cluster1_name,
		},
	)
	require.NoError(t, err)

	// Assert edit changes are reflected in the cluster
	getClusterResponse, err = workspacesClient.Clusters.GetCluster(
		ctx,
		clusters.GetClusterRequest{ClusterId: clusterId1},
	)
	require.NoError(t, err)
	assert.True(t, getClusterResponse.AutoterminationMinutes == 30)
	assert.True(t, getClusterResponse.NumWorkers == 2)
	assert.True(t, getClusterResponse.State == clusters.GetClusterResponseStateTerminated)

	// Start cluster 1
	err = workspacesClient.Clusters.StartCluster(
		ctx,
		clusters.StartClusterRequest{
			ClusterId: clusterId1,
		},
	)
	require.NoError(t, err)

	// Create cluster 2
	clusterCreateResponse2, err := workspacesClient.Clusters.CreateCluster(
		ctx,
		clusters.CreateClusterRequest{
			NumWorkers:             1,
			ClusterName:            cluster2_name,
			SparkVersion:           DEFAULT_RUNTIME_VERSION,
			AutoterminationMinutes: 15,
			NodeTypeId:             getDefaultNodeTypeForCloudEnv(),
		},
	)
	require.NoError(t, err)
	clusterId2 := clusterCreateResponse2.ClusterId

	// Wait until both cluster 1 and 2 are running
	err = waitForClusterRunning(ctx, *workspacesClient, t, clusterId1)
	require.NoError(t, err)
	err = waitForClusterRunning(ctx, *workspacesClient, t, clusterId2)
	require.NoError(t, err)

	// Assert cluster 1 is running
	getClusterResponse, err = workspacesClient.Clusters.GetCluster(
		ctx,
		clusters.GetClusterRequest{ClusterId: clusterId1},
	)
	require.NoError(t, err)
	assert.True(t, getClusterResponse.State == clusters.GetClusterResponseStateRunning)
	// Assert cluster 2 is running
	getClusterResponse, err = workspacesClient.Clusters.GetCluster(
		ctx,
		clusters.GetClusterRequest{ClusterId: clusterId2},
	)
	require.NoError(t, err)
	assert.True(t, getClusterResponse.State == clusters.GetClusterResponseStateRunning)

	// Resize num_workers in cluster 1
	err = workspacesClient.Clusters.ResizeCluster(
		ctx,
		clusters.ResizeClusterRequest{
			ClusterId:  clusterId1,
			NumWorkers: 3,
		},
	)
	require.NoError(t, err)
	// Assert num_workers in cluster 1 is 3 now
	getClusterResponse, err = workspacesClient.Clusters.GetCluster(
		ctx,
		clusters.GetClusterRequest{ClusterId: clusterId1},
	)
	require.NoError(t, err)
	assert.True(t, getClusterResponse.NumWorkers == 3)

	// Restart cluster 2
	err = workspacesClient.Clusters.RestartCluster(
		ctx,
		clusters.RestartClusterRequest{
			ClusterId: clusterId2,
		},
	)
	require.NoError(t, err)

	// Get events for cluster 1 and assert its non empty
	getEventsResponse, err := workspacesClient.Clusters.GetEvents(
		ctx,
		clusters.GetEventsRequest{
			ClusterId: clusterId1,
		},
	)
	require.NoError(t, err)
	assert.True(t, len(getEventsResponse.Events) > 0)

	// List clusters in workspace
	listClustersResponse, err := workspacesClient.Clusters.ListClusters(
		ctx,
		clusters.ListClustersRequest{},
	)
	require.NoError(t, err)
	numOccurancesOfTestCluster1 := 0
	numOccurancesOfTestCluster2 := 0
	for _, clusterInfo := range listClustersResponse.Clusters {
		if clusterInfo.ClusterId == clusterId1 {
			numOccurancesOfTestCluster1 += 1
		}
		if clusterInfo.ClusterId == clusterId2 {
			numOccurancesOfTestCluster2 += 1
		}
	}
	// The test clusters should only occur once in the list clusters response
	assert.True(t, numOccurancesOfTestCluster1 == 1)
	assert.True(t, numOccurancesOfTestCluster2 == 1)

	// Permanently delete the clusters
	err = workspacesClient.Clusters.PermanentDeleteCluster(
		ctx,
		clusters.PermanentDeleteClusterRequest{
			ClusterId: clusterId1,
		},
	)
	require.NoError(t, err)
	err = workspacesClient.Clusters.PermanentDeleteCluster(
		ctx,
		clusters.PermanentDeleteClusterRequest{
			ClusterId: clusterId2,
		},
	)
	require.NoError(t, err)
}

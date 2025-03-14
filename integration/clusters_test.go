package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/compute/v2"
	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO: Enable this test when LRO is implemented
//
// func sharedRunningCluster(t *testing.T, ctx context.Context,
// 	ClustersAPI *compute.ClustersClient) string {
// 	clusterId := GetEnvOrSkipTest(t, "TEST_GO_SDK_CLUSTER_ID")
// 	err := ClustersAPI.EnsureClusterIsRunning(ctx, clusterId)
// 	require.NoError(t, err)
// 	return clusterId
// }

// TODO: Enable this test when LRO is implemented
//
// func TestAccClustersCreateFailsWithTimeoutNoTranspile(t *testing.T) {
// 	ctx := workspaceTest(t)

// 	// Fetch list of spark runtime versions
// 	ClustersAPI, err := compute.NewClustersClient(nil)
// 	sparkVersions, err := ClustersAPI.SparkVersions(ctx)
// 	require.NoError(t, err)

// 	// Select the latest LTS version without Photon
// 	latest, err := sparkVersions.Select(compute.SparkVersionRequest{
// 		Latest:          true,
// 		LongTermSupport: true,
// 	})
// 	require.NoError(t, err)

// 	var clusterId string

// 	// Create a cluster with unreasonably low timeout
// 	_, err = ClustersAPI.CreateAndWait(ctx, compute.CreateCluster{
// 		ClusterName:            RandomName(t.Name()),
// 		SparkVersion:           latest,
// 		InstancePoolId:         GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),
// 		AutoterminationMinutes: 10,
// 		NumWorkers:             1,
// 	}, retries.Timeout[compute.ClusterDetails](15*time.Second),
// 		func(i *retries.Info[compute.ClusterDetails]) {
// 			if i.Info == nil {
// 				return
// 			}
// 			clusterId = i.Info.ClusterId
// 		})
// 	assert.True(t, strings.HasPrefix(err.Error(), "timed out: "))
// 	_, err = ClustersAPI.DeleteByClusterIdAndWait(ctx, clusterId)
// 	require.NoError(t, err)
// }

func TestAccClustersGetCorrectErrorMessageNoTranspile(t *testing.T) {
	ctx := workspaceTest(t)

	ClustersAPI, err := compute.NewClustersClient(nil)
	require.NoError(t, err)
	_, err = ClustersAPI.GetByClusterId(ctx, "123456789")
	assert.ErrorIs(t, err, databricks.ErrResourceDoesNotExist)
}

func TestAccAwsInstanceProfiles(t *testing.T) {
	ctx := workspaceTest(t)

	cfg := &config.Config{}
	InstanceProfilesAPI, err := compute.NewInstanceProfilesClient(cfg)
	require.NoError(t, err)
	if !cfg.IsAws() {
		t.Skipf("runs only on AWS")
	}

	arn := "arn:aws:iam::000000000000:instance-profile/abc"
	err = InstanceProfilesAPI.Add(ctx, compute.AddInstanceProfile{
		InstanceProfileArn: arn,
		SkipValidation:     true,
		IamRoleArn:         "arn:aws:iam::000000000000:role/bcd",
	})
	require.NoError(t, err)

	defer InstanceProfilesAPI.RemoveByInstanceProfileArn(ctx, arn)

	err = InstanceProfilesAPI.Edit(ctx, compute.InstanceProfile{
		InstanceProfileArn: arn,
		IamRoleArn:         "arn:aws:iam::000000000000:role/bcdf",
	})
	require.NoError(t, err)

	all, err := InstanceProfilesAPI.ListAll(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

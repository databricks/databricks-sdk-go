package internal

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func sharedRunningCluster(t *testing.T, ctx context.Context,
	w *databricks.WorkspaceClient,
) string {
	clusterId := GetEnvOrSkipTest(t, "TEST_GO_SDK_CLUSTER_ID")
	err := w.Clusters.EnsureClusterIsRunning(ctx, clusterId)
	require.NoError(t, err)
	return clusterId
}

func TestAccClustersCreateFailsWithTimeoutNoTranspile(t *testing.T) {
	ctx, w := workspaceTest(t)

	// Fetch list of spark runtime versions
	sparkVersions, err := w.Clusters.SparkVersions(ctx)
	require.NoError(t, err)

	// Select the latest LTS version without Photon
	latest, err := sparkVersions.Select(compute.SparkVersionRequest{
		Latest:          true,
		LongTermSupport: true,
	})
	require.NoError(t, err)

	var clusterId string

	// Create a cluster with unreasonably low timeout
	_, err = w.Clusters.CreateAndWait(ctx, compute.CreateCluster{
		ClusterName:            RandomName(t.Name()),
		SparkVersion:           latest,
		InstancePoolId:         GetEnvOrSkipTest(t, "TEST_INSTANCE_POOL_ID"),
		AutoterminationMinutes: 10,
		NumWorkers:             1,
	}, retries.Timeout[compute.ClusterDetails](15*time.Second),
		func(i *retries.Info[compute.ClusterDetails]) {
			if i.Info == nil {
				return
			}
			clusterId = i.Info.ClusterId
		})
	assert.True(t, strings.HasPrefix(err.Error(), "timed out: "))
	_, err = w.Clusters.DeleteByClusterIdAndWait(ctx, clusterId)
	require.NoError(t, err)
}

func TestAccClustersGetCorrectErrorMessageNoTranspile(t *testing.T) {
	ctx, w := workspaceTest(t)
	_, err := w.Clusters.GetByClusterId(ctx, "123456789")
	assert.ErrorIs(t, err, databricks.ErrResourceDoesNotExist)
}

func TestAccAwsInstanceProfiles(t *testing.T) {
	ctx, w := workspaceTest(t)
	if !IsCloud(environment.CloudAWS) {
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

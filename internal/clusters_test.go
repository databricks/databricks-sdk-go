package internal

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClustersCreateTimesOutAndCleansUp(t *testing.T) {
	const clusterID = "test-cluster"
	createRequest := compute.CreateCluster{
		ClusterName:            "test-cluster",
		SparkVersion:           "15.4.x-scala2.12",
		InstancePoolId:         "test-pool",
		AutoterminationMinutes: 10,
		NumWorkers:             1,
	}
	pending := compute.ClusterDetails{
		ClusterId:    clusterID,
		State:        compute.StatePending,
		StateMessage: "still provisioning",
	}

	qa.HTTPFixtures{
		{
			Method:          http.MethodPost,
			Resource:        "/api/2.1/clusters/create",
			ExpectedRequest: createRequest,
			Response: compute.CreateClusterResponse{
				ClusterId: clusterID,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.1/clusters/get?cluster_id=" + clusterID,
			Response: pending,
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.1/clusters/get?cluster_id=" + clusterID,
			Response: pending,
		},
		{
			Method:   http.MethodPost,
			Resource: "/api/2.1/clusters/delete",
			ExpectedRequest: compute.DeleteCluster{
				ClusterId: clusterID,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.1/clusters/get?cluster_id=" + clusterID,
			Response: compute.ClusterDetails{
				ClusterId: clusterID,
				State:     compute.StateTerminated,
			},
		},
	}.ApplyClient(t, func(ctx context.Context, apiClient *client.DatabricksClient) {
		clusters := compute.NewClusters(apiClient)
		createCtx, cancelCreate := context.WithCancel(ctx)
		defer cancelCreate()

		var gotStates []compute.State
		_, err := clusters.CreateAndWait(
			createCtx,
			createRequest,
			retries.Timeout[compute.ClusterDetails](time.Minute),
			retries.OnPoll(func(details *compute.ClusterDetails) {
				gotStates = append(gotStates, details.State)
				if len(gotStates) == 2 {
					cancelCreate()
				}
			}),
		)
		assert.EqualError(t, err, "timed out: still provisioning")
		assert.Equal(t, []compute.State{compute.StatePending, compute.StatePending}, gotStates)

		deleted, err := clusters.DeleteByClusterIdAndWait(ctx, clusterID)
		require.NoError(t, err)
		assert.Equal(t, clusterID, deleted.ClusterId)
		assert.Equal(t, compute.StateTerminated, deleted.State)
	})
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

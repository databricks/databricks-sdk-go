package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/compute/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccInstancePools(t *testing.T) {
	ctx := workspaceTest(t)

	ClustersAPI, err := compute.NewClustersClient(nil)
	require.NoError(t, err)
	smallest, err := ClustersAPI.SelectNodeType(ctx, compute.NodeTypeRequest{
		LocalDisk: true,
	})
	require.NoError(t, err)

	InstancePoolsAPI, err := compute.NewInstancePoolsClient(nil)
	require.NoError(t, err)
	created, err := InstancePoolsAPI.Create(ctx, compute.CreateInstancePool{
		InstancePoolName: RandomName("go-sdk-"),
		NodeTypeId:       smallest,
	})
	require.NoError(t, err)

	defer InstancePoolsAPI.DeleteByInstancePoolId(ctx, created.InstancePoolId)

	_, err = InstancePoolsAPI.Edit(ctx, compute.EditInstancePool{
		InstancePoolId:   created.InstancePoolId,
		InstancePoolName: RandomName("go-sdk-updated"),
		NodeTypeId:       smallest,
	})
	require.NoError(t, err)

	byId, err := InstancePoolsAPI.GetByInstancePoolId(ctx, created.InstancePoolId)
	require.NoError(t, err)

	all, err := InstancePoolsAPI.ListAll(ctx)
	require.NoError(t, err)

	names, err := InstancePoolsAPI.InstancePoolAndStatsInstancePoolNameToInstancePoolIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, byId.InstancePoolId, names[byId.InstancePoolName])

	byName, err := InstancePoolsAPI.GetByInstancePoolName(ctx, byId.InstancePoolName)
	require.NoError(t, err)
	assert.Equal(t, byName.InstancePoolId, byId.InstancePoolId)
}

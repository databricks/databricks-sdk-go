package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/instancepools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccInstancePools(t *testing.T) {
	ctx, w := workspaceTest(t)

	vms, err := w.Clusters.ListNodeTypes(ctx)
	require.NoError(t, err)

	smallest, err := vms.Smallest(clusters.NodeTypeRequest{
		LocalDisk: true,
	})
	require.NoError(t, err)

	created, err := w.InstancePools.Create(ctx, instancepools.CreateInstancePool{
		InstancePoolName: RandomName("go-sdk-"),
		NodeTypeId:       smallest,
	})
	require.NoError(t, err)

	defer w.InstancePools.DeleteByInstancePoolId(ctx, created.InstancePoolId)

	err = w.InstancePools.Edit(ctx, instancepools.EditInstancePool{
		InstancePoolId:   created.InstancePoolId,
		InstancePoolName: RandomName("go-sdk-updated"),
		NodeTypeId:       smallest,
	})
	require.NoError(t, err)

	byId, err := w.InstancePools.GetByInstancePoolId(ctx, created.InstancePoolId)
	require.NoError(t, err)

	all, err := w.InstancePools.ListAll(ctx)
	require.NoError(t, err)

	names, err := w.InstancePools.InstancePoolAndStatsInstancePoolNameToInstancePoolIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, byId.InstancePoolId, names[byId.InstancePoolName])

	byName, err := w.InstancePools.GetByInstancePoolName(ctx, byId.InstancePoolName)
	require.NoError(t, err)
	assert.Equal(t, byName.InstancePoolId, byId.InstancePoolId)
}

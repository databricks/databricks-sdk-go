package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/assert"
)

func TestAccInstancePools(t *testing.T) {
	ctx, w := workspaceTest(t)

	smallest, err := w.Clusters.SelectNodeType(ctx, compute.NodeTypeRequest{
		LocalDisk: true,
	})
	assert.NoError(t, err)

	created, err := w.InstancePools.Create(ctx, compute.CreateInstancePool{
		InstancePoolName: RandomName("go-sdk-"),
		NodeTypeId:       smallest,
	})
	assert.NoError(t, err)

	defer w.InstancePools.DeleteByInstancePoolId(ctx, created.InstancePoolId)

	err = w.InstancePools.Edit(ctx, compute.EditInstancePool{
		InstancePoolId:   created.InstancePoolId,
		InstancePoolName: RandomName("go-sdk-updated"),
		NodeTypeId:       smallest,
	})
	assert.NoError(t, err)

	byId, err := w.InstancePools.GetByInstancePoolId(ctx, created.InstancePoolId)
	assert.NoError(t, err)

	all, err := w.InstancePools.ListAll(ctx)
	assert.NoError(t, err)

	names, err := w.InstancePools.InstancePoolAndStatsInstancePoolNameToInstancePoolIdMap(ctx)
	assert.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, byId.InstancePoolId, names[byId.InstancePoolName])

	byName, err := w.InstancePools.GetByInstancePoolName(ctx, byId.InstancePoolName)
	assert.NoError(t, err)
	assert.Equal(t, byName.InstancePoolId, byId.InstancePoolId)
}

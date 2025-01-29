package integration

// TODO: Enable this test when LRO is implemented
//
// func TestAccInstancePools(t *testing.T) {
// 	ctx := workspaceTest(t)

// 	ClustersAPI, err := compute.NewClustersClient(nil)
// 	smallest, err := ClustersAPI.SelectNodeType(ctx, compute.NodeTypeRequest{
// 		LocalDisk: true,
// 	})
// 	require.NoError(t, err)

// 	created, err := w.InstancePools.Create(ctx, compute.CreateInstancePool{
// 		InstancePoolName: RandomName("go-sdk-"),
// 		NodeTypeId:       smallest,
// 	})
// 	require.NoError(t, err)

// 	defer w.InstancePools.DeleteByInstancePoolId(ctx, created.InstancePoolId)

// 	err = w.InstancePools.Edit(ctx, compute.EditInstancePool{
// 		InstancePoolId:   created.InstancePoolId,
// 		InstancePoolName: RandomName("go-sdk-updated"),
// 		NodeTypeId:       smallest,
// 	})
// 	require.NoError(t, err)

// 	byId, err := w.InstancePools.GetByInstancePoolId(ctx, created.InstancePoolId)
// 	require.NoError(t, err)

// 	all, err := w.InstancePools.ListAll(ctx)
// 	require.NoError(t, err)

// 	names, err := w.InstancePools.InstancePoolAndStatsInstancePoolNameToInstancePoolIdMap(ctx)
// 	require.NoError(t, err)
// 	assert.Equal(t, len(all), len(names))
// 	assert.Equal(t, byId.InstancePoolId, names[byId.InstancePoolName])

// 	byName, err := w.InstancePools.GetByInstancePoolName(ctx, byId.InstancePoolName)
// 	require.NoError(t, err)
// 	assert.Equal(t, byName.InstancePoolId, byId.InstancePoolId)
// }

package integration

// TODO: Enable this test when LRO is implemented
//
// func TestAccLibraries(t *testing.T) {
// 	ctx, w := workspaceTest(t)
// 	clusterId := sharedRunningCluster(t, ctx, w)

// 	err := w.Libraries.UpdateAndWait(ctx, compute.Update{
// 		ClusterId: clusterId,
// 		Install: []compute.Library{
// 			{
// 				Pypi: &compute.PythonPyPiLibrary{
// 					Package: "dbl-tempo",
// 				},
// 			},
// 		},
// 	})
// 	require.NoError(t, err)

// 	err = w.Libraries.UpdateAndWait(ctx, compute.Update{
// 		ClusterId: clusterId,
// 		Uninstall: []compute.Library{
// 			{
// 				Pypi: &compute.PythonPyPiLibrary{
// 					Package: "dbl-tempo",
// 				},
// 			},
// 		},
// 	})
// 	require.NoError(t, err)
// }

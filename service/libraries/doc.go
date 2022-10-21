// Cluster Libraries API
//
// Libraries can only be installed on a running cluster. Please use the
// high-level [LibrariesAPI.Update] instead of direct calls
// to [LibrariesAPI.Install], [LibrariesAPI.Uninstall],
// and [LibrariesAPI.ClusterStatus].
//
//	err := w.Libraries.UpdateAndWait(ctx, libraries.Update{
//		ClusterId: clusterId,
//		Install: []libraries.Library{
//			{
//				Pypi: &libraries.PythonPyPiLibrary{
//					Package: "dbl-tempo",
//				},
//			},
//		},
//	})
package libraries

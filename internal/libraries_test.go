package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/stretchr/testify/require"
)

func TestAccLibraries(t *testing.T) {
	ctx, w := workspaceTest(t)
	clusterId := sharedRunningClusterNoTranspile(t, ctx, w)

	err := w.Libraries.UpdateAndWait(ctx, compute.Update{
		ClusterId: clusterId,
		Install: []compute.Library{
			{
				Pypi: &compute.PythonPyPiLibrary{
					Package: "dbl-tempo",
				},
			},
		},
	})
	require.NoError(t, err)

	err = w.Libraries.UpdateAndWait(ctx, compute.Update{
		ClusterId: clusterId,
		Uninstall: []compute.Library{
			{
				Pypi: &compute.PythonPyPiLibrary{
					Package: "dbl-tempo",
				},
			},
		},
	})
	require.NoError(t, err)
}

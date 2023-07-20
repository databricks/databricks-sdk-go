package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xuxiaoshuo/databricks-sdk-go/service/compute"
)

func TestAccLibraries(t *testing.T) {
	ctx, w := workspaceTest(t)
	clusterId := sharedRunningCluster(t, ctx, w)

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

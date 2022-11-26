package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/libraries"
	"github.com/stretchr/testify/require"
)

func TestAccLibraries(t *testing.T) {
	ctx, w := workspaceTest(t)
	clusterId := sharedRunningCluster(t, ctx, w)

	err := w.Libraries.UpdateAndWait(ctx, libraries.Update{
		ClusterId: clusterId,
		Install: []libraries.Library{
			{
				Pypi: &libraries.PythonPyPiLibrary{
					Package: "dbl-tempo",
				},
			},
		},
	})
	require.NoError(t, err)
}

package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/libraries"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
)

func TestAccLibraries(t *testing.T) {
	GetEnvOrSkipTest(t, "CLOUD_ENV")
	ctx := context.Background()
	w := workspaces.MustNewClient()
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

	clusterId := createTestCluster(ctx, w, t)
	defer w.Clusters.PermanentDeleteByClusterId(ctx, clusterId)

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
	assert.NoError(t, err)
}

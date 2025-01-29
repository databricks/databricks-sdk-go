package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/catalog/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUcAccMetastoreAssignments(t *testing.T) {
	ctx, cfg := ucacctTest(t)
	AccountMetastoreAssignmentsAPI, err := catalog.NewAccountMetastoreAssignmentsClient(cfg)
	require.NoError(t, err)
	ws, err := AccountMetastoreAssignmentsAPI.ListAll(ctx, catalog.ListAccountMetastoreAssignmentsRequest{
		MetastoreId: GetEnvOrSkipTest(t, "TEST_METASTORE_ID"),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, ws)
}

package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUcAccMetastoreAssignments(t *testing.T) {
	ctx, a := ucacctTest(t)
	ws, err := a.MetastoreAssignments.ListAll(ctx, catalog.ListAccountMetastoreAssignmentsRequest{
		MetastoreId: GetEnvOrSkipTest(t, "TEST_METASTORE_ID"),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, ws)
}

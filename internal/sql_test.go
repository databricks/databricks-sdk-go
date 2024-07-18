package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/stretchr/testify/require"
)

func TestAccSqlQueryHistory(t *testing.T) {
	ctx, w := workspaceTest(t)
	_, err := w.QueryHistory.List(ctx, sql.ListQueryHistoryRequest{
		FilterBy: &sql.QueryFilter{
			QueryStartTimeRange: &sql.TimeRange{
				StartTimeMs: 1690243200000, // 25-07-2023
				EndTimeMs:   1690329600000, // 26-07-2023
			},
		},
	})
	require.NoError(t, err)
}

package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/stretchr/testify/assert"
)

func TestAccSqlQueryHistory(t *testing.T) {
	ctx, w := workspaceTest(t)
	_, err := w.QueryHistory.ListAll(ctx, sql.ListQueryHistoryRequest{
		FilterBy: &sql.QueryFilter{
			QueryStartTimeRange: &sql.TimeRange{
				StartTimeMs: 1690243200000, // 25-07-2023
				EndTimeMs:   1690329600000, // 26-07-2023
			},
		},
	})
	assert.NoError(t, err)
}

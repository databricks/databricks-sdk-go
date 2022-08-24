package jobs

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/qa"
)

func TestCancelAllRuns(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "POST",
			Resource: "/api/2.1/jobs/runs/cancel-all",
			ExpectedRequest: CancelAllRunsRequest{
				JobId: 123,
			},
			Status: 200,
		},
	}.Apply(t, func(ctx context.Context, cfg *databricks.Config) {
		jobs := &JobsAPI{
			client: client.New(cfg),
		}
		jobs.CancelAllRuns(ctx, CancelAllRunsRequest{
			JobId: 123,
		})
	})
}

package compute

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/stretchr/testify/assert"
)

func TestQueryValues(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.1/clusters/list?filter_by.is_pinned=true",
			Status:   200,
			Response: ListClustersResponse{},
		},
	}.ApplyClient(t, func(ctx context.Context, client *client.DatabricksClient) {
		clusters := NewClusters(client)
		it := clusters.List(ctx, ListClustersRequest{
			FilterBy: &ListClustersFilterBy{
				IsPinned: true,
			},
		})

		assert.False(t, it.HasNext(ctx))
	})
}

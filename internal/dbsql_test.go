package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/dbsql"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccQueries(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	w := workspaces.MustNewClient()
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

	srcs, err := w.DataSources.ListDataSources(ctx)
	require.NoError(t, err)
	if len(srcs) == 0 {
		t.Skipf("no sql warehouses found")
	}

	query, err := w.Queries.CreateQuery(ctx, dbsql.QueryPostContent{
		Name:         RandomName("go-sdk/test/"),
		DataSourceId: srcs[0].Id,
		Description:  "test query from Go SDK",
		Query:        "SHOW TABLES",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := w.Queries.DeleteQueryByQueryId(ctx, query.Id)
		require.NoError(t, err)
	})

	loaded, err := w.Queries.GetQueryByQueryId(ctx, query.Id)
	require.NoError(t, err)
	assert.Equal(t, query.Query, loaded.Query)
}

func TestAccDashboards(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	w := workspaces.MustNewClient()
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

	all, err := w.Dashboards.ListDashboardsAll(ctx, dbsql.ListDashboardsRequest{})
	require.NoError(t, err)
	t.Log(len(all))
}

func TestAccQueriesList(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	w := workspaces.MustNewClient()
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

	srcs, err := w.DataSources.ListDataSources(ctx)
	require.NoError(t, err)
	if len(srcs) == 0 {
		t.Skipf("no sql warehouses found")
	}
	for i := 0; i < 34; i++ {
		query, err := w.Queries.CreateQuery(ctx, dbsql.QueryPostContent{
			Name:         RandomName("go-sdk/test/"),
			DataSourceId: srcs[0].Id,
			Description:  "test query from Go SDK",
			Query:        "SHOW TABLES",
		})
		require.NoError(t, err)
		t.Cleanup(func() {
			err := w.Queries.DeleteQueryByQueryId(ctx, query.Id)
			require.NoError(t, err)
		})
	}
	var qs1, qs2, qs3 []dbsql.Query
	{
		result, err := w.Queries.Impl().ListQueries(ctx, dbsql.ListQueriesRequest{PageSize: 10})
		require.NoError(t, err)
		qs1 = result.Results
	}
	{
		result, err := w.Queries.ListQueriesAll(ctx, dbsql.ListQueriesRequest{})
		require.NoError(t, err)
		qs2 = result
	}
	{
		result, err := w.Queries.ListQueriesAll(ctx, dbsql.ListQueriesRequest{PageSize: 10})
		require.NoError(t, err)
		qs3 = result
	}
	for i := 0; i < len(qs1) && i < len(qs2); i++ {
		assert.Equal(t, qs1[i], qs2[i], "Query at index %d not equal", i)
		assert.Equal(t, qs1[i], qs3[i], "Query at index %d not equal", i)
	}
}

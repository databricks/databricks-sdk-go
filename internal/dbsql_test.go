package internal

import (
	"context"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/dbsql"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccQueries(t *testing.T) {
	// env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	// t.Log(env)
	os.Setenv("DATABRICKS_CONFIG_PROFILE", "azure-deco")
	os.Setenv("DATABRICKS_DEBUG_TRUNCATE_BYTES", "5096")
	ctx := context.Background()
	wsc := workspaces.New()

	srcs, err := wsc.DataSources.ListDataSources(ctx)
	require.NoError(t, err)
	if len(srcs) == 0 {
		t.Skipf("no sql warehouses found")
	}

	query, err := wsc.Queries.CreateQuery(ctx, dbsql.QueryPostContent{
		Name:         RandomName("go-sdk/test/"),
		DataSourceId: srcs[0].Id,
		Description:  "test query from Go SDK",
		Query:        "SHOW TABLES",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := wsc.Queries.DeleteQueryByQueryId(ctx, query.Id)
		require.NoError(t, err)
	})

	loaded, err := wsc.Queries.GetQueryByQueryId(ctx, query.Id)
	require.NoError(t, err)
	assert.Equal(t, query.Query, loaded.Query)

	nameToId, err := wsc.Queries.QueryNameToIdMap(ctx, dbsql.ListQueriesRequest{})
	require.NoError(t, err)
	assert.Equal(t, nameToId[loaded.Name], loaded.Id)

	byName, err := wsc.Queries.GetQueryByName(ctx, loaded.Name)
	require.NoError(t, err)
	assert.Equal(t, byName.Id, loaded.Id)
}

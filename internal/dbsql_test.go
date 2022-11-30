package internal

import (
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/dbsql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccQueries(t *testing.T) {
	ctx, w := workspaceTest(t)

	srcs, err := w.DataSources.ListDataSources(ctx)
	require.NoError(t, err)
	if len(srcs) == 0 {
		t.Skipf("no sql warehouses found")
	}

	// TODO: OpenAPI: CRUD names for operationId
	query, err := w.Queries.CreateQuery(ctx, dbsql.QueryPostContent{
		Name:         RandomName("go-sdk/test/"),
		DataSourceId: srcs[0].Id,
		Description:  "test query from Go SDK",
		Query:        "SHOW TABLES",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Queries.DeleteQueryByQueryId(ctx, query.Id)
		require.NoError(t, err)
	})

	byId, err := w.Queries.GetQueryByQueryId(ctx, query.Id)
	require.NoError(t, err)
	assert.Equal(t, query.Query, byId.Query)

	byName, err := w.Queries.GetByName(ctx, byId.Name)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	updated, err := w.Queries.UpdateQuery(ctx, dbsql.QueryPostContent{
		QueryId:      query.Id,
		Name:         RandomName("go-sdk-updated"),
		DataSourceId: srcs[0].Id,
		Description:  "UPDATED: test query from Go SDK",
		Query:        "SELECT 2+2",
	})
	require.NoError(t, err)
	assert.NotEqual(t, updated.Query, byId.Query)
}

func TestAccAlerts(t *testing.T) {
	ctx, w := workspaceTest(t)

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
	defer w.Queries.DeleteQueryByQueryId(ctx, query.Id)

	// TODO: OpenAPI: CRUD names for operationId
	alert, err := w.Alerts.CreateAlert(ctx, dbsql.EditAlert{
		Name:    RandomName("go-sdk-"),
		QueryId: query.Id,
	})
	require.NoError(t, err)
	defer w.Alerts.DeleteAlertByAlertId(ctx, alert.Id)

	err = w.Alerts.UpdateAlert(ctx, dbsql.EditAlert{
		AlertId: alert.Id,
		Name:    RandomName("go-sdk-updated-"),
		QueryId: query.Id,
	})
	require.NoError(t, err)

	byId, err := w.Alerts.GetAlertByAlertId(ctx, alert.Id)
	require.NoError(t, err)

	byName, err := w.Alerts.GetByName(ctx, byId.Name)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	all, err := w.Alerts.ListAlerts(ctx)
	require.NoError(t, err)

	names, err := w.Alerts.AlertNameToIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, alert.Id, names[byId.Name])

	schedule, err := w.Alerts.CreateSchedule(ctx, dbsql.CreateRefreshSchedule{
		AlertId:      alert.Id,
		Cron:         "5 4 * * *",
		DataSourceId: srcs[0].Id,
	})
	require.NoError(t, err)
	defer w.Alerts.DeleteScheduleByAlertIdAndScheduleId(ctx, alert.Id, schedule.Id)

	schedules, err := w.Alerts.ListSchedulesByAlertId(ctx, alert.Id)
	require.NoError(t, err)
	assert.True(t, len(schedules) >= 1)

	me, err := w.CurrentUser.Me(ctx)
	require.NoError(t, err)

	userId, err := strconv.ParseInt(me.Id, 10, 64)
	require.NoError(t, err)

	sub, err := w.Alerts.Subscribe(ctx, dbsql.CreateSubscription{
		AlertId: alert.Id,
		UserId:  userId,
	})
	require.NoError(t, err)

	allSubs, err := w.Alerts.GetSubscriptionsByAlertId(ctx, alert.Id)
	require.NoError(t, err)
	assert.True(t, len(allSubs) >= 1)

	err = w.Alerts.UnsubscribeByAlertIdAndSubscriptionId(ctx, alert.Id, sub.Id)
	require.NoError(t, err)
}

func TestAccDashboards(t *testing.T) {
	ctx, w := workspaceTest(t)

	// TODO: OpenAPI: CRUD names for operationId
	created, err := w.Dashboards.CreateDashboard(ctx, dbsql.CreateDashboardRequest{
		Name:                    RandomName("go-sdk-"),
		DashboardFiltersEnabled: false,
		IsDraft:                 true,
	})
	require.NoError(t, err)

	defer w.Dashboards.DeleteDashboardByDashboardId(ctx, created.Id)

	byId, err := w.Dashboards.GetDashboardByDashboardId(ctx, created.Id)
	require.NoError(t, err)

	byName, err := w.Dashboards.GetByName(ctx, byId.Name)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	all, err := w.Dashboards.ListDashboardsAll(ctx, dbsql.ListDashboardsRequest{})
	require.NoError(t, err)

	names, err := w.Dashboards.DashboardNameToIdMap(ctx, dbsql.ListDashboardsRequest{})
	require.NoError(t, err)
	assert.Equal(t, created.Id, names[byId.Name])
	assert.Equal(t, len(all), len(names))

	err = w.Dashboards.DeleteDashboardByDashboardId(ctx, created.Id)
	require.NoError(t, err)

	err = w.Dashboards.RestoreDashboard(ctx, dbsql.RestoreDashboardRequest{
		DashboardId: created.Id,
	})
	require.NoError(t, err)
}

func TestAccQueriesList(t *testing.T) {
	ctx, w := workspaceTest(t)

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

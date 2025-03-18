package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/sql/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccQueries(t *testing.T) {
	ctx := workspaceTest(t)

	DataSourcesAPI, err := sql.NewDataSourcesClient(nil)
	require.NoError(t, err)
	srcs, err := DataSourcesAPI.List(ctx)
	require.NoError(t, err)
	if len(srcs) == 0 {
		t.Skipf("no sql warehouses found")
	}

	QueriesAPI, err := sql.NewQueriesClient(nil)
	require.NoError(t, err)
	query, err := QueriesAPI.Create(ctx, sql.CreateQueryRequest{
		Query: &sql.CreateQueryRequestQuery{
			DisplayName: RandomName("go-sdk-test-"),
			WarehouseId: srcs[0].WarehouseId,
			Description: "test query from Go SDK",
			QueryText:   "SHOW TABLES",
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err = QueriesAPI.DeleteById(ctx, query.Id)
		require.NoError(t, err)
	})

	byId, err := QueriesAPI.GetById(ctx, query.Id)
	require.NoError(t, err)
	assert.Equal(t, query.QueryText, byId.QueryText)

	byName, err := QueriesAPI.GetByDisplayName(ctx, byId.DisplayName)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	updated, err := QueriesAPI.Update(ctx, sql.UpdateQueryRequest{
		Id: query.Id,
		Query: &sql.UpdateQueryRequestQuery{
			DisplayName: RandomName("go-sdk-updated"),
			Description: "UPDATED: test query from Go SDK",
			QueryText:   "SELECT 2+2",
		},
		UpdateMask: "display_name,description,query_text",
	})
	require.NoError(t, err)
	assert.NotEqual(t, updated.QueryText, byId.QueryText)
}

func TestAccAlerts(t *testing.T) {
	ctx := workspaceTest(t)

	DataSourcesAPI, err := sql.NewDataSourcesClient(nil)
	require.NoError(t, err)
	srcs, err := DataSourcesAPI.List(ctx)
	require.NoError(t, err)
	if len(srcs) == 0 {
		t.Skipf("no sql warehouses found")
	}

	QueriesAPI, err := sql.NewQueriesClient(nil)
	require.NoError(t, err)
	query, err := QueriesAPI.Create(ctx, sql.CreateQueryRequest{
		Query: &sql.CreateQueryRequestQuery{
			DisplayName: RandomName("go-sdk-test-"),
			WarehouseId: srcs[0].WarehouseId,
			Description: "test query from Go SDK",
			QueryText:   "SELECT 1",
		},
	})
	require.NoError(t, err)
	defer QueriesAPI.DeleteById(ctx, query.Id)

	AlertsAPI, err := sql.NewAlertsClient(nil)
	require.NoError(t, err)
	alert, err := AlertsAPI.Create(ctx, sql.CreateAlertRequest{
		Alert: &sql.CreateAlertRequestAlert{
			Condition: &sql.AlertCondition{
				Operand: &sql.AlertConditionOperand{
					Column: &sql.AlertOperandColumn{
						Name: "1",
					},
				},
				Op: sql.AlertOperatorEqual,
				Threshold: &sql.AlertConditionThreshold{
					Value: &sql.AlertOperandValue{
						DoubleValue: 1,
					},
				},
			},
			DisplayName: RandomName("go-sdk-"),
			QueryId:     query.Id,
		},
	})
	require.NoError(t, err)
	defer AlertsAPI.DeleteById(ctx, alert.Id)

	_, err = AlertsAPI.Update(ctx, sql.UpdateAlertRequest{
		Id: alert.Id,
		Alert: &sql.UpdateAlertRequestAlert{
			DisplayName: RandomName("go-sdk-updated"),
		},
		UpdateMask: "display_name",
	})
	require.NoError(t, err)

	byId, err := AlertsAPI.GetById(ctx, alert.Id)
	require.NoError(t, err)

	byName, err := AlertsAPI.GetByDisplayName(ctx, byId.DisplayName)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	all, err := AlertsAPI.ListAll(ctx, sql.ListAlertsRequest{})
	require.NoError(t, err)

	names, err := AlertsAPI.ListAlertsResponseAlertDisplayNameToIdMap(ctx, sql.ListAlertsRequest{})
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, alert.Id, names[byId.DisplayName])
}

func TestAccDashboards(t *testing.T) {
	ctx := workspaceTest(t)

	DashboardsAPI, err := sql.NewDashboardsClient(nil)
	require.NoError(t, err)
	created, err := DashboardsAPI.Create(ctx, sql.DashboardPostContent{
		Name: RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	defer DashboardsAPI.DeleteByDashboardId(ctx, created.Id)

	byId, err := DashboardsAPI.GetByDashboardId(ctx, created.Id)
	require.NoError(t, err)

	byName, err := DashboardsAPI.GetByName(ctx, byId.Name)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	all, err := DashboardsAPI.ListAll(ctx, sql.ListDashboardsRequest{})
	require.NoError(t, err)

	names, err := DashboardsAPI.DashboardNameToIdMap(ctx, sql.ListDashboardsRequest{})
	require.NoError(t, err)
	assert.Equal(t, created.Id, names[byId.Name])
	assert.Equal(t, len(all), len(names))

	_, err = DashboardsAPI.DeleteByDashboardId(ctx, created.Id)
	require.NoError(t, err)

	_, err = DashboardsAPI.Restore(ctx, sql.RestoreDashboardRequest{
		DashboardId: created.Id,
	})
	require.NoError(t, err)
}

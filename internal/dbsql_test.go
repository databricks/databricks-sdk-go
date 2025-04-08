package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccQueries(t *testing.T) {
	ctx, w := workspaceTest(t)

	srcs, err := w.DataSources.List(ctx)
	require.NoError(t, err)
	if len(srcs) == 0 {
		t.Skipf("no sql warehouses found")
	}

	query, err := w.Queries.Create(ctx, sql.CreateQueryRequest{
		Query: &sql.CreateQueryRequestQuery{
			DisplayName: RandomName("go-sdk-test-"),
			WarehouseId: srcs[0].WarehouseId,
			Description: "test query from Go SDK",
			QueryText:   "SHOW TABLES",
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Queries.DeleteById(ctx, query.Id)
		require.NoError(t, err)
	})

	byId, err := w.Queries.GetById(ctx, query.Id)
	require.NoError(t, err)
	assert.Equal(t, query.QueryText, byId.QueryText)

	byName, err := w.Queries.GetByDisplayName(ctx, byId.DisplayName)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	updated, err := w.Queries.Update(ctx, sql.UpdateQueryRequest{
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
	ctx, w := workspaceTest(t)

	srcs, err := w.DataSources.List(ctx)
	require.NoError(t, err)
	if len(srcs) == 0 {
		t.Skipf("no sql warehouses found")
	}

	query, err := w.Queries.Create(ctx, sql.CreateQueryRequest{
		Query: &sql.CreateQueryRequestQuery{
			DisplayName: RandomName("go-sdk-test-"),
			WarehouseId: srcs[0].WarehouseId,
			Description: "test query from Go SDK",
			QueryText:   "SELECT 1",
		},
	})
	require.NoError(t, err)
	defer w.Queries.DeleteById(ctx, query.Id)

	alert, err := w.Alerts.Create(ctx, sql.CreateAlertRequest{
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
	defer w.Alerts.DeleteById(ctx, alert.Id)

	_, err = w.Alerts.Update(ctx, sql.UpdateAlertRequest{
		Id: alert.Id,
		Alert: &sql.UpdateAlertRequestAlert{
			DisplayName: RandomName("go-sdk-updated"),
		},
		UpdateMask: "display_name",
	})
	require.NoError(t, err)

	byId, err := w.Alerts.GetById(ctx, alert.Id)
	require.NoError(t, err)

	byName, err := w.Alerts.GetByDisplayName(ctx, byId.DisplayName)
	require.NoError(t, err)
	assert.Equal(t, byId.Id, byName.Id)

	all, err := w.Alerts.ListAll(ctx, sql.ListAlertsRequest{})
	require.NoError(t, err)

	names, err := w.Alerts.ListAlertsResponseAlertDisplayNameToIdMap(ctx, sql.ListAlertsRequest{})
	require.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, alert.Id, names[byId.DisplayName])
}

// As of April 7, 2025: Official support for the legacy version of dashboards has ended. More context: https://docs.databricks.com/gcp/en/sql/user/dashboards
// func TestAccDashboards(t *testing.T) {
// 	ctx, w := workspaceTest(t)
// 	// As of April 7, 2025: Official support for the legacy version of dashboards has ended.
// 	if w.Config.IsGcp() {
// 		t.Skip("Legacy dashboard creation is not supported on GCP")
// 	}
// 	created, err := w.Dashboards.Create(ctx, sql.DashboardPostContent{
// 		Name: RandomName("go-sdk-"),
// 	})
// 	require.NoError(t, err)
// 	defer w.Dashboards.DeleteByDashboardId(ctx, created.Id)
// 	byId, err := w.Dashboards.GetByDashboardId(ctx, created.Id)
// 	require.NoError(t, err)
// 	byName, err := w.Dashboards.GetByName(ctx, byId.Name)
// 	require.NoError(t, err)
// 	assert.Equal(t, byId.Id, byName.Id)
// 	all, err := w.Dashboards.ListAll(ctx, sql.ListDashboardsRequest{})
// 	require.NoError(t, err)
// 	names, err := w.Dashboards.DashboardNameToIdMap(ctx, sql.ListDashboardsRequest{})
// 	require.NoError(t, err)
// 	assert.Equal(t, created.Id, names[byId.Name])
// 	assert.Equal(t, len(all), len(names))
// 	err = w.Dashboards.DeleteByDashboardId(ctx, created.Id)
// 	require.NoError(t, err)
// 	err = w.Dashboards.Restore(ctx, sql.RestoreDashboardRequest{
// 		DashboardId: created.Id,
// 	})
// 	require.NoError(t, err)
// }

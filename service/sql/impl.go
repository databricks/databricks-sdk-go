// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Alerts API methods
type alertsImpl struct {
	client *client.DatabricksClient
}

func (a *alertsImpl) Create(ctx context.Context, request CreateAlertRequest) (*Alert, error) {
	var alert Alert
	path := "/api/2.0/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &alert)
	return &alert, err
}

func (a *alertsImpl) Delete(ctx context.Context, request TrashAlertRequest) error {
	var empty Empty
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &empty)
	return err
}

func (a *alertsImpl) Get(ctx context.Context, request GetAlertRequest) (*Alert, error) {
	var alert Alert
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &alert)
	return &alert, err
}

// List alerts.
//
// Gets a list of alerts accessible to the user, ordered by creation time.
// **Warning:** Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *alertsImpl) List(ctx context.Context, request ListAlertsRequest) listing.Iterator[ListAlertsResponseAlert] {

	getNextPage := func(ctx context.Context, req ListAlertsRequest) (*ListAlertsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAlertsResponse) []ListAlertsResponseAlert {
		return resp.Results
	}
	getNextReq := func(resp *ListAlertsResponse) *ListAlertsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List alerts.
//
// Gets a list of alerts accessible to the user, ordered by creation time.
// **Warning:** Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *alertsImpl) ListAll(ctx context.Context, request ListAlertsRequest) ([]ListAlertsResponseAlert, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ListAlertsResponseAlert](ctx, iterator)
}

func (a *alertsImpl) internalList(ctx context.Context, request ListAlertsRequest) (*ListAlertsResponse, error) {
	var listAlertsResponse ListAlertsResponse
	path := "/api/2.0/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAlertsResponse)
	return &listAlertsResponse, err
}

func (a *alertsImpl) Update(ctx context.Context, request UpdateAlertRequest) (*Alert, error) {
	var alert Alert
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &alert)
	return &alert, err
}

// unexported type that holds implementations of just AlertsLegacy API methods
type alertsLegacyImpl struct {
	client *client.DatabricksClient
}

func (a *alertsLegacyImpl) Create(ctx context.Context, request CreateAlert) (*LegacyAlert, error) {
	var legacyAlert LegacyAlert
	path := "/api/2.0/preview/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyAlert)
	return &legacyAlert, err
}

func (a *alertsLegacyImpl) Delete(ctx context.Context, request DeleteAlertsLegacyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *alertsLegacyImpl) Get(ctx context.Context, request GetAlertsLegacyRequest) (*LegacyAlert, error) {
	var legacyAlert LegacyAlert
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &legacyAlert)
	return &legacyAlert, err
}

func (a *alertsLegacyImpl) List(ctx context.Context) ([]LegacyAlert, error) {
	var legacyAlertList []LegacyAlert
	path := "/api/2.0/preview/sql/alerts"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &legacyAlertList)
	return legacyAlertList, err
}

func (a *alertsLegacyImpl) Update(ctx context.Context, request EditAlert) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just DashboardWidgets API methods
type dashboardWidgetsImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardWidgetsImpl) Create(ctx context.Context, request CreateWidget) (*Widget, error) {
	var widget Widget
	path := "/api/2.0/preview/sql/widgets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &widget)
	return &widget, err
}

func (a *dashboardWidgetsImpl) Delete(ctx context.Context, request DeleteDashboardWidgetRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/widgets/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *dashboardWidgetsImpl) Update(ctx context.Context, request CreateWidget) (*Widget, error) {
	var widget Widget
	path := fmt.Sprintf("/api/2.0/preview/sql/widgets/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &widget)
	return &widget, err
}

// unexported type that holds implementations of just Dashboards API methods
type dashboardsImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardsImpl) Create(ctx context.Context, request DashboardPostContent) (*Dashboard, error) {
	var dashboard Dashboard
	path := "/api/2.0/preview/sql/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &dashboard)
	return &dashboard, err
}

func (a *dashboardsImpl) Delete(ctx context.Context, request DeleteDashboardRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *dashboardsImpl) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &dashboard)
	return &dashboard, err
}

// Get dashboard objects.
//
// Fetch a paginated list of dashboard objects.
//
// **Warning**: Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *dashboardsImpl) List(ctx context.Context, request ListDashboardsRequest) listing.Iterator[Dashboard] {

	request.Page = 1 // start iterating from the first page

	getNextPage := func(ctx context.Context, req ListDashboardsRequest) (*ListResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListResponse) []Dashboard {
		return resp.Results
	}
	getNextReq := func(resp *ListResponse) *ListDashboardsRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.Page = resp.Page + 1
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[Dashboard, string](
		iterator,
		func(item Dashboard) string {
			return item.Id
		})
	return dedupedIterator
}

// Get dashboard objects.
//
// Fetch a paginated list of dashboard objects.
//
// **Warning**: Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *dashboardsImpl) ListAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Dashboard, int](ctx, iterator, request.PageSize)

}

func (a *dashboardsImpl) internalList(ctx context.Context, request ListDashboardsRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/preview/sql/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listResponse)
	return &listResponse, err
}

func (a *dashboardsImpl) Restore(ctx context.Context, request RestoreDashboardRequest) error {
	var restoreResponse RestoreResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/trash/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &restoreResponse)
	return err
}

func (a *dashboardsImpl) Update(ctx context.Context, request DashboardEditContent) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &dashboard)
	return &dashboard, err
}

// unexported type that holds implementations of just DataSources API methods
type dataSourcesImpl struct {
	client *client.DatabricksClient
}

func (a *dataSourcesImpl) List(ctx context.Context) ([]DataSource, error) {
	var dataSourceList []DataSource
	path := "/api/2.0/preview/sql/data_sources"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &dataSourceList)
	return dataSourceList, err
}

// unexported type that holds implementations of just DbsqlPermissions API methods
type dbsqlPermissionsImpl struct {
	client *client.DatabricksClient
}

func (a *dbsqlPermissionsImpl) Get(ctx context.Context, request GetDbsqlPermissionRequest) (*GetResponse, error) {
	var getResponse GetResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getResponse)
	return &getResponse, err
}

func (a *dbsqlPermissionsImpl) Set(ctx context.Context, request SetRequest) (*SetResponse, error) {
	var setResponse SetResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &setResponse)
	return &setResponse, err
}

func (a *dbsqlPermissionsImpl) TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error) {
	var success Success
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v/transfer", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &success)
	return &success, err
}

// unexported type that holds implementations of just Queries API methods
type queriesImpl struct {
	client *client.DatabricksClient
}

func (a *queriesImpl) Create(ctx context.Context, request CreateQueryRequest) (*Query, error) {
	var query Query
	path := "/api/2.0/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &query)
	return &query, err
}

func (a *queriesImpl) Delete(ctx context.Context, request TrashQueryRequest) error {
	var empty Empty
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &empty)
	return err
}

func (a *queriesImpl) Get(ctx context.Context, request GetQueryRequest) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &query)
	return &query, err
}

// List queries.
//
// Gets a list of queries accessible to the user, ordered by creation time.
// **Warning:** Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *queriesImpl) List(ctx context.Context, request ListQueriesRequest) listing.Iterator[ListQueryObjectsResponseQuery] {

	getNextPage := func(ctx context.Context, req ListQueriesRequest) (*ListQueryObjectsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListQueryObjectsResponse) []ListQueryObjectsResponseQuery {
		return resp.Results
	}
	getNextReq := func(resp *ListQueryObjectsResponse) *ListQueriesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List queries.
//
// Gets a list of queries accessible to the user, ordered by creation time.
// **Warning:** Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *queriesImpl) ListAll(ctx context.Context, request ListQueriesRequest) ([]ListQueryObjectsResponseQuery, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ListQueryObjectsResponseQuery](ctx, iterator)
}

func (a *queriesImpl) internalList(ctx context.Context, request ListQueriesRequest) (*ListQueryObjectsResponse, error) {
	var listQueryObjectsResponse ListQueryObjectsResponse
	path := "/api/2.0/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listQueryObjectsResponse)
	return &listQueryObjectsResponse, err
}

// List visualizations on a query.
//
// Gets a list of visualizations on a query.
func (a *queriesImpl) ListVisualizations(ctx context.Context, request ListVisualizationsForQueryRequest) listing.Iterator[Visualization] {

	getNextPage := func(ctx context.Context, req ListVisualizationsForQueryRequest) (*ListVisualizationsForQueryResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListVisualizations(ctx, req)
	}
	getItems := func(resp *ListVisualizationsForQueryResponse) []Visualization {
		return resp.Results
	}
	getNextReq := func(resp *ListVisualizationsForQueryResponse) *ListVisualizationsForQueryRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List visualizations on a query.
//
// Gets a list of visualizations on a query.
func (a *queriesImpl) ListVisualizationsAll(ctx context.Context, request ListVisualizationsForQueryRequest) ([]Visualization, error) {
	iterator := a.ListVisualizations(ctx, request)
	return listing.ToSlice[Visualization](ctx, iterator)
}

func (a *queriesImpl) internalListVisualizations(ctx context.Context, request ListVisualizationsForQueryRequest) (*ListVisualizationsForQueryResponse, error) {
	var listVisualizationsForQueryResponse ListVisualizationsForQueryResponse
	path := fmt.Sprintf("/api/2.0/sql/queries/%v/visualizations", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listVisualizationsForQueryResponse)
	return &listVisualizationsForQueryResponse, err
}

func (a *queriesImpl) Update(ctx context.Context, request UpdateQueryRequest) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &query)
	return &query, err
}

// unexported type that holds implementations of just QueriesLegacy API methods
type queriesLegacyImpl struct {
	client *client.DatabricksClient
}

func (a *queriesLegacyImpl) Create(ctx context.Context, request QueryPostContent) (*LegacyQuery, error) {
	var legacyQuery LegacyQuery
	path := "/api/2.0/preview/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyQuery)
	return &legacyQuery, err
}

func (a *queriesLegacyImpl) Delete(ctx context.Context, request DeleteQueriesLegacyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *queriesLegacyImpl) Get(ctx context.Context, request GetQueriesLegacyRequest) (*LegacyQuery, error) {
	var legacyQuery LegacyQuery
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &legacyQuery)
	return &legacyQuery, err
}

// Get a list of queries.
//
// Gets a list of queries. Optionally, this list can be filtered by a search
// term.
//
// **Warning**: Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
//
// **Note**: A new version of the Databricks SQL API is now available. Please
// use :method:queries/list instead. [Learn more]
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
func (a *queriesLegacyImpl) List(ctx context.Context, request ListQueriesLegacyRequest) listing.Iterator[LegacyQuery] {

	request.Page = 1 // start iterating from the first page

	getNextPage := func(ctx context.Context, req ListQueriesLegacyRequest) (*QueryList, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *QueryList) []LegacyQuery {
		return resp.Results
	}
	getNextReq := func(resp *QueryList) *ListQueriesLegacyRequest {
		if len(getItems(resp)) == 0 {
			return nil
		}
		request.Page = resp.Page + 1
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	dedupedIterator := listing.NewDedupeIterator[LegacyQuery, string](
		iterator,
		func(item LegacyQuery) string {
			return item.Id
		})
	return dedupedIterator
}

// Get a list of queries.
//
// Gets a list of queries. Optionally, this list can be filtered by a search
// term.
//
// **Warning**: Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
//
// **Note**: A new version of the Databricks SQL API is now available. Please
// use :method:queries/list instead. [Learn more]
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
func (a *queriesLegacyImpl) ListAll(ctx context.Context, request ListQueriesLegacyRequest) ([]LegacyQuery, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[LegacyQuery, int](ctx, iterator, request.PageSize)

}

func (a *queriesLegacyImpl) internalList(ctx context.Context, request ListQueriesLegacyRequest) (*QueryList, error) {
	var queryList QueryList
	path := "/api/2.0/preview/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &queryList)
	return &queryList, err
}

func (a *queriesLegacyImpl) Restore(ctx context.Context, request RestoreQueriesLegacyRequest) error {
	var restoreResponse RestoreResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/trash/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &restoreResponse)
	return err
}

func (a *queriesLegacyImpl) Update(ctx context.Context, request QueryEditContent) (*LegacyQuery, error) {
	var legacyQuery LegacyQuery
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyQuery)
	return &legacyQuery, err
}

// unexported type that holds implementations of just QueryHistory API methods
type queryHistoryImpl struct {
	client *client.DatabricksClient
}

func (a *queryHistoryImpl) List(ctx context.Context, request ListQueryHistoryRequest) (*ListQueriesResponse, error) {
	var listQueriesResponse ListQueriesResponse
	path := "/api/2.0/sql/history/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listQueriesResponse)
	return &listQueriesResponse, err
}

// unexported type that holds implementations of just QueryVisualizations API methods
type queryVisualizationsImpl struct {
	client *client.DatabricksClient
}

func (a *queryVisualizationsImpl) Create(ctx context.Context, request CreateVisualizationRequest) (*Visualization, error) {
	var visualization Visualization
	path := "/api/2.0/sql/visualizations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &visualization)
	return &visualization, err
}

func (a *queryVisualizationsImpl) Delete(ctx context.Context, request DeleteVisualizationRequest) error {
	var empty Empty
	path := fmt.Sprintf("/api/2.0/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &empty)
	return err
}

func (a *queryVisualizationsImpl) Update(ctx context.Context, request UpdateVisualizationRequest) (*Visualization, error) {
	var visualization Visualization
	path := fmt.Sprintf("/api/2.0/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &visualization)
	return &visualization, err
}

// unexported type that holds implementations of just QueryVisualizationsLegacy API methods
type queryVisualizationsLegacyImpl struct {
	client *client.DatabricksClient
}

func (a *queryVisualizationsLegacyImpl) Create(ctx context.Context, request CreateQueryVisualizationsLegacyRequest) (*LegacyVisualization, error) {
	var legacyVisualization LegacyVisualization
	path := "/api/2.0/preview/sql/visualizations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyVisualization)
	return &legacyVisualization, err
}

func (a *queryVisualizationsLegacyImpl) Delete(ctx context.Context, request DeleteQueryVisualizationsLegacyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *queryVisualizationsLegacyImpl) Update(ctx context.Context, request LegacyVisualization) (*LegacyVisualization, error) {
	var legacyVisualization LegacyVisualization
	path := fmt.Sprintf("/api/2.0/preview/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyVisualization)
	return &legacyVisualization, err
}

// unexported type that holds implementations of just RedashConfig API methods
type redashConfigImpl struct {
	client *client.DatabricksClient
}

func (a *redashConfigImpl) GetConfig(ctx context.Context) (*ClientConfig, error) {
	var clientConfig ClientConfig
	path := "/api/2.0/redash-v2/config"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &clientConfig)
	return &clientConfig, err
}

// unexported type that holds implementations of just StatementExecution API methods
type statementExecutionImpl struct {
	client *client.DatabricksClient
}

func (a *statementExecutionImpl) CancelExecution(ctx context.Context, request CancelExecutionRequest) error {
	var cancelExecutionResponse CancelExecutionResponse
	path := fmt.Sprintf("/api/2.0/sql/statements/%v/cancel", request.StatementId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &cancelExecutionResponse)
	return err
}

func (a *statementExecutionImpl) ExecuteStatement(ctx context.Context, request ExecuteStatementRequest) (*StatementResponse, error) {
	var statementResponse StatementResponse
	path := "/api/2.0/sql/statements/"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &statementResponse)
	return &statementResponse, err
}

func (a *statementExecutionImpl) GetStatement(ctx context.Context, request GetStatementRequest) (*StatementResponse, error) {
	var statementResponse StatementResponse
	path := fmt.Sprintf("/api/2.0/sql/statements/%v", request.StatementId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &statementResponse)
	return &statementResponse, err
}

func (a *statementExecutionImpl) GetStatementResultChunkN(ctx context.Context, request GetStatementResultChunkNRequest) (*ResultData, error) {
	var resultData ResultData
	path := fmt.Sprintf("/api/2.0/sql/statements/%v/result/chunks/%v", request.StatementId, request.ChunkIndex)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &resultData)
	return &resultData, err
}

// unexported type that holds implementations of just Warehouses API methods
type warehousesImpl struct {
	client *client.DatabricksClient
}

func (a *warehousesImpl) Create(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	var createWarehouseResponse CreateWarehouseResponse
	path := "/api/2.0/sql/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createWarehouseResponse)
	return &createWarehouseResponse, err
}

func (a *warehousesImpl) Delete(ctx context.Context, request DeleteWarehouseRequest) error {
	var deleteWarehouseResponse DeleteWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteWarehouseResponse)
	return err
}

func (a *warehousesImpl) Edit(ctx context.Context, request EditWarehouseRequest) error {
	var editWarehouseResponse EditWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/edit", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &editWarehouseResponse)
	return err
}

func (a *warehousesImpl) Get(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error) {
	var getWarehouseResponse GetWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getWarehouseResponse)
	return &getWarehouseResponse, err
}

func (a *warehousesImpl) GetPermissionLevels(ctx context.Context, request GetWarehousePermissionLevelsRequest) (*GetWarehousePermissionLevelsResponse, error) {
	var getWarehousePermissionLevelsResponse GetWarehousePermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v/permissionLevels", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getWarehousePermissionLevelsResponse)
	return &getWarehousePermissionLevelsResponse, err
}

func (a *warehousesImpl) GetPermissions(ctx context.Context, request GetWarehousePermissionsRequest) (*WarehousePermissions, error) {
	var warehousePermissions WarehousePermissions
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &warehousePermissions)
	return &warehousePermissions, err
}

func (a *warehousesImpl) GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error) {
	var getWorkspaceWarehouseConfigResponse GetWorkspaceWarehouseConfigResponse
	path := "/api/2.0/sql/config/warehouses"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getWorkspaceWarehouseConfigResponse)
	return &getWorkspaceWarehouseConfigResponse, err
}

// List warehouses.
//
// Lists all SQL warehouses that a user has manager permissions on.
func (a *warehousesImpl) List(ctx context.Context, request ListWarehousesRequest) listing.Iterator[EndpointInfo] {

	getNextPage := func(ctx context.Context, req ListWarehousesRequest) (*ListWarehousesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListWarehousesResponse) []EndpointInfo {
		return resp.Warehouses
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List warehouses.
//
// Lists all SQL warehouses that a user has manager permissions on.
func (a *warehousesImpl) ListAll(ctx context.Context, request ListWarehousesRequest) ([]EndpointInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[EndpointInfo](ctx, iterator)
}

func (a *warehousesImpl) internalList(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error) {
	var listWarehousesResponse ListWarehousesResponse
	path := "/api/2.0/sql/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWarehousesResponse)
	return &listWarehousesResponse, err
}

func (a *warehousesImpl) SetPermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error) {
	var warehousePermissions WarehousePermissions
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &warehousePermissions)
	return &warehousePermissions, err
}

func (a *warehousesImpl) SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error {
	var setWorkspaceWarehouseConfigResponse SetWorkspaceWarehouseConfigResponse
	path := "/api/2.0/sql/config/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &setWorkspaceWarehouseConfigResponse)
	return err
}

func (a *warehousesImpl) Start(ctx context.Context, request StartRequest) error {
	var startWarehouseResponse StartWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/start", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &startWarehouseResponse)
	return err
}

func (a *warehousesImpl) Stop(ctx context.Context, request StopRequest) error {
	var stopWarehouseResponse StopWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/stop", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &stopWarehouseResponse)
	return err
}

func (a *warehousesImpl) UpdatePermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error) {
	var warehousePermissions WarehousePermissions
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &warehousePermissions)
	return &warehousePermissions, err
}

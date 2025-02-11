// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sqlpreview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

// unexported type that holds implementations of just AlertsLegacyPreview API methods
type alertsLegacyPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *alertsLegacyPreviewImpl) Create(ctx context.Context, request CreateAlert) (*LegacyAlert, error) {
	var legacyAlert LegacyAlert
	path := "/api/2.0preview/preview/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyAlert)
	return &legacyAlert, err
}

func (a *alertsLegacyPreviewImpl) Delete(ctx context.Context, request DeleteAlertsLegacyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/preview/sql/alerts/%v", request.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *alertsLegacyPreviewImpl) Get(ctx context.Context, request GetAlertsLegacyRequest) (*LegacyAlert, error) {
	var legacyAlert LegacyAlert
	path := fmt.Sprintf("/api/2.0preview/preview/sql/alerts/%v", request.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &legacyAlert)
	return &legacyAlert, err
}

func (a *alertsLegacyPreviewImpl) List(ctx context.Context) ([]LegacyAlert, error) {
	var legacyAlertList []LegacyAlert
	path := "/api/2.0preview/preview/sql/alerts"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &legacyAlertList)
	return legacyAlertList, err
}

func (a *alertsLegacyPreviewImpl) Update(ctx context.Context, request EditAlert) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0preview/preview/sql/alerts/%v", request.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just AlertsPreview API methods
type alertsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *alertsPreviewImpl) Create(ctx context.Context, request CreateAlertRequest) (*Alert, error) {
	var alert Alert
	path := "/api/2.0preview/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &alert)
	return &alert, err
}

func (a *alertsPreviewImpl) Delete(ctx context.Context, request TrashAlertRequest) error {
	var empty Empty
	path := fmt.Sprintf("/api/2.0preview/sql/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &empty)
	return err
}

func (a *alertsPreviewImpl) Get(ctx context.Context, request GetAlertRequest) (*Alert, error) {
	var alert Alert
	path := fmt.Sprintf("/api/2.0preview/sql/alerts/%v", request.Id)
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
func (a *alertsPreviewImpl) List(ctx context.Context, request ListAlertsRequest) listing.Iterator[ListAlertsResponseAlert] {

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
func (a *alertsPreviewImpl) ListAll(ctx context.Context, request ListAlertsRequest) ([]ListAlertsResponseAlert, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ListAlertsResponseAlert](ctx, iterator)
}
func (a *alertsPreviewImpl) internalList(ctx context.Context, request ListAlertsRequest) (*ListAlertsResponse, error) {
	var listAlertsResponse ListAlertsResponse
	path := "/api/2.0preview/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAlertsResponse)
	return &listAlertsResponse, err
}

func (a *alertsPreviewImpl) Update(ctx context.Context, request UpdateAlertRequest) (*Alert, error) {
	var alert Alert
	path := fmt.Sprintf("/api/2.0preview/sql/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &alert)
	return &alert, err
}

// unexported type that holds implementations of just DashboardWidgetsPreview API methods
type dashboardWidgetsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardWidgetsPreviewImpl) Create(ctx context.Context, request CreateWidget) (*Widget, error) {
	var widget Widget
	path := "/api/2.0preview/preview/sql/widgets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &widget)
	return &widget, err
}

func (a *dashboardWidgetsPreviewImpl) Delete(ctx context.Context, request DeleteDashboardWidgetRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/preview/sql/widgets/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *dashboardWidgetsPreviewImpl) Update(ctx context.Context, request CreateWidget) (*Widget, error) {
	var widget Widget
	path := fmt.Sprintf("/api/2.0preview/preview/sql/widgets/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &widget)
	return &widget, err
}

// unexported type that holds implementations of just DashboardsPreview API methods
type dashboardsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardsPreviewImpl) Create(ctx context.Context, request DashboardPostContent) (*Dashboard, error) {
	var dashboard Dashboard
	path := "/api/2.0preview/preview/sql/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &dashboard)
	return &dashboard, err
}

func (a *dashboardsPreviewImpl) Delete(ctx context.Context, request DeleteDashboardRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/preview/sql/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *dashboardsPreviewImpl) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0preview/preview/sql/dashboards/%v", request.DashboardId)
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
func (a *dashboardsPreviewImpl) List(ctx context.Context, request ListDashboardsRequest) listing.Iterator[Dashboard] {

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
func (a *dashboardsPreviewImpl) ListAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Dashboard, int](ctx, iterator, request.PageSize)

}
func (a *dashboardsPreviewImpl) internalList(ctx context.Context, request ListDashboardsRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0preview/preview/sql/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listResponse)
	return &listResponse, err
}

func (a *dashboardsPreviewImpl) Restore(ctx context.Context, request RestoreDashboardRequest) error {
	var restoreResponse RestoreResponse
	path := fmt.Sprintf("/api/2.0preview/preview/sql/dashboards/trash/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &restoreResponse)
	return err
}

func (a *dashboardsPreviewImpl) Update(ctx context.Context, request DashboardEditContent) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0preview/preview/sql/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &dashboard)
	return &dashboard, err
}

// unexported type that holds implementations of just DataSourcesPreview API methods
type dataSourcesPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *dataSourcesPreviewImpl) List(ctx context.Context) ([]DataSource, error) {
	var dataSourceList []DataSource
	path := "/api/2.0preview/preview/sql/data_sources"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &dataSourceList)
	return dataSourceList, err
}

// unexported type that holds implementations of just DbsqlPermissionsPreview API methods
type dbsqlPermissionsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *dbsqlPermissionsPreviewImpl) Get(ctx context.Context, request GetDbsqlPermissionRequest) (*GetResponse, error) {
	var getResponse GetResponse
	path := fmt.Sprintf("/api/2.0preview/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getResponse)
	return &getResponse, err
}

func (a *dbsqlPermissionsPreviewImpl) Set(ctx context.Context, request SetRequest) (*SetResponse, error) {
	var setResponse SetResponse
	path := fmt.Sprintf("/api/2.0preview/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &setResponse)
	return &setResponse, err
}

func (a *dbsqlPermissionsPreviewImpl) TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error) {
	var success Success
	path := fmt.Sprintf("/api/2.0preview/preview/sql/permissions/%v/%v/transfer", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &success)
	return &success, err
}

// unexported type that holds implementations of just QueriesLegacyPreview API methods
type queriesLegacyPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *queriesLegacyPreviewImpl) Create(ctx context.Context, request QueryPostContent) (*LegacyQuery, error) {
	var legacyQuery LegacyQuery
	path := "/api/2.0preview/preview/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyQuery)
	return &legacyQuery, err
}

func (a *queriesLegacyPreviewImpl) Delete(ctx context.Context, request DeleteQueriesLegacyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/preview/sql/queries/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *queriesLegacyPreviewImpl) Get(ctx context.Context, request GetQueriesLegacyRequest) (*LegacyQuery, error) {
	var legacyQuery LegacyQuery
	path := fmt.Sprintf("/api/2.0preview/preview/sql/queries/%v", request.QueryId)
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
func (a *queriesLegacyPreviewImpl) List(ctx context.Context, request ListQueriesLegacyRequest) listing.Iterator[LegacyQuery] {

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
func (a *queriesLegacyPreviewImpl) ListAll(ctx context.Context, request ListQueriesLegacyRequest) ([]LegacyQuery, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[LegacyQuery, int](ctx, iterator, request.PageSize)

}
func (a *queriesLegacyPreviewImpl) internalList(ctx context.Context, request ListQueriesLegacyRequest) (*QueryList, error) {
	var queryList QueryList
	path := "/api/2.0preview/preview/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &queryList)
	return &queryList, err
}

func (a *queriesLegacyPreviewImpl) Restore(ctx context.Context, request RestoreQueriesLegacyRequest) error {
	var restoreResponse RestoreResponse
	path := fmt.Sprintf("/api/2.0preview/preview/sql/queries/trash/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &restoreResponse)
	return err
}

func (a *queriesLegacyPreviewImpl) Update(ctx context.Context, request QueryEditContent) (*LegacyQuery, error) {
	var legacyQuery LegacyQuery
	path := fmt.Sprintf("/api/2.0preview/preview/sql/queries/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyQuery)
	return &legacyQuery, err
}

// unexported type that holds implementations of just QueriesPreview API methods
type queriesPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *queriesPreviewImpl) Create(ctx context.Context, request CreateQueryRequest) (*Query, error) {
	var query Query
	path := "/api/2.0preview/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &query)
	return &query, err
}

func (a *queriesPreviewImpl) Delete(ctx context.Context, request TrashQueryRequest) error {
	var empty Empty
	path := fmt.Sprintf("/api/2.0preview/sql/queries/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &empty)
	return err
}

func (a *queriesPreviewImpl) Get(ctx context.Context, request GetQueryRequest) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0preview/sql/queries/%v", request.Id)
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
func (a *queriesPreviewImpl) List(ctx context.Context, request ListQueriesRequest) listing.Iterator[ListQueryObjectsResponseQuery] {

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
func (a *queriesPreviewImpl) ListAll(ctx context.Context, request ListQueriesRequest) ([]ListQueryObjectsResponseQuery, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ListQueryObjectsResponseQuery](ctx, iterator)
}
func (a *queriesPreviewImpl) internalList(ctx context.Context, request ListQueriesRequest) (*ListQueryObjectsResponse, error) {
	var listQueryObjectsResponse ListQueryObjectsResponse
	path := "/api/2.0preview/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listQueryObjectsResponse)
	return &listQueryObjectsResponse, err
}

// List visualizations on a query.
//
// Gets a list of visualizations on a query.
func (a *queriesPreviewImpl) ListVisualizations(ctx context.Context, request ListVisualizationsForQueryRequest) listing.Iterator[Visualization] {

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
func (a *queriesPreviewImpl) ListVisualizationsAll(ctx context.Context, request ListVisualizationsForQueryRequest) ([]Visualization, error) {
	iterator := a.ListVisualizations(ctx, request)
	return listing.ToSlice[Visualization](ctx, iterator)
}
func (a *queriesPreviewImpl) internalListVisualizations(ctx context.Context, request ListVisualizationsForQueryRequest) (*ListVisualizationsForQueryResponse, error) {
	var listVisualizationsForQueryResponse ListVisualizationsForQueryResponse
	path := fmt.Sprintf("/api/2.0preview/sql/queries/%v/visualizations", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listVisualizationsForQueryResponse)
	return &listVisualizationsForQueryResponse, err
}

func (a *queriesPreviewImpl) Update(ctx context.Context, request UpdateQueryRequest) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0preview/sql/queries/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &query)
	return &query, err
}

// unexported type that holds implementations of just QueryHistoryPreview API methods
type queryHistoryPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *queryHistoryPreviewImpl) List(ctx context.Context, request ListQueryHistoryRequest) (*ListQueriesResponse, error) {
	var listQueriesResponse ListQueriesResponse
	path := "/api/2.0preview/sql/history/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listQueriesResponse)
	return &listQueriesResponse, err
}

// unexported type that holds implementations of just QueryVisualizationsLegacyPreview API methods
type queryVisualizationsLegacyPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *queryVisualizationsLegacyPreviewImpl) Create(ctx context.Context, request CreateQueryVisualizationsLegacyRequest) (*LegacyVisualization, error) {
	var legacyVisualization LegacyVisualization
	path := "/api/2.0preview/preview/sql/visualizations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyVisualization)
	return &legacyVisualization, err
}

func (a *queryVisualizationsLegacyPreviewImpl) Delete(ctx context.Context, request DeleteQueryVisualizationsLegacyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/preview/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *queryVisualizationsLegacyPreviewImpl) Update(ctx context.Context, request LegacyVisualization) (*LegacyVisualization, error) {
	var legacyVisualization LegacyVisualization
	path := fmt.Sprintf("/api/2.0preview/preview/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &legacyVisualization)
	return &legacyVisualization, err
}

// unexported type that holds implementations of just QueryVisualizationsPreview API methods
type queryVisualizationsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *queryVisualizationsPreviewImpl) Create(ctx context.Context, request CreateVisualizationRequest) (*Visualization, error) {
	var visualization Visualization
	path := "/api/2.0preview/sql/visualizations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &visualization)
	return &visualization, err
}

func (a *queryVisualizationsPreviewImpl) Delete(ctx context.Context, request DeleteVisualizationRequest) error {
	var empty Empty
	path := fmt.Sprintf("/api/2.0preview/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &empty)
	return err
}

func (a *queryVisualizationsPreviewImpl) Update(ctx context.Context, request UpdateVisualizationRequest) (*Visualization, error) {
	var visualization Visualization
	path := fmt.Sprintf("/api/2.0preview/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &visualization)
	return &visualization, err
}

// unexported type that holds implementations of just RedashConfigPreview API methods
type redashConfigPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *redashConfigPreviewImpl) GetConfig(ctx context.Context) (*ClientConfig, error) {
	var clientConfig ClientConfig
	path := "/api/2.0preview/redash-v2/config"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &clientConfig)
	return &clientConfig, err
}

// unexported type that holds implementations of just StatementExecutionPreview API methods
type statementExecutionPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *statementExecutionPreviewImpl) CancelExecution(ctx context.Context, request CancelExecutionRequest) error {
	var cancelExecutionResponse CancelExecutionResponse
	path := fmt.Sprintf("/api/2.0preview/sql/statements/%v/cancel", request.StatementId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &cancelExecutionResponse)
	return err
}

func (a *statementExecutionPreviewImpl) ExecuteStatement(ctx context.Context, request ExecuteStatementRequest) (*StatementResponse, error) {
	var statementResponse StatementResponse
	path := "/api/2.0preview/sql/statements/"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &statementResponse)
	return &statementResponse, err
}

func (a *statementExecutionPreviewImpl) GetStatement(ctx context.Context, request GetStatementRequest) (*StatementResponse, error) {
	var statementResponse StatementResponse
	path := fmt.Sprintf("/api/2.0preview/sql/statements/%v", request.StatementId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &statementResponse)
	return &statementResponse, err
}

func (a *statementExecutionPreviewImpl) GetStatementResultChunkN(ctx context.Context, request GetStatementResultChunkNRequest) (*ResultData, error) {
	var resultData ResultData
	path := fmt.Sprintf("/api/2.0preview/sql/statements/%v/result/chunks/%v", request.StatementId, request.ChunkIndex)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &resultData)
	return &resultData, err
}

// unexported type that holds implementations of just WarehousesPreview API methods
type warehousesPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *warehousesPreviewImpl) Create(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	var createWarehouseResponse CreateWarehouseResponse
	path := "/api/2.0preview/sql/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createWarehouseResponse)
	return &createWarehouseResponse, err
}

func (a *warehousesPreviewImpl) Delete(ctx context.Context, request DeleteWarehouseRequest) error {
	var deleteWarehouseResponse DeleteWarehouseResponse
	path := fmt.Sprintf("/api/2.0preview/sql/warehouses/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteWarehouseResponse)
	return err
}

func (a *warehousesPreviewImpl) Edit(ctx context.Context, request EditWarehouseRequest) error {
	var editWarehouseResponse EditWarehouseResponse
	path := fmt.Sprintf("/api/2.0preview/sql/warehouses/%v/edit", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &editWarehouseResponse)
	return err
}

func (a *warehousesPreviewImpl) Get(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error) {
	var getWarehouseResponse GetWarehouseResponse
	path := fmt.Sprintf("/api/2.0preview/sql/warehouses/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getWarehouseResponse)
	return &getWarehouseResponse, err
}

func (a *warehousesPreviewImpl) GetPermissionLevels(ctx context.Context, request GetWarehousePermissionLevelsRequest) (*GetWarehousePermissionLevelsResponse, error) {
	var getWarehousePermissionLevelsResponse GetWarehousePermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0preview/permissions/warehouses/%v/permissionLevels", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getWarehousePermissionLevelsResponse)
	return &getWarehousePermissionLevelsResponse, err
}

func (a *warehousesPreviewImpl) GetPermissions(ctx context.Context, request GetWarehousePermissionsRequest) (*WarehousePermissions, error) {
	var warehousePermissions WarehousePermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/warehouses/%v", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &warehousePermissions)
	return &warehousePermissions, err
}

func (a *warehousesPreviewImpl) GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error) {
	var getWorkspaceWarehouseConfigResponse GetWorkspaceWarehouseConfigResponse
	path := "/api/2.0preview/sql/config/warehouses"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getWorkspaceWarehouseConfigResponse)
	return &getWorkspaceWarehouseConfigResponse, err
}

// List warehouses.
//
// Lists all SQL warehouses that a user has manager permissions on.
func (a *warehousesPreviewImpl) List(ctx context.Context, request ListWarehousesRequest) listing.Iterator[EndpointInfo] {

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
func (a *warehousesPreviewImpl) ListAll(ctx context.Context, request ListWarehousesRequest) ([]EndpointInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[EndpointInfo](ctx, iterator)
}
func (a *warehousesPreviewImpl) internalList(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error) {
	var listWarehousesResponse ListWarehousesResponse
	path := "/api/2.0preview/sql/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listWarehousesResponse)
	return &listWarehousesResponse, err
}

func (a *warehousesPreviewImpl) SetPermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error) {
	var warehousePermissions WarehousePermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/warehouses/%v", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &warehousePermissions)
	return &warehousePermissions, err
}

func (a *warehousesPreviewImpl) SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error {
	var setWorkspaceWarehouseConfigResponse SetWorkspaceWarehouseConfigResponse
	path := "/api/2.0preview/sql/config/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &setWorkspaceWarehouseConfigResponse)
	return err
}

func (a *warehousesPreviewImpl) Start(ctx context.Context, request StartRequest) error {
	var startWarehouseResponse StartWarehouseResponse
	path := fmt.Sprintf("/api/2.0preview/sql/warehouses/%v/start", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &startWarehouseResponse)
	return err
}

func (a *warehousesPreviewImpl) Stop(ctx context.Context, request StopRequest) error {
	var stopWarehouseResponse StopWarehouseResponse
	path := fmt.Sprintf("/api/2.0preview/sql/warehouses/%v/stop", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &stopWarehouseResponse)
	return err
}

func (a *warehousesPreviewImpl) UpdatePermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error) {
	var warehousePermissions WarehousePermissions
	path := fmt.Sprintf("/api/2.0preview/permissions/warehouses/%v", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &warehousePermissions)
	return &warehousePermissions, err
}

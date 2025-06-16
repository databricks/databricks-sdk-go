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

	requestPb, pbErr := createAlertRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var alertPb alertPb
	path := "/api/2.0/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&alertPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := alertFromPb(&alertPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *alertsImpl) Delete(ctx context.Context, request TrashAlertRequest) error {

	requestPb, pbErr := trashAlertRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var emptyPb emptyPb
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&emptyPb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *alertsImpl) Get(ctx context.Context, request GetAlertRequest) (*Alert, error) {

	requestPb, pbErr := getAlertRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var alertPb alertPb
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&alertPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := alertFromPb(&alertPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets a list of alerts accessible to the user, ordered by creation time.
// **Warning:** Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *alertsImpl) ListAll(ctx context.Context, request ListAlertsRequest) ([]ListAlertsResponseAlert, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ListAlertsResponseAlert](ctx, iterator)
}

func (a *alertsImpl) internalList(ctx context.Context, request ListAlertsRequest) (*ListAlertsResponse, error) {

	requestPb, pbErr := listAlertsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAlertsResponsePb listAlertsResponsePb
	path := "/api/2.0/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listAlertsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listAlertsResponseFromPb(&listAlertsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *alertsImpl) Update(ctx context.Context, request UpdateAlertRequest) (*Alert, error) {

	requestPb, pbErr := updateAlertRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var alertPb alertPb
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&alertPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := alertFromPb(&alertPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just AlertsLegacy API methods
type alertsLegacyImpl struct {
	client *client.DatabricksClient
}

func (a *alertsLegacyImpl) Create(ctx context.Context, request CreateAlert) (*LegacyAlert, error) {

	requestPb, pbErr := createAlertToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyAlertPb legacyAlertPb
	path := "/api/2.0/preview/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&legacyAlertPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := legacyAlertFromPb(&legacyAlertPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *alertsLegacyImpl) Delete(ctx context.Context, request DeleteAlertsLegacyRequest) error {

	requestPb, pbErr := deleteAlertsLegacyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", requestPb.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *alertsLegacyImpl) Get(ctx context.Context, request GetAlertsLegacyRequest) (*LegacyAlert, error) {

	requestPb, pbErr := getAlertsLegacyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyAlertPb legacyAlertPb
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", requestPb.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&legacyAlertPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := legacyAlertFromPb(&legacyAlertPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *alertsLegacyImpl) List(ctx context.Context) ([]LegacyAlert, error) {

	var legacyAlertListPb []legacyAlertPb
	path := "/api/2.0/preview/sql/alerts"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&legacyAlertListPb,
	)
	if err != nil {
		return nil, err
	}
	var resp []LegacyAlert
	for _, item := range legacyAlertListPb {
		itemResp, err := legacyAlertFromPb(&item)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *itemResp)
	}

	return resp, err
}

func (a *alertsLegacyImpl) Update(ctx context.Context, request EditAlert) error {

	requestPb, pbErr := editAlertToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updateResponsePb updateResponsePb
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", requestPb.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updateResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just AlertsV2 API methods
type alertsV2Impl struct {
	client *client.DatabricksClient
}

func (a *alertsV2Impl) CreateAlert(ctx context.Context, request CreateAlertV2Request) (*AlertV2, error) {

	requestPb, pbErr := createAlertV2RequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var alertV2Pb alertV2Pb
	path := "/api/2.0/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb).Alert,
		&alertV2Pb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := alertV2FromPb(&alertV2Pb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *alertsV2Impl) GetAlert(ctx context.Context, request GetAlertV2Request) (*AlertV2, error) {

	requestPb, pbErr := getAlertV2RequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var alertV2Pb alertV2Pb
	path := fmt.Sprintf("/api/2.0/alerts/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&alertV2Pb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := alertV2FromPb(&alertV2Pb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets a list of alerts accessible to the user, ordered by creation time.
func (a *alertsV2Impl) ListAlerts(ctx context.Context, request ListAlertsV2Request) listing.Iterator[AlertV2] {

	getNextPage := func(ctx context.Context, req ListAlertsV2Request) (*ListAlertsV2Response, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListAlerts(ctx, req)
	}
	getItems := func(resp *ListAlertsV2Response) []AlertV2 {
		return resp.Results
	}
	getNextReq := func(resp *ListAlertsV2Response) *ListAlertsV2Request {
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

// Gets a list of alerts accessible to the user, ordered by creation time.
func (a *alertsV2Impl) ListAlertsAll(ctx context.Context, request ListAlertsV2Request) ([]AlertV2, error) {
	iterator := a.ListAlerts(ctx, request)
	return listing.ToSlice[AlertV2](ctx, iterator)
}

func (a *alertsV2Impl) internalListAlerts(ctx context.Context, request ListAlertsV2Request) (*ListAlertsV2Response, error) {

	requestPb, pbErr := listAlertsV2RequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAlertsV2ResponsePb listAlertsV2ResponsePb
	path := "/api/2.0/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listAlertsV2ResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listAlertsV2ResponseFromPb(&listAlertsV2ResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *alertsV2Impl) TrashAlert(ctx context.Context, request TrashAlertV2Request) error {

	requestPb, pbErr := trashAlertV2RequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var emptyPb emptyPb
	path := fmt.Sprintf("/api/2.0/alerts/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&emptyPb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *alertsV2Impl) UpdateAlert(ctx context.Context, request UpdateAlertV2Request) (*AlertV2, error) {

	requestPb, pbErr := updateAlertV2RequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var alertV2Pb alertV2Pb
	path := fmt.Sprintf("/api/2.0/alerts/%v", requestPb.Id)
	queryParams := make(map[string]any)
	if requestPb.UpdateMask != "" {
		queryParams["update_mask"] = requestPb.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb).Alert,
		&alertV2Pb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := alertV2FromPb(&alertV2Pb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just DashboardWidgets API methods
type dashboardWidgetsImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardWidgetsImpl) Create(ctx context.Context, request CreateWidget) (*Widget, error) {

	requestPb, pbErr := createWidgetToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var widgetPb widgetPb
	path := "/api/2.0/preview/sql/widgets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&widgetPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := widgetFromPb(&widgetPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dashboardWidgetsImpl) Delete(ctx context.Context, request DeleteDashboardWidgetRequest) error {

	requestPb, pbErr := deleteDashboardWidgetRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/preview/sql/widgets/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *dashboardWidgetsImpl) Update(ctx context.Context, request CreateWidget) (*Widget, error) {

	requestPb, pbErr := createWidgetToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var widgetPb widgetPb
	path := fmt.Sprintf("/api/2.0/preview/sql/widgets/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&widgetPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := widgetFromPb(&widgetPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Dashboards API methods
type dashboardsImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardsImpl) Create(ctx context.Context, request DashboardPostContent) (*Dashboard, error) {

	requestPb, pbErr := dashboardPostContentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb dashboardPb
	path := "/api/2.0/preview/sql/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&dashboardPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := dashboardFromPb(&dashboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dashboardsImpl) Delete(ctx context.Context, request DeleteDashboardRequest) error {

	requestPb, pbErr := deleteDashboardRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", requestPb.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *dashboardsImpl) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {

	requestPb, pbErr := getDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb dashboardPb
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", requestPb.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&dashboardPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := dashboardFromPb(&dashboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Fetch a paginated list of dashboard objects.
//
// **Warning**: Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *dashboardsImpl) ListAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[Dashboard, int](ctx, iterator, request.PageSize)

}

func (a *dashboardsImpl) internalList(ctx context.Context, request ListDashboardsRequest) (*ListResponse, error) {

	requestPb, pbErr := listDashboardsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listResponsePb listResponsePb
	path := "/api/2.0/preview/sql/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listResponseFromPb(&listResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dashboardsImpl) Restore(ctx context.Context, request RestoreDashboardRequest) error {

	requestPb, pbErr := restoreDashboardRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var restoreResponsePb restoreResponsePb
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/trash/%v", requestPb.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		&restoreResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *dashboardsImpl) Update(ctx context.Context, request DashboardEditContent) (*Dashboard, error) {

	requestPb, pbErr := dashboardEditContentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb dashboardPb
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", requestPb.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&dashboardPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := dashboardFromPb(&dashboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just DataSources API methods
type dataSourcesImpl struct {
	client *client.DatabricksClient
}

func (a *dataSourcesImpl) List(ctx context.Context) ([]DataSource, error) {

	var dataSourceListPb []dataSourcePb
	path := "/api/2.0/preview/sql/data_sources"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&dataSourceListPb,
	)
	if err != nil {
		return nil, err
	}
	var resp []DataSource
	for _, item := range dataSourceListPb {
		itemResp, err := dataSourceFromPb(&item)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *itemResp)
	}

	return resp, err
}

// unexported type that holds implementations of just DbsqlPermissions API methods
type dbsqlPermissionsImpl struct {
	client *client.DatabricksClient
}

func (a *dbsqlPermissionsImpl) Get(ctx context.Context, request GetDbsqlPermissionRequest) (*GetResponse, error) {

	requestPb, pbErr := getDbsqlPermissionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getResponsePb getResponsePb
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", requestPb.ObjectType, requestPb.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getResponseFromPb(&getResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dbsqlPermissionsImpl) Set(ctx context.Context, request SetRequest) (*SetResponse, error) {

	requestPb, pbErr := setRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var setResponsePb setResponsePb
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", requestPb.ObjectType, requestPb.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&setResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := setResponseFromPb(&setResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dbsqlPermissionsImpl) TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error) {

	requestPb, pbErr := transferOwnershipRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var successPb successPb
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v/transfer", requestPb.ObjectType, requestPb.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&successPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := successFromPb(&successPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Queries API methods
type queriesImpl struct {
	client *client.DatabricksClient
}

func (a *queriesImpl) Create(ctx context.Context, request CreateQueryRequest) (*Query, error) {

	requestPb, pbErr := createQueryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryPb queryPb
	path := "/api/2.0/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&queryPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := queryFromPb(&queryPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *queriesImpl) Delete(ctx context.Context, request TrashQueryRequest) error {

	requestPb, pbErr := trashQueryRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var emptyPb emptyPb
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&emptyPb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *queriesImpl) Get(ctx context.Context, request GetQueryRequest) (*Query, error) {

	requestPb, pbErr := getQueryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryPb queryPb
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&queryPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := queryFromPb(&queryPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets a list of queries accessible to the user, ordered by creation time.
// **Warning:** Calling this API concurrently 10 or more times could result in
// throttling, service degradation, or a temporary ban.
func (a *queriesImpl) ListAll(ctx context.Context, request ListQueriesRequest) ([]ListQueryObjectsResponseQuery, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ListQueryObjectsResponseQuery](ctx, iterator)
}

func (a *queriesImpl) internalList(ctx context.Context, request ListQueriesRequest) (*ListQueryObjectsResponse, error) {

	requestPb, pbErr := listQueriesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listQueryObjectsResponsePb listQueryObjectsResponsePb
	path := "/api/2.0/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listQueryObjectsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listQueryObjectsResponseFromPb(&listQueryObjectsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets a list of visualizations on a query.
func (a *queriesImpl) ListVisualizationsAll(ctx context.Context, request ListVisualizationsForQueryRequest) ([]Visualization, error) {
	iterator := a.ListVisualizations(ctx, request)
	return listing.ToSlice[Visualization](ctx, iterator)
}

func (a *queriesImpl) internalListVisualizations(ctx context.Context, request ListVisualizationsForQueryRequest) (*ListVisualizationsForQueryResponse, error) {

	requestPb, pbErr := listVisualizationsForQueryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listVisualizationsForQueryResponsePb listVisualizationsForQueryResponsePb
	path := fmt.Sprintf("/api/2.0/sql/queries/%v/visualizations", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listVisualizationsForQueryResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listVisualizationsForQueryResponseFromPb(&listVisualizationsForQueryResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *queriesImpl) Update(ctx context.Context, request UpdateQueryRequest) (*Query, error) {

	requestPb, pbErr := updateQueryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryPb queryPb
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&queryPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := queryFromPb(&queryPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just QueriesLegacy API methods
type queriesLegacyImpl struct {
	client *client.DatabricksClient
}

func (a *queriesLegacyImpl) Create(ctx context.Context, request QueryPostContent) (*LegacyQuery, error) {

	requestPb, pbErr := queryPostContentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyQueryPb legacyQueryPb
	path := "/api/2.0/preview/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&legacyQueryPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := legacyQueryFromPb(&legacyQueryPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *queriesLegacyImpl) Delete(ctx context.Context, request DeleteQueriesLegacyRequest) error {

	requestPb, pbErr := deleteQueriesLegacyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", requestPb.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *queriesLegacyImpl) Get(ctx context.Context, request GetQueriesLegacyRequest) (*LegacyQuery, error) {

	requestPb, pbErr := getQueriesLegacyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyQueryPb legacyQueryPb
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", requestPb.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&legacyQueryPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := legacyQueryFromPb(&legacyQueryPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

	requestPb, pbErr := listQueriesLegacyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryListPb queryListPb
	path := "/api/2.0/preview/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&queryListPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := queryListFromPb(&queryListPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *queriesLegacyImpl) Restore(ctx context.Context, request RestoreQueriesLegacyRequest) error {

	requestPb, pbErr := restoreQueriesLegacyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var restoreResponsePb restoreResponsePb
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/trash/%v", requestPb.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		&restoreResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *queriesLegacyImpl) Update(ctx context.Context, request QueryEditContent) (*LegacyQuery, error) {

	requestPb, pbErr := queryEditContentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyQueryPb legacyQueryPb
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", requestPb.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&legacyQueryPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := legacyQueryFromPb(&legacyQueryPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just QueryHistory API methods
type queryHistoryImpl struct {
	client *client.DatabricksClient
}

func (a *queryHistoryImpl) List(ctx context.Context, request ListQueryHistoryRequest) (*ListQueriesResponse, error) {

	requestPb, pbErr := listQueryHistoryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listQueriesResponsePb listQueriesResponsePb
	path := "/api/2.0/sql/history/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listQueriesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listQueriesResponseFromPb(&listQueriesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just QueryVisualizations API methods
type queryVisualizationsImpl struct {
	client *client.DatabricksClient
}

func (a *queryVisualizationsImpl) Create(ctx context.Context, request CreateVisualizationRequest) (*Visualization, error) {

	requestPb, pbErr := createVisualizationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var visualizationPb visualizationPb
	path := "/api/2.0/sql/visualizations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&visualizationPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := visualizationFromPb(&visualizationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *queryVisualizationsImpl) Delete(ctx context.Context, request DeleteVisualizationRequest) error {

	requestPb, pbErr := deleteVisualizationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var emptyPb emptyPb
	path := fmt.Sprintf("/api/2.0/sql/visualizations/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&emptyPb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *queryVisualizationsImpl) Update(ctx context.Context, request UpdateVisualizationRequest) (*Visualization, error) {

	requestPb, pbErr := updateVisualizationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var visualizationPb visualizationPb
	path := fmt.Sprintf("/api/2.0/sql/visualizations/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&visualizationPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := visualizationFromPb(&visualizationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just QueryVisualizationsLegacy API methods
type queryVisualizationsLegacyImpl struct {
	client *client.DatabricksClient
}

func (a *queryVisualizationsLegacyImpl) Create(ctx context.Context, request CreateQueryVisualizationsLegacyRequest) (*LegacyVisualization, error) {

	requestPb, pbErr := createQueryVisualizationsLegacyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyVisualizationPb legacyVisualizationPb
	path := "/api/2.0/preview/sql/visualizations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&legacyVisualizationPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := legacyVisualizationFromPb(&legacyVisualizationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *queryVisualizationsLegacyImpl) Delete(ctx context.Context, request DeleteQueryVisualizationsLegacyRequest) error {

	requestPb, pbErr := deleteQueryVisualizationsLegacyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/preview/sql/visualizations/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *queryVisualizationsLegacyImpl) Update(ctx context.Context, request LegacyVisualization) (*LegacyVisualization, error) {

	requestPb, pbErr := legacyVisualizationToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyVisualizationPb legacyVisualizationPb
	path := fmt.Sprintf("/api/2.0/preview/sql/visualizations/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&legacyVisualizationPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := legacyVisualizationFromPb(&legacyVisualizationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just RedashConfig API methods
type redashConfigImpl struct {
	client *client.DatabricksClient
}

func (a *redashConfigImpl) GetConfig(ctx context.Context) (*ClientConfig, error) {

	var clientConfigPb clientConfigPb
	path := "/api/2.0/redash-v2/config"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&clientConfigPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := clientConfigFromPb(&clientConfigPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just StatementExecution API methods
type statementExecutionImpl struct {
	client *client.DatabricksClient
}

func (a *statementExecutionImpl) CancelExecution(ctx context.Context, request CancelExecutionRequest) error {

	requestPb, pbErr := cancelExecutionRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var cancelExecutionResponsePb cancelExecutionResponsePb
	path := fmt.Sprintf("/api/2.0/sql/statements/%v/cancel", requestPb.StatementId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		&cancelExecutionResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *statementExecutionImpl) ExecuteStatement(ctx context.Context, request ExecuteStatementRequest) (*StatementResponse, error) {

	requestPb, pbErr := executeStatementRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var statementResponsePb statementResponsePb
	path := "/api/2.0/sql/statements/"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&statementResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := statementResponseFromPb(&statementResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *statementExecutionImpl) GetStatement(ctx context.Context, request GetStatementRequest) (*StatementResponse, error) {

	requestPb, pbErr := getStatementRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var statementResponsePb statementResponsePb
	path := fmt.Sprintf("/api/2.0/sql/statements/%v", requestPb.StatementId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&statementResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := statementResponseFromPb(&statementResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *statementExecutionImpl) GetStatementResultChunkN(ctx context.Context, request GetStatementResultChunkNRequest) (*ResultData, error) {

	requestPb, pbErr := getStatementResultChunkNRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var resultDataPb resultDataPb
	path := fmt.Sprintf("/api/2.0/sql/statements/%v/result/chunks/%v", requestPb.StatementId, requestPb.ChunkIndex)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&resultDataPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := resultDataFromPb(&resultDataPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Warehouses API methods
type warehousesImpl struct {
	client *client.DatabricksClient
}

func (a *warehousesImpl) Create(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error) {

	requestPb, pbErr := createWarehouseRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createWarehouseResponsePb createWarehouseResponsePb
	path := "/api/2.0/sql/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createWarehouseResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createWarehouseResponseFromPb(&createWarehouseResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *warehousesImpl) Delete(ctx context.Context, request DeleteWarehouseRequest) error {

	requestPb, pbErr := deleteWarehouseRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteWarehouseResponsePb deleteWarehouseResponsePb
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteWarehouseResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *warehousesImpl) Edit(ctx context.Context, request EditWarehouseRequest) error {

	requestPb, pbErr := editWarehouseRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var editWarehouseResponsePb editWarehouseResponsePb
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/edit", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&editWarehouseResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *warehousesImpl) Get(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error) {

	requestPb, pbErr := getWarehouseRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getWarehouseResponsePb getWarehouseResponsePb
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getWarehouseResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getWarehouseResponseFromPb(&getWarehouseResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *warehousesImpl) GetPermissionLevels(ctx context.Context, request GetWarehousePermissionLevelsRequest) (*GetWarehousePermissionLevelsResponse, error) {

	requestPb, pbErr := getWarehousePermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getWarehousePermissionLevelsResponsePb getWarehousePermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v/permissionLevels", requestPb.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getWarehousePermissionLevelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getWarehousePermissionLevelsResponseFromPb(&getWarehousePermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *warehousesImpl) GetPermissions(ctx context.Context, request GetWarehousePermissionsRequest) (*WarehousePermissions, error) {

	requestPb, pbErr := getWarehousePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var warehousePermissionsPb warehousePermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", requestPb.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&warehousePermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := warehousePermissionsFromPb(&warehousePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *warehousesImpl) GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error) {

	var getWorkspaceWarehouseConfigResponsePb getWorkspaceWarehouseConfigResponsePb
	path := "/api/2.0/sql/config/warehouses"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&getWorkspaceWarehouseConfigResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getWorkspaceWarehouseConfigResponseFromPb(&getWorkspaceWarehouseConfigResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Lists all SQL warehouses that a user has manager permissions on.
func (a *warehousesImpl) ListAll(ctx context.Context, request ListWarehousesRequest) ([]EndpointInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[EndpointInfo](ctx, iterator)
}

func (a *warehousesImpl) internalList(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error) {

	requestPb, pbErr := listWarehousesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listWarehousesResponsePb listWarehousesResponsePb
	path := "/api/2.0/sql/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listWarehousesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listWarehousesResponseFromPb(&listWarehousesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *warehousesImpl) SetPermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error) {

	requestPb, pbErr := warehousePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var warehousePermissionsPb warehousePermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", requestPb.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&warehousePermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := warehousePermissionsFromPb(&warehousePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *warehousesImpl) SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error {

	requestPb, pbErr := setWorkspaceWarehouseConfigRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var setWorkspaceWarehouseConfigResponsePb setWorkspaceWarehouseConfigResponsePb
	path := "/api/2.0/sql/config/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&setWorkspaceWarehouseConfigResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *warehousesImpl) Start(ctx context.Context, request StartRequest) error {

	requestPb, pbErr := startRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var startWarehouseResponsePb startWarehouseResponsePb
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/start", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		&startWarehouseResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *warehousesImpl) Stop(ctx context.Context, request StopRequest) error {

	requestPb, pbErr := stopRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var stopWarehouseResponsePb stopWarehouseResponsePb
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/stop", requestPb.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		&stopWarehouseResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *warehousesImpl) UpdatePermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error) {

	requestPb, pbErr := warehousePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var warehousePermissionsPb warehousePermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", requestPb.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&warehousePermissionsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := warehousePermissionsFromPb(&warehousePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

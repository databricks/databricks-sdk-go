// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/sql/sqlpb"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Alerts API methods
type alertsImpl struct {
	client *client.DatabricksClient
}

func (a *alertsImpl) Create(ctx context.Context, request CreateAlertRequest) (*Alert, error) {
	requestPb, pbErr := CreateAlertRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var alertPb sqlpb.AlertPb
	path := "/api/2.0/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := AlertFromPb(&alertPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *alertsImpl) Delete(ctx context.Context, request TrashAlertRequest) error {
	requestPb, pbErr := TrashAlertRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *alertsImpl) Get(ctx context.Context, request GetAlertRequest) (*Alert, error) {
	requestPb, pbErr := GetAlertRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var alertPb sqlpb.AlertPb
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := AlertFromPb(&alertPb)
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
	requestPb, pbErr := ListAlertsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAlertsResponsePb sqlpb.ListAlertsResponsePb
	path := "/api/2.0/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := ListAlertsResponseFromPb(&listAlertsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *alertsImpl) Update(ctx context.Context, request UpdateAlertRequest) (*Alert, error) {
	requestPb, pbErr := UpdateAlertRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var alertPb sqlpb.AlertPb
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := AlertFromPb(&alertPb)
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
	requestPb, pbErr := CreateAlertToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyAlertPb sqlpb.LegacyAlertPb
	path := "/api/2.0/preview/sql/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := LegacyAlertFromPb(&legacyAlertPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *alertsLegacyImpl) Delete(ctx context.Context, request DeleteAlertsLegacyRequest) error {
	requestPb, pbErr := DeleteAlertsLegacyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *alertsLegacyImpl) Get(ctx context.Context, request GetAlertsLegacyRequest) (*LegacyAlert, error) {
	requestPb, pbErr := GetAlertsLegacyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyAlertPb sqlpb.LegacyAlertPb
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := LegacyAlertFromPb(&legacyAlertPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *alertsLegacyImpl) List(ctx context.Context) ([]LegacyAlert, error) {

	var legacyAlertListPb []sqlpb.LegacyAlertPb
	path := "/api/2.0/preview/sql/alerts"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
		itemResp, err := LegacyAlertFromPb(&item)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *itemResp)
	}

	return resp, err
}

func (a *alertsLegacyImpl) Update(ctx context.Context, request EditAlert) error {
	requestPb, pbErr := EditAlertToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
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
	requestPb, pbErr := CreateAlertV2RequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var alertV2Pb sqlpb.AlertV2Pb
	path := "/api/2.0/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := AlertV2FromPb(&alertV2Pb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *alertsV2Impl) GetAlert(ctx context.Context, request GetAlertV2Request) (*AlertV2, error) {
	requestPb, pbErr := GetAlertV2RequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var alertV2Pb sqlpb.AlertV2Pb
	path := fmt.Sprintf("/api/2.0/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := AlertV2FromPb(&alertV2Pb)
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
	requestPb, pbErr := ListAlertsV2RequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAlertsV2ResponsePb sqlpb.ListAlertsV2ResponsePb
	path := "/api/2.0/alerts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := ListAlertsV2ResponseFromPb(&listAlertsV2ResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *alertsV2Impl) TrashAlert(ctx context.Context, request TrashAlertV2Request) error {
	requestPb, pbErr := TrashAlertV2RequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *alertsV2Impl) UpdateAlert(ctx context.Context, request UpdateAlertV2Request) (*AlertV2, error) {
	requestPb, pbErr := UpdateAlertV2RequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var alertV2Pb sqlpb.AlertV2Pb
	path := fmt.Sprintf("/api/2.0/alerts/%v", request.Id)
	queryParams := make(map[string]any)
	if requestPb.UpdateMask != "" {
		queryParams["update_mask"] = requestPb.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := AlertV2FromPb(&alertV2Pb)
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
	requestPb, pbErr := CreateWidgetToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var widgetPb sqlpb.WidgetPb
	path := "/api/2.0/preview/sql/widgets"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := WidgetFromPb(&widgetPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dashboardWidgetsImpl) Delete(ctx context.Context, request DeleteDashboardWidgetRequest) error {
	requestPb, pbErr := DeleteDashboardWidgetRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/sql/widgets/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *dashboardWidgetsImpl) Update(ctx context.Context, request UpdateWidgetRequest) (*Widget, error) {
	requestPb, pbErr := UpdateWidgetRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var widgetPb sqlpb.WidgetPb
	path := fmt.Sprintf("/api/2.0/preview/sql/widgets/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := WidgetFromPb(&widgetPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Dashboards API methods
type dashboardsImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardsImpl) Delete(ctx context.Context, request DeleteDashboardRequest) error {
	requestPb, pbErr := DeleteDashboardRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *dashboardsImpl) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	requestPb, pbErr := GetDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb sqlpb.DashboardPb
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := DashboardFromPb(&dashboardPb)
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
	return iterator
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
	requestPb, pbErr := ListDashboardsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listResponsePb sqlpb.ListResponsePb
	path := "/api/2.0/preview/sql/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := ListResponseFromPb(&listResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dashboardsImpl) Restore(ctx context.Context, request RestoreDashboardRequest) error {

	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/trash/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *dashboardsImpl) Update(ctx context.Context, request DashboardEditContent) (*Dashboard, error) {
	requestPb, pbErr := DashboardEditContentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb sqlpb.DashboardPb
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := DashboardFromPb(&dashboardPb)
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

	var dataSourceListPb []sqlpb.DataSourcePb
	path := "/api/2.0/preview/sql/data_sources"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
		itemResp, err := DataSourceFromPb(&item)
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
	requestPb, pbErr := GetDbsqlPermissionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getResponsePb sqlpb.GetResponsePb
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := GetResponseFromPb(&getResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dbsqlPermissionsImpl) Set(ctx context.Context, request SetRequest) (*SetResponse, error) {
	requestPb, pbErr := SetRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var setResponsePb sqlpb.SetResponsePb
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := SetResponseFromPb(&setResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *dbsqlPermissionsImpl) TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error) {
	requestPb, pbErr := TransferOwnershipRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var successPb sqlpb.SuccessPb
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v/transfer", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := SuccessFromPb(&successPb)
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
	requestPb, pbErr := CreateQueryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryPb sqlpb.QueryPb
	path := "/api/2.0/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := QueryFromPb(&queryPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *queriesImpl) Delete(ctx context.Context, request TrashQueryRequest) error {
	requestPb, pbErr := TrashQueryRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/sql/queries/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *queriesImpl) Get(ctx context.Context, request GetQueryRequest) (*Query, error) {
	requestPb, pbErr := GetQueryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryPb sqlpb.QueryPb
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := QueryFromPb(&queryPb)
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
	requestPb, pbErr := ListQueriesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listQueryObjectsResponsePb sqlpb.ListQueryObjectsResponsePb
	path := "/api/2.0/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := ListQueryObjectsResponseFromPb(&listQueryObjectsResponsePb)
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
	requestPb, pbErr := ListVisualizationsForQueryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listVisualizationsForQueryResponsePb sqlpb.ListVisualizationsForQueryResponsePb
	path := fmt.Sprintf("/api/2.0/sql/queries/%v/visualizations", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := ListVisualizationsForQueryResponseFromPb(&listVisualizationsForQueryResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *queriesImpl) Update(ctx context.Context, request UpdateQueryRequest) (*Query, error) {
	requestPb, pbErr := UpdateQueryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryPb sqlpb.QueryPb
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := QueryFromPb(&queryPb)
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
	requestPb, pbErr := QueryPostContentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyQueryPb sqlpb.LegacyQueryPb
	path := "/api/2.0/preview/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := LegacyQueryFromPb(&legacyQueryPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *queriesLegacyImpl) Delete(ctx context.Context, request DeleteQueriesLegacyRequest) error {
	requestPb, pbErr := DeleteQueriesLegacyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *queriesLegacyImpl) Get(ctx context.Context, request GetQueriesLegacyRequest) (*LegacyQuery, error) {
	requestPb, pbErr := GetQueriesLegacyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyQueryPb sqlpb.LegacyQueryPb
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := LegacyQueryFromPb(&legacyQueryPb)
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
	return iterator
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
	requestPb, pbErr := ListQueriesLegacyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryListPb sqlpb.QueryListPb
	path := "/api/2.0/preview/sql/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := QueryListFromPb(&queryListPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *queriesLegacyImpl) Restore(ctx context.Context, request RestoreQueriesLegacyRequest) error {

	path := fmt.Sprintf("/api/2.0/preview/sql/queries/trash/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *queriesLegacyImpl) Update(ctx context.Context, request QueryEditContent) (*LegacyQuery, error) {
	requestPb, pbErr := QueryEditContentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyQueryPb sqlpb.LegacyQueryPb
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := LegacyQueryFromPb(&legacyQueryPb)
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
	requestPb, pbErr := ListQueryHistoryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listQueriesResponsePb sqlpb.ListQueriesResponsePb
	path := "/api/2.0/sql/history/queries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := ListQueriesResponseFromPb(&listQueriesResponsePb)
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
	requestPb, pbErr := CreateVisualizationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var visualizationPb sqlpb.VisualizationPb
	path := "/api/2.0/sql/visualizations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := VisualizationFromPb(&visualizationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *queryVisualizationsImpl) Delete(ctx context.Context, request DeleteVisualizationRequest) error {
	requestPb, pbErr := DeleteVisualizationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *queryVisualizationsImpl) Update(ctx context.Context, request UpdateVisualizationRequest) (*Visualization, error) {
	requestPb, pbErr := UpdateVisualizationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var visualizationPb sqlpb.VisualizationPb
	path := fmt.Sprintf("/api/2.0/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := VisualizationFromPb(&visualizationPb)
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
	requestPb, pbErr := CreateQueryVisualizationsLegacyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyVisualizationPb sqlpb.LegacyVisualizationPb
	path := "/api/2.0/preview/sql/visualizations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := LegacyVisualizationFromPb(&legacyVisualizationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *queryVisualizationsLegacyImpl) Delete(ctx context.Context, request DeleteQueryVisualizationsLegacyRequest) error {
	requestPb, pbErr := DeleteQueryVisualizationsLegacyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/preview/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *queryVisualizationsLegacyImpl) Update(ctx context.Context, request LegacyVisualization) (*LegacyVisualization, error) {
	requestPb, pbErr := LegacyVisualizationToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var legacyVisualizationPb sqlpb.LegacyVisualizationPb
	path := fmt.Sprintf("/api/2.0/preview/sql/visualizations/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := LegacyVisualizationFromPb(&legacyVisualizationPb)
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

	var clientConfigPb sqlpb.ClientConfigPb
	path := "/api/2.0/redash-v2/config"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := ClientConfigFromPb(&clientConfigPb)
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

	path := fmt.Sprintf("/api/2.0/sql/statements/%v/cancel", request.StatementId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *statementExecutionImpl) ExecuteStatement(ctx context.Context, request ExecuteStatementRequest) (*StatementResponse, error) {
	requestPb, pbErr := ExecuteStatementRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var statementResponsePb sqlpb.StatementResponsePb
	path := "/api/2.0/sql/statements/"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := StatementResponseFromPb(&statementResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *statementExecutionImpl) GetStatement(ctx context.Context, request GetStatementRequest) (*StatementResponse, error) {
	requestPb, pbErr := GetStatementRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var statementResponsePb sqlpb.StatementResponsePb
	path := fmt.Sprintf("/api/2.0/sql/statements/%v", request.StatementId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := StatementResponseFromPb(&statementResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *statementExecutionImpl) GetStatementResultChunkN(ctx context.Context, request GetStatementResultChunkNRequest) (*ResultData, error) {
	requestPb, pbErr := GetStatementResultChunkNRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var resultDataPb sqlpb.ResultDataPb
	path := fmt.Sprintf("/api/2.0/sql/statements/%v/result/chunks/%v", request.StatementId, request.ChunkIndex)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := ResultDataFromPb(&resultDataPb)
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
	requestPb, pbErr := CreateWarehouseRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createWarehouseResponsePb sqlpb.CreateWarehouseResponsePb
	path := "/api/2.0/sql/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := CreateWarehouseResponseFromPb(&createWarehouseResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *warehousesImpl) Delete(ctx context.Context, request DeleteWarehouseRequest) error {
	requestPb, pbErr := DeleteWarehouseRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *warehousesImpl) Edit(ctx context.Context, request EditWarehouseRequest) error {
	requestPb, pbErr := EditWarehouseRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/edit", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *warehousesImpl) Get(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error) {
	requestPb, pbErr := GetWarehouseRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getWarehouseResponsePb sqlpb.GetWarehouseResponsePb
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := GetWarehouseResponseFromPb(&getWarehouseResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *warehousesImpl) GetPermissionLevels(ctx context.Context, request GetWarehousePermissionLevelsRequest) (*GetWarehousePermissionLevelsResponse, error) {
	requestPb, pbErr := GetWarehousePermissionLevelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getWarehousePermissionLevelsResponsePb sqlpb.GetWarehousePermissionLevelsResponsePb
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v/permissionLevels", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := GetWarehousePermissionLevelsResponseFromPb(&getWarehousePermissionLevelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *warehousesImpl) GetPermissions(ctx context.Context, request GetWarehousePermissionsRequest) (*WarehousePermissions, error) {
	requestPb, pbErr := GetWarehousePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var warehousePermissionsPb sqlpb.WarehousePermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := WarehousePermissionsFromPb(&warehousePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *warehousesImpl) GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error) {

	var getWorkspaceWarehouseConfigResponsePb sqlpb.GetWorkspaceWarehouseConfigResponsePb
	path := "/api/2.0/sql/config/warehouses"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := GetWorkspaceWarehouseConfigResponseFromPb(&getWorkspaceWarehouseConfigResponsePb)
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
	requestPb, pbErr := ListWarehousesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listWarehousesResponsePb sqlpb.ListWarehousesResponsePb
	path := "/api/2.0/sql/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := ListWarehousesResponseFromPb(&listWarehousesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *warehousesImpl) SetPermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error) {
	requestPb, pbErr := WarehousePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var warehousePermissionsPb sqlpb.WarehousePermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := WarehousePermissionsFromPb(&warehousePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *warehousesImpl) SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error {
	requestPb, pbErr := SetWorkspaceWarehouseConfigRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/sql/config/warehouses"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *warehousesImpl) Start(ctx context.Context, request StartRequest) error {

	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/start", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *warehousesImpl) Stop(ctx context.Context, request StopRequest) error {

	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/stop", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *warehousesImpl) UpdatePermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error) {
	requestPb, pbErr := WarehousePermissionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var warehousePermissionsPb sqlpb.WarehousePermissionsPb
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", request.WarehouseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
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
	resp, err := WarehousePermissionsFromPb(&warehousePermissionsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

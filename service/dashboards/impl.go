// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Lakeview API methods
type lakeviewImpl struct {
	client *client.DatabricksClient
}

func (a *lakeviewImpl) Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := "/api/2.0/lakeview/dashboards"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &dashboard)
	return &dashboard, err
}

func (a *lakeviewImpl) CreateSchedule(ctx context.Context, request CreateScheduleRequest) (*Schedule, error) {
	var schedule Schedule
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules", request.DashboardId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &schedule)
	return &schedule, err
}

func (a *lakeviewImpl) CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (*Subscription, error) {
	var subscription Subscription
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions", request.DashboardId, request.ScheduleId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &subscription)
	return &subscription, err
}

func (a *lakeviewImpl) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error {
	var deleteScheduleResponse DeleteScheduleResponse
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v", request.DashboardId, request.ScheduleId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteScheduleResponse)
	return err
}

func (a *lakeviewImpl) DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) error {
	var deleteSubscriptionResponse DeleteSubscriptionResponse
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions/%v", request.DashboardId, request.ScheduleId, request.SubscriptionId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteSubscriptionResponse)
	return err
}

func (a *lakeviewImpl) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v", request.DashboardId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &dashboard)
	return &dashboard, err
}

func (a *lakeviewImpl) GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error) {
	var publishedDashboard PublishedDashboard
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", request.DashboardId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &publishedDashboard)
	return &publishedDashboard, err
}

func (a *lakeviewImpl) GetSchedule(ctx context.Context, request GetScheduleRequest) (*Schedule, error) {
	var schedule Schedule
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v", request.DashboardId, request.ScheduleId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &schedule)
	return &schedule, err
}

func (a *lakeviewImpl) GetSubscription(ctx context.Context, request GetSubscriptionRequest) (*Subscription, error) {
	var subscription Subscription
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions/%v", request.DashboardId, request.ScheduleId, request.SubscriptionId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &subscription)
	return &subscription, err
}

func (a *lakeviewImpl) List(ctx context.Context, request ListDashboardsRequest) (*ListDashboardsResponse, error) {
	var listDashboardsResponse ListDashboardsResponse
	path := "/api/2.0/lakeview/dashboards"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listDashboardsResponse)
	return &listDashboardsResponse, err
}

func (a *lakeviewImpl) ListSchedules(ctx context.Context, request ListSchedulesRequest) (*ListSchedulesResponse, error) {
	var listSchedulesResponse ListSchedulesResponse
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules", request.DashboardId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listSchedulesResponse)
	return &listSchedulesResponse, err
}

func (a *lakeviewImpl) ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) (*ListSubscriptionsResponse, error) {
	var listSubscriptionsResponse ListSubscriptionsResponse
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions", request.DashboardId, request.ScheduleId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listSubscriptionsResponse)
	return &listSubscriptionsResponse, err
}

func (a *lakeviewImpl) Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := "/api/2.0/lakeview/dashboards/migrate"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &dashboard)
	return &dashboard, err
}

func (a *lakeviewImpl) Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error) {
	var publishedDashboard PublishedDashboard
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", request.DashboardId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &publishedDashboard)
	return &publishedDashboard, err
}

func (a *lakeviewImpl) Trash(ctx context.Context, request TrashDashboardRequest) error {
	var trashDashboardResponse TrashDashboardResponse
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v", request.DashboardId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &trashDashboardResponse)
	return err
}

func (a *lakeviewImpl) Unpublish(ctx context.Context, request UnpublishDashboardRequest) error {
	var unpublishDashboardResponse UnpublishDashboardResponse
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", request.DashboardId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &unpublishDashboardResponse)
	return err
}

func (a *lakeviewImpl) Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v", request.DashboardId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &dashboard)
	return &dashboard, err
}

func (a *lakeviewImpl) UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error) {
	var schedule Schedule
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v", request.DashboardId, request.ScheduleId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &schedule)
	return &schedule, err
}

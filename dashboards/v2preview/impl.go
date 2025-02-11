// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboardspreview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

// unexported type that holds implementations of just GeniePreview API methods
type geniePreviewImpl struct {
	client *client.DatabricksClient
}

func (a *geniePreviewImpl) CreateMessage(ctx context.Context, request GenieCreateConversationMessageRequest) (*GenieMessage, error) {
	var genieMessage GenieMessage
	path := fmt.Sprintf("/api/2.0preview/genie/spaces/%v/conversations/%v/messages", request.SpaceId, request.ConversationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &genieMessage)
	return &genieMessage, err
}

func (a *geniePreviewImpl) ExecuteMessageQuery(ctx context.Context, request GenieExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error) {
	var genieGetMessageQueryResultResponse GenieGetMessageQueryResultResponse
	path := fmt.Sprintf("/api/2.0preview/genie/spaces/%v/conversations/%v/messages/%v/execute-query", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &genieGetMessageQueryResultResponse)
	return &genieGetMessageQueryResultResponse, err
}

func (a *geniePreviewImpl) GetMessage(ctx context.Context, request GenieGetConversationMessageRequest) (*GenieMessage, error) {
	var genieMessage GenieMessage
	path := fmt.Sprintf("/api/2.0preview/genie/spaces/%v/conversations/%v/messages/%v", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieMessage)
	return &genieMessage, err
}

func (a *geniePreviewImpl) GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error) {
	var genieGetMessageQueryResultResponse GenieGetMessageQueryResultResponse
	path := fmt.Sprintf("/api/2.0preview/genie/spaces/%v/conversations/%v/messages/%v/query-result", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieGetMessageQueryResultResponse)
	return &genieGetMessageQueryResultResponse, err
}

func (a *geniePreviewImpl) GetMessageQueryResultByAttachment(ctx context.Context, request GenieGetQueryResultByAttachmentRequest) (*GenieGetMessageQueryResultResponse, error) {
	var genieGetMessageQueryResultResponse GenieGetMessageQueryResultResponse
	path := fmt.Sprintf("/api/2.0preview/genie/spaces/%v/conversations/%v/messages/%v/query-result/%v", request.SpaceId, request.ConversationId, request.MessageId, request.AttachmentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieGetMessageQueryResultResponse)
	return &genieGetMessageQueryResultResponse, err
}

func (a *geniePreviewImpl) StartConversation(ctx context.Context, request GenieStartConversationMessageRequest) (*GenieStartConversationResponse, error) {
	var genieStartConversationResponse GenieStartConversationResponse
	path := fmt.Sprintf("/api/2.0preview/genie/spaces/%v/start-conversation", request.SpaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &genieStartConversationResponse)
	return &genieStartConversationResponse, err
}

// unexported type that holds implementations of just LakeviewEmbeddedPreview API methods
type lakeviewEmbeddedPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *lakeviewEmbeddedPreviewImpl) GetPublishedDashboardEmbedded(ctx context.Context, request GetPublishedDashboardEmbeddedRequest) error {
	var getPublishedDashboardEmbeddedResponse GetPublishedDashboardEmbeddedResponse
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/published/embedded", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPublishedDashboardEmbeddedResponse)
	return err
}

// unexported type that holds implementations of just LakeviewPreview API methods
type lakeviewPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *lakeviewPreviewImpl) Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := "/api/2.0preview/lakeview/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Dashboard, &dashboard)
	return &dashboard, err
}

func (a *lakeviewPreviewImpl) CreateSchedule(ctx context.Context, request CreateScheduleRequest) (*Schedule, error) {
	var schedule Schedule
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/schedules", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Schedule, &schedule)
	return &schedule, err
}

func (a *lakeviewPreviewImpl) CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (*Subscription, error) {
	var subscription Subscription
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/schedules/%v/subscriptions", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Subscription, &subscription)
	return &subscription, err
}

func (a *lakeviewPreviewImpl) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error {
	var deleteScheduleResponse DeleteScheduleResponse
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/schedules/%v", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteScheduleResponse)
	return err
}

func (a *lakeviewPreviewImpl) DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) error {
	var deleteSubscriptionResponse DeleteSubscriptionResponse
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/schedules/%v/subscriptions/%v", request.DashboardId, request.ScheduleId, request.SubscriptionId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteSubscriptionResponse)
	return err
}

func (a *lakeviewPreviewImpl) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &dashboard)
	return &dashboard, err
}

func (a *lakeviewPreviewImpl) GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error) {
	var publishedDashboard PublishedDashboard
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/published", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &publishedDashboard)
	return &publishedDashboard, err
}

func (a *lakeviewPreviewImpl) GetSchedule(ctx context.Context, request GetScheduleRequest) (*Schedule, error) {
	var schedule Schedule
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/schedules/%v", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &schedule)
	return &schedule, err
}

func (a *lakeviewPreviewImpl) GetSubscription(ctx context.Context, request GetSubscriptionRequest) (*Subscription, error) {
	var subscription Subscription
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/schedules/%v/subscriptions/%v", request.DashboardId, request.ScheduleId, request.SubscriptionId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &subscription)
	return &subscription, err
}

// List dashboards.
func (a *lakeviewPreviewImpl) List(ctx context.Context, request ListDashboardsRequest) listing.Iterator[Dashboard] {

	getNextPage := func(ctx context.Context, req ListDashboardsRequest) (*ListDashboardsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListDashboardsResponse) []Dashboard {
		return resp.Dashboards
	}
	getNextReq := func(resp *ListDashboardsResponse) *ListDashboardsRequest {
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

// List dashboards.
func (a *lakeviewPreviewImpl) ListAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Dashboard](ctx, iterator)
}
func (a *lakeviewPreviewImpl) internalList(ctx context.Context, request ListDashboardsRequest) (*ListDashboardsResponse, error) {
	var listDashboardsResponse ListDashboardsResponse
	path := "/api/2.0preview/lakeview/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDashboardsResponse)
	return &listDashboardsResponse, err
}

// List dashboard schedules.
func (a *lakeviewPreviewImpl) ListSchedules(ctx context.Context, request ListSchedulesRequest) listing.Iterator[Schedule] {

	getNextPage := func(ctx context.Context, req ListSchedulesRequest) (*ListSchedulesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSchedules(ctx, req)
	}
	getItems := func(resp *ListSchedulesResponse) []Schedule {
		return resp.Schedules
	}
	getNextReq := func(resp *ListSchedulesResponse) *ListSchedulesRequest {
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

// List dashboard schedules.
func (a *lakeviewPreviewImpl) ListSchedulesAll(ctx context.Context, request ListSchedulesRequest) ([]Schedule, error) {
	iterator := a.ListSchedules(ctx, request)
	return listing.ToSlice[Schedule](ctx, iterator)
}
func (a *lakeviewPreviewImpl) internalListSchedules(ctx context.Context, request ListSchedulesRequest) (*ListSchedulesResponse, error) {
	var listSchedulesResponse ListSchedulesResponse
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/schedules", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSchedulesResponse)
	return &listSchedulesResponse, err
}

// List schedule subscriptions.
func (a *lakeviewPreviewImpl) ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) listing.Iterator[Subscription] {

	getNextPage := func(ctx context.Context, req ListSubscriptionsRequest) (*ListSubscriptionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSubscriptions(ctx, req)
	}
	getItems := func(resp *ListSubscriptionsResponse) []Subscription {
		return resp.Subscriptions
	}
	getNextReq := func(resp *ListSubscriptionsResponse) *ListSubscriptionsRequest {
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

// List schedule subscriptions.
func (a *lakeviewPreviewImpl) ListSubscriptionsAll(ctx context.Context, request ListSubscriptionsRequest) ([]Subscription, error) {
	iterator := a.ListSubscriptions(ctx, request)
	return listing.ToSlice[Subscription](ctx, iterator)
}
func (a *lakeviewPreviewImpl) internalListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) (*ListSubscriptionsResponse, error) {
	var listSubscriptionsResponse ListSubscriptionsResponse
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/schedules/%v/subscriptions", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSubscriptionsResponse)
	return &listSubscriptionsResponse, err
}

func (a *lakeviewPreviewImpl) Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := "/api/2.0preview/lakeview/dashboards/migrate"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &dashboard)
	return &dashboard, err
}

func (a *lakeviewPreviewImpl) Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error) {
	var publishedDashboard PublishedDashboard
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/published", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &publishedDashboard)
	return &publishedDashboard, err
}

func (a *lakeviewPreviewImpl) Trash(ctx context.Context, request TrashDashboardRequest) error {
	var trashDashboardResponse TrashDashboardResponse
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &trashDashboardResponse)
	return err
}

func (a *lakeviewPreviewImpl) Unpublish(ctx context.Context, request UnpublishDashboardRequest) error {
	var unpublishDashboardResponse UnpublishDashboardResponse
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/published", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &unpublishDashboardResponse)
	return err
}

func (a *lakeviewPreviewImpl) Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Dashboard, &dashboard)
	return &dashboard, err
}

func (a *lakeviewPreviewImpl) UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error) {
	var schedule Schedule
	path := fmt.Sprintf("/api/2.0preview/lakeview/dashboards/%v/schedules/%v", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request.Schedule, &schedule)
	return &schedule, err
}

// unexported type that holds implementations of just QueryExecutionPreview API methods
type queryExecutionPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *queryExecutionPreviewImpl) CancelPublishedQueryExecution(ctx context.Context, request CancelPublishedQueryExecutionRequest) (*CancelQueryExecutionResponse, error) {
	var cancelQueryExecutionResponse CancelQueryExecutionResponse
	path := "/api/2.0preview/lakeview-query/query/published"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &cancelQueryExecutionResponse)
	return &cancelQueryExecutionResponse, err
}

func (a *queryExecutionPreviewImpl) ExecutePublishedDashboardQuery(ctx context.Context, request ExecutePublishedDashboardQueryRequest) error {
	var executeQueryResponse ExecuteQueryResponse
	path := "/api/2.0preview/lakeview-query/query/published"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &executeQueryResponse)
	return err
}

func (a *queryExecutionPreviewImpl) PollPublishedQueryStatus(ctx context.Context, request PollPublishedQueryStatusRequest) (*PollQueryStatusResponse, error) {
	var pollQueryStatusResponse PollQueryStatusResponse
	path := "/api/2.0preview/lakeview-query/query/published"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &pollQueryStatusResponse)
	return &pollQueryStatusResponse, err
}

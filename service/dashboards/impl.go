// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Genie API methods
type genieImpl struct {
	client *client.DatabricksClient
}

func (a *genieImpl) CreateMessage(ctx context.Context, request GenieCreateConversationMessageRequest) (*GenieMessage, error) {

	requestPb, pbErr := genieCreateConversationMessageRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieMessagePb genieMessagePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages", requestPb.SpaceId, requestPb.ConversationId)
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
		&genieMessagePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := genieMessageFromPb(&genieMessagePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) DeleteConversation(ctx context.Context, request GenieDeleteConversationRequest) error {

	requestPb, pbErr := genieDeleteConversationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v", requestPb.SpaceId, requestPb.ConversationId)
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
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *genieImpl) ExecuteMessageAttachmentQuery(ctx context.Context, request GenieExecuteMessageAttachmentQueryRequest) (*GenieGetMessageQueryResultResponse, error) {

	requestPb, pbErr := genieExecuteMessageAttachmentQueryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieGetMessageQueryResultResponsePb genieGetMessageQueryResultResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/attachments/%v/execute-query", requestPb.SpaceId, requestPb.ConversationId, requestPb.MessageId, requestPb.AttachmentId)
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
		&genieGetMessageQueryResultResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := genieGetMessageQueryResultResponseFromPb(&genieGetMessageQueryResultResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) ExecuteMessageQuery(ctx context.Context, request GenieExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error) {

	requestPb, pbErr := genieExecuteMessageQueryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieGetMessageQueryResultResponsePb genieGetMessageQueryResultResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/execute-query", requestPb.SpaceId, requestPb.ConversationId, requestPb.MessageId)
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
		&genieGetMessageQueryResultResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := genieGetMessageQueryResultResponseFromPb(&genieGetMessageQueryResultResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) GetMessage(ctx context.Context, request GenieGetConversationMessageRequest) (*GenieMessage, error) {

	requestPb, pbErr := genieGetConversationMessageRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieMessagePb genieMessagePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v", requestPb.SpaceId, requestPb.ConversationId, requestPb.MessageId)
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
		&genieMessagePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := genieMessageFromPb(&genieMessagePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) GetMessageAttachmentQueryResult(ctx context.Context, request GenieGetMessageAttachmentQueryResultRequest) (*GenieGetMessageQueryResultResponse, error) {

	requestPb, pbErr := genieGetMessageAttachmentQueryResultRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieGetMessageQueryResultResponsePb genieGetMessageQueryResultResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/attachments/%v/query-result", requestPb.SpaceId, requestPb.ConversationId, requestPb.MessageId, requestPb.AttachmentId)
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
		&genieGetMessageQueryResultResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := genieGetMessageQueryResultResponseFromPb(&genieGetMessageQueryResultResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error) {

	requestPb, pbErr := genieGetMessageQueryResultRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieGetMessageQueryResultResponsePb genieGetMessageQueryResultResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/query-result", requestPb.SpaceId, requestPb.ConversationId, requestPb.MessageId)
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
		&genieGetMessageQueryResultResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := genieGetMessageQueryResultResponseFromPb(&genieGetMessageQueryResultResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) GetMessageQueryResultByAttachment(ctx context.Context, request GenieGetQueryResultByAttachmentRequest) (*GenieGetMessageQueryResultResponse, error) {

	requestPb, pbErr := genieGetQueryResultByAttachmentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieGetMessageQueryResultResponsePb genieGetMessageQueryResultResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/query-result/%v", requestPb.SpaceId, requestPb.ConversationId, requestPb.MessageId, requestPb.AttachmentId)
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
		&genieGetMessageQueryResultResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := genieGetMessageQueryResultResponseFromPb(&genieGetMessageQueryResultResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) GetSpace(ctx context.Context, request GenieGetSpaceRequest) (*GenieSpace, error) {

	requestPb, pbErr := genieGetSpaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieSpacePb genieSpacePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v", requestPb.SpaceId)
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
		&genieSpacePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := genieSpaceFromPb(&genieSpacePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) ListConversations(ctx context.Context, request GenieListConversationsRequest) (*GenieListConversationsResponse, error) {

	requestPb, pbErr := genieListConversationsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieListConversationsResponsePb genieListConversationsResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations", requestPb.SpaceId)
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
		&genieListConversationsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := genieListConversationsResponseFromPb(&genieListConversationsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) ListSpaces(ctx context.Context, request GenieListSpacesRequest) (*GenieListSpacesResponse, error) {

	requestPb, pbErr := genieListSpacesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieListSpacesResponsePb genieListSpacesResponsePb
	path := "/api/2.0/genie/spaces"
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
		&genieListSpacesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := genieListSpacesResponseFromPb(&genieListSpacesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) StartConversation(ctx context.Context, request GenieStartConversationMessageRequest) (*GenieStartConversationResponse, error) {

	requestPb, pbErr := genieStartConversationMessageRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieStartConversationResponsePb genieStartConversationResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/start-conversation", requestPb.SpaceId)
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
		&genieStartConversationResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := genieStartConversationResponseFromPb(&genieStartConversationResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) TrashSpace(ctx context.Context, request GenieTrashSpaceRequest) error {

	requestPb, pbErr := genieTrashSpaceRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/genie/spaces/%v", requestPb.SpaceId)
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
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just Lakeview API methods
type lakeviewImpl struct {
	client *client.DatabricksClient
}

func (a *lakeviewImpl) Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error) {

	requestPb, pbErr := createDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb dashboardPb
	path := "/api/2.0/lakeview/dashboards"
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
		(*requestPb).Dashboard,
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

func (a *lakeviewImpl) CreateSchedule(ctx context.Context, request CreateScheduleRequest) (*Schedule, error) {

	requestPb, pbErr := createScheduleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var schedulePb schedulePb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules", requestPb.DashboardId)
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
		(*requestPb).Schedule,
		&schedulePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := scheduleFromPb(&schedulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (*Subscription, error) {

	requestPb, pbErr := createSubscriptionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var subscriptionPb subscriptionPb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions", requestPb.DashboardId, requestPb.ScheduleId)
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
		(*requestPb).Subscription,
		&subscriptionPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := subscriptionFromPb(&subscriptionPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error {

	requestPb, pbErr := deleteScheduleRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v", requestPb.DashboardId, requestPb.ScheduleId)
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
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *lakeviewImpl) DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) error {

	requestPb, pbErr := deleteSubscriptionRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions/%v", requestPb.DashboardId, requestPb.ScheduleId, requestPb.SubscriptionId)
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
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *lakeviewImpl) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {

	requestPb, pbErr := getDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb dashboardPb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v", requestPb.DashboardId)
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

func (a *lakeviewImpl) GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error) {

	requestPb, pbErr := getPublishedDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var publishedDashboardPb publishedDashboardPb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", requestPb.DashboardId)
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
		&publishedDashboardPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := publishedDashboardFromPb(&publishedDashboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) GetSchedule(ctx context.Context, request GetScheduleRequest) (*Schedule, error) {

	requestPb, pbErr := getScheduleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var schedulePb schedulePb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v", requestPb.DashboardId, requestPb.ScheduleId)
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
		&schedulePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := scheduleFromPb(&schedulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) GetSubscription(ctx context.Context, request GetSubscriptionRequest) (*Subscription, error) {

	requestPb, pbErr := getSubscriptionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var subscriptionPb subscriptionPb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions/%v", requestPb.DashboardId, requestPb.ScheduleId, requestPb.SubscriptionId)
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
		&subscriptionPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := subscriptionFromPb(&subscriptionPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List dashboards.
func (a *lakeviewImpl) List(ctx context.Context, request ListDashboardsRequest) listing.Iterator[Dashboard] {

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
func (a *lakeviewImpl) ListAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Dashboard](ctx, iterator)
}

func (a *lakeviewImpl) internalList(ctx context.Context, request ListDashboardsRequest) (*ListDashboardsResponse, error) {

	requestPb, pbErr := listDashboardsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listDashboardsResponsePb listDashboardsResponsePb
	path := "/api/2.0/lakeview/dashboards"
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
		&listDashboardsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listDashboardsResponseFromPb(&listDashboardsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List dashboard schedules.
func (a *lakeviewImpl) ListSchedules(ctx context.Context, request ListSchedulesRequest) listing.Iterator[Schedule] {

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
func (a *lakeviewImpl) ListSchedulesAll(ctx context.Context, request ListSchedulesRequest) ([]Schedule, error) {
	iterator := a.ListSchedules(ctx, request)
	return listing.ToSlice[Schedule](ctx, iterator)
}

func (a *lakeviewImpl) internalListSchedules(ctx context.Context, request ListSchedulesRequest) (*ListSchedulesResponse, error) {

	requestPb, pbErr := listSchedulesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listSchedulesResponsePb listSchedulesResponsePb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules", requestPb.DashboardId)
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
		&listSchedulesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listSchedulesResponseFromPb(&listSchedulesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List schedule subscriptions.
func (a *lakeviewImpl) ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) listing.Iterator[Subscription] {

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
func (a *lakeviewImpl) ListSubscriptionsAll(ctx context.Context, request ListSubscriptionsRequest) ([]Subscription, error) {
	iterator := a.ListSubscriptions(ctx, request)
	return listing.ToSlice[Subscription](ctx, iterator)
}

func (a *lakeviewImpl) internalListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) (*ListSubscriptionsResponse, error) {

	requestPb, pbErr := listSubscriptionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listSubscriptionsResponsePb listSubscriptionsResponsePb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions", requestPb.DashboardId, requestPb.ScheduleId)
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
		&listSubscriptionsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listSubscriptionsResponseFromPb(&listSubscriptionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error) {

	requestPb, pbErr := migrateDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb dashboardPb
	path := "/api/2.0/lakeview/dashboards/migrate"
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

func (a *lakeviewImpl) Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error) {

	requestPb, pbErr := publishRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var publishedDashboardPb publishedDashboardPb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", requestPb.DashboardId)
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
		&publishedDashboardPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := publishedDashboardFromPb(&publishedDashboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) Trash(ctx context.Context, request TrashDashboardRequest) error {

	requestPb, pbErr := trashDashboardRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v", requestPb.DashboardId)
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
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *lakeviewImpl) Unpublish(ctx context.Context, request UnpublishDashboardRequest) error {

	requestPb, pbErr := unpublishDashboardRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", requestPb.DashboardId)
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
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *lakeviewImpl) Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error) {

	requestPb, pbErr := updateDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb dashboardPb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v", requestPb.DashboardId)
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
		(*requestPb).Dashboard,
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

func (a *lakeviewImpl) UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error) {

	requestPb, pbErr := updateScheduleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var schedulePb schedulePb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v", requestPb.DashboardId, requestPb.ScheduleId)
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
		(*requestPb).Schedule,
		&schedulePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := scheduleFromPb(&schedulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just LakeviewEmbedded API methods
type lakeviewEmbeddedImpl struct {
	client *client.DatabricksClient
}

func (a *lakeviewEmbeddedImpl) GetPublishedDashboardTokenInfo(ctx context.Context, request GetPublishedDashboardTokenInfoRequest) (*GetPublishedDashboardTokenInfoResponse, error) {

	requestPb, pbErr := getPublishedDashboardTokenInfoRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPublishedDashboardTokenInfoResponsePb getPublishedDashboardTokenInfoResponsePb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published/tokeninfo", requestPb.DashboardId)
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
		&getPublishedDashboardTokenInfoResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getPublishedDashboardTokenInfoResponseFromPb(&getPublishedDashboardTokenInfoResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/dashboards/dashboardspb"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Genie API methods
type genieImpl struct {
	client *client.DatabricksClient
}

func (a *genieImpl) CreateMessage(ctx context.Context, request GenieCreateConversationMessageRequest) (*GenieMessage, error) {
	requestPb, pbErr := GenieCreateConversationMessageRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieMessagePb dashboardspb.GenieMessagePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages", request.SpaceId, request.ConversationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&genieMessagePb,
	)
	resp, err := GenieMessageFromPb(&genieMessagePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) DeleteConversation(ctx context.Context, request GenieDeleteConversationRequest) error {
	requestPb, pbErr := GenieDeleteConversationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v", request.SpaceId, request.ConversationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *genieImpl) ExecuteMessageAttachmentQuery(ctx context.Context, request GenieExecuteMessageAttachmentQueryRequest) (*GenieGetMessageQueryResultResponse, error) {

	var genieGetMessageQueryResultResponsePb dashboardspb.GenieGetMessageQueryResultResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/attachments/%v/execute-query", request.SpaceId, request.ConversationId, request.MessageId, request.AttachmentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		&genieGetMessageQueryResultResponsePb,
	)
	resp, err := GenieGetMessageQueryResultResponseFromPb(&genieGetMessageQueryResultResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) ExecuteMessageQuery(ctx context.Context, request GenieExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error) {

	var genieGetMessageQueryResultResponsePb dashboardspb.GenieGetMessageQueryResultResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/execute-query", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		&genieGetMessageQueryResultResponsePb,
	)
	resp, err := GenieGetMessageQueryResultResponseFromPb(&genieGetMessageQueryResultResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) GetMessage(ctx context.Context, request GenieGetConversationMessageRequest) (*GenieMessage, error) {
	requestPb, pbErr := GenieGetConversationMessageRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieMessagePb dashboardspb.GenieMessagePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&genieMessagePb,
	)
	resp, err := GenieMessageFromPb(&genieMessagePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) GetMessageAttachmentQueryResult(ctx context.Context, request GenieGetMessageAttachmentQueryResultRequest) (*GenieGetMessageQueryResultResponse, error) {
	requestPb, pbErr := GenieGetMessageAttachmentQueryResultRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieGetMessageQueryResultResponsePb dashboardspb.GenieGetMessageQueryResultResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/attachments/%v/query-result", request.SpaceId, request.ConversationId, request.MessageId, request.AttachmentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&genieGetMessageQueryResultResponsePb,
	)
	resp, err := GenieGetMessageQueryResultResponseFromPb(&genieGetMessageQueryResultResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error) {
	requestPb, pbErr := GenieGetMessageQueryResultRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieGetMessageQueryResultResponsePb dashboardspb.GenieGetMessageQueryResultResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/query-result", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&genieGetMessageQueryResultResponsePb,
	)
	resp, err := GenieGetMessageQueryResultResponseFromPb(&genieGetMessageQueryResultResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) GetMessageQueryResultByAttachment(ctx context.Context, request GenieGetQueryResultByAttachmentRequest) (*GenieGetMessageQueryResultResponse, error) {
	requestPb, pbErr := GenieGetQueryResultByAttachmentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieGetMessageQueryResultResponsePb dashboardspb.GenieGetMessageQueryResultResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/query-result/%v", request.SpaceId, request.ConversationId, request.MessageId, request.AttachmentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&genieGetMessageQueryResultResponsePb,
	)
	resp, err := GenieGetMessageQueryResultResponseFromPb(&genieGetMessageQueryResultResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) GetSpace(ctx context.Context, request GenieGetSpaceRequest) (*GenieSpace, error) {
	requestPb, pbErr := GenieGetSpaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieSpacePb dashboardspb.GenieSpacePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v", request.SpaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&genieSpacePb,
	)
	resp, err := GenieSpaceFromPb(&genieSpacePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) ListConversations(ctx context.Context, request GenieListConversationsRequest) (*GenieListConversationsResponse, error) {
	requestPb, pbErr := GenieListConversationsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieListConversationsResponsePb dashboardspb.GenieListConversationsResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations", request.SpaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&genieListConversationsResponsePb,
	)
	resp, err := GenieListConversationsResponseFromPb(&genieListConversationsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) ListSpaces(ctx context.Context, request GenieListSpacesRequest) (*GenieListSpacesResponse, error) {
	requestPb, pbErr := GenieListSpacesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieListSpacesResponsePb dashboardspb.GenieListSpacesResponsePb
	path := "/api/2.0/genie/spaces"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&genieListSpacesResponsePb,
	)
	resp, err := GenieListSpacesResponseFromPb(&genieListSpacesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) StartConversation(ctx context.Context, request GenieStartConversationMessageRequest) (*GenieStartConversationResponse, error) {
	requestPb, pbErr := GenieStartConversationMessageRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var genieStartConversationResponsePb dashboardspb.GenieStartConversationResponsePb
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/start-conversation", request.SpaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&genieStartConversationResponsePb,
	)
	resp, err := GenieStartConversationResponseFromPb(&genieStartConversationResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *genieImpl) TrashSpace(ctx context.Context, request GenieTrashSpaceRequest) error {
	requestPb, pbErr := GenieTrashSpaceRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/genie/spaces/%v", request.SpaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just Lakeview API methods
type lakeviewImpl struct {
	client *client.DatabricksClient
}

func (a *lakeviewImpl) Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error) {
	requestPb, pbErr := CreateDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb dashboardspb.DashboardPb
	path := "/api/2.0/lakeview/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.Dashboard,
		&dashboardPb,
	)
	resp, err := DashboardFromPb(&dashboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) CreateSchedule(ctx context.Context, request CreateScheduleRequest) (*Schedule, error) {
	requestPb, pbErr := CreateScheduleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var schedulePb dashboardspb.SchedulePb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.Schedule,
		&schedulePb,
	)
	resp, err := ScheduleFromPb(&schedulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (*Subscription, error) {
	requestPb, pbErr := CreateSubscriptionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var subscriptionPb dashboardspb.SubscriptionPb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.Subscription,
		&subscriptionPb,
	)
	resp, err := SubscriptionFromPb(&subscriptionPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error {
	requestPb, pbErr := DeleteScheduleRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *lakeviewImpl) DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) error {
	requestPb, pbErr := DeleteSubscriptionRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions/%v", request.DashboardId, request.ScheduleId, request.SubscriptionId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *lakeviewImpl) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	requestPb, pbErr := GetDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb dashboardspb.DashboardPb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&dashboardPb,
	)
	resp, err := DashboardFromPb(&dashboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error) {
	requestPb, pbErr := GetPublishedDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var publishedDashboardPb dashboardspb.PublishedDashboardPb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&publishedDashboardPb,
	)
	resp, err := PublishedDashboardFromPb(&publishedDashboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) GetSchedule(ctx context.Context, request GetScheduleRequest) (*Schedule, error) {
	requestPb, pbErr := GetScheduleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var schedulePb dashboardspb.SchedulePb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&schedulePb,
	)
	resp, err := ScheduleFromPb(&schedulePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) GetSubscription(ctx context.Context, request GetSubscriptionRequest) (*Subscription, error) {
	requestPb, pbErr := GetSubscriptionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var subscriptionPb dashboardspb.SubscriptionPb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions/%v", request.DashboardId, request.ScheduleId, request.SubscriptionId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&subscriptionPb,
	)
	resp, err := SubscriptionFromPb(&subscriptionPb)
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
	requestPb, pbErr := ListDashboardsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listDashboardsResponsePb dashboardspb.ListDashboardsResponsePb
	path := "/api/2.0/lakeview/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listDashboardsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListDashboardsResponseFromPb(&listDashboardsResponsePb)
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
	requestPb, pbErr := ListSchedulesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listSchedulesResponsePb dashboardspb.ListSchedulesResponsePb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listSchedulesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListSchedulesResponseFromPb(&listSchedulesResponsePb)
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
	requestPb, pbErr := ListSubscriptionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listSubscriptionsResponsePb dashboardspb.ListSubscriptionsResponsePb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listSubscriptionsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListSubscriptionsResponseFromPb(&listSubscriptionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error) {
	requestPb, pbErr := MigrateDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb dashboardspb.DashboardPb
	path := "/api/2.0/lakeview/dashboards/migrate"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&dashboardPb,
	)
	resp, err := DashboardFromPb(&dashboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error) {
	requestPb, pbErr := PublishRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var publishedDashboardPb dashboardspb.PublishedDashboardPb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&publishedDashboardPb,
	)
	resp, err := PublishedDashboardFromPb(&publishedDashboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) Trash(ctx context.Context, request TrashDashboardRequest) error {
	requestPb, pbErr := TrashDashboardRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *lakeviewImpl) Unpublish(ctx context.Context, request UnpublishDashboardRequest) error {
	requestPb, pbErr := UnpublishDashboardRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *lakeviewImpl) Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error) {
	requestPb, pbErr := UpdateDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var dashboardPb dashboardspb.DashboardPb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb.Dashboard,
		&dashboardPb,
	)
	resp, err := DashboardFromPb(&dashboardPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *lakeviewImpl) UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error) {
	requestPb, pbErr := UpdateScheduleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var schedulePb dashboardspb.SchedulePb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb.Schedule,
		&schedulePb,
	)
	resp, err := ScheduleFromPb(&schedulePb)
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
	requestPb, pbErr := GetPublishedDashboardTokenInfoRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPublishedDashboardTokenInfoResponsePb dashboardspb.GetPublishedDashboardTokenInfoResponsePb
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published/tokeninfo", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getPublishedDashboardTokenInfoResponsePb,
	)
	resp, err := GetPublishedDashboardTokenInfoResponseFromPb(&getPublishedDashboardTokenInfoResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

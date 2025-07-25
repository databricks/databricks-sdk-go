// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Genie, Lakeview, Lakeview Embedded, etc.
package dashboards

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type GenieInterface interface {

	// WaitGetMessageGenieCompleted repeatedly calls [GenieAPI.GetMessage] and waits to reach COMPLETED state
	WaitGetMessageGenieCompleted(ctx context.Context, conversationId string, messageId string, spaceId string,
		timeout time.Duration, callback func(*GenieMessage)) (*GenieMessage, error)

	// Create new message in a [conversation](:method:genie/startconversation). The
	// AI response uses all previously created messages in the conversation to
	// respond.
	CreateMessage(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest) (*WaitGetMessageGenieCompleted[GenieMessage], error)

	// Calls [GenieAPIInterface.CreateMessage] and waits to reach COMPLETED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[GenieMessage](60*time.Minute) functional option.
	//
	// Deprecated: use [GenieAPIInterface.CreateMessage].Get() or [GenieAPIInterface.WaitGetMessageGenieCompleted]
	CreateMessageAndWait(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest, options ...retries.Option[GenieMessage]) (*GenieMessage, error)

	// Delete a conversation.
	DeleteConversation(ctx context.Context, request GenieDeleteConversationRequest) error

	// Delete a conversation.
	DeleteConversationBySpaceIdAndConversationId(ctx context.Context, spaceId string, conversationId string) error

	// Execute the SQL for a message query attachment. Use this API when the query
	// attachment has expired and needs to be re-executed.
	ExecuteMessageAttachmentQuery(ctx context.Context, request GenieExecuteMessageAttachmentQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// Execute the SQL query in the message.
	ExecuteMessageQuery(ctx context.Context, request GenieExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get message from conversation.
	GetMessage(ctx context.Context, request GenieGetConversationMessageRequest) (*GenieMessage, error)

	// Get message from conversation.
	GetMessageBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieMessage, error)

	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY` OR `COMPLETED`.
	GetMessageAttachmentQueryResult(ctx context.Context, request GenieGetMessageAttachmentQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY` OR `COMPLETED`.
	GetMessageAttachmentQueryResultBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error)

	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY`.
	GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY`.
	GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieGetMessageQueryResultResponse, error)

	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY` OR `COMPLETED`.
	GetMessageQueryResultByAttachment(ctx context.Context, request GenieGetQueryResultByAttachmentRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY` OR `COMPLETED`.
	GetMessageQueryResultByAttachmentBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error)

	// Get details of a Genie Space.
	GetSpace(ctx context.Context, request GenieGetSpaceRequest) (*GenieSpace, error)

	// Get details of a Genie Space.
	GetSpaceBySpaceId(ctx context.Context, spaceId string) (*GenieSpace, error)

	// Get a list of conversations in a Genie Space.
	ListConversations(ctx context.Context, request GenieListConversationsRequest) (*GenieListConversationsResponse, error)

	// Get a list of conversations in a Genie Space.
	ListConversationsBySpaceId(ctx context.Context, spaceId string) (*GenieListConversationsResponse, error)

	// Get list of Genie Spaces.
	ListSpaces(ctx context.Context, request GenieListSpacesRequest) (*GenieListSpacesResponse, error)

	// Start a new conversation.
	StartConversation(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest) (*WaitGetMessageGenieCompleted[GenieStartConversationResponse], error)

	// Calls [GenieAPIInterface.StartConversation] and waits to reach COMPLETED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[GenieMessage](60*time.Minute) functional option.
	//
	// Deprecated: use [GenieAPIInterface.StartConversation].Get() or [GenieAPIInterface.WaitGetMessageGenieCompleted]
	StartConversationAndWait(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest, options ...retries.Option[GenieMessage]) (*GenieMessage, error)

	// Move a Genie Space to the trash.
	TrashSpace(ctx context.Context, request GenieTrashSpaceRequest) error

	// Move a Genie Space to the trash.
	TrashSpaceBySpaceId(ctx context.Context, spaceId string) error
}

func NewGenie(client *client.DatabricksClient) *GenieAPI {
	return &GenieAPI{
		genieImpl: genieImpl{
			client: client,
		},
	}
}

// Genie provides a no-code experience for business users, powered by AI/BI.
// Analysts set up spaces that business users can use to ask questions using
// natural language. Genie uses data registered to Unity Catalog and requires at
// least CAN USE permission on a Pro or Serverless SQL warehouse. Also,
// Databricks Assistant must be enabled.
type GenieAPI struct {
	genieImpl
}

// WaitGetMessageGenieCompleted repeatedly calls [GenieAPI.GetMessage] and waits to reach COMPLETED state
func (a *GenieAPI) WaitGetMessageGenieCompleted(ctx context.Context, conversationId string, messageId string, spaceId string,
	timeout time.Duration, callback func(*GenieMessage)) (*GenieMessage, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[GenieMessage](ctx, timeout, func() (*GenieMessage, *retries.Err) {
		genieMessage, err := a.GetMessage(ctx, GenieGetConversationMessageRequest{
			ConversationId: conversationId,
			MessageId:      messageId,
			SpaceId:        spaceId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(genieMessage)
		}
		status := genieMessage.Status
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case MessageStatusCompleted: // target state
			return genieMessage, nil
		case MessageStatusFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				MessageStatusCompleted, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetMessageGenieCompleted is a wrapper that calls [GenieAPI.WaitGetMessageGenieCompleted] and waits to reach COMPLETED state.
type WaitGetMessageGenieCompleted[R any] struct {
	Response       *R
	ConversationId string `json:"conversation_id"`
	MessageId      string `json:"message_id"`
	SpaceId        string `json:"space_id"`
	Poll           func(time.Duration, func(*GenieMessage)) (*GenieMessage, error)
	callback       func(*GenieMessage)
	timeout        time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetMessageGenieCompleted[R]) OnProgress(callback func(*GenieMessage)) *WaitGetMessageGenieCompleted[R] {
	w.callback = callback
	return w
}

// Get the GenieMessage with the default timeout of 20 minutes.
func (w *WaitGetMessageGenieCompleted[R]) Get() (*GenieMessage, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the GenieMessage with custom timeout.
func (w *WaitGetMessageGenieCompleted[R]) GetWithTimeout(timeout time.Duration) (*GenieMessage, error) {
	return w.Poll(timeout, w.callback)
}

// Create new message in a [conversation](:method:genie/startconversation). The
// AI response uses all previously created messages in the conversation to
// respond.
func (a *GenieAPI) CreateMessage(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest) (*WaitGetMessageGenieCompleted[GenieMessage], error) {
	genieMessage, err := a.genieImpl.CreateMessage(ctx, genieCreateConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetMessageGenieCompleted[GenieMessage]{
		Response:       genieMessage,
		ConversationId: genieCreateConversationMessageRequest.ConversationId,
		MessageId:      genieMessage.MessageId,
		SpaceId:        genieCreateConversationMessageRequest.SpaceId,
		Poll: func(timeout time.Duration, callback func(*GenieMessage)) (*GenieMessage, error) {
			return a.WaitGetMessageGenieCompleted(ctx, genieCreateConversationMessageRequest.ConversationId, genieMessage.MessageId, genieCreateConversationMessageRequest.SpaceId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [GenieAPI.CreateMessage] and waits to reach COMPLETED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GenieMessage](60*time.Minute) functional option.
//
// Deprecated: use [GenieAPI.CreateMessage].Get() or [GenieAPI.WaitGetMessageGenieCompleted]
func (a *GenieAPI) CreateMessageAndWait(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest, options ...retries.Option[GenieMessage]) (*GenieMessage, error) {
	wait, err := a.CreateMessage(ctx, genieCreateConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[GenieMessage]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *GenieMessage) {
		for _, o := range options {
			o(&retries.Info[GenieMessage]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Delete a conversation.
func (a *GenieAPI) DeleteConversationBySpaceIdAndConversationId(ctx context.Context, spaceId string, conversationId string) error {
	return a.genieImpl.DeleteConversation(ctx, GenieDeleteConversationRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
	})
}

// Get message from conversation.
func (a *GenieAPI) GetMessageBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieMessage, error) {
	return a.genieImpl.GetMessage(ctx, GenieGetConversationMessageRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
	})
}

// Get the result of SQL query if the message has a query attachment. This is
// only available if a message has a query attachment and the message status is
// `EXECUTING_QUERY` OR `COMPLETED`.
func (a *GenieAPI) GetMessageAttachmentQueryResultBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error) {
	return a.genieImpl.GetMessageAttachmentQueryResult(ctx, GenieGetMessageAttachmentQueryResultRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
		AttachmentId:   attachmentId,
	})
}

// Get the result of SQL query if the message has a query attachment. This is
// only available if a message has a query attachment and the message status is
// `EXECUTING_QUERY`.
func (a *GenieAPI) GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieGetMessageQueryResultResponse, error) {
	return a.genieImpl.GetMessageQueryResult(ctx, GenieGetMessageQueryResultRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
	})
}

// Get the result of SQL query if the message has a query attachment. This is
// only available if a message has a query attachment and the message status is
// `EXECUTING_QUERY` OR `COMPLETED`.
func (a *GenieAPI) GetMessageQueryResultByAttachmentBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error) {
	return a.genieImpl.GetMessageQueryResultByAttachment(ctx, GenieGetQueryResultByAttachmentRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
		AttachmentId:   attachmentId,
	})
}

// Get details of a Genie Space.
func (a *GenieAPI) GetSpaceBySpaceId(ctx context.Context, spaceId string) (*GenieSpace, error) {
	return a.genieImpl.GetSpace(ctx, GenieGetSpaceRequest{
		SpaceId: spaceId,
	})
}

// Get a list of conversations in a Genie Space.
func (a *GenieAPI) ListConversationsBySpaceId(ctx context.Context, spaceId string) (*GenieListConversationsResponse, error) {
	return a.genieImpl.ListConversations(ctx, GenieListConversationsRequest{
		SpaceId: spaceId,
	})
}

// Start a new conversation.
func (a *GenieAPI) StartConversation(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest) (*WaitGetMessageGenieCompleted[GenieStartConversationResponse], error) {
	genieStartConversationResponse, err := a.genieImpl.StartConversation(ctx, genieStartConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetMessageGenieCompleted[GenieStartConversationResponse]{
		Response:       genieStartConversationResponse,
		ConversationId: genieStartConversationResponse.ConversationId,
		MessageId:      genieStartConversationResponse.MessageId,
		SpaceId:        genieStartConversationMessageRequest.SpaceId,
		Poll: func(timeout time.Duration, callback func(*GenieMessage)) (*GenieMessage, error) {
			return a.WaitGetMessageGenieCompleted(ctx, genieStartConversationResponse.ConversationId, genieStartConversationResponse.MessageId, genieStartConversationMessageRequest.SpaceId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [GenieAPI.StartConversation] and waits to reach COMPLETED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GenieMessage](60*time.Minute) functional option.
//
// Deprecated: use [GenieAPI.StartConversation].Get() or [GenieAPI.WaitGetMessageGenieCompleted]
func (a *GenieAPI) StartConversationAndWait(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest, options ...retries.Option[GenieMessage]) (*GenieMessage, error) {
	wait, err := a.StartConversation(ctx, genieStartConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[GenieMessage]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *GenieMessage) {
		for _, o := range options {
			o(&retries.Info[GenieMessage]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Move a Genie Space to the trash.
func (a *GenieAPI) TrashSpaceBySpaceId(ctx context.Context, spaceId string) error {
	return a.genieImpl.TrashSpace(ctx, GenieTrashSpaceRequest{
		SpaceId: spaceId,
	})
}

type LakeviewInterface interface {

	// Create a draft dashboard.
	Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error)

	// Create dashboard schedule.
	CreateSchedule(ctx context.Context, request CreateScheduleRequest) (*Schedule, error)

	// Create schedule subscription.
	CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (*Subscription, error)

	// Delete dashboard schedule.
	DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error

	// Delete dashboard schedule.
	DeleteScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) error

	// Delete schedule subscription.
	DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) error

	// Delete schedule subscription.
	DeleteSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) error

	// Get a draft dashboard.
	Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error)

	// Get a draft dashboard.
	GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error)

	// Get the current published dashboard.
	GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error)

	// Get the current published dashboard.
	GetPublishedByDashboardId(ctx context.Context, dashboardId string) (*PublishedDashboard, error)

	// Get dashboard schedule.
	GetSchedule(ctx context.Context, request GetScheduleRequest) (*Schedule, error)

	// Get dashboard schedule.
	GetScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*Schedule, error)

	// Get schedule subscription.
	GetSubscription(ctx context.Context, request GetSubscriptionRequest) (*Subscription, error)

	// Get schedule subscription.
	GetSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*Subscription, error)

	// List dashboards.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListDashboardsRequest) listing.Iterator[Dashboard]

	// List dashboards.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error)

	// List dashboard schedules.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSchedules(ctx context.Context, request ListSchedulesRequest) listing.Iterator[Schedule]

	// List dashboard schedules.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSchedulesAll(ctx context.Context, request ListSchedulesRequest) ([]Schedule, error)

	// List dashboard schedules.
	ListSchedulesByDashboardId(ctx context.Context, dashboardId string) (*ListSchedulesResponse, error)

	// List schedule subscriptions.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) listing.Iterator[Subscription]

	// List schedule subscriptions.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSubscriptionsAll(ctx context.Context, request ListSubscriptionsRequest) ([]Subscription, error)

	// List schedule subscriptions.
	ListSubscriptionsByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*ListSubscriptionsResponse, error)

	// Migrates a classic SQL dashboard to Lakeview.
	Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error)

	// Publish the current draft dashboard.
	Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error)

	// Trash a dashboard.
	Trash(ctx context.Context, request TrashDashboardRequest) error

	// Trash a dashboard.
	TrashByDashboardId(ctx context.Context, dashboardId string) error

	// Unpublish the dashboard.
	Unpublish(ctx context.Context, request UnpublishDashboardRequest) error

	// Unpublish the dashboard.
	UnpublishByDashboardId(ctx context.Context, dashboardId string) error

	// Update a draft dashboard.
	Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error)

	// Update dashboard schedule.
	UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error)
}

func NewLakeview(client *client.DatabricksClient) *LakeviewAPI {
	return &LakeviewAPI{
		lakeviewImpl: lakeviewImpl{
			client: client,
		},
	}
}

// These APIs provide specific management operations for Lakeview dashboards.
// Generic resource management can be done with Workspace API (import, export,
// get-status, list, delete).
type LakeviewAPI struct {
	lakeviewImpl
}

// Delete dashboard schedule.
func (a *LakeviewAPI) DeleteScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) error {
	return a.lakeviewImpl.DeleteSchedule(ctx, DeleteScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Delete schedule subscription.
func (a *LakeviewAPI) DeleteSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) error {
	return a.lakeviewImpl.DeleteSubscription(ctx, DeleteSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// Get a draft dashboard.
func (a *LakeviewAPI) GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error) {
	return a.lakeviewImpl.Get(ctx, GetDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get the current published dashboard.
func (a *LakeviewAPI) GetPublishedByDashboardId(ctx context.Context, dashboardId string) (*PublishedDashboard, error) {
	return a.lakeviewImpl.GetPublished(ctx, GetPublishedDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get dashboard schedule.
func (a *LakeviewAPI) GetScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*Schedule, error) {
	return a.lakeviewImpl.GetSchedule(ctx, GetScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Get schedule subscription.
func (a *LakeviewAPI) GetSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*Subscription, error) {
	return a.lakeviewImpl.GetSubscription(ctx, GetSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// List dashboard schedules.
func (a *LakeviewAPI) ListSchedulesByDashboardId(ctx context.Context, dashboardId string) (*ListSchedulesResponse, error) {
	return a.lakeviewImpl.internalListSchedules(ctx, ListSchedulesRequest{
		DashboardId: dashboardId,
	})
}

// List schedule subscriptions.
func (a *LakeviewAPI) ListSubscriptionsByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*ListSubscriptionsResponse, error) {
	return a.lakeviewImpl.internalListSubscriptions(ctx, ListSubscriptionsRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Trash a dashboard.
func (a *LakeviewAPI) TrashByDashboardId(ctx context.Context, dashboardId string) error {
	return a.lakeviewImpl.Trash(ctx, TrashDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Unpublish the dashboard.
func (a *LakeviewAPI) UnpublishByDashboardId(ctx context.Context, dashboardId string) error {
	return a.lakeviewImpl.Unpublish(ctx, UnpublishDashboardRequest{
		DashboardId: dashboardId,
	})
}

type LakeviewEmbeddedInterface interface {

	// Get a required authorization details and scopes of a published dashboard to
	// mint an OAuth token.
	GetPublishedDashboardTokenInfo(ctx context.Context, request GetPublishedDashboardTokenInfoRequest) (*GetPublishedDashboardTokenInfoResponse, error)

	// Get a required authorization details and scopes of a published dashboard to
	// mint an OAuth token.
	GetPublishedDashboardTokenInfoByDashboardId(ctx context.Context, dashboardId string) (*GetPublishedDashboardTokenInfoResponse, error)
}

func NewLakeviewEmbedded(client *client.DatabricksClient) *LakeviewEmbeddedAPI {
	return &LakeviewEmbeddedAPI{
		lakeviewEmbeddedImpl: lakeviewEmbeddedImpl{
			client: client,
		},
	}
}

// Token-based Lakeview APIs for embedding dashboards in external applications.
type LakeviewEmbeddedAPI struct {
	lakeviewEmbeddedImpl
}

// Get a required authorization details and scopes of a published dashboard to
// mint an OAuth token.
func (a *LakeviewEmbeddedAPI) GetPublishedDashboardTokenInfoByDashboardId(ctx context.Context, dashboardId string) (*GetPublishedDashboardTokenInfoResponse, error) {
	return a.lakeviewEmbeddedImpl.GetPublishedDashboardTokenInfo(ctx, GetPublishedDashboardTokenInfoRequest{
		DashboardId: dashboardId,
	})
}

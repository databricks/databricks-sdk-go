// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Genie, Lakeview, Lakeview Embedded, Query Execution, etc.
package dashboards

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type GenieInterface interface {
	// Create conversation message.
	//
	// Create new message in a [conversation](:method:genie/startconversation). The
	// AI response uses all previously created messages in the conversation to
	// respond.
	CreateMessage(ctx context.Context, request GenieCreateConversationMessageRequest) (*GenieCreateMessageWaiter, error)

	// Execute message attachment SQL query.
	//
	// Execute the SQL for a message query attachment. Use this API when the query
	// attachment has expired and needs to be re-executed.
	ExecuteMessageAttachmentQuery(ctx context.Context, request GenieExecuteMessageAttachmentQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// [Deprecated] Execute SQL query in a conversation message.
	//
	// Execute the SQL query in the message.
	ExecuteMessageQuery(ctx context.Context, request GenieExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get conversation message.
	//
	// Get message from conversation.
	GetMessage(ctx context.Context, request GenieGetConversationMessageRequest) (*GenieMessage, error)

	// Get conversation message.
	//
	// Get message from conversation.
	GetMessageBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieMessage, error)

	// Get message attachment SQL query result.
	//
	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY` OR `COMPLETED`.
	GetMessageAttachmentQueryResult(ctx context.Context, request GenieGetMessageAttachmentQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get message attachment SQL query result.
	//
	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY` OR `COMPLETED`.
	GetMessageAttachmentQueryResultBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error)

	// [Deprecated] Get conversation message SQL query result.
	//
	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY`.
	GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// [Deprecated] Get conversation message SQL query result.
	//
	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY`.
	GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieGetMessageQueryResultResponse, error)

	// [Deprecated] Get conversation message SQL query result.
	//
	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY` OR `COMPLETED`.
	GetMessageQueryResultByAttachment(ctx context.Context, request GenieGetQueryResultByAttachmentRequest) (*GenieGetMessageQueryResultResponse, error)

	// [Deprecated] Get conversation message SQL query result.
	//
	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY` OR `COMPLETED`.
	GetMessageQueryResultByAttachmentBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error)

	// Get Genie Space.
	//
	// Get details of a Genie Space.
	GetSpace(ctx context.Context, request GenieGetSpaceRequest) (*GenieSpace, error)

	// Get Genie Space.
	//
	// Get details of a Genie Space.
	GetSpaceBySpaceId(ctx context.Context, spaceId string) (*GenieSpace, error)

	// Start conversation.
	//
	// Start a new conversation.
	StartConversation(ctx context.Context, request GenieStartConversationMessageRequest) (*GenieStartConversationWaiter, error)
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

// Create conversation message.
//
// Create new message in a [conversation](:method:genie/startconversation). The
// AI response uses all previously created messages in the conversation to
// respond.
func (a *GenieAPI) CreateMessage(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest) (*GenieCreateMessageWaiter, error) {
	genieMessage, err := a.genieImpl.CreateMessage(ctx, genieCreateConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	return &GenieCreateMessageWaiter{
		Response:       genieMessage,
		conversationId: genieCreateConversationMessageRequest.ConversationId,
		messageId:      genieMessage.Id,
		spaceId:        genieCreateConversationMessageRequest.SpaceId,
		service:        a,
	}, nil
}

type GenieCreateMessageWaiter struct {
	// RawResponse is the raw response of the CreateMessage call.
	Response       *GenieMessage
	service        *GenieAPI
	conversationId string
	messageId      string
	spaceId        string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *GenieCreateMessageWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*GenieMessage, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[GenieMessage](ctx, opts.Timeout, func() (*GenieMessage, *retries.Err) {
		genieMessage, err := w.service.GetMessage(ctx, GenieGetConversationMessageRequest{
			ConversationId: w.conversationId,
			MessageId:      w.messageId,
			SpaceId:        w.spaceId,
		})
		if err != nil {
			return nil, retries.Halt(err)
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

// Get conversation message.
//
// Get message from conversation.
func (a *GenieAPI) GetMessageBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieMessage, error) {
	return a.genieImpl.GetMessage(ctx, GenieGetConversationMessageRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
	})
}

// Get message attachment SQL query result.
//
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

// [Deprecated] Get conversation message SQL query result.
//
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

// [Deprecated] Get conversation message SQL query result.
//
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

// Get Genie Space.
//
// Get details of a Genie Space.
func (a *GenieAPI) GetSpaceBySpaceId(ctx context.Context, spaceId string) (*GenieSpace, error) {
	return a.genieImpl.GetSpace(ctx, GenieGetSpaceRequest{
		SpaceId: spaceId,
	})
}

// Start conversation.
//
// Start a new conversation.
func (a *GenieAPI) StartConversation(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest) (*GenieStartConversationWaiter, error) {
	genieStartConversationResponse, err := a.genieImpl.StartConversation(ctx, genieStartConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	return &GenieStartConversationWaiter{
		Response:       genieStartConversationResponse,
		conversationId: genieStartConversationResponse.ConversationId,
		messageId:      genieStartConversationResponse.MessageId,
		spaceId:        genieStartConversationMessageRequest.SpaceId,
		service:        a,
	}, nil
}

type GenieStartConversationWaiter struct {
	// RawResponse is the raw response of the StartConversation call.
	Response       *GenieStartConversationResponse
	service        *GenieAPI
	conversationId string
	messageId      string
	spaceId        string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *GenieStartConversationWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*GenieMessage, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[GenieMessage](ctx, opts.Timeout, func() (*GenieMessage, *retries.Err) {
		genieMessage, err := w.service.GetMessage(ctx, GenieGetConversationMessageRequest{
			ConversationId: w.conversationId,
			MessageId:      w.messageId,
			SpaceId:        w.spaceId,
		})
		if err != nil {
			return nil, retries.Halt(err)
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

type LakeviewInterface interface {
	// Create dashboard.
	//
	// Create a draft dashboard.
	Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error)

	// Create dashboard schedule.
	CreateSchedule(ctx context.Context, request CreateScheduleRequest) (*Schedule, error)

	// Create schedule subscription.
	CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (*Subscription, error)

	// Delete dashboard schedule.
	DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) (*DeleteScheduleResponse, error)

	// Delete dashboard schedule.
	DeleteScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*DeleteScheduleResponse, error)

	// Delete schedule subscription.
	DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) (*DeleteSubscriptionResponse, error)

	// Delete schedule subscription.
	DeleteSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*DeleteSubscriptionResponse, error)

	// Get dashboard.
	//
	// Get a draft dashboard.
	Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error)

	// Get dashboard.
	//
	// Get a draft dashboard.
	GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error)

	// Get published dashboard.
	//
	// Get the current published dashboard.
	GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error)

	// Get published dashboard.
	//
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

	// Migrate dashboard.
	//
	// Migrates a classic SQL dashboard to Lakeview.
	Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error)

	// Publish dashboard.
	//
	// Publish the current draft dashboard.
	Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error)

	// Trash dashboard.
	//
	// Trash a dashboard.
	Trash(ctx context.Context, request TrashDashboardRequest) (*TrashDashboardResponse, error)

	// Trash dashboard.
	//
	// Trash a dashboard.
	TrashByDashboardId(ctx context.Context, dashboardId string) (*TrashDashboardResponse, error)

	// Unpublish dashboard.
	//
	// Unpublish the dashboard.
	Unpublish(ctx context.Context, request UnpublishDashboardRequest) (*UnpublishDashboardResponse, error)

	// Unpublish dashboard.
	//
	// Unpublish the dashboard.
	UnpublishByDashboardId(ctx context.Context, dashboardId string) (*UnpublishDashboardResponse, error)

	// Update dashboard.
	//
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
func (a *LakeviewAPI) DeleteScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*DeleteScheduleResponse, error) {
	return a.lakeviewImpl.DeleteSchedule(ctx, DeleteScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Delete schedule subscription.
func (a *LakeviewAPI) DeleteSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*DeleteSubscriptionResponse, error) {
	return a.lakeviewImpl.DeleteSubscription(ctx, DeleteSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// Get dashboard.
//
// Get a draft dashboard.
func (a *LakeviewAPI) GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error) {
	return a.lakeviewImpl.Get(ctx, GetDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get published dashboard.
//
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

// Trash dashboard.
//
// Trash a dashboard.
func (a *LakeviewAPI) TrashByDashboardId(ctx context.Context, dashboardId string) (*TrashDashboardResponse, error) {
	return a.lakeviewImpl.Trash(ctx, TrashDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Unpublish dashboard.
//
// Unpublish the dashboard.
func (a *LakeviewAPI) UnpublishByDashboardId(ctx context.Context, dashboardId string) (*UnpublishDashboardResponse, error) {
	return a.lakeviewImpl.Unpublish(ctx, UnpublishDashboardRequest{
		DashboardId: dashboardId,
	})
}

type LakeviewEmbeddedInterface interface {
	// Read a published dashboard in an embedded ui.
	//
	// Get the current published dashboard within an embedded context.
	GetPublishedDashboardEmbedded(ctx context.Context, request GetPublishedDashboardEmbeddedRequest) (*GetPublishedDashboardEmbeddedResponse, error)

	// Read a published dashboard in an embedded ui.
	//
	// Get the current published dashboard within an embedded context.
	GetPublishedDashboardEmbeddedByDashboardId(ctx context.Context, dashboardId string) (*GetPublishedDashboardEmbeddedResponse, error)
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

// Read a published dashboard in an embedded ui.
//
// Get the current published dashboard within an embedded context.
func (a *LakeviewEmbeddedAPI) GetPublishedDashboardEmbeddedByDashboardId(ctx context.Context, dashboardId string) (*GetPublishedDashboardEmbeddedResponse, error) {
	return a.lakeviewEmbeddedImpl.GetPublishedDashboardEmbedded(ctx, GetPublishedDashboardEmbeddedRequest{
		DashboardId: dashboardId,
	})
}

type QueryExecutionInterface interface {
	// Cancel the results for the a query for a published, embedded dashboard.
	CancelPublishedQueryExecution(ctx context.Context, request CancelPublishedQueryExecutionRequest) (*CancelQueryExecutionResponse, error)

	// Execute a query for a published dashboard.
	ExecutePublishedDashboardQuery(ctx context.Context, request ExecutePublishedDashboardQueryRequest) (*ExecuteQueryResponse, error)

	// Poll the results for the a query for a published, embedded dashboard.
	PollPublishedQueryStatus(ctx context.Context, request PollPublishedQueryStatusRequest) (*PollQueryStatusResponse, error)
}

func NewQueryExecution(client *client.DatabricksClient) *QueryExecutionAPI {
	return &QueryExecutionAPI{
		queryExecutionImpl: queryExecutionImpl{
			client: client,
		},
	}
}

// Query execution APIs for AI / BI Dashboards
type QueryExecutionAPI struct {
	queryExecutionImpl
}

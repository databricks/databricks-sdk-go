// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Genie Preview, Lakeview Embedded Preview, Lakeview Preview, Query Execution Preview, etc.
package dashboardspreview

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
)

type GeniePreviewInterface interface {

	// Create conversation message.
	//
	// Create new message in [conversation](:method:genie/startconversation). The AI
	// response uses all previously created messages in the conversation to respond.
	CreateMessage(ctx context.Context, request GenieCreateConversationMessageRequest) (*GenieMessage, error)

	// Execute SQL query in a conversation message.
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

	// Get conversation message SQL query result.
	//
	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY`.
	GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get conversation message SQL query result.
	//
	// Get the result of SQL query if the message has a query attachment. This is
	// only available if a message has a query attachment and the message status is
	// `EXECUTING_QUERY`.
	GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieGetMessageQueryResultResponse, error)

	// Get conversation message SQL query result by attachment id.
	//
	// Get the result of SQL query by attachment id This is only available if a
	// message has a query attachment and the message status is `EXECUTING_QUERY`.
	GetMessageQueryResultByAttachment(ctx context.Context, request GenieGetQueryResultByAttachmentRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get conversation message SQL query result by attachment id.
	//
	// Get the result of SQL query by attachment id This is only available if a
	// message has a query attachment and the message status is `EXECUTING_QUERY`.
	GetMessageQueryResultByAttachmentBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error)

	// Start conversation.
	//
	// Start a new conversation.
	StartConversation(ctx context.Context, request GenieStartConversationMessageRequest) (*GenieStartConversationResponse, error)
}

func NewGeniePreview(client *client.DatabricksClient) *GeniePreviewAPI {
	return &GeniePreviewAPI{
		geniePreviewImpl: geniePreviewImpl{
			client: client,
		},
	}
}

// Genie provides a no-code experience for business users, powered by AI/BI.
// Analysts set up spaces that business users can use to ask questions using
// natural language. Genie uses data registered to Unity Catalog and requires at
// least CAN USE permission on a Pro or Serverless SQL warehouse. Also,
// Databricks Assistant must be enabled.
type GeniePreviewAPI struct {
	geniePreviewImpl
}

// Get conversation message.
//
// Get message from conversation.
func (a *GeniePreviewAPI) GetMessageBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieMessage, error) {
	return a.geniePreviewImpl.GetMessage(ctx, GenieGetConversationMessageRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
	})
}

// Get conversation message SQL query result.
//
// Get the result of SQL query if the message has a query attachment. This is
// only available if a message has a query attachment and the message status is
// `EXECUTING_QUERY`.
func (a *GeniePreviewAPI) GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieGetMessageQueryResultResponse, error) {
	return a.geniePreviewImpl.GetMessageQueryResult(ctx, GenieGetMessageQueryResultRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
	})
}

// Get conversation message SQL query result by attachment id.
//
// Get the result of SQL query by attachment id This is only available if a
// message has a query attachment and the message status is `EXECUTING_QUERY`.
func (a *GeniePreviewAPI) GetMessageQueryResultByAttachmentBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error) {
	return a.geniePreviewImpl.GetMessageQueryResultByAttachment(ctx, GenieGetQueryResultByAttachmentRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
		AttachmentId:   attachmentId,
	})
}

type LakeviewEmbeddedPreviewInterface interface {

	// Read a published dashboard in an embedded ui.
	//
	// Get the current published dashboard within an embedded context.
	GetPublishedDashboardEmbedded(ctx context.Context, request GetPublishedDashboardEmbeddedRequest) error

	// Read a published dashboard in an embedded ui.
	//
	// Get the current published dashboard within an embedded context.
	GetPublishedDashboardEmbeddedByDashboardId(ctx context.Context, dashboardId string) error
}

func NewLakeviewEmbeddedPreview(client *client.DatabricksClient) *LakeviewEmbeddedPreviewAPI {
	return &LakeviewEmbeddedPreviewAPI{
		lakeviewEmbeddedPreviewImpl: lakeviewEmbeddedPreviewImpl{
			client: client,
		},
	}
}

// Token-based Lakeview APIs for embedding dashboards in external applications.
type LakeviewEmbeddedPreviewAPI struct {
	lakeviewEmbeddedPreviewImpl
}

// Read a published dashboard in an embedded ui.
//
// Get the current published dashboard within an embedded context.
func (a *LakeviewEmbeddedPreviewAPI) GetPublishedDashboardEmbeddedByDashboardId(ctx context.Context, dashboardId string) error {
	return a.lakeviewEmbeddedPreviewImpl.GetPublishedDashboardEmbedded(ctx, GetPublishedDashboardEmbeddedRequest{
		DashboardId: dashboardId,
	})
}

type LakeviewPreviewInterface interface {

	// Create dashboard.
	//
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
	Trash(ctx context.Context, request TrashDashboardRequest) error

	// Trash dashboard.
	//
	// Trash a dashboard.
	TrashByDashboardId(ctx context.Context, dashboardId string) error

	// Unpublish dashboard.
	//
	// Unpublish the dashboard.
	Unpublish(ctx context.Context, request UnpublishDashboardRequest) error

	// Unpublish dashboard.
	//
	// Unpublish the dashboard.
	UnpublishByDashboardId(ctx context.Context, dashboardId string) error

	// Update dashboard.
	//
	// Update a draft dashboard.
	Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error)

	// Update dashboard schedule.
	UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error)
}

func NewLakeviewPreview(client *client.DatabricksClient) *LakeviewPreviewAPI {
	return &LakeviewPreviewAPI{
		lakeviewPreviewImpl: lakeviewPreviewImpl{
			client: client,
		},
	}
}

// These APIs provide specific management operations for Lakeview dashboards.
// Generic resource management can be done with Workspace API (import, export,
// get-status, list, delete).
type LakeviewPreviewAPI struct {
	lakeviewPreviewImpl
}

// Delete dashboard schedule.
func (a *LakeviewPreviewAPI) DeleteScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) error {
	return a.lakeviewPreviewImpl.DeleteSchedule(ctx, DeleteScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Delete schedule subscription.
func (a *LakeviewPreviewAPI) DeleteSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) error {
	return a.lakeviewPreviewImpl.DeleteSubscription(ctx, DeleteSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// Get dashboard.
//
// Get a draft dashboard.
func (a *LakeviewPreviewAPI) GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error) {
	return a.lakeviewPreviewImpl.Get(ctx, GetDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get published dashboard.
//
// Get the current published dashboard.
func (a *LakeviewPreviewAPI) GetPublishedByDashboardId(ctx context.Context, dashboardId string) (*PublishedDashboard, error) {
	return a.lakeviewPreviewImpl.GetPublished(ctx, GetPublishedDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get dashboard schedule.
func (a *LakeviewPreviewAPI) GetScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*Schedule, error) {
	return a.lakeviewPreviewImpl.GetSchedule(ctx, GetScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Get schedule subscription.
func (a *LakeviewPreviewAPI) GetSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*Subscription, error) {
	return a.lakeviewPreviewImpl.GetSubscription(ctx, GetSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// List dashboard schedules.
func (a *LakeviewPreviewAPI) ListSchedulesByDashboardId(ctx context.Context, dashboardId string) (*ListSchedulesResponse, error) {
	return a.lakeviewPreviewImpl.internalListSchedules(ctx, ListSchedulesRequest{
		DashboardId: dashboardId,
	})
}

// List schedule subscriptions.
func (a *LakeviewPreviewAPI) ListSubscriptionsByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*ListSubscriptionsResponse, error) {
	return a.lakeviewPreviewImpl.internalListSubscriptions(ctx, ListSubscriptionsRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Trash dashboard.
//
// Trash a dashboard.
func (a *LakeviewPreviewAPI) TrashByDashboardId(ctx context.Context, dashboardId string) error {
	return a.lakeviewPreviewImpl.Trash(ctx, TrashDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Unpublish dashboard.
//
// Unpublish the dashboard.
func (a *LakeviewPreviewAPI) UnpublishByDashboardId(ctx context.Context, dashboardId string) error {
	return a.lakeviewPreviewImpl.Unpublish(ctx, UnpublishDashboardRequest{
		DashboardId: dashboardId,
	})
}

type QueryExecutionPreviewInterface interface {

	// Cancel the results for the a query for a published, embedded dashboard.
	CancelPublishedQueryExecution(ctx context.Context, request CancelPublishedQueryExecutionRequest) (*CancelQueryExecutionResponse, error)

	// Execute a query for a published dashboard.
	ExecutePublishedDashboardQuery(ctx context.Context, request ExecutePublishedDashboardQueryRequest) error

	// Poll the results for the a query for a published, embedded dashboard.
	PollPublishedQueryStatus(ctx context.Context, request PollPublishedQueryStatusRequest) (*PollQueryStatusResponse, error)
}

func NewQueryExecutionPreview(client *client.DatabricksClient) *QueryExecutionPreviewAPI {
	return &QueryExecutionPreviewAPI{
		queryExecutionPreviewImpl: queryExecutionPreviewImpl{
			client: client,
		},
	}
}

// Query execution APIs for AI / BI Dashboards
type QueryExecutionPreviewAPI struct {
	queryExecutionPreviewImpl
}

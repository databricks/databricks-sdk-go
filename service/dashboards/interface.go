// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"context"
)

// Genie provides a no-code experience for business users, powered by AI/BI.
// Analysts set up spaces that business users can use to ask questions using
// natural language. Genie uses data registered to Unity Catalog and requires at
// least CAN USE permission on a Pro or Serverless SQL warehouse. Also,
// Databricks Assistant must be enabled.
type GenieService interface {

	// Create conversation message.
	//
	// Create new message in [conversation](:method:genie/startconversation).
	// The AI response uses all previously created messages in the conversation
	// to respond.
	CreateMessage(ctx context.Context, request GenieCreateConversationMessageRequest) (*GenieMessage, error)

	// Execute SQL query in a conversation message.
	//
	// Execute the SQL query in the message.
	ExecuteMessageQuery(ctx context.Context, request GenieExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get conversation message.
	//
	// Get message from conversation.
	GetMessage(ctx context.Context, request GenieGetConversationMessageRequest) (*GenieMessage, error)

	// Get conversation message SQL query result.
	//
	// Get the result of SQL query if the message has a query attachment. This
	// is only available if a message has a query attachment and the message
	// status is `EXECUTING_QUERY`.
	GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get conversation message SQL query result by attachment id.
	//
	// Get the result of SQL query by attachment id This is only available if a
	// message has a query attachment and the message status is
	// `EXECUTING_QUERY`.
	GetMessageQueryResultByAttachment(ctx context.Context, request GenieGetQueryResultByAttachmentRequest) (*GenieGetMessageQueryResultResponse, error)

	// Start conversation.
	//
	// Start a new conversation.
	StartConversation(ctx context.Context, request GenieStartConversationMessageRequest) (*GenieStartConversationResponse, error)
}

// These APIs provide specific management operations for Lakeview dashboards.
// Generic resource management can be done with Workspace API (import, export,
// get-status, list, delete).
type LakeviewService interface {

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

	// Delete schedule subscription.
	DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) error

	// Get dashboard.
	//
	// Get a draft dashboard.
	Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error)

	// Get published dashboard.
	//
	// Get the current published dashboard.
	GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error)

	// Get dashboard schedule.
	GetSchedule(ctx context.Context, request GetScheduleRequest) (*Schedule, error)

	// Get schedule subscription.
	GetSubscription(ctx context.Context, request GetSubscriptionRequest) (*Subscription, error)

	// List dashboards.
	//
	// Use ListAll() to get all Dashboard instances, which will iterate over every result page.
	List(ctx context.Context, request ListDashboardsRequest) (*ListDashboardsResponse, error)

	// List dashboard schedules.
	//
	// Use ListSchedulesAll() to get all Schedule instances, which will iterate over every result page.
	ListSchedules(ctx context.Context, request ListSchedulesRequest) (*ListSchedulesResponse, error)

	// List schedule subscriptions.
	//
	// Use ListSubscriptionsAll() to get all Subscription instances, which will iterate over every result page.
	ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) (*ListSubscriptionsResponse, error)

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

	// Unpublish dashboard.
	//
	// Unpublish the dashboard.
	Unpublish(ctx context.Context, request UnpublishDashboardRequest) error

	// Update dashboard.
	//
	// Update a draft dashboard.
	Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error)

	// Update dashboard schedule.
	UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error)
}

// Token-based Lakeview APIs for embedding dashboards in external applications.
type LakeviewEmbeddedService interface {

	// Read a published dashboard in an embedded ui.
	//
	// Get the current published dashboard within an embedded context.
	GetPublishedDashboardEmbedded(ctx context.Context, request GetPublishedDashboardEmbeddedRequest) error
}

// Query execution APIs for AI / BI Dashboards
type QueryExecutionService interface {

	// Cancel the results for the a query for a published, embedded dashboard.
	CancelPublishedQueryExecution(ctx context.Context, request CancelPublishedQueryExecutionRequest) (*CancelQueryExecutionResponse, error)

	// Execute a query for a published dashboard.
	ExecutePublishedDashboardQuery(ctx context.Context, request ExecutePublishedDashboardQueryRequest) error

	// Poll the results for the a query for a published, embedded dashboard.
	PollPublishedQueryStatus(ctx context.Context, request PollPublishedQueryStatusRequest) (*PollQueryStatusResponse, error)
}

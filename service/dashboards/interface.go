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
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type GenieService interface {

	// Create new message in a [conversation](:method:genie/startconversation).
	// The AI response uses all previously created messages in the conversation
	// to respond.
	CreateMessage(ctx context.Context, request GenieCreateConversationMessageRequest) (*GenieMessage, error)

	// Delete a conversation.
	DeleteConversation(ctx context.Context, request GenieDeleteConversationRequest) error

	// Delete a conversation message.
	DeleteConversationMessage(ctx context.Context, request GenieDeleteConversationMessageRequest) error

	// Execute the SQL for a message query attachment. Use this API when the
	// query attachment has expired and needs to be re-executed.
	ExecuteMessageAttachmentQuery(ctx context.Context, request GenieExecuteMessageAttachmentQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// DEPRECATED: Use [Execute Message Attachment
	// Query](:method:genie/executemessageattachmentquery) instead.
	ExecuteMessageQuery(ctx context.Context, request GenieExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get message from conversation.
	GetMessage(ctx context.Context, request GenieGetConversationMessageRequest) (*GenieMessage, error)

	// Get the result of SQL query if the message has a query attachment. This
	// is only available if a message has a query attachment and the message
	// status is `EXECUTING_QUERY` OR `COMPLETED`.
	GetMessageAttachmentQueryResult(ctx context.Context, request GenieGetMessageAttachmentQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// DEPRECATED: Use [Get Message Attachment Query
	// Result](:method:genie/getmessageattachmentqueryresult) instead.
	GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// DEPRECATED: Use [Get Message Attachment Query
	// Result](:method:genie/getmessageattachmentqueryresult) instead.
	GetMessageQueryResultByAttachment(ctx context.Context, request GenieGetQueryResultByAttachmentRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get details of a Genie Space.
	GetSpace(ctx context.Context, request GenieGetSpaceRequest) (*GenieSpace, error)

	// List messages in a conversation
	ListConversationMessages(ctx context.Context, request GenieListConversationMessagesRequest) (*GenieListConversationMessagesResponse, error)

	// Get a list of conversations in a Genie Space.
	ListConversations(ctx context.Context, request GenieListConversationsRequest) (*GenieListConversationsResponse, error)

	// Get list of Genie Spaces.
	ListSpaces(ctx context.Context, request GenieListSpacesRequest) (*GenieListSpacesResponse, error)

	// Send feedback for a message.
	SendMessageFeedback(ctx context.Context, request GenieSendMessageFeedbackRequest) error

	// Start a new conversation.
	StartConversation(ctx context.Context, request GenieStartConversationMessageRequest) (*GenieStartConversationResponse, error)

	// Move a Genie Space to the trash.
	TrashSpace(ctx context.Context, request GenieTrashSpaceRequest) error
}

// These APIs provide specific management operations for Lakeview dashboards.
// Generic resource management can be done with Workspace API (import, export,
// get-status, list, delete).
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type LakeviewService interface {

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

	// Get a draft dashboard.
	Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error)

	// Get the current published dashboard.
	GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error)

	// Get dashboard schedule.
	GetSchedule(ctx context.Context, request GetScheduleRequest) (*Schedule, error)

	// Get schedule subscription.
	GetSubscription(ctx context.Context, request GetSubscriptionRequest) (*Subscription, error)

	// List dashboards.
	List(ctx context.Context, request ListDashboardsRequest) (*ListDashboardsResponse, error)

	// List dashboard schedules.
	ListSchedules(ctx context.Context, request ListSchedulesRequest) (*ListSchedulesResponse, error)

	// List schedule subscriptions.
	ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) (*ListSubscriptionsResponse, error)

	// Migrates a classic SQL dashboard to Lakeview.
	Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error)

	// Publish the current draft dashboard.
	Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error)

	// Trash a dashboard.
	Trash(ctx context.Context, request TrashDashboardRequest) error

	// Unpublish the dashboard.
	Unpublish(ctx context.Context, request UnpublishDashboardRequest) error

	// Update a draft dashboard.
	Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error)

	// Update dashboard schedule.
	UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error)
}

// Token-based Lakeview APIs for embedding dashboards in external applications.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type LakeviewEmbeddedService interface {

	// Get a required authorization details and scopes of a published dashboard
	// to mint an OAuth token.
	GetPublishedDashboardTokenInfo(ctx context.Context, request GetPublishedDashboardTokenInfoRequest) (*GetPublishedDashboardTokenInfoResponse, error)
}

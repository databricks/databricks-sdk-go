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

	// Sends a new message in a chat-mode
	// [conversation](:method:genie/startconversation). The AI response uses all
	// previously created messages in the conversation to respond.
	CreateMessage(ctx context.Context, request GenieCreateConversationMessageRequest) (*GenieMessage, error)

	// Create a comment on a conversation message.
	CreateMessageComment(ctx context.Context, request GenieCreateMessageCommentRequest) (*GenieMessageComment, error)

	// Creates a Genie space from a serialized payload.
	CreateSpace(ctx context.Context, request GenieCreateSpaceRequest) (*GenieSpace, error)

	// Delete a conversation.
	DeleteConversation(ctx context.Context, request GenieDeleteConversationRequest) error

	// Delete a conversation message.
	DeleteConversationMessage(ctx context.Context, request GenieDeleteConversationMessageRequest) error

	// Download a rendered image of a message visualization attachment. The
	// response body is the raw PNG image, not a JSON payload. This is only
	// available if the attachment is a visualization and the message status is
	// `COMPLETED`. This endpoint is not supported for Private Link workspaces.
	DownloadMessageAttachmentVisualization(ctx context.Context, request DownloadMessageAttachmentVisualizationRequest) (*DownloadMessageAttachmentVisualizationResponse, error)

	// Execute the SQL for a message query attachment. Use this API when the
	// query attachment has expired and needs to be re-executed.
	ExecuteMessageAttachmentQuery(ctx context.Context, request GenieExecuteMessageAttachmentQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// DEPRECATED: Use [Execute Message Attachment
	// Query](:method:genie/executemessageattachmentquery) instead.
	ExecuteMessageQuery(ctx context.Context, request GenieExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// Initiates a new SQL execution and returns a `download_id` and
	// `download_id_signature` that you can use to track the progress of the
	// download. The query result is stored in an external link and can be
	// retrieved using the [Get Download Full Query
	// Result](:method:genie/getdownloadfullqueryresult) API. Both `download_id`
	// and `download_id_signature` must be provided when calling the Get
	// endpoint.
	//
	// ----
	//
	// ### **Warning: Databricks strongly recommends that you protect the URLs
	// that are returned by the `EXTERNAL_LINKS` disposition.**
	//
	// When you use the `EXTERNAL_LINKS` disposition, a short-lived
	// cloud-storage URL is generated to download the results. The URL contains
	// temporary access credentials, so protect it and do not set an
	// `Authorization` header in the download request.
	//
	// See [Execute Statement](:method:statementexecution/executestatement) for
	// more details.
	//
	// ----
	GenerateDownloadFullQueryResult(ctx context.Context, request GenieGenerateDownloadFullQueryResultRequest) (*GenieGenerateDownloadFullQueryResultResponse, error)

	// Cancels an in-flight agent-mode response. `response_id` is the id
	// returned in the `response.created` event from the agent-mode responses
	// endpoint. The response stops at the next agent boundary and its terminal
	// state is returned.
	GenieCancelResponse(ctx context.Context, request GenieCancelResponseRequest) (*GenieMessage, error)

	// Creates and runs chat-mode evaluations for multiple benchmark questions
	// in a Genie space.
	GenieCreateEvalRun(ctx context.Context, request GenieCreateEvalRunRequest) (*GenieEvalRunResponse, error)

	// Get details for evaluation results.
	GenieGetEvalResultDetails(ctx context.Context, request GenieGetEvalResultDetailsRequest) (*GenieEvalResultDetails, error)

	// Get evaluation run details.
	GenieGetEvalRun(ctx context.Context, request GenieGetEvalRunRequest) (*GenieEvalRunResponse, error)

	// List evaluation results for a specific evaluation run.
	GenieListEvalResults(ctx context.Context, request GenieListEvalResultsRequest) (*GenieListEvalResultsResponse, error)

	// Lists all evaluation runs in a space.
	GenieListEvalRuns(ctx context.Context, request GenieListEvalRunsRequest) (*GenieListEvalRunsResponse, error)

	// After [Generating a Full Query Result
	// Download](:method:genie/generatedownloadfullqueryresult) and successfully
	// receiving a `download_id` and `download_id_signature`, use this API to
	// poll the download progress. Both `download_id` and
	// `download_id_signature` are required to call this endpoint. When the
	// download is complete, the API returns the result in the `EXTERNAL_LINKS`
	// disposition, containing one or more external links to the query result
	// files.
	//
	// ----
	//
	// ### **Warning: Databricks strongly recommends that you protect the URLs
	// that are returned by the `EXTERNAL_LINKS` disposition.**
	//
	// When you use the `EXTERNAL_LINKS` disposition, a short-lived
	// cloud-storage URL is generated to download the results. The URL contains
	// temporary access credentials, so protect it and do not set an
	// `Authorization` header in the download request.
	//
	// See [Execute Statement](:method:statementexecution/executestatement) for
	// more details.
	//
	// ----
	GetDownloadFullQueryResult(ctx context.Context, request GenieGetDownloadFullQueryResultRequest) (*GenieGetDownloadFullQueryResultResponse, error)

	// Gets a message from a chat-mode or agent-mode conversation. For a
	// complete agent-mode transcript, use the List conversation items endpoint.
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

	// List all comments across all messages in a conversation.
	ListConversationComments(ctx context.Context, request GenieListConversationCommentsRequest) (*GenieListConversationCommentsResponse, error)

	// Lists messages in a chat-mode or agent-mode conversation. Agent-mode
	// messages are returned as GenieMessage projections. Use the List
	// conversation items endpoint for the complete reasoning and tool-call
	// history.
	ListConversationMessages(ctx context.Context, request GenieListConversationMessagesRequest) (*GenieListConversationMessagesResponse, error)

	// Get a list of conversations in a Genie Space.
	ListConversations(ctx context.Context, request GenieListConversationsRequest) (*GenieListConversationsResponse, error)

	// List comments on a specific conversation message.
	ListMessageComments(ctx context.Context, request GenieListMessageCommentsRequest) (*GenieListMessageCommentsResponse, error)

	// Get list of Genie Spaces.
	ListSpaces(ctx context.Context, request GenieListSpacesRequest) (*GenieListSpacesResponse, error)

	// Sends feedback for a message in a chat-mode or agent-mode conversation.
	SendMessageFeedback(ctx context.Context, request GenieSendMessageFeedbackRequest) error

	// Starts a new chat-mode conversation and sends its first message.
	StartConversation(ctx context.Context, request GenieStartConversationMessageRequest) (*GenieStartConversationResponse, error)

	// Move a Genie Space to the trash.
	TrashSpace(ctx context.Context, request GenieTrashSpaceRequest) error

	// Updates a Genie space with a serialized payload.
	UpdateSpace(ctx context.Context, request GenieUpdateSpaceRequest) (*GenieSpace, error)
}

// These APIs provide specific management operations for Lakeview dashboards.
// Generic resource management can be done with Workspace API (import, export,
// get-status, list, delete).
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type LakeviewService interface {

	// Create a draft dashboard.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error)

	// Create dashboard schedule.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	CreateSchedule(ctx context.Context, request CreateScheduleRequest) (*Schedule, error)

	// Create schedule subscription.
	//
	// The caller must be a workspace user with one of the following
	// [entitlements]: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (*Subscription, error)

	// Delete dashboard schedule.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error

	// Delete schedule subscription.
	//
	// The caller must be a workspace user with one of the following
	// [entitlements]: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) error

	// Get a draft dashboard.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error)

	// Get the current published dashboard.
	//
	// The caller must be a workspace user with one of the following
	// entitlements: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error)

	// Get dashboard schedule.
	//
	// The caller must be a workspace user with one of the following
	// [entitlements]: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	GetSchedule(ctx context.Context, request GetScheduleRequest) (*Schedule, error)

	// Get schedule subscription.
	//
	// The caller must be a workspace user with one of the following
	// [entitlements]: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	GetSubscription(ctx context.Context, request GetSubscriptionRequest) (*Subscription, error)

	// List dashboards.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	List(ctx context.Context, request ListDashboardsRequest) (*ListDashboardsResponse, error)

	// List dashboard schedules.
	//
	// The caller must be a workspace user with one of the following
	// [entitlements]: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	ListSchedules(ctx context.Context, request ListSchedulesRequest) (*ListSchedulesResponse, error)

	// List schedule subscriptions.
	//
	// The caller must be a workspace user with one of the following
	// [entitlements]: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) (*ListSubscriptionsResponse, error)

	// Deprecated: Legacy dashboard migration is no longer supported. Use
	// Lakeview (AI/BI) dashboards instead.
	Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error)

	// Publish the current draft dashboard.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error)

	// Revert a dashboard's definition in draft mode to the last published
	// version.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	Revert(ctx context.Context, request RevertDashboardRequest) (*RevertDashboardResponse, error)

	// Trash a dashboard.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	Trash(ctx context.Context, request TrashDashboardRequest) error

	// Unpublish the dashboard.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	Unpublish(ctx context.Context, request UnpublishDashboardRequest) error

	// Update a draft dashboard.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error)

	// Update dashboard schedule.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error)
}

// Token-based Lakeview APIs for embedding dashboards in external applications.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type LakeviewEmbeddedService interface {

	// Get a required authorization details and scopes of a published dashboard
	// to mint an OAuth token.
	//
	// The caller must be a workspace user with one of the following
	// entitlements: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	GetPublishedDashboardTokenInfo(ctx context.Context, request GetPublishedDashboardTokenInfoRequest) (*GetPublishedDashboardTokenInfoResponse, error)
}

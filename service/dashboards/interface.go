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

	// Execute the SQL for a message query attachment. Use this API when the
	// query attachment has expired and needs to be re-executed.
	ExecuteMessageAttachmentQuery(ctx context.Context, request GenieExecuteMessageAttachmentQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// Execute the SQL query in the message.
	ExecuteMessageQuery(ctx context.Context, request GenieExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// Initiates a new SQL execution and returns a `download_id` that you can
	// use to track the progress of the download. The query result is stored in
	// an external link and can be retrieved using the [Get Download Full Query
	// Result](:method:genie/getdownloadfullqueryresult) API. Warning:
	// Databricks strongly recommends that you protect the URLs that are
	// returned by the `EXTERNAL_LINKS` disposition. See [Execute
	// Statement](:method:statementexecution/executestatement) for more details.
	GenerateDownloadFullQueryResult(ctx context.Context, request GenieGenerateDownloadFullQueryResultRequest) (*GenieGenerateDownloadFullQueryResultResponse, error)

	// After [Generating a Full Query Result
	// Download](:method:genie/getdownloadfullqueryresult) and successfully
	// receiving a `download_id`, use this API to poll the download progress.
	// When the download is complete, the API returns one or more external links
	// to the query result files. Warning: Databricks strongly recommends that
	// you protect the URLs that are returned by the `EXTERNAL_LINKS`
	// disposition. You must not set an Authorization header in download
	// requests. When using the `EXTERNAL_LINKS` disposition, Databricks returns
	// presigned URLs that grant temporary access to data. See [Execute
	// Statement](:method:statementexecution/executestatement) for more details.
	GetDownloadFullQueryResult(ctx context.Context, request GenieGetDownloadFullQueryResultRequest) (*GenieGetDownloadFullQueryResultResponse, error)

	// Get message from conversation.
	GetMessage(ctx context.Context, request GenieGetConversationMessageRequest) (*GenieMessage, error)

	// Get the result of SQL query if the message has a query attachment. This
	// is only available if a message has a query attachment and the message
	// status is `EXECUTING_QUERY` OR `COMPLETED`.
	GetMessageAttachmentQueryResult(ctx context.Context, request GenieGetMessageAttachmentQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get the result of SQL query if the message has a query attachment. This
	// is only available if a message has a query attachment and the message
	// status is `EXECUTING_QUERY`.
	GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get the result of SQL query if the message has a query attachment. This
	// is only available if a message has a query attachment and the message
	// status is `EXECUTING_QUERY` OR `COMPLETED`.
	GetMessageQueryResultByAttachment(ctx context.Context, request GenieGetQueryResultByAttachmentRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get details of a Genie Space.
	GetSpace(ctx context.Context, request GenieGetSpaceRequest) (*GenieSpace, error)

	// Get list of Genie Spaces.
	ListSpaces(ctx context.Context, request GenieListSpacesRequest) (*GenieListSpacesResponse, error)

	// Start a new conversation.
	StartConversation(ctx context.Context, request GenieStartConversationMessageRequest) (*GenieStartConversationResponse, error)
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
	// to mint an OAuth token. The `authorization_details` can be enriched to
	// apply additional restriction.
	//
	// Example: Adding the following `authorization_details` object to downscope
	// the viewer permission to specific table ``` { type:
	// "unity_catalog_privileges", privileges: ["SELECT"], object_type: "TABLE",
	// object_full_path: "main.default.testdata" } ```
	GetPublishedDashboardTokenInfo(ctx context.Context, request GetPublishedDashboardTokenInfoRequest) (*GetPublishedDashboardTokenInfoResponse, error)
}

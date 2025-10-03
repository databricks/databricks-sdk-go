// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/sql"
)

type AuthorizationDetails struct {
	// Represents downscoped permission rules with specific access rights. This
	// field is specific to `workspace_rule_set` constraint.
	GrantRules []AuthorizationDetailsGrantRule `json:"grant_rules,omitempty"`
	// The acl path of the tree store resource resource.
	ResourceLegacyAclPath string `json:"resource_legacy_acl_path,omitempty"`
	// The resource name to which the authorization rule applies. This field is
	// specific to `workspace_rule_set` constraint. Format:
	// `workspaces/{workspace_id}/dashboards/{dashboard_id}`
	ResourceName string `json:"resource_name,omitempty"`
	// The type of authorization downscoping policy. Ex: `workspace_rule_set`
	// defines access rules for a specific workspace resource
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AuthorizationDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AuthorizationDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AuthorizationDetailsGrantRule struct {
	// Permission sets for dashboard are defined in
	// iam-common/rbac-common/permission-sets/definitions/TreeStoreBasePermissionSets
	// Ex: `permissionSets/dashboard.runner`
	PermissionSet string `json:"permission_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AuthorizationDetailsGrantRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AuthorizationDetailsGrantRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateDashboardRequest struct {
	Dashboard Dashboard `json:"dashboard"`
}

type CreateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// The schedule to create. A dashboard is limited to 10 schedules.
	Schedule Schedule `json:"schedule"`
}

type CreateSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId string `json:"-" url:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId string `json:"-" url:"-"`
	// The subscription to create. A schedule is limited to 100 subscriptions.
	Subscription Subscription `json:"subscription"`
}

type CronSchedule struct {
	// A cron expression using quartz syntax. EX: `0 0 8 * * ?` represents
	// everyday at 8am. See [Cron Trigger] for details.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// A Java timezone id. The schedule will be resolved with respect to this
	// timezone. See [Java TimeZone] for details.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId string `json:"timezone_id"`
}

type Dashboard struct {
	// The timestamp of when the dashboard was created.
	CreateTime string `json:"create_time,omitempty"`
	// UUID identifying the dashboard.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The display name of the dashboard.
	DisplayName string `json:"display_name,omitempty"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read. This
	// field is excluded in List Dashboards responses.
	Etag string `json:"etag,omitempty"`
	// The state of the dashboard resource. Used for tracking trashed status.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash. This field is excluded in List
	// Dashboards responses.
	ParentPath string `json:"parent_path,omitempty"`
	// The workspace path of the dashboard asset, including the file name.
	// Exported dashboards always have the file extension `.lvdash.json`. This
	// field is excluded in List Dashboards responses.
	Path string `json:"path,omitempty"`
	// The contents of the dashboard in serialized string form. This field is
	// excluded in List Dashboards responses. Use the [get dashboard API] to
	// retrieve an example response, which includes the `serialized_dashboard`
	// field. This field provides the structure of the JSON string that
	// represents the dashboard's layout and components.
	//
	// [get dashboard API]: https://docs.databricks.com/api/workspace/lakeview/get
	SerializedDashboard string `json:"serialized_dashboard,omitempty"`
	// The timestamp of when the dashboard was last updated by the user. This
	// field is excluded in List Dashboards responses.
	UpdateTime string `json:"update_time,omitempty"`
	// The warehouse ID used to run the dashboard.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Dashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Dashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DashboardView string

const DashboardViewDashboardViewBasic DashboardView = `DASHBOARD_VIEW_BASIC`

// String representation for [fmt.Print]
func (f *DashboardView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DashboardView) Set(v string) error {
	switch v {
	case `DASHBOARD_VIEW_BASIC`:
		*f = DashboardView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DASHBOARD_VIEW_BASIC"`, v)
	}
}

// Values returns all possible values for DashboardView.
//
// There is no guarantee on the order of the values in the slice.
func (f *DashboardView) Values() []DashboardView {
	return []DashboardView{
		DashboardViewDashboardViewBasic,
	}
}

// Type always returns DashboardView to satisfy [pflag.Value] interface
func (f *DashboardView) Type() string {
	return "DashboardView"
}

type DeleteScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// The etag for the schedule. Optionally, it can be provided to verify that
	// the schedule has not been modified from its last retrieval.
	Etag string `json:"-" url:"etag,omitempty"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteScheduleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteScheduleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId string `json:"-" url:"-"`
	// The etag for the subscription. Can be optionally provided to ensure that
	// the subscription has not been modified since the last read.
	Etag string `json:"-" url:"etag,omitempty"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId string `json:"-" url:"-"`
	// UUID identifying the subscription.
	SubscriptionId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteSubscriptionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Genie AI Response
type GenieAttachment struct {
	// Attachment ID
	AttachmentId string `json:"attachment_id,omitempty"`
	// Query Attachment if Genie responds with a SQL query
	Query *GenieQueryAttachment `json:"query,omitempty"`
	// Follow-up questions suggested by Genie
	SuggestedQuestions *GenieSuggestedQuestionsAttachment `json:"suggested_questions,omitempty"`
	// Text Attachment if Genie responds with text
	Text *TextAttachment `json:"text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieConversation struct {
	// Conversation ID
	ConversationId string `json:"conversation_id"`
	// Timestamp when the message was created
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// Conversation ID. Legacy identifier, use conversation_id instead
	Id string `json:"id"`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Genie space ID
	SpaceId string `json:"space_id"`
	// Conversation title
	Title string `json:"title"`
	// ID of the user who created the conversation
	UserId int `json:"user_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieConversation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieConversation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieConversationSummary struct {
	ConversationId string `json:"conversation_id"`

	CreatedTimestamp int64 `json:"created_timestamp"`

	Title string `json:"title"`
}

type GenieCreateConversationMessageRequest struct {
	// User message content.
	Content string `json:"content"`
	// The ID associated with the conversation.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the Genie space where the conversation is started.
	SpaceId string `json:"-" url:"-"`
}

type GenieDeleteConversationMessageRequest struct {
	// The ID associated with the conversation.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the message to delete.
	MessageId string `json:"-" url:"-"`
	// The ID associated with the Genie space where the message is located.
	SpaceId string `json:"-" url:"-"`
}

type GenieDeleteConversationRequest struct {
	// The ID of the conversation to delete.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the Genie space where the conversation is located.
	SpaceId string `json:"-" url:"-"`
}

type GenieExecuteMessageAttachmentQueryRequest struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

type GenieExecuteMessageQueryRequest struct {
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

// Feedback containing rating and optional comment
type GenieFeedback struct {
	// The feedback rating
	Rating GenieFeedbackRating `json:"rating,omitempty"`
}

// Feedback rating for Genie messages
type GenieFeedbackRating string

const GenieFeedbackRatingNegative GenieFeedbackRating = `NEGATIVE`

const GenieFeedbackRatingNone GenieFeedbackRating = `NONE`

const GenieFeedbackRatingPositive GenieFeedbackRating = `POSITIVE`

// String representation for [fmt.Print]
func (f *GenieFeedbackRating) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GenieFeedbackRating) Set(v string) error {
	switch v {
	case `NEGATIVE`, `NONE`, `POSITIVE`:
		*f = GenieFeedbackRating(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NEGATIVE", "NONE", "POSITIVE"`, v)
	}
}

// Values returns all possible values for GenieFeedbackRating.
//
// There is no guarantee on the order of the values in the slice.
func (f *GenieFeedbackRating) Values() []GenieFeedbackRating {
	return []GenieFeedbackRating{
		GenieFeedbackRatingNegative,
		GenieFeedbackRatingNone,
		GenieFeedbackRatingPositive,
	}
}

// Type always returns GenieFeedbackRating to satisfy [pflag.Value] interface
func (f *GenieFeedbackRating) Type() string {
	return "GenieFeedbackRating"
}

type GenieGetConversationMessageRequest struct {
	// The ID associated with the target conversation.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the target message from the identified
	// conversation.
	MessageId string `json:"-" url:"-"`
	// The ID associated with the Genie space where the target conversation is
	// located.
	SpaceId string `json:"-" url:"-"`
}

type GenieGetMessageAttachmentQueryResultRequest struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

type GenieGetMessageQueryResultRequest struct {
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

type GenieGetMessageQueryResultResponse struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	StatementResponse *sql.StatementResponse `json:"statement_response,omitempty"`
}

type GenieGetQueryResultByAttachmentRequest struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

type GenieGetSpaceRequest struct {
	// The ID associated with the Genie space
	SpaceId string `json:"-" url:"-"`
}

type GenieListConversationMessagesRequest struct {
	// The ID of the conversation to list messages from
	ConversationId string `json:"-" url:"-"`
	// Maximum number of messages to return per page
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token to get the next page of results
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The ID associated with the Genie space where the conversation is located
	SpaceId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListConversationMessagesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListConversationMessagesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListConversationMessagesResponse struct {
	// List of messages in the conversation.
	Messages []GenieMessage `json:"messages,omitempty"`
	// The token to use for retrieving the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListConversationMessagesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListConversationMessagesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListConversationsRequest struct {
	// Include all conversations in the space across all users. Requires at
	// least CAN MANAGE permission on the space.
	IncludeAll bool `json:"-" url:"include_all,omitempty"`
	// Maximum number of conversations to return per page
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token to get the next page of results
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The ID of the Genie space to retrieve conversations from.
	SpaceId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListConversationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListConversationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListConversationsResponse struct {
	// List of conversations in the Genie space
	Conversations []GenieConversationSummary `json:"conversations,omitempty"`
	// Token to get the next page of results
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListConversationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListConversationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListSpacesRequest struct {
	// Maximum number of spaces to return per page
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token for getting the next page of results
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListSpacesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListSpacesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListSpacesResponse struct {
	// Token to get the next page of results
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of Genie spaces
	Spaces []GenieSpace `json:"spaces,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListSpacesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListSpacesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieMessage struct {
	// AI-generated response to the message
	Attachments []GenieAttachment `json:"attachments,omitempty"`
	// User message content
	Content string `json:"content"`
	// Conversation ID
	ConversationId string `json:"conversation_id"`
	// Timestamp when the message was created
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// Error message if Genie failed to respond to the message
	Error *MessageError `json:"error,omitempty"`
	// User feedback for the message if provided
	Feedback *GenieFeedback `json:"feedback,omitempty"`
	// Message ID. Legacy identifier, use message_id instead
	Id string `json:"id"`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Message ID
	MessageId string `json:"message_id"`
	// The result of SQL query if the message includes a query attachment.
	// Deprecated. Use `query_result_metadata` in `GenieQueryAttachment`
	// instead.
	QueryResult *Result `json:"query_result,omitempty"`
	// Genie space ID
	SpaceId string `json:"space_id"`

	Status MessageStatus `json:"status,omitempty"`
	// ID of the user who created the message
	UserId int64 `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieQueryAttachment struct {
	// Description of the query
	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`
	// Time when the user updated the query last
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`

	Parameters []QueryAttachmentParameter `json:"parameters,omitempty"`
	// AI generated SQL query
	Query string `json:"query,omitempty"`
	// Metadata associated with the query result.
	QueryResultMetadata *GenieResultMetadata `json:"query_result_metadata,omitempty"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId string `json:"statement_id,omitempty"`
	// Name of the query
	Title string `json:"title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieQueryAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieQueryAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieResultMetadata struct {
	// Indicates whether the result set is truncated.
	IsTruncated bool `json:"is_truncated,omitempty"`
	// The number of rows in the result set.
	RowCount int64 `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieResultMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieResultMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieSendMessageFeedbackRequest struct {
	// The ID associated with the conversation.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the message to provide feedback for.
	MessageId string `json:"-" url:"-"`
	// The rating (POSITIVE, NEGATIVE, or NONE).
	Rating GenieFeedbackRating `json:"rating"`
	// The ID associated with the Genie space where the message is located.
	SpaceId string `json:"-" url:"-"`
}

type GenieSpace struct {
	// Description of the Genie Space
	Description string `json:"description,omitempty"`
	// Genie space ID
	SpaceId string `json:"space_id"`
	// Title of the Genie Space
	Title string `json:"title"`
	// Warehouse associated with the Genie Space
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieSpace) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieSpace) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieStartConversationMessageRequest struct {
	// The text of the message that starts the conversation.
	Content string `json:"content"`
	// The ID associated with the Genie space where you want to start a
	// conversation.
	SpaceId string `json:"-" url:"-"`
}

type GenieStartConversationResponse struct {
	Conversation *GenieConversation `json:"conversation,omitempty"`
	// Conversation ID
	ConversationId string `json:"conversation_id"`

	Message *GenieMessage `json:"message,omitempty"`
	// Message ID
	MessageId string `json:"message_id"`
}

// Follow-up questions suggested by Genie
type GenieSuggestedQuestionsAttachment struct {
	// The suggested follow-up questions
	Questions []string `json:"questions,omitempty"`
}

type GenieTrashSpaceRequest struct {
	// The ID associated with the Genie space to be sent to the trash.
	SpaceId string `json:"-" url:"-"`
}

type GetDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
}

type GetPublishedDashboardRequest struct {
	// UUID identifying the published dashboard.
	DashboardId string `json:"-" url:"-"`
}

type GetPublishedDashboardTokenInfoRequest struct {
	// UUID identifying the published dashboard.
	DashboardId string `json:"-" url:"-"`
	// Provided external value to be included in the custom claim.
	ExternalValue string `json:"-" url:"external_value,omitempty"`
	// Provided external viewer id to be included in the custom claim.
	ExternalViewerId string `json:"-" url:"external_viewer_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPublishedDashboardTokenInfoRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedDashboardTokenInfoRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPublishedDashboardTokenInfoResponse struct {
	// Authorization constraints for accessing the published dashboard.
	// Currently includes `workspace_rule_set` and could be enriched with
	// `unity_catalog_privileges` before oAuth token generation.
	AuthorizationDetails []AuthorizationDetails `json:"authorization_details,omitempty"`
	// Custom claim generated from external_value and external_viewer_id.
	// Format:
	// `urn:aibi:external_data:<external_value>:<external_viewer_id>:<dashboard_id>`
	CustomClaim string `json:"custom_claim,omitempty"`
	// Scope defining access permissions.
	Scope string `json:"scope,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPublishedDashboardTokenInfoResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedDashboardTokenInfoResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" url:"-"`
}

type GetSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId string `json:"-" url:"-"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId string `json:"-" url:"-"`
	// UUID identifying the subscription.
	SubscriptionId string `json:"-" url:"-"`
}

type LifecycleState string

const LifecycleStateActive LifecycleState = `ACTIVE`

const LifecycleStateTrashed LifecycleState = `TRASHED`

// String representation for [fmt.Print]
func (f *LifecycleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LifecycleState) Set(v string) error {
	switch v {
	case `ACTIVE`, `TRASHED`:
		*f = LifecycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "TRASHED"`, v)
	}
}

// Values returns all possible values for LifecycleState.
//
// There is no guarantee on the order of the values in the slice.
func (f *LifecycleState) Values() []LifecycleState {
	return []LifecycleState{
		LifecycleStateActive,
		LifecycleStateTrashed,
	}
}

// Type always returns LifecycleState to satisfy [pflag.Value] interface
func (f *LifecycleState) Type() string {
	return "LifecycleState"
}

type ListDashboardsRequest struct {
	// The number of dashboards to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListDashboards` call. This token
	// can be used to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The flag to include dashboards located in the trash. If unspecified, only
	// active dashboards will be returned.
	ShowTrashed bool `json:"-" url:"show_trashed,omitempty"`
	// `DASHBOARD_VIEW_BASIC`only includes summary metadata from the dashboard.
	View DashboardView `json:"-" url:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDashboardsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDashboardsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDashboardsResponse struct {
	Dashboards []Dashboard `json:"dashboards,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDashboardsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDashboardsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSchedulesRequest struct {
	// UUID identifying the dashboard to which the schedules belongs.
	DashboardId string `json:"-" url:"-"`
	// The number of schedules to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSchedulesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchedulesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSchedulesResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent schedules.
	NextPageToken string `json:"next_page_token,omitempty"`

	Schedules []Schedule `json:"schedules,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSchedulesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchedulesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSubscriptionsRequest struct {
	// UUID identifying the dashboard which the subscriptions belongs.
	DashboardId string `json:"-" url:"-"`
	// The number of subscriptions to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListSubscriptions` call. Use this
	// to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// UUID identifying the schedule which the subscriptions belongs.
	ScheduleId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSubscriptionsResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent subscriptions.
	NextPageToken string `json:"next_page_token,omitempty"`

	Subscriptions []Subscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSubscriptionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSubscriptionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MessageError struct {
	Error string `json:"error,omitempty"`

	Type MessageErrorType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MessageError) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MessageError) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MessageErrorType string

const MessageErrorTypeBlockMultipleExecutionsException MessageErrorType = `BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION`

const MessageErrorTypeChatCompletionClientException MessageErrorType = `CHAT_COMPLETION_CLIENT_EXCEPTION`

const MessageErrorTypeChatCompletionClientTimeoutException MessageErrorType = `CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION`

const MessageErrorTypeChatCompletionNetworkException MessageErrorType = `CHAT_COMPLETION_NETWORK_EXCEPTION`

const MessageErrorTypeContentFilterException MessageErrorType = `CONTENT_FILTER_EXCEPTION`

const MessageErrorTypeContextExceededException MessageErrorType = `CONTEXT_EXCEEDED_EXCEPTION`

const MessageErrorTypeCouldNotGetModelDeploymentsException MessageErrorType = `COULD_NOT_GET_MODEL_DEPLOYMENTS_EXCEPTION`

const MessageErrorTypeCouldNotGetUcSchemaException MessageErrorType = `COULD_NOT_GET_UC_SCHEMA_EXCEPTION`

const MessageErrorTypeDeploymentNotFoundException MessageErrorType = `DEPLOYMENT_NOT_FOUND_EXCEPTION`

const MessageErrorTypeDescribeQueryInvalidSqlError MessageErrorType = `DESCRIBE_QUERY_INVALID_SQL_ERROR`

const MessageErrorTypeDescribeQueryTimeout MessageErrorType = `DESCRIBE_QUERY_TIMEOUT`

const MessageErrorTypeDescribeQueryUnexpectedFailure MessageErrorType = `DESCRIBE_QUERY_UNEXPECTED_FAILURE`

const MessageErrorTypeExceededMaxTokenLengthException MessageErrorType = `EXCEEDED_MAX_TOKEN_LENGTH_EXCEPTION`

const MessageErrorTypeFunctionsNotAvailableException MessageErrorType = `FUNCTIONS_NOT_AVAILABLE_EXCEPTION`

const MessageErrorTypeFunctionArgumentsInvalidException MessageErrorType = `FUNCTION_ARGUMENTS_INVALID_EXCEPTION`

const MessageErrorTypeFunctionArgumentsInvalidJsonException MessageErrorType = `FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION`

const MessageErrorTypeFunctionArgumentsInvalidTypeException MessageErrorType = `FUNCTION_ARGUMENTS_INVALID_TYPE_EXCEPTION`

const MessageErrorTypeFunctionCallMissingParameterException MessageErrorType = `FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION`

const MessageErrorTypeGeneratedSqlQueryTooLongException MessageErrorType = `GENERATED_SQL_QUERY_TOO_LONG_EXCEPTION`

const MessageErrorTypeGenericChatCompletionException MessageErrorType = `GENERIC_CHAT_COMPLETION_EXCEPTION`

const MessageErrorTypeGenericChatCompletionServiceException MessageErrorType = `GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION`

const MessageErrorTypeGenericSqlExecApiCallException MessageErrorType = `GENERIC_SQL_EXEC_API_CALL_EXCEPTION`

const MessageErrorTypeIllegalParameterDefinitionException MessageErrorType = `ILLEGAL_PARAMETER_DEFINITION_EXCEPTION`

const MessageErrorTypeInternalCatalogAssetCreationFailedException MessageErrorType = `INTERNAL_CATALOG_ASSET_CREATION_FAILED_EXCEPTION`

const MessageErrorTypeInternalCatalogAssetCreationOngoingException MessageErrorType = `INTERNAL_CATALOG_ASSET_CREATION_ONGOING_EXCEPTION`

const MessageErrorTypeInternalCatalogAssetCreationUnsupportedException MessageErrorType = `INTERNAL_CATALOG_ASSET_CREATION_UNSUPPORTED_EXCEPTION`

const MessageErrorTypeInternalCatalogMissingUcPathException MessageErrorType = `INTERNAL_CATALOG_MISSING_UC_PATH_EXCEPTION`

const MessageErrorTypeInternalCatalogPathOverlapException MessageErrorType = `INTERNAL_CATALOG_PATH_OVERLAP_EXCEPTION`

const MessageErrorTypeInvalidCertifiedAnswerFunctionException MessageErrorType = `INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION`

const MessageErrorTypeInvalidCertifiedAnswerIdentifierException MessageErrorType = `INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION`

const MessageErrorTypeInvalidChatCompletionArgumentsJsonException MessageErrorType = `INVALID_CHAT_COMPLETION_ARGUMENTS_JSON_EXCEPTION`

const MessageErrorTypeInvalidChatCompletionJsonException MessageErrorType = `INVALID_CHAT_COMPLETION_JSON_EXCEPTION`

const MessageErrorTypeInvalidCompletionRequestException MessageErrorType = `INVALID_COMPLETION_REQUEST_EXCEPTION`

const MessageErrorTypeInvalidFunctionCallException MessageErrorType = `INVALID_FUNCTION_CALL_EXCEPTION`

const MessageErrorTypeInvalidSqlMultipleDatasetReferencesException MessageErrorType = `INVALID_SQL_MULTIPLE_DATASET_REFERENCES_EXCEPTION`

const MessageErrorTypeInvalidSqlMultipleStatementsException MessageErrorType = `INVALID_SQL_MULTIPLE_STATEMENTS_EXCEPTION`

const MessageErrorTypeInvalidSqlUnknownTableException MessageErrorType = `INVALID_SQL_UNKNOWN_TABLE_EXCEPTION`

const MessageErrorTypeInvalidTableIdentifierException MessageErrorType = `INVALID_TABLE_IDENTIFIER_EXCEPTION`

const MessageErrorTypeLocalContextExceededException MessageErrorType = `LOCAL_CONTEXT_EXCEEDED_EXCEPTION`

const MessageErrorTypeMessageAttachmentTooLongError MessageErrorType = `MESSAGE_ATTACHMENT_TOO_LONG_ERROR`

const MessageErrorTypeMessageCancelledWhileExecutingException MessageErrorType = `MESSAGE_CANCELLED_WHILE_EXECUTING_EXCEPTION`

const MessageErrorTypeMessageDeletedWhileExecutingException MessageErrorType = `MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION`

const MessageErrorTypeMessageUpdatedWhileExecutingException MessageErrorType = `MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION`

const MessageErrorTypeMissingSqlQueryException MessageErrorType = `MISSING_SQL_QUERY_EXCEPTION`

const MessageErrorTypeNoDeploymentsAvailableToWorkspace MessageErrorType = `NO_DEPLOYMENTS_AVAILABLE_TO_WORKSPACE`

const MessageErrorTypeNoQueryToVisualizeException MessageErrorType = `NO_QUERY_TO_VISUALIZE_EXCEPTION`

const MessageErrorTypeNoTablesToQueryException MessageErrorType = `NO_TABLES_TO_QUERY_EXCEPTION`

const MessageErrorTypeRateLimitExceededGenericException MessageErrorType = `RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION`

const MessageErrorTypeRateLimitExceededSpecifiedWaitException MessageErrorType = `RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION`

const MessageErrorTypeReplyProcessTimeoutException MessageErrorType = `REPLY_PROCESS_TIMEOUT_EXCEPTION`

const MessageErrorTypeRetryableProcessingException MessageErrorType = `RETRYABLE_PROCESSING_EXCEPTION`

const MessageErrorTypeSqlExecutionException MessageErrorType = `SQL_EXECUTION_EXCEPTION`

const MessageErrorTypeStopProcessDueToAutoRegenerate MessageErrorType = `STOP_PROCESS_DUE_TO_AUTO_REGENERATE`

const MessageErrorTypeTablesMissingException MessageErrorType = `TABLES_MISSING_EXCEPTION`

const MessageErrorTypeTooManyCertifiedAnswersException MessageErrorType = `TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION`

const MessageErrorTypeTooManyTablesException MessageErrorType = `TOO_MANY_TABLES_EXCEPTION`

const MessageErrorTypeUnexpectedReplyProcessException MessageErrorType = `UNEXPECTED_REPLY_PROCESS_EXCEPTION`

const MessageErrorTypeUnknownAiModel MessageErrorType = `UNKNOWN_AI_MODEL`

const MessageErrorTypeWarehouseAccessMissingException MessageErrorType = `WAREHOUSE_ACCESS_MISSING_EXCEPTION`

const MessageErrorTypeWarehouseNotFoundException MessageErrorType = `WAREHOUSE_NOT_FOUND_EXCEPTION`

// String representation for [fmt.Print]
func (f *MessageErrorType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MessageErrorType) Set(v string) error {
	switch v {
	case `BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION`, `CHAT_COMPLETION_CLIENT_EXCEPTION`, `CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION`, `CHAT_COMPLETION_NETWORK_EXCEPTION`, `CONTENT_FILTER_EXCEPTION`, `CONTEXT_EXCEEDED_EXCEPTION`, `COULD_NOT_GET_MODEL_DEPLOYMENTS_EXCEPTION`, `COULD_NOT_GET_UC_SCHEMA_EXCEPTION`, `DEPLOYMENT_NOT_FOUND_EXCEPTION`, `DESCRIBE_QUERY_INVALID_SQL_ERROR`, `DESCRIBE_QUERY_TIMEOUT`, `DESCRIBE_QUERY_UNEXPECTED_FAILURE`, `EXCEEDED_MAX_TOKEN_LENGTH_EXCEPTION`, `FUNCTIONS_NOT_AVAILABLE_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_TYPE_EXCEPTION`, `FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION`, `GENERATED_SQL_QUERY_TOO_LONG_EXCEPTION`, `GENERIC_CHAT_COMPLETION_EXCEPTION`, `GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION`, `GENERIC_SQL_EXEC_API_CALL_EXCEPTION`, `ILLEGAL_PARAMETER_DEFINITION_EXCEPTION`, `INTERNAL_CATALOG_ASSET_CREATION_FAILED_EXCEPTION`, `INTERNAL_CATALOG_ASSET_CREATION_ONGOING_EXCEPTION`, `INTERNAL_CATALOG_ASSET_CREATION_UNSUPPORTED_EXCEPTION`, `INTERNAL_CATALOG_MISSING_UC_PATH_EXCEPTION`, `INTERNAL_CATALOG_PATH_OVERLAP_EXCEPTION`, `INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION`, `INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION`, `INVALID_CHAT_COMPLETION_ARGUMENTS_JSON_EXCEPTION`, `INVALID_CHAT_COMPLETION_JSON_EXCEPTION`, `INVALID_COMPLETION_REQUEST_EXCEPTION`, `INVALID_FUNCTION_CALL_EXCEPTION`, `INVALID_SQL_MULTIPLE_DATASET_REFERENCES_EXCEPTION`, `INVALID_SQL_MULTIPLE_STATEMENTS_EXCEPTION`, `INVALID_SQL_UNKNOWN_TABLE_EXCEPTION`, `INVALID_TABLE_IDENTIFIER_EXCEPTION`, `LOCAL_CONTEXT_EXCEEDED_EXCEPTION`, `MESSAGE_ATTACHMENT_TOO_LONG_ERROR`, `MESSAGE_CANCELLED_WHILE_EXECUTING_EXCEPTION`, `MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION`, `MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION`, `MISSING_SQL_QUERY_EXCEPTION`, `NO_DEPLOYMENTS_AVAILABLE_TO_WORKSPACE`, `NO_QUERY_TO_VISUALIZE_EXCEPTION`, `NO_TABLES_TO_QUERY_EXCEPTION`, `RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION`, `RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION`, `REPLY_PROCESS_TIMEOUT_EXCEPTION`, `RETRYABLE_PROCESSING_EXCEPTION`, `SQL_EXECUTION_EXCEPTION`, `STOP_PROCESS_DUE_TO_AUTO_REGENERATE`, `TABLES_MISSING_EXCEPTION`, `TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION`, `TOO_MANY_TABLES_EXCEPTION`, `UNEXPECTED_REPLY_PROCESS_EXCEPTION`, `UNKNOWN_AI_MODEL`, `WAREHOUSE_ACCESS_MISSING_EXCEPTION`, `WAREHOUSE_NOT_FOUND_EXCEPTION`:
		*f = MessageErrorType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION", "CHAT_COMPLETION_CLIENT_EXCEPTION", "CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION", "CHAT_COMPLETION_NETWORK_EXCEPTION", "CONTENT_FILTER_EXCEPTION", "CONTEXT_EXCEEDED_EXCEPTION", "COULD_NOT_GET_MODEL_DEPLOYMENTS_EXCEPTION", "COULD_NOT_GET_UC_SCHEMA_EXCEPTION", "DEPLOYMENT_NOT_FOUND_EXCEPTION", "DESCRIBE_QUERY_INVALID_SQL_ERROR", "DESCRIBE_QUERY_TIMEOUT", "DESCRIBE_QUERY_UNEXPECTED_FAILURE", "EXCEEDED_MAX_TOKEN_LENGTH_EXCEPTION", "FUNCTIONS_NOT_AVAILABLE_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_TYPE_EXCEPTION", "FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION", "GENERATED_SQL_QUERY_TOO_LONG_EXCEPTION", "GENERIC_CHAT_COMPLETION_EXCEPTION", "GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION", "GENERIC_SQL_EXEC_API_CALL_EXCEPTION", "ILLEGAL_PARAMETER_DEFINITION_EXCEPTION", "INTERNAL_CATALOG_ASSET_CREATION_FAILED_EXCEPTION", "INTERNAL_CATALOG_ASSET_CREATION_ONGOING_EXCEPTION", "INTERNAL_CATALOG_ASSET_CREATION_UNSUPPORTED_EXCEPTION", "INTERNAL_CATALOG_MISSING_UC_PATH_EXCEPTION", "INTERNAL_CATALOG_PATH_OVERLAP_EXCEPTION", "INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION", "INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION", "INVALID_CHAT_COMPLETION_ARGUMENTS_JSON_EXCEPTION", "INVALID_CHAT_COMPLETION_JSON_EXCEPTION", "INVALID_COMPLETION_REQUEST_EXCEPTION", "INVALID_FUNCTION_CALL_EXCEPTION", "INVALID_SQL_MULTIPLE_DATASET_REFERENCES_EXCEPTION", "INVALID_SQL_MULTIPLE_STATEMENTS_EXCEPTION", "INVALID_SQL_UNKNOWN_TABLE_EXCEPTION", "INVALID_TABLE_IDENTIFIER_EXCEPTION", "LOCAL_CONTEXT_EXCEEDED_EXCEPTION", "MESSAGE_ATTACHMENT_TOO_LONG_ERROR", "MESSAGE_CANCELLED_WHILE_EXECUTING_EXCEPTION", "MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION", "MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION", "MISSING_SQL_QUERY_EXCEPTION", "NO_DEPLOYMENTS_AVAILABLE_TO_WORKSPACE", "NO_QUERY_TO_VISUALIZE_EXCEPTION", "NO_TABLES_TO_QUERY_EXCEPTION", "RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION", "RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION", "REPLY_PROCESS_TIMEOUT_EXCEPTION", "RETRYABLE_PROCESSING_EXCEPTION", "SQL_EXECUTION_EXCEPTION", "STOP_PROCESS_DUE_TO_AUTO_REGENERATE", "TABLES_MISSING_EXCEPTION", "TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION", "TOO_MANY_TABLES_EXCEPTION", "UNEXPECTED_REPLY_PROCESS_EXCEPTION", "UNKNOWN_AI_MODEL", "WAREHOUSE_ACCESS_MISSING_EXCEPTION", "WAREHOUSE_NOT_FOUND_EXCEPTION"`, v)
	}
}

// Values returns all possible values for MessageErrorType.
//
// There is no guarantee on the order of the values in the slice.
func (f *MessageErrorType) Values() []MessageErrorType {
	return []MessageErrorType{
		MessageErrorTypeBlockMultipleExecutionsException,
		MessageErrorTypeChatCompletionClientException,
		MessageErrorTypeChatCompletionClientTimeoutException,
		MessageErrorTypeChatCompletionNetworkException,
		MessageErrorTypeContentFilterException,
		MessageErrorTypeContextExceededException,
		MessageErrorTypeCouldNotGetModelDeploymentsException,
		MessageErrorTypeCouldNotGetUcSchemaException,
		MessageErrorTypeDeploymentNotFoundException,
		MessageErrorTypeDescribeQueryInvalidSqlError,
		MessageErrorTypeDescribeQueryTimeout,
		MessageErrorTypeDescribeQueryUnexpectedFailure,
		MessageErrorTypeExceededMaxTokenLengthException,
		MessageErrorTypeFunctionsNotAvailableException,
		MessageErrorTypeFunctionArgumentsInvalidException,
		MessageErrorTypeFunctionArgumentsInvalidJsonException,
		MessageErrorTypeFunctionArgumentsInvalidTypeException,
		MessageErrorTypeFunctionCallMissingParameterException,
		MessageErrorTypeGeneratedSqlQueryTooLongException,
		MessageErrorTypeGenericChatCompletionException,
		MessageErrorTypeGenericChatCompletionServiceException,
		MessageErrorTypeGenericSqlExecApiCallException,
		MessageErrorTypeIllegalParameterDefinitionException,
		MessageErrorTypeInternalCatalogAssetCreationFailedException,
		MessageErrorTypeInternalCatalogAssetCreationOngoingException,
		MessageErrorTypeInternalCatalogAssetCreationUnsupportedException,
		MessageErrorTypeInternalCatalogMissingUcPathException,
		MessageErrorTypeInternalCatalogPathOverlapException,
		MessageErrorTypeInvalidCertifiedAnswerFunctionException,
		MessageErrorTypeInvalidCertifiedAnswerIdentifierException,
		MessageErrorTypeInvalidChatCompletionArgumentsJsonException,
		MessageErrorTypeInvalidChatCompletionJsonException,
		MessageErrorTypeInvalidCompletionRequestException,
		MessageErrorTypeInvalidFunctionCallException,
		MessageErrorTypeInvalidSqlMultipleDatasetReferencesException,
		MessageErrorTypeInvalidSqlMultipleStatementsException,
		MessageErrorTypeInvalidSqlUnknownTableException,
		MessageErrorTypeInvalidTableIdentifierException,
		MessageErrorTypeLocalContextExceededException,
		MessageErrorTypeMessageAttachmentTooLongError,
		MessageErrorTypeMessageCancelledWhileExecutingException,
		MessageErrorTypeMessageDeletedWhileExecutingException,
		MessageErrorTypeMessageUpdatedWhileExecutingException,
		MessageErrorTypeMissingSqlQueryException,
		MessageErrorTypeNoDeploymentsAvailableToWorkspace,
		MessageErrorTypeNoQueryToVisualizeException,
		MessageErrorTypeNoTablesToQueryException,
		MessageErrorTypeRateLimitExceededGenericException,
		MessageErrorTypeRateLimitExceededSpecifiedWaitException,
		MessageErrorTypeReplyProcessTimeoutException,
		MessageErrorTypeRetryableProcessingException,
		MessageErrorTypeSqlExecutionException,
		MessageErrorTypeStopProcessDueToAutoRegenerate,
		MessageErrorTypeTablesMissingException,
		MessageErrorTypeTooManyCertifiedAnswersException,
		MessageErrorTypeTooManyTablesException,
		MessageErrorTypeUnexpectedReplyProcessException,
		MessageErrorTypeUnknownAiModel,
		MessageErrorTypeWarehouseAccessMissingException,
		MessageErrorTypeWarehouseNotFoundException,
	}
}

// Type always returns MessageErrorType to satisfy [pflag.Value] interface
func (f *MessageErrorType) Type() string {
	return "MessageErrorType"
}

// MessageStatus. The possible values are: * `FETCHING_METADATA`: Fetching
// metadata from the data sources. * `FILTERING_CONTEXT`: Running smart context
// step to determine relevant context. * `ASKING_AI`: Waiting for the LLM to
// respond to the user's question. * `PENDING_WAREHOUSE`: Waiting for warehouse
// before the SQL query can start executing. * `EXECUTING_QUERY`: Executing a
// generated SQL query. Get the SQL query result by calling
// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
// API. * `FAILED`: The response generation or query execution failed. See
// `error` field. * `COMPLETED`: Message processing is completed. Results are in
// the `attachments` field. Get the SQL query result by calling
// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
// API. * `SUBMITTED`: Message has been submitted. * `QUERY_RESULT_EXPIRED`: SQL
// result is not available anymore. The user needs to rerun the query. Rerun the
// SQL query result by calling
// [executeMessageAttachmentQuery](:method:genie/executeMessageAttachmentQuery)
// API. * `CANCELLED`: Message has been cancelled.
type MessageStatus string

// Waiting for the LLM to respond to the user's question.
const MessageStatusAskingAi MessageStatus = `ASKING_AI`

// Message has been cancelled.
const MessageStatusCancelled MessageStatus = `CANCELLED`

// Message processing is completed. Results are in the `attachments` field. Get
// the SQL query result by calling
// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
// API.
const MessageStatusCompleted MessageStatus = `COMPLETED`

// Executing a generated SQL query. Get the SQL query result by calling
// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
// API.
const MessageStatusExecutingQuery MessageStatus = `EXECUTING_QUERY`

// The response generation or query execution failed. See `error` field.
const MessageStatusFailed MessageStatus = `FAILED`

// Fetching metadata from the data sources.
const MessageStatusFetchingMetadata MessageStatus = `FETCHING_METADATA`

// Running smart context step to determine relevant context.
const MessageStatusFilteringContext MessageStatus = `FILTERING_CONTEXT`

// Waiting for warehouse before the SQL query can start executing.
const MessageStatusPendingWarehouse MessageStatus = `PENDING_WAREHOUSE`

// SQL result is not available anymore. The user needs to rerun the query. Rerun
// the SQL query result by calling
// [executeMessageAttachmentQuery](:method:genie/executeMessageAttachmentQuery)
// API.
const MessageStatusQueryResultExpired MessageStatus = `QUERY_RESULT_EXPIRED`

// Message has been submitted.
const MessageStatusSubmitted MessageStatus = `SUBMITTED`

// String representation for [fmt.Print]
func (f *MessageStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MessageStatus) Set(v string) error {
	switch v {
	case `ASKING_AI`, `CANCELLED`, `COMPLETED`, `EXECUTING_QUERY`, `FAILED`, `FETCHING_METADATA`, `FILTERING_CONTEXT`, `PENDING_WAREHOUSE`, `QUERY_RESULT_EXPIRED`, `SUBMITTED`:
		*f = MessageStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASKING_AI", "CANCELLED", "COMPLETED", "EXECUTING_QUERY", "FAILED", "FETCHING_METADATA", "FILTERING_CONTEXT", "PENDING_WAREHOUSE", "QUERY_RESULT_EXPIRED", "SUBMITTED"`, v)
	}
}

// Values returns all possible values for MessageStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *MessageStatus) Values() []MessageStatus {
	return []MessageStatus{
		MessageStatusAskingAi,
		MessageStatusCancelled,
		MessageStatusCompleted,
		MessageStatusExecutingQuery,
		MessageStatusFailed,
		MessageStatusFetchingMetadata,
		MessageStatusFilteringContext,
		MessageStatusPendingWarehouse,
		MessageStatusQueryResultExpired,
		MessageStatusSubmitted,
	}
}

// Type always returns MessageStatus to satisfy [pflag.Value] interface
func (f *MessageStatus) Type() string {
	return "MessageStatus"
}

type MigrateDashboardRequest struct {
	// Display name for the new Lakeview dashboard.
	DisplayName string `json:"display_name,omitempty"`
	// The workspace path of the folder to contain the migrated Lakeview
	// dashboard.
	ParentPath string `json:"parent_path,omitempty"`
	// UUID of the dashboard to be migrated.
	SourceDashboardId string `json:"source_dashboard_id"`
	// Flag to indicate if mustache parameter syntax ({{ param }}) should be
	// auto-updated to named syntax (:param) when converting datasets in the
	// dashboard.
	UpdateParameterSyntax bool `json:"update_parameter_syntax,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MigrateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MigrateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PublishRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId string `json:"-" url:"-"`
	// Flag to indicate if the publisher's credentials should be embedded in the
	// published dashboard. These embedded credentials will be used to execute
	// the published dashboard's queries.
	EmbedCredentials bool `json:"embed_credentials,omitempty"`
	// The ID of the warehouse that can be used to override the warehouse which
	// was set in the draft.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PublishRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PublishedDashboard struct {
	// The display name of the published dashboard.
	DisplayName string `json:"display_name,omitempty"`
	// Indicates whether credentials are embedded in the published dashboard.
	EmbedCredentials bool `json:"embed_credentials,omitempty"`
	// The timestamp of when the published dashboard was last revised.
	RevisionCreateTime string `json:"revision_create_time,omitempty"`
	// The warehouse ID used to run the published dashboard.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PublishedDashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishedDashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryAttachmentParameter struct {
	Keyword string `json:"keyword,omitempty"`

	SqlType string `json:"sql_type,omitempty"`

	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryAttachmentParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryAttachmentParameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Result struct {
	// If result is truncated
	IsTruncated bool `json:"is_truncated,omitempty"`
	// Row count of the result
	RowCount int64 `json:"row_count,omitempty"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId string `json:"statement_id,omitempty"`
	// JWT corresponding to the statement contained in this result
	StatementIdSignature string `json:"statement_id_signature,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Result) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Result) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Schedule struct {
	// A timestamp indicating when the schedule was created.
	CreateTime string `json:"create_time,omitempty"`
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule CronSchedule `json:"cron_schedule"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The display name for schedule.
	DisplayName string `json:"display_name,omitempty"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag string `json:"etag,omitempty"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus SchedulePauseStatus `json:"pause_status,omitempty"`
	// UUID identifying the schedule.
	ScheduleId string `json:"schedule_id,omitempty"`
	// A timestamp indicating when the schedule was last updated.
	UpdateTime string `json:"update_time,omitempty"`
	// The warehouse id to run the dashboard with for the schedule.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Schedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Schedule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SchedulePauseStatus string

const SchedulePauseStatusPaused SchedulePauseStatus = `PAUSED`

const SchedulePauseStatusUnpaused SchedulePauseStatus = `UNPAUSED`

// String representation for [fmt.Print]
func (f *SchedulePauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SchedulePauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`:
		*f = SchedulePauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED"`, v)
	}
}

// Values returns all possible values for SchedulePauseStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *SchedulePauseStatus) Values() []SchedulePauseStatus {
	return []SchedulePauseStatus{
		SchedulePauseStatusPaused,
		SchedulePauseStatusUnpaused,
	}
}

// Type always returns SchedulePauseStatus to satisfy [pflag.Value] interface
func (f *SchedulePauseStatus) Type() string {
	return "SchedulePauseStatus"
}

type Subscriber struct {
	// The destination to receive the subscription email. This parameter is
	// mutually exclusive with `user_subscriber`.
	DestinationSubscriber *SubscriptionSubscriberDestination `json:"destination_subscriber,omitempty"`
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	UserSubscriber *SubscriptionSubscriberUser `json:"user_subscriber,omitempty"`
}

type Subscription struct {
	// A timestamp indicating when the subscription was created.
	CreateTime string `json:"create_time,omitempty"`
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	CreatedByUserId int64 `json:"created_by_user_id,omitempty"`
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The etag for the subscription. Must be left empty on create, can be
	// optionally provided on delete to ensure that the subscription has not
	// been deleted since the last read.
	Etag string `json:"etag,omitempty"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId string `json:"schedule_id,omitempty"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber Subscriber `json:"subscriber"`
	// UUID identifying the subscription.
	SubscriptionId string `json:"subscription_id,omitempty"`
	// A timestamp indicating when the subscription was last updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Subscription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Subscription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SubscriptionSubscriberDestination struct {
	// The canonical identifier of the destination to receive email
	// notification.
	DestinationId string `json:"destination_id"`
}

type SubscriptionSubscriberUser struct {
	// UserId of the subscriber.
	UserId int64 `json:"user_id"`
}

type TextAttachment struct {
	// AI generated message
	Content string `json:"content,omitempty"`

	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TextAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TextAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TrashDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
}

type UnpublishDashboardRequest struct {
	// UUID identifying the published dashboard.
	DashboardId string `json:"-" url:"-"`
}

type UpdateDashboardRequest struct {
	Dashboard Dashboard `json:"dashboard"`
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
}

type UpdateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// The schedule to update.
	Schedule Schedule `json:"schedule"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" url:"-"`
}

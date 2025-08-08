// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboardspb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/sql/sqlpb"
)

type AuthorizationDetailsPb struct {
	GrantRules            []AuthorizationDetailsGrantRulePb `json:"grant_rules,omitempty"`
	ResourceLegacyAclPath string                            `json:"resource_legacy_acl_path,omitempty"`
	ResourceName          string                            `json:"resource_name,omitempty"`
	Type                  string                            `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AuthorizationDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AuthorizationDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AuthorizationDetailsGrantRulePb struct {
	PermissionSet string `json:"permission_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AuthorizationDetailsGrantRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AuthorizationDetailsGrantRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateDashboardRequestPb struct {
	Dashboard DashboardPb `json:"dashboard"`
}

type CreateScheduleRequestPb struct {
	DashboardId string     `json:"-" url:"-"`
	Schedule    SchedulePb `json:"schedule"`
}

type CreateSubscriptionRequestPb struct {
	DashboardId  string         `json:"-" url:"-"`
	ScheduleId   string         `json:"-" url:"-"`
	Subscription SubscriptionPb `json:"subscription"`
}

type CronSchedulePb struct {
	QuartzCronExpression string `json:"quartz_cron_expression"`
	TimezoneId           string `json:"timezone_id"`
}

type DashboardPb struct {
	CreateTime          string           `json:"create_time,omitempty"`
	DashboardId         string           `json:"dashboard_id,omitempty"`
	DisplayName         string           `json:"display_name,omitempty"`
	Etag                string           `json:"etag,omitempty"`
	LifecycleState      LifecycleStatePb `json:"lifecycle_state,omitempty"`
	ParentPath          string           `json:"parent_path,omitempty"`
	Path                string           `json:"path,omitempty"`
	SerializedDashboard string           `json:"serialized_dashboard,omitempty"`
	UpdateTime          string           `json:"update_time,omitempty"`
	WarehouseId         string           `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DashboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DashboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DashboardViewPb string

const DashboardViewDashboardViewBasic DashboardViewPb = `DASHBOARD_VIEW_BASIC`

type DeleteScheduleRequestPb struct {
	DashboardId string `json:"-" url:"-"`
	Etag        string `json:"-" url:"etag,omitempty"`
	ScheduleId  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteScheduleRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteScheduleRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteSubscriptionRequestPb struct {
	DashboardId    string `json:"-" url:"-"`
	Etag           string `json:"-" url:"etag,omitempty"`
	ScheduleId     string `json:"-" url:"-"`
	SubscriptionId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteSubscriptionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteSubscriptionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieAttachmentPb struct {
	AttachmentId string                  `json:"attachment_id,omitempty"`
	Query        *GenieQueryAttachmentPb `json:"query,omitempty"`
	Text         *TextAttachmentPb       `json:"text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenieAttachmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenieAttachmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieConversationPb struct {
	ConversationId       string `json:"conversation_id"`
	CreatedTimestamp     int64  `json:"created_timestamp,omitempty"`
	Id                   string `json:"id"`
	LastUpdatedTimestamp int64  `json:"last_updated_timestamp,omitempty"`
	SpaceId              string `json:"space_id"`
	Title                string `json:"title"`
	UserId               int    `json:"user_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenieConversationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenieConversationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieConversationSummaryPb struct {
	ConversationId   string `json:"conversation_id"`
	CreatedTimestamp int64  `json:"created_timestamp"`
	Title            string `json:"title"`
}

type GenieCreateConversationMessageRequestPb struct {
	Content        string `json:"content"`
	ConversationId string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

type GenieDeleteConversationRequestPb struct {
	ConversationId string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

type GenieExecuteMessageAttachmentQueryRequestPb struct {
	AttachmentId   string `json:"-" url:"-"`
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

type GenieExecuteMessageQueryRequestPb struct {
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

type GenieGetConversationMessageRequestPb struct {
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

type GenieGetMessageAttachmentQueryResultRequestPb struct {
	AttachmentId   string `json:"-" url:"-"`
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

type GenieGetMessageQueryResultRequestPb struct {
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

type GenieGetMessageQueryResultResponsePb struct {
	StatementResponse *sqlpb.StatementResponsePb `json:"statement_response,omitempty"`
}

type GenieGetQueryResultByAttachmentRequestPb struct {
	AttachmentId   string `json:"-" url:"-"`
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

type GenieGetSpaceRequestPb struct {
	SpaceId string `json:"-" url:"-"`
}

type GenieListConversationsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`
	SpaceId   string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenieListConversationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenieListConversationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieListConversationsResponsePb struct {
	Conversations []GenieConversationSummaryPb `json:"conversations,omitempty"`
	NextPageToken string                       `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenieListConversationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenieListConversationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieListSpacesRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenieListSpacesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenieListSpacesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieListSpacesResponsePb struct {
	NextPageToken string         `json:"next_page_token,omitempty"`
	Spaces        []GenieSpacePb `json:"spaces,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenieListSpacesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenieListSpacesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieMessagePb struct {
	Attachments          []GenieAttachmentPb `json:"attachments,omitempty"`
	Content              string              `json:"content"`
	ConversationId       string              `json:"conversation_id"`
	CreatedTimestamp     int64               `json:"created_timestamp,omitempty"`
	Error                *MessageErrorPb     `json:"error,omitempty"`
	Id                   string              `json:"id"`
	LastUpdatedTimestamp int64               `json:"last_updated_timestamp,omitempty"`
	MessageId            string              `json:"message_id"`
	QueryResult          *ResultPb           `json:"query_result,omitempty"`
	SpaceId              string              `json:"space_id"`
	Status               MessageStatusPb     `json:"status,omitempty"`
	UserId               int64               `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenieMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenieMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieQueryAttachmentPb struct {
	Description          string                 `json:"description,omitempty"`
	Id                   string                 `json:"id,omitempty"`
	LastUpdatedTimestamp int64                  `json:"last_updated_timestamp,omitempty"`
	Query                string                 `json:"query,omitempty"`
	QueryResultMetadata  *GenieResultMetadataPb `json:"query_result_metadata,omitempty"`
	StatementId          string                 `json:"statement_id,omitempty"`
	Title                string                 `json:"title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenieQueryAttachmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenieQueryAttachmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieResultMetadataPb struct {
	IsTruncated bool  `json:"is_truncated,omitempty"`
	RowCount    int64 `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenieResultMetadataPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenieResultMetadataPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieSpacePb struct {
	Description string `json:"description,omitempty"`
	SpaceId     string `json:"space_id"`
	Title       string `json:"title"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenieSpacePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenieSpacePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieStartConversationMessageRequestPb struct {
	Content string `json:"content"`
	SpaceId string `json:"-" url:"-"`
}

type GenieStartConversationResponsePb struct {
	Conversation   *GenieConversationPb `json:"conversation,omitempty"`
	ConversationId string               `json:"conversation_id"`
	Message        *GenieMessagePb      `json:"message,omitempty"`
	MessageId      string               `json:"message_id"`
}

type GenieTrashSpaceRequestPb struct {
	SpaceId string `json:"-" url:"-"`
}

type GetDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

type GetPublishedDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

type GetPublishedDashboardTokenInfoRequestPb struct {
	DashboardId      string `json:"-" url:"-"`
	ExternalValue    string `json:"-" url:"external_value,omitempty"`
	ExternalViewerId string `json:"-" url:"external_viewer_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetPublishedDashboardTokenInfoRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetPublishedDashboardTokenInfoRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetPublishedDashboardTokenInfoResponsePb struct {
	AuthorizationDetails []AuthorizationDetailsPb `json:"authorization_details,omitempty"`
	CustomClaim          string                   `json:"custom_claim,omitempty"`
	Scope                string                   `json:"scope,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetPublishedDashboardTokenInfoResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetPublishedDashboardTokenInfoResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetScheduleRequestPb struct {
	DashboardId string `json:"-" url:"-"`
	ScheduleId  string `json:"-" url:"-"`
}

type GetSubscriptionRequestPb struct {
	DashboardId    string `json:"-" url:"-"`
	ScheduleId     string `json:"-" url:"-"`
	SubscriptionId string `json:"-" url:"-"`
}

type LifecycleStatePb string

const LifecycleStateActive LifecycleStatePb = `ACTIVE`

const LifecycleStateTrashed LifecycleStatePb = `TRASHED`

type ListDashboardsRequestPb struct {
	PageSize    int             `json:"-" url:"page_size,omitempty"`
	PageToken   string          `json:"-" url:"page_token,omitempty"`
	ShowTrashed bool            `json:"-" url:"show_trashed,omitempty"`
	View        DashboardViewPb `json:"-" url:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListDashboardsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListDashboardsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListDashboardsResponsePb struct {
	Dashboards    []DashboardPb `json:"dashboards,omitempty"`
	NextPageToken string        `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListDashboardsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListDashboardsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSchedulesRequestPb struct {
	DashboardId string `json:"-" url:"-"`
	PageSize    int    `json:"-" url:"page_size,omitempty"`
	PageToken   string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListSchedulesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListSchedulesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSchedulesResponsePb struct {
	NextPageToken string       `json:"next_page_token,omitempty"`
	Schedules     []SchedulePb `json:"schedules,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListSchedulesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListSchedulesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSubscriptionsRequestPb struct {
	DashboardId string `json:"-" url:"-"`
	PageSize    int    `json:"-" url:"page_size,omitempty"`
	PageToken   string `json:"-" url:"page_token,omitempty"`
	ScheduleId  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListSubscriptionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListSubscriptionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSubscriptionsResponsePb struct {
	NextPageToken string           `json:"next_page_token,omitempty"`
	Subscriptions []SubscriptionPb `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListSubscriptionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListSubscriptionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MessageErrorPb struct {
	Error string             `json:"error,omitempty"`
	Type  MessageErrorTypePb `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MessageErrorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MessageErrorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MessageErrorTypePb string

const MessageErrorTypeBlockMultipleExecutionsException MessageErrorTypePb = `BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION`

const MessageErrorTypeChatCompletionClientException MessageErrorTypePb = `CHAT_COMPLETION_CLIENT_EXCEPTION`

const MessageErrorTypeChatCompletionClientTimeoutException MessageErrorTypePb = `CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION`

const MessageErrorTypeChatCompletionNetworkException MessageErrorTypePb = `CHAT_COMPLETION_NETWORK_EXCEPTION`

const MessageErrorTypeContentFilterException MessageErrorTypePb = `CONTENT_FILTER_EXCEPTION`

const MessageErrorTypeContextExceededException MessageErrorTypePb = `CONTEXT_EXCEEDED_EXCEPTION`

const MessageErrorTypeCouldNotGetModelDeploymentsException MessageErrorTypePb = `COULD_NOT_GET_MODEL_DEPLOYMENTS_EXCEPTION`

const MessageErrorTypeCouldNotGetUcSchemaException MessageErrorTypePb = `COULD_NOT_GET_UC_SCHEMA_EXCEPTION`

const MessageErrorTypeDeploymentNotFoundException MessageErrorTypePb = `DEPLOYMENT_NOT_FOUND_EXCEPTION`

const MessageErrorTypeDescribeQueryInvalidSqlError MessageErrorTypePb = `DESCRIBE_QUERY_INVALID_SQL_ERROR`

const MessageErrorTypeDescribeQueryTimeout MessageErrorTypePb = `DESCRIBE_QUERY_TIMEOUT`

const MessageErrorTypeDescribeQueryUnexpectedFailure MessageErrorTypePb = `DESCRIBE_QUERY_UNEXPECTED_FAILURE`

const MessageErrorTypeFunctionsNotAvailableException MessageErrorTypePb = `FUNCTIONS_NOT_AVAILABLE_EXCEPTION`

const MessageErrorTypeFunctionArgumentsInvalidException MessageErrorTypePb = `FUNCTION_ARGUMENTS_INVALID_EXCEPTION`

const MessageErrorTypeFunctionArgumentsInvalidJsonException MessageErrorTypePb = `FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION`

const MessageErrorTypeFunctionArgumentsInvalidTypeException MessageErrorTypePb = `FUNCTION_ARGUMENTS_INVALID_TYPE_EXCEPTION`

const MessageErrorTypeFunctionCallMissingParameterException MessageErrorTypePb = `FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION`

const MessageErrorTypeGeneratedSqlQueryTooLongException MessageErrorTypePb = `GENERATED_SQL_QUERY_TOO_LONG_EXCEPTION`

const MessageErrorTypeGenericChatCompletionException MessageErrorTypePb = `GENERIC_CHAT_COMPLETION_EXCEPTION`

const MessageErrorTypeGenericChatCompletionServiceException MessageErrorTypePb = `GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION`

const MessageErrorTypeGenericSqlExecApiCallException MessageErrorTypePb = `GENERIC_SQL_EXEC_API_CALL_EXCEPTION`

const MessageErrorTypeIllegalParameterDefinitionException MessageErrorTypePb = `ILLEGAL_PARAMETER_DEFINITION_EXCEPTION`

const MessageErrorTypeInvalidCertifiedAnswerFunctionException MessageErrorTypePb = `INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION`

const MessageErrorTypeInvalidCertifiedAnswerIdentifierException MessageErrorTypePb = `INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION`

const MessageErrorTypeInvalidChatCompletionArgumentsJsonException MessageErrorTypePb = `INVALID_CHAT_COMPLETION_ARGUMENTS_JSON_EXCEPTION`

const MessageErrorTypeInvalidChatCompletionJsonException MessageErrorTypePb = `INVALID_CHAT_COMPLETION_JSON_EXCEPTION`

const MessageErrorTypeInvalidCompletionRequestException MessageErrorTypePb = `INVALID_COMPLETION_REQUEST_EXCEPTION`

const MessageErrorTypeInvalidFunctionCallException MessageErrorTypePb = `INVALID_FUNCTION_CALL_EXCEPTION`

const MessageErrorTypeInvalidSqlMultipleDatasetReferencesException MessageErrorTypePb = `INVALID_SQL_MULTIPLE_DATASET_REFERENCES_EXCEPTION`

const MessageErrorTypeInvalidSqlMultipleStatementsException MessageErrorTypePb = `INVALID_SQL_MULTIPLE_STATEMENTS_EXCEPTION`

const MessageErrorTypeInvalidSqlUnknownTableException MessageErrorTypePb = `INVALID_SQL_UNKNOWN_TABLE_EXCEPTION`

const MessageErrorTypeInvalidTableIdentifierException MessageErrorTypePb = `INVALID_TABLE_IDENTIFIER_EXCEPTION`

const MessageErrorTypeLocalContextExceededException MessageErrorTypePb = `LOCAL_CONTEXT_EXCEEDED_EXCEPTION`

const MessageErrorTypeMessageCancelledWhileExecutingException MessageErrorTypePb = `MESSAGE_CANCELLED_WHILE_EXECUTING_EXCEPTION`

const MessageErrorTypeMessageDeletedWhileExecutingException MessageErrorTypePb = `MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION`

const MessageErrorTypeMessageUpdatedWhileExecutingException MessageErrorTypePb = `MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION`

const MessageErrorTypeMissingSqlQueryException MessageErrorTypePb = `MISSING_SQL_QUERY_EXCEPTION`

const MessageErrorTypeNoDeploymentsAvailableToWorkspace MessageErrorTypePb = `NO_DEPLOYMENTS_AVAILABLE_TO_WORKSPACE`

const MessageErrorTypeNoQueryToVisualizeException MessageErrorTypePb = `NO_QUERY_TO_VISUALIZE_EXCEPTION`

const MessageErrorTypeNoTablesToQueryException MessageErrorTypePb = `NO_TABLES_TO_QUERY_EXCEPTION`

const MessageErrorTypeRateLimitExceededGenericException MessageErrorTypePb = `RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION`

const MessageErrorTypeRateLimitExceededSpecifiedWaitException MessageErrorTypePb = `RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION`

const MessageErrorTypeReplyProcessTimeoutException MessageErrorTypePb = `REPLY_PROCESS_TIMEOUT_EXCEPTION`

const MessageErrorTypeRetryableProcessingException MessageErrorTypePb = `RETRYABLE_PROCESSING_EXCEPTION`

const MessageErrorTypeSqlExecutionException MessageErrorTypePb = `SQL_EXECUTION_EXCEPTION`

const MessageErrorTypeStopProcessDueToAutoRegenerate MessageErrorTypePb = `STOP_PROCESS_DUE_TO_AUTO_REGENERATE`

const MessageErrorTypeTablesMissingException MessageErrorTypePb = `TABLES_MISSING_EXCEPTION`

const MessageErrorTypeTooManyCertifiedAnswersException MessageErrorTypePb = `TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION`

const MessageErrorTypeTooManyTablesException MessageErrorTypePb = `TOO_MANY_TABLES_EXCEPTION`

const MessageErrorTypeUnexpectedReplyProcessException MessageErrorTypePb = `UNEXPECTED_REPLY_PROCESS_EXCEPTION`

const MessageErrorTypeUnknownAiModel MessageErrorTypePb = `UNKNOWN_AI_MODEL`

const MessageErrorTypeWarehouseAccessMissingException MessageErrorTypePb = `WAREHOUSE_ACCESS_MISSING_EXCEPTION`

const MessageErrorTypeWarehouseNotFoundException MessageErrorTypePb = `WAREHOUSE_NOT_FOUND_EXCEPTION`

type MessageStatusPb string

// Waiting for the LLM to respond to the user's question.
const MessageStatusAskingAi MessageStatusPb = `ASKING_AI`

// Message has been cancelled.
const MessageStatusCancelled MessageStatusPb = `CANCELLED`

// Message processing is completed. Results are in the `attachments` field. Get
// the SQL query result by calling
// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
// API.
const MessageStatusCompleted MessageStatusPb = `COMPLETED`

// Executing a generated SQL query. Get the SQL query result by calling
// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
// API.
const MessageStatusExecutingQuery MessageStatusPb = `EXECUTING_QUERY`

// The response generation or query execution failed. See `error` field.
const MessageStatusFailed MessageStatusPb = `FAILED`

// Fetching metadata from the data sources.
const MessageStatusFetchingMetadata MessageStatusPb = `FETCHING_METADATA`

// Running smart context step to determine relevant context.
const MessageStatusFilteringContext MessageStatusPb = `FILTERING_CONTEXT`

// Waiting for warehouse before the SQL query can start executing.
const MessageStatusPendingWarehouse MessageStatusPb = `PENDING_WAREHOUSE`

// SQL result is not available anymore. The user needs to rerun the query. Rerun
// the SQL query result by calling
// [executeMessageAttachmentQuery](:method:genie/executeMessageAttachmentQuery)
// API.
const MessageStatusQueryResultExpired MessageStatusPb = `QUERY_RESULT_EXPIRED`

// Message has been submitted.
const MessageStatusSubmitted MessageStatusPb = `SUBMITTED`

type MigrateDashboardRequestPb struct {
	DisplayName           string `json:"display_name,omitempty"`
	ParentPath            string `json:"parent_path,omitempty"`
	SourceDashboardId     string `json:"source_dashboard_id"`
	UpdateParameterSyntax bool   `json:"update_parameter_syntax,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MigrateDashboardRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MigrateDashboardRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PublishRequestPb struct {
	DashboardId      string `json:"-" url:"-"`
	EmbedCredentials bool   `json:"embed_credentials,omitempty"`
	WarehouseId      string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PublishRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PublishRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PublishedDashboardPb struct {
	DisplayName        string `json:"display_name,omitempty"`
	EmbedCredentials   bool   `json:"embed_credentials,omitempty"`
	RevisionCreateTime string `json:"revision_create_time,omitempty"`
	WarehouseId        string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PublishedDashboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PublishedDashboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResultPb struct {
	IsTruncated bool   `json:"is_truncated,omitempty"`
	RowCount    int64  `json:"row_count,omitempty"`
	StatementId string `json:"statement_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SchedulePb struct {
	CreateTime   string                `json:"create_time,omitempty"`
	CronSchedule CronSchedulePb        `json:"cron_schedule"`
	DashboardId  string                `json:"dashboard_id,omitempty"`
	DisplayName  string                `json:"display_name,omitempty"`
	Etag         string                `json:"etag,omitempty"`
	PauseStatus  SchedulePauseStatusPb `json:"pause_status,omitempty"`
	ScheduleId   string                `json:"schedule_id,omitempty"`
	UpdateTime   string                `json:"update_time,omitempty"`
	WarehouseId  string                `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SchedulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SchedulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SchedulePauseStatusPb string

const SchedulePauseStatusPaused SchedulePauseStatusPb = `PAUSED`

const SchedulePauseStatusUnpaused SchedulePauseStatusPb = `UNPAUSED`

type SubscriberPb struct {
	DestinationSubscriber *SubscriptionSubscriberDestinationPb `json:"destination_subscriber,omitempty"`
	UserSubscriber        *SubscriptionSubscriberUserPb        `json:"user_subscriber,omitempty"`
}

type SubscriptionPb struct {
	CreateTime      string       `json:"create_time,omitempty"`
	CreatedByUserId int64        `json:"created_by_user_id,omitempty"`
	DashboardId     string       `json:"dashboard_id,omitempty"`
	Etag            string       `json:"etag,omitempty"`
	ScheduleId      string       `json:"schedule_id,omitempty"`
	Subscriber      SubscriberPb `json:"subscriber"`
	SubscriptionId  string       `json:"subscription_id,omitempty"`
	UpdateTime      string       `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SubscriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SubscriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SubscriptionSubscriberDestinationPb struct {
	DestinationId string `json:"destination_id"`
}

type SubscriptionSubscriberUserPb struct {
	UserId int64 `json:"user_id"`
}

type TextAttachmentPb struct {
	Content string `json:"content,omitempty"`
	Id      string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TextAttachmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TextAttachmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TrashDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

type TrashDashboardResponsePb struct {
}

type UnpublishDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

type UnpublishDashboardResponsePb struct {
}

type UpdateDashboardRequestPb struct {
	Dashboard   DashboardPb `json:"dashboard"`
	DashboardId string      `json:"-" url:"-"`
}

type UpdateScheduleRequestPb struct {
	DashboardId string     `json:"-" url:"-"`
	Schedule    SchedulePb `json:"schedule"`
	ScheduleId  string     `json:"-" url:"-"`
}

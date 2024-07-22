// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/sql"
)

type CreateDashboardRequest struct {
	// The display name of the dashboard.
	DisplayName string `json:"display_name"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash.
	ParentPath string `json:"parent_path,omitempty"`
	// The contents of the dashboard in serialized string form.
	SerializedDashboard string `json:"serialized_dashboard,omitempty"`
	// The warehouse ID used to run the dashboard.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateScheduleRequest struct {
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule CronSchedule `json:"cron_schedule"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// The display name for schedule.
	DisplayName string `json:"display_name,omitempty"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus SchedulePauseStatus `json:"pause_status,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateScheduleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateScheduleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId string `json:"-" url:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId string `json:"-" url:"-"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber Subscriber `json:"subscriber"`
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
	// ensure that the dashboard has not been modified since the last read.
	Etag string `json:"etag,omitempty"`
	// The state of the dashboard resource. Used for tracking trashed status.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash.
	ParentPath string `json:"parent_path,omitempty"`
	// The workspace path of the dashboard asset, including the file name.
	Path string `json:"path,omitempty"`
	// The contents of the dashboard in serialized string form.
	SerializedDashboard string `json:"serialized_dashboard,omitempty"`
	// The timestamp of when the dashboard was last updated by the user.
	UpdateTime string `json:"update_time,omitempty"`
	// The warehouse ID used to run the dashboard.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *Dashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Dashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DashboardView string

const DashboardViewDashboardViewBasic DashboardView = `DASHBOARD_VIEW_BASIC`

const DashboardViewDashboardViewFull DashboardView = `DASHBOARD_VIEW_FULL`

// String representation for [fmt.Print]
func (f *DashboardView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DashboardView) Set(v string) error {
	switch v {
	case `DASHBOARD_VIEW_BASIC`, `DASHBOARD_VIEW_FULL`:
		*f = DashboardView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DASHBOARD_VIEW_BASIC", "DASHBOARD_VIEW_FULL"`, v)
	}
}

// Type always returns DashboardView to satisfy [pflag.Value] interface
func (f *DashboardView) Type() string {
	return "DashboardView"
}

// Delete dashboard schedule
type DeleteScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// The etag for the schedule. Optionally, it can be provided to verify that
	// the schedule has not been modified from its last retrieval.
	Etag string `json:"-" url:"etag,omitempty"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteScheduleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteScheduleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteScheduleResponse struct {
}

// Delete schedule subscription
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

	ForceSendFields []string `json:"-"`
}

func (s *DeleteSubscriptionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteSubscriptionResponse struct {
}

// Execute SQL query in a conversation message
type ExecuteMessageQueryRequest struct {
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

// Genie AI Response
type GenieAttachment struct {
	Query *QueryAttachment `json:"query,omitempty"`

	Text *TextAttachment `json:"text,omitempty"`
}

type GenieConversation struct {
	// Timestamp when the message was created
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// Conversation ID
	Id string `json:"id"`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Genie space ID
	SpaceId string `json:"space_id"`
	// Conversation title
	Title string `json:"title"`
	// ID of the user who created the conversation
	UserId int `json:"user_id"`

	ForceSendFields []string `json:"-"`
}

func (s *GenieConversation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieConversation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieCreateConversationMessageRequest struct {
	// User message content.
	Content string `json:"content"`
	// The ID associated with the conversation.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the Genie space where the conversation is started.
	SpaceId string `json:"-" url:"-"`
}

// Get conversation message
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

// Get conversation message SQL query result
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

type GenieMessage struct {
	// AI produced response to the message
	Attachments []GenieAttachment `json:"attachments,omitempty"`
	// User message content
	Content string `json:"content"`
	// Conversation ID
	ConversationId string `json:"conversation_id"`
	// Timestamp when the message was created
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// Error message if AI failed to respond to the message
	Error *MessageError `json:"error,omitempty"`
	// Message ID
	Id string `json:"id"`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The result of SQL query if the message has a query attachment
	QueryResult *Result `json:"query_result,omitempty"`
	// Genie space ID
	SpaceId string `json:"space_id"`
	// MesssageStatus. The possible values are: * `FETCHING_METADATA`: Fetching
	// metadata from the data sources. * `ASKING_AI`: Waiting for the LLM to
	// respond to the users question. * `EXECUTING_QUERY`: Executing AI provided
	// SQL query. Get the SQL query result by calling
	// [getMessageQueryResult](:method:genie/getMessageQueryResult) API. *
	// `FAILED`: Generating a response or the executing the query failed. Please
	// see `error` field. * `COMPLETED`: Message processing is completed.
	// Results are in the `attachments` field. Get the SQL query result by
	// calling [getMessageQueryResult](:method:genie/getMessageQueryResult) API.
	// * `SUBMITTED`: Message has been submitted. * `QUERY_RESULT_EXPIRED`: SQL
	// result is not available anymore. The user needs to execute the query
	// again. * `CANCELLED`: Message has been cancelled.
	Status MessageStatus `json:"status,omitempty"`
	// ID of the user who created the message
	UserId int64 `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GenieMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieMessage) MarshalJSON() ([]byte, error) {
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

// Get dashboard
type GetDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
}

// Get published dashboard
type GetPublishedDashboardRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId string `json:"-" url:"-"`
}

// Get dashboard schedule
type GetScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" url:"-"`
}

// Get schedule subscription
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

// Type always returns LifecycleState to satisfy [pflag.Value] interface
func (f *LifecycleState) Type() string {
	return "LifecycleState"
}

// List dashboards
type ListDashboardsRequest struct {
	// The number of dashboards to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListDashboards` call. This token
	// can be used to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The flag to include dashboards located in the trash. If unspecified, only
	// active dashboards will be returned.
	ShowTrashed bool `json:"-" url:"show_trashed,omitempty"`
	// Indicates whether to include all metadata from the dashboard in the
	// response. If unset, the response defaults to `DASHBOARD_VIEW_BASIC` which
	// only includes summary metadata from the dashboard.
	View DashboardView `json:"-" url:"view,omitempty"`

	ForceSendFields []string `json:"-"`
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

	ForceSendFields []string `json:"-"`
}

func (s *ListDashboardsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDashboardsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List dashboard schedules
type ListSchedulesRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// The number of schedules to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-"`
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

	ForceSendFields []string `json:"-"`
}

func (s *ListSchedulesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchedulesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List schedule subscriptions
type ListSubscriptionsRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId string `json:"-" url:"-"`
	// The number of subscriptions to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListSubscriptions` call. Use this
	// to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
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

	ForceSendFields []string `json:"-"`
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

	ForceSendFields []string `json:"-"`
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

const MessageErrorTypeCouldNotGetUcSchemaException MessageErrorType = `COULD_NOT_GET_UC_SCHEMA_EXCEPTION`

const MessageErrorTypeDeploymentNotFoundException MessageErrorType = `DEPLOYMENT_NOT_FOUND_EXCEPTION`

const MessageErrorTypeFunctionsNotAvailableException MessageErrorType = `FUNCTIONS_NOT_AVAILABLE_EXCEPTION`

const MessageErrorTypeFunctionArgumentsInvalidException MessageErrorType = `FUNCTION_ARGUMENTS_INVALID_EXCEPTION`

const MessageErrorTypeFunctionArgumentsInvalidJsonException MessageErrorType = `FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION`

const MessageErrorTypeFunctionCallMissingParameterException MessageErrorType = `FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION`

const MessageErrorTypeGenericChatCompletionException MessageErrorType = `GENERIC_CHAT_COMPLETION_EXCEPTION`

const MessageErrorTypeGenericChatCompletionServiceException MessageErrorType = `GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION`

const MessageErrorTypeGenericSqlExecApiCallException MessageErrorType = `GENERIC_SQL_EXEC_API_CALL_EXCEPTION`

const MessageErrorTypeIllegalParameterDefinitionException MessageErrorType = `ILLEGAL_PARAMETER_DEFINITION_EXCEPTION`

const MessageErrorTypeInvalidCertifiedAnswerFunctionException MessageErrorType = `INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION`

const MessageErrorTypeInvalidCertifiedAnswerIdentifierException MessageErrorType = `INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION`

const MessageErrorTypeInvalidChatCompletionJsonException MessageErrorType = `INVALID_CHAT_COMPLETION_JSON_EXCEPTION`

const MessageErrorTypeInvalidCompletionRequestException MessageErrorType = `INVALID_COMPLETION_REQUEST_EXCEPTION`

const MessageErrorTypeInvalidFunctionCallException MessageErrorType = `INVALID_FUNCTION_CALL_EXCEPTION`

const MessageErrorTypeInvalidTableIdentifierException MessageErrorType = `INVALID_TABLE_IDENTIFIER_EXCEPTION`

const MessageErrorTypeLocalContextExceededException MessageErrorType = `LOCAL_CONTEXT_EXCEEDED_EXCEPTION`

const MessageErrorTypeMessageDeletedWhileExecutingException MessageErrorType = `MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION`

const MessageErrorTypeMessageUpdatedWhileExecutingException MessageErrorType = `MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION`

const MessageErrorTypeNoTablesToQueryException MessageErrorType = `NO_TABLES_TO_QUERY_EXCEPTION`

const MessageErrorTypeRateLimitExceededGenericException MessageErrorType = `RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION`

const MessageErrorTypeRateLimitExceededSpecifiedWaitException MessageErrorType = `RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION`

const MessageErrorTypeReplyProcessTimeoutException MessageErrorType = `REPLY_PROCESS_TIMEOUT_EXCEPTION`

const MessageErrorTypeRetryableProcessingException MessageErrorType = `RETRYABLE_PROCESSING_EXCEPTION`

const MessageErrorTypeSqlExecutionException MessageErrorType = `SQL_EXECUTION_EXCEPTION`

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
	case `BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION`, `CHAT_COMPLETION_CLIENT_EXCEPTION`, `CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION`, `CHAT_COMPLETION_NETWORK_EXCEPTION`, `CONTENT_FILTER_EXCEPTION`, `CONTEXT_EXCEEDED_EXCEPTION`, `COULD_NOT_GET_UC_SCHEMA_EXCEPTION`, `DEPLOYMENT_NOT_FOUND_EXCEPTION`, `FUNCTIONS_NOT_AVAILABLE_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION`, `FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION`, `GENERIC_CHAT_COMPLETION_EXCEPTION`, `GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION`, `GENERIC_SQL_EXEC_API_CALL_EXCEPTION`, `ILLEGAL_PARAMETER_DEFINITION_EXCEPTION`, `INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION`, `INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION`, `INVALID_CHAT_COMPLETION_JSON_EXCEPTION`, `INVALID_COMPLETION_REQUEST_EXCEPTION`, `INVALID_FUNCTION_CALL_EXCEPTION`, `INVALID_TABLE_IDENTIFIER_EXCEPTION`, `LOCAL_CONTEXT_EXCEEDED_EXCEPTION`, `MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION`, `MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION`, `NO_TABLES_TO_QUERY_EXCEPTION`, `RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION`, `RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION`, `REPLY_PROCESS_TIMEOUT_EXCEPTION`, `RETRYABLE_PROCESSING_EXCEPTION`, `SQL_EXECUTION_EXCEPTION`, `TABLES_MISSING_EXCEPTION`, `TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION`, `TOO_MANY_TABLES_EXCEPTION`, `UNEXPECTED_REPLY_PROCESS_EXCEPTION`, `UNKNOWN_AI_MODEL`, `WAREHOUSE_ACCESS_MISSING_EXCEPTION`, `WAREHOUSE_NOT_FOUND_EXCEPTION`:
		*f = MessageErrorType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION", "CHAT_COMPLETION_CLIENT_EXCEPTION", "CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION", "CHAT_COMPLETION_NETWORK_EXCEPTION", "CONTENT_FILTER_EXCEPTION", "CONTEXT_EXCEEDED_EXCEPTION", "COULD_NOT_GET_UC_SCHEMA_EXCEPTION", "DEPLOYMENT_NOT_FOUND_EXCEPTION", "FUNCTIONS_NOT_AVAILABLE_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION", "FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION", "GENERIC_CHAT_COMPLETION_EXCEPTION", "GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION", "GENERIC_SQL_EXEC_API_CALL_EXCEPTION", "ILLEGAL_PARAMETER_DEFINITION_EXCEPTION", "INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION", "INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION", "INVALID_CHAT_COMPLETION_JSON_EXCEPTION", "INVALID_COMPLETION_REQUEST_EXCEPTION", "INVALID_FUNCTION_CALL_EXCEPTION", "INVALID_TABLE_IDENTIFIER_EXCEPTION", "LOCAL_CONTEXT_EXCEEDED_EXCEPTION", "MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION", "MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION", "NO_TABLES_TO_QUERY_EXCEPTION", "RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION", "RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION", "REPLY_PROCESS_TIMEOUT_EXCEPTION", "RETRYABLE_PROCESSING_EXCEPTION", "SQL_EXECUTION_EXCEPTION", "TABLES_MISSING_EXCEPTION", "TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION", "TOO_MANY_TABLES_EXCEPTION", "UNEXPECTED_REPLY_PROCESS_EXCEPTION", "UNKNOWN_AI_MODEL", "WAREHOUSE_ACCESS_MISSING_EXCEPTION", "WAREHOUSE_NOT_FOUND_EXCEPTION"`, v)
	}
}

// Type always returns MessageErrorType to satisfy [pflag.Value] interface
func (f *MessageErrorType) Type() string {
	return "MessageErrorType"
}

// MesssageStatus. The possible values are: * `FETCHING_METADATA`: Fetching
// metadata from the data sources. * `ASKING_AI`: Waiting for the LLM to respond
// to the users question. * `EXECUTING_QUERY`: Executing AI provided SQL query.
// Get the SQL query result by calling
// [getMessageQueryResult](:method:genie/getMessageQueryResult) API. * `FAILED`:
// Generating a response or the executing the query failed. Please see `error`
// field. * `COMPLETED`: Message processing is completed. Results are in the
// `attachments` field. Get the SQL query result by calling
// [getMessageQueryResult](:method:genie/getMessageQueryResult) API. *
// `SUBMITTED`: Message has been submitted. * `QUERY_RESULT_EXPIRED`: SQL result
// is not available anymore. The user needs to execute the query again. *
// `CANCELLED`: Message has been cancelled.
type MessageStatus string

// Waiting for the LLM to respond to the users question.
const MessageStatusAskingAi MessageStatus = `ASKING_AI`

// Message has been cancelled.
const MessageStatusCancelled MessageStatus = `CANCELLED`

// Message processing is completed. Results are in the `attachments` field. Get
// the SQL query result by calling
// [getMessageQueryResult](:method:genie/getMessageQueryResult) API.
const MessageStatusCompleted MessageStatus = `COMPLETED`

// Executing AI provided SQL query. Get the SQL query result by calling
// [getMessageQueryResult](:method:genie/getMessageQueryResult) API.
const MessageStatusExecutingQuery MessageStatus = `EXECUTING_QUERY`

// Generating a response or the executing the query failed. Please see `error`
// field.
const MessageStatusFailed MessageStatus = `FAILED`

// Fetching metadata from the data sources.
const MessageStatusFetchingMetadata MessageStatus = `FETCHING_METADATA`

// SQL result is not available anymore. The user needs to execute the query
// again.
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
	case `ASKING_AI`, `CANCELLED`, `COMPLETED`, `EXECUTING_QUERY`, `FAILED`, `FETCHING_METADATA`, `QUERY_RESULT_EXPIRED`, `SUBMITTED`:
		*f = MessageStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASKING_AI", "CANCELLED", "COMPLETED", "EXECUTING_QUERY", "FAILED", "FETCHING_METADATA", "QUERY_RESULT_EXPIRED", "SUBMITTED"`, v)
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

	ForceSendFields []string `json:"-"`
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

	ForceSendFields []string `json:"-"`
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

	ForceSendFields []string `json:"-"`
}

func (s *PublishedDashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishedDashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryAttachment struct {
	// Description of the query
	Description string `json:"description,omitempty"`
	// If the query was created on an instruction (trusted asset) we link to the
	// id
	InstructionId string `json:"instruction_id,omitempty"`
	// Always store the title next to the id in case the original instruction
	// title changes or the instruction is deleted.
	InstructionTitle string `json:"instruction_title,omitempty"`
	// Time when the user updated the query last
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// AI generated SQL query
	Query string `json:"query,omitempty"`
	// Name of the query
	Title string `json:"title,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *QueryAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Result struct {
	// Row count of the result
	RowCount int64 `json:"row_count,omitempty"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId string `json:"statement_id,omitempty"`

	ForceSendFields []string `json:"-"`
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

	ForceSendFields []string `json:"-"`
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

	ForceSendFields []string `json:"-"`
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

	ForceSendFields []string `json:"-"`
}

func (s *TextAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TextAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Trash dashboard
type TrashDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
}

type TrashDashboardResponse struct {
}

// Unpublish dashboard
type UnpublishDashboardRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId string `json:"-" url:"-"`
}

type UnpublishDashboardResponse struct {
}

type UpdateDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
	// The display name of the dashboard.
	DisplayName string `json:"display_name,omitempty"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read.
	Etag string `json:"etag,omitempty"`
	// The contents of the dashboard in serialized string form.
	SerializedDashboard string `json:"serialized_dashboard,omitempty"`
	// The warehouse ID used to run the dashboard.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateScheduleRequest struct {
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule CronSchedule `json:"cron_schedule"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// The display name for schedule.
	DisplayName string `json:"display_name,omitempty"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag string `json:"etag,omitempty"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus SchedulePauseStatus `json:"pause_status,omitempty"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdateScheduleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateScheduleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/sql"
)

type AuthorizationDetails struct {
	// Represents downscoped permission rules with specific access rights. This
	// field is specific to `workspace_rule_set` constraint.
	// Wire name: 'grant_rules'
	GrantRules []AuthorizationDetailsGrantRule
	// The acl path of the tree store resource resource.
	// Wire name: 'resource_legacy_acl_path'
	ResourceLegacyAclPath string
	// The resource name to which the authorization rule applies. This field is
	// specific to `workspace_rule_set` constraint. Format:
	// `workspaces/{workspace_id}/dashboards/{dashboard_id}`
	// Wire name: 'resource_name'
	ResourceName string
	// The type of authorization downscoping policy. Ex: `workspace_rule_set`
	// defines access rules for a specific workspace resource
	// Wire name: 'type'
	Type string

	ForceSendFields []string `tf:"-"`
}

func (st *AuthorizationDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &authorizationDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := authorizationDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AuthorizationDetails) MarshalJSON() ([]byte, error) {
	pb, err := authorizationDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AuthorizationDetailsGrantRule struct {
	// Permission sets for dashboard are defined in
	// iam-common/rbac-common/permission-sets/definitions/TreeStoreBasePermissionSets
	// Ex: `permissionSets/dashboard.runner`
	// Wire name: 'permission_set'
	PermissionSet string

	ForceSendFields []string `tf:"-"`
}

func (st *AuthorizationDetailsGrantRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &authorizationDetailsGrantRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := authorizationDetailsGrantRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AuthorizationDetailsGrantRule) MarshalJSON() ([]byte, error) {
	pb, err := authorizationDetailsGrantRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Cancel the results for the a query for a published, embedded dashboard
type CancelPublishedQueryExecutionRequest struct {

	// Wire name: 'dashboard_name'
	DashboardName string `tf:"-"`

	// Wire name: 'dashboard_revision_id'
	DashboardRevisionId string `tf:"-"`
	// Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	// Wire name: 'tokens'
	Tokens []string `tf:"-"`
}

func (st *CancelPublishedQueryExecutionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelPublishedQueryExecutionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelPublishedQueryExecutionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelPublishedQueryExecutionRequest) MarshalJSON() ([]byte, error) {
	pb, err := cancelPublishedQueryExecutionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CancelQueryExecutionResponse struct {

	// Wire name: 'status'
	Status []CancelQueryExecutionResponseStatus
}

func (st *CancelQueryExecutionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelQueryExecutionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelQueryExecutionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelQueryExecutionResponse) MarshalJSON() ([]byte, error) {
	pb, err := cancelQueryExecutionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CancelQueryExecutionResponseStatus struct {
	// The token to poll for result asynchronously Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	// Wire name: 'data_token'
	DataToken string
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	// Wire name: 'pending'
	Pending *Empty
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	// Wire name: 'success'
	Success *Empty
}

func (st *CancelQueryExecutionResponseStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelQueryExecutionResponseStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelQueryExecutionResponseStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelQueryExecutionResponseStatus) MarshalJSON() ([]byte, error) {
	pb, err := cancelQueryExecutionResponseStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Create dashboard
type CreateDashboardRequest struct {

	// Wire name: 'dashboard'
	Dashboard Dashboard
}

func (st *CreateDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := createDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Create dashboard schedule
type CreateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`

	// Wire name: 'schedule'
	Schedule Schedule
}

func (st *CreateScheduleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createScheduleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createScheduleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateScheduleRequest) MarshalJSON() ([]byte, error) {
	pb, err := createScheduleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Create schedule subscription
type CreateSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	// Wire name: 'schedule_id'
	ScheduleId string `tf:"-"`

	// Wire name: 'subscription'
	Subscription Subscription
}

func (st *CreateSubscriptionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createSubscriptionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createSubscriptionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateSubscriptionRequest) MarshalJSON() ([]byte, error) {
	pb, err := createSubscriptionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CronSchedule struct {
	// A cron expression using quartz syntax. EX: `0 0 8 * * ?` represents
	// everyday at 8am. See [Cron Trigger] for details.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	// Wire name: 'quartz_cron_expression'
	QuartzCronExpression string
	// A Java timezone id. The schedule will be resolved with respect to this
	// timezone. See [Java TimeZone] for details.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	// Wire name: 'timezone_id'
	TimezoneId string
}

func (st *CronSchedule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cronSchedulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cronScheduleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CronSchedule) MarshalJSON() ([]byte, error) {
	pb, err := cronScheduleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Dashboard struct {
	// The timestamp of when the dashboard was created.
	// Wire name: 'create_time'
	CreateTime string
	// UUID identifying the dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string
	// The display name of the dashboard.
	// Wire name: 'display_name'
	DisplayName string
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read. This
	// field is excluded in List Dashboards responses.
	// Wire name: 'etag'
	Etag string
	// The state of the dashboard resource. Used for tracking trashed status.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash. This field is excluded in List
	// Dashboards responses.
	// Wire name: 'parent_path'
	ParentPath string
	// The workspace path of the dashboard asset, including the file name.
	// Exported dashboards always have the file extension `.lvdash.json`. This
	// field is excluded in List Dashboards responses.
	// Wire name: 'path'
	Path string
	// The contents of the dashboard in serialized string form. This field is
	// excluded in List Dashboards responses. Use the [get dashboard API] to
	// retrieve an example response, which includes the `serialized_dashboard`
	// field. This field provides the structure of the JSON string that
	// represents the dashboard's layout and components.
	//
	// [get dashboard API]: https://docs.databricks.com/api/workspace/lakeview/get
	// Wire name: 'serialized_dashboard'
	SerializedDashboard string
	// The timestamp of when the dashboard was last updated by the user. This
	// field is excluded in List Dashboards responses.
	// Wire name: 'update_time'
	UpdateTime string
	// The warehouse ID used to run the dashboard.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func (st *Dashboard) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dashboardPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dashboardFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Dashboard) MarshalJSON() ([]byte, error) {
	pb, err := dashboardToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DashboardView string
type dashboardViewPb string

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

// Type always returns DashboardView to satisfy [pflag.Value] interface
func (f *DashboardView) Type() string {
	return "DashboardView"
}

func dashboardViewToPb(st *DashboardView) (*dashboardViewPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dashboardViewPb(*st)
	return &pb, nil
}

func dashboardViewFromPb(pb *dashboardViewPb) (*DashboardView, error) {
	if pb == nil {
		return nil, nil
	}
	st := DashboardView(*pb)
	return &st, nil
}

// Delete dashboard schedule
type DeleteScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// The etag for the schedule. Optionally, it can be provided to verify that
	// the schedule has not been modified from its last retrieval.
	// Wire name: 'etag'
	Etag string `tf:"-"`
	// UUID identifying the schedule.
	// Wire name: 'schedule_id'
	ScheduleId string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteScheduleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteScheduleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteScheduleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteScheduleRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteScheduleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteScheduleResponse struct {
}

func (st *DeleteScheduleResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteScheduleResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteScheduleResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteScheduleResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteScheduleResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete schedule subscription
type DeleteSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// The etag for the subscription. Can be optionally provided to ensure that
	// the subscription has not been modified since the last read.
	// Wire name: 'etag'
	Etag string `tf:"-"`
	// UUID identifying the schedule which the subscription belongs.
	// Wire name: 'schedule_id'
	ScheduleId string `tf:"-"`
	// UUID identifying the subscription.
	// Wire name: 'subscription_id'
	SubscriptionId string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteSubscriptionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteSubscriptionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteSubscriptionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteSubscriptionRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteSubscriptionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteSubscriptionResponse struct {
}

func (st *DeleteSubscriptionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteSubscriptionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteSubscriptionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteSubscriptionResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteSubscriptionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Represents an empty message, similar to google.protobuf.Empty, which is not
// available in the firm right now.
type Empty struct {
}

func (st *Empty) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &emptyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := emptyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Empty) MarshalJSON() ([]byte, error) {
	pb, err := emptyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Execute query request for published Dashboards. Since published dashboards
// have the option of running as the publisher, the datasets, warehouse_id are
// excluded from the request and instead read from the source (lakeview-config)
// via the additional parameters (dashboardName and dashboardRevisionId)
type ExecutePublishedDashboardQueryRequest struct {
	// Dashboard name and revision_id is required to retrieve
	// PublishedDatasetDataModel which contains the list of datasets,
	// warehouse_id, and embedded_credentials
	// Wire name: 'dashboard_name'
	DashboardName string

	// Wire name: 'dashboard_revision_id'
	DashboardRevisionId string
	// A dashboard schedule can override the warehouse used as compute for
	// processing the published dashboard queries
	// Wire name: 'override_warehouse_id'
	OverrideWarehouseId string

	ForceSendFields []string `tf:"-"`
}

func (st *ExecutePublishedDashboardQueryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &executePublishedDashboardQueryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := executePublishedDashboardQueryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExecutePublishedDashboardQueryRequest) MarshalJSON() ([]byte, error) {
	pb, err := executePublishedDashboardQueryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExecuteQueryResponse struct {
}

func (st *ExecuteQueryResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &executeQueryResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := executeQueryResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExecuteQueryResponse) MarshalJSON() ([]byte, error) {
	pb, err := executeQueryResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Genie AI Response
type GenieAttachment struct {
	// Attachment ID
	// Wire name: 'attachment_id'
	AttachmentId string
	// Query Attachment if Genie responds with a SQL query
	// Wire name: 'query'
	Query *GenieQueryAttachment
	// Text Attachment if Genie responds with text
	// Wire name: 'text'
	Text *TextAttachment

	ForceSendFields []string `tf:"-"`
}

func (st *GenieAttachment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieAttachmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieAttachmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieAttachment) MarshalJSON() ([]byte, error) {
	pb, err := genieAttachmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieConversation struct {
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string
	// Timestamp when the message was created
	// Wire name: 'created_timestamp'
	CreatedTimestamp int64
	// Conversation ID. Legacy identifier, use conversation_id instead
	// Wire name: 'id'
	Id string
	// Timestamp when the message was last updated
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string
	// Conversation title
	// Wire name: 'title'
	Title string
	// ID of the user who created the conversation
	// Wire name: 'user_id'
	UserId int

	ForceSendFields []string `tf:"-"`
}

func (st *GenieConversation) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieConversationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieConversationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieConversation) MarshalJSON() ([]byte, error) {
	pb, err := genieConversationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieCreateConversationMessageRequest struct {
	// User message content.
	// Wire name: 'content'
	Content string
	// The ID associated with the conversation.
	// Wire name: 'conversation_id'
	ConversationId string `tf:"-"`
	// The ID associated with the Genie space where the conversation is started.
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func (st *GenieCreateConversationMessageRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieCreateConversationMessageRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieCreateConversationMessageRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieCreateConversationMessageRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieCreateConversationMessageRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Execute message attachment SQL query
type GenieExecuteMessageAttachmentQueryRequest struct {
	// Attachment ID
	// Wire name: 'attachment_id'
	AttachmentId string `tf:"-"`
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string `tf:"-"`
	// Message ID
	// Wire name: 'message_id'
	MessageId string `tf:"-"`
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func (st *GenieExecuteMessageAttachmentQueryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieExecuteMessageAttachmentQueryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieExecuteMessageAttachmentQueryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieExecuteMessageAttachmentQueryRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieExecuteMessageAttachmentQueryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// [Deprecated] Execute SQL query in a conversation message
type GenieExecuteMessageQueryRequest struct {
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string `tf:"-"`
	// Message ID
	// Wire name: 'message_id'
	MessageId string `tf:"-"`
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func (st *GenieExecuteMessageQueryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieExecuteMessageQueryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieExecuteMessageQueryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieExecuteMessageQueryRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieExecuteMessageQueryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Generate full query result download
type GenieGenerateDownloadFullQueryResultRequest struct {
	// Attachment ID
	// Wire name: 'attachment_id'
	AttachmentId string `tf:"-"`
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string `tf:"-"`
	// Message ID
	// Wire name: 'message_id'
	MessageId string `tf:"-"`
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func (st *GenieGenerateDownloadFullQueryResultRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieGenerateDownloadFullQueryResultRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieGenerateDownloadFullQueryResultRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieGenerateDownloadFullQueryResultRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieGenerateDownloadFullQueryResultRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieGenerateDownloadFullQueryResultResponse struct {
	// Download ID. Use this ID to track the download request in subsequent
	// polling calls
	// Wire name: 'download_id'
	DownloadId string

	ForceSendFields []string `tf:"-"`
}

func (st *GenieGenerateDownloadFullQueryResultResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieGenerateDownloadFullQueryResultResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieGenerateDownloadFullQueryResultResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieGenerateDownloadFullQueryResultResponse) MarshalJSON() ([]byte, error) {
	pb, err := genieGenerateDownloadFullQueryResultResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get conversation message
type GenieGetConversationMessageRequest struct {
	// The ID associated with the target conversation.
	// Wire name: 'conversation_id'
	ConversationId string `tf:"-"`
	// The ID associated with the target message from the identified
	// conversation.
	// Wire name: 'message_id'
	MessageId string `tf:"-"`
	// The ID associated with the Genie space where the target conversation is
	// located.
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func (st *GenieGetConversationMessageRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieGetConversationMessageRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieGetConversationMessageRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieGetConversationMessageRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieGetConversationMessageRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get download full query result
type GenieGetDownloadFullQueryResultRequest struct {
	// Attachment ID
	// Wire name: 'attachment_id'
	AttachmentId string `tf:"-"`
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string `tf:"-"`
	// Download ID. This ID is provided by the [Generate Download
	// endpoint](:method:genie/generateDownloadFullQueryResult)
	// Wire name: 'download_id'
	DownloadId string `tf:"-"`
	// Message ID
	// Wire name: 'message_id'
	MessageId string `tf:"-"`
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func (st *GenieGetDownloadFullQueryResultRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieGetDownloadFullQueryResultRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieGetDownloadFullQueryResultRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieGetDownloadFullQueryResultRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieGetDownloadFullQueryResultRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieGetDownloadFullQueryResultResponse struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	// Wire name: 'statement_response'
	StatementResponse *sql.StatementResponse
}

func (st *GenieGetDownloadFullQueryResultResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieGetDownloadFullQueryResultResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieGetDownloadFullQueryResultResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieGetDownloadFullQueryResultResponse) MarshalJSON() ([]byte, error) {
	pb, err := genieGetDownloadFullQueryResultResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get message attachment SQL query result
type GenieGetMessageAttachmentQueryResultRequest struct {
	// Attachment ID
	// Wire name: 'attachment_id'
	AttachmentId string `tf:"-"`
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string `tf:"-"`
	// Message ID
	// Wire name: 'message_id'
	MessageId string `tf:"-"`
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func (st *GenieGetMessageAttachmentQueryResultRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieGetMessageAttachmentQueryResultRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieGetMessageAttachmentQueryResultRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieGetMessageAttachmentQueryResultRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieGetMessageAttachmentQueryResultRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// [Deprecated] Get conversation message SQL query result
type GenieGetMessageQueryResultRequest struct {
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string `tf:"-"`
	// Message ID
	// Wire name: 'message_id'
	MessageId string `tf:"-"`
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func (st *GenieGetMessageQueryResultRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieGetMessageQueryResultRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieGetMessageQueryResultRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieGetMessageQueryResultRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieGetMessageQueryResultRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieGetMessageQueryResultResponse struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	// Wire name: 'statement_response'
	StatementResponse *sql.StatementResponse
}

func (st *GenieGetMessageQueryResultResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieGetMessageQueryResultResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieGetMessageQueryResultResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieGetMessageQueryResultResponse) MarshalJSON() ([]byte, error) {
	pb, err := genieGetMessageQueryResultResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// [Deprecated] Get conversation message SQL query result
type GenieGetQueryResultByAttachmentRequest struct {
	// Attachment ID
	// Wire name: 'attachment_id'
	AttachmentId string `tf:"-"`
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string `tf:"-"`
	// Message ID
	// Wire name: 'message_id'
	MessageId string `tf:"-"`
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func (st *GenieGetQueryResultByAttachmentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieGetQueryResultByAttachmentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieGetQueryResultByAttachmentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieGetQueryResultByAttachmentRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieGetQueryResultByAttachmentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get Genie Space
type GenieGetSpaceRequest struct {
	// The ID associated with the Genie space
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func (st *GenieGetSpaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieGetSpaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieGetSpaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieGetSpaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieGetSpaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieMessage struct {
	// AI-generated response to the message
	// Wire name: 'attachments'
	Attachments []GenieAttachment
	// User message content
	// Wire name: 'content'
	Content string
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string
	// Timestamp when the message was created
	// Wire name: 'created_timestamp'
	CreatedTimestamp int64
	// Error message if Genie failed to respond to the message
	// Wire name: 'error'
	Error *MessageError
	// Message ID. Legacy identifier, use message_id instead
	// Wire name: 'id'
	Id string
	// Timestamp when the message was last updated
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Message ID
	// Wire name: 'message_id'
	MessageId string
	// The result of SQL query if the message includes a query attachment.
	// Deprecated. Use `query_result_metadata` in `GenieQueryAttachment`
	// instead.
	// Wire name: 'query_result'
	QueryResult *Result
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string
	// MessageStatus. The possible values are: * `FETCHING_METADATA`: Fetching
	// metadata from the data sources. * `FILTERING_CONTEXT`: Running smart
	// context step to determine relevant context. * `ASKING_AI`: Waiting for
	// the LLM to respond to the user's question. * `PENDING_WAREHOUSE`: Waiting
	// for warehouse before the SQL query can start executing. *
	// `EXECUTING_QUERY`: Executing a generated SQL query. Get the SQL query
	// result by calling
	// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
	// API. * `FAILED`: The response generation or query execution failed. See
	// `error` field. * `COMPLETED`: Message processing is completed. Results
	// are in the `attachments` field. Get the SQL query result by calling
	// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
	// API. * `SUBMITTED`: Message has been submitted. * `QUERY_RESULT_EXPIRED`:
	// SQL result is not available anymore. The user needs to rerun the query.
	// Rerun the SQL query result by calling
	// [executeMessageAttachmentQuery](:method:genie/executeMessageAttachmentQuery)
	// API. * `CANCELLED`: Message has been cancelled.
	// Wire name: 'status'
	Status MessageStatus
	// ID of the user who created the message
	// Wire name: 'user_id'
	UserId int64

	ForceSendFields []string `tf:"-"`
}

func (st *GenieMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieMessage) MarshalJSON() ([]byte, error) {
	pb, err := genieMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieQueryAttachment struct {
	// Description of the query
	// Wire name: 'description'
	Description string

	// Wire name: 'id'
	Id string
	// Time when the user updated the query last
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// AI generated SQL query
	// Wire name: 'query'
	Query string
	// Metadata associated with the query result.
	// Wire name: 'query_result_metadata'
	QueryResultMetadata *GenieResultMetadata
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	// Wire name: 'statement_id'
	StatementId string
	// Name of the query
	// Wire name: 'title'
	Title string

	ForceSendFields []string `tf:"-"`
}

func (st *GenieQueryAttachment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieQueryAttachmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieQueryAttachmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieQueryAttachment) MarshalJSON() ([]byte, error) {
	pb, err := genieQueryAttachmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieResultMetadata struct {
	// Indicates whether the result set is truncated.
	// Wire name: 'is_truncated'
	IsTruncated bool
	// The number of rows in the result set.
	// Wire name: 'row_count'
	RowCount int64

	ForceSendFields []string `tf:"-"`
}

func (st *GenieResultMetadata) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieResultMetadataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieResultMetadataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieResultMetadata) MarshalJSON() ([]byte, error) {
	pb, err := genieResultMetadataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieSpace struct {
	// Description of the Genie Space
	// Wire name: 'description'
	Description string
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string
	// Title of the Genie Space
	// Wire name: 'title'
	Title string

	ForceSendFields []string `tf:"-"`
}

func (st *GenieSpace) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieSpacePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieSpaceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieSpace) MarshalJSON() ([]byte, error) {
	pb, err := genieSpaceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieStartConversationMessageRequest struct {
	// The text of the message that starts the conversation.
	// Wire name: 'content'
	Content string
	// The ID associated with the Genie space where you want to start a
	// conversation.
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func (st *GenieStartConversationMessageRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieStartConversationMessageRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieStartConversationMessageRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieStartConversationMessageRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieStartConversationMessageRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieStartConversationResponse struct {

	// Wire name: 'conversation'
	Conversation *GenieConversation
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string

	// Wire name: 'message'
	Message *GenieMessage
	// Message ID
	// Wire name: 'message_id'
	MessageId string
}

func (st *GenieStartConversationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieStartConversationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieStartConversationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieStartConversationResponse) MarshalJSON() ([]byte, error) {
	pb, err := genieStartConversationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get dashboard
type GetDashboardRequest struct {
	// UUID identifying the dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func (st *GetDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Read a published dashboard in an embedded ui.
type GetPublishedDashboardEmbeddedRequest struct {
	// UUID identifying the published dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func (st *GetPublishedDashboardEmbeddedRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPublishedDashboardEmbeddedRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPublishedDashboardEmbeddedRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPublishedDashboardEmbeddedRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPublishedDashboardEmbeddedRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetPublishedDashboardEmbeddedResponse struct {
}

func (st *GetPublishedDashboardEmbeddedResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPublishedDashboardEmbeddedResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPublishedDashboardEmbeddedResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPublishedDashboardEmbeddedResponse) MarshalJSON() ([]byte, error) {
	pb, err := getPublishedDashboardEmbeddedResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get published dashboard
type GetPublishedDashboardRequest struct {
	// UUID identifying the published dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func (st *GetPublishedDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPublishedDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPublishedDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPublishedDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPublishedDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Read an information of a published dashboard to mint an OAuth token.
type GetPublishedDashboardTokenInfoRequest struct {
	// UUID identifying the published dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// Provided external value to be included in the custom claim.
	// Wire name: 'external_value'
	ExternalValue string `tf:"-"`
	// Provided external viewer id to be included in the custom claim.
	// Wire name: 'external_viewer_id'
	ExternalViewerId string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetPublishedDashboardTokenInfoRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPublishedDashboardTokenInfoRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPublishedDashboardTokenInfoRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPublishedDashboardTokenInfoRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPublishedDashboardTokenInfoRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetPublishedDashboardTokenInfoResponse struct {
	// Authorization constraints for accessing the published dashboard.
	// Currently includes `workspace_rule_set` and could be enriched with
	// `unity_catalog_privileges` before oAuth token generation.
	// Wire name: 'authorization_details'
	AuthorizationDetails []AuthorizationDetails
	// Custom claim generated from external_value and external_viewer_id.
	// Format:
	// `urn:aibi:external_data:<external_value>:<external_viewer_id>:<dashboard_id>`
	// Wire name: 'custom_claim'
	CustomClaim string
	// Scope defining access permissions.
	// Wire name: 'scope'
	Scope string

	ForceSendFields []string `tf:"-"`
}

func (st *GetPublishedDashboardTokenInfoResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPublishedDashboardTokenInfoResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPublishedDashboardTokenInfoResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPublishedDashboardTokenInfoResponse) MarshalJSON() ([]byte, error) {
	pb, err := getPublishedDashboardTokenInfoResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get dashboard schedule
type GetScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// UUID identifying the schedule.
	// Wire name: 'schedule_id'
	ScheduleId string `tf:"-"`
}

func (st *GetScheduleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getScheduleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getScheduleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetScheduleRequest) MarshalJSON() ([]byte, error) {
	pb, err := getScheduleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get schedule subscription
type GetSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// UUID identifying the schedule which the subscription belongs.
	// Wire name: 'schedule_id'
	ScheduleId string `tf:"-"`
	// UUID identifying the subscription.
	// Wire name: 'subscription_id'
	SubscriptionId string `tf:"-"`
}

func (st *GetSubscriptionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getSubscriptionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getSubscriptionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetSubscriptionRequest) MarshalJSON() ([]byte, error) {
	pb, err := getSubscriptionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LifecycleState string
type lifecycleStatePb string

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

func lifecycleStateToPb(st *LifecycleState) (*lifecycleStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := lifecycleStatePb(*st)
	return &pb, nil
}

func lifecycleStateFromPb(pb *lifecycleStatePb) (*LifecycleState, error) {
	if pb == nil {
		return nil, nil
	}
	st := LifecycleState(*pb)
	return &st, nil
}

// List dashboards
type ListDashboardsRequest struct {
	// The number of dashboards to return per page.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// A page token, received from a previous `ListDashboards` call. This token
	// can be used to retrieve the subsequent page.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// The flag to include dashboards located in the trash. If unspecified, only
	// active dashboards will be returned.
	// Wire name: 'show_trashed'
	ShowTrashed bool `tf:"-"`
	// `DASHBOARD_VIEW_BASIC`only includes summary metadata from the dashboard.
	// Wire name: 'view'
	View DashboardView `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListDashboardsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listDashboardsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listDashboardsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListDashboardsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listDashboardsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListDashboardsResponse struct {

	// Wire name: 'dashboards'
	Dashboards []Dashboard
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func (st *ListDashboardsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listDashboardsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listDashboardsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListDashboardsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listDashboardsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List dashboard schedules
type ListSchedulesRequest struct {
	// UUID identifying the dashboard to which the schedules belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// The number of schedules to return per page.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListSchedulesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSchedulesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSchedulesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSchedulesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listSchedulesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListSchedulesResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent schedules.
	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'schedules'
	Schedules []Schedule

	ForceSendFields []string `tf:"-"`
}

func (st *ListSchedulesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSchedulesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSchedulesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSchedulesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listSchedulesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List schedule subscriptions
type ListSubscriptionsRequest struct {
	// UUID identifying the dashboard which the subscriptions belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// The number of subscriptions to return per page.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// A page token, received from a previous `ListSubscriptions` call. Use this
	// to retrieve the subsequent page.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// UUID identifying the schedule which the subscriptions belongs.
	// Wire name: 'schedule_id'
	ScheduleId string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSubscriptionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSubscriptionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listSubscriptionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListSubscriptionsResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent subscriptions.
	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'subscriptions'
	Subscriptions []Subscription

	ForceSendFields []string `tf:"-"`
}

func (st *ListSubscriptionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSubscriptionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSubscriptionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSubscriptionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listSubscriptionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type MessageError struct {

	// Wire name: 'error'
	Error string

	// Wire name: 'type'
	Type MessageErrorType

	ForceSendFields []string `tf:"-"`
}

func (st *MessageError) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &messageErrorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := messageErrorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MessageError) MarshalJSON() ([]byte, error) {
	pb, err := messageErrorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type MessageErrorType string
type messageErrorTypePb string

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
	case `BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION`, `CHAT_COMPLETION_CLIENT_EXCEPTION`, `CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION`, `CHAT_COMPLETION_NETWORK_EXCEPTION`, `CONTENT_FILTER_EXCEPTION`, `CONTEXT_EXCEEDED_EXCEPTION`, `COULD_NOT_GET_MODEL_DEPLOYMENTS_EXCEPTION`, `COULD_NOT_GET_UC_SCHEMA_EXCEPTION`, `DEPLOYMENT_NOT_FOUND_EXCEPTION`, `DESCRIBE_QUERY_INVALID_SQL_ERROR`, `DESCRIBE_QUERY_TIMEOUT`, `DESCRIBE_QUERY_UNEXPECTED_FAILURE`, `FUNCTIONS_NOT_AVAILABLE_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_TYPE_EXCEPTION`, `FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION`, `GENERATED_SQL_QUERY_TOO_LONG_EXCEPTION`, `GENERIC_CHAT_COMPLETION_EXCEPTION`, `GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION`, `GENERIC_SQL_EXEC_API_CALL_EXCEPTION`, `ILLEGAL_PARAMETER_DEFINITION_EXCEPTION`, `INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION`, `INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION`, `INVALID_CHAT_COMPLETION_ARGUMENTS_JSON_EXCEPTION`, `INVALID_CHAT_COMPLETION_JSON_EXCEPTION`, `INVALID_COMPLETION_REQUEST_EXCEPTION`, `INVALID_FUNCTION_CALL_EXCEPTION`, `INVALID_SQL_MULTIPLE_DATASET_REFERENCES_EXCEPTION`, `INVALID_SQL_MULTIPLE_STATEMENTS_EXCEPTION`, `INVALID_SQL_UNKNOWN_TABLE_EXCEPTION`, `INVALID_TABLE_IDENTIFIER_EXCEPTION`, `LOCAL_CONTEXT_EXCEEDED_EXCEPTION`, `MESSAGE_CANCELLED_WHILE_EXECUTING_EXCEPTION`, `MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION`, `MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION`, `MISSING_SQL_QUERY_EXCEPTION`, `NO_DEPLOYMENTS_AVAILABLE_TO_WORKSPACE`, `NO_QUERY_TO_VISUALIZE_EXCEPTION`, `NO_TABLES_TO_QUERY_EXCEPTION`, `RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION`, `RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION`, `REPLY_PROCESS_TIMEOUT_EXCEPTION`, `RETRYABLE_PROCESSING_EXCEPTION`, `SQL_EXECUTION_EXCEPTION`, `STOP_PROCESS_DUE_TO_AUTO_REGENERATE`, `TABLES_MISSING_EXCEPTION`, `TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION`, `TOO_MANY_TABLES_EXCEPTION`, `UNEXPECTED_REPLY_PROCESS_EXCEPTION`, `UNKNOWN_AI_MODEL`, `WAREHOUSE_ACCESS_MISSING_EXCEPTION`, `WAREHOUSE_NOT_FOUND_EXCEPTION`:
		*f = MessageErrorType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION", "CHAT_COMPLETION_CLIENT_EXCEPTION", "CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION", "CHAT_COMPLETION_NETWORK_EXCEPTION", "CONTENT_FILTER_EXCEPTION", "CONTEXT_EXCEEDED_EXCEPTION", "COULD_NOT_GET_MODEL_DEPLOYMENTS_EXCEPTION", "COULD_NOT_GET_UC_SCHEMA_EXCEPTION", "DEPLOYMENT_NOT_FOUND_EXCEPTION", "DESCRIBE_QUERY_INVALID_SQL_ERROR", "DESCRIBE_QUERY_TIMEOUT", "DESCRIBE_QUERY_UNEXPECTED_FAILURE", "FUNCTIONS_NOT_AVAILABLE_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_TYPE_EXCEPTION", "FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION", "GENERATED_SQL_QUERY_TOO_LONG_EXCEPTION", "GENERIC_CHAT_COMPLETION_EXCEPTION", "GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION", "GENERIC_SQL_EXEC_API_CALL_EXCEPTION", "ILLEGAL_PARAMETER_DEFINITION_EXCEPTION", "INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION", "INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION", "INVALID_CHAT_COMPLETION_ARGUMENTS_JSON_EXCEPTION", "INVALID_CHAT_COMPLETION_JSON_EXCEPTION", "INVALID_COMPLETION_REQUEST_EXCEPTION", "INVALID_FUNCTION_CALL_EXCEPTION", "INVALID_SQL_MULTIPLE_DATASET_REFERENCES_EXCEPTION", "INVALID_SQL_MULTIPLE_STATEMENTS_EXCEPTION", "INVALID_SQL_UNKNOWN_TABLE_EXCEPTION", "INVALID_TABLE_IDENTIFIER_EXCEPTION", "LOCAL_CONTEXT_EXCEEDED_EXCEPTION", "MESSAGE_CANCELLED_WHILE_EXECUTING_EXCEPTION", "MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION", "MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION", "MISSING_SQL_QUERY_EXCEPTION", "NO_DEPLOYMENTS_AVAILABLE_TO_WORKSPACE", "NO_QUERY_TO_VISUALIZE_EXCEPTION", "NO_TABLES_TO_QUERY_EXCEPTION", "RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION", "RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION", "REPLY_PROCESS_TIMEOUT_EXCEPTION", "RETRYABLE_PROCESSING_EXCEPTION", "SQL_EXECUTION_EXCEPTION", "STOP_PROCESS_DUE_TO_AUTO_REGENERATE", "TABLES_MISSING_EXCEPTION", "TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION", "TOO_MANY_TABLES_EXCEPTION", "UNEXPECTED_REPLY_PROCESS_EXCEPTION", "UNKNOWN_AI_MODEL", "WAREHOUSE_ACCESS_MISSING_EXCEPTION", "WAREHOUSE_NOT_FOUND_EXCEPTION"`, v)
	}
}

// Type always returns MessageErrorType to satisfy [pflag.Value] interface
func (f *MessageErrorType) Type() string {
	return "MessageErrorType"
}

func messageErrorTypeToPb(st *MessageErrorType) (*messageErrorTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := messageErrorTypePb(*st)
	return &pb, nil
}

func messageErrorTypeFromPb(pb *messageErrorTypePb) (*MessageErrorType, error) {
	if pb == nil {
		return nil, nil
	}
	st := MessageErrorType(*pb)
	return &st, nil
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
type messageStatusPb string

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

// Type always returns MessageStatus to satisfy [pflag.Value] interface
func (f *MessageStatus) Type() string {
	return "MessageStatus"
}

func messageStatusToPb(st *MessageStatus) (*messageStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := messageStatusPb(*st)
	return &pb, nil
}

func messageStatusFromPb(pb *messageStatusPb) (*MessageStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := MessageStatus(*pb)
	return &st, nil
}

type MigrateDashboardRequest struct {
	// Display name for the new Lakeview dashboard.
	// Wire name: 'display_name'
	DisplayName string
	// The workspace path of the folder to contain the migrated Lakeview
	// dashboard.
	// Wire name: 'parent_path'
	ParentPath string
	// UUID of the dashboard to be migrated.
	// Wire name: 'source_dashboard_id'
	SourceDashboardId string
	// Flag to indicate if mustache parameter syntax ({{ param }}) should be
	// auto-updated to named syntax (:param) when converting datasets in the
	// dashboard.
	// Wire name: 'update_parameter_syntax'
	UpdateParameterSyntax bool

	ForceSendFields []string `tf:"-"`
}

func (st *MigrateDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &migrateDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := migrateDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MigrateDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := migrateDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PendingStatus struct {
	// The token to poll for result asynchronously Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	// Wire name: 'data_token'
	DataToken string
}

func (st *PendingStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pendingStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pendingStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PendingStatus) MarshalJSON() ([]byte, error) {
	pb, err := pendingStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Poll the results for the a query for a published, embedded dashboard
type PollPublishedQueryStatusRequest struct {

	// Wire name: 'dashboard_name'
	DashboardName string `tf:"-"`

	// Wire name: 'dashboard_revision_id'
	DashboardRevisionId string `tf:"-"`
	// Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	// Wire name: 'tokens'
	Tokens []string `tf:"-"`
}

func (st *PollPublishedQueryStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pollPublishedQueryStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pollPublishedQueryStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PollPublishedQueryStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := pollPublishedQueryStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PollQueryStatusResponse struct {

	// Wire name: 'data'
	Data []PollQueryStatusResponseData
}

func (st *PollQueryStatusResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pollQueryStatusResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pollQueryStatusResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PollQueryStatusResponse) MarshalJSON() ([]byte, error) {
	pb, err := pollQueryStatusResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PollQueryStatusResponseData struct {

	// Wire name: 'status'
	Status QueryResponseStatus
}

func (st *PollQueryStatusResponseData) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pollQueryStatusResponseDataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pollQueryStatusResponseDataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PollQueryStatusResponseData) MarshalJSON() ([]byte, error) {
	pb, err := pollQueryStatusResponseDataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PublishRequest struct {
	// UUID identifying the dashboard to be published.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// Flag to indicate if the publisher's credentials should be embedded in the
	// published dashboard. These embedded credentials will be used to execute
	// the published dashboard's queries.
	// Wire name: 'embed_credentials'
	EmbedCredentials bool
	// The ID of the warehouse that can be used to override the warehouse which
	// was set in the draft.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func (st *PublishRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &publishRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := publishRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PublishRequest) MarshalJSON() ([]byte, error) {
	pb, err := publishRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PublishedDashboard struct {
	// The display name of the published dashboard.
	// Wire name: 'display_name'
	DisplayName string
	// Indicates whether credentials are embedded in the published dashboard.
	// Wire name: 'embed_credentials'
	EmbedCredentials bool
	// The timestamp of when the published dashboard was last revised.
	// Wire name: 'revision_create_time'
	RevisionCreateTime string
	// The warehouse ID used to run the published dashboard.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func (st *PublishedDashboard) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &publishedDashboardPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := publishedDashboardFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PublishedDashboard) MarshalJSON() ([]byte, error) {
	pb, err := publishedDashboardToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueryResponseStatus struct {
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	// Wire name: 'canceled'
	Canceled *Empty
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	// Wire name: 'closed'
	Closed *Empty

	// Wire name: 'pending'
	Pending *PendingStatus
	// The statement id in format(01eef5da-c56e-1f36-bafa-21906587d6ba) The
	// statement_id should be identical to data_token in SuccessStatus and
	// PendingStatus. This field is created for audit logging purpose to record
	// the statement_id of all QueryResponseStatus.
	// Wire name: 'statement_id'
	StatementId string

	// Wire name: 'success'
	Success *SuccessStatus

	ForceSendFields []string `tf:"-"`
}

func (st *QueryResponseStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryResponseStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryResponseStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryResponseStatus) MarshalJSON() ([]byte, error) {
	pb, err := queryResponseStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Result struct {
	// If result is truncated
	// Wire name: 'is_truncated'
	IsTruncated bool
	// Row count of the result
	// Wire name: 'row_count'
	RowCount int64
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	// Wire name: 'statement_id'
	StatementId string

	ForceSendFields []string `tf:"-"`
}

func (st *Result) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resultPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resultFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Result) MarshalJSON() ([]byte, error) {
	pb, err := resultToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Schedule struct {
	// A timestamp indicating when the schedule was created.
	// Wire name: 'create_time'
	CreateTime string
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	// Wire name: 'cron_schedule'
	CronSchedule CronSchedule
	// UUID identifying the dashboard to which the schedule belongs.
	// Wire name: 'dashboard_id'
	DashboardId string
	// The display name for schedule.
	// Wire name: 'display_name'
	DisplayName string
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	// Wire name: 'etag'
	Etag string
	// The status indicates whether this schedule is paused or not.
	// Wire name: 'pause_status'
	PauseStatus SchedulePauseStatus
	// UUID identifying the schedule.
	// Wire name: 'schedule_id'
	ScheduleId string
	// A timestamp indicating when the schedule was last updated.
	// Wire name: 'update_time'
	UpdateTime string
	// The warehouse id to run the dashboard with for the schedule.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func (st *Schedule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &schedulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := scheduleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Schedule) MarshalJSON() ([]byte, error) {
	pb, err := scheduleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SchedulePauseStatus string
type schedulePauseStatusPb string

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

func schedulePauseStatusToPb(st *SchedulePauseStatus) (*schedulePauseStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := schedulePauseStatusPb(*st)
	return &pb, nil
}

func schedulePauseStatusFromPb(pb *schedulePauseStatusPb) (*SchedulePauseStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := SchedulePauseStatus(*pb)
	return &st, nil
}

type Subscriber struct {
	// The destination to receive the subscription email. This parameter is
	// mutually exclusive with `user_subscriber`.
	// Wire name: 'destination_subscriber'
	DestinationSubscriber *SubscriptionSubscriberDestination
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	// Wire name: 'user_subscriber'
	UserSubscriber *SubscriptionSubscriberUser
}

func (st *Subscriber) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &subscriberPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := subscriberFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Subscriber) MarshalJSON() ([]byte, error) {
	pb, err := subscriberToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Subscription struct {
	// A timestamp indicating when the subscription was created.
	// Wire name: 'create_time'
	CreateTime string
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	// Wire name: 'created_by_user_id'
	CreatedByUserId int64
	// UUID identifying the dashboard to which the subscription belongs.
	// Wire name: 'dashboard_id'
	DashboardId string
	// The etag for the subscription. Must be left empty on create, can be
	// optionally provided on delete to ensure that the subscription has not
	// been deleted since the last read.
	// Wire name: 'etag'
	Etag string
	// UUID identifying the schedule to which the subscription belongs.
	// Wire name: 'schedule_id'
	ScheduleId string
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	// Wire name: 'subscriber'
	Subscriber Subscriber
	// UUID identifying the subscription.
	// Wire name: 'subscription_id'
	SubscriptionId string
	// A timestamp indicating when the subscription was last updated.
	// Wire name: 'update_time'
	UpdateTime string

	ForceSendFields []string `tf:"-"`
}

func (st *Subscription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &subscriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := subscriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Subscription) MarshalJSON() ([]byte, error) {
	pb, err := subscriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SubscriptionSubscriberDestination struct {
	// The canonical identifier of the destination to receive email
	// notification.
	// Wire name: 'destination_id'
	DestinationId string
}

func (st *SubscriptionSubscriberDestination) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &subscriptionSubscriberDestinationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := subscriptionSubscriberDestinationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SubscriptionSubscriberDestination) MarshalJSON() ([]byte, error) {
	pb, err := subscriptionSubscriberDestinationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SubscriptionSubscriberUser struct {
	// UserId of the subscriber.
	// Wire name: 'user_id'
	UserId int64
}

func (st *SubscriptionSubscriberUser) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &subscriptionSubscriberUserPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := subscriptionSubscriberUserFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SubscriptionSubscriberUser) MarshalJSON() ([]byte, error) {
	pb, err := subscriptionSubscriberUserToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SuccessStatus struct {
	// The token to poll for result asynchronously Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	// Wire name: 'data_token'
	DataToken string
	// Whether the query result is truncated (either by byte limit or row limit)
	// Wire name: 'truncated'
	Truncated bool

	ForceSendFields []string `tf:"-"`
}

func (st *SuccessStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &successStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := successStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SuccessStatus) MarshalJSON() ([]byte, error) {
	pb, err := successStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TextAttachment struct {
	// AI generated message
	// Wire name: 'content'
	Content string

	// Wire name: 'id'
	Id string

	ForceSendFields []string `tf:"-"`
}

func (st *TextAttachment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &textAttachmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := textAttachmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TextAttachment) MarshalJSON() ([]byte, error) {
	pb, err := textAttachmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Trash dashboard
type TrashDashboardRequest struct {
	// UUID identifying the dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func (st *TrashDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &trashDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := trashDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TrashDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := trashDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TrashDashboardResponse struct {
}

func (st *TrashDashboardResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &trashDashboardResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := trashDashboardResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TrashDashboardResponse) MarshalJSON() ([]byte, error) {
	pb, err := trashDashboardResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Unpublish dashboard
type UnpublishDashboardRequest struct {
	// UUID identifying the published dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func (st *UnpublishDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &unpublishDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := unpublishDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UnpublishDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := unpublishDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UnpublishDashboardResponse struct {
}

func (st *UnpublishDashboardResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &unpublishDashboardResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := unpublishDashboardResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UnpublishDashboardResponse) MarshalJSON() ([]byte, error) {
	pb, err := unpublishDashboardResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Update dashboard
type UpdateDashboardRequest struct {

	// Wire name: 'dashboard'
	Dashboard Dashboard
	// UUID identifying the dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func (st *UpdateDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Update dashboard schedule
type UpdateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`

	// Wire name: 'schedule'
	Schedule Schedule
	// UUID identifying the schedule.
	// Wire name: 'schedule_id'
	ScheduleId string `tf:"-"`
}

func (st *UpdateScheduleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateScheduleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateScheduleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateScheduleRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateScheduleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

func durationFromPb(s *string) (*time.Duration, error) {
	if s == nil {
		return nil, nil
	}
	d, err := time.ParseDuration(*s)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func timestampToPb(t *time.Time) (*string, error) {
	if t == nil {
		return nil, nil
	}
	s := t.Format(time.RFC3339)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func fieldMaskToPb(fm *[]string) (*string, error) {
	if fm == nil {
		return nil, nil
	}
	s := strings.Join(*fm, ",")
	return &s, nil
}

func fieldMaskFromPb(s *string) (*[]string, error) {
	if s == nil {
		return nil, nil
	}
	fm := strings.Split(*s, ",")
	return &fm, nil
}

// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/sql"
)

type AuthorizationDetails struct {
	// Represents downscoped permission rules with specific access rights. This
	// field is specific to `workspace_rule_set` constraint.
	// Wire name: 'grant_rules'
	GrantRules []AuthorizationDetailsGrantRule `json:"grant_rules,omitempty"`
	// The acl path of the tree store resource resource.
	// Wire name: 'resource_legacy_acl_path'
	ResourceLegacyAclPath string `json:"resource_legacy_acl_path,omitempty"`
	// The resource name to which the authorization rule applies. This field is
	// specific to `workspace_rule_set` constraint. Format:
	// `workspaces/{workspace_id}/dashboards/{dashboard_id}`
	// Wire name: 'resource_name'
	ResourceName string `json:"resource_name,omitempty"`
	// The type of authorization downscoping policy. Ex: `workspace_rule_set`
	// defines access rules for a specific workspace resource
	// Wire name: 'type'
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	PermissionSet string `json:"permission_set,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	DashboardName string `json:"-" tf:"-"`

	DashboardRevisionId string `json:"-" tf:"-"`
	// Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	Tokens []string `json:"-" tf:"-"`
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
	Status []CancelQueryExecutionResponseStatus `json:"status,omitempty"`
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
	DataToken string `json:"data_token"`
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	// Wire name: 'pending'
	Pending *Empty `json:"pending,omitempty"`
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	// Wire name: 'success'
	Success *Empty `json:"success,omitempty"`
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
	Dashboard Dashboard `json:"dashboard"`
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
	DashboardId string `json:"-" tf:"-"`

	// Wire name: 'schedule'
	Schedule Schedule `json:"schedule"`
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
	DashboardId string `json:"-" tf:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId string `json:"-" tf:"-"`

	// Wire name: 'subscription'
	Subscription Subscription `json:"subscription"`
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
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// A Java timezone id. The schedule will be resolved with respect to this
	// timezone. See [Java TimeZone] for details.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	// Wire name: 'timezone_id'
	TimezoneId string `json:"timezone_id"`
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
	CreateTime string `json:"create_time,omitempty"`
	// UUID identifying the dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id,omitempty"`
	// The display name of the dashboard.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read. This
	// field is excluded in List Dashboards responses.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// The state of the dashboard resource. Used for tracking trashed status.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash. This field is excluded in List
	// Dashboards responses.
	// Wire name: 'parent_path'
	ParentPath string `json:"parent_path,omitempty"`
	// The workspace path of the dashboard asset, including the file name.
	// Exported dashboards always have the file extension `.lvdash.json`. This
	// field is excluded in List Dashboards responses.
	// Wire name: 'path'
	Path string `json:"path,omitempty"`
	// The contents of the dashboard in serialized string form. This field is
	// excluded in List Dashboards responses. Use the [get dashboard API] to
	// retrieve an example response, which includes the `serialized_dashboard`
	// field. This field provides the structure of the JSON string that
	// represents the dashboard's layout and components.
	//
	// [get dashboard API]: https://docs.databricks.com/api/workspace/lakeview/get
	// Wire name: 'serialized_dashboard'
	SerializedDashboard string `json:"serialized_dashboard,omitempty"`
	// The timestamp of when the dashboard was last updated by the user. This
	// field is excluded in List Dashboards responses.
	// Wire name: 'update_time'
	UpdateTime string `json:"update_time,omitempty"`
	// The warehouse ID used to run the dashboard.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type DeleteConversationResponse struct {
}

func (st *DeleteConversationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteConversationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteConversationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteConversationResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteConversationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete dashboard schedule
type DeleteScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" tf:"-"`
	// The etag for the schedule. Optionally, it can be provided to verify that
	// the schedule has not been modified from its last retrieval.
	Etag string `json:"-" tf:"-"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	DashboardId string `json:"-" tf:"-"`
	// The etag for the subscription. Can be optionally provided to ensure that
	// the subscription has not been modified since the last read.
	Etag string `json:"-" tf:"-"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId string `json:"-" tf:"-"`
	// UUID identifying the subscription.
	SubscriptionId string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	DashboardName string `json:"dashboard_name"`

	// Wire name: 'dashboard_revision_id'
	DashboardRevisionId string `json:"dashboard_revision_id"`
	// A dashboard schedule can override the warehouse used as compute for
	// processing the published dashboard queries
	// Wire name: 'override_warehouse_id'
	OverrideWarehouseId string `json:"override_warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	AttachmentId string `json:"attachment_id,omitempty"`
	// Query Attachment if Genie responds with a SQL query
	// Wire name: 'query'
	Query *GenieQueryAttachment `json:"query,omitempty"`
	// Text Attachment if Genie responds with text
	// Wire name: 'text'
	Text *TextAttachment `json:"text,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	ConversationId string `json:"conversation_id"`
	// Timestamp when the message was created
	// Wire name: 'created_timestamp'
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// Conversation ID. Legacy identifier, use conversation_id instead
	// Wire name: 'id'
	Id string `json:"id"`
	// Timestamp when the message was last updated
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string `json:"space_id"`
	// Conversation title
	// Wire name: 'title'
	Title string `json:"title"`
	// ID of the user who created the conversation
	// Wire name: 'user_id'
	UserId int `json:"user_id"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	Content string `json:"content"`
	// The ID associated with the conversation.
	ConversationId string `json:"-" tf:"-"`
	// The ID associated with the Genie space where the conversation is started.
	SpaceId string `json:"-" tf:"-"`
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

// Delete conversation
type GenieDeleteConversationRequest struct {
	// The ID of the conversation to delete.
	ConversationId string `json:"-" tf:"-"`
	// The ID associated with the Genie space where the conversation is located.
	SpaceId string `json:"-" tf:"-"`
}

func (st *GenieDeleteConversationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieDeleteConversationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieDeleteConversationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieDeleteConversationRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieDeleteConversationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Execute message attachment SQL query
type GenieExecuteMessageAttachmentQueryRequest struct {
	// Attachment ID
	AttachmentId string `json:"-" tf:"-"`
	// Conversation ID
	ConversationId string `json:"-" tf:"-"`
	// Message ID
	MessageId string `json:"-" tf:"-"`
	// Genie space ID
	SpaceId string `json:"-" tf:"-"`
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
	ConversationId string `json:"-" tf:"-"`
	// Message ID
	MessageId string `json:"-" tf:"-"`
	// Genie space ID
	SpaceId string `json:"-" tf:"-"`
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
	AttachmentId string `json:"-" tf:"-"`
	// Conversation ID
	ConversationId string `json:"-" tf:"-"`
	// Message ID
	MessageId string `json:"-" tf:"-"`
	// Genie space ID
	SpaceId string `json:"-" tf:"-"`
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
	DownloadId string `json:"download_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	ConversationId string `json:"-" tf:"-"`
	// The ID associated with the target message from the identified
	// conversation.
	MessageId string `json:"-" tf:"-"`
	// The ID associated with the Genie space where the target conversation is
	// located.
	SpaceId string `json:"-" tf:"-"`
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
	AttachmentId string `json:"-" tf:"-"`
	// Conversation ID
	ConversationId string `json:"-" tf:"-"`
	// Download ID. This ID is provided by the [Generate Download
	// endpoint](:method:genie/generateDownloadFullQueryResult)
	DownloadId string `json:"-" tf:"-"`
	// Message ID
	MessageId string `json:"-" tf:"-"`
	// Genie space ID
	SpaceId string `json:"-" tf:"-"`
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
	StatementResponse *sql.StatementResponse `json:"statement_response,omitempty"`
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
	AttachmentId string `json:"-" tf:"-"`
	// Conversation ID
	ConversationId string `json:"-" tf:"-"`
	// Message ID
	MessageId string `json:"-" tf:"-"`
	// Genie space ID
	SpaceId string `json:"-" tf:"-"`
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
	ConversationId string `json:"-" tf:"-"`
	// Message ID
	MessageId string `json:"-" tf:"-"`
	// Genie space ID
	SpaceId string `json:"-" tf:"-"`
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
	StatementResponse *sql.StatementResponse `json:"statement_response,omitempty"`
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
	AttachmentId string `json:"-" tf:"-"`
	// Conversation ID
	ConversationId string `json:"-" tf:"-"`
	// Message ID
	MessageId string `json:"-" tf:"-"`
	// Genie space ID
	SpaceId string `json:"-" tf:"-"`
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
	SpaceId string `json:"-" tf:"-"`
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

// List Genie spaces
type GenieListSpacesRequest struct {
	// Maximum number of spaces to return per page
	PageSize int `json:"-" tf:"-"`
	// Pagination token for getting the next page of results
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GenieListSpacesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieListSpacesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieListSpacesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieListSpacesRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieListSpacesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieListSpacesResponse struct {
	// Token to get the next page of results
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of Genie spaces
	// Wire name: 'spaces'
	Spaces []GenieSpace `json:"spaces,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GenieListSpacesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieListSpacesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieListSpacesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieListSpacesResponse) MarshalJSON() ([]byte, error) {
	pb, err := genieListSpacesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenieMessage struct {
	// AI-generated response to the message
	// Wire name: 'attachments'
	Attachments []GenieAttachment `json:"attachments,omitempty"`
	// User message content
	// Wire name: 'content'
	Content string `json:"content"`
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string `json:"conversation_id"`
	// Timestamp when the message was created
	// Wire name: 'created_timestamp'
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// Error message if Genie failed to respond to the message
	// Wire name: 'error'
	Error *MessageError `json:"error,omitempty"`
	// Message ID. Legacy identifier, use message_id instead
	// Wire name: 'id'
	Id string `json:"id"`
	// Timestamp when the message was last updated
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Message ID
	// Wire name: 'message_id'
	MessageId string `json:"message_id"`
	// The result of SQL query if the message includes a query attachment.
	// Deprecated. Use `query_result_metadata` in `GenieQueryAttachment`
	// instead.
	// Wire name: 'query_result'
	QueryResult *Result `json:"query_result,omitempty"`
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string `json:"space_id"`
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
	Status MessageStatus `json:"status,omitempty"`
	// ID of the user who created the message
	// Wire name: 'user_id'
	UserId int64 `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	Description string `json:"description,omitempty"`

	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Time when the user updated the query last
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// AI generated SQL query
	// Wire name: 'query'
	Query string `json:"query,omitempty"`
	// Metadata associated with the query result.
	// Wire name: 'query_result_metadata'
	QueryResultMetadata *GenieResultMetadata `json:"query_result_metadata,omitempty"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	// Wire name: 'statement_id'
	StatementId string `json:"statement_id,omitempty"`
	// Name of the query
	// Wire name: 'title'
	Title string `json:"title,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	IsTruncated bool `json:"is_truncated,omitempty"`
	// The number of rows in the result set.
	// Wire name: 'row_count'
	RowCount int64 `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	Description string `json:"description,omitempty"`
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string `json:"space_id"`
	// Title of the Genie Space
	// Wire name: 'title'
	Title string `json:"title"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	Content string `json:"content"`
	// The ID associated with the Genie space where you want to start a
	// conversation.
	SpaceId string `json:"-" tf:"-"`
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
	Conversation *GenieConversation `json:"conversation,omitempty"`
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string `json:"conversation_id"`

	// Wire name: 'message'
	Message *GenieMessage `json:"message,omitempty"`
	// Message ID
	// Wire name: 'message_id'
	MessageId string `json:"message_id"`
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

// Trash Genie Space
type GenieTrashSpaceRequest struct {
	// The ID associated with the Genie space to be trashed.
	SpaceId string `json:"-" tf:"-"`
}

func (st *GenieTrashSpaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genieTrashSpaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genieTrashSpaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenieTrashSpaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := genieTrashSpaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get dashboard
type GetDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string `json:"-" tf:"-"`
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
	DashboardId string `json:"-" tf:"-"`
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
	DashboardId string `json:"-" tf:"-"`
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
	DashboardId string `json:"-" tf:"-"`
	// Provided external value to be included in the custom claim.
	ExternalValue string `json:"-" tf:"-"`
	// Provided external viewer id to be included in the custom claim.
	ExternalViewerId string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	AuthorizationDetails []AuthorizationDetails `json:"authorization_details,omitempty"`
	// Custom claim generated from external_value and external_viewer_id.
	// Format:
	// `urn:aibi:external_data:<external_value>:<external_viewer_id>:<dashboard_id>`
	// Wire name: 'custom_claim'
	CustomClaim string `json:"custom_claim,omitempty"`
	// Scope defining access permissions.
	// Wire name: 'scope'
	Scope string `json:"scope,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	DashboardId string `json:"-" tf:"-"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" tf:"-"`
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
	DashboardId string `json:"-" tf:"-"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId string `json:"-" tf:"-"`
	// UUID identifying the subscription.
	SubscriptionId string `json:"-" tf:"-"`
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

// List dashboards
type ListDashboardsRequest struct {
	// The number of dashboards to return per page.
	PageSize int `json:"-" tf:"-"`
	// A page token, received from a previous `ListDashboards` call. This token
	// can be used to retrieve the subsequent page.
	PageToken string `json:"-" tf:"-"`
	// The flag to include dashboards located in the trash. If unspecified, only
	// active dashboards will be returned.
	ShowTrashed bool `json:"-" tf:"-"`
	// `DASHBOARD_VIEW_BASIC`only includes summary metadata from the dashboard.
	View DashboardView `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	Dashboards []Dashboard `json:"dashboards,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	DashboardId string `json:"-" tf:"-"`
	// The number of schedules to return per page.
	PageSize int `json:"-" tf:"-"`
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'schedules'
	Schedules []Schedule `json:"schedules,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	DashboardId string `json:"-" tf:"-"`
	// The number of subscriptions to return per page.
	PageSize int `json:"-" tf:"-"`
	// A page token, received from a previous `ListSubscriptions` call. Use this
	// to retrieve the subsequent page.
	PageToken string `json:"-" tf:"-"`
	// UUID identifying the schedule which the subscriptions belongs.
	ScheduleId string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'subscriptions'
	Subscriptions []Subscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	Error string `json:"error,omitempty"`

	// Wire name: 'type'
	Type MessageErrorType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// The workspace path of the folder to contain the migrated Lakeview
	// dashboard.
	// Wire name: 'parent_path'
	ParentPath string `json:"parent_path,omitempty"`
	// UUID of the dashboard to be migrated.
	// Wire name: 'source_dashboard_id'
	SourceDashboardId string `json:"source_dashboard_id"`
	// Flag to indicate if mustache parameter syntax ({{ param }}) should be
	// auto-updated to named syntax (:param) when converting datasets in the
	// dashboard.
	// Wire name: 'update_parameter_syntax'
	UpdateParameterSyntax bool `json:"update_parameter_syntax,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	DataToken string `json:"data_token"`
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
	DashboardName string `json:"-" tf:"-"`

	DashboardRevisionId string `json:"-" tf:"-"`
	// Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	Tokens []string `json:"-" tf:"-"`
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
	Data []PollQueryStatusResponseData `json:"data,omitempty"`
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
	Status QueryResponseStatus `json:"status"`
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
	DashboardId string `json:"-" tf:"-"`
	// Flag to indicate if the publisher's credentials should be embedded in the
	// published dashboard. These embedded credentials will be used to execute
	// the published dashboard's queries.
	// Wire name: 'embed_credentials'
	EmbedCredentials bool `json:"embed_credentials,omitempty"`
	// The ID of the warehouse that can be used to override the warehouse which
	// was set in the draft.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	DisplayName string `json:"display_name,omitempty"`
	// Indicates whether credentials are embedded in the published dashboard.
	// Wire name: 'embed_credentials'
	EmbedCredentials bool `json:"embed_credentials,omitempty"`
	// The timestamp of when the published dashboard was last revised.
	// Wire name: 'revision_create_time'
	RevisionCreateTime string `json:"revision_create_time,omitempty"`
	// The warehouse ID used to run the published dashboard.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	Canceled *Empty `json:"canceled,omitempty"`
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	// Wire name: 'closed'
	Closed *Empty `json:"closed,omitempty"`

	// Wire name: 'pending'
	Pending *PendingStatus `json:"pending,omitempty"`
	// The statement id in format(01eef5da-c56e-1f36-bafa-21906587d6ba) The
	// statement_id should be identical to data_token in SuccessStatus and
	// PendingStatus. This field is created for audit logging purpose to record
	// the statement_id of all QueryResponseStatus.
	// Wire name: 'statement_id'
	StatementId string `json:"statement_id,omitempty"`

	// Wire name: 'success'
	Success *SuccessStatus `json:"success,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	IsTruncated bool `json:"is_truncated,omitempty"`
	// Row count of the result
	// Wire name: 'row_count'
	RowCount int64 `json:"row_count,omitempty"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	// Wire name: 'statement_id'
	StatementId string `json:"statement_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	CreateTime string `json:"create_time,omitempty"`
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	// Wire name: 'cron_schedule'
	CronSchedule CronSchedule `json:"cron_schedule"`
	// UUID identifying the dashboard to which the schedule belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id,omitempty"`
	// The display name for schedule.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// The status indicates whether this schedule is paused or not.
	// Wire name: 'pause_status'
	PauseStatus SchedulePauseStatus `json:"pause_status,omitempty"`
	// UUID identifying the schedule.
	// Wire name: 'schedule_id'
	ScheduleId string `json:"schedule_id,omitempty"`
	// A timestamp indicating when the schedule was last updated.
	// Wire name: 'update_time'
	UpdateTime string `json:"update_time,omitempty"`
	// The warehouse id to run the dashboard with for the schedule.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	// Wire name: 'destination_subscriber'
	DestinationSubscriber *SubscriptionSubscriberDestination `json:"destination_subscriber,omitempty"`
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	// Wire name: 'user_subscriber'
	UserSubscriber *SubscriptionSubscriberUser `json:"user_subscriber,omitempty"`
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
	CreateTime string `json:"create_time,omitempty"`
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	// Wire name: 'created_by_user_id'
	CreatedByUserId int64 `json:"created_by_user_id,omitempty"`
	// UUID identifying the dashboard to which the subscription belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id,omitempty"`
	// The etag for the subscription. Must be left empty on create, can be
	// optionally provided on delete to ensure that the subscription has not
	// been deleted since the last read.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// UUID identifying the schedule to which the subscription belongs.
	// Wire name: 'schedule_id'
	ScheduleId string `json:"schedule_id,omitempty"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	// Wire name: 'subscriber'
	Subscriber Subscriber `json:"subscriber"`
	// UUID identifying the subscription.
	// Wire name: 'subscription_id'
	SubscriptionId string `json:"subscription_id,omitempty"`
	// A timestamp indicating when the subscription was last updated.
	// Wire name: 'update_time'
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	DestinationId string `json:"destination_id"`
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
	UserId int64 `json:"user_id"`
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
	DataToken string `json:"data_token"`
	// Whether the query result is truncated (either by byte limit or row limit)
	// Wire name: 'truncated'
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	Content string `json:"content,omitempty"`

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	DashboardId string `json:"-" tf:"-"`
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

type TrashSpaceResponse struct {
}

func (st *TrashSpaceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &trashSpaceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := trashSpaceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TrashSpaceResponse) MarshalJSON() ([]byte, error) {
	pb, err := trashSpaceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Unpublish dashboard
type UnpublishDashboardRequest struct {
	// UUID identifying the published dashboard.
	DashboardId string `json:"-" tf:"-"`
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
	Dashboard Dashboard `json:"dashboard"`
	// UUID identifying the dashboard.
	DashboardId string `json:"-" tf:"-"`
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
	DashboardId string `json:"-" tf:"-"`

	// Wire name: 'schedule'
	Schedule Schedule `json:"schedule"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" tf:"-"`
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

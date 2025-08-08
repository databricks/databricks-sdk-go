// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/dashboards/dashboardspb"
	"github.com/databricks/databricks-sdk-go/service/sql"
)

type AuthorizationDetails struct {
	// Represents downscoped permission rules with specific access rights. This
	// field is specific to `workspace_rule_set` constraint.
	// Wire name: 'grant_rules'
	GrantRules []AuthorizationDetailsGrantRule ``
	// The acl path of the tree store resource resource.
	// Wire name: 'resource_legacy_acl_path'
	ResourceLegacyAclPath string ``
	// The resource name to which the authorization rule applies. This field is
	// specific to `workspace_rule_set` constraint. Format:
	// `workspaces/{workspace_id}/dashboards/{dashboard_id}`
	// Wire name: 'resource_name'
	ResourceName string ``
	// The type of authorization downscoping policy. Ex: `workspace_rule_set`
	// defines access rules for a specific workspace resource
	// Wire name: 'type'
	Type            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AuthorizationDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AuthorizationDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AuthorizationDetailsToPb(st *AuthorizationDetails) (*dashboardspb.AuthorizationDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.AuthorizationDetailsPb{}

	var grantRulesPb []dashboardspb.AuthorizationDetailsGrantRulePb
	for _, item := range st.GrantRules {
		itemPb, err := AuthorizationDetailsGrantRuleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			grantRulesPb = append(grantRulesPb, *itemPb)
		}
	}
	pb.GrantRules = grantRulesPb
	pb.ResourceLegacyAclPath = st.ResourceLegacyAclPath
	pb.ResourceName = st.ResourceName
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AuthorizationDetailsFromPb(pb *dashboardspb.AuthorizationDetailsPb) (*AuthorizationDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AuthorizationDetails{}

	var grantRulesField []AuthorizationDetailsGrantRule
	for _, itemPb := range pb.GrantRules {
		item, err := AuthorizationDetailsGrantRuleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			grantRulesField = append(grantRulesField, *item)
		}
	}
	st.GrantRules = grantRulesField
	st.ResourceLegacyAclPath = pb.ResourceLegacyAclPath
	st.ResourceName = pb.ResourceName
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AuthorizationDetailsGrantRule struct {
	// Permission sets for dashboard are defined in
	// iam-common/rbac-common/permission-sets/definitions/TreeStoreBasePermissionSets
	// Ex: `permissionSets/dashboard.runner`
	// Wire name: 'permission_set'
	PermissionSet   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AuthorizationDetailsGrantRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AuthorizationDetailsGrantRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AuthorizationDetailsGrantRuleToPb(st *AuthorizationDetailsGrantRule) (*dashboardspb.AuthorizationDetailsGrantRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.AuthorizationDetailsGrantRulePb{}
	pb.PermissionSet = st.PermissionSet

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AuthorizationDetailsGrantRuleFromPb(pb *dashboardspb.AuthorizationDetailsGrantRulePb) (*AuthorizationDetailsGrantRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AuthorizationDetailsGrantRule{}
	st.PermissionSet = pb.PermissionSet

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateDashboardRequest struct {

	// Wire name: 'dashboard'
	Dashboard Dashboard ``
}

func CreateDashboardRequestToPb(st *CreateDashboardRequest) (*dashboardspb.CreateDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.CreateDashboardRequestPb{}
	dashboardPb, err := DashboardToPb(&st.Dashboard)
	if err != nil {
		return nil, err
	}
	if dashboardPb != nil {
		pb.Dashboard = *dashboardPb
	}

	return pb, nil
}

func CreateDashboardRequestFromPb(pb *dashboardspb.CreateDashboardRequestPb) (*CreateDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDashboardRequest{}
	dashboardField, err := DashboardFromPb(&pb.Dashboard)
	if err != nil {
		return nil, err
	}
	if dashboardField != nil {
		st.Dashboard = *dashboardField
	}

	return st, nil
}

type CreateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// The schedule to create. A dashboard is limited to 10 schedules.
	// Wire name: 'schedule'
	Schedule Schedule ``
}

func CreateScheduleRequestToPb(st *CreateScheduleRequest) (*dashboardspb.CreateScheduleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.CreateScheduleRequestPb{}
	pb.DashboardId = st.DashboardId
	schedulePb, err := ScheduleToPb(&st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = *schedulePb
	}

	return pb, nil
}

func CreateScheduleRequestFromPb(pb *dashboardspb.CreateScheduleRequestPb) (*CreateScheduleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateScheduleRequest{}
	st.DashboardId = pb.DashboardId
	scheduleField, err := ScheduleFromPb(&pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = *scheduleField
	}

	return st, nil
}

type CreateSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	// Wire name: 'schedule_id'
	ScheduleId string `tf:"-"`
	// The subscription to create. A schedule is limited to 100 subscriptions.
	// Wire name: 'subscription'
	Subscription Subscription ``
}

func CreateSubscriptionRequestToPb(st *CreateSubscriptionRequest) (*dashboardspb.CreateSubscriptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.CreateSubscriptionRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.ScheduleId = st.ScheduleId
	subscriptionPb, err := SubscriptionToPb(&st.Subscription)
	if err != nil {
		return nil, err
	}
	if subscriptionPb != nil {
		pb.Subscription = *subscriptionPb
	}

	return pb, nil
}

func CreateSubscriptionRequestFromPb(pb *dashboardspb.CreateSubscriptionRequestPb) (*CreateSubscriptionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateSubscriptionRequest{}
	st.DashboardId = pb.DashboardId
	st.ScheduleId = pb.ScheduleId
	subscriptionField, err := SubscriptionFromPb(&pb.Subscription)
	if err != nil {
		return nil, err
	}
	if subscriptionField != nil {
		st.Subscription = *subscriptionField
	}

	return st, nil
}

type CronSchedule struct {
	// A cron expression using quartz syntax. EX: `0 0 8 * * ?` represents
	// everyday at 8am. See [Cron Trigger] for details.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	// Wire name: 'quartz_cron_expression'
	QuartzCronExpression string ``
	// A Java timezone id. The schedule will be resolved with respect to this
	// timezone. See [Java TimeZone] for details.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	// Wire name: 'timezone_id'
	TimezoneId string ``
}

func CronScheduleToPb(st *CronSchedule) (*dashboardspb.CronSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.CronSchedulePb{}
	pb.QuartzCronExpression = st.QuartzCronExpression
	pb.TimezoneId = st.TimezoneId

	return pb, nil
}

func CronScheduleFromPb(pb *dashboardspb.CronSchedulePb) (*CronSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CronSchedule{}
	st.QuartzCronExpression = pb.QuartzCronExpression
	st.TimezoneId = pb.TimezoneId

	return st, nil
}

type Dashboard struct {
	// The timestamp of when the dashboard was created.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// UUID identifying the dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string ``
	// The display name of the dashboard.
	// Wire name: 'display_name'
	DisplayName string ``
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read. This
	// field is excluded in List Dashboards responses.
	// Wire name: 'etag'
	Etag string ``
	// The state of the dashboard resource. Used for tracking trashed status.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState ``
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash. This field is excluded in List
	// Dashboards responses.
	// Wire name: 'parent_path'
	ParentPath string ``
	// The workspace path of the dashboard asset, including the file name.
	// Exported dashboards always have the file extension `.lvdash.json`. This
	// field is excluded in List Dashboards responses.
	// Wire name: 'path'
	Path string ``
	// The contents of the dashboard in serialized string form. This field is
	// excluded in List Dashboards responses. Use the [get dashboard API] to
	// retrieve an example response, which includes the `serialized_dashboard`
	// field. This field provides the structure of the JSON string that
	// represents the dashboard's layout and components.
	//
	// [get dashboard API]: https://docs.databricks.com/api/workspace/lakeview/get
	// Wire name: 'serialized_dashboard'
	SerializedDashboard string ``
	// The timestamp of when the dashboard was last updated by the user. This
	// field is excluded in List Dashboards responses.
	// Wire name: 'update_time'
	UpdateTime *time.Time ``
	// The warehouse ID used to run the dashboard.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *Dashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Dashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DashboardToPb(st *Dashboard) (*dashboardspb.DashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.DashboardPb{}
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
	pb.DashboardId = st.DashboardId
	pb.DisplayName = st.DisplayName
	pb.Etag = st.Etag
	lifecycleStatePb, err := LifecycleStateToPb(&st.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStatePb != nil {
		pb.LifecycleState = *lifecycleStatePb
	}
	pb.ParentPath = st.ParentPath
	pb.Path = st.Path
	pb.SerializedDashboard = st.SerializedDashboard
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DashboardFromPb(pb *dashboardspb.DashboardPb) (*Dashboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Dashboard{}
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
	st.DashboardId = pb.DashboardId
	st.DisplayName = pb.DisplayName
	st.Etag = pb.Etag
	lifecycleStateField, err := LifecycleStateFromPb(&pb.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStateField != nil {
		st.LifecycleState = *lifecycleStateField
	}
	st.ParentPath = pb.ParentPath
	st.Path = pb.Path
	st.SerializedDashboard = pb.SerializedDashboard
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func DashboardViewToPb(st *DashboardView) (*dashboardspb.DashboardViewPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dashboardspb.DashboardViewPb(*st)
	return &pb, nil
}

func DashboardViewFromPb(pb *dashboardspb.DashboardViewPb) (*DashboardView, error) {
	if pb == nil {
		return nil, nil
	}
	st := DashboardView(*pb)
	return &st, nil
}

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
	ScheduleId      string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteScheduleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteScheduleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteScheduleRequestToPb(st *DeleteScheduleRequest) (*dashboardspb.DeleteScheduleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.DeleteScheduleRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.Etag = st.Etag
	pb.ScheduleId = st.ScheduleId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteScheduleRequestFromPb(pb *dashboardspb.DeleteScheduleRequestPb) (*DeleteScheduleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteScheduleRequest{}
	st.DashboardId = pb.DashboardId
	st.Etag = pb.Etag
	st.ScheduleId = pb.ScheduleId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

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
	SubscriptionId  string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteSubscriptionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteSubscriptionRequestToPb(st *DeleteSubscriptionRequest) (*dashboardspb.DeleteSubscriptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.DeleteSubscriptionRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.Etag = st.Etag
	pb.ScheduleId = st.ScheduleId
	pb.SubscriptionId = st.SubscriptionId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteSubscriptionRequestFromPb(pb *dashboardspb.DeleteSubscriptionRequestPb) (*DeleteSubscriptionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSubscriptionRequest{}
	st.DashboardId = pb.DashboardId
	st.Etag = pb.Etag
	st.ScheduleId = pb.ScheduleId
	st.SubscriptionId = pb.SubscriptionId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Genie AI Response
type GenieAttachment struct {
	// Attachment ID
	// Wire name: 'attachment_id'
	AttachmentId string ``
	// Query Attachment if Genie responds with a SQL query
	// Wire name: 'query'
	Query *GenieQueryAttachment ``
	// Text Attachment if Genie responds with text
	// Wire name: 'text'
	Text            *TextAttachment ``
	ForceSendFields []string        `tf:"-"`
}

func (s *GenieAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenieAttachmentToPb(st *GenieAttachment) (*dashboardspb.GenieAttachmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieAttachmentPb{}
	pb.AttachmentId = st.AttachmentId
	queryPb, err := GenieQueryAttachmentToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}
	textPb, err := TextAttachmentToPb(st.Text)
	if err != nil {
		return nil, err
	}
	if textPb != nil {
		pb.Text = textPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GenieAttachmentFromPb(pb *dashboardspb.GenieAttachmentPb) (*GenieAttachment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieAttachment{}
	st.AttachmentId = pb.AttachmentId
	queryField, err := GenieQueryAttachmentFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}
	textField, err := TextAttachmentFromPb(pb.Text)
	if err != nil {
		return nil, err
	}
	if textField != nil {
		st.Text = textField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GenieConversation struct {
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string ``
	// Timestamp when the message was created
	// Wire name: 'created_timestamp'
	CreatedTimestamp int64 ``
	// Conversation ID. Legacy identifier, use conversation_id instead
	// Wire name: 'id'
	Id string ``
	// Timestamp when the message was last updated
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string ``
	// Conversation title
	// Wire name: 'title'
	Title string ``
	// ID of the user who created the conversation
	// Wire name: 'user_id'
	UserId          int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *GenieConversation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieConversation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenieConversationToPb(st *GenieConversation) (*dashboardspb.GenieConversationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieConversationPb{}
	pb.ConversationId = st.ConversationId
	pb.CreatedTimestamp = st.CreatedTimestamp
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.SpaceId = st.SpaceId
	pb.Title = st.Title
	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GenieConversationFromPb(pb *dashboardspb.GenieConversationPb) (*GenieConversation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieConversation{}
	st.ConversationId = pb.ConversationId
	st.CreatedTimestamp = pb.CreatedTimestamp
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.SpaceId = pb.SpaceId
	st.Title = pb.Title
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GenieConversationSummary struct {

	// Wire name: 'conversation_id'
	ConversationId string ``

	// Wire name: 'created_timestamp'
	CreatedTimestamp int64 ``

	// Wire name: 'title'
	Title string ``
}

func GenieConversationSummaryToPb(st *GenieConversationSummary) (*dashboardspb.GenieConversationSummaryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieConversationSummaryPb{}
	pb.ConversationId = st.ConversationId
	pb.CreatedTimestamp = st.CreatedTimestamp
	pb.Title = st.Title

	return pb, nil
}

func GenieConversationSummaryFromPb(pb *dashboardspb.GenieConversationSummaryPb) (*GenieConversationSummary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieConversationSummary{}
	st.ConversationId = pb.ConversationId
	st.CreatedTimestamp = pb.CreatedTimestamp
	st.Title = pb.Title

	return st, nil
}

type GenieCreateConversationMessageRequest struct {
	// User message content.
	// Wire name: 'content'
	Content string ``
	// The ID associated with the conversation.
	// Wire name: 'conversation_id'
	ConversationId string `tf:"-"`
	// The ID associated with the Genie space where the conversation is started.
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func GenieCreateConversationMessageRequestToPb(st *GenieCreateConversationMessageRequest) (*dashboardspb.GenieCreateConversationMessageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieCreateConversationMessageRequestPb{}
	pb.Content = st.Content
	pb.ConversationId = st.ConversationId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

func GenieCreateConversationMessageRequestFromPb(pb *dashboardspb.GenieCreateConversationMessageRequestPb) (*GenieCreateConversationMessageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieCreateConversationMessageRequest{}
	st.Content = pb.Content
	st.ConversationId = pb.ConversationId
	st.SpaceId = pb.SpaceId

	return st, nil
}

type GenieDeleteConversationRequest struct {
	// The ID of the conversation to delete.
	// Wire name: 'conversation_id'
	ConversationId string `tf:"-"`
	// The ID associated with the Genie space where the conversation is located.
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func GenieDeleteConversationRequestToPb(st *GenieDeleteConversationRequest) (*dashboardspb.GenieDeleteConversationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieDeleteConversationRequestPb{}
	pb.ConversationId = st.ConversationId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

func GenieDeleteConversationRequestFromPb(pb *dashboardspb.GenieDeleteConversationRequestPb) (*GenieDeleteConversationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieDeleteConversationRequest{}
	st.ConversationId = pb.ConversationId
	st.SpaceId = pb.SpaceId

	return st, nil
}

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

func GenieExecuteMessageAttachmentQueryRequestToPb(st *GenieExecuteMessageAttachmentQueryRequest) (*dashboardspb.GenieExecuteMessageAttachmentQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieExecuteMessageAttachmentQueryRequestPb{}
	pb.AttachmentId = st.AttachmentId
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

func GenieExecuteMessageAttachmentQueryRequestFromPb(pb *dashboardspb.GenieExecuteMessageAttachmentQueryRequestPb) (*GenieExecuteMessageAttachmentQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieExecuteMessageAttachmentQueryRequest{}
	st.AttachmentId = pb.AttachmentId
	st.ConversationId = pb.ConversationId
	st.MessageId = pb.MessageId
	st.SpaceId = pb.SpaceId

	return st, nil
}

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

func GenieExecuteMessageQueryRequestToPb(st *GenieExecuteMessageQueryRequest) (*dashboardspb.GenieExecuteMessageQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieExecuteMessageQueryRequestPb{}
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

func GenieExecuteMessageQueryRequestFromPb(pb *dashboardspb.GenieExecuteMessageQueryRequestPb) (*GenieExecuteMessageQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieExecuteMessageQueryRequest{}
	st.ConversationId = pb.ConversationId
	st.MessageId = pb.MessageId
	st.SpaceId = pb.SpaceId

	return st, nil
}

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

func GenieGetConversationMessageRequestToPb(st *GenieGetConversationMessageRequest) (*dashboardspb.GenieGetConversationMessageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieGetConversationMessageRequestPb{}
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

func GenieGetConversationMessageRequestFromPb(pb *dashboardspb.GenieGetConversationMessageRequestPb) (*GenieGetConversationMessageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetConversationMessageRequest{}
	st.ConversationId = pb.ConversationId
	st.MessageId = pb.MessageId
	st.SpaceId = pb.SpaceId

	return st, nil
}

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

func GenieGetMessageAttachmentQueryResultRequestToPb(st *GenieGetMessageAttachmentQueryResultRequest) (*dashboardspb.GenieGetMessageAttachmentQueryResultRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieGetMessageAttachmentQueryResultRequestPb{}
	pb.AttachmentId = st.AttachmentId
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

func GenieGetMessageAttachmentQueryResultRequestFromPb(pb *dashboardspb.GenieGetMessageAttachmentQueryResultRequestPb) (*GenieGetMessageAttachmentQueryResultRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetMessageAttachmentQueryResultRequest{}
	st.AttachmentId = pb.AttachmentId
	st.ConversationId = pb.ConversationId
	st.MessageId = pb.MessageId
	st.SpaceId = pb.SpaceId

	return st, nil
}

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

func GenieGetMessageQueryResultRequestToPb(st *GenieGetMessageQueryResultRequest) (*dashboardspb.GenieGetMessageQueryResultRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieGetMessageQueryResultRequestPb{}
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

func GenieGetMessageQueryResultRequestFromPb(pb *dashboardspb.GenieGetMessageQueryResultRequestPb) (*GenieGetMessageQueryResultRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetMessageQueryResultRequest{}
	st.ConversationId = pb.ConversationId
	st.MessageId = pb.MessageId
	st.SpaceId = pb.SpaceId

	return st, nil
}

type GenieGetMessageQueryResultResponse struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	// Wire name: 'statement_response'
	StatementResponse *sql.StatementResponse ``
}

func GenieGetMessageQueryResultResponseToPb(st *GenieGetMessageQueryResultResponse) (*dashboardspb.GenieGetMessageQueryResultResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieGetMessageQueryResultResponsePb{}
	statementResponsePb, err := sql.StatementResponseToPb(st.StatementResponse)
	if err != nil {
		return nil, err
	}
	if statementResponsePb != nil {
		pb.StatementResponse = statementResponsePb
	}

	return pb, nil
}

func GenieGetMessageQueryResultResponseFromPb(pb *dashboardspb.GenieGetMessageQueryResultResponsePb) (*GenieGetMessageQueryResultResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetMessageQueryResultResponse{}
	statementResponseField, err := sql.StatementResponseFromPb(pb.StatementResponse)
	if err != nil {
		return nil, err
	}
	if statementResponseField != nil {
		st.StatementResponse = statementResponseField
	}

	return st, nil
}

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

func GenieGetQueryResultByAttachmentRequestToPb(st *GenieGetQueryResultByAttachmentRequest) (*dashboardspb.GenieGetQueryResultByAttachmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieGetQueryResultByAttachmentRequestPb{}
	pb.AttachmentId = st.AttachmentId
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

func GenieGetQueryResultByAttachmentRequestFromPb(pb *dashboardspb.GenieGetQueryResultByAttachmentRequestPb) (*GenieGetQueryResultByAttachmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetQueryResultByAttachmentRequest{}
	st.AttachmentId = pb.AttachmentId
	st.ConversationId = pb.ConversationId
	st.MessageId = pb.MessageId
	st.SpaceId = pb.SpaceId

	return st, nil
}

type GenieGetSpaceRequest struct {
	// The ID associated with the Genie space
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func GenieGetSpaceRequestToPb(st *GenieGetSpaceRequest) (*dashboardspb.GenieGetSpaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieGetSpaceRequestPb{}
	pb.SpaceId = st.SpaceId

	return pb, nil
}

func GenieGetSpaceRequestFromPb(pb *dashboardspb.GenieGetSpaceRequestPb) (*GenieGetSpaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetSpaceRequest{}
	st.SpaceId = pb.SpaceId

	return st, nil
}

type GenieListConversationsRequest struct {
	// Maximum number of conversations to return per page
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Token to get the next page of results
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// The ID of the Genie space to retrieve conversations from.
	// Wire name: 'space_id'
	SpaceId         string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GenieListConversationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListConversationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenieListConversationsRequestToPb(st *GenieListConversationsRequest) (*dashboardspb.GenieListConversationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieListConversationsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.SpaceId = st.SpaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GenieListConversationsRequestFromPb(pb *dashboardspb.GenieListConversationsRequestPb) (*GenieListConversationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieListConversationsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.SpaceId = pb.SpaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GenieListConversationsResponse struct {
	// List of conversations in the Genie space
	// Wire name: 'conversations'
	Conversations []GenieConversationSummary ``
	// Token to get the next page of results
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GenieListConversationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListConversationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenieListConversationsResponseToPb(st *GenieListConversationsResponse) (*dashboardspb.GenieListConversationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieListConversationsResponsePb{}

	var conversationsPb []dashboardspb.GenieConversationSummaryPb
	for _, item := range st.Conversations {
		itemPb, err := GenieConversationSummaryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			conversationsPb = append(conversationsPb, *itemPb)
		}
	}
	pb.Conversations = conversationsPb
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GenieListConversationsResponseFromPb(pb *dashboardspb.GenieListConversationsResponsePb) (*GenieListConversationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieListConversationsResponse{}

	var conversationsField []GenieConversationSummary
	for _, itemPb := range pb.Conversations {
		item, err := GenieConversationSummaryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			conversationsField = append(conversationsField, *item)
		}
	}
	st.Conversations = conversationsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GenieListSpacesRequest struct {
	// Maximum number of spaces to return per page
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Pagination token for getting the next page of results
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GenieListSpacesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListSpacesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenieListSpacesRequestToPb(st *GenieListSpacesRequest) (*dashboardspb.GenieListSpacesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieListSpacesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GenieListSpacesRequestFromPb(pb *dashboardspb.GenieListSpacesRequestPb) (*GenieListSpacesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieListSpacesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GenieListSpacesResponse struct {
	// Token to get the next page of results
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// List of Genie spaces
	// Wire name: 'spaces'
	Spaces          []GenieSpace ``
	ForceSendFields []string     `tf:"-"`
}

func (s *GenieListSpacesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListSpacesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenieListSpacesResponseToPb(st *GenieListSpacesResponse) (*dashboardspb.GenieListSpacesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieListSpacesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var spacesPb []dashboardspb.GenieSpacePb
	for _, item := range st.Spaces {
		itemPb, err := GenieSpaceToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			spacesPb = append(spacesPb, *itemPb)
		}
	}
	pb.Spaces = spacesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GenieListSpacesResponseFromPb(pb *dashboardspb.GenieListSpacesResponsePb) (*GenieListSpacesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieListSpacesResponse{}
	st.NextPageToken = pb.NextPageToken

	var spacesField []GenieSpace
	for _, itemPb := range pb.Spaces {
		item, err := GenieSpaceFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			spacesField = append(spacesField, *item)
		}
	}
	st.Spaces = spacesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GenieMessage struct {
	// AI-generated response to the message
	// Wire name: 'attachments'
	Attachments []GenieAttachment ``
	// User message content
	// Wire name: 'content'
	Content string ``
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string ``
	// Timestamp when the message was created
	// Wire name: 'created_timestamp'
	CreatedTimestamp int64 ``
	// Error message if Genie failed to respond to the message
	// Wire name: 'error'
	Error *MessageError ``
	// Message ID. Legacy identifier, use message_id instead
	// Wire name: 'id'
	Id string ``
	// Timestamp when the message was last updated
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// Message ID
	// Wire name: 'message_id'
	MessageId string ``
	// The result of SQL query if the message includes a query attachment.
	// Deprecated. Use `query_result_metadata` in `GenieQueryAttachment`
	// instead.
	// Wire name: 'query_result'
	QueryResult *Result ``
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string ``

	// Wire name: 'status'
	Status MessageStatus ``
	// ID of the user who created the message
	// Wire name: 'user_id'
	UserId          int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *GenieMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenieMessageToPb(st *GenieMessage) (*dashboardspb.GenieMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieMessagePb{}

	var attachmentsPb []dashboardspb.GenieAttachmentPb
	for _, item := range st.Attachments {
		itemPb, err := GenieAttachmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			attachmentsPb = append(attachmentsPb, *itemPb)
		}
	}
	pb.Attachments = attachmentsPb
	pb.Content = st.Content
	pb.ConversationId = st.ConversationId
	pb.CreatedTimestamp = st.CreatedTimestamp
	errorPb, err := MessageErrorToPb(st.Error)
	if err != nil {
		return nil, err
	}
	if errorPb != nil {
		pb.Error = errorPb
	}
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.MessageId = st.MessageId
	queryResultPb, err := ResultToPb(st.QueryResult)
	if err != nil {
		return nil, err
	}
	if queryResultPb != nil {
		pb.QueryResult = queryResultPb
	}
	pb.SpaceId = st.SpaceId
	statusPb, err := MessageStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GenieMessageFromPb(pb *dashboardspb.GenieMessagePb) (*GenieMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieMessage{}

	var attachmentsField []GenieAttachment
	for _, itemPb := range pb.Attachments {
		item, err := GenieAttachmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			attachmentsField = append(attachmentsField, *item)
		}
	}
	st.Attachments = attachmentsField
	st.Content = pb.Content
	st.ConversationId = pb.ConversationId
	st.CreatedTimestamp = pb.CreatedTimestamp
	errorField, err := MessageErrorFromPb(pb.Error)
	if err != nil {
		return nil, err
	}
	if errorField != nil {
		st.Error = errorField
	}
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.MessageId = pb.MessageId
	queryResultField, err := ResultFromPb(pb.QueryResult)
	if err != nil {
		return nil, err
	}
	if queryResultField != nil {
		st.QueryResult = queryResultField
	}
	st.SpaceId = pb.SpaceId
	statusField, err := MessageStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GenieQueryAttachment struct {
	// Description of the query
	// Wire name: 'description'
	Description string ``

	// Wire name: 'id'
	Id string ``
	// Time when the user updated the query last
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// AI generated SQL query
	// Wire name: 'query'
	Query string ``
	// Metadata associated with the query result.
	// Wire name: 'query_result_metadata'
	QueryResultMetadata *GenieResultMetadata ``
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	// Wire name: 'statement_id'
	StatementId string ``
	// Name of the query
	// Wire name: 'title'
	Title           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GenieQueryAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieQueryAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenieQueryAttachmentToPb(st *GenieQueryAttachment) (*dashboardspb.GenieQueryAttachmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieQueryAttachmentPb{}
	pb.Description = st.Description
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.Query = st.Query
	queryResultMetadataPb, err := GenieResultMetadataToPb(st.QueryResultMetadata)
	if err != nil {
		return nil, err
	}
	if queryResultMetadataPb != nil {
		pb.QueryResultMetadata = queryResultMetadataPb
	}
	pb.StatementId = st.StatementId
	pb.Title = st.Title

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GenieQueryAttachmentFromPb(pb *dashboardspb.GenieQueryAttachmentPb) (*GenieQueryAttachment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieQueryAttachment{}
	st.Description = pb.Description
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.Query = pb.Query
	queryResultMetadataField, err := GenieResultMetadataFromPb(pb.QueryResultMetadata)
	if err != nil {
		return nil, err
	}
	if queryResultMetadataField != nil {
		st.QueryResultMetadata = queryResultMetadataField
	}
	st.StatementId = pb.StatementId
	st.Title = pb.Title

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GenieResultMetadata struct {
	// Indicates whether the result set is truncated.
	// Wire name: 'is_truncated'
	IsTruncated bool ``
	// The number of rows in the result set.
	// Wire name: 'row_count'
	RowCount        int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *GenieResultMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieResultMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenieResultMetadataToPb(st *GenieResultMetadata) (*dashboardspb.GenieResultMetadataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieResultMetadataPb{}
	pb.IsTruncated = st.IsTruncated
	pb.RowCount = st.RowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GenieResultMetadataFromPb(pb *dashboardspb.GenieResultMetadataPb) (*GenieResultMetadata, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieResultMetadata{}
	st.IsTruncated = pb.IsTruncated
	st.RowCount = pb.RowCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GenieSpace struct {
	// Description of the Genie Space
	// Wire name: 'description'
	Description string ``
	// Genie space ID
	// Wire name: 'space_id'
	SpaceId string ``
	// Title of the Genie Space
	// Wire name: 'title'
	Title           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GenieSpace) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieSpace) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GenieSpaceToPb(st *GenieSpace) (*dashboardspb.GenieSpacePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieSpacePb{}
	pb.Description = st.Description
	pb.SpaceId = st.SpaceId
	pb.Title = st.Title

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GenieSpaceFromPb(pb *dashboardspb.GenieSpacePb) (*GenieSpace, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieSpace{}
	st.Description = pb.Description
	st.SpaceId = pb.SpaceId
	st.Title = pb.Title

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GenieStartConversationMessageRequest struct {
	// The text of the message that starts the conversation.
	// Wire name: 'content'
	Content string ``
	// The ID associated with the Genie space where you want to start a
	// conversation.
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func GenieStartConversationMessageRequestToPb(st *GenieStartConversationMessageRequest) (*dashboardspb.GenieStartConversationMessageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieStartConversationMessageRequestPb{}
	pb.Content = st.Content
	pb.SpaceId = st.SpaceId

	return pb, nil
}

func GenieStartConversationMessageRequestFromPb(pb *dashboardspb.GenieStartConversationMessageRequestPb) (*GenieStartConversationMessageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieStartConversationMessageRequest{}
	st.Content = pb.Content
	st.SpaceId = pb.SpaceId

	return st, nil
}

type GenieStartConversationResponse struct {

	// Wire name: 'conversation'
	Conversation *GenieConversation ``
	// Conversation ID
	// Wire name: 'conversation_id'
	ConversationId string ``

	// Wire name: 'message'
	Message *GenieMessage ``
	// Message ID
	// Wire name: 'message_id'
	MessageId string ``
}

func GenieStartConversationResponseToPb(st *GenieStartConversationResponse) (*dashboardspb.GenieStartConversationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieStartConversationResponsePb{}
	conversationPb, err := GenieConversationToPb(st.Conversation)
	if err != nil {
		return nil, err
	}
	if conversationPb != nil {
		pb.Conversation = conversationPb
	}
	pb.ConversationId = st.ConversationId
	messagePb, err := GenieMessageToPb(st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = messagePb
	}
	pb.MessageId = st.MessageId

	return pb, nil
}

func GenieStartConversationResponseFromPb(pb *dashboardspb.GenieStartConversationResponsePb) (*GenieStartConversationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieStartConversationResponse{}
	conversationField, err := GenieConversationFromPb(pb.Conversation)
	if err != nil {
		return nil, err
	}
	if conversationField != nil {
		st.Conversation = conversationField
	}
	st.ConversationId = pb.ConversationId
	messageField, err := GenieMessageFromPb(pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = messageField
	}
	st.MessageId = pb.MessageId

	return st, nil
}

type GenieTrashSpaceRequest struct {
	// The ID associated with the Genie space to be sent to the trash.
	// Wire name: 'space_id'
	SpaceId string `tf:"-"`
}

func GenieTrashSpaceRequestToPb(st *GenieTrashSpaceRequest) (*dashboardspb.GenieTrashSpaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GenieTrashSpaceRequestPb{}
	pb.SpaceId = st.SpaceId

	return pb, nil
}

func GenieTrashSpaceRequestFromPb(pb *dashboardspb.GenieTrashSpaceRequestPb) (*GenieTrashSpaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieTrashSpaceRequest{}
	st.SpaceId = pb.SpaceId

	return st, nil
}

type GetDashboardRequest struct {
	// UUID identifying the dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func GetDashboardRequestToPb(st *GetDashboardRequest) (*dashboardspb.GetDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GetDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

func GetDashboardRequestFromPb(pb *dashboardspb.GetDashboardRequestPb) (*GetDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

type GetPublishedDashboardRequest struct {
	// UUID identifying the published dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func GetPublishedDashboardRequestToPb(st *GetPublishedDashboardRequest) (*dashboardspb.GetPublishedDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GetPublishedDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

func GetPublishedDashboardRequestFromPb(pb *dashboardspb.GetPublishedDashboardRequestPb) (*GetPublishedDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

type GetPublishedDashboardTokenInfoRequest struct {
	// UUID identifying the published dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// Provided external value to be included in the custom claim.
	// Wire name: 'external_value'
	ExternalValue string `tf:"-"`
	// Provided external viewer id to be included in the custom claim.
	// Wire name: 'external_viewer_id'
	ExternalViewerId string   `tf:"-"`
	ForceSendFields  []string `tf:"-"`
}

func (s *GetPublishedDashboardTokenInfoRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedDashboardTokenInfoRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetPublishedDashboardTokenInfoRequestToPb(st *GetPublishedDashboardTokenInfoRequest) (*dashboardspb.GetPublishedDashboardTokenInfoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GetPublishedDashboardTokenInfoRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.ExternalValue = st.ExternalValue
	pb.ExternalViewerId = st.ExternalViewerId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetPublishedDashboardTokenInfoRequestFromPb(pb *dashboardspb.GetPublishedDashboardTokenInfoRequestPb) (*GetPublishedDashboardTokenInfoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedDashboardTokenInfoRequest{}
	st.DashboardId = pb.DashboardId
	st.ExternalValue = pb.ExternalValue
	st.ExternalViewerId = pb.ExternalViewerId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetPublishedDashboardTokenInfoResponse struct {
	// Authorization constraints for accessing the published dashboard.
	// Currently includes `workspace_rule_set` and could be enriched with
	// `unity_catalog_privileges` before oAuth token generation.
	// Wire name: 'authorization_details'
	AuthorizationDetails []AuthorizationDetails ``
	// Custom claim generated from external_value and external_viewer_id.
	// Format:
	// `urn:aibi:external_data:<external_value>:<external_viewer_id>:<dashboard_id>`
	// Wire name: 'custom_claim'
	CustomClaim string ``
	// Scope defining access permissions.
	// Wire name: 'scope'
	Scope           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GetPublishedDashboardTokenInfoResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedDashboardTokenInfoResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetPublishedDashboardTokenInfoResponseToPb(st *GetPublishedDashboardTokenInfoResponse) (*dashboardspb.GetPublishedDashboardTokenInfoResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GetPublishedDashboardTokenInfoResponsePb{}

	var authorizationDetailsPb []dashboardspb.AuthorizationDetailsPb
	for _, item := range st.AuthorizationDetails {
		itemPb, err := AuthorizationDetailsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			authorizationDetailsPb = append(authorizationDetailsPb, *itemPb)
		}
	}
	pb.AuthorizationDetails = authorizationDetailsPb
	pb.CustomClaim = st.CustomClaim
	pb.Scope = st.Scope

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetPublishedDashboardTokenInfoResponseFromPb(pb *dashboardspb.GetPublishedDashboardTokenInfoResponsePb) (*GetPublishedDashboardTokenInfoResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedDashboardTokenInfoResponse{}

	var authorizationDetailsField []AuthorizationDetails
	for _, itemPb := range pb.AuthorizationDetails {
		item, err := AuthorizationDetailsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			authorizationDetailsField = append(authorizationDetailsField, *item)
		}
	}
	st.AuthorizationDetails = authorizationDetailsField
	st.CustomClaim = pb.CustomClaim
	st.Scope = pb.Scope

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// UUID identifying the schedule.
	// Wire name: 'schedule_id'
	ScheduleId string `tf:"-"`
}

func GetScheduleRequestToPb(st *GetScheduleRequest) (*dashboardspb.GetScheduleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GetScheduleRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.ScheduleId = st.ScheduleId

	return pb, nil
}

func GetScheduleRequestFromPb(pb *dashboardspb.GetScheduleRequestPb) (*GetScheduleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetScheduleRequest{}
	st.DashboardId = pb.DashboardId
	st.ScheduleId = pb.ScheduleId

	return st, nil
}

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

func GetSubscriptionRequestToPb(st *GetSubscriptionRequest) (*dashboardspb.GetSubscriptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.GetSubscriptionRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.ScheduleId = st.ScheduleId
	pb.SubscriptionId = st.SubscriptionId

	return pb, nil
}

func GetSubscriptionRequestFromPb(pb *dashboardspb.GetSubscriptionRequestPb) (*GetSubscriptionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSubscriptionRequest{}
	st.DashboardId = pb.DashboardId
	st.ScheduleId = pb.ScheduleId
	st.SubscriptionId = pb.SubscriptionId

	return st, nil
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

func LifecycleStateToPb(st *LifecycleState) (*dashboardspb.LifecycleStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dashboardspb.LifecycleStatePb(*st)
	return &pb, nil
}

func LifecycleStateFromPb(pb *dashboardspb.LifecycleStatePb) (*LifecycleState, error) {
	if pb == nil {
		return nil, nil
	}
	st := LifecycleState(*pb)
	return &st, nil
}

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
	View            DashboardView `tf:"-"`
	ForceSendFields []string      `tf:"-"`
}

func (s *ListDashboardsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDashboardsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListDashboardsRequestToPb(st *ListDashboardsRequest) (*dashboardspb.ListDashboardsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.ListDashboardsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.ShowTrashed = st.ShowTrashed
	viewPb, err := DashboardViewToPb(&st.View)
	if err != nil {
		return nil, err
	}
	if viewPb != nil {
		pb.View = *viewPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListDashboardsRequestFromPb(pb *dashboardspb.ListDashboardsRequestPb) (*ListDashboardsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDashboardsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.ShowTrashed = pb.ShowTrashed
	viewField, err := DashboardViewFromPb(&pb.View)
	if err != nil {
		return nil, err
	}
	if viewField != nil {
		st.View = *viewField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListDashboardsResponse struct {

	// Wire name: 'dashboards'
	Dashboards []Dashboard ``
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListDashboardsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDashboardsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListDashboardsResponseToPb(st *ListDashboardsResponse) (*dashboardspb.ListDashboardsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.ListDashboardsResponsePb{}

	var dashboardsPb []dashboardspb.DashboardPb
	for _, item := range st.Dashboards {
		itemPb, err := DashboardToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dashboardsPb = append(dashboardsPb, *itemPb)
		}
	}
	pb.Dashboards = dashboardsPb
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListDashboardsResponseFromPb(pb *dashboardspb.ListDashboardsResponsePb) (*ListDashboardsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDashboardsResponse{}

	var dashboardsField []Dashboard
	for _, itemPb := range pb.Dashboards {
		item, err := DashboardFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dashboardsField = append(dashboardsField, *item)
		}
	}
	st.Dashboards = dashboardsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

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
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListSchedulesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchedulesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListSchedulesRequestToPb(st *ListSchedulesRequest) (*dashboardspb.ListSchedulesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.ListSchedulesRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListSchedulesRequestFromPb(pb *dashboardspb.ListSchedulesRequestPb) (*ListSchedulesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSchedulesRequest{}
	st.DashboardId = pb.DashboardId
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListSchedulesResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent schedules.
	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'schedules'
	Schedules       []Schedule ``
	ForceSendFields []string   `tf:"-"`
}

func (s *ListSchedulesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchedulesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListSchedulesResponseToPb(st *ListSchedulesResponse) (*dashboardspb.ListSchedulesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.ListSchedulesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var schedulesPb []dashboardspb.SchedulePb
	for _, item := range st.Schedules {
		itemPb, err := ScheduleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			schedulesPb = append(schedulesPb, *itemPb)
		}
	}
	pb.Schedules = schedulesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListSchedulesResponseFromPb(pb *dashboardspb.ListSchedulesResponsePb) (*ListSchedulesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSchedulesResponse{}
	st.NextPageToken = pb.NextPageToken

	var schedulesField []Schedule
	for _, itemPb := range pb.Schedules {
		item, err := ScheduleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			schedulesField = append(schedulesField, *item)
		}
	}
	st.Schedules = schedulesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

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
	ScheduleId      string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListSubscriptionsRequestToPb(st *ListSubscriptionsRequest) (*dashboardspb.ListSubscriptionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.ListSubscriptionsRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.ScheduleId = st.ScheduleId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListSubscriptionsRequestFromPb(pb *dashboardspb.ListSubscriptionsRequestPb) (*ListSubscriptionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSubscriptionsRequest{}
	st.DashboardId = pb.DashboardId
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.ScheduleId = pb.ScheduleId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListSubscriptionsResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent subscriptions.
	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'subscriptions'
	Subscriptions   []Subscription ``
	ForceSendFields []string       `tf:"-"`
}

func (s *ListSubscriptionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSubscriptionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListSubscriptionsResponseToPb(st *ListSubscriptionsResponse) (*dashboardspb.ListSubscriptionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.ListSubscriptionsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var subscriptionsPb []dashboardspb.SubscriptionPb
	for _, item := range st.Subscriptions {
		itemPb, err := SubscriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			subscriptionsPb = append(subscriptionsPb, *itemPb)
		}
	}
	pb.Subscriptions = subscriptionsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListSubscriptionsResponseFromPb(pb *dashboardspb.ListSubscriptionsResponsePb) (*ListSubscriptionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSubscriptionsResponse{}
	st.NextPageToken = pb.NextPageToken

	var subscriptionsField []Subscription
	for _, itemPb := range pb.Subscriptions {
		item, err := SubscriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			subscriptionsField = append(subscriptionsField, *item)
		}
	}
	st.Subscriptions = subscriptionsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type MessageError struct {

	// Wire name: 'error'
	Error string ``

	// Wire name: 'type'
	Type            MessageErrorType ``
	ForceSendFields []string         `tf:"-"`
}

func (s *MessageError) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MessageError) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func MessageErrorToPb(st *MessageError) (*dashboardspb.MessageErrorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.MessageErrorPb{}
	pb.Error = st.Error
	typePb, err := MessageErrorTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func MessageErrorFromPb(pb *dashboardspb.MessageErrorPb) (*MessageError, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MessageError{}
	st.Error = pb.Error
	typeField, err := MessageErrorTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func MessageErrorTypeToPb(st *MessageErrorType) (*dashboardspb.MessageErrorTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dashboardspb.MessageErrorTypePb(*st)
	return &pb, nil
}

func MessageErrorTypeFromPb(pb *dashboardspb.MessageErrorTypePb) (*MessageErrorType, error) {
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

func MessageStatusToPb(st *MessageStatus) (*dashboardspb.MessageStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dashboardspb.MessageStatusPb(*st)
	return &pb, nil
}

func MessageStatusFromPb(pb *dashboardspb.MessageStatusPb) (*MessageStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := MessageStatus(*pb)
	return &st, nil
}

type MigrateDashboardRequest struct {
	// Display name for the new Lakeview dashboard.
	// Wire name: 'display_name'
	DisplayName string ``
	// The workspace path of the folder to contain the migrated Lakeview
	// dashboard.
	// Wire name: 'parent_path'
	ParentPath string ``
	// UUID of the dashboard to be migrated.
	// Wire name: 'source_dashboard_id'
	SourceDashboardId string ``
	// Flag to indicate if mustache parameter syntax ({{ param }}) should be
	// auto-updated to named syntax (:param) when converting datasets in the
	// dashboard.
	// Wire name: 'update_parameter_syntax'
	UpdateParameterSyntax bool     ``
	ForceSendFields       []string `tf:"-"`
}

func (s *MigrateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MigrateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func MigrateDashboardRequestToPb(st *MigrateDashboardRequest) (*dashboardspb.MigrateDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.MigrateDashboardRequestPb{}
	pb.DisplayName = st.DisplayName
	pb.ParentPath = st.ParentPath
	pb.SourceDashboardId = st.SourceDashboardId
	pb.UpdateParameterSyntax = st.UpdateParameterSyntax

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func MigrateDashboardRequestFromPb(pb *dashboardspb.MigrateDashboardRequestPb) (*MigrateDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MigrateDashboardRequest{}
	st.DisplayName = pb.DisplayName
	st.ParentPath = pb.ParentPath
	st.SourceDashboardId = pb.SourceDashboardId
	st.UpdateParameterSyntax = pb.UpdateParameterSyntax

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type PublishRequest struct {
	// UUID identifying the dashboard to be published.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// Flag to indicate if the publisher's credentials should be embedded in the
	// published dashboard. These embedded credentials will be used to execute
	// the published dashboard's queries.
	// Wire name: 'embed_credentials'
	EmbedCredentials bool ``
	// The ID of the warehouse that can be used to override the warehouse which
	// was set in the draft.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *PublishRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PublishRequestToPb(st *PublishRequest) (*dashboardspb.PublishRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.PublishRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.EmbedCredentials = st.EmbedCredentials
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PublishRequestFromPb(pb *dashboardspb.PublishRequestPb) (*PublishRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishRequest{}
	st.DashboardId = pb.DashboardId
	st.EmbedCredentials = pb.EmbedCredentials
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type PublishedDashboard struct {
	// The display name of the published dashboard.
	// Wire name: 'display_name'
	DisplayName string ``
	// Indicates whether credentials are embedded in the published dashboard.
	// Wire name: 'embed_credentials'
	EmbedCredentials bool ``
	// The timestamp of when the published dashboard was last revised.
	// Wire name: 'revision_create_time'
	RevisionCreateTime *time.Time ``
	// The warehouse ID used to run the published dashboard.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *PublishedDashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishedDashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PublishedDashboardToPb(st *PublishedDashboard) (*dashboardspb.PublishedDashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.PublishedDashboardPb{}
	pb.DisplayName = st.DisplayName
	pb.EmbedCredentials = st.EmbedCredentials
	revisionCreateTimePb, err := timestampToPb(st.RevisionCreateTime)
	if err != nil {
		return nil, err
	}
	if revisionCreateTimePb != nil {
		pb.RevisionCreateTime = *revisionCreateTimePb
	}
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PublishedDashboardFromPb(pb *dashboardspb.PublishedDashboardPb) (*PublishedDashboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishedDashboard{}
	st.DisplayName = pb.DisplayName
	st.EmbedCredentials = pb.EmbedCredentials
	revisionCreateTimeField, err := timestampFromPb(&pb.RevisionCreateTime)
	if err != nil {
		return nil, err
	}
	if revisionCreateTimeField != nil {
		st.RevisionCreateTime = revisionCreateTimeField
	}
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type Result struct {
	// If result is truncated
	// Wire name: 'is_truncated'
	IsTruncated bool ``
	// Row count of the result
	// Wire name: 'row_count'
	RowCount int64 ``
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	// Wire name: 'statement_id'
	StatementId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *Result) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Result) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ResultToPb(st *Result) (*dashboardspb.ResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.ResultPb{}
	pb.IsTruncated = st.IsTruncated
	pb.RowCount = st.RowCount
	pb.StatementId = st.StatementId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ResultFromPb(pb *dashboardspb.ResultPb) (*Result, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Result{}
	st.IsTruncated = pb.IsTruncated
	st.RowCount = pb.RowCount
	st.StatementId = pb.StatementId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type Schedule struct {
	// A timestamp indicating when the schedule was created.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	// Wire name: 'cron_schedule'
	CronSchedule CronSchedule ``
	// UUID identifying the dashboard to which the schedule belongs.
	// Wire name: 'dashboard_id'
	DashboardId string ``
	// The display name for schedule.
	// Wire name: 'display_name'
	DisplayName string ``
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	// Wire name: 'etag'
	Etag string ``
	// The status indicates whether this schedule is paused or not.
	// Wire name: 'pause_status'
	PauseStatus SchedulePauseStatus ``
	// UUID identifying the schedule.
	// Wire name: 'schedule_id'
	ScheduleId string ``
	// A timestamp indicating when the schedule was last updated.
	// Wire name: 'update_time'
	UpdateTime *time.Time ``
	// The warehouse id to run the dashboard with for the schedule.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *Schedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Schedule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ScheduleToPb(st *Schedule) (*dashboardspb.SchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.SchedulePb{}
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
	cronSchedulePb, err := CronScheduleToPb(&st.CronSchedule)
	if err != nil {
		return nil, err
	}
	if cronSchedulePb != nil {
		pb.CronSchedule = *cronSchedulePb
	}
	pb.DashboardId = st.DashboardId
	pb.DisplayName = st.DisplayName
	pb.Etag = st.Etag
	pauseStatusPb, err := SchedulePauseStatusToPb(&st.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusPb != nil {
		pb.PauseStatus = *pauseStatusPb
	}
	pb.ScheduleId = st.ScheduleId
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ScheduleFromPb(pb *dashboardspb.SchedulePb) (*Schedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Schedule{}
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
	cronScheduleField, err := CronScheduleFromPb(&pb.CronSchedule)
	if err != nil {
		return nil, err
	}
	if cronScheduleField != nil {
		st.CronSchedule = *cronScheduleField
	}
	st.DashboardId = pb.DashboardId
	st.DisplayName = pb.DisplayName
	st.Etag = pb.Etag
	pauseStatusField, err := SchedulePauseStatusFromPb(&pb.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusField != nil {
		st.PauseStatus = *pauseStatusField
	}
	st.ScheduleId = pb.ScheduleId
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func SchedulePauseStatusToPb(st *SchedulePauseStatus) (*dashboardspb.SchedulePauseStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dashboardspb.SchedulePauseStatusPb(*st)
	return &pb, nil
}

func SchedulePauseStatusFromPb(pb *dashboardspb.SchedulePauseStatusPb) (*SchedulePauseStatus, error) {
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
	DestinationSubscriber *SubscriptionSubscriberDestination ``
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	// Wire name: 'user_subscriber'
	UserSubscriber *SubscriptionSubscriberUser ``
}

func SubscriberToPb(st *Subscriber) (*dashboardspb.SubscriberPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.SubscriberPb{}
	destinationSubscriberPb, err := SubscriptionSubscriberDestinationToPb(st.DestinationSubscriber)
	if err != nil {
		return nil, err
	}
	if destinationSubscriberPb != nil {
		pb.DestinationSubscriber = destinationSubscriberPb
	}
	userSubscriberPb, err := SubscriptionSubscriberUserToPb(st.UserSubscriber)
	if err != nil {
		return nil, err
	}
	if userSubscriberPb != nil {
		pb.UserSubscriber = userSubscriberPb
	}

	return pb, nil
}

func SubscriberFromPb(pb *dashboardspb.SubscriberPb) (*Subscriber, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Subscriber{}
	destinationSubscriberField, err := SubscriptionSubscriberDestinationFromPb(pb.DestinationSubscriber)
	if err != nil {
		return nil, err
	}
	if destinationSubscriberField != nil {
		st.DestinationSubscriber = destinationSubscriberField
	}
	userSubscriberField, err := SubscriptionSubscriberUserFromPb(pb.UserSubscriber)
	if err != nil {
		return nil, err
	}
	if userSubscriberField != nil {
		st.UserSubscriber = userSubscriberField
	}

	return st, nil
}

type Subscription struct {
	// A timestamp indicating when the subscription was created.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	// Wire name: 'created_by_user_id'
	CreatedByUserId int64 ``
	// UUID identifying the dashboard to which the subscription belongs.
	// Wire name: 'dashboard_id'
	DashboardId string ``
	// The etag for the subscription. Must be left empty on create, can be
	// optionally provided on delete to ensure that the subscription has not
	// been deleted since the last read.
	// Wire name: 'etag'
	Etag string ``
	// UUID identifying the schedule to which the subscription belongs.
	// Wire name: 'schedule_id'
	ScheduleId string ``
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	// Wire name: 'subscriber'
	Subscriber Subscriber ``
	// UUID identifying the subscription.
	// Wire name: 'subscription_id'
	SubscriptionId string ``
	// A timestamp indicating when the subscription was last updated.
	// Wire name: 'update_time'
	UpdateTime      *time.Time ``
	ForceSendFields []string   `tf:"-"`
}

func (s *Subscription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Subscription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SubscriptionToPb(st *Subscription) (*dashboardspb.SubscriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.SubscriptionPb{}
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
	pb.CreatedByUserId = st.CreatedByUserId
	pb.DashboardId = st.DashboardId
	pb.Etag = st.Etag
	pb.ScheduleId = st.ScheduleId
	subscriberPb, err := SubscriberToPb(&st.Subscriber)
	if err != nil {
		return nil, err
	}
	if subscriberPb != nil {
		pb.Subscriber = *subscriberPb
	}
	pb.SubscriptionId = st.SubscriptionId
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SubscriptionFromPb(pb *dashboardspb.SubscriptionPb) (*Subscription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Subscription{}
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
	st.CreatedByUserId = pb.CreatedByUserId
	st.DashboardId = pb.DashboardId
	st.Etag = pb.Etag
	st.ScheduleId = pb.ScheduleId
	subscriberField, err := SubscriberFromPb(&pb.Subscriber)
	if err != nil {
		return nil, err
	}
	if subscriberField != nil {
		st.Subscriber = *subscriberField
	}
	st.SubscriptionId = pb.SubscriptionId
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type SubscriptionSubscriberDestination struct {
	// The canonical identifier of the destination to receive email
	// notification.
	// Wire name: 'destination_id'
	DestinationId string ``
}

func SubscriptionSubscriberDestinationToPb(st *SubscriptionSubscriberDestination) (*dashboardspb.SubscriptionSubscriberDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.SubscriptionSubscriberDestinationPb{}
	pb.DestinationId = st.DestinationId

	return pb, nil
}

func SubscriptionSubscriberDestinationFromPb(pb *dashboardspb.SubscriptionSubscriberDestinationPb) (*SubscriptionSubscriberDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubscriptionSubscriberDestination{}
	st.DestinationId = pb.DestinationId

	return st, nil
}

type SubscriptionSubscriberUser struct {
	// UserId of the subscriber.
	// Wire name: 'user_id'
	UserId int64 ``
}

func SubscriptionSubscriberUserToPb(st *SubscriptionSubscriberUser) (*dashboardspb.SubscriptionSubscriberUserPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.SubscriptionSubscriberUserPb{}
	pb.UserId = st.UserId

	return pb, nil
}

func SubscriptionSubscriberUserFromPb(pb *dashboardspb.SubscriptionSubscriberUserPb) (*SubscriptionSubscriberUser, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubscriptionSubscriberUser{}
	st.UserId = pb.UserId

	return st, nil
}

type TextAttachment struct {
	// AI generated message
	// Wire name: 'content'
	Content string ``

	// Wire name: 'id'
	Id              string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *TextAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TextAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TextAttachmentToPb(st *TextAttachment) (*dashboardspb.TextAttachmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.TextAttachmentPb{}
	pb.Content = st.Content
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TextAttachmentFromPb(pb *dashboardspb.TextAttachmentPb) (*TextAttachment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TextAttachment{}
	st.Content = pb.Content
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TrashDashboardRequest struct {
	// UUID identifying the dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func TrashDashboardRequestToPb(st *TrashDashboardRequest) (*dashboardspb.TrashDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.TrashDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

func TrashDashboardRequestFromPb(pb *dashboardspb.TrashDashboardRequestPb) (*TrashDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrashDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

type UnpublishDashboardRequest struct {
	// UUID identifying the published dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func UnpublishDashboardRequestToPb(st *UnpublishDashboardRequest) (*dashboardspb.UnpublishDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.UnpublishDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

func UnpublishDashboardRequestFromPb(pb *dashboardspb.UnpublishDashboardRequestPb) (*UnpublishDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnpublishDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

type UpdateDashboardRequest struct {

	// Wire name: 'dashboard'
	Dashboard Dashboard ``
	// UUID identifying the dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func UpdateDashboardRequestToPb(st *UpdateDashboardRequest) (*dashboardspb.UpdateDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.UpdateDashboardRequestPb{}
	dashboardPb, err := DashboardToPb(&st.Dashboard)
	if err != nil {
		return nil, err
	}
	if dashboardPb != nil {
		pb.Dashboard = *dashboardPb
	}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

func UpdateDashboardRequestFromPb(pb *dashboardspb.UpdateDashboardRequestPb) (*UpdateDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDashboardRequest{}
	dashboardField, err := DashboardFromPb(&pb.Dashboard)
	if err != nil {
		return nil, err
	}
	if dashboardField != nil {
		st.Dashboard = *dashboardField
	}
	st.DashboardId = pb.DashboardId

	return st, nil
}

type UpdateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// The schedule to update.
	// Wire name: 'schedule'
	Schedule Schedule ``
	// UUID identifying the schedule.
	// Wire name: 'schedule_id'
	ScheduleId string `tf:"-"`
}

func UpdateScheduleRequestToPb(st *UpdateScheduleRequest) (*dashboardspb.UpdateScheduleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardspb.UpdateScheduleRequestPb{}
	pb.DashboardId = st.DashboardId
	schedulePb, err := ScheduleToPb(&st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = *schedulePb
	}
	pb.ScheduleId = st.ScheduleId

	return pb, nil
}

func UpdateScheduleRequestFromPb(pb *dashboardspb.UpdateScheduleRequestPb) (*UpdateScheduleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateScheduleRequest{}
	st.DashboardId = pb.DashboardId
	scheduleField, err := ScheduleFromPb(&pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = *scheduleField
	}
	st.ScheduleId = pb.ScheduleId

	return st, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%.9fs", d.Seconds())
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

// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/sql"
)

func identity[T any](obj *T) (*T, error) {
	return obj, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

// Helper to strip trailing zeros in fractional part
func rstripZeros(s string) string {
	for len(s) > 0 && s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '.' {
		s = s[:len(s)-1]
	}
	return s
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

type AuthorizationDetails struct {
	// Represents downscoped permission rules with specific access rights. This
	// field is specific to `workspace_rule_set` constraint.
	GrantRules []AuthorizationDetailsGrantRule
	// The acl path of the tree store resource resource.
	ResourceLegacyAclPath string
	// The resource name to which the authorization rule applies. This field is
	// specific to `workspace_rule_set` constraint. Format:
	// `workspaces/{workspace_id}/dashboards/{dashboard_id}`
	ResourceName string
	// The type of authorization downscoping policy. Ex: `workspace_rule_set`
	// defines access rules for a specific workspace resource
	Type string

	ForceSendFields []string
}

func authorizationDetailsToPb(st *AuthorizationDetails) (*authorizationDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &authorizationDetailsPb{}

	var grantRulesPb []authorizationDetailsGrantRulePb
	for _, item := range st.GrantRules {
		itemPb, err := authorizationDetailsGrantRuleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			grantRulesPb = append(grantRulesPb, *itemPb)
		}
	}
	pb.GrantRules = grantRulesPb

	resourceLegacyAclPathPb, err := identity(&st.ResourceLegacyAclPath)
	if err != nil {
		return nil, err
	}
	if resourceLegacyAclPathPb != nil {
		pb.ResourceLegacyAclPath = *resourceLegacyAclPathPb
	}

	resourceNamePb, err := identity(&st.ResourceName)
	if err != nil {
		return nil, err
	}
	if resourceNamePb != nil {
		pb.ResourceName = *resourceNamePb
	}

	typePb, err := identity(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type authorizationDetailsPb struct {
	// Represents downscoped permission rules with specific access rights. This
	// field is specific to `workspace_rule_set` constraint.
	GrantRules []authorizationDetailsGrantRulePb `json:"grant_rules,omitempty"`
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

func authorizationDetailsFromPb(pb *authorizationDetailsPb) (*AuthorizationDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AuthorizationDetails{}

	var grantRulesField []AuthorizationDetailsGrantRule
	for _, itemPb := range pb.GrantRules {
		item, err := authorizationDetailsGrantRuleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			grantRulesField = append(grantRulesField, *item)
		}
	}
	st.GrantRules = grantRulesField
	resourceLegacyAclPathField, err := identity(&pb.ResourceLegacyAclPath)
	if err != nil {
		return nil, err
	}
	if resourceLegacyAclPathField != nil {
		st.ResourceLegacyAclPath = *resourceLegacyAclPathField
	}
	resourceNameField, err := identity(&pb.ResourceName)
	if err != nil {
		return nil, err
	}
	if resourceNameField != nil {
		st.ResourceName = *resourceNameField
	}
	typeField, err := identity(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *authorizationDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st authorizationDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AuthorizationDetailsGrantRule struct {
	// Permission sets for dashboard are defined in
	// iam-common/rbac-common/permission-sets/definitions/TreeStoreBasePermissionSets
	// Ex: `permissionSets/dashboard.runner`
	PermissionSet string

	ForceSendFields []string
}

func authorizationDetailsGrantRuleToPb(st *AuthorizationDetailsGrantRule) (*authorizationDetailsGrantRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &authorizationDetailsGrantRulePb{}
	permissionSetPb, err := identity(&st.PermissionSet)
	if err != nil {
		return nil, err
	}
	if permissionSetPb != nil {
		pb.PermissionSet = *permissionSetPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type authorizationDetailsGrantRulePb struct {
	// Permission sets for dashboard are defined in
	// iam-common/rbac-common/permission-sets/definitions/TreeStoreBasePermissionSets
	// Ex: `permissionSets/dashboard.runner`
	PermissionSet string `json:"permission_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func authorizationDetailsGrantRuleFromPb(pb *authorizationDetailsGrantRulePb) (*AuthorizationDetailsGrantRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AuthorizationDetailsGrantRule{}
	permissionSetField, err := identity(&pb.PermissionSet)
	if err != nil {
		return nil, err
	}
	if permissionSetField != nil {
		st.PermissionSet = *permissionSetField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *authorizationDetailsGrantRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st authorizationDetailsGrantRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Cancel the results for the a query for a published, embedded dashboard
type CancelPublishedQueryExecutionRequest struct {
	DashboardName string

	DashboardRevisionId string
	// Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	Tokens []string
}

func cancelPublishedQueryExecutionRequestToPb(st *CancelPublishedQueryExecutionRequest) (*cancelPublishedQueryExecutionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelPublishedQueryExecutionRequestPb{}
	dashboardNamePb, err := identity(&st.DashboardName)
	if err != nil {
		return nil, err
	}
	if dashboardNamePb != nil {
		pb.DashboardName = *dashboardNamePb
	}

	dashboardRevisionIdPb, err := identity(&st.DashboardRevisionId)
	if err != nil {
		return nil, err
	}
	if dashboardRevisionIdPb != nil {
		pb.DashboardRevisionId = *dashboardRevisionIdPb
	}

	var tokensPb []string
	for _, item := range st.Tokens {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tokensPb = append(tokensPb, *itemPb)
		}
	}
	pb.Tokens = tokensPb

	return pb, nil
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

type cancelPublishedQueryExecutionRequestPb struct {
	DashboardName string `json:"-" url:"dashboard_name"`

	DashboardRevisionId string `json:"-" url:"dashboard_revision_id"`
	// Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	Tokens []string `json:"-" url:"tokens,omitempty"`
}

func cancelPublishedQueryExecutionRequestFromPb(pb *cancelPublishedQueryExecutionRequestPb) (*CancelPublishedQueryExecutionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelPublishedQueryExecutionRequest{}
	dashboardNameField, err := identity(&pb.DashboardName)
	if err != nil {
		return nil, err
	}
	if dashboardNameField != nil {
		st.DashboardName = *dashboardNameField
	}
	dashboardRevisionIdField, err := identity(&pb.DashboardRevisionId)
	if err != nil {
		return nil, err
	}
	if dashboardRevisionIdField != nil {
		st.DashboardRevisionId = *dashboardRevisionIdField
	}

	var tokensField []string
	for _, itemPb := range pb.Tokens {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tokensField = append(tokensField, *item)
		}
	}
	st.Tokens = tokensField

	return st, nil
}

type CancelQueryExecutionResponse struct {
	Status []CancelQueryExecutionResponseStatus
}

func cancelQueryExecutionResponseToPb(st *CancelQueryExecutionResponse) (*cancelQueryExecutionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelQueryExecutionResponsePb{}

	var statusPb []cancelQueryExecutionResponseStatusPb
	for _, item := range st.Status {
		itemPb, err := cancelQueryExecutionResponseStatusToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			statusPb = append(statusPb, *itemPb)
		}
	}
	pb.Status = statusPb

	return pb, nil
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

type cancelQueryExecutionResponsePb struct {
	Status []cancelQueryExecutionResponseStatusPb `json:"status,omitempty"`
}

func cancelQueryExecutionResponseFromPb(pb *cancelQueryExecutionResponsePb) (*CancelQueryExecutionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelQueryExecutionResponse{}

	var statusField []CancelQueryExecutionResponseStatus
	for _, itemPb := range pb.Status {
		item, err := cancelQueryExecutionResponseStatusFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			statusField = append(statusField, *item)
		}
	}
	st.Status = statusField

	return st, nil
}

type CancelQueryExecutionResponseStatus struct {
	// The token to poll for result asynchronously Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	DataToken string
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	Pending *Empty
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	Success *Empty
}

func cancelQueryExecutionResponseStatusToPb(st *CancelQueryExecutionResponseStatus) (*cancelQueryExecutionResponseStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelQueryExecutionResponseStatusPb{}
	dataTokenPb, err := identity(&st.DataToken)
	if err != nil {
		return nil, err
	}
	if dataTokenPb != nil {
		pb.DataToken = *dataTokenPb
	}

	pendingPb, err := emptyToPb(st.Pending)
	if err != nil {
		return nil, err
	}
	if pendingPb != nil {
		pb.Pending = pendingPb
	}

	successPb, err := emptyToPb(st.Success)
	if err != nil {
		return nil, err
	}
	if successPb != nil {
		pb.Success = successPb
	}

	return pb, nil
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

type cancelQueryExecutionResponseStatusPb struct {
	// The token to poll for result asynchronously Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	DataToken string `json:"data_token"`
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	Pending *emptyPb `json:"pending,omitempty"`
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	Success *emptyPb `json:"success,omitempty"`
}

func cancelQueryExecutionResponseStatusFromPb(pb *cancelQueryExecutionResponseStatusPb) (*CancelQueryExecutionResponseStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelQueryExecutionResponseStatus{}
	dataTokenField, err := identity(&pb.DataToken)
	if err != nil {
		return nil, err
	}
	if dataTokenField != nil {
		st.DataToken = *dataTokenField
	}
	pendingField, err := emptyFromPb(pb.Pending)
	if err != nil {
		return nil, err
	}
	if pendingField != nil {
		st.Pending = pendingField
	}
	successField, err := emptyFromPb(pb.Success)
	if err != nil {
		return nil, err
	}
	if successField != nil {
		st.Success = successField
	}

	return st, nil
}

// Create dashboard
type CreateDashboardRequest struct {
	Dashboard Dashboard
}

func createDashboardRequestToPb(st *CreateDashboardRequest) (*createDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createDashboardRequestPb{}
	dashboardPb, err := dashboardToPb(&st.Dashboard)
	if err != nil {
		return nil, err
	}
	if dashboardPb != nil {
		pb.Dashboard = *dashboardPb
	}

	return pb, nil
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

type createDashboardRequestPb struct {
	Dashboard dashboardPb `json:"dashboard"`
}

func createDashboardRequestFromPb(pb *createDashboardRequestPb) (*CreateDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDashboardRequest{}
	dashboardField, err := dashboardFromPb(&pb.Dashboard)
	if err != nil {
		return nil, err
	}
	if dashboardField != nil {
		st.Dashboard = *dashboardField
	}

	return st, nil
}

// Create dashboard schedule
type CreateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string

	Schedule Schedule
}

func createScheduleRequestToPb(st *CreateScheduleRequest) (*createScheduleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createScheduleRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	schedulePb, err := scheduleToPb(&st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = *schedulePb
	}

	return pb, nil
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

type createScheduleRequestPb struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`

	Schedule schedulePb `json:"schedule"`
}

func createScheduleRequestFromPb(pb *createScheduleRequestPb) (*CreateScheduleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateScheduleRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	scheduleField, err := scheduleFromPb(&pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = *scheduleField
	}

	return st, nil
}

// Create schedule subscription
type CreateSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId string
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId string

	Subscription Subscription
}

func createSubscriptionRequestToPb(st *CreateSubscriptionRequest) (*createSubscriptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createSubscriptionRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	scheduleIdPb, err := identity(&st.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdPb != nil {
		pb.ScheduleId = *scheduleIdPb
	}

	subscriptionPb, err := subscriptionToPb(&st.Subscription)
	if err != nil {
		return nil, err
	}
	if subscriptionPb != nil {
		pb.Subscription = *subscriptionPb
	}

	return pb, nil
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

type createSubscriptionRequestPb struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId string `json:"-" url:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId string `json:"-" url:"-"`

	Subscription subscriptionPb `json:"subscription"`
}

func createSubscriptionRequestFromPb(pb *createSubscriptionRequestPb) (*CreateSubscriptionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateSubscriptionRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	scheduleIdField, err := identity(&pb.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdField != nil {
		st.ScheduleId = *scheduleIdField
	}
	subscriptionField, err := subscriptionFromPb(&pb.Subscription)
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
	QuartzCronExpression string
	// A Java timezone id. The schedule will be resolved with respect to this
	// timezone. See [Java TimeZone] for details.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId string
}

func cronScheduleToPb(st *CronSchedule) (*cronSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cronSchedulePb{}
	quartzCronExpressionPb, err := identity(&st.QuartzCronExpression)
	if err != nil {
		return nil, err
	}
	if quartzCronExpressionPb != nil {
		pb.QuartzCronExpression = *quartzCronExpressionPb
	}

	timezoneIdPb, err := identity(&st.TimezoneId)
	if err != nil {
		return nil, err
	}
	if timezoneIdPb != nil {
		pb.TimezoneId = *timezoneIdPb
	}

	return pb, nil
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

type cronSchedulePb struct {
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

func cronScheduleFromPb(pb *cronSchedulePb) (*CronSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CronSchedule{}
	quartzCronExpressionField, err := identity(&pb.QuartzCronExpression)
	if err != nil {
		return nil, err
	}
	if quartzCronExpressionField != nil {
		st.QuartzCronExpression = *quartzCronExpressionField
	}
	timezoneIdField, err := identity(&pb.TimezoneId)
	if err != nil {
		return nil, err
	}
	if timezoneIdField != nil {
		st.TimezoneId = *timezoneIdField
	}

	return st, nil
}

type Dashboard struct {
	// The timestamp of when the dashboard was created.
	CreateTime string
	// UUID identifying the dashboard.
	DashboardId string
	// The display name of the dashboard.
	DisplayName string
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read. This
	// field is excluded in List Dashboards responses.
	Etag string
	// The state of the dashboard resource. Used for tracking trashed status.
	LifecycleState LifecycleState
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash. This field is excluded in List
	// Dashboards responses.
	ParentPath string
	// The workspace path of the dashboard asset, including the file name.
	// Exported dashboards always have the file extension `.lvdash.json`. This
	// field is excluded in List Dashboards responses.
	Path string
	// The contents of the dashboard in serialized string form. This field is
	// excluded in List Dashboards responses. Use the [get dashboard API] to
	// retrieve an example response, which includes the `serialized_dashboard`
	// field. This field provides the structure of the JSON string that
	// represents the dashboard's layout and components.
	//
	// [get dashboard API]: https://docs.databricks.com/api/workspace/lakeview/get
	SerializedDashboard string
	// The timestamp of when the dashboard was last updated by the user. This
	// field is excluded in List Dashboards responses.
	UpdateTime string
	// The warehouse ID used to run the dashboard.
	WarehouseId string

	ForceSendFields []string
}

func dashboardToPb(st *Dashboard) (*dashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardPb{}
	createTimePb, err := identity(&st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	lifecycleStatePb, err := identity(&st.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStatePb != nil {
		pb.LifecycleState = *lifecycleStatePb
	}

	parentPathPb, err := identity(&st.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathPb != nil {
		pb.ParentPath = *parentPathPb
	}

	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	serializedDashboardPb, err := identity(&st.SerializedDashboard)
	if err != nil {
		return nil, err
	}
	if serializedDashboardPb != nil {
		pb.SerializedDashboard = *serializedDashboardPb
	}

	updateTimePb, err := identity(&st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type dashboardPb struct {
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

func dashboardFromPb(pb *dashboardPb) (*Dashboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Dashboard{}
	createTimeField, err := identity(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	lifecycleStateField, err := identity(&pb.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStateField != nil {
		st.LifecycleState = *lifecycleStateField
	}
	parentPathField, err := identity(&pb.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathField != nil {
		st.ParentPath = *parentPathField
	}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}
	serializedDashboardField, err := identity(&pb.SerializedDashboard)
	if err != nil {
		return nil, err
	}
	if serializedDashboardField != nil {
		st.SerializedDashboard = *serializedDashboardField
	}
	updateTimeField, err := identity(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
	}
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dashboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dashboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	DashboardId string
	// The etag for the schedule. Optionally, it can be provided to verify that
	// the schedule has not been modified from its last retrieval.
	Etag string
	// UUID identifying the schedule.
	ScheduleId string

	ForceSendFields []string
}

func deleteScheduleRequestToPb(st *DeleteScheduleRequest) (*deleteScheduleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteScheduleRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	scheduleIdPb, err := identity(&st.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdPb != nil {
		pb.ScheduleId = *scheduleIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type deleteScheduleRequestPb struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// The etag for the schedule. Optionally, it can be provided to verify that
	// the schedule has not been modified from its last retrieval.
	Etag string `json:"-" url:"etag,omitempty"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteScheduleRequestFromPb(pb *deleteScheduleRequestPb) (*DeleteScheduleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteScheduleRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	scheduleIdField, err := identity(&pb.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdField != nil {
		st.ScheduleId = *scheduleIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteScheduleRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteScheduleRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteScheduleResponse struct {
}

func deleteScheduleResponseToPb(st *DeleteScheduleResponse) (*deleteScheduleResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteScheduleResponsePb{}

	return pb, nil
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

type deleteScheduleResponsePb struct {
}

func deleteScheduleResponseFromPb(pb *deleteScheduleResponsePb) (*DeleteScheduleResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteScheduleResponse{}

	return st, nil
}

// Delete schedule subscription
type DeleteSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId string
	// The etag for the subscription. Can be optionally provided to ensure that
	// the subscription has not been modified since the last read.
	Etag string
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId string
	// UUID identifying the subscription.
	SubscriptionId string

	ForceSendFields []string
}

func deleteSubscriptionRequestToPb(st *DeleteSubscriptionRequest) (*deleteSubscriptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSubscriptionRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	scheduleIdPb, err := identity(&st.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdPb != nil {
		pb.ScheduleId = *scheduleIdPb
	}

	subscriptionIdPb, err := identity(&st.SubscriptionId)
	if err != nil {
		return nil, err
	}
	if subscriptionIdPb != nil {
		pb.SubscriptionId = *subscriptionIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type deleteSubscriptionRequestPb struct {
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

func deleteSubscriptionRequestFromPb(pb *deleteSubscriptionRequestPb) (*DeleteSubscriptionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSubscriptionRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	scheduleIdField, err := identity(&pb.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdField != nil {
		st.ScheduleId = *scheduleIdField
	}
	subscriptionIdField, err := identity(&pb.SubscriptionId)
	if err != nil {
		return nil, err
	}
	if subscriptionIdField != nil {
		st.SubscriptionId = *subscriptionIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteSubscriptionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteSubscriptionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteSubscriptionResponse struct {
}

func deleteSubscriptionResponseToPb(st *DeleteSubscriptionResponse) (*deleteSubscriptionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSubscriptionResponsePb{}

	return pb, nil
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

type deleteSubscriptionResponsePb struct {
}

func deleteSubscriptionResponseFromPb(pb *deleteSubscriptionResponsePb) (*DeleteSubscriptionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSubscriptionResponse{}

	return st, nil
}

// Represents an empty message, similar to google.protobuf.Empty, which is not
// available in the firm right now.
type Empty struct {
}

func emptyToPb(st *Empty) (*emptyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &emptyPb{}

	return pb, nil
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

type emptyPb struct {
}

func emptyFromPb(pb *emptyPb) (*Empty, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Empty{}

	return st, nil
}

// Execute query request for published Dashboards. Since published dashboards
// have the option of running as the publisher, the datasets, warehouse_id are
// excluded from the request and instead read from the source (lakeview-config)
// via the additional parameters (dashboardName and dashboardRevisionId)
type ExecutePublishedDashboardQueryRequest struct {
	// Dashboard name and revision_id is required to retrieve
	// PublishedDatasetDataModel which contains the list of datasets,
	// warehouse_id, and embedded_credentials
	DashboardName string

	DashboardRevisionId string
	// A dashboard schedule can override the warehouse used as compute for
	// processing the published dashboard queries
	OverrideWarehouseId string

	ForceSendFields []string
}

func executePublishedDashboardQueryRequestToPb(st *ExecutePublishedDashboardQueryRequest) (*executePublishedDashboardQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &executePublishedDashboardQueryRequestPb{}
	dashboardNamePb, err := identity(&st.DashboardName)
	if err != nil {
		return nil, err
	}
	if dashboardNamePb != nil {
		pb.DashboardName = *dashboardNamePb
	}

	dashboardRevisionIdPb, err := identity(&st.DashboardRevisionId)
	if err != nil {
		return nil, err
	}
	if dashboardRevisionIdPb != nil {
		pb.DashboardRevisionId = *dashboardRevisionIdPb
	}

	overrideWarehouseIdPb, err := identity(&st.OverrideWarehouseId)
	if err != nil {
		return nil, err
	}
	if overrideWarehouseIdPb != nil {
		pb.OverrideWarehouseId = *overrideWarehouseIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type executePublishedDashboardQueryRequestPb struct {
	// Dashboard name and revision_id is required to retrieve
	// PublishedDatasetDataModel which contains the list of datasets,
	// warehouse_id, and embedded_credentials
	DashboardName string `json:"dashboard_name"`

	DashboardRevisionId string `json:"dashboard_revision_id"`
	// A dashboard schedule can override the warehouse used as compute for
	// processing the published dashboard queries
	OverrideWarehouseId string `json:"override_warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func executePublishedDashboardQueryRequestFromPb(pb *executePublishedDashboardQueryRequestPb) (*ExecutePublishedDashboardQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExecutePublishedDashboardQueryRequest{}
	dashboardNameField, err := identity(&pb.DashboardName)
	if err != nil {
		return nil, err
	}
	if dashboardNameField != nil {
		st.DashboardName = *dashboardNameField
	}
	dashboardRevisionIdField, err := identity(&pb.DashboardRevisionId)
	if err != nil {
		return nil, err
	}
	if dashboardRevisionIdField != nil {
		st.DashboardRevisionId = *dashboardRevisionIdField
	}
	overrideWarehouseIdField, err := identity(&pb.OverrideWarehouseId)
	if err != nil {
		return nil, err
	}
	if overrideWarehouseIdField != nil {
		st.OverrideWarehouseId = *overrideWarehouseIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *executePublishedDashboardQueryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st executePublishedDashboardQueryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExecuteQueryResponse struct {
}

func executeQueryResponseToPb(st *ExecuteQueryResponse) (*executeQueryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &executeQueryResponsePb{}

	return pb, nil
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

type executeQueryResponsePb struct {
}

func executeQueryResponseFromPb(pb *executeQueryResponsePb) (*ExecuteQueryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExecuteQueryResponse{}

	return st, nil
}

// Genie AI Response
type GenieAttachment struct {
	// Attachment ID
	AttachmentId string
	// Query Attachment if Genie responds with a SQL query
	Query *GenieQueryAttachment
	// Text Attachment if Genie responds with text
	Text *TextAttachment

	ForceSendFields []string
}

func genieAttachmentToPb(st *GenieAttachment) (*genieAttachmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieAttachmentPb{}
	attachmentIdPb, err := identity(&st.AttachmentId)
	if err != nil {
		return nil, err
	}
	if attachmentIdPb != nil {
		pb.AttachmentId = *attachmentIdPb
	}

	queryPb, err := genieQueryAttachmentToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}

	textPb, err := textAttachmentToPb(st.Text)
	if err != nil {
		return nil, err
	}
	if textPb != nil {
		pb.Text = textPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type genieAttachmentPb struct {
	// Attachment ID
	AttachmentId string `json:"attachment_id,omitempty"`
	// Query Attachment if Genie responds with a SQL query
	Query *genieQueryAttachmentPb `json:"query,omitempty"`
	// Text Attachment if Genie responds with text
	Text *textAttachmentPb `json:"text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieAttachmentFromPb(pb *genieAttachmentPb) (*GenieAttachment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieAttachment{}
	attachmentIdField, err := identity(&pb.AttachmentId)
	if err != nil {
		return nil, err
	}
	if attachmentIdField != nil {
		st.AttachmentId = *attachmentIdField
	}
	queryField, err := genieQueryAttachmentFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}
	textField, err := textAttachmentFromPb(pb.Text)
	if err != nil {
		return nil, err
	}
	if textField != nil {
		st.Text = textField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genieAttachmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieAttachmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieConversation struct {
	// Conversation ID
	ConversationId string
	// Timestamp when the message was created
	CreatedTimestamp int64
	// Conversation ID. Legacy identifier, use conversation_id instead
	Id string
	// Timestamp when the message was last updated
	LastUpdatedTimestamp int64
	// Genie space ID
	SpaceId string
	// Conversation title
	Title string
	// ID of the user who created the conversation
	UserId int

	ForceSendFields []string
}

func genieConversationToPb(st *GenieConversation) (*genieConversationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieConversationPb{}
	conversationIdPb, err := identity(&st.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdPb != nil {
		pb.ConversationId = *conversationIdPb
	}

	createdTimestampPb, err := identity(&st.CreatedTimestamp)
	if err != nil {
		return nil, err
	}
	if createdTimestampPb != nil {
		pb.CreatedTimestamp = *createdTimestampPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	lastUpdatedTimestampPb, err := identity(&st.LastUpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	if lastUpdatedTimestampPb != nil {
		pb.LastUpdatedTimestamp = *lastUpdatedTimestampPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	titlePb, err := identity(&st.Title)
	if err != nil {
		return nil, err
	}
	if titlePb != nil {
		pb.Title = *titlePb
	}

	userIdPb, err := identity(&st.UserId)
	if err != nil {
		return nil, err
	}
	if userIdPb != nil {
		pb.UserId = *userIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type genieConversationPb struct {
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

func genieConversationFromPb(pb *genieConversationPb) (*GenieConversation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieConversation{}
	conversationIdField, err := identity(&pb.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdField != nil {
		st.ConversationId = *conversationIdField
	}
	createdTimestampField, err := identity(&pb.CreatedTimestamp)
	if err != nil {
		return nil, err
	}
	if createdTimestampField != nil {
		st.CreatedTimestamp = *createdTimestampField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	lastUpdatedTimestampField, err := identity(&pb.LastUpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	if lastUpdatedTimestampField != nil {
		st.LastUpdatedTimestamp = *lastUpdatedTimestampField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}
	titleField, err := identity(&pb.Title)
	if err != nil {
		return nil, err
	}
	if titleField != nil {
		st.Title = *titleField
	}
	userIdField, err := identity(&pb.UserId)
	if err != nil {
		return nil, err
	}
	if userIdField != nil {
		st.UserId = *userIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genieConversationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieConversationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieCreateConversationMessageRequest struct {
	// User message content.
	Content string
	// The ID associated with the conversation.
	ConversationId string
	// The ID associated with the Genie space where the conversation is started.
	SpaceId string
}

func genieCreateConversationMessageRequestToPb(st *GenieCreateConversationMessageRequest) (*genieCreateConversationMessageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieCreateConversationMessageRequestPb{}
	contentPb, err := identity(&st.Content)
	if err != nil {
		return nil, err
	}
	if contentPb != nil {
		pb.Content = *contentPb
	}

	conversationIdPb, err := identity(&st.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdPb != nil {
		pb.ConversationId = *conversationIdPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	return pb, nil
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

type genieCreateConversationMessageRequestPb struct {
	// User message content.
	Content string `json:"content"`
	// The ID associated with the conversation.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the Genie space where the conversation is started.
	SpaceId string `json:"-" url:"-"`
}

func genieCreateConversationMessageRequestFromPb(pb *genieCreateConversationMessageRequestPb) (*GenieCreateConversationMessageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieCreateConversationMessageRequest{}
	contentField, err := identity(&pb.Content)
	if err != nil {
		return nil, err
	}
	if contentField != nil {
		st.Content = *contentField
	}
	conversationIdField, err := identity(&pb.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdField != nil {
		st.ConversationId = *conversationIdField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}

	return st, nil
}

// Execute message attachment SQL query
type GenieExecuteMessageAttachmentQueryRequest struct {
	// Attachment ID
	AttachmentId string
	// Conversation ID
	ConversationId string
	// Message ID
	MessageId string
	// Genie space ID
	SpaceId string
}

func genieExecuteMessageAttachmentQueryRequestToPb(st *GenieExecuteMessageAttachmentQueryRequest) (*genieExecuteMessageAttachmentQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieExecuteMessageAttachmentQueryRequestPb{}
	attachmentIdPb, err := identity(&st.AttachmentId)
	if err != nil {
		return nil, err
	}
	if attachmentIdPb != nil {
		pb.AttachmentId = *attachmentIdPb
	}

	conversationIdPb, err := identity(&st.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdPb != nil {
		pb.ConversationId = *conversationIdPb
	}

	messageIdPb, err := identity(&st.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdPb != nil {
		pb.MessageId = *messageIdPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	return pb, nil
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

type genieExecuteMessageAttachmentQueryRequestPb struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func genieExecuteMessageAttachmentQueryRequestFromPb(pb *genieExecuteMessageAttachmentQueryRequestPb) (*GenieExecuteMessageAttachmentQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieExecuteMessageAttachmentQueryRequest{}
	attachmentIdField, err := identity(&pb.AttachmentId)
	if err != nil {
		return nil, err
	}
	if attachmentIdField != nil {
		st.AttachmentId = *attachmentIdField
	}
	conversationIdField, err := identity(&pb.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdField != nil {
		st.ConversationId = *conversationIdField
	}
	messageIdField, err := identity(&pb.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdField != nil {
		st.MessageId = *messageIdField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}

	return st, nil
}

// [Deprecated] Execute SQL query in a conversation message
type GenieExecuteMessageQueryRequest struct {
	// Conversation ID
	ConversationId string
	// Message ID
	MessageId string
	// Genie space ID
	SpaceId string
}

func genieExecuteMessageQueryRequestToPb(st *GenieExecuteMessageQueryRequest) (*genieExecuteMessageQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieExecuteMessageQueryRequestPb{}
	conversationIdPb, err := identity(&st.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdPb != nil {
		pb.ConversationId = *conversationIdPb
	}

	messageIdPb, err := identity(&st.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdPb != nil {
		pb.MessageId = *messageIdPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	return pb, nil
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

type genieExecuteMessageQueryRequestPb struct {
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func genieExecuteMessageQueryRequestFromPb(pb *genieExecuteMessageQueryRequestPb) (*GenieExecuteMessageQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieExecuteMessageQueryRequest{}
	conversationIdField, err := identity(&pb.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdField != nil {
		st.ConversationId = *conversationIdField
	}
	messageIdField, err := identity(&pb.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdField != nil {
		st.MessageId = *messageIdField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}

	return st, nil
}

// Generate full query result download
type GenieGenerateDownloadFullQueryResultRequest struct {
	// Attachment ID
	AttachmentId string
	// Conversation ID
	ConversationId string
	// Message ID
	MessageId string
	// Genie space ID
	SpaceId string
}

func genieGenerateDownloadFullQueryResultRequestToPb(st *GenieGenerateDownloadFullQueryResultRequest) (*genieGenerateDownloadFullQueryResultRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGenerateDownloadFullQueryResultRequestPb{}
	attachmentIdPb, err := identity(&st.AttachmentId)
	if err != nil {
		return nil, err
	}
	if attachmentIdPb != nil {
		pb.AttachmentId = *attachmentIdPb
	}

	conversationIdPb, err := identity(&st.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdPb != nil {
		pb.ConversationId = *conversationIdPb
	}

	messageIdPb, err := identity(&st.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdPb != nil {
		pb.MessageId = *messageIdPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	return pb, nil
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

type genieGenerateDownloadFullQueryResultRequestPb struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func genieGenerateDownloadFullQueryResultRequestFromPb(pb *genieGenerateDownloadFullQueryResultRequestPb) (*GenieGenerateDownloadFullQueryResultRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGenerateDownloadFullQueryResultRequest{}
	attachmentIdField, err := identity(&pb.AttachmentId)
	if err != nil {
		return nil, err
	}
	if attachmentIdField != nil {
		st.AttachmentId = *attachmentIdField
	}
	conversationIdField, err := identity(&pb.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdField != nil {
		st.ConversationId = *conversationIdField
	}
	messageIdField, err := identity(&pb.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdField != nil {
		st.MessageId = *messageIdField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}

	return st, nil
}

type GenieGenerateDownloadFullQueryResultResponse struct {
	// Download ID. Use this ID to track the download request in subsequent
	// polling calls
	DownloadId string

	ForceSendFields []string
}

func genieGenerateDownloadFullQueryResultResponseToPb(st *GenieGenerateDownloadFullQueryResultResponse) (*genieGenerateDownloadFullQueryResultResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGenerateDownloadFullQueryResultResponsePb{}
	downloadIdPb, err := identity(&st.DownloadId)
	if err != nil {
		return nil, err
	}
	if downloadIdPb != nil {
		pb.DownloadId = *downloadIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type genieGenerateDownloadFullQueryResultResponsePb struct {
	// Download ID. Use this ID to track the download request in subsequent
	// polling calls
	DownloadId string `json:"download_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieGenerateDownloadFullQueryResultResponseFromPb(pb *genieGenerateDownloadFullQueryResultResponsePb) (*GenieGenerateDownloadFullQueryResultResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGenerateDownloadFullQueryResultResponse{}
	downloadIdField, err := identity(&pb.DownloadId)
	if err != nil {
		return nil, err
	}
	if downloadIdField != nil {
		st.DownloadId = *downloadIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genieGenerateDownloadFullQueryResultResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieGenerateDownloadFullQueryResultResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get conversation message
type GenieGetConversationMessageRequest struct {
	// The ID associated with the target conversation.
	ConversationId string
	// The ID associated with the target message from the identified
	// conversation.
	MessageId string
	// The ID associated with the Genie space where the target conversation is
	// located.
	SpaceId string
}

func genieGetConversationMessageRequestToPb(st *GenieGetConversationMessageRequest) (*genieGetConversationMessageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetConversationMessageRequestPb{}
	conversationIdPb, err := identity(&st.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdPb != nil {
		pb.ConversationId = *conversationIdPb
	}

	messageIdPb, err := identity(&st.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdPb != nil {
		pb.MessageId = *messageIdPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	return pb, nil
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

type genieGetConversationMessageRequestPb struct {
	// The ID associated with the target conversation.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the target message from the identified
	// conversation.
	MessageId string `json:"-" url:"-"`
	// The ID associated with the Genie space where the target conversation is
	// located.
	SpaceId string `json:"-" url:"-"`
}

func genieGetConversationMessageRequestFromPb(pb *genieGetConversationMessageRequestPb) (*GenieGetConversationMessageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetConversationMessageRequest{}
	conversationIdField, err := identity(&pb.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdField != nil {
		st.ConversationId = *conversationIdField
	}
	messageIdField, err := identity(&pb.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdField != nil {
		st.MessageId = *messageIdField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}

	return st, nil
}

// Get download full query result
type GenieGetDownloadFullQueryResultRequest struct {
	// Attachment ID
	AttachmentId string
	// Conversation ID
	ConversationId string
	// Download ID. This ID is provided by the [Generate Download
	// endpoint](:method:genie/generateDownloadFullQueryResult)
	DownloadId string
	// Message ID
	MessageId string
	// Genie space ID
	SpaceId string
}

func genieGetDownloadFullQueryResultRequestToPb(st *GenieGetDownloadFullQueryResultRequest) (*genieGetDownloadFullQueryResultRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetDownloadFullQueryResultRequestPb{}
	attachmentIdPb, err := identity(&st.AttachmentId)
	if err != nil {
		return nil, err
	}
	if attachmentIdPb != nil {
		pb.AttachmentId = *attachmentIdPb
	}

	conversationIdPb, err := identity(&st.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdPb != nil {
		pb.ConversationId = *conversationIdPb
	}

	downloadIdPb, err := identity(&st.DownloadId)
	if err != nil {
		return nil, err
	}
	if downloadIdPb != nil {
		pb.DownloadId = *downloadIdPb
	}

	messageIdPb, err := identity(&st.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdPb != nil {
		pb.MessageId = *messageIdPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	return pb, nil
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

type genieGetDownloadFullQueryResultRequestPb struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Download ID. This ID is provided by the [Generate Download
	// endpoint](:method:genie/generateDownloadFullQueryResult)
	DownloadId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func genieGetDownloadFullQueryResultRequestFromPb(pb *genieGetDownloadFullQueryResultRequestPb) (*GenieGetDownloadFullQueryResultRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetDownloadFullQueryResultRequest{}
	attachmentIdField, err := identity(&pb.AttachmentId)
	if err != nil {
		return nil, err
	}
	if attachmentIdField != nil {
		st.AttachmentId = *attachmentIdField
	}
	conversationIdField, err := identity(&pb.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdField != nil {
		st.ConversationId = *conversationIdField
	}
	downloadIdField, err := identity(&pb.DownloadId)
	if err != nil {
		return nil, err
	}
	if downloadIdField != nil {
		st.DownloadId = *downloadIdField
	}
	messageIdField, err := identity(&pb.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdField != nil {
		st.MessageId = *messageIdField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}

	return st, nil
}

type GenieGetDownloadFullQueryResultResponse struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	StatementResponse *sql.StatementResponse
}

func genieGetDownloadFullQueryResultResponseToPb(st *GenieGetDownloadFullQueryResultResponse) (*genieGetDownloadFullQueryResultResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetDownloadFullQueryResultResponsePb{}
	statementResponsePb, err := sql.StatementResponseToPb(st.StatementResponse)
	if err != nil {
		return nil, err
	}
	if statementResponsePb != nil {
		pb.StatementResponse = statementResponsePb
	}

	return pb, nil
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

type genieGetDownloadFullQueryResultResponsePb struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	StatementResponse *sql.StatementResponsePb `json:"statement_response,omitempty"`
}

func genieGetDownloadFullQueryResultResponseFromPb(pb *genieGetDownloadFullQueryResultResponsePb) (*GenieGetDownloadFullQueryResultResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetDownloadFullQueryResultResponse{}
	statementResponseField, err := sql.StatementResponseFromPb(pb.StatementResponse)
	if err != nil {
		return nil, err
	}
	if statementResponseField != nil {
		st.StatementResponse = statementResponseField
	}

	return st, nil
}

// Get message attachment SQL query result
type GenieGetMessageAttachmentQueryResultRequest struct {
	// Attachment ID
	AttachmentId string
	// Conversation ID
	ConversationId string
	// Message ID
	MessageId string
	// Genie space ID
	SpaceId string
}

func genieGetMessageAttachmentQueryResultRequestToPb(st *GenieGetMessageAttachmentQueryResultRequest) (*genieGetMessageAttachmentQueryResultRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetMessageAttachmentQueryResultRequestPb{}
	attachmentIdPb, err := identity(&st.AttachmentId)
	if err != nil {
		return nil, err
	}
	if attachmentIdPb != nil {
		pb.AttachmentId = *attachmentIdPb
	}

	conversationIdPb, err := identity(&st.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdPb != nil {
		pb.ConversationId = *conversationIdPb
	}

	messageIdPb, err := identity(&st.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdPb != nil {
		pb.MessageId = *messageIdPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	return pb, nil
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

type genieGetMessageAttachmentQueryResultRequestPb struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func genieGetMessageAttachmentQueryResultRequestFromPb(pb *genieGetMessageAttachmentQueryResultRequestPb) (*GenieGetMessageAttachmentQueryResultRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetMessageAttachmentQueryResultRequest{}
	attachmentIdField, err := identity(&pb.AttachmentId)
	if err != nil {
		return nil, err
	}
	if attachmentIdField != nil {
		st.AttachmentId = *attachmentIdField
	}
	conversationIdField, err := identity(&pb.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdField != nil {
		st.ConversationId = *conversationIdField
	}
	messageIdField, err := identity(&pb.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdField != nil {
		st.MessageId = *messageIdField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}

	return st, nil
}

// [Deprecated] Get conversation message SQL query result
type GenieGetMessageQueryResultRequest struct {
	// Conversation ID
	ConversationId string
	// Message ID
	MessageId string
	// Genie space ID
	SpaceId string
}

func genieGetMessageQueryResultRequestToPb(st *GenieGetMessageQueryResultRequest) (*genieGetMessageQueryResultRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetMessageQueryResultRequestPb{}
	conversationIdPb, err := identity(&st.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdPb != nil {
		pb.ConversationId = *conversationIdPb
	}

	messageIdPb, err := identity(&st.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdPb != nil {
		pb.MessageId = *messageIdPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	return pb, nil
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

type genieGetMessageQueryResultRequestPb struct {
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func genieGetMessageQueryResultRequestFromPb(pb *genieGetMessageQueryResultRequestPb) (*GenieGetMessageQueryResultRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetMessageQueryResultRequest{}
	conversationIdField, err := identity(&pb.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdField != nil {
		st.ConversationId = *conversationIdField
	}
	messageIdField, err := identity(&pb.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdField != nil {
		st.MessageId = *messageIdField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}

	return st, nil
}

type GenieGetMessageQueryResultResponse struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	StatementResponse *sql.StatementResponse
}

func genieGetMessageQueryResultResponseToPb(st *GenieGetMessageQueryResultResponse) (*genieGetMessageQueryResultResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetMessageQueryResultResponsePb{}
	statementResponsePb, err := sql.StatementResponseToPb(st.StatementResponse)
	if err != nil {
		return nil, err
	}
	if statementResponsePb != nil {
		pb.StatementResponse = statementResponsePb
	}

	return pb, nil
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

type genieGetMessageQueryResultResponsePb struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	StatementResponse *sql.StatementResponsePb `json:"statement_response,omitempty"`
}

func genieGetMessageQueryResultResponseFromPb(pb *genieGetMessageQueryResultResponsePb) (*GenieGetMessageQueryResultResponse, error) {
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

// [Deprecated] Get conversation message SQL query result
type GenieGetQueryResultByAttachmentRequest struct {
	// Attachment ID
	AttachmentId string
	// Conversation ID
	ConversationId string
	// Message ID
	MessageId string
	// Genie space ID
	SpaceId string
}

func genieGetQueryResultByAttachmentRequestToPb(st *GenieGetQueryResultByAttachmentRequest) (*genieGetQueryResultByAttachmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetQueryResultByAttachmentRequestPb{}
	attachmentIdPb, err := identity(&st.AttachmentId)
	if err != nil {
		return nil, err
	}
	if attachmentIdPb != nil {
		pb.AttachmentId = *attachmentIdPb
	}

	conversationIdPb, err := identity(&st.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdPb != nil {
		pb.ConversationId = *conversationIdPb
	}

	messageIdPb, err := identity(&st.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdPb != nil {
		pb.MessageId = *messageIdPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	return pb, nil
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

type genieGetQueryResultByAttachmentRequestPb struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func genieGetQueryResultByAttachmentRequestFromPb(pb *genieGetQueryResultByAttachmentRequestPb) (*GenieGetQueryResultByAttachmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetQueryResultByAttachmentRequest{}
	attachmentIdField, err := identity(&pb.AttachmentId)
	if err != nil {
		return nil, err
	}
	if attachmentIdField != nil {
		st.AttachmentId = *attachmentIdField
	}
	conversationIdField, err := identity(&pb.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdField != nil {
		st.ConversationId = *conversationIdField
	}
	messageIdField, err := identity(&pb.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdField != nil {
		st.MessageId = *messageIdField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}

	return st, nil
}

// Get Genie Space
type GenieGetSpaceRequest struct {
	// The ID associated with the Genie space
	SpaceId string
}

func genieGetSpaceRequestToPb(st *GenieGetSpaceRequest) (*genieGetSpaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetSpaceRequestPb{}
	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	return pb, nil
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

type genieGetSpaceRequestPb struct {
	// The ID associated with the Genie space
	SpaceId string `json:"-" url:"-"`
}

func genieGetSpaceRequestFromPb(pb *genieGetSpaceRequestPb) (*GenieGetSpaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetSpaceRequest{}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}

	return st, nil
}

type GenieMessage struct {
	// AI-generated response to the message
	Attachments []GenieAttachment
	// User message content
	Content string
	// Conversation ID
	ConversationId string
	// Timestamp when the message was created
	CreatedTimestamp int64
	// Error message if Genie failed to respond to the message
	Error *MessageError
	// Message ID. Legacy identifier, use message_id instead
	Id string
	// Timestamp when the message was last updated
	LastUpdatedTimestamp int64
	// Message ID
	MessageId string
	// The result of SQL query if the message includes a query attachment.
	// Deprecated. Use `query_result_metadata` in `GenieQueryAttachment`
	// instead.
	QueryResult *Result
	// Genie space ID
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
	Status MessageStatus
	// ID of the user who created the message
	UserId int64

	ForceSendFields []string
}

func genieMessageToPb(st *GenieMessage) (*genieMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieMessagePb{}

	var attachmentsPb []genieAttachmentPb
	for _, item := range st.Attachments {
		itemPb, err := genieAttachmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			attachmentsPb = append(attachmentsPb, *itemPb)
		}
	}
	pb.Attachments = attachmentsPb

	contentPb, err := identity(&st.Content)
	if err != nil {
		return nil, err
	}
	if contentPb != nil {
		pb.Content = *contentPb
	}

	conversationIdPb, err := identity(&st.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdPb != nil {
		pb.ConversationId = *conversationIdPb
	}

	createdTimestampPb, err := identity(&st.CreatedTimestamp)
	if err != nil {
		return nil, err
	}
	if createdTimestampPb != nil {
		pb.CreatedTimestamp = *createdTimestampPb
	}

	errorPb, err := messageErrorToPb(st.Error)
	if err != nil {
		return nil, err
	}
	if errorPb != nil {
		pb.Error = errorPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	lastUpdatedTimestampPb, err := identity(&st.LastUpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	if lastUpdatedTimestampPb != nil {
		pb.LastUpdatedTimestamp = *lastUpdatedTimestampPb
	}

	messageIdPb, err := identity(&st.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdPb != nil {
		pb.MessageId = *messageIdPb
	}

	queryResultPb, err := resultToPb(st.QueryResult)
	if err != nil {
		return nil, err
	}
	if queryResultPb != nil {
		pb.QueryResult = queryResultPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	userIdPb, err := identity(&st.UserId)
	if err != nil {
		return nil, err
	}
	if userIdPb != nil {
		pb.UserId = *userIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type genieMessagePb struct {
	// AI-generated response to the message
	Attachments []genieAttachmentPb `json:"attachments,omitempty"`
	// User message content
	Content string `json:"content"`
	// Conversation ID
	ConversationId string `json:"conversation_id"`
	// Timestamp when the message was created
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// Error message if Genie failed to respond to the message
	Error *messageErrorPb `json:"error,omitempty"`
	// Message ID. Legacy identifier, use message_id instead
	Id string `json:"id"`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Message ID
	MessageId string `json:"message_id"`
	// The result of SQL query if the message includes a query attachment.
	// Deprecated. Use `query_result_metadata` in `GenieQueryAttachment`
	// instead.
	QueryResult *resultPb `json:"query_result,omitempty"`
	// Genie space ID
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
	Status MessageStatus `json:"status,omitempty"`
	// ID of the user who created the message
	UserId int64 `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieMessageFromPb(pb *genieMessagePb) (*GenieMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieMessage{}

	var attachmentsField []GenieAttachment
	for _, itemPb := range pb.Attachments {
		item, err := genieAttachmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			attachmentsField = append(attachmentsField, *item)
		}
	}
	st.Attachments = attachmentsField
	contentField, err := identity(&pb.Content)
	if err != nil {
		return nil, err
	}
	if contentField != nil {
		st.Content = *contentField
	}
	conversationIdField, err := identity(&pb.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdField != nil {
		st.ConversationId = *conversationIdField
	}
	createdTimestampField, err := identity(&pb.CreatedTimestamp)
	if err != nil {
		return nil, err
	}
	if createdTimestampField != nil {
		st.CreatedTimestamp = *createdTimestampField
	}
	errorField, err := messageErrorFromPb(pb.Error)
	if err != nil {
		return nil, err
	}
	if errorField != nil {
		st.Error = errorField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	lastUpdatedTimestampField, err := identity(&pb.LastUpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	if lastUpdatedTimestampField != nil {
		st.LastUpdatedTimestamp = *lastUpdatedTimestampField
	}
	messageIdField, err := identity(&pb.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdField != nil {
		st.MessageId = *messageIdField
	}
	queryResultField, err := resultFromPb(pb.QueryResult)
	if err != nil {
		return nil, err
	}
	if queryResultField != nil {
		st.QueryResult = queryResultField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	userIdField, err := identity(&pb.UserId)
	if err != nil {
		return nil, err
	}
	if userIdField != nil {
		st.UserId = *userIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genieMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieQueryAttachment struct {
	// Description of the query
	Description string

	Id string
	// Time when the user updated the query last
	LastUpdatedTimestamp int64
	// AI generated SQL query
	Query string
	// Metadata associated with the query result.
	QueryResultMetadata *GenieResultMetadata
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId string
	// Name of the query
	Title string

	ForceSendFields []string
}

func genieQueryAttachmentToPb(st *GenieQueryAttachment) (*genieQueryAttachmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieQueryAttachmentPb{}
	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	lastUpdatedTimestampPb, err := identity(&st.LastUpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	if lastUpdatedTimestampPb != nil {
		pb.LastUpdatedTimestamp = *lastUpdatedTimestampPb
	}

	queryPb, err := identity(&st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = *queryPb
	}

	queryResultMetadataPb, err := genieResultMetadataToPb(st.QueryResultMetadata)
	if err != nil {
		return nil, err
	}
	if queryResultMetadataPb != nil {
		pb.QueryResultMetadata = queryResultMetadataPb
	}

	statementIdPb, err := identity(&st.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdPb != nil {
		pb.StatementId = *statementIdPb
	}

	titlePb, err := identity(&st.Title)
	if err != nil {
		return nil, err
	}
	if titlePb != nil {
		pb.Title = *titlePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type genieQueryAttachmentPb struct {
	// Description of the query
	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`
	// Time when the user updated the query last
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// AI generated SQL query
	Query string `json:"query,omitempty"`
	// Metadata associated with the query result.
	QueryResultMetadata *genieResultMetadataPb `json:"query_result_metadata,omitempty"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId string `json:"statement_id,omitempty"`
	// Name of the query
	Title string `json:"title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieQueryAttachmentFromPb(pb *genieQueryAttachmentPb) (*GenieQueryAttachment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieQueryAttachment{}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	lastUpdatedTimestampField, err := identity(&pb.LastUpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	if lastUpdatedTimestampField != nil {
		st.LastUpdatedTimestamp = *lastUpdatedTimestampField
	}
	queryField, err := identity(&pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = *queryField
	}
	queryResultMetadataField, err := genieResultMetadataFromPb(pb.QueryResultMetadata)
	if err != nil {
		return nil, err
	}
	if queryResultMetadataField != nil {
		st.QueryResultMetadata = queryResultMetadataField
	}
	statementIdField, err := identity(&pb.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdField != nil {
		st.StatementId = *statementIdField
	}
	titleField, err := identity(&pb.Title)
	if err != nil {
		return nil, err
	}
	if titleField != nil {
		st.Title = *titleField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genieQueryAttachmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieQueryAttachmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieResultMetadata struct {
	// Indicates whether the result set is truncated.
	IsTruncated bool
	// The number of rows in the result set.
	RowCount int64

	ForceSendFields []string
}

func genieResultMetadataToPb(st *GenieResultMetadata) (*genieResultMetadataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieResultMetadataPb{}
	isTruncatedPb, err := identity(&st.IsTruncated)
	if err != nil {
		return nil, err
	}
	if isTruncatedPb != nil {
		pb.IsTruncated = *isTruncatedPb
	}

	rowCountPb, err := identity(&st.RowCount)
	if err != nil {
		return nil, err
	}
	if rowCountPb != nil {
		pb.RowCount = *rowCountPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type genieResultMetadataPb struct {
	// Indicates whether the result set is truncated.
	IsTruncated bool `json:"is_truncated,omitempty"`
	// The number of rows in the result set.
	RowCount int64 `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieResultMetadataFromPb(pb *genieResultMetadataPb) (*GenieResultMetadata, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieResultMetadata{}
	isTruncatedField, err := identity(&pb.IsTruncated)
	if err != nil {
		return nil, err
	}
	if isTruncatedField != nil {
		st.IsTruncated = *isTruncatedField
	}
	rowCountField, err := identity(&pb.RowCount)
	if err != nil {
		return nil, err
	}
	if rowCountField != nil {
		st.RowCount = *rowCountField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genieResultMetadataPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieResultMetadataPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieSpace struct {
	// Description of the Genie Space
	Description string
	// Genie space ID
	SpaceId string
	// Title of the Genie Space
	Title string

	ForceSendFields []string
}

func genieSpaceToPb(st *GenieSpace) (*genieSpacePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieSpacePb{}
	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	titlePb, err := identity(&st.Title)
	if err != nil {
		return nil, err
	}
	if titlePb != nil {
		pb.Title = *titlePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type genieSpacePb struct {
	// Description of the Genie Space
	Description string `json:"description,omitempty"`
	// Genie space ID
	SpaceId string `json:"space_id"`
	// Title of the Genie Space
	Title string `json:"title"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieSpaceFromPb(pb *genieSpacePb) (*GenieSpace, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieSpace{}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}
	titleField, err := identity(&pb.Title)
	if err != nil {
		return nil, err
	}
	if titleField != nil {
		st.Title = *titleField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genieSpacePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieSpacePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GenieStartConversationMessageRequest struct {
	// The text of the message that starts the conversation.
	Content string
	// The ID associated with the Genie space where you want to start a
	// conversation.
	SpaceId string
}

func genieStartConversationMessageRequestToPb(st *GenieStartConversationMessageRequest) (*genieStartConversationMessageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieStartConversationMessageRequestPb{}
	contentPb, err := identity(&st.Content)
	if err != nil {
		return nil, err
	}
	if contentPb != nil {
		pb.Content = *contentPb
	}

	spaceIdPb, err := identity(&st.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdPb != nil {
		pb.SpaceId = *spaceIdPb
	}

	return pb, nil
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

type genieStartConversationMessageRequestPb struct {
	// The text of the message that starts the conversation.
	Content string `json:"content"`
	// The ID associated with the Genie space where you want to start a
	// conversation.
	SpaceId string `json:"-" url:"-"`
}

func genieStartConversationMessageRequestFromPb(pb *genieStartConversationMessageRequestPb) (*GenieStartConversationMessageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieStartConversationMessageRequest{}
	contentField, err := identity(&pb.Content)
	if err != nil {
		return nil, err
	}
	if contentField != nil {
		st.Content = *contentField
	}
	spaceIdField, err := identity(&pb.SpaceId)
	if err != nil {
		return nil, err
	}
	if spaceIdField != nil {
		st.SpaceId = *spaceIdField
	}

	return st, nil
}

type GenieStartConversationResponse struct {
	Conversation *GenieConversation
	// Conversation ID
	ConversationId string

	Message *GenieMessage
	// Message ID
	MessageId string
}

func genieStartConversationResponseToPb(st *GenieStartConversationResponse) (*genieStartConversationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieStartConversationResponsePb{}
	conversationPb, err := genieConversationToPb(st.Conversation)
	if err != nil {
		return nil, err
	}
	if conversationPb != nil {
		pb.Conversation = conversationPb
	}

	conversationIdPb, err := identity(&st.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdPb != nil {
		pb.ConversationId = *conversationIdPb
	}

	messagePb, err := genieMessageToPb(st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = messagePb
	}

	messageIdPb, err := identity(&st.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdPb != nil {
		pb.MessageId = *messageIdPb
	}

	return pb, nil
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

type genieStartConversationResponsePb struct {
	Conversation *genieConversationPb `json:"conversation,omitempty"`
	// Conversation ID
	ConversationId string `json:"conversation_id"`

	Message *genieMessagePb `json:"message,omitempty"`
	// Message ID
	MessageId string `json:"message_id"`
}

func genieStartConversationResponseFromPb(pb *genieStartConversationResponsePb) (*GenieStartConversationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieStartConversationResponse{}
	conversationField, err := genieConversationFromPb(pb.Conversation)
	if err != nil {
		return nil, err
	}
	if conversationField != nil {
		st.Conversation = conversationField
	}
	conversationIdField, err := identity(&pb.ConversationId)
	if err != nil {
		return nil, err
	}
	if conversationIdField != nil {
		st.ConversationId = *conversationIdField
	}
	messageField, err := genieMessageFromPb(pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = messageField
	}
	messageIdField, err := identity(&pb.MessageId)
	if err != nil {
		return nil, err
	}
	if messageIdField != nil {
		st.MessageId = *messageIdField
	}

	return st, nil
}

// Get dashboard
type GetDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string
}

func getDashboardRequestToPb(st *GetDashboardRequest) (*getDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDashboardRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	return pb, nil
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

type getDashboardRequestPb struct {
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
}

func getDashboardRequestFromPb(pb *getDashboardRequestPb) (*GetDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDashboardRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}

	return st, nil
}

// Read a published dashboard in an embedded ui.
type GetPublishedDashboardEmbeddedRequest struct {
	// UUID identifying the published dashboard.
	DashboardId string
}

func getPublishedDashboardEmbeddedRequestToPb(st *GetPublishedDashboardEmbeddedRequest) (*getPublishedDashboardEmbeddedRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedDashboardEmbeddedRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	return pb, nil
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

type getPublishedDashboardEmbeddedRequestPb struct {
	// UUID identifying the published dashboard.
	DashboardId string `json:"-" url:"-"`
}

func getPublishedDashboardEmbeddedRequestFromPb(pb *getPublishedDashboardEmbeddedRequestPb) (*GetPublishedDashboardEmbeddedRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedDashboardEmbeddedRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}

	return st, nil
}

type GetPublishedDashboardEmbeddedResponse struct {
}

func getPublishedDashboardEmbeddedResponseToPb(st *GetPublishedDashboardEmbeddedResponse) (*getPublishedDashboardEmbeddedResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedDashboardEmbeddedResponsePb{}

	return pb, nil
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

type getPublishedDashboardEmbeddedResponsePb struct {
}

func getPublishedDashboardEmbeddedResponseFromPb(pb *getPublishedDashboardEmbeddedResponsePb) (*GetPublishedDashboardEmbeddedResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedDashboardEmbeddedResponse{}

	return st, nil
}

// Get published dashboard
type GetPublishedDashboardRequest struct {
	// UUID identifying the published dashboard.
	DashboardId string
}

func getPublishedDashboardRequestToPb(st *GetPublishedDashboardRequest) (*getPublishedDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedDashboardRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	return pb, nil
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

type getPublishedDashboardRequestPb struct {
	// UUID identifying the published dashboard.
	DashboardId string `json:"-" url:"-"`
}

func getPublishedDashboardRequestFromPb(pb *getPublishedDashboardRequestPb) (*GetPublishedDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedDashboardRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}

	return st, nil
}

// Read an information of a published dashboard to mint an OAuth token.
type GetPublishedDashboardTokenInfoRequest struct {
	// UUID identifying the published dashboard.
	DashboardId string
	// Provided external value to be included in the custom claim.
	ExternalValue string
	// Provided external viewer id to be included in the custom claim.
	ExternalViewerId string

	ForceSendFields []string
}

func getPublishedDashboardTokenInfoRequestToPb(st *GetPublishedDashboardTokenInfoRequest) (*getPublishedDashboardTokenInfoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedDashboardTokenInfoRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	externalValuePb, err := identity(&st.ExternalValue)
	if err != nil {
		return nil, err
	}
	if externalValuePb != nil {
		pb.ExternalValue = *externalValuePb
	}

	externalViewerIdPb, err := identity(&st.ExternalViewerId)
	if err != nil {
		return nil, err
	}
	if externalViewerIdPb != nil {
		pb.ExternalViewerId = *externalViewerIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getPublishedDashboardTokenInfoRequestPb struct {
	// UUID identifying the published dashboard.
	DashboardId string `json:"-" url:"-"`
	// Provided external value to be included in the custom claim.
	ExternalValue string `json:"-" url:"external_value,omitempty"`
	// Provided external viewer id to be included in the custom claim.
	ExternalViewerId string `json:"-" url:"external_viewer_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPublishedDashboardTokenInfoRequestFromPb(pb *getPublishedDashboardTokenInfoRequestPb) (*GetPublishedDashboardTokenInfoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedDashboardTokenInfoRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	externalValueField, err := identity(&pb.ExternalValue)
	if err != nil {
		return nil, err
	}
	if externalValueField != nil {
		st.ExternalValue = *externalValueField
	}
	externalViewerIdField, err := identity(&pb.ExternalViewerId)
	if err != nil {
		return nil, err
	}
	if externalViewerIdField != nil {
		st.ExternalViewerId = *externalViewerIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPublishedDashboardTokenInfoRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPublishedDashboardTokenInfoRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetPublishedDashboardTokenInfoResponse struct {
	// Authorization constraints for accessing the published dashboard.
	// Currently includes `workspace_rule_set` and could be enriched with
	// `unity_catalog_privileges` before oAuth token generation.
	AuthorizationDetails []AuthorizationDetails
	// Custom claim generated from external_value and external_viewer_id.
	// Format:
	// `urn:aibi:external_data:<external_value>:<external_viewer_id>:<dashboard_id>`
	CustomClaim string
	// Scope defining access permissions.
	Scope string

	ForceSendFields []string
}

func getPublishedDashboardTokenInfoResponseToPb(st *GetPublishedDashboardTokenInfoResponse) (*getPublishedDashboardTokenInfoResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedDashboardTokenInfoResponsePb{}

	var authorizationDetailsPb []authorizationDetailsPb
	for _, item := range st.AuthorizationDetails {
		itemPb, err := authorizationDetailsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			authorizationDetailsPb = append(authorizationDetailsPb, *itemPb)
		}
	}
	pb.AuthorizationDetails = authorizationDetailsPb

	customClaimPb, err := identity(&st.CustomClaim)
	if err != nil {
		return nil, err
	}
	if customClaimPb != nil {
		pb.CustomClaim = *customClaimPb
	}

	scopePb, err := identity(&st.Scope)
	if err != nil {
		return nil, err
	}
	if scopePb != nil {
		pb.Scope = *scopePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getPublishedDashboardTokenInfoResponsePb struct {
	// Authorization constraints for accessing the published dashboard.
	// Currently includes `workspace_rule_set` and could be enriched with
	// `unity_catalog_privileges` before oAuth token generation.
	AuthorizationDetails []authorizationDetailsPb `json:"authorization_details,omitempty"`
	// Custom claim generated from external_value and external_viewer_id.
	// Format:
	// `urn:aibi:external_data:<external_value>:<external_viewer_id>:<dashboard_id>`
	CustomClaim string `json:"custom_claim,omitempty"`
	// Scope defining access permissions.
	Scope string `json:"scope,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPublishedDashboardTokenInfoResponseFromPb(pb *getPublishedDashboardTokenInfoResponsePb) (*GetPublishedDashboardTokenInfoResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedDashboardTokenInfoResponse{}

	var authorizationDetailsField []AuthorizationDetails
	for _, itemPb := range pb.AuthorizationDetails {
		item, err := authorizationDetailsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			authorizationDetailsField = append(authorizationDetailsField, *item)
		}
	}
	st.AuthorizationDetails = authorizationDetailsField
	customClaimField, err := identity(&pb.CustomClaim)
	if err != nil {
		return nil, err
	}
	if customClaimField != nil {
		st.CustomClaim = *customClaimField
	}
	scopeField, err := identity(&pb.Scope)
	if err != nil {
		return nil, err
	}
	if scopeField != nil {
		st.Scope = *scopeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPublishedDashboardTokenInfoResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPublishedDashboardTokenInfoResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get dashboard schedule
type GetScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string
	// UUID identifying the schedule.
	ScheduleId string
}

func getScheduleRequestToPb(st *GetScheduleRequest) (*getScheduleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getScheduleRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	scheduleIdPb, err := identity(&st.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdPb != nil {
		pb.ScheduleId = *scheduleIdPb
	}

	return pb, nil
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

type getScheduleRequestPb struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" url:"-"`
}

func getScheduleRequestFromPb(pb *getScheduleRequestPb) (*GetScheduleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetScheduleRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	scheduleIdField, err := identity(&pb.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdField != nil {
		st.ScheduleId = *scheduleIdField
	}

	return st, nil
}

// Get schedule subscription
type GetSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId string
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId string
	// UUID identifying the subscription.
	SubscriptionId string
}

func getSubscriptionRequestToPb(st *GetSubscriptionRequest) (*getSubscriptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSubscriptionRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	scheduleIdPb, err := identity(&st.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdPb != nil {
		pb.ScheduleId = *scheduleIdPb
	}

	subscriptionIdPb, err := identity(&st.SubscriptionId)
	if err != nil {
		return nil, err
	}
	if subscriptionIdPb != nil {
		pb.SubscriptionId = *subscriptionIdPb
	}

	return pb, nil
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

type getSubscriptionRequestPb struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId string `json:"-" url:"-"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId string `json:"-" url:"-"`
	// UUID identifying the subscription.
	SubscriptionId string `json:"-" url:"-"`
}

func getSubscriptionRequestFromPb(pb *getSubscriptionRequestPb) (*GetSubscriptionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSubscriptionRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	scheduleIdField, err := identity(&pb.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdField != nil {
		st.ScheduleId = *scheduleIdField
	}
	subscriptionIdField, err := identity(&pb.SubscriptionId)
	if err != nil {
		return nil, err
	}
	if subscriptionIdField != nil {
		st.SubscriptionId = *subscriptionIdField
	}

	return st, nil
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
	PageSize int
	// A page token, received from a previous `ListDashboards` call. This token
	// can be used to retrieve the subsequent page.
	PageToken string
	// The flag to include dashboards located in the trash. If unspecified, only
	// active dashboards will be returned.
	ShowTrashed bool
	// `DASHBOARD_VIEW_BASIC`only includes summary metadata from the dashboard.
	View DashboardView

	ForceSendFields []string
}

func listDashboardsRequestToPb(st *ListDashboardsRequest) (*listDashboardsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDashboardsRequestPb{}
	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	showTrashedPb, err := identity(&st.ShowTrashed)
	if err != nil {
		return nil, err
	}
	if showTrashedPb != nil {
		pb.ShowTrashed = *showTrashedPb
	}

	viewPb, err := identity(&st.View)
	if err != nil {
		return nil, err
	}
	if viewPb != nil {
		pb.View = *viewPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listDashboardsRequestPb struct {
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

func listDashboardsRequestFromPb(pb *listDashboardsRequestPb) (*ListDashboardsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDashboardsRequest{}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	showTrashedField, err := identity(&pb.ShowTrashed)
	if err != nil {
		return nil, err
	}
	if showTrashedField != nil {
		st.ShowTrashed = *showTrashedField
	}
	viewField, err := identity(&pb.View)
	if err != nil {
		return nil, err
	}
	if viewField != nil {
		st.View = *viewField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listDashboardsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDashboardsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListDashboardsResponse struct {
	Dashboards []Dashboard
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	NextPageToken string

	ForceSendFields []string
}

func listDashboardsResponseToPb(st *ListDashboardsResponse) (*listDashboardsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDashboardsResponsePb{}

	var dashboardsPb []dashboardPb
	for _, item := range st.Dashboards {
		itemPb, err := dashboardToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dashboardsPb = append(dashboardsPb, *itemPb)
		}
	}
	pb.Dashboards = dashboardsPb

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listDashboardsResponsePb struct {
	Dashboards []dashboardPb `json:"dashboards,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listDashboardsResponseFromPb(pb *listDashboardsResponsePb) (*ListDashboardsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDashboardsResponse{}

	var dashboardsField []Dashboard
	for _, itemPb := range pb.Dashboards {
		item, err := dashboardFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dashboardsField = append(dashboardsField, *item)
		}
	}
	st.Dashboards = dashboardsField
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listDashboardsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDashboardsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List dashboard schedules
type ListSchedulesRequest struct {
	// UUID identifying the dashboard to which the schedules belongs.
	DashboardId string
	// The number of schedules to return per page.
	PageSize int
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	PageToken string

	ForceSendFields []string
}

func listSchedulesRequestToPb(st *ListSchedulesRequest) (*listSchedulesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSchedulesRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listSchedulesRequestPb struct {
	// UUID identifying the dashboard to which the schedules belongs.
	DashboardId string `json:"-" url:"-"`
	// The number of schedules to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSchedulesRequestFromPb(pb *listSchedulesRequestPb) (*ListSchedulesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSchedulesRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSchedulesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSchedulesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSchedulesResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent schedules.
	NextPageToken string

	Schedules []Schedule

	ForceSendFields []string
}

func listSchedulesResponseToPb(st *ListSchedulesResponse) (*listSchedulesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSchedulesResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	var schedulesPb []schedulePb
	for _, item := range st.Schedules {
		itemPb, err := scheduleToPb(&item)
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

type listSchedulesResponsePb struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent schedules.
	NextPageToken string `json:"next_page_token,omitempty"`

	Schedules []schedulePb `json:"schedules,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSchedulesResponseFromPb(pb *listSchedulesResponsePb) (*ListSchedulesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSchedulesResponse{}
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var schedulesField []Schedule
	for _, itemPb := range pb.Schedules {
		item, err := scheduleFromPb(&itemPb)
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

func (st *listSchedulesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSchedulesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List schedule subscriptions
type ListSubscriptionsRequest struct {
	// UUID identifying the dashboard which the subscriptions belongs.
	DashboardId string
	// The number of subscriptions to return per page.
	PageSize int
	// A page token, received from a previous `ListSubscriptions` call. Use this
	// to retrieve the subsequent page.
	PageToken string
	// UUID identifying the schedule which the subscriptions belongs.
	ScheduleId string

	ForceSendFields []string
}

func listSubscriptionsRequestToPb(st *ListSubscriptionsRequest) (*listSubscriptionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSubscriptionsRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	scheduleIdPb, err := identity(&st.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdPb != nil {
		pb.ScheduleId = *scheduleIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listSubscriptionsRequestPb struct {
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

func listSubscriptionsRequestFromPb(pb *listSubscriptionsRequestPb) (*ListSubscriptionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSubscriptionsRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	scheduleIdField, err := identity(&pb.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdField != nil {
		st.ScheduleId = *scheduleIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSubscriptionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSubscriptionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSubscriptionsResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent subscriptions.
	NextPageToken string

	Subscriptions []Subscription

	ForceSendFields []string
}

func listSubscriptionsResponseToPb(st *ListSubscriptionsResponse) (*listSubscriptionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSubscriptionsResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	var subscriptionsPb []subscriptionPb
	for _, item := range st.Subscriptions {
		itemPb, err := subscriptionToPb(&item)
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

type listSubscriptionsResponsePb struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent subscriptions.
	NextPageToken string `json:"next_page_token,omitempty"`

	Subscriptions []subscriptionPb `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSubscriptionsResponseFromPb(pb *listSubscriptionsResponsePb) (*ListSubscriptionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSubscriptionsResponse{}
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var subscriptionsField []Subscription
	for _, itemPb := range pb.Subscriptions {
		item, err := subscriptionFromPb(&itemPb)
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

func (st *listSubscriptionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSubscriptionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MessageError struct {
	Error string

	Type MessageErrorType

	ForceSendFields []string
}

func messageErrorToPb(st *MessageError) (*messageErrorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &messageErrorPb{}
	errorPb, err := identity(&st.Error)
	if err != nil {
		return nil, err
	}
	if errorPb != nil {
		pb.Error = *errorPb
	}

	typePb, err := identity(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type messageErrorPb struct {
	Error string `json:"error,omitempty"`

	Type MessageErrorType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func messageErrorFromPb(pb *messageErrorPb) (*MessageError, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MessageError{}
	errorField, err := identity(&pb.Error)
	if err != nil {
		return nil, err
	}
	if errorField != nil {
		st.Error = *errorField
	}
	typeField, err := identity(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *messageErrorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st messageErrorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

const MessageErrorTypeInvalidChatCompletionJsonException MessageErrorType = `INVALID_CHAT_COMPLETION_JSON_EXCEPTION`

const MessageErrorTypeInvalidCompletionRequestException MessageErrorType = `INVALID_COMPLETION_REQUEST_EXCEPTION`

const MessageErrorTypeInvalidFunctionCallException MessageErrorType = `INVALID_FUNCTION_CALL_EXCEPTION`

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
	case `BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION`, `CHAT_COMPLETION_CLIENT_EXCEPTION`, `CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION`, `CHAT_COMPLETION_NETWORK_EXCEPTION`, `CONTENT_FILTER_EXCEPTION`, `CONTEXT_EXCEEDED_EXCEPTION`, `COULD_NOT_GET_MODEL_DEPLOYMENTS_EXCEPTION`, `COULD_NOT_GET_UC_SCHEMA_EXCEPTION`, `DEPLOYMENT_NOT_FOUND_EXCEPTION`, `FUNCTIONS_NOT_AVAILABLE_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_TYPE_EXCEPTION`, `FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION`, `GENERATED_SQL_QUERY_TOO_LONG_EXCEPTION`, `GENERIC_CHAT_COMPLETION_EXCEPTION`, `GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION`, `GENERIC_SQL_EXEC_API_CALL_EXCEPTION`, `ILLEGAL_PARAMETER_DEFINITION_EXCEPTION`, `INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION`, `INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION`, `INVALID_CHAT_COMPLETION_JSON_EXCEPTION`, `INVALID_COMPLETION_REQUEST_EXCEPTION`, `INVALID_FUNCTION_CALL_EXCEPTION`, `INVALID_TABLE_IDENTIFIER_EXCEPTION`, `LOCAL_CONTEXT_EXCEEDED_EXCEPTION`, `MESSAGE_CANCELLED_WHILE_EXECUTING_EXCEPTION`, `MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION`, `MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION`, `MISSING_SQL_QUERY_EXCEPTION`, `NO_DEPLOYMENTS_AVAILABLE_TO_WORKSPACE`, `NO_QUERY_TO_VISUALIZE_EXCEPTION`, `NO_TABLES_TO_QUERY_EXCEPTION`, `RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION`, `RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION`, `REPLY_PROCESS_TIMEOUT_EXCEPTION`, `RETRYABLE_PROCESSING_EXCEPTION`, `SQL_EXECUTION_EXCEPTION`, `STOP_PROCESS_DUE_TO_AUTO_REGENERATE`, `TABLES_MISSING_EXCEPTION`, `TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION`, `TOO_MANY_TABLES_EXCEPTION`, `UNEXPECTED_REPLY_PROCESS_EXCEPTION`, `UNKNOWN_AI_MODEL`, `WAREHOUSE_ACCESS_MISSING_EXCEPTION`, `WAREHOUSE_NOT_FOUND_EXCEPTION`:
		*f = MessageErrorType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION", "CHAT_COMPLETION_CLIENT_EXCEPTION", "CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION", "CHAT_COMPLETION_NETWORK_EXCEPTION", "CONTENT_FILTER_EXCEPTION", "CONTEXT_EXCEEDED_EXCEPTION", "COULD_NOT_GET_MODEL_DEPLOYMENTS_EXCEPTION", "COULD_NOT_GET_UC_SCHEMA_EXCEPTION", "DEPLOYMENT_NOT_FOUND_EXCEPTION", "FUNCTIONS_NOT_AVAILABLE_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_TYPE_EXCEPTION", "FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION", "GENERATED_SQL_QUERY_TOO_LONG_EXCEPTION", "GENERIC_CHAT_COMPLETION_EXCEPTION", "GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION", "GENERIC_SQL_EXEC_API_CALL_EXCEPTION", "ILLEGAL_PARAMETER_DEFINITION_EXCEPTION", "INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION", "INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION", "INVALID_CHAT_COMPLETION_JSON_EXCEPTION", "INVALID_COMPLETION_REQUEST_EXCEPTION", "INVALID_FUNCTION_CALL_EXCEPTION", "INVALID_TABLE_IDENTIFIER_EXCEPTION", "LOCAL_CONTEXT_EXCEEDED_EXCEPTION", "MESSAGE_CANCELLED_WHILE_EXECUTING_EXCEPTION", "MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION", "MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION", "MISSING_SQL_QUERY_EXCEPTION", "NO_DEPLOYMENTS_AVAILABLE_TO_WORKSPACE", "NO_QUERY_TO_VISUALIZE_EXCEPTION", "NO_TABLES_TO_QUERY_EXCEPTION", "RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION", "RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION", "REPLY_PROCESS_TIMEOUT_EXCEPTION", "RETRYABLE_PROCESSING_EXCEPTION", "SQL_EXECUTION_EXCEPTION", "STOP_PROCESS_DUE_TO_AUTO_REGENERATE", "TABLES_MISSING_EXCEPTION", "TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION", "TOO_MANY_TABLES_EXCEPTION", "UNEXPECTED_REPLY_PROCESS_EXCEPTION", "UNKNOWN_AI_MODEL", "WAREHOUSE_ACCESS_MISSING_EXCEPTION", "WAREHOUSE_NOT_FOUND_EXCEPTION"`, v)
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
	DisplayName string
	// The workspace path of the folder to contain the migrated Lakeview
	// dashboard.
	ParentPath string
	// UUID of the dashboard to be migrated.
	SourceDashboardId string
	// Flag to indicate if mustache parameter syntax ({{ param }}) should be
	// auto-updated to named syntax (:param) when converting datasets in the
	// dashboard.
	UpdateParameterSyntax bool

	ForceSendFields []string
}

func migrateDashboardRequestToPb(st *MigrateDashboardRequest) (*migrateDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &migrateDashboardRequestPb{}
	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	parentPathPb, err := identity(&st.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathPb != nil {
		pb.ParentPath = *parentPathPb
	}

	sourceDashboardIdPb, err := identity(&st.SourceDashboardId)
	if err != nil {
		return nil, err
	}
	if sourceDashboardIdPb != nil {
		pb.SourceDashboardId = *sourceDashboardIdPb
	}

	updateParameterSyntaxPb, err := identity(&st.UpdateParameterSyntax)
	if err != nil {
		return nil, err
	}
	if updateParameterSyntaxPb != nil {
		pb.UpdateParameterSyntax = *updateParameterSyntaxPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type migrateDashboardRequestPb struct {
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

func migrateDashboardRequestFromPb(pb *migrateDashboardRequestPb) (*MigrateDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MigrateDashboardRequest{}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	parentPathField, err := identity(&pb.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathField != nil {
		st.ParentPath = *parentPathField
	}
	sourceDashboardIdField, err := identity(&pb.SourceDashboardId)
	if err != nil {
		return nil, err
	}
	if sourceDashboardIdField != nil {
		st.SourceDashboardId = *sourceDashboardIdField
	}
	updateParameterSyntaxField, err := identity(&pb.UpdateParameterSyntax)
	if err != nil {
		return nil, err
	}
	if updateParameterSyntaxField != nil {
		st.UpdateParameterSyntax = *updateParameterSyntaxField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *migrateDashboardRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st migrateDashboardRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PendingStatus struct {
	// The token to poll for result asynchronously Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	DataToken string
}

func pendingStatusToPb(st *PendingStatus) (*pendingStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pendingStatusPb{}
	dataTokenPb, err := identity(&st.DataToken)
	if err != nil {
		return nil, err
	}
	if dataTokenPb != nil {
		pb.DataToken = *dataTokenPb
	}

	return pb, nil
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

type pendingStatusPb struct {
	// The token to poll for result asynchronously Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	DataToken string `json:"data_token"`
}

func pendingStatusFromPb(pb *pendingStatusPb) (*PendingStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PendingStatus{}
	dataTokenField, err := identity(&pb.DataToken)
	if err != nil {
		return nil, err
	}
	if dataTokenField != nil {
		st.DataToken = *dataTokenField
	}

	return st, nil
}

// Poll the results for the a query for a published, embedded dashboard
type PollPublishedQueryStatusRequest struct {
	DashboardName string

	DashboardRevisionId string
	// Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	Tokens []string
}

func pollPublishedQueryStatusRequestToPb(st *PollPublishedQueryStatusRequest) (*pollPublishedQueryStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pollPublishedQueryStatusRequestPb{}
	dashboardNamePb, err := identity(&st.DashboardName)
	if err != nil {
		return nil, err
	}
	if dashboardNamePb != nil {
		pb.DashboardName = *dashboardNamePb
	}

	dashboardRevisionIdPb, err := identity(&st.DashboardRevisionId)
	if err != nil {
		return nil, err
	}
	if dashboardRevisionIdPb != nil {
		pb.DashboardRevisionId = *dashboardRevisionIdPb
	}

	var tokensPb []string
	for _, item := range st.Tokens {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tokensPb = append(tokensPb, *itemPb)
		}
	}
	pb.Tokens = tokensPb

	return pb, nil
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

type pollPublishedQueryStatusRequestPb struct {
	DashboardName string `json:"-" url:"dashboard_name"`

	DashboardRevisionId string `json:"-" url:"dashboard_revision_id"`
	// Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	Tokens []string `json:"-" url:"tokens,omitempty"`
}

func pollPublishedQueryStatusRequestFromPb(pb *pollPublishedQueryStatusRequestPb) (*PollPublishedQueryStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PollPublishedQueryStatusRequest{}
	dashboardNameField, err := identity(&pb.DashboardName)
	if err != nil {
		return nil, err
	}
	if dashboardNameField != nil {
		st.DashboardName = *dashboardNameField
	}
	dashboardRevisionIdField, err := identity(&pb.DashboardRevisionId)
	if err != nil {
		return nil, err
	}
	if dashboardRevisionIdField != nil {
		st.DashboardRevisionId = *dashboardRevisionIdField
	}

	var tokensField []string
	for _, itemPb := range pb.Tokens {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tokensField = append(tokensField, *item)
		}
	}
	st.Tokens = tokensField

	return st, nil
}

type PollQueryStatusResponse struct {
	Data []PollQueryStatusResponseData
}

func pollQueryStatusResponseToPb(st *PollQueryStatusResponse) (*pollQueryStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pollQueryStatusResponsePb{}

	var dataPb []pollQueryStatusResponseDataPb
	for _, item := range st.Data {
		itemPb, err := pollQueryStatusResponseDataToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dataPb = append(dataPb, *itemPb)
		}
	}
	pb.Data = dataPb

	return pb, nil
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

type pollQueryStatusResponsePb struct {
	Data []pollQueryStatusResponseDataPb `json:"data,omitempty"`
}

func pollQueryStatusResponseFromPb(pb *pollQueryStatusResponsePb) (*PollQueryStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PollQueryStatusResponse{}

	var dataField []PollQueryStatusResponseData
	for _, itemPb := range pb.Data {
		item, err := pollQueryStatusResponseDataFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dataField = append(dataField, *item)
		}
	}
	st.Data = dataField

	return st, nil
}

type PollQueryStatusResponseData struct {
	Status QueryResponseStatus
}

func pollQueryStatusResponseDataToPb(st *PollQueryStatusResponseData) (*pollQueryStatusResponseDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pollQueryStatusResponseDataPb{}
	statusPb, err := queryResponseStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	return pb, nil
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

type pollQueryStatusResponseDataPb struct {
	Status queryResponseStatusPb `json:"status"`
}

func pollQueryStatusResponseDataFromPb(pb *pollQueryStatusResponseDataPb) (*PollQueryStatusResponseData, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PollQueryStatusResponseData{}
	statusField, err := queryResponseStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	return st, nil
}

type PublishRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId string
	// Flag to indicate if the publisher's credentials should be embedded in the
	// published dashboard. These embedded credentials will be used to execute
	// the published dashboard's queries.
	EmbedCredentials bool
	// The ID of the warehouse that can be used to override the warehouse which
	// was set in the draft.
	WarehouseId string

	ForceSendFields []string
}

func publishRequestToPb(st *PublishRequest) (*publishRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &publishRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	embedCredentialsPb, err := identity(&st.EmbedCredentials)
	if err != nil {
		return nil, err
	}
	if embedCredentialsPb != nil {
		pb.EmbedCredentials = *embedCredentialsPb
	}

	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type publishRequestPb struct {
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

func publishRequestFromPb(pb *publishRequestPb) (*PublishRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	embedCredentialsField, err := identity(&pb.EmbedCredentials)
	if err != nil {
		return nil, err
	}
	if embedCredentialsField != nil {
		st.EmbedCredentials = *embedCredentialsField
	}
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *publishRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st publishRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PublishedDashboard struct {
	// The display name of the published dashboard.
	DisplayName string
	// Indicates whether credentials are embedded in the published dashboard.
	EmbedCredentials bool
	// The timestamp of when the published dashboard was last revised.
	RevisionCreateTime string
	// The warehouse ID used to run the published dashboard.
	WarehouseId string

	ForceSendFields []string
}

func publishedDashboardToPb(st *PublishedDashboard) (*publishedDashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &publishedDashboardPb{}
	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	embedCredentialsPb, err := identity(&st.EmbedCredentials)
	if err != nil {
		return nil, err
	}
	if embedCredentialsPb != nil {
		pb.EmbedCredentials = *embedCredentialsPb
	}

	revisionCreateTimePb, err := identity(&st.RevisionCreateTime)
	if err != nil {
		return nil, err
	}
	if revisionCreateTimePb != nil {
		pb.RevisionCreateTime = *revisionCreateTimePb
	}

	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type publishedDashboardPb struct {
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

func publishedDashboardFromPb(pb *publishedDashboardPb) (*PublishedDashboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishedDashboard{}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	embedCredentialsField, err := identity(&pb.EmbedCredentials)
	if err != nil {
		return nil, err
	}
	if embedCredentialsField != nil {
		st.EmbedCredentials = *embedCredentialsField
	}
	revisionCreateTimeField, err := identity(&pb.RevisionCreateTime)
	if err != nil {
		return nil, err
	}
	if revisionCreateTimeField != nil {
		st.RevisionCreateTime = *revisionCreateTimeField
	}
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *publishedDashboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st publishedDashboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryResponseStatus struct {
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	Canceled *Empty
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	Closed *Empty

	Pending *PendingStatus
	// The statement id in format(01eef5da-c56e-1f36-bafa-21906587d6ba) The
	// statement_id should be identical to data_token in SuccessStatus and
	// PendingStatus. This field is created for audit logging purpose to record
	// the statement_id of all QueryResponseStatus.
	StatementId string

	Success *SuccessStatus

	ForceSendFields []string
}

func queryResponseStatusToPb(st *QueryResponseStatus) (*queryResponseStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryResponseStatusPb{}
	canceledPb, err := emptyToPb(st.Canceled)
	if err != nil {
		return nil, err
	}
	if canceledPb != nil {
		pb.Canceled = canceledPb
	}

	closedPb, err := emptyToPb(st.Closed)
	if err != nil {
		return nil, err
	}
	if closedPb != nil {
		pb.Closed = closedPb
	}

	pendingPb, err := pendingStatusToPb(st.Pending)
	if err != nil {
		return nil, err
	}
	if pendingPb != nil {
		pb.Pending = pendingPb
	}

	statementIdPb, err := identity(&st.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdPb != nil {
		pb.StatementId = *statementIdPb
	}

	successPb, err := successStatusToPb(st.Success)
	if err != nil {
		return nil, err
	}
	if successPb != nil {
		pb.Success = successPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type queryResponseStatusPb struct {
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	Canceled *emptyPb `json:"canceled,omitempty"`
	// Represents an empty message, similar to google.protobuf.Empty, which is
	// not available in the firm right now.
	Closed *emptyPb `json:"closed,omitempty"`

	Pending *pendingStatusPb `json:"pending,omitempty"`
	// The statement id in format(01eef5da-c56e-1f36-bafa-21906587d6ba) The
	// statement_id should be identical to data_token in SuccessStatus and
	// PendingStatus. This field is created for audit logging purpose to record
	// the statement_id of all QueryResponseStatus.
	StatementId string `json:"statement_id,omitempty"`

	Success *successStatusPb `json:"success,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryResponseStatusFromPb(pb *queryResponseStatusPb) (*QueryResponseStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryResponseStatus{}
	canceledField, err := emptyFromPb(pb.Canceled)
	if err != nil {
		return nil, err
	}
	if canceledField != nil {
		st.Canceled = canceledField
	}
	closedField, err := emptyFromPb(pb.Closed)
	if err != nil {
		return nil, err
	}
	if closedField != nil {
		st.Closed = closedField
	}
	pendingField, err := pendingStatusFromPb(pb.Pending)
	if err != nil {
		return nil, err
	}
	if pendingField != nil {
		st.Pending = pendingField
	}
	statementIdField, err := identity(&pb.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdField != nil {
		st.StatementId = *statementIdField
	}
	successField, err := successStatusFromPb(pb.Success)
	if err != nil {
		return nil, err
	}
	if successField != nil {
		st.Success = successField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryResponseStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryResponseStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Result struct {
	// If result is truncated
	IsTruncated bool
	// Row count of the result
	RowCount int64
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId string

	ForceSendFields []string
}

func resultToPb(st *Result) (*resultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultPb{}
	isTruncatedPb, err := identity(&st.IsTruncated)
	if err != nil {
		return nil, err
	}
	if isTruncatedPb != nil {
		pb.IsTruncated = *isTruncatedPb
	}

	rowCountPb, err := identity(&st.RowCount)
	if err != nil {
		return nil, err
	}
	if rowCountPb != nil {
		pb.RowCount = *rowCountPb
	}

	statementIdPb, err := identity(&st.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdPb != nil {
		pb.StatementId = *statementIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type resultPb struct {
	// If result is truncated
	IsTruncated bool `json:"is_truncated,omitempty"`
	// Row count of the result
	RowCount int64 `json:"row_count,omitempty"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId string `json:"statement_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resultFromPb(pb *resultPb) (*Result, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Result{}
	isTruncatedField, err := identity(&pb.IsTruncated)
	if err != nil {
		return nil, err
	}
	if isTruncatedField != nil {
		st.IsTruncated = *isTruncatedField
	}
	rowCountField, err := identity(&pb.RowCount)
	if err != nil {
		return nil, err
	}
	if rowCountField != nil {
		st.RowCount = *rowCountField
	}
	statementIdField, err := identity(&pb.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdField != nil {
		st.StatementId = *statementIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Schedule struct {
	// A timestamp indicating when the schedule was created.
	CreateTime string
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule CronSchedule
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string
	// The display name for schedule.
	DisplayName string
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag string
	// The status indicates whether this schedule is paused or not.
	PauseStatus SchedulePauseStatus
	// UUID identifying the schedule.
	ScheduleId string
	// A timestamp indicating when the schedule was last updated.
	UpdateTime string
	// The warehouse id to run the dashboard with for the schedule.
	WarehouseId string

	ForceSendFields []string
}

func scheduleToPb(st *Schedule) (*schedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &schedulePb{}
	createTimePb, err := identity(&st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	cronSchedulePb, err := cronScheduleToPb(&st.CronSchedule)
	if err != nil {
		return nil, err
	}
	if cronSchedulePb != nil {
		pb.CronSchedule = *cronSchedulePb
	}

	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pauseStatusPb, err := identity(&st.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusPb != nil {
		pb.PauseStatus = *pauseStatusPb
	}

	scheduleIdPb, err := identity(&st.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdPb != nil {
		pb.ScheduleId = *scheduleIdPb
	}

	updateTimePb, err := identity(&st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type schedulePb struct {
	// A timestamp indicating when the schedule was created.
	CreateTime string `json:"create_time,omitempty"`
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule cronSchedulePb `json:"cron_schedule"`
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

func scheduleFromPb(pb *schedulePb) (*Schedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Schedule{}
	createTimeField, err := identity(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	cronScheduleField, err := cronScheduleFromPb(&pb.CronSchedule)
	if err != nil {
		return nil, err
	}
	if cronScheduleField != nil {
		st.CronSchedule = *cronScheduleField
	}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	pauseStatusField, err := identity(&pb.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusField != nil {
		st.PauseStatus = *pauseStatusField
	}
	scheduleIdField, err := identity(&pb.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdField != nil {
		st.ScheduleId = *scheduleIdField
	}
	updateTimeField, err := identity(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
	}
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *schedulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st schedulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	DestinationSubscriber *SubscriptionSubscriberDestination
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	UserSubscriber *SubscriptionSubscriberUser
}

func subscriberToPb(st *Subscriber) (*subscriberPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &subscriberPb{}
	destinationSubscriberPb, err := subscriptionSubscriberDestinationToPb(st.DestinationSubscriber)
	if err != nil {
		return nil, err
	}
	if destinationSubscriberPb != nil {
		pb.DestinationSubscriber = destinationSubscriberPb
	}

	userSubscriberPb, err := subscriptionSubscriberUserToPb(st.UserSubscriber)
	if err != nil {
		return nil, err
	}
	if userSubscriberPb != nil {
		pb.UserSubscriber = userSubscriberPb
	}

	return pb, nil
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

type subscriberPb struct {
	// The destination to receive the subscription email. This parameter is
	// mutually exclusive with `user_subscriber`.
	DestinationSubscriber *subscriptionSubscriberDestinationPb `json:"destination_subscriber,omitempty"`
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	UserSubscriber *subscriptionSubscriberUserPb `json:"user_subscriber,omitempty"`
}

func subscriberFromPb(pb *subscriberPb) (*Subscriber, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Subscriber{}
	destinationSubscriberField, err := subscriptionSubscriberDestinationFromPb(pb.DestinationSubscriber)
	if err != nil {
		return nil, err
	}
	if destinationSubscriberField != nil {
		st.DestinationSubscriber = destinationSubscriberField
	}
	userSubscriberField, err := subscriptionSubscriberUserFromPb(pb.UserSubscriber)
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
	CreateTime string
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	CreatedByUserId int64
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId string
	// The etag for the subscription. Must be left empty on create, can be
	// optionally provided on delete to ensure that the subscription has not
	// been deleted since the last read.
	Etag string
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId string
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber Subscriber
	// UUID identifying the subscription.
	SubscriptionId string
	// A timestamp indicating when the subscription was last updated.
	UpdateTime string

	ForceSendFields []string
}

func subscriptionToPb(st *Subscription) (*subscriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &subscriptionPb{}
	createTimePb, err := identity(&st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	createdByUserIdPb, err := identity(&st.CreatedByUserId)
	if err != nil {
		return nil, err
	}
	if createdByUserIdPb != nil {
		pb.CreatedByUserId = *createdByUserIdPb
	}

	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	scheduleIdPb, err := identity(&st.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdPb != nil {
		pb.ScheduleId = *scheduleIdPb
	}

	subscriberPb, err := subscriberToPb(&st.Subscriber)
	if err != nil {
		return nil, err
	}
	if subscriberPb != nil {
		pb.Subscriber = *subscriberPb
	}

	subscriptionIdPb, err := identity(&st.SubscriptionId)
	if err != nil {
		return nil, err
	}
	if subscriptionIdPb != nil {
		pb.SubscriptionId = *subscriptionIdPb
	}

	updateTimePb, err := identity(&st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type subscriptionPb struct {
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
	Subscriber subscriberPb `json:"subscriber"`
	// UUID identifying the subscription.
	SubscriptionId string `json:"subscription_id,omitempty"`
	// A timestamp indicating when the subscription was last updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func subscriptionFromPb(pb *subscriptionPb) (*Subscription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Subscription{}
	createTimeField, err := identity(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	createdByUserIdField, err := identity(&pb.CreatedByUserId)
	if err != nil {
		return nil, err
	}
	if createdByUserIdField != nil {
		st.CreatedByUserId = *createdByUserIdField
	}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	scheduleIdField, err := identity(&pb.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdField != nil {
		st.ScheduleId = *scheduleIdField
	}
	subscriberField, err := subscriberFromPb(&pb.Subscriber)
	if err != nil {
		return nil, err
	}
	if subscriberField != nil {
		st.Subscriber = *subscriberField
	}
	subscriptionIdField, err := identity(&pb.SubscriptionId)
	if err != nil {
		return nil, err
	}
	if subscriptionIdField != nil {
		st.SubscriptionId = *subscriptionIdField
	}
	updateTimeField, err := identity(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *subscriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st subscriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SubscriptionSubscriberDestination struct {
	// The canonical identifier of the destination to receive email
	// notification.
	DestinationId string
}

func subscriptionSubscriberDestinationToPb(st *SubscriptionSubscriberDestination) (*subscriptionSubscriberDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &subscriptionSubscriberDestinationPb{}
	destinationIdPb, err := identity(&st.DestinationId)
	if err != nil {
		return nil, err
	}
	if destinationIdPb != nil {
		pb.DestinationId = *destinationIdPb
	}

	return pb, nil
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

type subscriptionSubscriberDestinationPb struct {
	// The canonical identifier of the destination to receive email
	// notification.
	DestinationId string `json:"destination_id"`
}

func subscriptionSubscriberDestinationFromPb(pb *subscriptionSubscriberDestinationPb) (*SubscriptionSubscriberDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubscriptionSubscriberDestination{}
	destinationIdField, err := identity(&pb.DestinationId)
	if err != nil {
		return nil, err
	}
	if destinationIdField != nil {
		st.DestinationId = *destinationIdField
	}

	return st, nil
}

type SubscriptionSubscriberUser struct {
	// UserId of the subscriber.
	UserId int64
}

func subscriptionSubscriberUserToPb(st *SubscriptionSubscriberUser) (*subscriptionSubscriberUserPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &subscriptionSubscriberUserPb{}
	userIdPb, err := identity(&st.UserId)
	if err != nil {
		return nil, err
	}
	if userIdPb != nil {
		pb.UserId = *userIdPb
	}

	return pb, nil
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

type subscriptionSubscriberUserPb struct {
	// UserId of the subscriber.
	UserId int64 `json:"user_id"`
}

func subscriptionSubscriberUserFromPb(pb *subscriptionSubscriberUserPb) (*SubscriptionSubscriberUser, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubscriptionSubscriberUser{}
	userIdField, err := identity(&pb.UserId)
	if err != nil {
		return nil, err
	}
	if userIdField != nil {
		st.UserId = *userIdField
	}

	return st, nil
}

type SuccessStatus struct {
	// The token to poll for result asynchronously Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	DataToken string
	// Whether the query result is truncated (either by byte limit or row limit)
	Truncated bool

	ForceSendFields []string
}

func successStatusToPb(st *SuccessStatus) (*successStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &successStatusPb{}
	dataTokenPb, err := identity(&st.DataToken)
	if err != nil {
		return nil, err
	}
	if dataTokenPb != nil {
		pb.DataToken = *dataTokenPb
	}

	truncatedPb, err := identity(&st.Truncated)
	if err != nil {
		return nil, err
	}
	if truncatedPb != nil {
		pb.Truncated = *truncatedPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type successStatusPb struct {
	// The token to poll for result asynchronously Example:
	// EC0A..ChAB7WCEn_4Qo4vkLqEbXsxxEgh3Y2pbWw45WhoQXgZSQo9aS5q2ZvFcbvbx9CgA-PAEAQ
	DataToken string `json:"data_token"`
	// Whether the query result is truncated (either by byte limit or row limit)
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func successStatusFromPb(pb *successStatusPb) (*SuccessStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SuccessStatus{}
	dataTokenField, err := identity(&pb.DataToken)
	if err != nil {
		return nil, err
	}
	if dataTokenField != nil {
		st.DataToken = *dataTokenField
	}
	truncatedField, err := identity(&pb.Truncated)
	if err != nil {
		return nil, err
	}
	if truncatedField != nil {
		st.Truncated = *truncatedField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *successStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st successStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TextAttachment struct {
	// AI generated message
	Content string

	Id string

	ForceSendFields []string
}

func textAttachmentToPb(st *TextAttachment) (*textAttachmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &textAttachmentPb{}
	contentPb, err := identity(&st.Content)
	if err != nil {
		return nil, err
	}
	if contentPb != nil {
		pb.Content = *contentPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type textAttachmentPb struct {
	// AI generated message
	Content string `json:"content,omitempty"`

	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func textAttachmentFromPb(pb *textAttachmentPb) (*TextAttachment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TextAttachment{}
	contentField, err := identity(&pb.Content)
	if err != nil {
		return nil, err
	}
	if contentField != nil {
		st.Content = *contentField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *textAttachmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st textAttachmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Trash dashboard
type TrashDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string
}

func trashDashboardRequestToPb(st *TrashDashboardRequest) (*trashDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashDashboardRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	return pb, nil
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

type trashDashboardRequestPb struct {
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
}

func trashDashboardRequestFromPb(pb *trashDashboardRequestPb) (*TrashDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrashDashboardRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}

	return st, nil
}

type TrashDashboardResponse struct {
}

func trashDashboardResponseToPb(st *TrashDashboardResponse) (*trashDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashDashboardResponsePb{}

	return pb, nil
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

type trashDashboardResponsePb struct {
}

func trashDashboardResponseFromPb(pb *trashDashboardResponsePb) (*TrashDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrashDashboardResponse{}

	return st, nil
}

// Unpublish dashboard
type UnpublishDashboardRequest struct {
	// UUID identifying the published dashboard.
	DashboardId string
}

func unpublishDashboardRequestToPb(st *UnpublishDashboardRequest) (*unpublishDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &unpublishDashboardRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	return pb, nil
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

type unpublishDashboardRequestPb struct {
	// UUID identifying the published dashboard.
	DashboardId string `json:"-" url:"-"`
}

func unpublishDashboardRequestFromPb(pb *unpublishDashboardRequestPb) (*UnpublishDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnpublishDashboardRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}

	return st, nil
}

type UnpublishDashboardResponse struct {
}

func unpublishDashboardResponseToPb(st *UnpublishDashboardResponse) (*unpublishDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &unpublishDashboardResponsePb{}

	return pb, nil
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

type unpublishDashboardResponsePb struct {
}

func unpublishDashboardResponseFromPb(pb *unpublishDashboardResponsePb) (*UnpublishDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnpublishDashboardResponse{}

	return st, nil
}

// Update dashboard
type UpdateDashboardRequest struct {
	Dashboard Dashboard
	// UUID identifying the dashboard.
	DashboardId string
}

func updateDashboardRequestToPb(st *UpdateDashboardRequest) (*updateDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateDashboardRequestPb{}
	dashboardPb, err := dashboardToPb(&st.Dashboard)
	if err != nil {
		return nil, err
	}
	if dashboardPb != nil {
		pb.Dashboard = *dashboardPb
	}

	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	return pb, nil
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

type updateDashboardRequestPb struct {
	Dashboard dashboardPb `json:"dashboard"`
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
}

func updateDashboardRequestFromPb(pb *updateDashboardRequestPb) (*UpdateDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDashboardRequest{}
	dashboardField, err := dashboardFromPb(&pb.Dashboard)
	if err != nil {
		return nil, err
	}
	if dashboardField != nil {
		st.Dashboard = *dashboardField
	}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}

	return st, nil
}

// Update dashboard schedule
type UpdateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string

	Schedule Schedule
	// UUID identifying the schedule.
	ScheduleId string
}

func updateScheduleRequestToPb(st *UpdateScheduleRequest) (*updateScheduleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateScheduleRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	schedulePb, err := scheduleToPb(&st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = *schedulePb
	}

	scheduleIdPb, err := identity(&st.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdPb != nil {
		pb.ScheduleId = *scheduleIdPb
	}

	return pb, nil
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

type updateScheduleRequestPb struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`

	Schedule schedulePb `json:"schedule"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" url:"-"`
}

func updateScheduleRequestFromPb(pb *updateScheduleRequestPb) (*UpdateScheduleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateScheduleRequest{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	scheduleField, err := scheduleFromPb(&pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = *scheduleField
	}
	scheduleIdField, err := identity(&pb.ScheduleId)
	if err != nil {
		return nil, err
	}
	if scheduleIdField != nil {
		st.ScheduleId = *scheduleIdField
	}

	return st, nil
}

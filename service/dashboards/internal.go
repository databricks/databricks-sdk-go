// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/sql"
)

func authorizationDetailsToPb(st *AuthorizationDetails) (*authorizationDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &authorizationDetailsPb{}
	pb.GrantRules = st.GrantRules
	pb.ResourceLegacyAclPath = st.ResourceLegacyAclPath
	pb.ResourceName = st.ResourceName
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type authorizationDetailsPb struct {
	GrantRules            []AuthorizationDetailsGrantRule `json:"grant_rules,omitempty"`
	ResourceLegacyAclPath string                          `json:"resource_legacy_acl_path,omitempty"`
	ResourceName          string                          `json:"resource_name,omitempty"`
	Type                  string                          `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func authorizationDetailsFromPb(pb *authorizationDetailsPb) (*AuthorizationDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AuthorizationDetails{}
	st.GrantRules = pb.GrantRules
	st.ResourceLegacyAclPath = pb.ResourceLegacyAclPath
	st.ResourceName = pb.ResourceName
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *authorizationDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st authorizationDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func authorizationDetailsGrantRuleToPb(st *AuthorizationDetailsGrantRule) (*authorizationDetailsGrantRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &authorizationDetailsGrantRulePb{}
	pb.PermissionSet = st.PermissionSet

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type authorizationDetailsGrantRulePb struct {
	PermissionSet string `json:"permission_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func authorizationDetailsGrantRuleFromPb(pb *authorizationDetailsGrantRulePb) (*AuthorizationDetailsGrantRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AuthorizationDetailsGrantRule{}
	st.PermissionSet = pb.PermissionSet

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *authorizationDetailsGrantRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st authorizationDetailsGrantRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cancelPublishedQueryExecutionRequestToPb(st *CancelPublishedQueryExecutionRequest) (*cancelPublishedQueryExecutionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelPublishedQueryExecutionRequestPb{}
	pb.DashboardName = st.DashboardName
	pb.DashboardRevisionId = st.DashboardRevisionId
	pb.Tokens = st.Tokens

	return pb, nil
}

type cancelPublishedQueryExecutionRequestPb struct {
	DashboardName       string   `json:"-" url:"dashboard_name"`
	DashboardRevisionId string   `json:"-" url:"dashboard_revision_id"`
	Tokens              []string `json:"-" url:"tokens,omitempty"`
}

func cancelPublishedQueryExecutionRequestFromPb(pb *cancelPublishedQueryExecutionRequestPb) (*CancelPublishedQueryExecutionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelPublishedQueryExecutionRequest{}
	st.DashboardName = pb.DashboardName
	st.DashboardRevisionId = pb.DashboardRevisionId
	st.Tokens = pb.Tokens

	return st, nil
}

func cancelQueryExecutionResponseToPb(st *CancelQueryExecutionResponse) (*cancelQueryExecutionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelQueryExecutionResponsePb{}
	pb.Status = st.Status

	return pb, nil
}

type cancelQueryExecutionResponsePb struct {
	Status []CancelQueryExecutionResponseStatus `json:"status,omitempty"`
}

func cancelQueryExecutionResponseFromPb(pb *cancelQueryExecutionResponsePb) (*CancelQueryExecutionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelQueryExecutionResponse{}
	st.Status = pb.Status

	return st, nil
}

func cancelQueryExecutionResponseStatusToPb(st *CancelQueryExecutionResponseStatus) (*cancelQueryExecutionResponseStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelQueryExecutionResponseStatusPb{}
	pb.DataToken = st.DataToken
	pb.Pending = st.Pending
	pb.Success = st.Success

	return pb, nil
}

type cancelQueryExecutionResponseStatusPb struct {
	DataToken string `json:"data_token"`
	Pending   *Empty `json:"pending,omitempty"`
	Success   *Empty `json:"success,omitempty"`
}

func cancelQueryExecutionResponseStatusFromPb(pb *cancelQueryExecutionResponseStatusPb) (*CancelQueryExecutionResponseStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelQueryExecutionResponseStatus{}
	st.DataToken = pb.DataToken
	st.Pending = pb.Pending
	st.Success = pb.Success

	return st, nil
}

func createDashboardRequestToPb(st *CreateDashboardRequest) (*createDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createDashboardRequestPb{}
	pb.Dashboard = st.Dashboard

	return pb, nil
}

type createDashboardRequestPb struct {
	Dashboard Dashboard `json:"dashboard"`
}

func createDashboardRequestFromPb(pb *createDashboardRequestPb) (*CreateDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDashboardRequest{}
	st.Dashboard = pb.Dashboard

	return st, nil
}

func createScheduleRequestToPb(st *CreateScheduleRequest) (*createScheduleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createScheduleRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.Schedule = st.Schedule

	return pb, nil
}

type createScheduleRequestPb struct {
	DashboardId string   `json:"-" url:"-"`
	Schedule    Schedule `json:"schedule"`
}

func createScheduleRequestFromPb(pb *createScheduleRequestPb) (*CreateScheduleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateScheduleRequest{}
	st.DashboardId = pb.DashboardId
	st.Schedule = pb.Schedule

	return st, nil
}

func createSubscriptionRequestToPb(st *CreateSubscriptionRequest) (*createSubscriptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createSubscriptionRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.ScheduleId = st.ScheduleId
	pb.Subscription = st.Subscription

	return pb, nil
}

type createSubscriptionRequestPb struct {
	DashboardId  string       `json:"-" url:"-"`
	ScheduleId   string       `json:"-" url:"-"`
	Subscription Subscription `json:"subscription"`
}

func createSubscriptionRequestFromPb(pb *createSubscriptionRequestPb) (*CreateSubscriptionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateSubscriptionRequest{}
	st.DashboardId = pb.DashboardId
	st.ScheduleId = pb.ScheduleId
	st.Subscription = pb.Subscription

	return st, nil
}

func cronScheduleToPb(st *CronSchedule) (*cronSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cronSchedulePb{}
	pb.QuartzCronExpression = st.QuartzCronExpression
	pb.TimezoneId = st.TimezoneId

	return pb, nil
}

type cronSchedulePb struct {
	QuartzCronExpression string `json:"quartz_cron_expression"`
	TimezoneId           string `json:"timezone_id"`
}

func cronScheduleFromPb(pb *cronSchedulePb) (*CronSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CronSchedule{}
	st.QuartzCronExpression = pb.QuartzCronExpression
	st.TimezoneId = pb.TimezoneId

	return st, nil
}

func dashboardToPb(st *Dashboard) (*dashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardPb{}
	pb.CreateTime = st.CreateTime
	pb.DashboardId = st.DashboardId
	pb.DisplayName = st.DisplayName
	pb.Etag = st.Etag
	pb.LifecycleState = st.LifecycleState
	pb.ParentPath = st.ParentPath
	pb.Path = st.Path
	pb.SerializedDashboard = st.SerializedDashboard
	pb.UpdateTime = st.UpdateTime
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dashboardPb struct {
	CreateTime          string         `json:"create_time,omitempty"`
	DashboardId         string         `json:"dashboard_id,omitempty"`
	DisplayName         string         `json:"display_name,omitempty"`
	Etag                string         `json:"etag,omitempty"`
	LifecycleState      LifecycleState `json:"lifecycle_state,omitempty"`
	ParentPath          string         `json:"parent_path,omitempty"`
	Path                string         `json:"path,omitempty"`
	SerializedDashboard string         `json:"serialized_dashboard,omitempty"`
	UpdateTime          string         `json:"update_time,omitempty"`
	WarehouseId         string         `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dashboardFromPb(pb *dashboardPb) (*Dashboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Dashboard{}
	st.CreateTime = pb.CreateTime
	st.DashboardId = pb.DashboardId
	st.DisplayName = pb.DisplayName
	st.Etag = pb.Etag
	st.LifecycleState = pb.LifecycleState
	st.ParentPath = pb.ParentPath
	st.Path = pb.Path
	st.SerializedDashboard = pb.SerializedDashboard
	st.UpdateTime = pb.UpdateTime
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dashboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dashboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteScheduleRequestToPb(st *DeleteScheduleRequest) (*deleteScheduleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteScheduleRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.Etag = st.Etag
	pb.ScheduleId = st.ScheduleId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteScheduleRequestPb struct {
	DashboardId string `json:"-" url:"-"`
	Etag        string `json:"-" url:"etag,omitempty"`
	ScheduleId  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteScheduleRequestFromPb(pb *deleteScheduleRequestPb) (*DeleteScheduleRequest, error) {
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

func (st *deleteScheduleRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteScheduleRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteScheduleResponseToPb(st *DeleteScheduleResponse) (*deleteScheduleResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteScheduleResponsePb{}

	return pb, nil
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

func deleteSubscriptionRequestToPb(st *DeleteSubscriptionRequest) (*deleteSubscriptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSubscriptionRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.Etag = st.Etag
	pb.ScheduleId = st.ScheduleId
	pb.SubscriptionId = st.SubscriptionId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteSubscriptionRequestPb struct {
	DashboardId    string `json:"-" url:"-"`
	Etag           string `json:"-" url:"etag,omitempty"`
	ScheduleId     string `json:"-" url:"-"`
	SubscriptionId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteSubscriptionRequestFromPb(pb *deleteSubscriptionRequestPb) (*DeleteSubscriptionRequest, error) {
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

func (st *deleteSubscriptionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteSubscriptionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteSubscriptionResponseToPb(st *DeleteSubscriptionResponse) (*deleteSubscriptionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSubscriptionResponsePb{}

	return pb, nil
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

func emptyToPb(st *Empty) (*emptyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &emptyPb{}

	return pb, nil
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

func executePublishedDashboardQueryRequestToPb(st *ExecutePublishedDashboardQueryRequest) (*executePublishedDashboardQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &executePublishedDashboardQueryRequestPb{}
	pb.DashboardName = st.DashboardName
	pb.DashboardRevisionId = st.DashboardRevisionId
	pb.OverrideWarehouseId = st.OverrideWarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type executePublishedDashboardQueryRequestPb struct {
	DashboardName       string `json:"dashboard_name"`
	DashboardRevisionId string `json:"dashboard_revision_id"`
	OverrideWarehouseId string `json:"override_warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func executePublishedDashboardQueryRequestFromPb(pb *executePublishedDashboardQueryRequestPb) (*ExecutePublishedDashboardQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExecutePublishedDashboardQueryRequest{}
	st.DashboardName = pb.DashboardName
	st.DashboardRevisionId = pb.DashboardRevisionId
	st.OverrideWarehouseId = pb.OverrideWarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *executePublishedDashboardQueryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st executePublishedDashboardQueryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func executeQueryResponseToPb(st *ExecuteQueryResponse) (*executeQueryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &executeQueryResponsePb{}

	return pb, nil
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

func genieAttachmentToPb(st *GenieAttachment) (*genieAttachmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieAttachmentPb{}
	pb.AttachmentId = st.AttachmentId
	pb.Query = st.Query
	pb.Text = st.Text

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type genieAttachmentPb struct {
	AttachmentId string                `json:"attachment_id,omitempty"`
	Query        *GenieQueryAttachment `json:"query,omitempty"`
	Text         *TextAttachment       `json:"text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieAttachmentFromPb(pb *genieAttachmentPb) (*GenieAttachment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieAttachment{}
	st.AttachmentId = pb.AttachmentId
	st.Query = pb.Query
	st.Text = pb.Text

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genieAttachmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieAttachmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func genieConversationToPb(st *GenieConversation) (*genieConversationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieConversationPb{}
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

type genieConversationPb struct {
	ConversationId       string `json:"conversation_id"`
	CreatedTimestamp     int64  `json:"created_timestamp,omitempty"`
	Id                   string `json:"id"`
	LastUpdatedTimestamp int64  `json:"last_updated_timestamp,omitempty"`
	SpaceId              string `json:"space_id"`
	Title                string `json:"title"`
	UserId               int    `json:"user_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieConversationFromPb(pb *genieConversationPb) (*GenieConversation, error) {
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

func (st *genieConversationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieConversationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func genieCreateConversationMessageRequestToPb(st *GenieCreateConversationMessageRequest) (*genieCreateConversationMessageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieCreateConversationMessageRequestPb{}
	pb.Content = st.Content
	pb.ConversationId = st.ConversationId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

type genieCreateConversationMessageRequestPb struct {
	Content        string `json:"content"`
	ConversationId string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

func genieCreateConversationMessageRequestFromPb(pb *genieCreateConversationMessageRequestPb) (*GenieCreateConversationMessageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieCreateConversationMessageRequest{}
	st.Content = pb.Content
	st.ConversationId = pb.ConversationId
	st.SpaceId = pb.SpaceId

	return st, nil
}

func genieExecuteMessageAttachmentQueryRequestToPb(st *GenieExecuteMessageAttachmentQueryRequest) (*genieExecuteMessageAttachmentQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieExecuteMessageAttachmentQueryRequestPb{}
	pb.AttachmentId = st.AttachmentId
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

type genieExecuteMessageAttachmentQueryRequestPb struct {
	AttachmentId   string `json:"-" url:"-"`
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

func genieExecuteMessageAttachmentQueryRequestFromPb(pb *genieExecuteMessageAttachmentQueryRequestPb) (*GenieExecuteMessageAttachmentQueryRequest, error) {
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

func genieExecuteMessageQueryRequestToPb(st *GenieExecuteMessageQueryRequest) (*genieExecuteMessageQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieExecuteMessageQueryRequestPb{}
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

type genieExecuteMessageQueryRequestPb struct {
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

func genieExecuteMessageQueryRequestFromPb(pb *genieExecuteMessageQueryRequestPb) (*GenieExecuteMessageQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieExecuteMessageQueryRequest{}
	st.ConversationId = pb.ConversationId
	st.MessageId = pb.MessageId
	st.SpaceId = pb.SpaceId

	return st, nil
}

func genieGenerateDownloadFullQueryResultRequestToPb(st *GenieGenerateDownloadFullQueryResultRequest) (*genieGenerateDownloadFullQueryResultRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGenerateDownloadFullQueryResultRequestPb{}
	pb.AttachmentId = st.AttachmentId
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

type genieGenerateDownloadFullQueryResultRequestPb struct {
	AttachmentId   string `json:"-" url:"-"`
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

func genieGenerateDownloadFullQueryResultRequestFromPb(pb *genieGenerateDownloadFullQueryResultRequestPb) (*GenieGenerateDownloadFullQueryResultRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGenerateDownloadFullQueryResultRequest{}
	st.AttachmentId = pb.AttachmentId
	st.ConversationId = pb.ConversationId
	st.MessageId = pb.MessageId
	st.SpaceId = pb.SpaceId

	return st, nil
}

func genieGenerateDownloadFullQueryResultResponseToPb(st *GenieGenerateDownloadFullQueryResultResponse) (*genieGenerateDownloadFullQueryResultResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGenerateDownloadFullQueryResultResponsePb{}
	pb.DownloadId = st.DownloadId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type genieGenerateDownloadFullQueryResultResponsePb struct {
	DownloadId string `json:"download_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieGenerateDownloadFullQueryResultResponseFromPb(pb *genieGenerateDownloadFullQueryResultResponsePb) (*GenieGenerateDownloadFullQueryResultResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGenerateDownloadFullQueryResultResponse{}
	st.DownloadId = pb.DownloadId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genieGenerateDownloadFullQueryResultResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieGenerateDownloadFullQueryResultResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func genieGetConversationMessageRequestToPb(st *GenieGetConversationMessageRequest) (*genieGetConversationMessageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetConversationMessageRequestPb{}
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

type genieGetConversationMessageRequestPb struct {
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

func genieGetConversationMessageRequestFromPb(pb *genieGetConversationMessageRequestPb) (*GenieGetConversationMessageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetConversationMessageRequest{}
	st.ConversationId = pb.ConversationId
	st.MessageId = pb.MessageId
	st.SpaceId = pb.SpaceId

	return st, nil
}

func genieGetDownloadFullQueryResultRequestToPb(st *GenieGetDownloadFullQueryResultRequest) (*genieGetDownloadFullQueryResultRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetDownloadFullQueryResultRequestPb{}
	pb.AttachmentId = st.AttachmentId
	pb.ConversationId = st.ConversationId
	pb.DownloadId = st.DownloadId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

type genieGetDownloadFullQueryResultRequestPb struct {
	AttachmentId   string `json:"-" url:"-"`
	ConversationId string `json:"-" url:"-"`
	DownloadId     string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

func genieGetDownloadFullQueryResultRequestFromPb(pb *genieGetDownloadFullQueryResultRequestPb) (*GenieGetDownloadFullQueryResultRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetDownloadFullQueryResultRequest{}
	st.AttachmentId = pb.AttachmentId
	st.ConversationId = pb.ConversationId
	st.DownloadId = pb.DownloadId
	st.MessageId = pb.MessageId
	st.SpaceId = pb.SpaceId

	return st, nil
}

func genieGetDownloadFullQueryResultResponseToPb(st *GenieGetDownloadFullQueryResultResponse) (*genieGetDownloadFullQueryResultResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetDownloadFullQueryResultResponsePb{}
	pb.StatementResponse = st.StatementResponse

	return pb, nil
}

type genieGetDownloadFullQueryResultResponsePb struct {
	StatementResponse *sql.StatementResponse `json:"statement_response,omitempty"`
}

func genieGetDownloadFullQueryResultResponseFromPb(pb *genieGetDownloadFullQueryResultResponsePb) (*GenieGetDownloadFullQueryResultResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetDownloadFullQueryResultResponse{}
	st.StatementResponse = pb.StatementResponse

	return st, nil
}

func genieGetMessageAttachmentQueryResultRequestToPb(st *GenieGetMessageAttachmentQueryResultRequest) (*genieGetMessageAttachmentQueryResultRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetMessageAttachmentQueryResultRequestPb{}
	pb.AttachmentId = st.AttachmentId
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

type genieGetMessageAttachmentQueryResultRequestPb struct {
	AttachmentId   string `json:"-" url:"-"`
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

func genieGetMessageAttachmentQueryResultRequestFromPb(pb *genieGetMessageAttachmentQueryResultRequestPb) (*GenieGetMessageAttachmentQueryResultRequest, error) {
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

func genieGetMessageQueryResultRequestToPb(st *GenieGetMessageQueryResultRequest) (*genieGetMessageQueryResultRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetMessageQueryResultRequestPb{}
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

type genieGetMessageQueryResultRequestPb struct {
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

func genieGetMessageQueryResultRequestFromPb(pb *genieGetMessageQueryResultRequestPb) (*GenieGetMessageQueryResultRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetMessageQueryResultRequest{}
	st.ConversationId = pb.ConversationId
	st.MessageId = pb.MessageId
	st.SpaceId = pb.SpaceId

	return st, nil
}

func genieGetMessageQueryResultResponseToPb(st *GenieGetMessageQueryResultResponse) (*genieGetMessageQueryResultResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetMessageQueryResultResponsePb{}
	pb.StatementResponse = st.StatementResponse

	return pb, nil
}

type genieGetMessageQueryResultResponsePb struct {
	StatementResponse *sql.StatementResponse `json:"statement_response,omitempty"`
}

func genieGetMessageQueryResultResponseFromPb(pb *genieGetMessageQueryResultResponsePb) (*GenieGetMessageQueryResultResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetMessageQueryResultResponse{}
	st.StatementResponse = pb.StatementResponse

	return st, nil
}

func genieGetQueryResultByAttachmentRequestToPb(st *GenieGetQueryResultByAttachmentRequest) (*genieGetQueryResultByAttachmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetQueryResultByAttachmentRequestPb{}
	pb.AttachmentId = st.AttachmentId
	pb.ConversationId = st.ConversationId
	pb.MessageId = st.MessageId
	pb.SpaceId = st.SpaceId

	return pb, nil
}

type genieGetQueryResultByAttachmentRequestPb struct {
	AttachmentId   string `json:"-" url:"-"`
	ConversationId string `json:"-" url:"-"`
	MessageId      string `json:"-" url:"-"`
	SpaceId        string `json:"-" url:"-"`
}

func genieGetQueryResultByAttachmentRequestFromPb(pb *genieGetQueryResultByAttachmentRequestPb) (*GenieGetQueryResultByAttachmentRequest, error) {
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

func genieGetSpaceRequestToPb(st *GenieGetSpaceRequest) (*genieGetSpaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieGetSpaceRequestPb{}
	pb.SpaceId = st.SpaceId

	return pb, nil
}

type genieGetSpaceRequestPb struct {
	SpaceId string `json:"-" url:"-"`
}

func genieGetSpaceRequestFromPb(pb *genieGetSpaceRequestPb) (*GenieGetSpaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieGetSpaceRequest{}
	st.SpaceId = pb.SpaceId

	return st, nil
}

func genieMessageToPb(st *GenieMessage) (*genieMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieMessagePb{}
	pb.Attachments = st.Attachments
	pb.Content = st.Content
	pb.ConversationId = st.ConversationId
	pb.CreatedTimestamp = st.CreatedTimestamp
	pb.Error = st.Error
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.MessageId = st.MessageId
	pb.QueryResult = st.QueryResult
	pb.SpaceId = st.SpaceId
	pb.Status = st.Status
	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type genieMessagePb struct {
	Attachments          []GenieAttachment `json:"attachments,omitempty"`
	Content              string            `json:"content"`
	ConversationId       string            `json:"conversation_id"`
	CreatedTimestamp     int64             `json:"created_timestamp,omitempty"`
	Error                *MessageError     `json:"error,omitempty"`
	Id                   string            `json:"id"`
	LastUpdatedTimestamp int64             `json:"last_updated_timestamp,omitempty"`
	MessageId            string            `json:"message_id"`
	QueryResult          *Result           `json:"query_result,omitempty"`
	SpaceId              string            `json:"space_id"`
	Status               MessageStatus     `json:"status,omitempty"`
	UserId               int64             `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieMessageFromPb(pb *genieMessagePb) (*GenieMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieMessage{}
	st.Attachments = pb.Attachments
	st.Content = pb.Content
	st.ConversationId = pb.ConversationId
	st.CreatedTimestamp = pb.CreatedTimestamp
	st.Error = pb.Error
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.MessageId = pb.MessageId
	st.QueryResult = pb.QueryResult
	st.SpaceId = pb.SpaceId
	st.Status = pb.Status
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genieMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func genieQueryAttachmentToPb(st *GenieQueryAttachment) (*genieQueryAttachmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieQueryAttachmentPb{}
	pb.Description = st.Description
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.Query = st.Query
	pb.QueryResultMetadata = st.QueryResultMetadata
	pb.StatementId = st.StatementId
	pb.Title = st.Title

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type genieQueryAttachmentPb struct {
	Description          string               `json:"description,omitempty"`
	Id                   string               `json:"id,omitempty"`
	LastUpdatedTimestamp int64                `json:"last_updated_timestamp,omitempty"`
	Query                string               `json:"query,omitempty"`
	QueryResultMetadata  *GenieResultMetadata `json:"query_result_metadata,omitempty"`
	StatementId          string               `json:"statement_id,omitempty"`
	Title                string               `json:"title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieQueryAttachmentFromPb(pb *genieQueryAttachmentPb) (*GenieQueryAttachment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieQueryAttachment{}
	st.Description = pb.Description
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.Query = pb.Query
	st.QueryResultMetadata = pb.QueryResultMetadata
	st.StatementId = pb.StatementId
	st.Title = pb.Title

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genieQueryAttachmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieQueryAttachmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func genieResultMetadataToPb(st *GenieResultMetadata) (*genieResultMetadataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieResultMetadataPb{}
	pb.IsTruncated = st.IsTruncated
	pb.RowCount = st.RowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type genieResultMetadataPb struct {
	IsTruncated bool  `json:"is_truncated,omitempty"`
	RowCount    int64 `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieResultMetadataFromPb(pb *genieResultMetadataPb) (*GenieResultMetadata, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieResultMetadata{}
	st.IsTruncated = pb.IsTruncated
	st.RowCount = pb.RowCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genieResultMetadataPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieResultMetadataPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func genieSpaceToPb(st *GenieSpace) (*genieSpacePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieSpacePb{}
	pb.Description = st.Description
	pb.SpaceId = st.SpaceId
	pb.Title = st.Title

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type genieSpacePb struct {
	Description string `json:"description,omitempty"`
	SpaceId     string `json:"space_id"`
	Title       string `json:"title"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genieSpaceFromPb(pb *genieSpacePb) (*GenieSpace, error) {
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

func (st *genieSpacePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genieSpacePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func genieStartConversationMessageRequestToPb(st *GenieStartConversationMessageRequest) (*genieStartConversationMessageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieStartConversationMessageRequestPb{}
	pb.Content = st.Content
	pb.SpaceId = st.SpaceId

	return pb, nil
}

type genieStartConversationMessageRequestPb struct {
	Content string `json:"content"`
	SpaceId string `json:"-" url:"-"`
}

func genieStartConversationMessageRequestFromPb(pb *genieStartConversationMessageRequestPb) (*GenieStartConversationMessageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieStartConversationMessageRequest{}
	st.Content = pb.Content
	st.SpaceId = pb.SpaceId

	return st, nil
}

func genieStartConversationResponseToPb(st *GenieStartConversationResponse) (*genieStartConversationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genieStartConversationResponsePb{}
	pb.Conversation = st.Conversation
	pb.ConversationId = st.ConversationId
	pb.Message = st.Message
	pb.MessageId = st.MessageId

	return pb, nil
}

type genieStartConversationResponsePb struct {
	Conversation   *GenieConversation `json:"conversation,omitempty"`
	ConversationId string             `json:"conversation_id"`
	Message        *GenieMessage      `json:"message,omitempty"`
	MessageId      string             `json:"message_id"`
}

func genieStartConversationResponseFromPb(pb *genieStartConversationResponsePb) (*GenieStartConversationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenieStartConversationResponse{}
	st.Conversation = pb.Conversation
	st.ConversationId = pb.ConversationId
	st.Message = pb.Message
	st.MessageId = pb.MessageId

	return st, nil
}

func getDashboardRequestToPb(st *GetDashboardRequest) (*getDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

type getDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

func getDashboardRequestFromPb(pb *getDashboardRequestPb) (*GetDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

func getPublishedDashboardEmbeddedRequestToPb(st *GetPublishedDashboardEmbeddedRequest) (*getPublishedDashboardEmbeddedRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedDashboardEmbeddedRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

type getPublishedDashboardEmbeddedRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

func getPublishedDashboardEmbeddedRequestFromPb(pb *getPublishedDashboardEmbeddedRequestPb) (*GetPublishedDashboardEmbeddedRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedDashboardEmbeddedRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

func getPublishedDashboardEmbeddedResponseToPb(st *GetPublishedDashboardEmbeddedResponse) (*getPublishedDashboardEmbeddedResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedDashboardEmbeddedResponsePb{}

	return pb, nil
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

func getPublishedDashboardRequestToPb(st *GetPublishedDashboardRequest) (*getPublishedDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

type getPublishedDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

func getPublishedDashboardRequestFromPb(pb *getPublishedDashboardRequestPb) (*GetPublishedDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

func getPublishedDashboardTokenInfoRequestToPb(st *GetPublishedDashboardTokenInfoRequest) (*getPublishedDashboardTokenInfoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedDashboardTokenInfoRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.ExternalValue = st.ExternalValue
	pb.ExternalViewerId = st.ExternalViewerId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getPublishedDashboardTokenInfoRequestPb struct {
	DashboardId      string `json:"-" url:"-"`
	ExternalValue    string `json:"-" url:"external_value,omitempty"`
	ExternalViewerId string `json:"-" url:"external_viewer_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPublishedDashboardTokenInfoRequestFromPb(pb *getPublishedDashboardTokenInfoRequestPb) (*GetPublishedDashboardTokenInfoRequest, error) {
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

func (st *getPublishedDashboardTokenInfoRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPublishedDashboardTokenInfoRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getPublishedDashboardTokenInfoResponseToPb(st *GetPublishedDashboardTokenInfoResponse) (*getPublishedDashboardTokenInfoResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedDashboardTokenInfoResponsePb{}
	pb.AuthorizationDetails = st.AuthorizationDetails
	pb.CustomClaim = st.CustomClaim
	pb.Scope = st.Scope

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getPublishedDashboardTokenInfoResponsePb struct {
	AuthorizationDetails []AuthorizationDetails `json:"authorization_details,omitempty"`
	CustomClaim          string                 `json:"custom_claim,omitempty"`
	Scope                string                 `json:"scope,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPublishedDashboardTokenInfoResponseFromPb(pb *getPublishedDashboardTokenInfoResponsePb) (*GetPublishedDashboardTokenInfoResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedDashboardTokenInfoResponse{}
	st.AuthorizationDetails = pb.AuthorizationDetails
	st.CustomClaim = pb.CustomClaim
	st.Scope = pb.Scope

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPublishedDashboardTokenInfoResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPublishedDashboardTokenInfoResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getScheduleRequestToPb(st *GetScheduleRequest) (*getScheduleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getScheduleRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.ScheduleId = st.ScheduleId

	return pb, nil
}

type getScheduleRequestPb struct {
	DashboardId string `json:"-" url:"-"`
	ScheduleId  string `json:"-" url:"-"`
}

func getScheduleRequestFromPb(pb *getScheduleRequestPb) (*GetScheduleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetScheduleRequest{}
	st.DashboardId = pb.DashboardId
	st.ScheduleId = pb.ScheduleId

	return st, nil
}

func getSubscriptionRequestToPb(st *GetSubscriptionRequest) (*getSubscriptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSubscriptionRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.ScheduleId = st.ScheduleId
	pb.SubscriptionId = st.SubscriptionId

	return pb, nil
}

type getSubscriptionRequestPb struct {
	DashboardId    string `json:"-" url:"-"`
	ScheduleId     string `json:"-" url:"-"`
	SubscriptionId string `json:"-" url:"-"`
}

func getSubscriptionRequestFromPb(pb *getSubscriptionRequestPb) (*GetSubscriptionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSubscriptionRequest{}
	st.DashboardId = pb.DashboardId
	st.ScheduleId = pb.ScheduleId
	st.SubscriptionId = pb.SubscriptionId

	return st, nil
}

func listDashboardsRequestToPb(st *ListDashboardsRequest) (*listDashboardsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDashboardsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.ShowTrashed = st.ShowTrashed
	pb.View = st.View

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listDashboardsRequestPb struct {
	PageSize    int           `json:"-" url:"page_size,omitempty"`
	PageToken   string        `json:"-" url:"page_token,omitempty"`
	ShowTrashed bool          `json:"-" url:"show_trashed,omitempty"`
	View        DashboardView `json:"-" url:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listDashboardsRequestFromPb(pb *listDashboardsRequestPb) (*ListDashboardsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDashboardsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.ShowTrashed = pb.ShowTrashed
	st.View = pb.View

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listDashboardsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDashboardsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listDashboardsResponseToPb(st *ListDashboardsResponse) (*listDashboardsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDashboardsResponsePb{}
	pb.Dashboards = st.Dashboards
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listDashboardsResponsePb struct {
	Dashboards    []Dashboard `json:"dashboards,omitempty"`
	NextPageToken string      `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listDashboardsResponseFromPb(pb *listDashboardsResponsePb) (*ListDashboardsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDashboardsResponse{}
	st.Dashboards = pb.Dashboards
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listDashboardsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDashboardsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listSchedulesRequestToPb(st *ListSchedulesRequest) (*listSchedulesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSchedulesRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listSchedulesRequestPb struct {
	DashboardId string `json:"-" url:"-"`
	PageSize    int    `json:"-" url:"page_size,omitempty"`
	PageToken   string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSchedulesRequestFromPb(pb *listSchedulesRequestPb) (*ListSchedulesRequest, error) {
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

func (st *listSchedulesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSchedulesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listSchedulesResponseToPb(st *ListSchedulesResponse) (*listSchedulesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSchedulesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Schedules = st.Schedules

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listSchedulesResponsePb struct {
	NextPageToken string     `json:"next_page_token,omitempty"`
	Schedules     []Schedule `json:"schedules,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSchedulesResponseFromPb(pb *listSchedulesResponsePb) (*ListSchedulesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSchedulesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Schedules = pb.Schedules

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSchedulesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSchedulesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listSubscriptionsRequestToPb(st *ListSubscriptionsRequest) (*listSubscriptionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSubscriptionsRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.ScheduleId = st.ScheduleId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listSubscriptionsRequestPb struct {
	DashboardId string `json:"-" url:"-"`
	PageSize    int    `json:"-" url:"page_size,omitempty"`
	PageToken   string `json:"-" url:"page_token,omitempty"`
	ScheduleId  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSubscriptionsRequestFromPb(pb *listSubscriptionsRequestPb) (*ListSubscriptionsRequest, error) {
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

func (st *listSubscriptionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSubscriptionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listSubscriptionsResponseToPb(st *ListSubscriptionsResponse) (*listSubscriptionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSubscriptionsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Subscriptions = st.Subscriptions

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listSubscriptionsResponsePb struct {
	NextPageToken string         `json:"next_page_token,omitempty"`
	Subscriptions []Subscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSubscriptionsResponseFromPb(pb *listSubscriptionsResponsePb) (*ListSubscriptionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSubscriptionsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Subscriptions = pb.Subscriptions

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSubscriptionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSubscriptionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func messageErrorToPb(st *MessageError) (*messageErrorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &messageErrorPb{}
	pb.Error = st.Error
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type messageErrorPb struct {
	Error string           `json:"error,omitempty"`
	Type  MessageErrorType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func messageErrorFromPb(pb *messageErrorPb) (*MessageError, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MessageError{}
	st.Error = pb.Error
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *messageErrorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st messageErrorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func migrateDashboardRequestToPb(st *MigrateDashboardRequest) (*migrateDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &migrateDashboardRequestPb{}
	pb.DisplayName = st.DisplayName
	pb.ParentPath = st.ParentPath
	pb.SourceDashboardId = st.SourceDashboardId
	pb.UpdateParameterSyntax = st.UpdateParameterSyntax

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type migrateDashboardRequestPb struct {
	DisplayName           string `json:"display_name,omitempty"`
	ParentPath            string `json:"parent_path,omitempty"`
	SourceDashboardId     string `json:"source_dashboard_id"`
	UpdateParameterSyntax bool   `json:"update_parameter_syntax,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func migrateDashboardRequestFromPb(pb *migrateDashboardRequestPb) (*MigrateDashboardRequest, error) {
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

func (st *migrateDashboardRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st migrateDashboardRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pendingStatusToPb(st *PendingStatus) (*pendingStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pendingStatusPb{}
	pb.DataToken = st.DataToken

	return pb, nil
}

type pendingStatusPb struct {
	DataToken string `json:"data_token"`
}

func pendingStatusFromPb(pb *pendingStatusPb) (*PendingStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PendingStatus{}
	st.DataToken = pb.DataToken

	return st, nil
}

func pollPublishedQueryStatusRequestToPb(st *PollPublishedQueryStatusRequest) (*pollPublishedQueryStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pollPublishedQueryStatusRequestPb{}
	pb.DashboardName = st.DashboardName
	pb.DashboardRevisionId = st.DashboardRevisionId
	pb.Tokens = st.Tokens

	return pb, nil
}

type pollPublishedQueryStatusRequestPb struct {
	DashboardName       string   `json:"-" url:"dashboard_name"`
	DashboardRevisionId string   `json:"-" url:"dashboard_revision_id"`
	Tokens              []string `json:"-" url:"tokens,omitempty"`
}

func pollPublishedQueryStatusRequestFromPb(pb *pollPublishedQueryStatusRequestPb) (*PollPublishedQueryStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PollPublishedQueryStatusRequest{}
	st.DashboardName = pb.DashboardName
	st.DashboardRevisionId = pb.DashboardRevisionId
	st.Tokens = pb.Tokens

	return st, nil
}

func pollQueryStatusResponseToPb(st *PollQueryStatusResponse) (*pollQueryStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pollQueryStatusResponsePb{}
	pb.Data = st.Data

	return pb, nil
}

type pollQueryStatusResponsePb struct {
	Data []PollQueryStatusResponseData `json:"data,omitempty"`
}

func pollQueryStatusResponseFromPb(pb *pollQueryStatusResponsePb) (*PollQueryStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PollQueryStatusResponse{}
	st.Data = pb.Data

	return st, nil
}

func pollQueryStatusResponseDataToPb(st *PollQueryStatusResponseData) (*pollQueryStatusResponseDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pollQueryStatusResponseDataPb{}
	pb.Status = st.Status

	return pb, nil
}

type pollQueryStatusResponseDataPb struct {
	Status QueryResponseStatus `json:"status"`
}

func pollQueryStatusResponseDataFromPb(pb *pollQueryStatusResponseDataPb) (*PollQueryStatusResponseData, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PollQueryStatusResponseData{}
	st.Status = pb.Status

	return st, nil
}

func publishRequestToPb(st *PublishRequest) (*publishRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &publishRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.EmbedCredentials = st.EmbedCredentials
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type publishRequestPb struct {
	DashboardId      string `json:"-" url:"-"`
	EmbedCredentials bool   `json:"embed_credentials,omitempty"`
	WarehouseId      string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func publishRequestFromPb(pb *publishRequestPb) (*PublishRequest, error) {
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

func (st *publishRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st publishRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func publishedDashboardToPb(st *PublishedDashboard) (*publishedDashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &publishedDashboardPb{}
	pb.DisplayName = st.DisplayName
	pb.EmbedCredentials = st.EmbedCredentials
	pb.RevisionCreateTime = st.RevisionCreateTime
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type publishedDashboardPb struct {
	DisplayName        string `json:"display_name,omitempty"`
	EmbedCredentials   bool   `json:"embed_credentials,omitempty"`
	RevisionCreateTime string `json:"revision_create_time,omitempty"`
	WarehouseId        string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func publishedDashboardFromPb(pb *publishedDashboardPb) (*PublishedDashboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishedDashboard{}
	st.DisplayName = pb.DisplayName
	st.EmbedCredentials = pb.EmbedCredentials
	st.RevisionCreateTime = pb.RevisionCreateTime
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *publishedDashboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st publishedDashboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func queryResponseStatusToPb(st *QueryResponseStatus) (*queryResponseStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryResponseStatusPb{}
	pb.Canceled = st.Canceled
	pb.Closed = st.Closed
	pb.Pending = st.Pending
	pb.StatementId = st.StatementId
	pb.Success = st.Success

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type queryResponseStatusPb struct {
	Canceled    *Empty         `json:"canceled,omitempty"`
	Closed      *Empty         `json:"closed,omitempty"`
	Pending     *PendingStatus `json:"pending,omitempty"`
	StatementId string         `json:"statement_id,omitempty"`
	Success     *SuccessStatus `json:"success,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryResponseStatusFromPb(pb *queryResponseStatusPb) (*QueryResponseStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryResponseStatus{}
	st.Canceled = pb.Canceled
	st.Closed = pb.Closed
	st.Pending = pb.Pending
	st.StatementId = pb.StatementId
	st.Success = pb.Success

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryResponseStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryResponseStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func resultToPb(st *Result) (*resultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultPb{}
	pb.IsTruncated = st.IsTruncated
	pb.RowCount = st.RowCount
	pb.StatementId = st.StatementId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type resultPb struct {
	IsTruncated bool   `json:"is_truncated,omitempty"`
	RowCount    int64  `json:"row_count,omitempty"`
	StatementId string `json:"statement_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resultFromPb(pb *resultPb) (*Result, error) {
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

func (st *resultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func scheduleToPb(st *Schedule) (*schedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &schedulePb{}
	pb.CreateTime = st.CreateTime
	pb.CronSchedule = st.CronSchedule
	pb.DashboardId = st.DashboardId
	pb.DisplayName = st.DisplayName
	pb.Etag = st.Etag
	pb.PauseStatus = st.PauseStatus
	pb.ScheduleId = st.ScheduleId
	pb.UpdateTime = st.UpdateTime
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type schedulePb struct {
	CreateTime   string              `json:"create_time,omitempty"`
	CronSchedule CronSchedule        `json:"cron_schedule"`
	DashboardId  string              `json:"dashboard_id,omitempty"`
	DisplayName  string              `json:"display_name,omitempty"`
	Etag         string              `json:"etag,omitempty"`
	PauseStatus  SchedulePauseStatus `json:"pause_status,omitempty"`
	ScheduleId   string              `json:"schedule_id,omitempty"`
	UpdateTime   string              `json:"update_time,omitempty"`
	WarehouseId  string              `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func scheduleFromPb(pb *schedulePb) (*Schedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Schedule{}
	st.CreateTime = pb.CreateTime
	st.CronSchedule = pb.CronSchedule
	st.DashboardId = pb.DashboardId
	st.DisplayName = pb.DisplayName
	st.Etag = pb.Etag
	st.PauseStatus = pb.PauseStatus
	st.ScheduleId = pb.ScheduleId
	st.UpdateTime = pb.UpdateTime
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *schedulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st schedulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func subscriberToPb(st *Subscriber) (*subscriberPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &subscriberPb{}
	pb.DestinationSubscriber = st.DestinationSubscriber
	pb.UserSubscriber = st.UserSubscriber

	return pb, nil
}

type subscriberPb struct {
	DestinationSubscriber *SubscriptionSubscriberDestination `json:"destination_subscriber,omitempty"`
	UserSubscriber        *SubscriptionSubscriberUser        `json:"user_subscriber,omitempty"`
}

func subscriberFromPb(pb *subscriberPb) (*Subscriber, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Subscriber{}
	st.DestinationSubscriber = pb.DestinationSubscriber
	st.UserSubscriber = pb.UserSubscriber

	return st, nil
}

func subscriptionToPb(st *Subscription) (*subscriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &subscriptionPb{}
	pb.CreateTime = st.CreateTime
	pb.CreatedByUserId = st.CreatedByUserId
	pb.DashboardId = st.DashboardId
	pb.Etag = st.Etag
	pb.ScheduleId = st.ScheduleId
	pb.Subscriber = st.Subscriber
	pb.SubscriptionId = st.SubscriptionId
	pb.UpdateTime = st.UpdateTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type subscriptionPb struct {
	CreateTime      string     `json:"create_time,omitempty"`
	CreatedByUserId int64      `json:"created_by_user_id,omitempty"`
	DashboardId     string     `json:"dashboard_id,omitempty"`
	Etag            string     `json:"etag,omitempty"`
	ScheduleId      string     `json:"schedule_id,omitempty"`
	Subscriber      Subscriber `json:"subscriber"`
	SubscriptionId  string     `json:"subscription_id,omitempty"`
	UpdateTime      string     `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func subscriptionFromPb(pb *subscriptionPb) (*Subscription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Subscription{}
	st.CreateTime = pb.CreateTime
	st.CreatedByUserId = pb.CreatedByUserId
	st.DashboardId = pb.DashboardId
	st.Etag = pb.Etag
	st.ScheduleId = pb.ScheduleId
	st.Subscriber = pb.Subscriber
	st.SubscriptionId = pb.SubscriptionId
	st.UpdateTime = pb.UpdateTime

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *subscriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st subscriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func subscriptionSubscriberDestinationToPb(st *SubscriptionSubscriberDestination) (*subscriptionSubscriberDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &subscriptionSubscriberDestinationPb{}
	pb.DestinationId = st.DestinationId

	return pb, nil
}

type subscriptionSubscriberDestinationPb struct {
	DestinationId string `json:"destination_id"`
}

func subscriptionSubscriberDestinationFromPb(pb *subscriptionSubscriberDestinationPb) (*SubscriptionSubscriberDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubscriptionSubscriberDestination{}
	st.DestinationId = pb.DestinationId

	return st, nil
}

func subscriptionSubscriberUserToPb(st *SubscriptionSubscriberUser) (*subscriptionSubscriberUserPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &subscriptionSubscriberUserPb{}
	pb.UserId = st.UserId

	return pb, nil
}

type subscriptionSubscriberUserPb struct {
	UserId int64 `json:"user_id"`
}

func subscriptionSubscriberUserFromPb(pb *subscriptionSubscriberUserPb) (*SubscriptionSubscriberUser, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubscriptionSubscriberUser{}
	st.UserId = pb.UserId

	return st, nil
}

func successStatusToPb(st *SuccessStatus) (*successStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &successStatusPb{}
	pb.DataToken = st.DataToken
	pb.Truncated = st.Truncated

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type successStatusPb struct {
	DataToken string `json:"data_token"`
	Truncated bool   `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func successStatusFromPb(pb *successStatusPb) (*SuccessStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SuccessStatus{}
	st.DataToken = pb.DataToken
	st.Truncated = pb.Truncated

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *successStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st successStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func textAttachmentToPb(st *TextAttachment) (*textAttachmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &textAttachmentPb{}
	pb.Content = st.Content
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type textAttachmentPb struct {
	Content string `json:"content,omitempty"`
	Id      string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func textAttachmentFromPb(pb *textAttachmentPb) (*TextAttachment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TextAttachment{}
	st.Content = pb.Content
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *textAttachmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st textAttachmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func trashDashboardRequestToPb(st *TrashDashboardRequest) (*trashDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

type trashDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

func trashDashboardRequestFromPb(pb *trashDashboardRequestPb) (*TrashDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrashDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

func trashDashboardResponseToPb(st *TrashDashboardResponse) (*trashDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashDashboardResponsePb{}

	return pb, nil
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

func unpublishDashboardRequestToPb(st *UnpublishDashboardRequest) (*unpublishDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &unpublishDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

type unpublishDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

func unpublishDashboardRequestFromPb(pb *unpublishDashboardRequestPb) (*UnpublishDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnpublishDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

func unpublishDashboardResponseToPb(st *UnpublishDashboardResponse) (*unpublishDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &unpublishDashboardResponsePb{}

	return pb, nil
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

func updateDashboardRequestToPb(st *UpdateDashboardRequest) (*updateDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateDashboardRequestPb{}
	pb.Dashboard = st.Dashboard
	pb.DashboardId = st.DashboardId

	return pb, nil
}

type updateDashboardRequestPb struct {
	Dashboard   Dashboard `json:"dashboard"`
	DashboardId string    `json:"-" url:"-"`
}

func updateDashboardRequestFromPb(pb *updateDashboardRequestPb) (*UpdateDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDashboardRequest{}
	st.Dashboard = pb.Dashboard
	st.DashboardId = pb.DashboardId

	return st, nil
}

func updateScheduleRequestToPb(st *UpdateScheduleRequest) (*updateScheduleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateScheduleRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.Schedule = st.Schedule
	pb.ScheduleId = st.ScheduleId

	return pb, nil
}

type updateScheduleRequestPb struct {
	DashboardId string   `json:"-" url:"-"`
	Schedule    Schedule `json:"schedule"`
	ScheduleId  string   `json:"-" url:"-"`
}

func updateScheduleRequestFromPb(pb *updateScheduleRequestPb) (*UpdateScheduleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateScheduleRequest{}
	st.DashboardId = pb.DashboardId
	st.Schedule = pb.Schedule
	st.ScheduleId = pb.ScheduleId

	return st, nil
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

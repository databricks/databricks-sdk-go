// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Genie, Lakeview, Lakeview Embedded, Query Execution, etc.
package dashboards

import (
	"context"
)

type GenieAPI struct {
	genieImpl
}

// Get conversation message.
//
// Get message from conversation.
func (a *GenieAPI) GetMessageBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieMessage, error) {
	return a.genieImpl.GetMessage(ctx, GenieGetConversationMessageRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
	})
}

// Get message attachment SQL query result.
//
// Get the result of SQL query if the message has a query attachment. This is
// only available if a message has a query attachment and the message status is
// `EXECUTING_QUERY` OR `COMPLETED`.
func (a *GenieAPI) GetMessageAttachmentQueryResultBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error) {
	return a.genieImpl.GetMessageAttachmentQueryResult(ctx, GenieGetMessageAttachmentQueryResultRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
		AttachmentId:   attachmentId,
	})
}

// [Deprecated] Get conversation message SQL query result.
//
// Get the result of SQL query if the message has a query attachment. This is
// only available if a message has a query attachment and the message status is
// `EXECUTING_QUERY`.
func (a *GenieAPI) GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieGetMessageQueryResultResponse, error) {
	return a.genieImpl.GetMessageQueryResult(ctx, GenieGetMessageQueryResultRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
	})
}

// [Deprecated] Get conversation message SQL query result.
//
// Get the result of SQL query if the message has a query attachment. This is
// only available if a message has a query attachment and the message status is
// `EXECUTING_QUERY` OR `COMPLETED`.
func (a *GenieAPI) GetMessageQueryResultByAttachmentBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error) {
	return a.genieImpl.GetMessageQueryResultByAttachment(ctx, GenieGetQueryResultByAttachmentRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
		AttachmentId:   attachmentId,
	})
}

// Get Genie Space.
//
// Get details of a Genie Space.
func (a *GenieAPI) GetSpaceBySpaceId(ctx context.Context, spaceId string) (*GenieSpace, error) {
	return a.genieImpl.GetSpace(ctx, GenieGetSpaceRequest{
		SpaceId: spaceId,
	})
}

type LakeviewAPI struct {
	lakeviewImpl
}

// Delete dashboard schedule.
func (a *LakeviewAPI) DeleteScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*DeleteScheduleResponse, error) {
	return a.lakeviewImpl.DeleteSchedule(ctx, DeleteScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Delete schedule subscription.
func (a *LakeviewAPI) DeleteSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*DeleteSubscriptionResponse, error) {
	return a.lakeviewImpl.DeleteSubscription(ctx, DeleteSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// Get dashboard.
//
// Get a draft dashboard.
func (a *LakeviewAPI) GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error) {
	return a.lakeviewImpl.Get(ctx, GetDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get published dashboard.
//
// Get the current published dashboard.
func (a *LakeviewAPI) GetPublishedByDashboardId(ctx context.Context, dashboardId string) (*PublishedDashboard, error) {
	return a.lakeviewImpl.GetPublished(ctx, GetPublishedDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get dashboard schedule.
func (a *LakeviewAPI) GetScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*Schedule, error) {
	return a.lakeviewImpl.GetSchedule(ctx, GetScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Get schedule subscription.
func (a *LakeviewAPI) GetSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*Subscription, error) {
	return a.lakeviewImpl.GetSubscription(ctx, GetSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// List dashboard schedules.
func (a *LakeviewAPI) ListSchedulesByDashboardId(ctx context.Context, dashboardId string) (*ListSchedulesResponse, error) {
	return a.lakeviewImpl.internalListSchedules(ctx, ListSchedulesRequest{
		DashboardId: dashboardId,
	})
}

// List schedule subscriptions.
func (a *LakeviewAPI) ListSubscriptionsByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*ListSubscriptionsResponse, error) {
	return a.lakeviewImpl.internalListSubscriptions(ctx, ListSubscriptionsRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Trash dashboard.
//
// Trash a dashboard.
func (a *LakeviewAPI) TrashByDashboardId(ctx context.Context, dashboardId string) (*TrashDashboardResponse, error) {
	return a.lakeviewImpl.Trash(ctx, TrashDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Unpublish dashboard.
//
// Unpublish the dashboard.
func (a *LakeviewAPI) UnpublishByDashboardId(ctx context.Context, dashboardId string) (*UnpublishDashboardResponse, error) {
	return a.lakeviewImpl.Unpublish(ctx, UnpublishDashboardRequest{
		DashboardId: dashboardId,
	})
}

type LakeviewEmbeddedAPI struct {
	lakeviewEmbeddedImpl
}

// Read a published dashboard in an embedded ui.
//
// Get the current published dashboard within an embedded context.
func (a *LakeviewEmbeddedAPI) GetPublishedDashboardEmbeddedByDashboardId(ctx context.Context, dashboardId string) (*GetPublishedDashboardEmbeddedResponse, error) {
	return a.lakeviewEmbeddedImpl.GetPublishedDashboardEmbedded(ctx, GetPublishedDashboardEmbeddedRequest{
		DashboardId: dashboardId,
	})
}

type QueryExecutionAPI struct {
	queryExecutionImpl
}

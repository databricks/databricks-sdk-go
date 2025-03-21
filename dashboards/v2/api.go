// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Genie, Lakeview, Lakeview Embedded, Query Execution, etc.
package dashboards

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type genieBaseClient struct {
	genieImpl
}

// Create conversation message.
//
// Create new message in a [conversation](:method:genie/startconversation). The
// AI response uses all previously created messages in the conversation to
// respond.
func (a *genieBaseClient) CreateMessage(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest) (*GenieCreateMessageWaiter, error) {
	genieMessage, err := a.genieImpl.CreateMessage(ctx, genieCreateConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	return &GenieCreateMessageWaiter{
		RawResponse:    genieMessage,
		conversationId: genieCreateConversationMessageRequest.ConversationId,
		messageId:      genieMessage.Id,
		spaceId:        genieCreateConversationMessageRequest.SpaceId,
		service:        a,
	}, nil
}

type GenieCreateMessageWaiter struct {
	// RawResponse is the raw response of the CreateMessage call.
	RawResponse    *GenieMessage
	service        *genieBaseClient
	conversationId string
	messageId      string
	spaceId        string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *GenieCreateMessageWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*GenieMessage, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[GenieMessage](ctx, opts.Timeout, func() (*GenieMessage, *retries.Err) {
		genieMessage, err := w.service.GetMessage(ctx, GenieGetConversationMessageRequest{
			ConversationId: w.conversationId,
			MessageId:      w.messageId,
			SpaceId:        w.spaceId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := genieMessage.Status
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case MessageStatusCompleted: // target state
			return genieMessage, nil
		case MessageStatusFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				MessageStatusCompleted, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})

}

// Get conversation message.
//
// Get message from conversation.
func (a *genieBaseClient) GetMessageBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieMessage, error) {
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
func (a *genieBaseClient) GetMessageAttachmentQueryResultBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error) {
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
func (a *genieBaseClient) GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieGetMessageQueryResultResponse, error) {
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
func (a *genieBaseClient) GetMessageQueryResultByAttachmentBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error) {
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
func (a *genieBaseClient) GetSpaceBySpaceId(ctx context.Context, spaceId string) (*GenieSpace, error) {
	return a.genieImpl.GetSpace(ctx, GenieGetSpaceRequest{
		SpaceId: spaceId,
	})
}

// Start conversation.
//
// Start a new conversation.
func (a *genieBaseClient) StartConversation(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest) (*GenieStartConversationWaiter, error) {
	genieStartConversationResponse, err := a.genieImpl.StartConversation(ctx, genieStartConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	return &GenieStartConversationWaiter{
		RawResponse:    genieStartConversationResponse,
		conversationId: genieStartConversationResponse.ConversationId,
		messageId:      genieStartConversationResponse.MessageId,
		spaceId:        genieStartConversationMessageRequest.SpaceId,
		service:        a,
	}, nil
}

type GenieStartConversationWaiter struct {
	// RawResponse is the raw response of the StartConversation call.
	RawResponse    *GenieStartConversationResponse
	service        *genieBaseClient
	conversationId string
	messageId      string
	spaceId        string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *GenieStartConversationWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*GenieMessage, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[GenieMessage](ctx, opts.Timeout, func() (*GenieMessage, *retries.Err) {
		genieMessage, err := w.service.GetMessage(ctx, GenieGetConversationMessageRequest{
			ConversationId: w.conversationId,
			MessageId:      w.messageId,
			SpaceId:        w.spaceId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := genieMessage.Status
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case MessageStatusCompleted: // target state
			return genieMessage, nil
		case MessageStatusFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				MessageStatusCompleted, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})

}

type lakeviewBaseClient struct {
	lakeviewImpl
}

// Delete dashboard schedule.
func (a *lakeviewBaseClient) DeleteScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*DeleteScheduleResponse, error) {
	return a.lakeviewImpl.DeleteSchedule(ctx, DeleteScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Delete schedule subscription.
func (a *lakeviewBaseClient) DeleteSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*DeleteSubscriptionResponse, error) {
	return a.lakeviewImpl.DeleteSubscription(ctx, DeleteSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// Get dashboard.
//
// Get a draft dashboard.
func (a *lakeviewBaseClient) GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error) {
	return a.lakeviewImpl.Get(ctx, GetDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get published dashboard.
//
// Get the current published dashboard.
func (a *lakeviewBaseClient) GetPublishedByDashboardId(ctx context.Context, dashboardId string) (*PublishedDashboard, error) {
	return a.lakeviewImpl.GetPublished(ctx, GetPublishedDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get dashboard schedule.
func (a *lakeviewBaseClient) GetScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*Schedule, error) {
	return a.lakeviewImpl.GetSchedule(ctx, GetScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Get schedule subscription.
func (a *lakeviewBaseClient) GetSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*Subscription, error) {
	return a.lakeviewImpl.GetSubscription(ctx, GetSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// List dashboard schedules.
func (a *lakeviewBaseClient) ListSchedulesByDashboardId(ctx context.Context, dashboardId string) (*ListSchedulesResponse, error) {
	return a.lakeviewImpl.internalListSchedules(ctx, ListSchedulesRequest{
		DashboardId: dashboardId,
	})
}

// List schedule subscriptions.
func (a *lakeviewBaseClient) ListSubscriptionsByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*ListSubscriptionsResponse, error) {
	return a.lakeviewImpl.internalListSubscriptions(ctx, ListSubscriptionsRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Trash dashboard.
//
// Trash a dashboard.
func (a *lakeviewBaseClient) TrashByDashboardId(ctx context.Context, dashboardId string) (*TrashDashboardResponse, error) {
	return a.lakeviewImpl.Trash(ctx, TrashDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Unpublish dashboard.
//
// Unpublish the dashboard.
func (a *lakeviewBaseClient) UnpublishByDashboardId(ctx context.Context, dashboardId string) (*UnpublishDashboardResponse, error) {
	return a.lakeviewImpl.Unpublish(ctx, UnpublishDashboardRequest{
		DashboardId: dashboardId,
	})
}

type lakeviewEmbeddedBaseClient struct {
	lakeviewEmbeddedImpl
}

// Read a published dashboard in an embedded ui.
//
// Get the current published dashboard within an embedded context.
func (a *lakeviewEmbeddedBaseClient) GetPublishedDashboardEmbeddedByDashboardId(ctx context.Context, dashboardId string) (*GetPublishedDashboardEmbeddedResponse, error) {
	return a.lakeviewEmbeddedImpl.GetPublishedDashboardEmbedded(ctx, GetPublishedDashboardEmbeddedRequest{
		DashboardId: dashboardId,
	})
}

type queryExecutionBaseClient struct {
	queryExecutionImpl
}

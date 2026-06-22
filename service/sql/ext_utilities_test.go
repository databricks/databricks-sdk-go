package sql

import (
	"context"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/qa"
)

func TestExecuteAndWait_DefaultTimeout(t *testing.T) {
	client, server := qa.HTTPFixtures{
		{
			Method:   "POST",
			Resource: "/api/2.0/sql/statements",
			Response: StatementResponse{
				StatementId: "123",
				Status: &StatementStatus{
					State: StatementStateSucceeded,
				},
			},
		},
	}.Client(t)
	defer server.Close()

	ctx := context.Background()
	api := NewStatementExecution(client)
	resp, err := api.ExecuteAndWait(ctx, ExecuteStatementRequest{
		WarehouseId: "wh-1",
		Statement:   "SELECT 1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if resp.StatementId != "123" {
		t.Errorf("expected statement id 123, got %s", resp.StatementId)
	}
}

func TestExecuteAndWait_CustomTimeoutExceeded(t *testing.T) {
	client, server := qa.HTTPFixtures{
		{
			Method:   "POST",
			Resource: "/api/2.0/sql/statements",
			Response: StatementResponse{
				StatementId: "123",
				Status: &StatementStatus{
					State: StatementStatePending,
				},
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/sql/statements/123",
			Response: StatementResponse{
				StatementId: "123",
				Status: &StatementStatus{
					State: StatementStateRunning,
				},
			},
		},
	}.Client(t)
	defer server.Close()

	ctx := context.Background()
	api := NewStatementExecution(client)

	start := time.Now()
	_, err := api.ExecuteAndWait(ctx, ExecuteStatementRequest{
		WarehouseId: "wh-1",
		Statement:   "SELECT 1",
		WaitTimeout: "500ms",
	})
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}

	// We expect the polling to timeout around 500ms.
	// Allow some buffer (e.g., up to 2s) to prevent transient test failures.
	if elapsed > 2*time.Second {
		t.Errorf("expected poll to timeout quickly, took %s", elapsed)
	}
}

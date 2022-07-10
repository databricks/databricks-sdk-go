package commands

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/databricks/sdk-go/databricks"
	"github.com/databricks/sdk-go/databricks/apierr"
	"github.com/databricks/sdk-go/databricks/qa"
	"github.com/databricks/sdk-go/service/clusters"
)

// Test interface compliance
var _ CommandExecutor = (*CommandsAPI)(nil)

func commonFixtureWithStatusResponse(response Command) qa.HTTPFixtures {
	return []qa.HTTPFixture{
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				ClusterID:              "abc",
				NumWorkers:             100,
				ClusterName:            "Shared Autoscaling",
				SparkVersion:           "7.1-scala12",
				NodeTypeID:             "i3.xlarge",
				AutoterminationMinutes: 15,
				State:                  clusters.ClusterStateRunning,
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			ExpectedRequest: map[string]interface{}{
				"clusterId": "abc",
				"language":  "python",
			},
			Response: Command{
				ID: "123",
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/1.2/contexts/status?clusterId=abc&contextId=123",
			ReuseRequest: true,
			Response: Command{
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			ExpectedRequest: genericCommandRequest{
				Language:  "python",
				ClusterID: "abc",
				ContextID: "123",
				Command:   "print(\"done\")\n",
			},
			Response: Command{
				ID: "234",
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/1.2/commands/status?clusterId=abc&commandId=234&contextId=123",
			Response:     response,
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/destroy",
			ExpectedRequest: genericCommandRequest{
				ClusterID: "abc",
				ContextID: "123",
			},
		},
	}
}

func TestCommandWithExecutionError(t *testing.T) {
	cfg, server := commonFixtureWithStatusResponse(Command{
		Status: "Finished",
		Results: &CommandResults{
			ResultType: "error",
			Cause: `
---
ExecutionError: An error occurred
StatusCode=400
StatusDescription=BadRequest
---
			`,
		},
	}).Config(t)
	defer server.Close()
	ctx := context.Background()
	commands := New(cfg)

	result := commands.Execute(ctx, "abc", "python", `print("done")`)
	assert.Equal(t, true, result.Failed())
	assert.Equal(t, "An error occurred\nStatusCode=400\nStatusDescription=BadRequest", result.Error())
}

func TestCommandWithEmptyErrorMessageUsesSummary(t *testing.T) {
	cfg, server := commonFixtureWithStatusResponse(Command{
		Status: "Finished",
		Results: &CommandResults{
			ResultType: "error",
			Cause: `
---
ErrorCode=
ErrorMessage=
    other text
---
			`,
			Summary: "Proper error",
		},
	}).Config(t)
	defer server.Close()
	ctx := context.Background()
	commands := New(cfg)

	result := commands.Execute(ctx, "abc", "python", `print("done")`)
	assert.Equal(t, true, result.Failed())
	assert.Equal(t, "Proper error", result.Error())
}

func TestCommandWithErrorMessage(t *testing.T) {
	cfg, server := commonFixtureWithStatusResponse(Command{
		Status: "Finished",
		Results: &CommandResults{
			ResultType: "error",
			Cause: `
---
ErrorCode=
ErrorMessage=An error occurred
---
			`,
		},
	}).Config(t)
	defer server.Close()
	ctx := context.Background()
	commands := New(cfg)

	result := commands.Execute(ctx, "abc", "python", `print("done")`)
	assert.Equal(t, true, result.Failed())
	assert.Equal(t, "An error occurred", result.Error())
}

func TestCommandWithExceptionMessage(t *testing.T) {
	cfg, server := commonFixtureWithStatusResponse(Command{
		Status: "Finished",
		Results: &CommandResults{
			ResultType: "error",
			Summary:    "Exception: An error occurred",
		},
	}).Config(t)
	defer server.Close()
	ctx := context.Background()
	commands := New(cfg)

	result := commands.Execute(ctx, "abc", "python", `print("done")`)
	assert.Equal(t, true, result.Failed())
	assert.Equal(t, "An error occurred", result.Error())
}

func TestSomeCommands(t *testing.T) {
	cfg, server := commonFixtureWithStatusResponse(Command{
		Status: "Finished",
		Results: &CommandResults{
			ResultType: "text",
			Data:       "done",
		},
	}).Config(t)
	defer server.Close()
	ctx := context.Background()
	commands := New(cfg)

	result := commands.Execute(ctx, "abc", "python", `print("done")`)
	assert.Equal(t, false, result.Failed())
	assert.Equal(t, "done", result.Text())
}

func TestCommandsAPIExecute_FailGettingCluster(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}.Apply(t, func(ctx context.Context, cfg *databricks.Config) {
		commands := New(cfg)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_StoppedCluster(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "TERMINATED",
			},
		},
	}.Apply(t, func(ctx context.Context, cfg *databricks.Config) {
		commands := New(cfg)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Cluster abc has to be running or resizing, but is TERMINATED")
	})
}

func TestCommandsAPIExecute_FailToCreateContext(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}.Apply(t, func(ctx context.Context, cfg *databricks.Config) {
		commands := New(cfg)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToWaitForContext(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}.Apply(t, func(ctx context.Context, cfg *databricks.Config) {
		commands := New(cfg)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToCreateCommand(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: Command{
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}.Apply(t, func(ctx context.Context, cfg *databricks.Config) {
		commands := New(cfg)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToWaitForCommand(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: Command{
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/commands/status?clusterId=abc&commandId=abc&contextId=abc",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}.Apply(t, func(ctx context.Context, cfg *databricks.Config) {
		commands := New(cfg)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToGetCommand(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: Command{
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/commands/status?clusterId=abc&commandId=abc&contextId=abc",
			Response: Command{
				Status: "Finished",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/commands/status?clusterId=abc&commandId=abc&contextId=abc",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}.Apply(t, func(ctx context.Context, cfg *databricks.Config) {
		commands := New(cfg)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_FailToDeleteContext(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: Command{
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/1.2/commands/status?clusterId=abc&commandId=abc&contextId=abc",
			Response: Command{
				Status: "Finished",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/destroy",
			Status:   417,
			Response: apierr.APIError{
				Message: "Does not compute",
			},
		},
	}.Apply(t, func(ctx context.Context, cfg *databricks.Config) {
		commands := New(cfg)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Does not compute")
	})
}

func TestCommandsAPIExecute_NoCommandResults(t *testing.T) {
	qa.HTTPFixtures{
		{
			Method:   "GET",
			Resource: "/api/2.0/clusters/get?cluster_id=abc",
			Response: clusters.ClusterInfo{
				State: "RUNNING",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/create",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:   "GET",
			Resource: "/api/1.2/contexts/status?clusterId=abc&contextId=abc",
			Response: Command{
				Status: "Running",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/commands/execute",
			Response: Command{
				ID: "abc",
			},
		},
		{
			Method:       "GET",
			ReuseRequest: true,
			Resource:     "/api/1.2/commands/status?clusterId=abc&commandId=abc&contextId=abc",
			Response: Command{
				Status: "Finished",
			},
		},
		{
			Method:   "POST",
			Resource: "/api/1.2/contexts/destroy",
		},
	}.Apply(t, func(ctx context.Context, cfg *databricks.Config) {
		commands := New(cfg)
		cr := commands.Execute(ctx, "abc", "cobol", "Hello?")
		assert.EqualError(t, cr.Err(), "Command has no results")
	})
}

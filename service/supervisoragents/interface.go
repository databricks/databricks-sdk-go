// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package supervisoragents

import (
	"context"
)

// Manage Supervisor Agents and related resources.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type SupervisorAgentsService interface {

	// Creates a new Supervisor Agent.
	CreateSupervisorAgent(ctx context.Context, request CreateSupervisorAgentRequest) (*SupervisorAgent, error)

	// Creates a Tool under a Supervisor Agent. Specify one of "genie_space",
	// "knowledge_assistant", "uc_function", "uc_connection", "app", "volume",
	// "lakeview_dashboard", "uc_table", "vector_search_index", "catalog",
	// "schema" in the request body.
	CreateTool(ctx context.Context, request CreateToolRequest) (*Tool, error)

	// Deletes a Supervisor Agent.
	DeleteSupervisorAgent(ctx context.Context, request DeleteSupervisorAgentRequest) error

	// Deletes a Tool.
	DeleteTool(ctx context.Context, request DeleteToolRequest) error

	// Gets a Supervisor Agent.
	GetSupervisorAgent(ctx context.Context, request GetSupervisorAgentRequest) (*SupervisorAgent, error)

	// Gets a Tool.
	GetTool(ctx context.Context, request GetToolRequest) (*Tool, error)

	// Lists Supervisor Agents.
	ListSupervisorAgents(ctx context.Context, request ListSupervisorAgentsRequest) (*ListSupervisorAgentsResponse, error)

	// Lists Tools under a Supervisor Agent.
	ListTools(ctx context.Context, request ListToolsRequest) (*ListToolsResponse, error)

	// Updates a Supervisor Agent. The fields that are required depend on the
	// paths specified in `update_mask`. Only fields included in the mask will
	// be updated.
	UpdateSupervisorAgent(ctx context.Context, request UpdateSupervisorAgentRequest) (*SupervisorAgent, error)

	// Updates a Tool. Only the `description` field can be updated. To change
	// immutable fields such as tool type, spec, or tool ID, delete the tool and
	// recreate it.
	UpdateTool(ctx context.Context, request UpdateToolRequest) (*Tool, error)
}

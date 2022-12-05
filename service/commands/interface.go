// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package commands

import (
	"context"
)

// This API allows execution of Python, Scala, SQL, or R commands on running
// Databricks Clusters.
type CommandExecutionService interface {

	// Cancel a command.
	//
	// Cancels a currently running command within an execution context.
	//
	// The command ID is obtained from a prior successful call to __execute__.
	Cancel(ctx context.Context, request CancelCommand) error

	// Get command info.
	//
	// Gets the status of and, if available, the results from a currently
	// executing command.
	//
	// The command ID is obtained from a prior successful call to __execute__.
	CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error)

	// Get status.
	//
	// Gets the status for an execution context.
	ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error)

	// Create an execution context.
	//
	// Creates an execution context for running cluster commands.
	//
	// If successful, this method returns the ID of the new execution context.
	Create(ctx context.Context, request CreateContext) (*Created, error)

	// Delete an execution context.
	//
	// Deletes an execution context.
	Destroy(ctx context.Context, request DestroyContext) error

	// Run a command.
	//
	// Runs a cluster command in the given execution context, using the provided
	// language.
	//
	// If successful, it returns an ID for tracking the status of the command's
	// execution.
	Execute(ctx context.Context, request Command) (*Created, error)
}

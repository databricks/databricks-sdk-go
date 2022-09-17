// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package commands

import (
	"context"
	"time"
)

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type CommandExecutionService interface {
	Cancel(ctx context.Context, request CancelCommand) error

	// CancelAndWait calls Cancel() and waits to reach Cancelled state
	//
	// This method is generated by Databricks SDK Code Generator.
	CancelAndWait(ctx context.Context, request CancelCommand, timeout ...time.Duration) (*CommandStatusResponse, error)

	CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error)

	ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error)

	Create(ctx context.Context, request CreateContext) (*Created, error)

	// CreateAndWait calls Create() and waits to reach Running state
	//
	// This method is generated by Databricks SDK Code Generator.
	CreateAndWait(ctx context.Context, request CreateContext, timeout ...time.Duration) (*ContextStatusResponse, error)

	Destroy(ctx context.Context, request DestroyRequest) error

	Execute(ctx context.Context, request Command) (*Created, error)

	// ExecuteAndWait calls Execute() and waits to reach Finished or Error state
	//
	// This method is generated by Databricks SDK Code Generator.
	ExecuteAndWait(ctx context.Context, request Command, timeout ...time.Duration) (*CommandStatusResponse, error)
}

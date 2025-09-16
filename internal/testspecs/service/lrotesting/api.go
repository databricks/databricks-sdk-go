// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Test service for Long Running Operations
package lrotesting

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/internal/testspecs/service/common"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/common/lro"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type LroTestingInterface interface {
	CancelOperation(ctx context.Context, request CancelOperationRequest) error

	CreateTestResource(ctx context.Context, request CreateTestResourceRequest) (*CreateTestResourceOperation, error)

	GetOperation(ctx context.Context, request GetOperationRequest) (*common.Operation, error)

	// Simple method to get test resource
	GetTestResource(ctx context.Context, request GetTestResourceRequest) (*TestResource, error)
}

func NewLroTesting(client *client.DatabricksClient) *LroTestingAPI {
	return &LroTestingAPI{
		lroTestingImpl: lroTestingImpl{
			client: client,
		},
	}
}

// Test service for Long Running Operations
type LroTestingAPI struct {
	lroTestingImpl
}

func (a *LroTestingAPI) CreateTestResource(ctx context.Context, request CreateTestResourceRequest) (*CreateTestResourceOperation, error) {
	operation, err := a.lroTestingImpl.CreateTestResource(ctx, request)
	if err != nil {
		return nil, err
	}
	return &CreateTestResourceOperation{
		impl:      &a.lroTestingImpl,
		operation: operation,
	}, nil
}

type CreateTestResourceOperation struct {
	impl      *lroTestingImpl
	operation *common.Operation
}

// Wait blocks until the long-running operation is completed with default 20 min
// timeout, the timeout can be overridden in the opts. If the operation didn't
// finished within the timeout, this function will through an error of type
// ErrTimedOut, otherwise successful response and any errors encountered.
func (a *CreateTestResourceOperation) Wait(ctx context.Context, opts *lro.LroOptions) (*TestResource, error) {
	timeout := 20 * time.Minute // default timeout per LRO spec
	if opts != nil && opts.Timeout > 0 {
		timeout = opts.Timeout
	}

	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[TestResource](ctx, timeout, func() (*TestResource, *retries.Err) {
		operation, err := a.impl.GetOperation(ctx, GetOperationRequest{
			Name: a.operation.Name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}

		// Update local operation state
		a.operation = operation

		if !operation.Done {
			return nil, retries.Continues("operation still in progress")
		}

		if operation.Error != nil {
			var errorMsg string
			if operation.Error.Message != "" {
				errorMsg = operation.Error.Message
			} else {
				errorMsg = "unknown error"
			}

			if operation.Error.ErrorCode != "" {
				errorMsg = fmt.Sprintf("[%s] %s", operation.Error.ErrorCode, errorMsg)
			}

			return nil, retries.Halt(fmt.Errorf("operation failed: %s", errorMsg))
		}

		// Operation completed successfully, unmarshal response
		if operation.Response == nil {
			return nil, retries.Halt(fmt.Errorf("operation completed but no response available"))
		}

		var testResource TestResource
		err = json.Unmarshal(operation.Response, &testResource)
		if err != nil {
			return nil, retries.Halt(fmt.Errorf("failed to unmarshal testResource response: %w", err))
		}

		return &testResource, nil
	})
}

// Starts asynchronous cancellation on a long-running operation. The server
// makes a best effort to cancel the operation, but success is not guaranteed.
func (a *CreateTestResourceOperation) Cancel(ctx context.Context) error {
	return a.impl.CancelOperation(ctx, CancelOperationRequest{
		Name: a.operation.Name,
	})
}

// Name returns the name of the long-running operation. The name is assigned
// by the server and is unique within the service from which the operation is created.
func (a *CreateTestResourceOperation) Name() string {
	return a.operation.Name
}

// Metadata returns metadata associated with the long-running operation.
// If the metadata is not available, the returned metadata and error are both nil.
func (a *CreateTestResourceOperation) Metadata() (*TestResourceOperationMetadata, error) {
	if a.operation.Metadata == nil {
		return nil, nil
	}

	var metadata TestResourceOperationMetadata
	err := json.Unmarshal(a.operation.Metadata, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal operation metadata: %w", err)
	}

	return &metadata, nil
}

// Done reports whether the long-running operation has completed.
func (a *CreateTestResourceOperation) Done() (bool, error) {
	// Refresh the operation state first
	operation, err := a.impl.GetOperation(context.Background(), GetOperationRequest{
		Name: a.operation.Name,
	})
	if err != nil {
		return false, err
	}

	// Update local operation state
	a.operation = operation

	return operation.Done, nil
}

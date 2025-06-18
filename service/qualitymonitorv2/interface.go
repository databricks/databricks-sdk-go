// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package qualitymonitorv2

import (
	"context"
)

// Manage data quality of UC objects (currently support `schema`)
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type QualityMonitorV2Service interface {

	// Create a quality monitor on UC object
	CreateQualityMonitor(ctx context.Context, request CreateQualityMonitorRequest) (*QualityMonitor, error)

	// Delete a quality monitor on UC object
	DeleteQualityMonitor(ctx context.Context, request DeleteQualityMonitorRequest) error

	// Read a quality monitor on UC object
	GetQualityMonitor(ctx context.Context, request GetQualityMonitorRequest) (*QualityMonitor, error)

	// (Unimplemented) List quality monitors
	ListQualityMonitor(ctx context.Context, request ListQualityMonitorRequest) (*ListQualityMonitorResponse, error)

	// (Unimplemented) Update a quality monitor on UC object
	UpdateQualityMonitor(ctx context.Context, request UpdateQualityMonitorRequest) (*QualityMonitor, error)
}

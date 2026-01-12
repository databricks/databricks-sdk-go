// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package qualitymonitorv2

import (
	"context"
)

// [DEPRECATED] This API is deprecated. Please use the Data Quality Monitoring
// API instead (REST: /api/data-quality/v1/monitors). Manage data quality of UC
// objects (currently support `schema`).
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type QualityMonitorV2Service interface {

	// [DEPRECATED] Create a quality monitor on UC object. Use Data Quality
	// Monitoring API instead.
	CreateQualityMonitor(ctx context.Context, request CreateQualityMonitorRequest) (*QualityMonitor, error)

	// [DEPRECATED] Delete a quality monitor on UC object. Use Data Quality
	// Monitoring API instead.
	DeleteQualityMonitor(ctx context.Context, request DeleteQualityMonitorRequest) error

	// [DEPRECATED] Read a quality monitor on UC object. Use Data Quality
	// Monitoring API instead.
	GetQualityMonitor(ctx context.Context, request GetQualityMonitorRequest) (*QualityMonitor, error)

	// [DEPRECATED] (Unimplemented) List quality monitors. Use Data Quality
	// Monitoring API instead.
	ListQualityMonitor(ctx context.Context, request ListQualityMonitorRequest) (*ListQualityMonitorResponse, error)

	// [DEPRECATED] (Unimplemented) Update a quality monitor on UC object. Use
	// Data Quality Monitoring API instead.
	UpdateQualityMonitor(ctx context.Context, request UpdateQualityMonitorRequest) (*QualityMonitor, error)
}

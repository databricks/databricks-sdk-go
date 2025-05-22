// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package qualitymonitor

import (
	"context"
)

// Manage data quality of UC objects (currently support `schema`)
type QualityMonitorsV2Service interface {

	// Create a quality monitor.
	//
	// Create a quality monitor on UC object
	CreateQualityMonitor(ctx context.Context, request CreateQualityMonitorRequest) (*QualityMonitor, error)

	// Delete a quality monitor.
	//
	// Delete a quality monitor on UC object
	DeleteQualityMonitor(ctx context.Context, request DeleteQualityMonitorRequest) error

	// Read a quality monitor.
	//
	// Read a quality monitor on UC object
	GetQualityMonitor(ctx context.Context, request GetQualityMonitorRequest) (*QualityMonitor, error)

	// List quality monitors.
	//
	// (Unimplemented) List quality monitors
	//
	// Use ListQualityMonitorAll() to get all QualityMonitor instances, which will iterate over every result page.
	ListQualityMonitor(ctx context.Context, request ListQualityMonitorRequest) (*ListQualityMonitorResponse, error)

	// Update a quality monitor.
	//
	// Update a quality monitor on UC object
	UpdateQualityMonitor(ctx context.Context, request UpdateQualityMonitorRequest) (*QualityMonitor, error)
}

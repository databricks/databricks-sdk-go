// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dataquality

import (
	"context"
)

// Manage the data quality of Unity Catalog objects (currently support `schema`
// and `table`)
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DataQualityService interface {

	// Cancels a data quality monitor refresh. Currently only supported for the
	// `table` `object_type`.
	CancelRefresh(ctx context.Context, request CancelRefreshRequest) (*CancelRefreshResponse, error)

	// Create a data quality monitor on a Unity Catalog object. The caller must
	// provide either `anomaly_detection_config` for a schema monitor or
	// `data_profiling_config` for a table monitor.
	//
	// For the `table` `object_type`, the caller must either: 1. be an owner of
	// the table's parent catalog, have **USE_SCHEMA** on the table's parent
	// schema, and have **SELECT** access on the table 2. have **USE_CATALOG**
	// on the table's parent catalog, be an owner of the table's parent schema,
	// and have **SELECT** access on the table. 3. have the following
	// permissions: - **USE_CATALOG** on the table's parent catalog -
	// **USE_SCHEMA** on the table's parent schema - be an owner of the table.
	//
	// Workspace assets, such as the dashboard, will be created in the workspace
	// where this call was made.
	CreateMonitor(ctx context.Context, request CreateMonitorRequest) (*Monitor, error)

	// Creates a refresh. Currently only supported for the `table`
	// `object_type`.
	//
	// The caller must either: 1. be an owner of the table's parent catalog 2.
	// have **USE_CATALOG** on the table's parent catalog and be an owner of the
	// table's parent schema 3. have the following permissions: -
	// **USE_CATALOG** on the table's parent catalog - **USE_SCHEMA** on the
	// table's parent schema - be an owner of the table
	CreateRefresh(ctx context.Context, request CreateRefreshRequest) (*Refresh, error)

	// Delete a data quality monitor on Unity Catalog object.
	//
	// For the `table` `object_type`, the caller must either: 1. be an owner of
	// the table's parent catalog 2. have **USE_CATALOG** on the table's parent
	// catalog and be an owner of the table's parent schema 3. have the
	// following permissions: - **USE_CATALOG** on the table's parent catalog -
	// **USE_SCHEMA** on the table's parent schema - be an owner of the table.
	//
	// Note that the metric tables and dashboard will not be deleted as part of
	// this call; those assets must be manually cleaned up (if desired).
	DeleteMonitor(ctx context.Context, request DeleteMonitorRequest) error

	// (Unimplemented) Delete a refresh
	DeleteRefresh(ctx context.Context, request DeleteRefreshRequest) error

	// Read a data quality monitor on Unity Catalog object.
	//
	// For the `table` `object_type`, the caller must either: 1. be an owner of
	// the table's parent catalog 2. have **USE_CATALOG** on the table's parent
	// catalog and be an owner of the table's parent schema. 3. have the
	// following permissions: - **USE_CATALOG** on the table's parent catalog -
	// **USE_SCHEMA** on the table's parent schema - **SELECT** privilege on the
	// table.
	//
	// The returned information includes configuration values, as well as
	// information on assets created by the monitor. Some information (e.g.,
	// dashboard) may be filtered out if the caller is in a different workspace
	// than where the monitor was created.
	GetMonitor(ctx context.Context, request GetMonitorRequest) (*Monitor, error)

	// Get data quality monitor refresh.
	//
	// For the `table` `object_type`, the caller must either: 1. be an owner of
	// the table's parent catalog 2. have **USE_CATALOG** on the table's parent
	// catalog and be an owner of the table's parent schema 3. have the
	// following permissions: - **USE_CATALOG** on the table's parent catalog -
	// **USE_SCHEMA** on the table's parent schema - **SELECT** privilege on the
	// table.
	GetRefresh(ctx context.Context, request GetRefreshRequest) (*Refresh, error)

	// (Unimplemented) List data quality monitors.
	ListMonitor(ctx context.Context, request ListMonitorRequest) (*ListMonitorResponse, error)

	// List data quality monitor refreshes.
	//
	// For the `table` `object_type`, the caller must either: 1. be an owner of
	// the table's parent catalog 2. have **USE_CATALOG** on the table's parent
	// catalog and be an owner of the table's parent schema 3. have the
	// following permissions: - **USE_CATALOG** on the table's parent catalog -
	// **USE_SCHEMA** on the table's parent schema - **SELECT** privilege on the
	// table.
	ListRefresh(ctx context.Context, request ListRefreshRequest) (*ListRefreshResponse, error)

	// Update a data quality monitor on Unity Catalog object.
	//
	// For the `table` `object_type`, The caller must either: 1. be an owner of
	// the table's parent catalog 2. have **USE_CATALOG** on the table's parent
	// catalog and be an owner of the table's parent schema 3. have the
	// following permissions: - **USE_CATALOG** on the table's parent catalog -
	// **USE_SCHEMA** on the table's parent schema - be an owner of the table.
	UpdateMonitor(ctx context.Context, request UpdateMonitorRequest) (*Monitor, error)

	// (Unimplemented) Update a refresh
	UpdateRefresh(ctx context.Context, request UpdateRefreshRequest) (*Refresh, error)
}

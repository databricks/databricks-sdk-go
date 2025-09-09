// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database

import (
	"context"
)

// Database Instances provide access to a database via REST API or direct SQL.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DatabaseService interface {

	// Create a Database Catalog.
	CreateDatabaseCatalog(ctx context.Context, request CreateDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Create a Database Instance.
	CreateDatabaseInstance(ctx context.Context, request CreateDatabaseInstanceRequest) (*DatabaseInstance, error)

	// Create a role for a Database Instance.
	CreateDatabaseInstanceRole(ctx context.Context, request CreateDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error)

	// Create a Database Table. Useful for registering pre-existing PG tables in
	// UC. See CreateSyncedDatabaseTable for creating synced tables in PG from a
	// source table in UC.
	CreateDatabaseTable(ctx context.Context, request CreateDatabaseTableRequest) (*DatabaseTable, error)

	// Create a Synced Database Table.
	CreateSyncedDatabaseTable(ctx context.Context, request CreateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	// Delete a Database Catalog.
	DeleteDatabaseCatalog(ctx context.Context, request DeleteDatabaseCatalogRequest) error

	// Delete a Database Instance.
	DeleteDatabaseInstance(ctx context.Context, request DeleteDatabaseInstanceRequest) error

	// Deletes a role for a Database Instance.
	DeleteDatabaseInstanceRole(ctx context.Context, request DeleteDatabaseInstanceRoleRequest) error

	// Delete a Database Table.
	DeleteDatabaseTable(ctx context.Context, request DeleteDatabaseTableRequest) error

	// Delete a Synced Database Table.
	DeleteSyncedDatabaseTable(ctx context.Context, request DeleteSyncedDatabaseTableRequest) error

	// Failover the primary node of a Database Instance to a secondary.
	FailoverDatabaseInstance(ctx context.Context, request FailoverDatabaseInstanceRequest) (*DatabaseInstance, error)

	// Find a Database Instance by uid.
	FindDatabaseInstanceByUid(ctx context.Context, request FindDatabaseInstanceByUidRequest) (*DatabaseInstance, error)

	// Generates a credential that can be used to access database instances.
	GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error)

	// Get a Database Catalog.
	GetDatabaseCatalog(ctx context.Context, request GetDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Get a Database Instance.
	GetDatabaseInstance(ctx context.Context, request GetDatabaseInstanceRequest) (*DatabaseInstance, error)

	// Gets a role for a Database Instance.
	GetDatabaseInstanceRole(ctx context.Context, request GetDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error)

	// Get a Database Table.
	GetDatabaseTable(ctx context.Context, request GetDatabaseTableRequest) (*DatabaseTable, error)

	// Get a Synced Database Table.
	GetSyncedDatabaseTable(ctx context.Context, request GetSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	ListDatabaseCatalogs(ctx context.Context, request ListDatabaseCatalogsRequest) (*ListDatabaseCatalogsResponse, error)

	// START OF PG ROLE APIs Section These APIs are marked a PUBLIC with stage <
	// PUBLIC_PREVIEW. With more recent Lakebase V2 plans, we don't plan to ever
	// advance these to PUBLIC_PREVIEW. These APIs will remain effectively
	// undocumented/UI-only and we'll aim for a new public roles API as part of
	// V2 PuPr.
	ListDatabaseInstanceRoles(ctx context.Context, request ListDatabaseInstanceRolesRequest) (*ListDatabaseInstanceRolesResponse, error)

	// List Database Instances.
	ListDatabaseInstances(ctx context.Context, request ListDatabaseInstancesRequest) (*ListDatabaseInstancesResponse, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	ListSyncedDatabaseTables(ctx context.Context, request ListSyncedDatabaseTablesRequest) (*ListSyncedDatabaseTablesResponse, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	UpdateDatabaseCatalog(ctx context.Context, request UpdateDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Update a Database Instance.
	UpdateDatabaseInstance(ctx context.Context, request UpdateDatabaseInstanceRequest) (*DatabaseInstance, error)

	// Update a role for a Database Instance.
	UpdateDatabaseInstanceRole(ctx context.Context, request UpdateDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	UpdateSyncedDatabaseTable(ctx context.Context, request UpdateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)
}

// Database Projects provide access to a database via REST API or direct SQL.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DatabaseProjectService interface {

	// Create a Database Branch.
	CreateDatabaseBranch(ctx context.Context, request CreateDatabaseBranchRequest) (*DatabaseBranch, error)

	// Create a Database Endpoint.
	CreateDatabaseEndpoint(ctx context.Context, request CreateDatabaseEndpointRequest) (*DatabaseEndpoint, error)

	// Create a Database Project.
	CreateDatabaseProject(ctx context.Context, request CreateDatabaseProjectRequest) (*DatabaseProject, error)

	// Delete a Database Branch.
	DeleteDatabaseBranch(ctx context.Context, request DeleteDatabaseBranchRequest) error

	// Delete a Database Endpoint.
	DeleteDatabaseEndpoint(ctx context.Context, request DeleteDatabaseEndpointRequest) error

	// Delete a Database Project.
	DeleteDatabaseProject(ctx context.Context, request DeleteDatabaseProjectRequest) error

	// Get a Database Branch.
	GetDatabaseBranch(ctx context.Context, request GetDatabaseBranchRequest) (*DatabaseBranch, error)

	// Get a Database Endpoint.
	GetDatabaseEndpoint(ctx context.Context, request GetDatabaseEndpointRequest) (*DatabaseEndpoint, error)

	// Get a Database Project.
	GetDatabaseProject(ctx context.Context, request GetDatabaseProjectRequest) (*DatabaseProject, error)

	// List Database Branches.
	ListDatabaseBranches(ctx context.Context, request ListDatabaseBranchesRequest) (*ListDatabaseBranchesResponse, error)

	// List Database Endpoints.
	ListDatabaseEndpoints(ctx context.Context, request ListDatabaseEndpointsRequest) (*ListDatabaseEndpointsResponse, error)

	// List Database Instances.
	ListDatabaseProjects(ctx context.Context, request ListDatabaseProjectsRequest) (*ListDatabaseProjectsResponse, error)

	// Restart a Database Endpoint. TODO: should return
	// databricks.longrunning.Operation
	RestartDatabaseEndpoint(ctx context.Context, request RestartDatabaseEndpointRequest) (*DatabaseEndpoint, error)

	// Update a Database Branch.
	UpdateDatabaseBranch(ctx context.Context, request UpdateDatabaseBranchRequest) (*DatabaseBranch, error)

	// Update a Database Endpoint. TODO: should return
	// databricks.longrunning.Operation {
	UpdateDatabaseEndpoint(ctx context.Context, request UpdateDatabaseEndpointRequest) (*DatabaseEndpoint, error)

	// Update a Database Project.
	UpdateDatabaseProject(ctx context.Context, request UpdateDatabaseProjectRequest) (*DatabaseProject, error)
}

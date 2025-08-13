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

	// This API is currently unimplemented, but exposed for Terraform support.
	UpdateSyncedDatabaseTable(ctx context.Context, request UpdateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)
}

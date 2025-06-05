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

	// Create a Database Table.
	CreateDatabaseTable(ctx context.Context, request CreateDatabaseTableRequest) (*DatabaseTable, error)

	// Create a Synced Database Table.
	CreateSyncedDatabaseTable(ctx context.Context, request CreateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	// Delete a Database Catalog.
	DeleteDatabaseCatalog(ctx context.Context, request DeleteDatabaseCatalogRequest) error

	// Delete a Database Instance.
	DeleteDatabaseInstance(ctx context.Context, request DeleteDatabaseInstanceRequest) error

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

	// Get a Database Table.
	GetDatabaseTable(ctx context.Context, request GetDatabaseTableRequest) (*DatabaseTable, error)

	// Get a Synced Database Table.
	GetSyncedDatabaseTable(ctx context.Context, request GetSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	// List Database Instances.
	ListDatabaseInstances(ctx context.Context, request ListDatabaseInstancesRequest) (*ListDatabaseInstancesResponse, error)

	// Update a Database Instance.
	UpdateDatabaseInstance(ctx context.Context, request UpdateDatabaseInstanceRequest) (*DatabaseInstance, error)
}

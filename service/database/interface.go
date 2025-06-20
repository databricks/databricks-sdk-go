// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database

import (
	"context"
)

// Database Instances provide access to a database via REST API or direct SQL.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DatabaseService interface {
	CreateDatabaseCatalog(ctx context.Context, request CreateDatabaseCatalogRequest) (*DatabaseCatalog, error)

	CreateDatabaseInstance(ctx context.Context, request CreateDatabaseInstanceRequest) (*DatabaseInstance, error)

	CreateDatabaseTable(ctx context.Context, request CreateDatabaseTableRequest) (*DatabaseTable, error)

	CreateSyncedDatabaseTable(ctx context.Context, request CreateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	DeleteDatabaseCatalog(ctx context.Context, request DeleteDatabaseCatalogRequest) error

	DeleteDatabaseInstance(ctx context.Context, request DeleteDatabaseInstanceRequest) error

	DeleteDatabaseTable(ctx context.Context, request DeleteDatabaseTableRequest) error

	DeleteSyncedDatabaseTable(ctx context.Context, request DeleteSyncedDatabaseTableRequest) error

	FindDatabaseInstanceByUid(ctx context.Context, request FindDatabaseInstanceByUidRequest) (*DatabaseInstance, error)

	GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error)

	GetDatabaseCatalog(ctx context.Context, request GetDatabaseCatalogRequest) (*DatabaseCatalog, error)

	GetDatabaseInstance(ctx context.Context, request GetDatabaseInstanceRequest) (*DatabaseInstance, error)

	GetDatabaseTable(ctx context.Context, request GetDatabaseTableRequest) (*DatabaseTable, error)

	GetSyncedDatabaseTable(ctx context.Context, request GetSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	ListDatabaseInstances(ctx context.Context, request ListDatabaseInstancesRequest) (*ListDatabaseInstancesResponse, error)

	UpdateDatabaseInstance(ctx context.Context, request UpdateDatabaseInstanceRequest) (*DatabaseInstance, error)
}

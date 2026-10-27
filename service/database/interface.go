// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database

import (
	"context"
)

// Database Instances provide access to a database via REST API or direct SQL.
//
// For new workloads, Databricks recommends the [Postgres API], which supports
// the latest Lakebase features.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Postgres API]: https://docs.databricks.com/api/postgres/v1/postgres
type DatabaseService interface {

	// Create a Database Catalog.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Create a catalog].
	//
	// [Create a catalog]: https://docs.databricks.com/api/postgres/v1/create-catalog
	CreateDatabaseCatalog(ctx context.Context, request CreateDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Create a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Create a project].
	//
	// [Create a project]: https://docs.databricks.com/api/postgres/v1/create-project
	CreateDatabaseInstance(ctx context.Context, request CreateDatabaseInstanceRequest) (*DatabaseInstance, error)

	// Create a role for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Create a role].
	//
	// [Create a role]: https://docs.databricks.com/api/postgres/v1/create-role
	CreateDatabaseInstanceRole(ctx context.Context, request CreateDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error)

	// Create a Database Table. Useful for registering pre-existing PG tables in
	// UC. See CreateSyncedDatabaseTable for creating synced tables in PG from a
	// source table in UC.
	CreateDatabaseTable(ctx context.Context, request CreateDatabaseTableRequest) (*DatabaseTable, error)

	// Create a Synced Database Table.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Create a synced table].
	//
	// [Create a synced table]: https://docs.databricks.com/api/postgres/v1/create-synced-table
	CreateSyncedDatabaseTable(ctx context.Context, request CreateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	// Delete a Database Catalog.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Delete a catalog].
	//
	// [Delete a catalog]: https://docs.databricks.com/api/postgres/v1/delete-catalog
	DeleteDatabaseCatalog(ctx context.Context, request DeleteDatabaseCatalogRequest) error

	// Delete a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Delete a project].
	//
	// [Delete a project]: https://docs.databricks.com/api/postgres/v1/delete-project
	DeleteDatabaseInstance(ctx context.Context, request DeleteDatabaseInstanceRequest) error

	// Deletes a role for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Delete a role].
	//
	// [Delete a role]: https://docs.databricks.com/api/postgres/v1/delete-role
	DeleteDatabaseInstanceRole(ctx context.Context, request DeleteDatabaseInstanceRoleRequest) error

	// Delete a Database Table.
	DeleteDatabaseTable(ctx context.Context, request DeleteDatabaseTableRequest) error

	// Delete a Synced Database Table.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Delete a synced table].
	//
	// [Delete a synced table]: https://docs.databricks.com/api/postgres/v1/delete-synced-table
	DeleteSyncedDatabaseTable(ctx context.Context, request DeleteSyncedDatabaseTableRequest) error

	// Find a Database Instance by uid.
	FindDatabaseInstanceByUid(ctx context.Context, request FindDatabaseInstanceByUidRequest) (*DatabaseInstance, error)

	// Generates a credential that can be used to access database instances.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Generate a database credential].
	//
	// [Generate a database credential]: https://docs.databricks.com/api/postgres/v1/generate-database-credential
	GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error)

	// Get a Database Catalog.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Get a catalog].
	//
	// [Get a catalog]: https://docs.databricks.com/api/postgres/v1/get-catalog
	GetDatabaseCatalog(ctx context.Context, request GetDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Get a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Get a project].
	//
	// [Get a project]: https://docs.databricks.com/api/postgres/v1/get-project
	GetDatabaseInstance(ctx context.Context, request GetDatabaseInstanceRequest) (*DatabaseInstance, error)

	// Gets a role for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Get a role].
	//
	// [Get a role]: https://docs.databricks.com/api/postgres/v1/get-role
	GetDatabaseInstanceRole(ctx context.Context, request GetDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error)

	// Get a Database Table.
	GetDatabaseTable(ctx context.Context, request GetDatabaseTableRequest) (*DatabaseTable, error)

	// Get a Synced Database Table.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [Get a synced table].
	//
	// [Get a synced table]: https://docs.databricks.com/api/postgres/v1/get-synced-table
	GetSyncedDatabaseTable(ctx context.Context, request GetSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	ListDatabaseCatalogs(ctx context.Context, request ListDatabaseCatalogsRequest) (*ListDatabaseCatalogsResponse, error)

	// List roles for a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [List roles].
	//
	// [List roles]: https://docs.databricks.com/api/postgres/v1/list-roles
	ListDatabaseInstanceRoles(ctx context.Context, request ListDatabaseInstanceRolesRequest) (*ListDatabaseInstanceRolesResponse, error)

	// List Database Instances.
	//
	// For new workloads, Databricks recommends the Postgres API equivalent:
	// [List projects].
	//
	// [List projects]: https://docs.databricks.com/api/postgres/v1/list-projects
	ListDatabaseInstances(ctx context.Context, request ListDatabaseInstancesRequest) (*ListDatabaseInstancesResponse, error)

	// List synced database tables in a Database Instance.
	ListSyncedDatabaseTables(ctx context.Context, request ListSyncedDatabaseTablesRequest) (*ListSyncedDatabaseTablesResponse, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	UpdateDatabaseCatalog(ctx context.Context, request UpdateDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Update a Database Instance.
	//
	// For new workloads, Databricks recommends the Postgres API equivalents:
	// [Update a project] and [Update an endpoint].
	//
	// [Update a project]: https://docs.databricks.com/api/postgres/v1/update-project
	// [Update an endpoint]: https://docs.databricks.com/api/postgres/v1/update-endpoint
	UpdateDatabaseInstance(ctx context.Context, request UpdateDatabaseInstanceRequest) (*DatabaseInstance, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	UpdateSyncedDatabaseTable(ctx context.Context, request UpdateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)
}

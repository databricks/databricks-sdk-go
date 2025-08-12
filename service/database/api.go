// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// Database Instances provide access to a database via REST API or direct SQL.
package database

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type DatabaseInterface interface {

	// WaitGetDatabaseInstanceDatabaseAvailable repeatedly calls [DatabaseAPI.GetDatabaseInstance] and waits to reach AVAILABLE state
	WaitGetDatabaseInstanceDatabaseAvailable(ctx context.Context, name string,
		timeout time.Duration, callback func(*DatabaseInstance)) (*DatabaseInstance, error)

	// Create a Database Catalog.
	CreateDatabaseCatalog(ctx context.Context, request CreateDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Create a Database Instance.
	CreateDatabaseInstance(ctx context.Context, createDatabaseInstanceRequest CreateDatabaseInstanceRequest) (*WaitGetDatabaseInstanceDatabaseAvailable[DatabaseInstance], error)

	// Calls [DatabaseAPIInterface.CreateDatabaseInstance] and waits to reach AVAILABLE state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[DatabaseInstance](60*time.Minute) functional option.
	//
	// Deprecated: use [DatabaseAPIInterface.CreateDatabaseInstance].Get() or [DatabaseAPIInterface.WaitGetDatabaseInstanceDatabaseAvailable]
	CreateDatabaseInstanceAndWait(ctx context.Context, createDatabaseInstanceRequest CreateDatabaseInstanceRequest, options ...retries.Option[DatabaseInstance]) (*DatabaseInstance, error)

	// Create a role for a Database Instance.
	CreateDatabaseInstanceRole(ctx context.Context, request CreateDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error)

	// Create a Database Table. Useful for registering pre-existing PG tables in UC.
	// See CreateSyncedDatabaseTable for creating synced tables in PG from a source
	// table in UC.
	CreateDatabaseTable(ctx context.Context, request CreateDatabaseTableRequest) (*DatabaseTable, error)

	// Create a Synced Database Table.
	CreateSyncedDatabaseTable(ctx context.Context, request CreateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	// Delete a Database Catalog.
	DeleteDatabaseCatalog(ctx context.Context, request DeleteDatabaseCatalogRequest) error

	// Delete a Database Catalog.
	DeleteDatabaseCatalogByName(ctx context.Context, name string) error

	// Delete a Database Instance.
	DeleteDatabaseInstance(ctx context.Context, request DeleteDatabaseInstanceRequest) error

	// Delete a Database Instance.
	DeleteDatabaseInstanceByName(ctx context.Context, name string) error

	// Deletes a role for a Database Instance.
	DeleteDatabaseInstanceRole(ctx context.Context, request DeleteDatabaseInstanceRoleRequest) error

	// Deletes a role for a Database Instance.
	DeleteDatabaseInstanceRoleByInstanceNameAndName(ctx context.Context, instanceName string, name string) error

	// Delete a Database Table.
	DeleteDatabaseTable(ctx context.Context, request DeleteDatabaseTableRequest) error

	// Delete a Database Table.
	DeleteDatabaseTableByName(ctx context.Context, name string) error

	// Delete a Synced Database Table.
	DeleteSyncedDatabaseTable(ctx context.Context, request DeleteSyncedDatabaseTableRequest) error

	// Delete a Synced Database Table.
	DeleteSyncedDatabaseTableByName(ctx context.Context, name string) error

	// Find a Database Instance by uid.
	FindDatabaseInstanceByUid(ctx context.Context, request FindDatabaseInstanceByUidRequest) (*DatabaseInstance, error)

	// Generates a credential that can be used to access database instances.
	GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error)

	// Get a Database Catalog.
	GetDatabaseCatalog(ctx context.Context, request GetDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Get a Database Catalog.
	GetDatabaseCatalogByName(ctx context.Context, name string) (*DatabaseCatalog, error)

	// Get a Database Instance.
	GetDatabaseInstance(ctx context.Context, request GetDatabaseInstanceRequest) (*DatabaseInstance, error)

	// Get a Database Instance.
	GetDatabaseInstanceByName(ctx context.Context, name string) (*DatabaseInstance, error)

	// Gets a role for a Database Instance.
	GetDatabaseInstanceRole(ctx context.Context, request GetDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error)

	// Gets a role for a Database Instance.
	GetDatabaseInstanceRoleByInstanceNameAndName(ctx context.Context, instanceName string, name string) (*DatabaseInstanceRole, error)

	// Get a Database Table.
	GetDatabaseTable(ctx context.Context, request GetDatabaseTableRequest) (*DatabaseTable, error)

	// Get a Database Table.
	GetDatabaseTableByName(ctx context.Context, name string) (*DatabaseTable, error)

	// Get a Synced Database Table.
	GetSyncedDatabaseTable(ctx context.Context, request GetSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)

	// Get a Synced Database Table.
	GetSyncedDatabaseTableByName(ctx context.Context, name string) (*SyncedDatabaseTable, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabaseCatalogs(ctx context.Context, request ListDatabaseCatalogsRequest) listing.Iterator[DatabaseCatalog]

	// This API is currently unimplemented, but exposed for Terraform support.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabaseCatalogsAll(ctx context.Context, request ListDatabaseCatalogsRequest) ([]DatabaseCatalog, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	ListDatabaseCatalogsByInstanceName(ctx context.Context, instanceName string) (*ListDatabaseCatalogsResponse, error)

	// START OF PG ROLE APIs Section These APIs are marked a PUBLIC with stage <
	// PUBLIC_PREVIEW. With more recent Lakebase V2 plans, we don't plan to ever
	// advance these to PUBLIC_PREVIEW. These APIs will remain effectively
	// undocumented/UI-only and we'll aim for a new public roles API as part of V2
	// PuPr.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabaseInstanceRoles(ctx context.Context, request ListDatabaseInstanceRolesRequest) listing.Iterator[DatabaseInstanceRole]

	// START OF PG ROLE APIs Section These APIs are marked a PUBLIC with stage <
	// PUBLIC_PREVIEW. With more recent Lakebase V2 plans, we don't plan to ever
	// advance these to PUBLIC_PREVIEW. These APIs will remain effectively
	// undocumented/UI-only and we'll aim for a new public roles API as part of V2
	// PuPr.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabaseInstanceRolesAll(ctx context.Context, request ListDatabaseInstanceRolesRequest) ([]DatabaseInstanceRole, error)

	// START OF PG ROLE APIs Section These APIs are marked a PUBLIC with stage <
	// PUBLIC_PREVIEW. With more recent Lakebase V2 plans, we don't plan to ever
	// advance these to PUBLIC_PREVIEW. These APIs will remain effectively
	// undocumented/UI-only and we'll aim for a new public roles API as part of V2
	// PuPr.
	ListDatabaseInstanceRolesByInstanceName(ctx context.Context, instanceName string) (*ListDatabaseInstanceRolesResponse, error)

	// List Database Instances.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabaseInstances(ctx context.Context, request ListDatabaseInstancesRequest) listing.Iterator[DatabaseInstance]

	// List Database Instances.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListDatabaseInstancesAll(ctx context.Context, request ListDatabaseInstancesRequest) ([]DatabaseInstance, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSyncedDatabaseTables(ctx context.Context, request ListSyncedDatabaseTablesRequest) listing.Iterator[SyncedDatabaseTable]

	// This API is currently unimplemented, but exposed for Terraform support.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSyncedDatabaseTablesAll(ctx context.Context, request ListSyncedDatabaseTablesRequest) ([]SyncedDatabaseTable, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	ListSyncedDatabaseTablesByInstanceName(ctx context.Context, instanceName string) (*ListSyncedDatabaseTablesResponse, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	UpdateDatabaseCatalog(ctx context.Context, request UpdateDatabaseCatalogRequest) (*DatabaseCatalog, error)

	// Update a Database Instance.
	UpdateDatabaseInstance(ctx context.Context, request UpdateDatabaseInstanceRequest) (*DatabaseInstance, error)

	// This API is currently unimplemented, but exposed for Terraform support.
	UpdateSyncedDatabaseTable(ctx context.Context, request UpdateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error)
}

func NewDatabase(client *client.DatabricksClient) *DatabaseAPI {
	return &DatabaseAPI{
		databaseImpl: databaseImpl{
			client: client,
		},
	}
}

// Database Instances provide access to a database via REST API or direct SQL.
type DatabaseAPI struct {
	databaseImpl
}

// WaitGetDatabaseInstanceDatabaseAvailable repeatedly calls [DatabaseAPI.GetDatabaseInstance] and waits to reach AVAILABLE state
func (a *DatabaseAPI) WaitGetDatabaseInstanceDatabaseAvailable(ctx context.Context, name string,
	timeout time.Duration, callback func(*DatabaseInstance)) (*DatabaseInstance, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[DatabaseInstance](ctx, timeout, func() (*DatabaseInstance, *retries.Err) {
		databaseInstance, err := a.GetDatabaseInstance(ctx, GetDatabaseInstanceRequest{
			Name: name,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(databaseInstance)
		}
		status := databaseInstance.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case DatabaseInstanceStateAvailable: // target state
			return databaseInstance, nil
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetDatabaseInstanceDatabaseAvailable is a wrapper that calls [DatabaseAPI.WaitGetDatabaseInstanceDatabaseAvailable] and waits to reach AVAILABLE state.
type WaitGetDatabaseInstanceDatabaseAvailable[R any] struct {
	Response *R
	Name     string `json:"name"`
	Poll     func(time.Duration, func(*DatabaseInstance)) (*DatabaseInstance, error)
	callback func(*DatabaseInstance)
	timeout  time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetDatabaseInstanceDatabaseAvailable[R]) OnProgress(callback func(*DatabaseInstance)) *WaitGetDatabaseInstanceDatabaseAvailable[R] {
	w.callback = callback
	return w
}

// Get the DatabaseInstance with the default timeout of 20 minutes.
func (w *WaitGetDatabaseInstanceDatabaseAvailable[R]) Get() (*DatabaseInstance, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the DatabaseInstance with custom timeout.
func (w *WaitGetDatabaseInstanceDatabaseAvailable[R]) GetWithTimeout(timeout time.Duration) (*DatabaseInstance, error) {
	return w.Poll(timeout, w.callback)
}

// Create a Database Instance.
func (a *DatabaseAPI) CreateDatabaseInstance(ctx context.Context, createDatabaseInstanceRequest CreateDatabaseInstanceRequest) (*WaitGetDatabaseInstanceDatabaseAvailable[DatabaseInstance], error) {
	databaseInstance, err := a.databaseImpl.CreateDatabaseInstance(ctx, createDatabaseInstanceRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetDatabaseInstanceDatabaseAvailable[DatabaseInstance]{
		Response: databaseInstance,
		Name:     databaseInstance.Name,
		Poll: func(timeout time.Duration, callback func(*DatabaseInstance)) (*DatabaseInstance, error) {
			return a.WaitGetDatabaseInstanceDatabaseAvailable(ctx, databaseInstance.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [DatabaseAPI.CreateDatabaseInstance] and waits to reach AVAILABLE state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[DatabaseInstance](60*time.Minute) functional option.
//
// Deprecated: use [DatabaseAPI.CreateDatabaseInstance].Get() or [DatabaseAPI.WaitGetDatabaseInstanceDatabaseAvailable]
func (a *DatabaseAPI) CreateDatabaseInstanceAndWait(ctx context.Context, createDatabaseInstanceRequest CreateDatabaseInstanceRequest, options ...retries.Option[DatabaseInstance]) (*DatabaseInstance, error) {
	wait, err := a.CreateDatabaseInstance(ctx, createDatabaseInstanceRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[DatabaseInstance]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *DatabaseInstance) {
		for _, o := range options {
			o(&retries.Info[DatabaseInstance]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Delete a Database Catalog.
func (a *DatabaseAPI) DeleteDatabaseCatalogByName(ctx context.Context, name string) error {
	return a.databaseImpl.DeleteDatabaseCatalog(ctx, DeleteDatabaseCatalogRequest{
		Name: name,
	})
}

// Delete a Database Instance.
func (a *DatabaseAPI) DeleteDatabaseInstanceByName(ctx context.Context, name string) error {
	return a.databaseImpl.DeleteDatabaseInstance(ctx, DeleteDatabaseInstanceRequest{
		Name: name,
	})
}

// Deletes a role for a Database Instance.
func (a *DatabaseAPI) DeleteDatabaseInstanceRoleByInstanceNameAndName(ctx context.Context, instanceName string, name string) error {
	return a.databaseImpl.DeleteDatabaseInstanceRole(ctx, DeleteDatabaseInstanceRoleRequest{
		InstanceName: instanceName,
		Name:         name,
	})
}

// Delete a Database Table.
func (a *DatabaseAPI) DeleteDatabaseTableByName(ctx context.Context, name string) error {
	return a.databaseImpl.DeleteDatabaseTable(ctx, DeleteDatabaseTableRequest{
		Name: name,
	})
}

// Delete a Synced Database Table.
func (a *DatabaseAPI) DeleteSyncedDatabaseTableByName(ctx context.Context, name string) error {
	return a.databaseImpl.DeleteSyncedDatabaseTable(ctx, DeleteSyncedDatabaseTableRequest{
		Name: name,
	})
}

// Get a Database Catalog.
func (a *DatabaseAPI) GetDatabaseCatalogByName(ctx context.Context, name string) (*DatabaseCatalog, error) {
	return a.databaseImpl.GetDatabaseCatalog(ctx, GetDatabaseCatalogRequest{
		Name: name,
	})
}

// Get a Database Instance.
func (a *DatabaseAPI) GetDatabaseInstanceByName(ctx context.Context, name string) (*DatabaseInstance, error) {
	return a.databaseImpl.GetDatabaseInstance(ctx, GetDatabaseInstanceRequest{
		Name: name,
	})
}

// Gets a role for a Database Instance.
func (a *DatabaseAPI) GetDatabaseInstanceRoleByInstanceNameAndName(ctx context.Context, instanceName string, name string) (*DatabaseInstanceRole, error) {
	return a.databaseImpl.GetDatabaseInstanceRole(ctx, GetDatabaseInstanceRoleRequest{
		InstanceName: instanceName,
		Name:         name,
	})
}

// Get a Database Table.
func (a *DatabaseAPI) GetDatabaseTableByName(ctx context.Context, name string) (*DatabaseTable, error) {
	return a.databaseImpl.GetDatabaseTable(ctx, GetDatabaseTableRequest{
		Name: name,
	})
}

// Get a Synced Database Table.
func (a *DatabaseAPI) GetSyncedDatabaseTableByName(ctx context.Context, name string) (*SyncedDatabaseTable, error) {
	return a.databaseImpl.GetSyncedDatabaseTable(ctx, GetSyncedDatabaseTableRequest{
		Name: name,
	})
}

// This API is currently unimplemented, but exposed for Terraform support.
func (a *DatabaseAPI) ListDatabaseCatalogsByInstanceName(ctx context.Context, instanceName string) (*ListDatabaseCatalogsResponse, error) {
	return a.databaseImpl.internalListDatabaseCatalogs(ctx, ListDatabaseCatalogsRequest{
		InstanceName: instanceName,
	})
}

// START OF PG ROLE APIs Section These APIs are marked a PUBLIC with stage <
// PUBLIC_PREVIEW. With more recent Lakebase V2 plans, we don't plan to ever
// advance these to PUBLIC_PREVIEW. These APIs will remain effectively
// undocumented/UI-only and we'll aim for a new public roles API as part of V2
// PuPr.
func (a *DatabaseAPI) ListDatabaseInstanceRolesByInstanceName(ctx context.Context, instanceName string) (*ListDatabaseInstanceRolesResponse, error) {
	return a.databaseImpl.internalListDatabaseInstanceRoles(ctx, ListDatabaseInstanceRolesRequest{
		InstanceName: instanceName,
	})
}

// This API is currently unimplemented, but exposed for Terraform support.
func (a *DatabaseAPI) ListSyncedDatabaseTablesByInstanceName(ctx context.Context, instanceName string) (*ListSyncedDatabaseTablesResponse, error) {
	return a.databaseImpl.internalListSyncedDatabaseTables(ctx, ListSyncedDatabaseTablesRequest{
		InstanceName: instanceName,
	})
}

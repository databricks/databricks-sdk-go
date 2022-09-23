// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package warehouses

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/retries"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewWarehouses(client *client.DatabricksClient) WarehousesService {
	return &WarehousesAPI{
		client: client,
	}
}

type WarehousesAPI struct {
	client *client.DatabricksClient
}

// Creates a new SQL warehouse.
func (a *WarehousesAPI) CreateWarehouse(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	var createWarehouseResponse CreateWarehouseResponse
	path := "/api/2.0/sql/warehouses"
	err := a.client.Post(ctx, path, request, &createWarehouseResponse)
	return &createWarehouseResponse, err
}

// CreateWarehouse and wait to reach RUNNING state
func (a *WarehousesAPI) CreateWarehouseAndWait(ctx context.Context, createWarehouseRequest CreateWarehouseRequest, timeout ...time.Duration) (*GetWarehouseResponse, error) {
	createWarehouseResponse, err := a.CreateWarehouse(ctx, createWarehouseRequest)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[GetWarehouseResponse](ctx, timeout[0], func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: createWarehouseResponse.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := getWarehouseResponse.State
		statusMessage := getWarehouseResponse.Health.Summary
		switch status {
		case GetWarehouseResponseStateRunning: // target state
			return getWarehouseResponse, nil
		case GetWarehouseResponseStateStopped, GetWarehouseResponseStateDeleted:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetWarehouseResponseStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Deletes a SQL warehouse.
func (a *WarehousesAPI) DeleteWarehouse(ctx context.Context, request DeleteWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

// DeleteWarehouse and wait to reach DELETED state
func (a *WarehousesAPI) DeleteWarehouseAndWait(ctx context.Context, deleteWarehouseRequest DeleteWarehouseRequest, timeout ...time.Duration) (*GetWarehouseResponse, error) {
	err := a.DeleteWarehouse(ctx, deleteWarehouseRequest)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[GetWarehouseResponse](ctx, timeout[0], func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: deleteWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := getWarehouseResponse.State
		statusMessage := getWarehouseResponse.Health.Summary
		switch status {
		case GetWarehouseResponseStateDeleted: // target state
			return getWarehouseResponse, nil
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Deletes a SQL warehouse.
func (a *WarehousesAPI) DeleteWarehouseById(ctx context.Context, id string) error {
	return a.DeleteWarehouse(ctx, DeleteWarehouseRequest{
		Id: id,
	})
}

func (a *WarehousesAPI) DeleteWarehouseByIdAndWait(ctx context.Context, id string, timeout ...time.Duration) (*GetWarehouseResponse, error) {
	return a.DeleteWarehouseAndWait(ctx, DeleteWarehouseRequest{
		Id: id,
	}, timeout...)
}

// Edits a SQL warehouse.
func (a *WarehousesAPI) EditWarehouse(ctx context.Context, request EditWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/edit", request.Id)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// EditWarehouse and wait to reach RUNNING state
func (a *WarehousesAPI) EditWarehouseAndWait(ctx context.Context, editWarehouseRequest EditWarehouseRequest, timeout ...time.Duration) (*GetWarehouseResponse, error) {
	err := a.EditWarehouse(ctx, editWarehouseRequest)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[GetWarehouseResponse](ctx, timeout[0], func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: editWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := getWarehouseResponse.State
		statusMessage := getWarehouseResponse.Health.Summary
		switch status {
		case GetWarehouseResponseStateRunning: // target state
			return getWarehouseResponse, nil
		case GetWarehouseResponseStateStopped, GetWarehouseResponseStateDeleted:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetWarehouseResponseStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Gets the information for a single SQL warehouse.
func (a *WarehousesAPI) GetWarehouse(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error) {
	var getWarehouseResponse GetWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	err := a.client.Get(ctx, path, request, &getWarehouseResponse)
	return &getWarehouseResponse, err
}

// GetWarehouse and wait to reach RUNNING state
func (a *WarehousesAPI) GetWarehouseAndWait(ctx context.Context, getWarehouseRequest GetWarehouseRequest, timeout ...time.Duration) (*GetWarehouseResponse, error) {
	getWarehouseResponse, err := a.GetWarehouse(ctx, getWarehouseRequest)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[GetWarehouseResponse](ctx, timeout[0], func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: getWarehouseResponse.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := getWarehouseResponse.State
		statusMessage := getWarehouseResponse.Health.Summary
		switch status {
		case GetWarehouseResponseStateRunning: // target state
			return getWarehouseResponse, nil
		case GetWarehouseResponseStateStopped, GetWarehouseResponseStateDeleted:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetWarehouseResponseStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Gets the information for a single SQL warehouse.
func (a *WarehousesAPI) GetWarehouseById(ctx context.Context, id string) (*GetWarehouseResponse, error) {
	return a.GetWarehouse(ctx, GetWarehouseRequest{
		Id: id,
	})
}

func (a *WarehousesAPI) GetWarehouseByIdAndWait(ctx context.Context, id string, timeout ...time.Duration) (*GetWarehouseResponse, error) {
	return a.GetWarehouseAndWait(ctx, GetWarehouseRequest{
		Id: id,
	}, timeout...)
}

// Gets the workspace level configuration that is shared by all SQL warehouses
// in a workspace.
func (a *WarehousesAPI) GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error) {
	var getWorkspaceWarehouseConfigResponse GetWorkspaceWarehouseConfigResponse
	path := "/api/2.0/sql/config/warehouses"
	err := a.client.Get(ctx, path, nil, &getWorkspaceWarehouseConfigResponse)
	return &getWorkspaceWarehouseConfigResponse, err
}

// Lists all SQL warehouse a user has manager permissions for.
func (a *WarehousesAPI) ListWarehouses(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error) {
	var listWarehousesResponse ListWarehousesResponse
	path := "/api/2.0/sql/warehouses"
	err := a.client.Get(ctx, path, request, &listWarehousesResponse)
	return &listWarehousesResponse, err
}

// Sets the workspace level configuration that is shared by all SQL warehouses
// in a workspace.
func (a *WarehousesAPI) SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error {
	path := "/api/2.0/sql/config/warehouses"
	err := a.client.Put(ctx, path, request)
	return err
}

// Starts a SQL warehouse.
func (a *WarehousesAPI) StartWarehouse(ctx context.Context, request StartWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/start", request.Id)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// StartWarehouse and wait to reach RUNNING state
func (a *WarehousesAPI) StartWarehouseAndWait(ctx context.Context, startWarehouseRequest StartWarehouseRequest, timeout ...time.Duration) (*GetWarehouseResponse, error) {
	err := a.StartWarehouse(ctx, startWarehouseRequest)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[GetWarehouseResponse](ctx, timeout[0], func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: startWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := getWarehouseResponse.State
		statusMessage := getWarehouseResponse.Health.Summary
		switch status {
		case GetWarehouseResponseStateRunning: // target state
			return getWarehouseResponse, nil
		case GetWarehouseResponseStateStopped, GetWarehouseResponseStateDeleted:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				GetWarehouseResponseStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Stops a SQL warehouse.
func (a *WarehousesAPI) StopWarehouse(ctx context.Context, request StopWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/stop", request.Id)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// StopWarehouse and wait to reach STOPPED state
func (a *WarehousesAPI) StopWarehouseAndWait(ctx context.Context, stopWarehouseRequest StopWarehouseRequest, timeout ...time.Duration) (*GetWarehouseResponse, error) {
	err := a.StopWarehouse(ctx, stopWarehouseRequest)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[GetWarehouseResponse](ctx, timeout[0], func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: stopWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := getWarehouseResponse.State
		statusMessage := getWarehouseResponse.Health.Summary
		switch status {
		case GetWarehouseResponseStateStopped: // target state
			return getWarehouseResponse, nil
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

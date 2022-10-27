// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package warehouses

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
	"github.com/databricks/databricks-sdk-go/retries"
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

// CreateWarehouseTimeout overrides the default timeout of 20 minutes to reach RUNNING state
func CreateWarehouseTimeout(dur time.Duration) retries.Option[GetWarehouseResponse] {
	return retries.Timeout[GetWarehouseResponse](dur)
}

// CreateWarehouse and wait to reach RUNNING state
func (a *WarehousesAPI) CreateWarehouseAndWait(ctx context.Context, createWarehouseRequest CreateWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	createWarehouseResponse, err := a.CreateWarehouse(ctx, createWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: createWarehouseResponse.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    *getWarehouseResponse,
				Timeout: i.Timeout,
			})
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

// DeleteWarehouseTimeout overrides the default timeout of 20 minutes to reach DELETED state
func DeleteWarehouseTimeout(dur time.Duration) retries.Option[GetWarehouseResponse] {
	return retries.Timeout[GetWarehouseResponse](dur)
}

// DeleteWarehouse and wait to reach DELETED state
func (a *WarehousesAPI) DeleteWarehouseAndWait(ctx context.Context, deleteWarehouseRequest DeleteWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.DeleteWarehouse(ctx, deleteWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: deleteWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    *getWarehouseResponse,
				Timeout: i.Timeout,
			})
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

func (a *WarehousesAPI) DeleteWarehouseByIdAndWait(ctx context.Context, id string, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	return a.DeleteWarehouseAndWait(ctx, DeleteWarehouseRequest{
		Id: id,
	}, options...)
}

// Edits a SQL warehouse.
func (a *WarehousesAPI) EditWarehouse(ctx context.Context, request EditWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/edit", request.Id)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// EditWarehouseTimeout overrides the default timeout of 20 minutes to reach RUNNING state
func EditWarehouseTimeout(dur time.Duration) retries.Option[GetWarehouseResponse] {
	return retries.Timeout[GetWarehouseResponse](dur)
}

// EditWarehouse and wait to reach RUNNING state
func (a *WarehousesAPI) EditWarehouseAndWait(ctx context.Context, editWarehouseRequest EditWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.EditWarehouse(ctx, editWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: editWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    *getWarehouseResponse,
				Timeout: i.Timeout,
			})
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

// GetWarehouseTimeout overrides the default timeout of 20 minutes to reach RUNNING state
func GetWarehouseTimeout(dur time.Duration) retries.Option[GetWarehouseResponse] {
	return retries.Timeout[GetWarehouseResponse](dur)
}

// GetWarehouse and wait to reach RUNNING state
func (a *WarehousesAPI) GetWarehouseAndWait(ctx context.Context, getWarehouseRequest GetWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	getWarehouseResponse, err := a.GetWarehouse(ctx, getWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: getWarehouseResponse.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    *getWarehouseResponse,
				Timeout: i.Timeout,
			})
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

func (a *WarehousesAPI) GetWarehouseByIdAndWait(ctx context.Context, id string, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	return a.GetWarehouseAndWait(ctx, GetWarehouseRequest{
		Id: id,
	}, options...)
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
//
// Use ListWarehousesAll() to get all EndpointInfo instances
func (a *WarehousesAPI) ListWarehouses(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error) {
	var listWarehousesResponse ListWarehousesResponse
	path := "/api/2.0/sql/warehouses"
	err := a.client.Get(ctx, path, request, &listWarehousesResponse)
	return &listWarehousesResponse, err
}

// ListWarehousesAll returns all EndpointInfo instances. This method exists for consistency purposes.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WarehousesAPI) ListWarehousesAll(ctx context.Context, request ListWarehousesRequest) ([]EndpointInfo, error) {
	response, err := a.ListWarehouses(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Warehouses, nil
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

// StartWarehouseTimeout overrides the default timeout of 20 minutes to reach RUNNING state
func StartWarehouseTimeout(dur time.Duration) retries.Option[GetWarehouseResponse] {
	return retries.Timeout[GetWarehouseResponse](dur)
}

// StartWarehouse and wait to reach RUNNING state
func (a *WarehousesAPI) StartWarehouseAndWait(ctx context.Context, startWarehouseRequest StartWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.StartWarehouse(ctx, startWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: startWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    *getWarehouseResponse,
				Timeout: i.Timeout,
			})
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

// StopWarehouseTimeout overrides the default timeout of 20 minutes to reach STOPPED state
func StopWarehouseTimeout(dur time.Duration) retries.Option[GetWarehouseResponse] {
	return retries.Timeout[GetWarehouseResponse](dur)
}

// StopWarehouse and wait to reach STOPPED state
func (a *WarehousesAPI) StopWarehouseAndWait(ctx context.Context, stopWarehouseRequest StopWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.StopWarehouse(ctx, stopWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.GetWarehouse(ctx, GetWarehouseRequest{
			Id: stopWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    *getWarehouseResponse,
				Timeout: i.Timeout,
			})
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

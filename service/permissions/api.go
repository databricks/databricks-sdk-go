// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewPermissions(client *client.DatabricksClient) PermissionsService {
	return &PermissionsAPI{
		client: client,
	}
}

type PermissionsAPI struct {
	client *client.DatabricksClient
}

func (a *PermissionsAPI) GetObjectPermissions(ctx context.Context, request GetObjectPermissionsRequest) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Get(ctx, path, request, &objectPermissions)
	return &objectPermissions, err
}

func (a *PermissionsAPI) GetObjectPermissionsByObjectTypeAndObjectId(ctx context.Context, objectType string, objectId string) (*ObjectPermissions, error) {
	return a.GetObjectPermissions(ctx, GetObjectPermissionsRequest{
		ObjectType: objectType,
		ObjectId:   objectId,
	})
}

// <needs content>
func (a *PermissionsAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error) {
	var getPermissionLevelsResponse GetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", request.RequestObjectType, request.RequestObjectId)
	err := a.client.Get(ctx, path, request, &getPermissionLevelsResponse)
	return &getPermissionLevelsResponse, err
}

// <needs content>
func (a *PermissionsAPI) GetPermissionLevelsByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*GetPermissionLevelsResponse, error) {
	return a.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		RequestObjectType: requestObjectType,
		RequestObjectId:   requestObjectId,
	})
}

// <needs content>
func (a *PermissionsAPI) SetObjectPermissions(ctx context.Context, request SetObjectPermissions) error {
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Put(ctx, path, request)
	return err
}

// <needs content>
func (a *PermissionsAPI) UpdateObjectPermissions(ctx context.Context, request UpdateObjectPermissions) error {
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Patch(ctx, path, request)
	return err
}

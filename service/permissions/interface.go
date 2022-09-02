// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

import (
	"context"
)



type PermissionsService interface {
    
    GetObjectPermissions(ctx context.Context, getObjectPermissionsRequest GetObjectPermissionsRequest) (*ObjectPermissions, error)
    
    GetPermissionLevels(ctx context.Context, getPermissionLevelsRequest GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error)
    
    SetObjectPermissions(ctx context.Context, setObjectPermissionsRequest SetObjectPermissionsRequest) error
    
    UpdateObjectPermissions(ctx context.Context, updateObjectPermissionsRequest UpdateObjectPermissionsRequest) error
	GetPermissionLevelsByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*GetPermissionLevelsResponse, error)
	GetObjectPermissionsByObjectTypeAndObjectId(ctx context.Context, objectType string, objectId string) (*ObjectPermissions, error)
}

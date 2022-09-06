// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

import (
	"context"
)

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type PermissionsService interface {
	GetObjectPermissions(ctx context.Context, getObjectPermissionsRequest GetObjectPermissionsRequest) (*ObjectPermissions, error)

	GetObjectPermissionsByObjectTypeAndObjectId(ctx context.Context, objectType string, objectId string) (*ObjectPermissions, error)

	GetPermissionLevels(ctx context.Context, getPermissionLevelsRequest GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error)

	GetPermissionLevelsByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*GetPermissionLevelsResponse, error)

	SetObjectPermissions(ctx context.Context, setObjectPermissionsRequest SetObjectPermissionsRequest) error

	UpdateObjectPermissions(ctx context.Context, updateObjectPermissionsRequest UpdateObjectPermissionsRequest) error
}

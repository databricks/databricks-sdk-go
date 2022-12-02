// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspaceconf

import (
	"context"
)

// This API allows updating known workspace settings for advanced users.
type WorkspaceConfService interface {

	// Check configuration status
	//
	// Gets the configuration status for a workspace.
	GetStatus(ctx context.Context, request GetStatus) (*WorkspaceConf, error)

	// Enable/disable features
	//
	// Sets the configuration status for a workspace, including enabling or
	// disabling it.
	SetStatus(ctx context.Context, request WorkspaceConf) error
}

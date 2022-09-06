// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspaceconf

import (
	"context"
)

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type WorkspaceConfService interface {
	GetStatus(ctx context.Context, getStatusRequest GetStatusRequest) (*WorkspaceConf, error)

	SetStatus(ctx context.Context, workspaceConf WorkspaceConf) error
}

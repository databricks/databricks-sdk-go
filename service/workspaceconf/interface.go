// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspaceconf

import (
	"context"
	
)



type WorkspaceConfService interface {
    
    GetStatus(ctx context.Context, getStatusRequest GetStatusRequest) (*WorkspaceConf, error)
    
    SetStatus(ctx context.Context, workspaceConf WorkspaceConf) error
}

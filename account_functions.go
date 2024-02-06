package databricks

import (
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
)

type AccountClientUtilities interface {
	// GetWorkspaceClient returns a WorkspaceClient for the given workspace. The
	// workspace can be fetched by calling w.Workspaces.Get() or w.Workspaces.List().
	//
	// The config used for the workspace is identical to that used for the account,
	// except that the host is set to the workspace host, and the account ID is
	// not set.
	//
	// Example:
	//
	//	a, err := databricks.NewAccountClient()
	//	if err != nil {
	//		panic(err)
	//	}
	//	ctx := context.Background()
	//	workspaces, err := a.Workspaces.List(ctx)
	//	if err != nil {
	//		panic(err)
	//	}
	//	w, err := a.Client.GetWorkspaceClient(workspaces[0])
	//	if err != nil {
	//		panic(err)
	//	}
	//	me, err := w.CurrentUser.Me(ctx)
	GetWorkspaceClient(w provisioning.Workspace) (*WorkspaceClient, error)
}

type accountClientUtilitiesImpl struct {
	config *config.Config
}

func NewAccountClientUtilities(c *config.Config) AccountClientUtilities {
	return &accountClientUtilitiesImpl{config: c}
}

// GetWorkspaceClient returns a WorkspaceClient for the given workspace. The
// workspace can be fetched by calling w.Workspaces.Get() or w.Workspaces.List().
//
// The config used for the workspace is identical to that used for the account,
// except that the host is set to the workspace host, and the account ID is
// not set.
//
// Example:
//
//	a, err := databricks.NewAccountClient()
//	if err != nil {
//		panic(err)
//	}
//	ctx := context.Background()
//	workspaces, err := a.Workspaces.List(ctx)
//	if err != nil {
//		panic(err)
//	}
//	w, err := a.Utilities.GetWorkspaceClient(workspaces[0])
//	if err != nil {
//		panic(err)
//	}
//	me, err := w.CurrentUser.Me(ctx)
func (a accountClientUtilitiesImpl) GetWorkspaceClient(w provisioning.Workspace) (*WorkspaceClient, error) {
	host := a.config.Environment().DeploymentURL(w.DeploymentName)
	cfg, err := a.config.NewWithWorkspaceHost(host)
	if err != nil {
		return nil, err
	}
	cfg.AzureResourceID = w.AzureResourceId()
	return NewWorkspaceClient((*Config)(cfg))
}

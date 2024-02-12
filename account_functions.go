package databricks

import "github.com/databricks/databricks-sdk-go/service/provisioning"

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
//	w, err := a.GetWorkspaceClient(workspaces[0])
//	if err != nil {
//		panic(err)
//	}
//	me, err := w.CurrentUser.Me(ctx)
func (c *AccountClient) GetWorkspaceClient(ws provisioning.Workspace) (*WorkspaceClient, error) {
	host := c.Config.Environment().DeploymentURL(ws.DeploymentName)
	cfg, err := c.Config.NewWithWorkspaceHost(host)
	if err != nil {
		return nil, err
	}
	cfg.AzureResourceID = ws.AzureResourceId()
	w, err := NewWorkspaceClient((*Config)(cfg))
	if err != nil {
		return nil, err
	}
	return w, nil
}

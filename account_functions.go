package databricks

import "github.com/databricks/databricks-sdk-go/service/provisioning"

// GetWorkspaceClient returns a WorkspaceClient for the given deployment name.
// The deployment name for a workspace is present in the deploymentName field of
// the Workspace struct, which can be obtained by calling w.Workspaces.Get()
// or w.Workspaces.List().
//
// The config used for the workspace is identical to that used for the account,
// except that the host is set to the workspace host, and the account ID is
// not set.
func (c *AccountClient) GetWorkspaceClient(w provisioning.Workspace) (*WorkspaceClient, error) {
	host := c.Config.Environment().DeploymentURL(w.DeploymentName)
	cfg, err := c.Config.NewWithWorkspaceHost(host)
	if err != nil {
		return nil, err
	}
	cfg.AzureResourceID = w.AzureResourceID()
	return NewWorkspaceClient((*Config)(cfg))
}

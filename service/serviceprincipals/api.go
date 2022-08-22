// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serviceprincipals

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)


type ServiceprincipalsService interface {
    // Delete one service principal 
    DeleteServicePrincipal(ctx context.Context, deleteServicePrincipalRequest DeleteServicePrincipalRequest) error
    // Fetch information of one service principal 
    FetchServicePrincipal(ctx context.Context, fetchServicePrincipalRequest FetchServicePrincipalRequest) (*ServicePrincipal, error)
    // Get multiple service principals associated with a &lt;Workspace&gt;. 
    ListServicePrincipals(ctx context.Context, listServicePrincipalsRequest ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error)
    // Create one service principal in the &lt;Workspace&gt;. 
    NewServicePrincipal(ctx context.Context, servicePrincipal ServicePrincipal) (*ServicePrincipal, error)
    // Partially update details of one service principal. 
    PatchServicePrincipal(ctx context.Context, partialUpdate PartialUpdate) error
    // Update details of one service principal. 
    ReplaceServicePrincipal(ctx context.Context, servicePrincipal ServicePrincipal) error
	DeleteServicePrincipalById(ctx context.Context, id string) error
	FetchServicePrincipalById(ctx context.Context, id string) (*ServicePrincipal, error)
}

func New(client *client.DatabricksClient) ServiceprincipalsService {
	return &ServiceprincipalsAPI{
		client: client,
	}
}

type ServiceprincipalsAPI struct {
	client *client.DatabricksClient
}

// Delete one service principal 
func (a *ServiceprincipalsAPI) DeleteServicePrincipal(ctx context.Context, request DeleteServicePrincipalRequest) error {
	path := "/api/2.0/preview/scim/v2/ServicePrincipals/"+request.Id
	err := a.client.Delete(ctx, path, request)
	return err
}

// Fetch information of one service principal 
func (a *ServiceprincipalsAPI) FetchServicePrincipal(ctx context.Context, request FetchServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0/preview/scim/v2/ServicePrincipals/"+request.Id
	err := a.client.Get(ctx, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

// Get multiple service principals associated with a &lt;Workspace&gt;. 
func (a *ServiceprincipalsAPI) ListServicePrincipals(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	err := a.client.Get(ctx, path, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

// Create one service principal in the &lt;Workspace&gt;. 
func (a *ServiceprincipalsAPI) NewServicePrincipal(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	err := a.client.Post(ctx, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

// Partially update details of one service principal. 
func (a *ServiceprincipalsAPI) PatchServicePrincipal(ctx context.Context, request PartialUpdate) error {
	path := "/api/2.0/preview/scim/v2/ServicePrincipals/"+request.Id
	err := a.client.Patch(ctx, path, request)
	return err
}

// Update details of one service principal. 
func (a *ServiceprincipalsAPI) ReplaceServicePrincipal(ctx context.Context, request ServicePrincipal) error {
	path := "/api/2.0/preview/scim/v2/ServicePrincipals/"+request.Id
	err := a.client.Put(ctx, path, request)
	return err
}


func (a *ServiceprincipalsAPI) DeleteServicePrincipalById(ctx context.Context, id string) error {
	return a.DeleteServicePrincipal(ctx, DeleteServicePrincipalRequest{
		Id: id,
	})
}

func (a *ServiceprincipalsAPI) FetchServicePrincipalById(ctx context.Context, id string) (*ServicePrincipal, error) {
	return a.FetchServicePrincipal(ctx, FetchServicePrincipalRequest{
		Id: id,
	})
}

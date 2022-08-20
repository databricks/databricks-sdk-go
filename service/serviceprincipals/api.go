// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serviceprincipals

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func New(cfg *databricks.Config) ServiceprincipalsService {
	return &ServiceprincipalsAPI{
		client: client.New(cfg),
	}
}


type ServiceprincipalsService interface {
    // Delete one service principal 
    DeleteServicePrincipal(ctx context.Context, deleteServicePrincipalRequest DeleteServicePrincipalRequest) error
    // Fetch information of one service principal 
    FetchServicePrincipal(ctx context.Context, fetchServicePrincipalRequest FetchServicePrincipalRequest) (*ServicePrincipal, error)
    // Get multiple service principals associated with a &lt;Workspace&gt;. 
    ListServicePrincipals(ctx context.Context, listServicePrincipalsRequest ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error)
    // Create one service principal in the &lt;Workspace&gt;. 
    NewServicePrincipal(ctx context.Context, createServicePrincipalRequest CreateServicePrincipalRequest) (*ServicePrincipal, error)
    // Update details of one service principal. 
    UpdateServicePrincipal(ctx context.Context, updateDisplayNameRequest UpdateDisplayNameRequest) error
	FetchServicePrincipalByServicePrincipalId(ctx context.Context, servicePrincipalId string) (*ServicePrincipal, error)
	DeleteServicePrincipalByServicePrincipalId(ctx context.Context, servicePrincipalId string) error
}

type ServiceprincipalsAPI struct {
	client *client.DatabricksClient
}

// Delete one service principal 
func (a *ServiceprincipalsAPI) DeleteServicePrincipal(ctx context.Context, in DeleteServicePrincipalRequest) error {
	
	err := a.client.Delete(ctx, "/scim/v2/ServicePrincipals/{service_principal_id}", in)
	return err
}

// Fetch information of one service principal 
func (a *ServiceprincipalsAPI) FetchServicePrincipal(ctx context.Context, in FetchServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	err := a.client.Get(ctx, "/scim/v2/ServicePrincipals/{service_principal_id}", in, &servicePrincipal)
	return &servicePrincipal, err
}

// Get multiple service principals associated with a &lt;Workspace&gt;. 
func (a *ServiceprincipalsAPI) ListServicePrincipals(ctx context.Context, in ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	err := a.client.Get(ctx, "/scim/v2/ServicePrincipals", in, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

// Create one service principal in the &lt;Workspace&gt;. 
func (a *ServiceprincipalsAPI) NewServicePrincipal(ctx context.Context, in CreateServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	err := a.client.Post(ctx, "/scim/v2/ServicePrincipals", in, &servicePrincipal)
	return &servicePrincipal, err
}

// Update details of one service principal. 
func (a *ServiceprincipalsAPI) UpdateServicePrincipal(ctx context.Context, in UpdateDisplayNameRequest) error {
	
	err := a.client.Patch(ctx, "/scim/v2/ServicePrincipals/{service_principal_id}", in)
	return err
}


func (a *ServiceprincipalsAPI) FetchServicePrincipalByServicePrincipalId(ctx context.Context, servicePrincipalId string) (*ServicePrincipal, error) {
	return a.FetchServicePrincipal(ctx, FetchServicePrincipalRequest{
		ServicePrincipalId: servicePrincipalId,
	})
}

func (a *ServiceprincipalsAPI) DeleteServicePrincipalByServicePrincipalId(ctx context.Context, servicePrincipalId string) error {
	return a.DeleteServicePrincipal(ctx, DeleteServicePrincipalRequest{
		ServicePrincipalId: servicePrincipalId,
	})
}

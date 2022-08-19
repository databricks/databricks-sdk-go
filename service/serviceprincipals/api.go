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

// 
type ServiceprincipalsService interface {
	
	// Delete one service principal
	DeleteServicePrincipal(ctx context.Context, deleteServicePrincipalRequest DeleteServicePrincipalRequest) error
	
	// Fetch information of one service principal
	FetchServicePrincipal(ctx context.Context, fetchServicePrincipalRequest FetchServicePrincipalRequest) (*ServicePrincipal, error)
	
	// Get multiple service principals associated with a &lt;Workspace&gt;.
	ListServicePrincipals(ctx context.Context, listServicePrincipalsRequest ListServicePrincipalsRequest) (*ListServicePrincipalsResponse, error)
	
	// Create one service principal in the &lt;Workspace&gt;.
	NewServicePrincipal(ctx context.Context, newServicePrincipalRequest NewServicePrincipalRequest) (*ServicePrincipal, error)
	
	// Update details of one service principal.
	UpdateServicePrincipal(ctx context.Context, updateDisplayNameRequest UpdateDisplayNameRequest) error
	
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
func (a *ServiceprincipalsAPI) ListServicePrincipals(ctx context.Context, in ListServicePrincipalsRequest) (*ListServicePrincipalsResponse, error) {
	var listServicePrincipalsResponse ListServicePrincipalsResponse
	err := a.client.Get(ctx, "/scim/v2/ServicePrincipals", in, &listServicePrincipalsResponse)
	return &listServicePrincipalsResponse, err
}

// Create one service principal in the &lt;Workspace&gt;.
func (a *ServiceprincipalsAPI) NewServicePrincipal(ctx context.Context, in NewServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	err := a.client.Post(ctx, "/scim/v2/ServicePrincipals", in, &servicePrincipal)
	return &servicePrincipal, err
}

// Update details of one service principal.
func (a *ServiceprincipalsAPI) UpdateServicePrincipal(ctx context.Context, in UpdateDisplayNameRequest) error {
	
	err := a.client.Patch(ctx, "/scim/v2/ServicePrincipals/{service_principal_id}", in)
	return err
}

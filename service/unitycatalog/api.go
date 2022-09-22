// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package unitycatalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewCatalogs(client *client.DatabricksClient) CatalogsService {
	return &CatalogsAPI{
		client: client,
	}
}

type CatalogsAPI struct {
	client *client.DatabricksClient
}

// Create a catalog
//
// Creates a new catalog instance in the parent Metastore if the caller is a
// Metastore admin or has the CREATE CATALOG privilege.
func (a *CatalogsAPI) Create(ctx context.Context, request CreateCatalog) (*CreateCatalogResponse, error) {
	var createCatalogResponse CreateCatalogResponse
	path := "/api/2.1/unity-catalog/catalogs"
	err := a.client.Post(ctx, path, request, &createCatalogResponse)
	return &createCatalogResponse, err
}

// Delete a catalog
//
// Deletes the catalog that matches the supplied name. The caller must be a
// Metastore admin or the owner of the catalog.
func (a *CatalogsAPI) DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a catalog
//
// Deletes the catalog that matches the supplied name. The caller must be a
// Metastore admin or the owner of the catalog.
func (a *CatalogsAPI) DeleteCatalogByName(ctx context.Context, name string) error {
	return a.DeleteCatalog(ctx, DeleteCatalogRequest{
		Name: name,
	})
}

// Get a catalog
//
// Gets an array of all catalogs in the current Metastore for which the user is
// an admin or Catalog owner, or has the USAGE privilege set for their account.
func (a *CatalogsAPI) GetCatalog(ctx context.Context, request GetCatalogRequest) (*GetCatalogResponse, error) {
	var getCatalogResponse GetCatalogResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getCatalogResponse)
	return &getCatalogResponse, err
}

// Get a catalog
//
// Gets an array of all catalogs in the current Metastore for which the user is
// an admin or Catalog owner, or has the USAGE privilege set for their account.
func (a *CatalogsAPI) GetCatalogByName(ctx context.Context, name string) (*GetCatalogResponse, error) {
	return a.GetCatalog(ctx, GetCatalogRequest{
		Name: name,
	})
}

// List catalogs
//
// Gets an array of External Locations (ExternalLocationInfo objects) from the
// Metastore. The caller must be a Metastore admin, is the owner of the External
// Location, or has privileges to access the External Location.
func (a *CatalogsAPI) List(ctx context.Context) (*ListCatalogsResponse, error) {
	var listCatalogsResponse ListCatalogsResponse
	path := "/api/2.1/unity-catalog/catalogs"
	err := a.client.Get(ctx, path, nil, &listCatalogsResponse)
	return &listCatalogsResponse, err
}

func (a *CatalogsAPI) ListAll(ctx context.Context) ([]CatalogInfo, error) {
	response, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Catalogs, nil
}

// Update a catalog
//
// Updates the catalog that matches the supplied name. The caller must be either
// the owner of the catalog, or a Metastore admin (when changing the owner field
// of the catalog).
func (a *CatalogsAPI) Update(ctx context.Context, request UpdateCatalog) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewExternalLocations(client *client.DatabricksClient) ExternalLocationsService {
	return &ExternalLocationsAPI{
		client: client,
	}
}

type ExternalLocationsAPI struct {
	client *client.DatabricksClient
}

// Create an external location
//
// Creates a new External Location entry in the Metastore. The caller must be a
// Metastore admin or have the CREATE EXTERNAL LOCATION privilege on the
// Metastore.
func (a *ExternalLocationsAPI) Create(ctx context.Context, request CreateExternalLocation) (*CreateExternalLocationResponse, error) {
	var createExternalLocationResponse CreateExternalLocationResponse
	path := "/api/2.1/unity-catalog/external-locations"
	err := a.client.Post(ctx, path, request, &createExternalLocationResponse)
	return &createExternalLocationResponse, err
}

// Delete an external location
//
// Deletes the specified external location from the Metastore. The caller must
// be the owner of the external location.
func (a *ExternalLocationsAPI) DeleteExternalLocation(ctx context.Context, request DeleteExternalLocationRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete an external location
//
// Deletes the specified external location from the Metastore. The caller must
// be the owner of the external location.
func (a *ExternalLocationsAPI) DeleteExternalLocationByName(ctx context.Context, name string) error {
	return a.DeleteExternalLocation(ctx, DeleteExternalLocationRequest{
		Name: name,
	})
}

// Get an external location
//
// Gets an external location from the Metastore. The caller must be either a
// Metastore admin, the owner of the external location, or has an appropriate
// privilege level on the Metastore.
func (a *ExternalLocationsAPI) GetExternalLocation(ctx context.Context, request GetExternalLocationRequest) (*GetExternalLocationResponse, error) {
	var getExternalLocationResponse GetExternalLocationResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getExternalLocationResponse)
	return &getExternalLocationResponse, err
}

// Get an external location
//
// Gets an external location from the Metastore. The caller must be either a
// Metastore admin, the owner of the external location, or has an appropriate
// privilege level on the Metastore.
func (a *ExternalLocationsAPI) GetExternalLocationByName(ctx context.Context, name string) (*GetExternalLocationResponse, error) {
	return a.GetExternalLocation(ctx, GetExternalLocationRequest{
		Name: name,
	})
}

// List external locations
//
// Gets an array of External Locations (ExternalLocationInfo objects) from the
// Metastore. The caller must be a Metastore admin, is the owner of the external
// location, or has privileges to access the external location.
func (a *ExternalLocationsAPI) List(ctx context.Context) (*ListExternalLocationsResponse, error) {
	var listExternalLocationsResponse ListExternalLocationsResponse
	path := "/api/2.1/unity-catalog/external-locations"
	err := a.client.Get(ctx, path, nil, &listExternalLocationsResponse)
	return &listExternalLocationsResponse, err
}

func (a *ExternalLocationsAPI) ListAll(ctx context.Context) ([]ExternalLocationInfo, error) {
	response, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.ExternalLocations, nil
}

// Update an external location
//
// Updates an external location in the Metastore. The caller must be the owner
// of the externa location, or be a Metastore admin. In the second case, the
// admin can only update the name of the external location.
func (a *ExternalLocationsAPI) Update(ctx context.Context, request UpdateExternalLocation) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewGrants(client *client.DatabricksClient) GrantsService {
	return &GrantsAPI{
		client: client,
	}
}

type GrantsAPI struct {
	client *client.DatabricksClient
}

// Update permissions
//
// Updates the permissions for a Securable type.
func (a *GrantsAPI) UpdatePermissions(ctx context.Context, request UpdatePermissions) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	err := a.client.Patch(ctx, path, request)
	return err
}

// Get permissions
//
// Gets the permissions for a Securable type.
func (a *GrantsAPI) GetPermissions(ctx context.Context, request GetPermissionsRequest) (*GetPermissionsResponse, error) {
	var getPermissionsResponse GetPermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	err := a.client.Get(ctx, path, request, &getPermissionsResponse)
	return &getPermissionsResponse, err
}

// Get permissions
//
// Gets the permissions for a Securable type.
func (a *GrantsAPI) GetPermissionsBySecurableTypeAndFullName(ctx context.Context, securableType string, fullName string) (*GetPermissionsResponse, error) {
	return a.GetPermissions(ctx, GetPermissionsRequest{
		SecurableType: securableType,
		FullName:      fullName,
	})
}

func NewMetastores(client *client.DatabricksClient) MetastoresService {
	return &MetastoresAPI{
		client: client,
	}
}

type MetastoresAPI struct {
	client *client.DatabricksClient
}

// Create a Metastore
//
// Creates a new Metastore based on a provided name and storage root path.
func (a *MetastoresAPI) Create(ctx context.Context, request CreateMetastore) (*CreateMetastoreResponse, error) {
	var createMetastoreResponse CreateMetastoreResponse
	path := "/api/2.1/unity-catalog/metastores"
	err := a.client.Post(ctx, path, request, &createMetastoreResponse)
	return &createMetastoreResponse, err
}

// Create an assignment
//
// Creates a new Metastore assignment. If an assignment for the same
// __workspace_id__ exists, it will be overwritten by the new __metastore_id__
// and __default_catalog_name__. The caller must be an account admin.
func (a *MetastoresAPI) CreateMetastoreAssignment(ctx context.Context, request CreateMetastoreAssignment) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Put(ctx, path, request)
	return err
}

// Delete a Metastore
//
// Deletes a Metastore. The caller must be a Metastore admin.
func (a *MetastoresAPI) DeleteMetastore(ctx context.Context, request DeleteMetastoreRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a Metastore
//
// Deletes a Metastore. The caller must be a Metastore admin.
func (a *MetastoresAPI) DeleteMetastoreById(ctx context.Context, id string) error {
	return a.DeleteMetastore(ctx, DeleteMetastoreRequest{
		Id: id,
	})
}

// Delete an assignment
//
// Deletes a Metastore assignment. The caller must be an account administrator.
func (a *MetastoresAPI) DeleteMetastoreAssignment(ctx context.Context, request DeleteMetastoreAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete an assignment
//
// Deletes a Metastore assignment. The caller must be an account administrator.
func (a *MetastoresAPI) DeleteMetastoreAssignmentByWorkspaceId(ctx context.Context, workspaceId int) error {
	return a.DeleteMetastoreAssignment(ctx, DeleteMetastoreAssignmentRequest{
		WorkspaceId: workspaceId,
	})
}

// Get a Metastore
//
// Gets a Metastore that matches the supplied ID. The caller must be a Metastore
// admin to retrieve this info.
func (a *MetastoresAPI) GetMetastore(ctx context.Context, request GetMetastoreRequest) (*GetMetastoreResponse, error) {
	var getMetastoreResponse GetMetastoreResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Get(ctx, path, request, &getMetastoreResponse)
	return &getMetastoreResponse, err
}

// Get a Metastore
//
// Gets a Metastore that matches the supplied ID. The caller must be a Metastore
// admin to retrieve this info.
func (a *MetastoresAPI) GetMetastoreById(ctx context.Context, id string) (*GetMetastoreResponse, error) {
	return a.GetMetastore(ctx, GetMetastoreRequest{
		Id: id,
	})
}

// Get a summary
//
// Gets information about a Metastore. This summary includes the storage
// credential, the cloud vendor, the cloud region, and the global Metastore ID.
func (a *MetastoresAPI) GetMetastoreSummary(ctx context.Context) (*GetMetastoreSummaryResponse, error) {
	var getMetastoreSummaryResponse GetMetastoreSummaryResponse
	path := "/api/2.1/unity-catalog/metastore_summary"
	err := a.client.Get(ctx, path, nil, &getMetastoreSummaryResponse)
	return &getMetastoreSummaryResponse, err
}

// List Metastores
//
// Gets an array of the available Metastores (as MetastoreInfo objects). The
// caller must be an admin to retrieve this info.
func (a *MetastoresAPI) List(ctx context.Context) (*ListMetastoresResponse, error) {
	var listMetastoresResponse ListMetastoresResponse
	path := "/api/2.1/unity-catalog/metastores"
	err := a.client.Get(ctx, path, nil, &listMetastoresResponse)
	return &listMetastoresResponse, err
}

func (a *MetastoresAPI) ListAll(ctx context.Context) ([]MetastoreInfo, error) {
	response, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Metastores, nil
}

// Update a Metastore
//
// Updates information for a specific Metastore. The caller must be a Metastore
// admin.
func (a *MetastoresAPI) Update(ctx context.Context, request UpdateMetastore) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}

// Update an assignment
//
// Updates a Metastore assignment. This operation can be used to update
// __metastore_id__ or __default_catalog_name__ for a specified Workspace, if
// the Workspace is already assigned a Metastore. The caller must be an account
// admin to update __metastore_id__; otherwise, the caller can be a Workspace
// admin.
func (a *MetastoresAPI) UpdateMetastoreAssignment(ctx context.Context, request UpdateMetastoreAssignment) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewProviders(client *client.DatabricksClient) ProvidersService {
	return &ProvidersAPI{
		client: client,
	}
}

type ProvidersAPI struct {
	client *client.DatabricksClient
}

// Create an auth provider
//
// Creates a new authentication provider minimally based on a name and
// authentication type. The caller must be an admin on the Metastore.
func (a *ProvidersAPI) Create(ctx context.Context, request CreateProvider) (*CreateProviderResponse, error) {
	var createProviderResponse CreateProviderResponse
	path := "/api/2.1/unity-catalog/providers"
	err := a.client.Post(ctx, path, request, &createProviderResponse)
	return &createProviderResponse, err
}

// Delete a provider
//
// Deletes an authentication provider, if the caller is a Metastore admin or is
// the owner of the provider.
func (a *ProvidersAPI) DeleteProvider(ctx context.Context, request DeleteProviderRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a provider
//
// Deletes an authentication provider, if the caller is a Metastore admin or is
// the owner of the provider.
func (a *ProvidersAPI) DeleteProviderByName(ctx context.Context, name string) error {
	return a.DeleteProvider(ctx, DeleteProviderRequest{
		Name: name,
	})
}

// Get a provider
//
// Gets a specific authentication provider. The caller must supply the name of
// the provider, and must either be a Metastore admin or the owner of the
// provider.
func (a *ProvidersAPI) GetProvider(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	var getProviderResponse GetProviderResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getProviderResponse)
	return &getProviderResponse, err
}

// Get a provider
//
// Gets a specific authentication provider. The caller must supply the name of
// the provider, and must either be a Metastore admin or the owner of the
// provider.
func (a *ProvidersAPI) GetProviderByName(ctx context.Context, name string) (*GetProviderResponse, error) {
	return a.GetProvider(ctx, GetProviderRequest{
		Name: name,
	})
}

// List providers
//
// Gets an array of available authentication providers. The caller must either
// be a Metastore admin or the owner of the providers. Providers not owned by
// the caller are not included in the response.
func (a *ProvidersAPI) List(ctx context.Context) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.1/unity-catalog/providers"
	err := a.client.Get(ctx, path, nil, &listProvidersResponse)
	return &listProvidersResponse, err
}

func (a *ProvidersAPI) ListAll(ctx context.Context) ([]ProviderInfo, error) {
	response, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Providers, nil
}

// List shares
//
// Gets an array of all shares within the Metastore where:
//
// * the caller is a Metastore admin, or * the caller is the owner.
func (a *ProvidersAPI) ListShares(ctx context.Context, request ListSharesRequest) (*ListProviderSharesResponse, error) {
	var listProviderSharesResponse ListProviderSharesResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v/shares", request.Name)
	err := a.client.Get(ctx, path, request, &listProviderSharesResponse)
	return &listProviderSharesResponse, err
}

// List shares
//
// Gets an array of all shares within the Metastore where:
//
// * the caller is a Metastore admin, or * the caller is the owner.
func (a *ProvidersAPI) ListSharesByName(ctx context.Context, name string) (*ListProviderSharesResponse, error) {
	return a.ListShares(ctx, ListSharesRequest{
		Name: name,
	})
}

// Update a provider
//
// Updates the information for an authentication provider, if the caller is a
// Metastore admin or is the owner of the provider. If the update changes the
// provider name, the caller must be both a Metastore admin and the owner of the
// provider.
func (a *ProvidersAPI) Update(ctx context.Context, request UpdateProvider) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewRecipientActivation(client *client.DatabricksClient) RecipientActivationService {
	return &RecipientActivationAPI{
		client: client,
	}
}

type RecipientActivationAPI struct {
	client *client.DatabricksClient
}

// Get a share activation URL
//
// Gets information about an Activation URL.
func (a *RecipientActivationAPI) GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation_info/%v", request.ActivationUrl)
	err := a.client.Get(ctx, path, request, nil)
	return err
}

// Get a share activation URL
//
// Gets information about an Activation URL.
func (a *RecipientActivationAPI) GetActivationUrlInfoByActivationUrl(ctx context.Context, activationUrl string) error {
	return a.GetActivationUrlInfo(ctx, GetActivationUrlInfoRequest{
		ActivationUrl: activationUrl,
	})
}

// Get an access token
//
// RPC to retrieve access token with an activation token. This is a public API
// without any authentication.
func (a *RecipientActivationAPI) RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error) {
	var retrieveTokenResponse RetrieveTokenResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation/%v", request.ActivationUrl)
	err := a.client.Get(ctx, path, request, &retrieveTokenResponse)
	return &retrieveTokenResponse, err
}

// Get an access token
//
// RPC to retrieve access token with an activation token. This is a public API
// without any authentication.
func (a *RecipientActivationAPI) RetrieveTokenByActivationUrl(ctx context.Context, activationUrl string) (*RetrieveTokenResponse, error) {
	return a.RetrieveToken(ctx, RetrieveTokenRequest{
		ActivationUrl: activationUrl,
	})
}

func NewRecipients(client *client.DatabricksClient) RecipientsService {
	return &RecipientsAPI{
		client: client,
	}
}

type RecipientsAPI struct {
	client *client.DatabricksClient
}

// Create a share recipient
//
// Creates a new recipient with the delta sharing authentication type in the
// Metastore. The caller must be a Metastore admin or has the CREATE RECIPIENT
// privilege on the Metastore.
func (a *RecipientsAPI) Create(ctx context.Context, request CreateRecipient) (*CreateRecipientResponse, error) {
	var createRecipientResponse CreateRecipientResponse
	path := "/api/2.1/unity-catalog/recipients"
	err := a.client.Post(ctx, path, request, &createRecipientResponse)
	return &createRecipientResponse, err
}

// Delete a share recipient
//
// Deletes the specified recipient from the Metastore. The caller must be the
// owner of the recipient.
func (a *RecipientsAPI) DeleteRecipient(ctx context.Context, request DeleteRecipientRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a share recipient
//
// Deletes the specified recipient from the Metastore. The caller must be the
// owner of the recipient.
func (a *RecipientsAPI) DeleteRecipientByName(ctx context.Context, name string) error {
	return a.DeleteRecipient(ctx, DeleteRecipientRequest{
		Name: name,
	})
}

// Get a share recipient
//
// Gets a share recipient from the Metastore if:
//
// * the caller is the owner of the share recipient, or: * is a Metastore admin
func (a *RecipientsAPI) GetRecipient(ctx context.Context, request GetRecipientRequest) (*GetRecipientResponse, error) {
	var getRecipientResponse GetRecipientResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getRecipientResponse)
	return &getRecipientResponse, err
}

// Get a share recipient
//
// Gets a share recipient from the Metastore if:
//
// * the caller is the owner of the share recipient, or: * is a Metastore admin
func (a *RecipientsAPI) GetRecipientByName(ctx context.Context, name string) (*GetRecipientResponse, error) {
	return a.GetRecipient(ctx, GetRecipientRequest{
		Name: name,
	})
}

// Get share permissions
//
// Gets the share permissions for the specified Recipient. The caller must be a
// Metastore admin or the owner of the Recipient.
func (a *RecipientsAPI) GetRecipientSharePermissions(ctx context.Context, request GetRecipientSharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error) {
	var getRecipientSharePermissionsResponse GetRecipientSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/share-permissions", request.Name)
	err := a.client.Get(ctx, path, request, &getRecipientSharePermissionsResponse)
	return &getRecipientSharePermissionsResponse, err
}

// Get share permissions
//
// Gets the share permissions for the specified Recipient. The caller must be a
// Metastore admin or the owner of the Recipient.
func (a *RecipientsAPI) GetRecipientSharePermissionsByName(ctx context.Context, name string) (*GetRecipientSharePermissionsResponse, error) {
	return a.GetRecipientSharePermissions(ctx, GetRecipientSharePermissionsRequest{
		Name: name,
	})
}

// List share recipients
//
// Gets an array of all share recipients within the current Metastore where:
//
// * the caller is a Metastore admin, or * the caller is the owner.
func (a *RecipientsAPI) List(ctx context.Context) (*ListRecipientsResponse, error) {
	var listRecipientsResponse ListRecipientsResponse
	path := "/api/2.1/unity-catalog/recipients"
	err := a.client.Get(ctx, path, nil, &listRecipientsResponse)
	return &listRecipientsResponse, err
}

func (a *RecipientsAPI) ListAll(ctx context.Context) ([]RecipientInfo, error) {
	response, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Recipients, nil
}

// Rotate a token
//
// Refreshes the specified recipient's delta sharing authentication token with
// the provided token info. The caller must be the owner of the recipient.
func (a *RecipientsAPI) RotateRecipientToken(ctx context.Context, request RotateRecipientToken) (*RotateRecipientTokenResponse, error) {
	var rotateRecipientTokenResponse RotateRecipientTokenResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/rotate-token", request.Name)
	err := a.client.Post(ctx, path, request, &rotateRecipientTokenResponse)
	return &rotateRecipientTokenResponse, err
}

// Update a share recipient
//
// Updates an existing recipient in the Metastore. The caller must be a
// Metastore admin or the owner of the recipient. If the recipient name will be
// updated, the user must be both a Metastore admin and the owner of the
// recipient.
func (a *RecipientsAPI) Update(ctx context.Context, request UpdateRecipient) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewSchemas(client *client.DatabricksClient) SchemasService {
	return &SchemasAPI{
		client: client,
	}
}

type SchemasAPI struct {
	client *client.DatabricksClient
}

// Create a schema
//
// Creates a new schema for catalog in the Metatastore. The caller must be a
// Metastore admin, or have the CREATE privilege in the parentcatalog.
func (a *SchemasAPI) Create(ctx context.Context, request CreateSchema) (*CreateSchemaResponse, error) {
	var createSchemaResponse CreateSchemaResponse
	path := "/api/2.1/unity-catalog/schemas"
	err := a.client.Post(ctx, path, request, &createSchemaResponse)
	return &createSchemaResponse, err
}

// Delete a schema
//
// Deletes the specified schema from the parent catalog. The caller must be the
// owner of the schema or an owner of the parent catalog.
func (a *SchemasAPI) DeleteSchema(ctx context.Context, request DeleteSchemaRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a schema
//
// Deletes the specified schema from the parent catalog. The caller must be the
// owner of the schema or an owner of the parent catalog.
func (a *SchemasAPI) DeleteSchemaByFullName(ctx context.Context, fullName string) error {
	return a.DeleteSchema(ctx, DeleteSchemaRequest{
		FullName: fullName,
	})
}

// Get a schema
//
// Gets the specified schema for a catalog in the Metastore. The caller must be
// a Metastore admin, the owner of the schema, or a user that has the USAGE
// privilege on the schema.
func (a *SchemasAPI) GetSchema(ctx context.Context, request GetSchemaRequest) (*GetSchemaResponse, error) {
	var getSchemaResponse GetSchemaResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	err := a.client.Get(ctx, path, request, &getSchemaResponse)
	return &getSchemaResponse, err
}

// Get a schema
//
// Gets the specified schema for a catalog in the Metastore. The caller must be
// a Metastore admin, the owner of the schema, or a user that has the USAGE
// privilege on the schema.
func (a *SchemasAPI) GetSchemaByFullName(ctx context.Context, fullName string) (*GetSchemaResponse, error) {
	return a.GetSchema(ctx, GetSchemaRequest{
		FullName: fullName,
	})
}

// List schemas
//
// Gets an array of schemas for catalog in the Metastore. If the caller is the
// Metastore admin or the owner of the parent catalog, all schemas for the
// catalog will be retrieved. Otherwise, only schemas owned by the caller (or
// for which the caller has the USAGE privilege) will be retrieved.
func (a *SchemasAPI) List(ctx context.Context, request ListRequest) (*ListSchemasResponse, error) {
	var listSchemasResponse ListSchemasResponse
	path := "/api/2.1/unity-catalog/schemas"
	err := a.client.Get(ctx, path, request, &listSchemasResponse)
	return &listSchemasResponse, err
}

func (a *SchemasAPI) ListAll(ctx context.Context, request ListRequest) ([]SchemaInfo, error) {
	response, err := a.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Schemas, nil
}

// Update a schema
//
// Updates a schema for a catalog. The caller must be the owner of the schema.
// If the caller is a Metastore admin, only the __owner__ field can be changed
// in the update. If the __name__ field must be updated, the caller must be a
// Metastore admin or have the CREATE privilege on the parent catalog.
func (a *SchemasAPI) Update(ctx context.Context, request UpdateSchema) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewShares(client *client.DatabricksClient) SharesService {
	return &SharesAPI{
		client: client,
	}
}

type SharesAPI struct {
	client *client.DatabricksClient
}

// Create a share
//
// Creates a new share for data objects. Data objects can be added at this time
// or after creation with **update**. The caller must be a Metastore admin or
// have the CREATE SHARE privilege on the Metastore.
func (a *SharesAPI) Create(ctx context.Context, request CreateShare) (*CreateShareResponse, error) {
	var createShareResponse CreateShareResponse
	path := "/api/2.1/unity-catalog/shares"
	err := a.client.Post(ctx, path, request, &createShareResponse)
	return &createShareResponse, err
}

// Delete a share
//
// Deletes a data object share from the Metastore. The caller must be an owner
// of the share.
func (a *SharesAPI) DeleteShare(ctx context.Context, request DeleteShareRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a share
//
// Deletes a data object share from the Metastore. The caller must be an owner
// of the share.
func (a *SharesAPI) DeleteShareByName(ctx context.Context, name string) error {
	return a.DeleteShare(ctx, DeleteShareRequest{
		Name: name,
	})
}

// Get a share
//
// Gets a data object share from the Metastore. The caller must be a Metastore
// admin or the owner of the share.
func (a *SharesAPI) GetShare(ctx context.Context, request GetShareRequest) (*GetShareResponse, error) {
	var getShareResponse GetShareResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getShareResponse)
	return &getShareResponse, err
}

// Get a share
//
// Gets a data object share from the Metastore. The caller must be a Metastore
// admin or the owner of the share.
func (a *SharesAPI) GetShareByName(ctx context.Context, name string) (*GetShareResponse, error) {
	return a.GetShare(ctx, GetShareRequest{
		Name: name,
	})
}

// Get permissions
//
// Gets the permissions for a data share from the Metastore. The caller must be
// a Metastore admin or the owner of the share.
func (a *SharesAPI) GetSharePermissions(ctx context.Context, request GetSharePermissionsRequest) (*GetSharePermissionsResponse, error) {
	var getSharePermissionsResponse GetSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	err := a.client.Get(ctx, path, request, &getSharePermissionsResponse)
	return &getSharePermissionsResponse, err
}

// Get permissions
//
// Gets the permissions for a data share from the Metastore. The caller must be
// a Metastore admin or the owner of the share.
func (a *SharesAPI) GetSharePermissionsByName(ctx context.Context, name string) (*GetSharePermissionsResponse, error) {
	return a.GetSharePermissions(ctx, GetSharePermissionsRequest{
		Name: name,
	})
}

// List shares
//
// Gets an array of data object shares from the Metastore. The caller must be a
// Metastore admin or the owner of the share.
func (a *SharesAPI) List(ctx context.Context) (*ListSharesResponse, error) {
	var listSharesResponse ListSharesResponse
	path := "/api/2.1/unity-catalog/shares"
	err := a.client.Get(ctx, path, nil, &listSharesResponse)
	return &listSharesResponse, err
}

func (a *SharesAPI) ListAll(ctx context.Context) ([]ShareInfo, error) {
	response, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Shares, nil
}

// Update a share
//
// Updates the share with the changes and data objects in the request. The
// caller must be the owner of the share or a Metastore admin.
//
// When the caller is a Metastore admin, only the __owner__ field can be
// updated.
//
// In the case that the Share name is changed, **updateShare** requires that the
// caller is both the share owner and a Metastore admin.
//
// For each table that is added through this method, the share owner must also
// have SELECT privilege on the table. This privilege must be maintained
// indefinitely for recipients to be able to access the table. Typically, you
// should use a group as the share owner.
//
// Table removals through **update** do not require additional privileges.
func (a *SharesAPI) Update(ctx context.Context, request UpdateShare) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

// Update permissions
//
// Updates the permissions for a data share in the Metastore. The caller must be
// a Metastore admin or an owner of the share.
//
// For new recipient grants, the user must also be the owner of the recipients.
// recipient revocations do not require additional privileges.
func (a *SharesAPI) UpdateSharePermissions(ctx context.Context, request UpdateSharePermissions) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewStorageCredentials(client *client.DatabricksClient) StorageCredentialsService {
	return &StorageCredentialsAPI{
		client: client,
	}
}

type StorageCredentialsAPI struct {
	client *client.DatabricksClient
}

// Create credentials
//
// Creates a new storage credential. The request object is specific to the
// cloud:
//
// * **AwsIamRole** for AWS credentials * **AzureServicePrincipal** for Azure
// credentials * **GcpServiceAcountKey** for GCP credentials.
//
// The caller must be a Metastore admin and have the CREATE STORAGE CREDENTIAL
// privilege on the Metastore.
func (a *StorageCredentialsAPI) Create(ctx context.Context, request CreateStorageCredential) (*CreateStorageCredentialResponse, error) {
	var createStorageCredentialResponse CreateStorageCredentialResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	err := a.client.Post(ctx, path, request, &createStorageCredentialResponse)
	return &createStorageCredentialResponse, err
}

// Delete a credential
//
// Deletes a storage credential from the Metastore. The caller must be an owner
// of the storage credential.
func (a *StorageCredentialsAPI) DeleteStorageCredential(ctx context.Context, request DeleteStorageCredentialRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a credential
//
// Deletes a storage credential from the Metastore. The caller must be an owner
// of the storage credential.
func (a *StorageCredentialsAPI) DeleteStorageCredentialByName(ctx context.Context, name string) error {
	return a.DeleteStorageCredential(ctx, DeleteStorageCredentialRequest{
		Name: name,
	})
}

// Get a credential
//
// Gets a storage credential from the Metastore. The caller must be a Metastore
// admin, the owner of the storage credential, or have a level of privilege on
// the storage credential.
func (a *StorageCredentialsAPI) GetStorageCredentials(ctx context.Context, request GetStorageCredentialsRequest) (*GetStorageCredentialResponse, error) {
	var getStorageCredentialResponse GetStorageCredentialResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getStorageCredentialResponse)
	return &getStorageCredentialResponse, err
}

// Get a credential
//
// Gets a storage credential from the Metastore. The caller must be a Metastore
// admin, the owner of the storage credential, or have a level of privilege on
// the storage credential.
func (a *StorageCredentialsAPI) GetStorageCredentialsByName(ctx context.Context, name string) (*GetStorageCredentialResponse, error) {
	return a.GetStorageCredentials(ctx, GetStorageCredentialsRequest{
		Name: name,
	})
}

// List credentials
//
// Gets an array of storage credentials (as StorageCredentialInfo objects). The
// array is limited to only those storage credentials the caller has the
// privilege level to access. If the caller is a Metastore admin, all storage
// credentials will be retrieved.
func (a *StorageCredentialsAPI) List(ctx context.Context) (*ListStorageCredentialsResponse, error) {
	var listStorageCredentialsResponse ListStorageCredentialsResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	err := a.client.Get(ctx, path, nil, &listStorageCredentialsResponse)
	return &listStorageCredentialsResponse, err
}

func (a *StorageCredentialsAPI) ListAll(ctx context.Context) ([]StorageCredentialInfo, error) {
	response, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.StorageCredentials, nil
}

// Update a credential
//
// Updates a storage credential on the Metastore. The caller must be the owner
// of the storage credential. If the caller is a Metastore admin, only the
// __owner__ credential can be changed.
func (a *StorageCredentialsAPI) Update(ctx context.Context, request UpdateStorageCredential) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewTables(client *client.DatabricksClient) TablesService {
	return &TablesAPI{
		client: client,
	}
}

type TablesAPI struct {
	client *client.DatabricksClient
}

// Create a table
//
// Creates a new table in the Metastore for a specific catalog and schema.
// **WARNING**: Do not use this API at this time.
//
// The caller must be the owner of or have the USAGE privilege for both the
// parent catalog and schema, or be the owner of the parent schema (or have the
// CREATE privilege on it).
//
// If the new table has a __table_type__ of EXTERNAL specified, the user must be
// a Metastore admin or meet the permissions requirements of the storage
// credential or the external location.
func (a *TablesAPI) Create(ctx context.Context, request CreateTable) (*CreateTableResponse, error) {
	var createTableResponse CreateTableResponse
	path := "/api/2.1/unity-catalog/tables"
	err := a.client.Post(ctx, path, request, &createTableResponse)
	return &createTableResponse, err
}

// Create a staging table
//
// Creates a new staging table for a schema. The caller must have both the USAGE
// privilege on the parent Catalog and the USAGE and CREATE privileges on the
// parent schema.
func (a *TablesAPI) CreateStagingTable(ctx context.Context, request CreateStagingTable) (*CreateStagingTableResponse, error) {
	var createStagingTableResponse CreateStagingTableResponse
	path := "/api/2.1/unity-catalog/staging-tables"
	err := a.client.Post(ctx, path, request, &createStagingTableResponse)
	return &createStagingTableResponse, err
}

// Delete a table
//
// Deletes a table from the specified parent catalog and schema. The caller must
// be the owner of the parent catalog, have the USAGE privilege on the parent
// catalog and be the owner of the parent schema, or be the owner of the table
// and have the USAGE privilege on both the parent catalog and schema.
func (a *TablesAPI) DeleteTable(ctx context.Context, request DeleteTableRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a table
//
// Deletes a table from the specified parent catalog and schema. The caller must
// be the owner of the parent catalog, have the USAGE privilege on the parent
// catalog and be the owner of the parent schema, or be the owner of the table
// and have the USAGE privilege on both the parent catalog and schema.
func (a *TablesAPI) DeleteTableByFullName(ctx context.Context, fullName string) error {
	return a.DeleteTable(ctx, DeleteTableRequest{
		FullName: fullName,
	})
}

// Get a table
//
// Gets a table from the Metastore for a specific catalog and schema. The caller
// must be a Metastore admin, be the owner of the table and have the USAGE
// privilege on both the parent catalog and schema, or be the owner of the table
// and have the SELECT privilege on it as well.
func (a *TablesAPI) GetTable(ctx context.Context, request GetTableRequest) (*GetTableResponse, error) {
	var getTableResponse GetTableResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	err := a.client.Get(ctx, path, request, &getTableResponse)
	return &getTableResponse, err
}

// Get a table
//
// Gets a table from the Metastore for a specific catalog and schema. The caller
// must be a Metastore admin, be the owner of the table and have the USAGE
// privilege on both the parent catalog and schema, or be the owner of the table
// and have the SELECT privilege on it as well.
func (a *TablesAPI) GetTableByFullName(ctx context.Context, fullName string) (*GetTableResponse, error) {
	return a.GetTable(ctx, GetTableRequest{
		FullName: fullName,
	})
}

// List tables
//
// Gets an array of all tables for the current Metastore under the parent
// catalog and schema. The caller must be a Metastore admin or an owner of (or
// have the SELECT privilege on) the table. For the latter case, the caller must
// also be the owner or have the USAGE privilege on the parent catalog and
// schema.
func (a *TablesAPI) List(ctx context.Context, request ListRequest) (*ListTablesResponse, error) {
	var listTablesResponse ListTablesResponse
	path := "/api/2.1/unity-catalog/tables"
	err := a.client.Get(ctx, path, request, &listTablesResponse)
	return &listTablesResponse, err
}

func (a *TablesAPI) ListAll(ctx context.Context, request ListRequest) ([]TableInfo, error) {
	response, err := a.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Tables, nil
}

// List table summaries
//
// Gets an array of summaries for tables for a schema and catalog within the
// Metastore. The table summaries returned are either:
//
// * summaries for all tables (within the current Metastore and parent catalog
// and schema), when the user is a Metastore admin, or: * summaries for all
// tables and schemas (within the current Metastore and parent catalog) for
// which the user has ownership or the SELECT privilege on the Table and
// ownership or USAGE privilege on the Schema, provided that the user also has
// ownership or the USAGE privilege on the parent Catalog
func (a *TablesAPI) ListTableSummaries(ctx context.Context, request ListTableSummariesRequest) (*ListTableSummariesResponse, error) {
	var listTableSummariesResponse ListTableSummariesResponse
	path := "/api/2.1/unity-catalog/table-summaries"
	err := a.client.Get(ctx, path, request, &listTableSummariesResponse)
	return &listTableSummariesResponse, err
}

// Update a table
//
// Updates a table in the specified catalog and schema. The caller must be the
// owner of have the USAGE privilege on the parent catalog and schema, or, if
// changing the owner, be a Metastore admin as well.
func (a *TablesAPI) Update(ctx context.Context, request UpdateTable) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewUnityFiles(client *client.DatabricksClient) UnityFilesService {
	return &UnityFilesAPI{
		client: client,
	}
}

type UnityFilesAPI struct {
	client *client.DatabricksClient
}

// List files
//
// List the files sound at the supplied URL.
func (a *UnityFilesAPI) ListFiles(ctx context.Context, request ListFilesRequest) (*ListFilesResponse, error) {
	var listFilesResponse ListFilesResponse
	path := "/api/2.1/unity-catalog/files"
	err := a.client.Get(ctx, path, request, &listFilesResponse)
	return &listFilesResponse, err
}

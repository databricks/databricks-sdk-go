// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just AccountMetastoreAssignments API methods
type accountMetastoreAssignmentsImpl struct {
	client *client.DatabricksClient
}

func (a *accountMetastoreAssignmentsImpl) Create(ctx context.Context, request AccountsCreateMetastoreAssignment) error {
	var createResponse CreateResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createResponse)
	return err
}

func (a *accountMetastoreAssignmentsImpl) Delete(ctx context.Context, request DeleteAccountMetastoreAssignmentRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *accountMetastoreAssignmentsImpl) Get(ctx context.Context, request GetAccountMetastoreAssignmentRequest) (*AccountsMetastoreAssignment, error) {
	var accountsMetastoreAssignment AccountsMetastoreAssignment
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastore", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &accountsMetastoreAssignment)
	return &accountsMetastoreAssignment, err
}

// Get all workspaces assigned to a metastore.
//
// Gets a list of all Databricks workspace IDs that have been assigned to given
// metastore.
func (a *accountMetastoreAssignmentsImpl) List(ctx context.Context, request ListAccountMetastoreAssignmentsRequest) listing.Iterator[int64] {

	getNextPage := func(ctx context.Context, req ListAccountMetastoreAssignmentsRequest) (*ListAccountMetastoreAssignmentsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAccountMetastoreAssignmentsResponse) []int64 {
		return resp.WorkspaceIds
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get all workspaces assigned to a metastore.
//
// Gets a list of all Databricks workspace IDs that have been assigned to given
// metastore.
func (a *accountMetastoreAssignmentsImpl) ListAll(ctx context.Context, request ListAccountMetastoreAssignmentsRequest) ([]int64, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[int64](ctx, iterator)
}

func (a *accountMetastoreAssignmentsImpl) internalList(ctx context.Context, request ListAccountMetastoreAssignmentsRequest) (*ListAccountMetastoreAssignmentsResponse, error) {
	var listAccountMetastoreAssignmentsResponse ListAccountMetastoreAssignmentsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/workspaces", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAccountMetastoreAssignmentsResponse)
	return &listAccountMetastoreAssignmentsResponse, err
}

func (a *accountMetastoreAssignmentsImpl) Update(ctx context.Context, request AccountsUpdateMetastoreAssignment) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just AccountMetastores API methods
type accountMetastoresImpl struct {
	client *client.DatabricksClient
}

func (a *accountMetastoresImpl) Create(ctx context.Context, request AccountsCreateMetastore) (*AccountsMetastoreInfo, error) {
	var accountsMetastoreInfo AccountsMetastoreInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &accountsMetastoreInfo)
	return &accountsMetastoreInfo, err
}

func (a *accountMetastoresImpl) Delete(ctx context.Context, request DeleteAccountMetastoreRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *accountMetastoresImpl) Get(ctx context.Context, request GetAccountMetastoreRequest) (*AccountsMetastoreInfo, error) {
	var accountsMetastoreInfo AccountsMetastoreInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &accountsMetastoreInfo)
	return &accountsMetastoreInfo, err
}

// Get all metastores associated with an account.
//
// Gets all Unity Catalog metastores associated with an account specified by ID.
func (a *accountMetastoresImpl) List(ctx context.Context) listing.Iterator[MetastoreInfo] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListMetastoresResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *ListMetastoresResponse) []MetastoreInfo {
		return resp.Metastores
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get all metastores associated with an account.
//
// Gets all Unity Catalog metastores associated with an account specified by ID.
func (a *accountMetastoresImpl) ListAll(ctx context.Context) ([]MetastoreInfo, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[MetastoreInfo](ctx, iterator)
}

func (a *accountMetastoresImpl) internalList(ctx context.Context) (*ListMetastoresResponse, error) {
	var listMetastoresResponse ListMetastoresResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listMetastoresResponse)
	return &listMetastoresResponse, err
}

func (a *accountMetastoresImpl) Update(ctx context.Context, request AccountsUpdateMetastore) (*AccountsMetastoreInfo, error) {
	var accountsMetastoreInfo AccountsMetastoreInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &accountsMetastoreInfo)
	return &accountsMetastoreInfo, err
}

// unexported type that holds implementations of just AccountStorageCredentials API methods
type accountStorageCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *accountStorageCredentialsImpl) Create(ctx context.Context, request AccountsCreateStorageCredential) (*AccountsStorageCredentialInfo, error) {
	var accountsStorageCredentialInfo AccountsStorageCredentialInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &accountsStorageCredentialInfo)
	return &accountsStorageCredentialInfo, err
}

func (a *accountStorageCredentialsImpl) Delete(ctx context.Context, request DeleteAccountStorageCredentialRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/%v", a.client.ConfiguredAccountID(), request.MetastoreId, request.StorageCredentialName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *accountStorageCredentialsImpl) Get(ctx context.Context, request GetAccountStorageCredentialRequest) (*AccountsStorageCredentialInfo, error) {
	var accountsStorageCredentialInfo AccountsStorageCredentialInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/%v", a.client.ConfiguredAccountID(), request.MetastoreId, request.StorageCredentialName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &accountsStorageCredentialInfo)
	return &accountsStorageCredentialInfo, err
}

// Get all storage credentials assigned to a metastore.
//
// Gets a list of all storage credentials that have been assigned to given
// metastore.
func (a *accountStorageCredentialsImpl) List(ctx context.Context, request ListAccountStorageCredentialsRequest) listing.Iterator[StorageCredentialInfo] {

	getNextPage := func(ctx context.Context, req ListAccountStorageCredentialsRequest) (*ListAccountStorageCredentialsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListAccountStorageCredentialsResponse) []StorageCredentialInfo {
		return resp.StorageCredentials
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get all storage credentials assigned to a metastore.
//
// Gets a list of all storage credentials that have been assigned to given
// metastore.
func (a *accountStorageCredentialsImpl) ListAll(ctx context.Context, request ListAccountStorageCredentialsRequest) ([]StorageCredentialInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[StorageCredentialInfo](ctx, iterator)
}

func (a *accountStorageCredentialsImpl) internalList(ctx context.Context, request ListAccountStorageCredentialsRequest) (*ListAccountStorageCredentialsResponse, error) {
	var listAccountStorageCredentialsResponse ListAccountStorageCredentialsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listAccountStorageCredentialsResponse)
	return &listAccountStorageCredentialsResponse, err
}

func (a *accountStorageCredentialsImpl) Update(ctx context.Context, request AccountsUpdateStorageCredential) (*AccountsStorageCredentialInfo, error) {
	var accountsStorageCredentialInfo AccountsStorageCredentialInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/%v", a.client.ConfiguredAccountID(), request.MetastoreId, request.StorageCredentialName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &accountsStorageCredentialInfo)
	return &accountsStorageCredentialInfo, err
}

// unexported type that holds implementations of just ArtifactAllowlists API methods
type artifactAllowlistsImpl struct {
	client *client.DatabricksClient
}

func (a *artifactAllowlistsImpl) Get(ctx context.Context, request GetArtifactAllowlistRequest) (*ArtifactAllowlistInfo, error) {
	var artifactAllowlistInfo ArtifactAllowlistInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/artifact-allowlists/%v", request.ArtifactType)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &artifactAllowlistInfo)
	return &artifactAllowlistInfo, err
}

func (a *artifactAllowlistsImpl) Update(ctx context.Context, request SetArtifactAllowlist) (*ArtifactAllowlistInfo, error) {
	var artifactAllowlistInfo ArtifactAllowlistInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/artifact-allowlists/%v", request.ArtifactType)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &artifactAllowlistInfo)
	return &artifactAllowlistInfo, err
}

// unexported type that holds implementations of just Catalogs API methods
type catalogsImpl struct {
	client *client.DatabricksClient
}

func (a *catalogsImpl) Create(ctx context.Context, request CreateCatalog) (*CatalogInfo, error) {
	var catalogInfo CatalogInfo
	path := "/api/2.1/unity-catalog/catalogs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &catalogInfo)
	return &catalogInfo, err
}

func (a *catalogsImpl) Delete(ctx context.Context, request DeleteCatalogRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *catalogsImpl) Get(ctx context.Context, request GetCatalogRequest) (*CatalogInfo, error) {
	var catalogInfo CatalogInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &catalogInfo)
	return &catalogInfo, err
}

// List catalogs.
//
// Gets an array of catalogs in the metastore. If the caller is the metastore
// admin, all catalogs will be retrieved. Otherwise, only catalogs owned by the
// caller (or for which the caller has the **USE_CATALOG** privilege) will be
// retrieved. There is no guarantee of a specific ordering of the elements in
// the array.
func (a *catalogsImpl) List(ctx context.Context, request ListCatalogsRequest) listing.Iterator[CatalogInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListCatalogsRequest) (*ListCatalogsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCatalogsResponse) []CatalogInfo {
		return resp.Catalogs
	}
	getNextReq := func(resp *ListCatalogsResponse) *ListCatalogsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List catalogs.
//
// Gets an array of catalogs in the metastore. If the caller is the metastore
// admin, all catalogs will be retrieved. Otherwise, only catalogs owned by the
// caller (or for which the caller has the **USE_CATALOG** privilege) will be
// retrieved. There is no guarantee of a specific ordering of the elements in
// the array.
func (a *catalogsImpl) ListAll(ctx context.Context, request ListCatalogsRequest) ([]CatalogInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CatalogInfo](ctx, iterator)
}

func (a *catalogsImpl) internalList(ctx context.Context, request ListCatalogsRequest) (*ListCatalogsResponse, error) {
	var listCatalogsResponse ListCatalogsResponse
	path := "/api/2.1/unity-catalog/catalogs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCatalogsResponse)
	return &listCatalogsResponse, err
}

func (a *catalogsImpl) Update(ctx context.Context, request UpdateCatalog) (*CatalogInfo, error) {
	var catalogInfo CatalogInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &catalogInfo)
	return &catalogInfo, err
}

// unexported type that holds implementations of just Connections API methods
type connectionsImpl struct {
	client *client.DatabricksClient
}

func (a *connectionsImpl) Create(ctx context.Context, request CreateConnection) (*ConnectionInfo, error) {
	var connectionInfo ConnectionInfo
	path := "/api/2.1/unity-catalog/connections"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &connectionInfo)
	return &connectionInfo, err
}

func (a *connectionsImpl) Delete(ctx context.Context, request DeleteConnectionRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/connections/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *connectionsImpl) Get(ctx context.Context, request GetConnectionRequest) (*ConnectionInfo, error) {
	var connectionInfo ConnectionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/connections/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &connectionInfo)
	return &connectionInfo, err
}

// List connections.
//
// List all connections.
func (a *connectionsImpl) List(ctx context.Context, request ListConnectionsRequest) listing.Iterator[ConnectionInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListConnectionsRequest) (*ListConnectionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListConnectionsResponse) []ConnectionInfo {
		return resp.Connections
	}
	getNextReq := func(resp *ListConnectionsResponse) *ListConnectionsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List connections.
//
// List all connections.
func (a *connectionsImpl) ListAll(ctx context.Context, request ListConnectionsRequest) ([]ConnectionInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ConnectionInfo](ctx, iterator)
}

func (a *connectionsImpl) internalList(ctx context.Context, request ListConnectionsRequest) (*ListConnectionsResponse, error) {
	var listConnectionsResponse ListConnectionsResponse
	path := "/api/2.1/unity-catalog/connections"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listConnectionsResponse)
	return &listConnectionsResponse, err
}

func (a *connectionsImpl) Update(ctx context.Context, request UpdateConnection) (*ConnectionInfo, error) {
	var connectionInfo ConnectionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/connections/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &connectionInfo)
	return &connectionInfo, err
}

// unexported type that holds implementations of just Credentials API methods
type credentialsImpl struct {
	client *client.DatabricksClient
}

func (a *credentialsImpl) CreateCredential(ctx context.Context, request CreateCredentialRequest) (*CredentialInfo, error) {
	var credentialInfo CredentialInfo
	path := "/api/2.1/unity-catalog/credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &credentialInfo)
	return &credentialInfo, err
}

func (a *credentialsImpl) DeleteCredential(ctx context.Context, request DeleteCredentialRequest) error {
	var deleteCredentialResponse DeleteCredentialResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/credentials/%v", request.NameArg)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteCredentialResponse)
	return err
}

func (a *credentialsImpl) GenerateTemporaryServiceCredential(ctx context.Context, request GenerateTemporaryServiceCredentialRequest) (*TemporaryCredentials, error) {
	var temporaryCredentials TemporaryCredentials
	path := "/api/2.1/unity-catalog/temporary-service-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &temporaryCredentials)
	return &temporaryCredentials, err
}

func (a *credentialsImpl) GetCredential(ctx context.Context, request GetCredentialRequest) (*CredentialInfo, error) {
	var credentialInfo CredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/credentials/%v", request.NameArg)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &credentialInfo)
	return &credentialInfo, err
}

// List credentials.
//
// Gets an array of credentials (as __CredentialInfo__ objects).
//
// The array is limited to only the credentials that the caller has permission
// to access. If the caller is a metastore admin, retrieval of credentials is
// unrestricted. There is no guarantee of a specific ordering of the elements in
// the array.
func (a *credentialsImpl) ListCredentials(ctx context.Context, request ListCredentialsRequest) listing.Iterator[CredentialInfo] {

	getNextPage := func(ctx context.Context, req ListCredentialsRequest) (*ListCredentialsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListCredentials(ctx, req)
	}
	getItems := func(resp *ListCredentialsResponse) []CredentialInfo {
		return resp.Credentials
	}
	getNextReq := func(resp *ListCredentialsResponse) *ListCredentialsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List credentials.
//
// Gets an array of credentials (as __CredentialInfo__ objects).
//
// The array is limited to only the credentials that the caller has permission
// to access. If the caller is a metastore admin, retrieval of credentials is
// unrestricted. There is no guarantee of a specific ordering of the elements in
// the array.
func (a *credentialsImpl) ListCredentialsAll(ctx context.Context, request ListCredentialsRequest) ([]CredentialInfo, error) {
	iterator := a.ListCredentials(ctx, request)
	return listing.ToSlice[CredentialInfo](ctx, iterator)
}

func (a *credentialsImpl) internalListCredentials(ctx context.Context, request ListCredentialsRequest) (*ListCredentialsResponse, error) {
	var listCredentialsResponse ListCredentialsResponse
	path := "/api/2.1/unity-catalog/credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCredentialsResponse)
	return &listCredentialsResponse, err
}

func (a *credentialsImpl) UpdateCredential(ctx context.Context, request UpdateCredentialRequest) (*CredentialInfo, error) {
	var credentialInfo CredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/credentials/%v", request.NameArg)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &credentialInfo)
	return &credentialInfo, err
}

func (a *credentialsImpl) ValidateCredential(ctx context.Context, request ValidateCredentialRequest) (*ValidateCredentialResponse, error) {
	var validateCredentialResponse ValidateCredentialResponse
	path := "/api/2.1/unity-catalog/validate-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &validateCredentialResponse)
	return &validateCredentialResponse, err
}

// unexported type that holds implementations of just ExternalLocations API methods
type externalLocationsImpl struct {
	client *client.DatabricksClient
}

func (a *externalLocationsImpl) Create(ctx context.Context, request CreateExternalLocation) (*ExternalLocationInfo, error) {
	var externalLocationInfo ExternalLocationInfo
	path := "/api/2.1/unity-catalog/external-locations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &externalLocationInfo)
	return &externalLocationInfo, err
}

func (a *externalLocationsImpl) Delete(ctx context.Context, request DeleteExternalLocationRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *externalLocationsImpl) Get(ctx context.Context, request GetExternalLocationRequest) (*ExternalLocationInfo, error) {
	var externalLocationInfo ExternalLocationInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &externalLocationInfo)
	return &externalLocationInfo, err
}

// List external locations.
//
// Gets an array of external locations (__ExternalLocationInfo__ objects) from
// the metastore. The caller must be a metastore admin, the owner of the
// external location, or a user that has some privilege on the external
// location. There is no guarantee of a specific ordering of the elements in the
// array.
func (a *externalLocationsImpl) List(ctx context.Context, request ListExternalLocationsRequest) listing.Iterator[ExternalLocationInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListExternalLocationsRequest) (*ListExternalLocationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListExternalLocationsResponse) []ExternalLocationInfo {
		return resp.ExternalLocations
	}
	getNextReq := func(resp *ListExternalLocationsResponse) *ListExternalLocationsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List external locations.
//
// Gets an array of external locations (__ExternalLocationInfo__ objects) from
// the metastore. The caller must be a metastore admin, the owner of the
// external location, or a user that has some privilege on the external
// location. There is no guarantee of a specific ordering of the elements in the
// array.
func (a *externalLocationsImpl) ListAll(ctx context.Context, request ListExternalLocationsRequest) ([]ExternalLocationInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ExternalLocationInfo](ctx, iterator)
}

func (a *externalLocationsImpl) internalList(ctx context.Context, request ListExternalLocationsRequest) (*ListExternalLocationsResponse, error) {
	var listExternalLocationsResponse ListExternalLocationsResponse
	path := "/api/2.1/unity-catalog/external-locations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listExternalLocationsResponse)
	return &listExternalLocationsResponse, err
}

func (a *externalLocationsImpl) Update(ctx context.Context, request UpdateExternalLocation) (*ExternalLocationInfo, error) {
	var externalLocationInfo ExternalLocationInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &externalLocationInfo)
	return &externalLocationInfo, err
}

// unexported type that holds implementations of just Functions API methods
type functionsImpl struct {
	client *client.DatabricksClient
}

func (a *functionsImpl) Create(ctx context.Context, request CreateFunctionRequest) (*FunctionInfo, error) {
	var functionInfo FunctionInfo
	path := "/api/2.1/unity-catalog/functions"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &functionInfo)
	return &functionInfo, err
}

func (a *functionsImpl) Delete(ctx context.Context, request DeleteFunctionRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *functionsImpl) Get(ctx context.Context, request GetFunctionRequest) (*FunctionInfo, error) {
	var functionInfo FunctionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &functionInfo)
	return &functionInfo, err
}

// List functions.
//
// List functions within the specified parent catalog and schema. If the user is
// a metastore admin, all functions are returned in the output list. Otherwise,
// the user must have the **USE_CATALOG** privilege on the catalog and the
// **USE_SCHEMA** privilege on the schema, and the output list contains only
// functions for which either the user has the **EXECUTE** privilege or the user
// is the owner. There is no guarantee of a specific ordering of the elements in
// the array.
func (a *functionsImpl) List(ctx context.Context, request ListFunctionsRequest) listing.Iterator[FunctionInfo] {

	getNextPage := func(ctx context.Context, req ListFunctionsRequest) (*ListFunctionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListFunctionsResponse) []FunctionInfo {
		return resp.Functions
	}
	getNextReq := func(resp *ListFunctionsResponse) *ListFunctionsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List functions.
//
// List functions within the specified parent catalog and schema. If the user is
// a metastore admin, all functions are returned in the output list. Otherwise,
// the user must have the **USE_CATALOG** privilege on the catalog and the
// **USE_SCHEMA** privilege on the schema, and the output list contains only
// functions for which either the user has the **EXECUTE** privilege or the user
// is the owner. There is no guarantee of a specific ordering of the elements in
// the array.
func (a *functionsImpl) ListAll(ctx context.Context, request ListFunctionsRequest) ([]FunctionInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[FunctionInfo](ctx, iterator)
}

func (a *functionsImpl) internalList(ctx context.Context, request ListFunctionsRequest) (*ListFunctionsResponse, error) {
	var listFunctionsResponse ListFunctionsResponse
	path := "/api/2.1/unity-catalog/functions"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listFunctionsResponse)
	return &listFunctionsResponse, err
}

func (a *functionsImpl) Update(ctx context.Context, request UpdateFunction) (*FunctionInfo, error) {
	var functionInfo FunctionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &functionInfo)
	return &functionInfo, err
}

// unexported type that holds implementations of just Grants API methods
type grantsImpl struct {
	client *client.DatabricksClient
}

func (a *grantsImpl) Get(ctx context.Context, request GetGrantRequest) (*PermissionsList, error) {
	var permissionsList PermissionsList
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &permissionsList)
	return &permissionsList, err
}

func (a *grantsImpl) GetEffective(ctx context.Context, request GetEffectiveRequest) (*EffectivePermissionsList, error) {
	var effectivePermissionsList EffectivePermissionsList
	path := fmt.Sprintf("/api/2.1/unity-catalog/effective-permissions/%v/%v", request.SecurableType, request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &effectivePermissionsList)
	return &effectivePermissionsList, err
}

func (a *grantsImpl) Update(ctx context.Context, request UpdatePermissions) (*PermissionsList, error) {
	var permissionsList PermissionsList
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &permissionsList)
	return &permissionsList, err
}

// unexported type that holds implementations of just Metastores API methods
type metastoresImpl struct {
	client *client.DatabricksClient
}

func (a *metastoresImpl) Assign(ctx context.Context, request CreateMetastoreAssignment) error {
	var assignResponse AssignResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &assignResponse)
	return err
}

func (a *metastoresImpl) Create(ctx context.Context, request CreateMetastore) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := "/api/2.1/unity-catalog/metastores"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *metastoresImpl) Current(ctx context.Context) (*MetastoreAssignment, error) {
	var metastoreAssignment MetastoreAssignment
	path := "/api/2.1/unity-catalog/current-metastore-assignment"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &metastoreAssignment)
	return &metastoreAssignment, err
}

func (a *metastoresImpl) Delete(ctx context.Context, request DeleteMetastoreRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *metastoresImpl) Get(ctx context.Context, request GetMetastoreRequest) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &metastoreInfo)
	return &metastoreInfo, err
}

// List metastores.
//
// Gets an array of the available metastores (as __MetastoreInfo__ objects). The
// caller must be an admin to retrieve this info. There is no guarantee of a
// specific ordering of the elements in the array.
func (a *metastoresImpl) List(ctx context.Context) listing.Iterator[MetastoreInfo] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListMetastoresResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *ListMetastoresResponse) []MetastoreInfo {
		return resp.Metastores
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List metastores.
//
// Gets an array of the available metastores (as __MetastoreInfo__ objects). The
// caller must be an admin to retrieve this info. There is no guarantee of a
// specific ordering of the elements in the array.
func (a *metastoresImpl) ListAll(ctx context.Context) ([]MetastoreInfo, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[MetastoreInfo](ctx, iterator)
}

func (a *metastoresImpl) internalList(ctx context.Context) (*ListMetastoresResponse, error) {
	var listMetastoresResponse ListMetastoresResponse
	path := "/api/2.1/unity-catalog/metastores"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listMetastoresResponse)
	return &listMetastoresResponse, err
}

func (a *metastoresImpl) Summary(ctx context.Context) (*GetMetastoreSummaryResponse, error) {
	var getMetastoreSummaryResponse GetMetastoreSummaryResponse
	path := "/api/2.1/unity-catalog/metastore_summary"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getMetastoreSummaryResponse)
	return &getMetastoreSummaryResponse, err
}

func (a *metastoresImpl) Unassign(ctx context.Context, request UnassignRequest) error {
	var unassignResponse UnassignResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &unassignResponse)
	return err
}

func (a *metastoresImpl) Update(ctx context.Context, request UpdateMetastore) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *metastoresImpl) UpdateAssignment(ctx context.Context, request UpdateMetastoreAssignment) error {
	var updateAssignmentResponse UpdateAssignmentResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateAssignmentResponse)
	return err
}

// unexported type that holds implementations of just ModelVersions API methods
type modelVersionsImpl struct {
	client *client.DatabricksClient
}

func (a *modelVersionsImpl) Delete(ctx context.Context, request DeleteModelVersionRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions/%v", request.FullName, request.Version)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *modelVersionsImpl) Get(ctx context.Context, request GetModelVersionRequest) (*ModelVersionInfo, error) {
	var modelVersionInfo ModelVersionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions/%v", request.FullName, request.Version)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &modelVersionInfo)
	return &modelVersionInfo, err
}

func (a *modelVersionsImpl) GetByAlias(ctx context.Context, request GetByAliasRequest) (*ModelVersionInfo, error) {
	var modelVersionInfo ModelVersionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/aliases/%v", request.FullName, request.Alias)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &modelVersionInfo)
	return &modelVersionInfo, err
}

// List Model Versions.
//
// List model versions. You can list model versions under a particular schema,
// or list all model versions in the current metastore.
//
// The returned models are filtered based on the privileges of the calling user.
// For example, the metastore admin is able to list all the model versions. A
// regular user needs to be the owner or have the **EXECUTE** privilege on the
// parent registered model to recieve the model versions in the response. For
// the latter case, the caller must also be the owner or have the
// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
// privilege on the parent schema.
//
// There is no guarantee of a specific ordering of the elements in the response.
// The elements in the response will not contain any aliases or tags.
func (a *modelVersionsImpl) List(ctx context.Context, request ListModelVersionsRequest) listing.Iterator[ModelVersionInfo] {

	getNextPage := func(ctx context.Context, req ListModelVersionsRequest) (*ListModelVersionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListModelVersionsResponse) []ModelVersionInfo {
		return resp.ModelVersions
	}
	getNextReq := func(resp *ListModelVersionsResponse) *ListModelVersionsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List Model Versions.
//
// List model versions. You can list model versions under a particular schema,
// or list all model versions in the current metastore.
//
// The returned models are filtered based on the privileges of the calling user.
// For example, the metastore admin is able to list all the model versions. A
// regular user needs to be the owner or have the **EXECUTE** privilege on the
// parent registered model to recieve the model versions in the response. For
// the latter case, the caller must also be the owner or have the
// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
// privilege on the parent schema.
//
// There is no guarantee of a specific ordering of the elements in the response.
// The elements in the response will not contain any aliases or tags.
func (a *modelVersionsImpl) ListAll(ctx context.Context, request ListModelVersionsRequest) ([]ModelVersionInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ModelVersionInfo](ctx, iterator)
}

func (a *modelVersionsImpl) internalList(ctx context.Context, request ListModelVersionsRequest) (*ListModelVersionsResponse, error) {
	var listModelVersionsResponse ListModelVersionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listModelVersionsResponse)
	return &listModelVersionsResponse, err
}

func (a *modelVersionsImpl) Update(ctx context.Context, request UpdateModelVersionRequest) (*ModelVersionInfo, error) {
	var modelVersionInfo ModelVersionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions/%v", request.FullName, request.Version)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &modelVersionInfo)
	return &modelVersionInfo, err
}

// unexported type that holds implementations of just OnlineTables API methods
type onlineTablesImpl struct {
	client *client.DatabricksClient
}

func (a *onlineTablesImpl) Create(ctx context.Context, request CreateOnlineTableRequest) (*OnlineTable, error) {
	var onlineTable OnlineTable
	path := "/api/2.0/online-tables"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Table, &onlineTable)
	return &onlineTable, err
}

func (a *onlineTablesImpl) Delete(ctx context.Context, request DeleteOnlineTableRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/online-tables/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *onlineTablesImpl) Get(ctx context.Context, request GetOnlineTableRequest) (*OnlineTable, error) {
	var onlineTable OnlineTable
	path := fmt.Sprintf("/api/2.0/online-tables/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &onlineTable)
	return &onlineTable, err
}

// unexported type that holds implementations of just QualityMonitors API methods
type qualityMonitorsImpl struct {
	client *client.DatabricksClient
}

func (a *qualityMonitorsImpl) CancelRefresh(ctx context.Context, request CancelRefreshRequest) error {
	var cancelRefreshResponse CancelRefreshResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes/%v/cancel", request.TableName, request.RefreshId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &cancelRefreshResponse)
	return err
}

func (a *qualityMonitorsImpl) Create(ctx context.Context, request CreateMonitor) (*MonitorInfo, error) {
	var monitorInfo MonitorInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &monitorInfo)
	return &monitorInfo, err
}

func (a *qualityMonitorsImpl) Delete(ctx context.Context, request DeleteQualityMonitorRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *qualityMonitorsImpl) Get(ctx context.Context, request GetQualityMonitorRequest) (*MonitorInfo, error) {
	var monitorInfo MonitorInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &monitorInfo)
	return &monitorInfo, err
}

func (a *qualityMonitorsImpl) GetRefresh(ctx context.Context, request GetRefreshRequest) (*MonitorRefreshInfo, error) {
	var monitorRefreshInfo MonitorRefreshInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes/%v", request.TableName, request.RefreshId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &monitorRefreshInfo)
	return &monitorRefreshInfo, err
}

func (a *qualityMonitorsImpl) ListRefreshes(ctx context.Context, request ListRefreshesRequest) (*MonitorRefreshListResponse, error) {
	var monitorRefreshListResponse MonitorRefreshListResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &monitorRefreshListResponse)
	return &monitorRefreshListResponse, err
}

func (a *qualityMonitorsImpl) RegenerateDashboard(ctx context.Context, request RegenerateDashboardRequest) (*RegenerateDashboardResponse, error) {
	var regenerateDashboardResponse RegenerateDashboardResponse
	path := fmt.Sprintf("/api/2.1/quality-monitoring/tables/%v/monitor/dashboard", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &regenerateDashboardResponse)
	return &regenerateDashboardResponse, err
}

func (a *qualityMonitorsImpl) RunRefresh(ctx context.Context, request RunRefreshRequest) (*MonitorRefreshInfo, error) {
	var monitorRefreshInfo MonitorRefreshInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, nil, &monitorRefreshInfo)
	return &monitorRefreshInfo, err
}

func (a *qualityMonitorsImpl) Update(ctx context.Context, request UpdateMonitor) (*MonitorInfo, error) {
	var monitorInfo MonitorInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &monitorInfo)
	return &monitorInfo, err
}

// unexported type that holds implementations of just RegisteredModels API methods
type registeredModelsImpl struct {
	client *client.DatabricksClient
}

func (a *registeredModelsImpl) Create(ctx context.Context, request CreateRegisteredModelRequest) (*RegisteredModelInfo, error) {
	var registeredModelInfo RegisteredModelInfo
	path := "/api/2.1/unity-catalog/models"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &registeredModelInfo)
	return &registeredModelInfo, err
}

func (a *registeredModelsImpl) Delete(ctx context.Context, request DeleteRegisteredModelRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *registeredModelsImpl) DeleteAlias(ctx context.Context, request DeleteAliasRequest) error {
	var deleteAliasResponse DeleteAliasResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/aliases/%v", request.FullName, request.Alias)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteAliasResponse)
	return err
}

func (a *registeredModelsImpl) Get(ctx context.Context, request GetRegisteredModelRequest) (*RegisteredModelInfo, error) {
	var registeredModelInfo RegisteredModelInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &registeredModelInfo)
	return &registeredModelInfo, err
}

// List Registered Models.
//
// List registered models. You can list registered models under a particular
// schema, or list all registered models in the current metastore.
//
// The returned models are filtered based on the privileges of the calling user.
// For example, the metastore admin is able to list all the registered models. A
// regular user needs to be the owner or have the **EXECUTE** privilege on the
// registered model to recieve the registered models in the response. For the
// latter case, the caller must also be the owner or have the **USE_CATALOG**
// privilege on the parent catalog and the **USE_SCHEMA** privilege on the
// parent schema.
//
// There is no guarantee of a specific ordering of the elements in the response.
func (a *registeredModelsImpl) List(ctx context.Context, request ListRegisteredModelsRequest) listing.Iterator[RegisteredModelInfo] {

	getNextPage := func(ctx context.Context, req ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListRegisteredModelsResponse) []RegisteredModelInfo {
		return resp.RegisteredModels
	}
	getNextReq := func(resp *ListRegisteredModelsResponse) *ListRegisteredModelsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List Registered Models.
//
// List registered models. You can list registered models under a particular
// schema, or list all registered models in the current metastore.
//
// The returned models are filtered based on the privileges of the calling user.
// For example, the metastore admin is able to list all the registered models. A
// regular user needs to be the owner or have the **EXECUTE** privilege on the
// registered model to recieve the registered models in the response. For the
// latter case, the caller must also be the owner or have the **USE_CATALOG**
// privilege on the parent catalog and the **USE_SCHEMA** privilege on the
// parent schema.
//
// There is no guarantee of a specific ordering of the elements in the response.
func (a *registeredModelsImpl) ListAll(ctx context.Context, request ListRegisteredModelsRequest) ([]RegisteredModelInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[RegisteredModelInfo](ctx, iterator)
}

func (a *registeredModelsImpl) internalList(ctx context.Context, request ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error) {
	var listRegisteredModelsResponse ListRegisteredModelsResponse
	path := "/api/2.1/unity-catalog/models"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRegisteredModelsResponse)
	return &listRegisteredModelsResponse, err
}

func (a *registeredModelsImpl) SetAlias(ctx context.Context, request SetRegisteredModelAliasRequest) (*RegisteredModelAlias, error) {
	var registeredModelAlias RegisteredModelAlias
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/aliases/%v", request.FullName, request.Alias)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &registeredModelAlias)
	return &registeredModelAlias, err
}

func (a *registeredModelsImpl) Update(ctx context.Context, request UpdateRegisteredModelRequest) (*RegisteredModelInfo, error) {
	var registeredModelInfo RegisteredModelInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &registeredModelInfo)
	return &registeredModelInfo, err
}

// unexported type that holds implementations of just ResourceQuotas API methods
type resourceQuotasImpl struct {
	client *client.DatabricksClient
}

func (a *resourceQuotasImpl) GetQuota(ctx context.Context, request GetQuotaRequest) (*GetQuotaResponse, error) {
	var getQuotaResponse GetQuotaResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/resource-quotas/%v/%v/%v", request.ParentSecurableType, request.ParentFullName, request.QuotaName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getQuotaResponse)
	return &getQuotaResponse, err
}

// List all resource quotas under a metastore.
//
// ListQuotas returns all quota values under the metastore. There are no SLAs on
// the freshness of the counts returned. This API does not trigger a refresh of
// quota counts.
func (a *resourceQuotasImpl) ListQuotas(ctx context.Context, request ListQuotasRequest) listing.Iterator[QuotaInfo] {

	getNextPage := func(ctx context.Context, req ListQuotasRequest) (*ListQuotasResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListQuotas(ctx, req)
	}
	getItems := func(resp *ListQuotasResponse) []QuotaInfo {
		return resp.Quotas
	}
	getNextReq := func(resp *ListQuotasResponse) *ListQuotasRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List all resource quotas under a metastore.
//
// ListQuotas returns all quota values under the metastore. There are no SLAs on
// the freshness of the counts returned. This API does not trigger a refresh of
// quota counts.
func (a *resourceQuotasImpl) ListQuotasAll(ctx context.Context, request ListQuotasRequest) ([]QuotaInfo, error) {
	iterator := a.ListQuotas(ctx, request)
	return listing.ToSlice[QuotaInfo](ctx, iterator)
}

func (a *resourceQuotasImpl) internalListQuotas(ctx context.Context, request ListQuotasRequest) (*ListQuotasResponse, error) {
	var listQuotasResponse ListQuotasResponse
	path := "/api/2.1/unity-catalog/resource-quotas/all-resource-quotas"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listQuotasResponse)
	return &listQuotasResponse, err
}

// unexported type that holds implementations of just Schemas API methods
type schemasImpl struct {
	client *client.DatabricksClient
}

func (a *schemasImpl) Create(ctx context.Context, request CreateSchema) (*SchemaInfo, error) {
	var schemaInfo SchemaInfo
	path := "/api/2.1/unity-catalog/schemas"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &schemaInfo)
	return &schemaInfo, err
}

func (a *schemasImpl) Delete(ctx context.Context, request DeleteSchemaRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *schemasImpl) Get(ctx context.Context, request GetSchemaRequest) (*SchemaInfo, error) {
	var schemaInfo SchemaInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &schemaInfo)
	return &schemaInfo, err
}

// List schemas.
//
// Gets an array of schemas for a catalog in the metastore. If the caller is the
// metastore admin or the owner of the parent catalog, all schemas for the
// catalog will be retrieved. Otherwise, only schemas owned by the caller (or
// for which the caller has the **USE_SCHEMA** privilege) will be retrieved.
// There is no guarantee of a specific ordering of the elements in the array.
func (a *schemasImpl) List(ctx context.Context, request ListSchemasRequest) listing.Iterator[SchemaInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListSchemasRequest) (*ListSchemasResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListSchemasResponse) []SchemaInfo {
		return resp.Schemas
	}
	getNextReq := func(resp *ListSchemasResponse) *ListSchemasRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List schemas.
//
// Gets an array of schemas for a catalog in the metastore. If the caller is the
// metastore admin or the owner of the parent catalog, all schemas for the
// catalog will be retrieved. Otherwise, only schemas owned by the caller (or
// for which the caller has the **USE_SCHEMA** privilege) will be retrieved.
// There is no guarantee of a specific ordering of the elements in the array.
func (a *schemasImpl) ListAll(ctx context.Context, request ListSchemasRequest) ([]SchemaInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[SchemaInfo](ctx, iterator)
}

func (a *schemasImpl) internalList(ctx context.Context, request ListSchemasRequest) (*ListSchemasResponse, error) {
	var listSchemasResponse ListSchemasResponse
	path := "/api/2.1/unity-catalog/schemas"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSchemasResponse)
	return &listSchemasResponse, err
}

func (a *schemasImpl) Update(ctx context.Context, request UpdateSchema) (*SchemaInfo, error) {
	var schemaInfo SchemaInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &schemaInfo)
	return &schemaInfo, err
}

// unexported type that holds implementations of just StorageCredentials API methods
type storageCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *storageCredentialsImpl) Create(ctx context.Context, request CreateStorageCredential) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := "/api/2.1/unity-catalog/storage-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

func (a *storageCredentialsImpl) Delete(ctx context.Context, request DeleteStorageCredentialRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *storageCredentialsImpl) Get(ctx context.Context, request GetStorageCredentialRequest) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

// List credentials.
//
// Gets an array of storage credentials (as __StorageCredentialInfo__ objects).
// The array is limited to only those storage credentials the caller has
// permission to access. If the caller is a metastore admin, retrieval of
// credentials is unrestricted. There is no guarantee of a specific ordering of
// the elements in the array.
func (a *storageCredentialsImpl) List(ctx context.Context, request ListStorageCredentialsRequest) listing.Iterator[StorageCredentialInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListStorageCredentialsRequest) (*ListStorageCredentialsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListStorageCredentialsResponse) []StorageCredentialInfo {
		return resp.StorageCredentials
	}
	getNextReq := func(resp *ListStorageCredentialsResponse) *ListStorageCredentialsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List credentials.
//
// Gets an array of storage credentials (as __StorageCredentialInfo__ objects).
// The array is limited to only those storage credentials the caller has
// permission to access. If the caller is a metastore admin, retrieval of
// credentials is unrestricted. There is no guarantee of a specific ordering of
// the elements in the array.
func (a *storageCredentialsImpl) ListAll(ctx context.Context, request ListStorageCredentialsRequest) ([]StorageCredentialInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[StorageCredentialInfo](ctx, iterator)
}

func (a *storageCredentialsImpl) internalList(ctx context.Context, request ListStorageCredentialsRequest) (*ListStorageCredentialsResponse, error) {
	var listStorageCredentialsResponse ListStorageCredentialsResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listStorageCredentialsResponse)
	return &listStorageCredentialsResponse, err
}

func (a *storageCredentialsImpl) Update(ctx context.Context, request UpdateStorageCredential) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

func (a *storageCredentialsImpl) Validate(ctx context.Context, request ValidateStorageCredential) (*ValidateStorageCredentialResponse, error) {
	var validateStorageCredentialResponse ValidateStorageCredentialResponse
	path := "/api/2.1/unity-catalog/validate-storage-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &validateStorageCredentialResponse)
	return &validateStorageCredentialResponse, err
}

// unexported type that holds implementations of just SystemSchemas API methods
type systemSchemasImpl struct {
	client *client.DatabricksClient
}

func (a *systemSchemasImpl) Disable(ctx context.Context, request DisableRequest) error {
	var disableResponse DisableResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v/systemschemas/%v", request.MetastoreId, request.SchemaName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &disableResponse)
	return err
}

func (a *systemSchemasImpl) Enable(ctx context.Context, request EnableRequest) error {
	var enableResponse EnableResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v/systemschemas/%v", request.MetastoreId, request.SchemaName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, nil, &enableResponse)
	return err
}

// List system schemas.
//
// Gets an array of system schemas for a metastore. The caller must be an
// account admin or a metastore admin.
func (a *systemSchemasImpl) List(ctx context.Context, request ListSystemSchemasRequest) listing.Iterator[SystemSchemaInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListSystemSchemasRequest) (*ListSystemSchemasResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListSystemSchemasResponse) []SystemSchemaInfo {
		return resp.Schemas
	}
	getNextReq := func(resp *ListSystemSchemasResponse) *ListSystemSchemasRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List system schemas.
//
// Gets an array of system schemas for a metastore. The caller must be an
// account admin or a metastore admin.
func (a *systemSchemasImpl) ListAll(ctx context.Context, request ListSystemSchemasRequest) ([]SystemSchemaInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[SystemSchemaInfo](ctx, iterator)
}

func (a *systemSchemasImpl) internalList(ctx context.Context, request ListSystemSchemasRequest) (*ListSystemSchemasResponse, error) {
	var listSystemSchemasResponse ListSystemSchemasResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v/systemschemas", request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSystemSchemasResponse)
	return &listSystemSchemasResponse, err
}

// unexported type that holds implementations of just TableConstraints API methods
type tableConstraintsImpl struct {
	client *client.DatabricksClient
}

func (a *tableConstraintsImpl) Create(ctx context.Context, request CreateTableConstraint) (*TableConstraint, error) {
	var tableConstraint TableConstraint
	path := "/api/2.1/unity-catalog/constraints"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &tableConstraint)
	return &tableConstraint, err
}

func (a *tableConstraintsImpl) Delete(ctx context.Context, request DeleteTableConstraintRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/constraints/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

// unexported type that holds implementations of just Tables API methods
type tablesImpl struct {
	client *client.DatabricksClient
}

func (a *tablesImpl) Delete(ctx context.Context, request DeleteTableRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *tablesImpl) Exists(ctx context.Context, request ExistsRequest) (*TableExistsResponse, error) {
	var tableExistsResponse TableExistsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/exists", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &tableExistsResponse)
	return &tableExistsResponse, err
}

func (a *tablesImpl) Get(ctx context.Context, request GetTableRequest) (*TableInfo, error) {
	var tableInfo TableInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &tableInfo)
	return &tableInfo, err
}

// List tables.
//
// Gets an array of all tables for the current metastore under the parent
// catalog and schema. The caller must be a metastore admin or an owner of (or
// have the **SELECT** privilege on) the table. For the latter case, the caller
// must also be the owner or have the **USE_CATALOG** privilege on the parent
// catalog and the **USE_SCHEMA** privilege on the parent schema. There is no
// guarantee of a specific ordering of the elements in the array.
func (a *tablesImpl) List(ctx context.Context, request ListTablesRequest) listing.Iterator[TableInfo] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListTablesRequest) (*ListTablesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListTablesResponse) []TableInfo {
		return resp.Tables
	}
	getNextReq := func(resp *ListTablesResponse) *ListTablesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List tables.
//
// Gets an array of all tables for the current metastore under the parent
// catalog and schema. The caller must be a metastore admin or an owner of (or
// have the **SELECT** privilege on) the table. For the latter case, the caller
// must also be the owner or have the **USE_CATALOG** privilege on the parent
// catalog and the **USE_SCHEMA** privilege on the parent schema. There is no
// guarantee of a specific ordering of the elements in the array.
func (a *tablesImpl) ListAll(ctx context.Context, request ListTablesRequest) ([]TableInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[TableInfo](ctx, iterator)
}

func (a *tablesImpl) internalList(ctx context.Context, request ListTablesRequest) (*ListTablesResponse, error) {
	var listTablesResponse ListTablesResponse
	path := "/api/2.1/unity-catalog/tables"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listTablesResponse)
	return &listTablesResponse, err
}

// List table summaries.
//
// Gets an array of summaries for tables for a schema and catalog within the
// metastore. The table summaries returned are either:
//
// * summaries for tables (within the current metastore and parent catalog and
// schema), when the user is a metastore admin, or: * summaries for tables and
// schemas (within the current metastore and parent catalog) for which the user
// has ownership or the **SELECT** privilege on the table and ownership or
// **USE_SCHEMA** privilege on the schema, provided that the user also has
// ownership or the **USE_CATALOG** privilege on the parent catalog.
//
// There is no guarantee of a specific ordering of the elements in the array.
func (a *tablesImpl) ListSummaries(ctx context.Context, request ListSummariesRequest) listing.Iterator[TableSummary] {

	request.ForceSendFields = append(request.ForceSendFields, "MaxResults")

	getNextPage := func(ctx context.Context, req ListSummariesRequest) (*ListTableSummariesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSummaries(ctx, req)
	}
	getItems := func(resp *ListTableSummariesResponse) []TableSummary {
		return resp.Tables
	}
	getNextReq := func(resp *ListTableSummariesResponse) *ListSummariesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List table summaries.
//
// Gets an array of summaries for tables for a schema and catalog within the
// metastore. The table summaries returned are either:
//
// * summaries for tables (within the current metastore and parent catalog and
// schema), when the user is a metastore admin, or: * summaries for tables and
// schemas (within the current metastore and parent catalog) for which the user
// has ownership or the **SELECT** privilege on the table and ownership or
// **USE_SCHEMA** privilege on the schema, provided that the user also has
// ownership or the **USE_CATALOG** privilege on the parent catalog.
//
// There is no guarantee of a specific ordering of the elements in the array.
func (a *tablesImpl) ListSummariesAll(ctx context.Context, request ListSummariesRequest) ([]TableSummary, error) {
	iterator := a.ListSummaries(ctx, request)
	return listing.ToSlice[TableSummary](ctx, iterator)
}

func (a *tablesImpl) internalListSummaries(ctx context.Context, request ListSummariesRequest) (*ListTableSummariesResponse, error) {
	var listTableSummariesResponse ListTableSummariesResponse
	path := "/api/2.1/unity-catalog/table-summaries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listTableSummariesResponse)
	return &listTableSummariesResponse, err
}

func (a *tablesImpl) Update(ctx context.Context, request UpdateTableRequest) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just TemporaryTableCredentials API methods
type temporaryTableCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *temporaryTableCredentialsImpl) GenerateTemporaryTableCredentials(ctx context.Context, request GenerateTemporaryTableCredentialRequest) (*GenerateTemporaryTableCredentialResponse, error) {
	var generateTemporaryTableCredentialResponse GenerateTemporaryTableCredentialResponse
	path := "/api/2.0/unity-catalog/temporary-table-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &generateTemporaryTableCredentialResponse)
	return &generateTemporaryTableCredentialResponse, err
}

// unexported type that holds implementations of just Volumes API methods
type volumesImpl struct {
	client *client.DatabricksClient
}

func (a *volumesImpl) Create(ctx context.Context, request CreateVolumeRequestContent) (*VolumeInfo, error) {
	var volumeInfo VolumeInfo
	path := "/api/2.1/unity-catalog/volumes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &volumeInfo)
	return &volumeInfo, err
}

func (a *volumesImpl) Delete(ctx context.Context, request DeleteVolumeRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

// List Volumes.
//
// Gets an array of volumes for the current metastore under the parent catalog
// and schema.
//
// The returned volumes are filtered based on the privileges of the calling
// user. For example, the metastore admin is able to list all the volumes. A
// regular user needs to be the owner or have the **READ VOLUME** privilege on
// the volume to recieve the volumes in the response. For the latter case, the
// caller must also be the owner or have the **USE_CATALOG** privilege on the
// parent catalog and the **USE_SCHEMA** privilege on the parent schema.
//
// There is no guarantee of a specific ordering of the elements in the array.
func (a *volumesImpl) List(ctx context.Context, request ListVolumesRequest) listing.Iterator[VolumeInfo] {

	getNextPage := func(ctx context.Context, req ListVolumesRequest) (*ListVolumesResponseContent, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListVolumesResponseContent) []VolumeInfo {
		return resp.Volumes
	}
	getNextReq := func(resp *ListVolumesResponseContent) *ListVolumesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List Volumes.
//
// Gets an array of volumes for the current metastore under the parent catalog
// and schema.
//
// The returned volumes are filtered based on the privileges of the calling
// user. For example, the metastore admin is able to list all the volumes. A
// regular user needs to be the owner or have the **READ VOLUME** privilege on
// the volume to recieve the volumes in the response. For the latter case, the
// caller must also be the owner or have the **USE_CATALOG** privilege on the
// parent catalog and the **USE_SCHEMA** privilege on the parent schema.
//
// There is no guarantee of a specific ordering of the elements in the array.
func (a *volumesImpl) ListAll(ctx context.Context, request ListVolumesRequest) ([]VolumeInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[VolumeInfo](ctx, iterator)
}

func (a *volumesImpl) internalList(ctx context.Context, request ListVolumesRequest) (*ListVolumesResponseContent, error) {
	var listVolumesResponseContent ListVolumesResponseContent
	path := "/api/2.1/unity-catalog/volumes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listVolumesResponseContent)
	return &listVolumesResponseContent, err
}

func (a *volumesImpl) Read(ctx context.Context, request ReadVolumeRequest) (*VolumeInfo, error) {
	var volumeInfo VolumeInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &volumeInfo)
	return &volumeInfo, err
}

func (a *volumesImpl) Update(ctx context.Context, request UpdateVolumeRequestContent) (*VolumeInfo, error) {
	var volumeInfo VolumeInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &volumeInfo)
	return &volumeInfo, err
}

// unexported type that holds implementations of just WorkspaceBindings API methods
type workspaceBindingsImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceBindingsImpl) Get(ctx context.Context, request GetWorkspaceBindingRequest) (*CurrentWorkspaceBindings, error) {
	var currentWorkspaceBindings CurrentWorkspaceBindings
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspace-bindings/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &currentWorkspaceBindings)
	return &currentWorkspaceBindings, err
}

// Get securable workspace bindings.
//
// Gets workspace bindings of the securable. The caller must be a metastore
// admin or an owner of the securable.
func (a *workspaceBindingsImpl) GetBindings(ctx context.Context, request GetBindingsRequest) listing.Iterator[WorkspaceBinding] {

	getNextPage := func(ctx context.Context, req GetBindingsRequest) (*WorkspaceBindingsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalGetBindings(ctx, req)
	}
	getItems := func(resp *WorkspaceBindingsResponse) []WorkspaceBinding {
		return resp.Bindings
	}
	getNextReq := func(resp *WorkspaceBindingsResponse) *GetBindingsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Get securable workspace bindings.
//
// Gets workspace bindings of the securable. The caller must be a metastore
// admin or an owner of the securable.
func (a *workspaceBindingsImpl) GetBindingsAll(ctx context.Context, request GetBindingsRequest) ([]WorkspaceBinding, error) {
	iterator := a.GetBindings(ctx, request)
	return listing.ToSlice[WorkspaceBinding](ctx, iterator)
}

func (a *workspaceBindingsImpl) internalGetBindings(ctx context.Context, request GetBindingsRequest) (*WorkspaceBindingsResponse, error) {
	var workspaceBindingsResponse WorkspaceBindingsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/bindings/%v/%v", request.SecurableType, request.SecurableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspaceBindingsResponse)
	return &workspaceBindingsResponse, err
}

func (a *workspaceBindingsImpl) Update(ctx context.Context, request UpdateWorkspaceBindings) (*CurrentWorkspaceBindings, error) {
	var currentWorkspaceBindings CurrentWorkspaceBindings
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspace-bindings/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &currentWorkspaceBindings)
	return &currentWorkspaceBindings, err
}

func (a *workspaceBindingsImpl) UpdateBindings(ctx context.Context, request UpdateWorkspaceBindingsParameters) (*WorkspaceBindingsResponse, error) {
	var workspaceBindingsResponse WorkspaceBindingsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/bindings/%v/%v", request.SecurableType, request.SecurableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &workspaceBindingsResponse)
	return &workspaceBindingsResponse, err
}

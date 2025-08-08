// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/catalog/catalogpb"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just AccountMetastoreAssignments API methods
type accountMetastoreAssignmentsImpl struct {
	client *client.DatabricksClient
}

func (a *accountMetastoreAssignmentsImpl) Create(ctx context.Context, request AccountsCreateMetastoreAssignment) error {
	requestPb, pbErr := AccountsCreateMetastoreAssignmentToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *accountMetastoreAssignmentsImpl) Delete(ctx context.Context, request DeleteAccountMetastoreAssignmentRequest) error {
	requestPb, pbErr := DeleteAccountMetastoreAssignmentRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *accountMetastoreAssignmentsImpl) Get(ctx context.Context, request GetAccountMetastoreAssignmentRequest) (*AccountsMetastoreAssignment, error) {
	requestPb, pbErr := GetAccountMetastoreAssignmentRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountsMetastoreAssignmentPb catalogpb.AccountsMetastoreAssignmentPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastore", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&accountsMetastoreAssignmentPb,
	)
	resp, err := AccountsMetastoreAssignmentFromPb(&accountsMetastoreAssignmentPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets a list of all Databricks workspace IDs that have been assigned to given
// metastore.
func (a *accountMetastoreAssignmentsImpl) ListAll(ctx context.Context, request ListAccountMetastoreAssignmentsRequest) ([]int64, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[int64](ctx, iterator)
}

func (a *accountMetastoreAssignmentsImpl) internalList(ctx context.Context, request ListAccountMetastoreAssignmentsRequest) (*ListAccountMetastoreAssignmentsResponse, error) {
	requestPb, pbErr := ListAccountMetastoreAssignmentsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAccountMetastoreAssignmentsResponsePb catalogpb.ListAccountMetastoreAssignmentsResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/workspaces", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listAccountMetastoreAssignmentsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListAccountMetastoreAssignmentsResponseFromPb(&listAccountMetastoreAssignmentsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountMetastoreAssignmentsImpl) Update(ctx context.Context, request AccountsUpdateMetastoreAssignment) error {
	requestPb, pbErr := AccountsUpdateMetastoreAssignmentToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just AccountMetastores API methods
type accountMetastoresImpl struct {
	client *client.DatabricksClient
}

func (a *accountMetastoresImpl) Create(ctx context.Context, request AccountsCreateMetastore) (*AccountsMetastoreInfo, error) {
	requestPb, pbErr := AccountsCreateMetastoreToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountsMetastoreInfoPb catalogpb.AccountsMetastoreInfoPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&accountsMetastoreInfoPb,
	)
	resp, err := AccountsMetastoreInfoFromPb(&accountsMetastoreInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountMetastoresImpl) Delete(ctx context.Context, request DeleteAccountMetastoreRequest) error {
	requestPb, pbErr := DeleteAccountMetastoreRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *accountMetastoresImpl) Get(ctx context.Context, request GetAccountMetastoreRequest) (*AccountsMetastoreInfo, error) {
	requestPb, pbErr := GetAccountMetastoreRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountsMetastoreInfoPb catalogpb.AccountsMetastoreInfoPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&accountsMetastoreInfoPb,
	)
	resp, err := AccountsMetastoreInfoFromPb(&accountsMetastoreInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets all Unity Catalog metastores associated with an account specified by ID.
func (a *accountMetastoresImpl) ListAll(ctx context.Context) ([]MetastoreInfo, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[MetastoreInfo](ctx, iterator)
}

func (a *accountMetastoresImpl) internalList(ctx context.Context) (*ListMetastoresResponse, error) {

	var listMetastoresResponsePb catalogpb.ListMetastoresResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&listMetastoresResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListMetastoresResponseFromPb(&listMetastoresResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountMetastoresImpl) Update(ctx context.Context, request AccountsUpdateMetastore) (*AccountsMetastoreInfo, error) {
	requestPb, pbErr := AccountsUpdateMetastoreToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountsMetastoreInfoPb catalogpb.AccountsMetastoreInfoPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&accountsMetastoreInfoPb,
	)
	resp, err := AccountsMetastoreInfoFromPb(&accountsMetastoreInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just AccountStorageCredentials API methods
type accountStorageCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *accountStorageCredentialsImpl) Create(ctx context.Context, request AccountsCreateStorageCredential) (*AccountsStorageCredentialInfo, error) {
	requestPb, pbErr := AccountsCreateStorageCredentialToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountsStorageCredentialInfoPb catalogpb.AccountsStorageCredentialInfoPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&accountsStorageCredentialInfoPb,
	)
	resp, err := AccountsStorageCredentialInfoFromPb(&accountsStorageCredentialInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountStorageCredentialsImpl) Delete(ctx context.Context, request DeleteAccountStorageCredentialRequest) error {
	requestPb, pbErr := DeleteAccountStorageCredentialRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/%v", a.client.ConfiguredAccountID(), request.MetastoreId, request.StorageCredentialName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *accountStorageCredentialsImpl) Get(ctx context.Context, request GetAccountStorageCredentialRequest) (*AccountsStorageCredentialInfo, error) {
	requestPb, pbErr := GetAccountStorageCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountsStorageCredentialInfoPb catalogpb.AccountsStorageCredentialInfoPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/%v", a.client.ConfiguredAccountID(), request.MetastoreId, request.StorageCredentialName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&accountsStorageCredentialInfoPb,
	)
	resp, err := AccountsStorageCredentialInfoFromPb(&accountsStorageCredentialInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Gets a list of all storage credentials that have been assigned to given
// metastore.
func (a *accountStorageCredentialsImpl) ListAll(ctx context.Context, request ListAccountStorageCredentialsRequest) ([]StorageCredentialInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[StorageCredentialInfo](ctx, iterator)
}

func (a *accountStorageCredentialsImpl) internalList(ctx context.Context, request ListAccountStorageCredentialsRequest) (*ListAccountStorageCredentialsResponse, error) {
	requestPb, pbErr := ListAccountStorageCredentialsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listAccountStorageCredentialsResponsePb catalogpb.ListAccountStorageCredentialsResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials", a.client.ConfiguredAccountID(), request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listAccountStorageCredentialsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListAccountStorageCredentialsResponseFromPb(&listAccountStorageCredentialsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *accountStorageCredentialsImpl) Update(ctx context.Context, request AccountsUpdateStorageCredential) (*AccountsStorageCredentialInfo, error) {
	requestPb, pbErr := AccountsUpdateStorageCredentialToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var accountsStorageCredentialInfoPb catalogpb.AccountsStorageCredentialInfoPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/%v", a.client.ConfiguredAccountID(), request.MetastoreId, request.StorageCredentialName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&accountsStorageCredentialInfoPb,
	)
	resp, err := AccountsStorageCredentialInfoFromPb(&accountsStorageCredentialInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ArtifactAllowlists API methods
type artifactAllowlistsImpl struct {
	client *client.DatabricksClient
}

func (a *artifactAllowlistsImpl) Get(ctx context.Context, request GetArtifactAllowlistRequest) (*ArtifactAllowlistInfo, error) {
	requestPb, pbErr := GetArtifactAllowlistRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var artifactAllowlistInfoPb catalogpb.ArtifactAllowlistInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/artifact-allowlists/%v", request.ArtifactType)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&artifactAllowlistInfoPb,
	)
	resp, err := ArtifactAllowlistInfoFromPb(&artifactAllowlistInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *artifactAllowlistsImpl) Update(ctx context.Context, request SetArtifactAllowlist) (*ArtifactAllowlistInfo, error) {
	requestPb, pbErr := SetArtifactAllowlistToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var artifactAllowlistInfoPb catalogpb.ArtifactAllowlistInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/artifact-allowlists/%v", request.ArtifactType)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&artifactAllowlistInfoPb,
	)
	resp, err := ArtifactAllowlistInfoFromPb(&artifactAllowlistInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Catalogs API methods
type catalogsImpl struct {
	client *client.DatabricksClient
}

func (a *catalogsImpl) Create(ctx context.Context, request CreateCatalog) (*CatalogInfo, error) {
	requestPb, pbErr := CreateCatalogToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var catalogInfoPb catalogpb.CatalogInfoPb
	path := "/api/2.1/unity-catalog/catalogs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&catalogInfoPb,
	)
	resp, err := CatalogInfoFromPb(&catalogInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *catalogsImpl) Delete(ctx context.Context, request DeleteCatalogRequest) error {
	requestPb, pbErr := DeleteCatalogRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *catalogsImpl) Get(ctx context.Context, request GetCatalogRequest) (*CatalogInfo, error) {
	requestPb, pbErr := GetCatalogRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var catalogInfoPb catalogpb.CatalogInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&catalogInfoPb,
	)
	resp, err := CatalogInfoFromPb(&catalogInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets an array of catalogs in the metastore. If the caller is the metastore
// admin, all catalogs will be retrieved. Otherwise, only catalogs owned by the
// caller (or for which the caller has the **USE_CATALOG** privilege) will be
// retrieved. There is no guarantee of a specific ordering of the elements in
// the array.
func (a *catalogsImpl) List(ctx context.Context, request ListCatalogsRequest) listing.Iterator[CatalogInfo] {

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
	requestPb, pbErr := ListCatalogsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listCatalogsResponsePb catalogpb.ListCatalogsResponsePb
	path := "/api/2.1/unity-catalog/catalogs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listCatalogsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListCatalogsResponseFromPb(&listCatalogsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *catalogsImpl) Update(ctx context.Context, request UpdateCatalog) (*CatalogInfo, error) {
	requestPb, pbErr := UpdateCatalogToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var catalogInfoPb catalogpb.CatalogInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&catalogInfoPb,
	)
	resp, err := CatalogInfoFromPb(&catalogInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Connections API methods
type connectionsImpl struct {
	client *client.DatabricksClient
}

func (a *connectionsImpl) Create(ctx context.Context, request CreateConnection) (*ConnectionInfo, error) {
	requestPb, pbErr := CreateConnectionToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var connectionInfoPb catalogpb.ConnectionInfoPb
	path := "/api/2.1/unity-catalog/connections"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&connectionInfoPb,
	)
	resp, err := ConnectionInfoFromPb(&connectionInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *connectionsImpl) Delete(ctx context.Context, request DeleteConnectionRequest) error {
	requestPb, pbErr := DeleteConnectionRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/connections/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *connectionsImpl) Get(ctx context.Context, request GetConnectionRequest) (*ConnectionInfo, error) {
	requestPb, pbErr := GetConnectionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var connectionInfoPb catalogpb.ConnectionInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/connections/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&connectionInfoPb,
	)
	resp, err := ConnectionInfoFromPb(&connectionInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List all connections.
func (a *connectionsImpl) List(ctx context.Context, request ListConnectionsRequest) listing.Iterator[ConnectionInfo] {

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

// List all connections.
func (a *connectionsImpl) ListAll(ctx context.Context, request ListConnectionsRequest) ([]ConnectionInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[ConnectionInfo](ctx, iterator)
}

func (a *connectionsImpl) internalList(ctx context.Context, request ListConnectionsRequest) (*ListConnectionsResponse, error) {
	requestPb, pbErr := ListConnectionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listConnectionsResponsePb catalogpb.ListConnectionsResponsePb
	path := "/api/2.1/unity-catalog/connections"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listConnectionsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListConnectionsResponseFromPb(&listConnectionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *connectionsImpl) Update(ctx context.Context, request UpdateConnection) (*ConnectionInfo, error) {
	requestPb, pbErr := UpdateConnectionToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var connectionInfoPb catalogpb.ConnectionInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/connections/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&connectionInfoPb,
	)
	resp, err := ConnectionInfoFromPb(&connectionInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Credentials API methods
type credentialsImpl struct {
	client *client.DatabricksClient
}

func (a *credentialsImpl) CreateCredential(ctx context.Context, request CreateCredentialRequest) (*CredentialInfo, error) {
	requestPb, pbErr := CreateCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var credentialInfoPb catalogpb.CredentialInfoPb
	path := "/api/2.1/unity-catalog/credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&credentialInfoPb,
	)
	resp, err := CredentialInfoFromPb(&credentialInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *credentialsImpl) DeleteCredential(ctx context.Context, request DeleteCredentialRequest) error {
	requestPb, pbErr := DeleteCredentialRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/credentials/%v", request.NameArg)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *credentialsImpl) GenerateTemporaryServiceCredential(ctx context.Context, request GenerateTemporaryServiceCredentialRequest) (*TemporaryCredentials, error) {
	requestPb, pbErr := GenerateTemporaryServiceCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var temporaryCredentialsPb catalogpb.TemporaryCredentialsPb
	path := "/api/2.1/unity-catalog/temporary-service-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&temporaryCredentialsPb,
	)
	resp, err := TemporaryCredentialsFromPb(&temporaryCredentialsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *credentialsImpl) GetCredential(ctx context.Context, request GetCredentialRequest) (*CredentialInfo, error) {
	requestPb, pbErr := GetCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var credentialInfoPb catalogpb.CredentialInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/credentials/%v", request.NameArg)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&credentialInfoPb,
	)
	resp, err := CredentialInfoFromPb(&credentialInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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
	requestPb, pbErr := ListCredentialsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listCredentialsResponsePb catalogpb.ListCredentialsResponsePb
	path := "/api/2.1/unity-catalog/credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listCredentialsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListCredentialsResponseFromPb(&listCredentialsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *credentialsImpl) UpdateCredential(ctx context.Context, request UpdateCredentialRequest) (*CredentialInfo, error) {
	requestPb, pbErr := UpdateCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var credentialInfoPb catalogpb.CredentialInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/credentials/%v", request.NameArg)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&credentialInfoPb,
	)
	resp, err := CredentialInfoFromPb(&credentialInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *credentialsImpl) ValidateCredential(ctx context.Context, request ValidateCredentialRequest) (*ValidateCredentialResponse, error) {
	requestPb, pbErr := ValidateCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var validateCredentialResponsePb catalogpb.ValidateCredentialResponsePb
	path := "/api/2.1/unity-catalog/validate-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&validateCredentialResponsePb,
	)
	resp, err := ValidateCredentialResponseFromPb(&validateCredentialResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ExternalLineage API methods
type externalLineageImpl struct {
	client *client.DatabricksClient
}

func (a *externalLineageImpl) CreateExternalLineageRelationship(ctx context.Context, request CreateExternalLineageRelationshipRequest) (*ExternalLineageRelationship, error) {
	requestPb, pbErr := CreateExternalLineageRelationshipRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var externalLineageRelationshipPb catalogpb.ExternalLineageRelationshipPb
	path := "/api/2.0/lineage-tracking/external-lineage"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.ExternalLineageRelationship,
		&externalLineageRelationshipPb,
	)
	resp, err := ExternalLineageRelationshipFromPb(&externalLineageRelationshipPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *externalLineageImpl) DeleteExternalLineageRelationship(ctx context.Context, request DeleteExternalLineageRelationshipRequest) error {
	requestPb, pbErr := DeleteExternalLineageRelationshipRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := "/api/2.0/lineage-tracking/external-lineage"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// Lists external lineage relationships of a Databricks object or external
// metadata given a supplied direction.
func (a *externalLineageImpl) ListExternalLineageRelationships(ctx context.Context, request ListExternalLineageRelationshipsRequest) listing.Iterator[ExternalLineageInfo] {

	getNextPage := func(ctx context.Context, req ListExternalLineageRelationshipsRequest) (*ListExternalLineageRelationshipsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListExternalLineageRelationships(ctx, req)
	}
	getItems := func(resp *ListExternalLineageRelationshipsResponse) []ExternalLineageInfo {
		return resp.ExternalLineageRelationships
	}
	getNextReq := func(resp *ListExternalLineageRelationshipsResponse) *ListExternalLineageRelationshipsRequest {
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

// Lists external lineage relationships of a Databricks object or external
// metadata given a supplied direction.
func (a *externalLineageImpl) ListExternalLineageRelationshipsAll(ctx context.Context, request ListExternalLineageRelationshipsRequest) ([]ExternalLineageInfo, error) {
	iterator := a.ListExternalLineageRelationships(ctx, request)
	return listing.ToSlice[ExternalLineageInfo](ctx, iterator)
}

func (a *externalLineageImpl) internalListExternalLineageRelationships(ctx context.Context, request ListExternalLineageRelationshipsRequest) (*ListExternalLineageRelationshipsResponse, error) {
	requestPb, pbErr := ListExternalLineageRelationshipsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listExternalLineageRelationshipsResponsePb catalogpb.ListExternalLineageRelationshipsResponsePb
	path := "/api/2.0/lineage-tracking/external-lineage"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listExternalLineageRelationshipsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListExternalLineageRelationshipsResponseFromPb(&listExternalLineageRelationshipsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *externalLineageImpl) UpdateExternalLineageRelationship(ctx context.Context, request UpdateExternalLineageRelationshipRequest) (*ExternalLineageRelationship, error) {
	requestPb, pbErr := UpdateExternalLineageRelationshipRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var externalLineageRelationshipPb catalogpb.ExternalLineageRelationshipPb
	path := "/api/2.0/lineage-tracking/external-lineage"
	queryParams := make(map[string]any)
	if requestPb.UpdateMask != "" {
		queryParams["update_mask"] = requestPb.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb.ExternalLineageRelationship,
		&externalLineageRelationshipPb,
	)
	resp, err := ExternalLineageRelationshipFromPb(&externalLineageRelationshipPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ExternalLocations API methods
type externalLocationsImpl struct {
	client *client.DatabricksClient
}

func (a *externalLocationsImpl) Create(ctx context.Context, request CreateExternalLocation) (*ExternalLocationInfo, error) {
	requestPb, pbErr := CreateExternalLocationToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var externalLocationInfoPb catalogpb.ExternalLocationInfoPb
	path := "/api/2.1/unity-catalog/external-locations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&externalLocationInfoPb,
	)
	resp, err := ExternalLocationInfoFromPb(&externalLocationInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *externalLocationsImpl) Delete(ctx context.Context, request DeleteExternalLocationRequest) error {
	requestPb, pbErr := DeleteExternalLocationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *externalLocationsImpl) Get(ctx context.Context, request GetExternalLocationRequest) (*ExternalLocationInfo, error) {
	requestPb, pbErr := GetExternalLocationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var externalLocationInfoPb catalogpb.ExternalLocationInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&externalLocationInfoPb,
	)
	resp, err := ExternalLocationInfoFromPb(&externalLocationInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets an array of external locations (__ExternalLocationInfo__ objects) from
// the metastore. The caller must be a metastore admin, the owner of the
// external location, or a user that has some privilege on the external
// location. There is no guarantee of a specific ordering of the elements in the
// array.
func (a *externalLocationsImpl) List(ctx context.Context, request ListExternalLocationsRequest) listing.Iterator[ExternalLocationInfo] {

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
	requestPb, pbErr := ListExternalLocationsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listExternalLocationsResponsePb catalogpb.ListExternalLocationsResponsePb
	path := "/api/2.1/unity-catalog/external-locations"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listExternalLocationsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListExternalLocationsResponseFromPb(&listExternalLocationsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *externalLocationsImpl) Update(ctx context.Context, request UpdateExternalLocation) (*ExternalLocationInfo, error) {
	requestPb, pbErr := UpdateExternalLocationToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var externalLocationInfoPb catalogpb.ExternalLocationInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&externalLocationInfoPb,
	)
	resp, err := ExternalLocationInfoFromPb(&externalLocationInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ExternalMetadata API methods
type externalMetadataImpl struct {
	client *client.DatabricksClient
}

func (a *externalMetadataImpl) CreateExternalMetadata(ctx context.Context, request CreateExternalMetadataRequest) (*ExternalMetadata, error) {
	requestPb, pbErr := CreateExternalMetadataRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var externalMetadataPb catalogpb.ExternalMetadataPb
	path := "/api/2.0/lineage-tracking/external-metadata"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.ExternalMetadata,
		&externalMetadataPb,
	)
	resp, err := ExternalMetadataFromPb(&externalMetadataPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *externalMetadataImpl) DeleteExternalMetadata(ctx context.Context, request DeleteExternalMetadataRequest) error {
	requestPb, pbErr := DeleteExternalMetadataRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/lineage-tracking/external-metadata/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *externalMetadataImpl) GetExternalMetadata(ctx context.Context, request GetExternalMetadataRequest) (*ExternalMetadata, error) {
	requestPb, pbErr := GetExternalMetadataRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var externalMetadataPb catalogpb.ExternalMetadataPb
	path := fmt.Sprintf("/api/2.0/lineage-tracking/external-metadata/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&externalMetadataPb,
	)
	resp, err := ExternalMetadataFromPb(&externalMetadataPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets an array of external metadata objects in the metastore. If the caller is
// the metastore admin, all external metadata objects will be retrieved.
// Otherwise, only external metadata objects that the caller has **BROWSE** on
// will be retrieved. There is no guarantee of a specific ordering of the
// elements in the array.
func (a *externalMetadataImpl) ListExternalMetadata(ctx context.Context, request ListExternalMetadataRequest) listing.Iterator[ExternalMetadata] {

	getNextPage := func(ctx context.Context, req ListExternalMetadataRequest) (*ListExternalMetadataResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListExternalMetadata(ctx, req)
	}
	getItems := func(resp *ListExternalMetadataResponse) []ExternalMetadata {
		return resp.ExternalMetadata
	}
	getNextReq := func(resp *ListExternalMetadataResponse) *ListExternalMetadataRequest {
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

// Gets an array of external metadata objects in the metastore. If the caller is
// the metastore admin, all external metadata objects will be retrieved.
// Otherwise, only external metadata objects that the caller has **BROWSE** on
// will be retrieved. There is no guarantee of a specific ordering of the
// elements in the array.
func (a *externalMetadataImpl) ListExternalMetadataAll(ctx context.Context, request ListExternalMetadataRequest) ([]ExternalMetadata, error) {
	iterator := a.ListExternalMetadata(ctx, request)
	return listing.ToSlice[ExternalMetadata](ctx, iterator)
}

func (a *externalMetadataImpl) internalListExternalMetadata(ctx context.Context, request ListExternalMetadataRequest) (*ListExternalMetadataResponse, error) {
	requestPb, pbErr := ListExternalMetadataRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listExternalMetadataResponsePb catalogpb.ListExternalMetadataResponsePb
	path := "/api/2.0/lineage-tracking/external-metadata"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listExternalMetadataResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListExternalMetadataResponseFromPb(&listExternalMetadataResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *externalMetadataImpl) UpdateExternalMetadata(ctx context.Context, request UpdateExternalMetadataRequest) (*ExternalMetadata, error) {
	requestPb, pbErr := UpdateExternalMetadataRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var externalMetadataPb catalogpb.ExternalMetadataPb
	path := fmt.Sprintf("/api/2.0/lineage-tracking/external-metadata/%v", request.Name)
	queryParams := make(map[string]any)
	if requestPb.UpdateMask != "" {
		queryParams["update_mask"] = requestPb.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb.ExternalMetadata,
		&externalMetadataPb,
	)
	resp, err := ExternalMetadataFromPb(&externalMetadataPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Functions API methods
type functionsImpl struct {
	client *client.DatabricksClient
}

func (a *functionsImpl) Create(ctx context.Context, request CreateFunctionRequest) (*FunctionInfo, error) {
	requestPb, pbErr := CreateFunctionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var functionInfoPb catalogpb.FunctionInfoPb
	path := "/api/2.1/unity-catalog/functions"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&functionInfoPb,
	)
	resp, err := FunctionInfoFromPb(&functionInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *functionsImpl) Delete(ctx context.Context, request DeleteFunctionRequest) error {
	requestPb, pbErr := DeleteFunctionRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *functionsImpl) Get(ctx context.Context, request GetFunctionRequest) (*FunctionInfo, error) {
	requestPb, pbErr := GetFunctionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var functionInfoPb catalogpb.FunctionInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&functionInfoPb,
	)
	resp, err := FunctionInfoFromPb(&functionInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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
	requestPb, pbErr := ListFunctionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listFunctionsResponsePb catalogpb.ListFunctionsResponsePb
	path := "/api/2.1/unity-catalog/functions"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listFunctionsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListFunctionsResponseFromPb(&listFunctionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *functionsImpl) Update(ctx context.Context, request UpdateFunction) (*FunctionInfo, error) {
	requestPb, pbErr := UpdateFunctionToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var functionInfoPb catalogpb.FunctionInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&functionInfoPb,
	)
	resp, err := FunctionInfoFromPb(&functionInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Grants API methods
type grantsImpl struct {
	client *client.DatabricksClient
}

func (a *grantsImpl) Get(ctx context.Context, request GetGrantRequest) (*GetPermissionsResponse, error) {
	requestPb, pbErr := GetGrantRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getPermissionsResponsePb catalogpb.GetPermissionsResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getPermissionsResponsePb,
	)
	resp, err := GetPermissionsResponseFromPb(&getPermissionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *grantsImpl) GetEffective(ctx context.Context, request GetEffectiveRequest) (*EffectivePermissionsList, error) {
	requestPb, pbErr := GetEffectiveRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var effectivePermissionsListPb catalogpb.EffectivePermissionsListPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/effective-permissions/%v/%v", request.SecurableType, request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&effectivePermissionsListPb,
	)
	resp, err := EffectivePermissionsListFromPb(&effectivePermissionsListPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *grantsImpl) Update(ctx context.Context, request UpdatePermissions) (*UpdatePermissionsResponse, error) {
	requestPb, pbErr := UpdatePermissionsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updatePermissionsResponsePb catalogpb.UpdatePermissionsResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&updatePermissionsResponsePb,
	)
	resp, err := UpdatePermissionsResponseFromPb(&updatePermissionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Metastores API methods
type metastoresImpl struct {
	client *client.DatabricksClient
}

func (a *metastoresImpl) Assign(ctx context.Context, request CreateMetastoreAssignment) error {
	requestPb, pbErr := CreateMetastoreAssignmentToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *metastoresImpl) Create(ctx context.Context, request CreateMetastore) (*MetastoreInfo, error) {
	requestPb, pbErr := CreateMetastoreToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var metastoreInfoPb catalogpb.MetastoreInfoPb
	path := "/api/2.1/unity-catalog/metastores"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&metastoreInfoPb,
	)
	resp, err := MetastoreInfoFromPb(&metastoreInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *metastoresImpl) Current(ctx context.Context) (*MetastoreAssignment, error) {

	var metastoreAssignmentPb catalogpb.MetastoreAssignmentPb
	path := "/api/2.1/unity-catalog/current-metastore-assignment"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&metastoreAssignmentPb,
	)
	resp, err := MetastoreAssignmentFromPb(&metastoreAssignmentPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *metastoresImpl) Delete(ctx context.Context, request DeleteMetastoreRequest) error {
	requestPb, pbErr := DeleteMetastoreRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *metastoresImpl) Get(ctx context.Context, request GetMetastoreRequest) (*MetastoreInfo, error) {
	requestPb, pbErr := GetMetastoreRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var metastoreInfoPb catalogpb.MetastoreInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&metastoreInfoPb,
	)
	resp, err := MetastoreInfoFromPb(&metastoreInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets an array of the available metastores (as __MetastoreInfo__ objects). The
// caller must be an admin to retrieve this info. There is no guarantee of a
// specific ordering of the elements in the array.
func (a *metastoresImpl) List(ctx context.Context, request ListMetastoresRequest) listing.Iterator[MetastoreInfo] {

	getNextPage := func(ctx context.Context, req ListMetastoresRequest) (*ListMetastoresResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListMetastoresResponse) []MetastoreInfo {
		return resp.Metastores
	}
	getNextReq := func(resp *ListMetastoresResponse) *ListMetastoresRequest {
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

// Gets an array of the available metastores (as __MetastoreInfo__ objects). The
// caller must be an admin to retrieve this info. There is no guarantee of a
// specific ordering of the elements in the array.
func (a *metastoresImpl) ListAll(ctx context.Context, request ListMetastoresRequest) ([]MetastoreInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[MetastoreInfo](ctx, iterator)
}

func (a *metastoresImpl) internalList(ctx context.Context, request ListMetastoresRequest) (*ListMetastoresResponse, error) {
	requestPb, pbErr := ListMetastoresRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listMetastoresResponsePb catalogpb.ListMetastoresResponsePb
	path := "/api/2.1/unity-catalog/metastores"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listMetastoresResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListMetastoresResponseFromPb(&listMetastoresResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *metastoresImpl) Summary(ctx context.Context) (*GetMetastoreSummaryResponse, error) {

	var getMetastoreSummaryResponsePb catalogpb.GetMetastoreSummaryResponsePb
	path := "/api/2.1/unity-catalog/metastore_summary"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&getMetastoreSummaryResponsePb,
	)
	resp, err := GetMetastoreSummaryResponseFromPb(&getMetastoreSummaryResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *metastoresImpl) Unassign(ctx context.Context, request UnassignRequest) error {
	requestPb, pbErr := UnassignRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *metastoresImpl) Update(ctx context.Context, request UpdateMetastore) (*MetastoreInfo, error) {
	requestPb, pbErr := UpdateMetastoreToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var metastoreInfoPb catalogpb.MetastoreInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&metastoreInfoPb,
	)
	resp, err := MetastoreInfoFromPb(&metastoreInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *metastoresImpl) UpdateAssignment(ctx context.Context, request UpdateMetastoreAssignment) error {
	requestPb, pbErr := UpdateMetastoreAssignmentToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just ModelVersions API methods
type modelVersionsImpl struct {
	client *client.DatabricksClient
}

func (a *modelVersionsImpl) Delete(ctx context.Context, request DeleteModelVersionRequest) error {
	requestPb, pbErr := DeleteModelVersionRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions/%v", request.FullName, request.Version)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *modelVersionsImpl) Get(ctx context.Context, request GetModelVersionRequest) (*ModelVersionInfo, error) {
	requestPb, pbErr := GetModelVersionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var modelVersionInfoPb catalogpb.ModelVersionInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions/%v", request.FullName, request.Version)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&modelVersionInfoPb,
	)
	resp, err := ModelVersionInfoFromPb(&modelVersionInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelVersionsImpl) GetByAlias(ctx context.Context, request GetByAliasRequest) (*ModelVersionInfo, error) {
	requestPb, pbErr := GetByAliasRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var modelVersionInfoPb catalogpb.ModelVersionInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/aliases/%v", request.FullName, request.Alias)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&modelVersionInfoPb,
	)
	resp, err := ModelVersionInfoFromPb(&modelVersionInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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
	requestPb, pbErr := ListModelVersionsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listModelVersionsResponsePb catalogpb.ListModelVersionsResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listModelVersionsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListModelVersionsResponseFromPb(&listModelVersionsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *modelVersionsImpl) Update(ctx context.Context, request UpdateModelVersionRequest) (*ModelVersionInfo, error) {
	requestPb, pbErr := UpdateModelVersionRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var modelVersionInfoPb catalogpb.ModelVersionInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions/%v", request.FullName, request.Version)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&modelVersionInfoPb,
	)
	resp, err := ModelVersionInfoFromPb(&modelVersionInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just OnlineTables API methods
type onlineTablesImpl struct {
	client *client.DatabricksClient
}

func (a *onlineTablesImpl) Create(ctx context.Context, request CreateOnlineTableRequest) (*OnlineTable, error) {
	requestPb, pbErr := CreateOnlineTableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var onlineTablePb catalogpb.OnlineTablePb
	path := "/api/2.0/online-tables"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.Table,
		&onlineTablePb,
	)
	resp, err := OnlineTableFromPb(&onlineTablePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *onlineTablesImpl) Delete(ctx context.Context, request DeleteOnlineTableRequest) error {
	requestPb, pbErr := DeleteOnlineTableRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/online-tables/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *onlineTablesImpl) Get(ctx context.Context, request GetOnlineTableRequest) (*OnlineTable, error) {
	requestPb, pbErr := GetOnlineTableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var onlineTablePb catalogpb.OnlineTablePb
	path := fmt.Sprintf("/api/2.0/online-tables/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&onlineTablePb,
	)
	resp, err := OnlineTableFromPb(&onlineTablePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just QualityMonitors API methods
type qualityMonitorsImpl struct {
	client *client.DatabricksClient
}

func (a *qualityMonitorsImpl) CancelRefresh(ctx context.Context, request CancelRefreshRequest) error {

	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes/%v/cancel", request.TableName, request.RefreshId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		nil,
	)

	return err
}

func (a *qualityMonitorsImpl) Create(ctx context.Context, request CreateMonitor) (*MonitorInfo, error) {
	requestPb, pbErr := CreateMonitorToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var monitorInfoPb catalogpb.MonitorInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&monitorInfoPb,
	)
	resp, err := MonitorInfoFromPb(&monitorInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *qualityMonitorsImpl) Delete(ctx context.Context, request DeleteQualityMonitorRequest) (*DeleteMonitorResponse, error) {
	requestPb, pbErr := DeleteQualityMonitorRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteMonitorResponsePb catalogpb.DeleteMonitorResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteMonitorResponsePb,
	)
	resp, err := DeleteMonitorResponseFromPb(&deleteMonitorResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *qualityMonitorsImpl) Get(ctx context.Context, request GetQualityMonitorRequest) (*MonitorInfo, error) {
	requestPb, pbErr := GetQualityMonitorRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var monitorInfoPb catalogpb.MonitorInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&monitorInfoPb,
	)
	resp, err := MonitorInfoFromPb(&monitorInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *qualityMonitorsImpl) GetRefresh(ctx context.Context, request GetRefreshRequest) (*MonitorRefreshInfo, error) {
	requestPb, pbErr := GetRefreshRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var monitorRefreshInfoPb catalogpb.MonitorRefreshInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes/%v", request.TableName, request.RefreshId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&monitorRefreshInfoPb,
	)
	resp, err := MonitorRefreshInfoFromPb(&monitorRefreshInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *qualityMonitorsImpl) ListRefreshes(ctx context.Context, request ListRefreshesRequest) (*MonitorRefreshListResponse, error) {
	requestPb, pbErr := ListRefreshesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var monitorRefreshListResponsePb catalogpb.MonitorRefreshListResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&monitorRefreshListResponsePb,
	)
	resp, err := MonitorRefreshListResponseFromPb(&monitorRefreshListResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *qualityMonitorsImpl) RegenerateDashboard(ctx context.Context, request RegenerateDashboardRequest) (*RegenerateDashboardResponse, error) {
	requestPb, pbErr := RegenerateDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var regenerateDashboardResponsePb catalogpb.RegenerateDashboardResponsePb
	path := fmt.Sprintf("/api/2.1/quality-monitoring/tables/%v/monitor/dashboard", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&regenerateDashboardResponsePb,
	)
	resp, err := RegenerateDashboardResponseFromPb(&regenerateDashboardResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *qualityMonitorsImpl) RunRefresh(ctx context.Context, request RunRefreshRequest) (*MonitorRefreshInfo, error) {

	var monitorRefreshInfoPb catalogpb.MonitorRefreshInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		&monitorRefreshInfoPb,
	)
	resp, err := MonitorRefreshInfoFromPb(&monitorRefreshInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *qualityMonitorsImpl) Update(ctx context.Context, request UpdateMonitor) (*MonitorInfo, error) {
	requestPb, pbErr := UpdateMonitorToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var monitorInfoPb catalogpb.MonitorInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&monitorInfoPb,
	)
	resp, err := MonitorInfoFromPb(&monitorInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just RegisteredModels API methods
type registeredModelsImpl struct {
	client *client.DatabricksClient
}

func (a *registeredModelsImpl) Create(ctx context.Context, request CreateRegisteredModelRequest) (*RegisteredModelInfo, error) {
	requestPb, pbErr := CreateRegisteredModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var registeredModelInfoPb catalogpb.RegisteredModelInfoPb
	path := "/api/2.1/unity-catalog/models"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&registeredModelInfoPb,
	)
	resp, err := RegisteredModelInfoFromPb(&registeredModelInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *registeredModelsImpl) Delete(ctx context.Context, request DeleteRegisteredModelRequest) error {
	requestPb, pbErr := DeleteRegisteredModelRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *registeredModelsImpl) DeleteAlias(ctx context.Context, request DeleteAliasRequest) error {
	requestPb, pbErr := DeleteAliasRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/aliases/%v", request.FullName, request.Alias)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *registeredModelsImpl) Get(ctx context.Context, request GetRegisteredModelRequest) (*RegisteredModelInfo, error) {
	requestPb, pbErr := GetRegisteredModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var registeredModelInfoPb catalogpb.RegisteredModelInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&registeredModelInfoPb,
	)
	resp, err := RegisteredModelInfoFromPb(&registeredModelInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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
	requestPb, pbErr := ListRegisteredModelsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listRegisteredModelsResponsePb catalogpb.ListRegisteredModelsResponsePb
	path := "/api/2.1/unity-catalog/models"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listRegisteredModelsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListRegisteredModelsResponseFromPb(&listRegisteredModelsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *registeredModelsImpl) SetAlias(ctx context.Context, request SetRegisteredModelAliasRequest) (*RegisteredModelAlias, error) {
	requestPb, pbErr := SetRegisteredModelAliasRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var registeredModelAliasPb catalogpb.RegisteredModelAliasPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/aliases/%v", request.FullName, request.Alias)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		&registeredModelAliasPb,
	)
	resp, err := RegisteredModelAliasFromPb(&registeredModelAliasPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *registeredModelsImpl) Update(ctx context.Context, request UpdateRegisteredModelRequest) (*RegisteredModelInfo, error) {
	requestPb, pbErr := UpdateRegisteredModelRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var registeredModelInfoPb catalogpb.RegisteredModelInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&registeredModelInfoPb,
	)
	resp, err := RegisteredModelInfoFromPb(&registeredModelInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just ResourceQuotas API methods
type resourceQuotasImpl struct {
	client *client.DatabricksClient
}

func (a *resourceQuotasImpl) GetQuota(ctx context.Context, request GetQuotaRequest) (*GetQuotaResponse, error) {
	requestPb, pbErr := GetQuotaRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getQuotaResponsePb catalogpb.GetQuotaResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/resource-quotas/%v/%v/%v", request.ParentSecurableType, request.ParentFullName, request.QuotaName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getQuotaResponsePb,
	)
	resp, err := GetQuotaResponseFromPb(&getQuotaResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// ListQuotas returns all quota values under the metastore. There are no SLAs on
// the freshness of the counts returned. This API does not trigger a refresh of
// quota counts.
func (a *resourceQuotasImpl) ListQuotasAll(ctx context.Context, request ListQuotasRequest) ([]QuotaInfo, error) {
	iterator := a.ListQuotas(ctx, request)
	return listing.ToSlice[QuotaInfo](ctx, iterator)
}

func (a *resourceQuotasImpl) internalListQuotas(ctx context.Context, request ListQuotasRequest) (*ListQuotasResponse, error) {
	requestPb, pbErr := ListQuotasRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listQuotasResponsePb catalogpb.ListQuotasResponsePb
	path := "/api/2.1/unity-catalog/resource-quotas/all-resource-quotas"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listQuotasResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListQuotasResponseFromPb(&listQuotasResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Schemas API methods
type schemasImpl struct {
	client *client.DatabricksClient
}

func (a *schemasImpl) Create(ctx context.Context, request CreateSchema) (*SchemaInfo, error) {
	requestPb, pbErr := CreateSchemaToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var schemaInfoPb catalogpb.SchemaInfoPb
	path := "/api/2.1/unity-catalog/schemas"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&schemaInfoPb,
	)
	resp, err := SchemaInfoFromPb(&schemaInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *schemasImpl) Delete(ctx context.Context, request DeleteSchemaRequest) error {
	requestPb, pbErr := DeleteSchemaRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *schemasImpl) Get(ctx context.Context, request GetSchemaRequest) (*SchemaInfo, error) {
	requestPb, pbErr := GetSchemaRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var schemaInfoPb catalogpb.SchemaInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&schemaInfoPb,
	)
	resp, err := SchemaInfoFromPb(&schemaInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets an array of schemas for a catalog in the metastore. If the caller is the
// metastore admin or the owner of the parent catalog, all schemas for the
// catalog will be retrieved. Otherwise, only schemas owned by the caller (or
// for which the caller has the **USE_SCHEMA** privilege) will be retrieved.
// There is no guarantee of a specific ordering of the elements in the array.
func (a *schemasImpl) List(ctx context.Context, request ListSchemasRequest) listing.Iterator[SchemaInfo] {

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
	requestPb, pbErr := ListSchemasRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listSchemasResponsePb catalogpb.ListSchemasResponsePb
	path := "/api/2.1/unity-catalog/schemas"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listSchemasResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListSchemasResponseFromPb(&listSchemasResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *schemasImpl) Update(ctx context.Context, request UpdateSchema) (*SchemaInfo, error) {
	requestPb, pbErr := UpdateSchemaToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var schemaInfoPb catalogpb.SchemaInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&schemaInfoPb,
	)
	resp, err := SchemaInfoFromPb(&schemaInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just StorageCredentials API methods
type storageCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *storageCredentialsImpl) Create(ctx context.Context, request CreateStorageCredential) (*StorageCredentialInfo, error) {
	requestPb, pbErr := CreateStorageCredentialToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var storageCredentialInfoPb catalogpb.StorageCredentialInfoPb
	path := "/api/2.1/unity-catalog/storage-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&storageCredentialInfoPb,
	)
	resp, err := StorageCredentialInfoFromPb(&storageCredentialInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *storageCredentialsImpl) Delete(ctx context.Context, request DeleteStorageCredentialRequest) error {
	requestPb, pbErr := DeleteStorageCredentialRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *storageCredentialsImpl) Get(ctx context.Context, request GetStorageCredentialRequest) (*StorageCredentialInfo, error) {
	requestPb, pbErr := GetStorageCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var storageCredentialInfoPb catalogpb.StorageCredentialInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&storageCredentialInfoPb,
	)
	resp, err := StorageCredentialInfoFromPb(&storageCredentialInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets an array of storage credentials (as __StorageCredentialInfo__ objects).
// The array is limited to only those storage credentials the caller has
// permission to access. If the caller is a metastore admin, retrieval of
// credentials is unrestricted. There is no guarantee of a specific ordering of
// the elements in the array.
func (a *storageCredentialsImpl) List(ctx context.Context, request ListStorageCredentialsRequest) listing.Iterator[StorageCredentialInfo] {

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
	requestPb, pbErr := ListStorageCredentialsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listStorageCredentialsResponsePb catalogpb.ListStorageCredentialsResponsePb
	path := "/api/2.1/unity-catalog/storage-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listStorageCredentialsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListStorageCredentialsResponseFromPb(&listStorageCredentialsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *storageCredentialsImpl) Update(ctx context.Context, request UpdateStorageCredential) (*StorageCredentialInfo, error) {
	requestPb, pbErr := UpdateStorageCredentialToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var storageCredentialInfoPb catalogpb.StorageCredentialInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&storageCredentialInfoPb,
	)
	resp, err := StorageCredentialInfoFromPb(&storageCredentialInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *storageCredentialsImpl) Validate(ctx context.Context, request ValidateStorageCredential) (*ValidateStorageCredentialResponse, error) {
	requestPb, pbErr := ValidateStorageCredentialToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var validateStorageCredentialResponsePb catalogpb.ValidateStorageCredentialResponsePb
	path := "/api/2.1/unity-catalog/validate-storage-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&validateStorageCredentialResponsePb,
	)
	resp, err := ValidateStorageCredentialResponseFromPb(&validateStorageCredentialResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just SystemSchemas API methods
type systemSchemasImpl struct {
	client *client.DatabricksClient
}

func (a *systemSchemasImpl) Disable(ctx context.Context, request DisableRequest) error {
	requestPb, pbErr := DisableRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v/systemschemas/%v", request.MetastoreId, request.SchemaName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *systemSchemasImpl) Enable(ctx context.Context, request EnableRequest) error {
	requestPb, pbErr := EnableRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v/systemschemas/%v", request.MetastoreId, request.SchemaName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// Gets an array of system schemas for a metastore. The caller must be an
// account admin or a metastore admin.
func (a *systemSchemasImpl) List(ctx context.Context, request ListSystemSchemasRequest) listing.Iterator[SystemSchemaInfo] {

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

// Gets an array of system schemas for a metastore. The caller must be an
// account admin or a metastore admin.
func (a *systemSchemasImpl) ListAll(ctx context.Context, request ListSystemSchemasRequest) ([]SystemSchemaInfo, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[SystemSchemaInfo](ctx, iterator)
}

func (a *systemSchemasImpl) internalList(ctx context.Context, request ListSystemSchemasRequest) (*ListSystemSchemasResponse, error) {
	requestPb, pbErr := ListSystemSchemasRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listSystemSchemasResponsePb catalogpb.ListSystemSchemasResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v/systemschemas", request.MetastoreId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listSystemSchemasResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListSystemSchemasResponseFromPb(&listSystemSchemasResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just TableConstraints API methods
type tableConstraintsImpl struct {
	client *client.DatabricksClient
}

func (a *tableConstraintsImpl) Create(ctx context.Context, request CreateTableConstraint) (*TableConstraint, error) {
	requestPb, pbErr := CreateTableConstraintToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var tableConstraintPb catalogpb.TableConstraintPb
	path := "/api/2.1/unity-catalog/constraints"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&tableConstraintPb,
	)
	resp, err := TableConstraintFromPb(&tableConstraintPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tableConstraintsImpl) Delete(ctx context.Context, request DeleteTableConstraintRequest) error {
	requestPb, pbErr := DeleteTableConstraintRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/constraints/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just Tables API methods
type tablesImpl struct {
	client *client.DatabricksClient
}

func (a *tablesImpl) Delete(ctx context.Context, request DeleteTableRequest) error {
	requestPb, pbErr := DeleteTableRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *tablesImpl) Exists(ctx context.Context, request ExistsRequest) (*TableExistsResponse, error) {
	requestPb, pbErr := ExistsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var tableExistsResponsePb catalogpb.TableExistsResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/exists", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&tableExistsResponsePb,
	)
	resp, err := TableExistsResponseFromPb(&tableExistsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tablesImpl) Get(ctx context.Context, request GetTableRequest) (*TableInfo, error) {
	requestPb, pbErr := GetTableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var tableInfoPb catalogpb.TableInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&tableInfoPb,
	)
	resp, err := TableInfoFromPb(&tableInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets an array of all tables for the current metastore under the parent
// catalog and schema. The caller must be a metastore admin or an owner of (or
// have the **SELECT** privilege on) the table. For the latter case, the caller
// must also be the owner or have the **USE_CATALOG** privilege on the parent
// catalog and the **USE_SCHEMA** privilege on the parent schema. There is no
// guarantee of a specific ordering of the elements in the array.
func (a *tablesImpl) List(ctx context.Context, request ListTablesRequest) listing.Iterator[TableInfo] {

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
	requestPb, pbErr := ListTablesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listTablesResponsePb catalogpb.ListTablesResponsePb
	path := "/api/2.1/unity-catalog/tables"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listTablesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListTablesResponseFromPb(&listTablesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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
	requestPb, pbErr := ListSummariesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listTableSummariesResponsePb catalogpb.ListTableSummariesResponsePb
	path := "/api/2.1/unity-catalog/table-summaries"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listTableSummariesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListTableSummariesResponseFromPb(&listTableSummariesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *tablesImpl) Update(ctx context.Context, request UpdateTableRequest) error {
	requestPb, pbErr := UpdateTableRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just TemporaryTableCredentials API methods
type temporaryTableCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *temporaryTableCredentialsImpl) GenerateTemporaryTableCredentials(ctx context.Context, request GenerateTemporaryTableCredentialRequest) (*GenerateTemporaryTableCredentialResponse, error) {
	requestPb, pbErr := GenerateTemporaryTableCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var generateTemporaryTableCredentialResponsePb catalogpb.GenerateTemporaryTableCredentialResponsePb
	path := "/api/2.0/unity-catalog/temporary-table-credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&generateTemporaryTableCredentialResponsePb,
	)
	resp, err := GenerateTemporaryTableCredentialResponseFromPb(&generateTemporaryTableCredentialResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just Volumes API methods
type volumesImpl struct {
	client *client.DatabricksClient
}

func (a *volumesImpl) Create(ctx context.Context, request CreateVolumeRequestContent) (*VolumeInfo, error) {
	requestPb, pbErr := CreateVolumeRequestContentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var volumeInfoPb catalogpb.VolumeInfoPb
	path := "/api/2.1/unity-catalog/volumes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&volumeInfoPb,
	)
	resp, err := VolumeInfoFromPb(&volumeInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *volumesImpl) Delete(ctx context.Context, request DeleteVolumeRequest) error {
	requestPb, pbErr := DeleteVolumeRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

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
	requestPb, pbErr := ListVolumesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listVolumesResponseContentPb catalogpb.ListVolumesResponseContentPb
	path := "/api/2.1/unity-catalog/volumes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listVolumesResponseContentPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListVolumesResponseContentFromPb(&listVolumesResponseContentPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *volumesImpl) Read(ctx context.Context, request ReadVolumeRequest) (*VolumeInfo, error) {
	requestPb, pbErr := ReadVolumeRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var volumeInfoPb catalogpb.VolumeInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&volumeInfoPb,
	)
	resp, err := VolumeInfoFromPb(&volumeInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *volumesImpl) Update(ctx context.Context, request UpdateVolumeRequestContent) (*VolumeInfo, error) {
	requestPb, pbErr := UpdateVolumeRequestContentToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var volumeInfoPb catalogpb.VolumeInfoPb
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&volumeInfoPb,
	)
	resp, err := VolumeInfoFromPb(&volumeInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just WorkspaceBindings API methods
type workspaceBindingsImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceBindingsImpl) Get(ctx context.Context, request GetWorkspaceBindingRequest) (*GetCatalogWorkspaceBindingsResponse, error) {
	requestPb, pbErr := GetWorkspaceBindingRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getCatalogWorkspaceBindingsResponsePb catalogpb.GetCatalogWorkspaceBindingsResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspace-bindings/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getCatalogWorkspaceBindingsResponsePb,
	)
	resp, err := GetCatalogWorkspaceBindingsResponseFromPb(&getCatalogWorkspaceBindingsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets workspace bindings of the securable. The caller must be a metastore
// admin or an owner of the securable.
func (a *workspaceBindingsImpl) GetBindings(ctx context.Context, request GetBindingsRequest) listing.Iterator[WorkspaceBinding] {

	getNextPage := func(ctx context.Context, req GetBindingsRequest) (*GetWorkspaceBindingsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalGetBindings(ctx, req)
	}
	getItems := func(resp *GetWorkspaceBindingsResponse) []WorkspaceBinding {
		return resp.Bindings
	}
	getNextReq := func(resp *GetWorkspaceBindingsResponse) *GetBindingsRequest {
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

// Gets workspace bindings of the securable. The caller must be a metastore
// admin or an owner of the securable.
func (a *workspaceBindingsImpl) GetBindingsAll(ctx context.Context, request GetBindingsRequest) ([]WorkspaceBinding, error) {
	iterator := a.GetBindings(ctx, request)
	return listing.ToSlice[WorkspaceBinding](ctx, iterator)
}

func (a *workspaceBindingsImpl) internalGetBindings(ctx context.Context, request GetBindingsRequest) (*GetWorkspaceBindingsResponse, error) {
	requestPb, pbErr := GetBindingsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getWorkspaceBindingsResponsePb catalogpb.GetWorkspaceBindingsResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/bindings/%v/%v", request.SecurableType, request.SecurableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&getWorkspaceBindingsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := GetWorkspaceBindingsResponseFromPb(&getWorkspaceBindingsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceBindingsImpl) Update(ctx context.Context, request UpdateWorkspaceBindings) (*UpdateCatalogWorkspaceBindingsResponse, error) {
	requestPb, pbErr := UpdateWorkspaceBindingsToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateCatalogWorkspaceBindingsResponsePb catalogpb.UpdateCatalogWorkspaceBindingsResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspace-bindings/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&updateCatalogWorkspaceBindingsResponsePb,
	)
	resp, err := UpdateCatalogWorkspaceBindingsResponseFromPb(&updateCatalogWorkspaceBindingsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspaceBindingsImpl) UpdateBindings(ctx context.Context, request UpdateWorkspaceBindingsParameters) (*UpdateWorkspaceBindingsResponse, error) {
	requestPb, pbErr := UpdateWorkspaceBindingsParametersToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateWorkspaceBindingsResponsePb catalogpb.UpdateWorkspaceBindingsResponsePb
	path := fmt.Sprintf("/api/2.1/unity-catalog/bindings/%v/%v", request.SecurableType, request.SecurableName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&updateWorkspaceBindingsResponsePb,
	)
	resp, err := UpdateWorkspaceBindingsResponseFromPb(&updateWorkspaceBindingsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

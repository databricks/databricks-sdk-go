// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package modelregistry

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

type ModelregistryService interface {
	CreateModelVersion(ctx context.Context, createModelVersionRequest CreateModelVersionRequest) (*CreateModelVersionResponse, error)
	// Throws ``RESOURCE_ALREADY_EXISTS`` if a registered model with the given
	// name exists.
	CreateRegisteredModel(ctx context.Context, createRegisteredModelRequest CreateRegisteredModelRequest) (*CreateRegisteredModelResponse, error)

	DeleteModelVersion(ctx context.Context, deleteModelVersionRequest DeleteModelVersionRequest) error

	DeleteModelVersionTag(ctx context.Context, deleteModelVersionTagRequest DeleteModelVersionTagRequest) error

	DeleteRegisteredModel(ctx context.Context, deleteRegisteredModelRequest DeleteRegisteredModelRequest) error

	DeleteRegisteredModelTag(ctx context.Context, deleteRegisteredModelTagRequest DeleteRegisteredModelTagRequest) error

	GetLatestVersions(ctx context.Context, getLatestVersionsRequest GetLatestVersionsRequest) (*GetLatestVersionsResponse, error)

	GetModelVersion(ctx context.Context, getModelVersionRequest GetModelVersionRequest) (*GetModelVersionResponse, error)

	GetModelVersionDownloadUri(ctx context.Context, getModelVersionDownloadUriRequest GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error)

	GetRegisteredModel(ctx context.Context, getRegisteredModelRequest GetRegisteredModelRequest) (*GetRegisteredModelResponse, error)

	ListRegisteredModels(ctx context.Context, listRegisteredModelsRequest ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error)

	RenameRegisteredModel(ctx context.Context, renameRegisteredModelRequest RenameRegisteredModelRequest) (*RenameRegisteredModelResponse, error)

	SearchModelVersions(ctx context.Context, searchModelVersionsRequest SearchModelVersionsRequest) (*SearchModelVersionsResponse, error)

	SearchRegisteredModels(ctx context.Context, searchRegisteredModelsRequest SearchRegisteredModelsRequest) (*SearchRegisteredModelsResponse, error)

	SetModelVersionTag(ctx context.Context, setModelVersionTagRequest SetModelVersionTagRequest) error

	SetRegisteredModelTag(ctx context.Context, setRegisteredModelTagRequest SetRegisteredModelTagRequest) error

	TransitionModelVersionStage(ctx context.Context, transitionModelVersionStageRequest TransitionModelVersionStageRequest) (*TransitionModelVersionStageResponse, error)

	UpdateModelVersion(ctx context.Context, updateModelVersionRequest UpdateModelVersionRequest) error

	UpdateRegisteredModel(ctx context.Context, updateRegisteredModelRequest UpdateRegisteredModelRequest) error
}

func New(client *client.DatabricksClient) ModelregistryService {
	return &ModelregistryAPI{
		client: client,
	}
}

type ModelregistryAPI struct {
	client *client.DatabricksClient
}

func (a *ModelregistryAPI) CreateModelVersion(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error) {
	var createModelVersionResponse CreateModelVersionResponse
	path := "/api/2.0/preview/mlflow/model-versions/create"
	err := a.client.Post(ctx, path, request, &createModelVersionResponse)
	return &createModelVersionResponse, err
}

// Throws ``RESOURCE_ALREADY_EXISTS`` if a registered model with the given name
// exists.
func (a *ModelregistryAPI) CreateRegisteredModel(ctx context.Context, request CreateRegisteredModelRequest) (*CreateRegisteredModelResponse, error) {
	var createRegisteredModelResponse CreateRegisteredModelResponse
	path := "/api/2.0/preview/mlflow/registered-models/create"
	err := a.client.Post(ctx, path, request, &createRegisteredModelResponse)
	return &createRegisteredModelResponse, err
}

func (a *ModelregistryAPI) DeleteModelVersion(ctx context.Context, request DeleteModelVersionRequest) error {
	path := "/api/2.0/preview/mlflow/model-versions/delete"
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *ModelregistryAPI) DeleteModelVersionTag(ctx context.Context, request DeleteModelVersionTagRequest) error {
	path := "/api/2.0/preview/mlflow/model-versions/delete-tag"
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *ModelregistryAPI) DeleteRegisteredModel(ctx context.Context, request DeleteRegisteredModelRequest) error {
	path := "/api/2.0/preview/mlflow/registered-models/delete"
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *ModelregistryAPI) DeleteRegisteredModelTag(ctx context.Context, request DeleteRegisteredModelTagRequest) error {
	path := "/api/2.0/preview/mlflow/registered-models/delete-tag"
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *ModelregistryAPI) GetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) (*GetLatestVersionsResponse, error) {
	var getLatestVersionsResponse GetLatestVersionsResponse
	path := "/api/2.0/mlflow/registered-models/get-latest-versions"
	err := a.client.Post(ctx, path, request, &getLatestVersionsResponse)
	return &getLatestVersionsResponse, err
}

func (a *ModelregistryAPI) GetModelVersion(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error) {
	var getModelVersionResponse GetModelVersionResponse
	path := "/api/2.0/preview/mlflow/model-versions/get"
	err := a.client.Get(ctx, path, request, &getModelVersionResponse)
	return &getModelVersionResponse, err
}

func (a *ModelregistryAPI) GetModelVersionDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error) {
	var getModelVersionDownloadUriResponse GetModelVersionDownloadUriResponse
	path := "/api/2.0/preview/mlflow/model-versions/get-download-uri"
	err := a.client.Get(ctx, path, request, &getModelVersionDownloadUriResponse)
	return &getModelVersionDownloadUriResponse, err
}

func (a *ModelregistryAPI) GetRegisteredModel(ctx context.Context, request GetRegisteredModelRequest) (*GetRegisteredModelResponse, error) {
	var getRegisteredModelResponse GetRegisteredModelResponse
	path := "/api/2.0/preview/mlflow/registered-models/get"
	err := a.client.Get(ctx, path, request, &getRegisteredModelResponse)
	return &getRegisteredModelResponse, err
}

func (a *ModelregistryAPI) ListRegisteredModels(ctx context.Context, request ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error) {
	var listRegisteredModelsResponse ListRegisteredModelsResponse
	path := "/api/2.0/preview/mlflow/registered-models/list"
	err := a.client.Get(ctx, path, request, &listRegisteredModelsResponse)
	return &listRegisteredModelsResponse, err
}

func (a *ModelregistryAPI) RenameRegisteredModel(ctx context.Context, request RenameRegisteredModelRequest) (*RenameRegisteredModelResponse, error) {
	var renameRegisteredModelResponse RenameRegisteredModelResponse
	path := "/api/2.0/preview/mlflow/registered-models/rename"
	err := a.client.Post(ctx, path, request, &renameRegisteredModelResponse)
	return &renameRegisteredModelResponse, err
}

func (a *ModelregistryAPI) SearchModelVersions(ctx context.Context, request SearchModelVersionsRequest) (*SearchModelVersionsResponse, error) {
	var searchModelVersionsResponse SearchModelVersionsResponse
	path := "/api/2.0/preview/mlflow/model-versions/search"
	err := a.client.Get(ctx, path, request, &searchModelVersionsResponse)
	return &searchModelVersionsResponse, err
}

func (a *ModelregistryAPI) SearchRegisteredModels(ctx context.Context, request SearchRegisteredModelsRequest) (*SearchRegisteredModelsResponse, error) {
	var searchRegisteredModelsResponse SearchRegisteredModelsResponse
	path := "/api/2.0/preview/mlflow/registered-models/search"
	err := a.client.Get(ctx, path, request, &searchRegisteredModelsResponse)
	return &searchRegisteredModelsResponse, err
}

func (a *ModelregistryAPI) SetModelVersionTag(ctx context.Context, request SetModelVersionTagRequest) error {
	path := "/api/2.0/preview/mlflow/model-versions/set-tag"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *ModelregistryAPI) SetRegisteredModelTag(ctx context.Context, request SetRegisteredModelTagRequest) error {
	path := "/api/2.0/preview/mlflow/registered-models/set-tag"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *ModelregistryAPI) TransitionModelVersionStage(ctx context.Context, request TransitionModelVersionStageRequest) (*TransitionModelVersionStageResponse, error) {
	var transitionModelVersionStageResponse TransitionModelVersionStageResponse
	path := "/api/2.0/preview/mlflow/model-versions/transition-stage"
	err := a.client.Post(ctx, path, request, &transitionModelVersionStageResponse)
	return &transitionModelVersionStageResponse, err
}

func (a *ModelregistryAPI) UpdateModelVersion(ctx context.Context, request UpdateModelVersionRequest) error {
	path := "/api/2.0/preview/mlflow/model-versions/update"
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *ModelregistryAPI) UpdateRegisteredModel(ctx context.Context, request UpdateRegisteredModelRequest) error {
	path := "/api/2.0/preview/mlflow/registered-models/update"
	err := a.client.Patch(ctx, path, request)
	return err
}

// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlflowext

import (
	"context"
	

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// These endpoints are modified versions of the MLflow API that accept
// additional input parameters or return additional information.
type MlflowextService interface {
    
    GetRegisteredModel(ctx context.Context, getRegisteredModelRequest GetRegisteredModelRequest) (*GetRegisteredModelResponse, error)
    // Transition a model version&#39;s stage. This is a &lt;Workspace&gt; version of the
    // [MLflow
    // endpoint](https://www.mlflow.org/docs/latest/rest-api.html#transition-modelversion-stage)
    // that also accepts a comment associated with the transition to be
    // recorded.
    TransitionModelVersionStage(ctx context.Context, transitionModelVersionStageRequest TransitionModelVersionStageRequest) (*TransitionModelVersionStageResponse, error)
}

func New(client *client.DatabricksClient) MlflowextService {
	return &MlflowextAPI{
		client: client,
	}
}

type MlflowextAPI struct {
	client *client.DatabricksClient
}


func (a *MlflowextAPI) GetRegisteredModel(ctx context.Context, request GetRegisteredModelRequest) (*GetRegisteredModelResponse, error) {
	var getRegisteredModelResponse GetRegisteredModelResponse
	path := "/api/2.0/mlflow/databricks/registered-models/get"
	err := a.client.Get(ctx, path, request, &getRegisteredModelResponse)
	return &getRegisteredModelResponse, err
}

// Transition a model version&#39;s stage. This is a &lt;Workspace&gt; version of the
// [MLflow
// endpoint](https://www.mlflow.org/docs/latest/rest-api.html#transition-modelversion-stage)
// that also accepts a comment associated with the transition to be recorded.
func (a *MlflowextAPI) TransitionModelVersionStage(ctx context.Context, request TransitionModelVersionStageRequest) (*TransitionModelVersionStageResponse, error) {
	var transitionModelVersionStageResponse TransitionModelVersionStageResponse
	path := "/api/2.0/mlflow/databricks/model-versions/transition-stage"
	err := a.client.Post(ctx, path, request, &transitionModelVersionStageResponse)
	return &transitionModelVersionStageResponse, err
}


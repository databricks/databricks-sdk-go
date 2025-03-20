// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Experiments, Forecasting, Model Registry, etc.
package ml

import (
	"context"
)

// Experiments are the primary unit of organization in MLflow; all MLflow runs
// belong to an experiment. Each experiment lets you visualize, search, and
// compare runs, as well as download run artifacts or metadata for analysis in
// other tools. Experiments are maintained in a Databricks hosted MLflow
// tracking server.
//
// Experiments are located in the workspace file tree. You manage experiments
// using the same tools you use to manage other workspace objects such as
// folders, notebooks, and libraries.
type ExperimentsAPI struct {
	experimentsImpl
}

// Get experiment permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *ExperimentsAPI) GetPermissionLevelsByExperimentId(ctx context.Context, experimentId string) (*GetExperimentPermissionLevelsResponse, error) {
	return a.experimentsImpl.GetPermissionLevels(ctx, GetExperimentPermissionLevelsRequest{
		ExperimentId: experimentId,
	})
}

// Get experiment permissions.
//
// Gets the permissions of an experiment. Experiments can inherit permissions
// from their root object.
func (a *ExperimentsAPI) GetPermissionsByExperimentId(ctx context.Context, experimentId string) (*ExperimentPermissions, error) {
	return a.experimentsImpl.GetPermissions(ctx, GetExperimentPermissionsRequest{
		ExperimentId: experimentId,
	})
}

// The Forecasting API allows you to create and get serverless forecasting
// experiments
type ForecastingAPI struct {
	forecastingImpl
}

// Get a forecasting experiment.
//
// Public RPC to get forecasting experiment
func (a *ForecastingAPI) GetExperimentByExperimentId(ctx context.Context, experimentId string) (*ForecastingExperiment, error) {
	return a.forecastingImpl.GetExperiment(ctx, GetForecastingExperimentRequest{
		ExperimentId: experimentId,
	})
}

// Note: This API reference documents APIs for the Workspace Model Registry.
// Databricks recommends using [Models in Unity
// Catalog](/api/workspace/registeredmodels) instead. Models in Unity Catalog
// provides centralized model governance, cross-workspace access, lineage, and
// deployment. Workspace Model Registry will be deprecated in the future.
//
// The Workspace Model Registry is a centralized model repository and a UI and
// set of APIs that enable you to manage the full lifecycle of MLflow Models.
type ModelRegistryAPI struct {
	modelRegistryImpl
}

// Get registered model permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *ModelRegistryAPI) GetPermissionLevelsByRegisteredModelId(ctx context.Context, registeredModelId string) (*GetRegisteredModelPermissionLevelsResponse, error) {
	return a.modelRegistryImpl.GetPermissionLevels(ctx, GetRegisteredModelPermissionLevelsRequest{
		RegisteredModelId: registeredModelId,
	})
}

// Get registered model permissions.
//
// Gets the permissions of a registered model. Registered models can inherit
// permissions from their root object.
func (a *ModelRegistryAPI) GetPermissionsByRegisteredModelId(ctx context.Context, registeredModelId string) (*RegisteredModelPermissions, error) {
	return a.modelRegistryImpl.GetPermissions(ctx, GetRegisteredModelPermissionsRequest{
		RegisteredModelId: registeredModelId,
	})
}

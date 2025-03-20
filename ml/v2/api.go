// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Experiments, Forecasting, Model Registry, etc.
package ml

import (
	"context"
)

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

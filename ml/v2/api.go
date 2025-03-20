// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Experiments, Forecasting, Model Registry, etc.
package ml

import (
	"context"
)

type experimentsBaseClient struct {
	experimentsImpl
}

// Get experiment permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *experimentsBaseClient) GetPermissionLevelsByExperimentId(ctx context.Context, experimentId string) (*GetExperimentPermissionLevelsResponse, error) {
	return a.experimentsImpl.GetPermissionLevels(ctx, GetExperimentPermissionLevelsRequest{
		ExperimentId: experimentId,
	})
}

// Get experiment permissions.
//
// Gets the permissions of an experiment. Experiments can inherit permissions
// from their root object.
func (a *experimentsBaseClient) GetPermissionsByExperimentId(ctx context.Context, experimentId string) (*ExperimentPermissions, error) {
	return a.experimentsImpl.GetPermissions(ctx, GetExperimentPermissionsRequest{
		ExperimentId: experimentId,
	})
}

type forecastingBaseClient struct {
	forecastingImpl
}

// Get a forecasting experiment.
//
// Public RPC to get forecasting experiment
func (a *forecastingBaseClient) GetExperimentByExperimentId(ctx context.Context, experimentId string) (*ForecastingExperiment, error) {
	return a.forecastingImpl.GetExperiment(ctx, GetForecastingExperimentRequest{
		ExperimentId: experimentId,
	})
}

type modelRegistryBaseClient struct {
	modelRegistryImpl
}

// Get registered model permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *modelRegistryBaseClient) GetPermissionLevelsByRegisteredModelId(ctx context.Context, registeredModelId string) (*GetRegisteredModelPermissionLevelsResponse, error) {
	return a.modelRegistryImpl.GetPermissionLevels(ctx, GetRegisteredModelPermissionLevelsRequest{
		RegisteredModelId: registeredModelId,
	})
}

// Get registered model permissions.
//
// Gets the permissions of a registered model. Registered models can inherit
// permissions from their root object.
func (a *modelRegistryBaseClient) GetPermissionsByRegisteredModelId(ctx context.Context, registeredModelId string) (*RegisteredModelPermissions, error) {
	return a.modelRegistryImpl.GetPermissions(ctx, GetRegisteredModelPermissionsRequest{
		RegisteredModelId: registeredModelId,
	})
}

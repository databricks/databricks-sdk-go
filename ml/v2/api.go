// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Experiments, Forecasting, Model Registry, etc.
package ml

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
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

// Create a forecasting experiment.
//
// Creates a serverless forecasting experiment. Returns the experiment ID.
func (a *forecastingBaseClient) CreateExperiment(ctx context.Context, createForecastingExperimentRequest CreateForecastingExperimentRequest) (*ForecastingCreateExperimentWaiter, error) {
	createForecastingExperimentResponse, err := a.forecastingImpl.CreateExperiment(ctx, createForecastingExperimentRequest)
	if err != nil {
		return nil, err
	}
	return &ForecastingCreateExperimentWaiter{
		RawResponse:  createForecastingExperimentResponse,
		experimentId: createForecastingExperimentResponse.ExperimentId,
		service:      a,
	}, nil
}

type ForecastingCreateExperimentWaiter struct {
	// RawResponse is the raw response of the CreateExperiment call.
	RawResponse  *CreateForecastingExperimentResponse
	service      *forecastingBaseClient
	experimentId string
}

// Polls the server until the operation reaches a terminal state, encounters an error, or reaches a timeout defaults to 20 min.
// This method will return an error if a failure state is reached.
func (w *ForecastingCreateExperimentWaiter) WaitUntilDone(ctx context.Context, opts *retries.WaitUntilDoneOptions) (*ForecastingExperiment, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	if opts == nil {
		opts = &retries.WaitUntilDoneOptions{}
	}
	if opts.Timeout == 0 {
		opts.Timeout = 20 * time.Minute
	}

	return retries.Poll[ForecastingExperiment](ctx, opts.Timeout, func() (*ForecastingExperiment, *retries.Err) {
		forecastingExperiment, err := w.service.GetExperiment(ctx, GetForecastingExperimentRequest{
			ExperimentId: w.experimentId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := forecastingExperiment.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case ForecastingExperimentStateSucceeded: // target state
			return forecastingExperiment, nil
		case ForecastingExperimentStateFailed, ForecastingExperimentStateCancelled:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ForecastingExperimentStateSucceeded, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})

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

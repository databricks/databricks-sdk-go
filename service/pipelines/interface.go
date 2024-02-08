// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"context"
)

// The Delta Live Tables API allows you to create, edit, delete, start, and view
// details about pipelines.
//
// Delta Live Tables is a framework for building reliable, maintainable, and
// testable data processing pipelines. You define the transformations to perform
// on your data, and Delta Live Tables manages task orchestration, cluster
// management, monitoring, data quality, and error handling.
//
// Instead of defining your data pipelines using a series of separate Apache
// Spark tasks, Delta Live Tables manages how your data is transformed based on
// a target schema you define for each processing step. You can also enforce
// data quality with Delta Live Tables expectations. Expectations allow you to
// define expected data quality and specify how to handle records that fail
// those expectations.
type PipelinesService interface {

	// Create a pipeline.
	//
	// Creates a new data processing pipeline based on the requested
	// configuration. If successful, this method returns the ID of the new
	// pipeline.
	Create(ctx context.Context, request CreatePipeline) (*CreatePipelineResponse, error)

	// Delete a pipeline.
	//
	// Deletes a pipeline.
	Delete(ctx context.Context, request DeletePipelineRequest) error

	// Get a pipeline.
	Get(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error)

	// Get pipeline permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetPipelinePermissionLevelsRequest) (*GetPipelinePermissionLevelsResponse, error)

	// Get pipeline permissions.
	//
	// Gets the permissions of a pipeline. Pipelines can inherit permissions
	// from their root object.
	GetPermissions(ctx context.Context, request GetPipelinePermissionsRequest) (*PipelinePermissions, error)

	// Get a pipeline update.
	//
	// Gets an update from an active pipeline.
	GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error)

	// List pipeline events.
	//
	// Retrieves events for a pipeline.
	//
	// Use ListPipelineEventsAll() to get all PipelineEvent instances, which will iterate over every result page.
	ListPipelineEvents(ctx context.Context, request ListPipelineEventsRequest) (*ListPipelineEventsResponse, error)

	// List pipelines.
	//
	// Lists pipelines defined in the Delta Live Tables system.
	//
	// Use ListPipelinesAll() to get all PipelineStateInfo instances, which will iterate over every result page.
	ListPipelines(ctx context.Context, request ListPipelinesRequest) (*ListPipelinesResponse, error)

	// List pipeline updates.
	//
	// List updates for an active pipeline.
	ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error)

	// Set pipeline permissions.
	//
	// Sets permissions on a pipeline. Pipelines can inherit permissions from
	// their root object.
	SetPermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error)

	// Start a pipeline.
	//
	// Starts a new update for the pipeline. If there is already an active
	// update for the pipeline, the request will fail and the active update will
	// remain running.
	StartUpdate(ctx context.Context, request StartUpdate) (*StartUpdateResponse, error)

	// Stop a pipeline.
	//
	// Stops the pipeline by canceling the active update. If there is no active
	// update for the pipeline, this request is a no-op.
	Stop(ctx context.Context, request StopRequest) error

	// Edit a pipeline.
	//
	// Updates a pipeline with the supplied configuration.
	Update(ctx context.Context, request EditPipeline) error

	// Update pipeline permissions.
	//
	// Updates the permissions on a pipeline. Pipelines can inherit permissions
	// from their root object.
	UpdatePermissions(ctx context.Context, request PipelinePermissionsRequest) (*PipelinePermissions, error)
}

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

	// Create a pipeline
	//
	// Creates a new data processing pipeline based on the requested
	// configuration. If successful, this method returns the ID of the new
	// pipeline.
	CreatePipeline(ctx context.Context, request CreatePipeline) (*CreatePipelineResponse, error)

	// Delete a pipeline
	//
	// Deletes a pipeline.
	DeletePipeline(ctx context.Context, request DeletePipelineRequest) error

	// Get a pipeline
	GetPipeline(ctx context.Context, request GetPipelineRequest) (*GetPipelineResponse, error)

	// Get a pipeline update
	//
	// Gets an update from an active pipeline.
	GetUpdate(ctx context.Context, request GetUpdateRequest) (*GetUpdateResponse, error)

	// List pipeline updates
	//
	// List updates for an active pipeline.
	ListUpdates(ctx context.Context, request ListUpdatesRequest) (*ListUpdatesResponse, error)

	// Reset a pipeline
	//
	// Resets a pipeline.
	ResetPipeline(ctx context.Context, request ResetPipelineRequest) error

	// Queue a pipeline update
	//
	// Starts or queues a pipeline update.
	StartUpdate(ctx context.Context, request StartUpdate) (*StartUpdateResponse, error)

	// Stop a pipeline
	//
	// Stops a pipeline.
	StopPipeline(ctx context.Context, request StopPipelineRequest) error

	// Edit a pipeline
	//
	// Updates a pipeline with the supplied configuration.
	UpdatePipeline(ctx context.Context, request EditPipeline) error
}

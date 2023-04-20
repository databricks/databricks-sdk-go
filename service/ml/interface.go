// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"context"
)

type ExperimentsService interface {

	// Create experiment.
	//
	// Creates an experiment with a name. Returns the ID of the newly created
	// experiment. Validates that another experiment with the same name does not
	// already exist and fails if another experiment with the same name already
	// exists.
	//
	// Throws `RESOURCE_ALREADY_EXISTS` if a experiment with the given name
	// exists.
	CreateExperiment(ctx context.Context, request CreateExperiment) (*CreateExperimentResponse, error)

	// Create a run.
	//
	// Creates a new run within an experiment. A run is usually a single
	// execution of a machine learning or data ETL pipeline. MLflow uses runs to
	// track the `mlflowParam`, `mlflowMetric` and `mlflowRunTag` associated
	// with a single execution.
	CreateRun(ctx context.Context, request CreateRun) (*CreateRunResponse, error)

	// Delete an experiment.
	//
	// Marks an experiment and associated metadata, runs, metrics, params, and
	// tags for deletion. If the experiment uses FileStore, artifacts associated
	// with experiment are also deleted.
	DeleteExperiment(ctx context.Context, request DeleteExperiment) error

	// Delete a run.
	//
	// Marks a run for deletion.
	DeleteRun(ctx context.Context, request DeleteRun) error

	// Delete a tag.
	//
	// Deletes a tag on a run. Tags are run metadata that can be updated during
	// a run and after a run completes.
	DeleteTag(ctx context.Context, request DeleteTag) error

	// Get metadata.
	//
	// "Gets metadata for an experiment.
	//
	// This endpoint will return deleted experiments, but prefers the active
	// experiment if an active and deleted experiment share the same name. If
	// multiple deleted experiments share the same name, the API will return one
	// of them.
	//
	// Throws `RESOURCE_DOES_NOT_EXIST` if no experiment with the specified name
	// exists.S
	GetByName(ctx context.Context, request GetByNameRequest) (*GetExperimentByNameResponse, error)

	// Get an experiment.
	//
	// Gets metadata for an experiment. This method works on deleted
	// experiments.
	GetExperiment(ctx context.Context, request GetExperimentRequest) (*Experiment, error)

	// Get history of a given metric within a run.
	//
	// Gets a list of all values for the specified metric for a given run.
	GetHistory(ctx context.Context, request GetHistoryRequest) (*GetMetricHistoryResponse, error)

	// Get a run.
	//
	// "Gets the metadata, metrics, params, and tags for a run. In the case
	// where multiple metrics with the same key are logged for a run, return
	// only the value with the latest timestamp.
	//
	// If there are multiple values with the latest timestamp, return the
	// maximum of these values.
	GetRun(ctx context.Context, request GetRunRequest) (*GetRunResponse, error)

	// Get all artifacts.
	//
	// List artifacts for a run. Takes an optional `artifact_path` prefix. If it
	// is specified, the response contains only artifacts with the specified
	// prefix.",
	//
	// Use ListArtifactsAll() to get all FileInfo instances, which will iterate over every result page.
	ListArtifacts(ctx context.Context, request ListArtifactsRequest) (*ListArtifactsResponse, error)

	// List experiments.
	//
	// Gets a list of all experiments.
	//
	// Use ListExperimentsAll() to get all Experiment instances, which will iterate over every result page.
	ListExperiments(ctx context.Context, request ListExperimentsRequest) (*ListExperimentsResponse, error)

	// Log a batch.
	//
	// Logs a batch of metrics, params, and tags for a run. If any data failed
	// to be persisted, the server will respond with an error (non-200 status
	// code).
	//
	// In case of error (due to internal server error or an invalid request),
	// partial data may be written.
	//
	// You can write metrics, params, and tags in interleaving fashion, but
	// within a given entity type are guaranteed to follow the order specified
	// in the request body.
	//
	// The overwrite behavior for metrics, params, and tags is as follows:
	//
	// * Metrics: metric values are never overwritten. Logging a metric (key,
	// value, timestamp) appends to the set of values for the metric with the
	// provided key.
	//
	// * Tags: tag values can be overwritten by successive writes to the same
	// tag key. That is, if multiple tag values with the same key are provided
	// in the same API request, the last-provided tag value is written. Logging
	// the same tag (key, value) is permitted. Specifically, logging a tag is
	// idempotent.
	//
	// * Parameters: once written, param values cannot be changed (attempting to
	// overwrite a param value will result in an error). However, logging the
	// same param (key, value) is permitted. Specifically, logging a param is
	// idempotent.
	//
	// Request Limits ------------------------------- A single JSON-serialized
	// API request may be up to 1 MB in size and contain:
	//
	// * No more than 1000 metrics, params, and tags in total * Up to 1000
	// metrics - Up to 100 params * Up to 100 tags
	//
	// For example, a valid request might contain 900 metrics, 50 params, and 50
	// tags, but logging 900 metrics, 50 params, and 51 tags is invalid.
	//
	// The following limits also apply to metric, param, and tag keys and
	// values:
	//
	// * Metric keyes, param keys, and tag keys can be up to 250 characters in
	// length * Parameter and tag values can be up to 250 characters in length
	LogBatch(ctx context.Context, request LogBatch) error

	// Log a metric.
	//
	// Logs a metric for a run. A metric is a key-value pair (string key, float
	// value) with an associated timestamp. Examples include the various metrics
	// that represent ML model accuracy. A metric can be logged multiple times.
	LogMetric(ctx context.Context, request LogMetric) error

	// Log a model.
	//
	// **NOTE:** Experimental: This API may change or be removed in a future
	// release without warning.
	LogModel(ctx context.Context, request LogModel) error

	// Log a param.
	//
	// Logs a param used for a run. A param is a key-value pair (string key,
	// string value). Examples include hyperparameters used for ML model
	// training and constant dates and values used in an ETL pipeline. A param
	// can be logged only once for a run.
	LogParam(ctx context.Context, request LogParam) error

	// Restores an experiment.
	//
	// "Restore an experiment marked for deletion. This also restores associated
	// metadata, runs, metrics, params, and tags. If experiment uses FileStore,
	// underlying artifacts associated with experiment are also restored.
	//
	// Throws `RESOURCE_DOES_NOT_EXIST` if experiment was never created or was
	// permanently deleted.",
	RestoreExperiment(ctx context.Context, request RestoreExperiment) error

	// Restore a run.
	//
	// Restores a deleted run.
	RestoreRun(ctx context.Context, request RestoreRun) error

	// Search experiments.
	//
	// Searches for experiments that satisfy specified search criteria.
	//
	// Use SearchExperimentsAll() to get all Experiment instances, which will iterate over every result page.
	SearchExperiments(ctx context.Context, request SearchExperiments) (*SearchExperimentsResponse, error)

	// Search for runs.
	//
	// Searches for runs that satisfy expressions.
	//
	// Search expressions can use `mlflowMetric` and `mlflowParam` keys.",
	//
	// Use SearchRunsAll() to get all Run instances, which will iterate over every result page.
	SearchRuns(ctx context.Context, request SearchRuns) (*SearchRunsResponse, error)

	// Set a tag.
	//
	// Sets a tag on an experiment. Experiment tags are metadata that can be
	// updated.
	SetExperimentTag(ctx context.Context, request SetExperimentTag) error

	// Set a tag.
	//
	// Sets a tag on a run. Tags are run metadata that can be updated during a
	// run and after a run completes.
	SetTag(ctx context.Context, request SetTag) error

	// Update an experiment.
	//
	// Updates experiment metadata.
	UpdateExperiment(ctx context.Context, request UpdateExperiment) error

	// Update a run.
	//
	// Updates run metadata.
	UpdateRun(ctx context.Context, request UpdateRun) (*UpdateRunResponse, error)
}

type ModelRegistryService interface {

	// Approve transition request.
	//
	// Approves a model version stage transition request.
	ApproveTransitionRequest(ctx context.Context, request ApproveTransitionRequest) (*ApproveTransitionRequestResponse, error)

	// Post a comment.
	//
	// Posts a comment on a model version. A comment can be submitted either by
	// a user or programmatically to display relevant information about the
	// model. For example, test results or deployment errors.
	CreateComment(ctx context.Context, request CreateComment) (*CreateCommentResponse, error)

	// Create a model.
	//
	// Creates a new registered model with the name specified in the request
	// body.
	//
	// Throws `RESOURCE_ALREADY_EXISTS` if a registered model with the given
	// name exists.
	CreateModel(ctx context.Context, request CreateModelRequest) (*CreateModelResponse, error)

	// Create a model version.
	//
	// Creates a model version.
	CreateModelVersion(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error)

	// Make a transition request.
	//
	// Creates a model version stage transition request.
	CreateTransitionRequest(ctx context.Context, request CreateTransitionRequest) (*CreateTransitionRequestResponse, error)

	// Create a webhook.
	//
	// **NOTE**: This endpoint is in Public Preview.
	//
	// Creates a registry webhook.
	CreateWebhook(ctx context.Context, request CreateRegistryWebhook) (*CreateWebhookResponse, error)

	// Delete a comment.
	//
	// Deletes a comment on a model version.
	DeleteComment(ctx context.Context, request DeleteCommentRequest) error

	// Delete a model.
	//
	// Deletes a registered model.
	DeleteModel(ctx context.Context, request DeleteModelRequest) error

	// Delete a model tag.
	//
	// Deletes the tag for a registered model.
	DeleteModelTag(ctx context.Context, request DeleteModelTagRequest) error

	// Delete a model version.
	//
	// Deletes a model version.
	DeleteModelVersion(ctx context.Context, request DeleteModelVersionRequest) error

	// Delete a model version tag.
	//
	// Deletes a model version tag.
	DeleteModelVersionTag(ctx context.Context, request DeleteModelVersionTagRequest) error

	// Delete a ransition request.
	//
	// Cancels a model version stage transition request.
	DeleteTransitionRequest(ctx context.Context, request DeleteTransitionRequestRequest) error

	// Delete a webhook.
	//
	// **NOTE:** This endpoint is in Public Preview.
	//
	// Deletes a registry webhook.
	DeleteWebhook(ctx context.Context, request DeleteWebhookRequest) error

	// Get the latest version.
	//
	// Gets the latest version of a registered model.
	//
	// Use GetLatestVersionsAll() to get all ModelVersion instances
	GetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) (*GetLatestVersionsResponse, error)

	// Get model.
	//
	// Get the details of a model. This is a Databricks Workspace version of the
	// [MLflow endpoint] that also returns the model's Databricks Workspace ID
	// and the permission level of the requesting user on the model.
	//
	// [MLflow endpoint]: https://www.mlflow.org/docs/latest/rest-api.html#get-registeredmodel
	GetModel(ctx context.Context, request GetModelRequest) (*GetModelResponse, error)

	// Get a model version.
	//
	// Get a model version.
	GetModelVersion(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error)

	// Get a model version URI.
	//
	// Gets a URI to download the model version.
	GetModelVersionDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error)

	// List models.
	//
	// Lists all available registered models, up to the limit specified in
	// __max_results__.
	//
	// Use ListModelsAll() to get all Model instances, which will iterate over every result page.
	ListModels(ctx context.Context, request ListModelsRequest) (*ListModelsResponse, error)

	// List transition requests.
	//
	// Gets a list of all open stage transition requests for the model version.
	//
	// Use ListTransitionRequestsAll() to get all Activity instances
	ListTransitionRequests(ctx context.Context, request ListTransitionRequestsRequest) (*ListTransitionRequestsResponse, error)

	// List registry webhooks.
	//
	// **NOTE:** This endpoint is in Public Preview.
	//
	// Lists all registry webhooks.
	//
	// Use ListWebhooksAll() to get all RegistryWebhook instances, which will iterate over every result page.
	ListWebhooks(ctx context.Context, request ListWebhooksRequest) (*ListRegistryWebhooks, error)

	// Reject a transition request.
	//
	// Rejects a model version stage transition request.
	RejectTransitionRequest(ctx context.Context, request RejectTransitionRequest) (*RejectTransitionRequestResponse, error)

	// Rename a model.
	//
	// Renames a registered model.
	RenameModel(ctx context.Context, request RenameModelRequest) (*RenameModelResponse, error)

	// Searches model versions.
	//
	// Searches for specific model versions based on the supplied __filter__.
	//
	// Use SearchModelVersionsAll() to get all ModelVersion instances, which will iterate over every result page.
	SearchModelVersions(ctx context.Context, request SearchModelVersionsRequest) (*SearchModelVersionsResponse, error)

	// Search models.
	//
	// Search for registered models based on the specified __filter__.
	//
	// Use SearchModelsAll() to get all Model instances, which will iterate over every result page.
	SearchModels(ctx context.Context, request SearchModelsRequest) (*SearchModelsResponse, error)

	// Set a tag.
	//
	// Sets a tag on a registered model.
	SetModelTag(ctx context.Context, request SetModelTagRequest) error

	// Set a version tag.
	//
	// Sets a model version tag.
	SetModelVersionTag(ctx context.Context, request SetModelVersionTagRequest) error

	// Test a webhook.
	//
	// **NOTE:** This endpoint is in Public Preview.
	//
	// Tests a registry webhook.
	TestRegistryWebhook(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error)

	// Transition a stage.
	//
	// Transition a model version's stage. This is a Databricks Workspace
	// version of the [MLflow endpoint] that also accepts a comment associated
	// with the transition to be recorded.",
	//
	// [MLflow endpoint]: https://www.mlflow.org/docs/latest/rest-api.html#transition-modelversion-stage
	TransitionStage(ctx context.Context, request TransitionModelVersionStageDatabricks) (*TransitionStageResponse, error)

	// Update a comment.
	//
	// Post an edit to a comment on a model version.
	UpdateComment(ctx context.Context, request UpdateComment) (*UpdateCommentResponse, error)

	// Update model.
	//
	// Updates a registered model.
	UpdateModel(ctx context.Context, request UpdateModelRequest) error

	// Update model version.
	//
	// Updates the model version.
	UpdateModelVersion(ctx context.Context, request UpdateModelVersionRequest) error

	// Update a webhook.
	//
	// **NOTE:** This endpoint is in Public Preview.
	//
	// Updates a registry webhook.
	UpdateWebhook(ctx context.Context, request UpdateRegistryWebhook) error
}

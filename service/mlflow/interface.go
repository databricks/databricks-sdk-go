// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlflow

import (
	"context"
)

type ExperimentsService interface {

	// Create experiment
	//
	// Creates an experiment with a name. Returns the ID of the newly created
	// experiment. Validates that another experiment with the same name does not
	// already exist and fails if another experiment with the same name already
	// exists.
	//
	// Throws ``RESOURCE_ALREADY_EXISTS`` if a experiment with the given name
	// exists.
	Create(ctx context.Context, request CreateExperiment) (*CreateExperimentResponse, error)

	// Delete an experiment
	//
	// Marks an experiment and associated metadata, runs, metrics, params, and
	// tags for deletion. If the experiment uses FileStore, artifacts associated
	// with experiment are also deleted.
	Delete(ctx context.Context, request DeleteExperiment) error

	// Get an experiment
	//
	// Gets metadata for an experiment. This method works on deleted
	// experiments.
	Get(ctx context.Context, request GetExperimentRequest) (*Experiment, error)

	// Get metadata
	//
	// "Gets metadata for an experiment.
	//
	// This endpoint will return deleted experiments, but prefers the active
	// experiment if an active and deleted experiment share the same name. If
	// multiple deleted\nexperiments share the same name, the API will return
	// one of them.
	//
	// Throws ``RESOURCE_DOES_NOT_EXIST`` if no experiment with the specified
	// name exists.S
	GetByName(ctx context.Context, request GetByNameRequest) (*GetExperimentByNameResponse, error)

	// List experiments
	//
	// Gets a list of all experiments.
	//
	// Use ListAll() to get all Experiment instances, which will iterate over every result page.
	List(ctx context.Context, request ListExperimentsRequest) (*ListExperimentsResponse, error)

	// Restores an experiment
	//
	// "Restore an experiment marked for deletion. This also
	// restores\nassociated metadata, runs, metrics, params, and tags. If
	// experiment uses FileStore, underlying\nartifacts associated with
	// experiment are also restored.\n\nThrows ``RESOURCE_DOES_NOT_EXIST`` if
	// experiment was never created or was permanently deleted.",
	Restore(ctx context.Context, request RestoreExperiment) error

	// Search experiments
	//
	// Searches for experiments that satisfy specified search criteria.
	//
	// Use SearchAll() to get all Experiment instances, which will iterate over every result page.
	Search(ctx context.Context, request SearchExperiments) (*SearchExperimentsResponse, error)

	// Set a tag
	//
	// Sets a tag on an experiment. Experiment tags are metadata that can be
	// updated.
	SetExperimentTag(ctx context.Context, request SetExperimentTag) error

	// Update an experiment
	//
	// Updates experiment metadata.
	Update(ctx context.Context, request UpdateExperiment) error
}

type MLflowArtifactsService interface {

	// Get all artifacts
	//
	// List artifacts for a run. Takes an optional ``artifact_path`` prefix. If
	// it is specified, the response contains only artifacts with the specified
	// prefix.",
	//
	// Use ListAll() to get all FileInfo instances, which will iterate over every result page.
	List(ctx context.Context, request ListArtifactsRequest) (*ListArtifactsResponse, error)
}

// These endpoints are modified versions of the MLflow API that accept
// additional input parameters or return additional information.
type MLflowDatabricksService interface {

	// Get model
	//
	// Get the details of a model. This is a Databricks Workspace version of the
	// [MLflow
	// endpoint](https://www.mlflow.org/docs/latest/rest-api.html#get-registeredmodel)
	// that also returns the model's Databricks Workspace ID and the permission
	// level of the requesting user on the model.
	Get(ctx context.Context, request GetRequest) (*GetResponse, error)

	// Transition a stage
	//
	// Transition a model version's stage. This is a Databricks Workspace
	// version of the [MLflow
	// endpoint](https://www.mlflow.org/docs/latest/rest-api.html#transition-modelversion-stage)
	// that also accepts a comment associated with the transition to be
	// recorded.",
	TransitionStage(ctx context.Context, request TransitionModelVersionStageDatabricks) (*TransitionStageResponse, error)
}

type MLflowMetricsService interface {

	// Get all history
	//
	// Gets a list of all values for the specified metric for a given run.
	GetHistory(ctx context.Context, request GetHistoryRequest) (*GetMetricHistoryResponse, error)
}

type MLflowRunsService interface {

	// Create a run
	//
	// Creates a new run within an experiment. A run is usually a single
	// execution of a machine learning or data ETL pipeline. MLflow uses runs to
	// track the `mlflowParam`, `mlflowMetric` and `mlflowRunTag` associated
	// with a single execution.
	Create(ctx context.Context, request CreateRun) (*CreateRunResponse, error)

	// Delete a run
	//
	// Marks a run for deletion.
	Delete(ctx context.Context, request DeleteRun) error

	// Delete a tag
	//
	// Deletes a tag on a run. Tags are run metadata that can be updated during
	// a run and after a run completes.
	DeleteTag(ctx context.Context, request DeleteTag) error

	// Get a run
	//
	// "Gets the metadata, metrics, params, and tags for a run. In the case
	// where multiple metrics with the same key are logged for a run, return
	// only the value with the latest timestamp.
	//
	// If there are multiple values with the latest timestamp, return the
	// maximum of these values.
	Get(ctx context.Context, request GetRunRequest) (*GetRunResponse, error)

	// Log a batch
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
	// in the request body. That is, for an API request like:
	//
	// ``` { "run_id": "2a14ed5c6a87499199e0106c3501eab8", "metrics": [ { "key":
	// "mae", "value": 2.5, "timestamp": 1552550804 }, { "key": "rmse", "value":
	// 2.7, "timestamp": 1552550804 }], "params": [ { \"key\": \"model_class\",
	// \"value\": \"LogisticRegression\" }] } ```
	//
	// the server is guaranteed to write metric "rmse" after "mae", though it
	// may write the param __model_class__ before both metrics, after "mae", or
	// after both metrics.
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
	// metrics\n- Up to 100 params * Up to 100 tags
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

	// Log a metric
	//
	// Logs a metric for a run. A metric is a key-value pair (string key, float
	// value) with an associated timestamp. Examples include the various metrics
	// that represent ML model accuracy. A metric can be logged multiple times.
	LogMetric(ctx context.Context, request LogMetric) error

	// Log a model
	//
	// **NOTE:** Experimental: This API may change or be removed in a future
	// release without warning.
	LogModel(ctx context.Context, request LogModel) error

	// Log a param
	//
	// Logs a param used for a run. A param is a key-value pair (string key,
	// string value). Examples include hyperparameters used for ML model
	// training and constant dates and values used in an ETL pipeline. A param
	// can be logged only once for a run.
	LogParameter(ctx context.Context, request LogParam) error

	// Restore a run
	//
	// Restores a deleted run.
	Restore(ctx context.Context, request RestoreRun) error

	// Search for runs
	//
	// Searches for runs that satisfy expressions.
	//
	// Search expressions can use `mlflowMetric` and `mlflowParam` keys.",
	//
	// Use SearchAll() to get all Run instances, which will iterate over every result page.
	Search(ctx context.Context, request SearchRuns) (*SearchRunsResponse, error)

	// Set a tag
	//
	// Sets a tag on a run. Tags are run metadata that can be updated during a
	// run and after a run completes.
	SetTag(ctx context.Context, request SetTag) error

	// Update a run
	//
	// Updates run metadata.
	Update(ctx context.Context, request UpdateRun) (*UpdateRunResponse, error)
}

type ModelVersionCommentsService interface {

	// Post a comment
	//
	// Posts a comment on a model version. A comment can be submitted either by
	// a user or programmatically to display relevant information about the
	// model. For example, test results or deployment errors.
	Create(ctx context.Context, request CreateComment) (*CreateResponse, error)

	// Delete a comment
	//
	// Deletes a comment on a model version.
	Delete(ctx context.Context, request DeleteRequest) error

	// Update a comment
	//
	// Post an edit to a comment on a model version.
	Update(ctx context.Context, request UpdateComment) (*UpdateResponse, error)
}

type ModelVersionsService interface {

	// Create a model version
	//
	// Creates a model version.
	Create(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error)

	// Delete a model version.
	//
	// Deletes a model version.
	Delete(ctx context.Context, request DeleteModelVersionRequest) error

	// Delete a model version tag
	//
	// Deletes a model version tag.
	DeleteTag(ctx context.Context, request DeleteModelVersionTagRequest) error

	// Get a model version
	//
	// Get a model version.
	Get(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error)

	// Get a model version URI
	//
	// Gets a URI to download the model version.
	GetDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error)

	// Searches model versions
	//
	// Searches for specific model versions based on the supplied __filter__.
	//
	// Use SearchAll() to get all ModelVersion instances, which will iterate over every result page.
	Search(ctx context.Context, request SearchModelVersionsRequest) (*SearchModelVersionsResponse, error)

	// Set a version tag
	//
	// Sets a model version tag.
	SetTag(ctx context.Context, request SetModelVersionTagRequest) error

	// Transition a stage
	//
	// Transition to the next model stage.
	TransitionStage(ctx context.Context, request TransitionModelVersionStage) (*TransitionModelVersionStageResponse, error)

	// Update model version
	//
	// Updates the model version.
	Update(ctx context.Context, request UpdateModelVersionRequest) error
}

type RegisteredModelsService interface {

	// Create a model
	//
	// Creates a new registered model with the name specified in the request
	// body.
	//
	// Throws ``RESOURCE_ALREADY_EXISTS`` if a registered model with the given
	// name exists.
	Create(ctx context.Context, request CreateRegisteredModelRequest) (*CreateRegisteredModelResponse, error)

	// Delete a model
	//
	// Deletes a registered model.
	Delete(ctx context.Context, request DeleteRegisteredModelRequest) error

	// Delete a model tag
	//
	// Deletes the tag for a registered model.
	DeleteTag(ctx context.Context, request DeleteRegisteredModelTagRequest) error

	// Get a model
	//
	// Gets the registered model that matches the specified ID.
	Get(ctx context.Context, request GetRegisteredModelRequest) (*GetRegisteredModelResponse, error)

	// Get the latest version
	//
	// Gets the latest version of a registered model.
	//
	// Use GetLatestVersionsAll() to get all ModelVersion instances
	GetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) (*GetLatestVersionsResponse, error)

	// List models
	//
	// Lists all available registered models, up to the limit specified in
	// __max_results__.
	//
	// Use ListAll() to get all RegisteredModel instances, which will iterate over every result page.
	List(ctx context.Context, request ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error)

	// Rename a model
	//
	// Renames a registered model.
	Rename(ctx context.Context, request RenameRegisteredModelRequest) (*RenameRegisteredModelResponse, error)

	// Search models
	//
	// Search for registered models based on the specified __filter__.
	//
	// Use SearchAll() to get all RegisteredModel instances, which will iterate over every result page.
	Search(ctx context.Context, request SearchRegisteredModelsRequest) (*SearchRegisteredModelsResponse, error)

	// Set a tag
	//
	// Sets a tag on a registered model.
	SetTag(ctx context.Context, request SetRegisteredModelTagRequest) error

	// Update model
	//
	// Updates a registered model.
	Update(ctx context.Context, request UpdateRegisteredModelRequest) error
}

type RegistryWebhooksService interface {

	// Create a webhook
	//
	// **NOTE**: This endpoint is in Public Preview.
	//
	// Creates a registry webhook.
	Create(ctx context.Context, request CreateRegistryWebhook) (*CreateResponse, error)

	// Delete a webhook
	//
	// **NOTE:** This endpoint is in Public Preview.
	//
	// Deletes a registry webhook.
	Delete(ctx context.Context, request DeleteRequest) error

	// List registry webhooks
	//
	// **NOTE:** This endpoint is in Public Preview.
	//
	// Lists all registry webhooks.
	//
	// Use ListAll() to get all RegistryWebhook instances, which will iterate over every result page.
	List(ctx context.Context, request ListRequest) (*ListRegistryWebhooks, error)

	// Test a webhook
	//
	// **NOTE:** This endpoint is in Public Preview.
	//
	// Tests a registry webhook.
	Test(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error)

	// Update a webhook
	//
	// **NOTE:** This endpoint is in Public Preview.
	//
	// Updates a registry webhook.
	Update(ctx context.Context, request UpdateRegistryWebhook) error
}

type TransitionRequestsService interface {

	// Approve transition requests
	//
	// Approves a model version stage transition request.
	Approve(ctx context.Context, request ApproveTransitionRequest) (*ApproveResponse, error)

	// Make a transition request
	//
	// Creates a model version stage transition request.
	Create(ctx context.Context, request CreateTransitionRequest) (*CreateResponse, error)

	// Delete a ransition request
	//
	// Cancels a model version stage transition request.
	Delete(ctx context.Context, request DeleteRequest) error

	// List transition requests
	//
	// Gets a list of all open stage transition requests for the model version.
	//
	// Use ListAll() to get all Activity instances
	List(ctx context.Context, request ListRequest) (*ListResponse, error)

	// Reject a transition request
	//
	// Rejects a model version stage transition request.
	Reject(ctx context.Context, request RejectTransitionRequest) (*RejectResponse, error)
}

// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlflow

import (
	"context"
)

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type ExperimentsService interface {

	// Create an experiment with a name. Returns the ID of the newly created
	// experiment. Validates that another experiment with the same name does not
	// already exist and fails if another experiment with the same name already
	// exists. Throws ``RESOURCE_ALREADY_EXISTS`` if a experiment with the given
	// name exists.
	Create(ctx context.Context, request CreateExperiment) (*CreateExperimentResponse, error)

	// Mark an experiment and associated metadata, runs, metrics, params, and
	// tags for deletion. If the experiment uses FileStore, artifacts associated
	// with experiment are also deleted.
	Delete(ctx context.Context, request DeleteExperiment) error

	// DeleteByExperimentId calls Delete, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteByExperimentId(ctx context.Context, experimentId string) error

	// Get metadata for an experiment. This method works on deleted experiments.
	Get(ctx context.Context, request GetExperimentRequest) (*Experiment, error)

	// GetByExperimentId calls Get, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByExperimentId(ctx context.Context, experimentId string) (*Experiment, error)

	// Get metadata for an experiment. This endpoint will return deleted
	// experiments, but prefers the active experiment if an active and deleted
	// experiment share the same name. If multiple deleted experiments share the
	// same name, the API will return one of them. Throws
	// ``RESOURCE_DOES_NOT_EXIST`` if no experiment with the specified name
	// exists.
	GetByName(ctx context.Context, request GetByNameRequest) (*GetExperimentByNameResponse, error)

	// GetByNameByExperimentName calls GetByName, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByNameByExperimentName(ctx context.Context, experimentName string) (*GetExperimentByNameResponse, error)

	// Get a list of all experiments.
	List(ctx context.Context, request ListExperimentsRequest) (*ListExperimentsResponse, error)

	// Restore an experiment marked for deletion. This also restores associated
	// metadata, runs, metrics, params, and tags. If experiment uses FileStore,
	// underlying artifacts associated with experiment are also restored. Throws
	// ``RESOURCE_DOES_NOT_EXIST`` if experiment was never created or was
	// permanently deleted.
	Restore(ctx context.Context, request RestoreExperiment) error

	// RestoreByExperimentId calls Restore, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	RestoreByExperimentId(ctx context.Context, experimentId string) error

	// Search for experiments that satisfy specified search criteria.
	Search(ctx context.Context, request SearchExperiments) (*SearchExperimentsResponse, error)

	// Set a tag on an experiment. Experiment tags are metadata that can be
	// updated.
	SetExperimentTag(ctx context.Context, request SetExperimentTag) error

	// Update experiment metadata.
	Update(ctx context.Context, request UpdateExperiment) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type MLflowArtifactsService interface {

	// List artifacts for a run. Takes an optional ``artifact_path`` prefix
	// which if specified, the response contains only artifacts with the
	// specified prefix.
	List(ctx context.Context, request ListArtifactsRequest) (*ListArtifactsResponse, error)
}

// These endpoints are modified versions of the MLflow API that accept
// additional input parameters or return additional information.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type MLflowDatabricksService interface {
	Get(ctx context.Context, request GetRequest) (*GetResponse, error)

	// GetByName calls Get, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByName(ctx context.Context, name string) (*GetResponse, error)

	// Transition a model version's stage. This is a <Workspace> version of the
	// [MLflow
	// endpoint](https://www.mlflow.org/docs/latest/rest-api.html#transition-modelversion-stage)
	// that also accepts a comment associated with the transition to be
	// recorded.
	TransitionStage(ctx context.Context, request TransitionModelVersionStageDatabricks) (*TransitionStageResponse, error)
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type MLflowMetricsService interface {

	// Get a list of all values for the specified metric for a given run.
	GetHistory(ctx context.Context, request GetHistoryRequest) (*GetMetricHistoryResponse, error)
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type MLflowRunsService interface {

	// Create a new run within an experiment. A run is usually a single
	// execution of a machine learning or data ETL pipeline. MLflow uses runs to
	// track :ref:`mlflowParam`, :ref:`mlflowMetric`, and :ref:`mlflowRunTag`
	// associated with a single execution.
	Create(ctx context.Context, request CreateRun) (*CreateRunResponse, error)

	// Mark a run for deletion.
	Delete(ctx context.Context, request DeleteRun) error

	// DeleteByRunId calls Delete, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteByRunId(ctx context.Context, runId string) error

	// Delete a tag on a run. Tags are run metadata that can be updated during a
	// run and after a run completes.
	DeleteTag(ctx context.Context, request DeleteTag) error

	// Get metadata, metrics, params, and tags for a run. In the case where
	// multiple metrics with the same key are logged for a run, return only the
	// value with the latest timestamp. If there are multiple values with the
	// latest timestamp, return the maximum of these values.
	Get(ctx context.Context, request GetRunRequest) (*GetRunResponse, error)

	// Log a batch of metrics, params, and tags for a run. If any data failed to
	// be persisted, the server will respond with an error (non-200 status
	// code). In case of error (due to internal server error or an invalid
	// request), partial data may be written. You can write metrics, params, and
	// tags in interleaving fashion, but within a given entity type are
	// guaranteed to follow the order specified in the request body. That is,
	// for an API request like .. code-block:: json { "run_id":
	// "2a14ed5c6a87499199e0106c3501eab8", "metrics": [ {"key": "mae", "value":
	// 2.5, "timestamp": 1552550804}, {"key": "rmse", "value": 2.7, "timestamp":
	// 1552550804}, ], "params": [ {"key": "model_class", "value":
	// "LogisticRegression"}, ] } the server is guaranteed to write metric
	// "rmse" after "mae", though it may write param "model_class" before both
	// metrics, after "mae", or after both metrics. The overwrite behavior for
	// metrics, params, and tags is as follows: - Metrics: metric values are
	// never overwritten. Logging a metric (key, value, timestamp) appends to
	// the set of values for the metric with the provided key. - Tags: tag
	// values can be overwritten by successive writes to the same tag key. That
	// is, if multiple tag values with the same key are provided in the same API
	// request, the last-provided tag value is written. Logging the same tag
	// (key, value) is permitted - that is, logging a tag is idempotent. -
	// Params: once written, param values cannot be changed (attempting to
	// overwrite a param value will result in an error). However, logging the
	// same param (key, value) is permitted - that is, logging a param is
	// idempotent. Request Limits -------------- A single JSON-serialized API
	// request may be up to 1 MB in size and contain: - No more than 1000
	// metrics, params, and tags in total - Up to 1000 metrics - Up to 100
	// params - Up to 100 tags For example, a valid request might contain 900
	// metrics, 50 params, and 50 tags, but logging 900 metrics, 50 params, and
	// 51 tags is invalid. The following limits also apply to metric, param, and
	// tag keys and values: - Metric, param, and tag keys can be up to 250
	// characters in length - Param and tag values can be up to 250 characters
	// in length
	LogBatch(ctx context.Context, request LogBatch) error

	// Log a metric for a run. A metric is a key-value pair (string key, float
	// value) with an associated timestamp. Examples include the various metrics
	// that represent ML model accuracy. A metric can be logged multiple times.
	LogMetric(ctx context.Context, request LogMetric) error

	// .. note:: Experimental: This API may change or be removed in a future
	// release without warning.
	LogModel(ctx context.Context, request LogModel) error

	// Log a param used for a run. A param is a key-value pair (string key,
	// string value). Examples include hyperparameters used for ML model
	// training and constant dates and values used in an ETL pipeline. A param
	// can be logged only once for a run.
	LogParameter(ctx context.Context, request LogParam) error

	// Restore a deleted run.
	Restore(ctx context.Context, request RestoreRun) error

	// RestoreByRunId calls Restore, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	RestoreByRunId(ctx context.Context, runId string) error

	// Search for runs that satisfy expressions. Search expressions can use
	// :ref:`mlflowMetric` and :ref:`mlflowParam` keys.
	Search(ctx context.Context, request SearchRuns) (*SearchRunsResponse, error)

	// Set a tag on a run. Tags are run metadata that can be updated during a
	// run and after a run completes.
	SetTag(ctx context.Context, request SetTag) error

	// Update run metadata.
	Update(ctx context.Context, request UpdateRun) (*UpdateRunResponse, error)
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type ModelVersionCommentsService interface {

	// Make a comment on a model version. A comment can be submitted either by a
	// user or programmatically to display relevant information about the model.
	// For example, test results or deployment errors.
	Create(ctx context.Context, request CreateComment) (*CreateResponse, error)

	// Delete a comment on a model version.
	Delete(ctx context.Context, request DeleteRequest) error

	// DeleteById calls Delete, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteById(ctx context.Context, id string) error

	// Edit a comment on a model version.
	Update(ctx context.Context, request UpdateComment) (*UpdateResponse, error)
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type ModelVersionsService interface {
	Create(ctx context.Context, request CreateModelVersionRequest) (*CreateModelVersionResponse, error)

	Delete(ctx context.Context, request DeleteModelVersionRequest) error

	DeleteTag(ctx context.Context, request DeleteModelVersionTagRequest) error

	Get(ctx context.Context, request GetModelVersionRequest) (*GetModelVersionResponse, error)

	GetDownloadUri(ctx context.Context, request GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error)

	Search(ctx context.Context, request SearchModelVersionsRequest) (*SearchModelVersionsResponse, error)

	SetTag(ctx context.Context, request SetModelVersionTagRequest) error

	TransitionStage(ctx context.Context, request TransitionModelVersionStage) (*TransitionModelVersionStageResponse, error)

	Update(ctx context.Context, request UpdateModelVersionRequest) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type RegisteredModelsService interface {

	// Throws ``RESOURCE_ALREADY_EXISTS`` if a registered model with the given
	// name exists.
	Create(ctx context.Context, request CreateRegisteredModelRequest) (*CreateRegisteredModelResponse, error)

	Delete(ctx context.Context, request DeleteRegisteredModelRequest) error

	// DeleteByName calls Delete, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteByName(ctx context.Context, name string) error

	DeleteTag(ctx context.Context, request DeleteRegisteredModelTagRequest) error

	Get(ctx context.Context, request GetRegisteredModelRequest) (*GetRegisteredModelResponse, error)

	// GetByName calls Get, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByName(ctx context.Context, name string) (*GetRegisteredModelResponse, error)

	GetLatestVersions(ctx context.Context, request GetLatestVersionsRequest) (*GetLatestVersionsResponse, error)

	List(ctx context.Context, request ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error)

	Rename(ctx context.Context, request RenameRegisteredModelRequest) (*RenameRegisteredModelResponse, error)

	Search(ctx context.Context, request SearchRegisteredModelsRequest) (*SearchRegisteredModelsResponse, error)

	SetTag(ctx context.Context, request SetRegisteredModelTagRequest) error

	Update(ctx context.Context, request UpdateRegisteredModelRequest) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type RegistryWebhooksService interface {

	// This endpoint is in Public Preview. Create a registry webhook.
	Create(ctx context.Context, request CreateRegistryWebhook) (*CreateResponse, error)

	// This endpoint is in Public Preview. Delete a registry webhook.
	Delete(ctx context.Context, request DeleteRequest) error

	// DeleteById calls Delete, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteById(ctx context.Context, id string) error

	// This endpoint is in Public Preview. List registry webhooks.
	List(ctx context.Context, request ListRequest) (*ListResponse, error)

	// This endpoint is in Public Preview. Test a registry webhook.
	Test(ctx context.Context, request TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error)

	// This endpoint is in Public Preview. Update a registry webhook.
	Update(ctx context.Context, request UpdateRegistryWebhook) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type TransitionRequestsService interface {

	// Approve model version stage transition request.
	Approve(ctx context.Context, request ApproveTransitionRequest) (*ApproveResponse, error)

	// Make a model version stage transition request.
	Create(ctx context.Context, request CreateTransitionRequest) (*CreateResponse, error)

	// Cancel model version stage transition request.
	Delete(ctx context.Context, request DeleteRequest) error

	// Get all open stage transition requests for the model version.
	List(ctx context.Context, request ListRequest) (*ListResponse, error)

	// Reject model version stage transition request.
	Reject(ctx context.Context, request RejectTransitionRequest) (*RejectResponse, error)
}

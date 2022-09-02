// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlflow

import (
	"context"
)



type ExperimentsService interface {
    // Create an experiment with a name. Returns the ID of the newly created
    // experiment. Validates that another experiment with the same name does not
    // already exist and fails if another experiment with the same name already
    // exists. Throws ``RESOURCE_ALREADY_EXISTS`` if a experiment with the given
    // name exists.
    Create(ctx context.Context, createExperimentRequest CreateExperimentRequest) (*CreateExperimentResponse, error)
    // Mark an experiment and associated metadata, runs, metrics, params, and
    // tags for deletion. If the experiment uses FileStore, artifacts associated
    // with experiment are also deleted.
    Delete(ctx context.Context, deleteExperimentRequest DeleteExperimentRequest) error
    // Get metadata for an experiment. This method works on deleted experiments.
    Get(ctx context.Context, getExperimentRequest GetExperimentRequest) (*GetExperimentResponse, error)
    // Get metadata for an experiment. This endpoint will return deleted
    // experiments, but prefers the active experiment if an active and deleted
    // experiment share the same name. If multiple deleted experiments share the
    // same name, the API will return one of them. Throws
    // ``RESOURCE_DOES_NOT_EXIST`` if no experiment with the specified name
    // exists.
    GetByName(ctx context.Context, getExperimentByNameRequest GetExperimentByNameRequest) (*GetExperimentByNameResponse, error)
    // Get a list of all experiments.
    List(ctx context.Context, listExperimentsRequest ListExperimentsRequest) (*ListExperimentsResponse, error)
    // Restore an experiment marked for deletion. This also restores associated
    // metadata, runs, metrics, params, and tags. If experiment uses FileStore,
    // underlying artifacts associated with experiment are also restored. Throws
    // ``RESOURCE_DOES_NOT_EXIST`` if experiment was never created or was
    // permanently deleted.
    Restore(ctx context.Context, restoreExperimentRequest RestoreExperimentRequest) error
    // Search for experiments that satisfy specified search criteria.
    Search(ctx context.Context, searchExperimentsRequest SearchExperimentsRequest) (*SearchExperimentsResponse, error)
    // Set a tag on an experiment. Experiment tags are metadata that can be
    // updated.
    SetExperimentTag(ctx context.Context, setExperimentTagRequest SetExperimentTagRequest) error
    // Update experiment metadata.
    Update(ctx context.Context, updateExperimentRequest UpdateExperimentRequest) error
}


type MLflowArtifactsService interface {
    // List artifacts for a run. Takes an optional ``artifact_path`` prefix
    // which if specified, the response contains only artifacts with the
    // specified prefix.
    List(ctx context.Context, listArtifactsRequest ListArtifactsRequest) (*ListArtifactsResponse, error)
}

// These endpoints are modified versions of the MLflow API that accept
// additional input parameters or return additional information.
type MLflowDatabricksService interface {
    
    Get(ctx context.Context, getRegisteredModelRequest GetRegisteredModelRequest) (*GetRegisteredModelResponse, error)
    // Transition a model version&#39;s stage. This is a &lt;Workspace&gt; version of the
    // [MLflow
    // endpoint](https://www.mlflow.org/docs/latest/rest-api.html#transition-modelversion-stage)
    // that also accepts a comment associated with the transition to be
    // recorded.
    TransitionStage(ctx context.Context, transitionModelVersionStageRequest TransitionModelVersionStageRequest) (*TransitionModelVersionStageResponse, error)
}


type MLflowMetricsService interface {
    // Get a list of all values for the specified metric for a given run.
    GetHistory(ctx context.Context, getMetricHistoryRequest GetMetricHistoryRequest) (*GetMetricHistoryResponse, error)
}


type MLflowRunsService interface {
    // Create a new run within an experiment. A run is usually a single
    // execution of a machine learning or data ETL pipeline. MLflow uses runs to
    // track :ref:`mlflowParam`, :ref:`mlflowMetric`, and :ref:`mlflowRunTag`
    // associated with a single execution.
    Create(ctx context.Context, createRunRequest CreateRunRequest) (*CreateRunResponse, error)
    // Mark a run for deletion.
    Delete(ctx context.Context, deleteRunRequest DeleteRunRequest) error
    // Delete a tag on a run. Tags are run metadata that can be updated during a
    // run and after a run completes.
    DeleteTag(ctx context.Context, deleteTagRequest DeleteTagRequest) error
    // Get metadata, metrics, params, and tags for a run. In the case where
    // multiple metrics with the same key are logged for a run, return only the
    // value with the latest timestamp. If there are multiple values with the
    // latest timestamp, return the maximum of these values.
    Get(ctx context.Context, getRunRequest GetRunRequest) (*GetRunResponse, error)
    // Log a batch of metrics, params, and tags for a run. If any data failed to
    // be persisted, the server will respond with an error (non-200 status
    // code). In case of error (due to internal server error or an invalid
    // request), partial data may be written. You can write metrics, params, and
    // tags in interleaving fashion, but within a given entity type are
    // guaranteed to follow the order specified in the request body. That is,
    // for an API request like .. code-block:: json { &#34;run_id&#34;:
    // &#34;2a14ed5c6a87499199e0106c3501eab8&#34;, &#34;metrics&#34;: [ {&#34;key&#34;: &#34;mae&#34;, &#34;value&#34;:
    // 2.5, &#34;timestamp&#34;: 1552550804}, {&#34;key&#34;: &#34;rmse&#34;, &#34;value&#34;: 2.7, &#34;timestamp&#34;:
    // 1552550804}, ], &#34;params&#34;: [ {&#34;key&#34;: &#34;model_class&#34;, &#34;value&#34;:
    // &#34;LogisticRegression&#34;}, ] } the server is guaranteed to write metric
    // &#34;rmse&#34; after &#34;mae&#34;, though it may write param &#34;model_class&#34; before both
    // metrics, after &#34;mae&#34;, or after both metrics. The overwrite behavior for
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
    LogBatch(ctx context.Context, logBatchRequest LogBatchRequest) error
    // Log a metric for a run. A metric is a key-value pair (string key, float
    // value) with an associated timestamp. Examples include the various metrics
    // that represent ML model accuracy. A metric can be logged multiple times.
    LogMetric(ctx context.Context, logMetricRequest LogMetricRequest) error
    // .. note:: Experimental: This API may change or be removed in a future
    // release without warning.
    LogModel(ctx context.Context, logModelRequest LogModelRequest) error
    // Log a param used for a run. A param is a key-value pair (string key,
    // string value). Examples include hyperparameters used for ML model
    // training and constant dates and values used in an ETL pipeline. A param
    // can be logged only once for a run.
    LogParameter(ctx context.Context, logParamRequest LogParamRequest) error
    // Restore a deleted run.
    Restore(ctx context.Context, restoreRunRequest RestoreRunRequest) error
    // Search for runs that satisfy expressions. Search expressions can use
    // :ref:`mlflowMetric` and :ref:`mlflowParam` keys.
    Search(ctx context.Context, searchRunsRequest SearchRunsRequest) (*SearchRunsResponse, error)
    // Set a tag on a run. Tags are run metadata that can be updated during a
    // run and after a run completes.
    SetTag(ctx context.Context, setTagRequest SetTagRequest) error
    // Update run metadata.
    Update(ctx context.Context, updateRunRequest UpdateRunRequest) (*UpdateRunResponse, error)
}


type ModelVersionCommentsService interface {
    // Make a comment on a model version. A comment can be submitted either by a
    // user or programmatically to display relevant information about the model.
    // For example, test results or deployment errors.
    Create(ctx context.Context, createCommentRequest CreateCommentRequest) (*CreateCommentResponse, error)
    // Delete a comment on a model version.
    Delete(ctx context.Context, deleteCommentRequest DeleteCommentRequest) error
    // Edit a comment on a model version.
    Update(ctx context.Context, updateCommentRequest UpdateCommentRequest) (*UpdateCommentResponse, error)
}


type ModelVersionsService interface {
    
    Create(ctx context.Context, createModelVersionRequest CreateModelVersionRequest) (*CreateModelVersionResponse, error)
    
    Delete(ctx context.Context, deleteModelVersionRequest DeleteModelVersionRequest) error
    
    DeleteTag(ctx context.Context, deleteModelVersionTagRequest DeleteModelVersionTagRequest) error
    
    Get(ctx context.Context, getModelVersionRequest GetModelVersionRequest) (*GetModelVersionResponse, error)
    
    GetDownloadUri(ctx context.Context, getModelVersionDownloadUriRequest GetModelVersionDownloadUriRequest) (*GetModelVersionDownloadUriResponse, error)
    
    Search(ctx context.Context, searchModelVersionsRequest SearchModelVersionsRequest) (*SearchModelVersionsResponse, error)
    
    SetTag(ctx context.Context, setModelVersionTagRequest SetModelVersionTagRequest) error
    
    TransitionStage(ctx context.Context, transitionModelVersionStageRequest TransitionModelVersionStageRequest) (*TransitionModelVersionStageResponse, error)
    
    Update(ctx context.Context, updateModelVersionRequest UpdateModelVersionRequest) error
}


type RegisteredModelsService interface {
    // Throws ``RESOURCE_ALREADY_EXISTS`` if a registered model with the given
    // name exists.
    Create(ctx context.Context, createRegisteredModelRequest CreateRegisteredModelRequest) (*CreateRegisteredModelResponse, error)
    
    Delete(ctx context.Context, deleteRegisteredModelRequest DeleteRegisteredModelRequest) error
    
    DeleteTag(ctx context.Context, deleteRegisteredModelTagRequest DeleteRegisteredModelTagRequest) error
    
    Get(ctx context.Context, getRegisteredModelRequest GetRegisteredModelRequest) (*GetRegisteredModelResponse, error)
    
    GetLatestVersions(ctx context.Context, getLatestVersionsRequest GetLatestVersionsRequest) (*GetLatestVersionsResponse, error)
    
    List(ctx context.Context, listRegisteredModelsRequest ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error)
    
    Rename(ctx context.Context, renameRegisteredModelRequest RenameRegisteredModelRequest) (*RenameRegisteredModelResponse, error)
    
    Search(ctx context.Context, searchRegisteredModelsRequest SearchRegisteredModelsRequest) (*SearchRegisteredModelsResponse, error)
    
    SetTag(ctx context.Context, setRegisteredModelTagRequest SetRegisteredModelTagRequest) error
    
    Update(ctx context.Context, updateRegisteredModelRequest UpdateRegisteredModelRequest) error
}


type RegistryWebhooksService interface {
    // This endpoint is in Public Preview. Create a registry webhook.
    Create(ctx context.Context, createRegistryWebhookRequest CreateRegistryWebhookRequest) (*CreateRegistryWebhookResponse, error)
    // This endpoint is in Public Preview. Delete a registry webhook.
    Delete(ctx context.Context, deleteRegistryWebhookRequest DeleteRegistryWebhookRequest) error
    // This endpoint is in Public Preview. List registry webhooks.
    List(ctx context.Context, listRegistryWebhooksRequest ListRegistryWebhooksRequest) (*ListRegistryWebhooksResponse, error)
    // This endpoint is in Public Preview. Test a registry webhook.
    Test(ctx context.Context, testRegistryWebhookRequest TestRegistryWebhookRequest) (*TestRegistryWebhookResponse, error)
    // This endpoint is in Public Preview. Update a registry webhook.
    Update(ctx context.Context, updateRegistryWebhookRequest UpdateRegistryWebhookRequest) error
}


type TransitionRequestsService interface {
    // Approve model version stage transition request.
    Approve(ctx context.Context, approveTransitionRequestRequest ApproveTransitionRequestRequest) (*ApproveTransitionRequestResponse, error)
    // Make a model version stage transition request.
    Create(ctx context.Context, createTransitionRequestRequest CreateTransitionRequestRequest) (*CreateTransitionRequestResponse, error)
    // Cancel model version stage transition request.
    Delete(ctx context.Context, deleteTransitionRequestRequest DeleteTransitionRequestRequest) error
    // Get all open stage transition requests for the model version.
    List(ctx context.Context, getTransitionRequestsRequest GetTransitionRequestsRequest) (*GetTransitionRequestsResponse, error)
    // Reject model version stage transition request.
    Reject(ctx context.Context, rejectTransitionRequestRequest RejectTransitionRequestRequest) (*RejectTransitionRequestResponse, error)
}

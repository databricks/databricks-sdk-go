// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlflowext

// all definitions in this file are in alphabetical order

type GetRegisteredModelRequest struct {
    // Name of the model.
    Name string ` url:"name,omitempty"`
}


type GetRegisteredModelResponse struct {
    
    RegisteredModel *RegisteredModelDatabricks `json:"registered_model,omitempty"`
}


type ModelVersion struct {
    // Creation time of the object, as a Unix timestamp in milliseconds.
    CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
    
    CurrentStage Stage `json:"current_stage,omitempty"`
    // User-specified description for the object.
    Description string `json:"description,omitempty"`
    // Time of the object at last update, as a Unix timestamp in milliseconds.
    LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
    // Name of the model.
    Name string `json:"name,omitempty"`
    // Unique identifier for the MLflow tracking run associated with the source
    // model artifacts.
    RunId string `json:"run_id,omitempty"`
    // URL of the run associated with the model artifacts, potentially in
    // another workspace.
    RunLink string `json:"run_link,omitempty"`
    // URI that indicates the location of the source model artifacts. This is
    // used when creating the model version.
    Source string `json:"source,omitempty"`
    
    Status Status `json:"status,omitempty"`
    // Details on the current status, for example why registration failed.
    StatusMessage string `json:"status_message,omitempty"`
    
    Tags []ModelVersionTag `json:"tags,omitempty"`
    // The username of the user that created the object.
    UserId string `json:"user_id,omitempty"`
    // Version of the model.
    Version string `json:"version,omitempty"`
}


type ModelVersionDatabricks struct {
    // Creation time of the object, as a Unix timestamp in milliseconds.
    CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
    
    CurrentStage Stage `json:"current_stage,omitempty"`
    // User-specified description for the object.
    Description string `json:"description,omitempty"`
    // Time of the object at last update, as a Unix timestamp in milliseconds.
    LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
    // Name of the model.
    Name string `json:"name,omitempty"`
    
    PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
    // Unique identifier for the MLflow tracking run associated with the source
    // model artifacts.
    RunId string `json:"run_id,omitempty"`
    // URL of the run associated with the model artifacts, potentially in
    // another workspace.
    RunLink string `json:"run_link,omitempty"`
    // URI that indicates the location of the source model artifacts. This is
    // used when creating the model version.
    Source string `json:"source,omitempty"`
    
    Status Status `json:"status,omitempty"`
    // Details on the current status, for example why registration failed.
    StatusMessage string `json:"status_message,omitempty"`
    
    Tags []ModelVersionTag `json:"tags,omitempty"`
    // The username of the user that created the object.
    UserId string `json:"user_id,omitempty"`
    // Version of the model.
    Version string `json:"version,omitempty"`
}


type ModelVersionTag struct {
    // Key for the tag.
    Key string `json:"key,omitempty"`
    // Value for the tag.
    Value string `json:"value,omitempty"`
}

// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).
type PermissionLevel string

// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).
const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`
// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).
const PermissionLevelCanEdit PermissionLevel = `CAN_EDIT`
// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).
const PermissionLevelCanRead PermissionLevel = `CAN_READ`
// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).
const PermissionLevelCanManageStagingVersions PermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`
// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).
const PermissionLevelCanManageProductionVersions PermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

type RegisteredModelDatabricks struct {
    // Creation time of the object, as a Unix timestamp in milliseconds.
    CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
    // User-specified description for the object.
    Description string `json:"description,omitempty"`
    // Unique identifier for the object.
    Id string `json:"id,omitempty"`
    // Time of the object at last update, as a Unix timestamp in milliseconds.
    LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
    // Array of model versions, each the latest version for its stage.
    LatestVersions []ModelVersion `json:"latest_versions,omitempty"`
    // Name of the model.
    Name string `json:"name,omitempty"`
    
    PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
    // Array of tags associated with the model.
    Tags []RegisteredModelTag `json:"tags,omitempty"`
    // The username of the user that created the object.
    UserId string `json:"user_id,omitempty"`
}


type RegisteredModelTag struct {
    // Key for the tag.
    Key string `json:"key,omitempty"`
    // Value for the tag.
    Value string `json:"value,omitempty"`
}


type TransitionModelVersionStageRequest struct {
    // Specifies whether to archive all current model versions in the target
    // stage.
    ArchiveExistingVersions bool `json:"archive_existing_versions"`
    // User-provided comment on the action.
    Comment string `json:"comment,omitempty"`
    // Name of the model.
    Name string `json:"name"`
    
    Stage Stage `json:"stage"`
    // Version of the model.
    Version string `json:"version"`
}


type TransitionModelVersionStageResponse struct {
    
    ModelVersion *ModelVersionDatabricks `json:"model_version,omitempty"`
}

// Stage of the model version. Valid values are: * `None`: The initial stage of
// a model version. * `Staging`: Staging or pre-production stage. *
// `Production`: Production stage. * `Archived`: Archived stage.
type Stage string

// Stage of the model version. Valid values are: * `None`: The initial stage of
// a model version. * `Staging`: Staging or pre-production stage. *
// `Production`: Production stage. * `Archived`: Archived stage.
const StageNone Stage = `None`
// Stage of the model version. Valid values are: * `None`: The initial stage of
// a model version. * `Staging`: Staging or pre-production stage. *
// `Production`: Production stage. * `Archived`: Archived stage.
const StageStaging Stage = `Staging`
// Stage of the model version. Valid values are: * `None`: The initial stage of
// a model version. * `Staging`: Staging or pre-production stage. *
// `Production`: Production stage. * `Archived`: Archived stage.
const StageProduction Stage = `Production`
// Stage of the model version. Valid values are: * `None`: The initial stage of
// a model version. * `Staging`: Staging or pre-production stage. *
// `Production`: Production stage. * `Archived`: Archived stage.
const StageArchived Stage = `Archived`
// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks. * `FAILED_REGISTRATION`: Request to register a new model
// version has failed. * `READY`: Model version is ready for use.
type Status string

// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks. * `FAILED_REGISTRATION`: Request to register a new model
// version has failed. * `READY`: Model version is ready for use.
const StatusPendingRegistration Status = `PENDING_REGISTRATION`
// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks. * `FAILED_REGISTRATION`: Request to register a new model
// version has failed. * `READY`: Model version is ready for use.
const StatusFailedRegistration Status = `FAILED_REGISTRATION`
// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks. * `FAILED_REGISTRATION`: Request to register a new model
// version has failed. * `READY`: Model version is ready for use.
const StatusReady Status = `READY`
// Array of tags that are associated with the model version.
type VersionTags []ModelVersionTag


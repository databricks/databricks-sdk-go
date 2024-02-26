# Version changelog

## 0.33.0

Internal Changes:

* Add helper function to get header fields ([#822](https://github.com/databricks/databricks-sdk-go/pull/822)).
* Add Int64 to header type injection ([#819](https://github.com/databricks/databricks-sdk-go/pull/819)).

API Changes:

 * Changed `Update` method for [w.LakehouseMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#LakehouseMonitorsAPI) workspace-level service with new required argument order.
 * Added [w.OnlineTables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTablesAPI) workspace-level service.
 * Removed `AssetsDir` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor).
 * Added [catalog.ContinuousUpdateStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ContinuousUpdateStatus).
 * Added [catalog.DeleteOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteOnlineTableRequest).
 * Added [catalog.FailedStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FailedStatus).
 * Added [catalog.GetOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetOnlineTableRequest).
 * Added [catalog.OnlineTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTable).
 * Added [catalog.OnlineTableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTableSpec).
 * Added [catalog.OnlineTableState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTableState).
 * Added [catalog.OnlineTableStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTableStatus).
 * Added [catalog.PipelineProgress](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PipelineProgress).
 * Added [catalog.ProvisioningStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningStatus).
 * Added [catalog.TriggeredUpdateStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TriggeredUpdateStatus).
 * Added [catalog.ViewData](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ViewData).
 * Added `ContentLength` field for [files.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadResponse).
 * Added `ContentType` field for [files.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadResponse).
 * Added `LastModified` field for [files.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadResponse).
 * Changed `LastModified` field for [files.GetMetadataResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#GetMetadataResponse) to [files.LastModifiedHttpDate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#LastModifiedHttpDate).
 * Added [files.LastModifiedHttpDate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#LastModifiedHttpDate).
 * Removed `Config` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `Ai21labsConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `AnthropicConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `AwsBedrockConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `CohereConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `DatabricksModelServingConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `OpenaiConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `PalmConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Removed [serving.ExternalModelConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModelConfig).
 * Added `MaxProvisionedThroughput` field for [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
 * Added `MinProvisionedThroughput` field for [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
 * Added `MaxProvisionedThroughput` field for [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).
 * Added `MinProvisionedThroughput` field for [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).

OpenAPI SHA: cdd76a98a4fca7008572b3a94427566dd286c63b, Date: 2024-02-19

## 0.32.0

Major Changes:

* Generate fields for headers in responses ([#812](https://github.com/databricks/databricks-sdk-go/pull/812)).

Other Changes:

* Add `CurrentWorkspaceID()` method on WorkspaceClient ([#808](https://github.com/databricks/databricks-sdk-go/pull/808)).
* Update OpenAPI spec ([#816](https://github.com/databricks/databricks-sdk-go/pull/816)).

Internal Changes:

* Do not generate erroneous "r" and "w" samples ([#811](https://github.com/databricks/databricks-sdk-go/pull/811)).

API Changes:

 * Changed `Delete` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Changed `Get` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Changed `Update` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Changed `Delete` method for [w.Volumes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumesAPI) workspace-level service with new required argument order.
 * Changed `Read` method for [w.Volumes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumesAPI) workspace-level service with new required argument order.
 * Changed `Update` method for [w.Volumes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumesAPI) workspace-level service with new required argument order.
 * Removed `NameArg` field for [catalog.DeleteConnectionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteConnectionRequest).
 * Added `Name` field for [catalog.DeleteConnectionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteConnectionRequest).
 * Removed `FullNameArg` field for [catalog.DeleteVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteVolumeRequest).
 * Added `Name` field for [catalog.DeleteVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteVolumeRequest).
 * Removed `NameArg` field for [catalog.GetConnectionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetConnectionRequest).
 * Added `Name` field for [catalog.GetConnectionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetConnectionRequest).
 * Added `MaxResults` field for [catalog.ListVolumesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesRequest).
 * Added `PageToken` field for [catalog.ListVolumesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesRequest).
 * Added `NextPageToken` field for [catalog.ListVolumesResponseContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesResponseContent).
 * Removed `FullNameArg` field for [catalog.ReadVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ReadVolumeRequest).
 * Added `Name` field for [catalog.ReadVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ReadVolumeRequest).
 * Removed `NameArg` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Added `Name` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Removed `FullNameArg` field for [catalog.UpdateVolumeRequestContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateVolumeRequestContent).
 * Added `Name` field for [catalog.UpdateVolumeRequestContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateVolumeRequestContent).
 * Removed `GetStatus` method for [w.Files](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FilesAPI) workspace-level service.
 * Added `GetDirectoryMetadata` method for [w.Files](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FilesAPI) workspace-level service.
 * Added `GetMetadata` method for [w.Files](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FilesAPI) workspace-level service.
 * Added [files.FileSize](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FileSize).
 * Added [files.GetDirectoryMetadataRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#GetDirectoryMetadataRequest).
 * Added [files.GetMetadataRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#GetMetadataRequest).
 * Added [files.GetMetadataResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#GetMetadataResponse).
 * Removed `TriggerHistory` field for [jobs.Job](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Job).
 * Removed [jobs.TriggerEvaluation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerEvaluation).
 * Removed [jobs.TriggerHistory](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerHistory).
 * Added `Table` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings).
 * Added [jobs.Condition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Condition).
 * Added [jobs.TableTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TableTriggerConfiguration).
 * Changed `Delete` method for [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service with new required argument order.
 * Changed `Get` method for [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service with new required argument order.
 * Changed `Update` method for [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service with new required argument order.
 * Removed `NameArg` field for [sharing.DeleteCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#DeleteCleanRoomRequest).
 * Added `Name` field for [sharing.DeleteCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#DeleteCleanRoomRequest).
 * Removed `NameArg` field for [sharing.GetCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#GetCleanRoomRequest).
 * Added `Name` field for [sharing.GetCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#GetCleanRoomRequest).
 * Removed `NameArg` field for [sharing.UpdateCleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateCleanRoom).
 * Added `Name` field for [sharing.UpdateCleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateCleanRoom).
 * Added `EnumOptions` field for [sql.Parameter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Parameter).
 * Added `MultiValuesOptions` field for [sql.Parameter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Parameter).
 * Added `QueryId` field for [sql.Parameter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Parameter).
 * Added [sql.MultiValuesOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#MultiValuesOptions).

OpenAPI SHA: c40670f5a2055c92cf0a6aac92a5bccebfb80866, Date: 2024-02-14

## 0.31.0

* Support creating a new workspace client from an account client ([#792](https://github.com/databricks/databricks-sdk-go/pull/792)). Please see the example:
```
// GetWorkspaceClient returns a WorkspaceClient for the given workspace. The
// workspace can be fetched by calling w.Workspaces.Get() or w.Workspaces.List().
//
// The config used for the workspace is identical to that used for the account,
// except that the host is set to the workspace host, and the account ID is
// not set.

a, err := databricks.NewAccountClient()
if err != nil {
  panic(err)
}
ctx := context.Background()
workspaces, err := a.Workspaces.List(ctx)
if err != nil {
  panic(err)
}
w, err := a.GetWorkspaceClient(workspaces[0])
if err != nil {
  panic(err)
}
me, err := w.CurrentUser.Me(ctx)
```
* Added support to select Spark version with Photon ([#799](https://github.com/databricks/databricks-sdk-go/pull/799)). Please Note: Photon selection is disabled by default. To enable it, please use `Photon: true` in request. Example:
````
	version, err := sparkVersions.Select(compute.SparkVersionRequest{
		Photon:          true,
	})
````


API Changes:

Additions:

 * Added the following Requests: 
    - [catalog.CancelRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CancelRefreshRequest).
    - [catalog.GetRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRefreshRequest).
    - [catalog.ListRefreshesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRefreshesRequest).
    - [settings.DeleteRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteRestrictWorkspaceAdminsSettingRequest).
    - [settings.GetDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetDefaultNamespaceSettingRequest).
    - [settings.GetPersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetPersonalComputeSettingRequest).
    - [settings.GetRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetRestrictWorkspaceAdminsSettingRequest).
    - [settings.UpdateDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDefaultNamespaceSettingRequest).
    - [settings.UpdateRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateRestrictWorkspaceAdminsSettingRequest).
    - [catalog.RunRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RunRefreshRequest).
    - [settings.DeleteDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultNamespaceSettingRequest).
    - [files.CreateDirectoryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#CreateDirectoryRequest).
    - [files.DeleteDirectoryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DeleteDirectoryRequest).
    - [files.ListDirectoryContentsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#ListDirectoryContentsRequest).
 * Added the following Responses: 
    - [settings.DeleteRestrictWorkspaceAdminsSettingResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteRestrictWorkspaceAdminsSettingResponse).
    - [settings.DeleteDefaultNamespaceSettingResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultNamespaceSettingResponse).
    - [files.ListDirectoryResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#ListDirectoryResponse).
 * Added `CancelRefresh`, `GetRefresh`, `ListRefreshes` and `RunRefresh` method for [w.LakehouseMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#LakehouseMonitorsAPI) workspace-level service.
 * Added `Abfss` and `Gcs` field for [compute.InitScriptInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfo).
 * Added `CreateDirectory`, `DeleteDirectory` and `ListDirectoryContents` method for [w.Files](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FilesAPI) workspace-level service.
 * Added `Source` field for [jobs.DbtTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DbtTask) and [jobs.SqlTaskFile](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SqlTaskFile).
 * Added [settings.RestrictWorkspaceAdminsMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#RestrictWorkspaceAdminsMessage).
 * Jobs: 
    - [jobs.ForEachStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ForEachStats).
    - [jobs.ForEachTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ForEachTask).
    - [jobs.ForEachTaskErrorMessageStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ForEachTaskErrorMessageStats).
    - [jobs.ForEachTaskTaskRunStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ForEachTaskTaskRunStats).
    - [jobs.RunForEachTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunForEachTask).
 * Pipelines: 
    - [pipelines.PipelineClusterAutoscale](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineClusterAutoscale).
    - [pipelines.PipelineClusterAutoscaleMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineClusterAutoscaleMode).
 * Settings: 
    - [settings.RestrictWorkspaceAdminsMessageStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#RestrictWorkspaceAdminsMessageStatus).
    - [settings.RestrictWorkspaceAdminsSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#RestrictWorkspaceAdminsSetting).
 * Catalog: 
    - [catalog.MonitorRefreshInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfo).
    - [catalog.MonitorRefreshInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfoState).
 * Added `GetPersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service.
 * Added the following fields: 
  - `DeltaSyncIndexSpec` for [vectorsearch.CreateVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateVectorIndexRequest).
  - `FileType` for [workspace.ExportResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ExportResponse).
  - `ResourceId` for [workspace.ObjectInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ObjectInfo).
 * Added `ZoneId` field for [compute.GcpAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GcpAttributes).
 * Added `ForEachTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask), [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask) and [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Added `DeleteDefaultNamespaceSetting`, `DeleteRestrictWorkspaceAdminsSetting`, `GetDefaultNamespaceSetting`, `GetRestrictWorkspaceAdminsSetting`, `UpdateDefaultNamespaceSetting` and `UpdateRestrictWorkspaceAdminsSetting` method for [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
 * Added `FieldMask` field for [settings.UpdatePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdatePersonalComputeSettingRequest).
 * Added `UsePreemptibleExecutors` field for [compute.GcpAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GcpAttributes).
 * Misc:
    - [compute.Adlsgen2Info](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Adlsgen2Info).
    - [compute.GcsStorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GcsStorageInfo).
    - [files.DirectoryEntry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DirectoryEntry).
    - [files.PageToken](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#PageToken).

Changes:
 * Changed `MaxWorkers` and `MinWorkers`  field for [compute.AutoScale](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#AutoScale) to no longer be required.
 * Changed `Destination` field for [compute.LocalFileInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#LocalFileInfo), [compute.S3StorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#S3StorageInfo), [compute.VolumesStorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#VolumesStorageInfo), [compute.WorkspaceStorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#WorkspaceStorageInfo) to be required.
 * Changed `Destination` field for [compute.DbfsStorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#DbfsStorageInfo) to be required.
 * Changed `Clients` field for [compute.WorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#WorkloadType) to be required.
 * Changed `Autoscale` field for [pipelines.PipelineCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineCluster) to [pipelines.PipelineClusterAutoscale](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineClusterAutoscale).
 * Changed `DeletePersonalComputeSetting` and `UpdatePersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service with new required argument order.
 * Changed `AllowMissing` and `Setting` field for [settings.UpdatePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdatePersonalComputeSettingRequest) to be required.
 * Changed `Etag` field for [settings.DeletePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeletePersonalComputeSettingRequest) to no longer be required.

Removals: 
 * Removed `Name` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection), [catalog.UpdateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMetastore), [catalog.UpdateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateRegisteredModelRequest), [catalog.UpdateSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateSchema), [catalog.UpdateVolumeRequestContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateVolumeRequestContent).
 * Settings: 
    - [settings.DeleteDefaultWorkspaceNamespaceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultWorkspaceNamespaceResponse).
    - [settings.DeleteDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultWorkspaceNamespaceRequest).
    - [settings.ReadDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadDefaultWorkspaceNamespaceRequest).
    - [settings.UpdateDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDefaultWorkspaceNamespaceRequest).
    - [settings.ReadPersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadPersonalComputeSettingRequest).
 * Removed `ReadDefaultWorkspaceNamespace`, `UpdateDefaultWorkspaceNamespace` and `DeleteDefaultWorkspaceNamespace` method for [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
 * Removed `Reset` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Removed [pipelines.ResetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ResetRequest).
 * Removed `ReadPersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service.
 * Removed `DeltaSyncVectorIndexSpec` field for [vectorsearch.CreateVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateVectorIndexRequest).

 Internal Changes:
* Added Support for HEAD operations ([#802](https://github.com/databricks/databricks-sdk-go/pull/802)).
* Updated actions/setup-go to v5 ([#784](https://github.com/databricks/databricks-sdk-go/pull/784)).
* Retry update of catalog in test to avoid flakiness ([#788](https://github.com/databricks/databricks-sdk-go/pull/788)).
* Skip AccountClient_GetWorkspaceClient() test in Azure/GCP ([#798](https://github.com/databricks/databricks-sdk-go/pull/798)).
* SDK Generation + Fix backwards incompatible changes ([#806](https://github.com/databricks/databricks-sdk-go/pull/806)).
* Fixed `any` references when generating code with circular dependencies ([#805](https://github.com/databricks/databricks-sdk-go/pull/805)).
* Fixed stack overflow on recursive schemas ([#801](https://github.com/databricks/databricks-sdk-go/pull/801)).

Dependency updates:
 * Bump google.golang.org/api from 0.154.0 to 0.161.0 ([#794](https://github.com/databricks/databricks-sdk-go/pull/794)).

OpenAPI SHA: cadf1693527b365728a55ff06a0e38ce5740c9f7, Date: 2024-02-08


## 0.30.1

Major changes:

* Enable Databricks CLI auth for all clouds ([#783](https://github.com/databricks/databricks-sdk-go/pull/783)).

Other changes:

* Update codecov/codecov-action to v3 ([#785](https://github.com/databricks/databricks-sdk-go/pull/785)).
* MustUseJson must be true if request is a map ([#786](https://github.com/databricks/databricks-sdk-go/pull/786)).


## 0.30.0

Other changes:

* Added RequiredPositionalArguments method for codegen ([#773](https://github.com/databricks/databricks-sdk-go/pull/773)).
* Support locking in integration tests ([#776](https://github.com/databricks/databricks-sdk-go/pull/776)).
* Update OpenAPI spec and fix incompatible changes ([#778](https://github.com/databricks/databricks-sdk-go/pull/778)).

API Changes:

 * Added `Exists` method for [w.Tables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TablesAPI) workspace-level service.
 * Added [w.LakehouseMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#LakehouseMonitorsAPI) workspace-level service.
 * Added [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor), [catalog.DeleteLakehouseMonitorRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteLakehouseMonitorRequest), [catalog.ExistsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExistsRequest), [catalog.GetLakehouseMonitorRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetLakehouseMonitorRequest), [catalog.MonitorCronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCronSchedule), [catalog.MonitorCronSchedulePauseStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCronSchedulePauseStatus), [catalog.MonitorCustomMetric](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCustomMetric), [catalog.MonitorCustomMetricType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCustomMetricType), [catalog.MonitorDataClassificationConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorDataClassificationConfig), [catalog.MonitorDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorDestinations), [catalog.MonitorInferenceLogProfileType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInferenceLogProfileType), [catalog.MonitorInferenceLogProfileTypeProblemType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInferenceLogProfileTypeProblemType), [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo), [catalog.MonitorInfoStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfoStatus), [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig), [catalog.MonitorTimeSeriesProfileType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorTimeSeriesProfileType), [catalog.TableExistsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableExistsResponse) and [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor). 
 * Added `InitScripts` field for [pipelines.PipelineCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineCluster).
 * Added `ValidateOnly` field for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
 * Added `ValidateOnly` field for [pipelines.UpdateInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#UpdateInfo).
 * Changed `CreateOboToken` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service with new required argument order.
 * Changed `Get` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service to return [settings.GetTokenResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetTokenResponse).
 * Changed `LifetimeSeconds` field for [settings.CreateOboTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateOboTokenRequest) to no longer be required.
 * Added [settings.GetTokenResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetTokenResponse).
 * Changed `Create` method for [w.Dashboards](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardsAPI) workspace-level service . New request type is [sql.DashboardPostContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardPostContent).
 * Added `Update` method for [w.Dashboards](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardsAPI) workspace-level service.
 * Removed [sql.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateDashboardRequest).
 * Added `HttpHeaders` field for [sql.ExternalLink](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExternalLink).
 * Added `RunAsRole` field for [sql.QueryEditContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryEditContent).
 * Added [sql.DashboardEditContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardEditContent).
 * Added [sql.DashboardPostContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardPostContent).

OpenAPI SHA: e05401ed5dd4974c5333d737ec308a7d451f749f, Date: 2024-01-23

## 0.29.1

This patch release contains two small changes:
* Retry on Status Code 503 ([#733](https://github.com/databricks/databricks-sdk-go/pull/733)), improving the stability of the SDK in light of transient API unavailability.
* Simplify mocking of iterator and waiter objects ([#769](https://github.com/databricks/databricks-sdk-go/pull/769), [#770](https://github.com/databricks/databricks-sdk-go/pull/770)). See the [Testing section of the README.md](https://github.com/databricks/databricks-sdk-go#testing) for usage information and examples.
## 0.29.0

* Extract API interfaces for all services and generate mock clients ([#740](https://github.com/databricks/databricks-sdk-go/pull/740)).
* Handle empty response for errors ([#756](https://github.com/databricks/databricks-sdk-go/pull/756)).
* Update SDK to OpenAPI spec + Fix tests ([#755](https://github.com/databricks/databricks-sdk-go/pull/755)).
* Add utility to retry on specific errors ([#757](https://github.com/databricks/databricks-sdk-go/pull/757)).
* Integration test fixes for TestMwsAccWorkspaces ([#763](https://github.com/databricks/databricks-sdk-go/pull/763)) and TestMwsAccUsageDownload ([#764](https://github.com/databricks/databricks-sdk-go/pull/764)).

Note: This release contains breaking changes, please see below.

API Changes:

 * [Breaking] Changed `List` method for [w.ExternalLocations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationsAPI) workspace-level service to require request of [catalog.ListExternalLocationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsRequest), [w.StorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialsAPI) workspace-level service to require request of [catalog.ListStorageCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListStorageCredentialsRequest) and [w.Tokens](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokensAPI) workspace-level service to return [settings.ListPublicTokensResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListPublicTokensResponse).
 * Added `NextPageToken` field for [catalog.ListExternalLocationsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsResponse), [catalog.ListFunctionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsResponse), [catalog.ListSchemasResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasResponse) and [catalog.ListStorageCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListStorageCredentialsResponse).
 * Added `MaxResults` field for [catalog.ListFunctionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsRequest) and [catalog.ListSchemasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasRequest).
 * Added `PageToken` field for [catalog.ListFunctionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsRequest) and [catalog.ListSchemasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasRequest).
 * Added `OmitColumns` and `OmitProperties` field for [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest).
 * Added [catalog.ListExternalLocationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsRequest).
 * Added [catalog.ListStorageCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListStorageCredentialsRequest).
 * Added [settings.ListPublicTokensResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListPublicTokensResponse).
 * Added [dashboards](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards) package.
 * Added [vectorsearch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch) package.

OpenAPI SHA: a7a9dc025bb80303e676bf3708942c6aa06689f1, Date: 2024-01-04
Dependency updates:

 * Bump google.golang.org/api from 0.153.0 to 0.154.0 ([#741](https://github.com/databricks/databricks-sdk-go/pull/741)).
 * Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /examples/slog ([#747](https://github.com/databricks/databricks-sdk-go/pull/747)) and /examples/zerolog ([#748](https://github.com/databricks/databricks-sdk-go/pull/748)).
 * Bump golang.org/x/crypto from 0.16.0 to 0.17.0 ([#749](https://github.com/databricks/databricks-sdk-go/pull/749)).

## 0.28.1

This is a bugfix release that improves debuggability of unexpected errors ([#744](https://github.com/databricks/databricks-sdk-go/pull/744), [#750](https://github.com/databricks/databricks-sdk-go/pull/750)). When the SDK cannot deserialize a response from the Databricks REST API, the resulting error will include debugging information and instructions on how to submit a bug report to the SDK.


## 0.28.0

* Consolidate usage of named sorting ([#736](https://github.com/databricks/databricks-sdk-go/pull/736)).
* Remove redundant retries for token refreshes as they're already handled in `httpclient` ([#729](https://github.com/databricks/databricks-sdk-go/pull/729)).
* Skip auth, config loading, and rate limits when fixture transport is present in the config ([#739](https://github.com/databricks/databricks-sdk-go/pull/739)).
* Generate SDK to latest OpenAPI specification with bugfix ([#742](https://github.com/databricks/databricks-sdk-go/pull/742)).

API Changes:

 * Added `AzureWorkspaceInfo` field for [provisioning.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Workspace).
 * Added [provisioning.AzureWorkspaceInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#AzureWorkspaceInfo).
 * Changed `UpdateConfig` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service with new required argument order.
 * Changed `ServedEntities` field for [serving.EndpointCoreConfigInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigInput) to no longer be required.
 * Changed `Create` method for [a.AccountIpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountIpAccessListsAPI) account-level service with new required argument order.
 * Changed `Replace` method for [a.AccountIpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountIpAccessListsAPI) account-level service with new required argument order.
 * Changed `Update` method for [a.AccountIpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountIpAccessListsAPI) account-level service with new required argument order.
 * Changed `Create` method for [w.IpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#IpAccessListsAPI) workspace-level service with new required argument order.
 * Changed `Replace` method for [w.IpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#IpAccessListsAPI) workspace-level service with new required argument order.
 * Changed `Update` method for [w.IpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#IpAccessListsAPI) workspace-level service with new required argument order.
 * Changed `IpAddresses` field for [settings.CreateIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateIpAccessList) to no longer be required.
 * Changed `IpAddresses` field for [settings.ReplaceIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReplaceIpAccessList) to no longer be required.
 * Removed `ListId` field for [settings.ReplaceIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReplaceIpAccessList).
 * Changed `Enabled` field for [settings.UpdateIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateIpAccessList) to no longer be required.
 * Changed `IpAddresses` field for [settings.UpdateIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateIpAccessList) to no longer be required.
 * Changed `Label` field for [settings.UpdateIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateIpAccessList) to no longer be required.
 * Removed `ListId` field for [settings.UpdateIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateIpAccessList).
 * Changed `ListType` field for [settings.UpdateIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateIpAccessList) to no longer be required.

OpenAPI SHA: d3853c8dee5806d04da2ae8910f273ffb35719a5, Date: 2023-12-14

## 0.27.0


API Changes:

 * Changed `Update` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Added `CloudflareApiToken` field for [catalog.CreateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateStorageCredential).
 * Added `CloudflareApiToken` field for [catalog.StorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfo).
 * Changed `Name` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog) to be required.
 * Added `NewName` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog).
 * Changed `Name` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection) to no longer be required.
 * Added `NewName` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Changed `Name` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation) to be required.
 * Added `NewName` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
 * Added `NewName` field for [catalog.UpdateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMetastore).
 * Added `NewName` field for [catalog.UpdateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateRegisteredModelRequest).
 * Added `NewName` field for [catalog.UpdateSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateSchema).
 * Changed `Name` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential) to be required.
 * Added `CloudflareApiToken` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential).
 * Added `NewName` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential).
 * Added `NewName` field for [catalog.UpdateVolumeRequestContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateVolumeRequestContent).
 * Added `CloudflareApiToken` field for [catalog.ValidateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateStorageCredential).
 * Added [catalog.CloudflareApiToken](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CloudflareApiToken).
 * Removed `Continuous` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
 * Removed `Continuous` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
 * Changed `JobParameters` field for [jobs.RunJobTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunJobTask) to [jobs.ParamPairs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ParamPairs).
 * Added `RunIf` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `RunJobTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Changed `UpdateConfig` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service with new required argument order.
 * Added `Put` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `RateLimits` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
 * Changed `ServedModels` field for [serving.EndpointCoreConfigInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigInput) to no longer be required.
 * Added `AutoCaptureConfig` field for [serving.EndpointCoreConfigInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigInput).
 * Added `ServedEntities` field for [serving.EndpointCoreConfigInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigInput).
 * Added `AutoCaptureConfig` field for [serving.EndpointCoreConfigOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigOutput).
 * Added `ServedEntities` field for [serving.EndpointCoreConfigOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigOutput).
 * Added `ServedEntities` field for [serving.EndpointCoreConfigSummary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigSummary).
 * Added `ServedEntities` field for [serving.EndpointPendingConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointPendingConfig).
 * Added `ExtraParams` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `Input` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `MaxTokens` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `Messages` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `N` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `Prompt` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `Stop` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `Stream` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `Temperature` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Changed `Predictions` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse) to no longer be required.
 * Added `Choices` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Added `Created` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Added `Data` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Added `Id` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Added `Model` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Added `Object` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Added `Usage` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Changed `WorkloadSize` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput) to [serving.ServedModelInputWorkloadSize](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadSize).
 * Changed `WorkloadType` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput) to [serving.ServedModelInputWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadType).
 * Added `Task` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
 * Added `Task` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
 * Added [serving.Ai21LabsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#Ai21LabsConfig).
 * Added [serving.AnthropicConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AnthropicConfig).
 * Added [serving.AutoCaptureConfigInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AutoCaptureConfigInput).
 * Added [serving.AutoCaptureConfigOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AutoCaptureConfigOutput).
 * Added [serving.AutoCaptureState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AutoCaptureState).
 * Added [serving.AwsBedrockConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AwsBedrockConfig).
 * Added [serving.AwsBedrockConfigBedrockProvider](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AwsBedrockConfigBedrockProvider).
 * Added [serving.ChatMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ChatMessage).
 * Added [serving.ChatMessageRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ChatMessageRole).
 * Added [serving.CohereConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CohereConfig).
 * Added [serving.DatabricksModelServingConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DatabricksModelServingConfig).
 * Added [serving.EmbeddingsV1ResponseEmbeddingElement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EmbeddingsV1ResponseEmbeddingElement).
 * Added [serving.EmbeddingsV1ResponseEmbeddingElementObject](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EmbeddingsV1ResponseEmbeddingElementObject).
 * Added [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added [serving.ExternalModelConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModelConfig).
 * Added [serving.ExternalModelProvider](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModelProvider).
 * Added [serving.ExternalModelUsageElement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModelUsageElement).
 * Added [serving.FoundationModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#FoundationModel).
 * Added [serving.OpenAiConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#OpenAiConfig).
 * Added [serving.PaLmConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PaLmConfig).
 * Added [serving.PayloadTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PayloadTable).
 * Added [serving.PutRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PutRequest).
 * Added [serving.PutResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PutResponse).
 * Added [serving.QueryEndpointResponseObject](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponseObject).
 * Added [serving.RateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#RateLimit).
 * Added [serving.RateLimitKey](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#RateLimitKey).
 * Added [serving.RateLimitRenewalPeriod](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#RateLimitRenewalPeriod).
 * Added [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
 * Added [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).
 * Added [serving.ServedEntitySpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntitySpec).
 * Added [serving.ServedModelInputWorkloadSize](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadSize).
 * Added [serving.ServedModelInputWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadType).
 * Added [serving.V1ResponseChoiceElement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#V1ResponseChoiceElement).
 * Removed [a.AccountNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyAPI) account-level service.
 * Removed [settings.AccountNetworkPolicyMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyMessage).
 * Removed [settings.DeleteAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyRequest).
 * Removed [settings.DeleteAccountNetworkPolicyResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyResponse).
 * Removed [settings.ReadAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadAccountNetworkPolicyRequest).
 * Removed [settings.UpdateAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateAccountNetworkPolicyRequest).
 * Removed `Name` field for [sharing.UpdateCleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateCleanRoom).
 * Changed `Name` field for [sharing.UpdateProvider](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateProvider) to be required.
 * Added `NewName` field for [sharing.UpdateProvider](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateProvider).
 * Changed `Name` field for [sharing.UpdateRecipient](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateRecipient) to be required.
 * Added `NewName` field for [sharing.UpdateRecipient](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateRecipient).
 * Changed `Name` field for [sharing.UpdateShare](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateShare) to be required.
 * Added `NewName` field for [sharing.UpdateShare](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateShare).
 * Added `StatementIds` field for [sql.QueryFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryFilter).
 * Added [sql.StatementId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementId).

OpenAPI SHA: 63caa3cb0c05045e81d3dcf2451fa990d8670f36, Date: 2023-12-12
Dependency updates:

 * Bumped google.golang.org/api from 0.152.0 to 0.153.0 ([#731](https://github.com/databricks/databricks-sdk-go/pull/731)).

## 0.26.2

This is a bugfix release, including a fix correcting issues with OAuth flows, due to a bug with the propagation of the response status in `httpclient`'s `RoundTrip()` implementation. This fixes the `failed during request visitor: token: oauth2: cannot fetch token: Response: {...}` error.

All fixes:

* Migrate Azure MSI & Metadata Service token sources to `httpclient` and add 100% test coverage ([#709](https://github.com/databricks/databricks-sdk-go/pull/709)).
* Added `config.NewAzureCliTokenSource` and `config.NewAzureMsiTokenSource` constructors ([#727](https://github.com/databricks/databricks-sdk-go/pull/727)).
* Use per-config refresh context for OAuth tokens ([#728](https://github.com/databricks/databricks-sdk-go/pull/728)).


## 0.26.1

Minor changes:
* Support overriding DatabricksEnvironment ([#723](https://github.com/databricks/databricks-sdk-go/pull/723)).
* Detect `Accept` header in `httpclient.WithResponseUnmarshal` ([#710](https://github.com/databricks/databricks-sdk-go/pull/710)).
* Detect `Content-Type` header in `newRequestBody` for `httpclient` ([#711](https://github.com/databricks/databricks-sdk-go/pull/711)).

Bug fixes:
* Retry request on `REQUEST_LIMIT_EXCEEDED` error returned by the SCIM API ([#721](https://github.com/databricks/databricks-sdk-go/pull/721)).
* Match retry logic of pre-refactor SDK ([#722](https://github.com/databricks/databricks-sdk-go/pull/722)).


## 0.26.0

Major changes:

* There has been a major overhaul of error handling. Users can now compare errors in API responses to the well-known error responses defined in the `apierr` package and reexported in the `databricks` package. Users can check whether a specific error was returned, for example `errors.Is(err, databricks.ErrResourceAlreadyExists)`, rather than converting the error to `*APIError` to check the status code and error code. This change is backwards-compatible; users do not need to modify existing error-handling code when upgrading the SDK. See [#682](https://github.com/databricks/databricks-sdk-go/pull/682) and [#703](https://github.com/databricks/databricks-sdk-go/pull/703) for the changes and https://github.com/databricks/databricks-sdk-go/blob/main/error_alias.go for the full set of errors.

Bug fixes:

* Handle "no configuration file found at" error during databricks-cli authentication ([#707](https://github.com/databricks/databricks-sdk-go/pull/707)).
* Introduce `DatabricksEnvironment` and fix Azure MSI auth from ACR, where IMDS doesn't give host environment information ([#700](https://github.com/databricks/databricks-sdk-go/pull/700)).
* Fix SCIM Pagination default parameters in the Go SDK ([#717](https://github.com/databricks/databricks-sdk-go/pull/717)).

Other changes:

* Update `slog` example with the correct interface ([#694](https://github.com/databricks/databricks-sdk-go/pull/694)).
* Fixed typo in error message for unknown azure environment ([#701](https://github.com/databricks/databricks-sdk-go/pull/701)).
* Allow injection of HTTP transport to enable HTTP replayer pattern ([#697](https://github.com/databricks/databricks-sdk-go/pull/697)).
* Decouple HTTP retries and error mapping mechanics from `DatabricksClient` into `httpclient.ApiClient` ([#699](https://github.com/databricks/databricks-sdk-go/pull/699), [#702](https://github.com/databricks/databricks-sdk-go/pull/702), [#712](https://github.com/databricks/databricks-sdk-go/pull/712)).
* Port `qa.HTTPFixtures` to faster transport-level stubs ([#708](https://github.com/databricks/databricks-sdk-go/pull/708)).

API Changes:

 * Removed `EnableOptimization` method for [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service.
 * Added `PipelineId` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
 * Added `EnablePredictiveOptimization` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog) and [catalog.UpdateSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateSchema).
 * Removed [catalog.UpdatePredictiveOptimization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePredictiveOptimization) and [catalog.UpdatePredictiveOptimizationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePredictiveOptimizationResponse).
 * Added `Description` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob) and [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `ListNetworkConnectivityConfigurations` and `ListPrivateEndpointRules` method for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service.
 * Added [settings.ListNccAzurePrivateEndpointRulesResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListNccAzurePrivateEndpointRulesResponse), [settings.ListNetworkConnectivityConfigurationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListNetworkConnectivityConfigurationsRequest), [settings.ListNetworkConnectivityConfigurationsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListNetworkConnectivityConfigurationsResponse), and [settings.ListPrivateEndpointRulesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListPrivateEndpointRulesRequest).
 * Added `StringSharedAs` field for [sharing.SharedDataObject](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharedDataObject).

Internal changes:

* Added `contains` method in OpenAPI Generator ([#690](https://github.com/databricks/databricks-sdk-go/pull/690)).
* Skip recipients tests in Azure ([#692](https://github.com/databricks/databricks-sdk-go/pull/692)).
* Allow Files API tests to run in UC environments ([#695](https://github.com/databricks/databricks-sdk-go/pull/695)).
* More cleanup in Unity Catalog integration test ([#719](https://github.com/databricks/databricks-sdk-go/pull/719)).

OpenAPI SHA: 22f09783eb8a84d52026f856be3b2068f9498db3, Date: 2023-11-23
Dependency updates:

 * Bump golang.org/x/oauth2 from 0.13.0 to 0.14.0 ([#689](https://github.com/databricks/databricks-sdk-go/pull/689)).
 * Bump google.golang.org/api from 0.150.0 to 0.151.0 ([#698](https://github.com/databricks/databricks-sdk-go/pull/698)).
 * Bump the OpenAPI Spec ([#706](https://github.com/databricks/databricks-sdk-go/pull/706)).
 * Bump golang.org/x/oauth2 from 0.14.0 to 0.15.0 ([#715](https://github.com/databricks/databricks-sdk-go/pull/715)).
 * Bump golang.org/x/time from 0.4.0 to 0.5.0 ([#714](https://github.com/databricks/databricks-sdk-go/pull/714)).
 * Bump google.golang.org/api from 0.151.0 to 0.152.0 ([#716](https://github.com/databricks/databricks-sdk-go/pull/716)).


## 0.25.0

* Make sure path parameters are first in order in RequiredFields ([#669](https://github.com/databricks/databricks-sdk-go/pull/669)).
* Added Field.IsRequestBodyField method for code generation ([#670](https://github.com/databricks/databricks-sdk-go/pull/670)).
* Added regressions question to the issue template ([#676](https://github.com/databricks/databricks-sdk-go/pull/676)).
* Added telemetry for CI/CD platform to useragent ([#665](https://github.com/databricks/databricks-sdk-go/pull/665)).
* Skiped GCP Integration Tests using Statement Execution API ([#678](https://github.com/databricks/databricks-sdk-go/pull/678)).
* Added more detailed error message on default credentials not found error ([#679](https://github.com/databricks/databricks-sdk-go/pull/679)).
* Updated SDK to latest OpenAPI Spec ([#685](https://github.com/databricks/databricks-sdk-go/pull/685)).

API Changes:

 * Changed `Create` method for [w.Functions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionsAPI) and [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service with new required argument order.
 * Changed `InputParams` field for [catalog.CreateFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateFunction) and [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo) to [catalog.FunctionParameterInfos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionParameterInfos).
 * Changed `Properties` field for [catalog.CreateFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateFunction) and [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo) to `string`.
 * Changed `ReturnParams` field for [catalog.CreateFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateFunction) and [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo) to [catalog.FunctionParameterInfos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionParameterInfos)
 * Changed `StorageRoot` field for [catalog.CreateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMetastore) to no longer be required.
 * Added `SkipValidation` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
 * Added `Libraries` field for [compute.CreatePolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreatePolicy), [compute.EditPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditPolicy) and [compute.Policy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Policy).
 * Added [compute.InitScriptEventDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptEventDetails).
 * Added `InitScripts` field for [compute.EventDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventDetails).
 * Added `File` field for [compute.InitScriptInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfo).
 * Added `ZoneId` field for [compute.InstancePoolGcpAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolGcpAttributes).
 * Added `IncludeResolvedValues` field for [jobs.GetRunRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetRunRequest).
 * Added `EditMode` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob) and [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `NetworkConnectivityConfigId` field for [provisioning.UpdateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#UpdateWorkspaceRequest).
 * Added `ContainerLogs` and `ExtraInfo` field for [serving.DeploymentStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DeploymentStatus).
 * Added [catalog.CreateFunctionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateFunctionRequest), [catalog.DependencyList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DependencyList) and [catalog.FunctionParameterInfos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionParameterInfos).
 * Added [compute.InitScriptExecutionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptExecutionDetails), [compute.InitScriptExecutionDetailsStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptExecutionDetailsStatus), [compute.InitScriptInfoAndExecutionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfoAndExecutionDetails), [compute.LocalFileInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#LocalFileInfo).
 * Added [jobs.CreateJobEditMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJobEditMode) and [jobs.JobSettingsEditMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettingsEditMode).
 * Added `DeleteApp`, `GetApp`, `GetAppDeploymentStatus`, `GetApps` and `GetEvents` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added [serving.AppEvents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppEvents), [serving.AppServiceStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppServiceStatus), [serving.DeleteAppResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DeleteAppResponse), [serving.GetAppDeploymentStatusRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetAppDeploymentStatusRequest), [serving.GetAppResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetAppResponse), [serving.GetEventsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetEventsRequest), [serving.ListAppEventsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppEventsResponse) and [serving.ListAppsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppsResponse).
 * Added [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service.
 * Added [settings.CreateNetworkConnectivityConfigRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateNetworkConnectivityConfigRequest), [settings.CreatePrivateEndpointRuleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRuleRequest), [settings.CreatePrivateEndpointRuleRequestGroupId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRuleRequestGroupId), [settings.DeleteNetworkConnectivityConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteNetworkConnectivityConfigurationRequest), [settings.DeletePrivateEndpointRuleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeletePrivateEndpointRuleRequest), [settings.GetNetworkConnectivityConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetNetworkConnectivityConfigurationRequest), [settings.GetPrivateEndpointRuleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetPrivateEndpointRuleRequest), [settings.NccAzurePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRule), [settings.NccAzurePrivateEndpointRuleConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRuleConnectionState), [settings.NccAzurePrivateEndpointRuleGroupId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRuleGroupId), [settings.NccAzureServiceEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzureServiceEndpointRule), [settings.NccEgressConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccEgressConfig), [settings.NccEgressDefaultRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccEgressDefaultRules), [settings.NccEgressTargetRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccEgressTargetRules), [settings.NetworkConnectivityConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityConfiguration).
 * Removed `Delete`, `Get`,  method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Removed [jobs.JobSettingsUiState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettingsUiState) and [jobs.CreateJobUiState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJobUiState).
 * Removed [a.OAuthEnrollment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#OAuthEnrollmentAPI) account-level service.
 * Removed [oauth2.CreateOAuthEnrollment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateOAuthEnrollment) and [oauth2.OAuthEnrollmentStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#OAuthEnrollmentStatus).
 * Removed `UiState` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob) and [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).

OpenAPI SHA: e7b127cb07af8dd4d8c61c7cc045c8910cdbb02a, Date: 2023-11-08
Dependency updates:

 * Bump google.golang.org/api from 0.146.0 to 0.150.0 ([#683](https://github.com/databricks/databricks-sdk-go/pull/683)).
 * Bump golang.org/x/mod from 0.13.0 to 0.14.0 ([#681](https://github.com/databricks/databricks-sdk-go/pull/681)).
 * Bump google.golang.org/grpc from 1.58.2 to 1.58.3 in /examples/slog ([#672](https://github.com/databricks/databricks-sdk-go/pull/672)).
 * Bump google.golang.org/grpc to 1.58.3 in /examples/zerolog ([#684](https://github.com/databricks/databricks-sdk-go/pull/684)).
 * Bump golang.org/x/time from 0.3.0 to 0.4.0 ([#680](https://github.com/databricks/databricks-sdk-go/pull/680)).

## 0.24.0

* Implemented Iterator support for paginated endpoints or endpoints returning a list ([#543](https://github.com/databricks/databricks-sdk-go/pull/543)). The Iterator interface allows users to get the next resource in the iterator and to check whether another resource is available in the iterator. Iterators for paginated endpoints fetch pages lazily, allowing users to only fetch the pages needed for their use case.
* Removed `photon` and `graviton` selectors in `compute.SparkVersionRequ…` ([#622](https://github.com/databricks/databricks-sdk-go/pull/622)). Going forward, photon is determined by the `RuntimeEngine` field in `compute.CreateCluster`, and graviton is chosen depending on the `aws_instance_type` field in `compute.CreateCluster`.

API Changes:

 * Added `Attributes`, `Count`, `ExcludedAttributes`, `Filter`, `SortBy`, `SortOrder`, and `StartIndex` fields to [iam.GetAccountUserRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GetAccountUserRequest) and [iam.GetUserRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GetUserRequest).
 * Added `Schemas` field to [iam.Group](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#Group), [iam.ListGroupsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ListGroupsResponse), [iam.ListServicePrincipalResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ListServicePrincipalResponse), [iam.ListUsersResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ListUsersResponse), [iam.ServicePrincipal](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ServicePrincipal), and [iam.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#User).
 * Added [iam.GetSortOrder](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GetSortOrder).
 * Added [iam.GroupSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GroupSchema).
 * Added [iam.ListResponseSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ListResponseSchema).
 * Added [iam.ServicePrincipalSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ServicePrincipalSchema).
 * Added [iam.UserSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UserSchema).
 * Added `WebhookNotifications` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service and related methods.
 * Added [a.AccountNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyAPI) account-level service and related methods.

Internal SDK Changes:
* Update to actions/checkout@v4 ([#650](https://github.com/databricks/databricks-sdk-go/pull/650)).
* Skip unshallow step in test workflow ([#649](https://github.com/databricks/databricks-sdk-go/pull/649)).
* Add integration tests for `Jobs`: `ListRuns` ([#645](https://github.com/databricks/databricks-sdk-go/pull/645)).
* Only log pkg.Load after checking whether the operation is tagged ([#655](https://github.com/databricks/databricks-sdk-go/pull/655)).
* Prefix library paths with the target directory to construct absolute paths ([#656](https://github.com/databricks/databricks-sdk-go/pull/656)).
* Fixed HasRequiredNonBodyField method ([#660](https://github.com/databricks/databricks-sdk-go/pull/660)).
* Added CanSetRequiredFieldsFromJson method for OpenAPI generator ([#661](https://github.com/databricks/databricks-sdk-go/pull/661)).
* Add integration tests for `ServicePrincipals`: `Patch` ([#662](https://github.com/databricks/databricks-sdk-go/pull/662)).
* Add integration tests for `Users`: `Patch`, `Update` ([#663](https://github.com/databricks/databricks-sdk-go/pull/663)).
* Enforce running `az login --service-principal` on nightly runs ([#659](https://github.com/databricks/databricks-sdk-go/pull/659)).
* Add integration tests for `Connections`: `Create`, `Delete`, `Get`, `List`, `Update` ([#653](https://github.com/databricks/databricks-sdk-go/pull/653)).

OpenAPI SHA: 5903bb39137fd76ac384b2044e425f9c56840e00, Date: 2023-10-23
## 0.23.0

Note: this release includes custom JSON marshalling that captures whether or
not a value is set and should be included in the marshalled output, if its
value is equal to Go's zero value (i.e. `0` for an int field).

* Add ForceSendFields and a custom marshaller ([#615](https://github.com/databricks/databricks-sdk-go/pull/615)).
* Support text/plain response as streaming request body ([#638](https://github.com/databricks/databricks-sdk-go/pull/638)).
* Added log statement if OIDC response is non 200 ([#644](https://github.com/databricks/databricks-sdk-go/pull/644)).

API Changes:

 * Changed `Download` method for [a.BillableUsage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BillableUsageAPI) account-level service to start returning [billing.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#DownloadResponse).
 * Added [billing.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#DownloadResponse).
 * Changed `Delete` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service with new required argument order.
 * Changed `Get` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service with new required argument order.
 * Changed `Update` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service with new required argument order.
 * Added `GetBindings` method for [w.WorkspaceBindings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsAPI) workspace-level service.
 * Added `UpdateBindings` method for [w.WorkspaceBindings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsAPI) workspace-level service.
 * Removed `Name` field for [catalog.AccountsUpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsUpdateStorageCredential).
 * Added `StorageCredentialName` field for [catalog.AccountsUpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsUpdateStorageCredential).
 * Removed `Name` field for [catalog.DeleteAccountStorageCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAccountStorageCredentialRequest).
 * Added `StorageCredentialName` field for [catalog.DeleteAccountStorageCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAccountStorageCredentialRequest).
 * Removed `Name` field for [catalog.GetAccountStorageCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetAccountStorageCredentialRequest).
 * Added `StorageCredentialName` field for [catalog.GetAccountStorageCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetAccountStorageCredentialRequest).
 * Added `Owner` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Added [catalog.GetBindingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetBindingsRequest).
 * Added [catalog.UpdateWorkspaceBindingsParameters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateWorkspaceBindingsParameters).
 * Added [catalog.WorkspaceBinding](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBinding).
 * Added [catalog.WorkspaceBindingBindingType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingBindingType).
 * Added [catalog.WorkspaceBindingsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsResponse).
 * Added `Spec` field for [compute.ClusterDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterDetails).
 * Added `ApplyPolicyDefaultValues` field for [compute.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterSpec).
 * Removed `AwsAttributes` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed `AzureAttributes` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed `DiskSpec` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed `EnableElasticDisk` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed `GcpAttributes` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed `PreloadedDockerImages` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed `PreloadedSparkVersions` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Added `Deployment` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `UiState` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `Deployment` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `UiState` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Removed `ConditionTask` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
 * Added `WebhookNotifications` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Added [jobs.CreateJobUiState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJobUiState).
 * Added [jobs.JobDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobDeployment).
 * Added [jobs.JobDeploymentKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobDeploymentKind).
 * Added [jobs.JobSettingsUiState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettingsUiState).
 * Added `WorkloadType` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
 * Added `WorkloadType` field for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput).
 * Removed [a.AccountNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyAPI) account-level service.
 * Changed `List` method for [w.IpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#IpAccessListsAPI) workspace-level service to return [settings.ListIpAccessListResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListIpAccessListResponse).
 * Removed [settings.AccountNetworkPolicyMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyMessage).
 * Removed [settings.DeleteAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyRequest).
 * Removed [settings.DeleteAccountNetworkPolicyResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyResponse).
 * Removed `IpAccessLists` field for [settings.GetIpAccessListResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetIpAccessListResponse).
 * Added `IpAccessList` field for [settings.GetIpAccessListResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetIpAccessListResponse).
 * Removed [settings.ReadAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadAccountNetworkPolicyRequest).
 * Removed [settings.UpdateAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateAccountNetworkPolicyRequest).
 * Added [settings.ListIpAccessListResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListIpAccessListResponse).

OpenAPI SHA: 493a76554afd3afdd15dc858773d01643f80352a, Date: 2023-10-12

Dependency updates:

 * Bump golang.org/x/mod from 0.12.0 to 0.13.0 ([#639](https://github.com/databricks/databricks-sdk-go/pull/639)).
 * Bump google.golang.org/api from 0.140.0 to 0.146.0 ([#643](https://github.com/databricks/databricks-sdk-go/pull/643)).
 * Bump golang.org/x/net from 0.15.0 to 0.17.0 ([#646](https://github.com/databricks/databricks-sdk-go/pull/646)).
 * Bump golang.org/x/oauth2 from 0.12.0 to 0.13.0 ([#642](https://github.com/databricks/databricks-sdk-go/pull/642)).
 * Bump golang.org/x/net from 0.8.0 to 0.17.0 in /examples/zerolog ([#647](https://github.com/databricks/databricks-sdk-go/pull/647)).
 * Bump golang.org/x/net from 0.8.0 to 0.17.0 in /examples/slog ([#648](https://github.com/databricks/databricks-sdk-go/pull/648)).

## 0.22.0

Breaking API Changes:

 * Changed `List` method for [a.AccountMetastoreAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoreAssignmentsAPI) account-level service to return [catalog.ListAccountMetastoreAssignmentsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListAccountMetastoreAssignmentsResponse).

API Changes:

 * Added [catalog.ListAccountMetastoreAssignmentsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListAccountMetastoreAssignmentsResponse).

Internal SDK Changes:

* Add support for template libraries ([#635](https://github.com/databricks/databricks-sdk-go/pull/635)).

OpenAPI SHA: bcbf6e851e3d82fd910940910dd31c10c059746c, Date: 2023-10-02

## 0.21.0

Breaking API Changes:

 * Changed `ArtifactMatchers` field for [catalog.ArtifactAllowlistInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactAllowlistInfo) to [catalog.ArtifactMatcherList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactMatcherList).
 * Removed `Owner` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Changed `ArtifactMatchers` field for [catalog.SetArtifactAllowlist](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SetArtifactAllowlist) to [catalog.ArtifactMatcherList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactMatcherList).
 * Removed `Options` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog).
 * Changed `CancelAllRuns` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service with new required argument order.
 * Changed `JobId` field for [jobs.CancelAllRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CancelAllRuns) to no longer be required.
 * Changed `JobParameters` field for [jobs.RunNow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunNow) to [jobs.ParamPairs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ParamPairs).
 * Changed `Query` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service . New request type is [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Removed [serving.QueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryRequest).
 * Changed `ExchangeToken` method for [w.CredentialsManager](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CredentialsManagerAPI) workspace-level service with new required argument order.
 * Removed [settings.CredentialPartitionId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CredentialPartitionId).
 * Changed `TokenType` field for [settings.ExchangeToken](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeToken) to [settings.TokenType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenType).
 * Removed `CredentialPartitionId` field for [settings.ExchangeTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeTokenRequest).
 * Changed `List` method for [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service to require request of [sharing.ListCleanRoomsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListCleanRoomsRequest).

API Changes:

 * Added `AllQueuedRuns` field for [jobs.CancelAllRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CancelAllRuns).
 * Added `Queue` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `Queue` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `JobParameters` field for [jobs.RepairRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RepairRun).
 * Added `Queue` field for [jobs.RunNow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunNow).
 * Added `JobParameters` field for [jobs.RunParameters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunParameters).
 * Added `QueueReason` field for [jobs.RunState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunState).
 * Added `QueueDuration` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
 * Added `Queue` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added [jobs.QueueSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#QueueSettings).
 * Added `Notifications` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
 * Added `Notifications` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
 * Added `Notifications` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
 * Added [pipelines.Notifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#Notifications).
 * Added [serving.DataframeSplitInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DataframeSplitInput).
 * Added [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
 * Added `PartitionId` field for [settings.ExchangeTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeTokenRequest).
 * Added [settings.DefaultNamespaceSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DefaultNamespaceSetting).
 * Added [settings.DeleteDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultWorkspaceNamespaceRequest).
 * Added [settings.DeleteDefaultWorkspaceNamespaceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultWorkspaceNamespaceResponse).
 * Added [settings.PartitionId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#PartitionId).
 * Added [settings.ReadDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadDefaultWorkspaceNamespaceRequest).
 * Added [settings.StringMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#StringMessage).
 * Added [settings.UpdateDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDefaultWorkspaceNamespaceRequest).
 * Added `NextPageToken` field for [sharing.ListCleanRoomsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListCleanRoomsResponse).
 * Added [sharing.ListCleanRoomsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListCleanRoomsRequest).

SDK Internal Changes:

* Remove use of cancel by job id in integration test ([#632](https://github.com/databricks/databricks-sdk-go/pull/632)).


OpenAPI SHA: 94ddf8ff02df271daebdc5f001075e1ca2ce080d, Date: 2023-09-27

## 0.20.0

* Adjust token expiry window to 40 seconds on Azure ([#617](https://github.com/databricks/databricks-sdk-go/pull/617)).
* Rename variables in Azure code to match across auth types ([#618](https://github.com/databricks/databricks-sdk-go/pull/618)).
* Add details to APIError ([#613](https://github.com/databricks/databricks-sdk-go/pull/613)).
* Fixed code generation of primitive types ([#623](https://github.com/databricks/databricks-sdk-go/pull/623)).
* Updated SDK to changes in OpenAPI specification ([#625](https://github.com/databricks/databricks-sdk-go/pull/625)).

API Changes:

 * Changed `List` method for [a.AccountMetastoreAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoreAssignmentsAPI) account-level service to return [catalog.WorkspaceIdList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceIdList).
 * Added [catalog.WorkspaceId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceId).
 * Added [a.OAuthPublishedApps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#OAuthPublishedAppsAPI) account-level service.
 * Added [oauth2.GetPublishedAppsOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#GetPublishedAppsOutput).
 * Added [oauth2.ListOAuthPublishedAppsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListOAuthPublishedAppsRequest).
 * Added [oauth2.PublishedAppOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#PublishedAppOutput).
 * Added `Patch` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `Tags` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
 * Added `Tags` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
 * Added `Tags` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
 * Added [serving.EndpointTag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointTag).
 * Added [serving.PatchServingEndpointTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PatchServingEndpointTags).
 * Added [w.CredentialsManager](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CredentialsManagerAPI) workspace-level service.
 * Added [settings.CredentialPartitionId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CredentialPartitionId).
 * Added [settings.ExchangeToken](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeToken).
 * Added [settings.ExchangeTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeTokenRequest).
 * Added [settings.ExchangeTokenResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeTokenResponse).
 * Added [settings.TokenType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenType).
 * Changed `ExecuteStatement` method for [w.StatementExecution](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementExecutionAPI) workspace-level service with new required argument order.
 * Added `EmptyResultState` field for [sql.AlertOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertOptions).
 * Removed [sql.ChunkInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ChunkInfo).
 * Changed `OnWaitTimeout` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest) to [sql.ExecuteStatementRequestOnWaitTimeout](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequestOnWaitTimeout).
 * Changed `Statement` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest) to be required.
 * Changed `WarehouseId` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest) to be required.
 * Changed `Chunks` field for [sql.ResultManifest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ResultManifest) to [sql.BaseChunkInfoList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#BaseChunkInfoList).
 * Added `Truncated` field for [sql.ResultManifest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ResultManifest).
 * Removed [sql.TimeoutAction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TimeoutAction).
 * Added [sql.AlertOptionsEmptyResultState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertOptionsEmptyResultState).
 * Added [sql.BaseChunkInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#BaseChunkInfo).
 * Added [sql.ExecuteStatementRequestOnWaitTimeout](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequestOnWaitTimeout).

OpenAPI SHA: 51e3d27c0553c891bc418bd0cd07c9240e0476dd, Date: 2023-09-13
Dependency updates:

 * Bump golang.org/x/oauth2 from 0.11.0 to 0.12.0 ([#606](https://github.com/databricks/databricks-sdk-go/pull/606)).
 * Bump google.golang.org/api from 0.138.0 to 0.140.0 ([#612](https://github.com/databricks/databricks-sdk-go/pull/612)).

## 0.19.2

* Fixed case where retry of a request without body fails ([#614](https://github.com/databricks/databricks-sdk-go/pull/614)).


## 0.19.1

* Handled Azure authentication when WorkspaceResourceID is provided ([#597](https://github.com/databricks/databricks-sdk-go/pull/597)).
* Reverted error message changes from client.go ([#604](https://github.com/databricks/databricks-sdk-go/pull/604)).
* Reused tokens after first call to Azure CLI ([#605](https://github.com/databricks/databricks-sdk-go/pull/605)).
* Updated error message assertions ([#607](https://github.com/databricks/databricks-sdk-go/pull/607)).


## 0.19.0

* Added `ErrNotWorkspaceClient` ([#596](https://github.com/databricks/databricks-sdk-go/pull/596)).
* Fix loading of databrickscfg with a password containing a hash ([#595](https://github.com/databricks/databricks-sdk-go/pull/595)).
* Use an operation's request type name if specified ([#598](https://github.com/databricks/databricks-sdk-go/pull/598)).
* Update OpenAPI spec ([#600](https://github.com/databricks/databricks-sdk-go/pull/600)).

API Changes:

* Renamed permissions APIs to no longer include the service name, for example:
  * `GetJobPermissionLevels` -> `GetPermissionLevels`
  * `GetJobPermissions` -> `GetPermissions`
  * `SetJobPermissions` -> `SetPermissions`
  * `UpdateJobPermissions` -> `UpdatePermissions`
* Changed `Create` method for [w.Volumes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumesAPI) workspace-level service with new required argument order.
* Added `SupportsElasticDisk` field for [compute.NodeType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#NodeType).
* Changed `Create` method for [w.Dashboards](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardsAPI) workspace-level service with new required argument order.
* Added [w.DashboardWidgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardWidgetsAPI) workspace-level service.
* Added [w.QueryVisualizations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryVisualizationsAPI) workspace-level service.
* Changed `Name` field for [sql.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateDashboardRequest) to be required.
* Added `DashboardFiltersEnabled` field for [sql.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateDashboardRequest).
* Added `RunAsRole` field for [sql.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateDashboardRequest).
* Added `RunAsRole` field for [sql.Query](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Query).
* Added `RunAsRole` field for [sql.QueryPostContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryPostContent).
* Removed `DashboardId` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions).
* Changed `Position` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions) to [sql.WidgetPosition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetPosition).
* Removed `Text` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions).
* Added `Description` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions).
* Added `Title` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions).
* Added [sql.CreateQueryVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateQueryVisualizationRequest).
* Added [sql.CreateWidget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateWidget).
* Added [sql.DeleteDashboardWidgetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteDashboardWidgetRequest).
* Added [sql.DeleteQueryVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteQueryVisualizationRequest).
* Added [sql.RunAsRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#RunAsRole).
* Added [sql.WidgetPosition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetPosition).

OpenAPI SHA: 09a7fa63d9ae243e5407941f200960ca14d48b07, Date: 2023-09-04

Dependency updates:

* Bump google.golang.org/api from 0.136.0 to 0.138.0 ([#587](https://github.com/databricks/databricks-sdk-go/pull/587)).
## 0.18.0

Breaking Changes:
* Added support Files API (application/octet-stream) in OpenAPI Spec ([#572](https://github.com/databricks/databricks-sdk-go/pull/572)). The signatures of `Upload`, `Download` and `Delete` have changed; these methods now take `UploadRequest`, `DownloadRequest`, and `DeleteRequest` structures, respectively. Shortcut methods are generated for `DownloadByFileName` and `DeleteByFileName` for convenience. The `WriteFile` and `ReadFile` methods are removed.
* Propagated Request Headers to client.go ([#589](https://github.com/databricks/databricks-sdk-go/pull/589)). The `Do` method of `Client` now requires an additional parameter for request headers. Add headers to requests using this parameter, or pass `nil` if no headers are needed.

Breaking API Changes:
 * Removed [w.SecurableTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableTagsAPI) workspace-level service and all associated structures.
 * Removed [w.SubentityTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SubentityTagsAPI) workspace-level service and all associated structures.
 * Renamed `ProvisioningState` field to `ProvisioningInfo` for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Changed [catalog.ProvisioningState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningState) to [catalog.ProvisioningInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningInfoState).
 * Removed `InstancePoolFleetAttributes` field for [compute.CreateInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateInstancePool).
 * Removed `InstancePoolFleetAttributes` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed [compute.FleetLaunchTemplateOverride](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetLaunchTemplateOverride).
 * Removed [compute.FleetOnDemandOption](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetOnDemandOption).
 * Removed [compute.FleetOnDemandOptionAllocationStrategy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetOnDemandOptionAllocationStrategy).
 * Removed [compute.FleetSpotOption](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetSpotOption).
 * Removed [compute.FleetSpotOptionAllocationStrategy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetSpotOptionAllocationStrategy).
 * Removed `InstancePoolFleetAttributes` field for [compute.GetInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePool).
 * Removed `InstancePoolFleetAttributes` field for [compute.InstancePoolAndStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAndStats).
 * Removed [compute.InstancePoolFleetAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolFleetAttributes).
 * Changed `GetByName` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service to return [ml.GetExperimentResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentResponse).
 * Changed `GetExperiment` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service to return [ml.GetExperimentResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentResponse).
 * Renamed [ml.GetExperimentByNameResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentByNameResponse) to [ml.GetExperimentResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentResponse).

API Changes:
 * Changed `List` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.StorageCredentialInfoList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfoList).
 * Added [w.ModelVersions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionsAPI) workspace-level service.
 * Added [w.RegisteredModels](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelsAPI) workspace-level service.
 * Added `BrowseOnly` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `FullName` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `ProvisioningInfo` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `SecurableKind` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `SecurableType` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `Options` field for [catalog.CreateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCatalog).
 * Added `Options` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog).
 * Added [catalog.CreateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateRegisteredModelRequest).
 * Added [catalog.DeleteAliasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAliasRequest).
 * Added [catalog.DeleteModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteModelVersionRequest).
 * Added [catalog.DeleteRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteRegisteredModelRequest).
 * Added [catalog.GetByAliasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetByAliasRequest).
 * Added [catalog.GetModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelVersionRequest).
 * Added [catalog.GetRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRegisteredModelRequest).
 * Added [catalog.ListModelVersionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsRequest).
 * Added [catalog.ListModelVersionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsResponse).
 * Added [catalog.ListRegisteredModelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsRequest).
 * Added [catalog.ListRegisteredModelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsResponse).
 * Added [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo).
 * Added [catalog.ModelVersionInfoStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfoStatus).
 * Added [catalog.ProvisioningInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningInfo).
 * Added [catalog.ProvisioningInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningInfoState).
 * Added [catalog.RegisteredModelAlias](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelAlias).
 * Added [catalog.RegisteredModelInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelInfo).
 * Added [catalog.SetRegisteredModelAliasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SetRegisteredModelAliasRequest).
 * Added [catalog.UpdateModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateModelVersionRequest).
 * Added [catalog.UpdateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateRegisteredModelRequest).
 * Added `Volumes` field for [compute.InitScriptInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfo).
 * Added [compute.VolumesStorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#VolumesStorageInfo).
 * Added [w.Files](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FilesAPI) workspace-level service.
 * Added [files.DeleteFileRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DeleteFileRequest).
 * Added [files.DownloadRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadRequest).
 * Added [files.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadResponse).
 * Added [files.UploadRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#UploadRequest).
 * Added `CustomTags` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
 * Added `CustomTags` field for [provisioning.UpdateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#UpdateWorkspaceRequest).
 * Added `CustomTags` field for [provisioning.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Workspace).
 * Added [provisioning.CustomTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CustomTags).
 * Added `Parameters` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest).
 * Added `RowLimit` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest).
 * Added [sql.StatementParameterListItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementParameterListItem).

SDK Internal Changes:
* Added support for io.ReadCloser for request and response bodies ([#590](https://github.com/databricks/databricks-sdk-go/pull/590)).
* Fixed nightly test failures for Go SDK ([#592](https://github.com/databricks/databricks-sdk-go/pull/592)).

OpenAPI Generator Changes:
* Improved error handling for ill-defined parameters ([#582](https://github.com/databricks/databricks-sdk-go/pull/582)).
* Switched to logging consistently rather than printing to stdout ([#588](https://github.com/databricks/databricks-sdk-go/pull/588)).

OpenAPI SHA: 5d0ccbb790d341eae8e85321a685a9e9e2d5bf24, Date: 2023-08-29

## 0.17.0

* Introduced Artifact Allowlist, Securable Tags, and Subentity Tags services.
* Introduced DeleteRuns and RestoreRuns methods in the Experiments API.
* Introduced the GetSecret method in the Secrets API.
* Renamed Auto Maintenance to Predictive Optimization.
* Set necessary headers when authenticating via Azure CLI ([#584](https://github.com/databricks/databricks-sdk-go/pull/584)).

New Services:

 * Added [w.ArtifactAllowlists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactAllowlistsAPI) workspace-level service.
 * Added [w.SecurableTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableTagsAPI) workspace-level service.
 * Added [w.SubentityTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SubentityTagsAPI) workspace-level service.
 * Added [catalog.ArtifactAllowlistInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactAllowlistInfo).
 * Added [catalog.ArtifactMatcher](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactMatcher).
 * Added [catalog.ArtifactType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactType).
 * Added [catalog.GetArtifactAllowlistRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetArtifactAllowlistRequest).
 * Added [catalog.ListSecurableTagsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSecurableTagsRequest).
 * Added [catalog.ListSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSecurableType).
 * Added [catalog.ListSubentityTagsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSubentityTagsRequest).
 * Added [catalog.MatchType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MatchType).
 * Added [catalog.SetArtifactAllowlist](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SetArtifactAllowlist).
 * Added [catalog.TagChanges](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagChanges).
 * Added [catalog.TagKeyValuePair](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagKeyValuePair).
 * Added [catalog.TagSecurable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSecurable).
 * Added [catalog.TagSecurableAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSecurableAssignment).
 * Added [catalog.TagSecurableAssignmentsList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSecurableAssignmentsList).
 * Added [catalog.TagSubentity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSubentity).
 * Added [catalog.TagSubentityAssignmentsList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSubentityAssignmentsList).
 * Added [catalog.TagsSubentityAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagsSubentityAssignment).
 * Added [catalog.UpdateSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateSecurableType).
 * Added [catalog.UpdateTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateTags).

New APIs:

 * Added `DeleteRuns` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `RestoreRuns` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added [ml.DeleteRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteRuns).
 * Added [ml.DeleteRunsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteRunsResponse).
 * Added [ml.RestoreRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RestoreRuns).
 * Added [ml.RestoreRunsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RestoreRunsResponse).
 * Added `GetSecret` method for [w.Secrets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#SecretsAPI) workspace-level service.
 * Added [workspace.GetSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetSecretRequest).
 * Added [workspace.GetSecretResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetSecretResponse).

Service Renames:

 * Renamed `EffectiveAutoMaintenanceFlag` field to `EffectivePredictiveOptimizationFlag` for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Renamed `EnableAutoMaintenance` field to `EnablePredictiveOptimization` for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Renamed [catalog.EffectiveAutoMaintenanceFlag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectiveAutoMaintenanceFlag) to [catalog.EffectivePredictiveOptimizationFlag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectivePredictiveOptimizationFlag).
 * Renamed [catalog.EffectiveAutoMaintenanceFlagInheritedFromType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectiveAutoMaintenanceFlagInheritedFromType) to [catalog.EffectivePredictiveOptimizationFlagInheritedFromType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectivePredictiveOptimizationFlagInheritedFromType).
 * Renamed [catalog.EnableAutoMaintenance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EnableAutoMaintenance) to [catalog.EnablePredictiveOptimization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EnablePredictiveOptimization).
 * Renamed `EffectiveAutoMaintenanceFlag` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo) to `EffectivePredictiveOptimizationFlag` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo).
 * Renamed `EnableAutoMaintenance` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo) to `EnablePredictiveOptimization` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo).
 * Renamed `EffectiveAutoMaintenanceFlag` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo) to `EffectivePredictiveOptimizationFlag` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
 * Renamed `EnableAutoMaintenance` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo) to `EnablePredictiveOptimization` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).

OpenAPI SHA: beff621d7b3e1d59244e2e34fc53a496f310e130, Date: 2023-08-17

## 0.16.0

* Added ability to generate flat names ([#568](https://github.com/databricks/databricks-sdk-go/pull/568)).

API Changes:

 * Changed `Create` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.AccountsStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsStorageCredentialInfo).
 * Changed `Get` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.AccountsStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsStorageCredentialInfo).
 * Changed `Update` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.AccountsStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsStorageCredentialInfo).
 * Changed `Create` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Changed `Update` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Removed `OptionsKvpairs` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Removed `PropertiesKvpairs` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `Options` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `Properties` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `ProvisioningState` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `SecurableKind` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `SecurableType` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Removed `OptionsKvpairs` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Removed `PropertiesKvpairs` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Added `Options` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Added `Properties` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Changed `Algorithm` field for [catalog.SseEncryptionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SseEncryptionDetails) to no longer be required.
 * Removed `OptionsKvpairs` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Added `Options` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Added [catalog.AccountsStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsStorageCredentialInfo).
 * Added [catalog.ConnectionInfoSecurableKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfoSecurableKind).
 * Added [catalog.ProvisioningState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningState).
 * Added `DataSecurityMode` field for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
 * Added `DockerImage` field for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
 * Added `SingleUserName` field for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
 * Removed `Schema` field for [iam.PartialUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PartialUpdate).
 * Added `Schemas` field for [iam.PartialUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PartialUpdate).

OpenAPI SHA: 1e3533f94335f0e6c5d9262bc1fea95b3ddcb0e1, Date: 2023-08-11
Dependency updates:

 * Bump golang.org/x/oauth2 from 0.10.0 to 0.11.0 ([#574](https://github.com/databricks/databricks-sdk-go/pull/574)).
 * Bump google.golang.org/api from 0.134.0 to 0.135.0 ([#575](https://github.com/databricks/databricks-sdk-go/pull/575)).
 * Bump google.golang.org/api from 0.135.0 to 0.136.0 ([#576](https://github.com/databricks/databricks-sdk-go/pull/576)).

## 0.15.0

To simplify documentation and management of object permissions, this release features a major reorganization of how permissions APIs are structured in the SDK. Rather than using a single `Permissions.Get()` API for all services, each service supporting permissions has its own permissions APIs. Follow these steps to migrate to the current SDK:

 * Change `w.Permissions.Get()` and `w.Permissions.GetByRequestObjectIdAndRequestObjectType()` to `w.<Service>.Get<Service>Permissions()`
 * Change `w.Permissions.GetPermissionLevels()` to `w.<Service>.GetServicePermissionLevels()`
 * Change `w.Permissions.Set()` to `w.<Service>.Set<Service>Permissions()`
 * Change `w.Permissions.Update()` to `w.<Service>.Update<Service>Permissions()`

API Changes:

 * Added `GetClusterPolicyPermissionLevels` method for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service.
 * Added `GetClusterPolicyPermissions` method for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service.
 * Added `SetClusterPolicyPermissions` method for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service.
 * Added `UpdateClusterPolicyPermissions` method for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service.
 * Added `GetClusterPermissionLevels` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `GetClusterPermissions` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `SetClusterPermissions` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `UpdateClusterPermissions` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `GetInstancePoolPermissionLevels` method for [w.InstancePools](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolsAPI) workspace-level service.
 * Added `GetInstancePoolPermissions` method for [w.InstancePools](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolsAPI) workspace-level service.
 * Added `SetInstancePoolPermissions` method for [w.InstancePools](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolsAPI) workspace-level service.
 * Added `UpdateInstancePoolPermissions` method for [w.InstancePools](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolsAPI) workspace-level service.
 * Added [compute.ClusterAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAccessControlRequest).
 * Added [compute.ClusterAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAccessControlResponse).
 * Added [compute.ClusterPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermission).
 * Added [compute.ClusterPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermissionLevel).
 * Added [compute.ClusterPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermissions).
 * Added [compute.ClusterPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermissionsDescription).
 * Added [compute.ClusterPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermissionsRequest).
 * Added [compute.ClusterPolicyAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyAccessControlRequest).
 * Added [compute.ClusterPolicyAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyAccessControlResponse).
 * Added [compute.ClusterPolicyPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermission).
 * Added [compute.ClusterPolicyPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermissionLevel).
 * Added [compute.ClusterPolicyPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermissions).
 * Added [compute.ClusterPolicyPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermissionsDescription).
 * Added [compute.ClusterPolicyPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermissionsRequest).
 * Added [compute.GetClusterPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPermissionLevelsRequest).
 * Added [compute.GetClusterPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPermissionLevelsResponse).
 * Added [compute.GetClusterPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPermissionsRequest).
 * Added [compute.GetClusterPolicyPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPolicyPermissionLevelsRequest).
 * Added [compute.GetClusterPolicyPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPolicyPermissionLevelsResponse).
 * Added [compute.GetClusterPolicyPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPolicyPermissionsRequest).
 * Added [compute.GetInstancePoolPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePoolPermissionLevelsRequest).
 * Added [compute.GetInstancePoolPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePoolPermissionLevelsResponse).
 * Added [compute.GetInstancePoolPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePoolPermissionsRequest).
 * Added [compute.InstancePoolAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAccessControlRequest).
 * Added [compute.InstancePoolAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAccessControlResponse).
 * Added [compute.InstancePoolPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermission).
 * Added [compute.InstancePoolPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermissionLevel).
 * Added [compute.InstancePoolPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermissions).
 * Added [compute.InstancePoolPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermissionsDescription).
 * Added [compute.InstancePoolPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermissionsRequest).
 * Changed `Set` method for [w.Permissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionsAPI) workspace-level service to start returning [iam.ObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ObjectPermissions).
 * Changed `Update` method for [w.Permissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionsAPI) workspace-level service to start returning [iam.ObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ObjectPermissions).
 * Added `GetPasswordPermissionLevels` method for [w.Users](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersAPI) workspace-level service.
 * Added `GetPasswordPermissions` method for [w.Users](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersAPI) workspace-level service.
 * Added `SetPasswordPermissions` method for [w.Users](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersAPI) workspace-level service.
 * Added `UpdatePasswordPermissions` method for [w.Users](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersAPI) workspace-level service.
 * Added `DisplayName` field for [iam.AccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#AccessControlResponse).
 * Changed `Roles` field for [iam.GetAssignableRolesForResourceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GetAssignableRolesForResourceResponse) to [iam.RoleList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#RoleList).
 * Added [iam.GetPasswordPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GetPasswordPermissionLevelsResponse).
 * Added [iam.PasswordAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordAccessControlRequest).
 * Added [iam.PasswordAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordAccessControlResponse).
 * Added [iam.PasswordPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermission).
 * Added [iam.PasswordPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermissionLevel).
 * Added [iam.PasswordPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermissions).
 * Added [iam.PasswordPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermissionsDescription).
 * Added [iam.PasswordPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermissionsRequest).
 * Added [iam.Role](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#Role).
 * Added `GetJobPermissionLevels` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service.
 * Added `GetJobPermissions` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service.
 * Added `SetJobPermissions` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service.
 * Added `UpdateJobPermissions` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service.
 * Added [jobs.GetJobPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetJobPermissionLevelsRequest).
 * Added [jobs.GetJobPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetJobPermissionLevelsResponse).
 * Added [jobs.GetJobPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetJobPermissionsRequest).
 * Added [jobs.JobAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobAccessControlRequest).
 * Added [jobs.JobAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobAccessControlResponse).
 * Added [jobs.JobPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermission).
 * Added [jobs.JobPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermissionLevel).
 * Added [jobs.JobPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermissions).
 * Added [jobs.JobPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermissionsDescription).
 * Added [jobs.JobPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermissionsRequest).
 * Added `GetExperimentPermissionLevels` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `GetExperimentPermissions` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `SetExperimentPermissions` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `UpdateExperimentPermissions` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `GetRegisteredModelPermissionLevels` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service.
 * Added `GetRegisteredModelPermissions` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service.
 * Added `SetRegisteredModelPermissions` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service.
 * Added `UpdateRegisteredModelPermissions` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service.
 * Added [ml.ExperimentAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentAccessControlRequest).
 * Added [ml.ExperimentAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentAccessControlResponse).
 * Added [ml.ExperimentPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermission).
 * Added [ml.ExperimentPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermissionLevel).
 * Added [ml.ExperimentPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermissions).
 * Added [ml.ExperimentPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermissionsDescription).
 * Added [ml.ExperimentPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermissionsRequest).
 * Added [ml.GetExperimentPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentPermissionLevelsRequest).
 * Added [ml.GetExperimentPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentPermissionLevelsResponse).
 * Added [ml.GetExperimentPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentPermissionsRequest).
 * Added [ml.GetRegisteredModelPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetRegisteredModelPermissionLevelsRequest).
 * Added [ml.GetRegisteredModelPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetRegisteredModelPermissionLevelsResponse).
 * Added [ml.GetRegisteredModelPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetRegisteredModelPermissionsRequest).
 * Added [ml.RegisteredModelAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelAccessControlRequest).
 * Added [ml.RegisteredModelAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelAccessControlResponse).
 * Added [ml.RegisteredModelPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermission).
 * Added [ml.RegisteredModelPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermissionLevel).
 * Added [ml.RegisteredModelPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermissions).
 * Added [ml.RegisteredModelPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermissionsDescription).
 * Added [ml.RegisteredModelPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermissionsRequest).
 * Added `Scopes` field for [oauth2.CreateCustomAppIntegration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateCustomAppIntegration).
 * Added `GetPipelinePermissionLevels` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Added `GetPipelinePermissions` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Added `SetPipelinePermissions` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Added `UpdatePipelinePermissions` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Added [pipelines.GetPipelinePermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelinePermissionLevelsRequest).
 * Added [pipelines.GetPipelinePermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelinePermissionLevelsResponse).
 * Added [pipelines.GetPipelinePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelinePermissionsRequest).
 * Added [pipelines.PipelineAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineAccessControlRequest).
 * Added [pipelines.PipelineAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineAccessControlResponse).
 * Added [pipelines.PipelinePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermission).
 * Added [pipelines.PipelinePermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermissionLevel).
 * Added [pipelines.PipelinePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermissions).
 * Added [pipelines.PipelinePermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermissionsDescription).
 * Added [pipelines.PipelinePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermissionsRequest).
 * Added `GcpManagedNetworkConfig` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
 * Added `GkeConfig` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
 * Added `GetServingEndpointPermissionLevels` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `GetServingEndpointPermissions` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `SetServingEndpointPermissions` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `UpdateServingEndpointPermissions` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `InstanceProfileArn` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
 * Added `InstanceProfileArn` field for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput).
 * Added [serving.GetServingEndpointPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetServingEndpointPermissionLevelsRequest).
 * Added [serving.GetServingEndpointPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetServingEndpointPermissionLevelsResponse).
 * Added [serving.GetServingEndpointPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetServingEndpointPermissionsRequest).
 * Added [serving.ServingEndpointAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointAccessControlRequest).
 * Added [serving.ServingEndpointAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointAccessControlResponse).
 * Added [serving.ServingEndpointPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermission).
 * Added [serving.ServingEndpointPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermissionLevel).
 * Added [serving.ServingEndpointPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermissions).
 * Added [serving.ServingEndpointPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermissionsDescription).
 * Added [serving.ServingEndpointPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermissionsRequest).
 * Added `GetTokenPermissionLevels` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
 * Added `GetTokenPermissions` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
 * Added `SetTokenPermissions` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
 * Added `UpdateTokenPermissions` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
 * Added [settings.GetTokenPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetTokenPermissionLevelsResponse).
 * Added [settings.TokenAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenAccessControlRequest).
 * Added [settings.TokenAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenAccessControlResponse).
 * Added [settings.TokenPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermission).
 * Added [settings.TokenPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermissionLevel).
 * Added [settings.TokenPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermissions).
 * Added [settings.TokenPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermissionsDescription).
 * Added [settings.TokenPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermissionsRequest).
 * Added `GetWarehousePermissionLevels` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
 * Added `GetWarehousePermissions` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
 * Added `SetWarehousePermissions` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
 * Added `UpdateWarehousePermissions` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
 * Added `CanSubscribeToLiveQuery` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
 * Removed `QueuedOverloadTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Removed `QueuedProvisioningTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Removed `TotalFilesCount` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Removed `TotalPartitionsCount` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `MetadataTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `OverloadingQueueStartTimestamp` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `PlanningPhases` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `PlanningTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `ProvisioningQueueStartTimestamp` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `PrunedBytes` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `PrunedFilesCount` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `QueryCompilationStartTimestamp` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `QueryExecutionTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added [sql.GetWarehousePermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehousePermissionLevelsRequest).
 * Added [sql.GetWarehousePermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehousePermissionLevelsResponse).
 * Added [sql.GetWarehousePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehousePermissionsRequest).
 * Added [sql.WarehouseAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehouseAccessControlRequest).
 * Added [sql.WarehouseAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehouseAccessControlResponse).
 * Added [sql.WarehousePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermission).
 * Added [sql.WarehousePermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissionLevel).
 * Added [sql.WarehousePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissions).
 * Added [sql.WarehousePermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissionsDescription).
 * Added [sql.WarehousePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissionsRequest).
 * Added `GetRepoPermissionLevels` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Added `GetRepoPermissions` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Added `SetRepoPermissions` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Added `UpdateRepoPermissions` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Added `GetWorkspaceObjectPermissionLevels` method for [w.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceAPI) workspace-level service.
 * Added `GetWorkspaceObjectPermissions` method for [w.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceAPI) workspace-level service.
 * Added `SetWorkspaceObjectPermissions` method for [w.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceAPI) workspace-level service.
 * Added `UpdateWorkspaceObjectPermissions` method for [w.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceAPI) workspace-level service.
 * Added [workspace.GetRepoPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetRepoPermissionLevelsRequest).
 * Added [workspace.GetRepoPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetRepoPermissionLevelsResponse).
 * Added [workspace.GetRepoPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetRepoPermissionsRequest).
 * Added [workspace.GetWorkspaceObjectPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetWorkspaceObjectPermissionLevelsRequest).
 * Added [workspace.GetWorkspaceObjectPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetWorkspaceObjectPermissionLevelsResponse).
 * Added [workspace.GetWorkspaceObjectPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetWorkspaceObjectPermissionsRequest).
 * Added [workspace.RepoAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoAccessControlRequest).
 * Added [workspace.RepoAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoAccessControlResponse).
 * Added [workspace.RepoPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermission).
 * Added [workspace.RepoPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermissionLevel).
 * Added [workspace.RepoPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermissions).
 * Added [workspace.RepoPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermissionsDescription).
 * Added [workspace.RepoPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermissionsRequest).
 * Added [workspace.WorkspaceObjectAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectAccessControlRequest).
 * Added [workspace.WorkspaceObjectAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectAccessControlResponse).
 * Added [workspace.WorkspaceObjectPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermission).
 * Added [workspace.WorkspaceObjectPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermissionLevel).
 * Added [workspace.WorkspaceObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermissions).
 * Added [workspace.WorkspaceObjectPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermissionsDescription).
 * Added [workspace.WorkspaceObjectPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermissionsRequest).

OpenAPI SHA: 386b65ecdc825b9c3ed4aa7ca88e2e5baf9d87df, Date: 2023-08-07
Dependency updates:

 * Bump google.golang.org/api from 0.131.0 to 0.134.0 ([#564](https://github.com/databricks/databricks-sdk-go/pull/564)).

## 0.14.1

* Handle nested query parameters in Client.attempt ([#559](https://github.com/databricks/databricks-sdk-go/pull/559)).
* Support x-databricks-path-style overrides at the operation level ([#562](https://github.com/databricks/databricks-sdk-go/pull/562)).


## 0.14.0

* Fixed names of keyword arguments in examples ([#560](https://github.com/databricks/databricks-sdk-go/pull/560)).

API Changes:

 * Changed `Create` method for [a.AccountMetastoreAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoreAssignmentsAPI) account-level service to no longer return [catalog.CreateMetastoreAssignmentsResponseItemList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMetastoreAssignmentsResponseItemList).
 * Added `ConnectionName` field for [catalog.CreateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCatalog).
 * Added `AccessPoint` field for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
 * Added `EncryptionDetails` field for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
 * Removed [catalog.CreateMetastoreAssignmentsResponseItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMetastoreAssignmentsResponseItem).
 * Added `AccessPoint` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
 * Added `EncryptionDetails` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
 * Added `AccessPoint` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
 * Added `EncryptionDetails` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
 * Added `AccessPoint` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
 * Added `EncryptionDetails` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
 * Added `AccessPoint` field for [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo).
 * Added `EncryptionDetails` field for [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo).
 * Added [catalog.EncryptionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EncryptionDetails).
 * Added [catalog.SseEncryptionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SseEncryptionDetails).
 * Added [catalog.SseEncryptionDetailsAlgorithm](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SseEncryptionDetailsAlgorithm).
 * Added [a.AccountNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyAPI) account-level service.
 * Added [settings.AccountNetworkPolicyMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyMessage).
 * Added [settings.DeleteAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyRequest).
 * Added [settings.DeleteAccountNetworkPolicyResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyResponse).
 * Added [settings.ReadAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadAccountNetworkPolicyRequest).
 * Added [settings.UpdateAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateAccountNetworkPolicyRequest).

OpenAPI SHA: fbdd0fa3e83fed2c798a58d376529bdb1285b915, Date: 2023-07-26

## 0.13.0

* Add issue templates ([#539](https://github.com/databricks/databricks-sdk-go/pull/539)).
* Added HasRequiredNonBodyField method ([#536](https://github.com/databricks/databricks-sdk-go/pull/536)).
* Make Azure MSI auth account compatible ([#544](https://github.com/databricks/databricks-sdk-go/pull/544)).
* Refactor Handling of Name<->ID Mapping in OpenAPI Generator ([#547](https://github.com/databricks/databricks-sdk-go/pull/547)).
* Regenerate Go SDK from current OpenAPI Specification ([#549](https://github.com/databricks/databricks-sdk-go/pull/549)).
* Parse Camel Case and Pascal Case Enum Values ([#550](https://github.com/databricks/databricks-sdk-go/pull/550)).
* Prepare for auto-releaser infra ([#554](https://github.com/databricks/databricks-sdk-go/pull/554)).
* Added SCIM Patch Acceptance Tests ([#540](https://github.com/databricks/databricks-sdk-go/pull/540)).

API Changes:

 * Removed `Maintenance` method for [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service.
 * Added `EnableOptimization` method for [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service.
 * Added `Update` method for [w.Tables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TablesAPI) workspace-level service.
 * Added `Force` field for [catalog.DeleteAccountMetastoreRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAccountMetastoreRequest).
 * Added `Force` field for [catalog.DeleteAccountStorageCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAccountStorageCredentialRequest).
 * Removed [catalog.UpdateAutoMaintenance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateAutoMaintenance).
 * Removed [catalog.UpdateAutoMaintenanceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateAutoMaintenanceResponse).
 * Added [catalog.UpdatePredictiveOptimization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePredictiveOptimization).
 * Added [catalog.UpdatePredictiveOptimizationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePredictiveOptimizationResponse).
 * Added [catalog.UpdateTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateTableRequest).
 * Added `Schema` field for [iam.PartialUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PartialUpdate).
 * Added [iam.PatchSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PatchSchema).
 * Added `TriggerInfo` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
 * Added `Health` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `JobSource` field for [jobs.GitSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GitSource).
 * Added `OnDurationWarningThresholdExceeded` field for [jobs.JobEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEmailNotifications).
 * Added `Health` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `TriggerInfo` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
 * Added `RunJobOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
 * Added `RunJobTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
 * Added `EmailNotifications` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added `Health` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added `EmailNotifications` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `Health` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `NotificationSettings` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `Health` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Added `RunJobTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Added `OnDurationWarningThresholdExceeded` field for [jobs.TaskEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TaskEmailNotifications).
 * Added `OnDurationWarningThresholdExceeded` field for [jobs.WebhookNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookNotifications).
 * Added [jobs.JobSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSource).
 * Added [jobs.JobSourceDirtyState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSourceDirtyState).
 * Added [jobs.JobsHealthMetric](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthMetric).
 * Added [jobs.JobsHealthOperator](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthOperator).
 * Added [jobs.JobsHealthRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthRule).
 * Added [jobs.JobsHealthRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthRules).
 * Added [jobs.RunJobOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunJobOutput).
 * Added [jobs.RunJobTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunJobTask).
 * Added [jobs.TriggerInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerInfo).
 * Added [jobs.WebhookNotificationsOnDurationWarningThresholdExceededItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookNotificationsOnDurationWarningThresholdExceededItem).
 * Removed `Whl` field for [pipelines.PipelineLibrary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineLibrary).
 * Changed `DeletePersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service with new required argument order.
 * Changed `ReadPersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service with new required argument order.
 * Changed `Etag` field for [settings.DeletePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeletePersonalComputeSettingRequest) to be required.
 * Changed `Etag` field for [settings.ReadPersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadPersonalComputeSettingRequest) to be required.
 * Added [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service.
 * Added [sharing.CentralCleanRoomInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CentralCleanRoomInfo).
 * Added [sharing.CleanRoomAssetInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomAssetInfo).
 * Added [sharing.CleanRoomCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomCatalog).
 * Added [sharing.CleanRoomCatalogUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomCatalogUpdate).
 * Added [sharing.CleanRoomCollaboratorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomCollaboratorInfo).
 * Added [sharing.CleanRoomInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomInfo).
 * Added [sharing.CleanRoomNotebookInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomNotebookInfo).
 * Added [sharing.CleanRoomTableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomTableInfo).
 * Added [sharing.ColumnInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ColumnInfo).
 * Added [sharing.ColumnMask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ColumnMask).
 * Added [sharing.ColumnTypeName](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ColumnTypeName).
 * Added [sharing.CreateCleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CreateCleanRoom).
 * Added [sharing.DeleteCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#DeleteCleanRoomRequest).
 * Added [sharing.GetCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#GetCleanRoomRequest).
 * Added [sharing.ListCleanRoomsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListCleanRoomsResponse).
 * Added [sharing.UpdateCleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateCleanRoom).
 * Changed `Query` field for [sql.Alert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Alert) to [sql.AlertQuery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertQuery).
 * Changed `Value` field for [sql.AlertOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertOptions) to `any`.
 * Removed `IsDbAdmin` field for [sql.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#User).
 * Removed `ProfileImageUrl` field for [sql.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#User).
 * Added [sql.AlertQuery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertQuery).

OpenAPI SHA: 0a1949ba96f71680dad30e06973eaae85b1307bb, Date: 2023-07-18
Dependency updates:

 * Bump golang.org/x/oauth2 from 0.9.0 to 0.10.0 ([#545](https://github.com/databricks/databricks-sdk-go/pull/545)).
 * Bump golang.org/x/mod from 0.11.0 to 0.12.0 ([#541](https://github.com/databricks/databricks-sdk-go/pull/541)).
 * Bump google.golang.org/api from 0.129.0 to 0.130.0 ([#542](https://github.com/databricks/databricks-sdk-go/pull/542)).
 * Bump google.golang.org/api from 0.130.0 to 0.131.0 ([#551](https://github.com/databricks/databricks-sdk-go/pull/551)).

## 0.12.0

 * Change release type
 * Regenerate SDK according to OpenAPI spec ([#534](https://github.com/databricks/databricks-sdk-go/pull/534)).
 * Bump google.golang.org/api from 0.128.0 to 0.129.0 ([#533](https://github.com/databricks/databricks-sdk-go/pull/533)).

## 0.11.0

 * Regenerated from OpenAPI spec ([#527](https://github.com/databricks/databricks-sdk-go/pull/527)). Includes bugfixes for System Tables.
 * Support accounts-dod as account endpoint ([#530](https://github.com/databricks/databricks-sdk-go/pull/530)).

Breaking API Changes:

 - Renamed compute.ClusterInfo to compute.ClusterDetails
 - Renamed compue.BaseClusterInfo to compute.ClusterSpec
 - Renamed jobs.RunSubmitTaskSettings to jobs.SubmitTask
 - Renamed jobs.JobTaskSettings to jobs.Task
 - Renamed jobs.CreateJobFormat to jobs.Format
 - Renamed jobs.JobsWebhookNotifications to jobs.WebhookNotifications
 - Renamed jobs.CronSchedulePauseStatus to jobs.PauseStatus
 - Renamed jobs.FileArrivalTriggerSettings to jobs.FileArrivalTriggerConfiguration
 - Renamed jobs.GitSourceGitProvider to jobs.GitProvider
 - Consolidated jobs.NotebookTaskSource, jobs.SparkPythonTaskSource to jobs.Source
 - Consolidated jobs.JobWebhookNotificationsOnFailureItem, jobs.JobWebhookNotificationsOnStartItem, jobs.JobWebhookNotificationsOnSuccessItem to jobs.Webhook
 - Renamed jobs.TaskDependenciesItem to job.TaskDependency
 - jobs.SparkSubmitTask.Widgets changed from *SqlDashboardWidgetOutput to []SqlDashboardWidgetOutput
 - (a *SystemSchemasAPI) DisableByMetastoreIdAndSchemaName changed from func(context.Context, string, string) error to func(context.Context, string, DisableSchemaName) error
 - (a *SystemSchemasAPI) Enable changed from func(context.Context) error to func(context.Context, EnableRequest) error
 - catalog.CatalogInfo.OptionsKvpairs changed from *OptionsKvPairs to map[string]string
 - catalog.CreateConnection.OptionsKvpairs changed from *OptionsKvPairs to map[string]string
 - catalog.UpdateConnection.OptionsKvpairs changed from OptionsKvPairs to map[string]string
 - catalog.DisableRequest.SchemaName changed from string to DisableSchemaName
 - catalog.OptionsKvPairs removed
 - catalog.SystemSchemaInfoStateDisableinitialized renamed to catalog.SystemSchemaInfoStateDisabledInitialized
 - catalog.SystemSchemaInfoStateEnabledcompleted renamed to catalog.SystemSchemaInfoStateEnabledCompleted
 - catalog.SystemSchemaInfoStateEnabledinitialized renamed to catalog.SystemSchemaInfoStateEnabledInitialized
 - settings.ReadPersonalComputeSettingsResponse renamed to settings.PersonalComputeSetting
 - workpace.CreateScope.KeyvaultMetadata removed

Other API Changes:

 - catalog.CatalogInfo.ConnectionName added
 - catalog.CatalogInfo.Options added
 - catalog.DisableSchemaName added
 - catalog.EnableRequest added
 - catalog.PrivilegeUseMarketplaceAssets added
 - compute.ClusterEvent added
 - compute.ComputeSpec added
 - compute.ComputeSpecKind added
 - jobs.ConditionTask added
 - jobs.ConditionTaskOp added
 - jobs.CreateJob.Compute added
 - jobs.JobCompute added
 - jobs.JobSettings.Compute added
 - jobs.RunConditionTask added
 - jobs.RunOutput.ConditionTask added
 - jobs.RunTask.ConditionTask added
 - serving.ServedModelInput.EnvironmentVars added
 - serving.ServedModelOutput.EnvironmentVars added
 - settings.DeletePersonalComputeSetting added
 - settings.UpdatePersonalComputeSetting added
 - sharing.PrivilegeUseMarketplaceAssets added
 - workspace.CreateScope.BackendAzureKeyVault added

## 0.10.1

 * Added ClusterID field for Config ([#524](https://github.com/databricks/databricks-sdk-go/pull/524)).
 * Added OnlyPayload property for method ([#522](https://github.com/databricks/databricks-sdk-go/pull/522)).
 * Renamed IsJsonOnly method ([#523](https://github.com/databricks/databricks-sdk-go/pull/523)).

Dependency updates:

* Bump google.golang.org/api from 0.127.0 to 0.128.0 ([#521](https://github.com/databricks/databricks-sdk-go/pull/521)).

## 0.10.0

 * Added log level support to SimpleLogger. Default logger now prints only `INFO` level messages. To replicate more verbose behavior from the previous versions, set the `DEBUG` level in `SimpleLogger` explicitly ([#426](https://github.com/databricks/databricks-sdk-go/pull/426)).
 * Added `Upload` and `Download` methods where applicable ([#423](https://github.com/databricks/databricks-sdk-go/pull/423)).
 * Added auth tests for a databrickscfg file with an empty DEFAULT profile ([#496](https://github.com/databricks/databricks-sdk-go/pull/496)).
 * Added more integration tests and examples ([#503](https://github.com/databricks/databricks-sdk-go/pull/503)).
 * Added preview level functions for services and methods to OpenAPI code ([#497](https://github.com/databricks/databricks-sdk-go/pull/497)).
 * Added IsAllRequiredFieldsPrimitive helper function ([#404](https://github.com/databricks/databricks-sdk-go/pull/404)).
 * Added isCrudCreate method ([#416](https://github.com/databricks/databricks-sdk-go/pull/416)).
 * Added retries for getting refreshable token ([#500](https://github.com/databricks/databricks-sdk-go/pull/500)).
 * Added type update when entity fields are updated ([#411](https://github.com/databricks/databricks-sdk-go/pull/411)).
 * Don't panic but return error from NewAccountClient ([#422](https://github.com/databricks/databricks-sdk-go/pull/422)).
 * Error when opening a DBFS directory for reading ([#415](https://github.com/databricks/databricks-sdk-go/pull/415)).
 * Fixed cleanup logic in SQL warehouse integration test ([#400](https://github.com/databricks/databricks-sdk-go/pull/400)).
 * Fixed error handling in cluster utility function ([#399](https://github.com/databricks/databricks-sdk-go/pull/399)).
 * Fixed example in README.md ([#494](https://github.com/databricks/databricks-sdk-go/pull/494)).
 * Fixed nondeterminism in workspace filesystem integration test ([#401](https://github.com/databricks/databricks-sdk-go/pull/401)).
 * Improve command execution interface ([#410](https://github.com/databricks/databricks-sdk-go/pull/410)).
 * Introduce waiters as top-level methods ([#408](https://github.com/databricks/databricks-sdk-go/pull/408)).
 * Regenerate examples if configured ([#518](https://github.com/databricks/databricks-sdk-go/pull/518)).
 * Respect limit field passed in ListRequests ([#407](https://github.com/databricks/databricks-sdk-go/pull/407)).
 * Updated from OpenAPI spec ([#412](https://github.com/databricks/databricks-sdk-go/pull/412), [#413](https://github.com/databricks/databricks-sdk-go/pull/413), [#421](https://github.com/databricks/databricks-sdk-go/pull/421), [#519](https://github.com/databricks/databricks-sdk-go/pull/519), [#424](https://github.com/databricks/databricks-sdk-go/pull/424)).
 * Updated API test template with page tokens ([#417](https://github.com/databricks/databricks-sdk-go/pull/417)).
 * Use constants instead of hardcoding strings ([#402](https://github.com/databricks/databricks-sdk-go/pull/402)).
 * Use `x-databricks-is-accounts` flag to determine whether a service is an account level service ([#420](https://github.com/databricks/databricks-sdk-go/pull/420)).

API Changes:

github.com/databricks/databricks-sdk-go/service/serving
  - (*ServingEndpointsAPI).Create: changed from func(context.Context, CreateServingEndpoint) (*ServingEndpointDetailed, error) to func(context.Context, CreateServingEndpoint) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error)
  - (*ServingEndpointsAPI).List: removed
  - (*ServingEndpointsAPI).UpdateConfig: changed from func(context.Context, EndpointCoreConfigInput) (*ServingEndpointDetailed, error) to func(context.Context, EndpointCoreConfigInput) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error)
  Compatible changes:
  - (*ServingEndpointsAPI).ListAll: added
  - (*ServingEndpointsAPI).WaitGetServingEndpointNotUpdating: added
  - WaitGetServingEndpointNotUpdating: added

github.com/databricks/databricks-sdk-go
  - WorkspaceClient.CommandExecutor: removed
  - WorkspaceClient.CommandExecution: added
  - WorkspaceClient.Connections: added
  - WorkspaceClient.Files: added
  - WorkspaceClient.SystemSchemas: added

github.com/databricks/databricks-sdk-go/service/sql
  - (*WarehousesAPI).Create: changed from func(context.Context, CreateWarehouseRequest) (*CreateWarehouseResponse, error) to func(context.Context, CreateWarehouseRequest) (*WaitGetWarehouseRunning[CreateWarehouseResponse], error)
  - (*WarehousesAPI).DeleteAndWait: removed
  - (*WarehousesAPI).DeleteByIdAndWait: removed
  - (*WarehousesAPI).Edit: changed from func(context.Context, EditWarehouseRequest) error to func(context.Context, EditWarehouseRequest) (*WaitGetWarehouseRunning[any], error)
  - (*WarehousesAPI).GetAndWait: removed
  - (*WarehousesAPI).GetByIdAndWait: removed
  - (*WarehousesAPI).Start: changed from func(context.Context, StartRequest) error to func(context.Context, StartRequest) (*WaitGetWarehouseRunning[any], error)
  - (*WarehousesAPI).Stop: changed from func(context.Context, StopRequest) error to func(context.Context, StopRequest) (*WaitGetWarehouseStopped[any], error)
  - (*StatementExecutionAPI).ExecuteAndWait: added
  - (*WarehousesAPI).WaitGetWarehouseRunning: added
  - (*WarehousesAPI).WaitGetWarehouseStopped: added
  - FormatCsv: added
  - WaitGetWarehouseRunning: added
  - WaitGetWarehouseStopped: added

github.com/databricks/databricks-sdk-go/service/jobs
  - (*JobsAPI).CancelRun: changed from func(context.Context, CancelRun) error to func(context.Context, CancelRun) (*WaitGetRunJobTerminatedOrSkipped[any], error)
  - (*JobsAPI).GetRunAndWait: removed
  - (*JobsAPI).RepairRun: changed from func(context.Context, RepairRun) (*RepairRunResponse, error) to func(context.Context, RepairRun) (*WaitGetRunJobTerminatedOrSkipped[RepairRunResponse], error)
  - (*JobsAPI).RunNow: changed from func(context.Context, RunNow) (*RunNowResponse, error) to func(context.Context, RunNow) (*WaitGetRunJobTerminatedOrSkipped[RunNowResponse], error)
  - (*JobsAPI).Submit: changed from func(context.Context, SubmitRun) (*SubmitRunResponse, error) to func(context.Context, SubmitRun) (*WaitGetRunJobTerminatedOrSkipped[SubmitRunResponse], error)
  - (*JobsAPI).WaitGetRunJobTerminatedOrSkipped: added
  - CreateJob.RunAs: added
  - JobRunAs: added
  - JobSettings.RunAs: added
  - WaitGetRunJobTerminatedOrSkipped: added

github.com/databricks/databricks-sdk-go/service/compute
  - (*ClustersAPI).Create: changed from func(context.Context, CreateCluster) (*CreateClusterResponse, error) to func(context.Context, CreateCluster) (*WaitGetClusterRunning[CreateClusterResponse], error)
  - (*ClustersAPI).Delete: changed from func(context.Context, DeleteCluster) error to func(context.Context, DeleteCluster) (*WaitGetClusterTerminated[any], error)
  - (*ClustersAPI).Edit: changed from func(context.Context, EditCluster) error to func(context.Context, EditCluster) (*WaitGetClusterRunning[any], error)
  - (*ClustersAPI).GetAndWait: removed
  - (*ClustersAPI).GetByClusterIdAndWait: removed
  - (*ClustersAPI).Resize: changed from func(context.Context, ResizeCluster) error to func(context.Context, ResizeCluster) (*WaitGetClusterRunning[any], error)
  - (*ClustersAPI).Restart: changed from func(context.Context, RestartCluster) error to func(context.Context, RestartCluster) (*WaitGetClusterRunning[any], error)
  - (*ClustersAPI).Start: changed from func(context.Context, StartCluster) error to func(context.Context, StartCluster) (*WaitGetClusterRunning[any], error)
  - (*CommandExecutionAPI).Cancel: changed from func(context.Context, CancelCommand) error to func(context.Context, CancelCommand) (*WaitCommandStatusCommandExecutionCancelled[any], error)
  - (*CommandExecutionAPI).Create: changed from func(context.Context, CreateContext) (*Created, error) to func(context.Context, CreateContext) (*WaitContextStatusCommandExecutionRunning[Created], error)
  - (*CommandExecutionAPI).Execute: changed from func(context.Context, Command) (*Created, error) to func(context.Context, Command) (*WaitCommandStatusCommandExecutionFinishedOrError[Created], error)
  - (*ClustersAPI).WaitGetClusterRunning: added
  - (*ClustersAPI).WaitGetClusterTerminated: added
  - (*CommandExecutionAPI).Start: added
  - (*CommandExecutionAPI).WaitCommandStatusCommandExecutionCancelled: added
  - (*CommandExecutionAPI).WaitCommandStatusCommandExecutionFinishedOrError: added
  - (*CommandExecutionAPI).WaitContextStatusCommandExecutionRunning: added
  - BaseClusterInfo.DataSecurityMode: added
  - BaseClusterInfo.DockerImage: added
  - BaseClusterInfo.SingleUserName: added
  - ClusterAttributes.DataSecurityMode: added
  - ClusterAttributes.DockerImage: added
  - ClusterAttributes.SingleUserName: added
  - ClusterInfo.DockerImage: added
  - CommandExecutorV2: added
  - EditCluster.DataSecurityMode: added
  - EditCluster.DockerImage: added
  - EditCluster.SingleUserName: added
  - WaitCommandStatusCommandExecutionCancelled: added
  - WaitCommandStatusCommandExecutionFinishedOrError: added
  - WaitContextStatusCommandExecutionRunning: added
  - WaitGetClusterRunning: added
  - WaitGetClusterTerminated: added

github.com/databricks/databricks-sdk-go/service/pipelines
  - (*PipelinesAPI).GetAndWait: removed
  - (*PipelinesAPI).GetByPipelineIdAndWait: removed
  - (*PipelinesAPI).Reset: changed from func(context.Context, ResetRequest) error to func(context.Context, ResetRequest) (*WaitGetPipelineRunning[any], error)
  - (*PipelinesAPI).Stop: changed from func(context.Context, StopRequest) error to func(context.Context, StopRequest) (*WaitGetPipelineIdle[any], error)
  - (*PipelinesAPI).WaitGetPipelineIdle: added
  - (*PipelinesAPI).WaitGetPipelineRunning: added
  - WaitGetPipelineIdle: added
  - WaitGetPipelineRunning: added

github.com/databricks/databricks-sdk-go/service/sharing
  - (*ProvidersAPI).ListShares: removed
  - (*ProvidersAPI).ListSharesAll: added

github.com/databricks/databricks-sdk-go/service/workspace
  - ExportFormatAuto: removed
  - ExportRequest.DirectDownload: removed
  - Import.Format: changed from ExportFormat to ImportFormat
  - (*WorkspaceAPI).Download: added
  - (*WorkspaceAPI).ReadFile: added
  - (*WorkspaceAPI).Upload: added
  - (*WorkspaceAPI).WriteFile: added
  - DownloadFormat: added
  - DownloadOption: added
  - ImportFormat: added
  - ImportFormatAuto: added
  - ImportFormatDbc: added
  - ImportFormatHtml: added
  - ImportFormatJupyter: added
  - ImportFormatRMarkdown: added
  - ImportFormatSource: added
  - UploadFormat: added
  - UploadLanguage: added
  - UploadOption: added
  - UploadOverwrite: added

github.com/databricks/databricks-sdk-go/service/catalog
  - (*FunctionsAPI).List: removed
  - (*MetastoresAPI).Assign: changed from func(context.Context, CreateMetastoreAssignment) error to func(context.Context, CreateMetastoreAssignment) error
  - (*MetastoresAPI).Create: changed from func(context.Context, CreateMetastore) (*MetastoreInfo, error) to func(context.Context, CreateMetastore) (*MetastoreInfo, error)
  - (*MetastoresAPI).Update: changed from func(context.Context, UpdateMetastore) (*MetastoreInfo, error) to func(context.Context, UpdateMetastore) (*MetastoreInfo, error)
  - (*MetastoresAPI).UpdateAssignment: changed from func(context.Context, UpdateMetastoreAssignment) error to func(context.Context, UpdateMetastoreAssignment) error
  - (*StorageCredentialsAPI).Create: changed from func(context.Context, CreateStorageCredential) (*StorageCredentialInfo, error) to func(context.Context, CreateStorageCredential) (*StorageCredentialInfo, error)
  - (*StorageCredentialsAPI).Update: changed from func(context.Context, UpdateStorageCredential) (*StorageCredentialInfo, error) to func(context.Context, UpdateStorageCredential) (*StorageCredentialInfo, error)
  - (*TablesAPI).ListSummaries: removed
  - CreateMetastore.Name: removed
  - CreateMetastore.Region: removed
  - CreateMetastore.StorageRoot: removed
  - CreateMetastore: changed from CreateMetastore to CreateMetastore
  - CreateMetastoreAssignment.DefaultCatalogName: removed
  - CreateMetastoreAssignment: changed from CreateMetastoreAssignment to CreateMetastoreAssignment
  - CreateStorageCredential.AwsIamRole: removed
  - CreateStorageCredential.AzureServicePrincipal: removed
  - CreateStorageCredential.Comment: removed
  - CreateStorageCredential.GcpServiceAccountKey: removed
  - CreateStorageCredential.Name: removed
  - CreateStorageCredential.ReadOnly: removed
  - CreateStorageCredential.SkipValidation: removed
  - CreateStorageCredential: changed from CreateStorageCredential to CreateStorageCredential
  - GcpServiceAccountKey: removed
  - ListFunctionsResponse.Schemas: removed
  - MetastoreAssignment.WorkspaceId: changed from string to int64
  - StorageCredentialInfo.GcpServiceAccountKey: removed
  - UpdateMetastore.DeltaSharingOrganizationName: removed
  - UpdateMetastore.DeltaSharingRecipientTokenLifetimeInSeconds: removed
  - UpdateMetastore.DeltaSharingScope: removed
  - UpdateMetastore.Id: removed
  - UpdateMetastore.Name: removed
  - UpdateMetastore.Owner: removed
  - UpdateMetastore.PrivilegeModelVersion: removed
  - UpdateMetastore.StorageRootCredentialId: removed
  - UpdateMetastore: changed from UpdateMetastore to UpdateMetastore
  - UpdateMetastoreAssignment.DefaultCatalogName: removed
  - UpdateMetastoreAssignment: changed from UpdateMetastoreAssignment to UpdateMetastoreAssignment
  - UpdateStorageCredential.AwsIamRole: removed
  - UpdateStorageCredential.AzureServicePrincipal: removed
  - UpdateStorageCredential.Comment: removed
  - UpdateStorageCredential.Force: removed
  - UpdateStorageCredential.GcpServiceAccountKey: removed
  - UpdateStorageCredential.Owner: removed
  - UpdateStorageCredential.ReadOnly: removed
  - UpdateStorageCredential.SkipValidation: removed
  - UpdateStorageCredential: changed from UpdateStorageCredential to UpdateStorageCredential
  - ValidateStorageCredential.GcpServiceAccountKey: removed
  - (*FunctionsAPI).FunctionInfoNameToFullNameMap: added
  - (*FunctionsAPI).ListAll: added
  - (*TablesAPI).ListSummariesAll: added
  - AccountsCreateMetastore: added
  - AccountsCreateMetastoreAssignment: added
  - AccountsCreateStorageCredential: added
  - AccountsUpdateMetastore: added
  - AccountsUpdateMetastoreAssignment: added
  - AccountsUpdateStorageCredential: added
  - AzureManagedIdentity: added
  - ConnectionInfo: added
  - ConnectionType: added
  - ConnectionTypeDatabricks: added
  - ConnectionTypeMysql: added
  - ConnectionTypePostgresql: added
  - ConnectionTypeRedshift: added
  - ConnectionTypeSnowflake: added
  - ConnectionTypeSqldw: added
  - ConnectionTypeSqlserver: added
  - ConnectionsAPI: added
  - ConnectionsService: added
  - CreateConnection: added
  - CreateMetastore.MetastoreInfo: added
  - CreateMetastoreAssignment.MetastoreAssignment: added
  - CreateStorageCredential.CredentialInfo: added
  - CredentialType: added
  - CredentialTypeUsernamePassword: added
  - DatabricksGcpServiceAccountResponse: added
  - DeleteConnectionRequest: added
  - DisableRequest: added
  - GetConnectionRequest: added
  - ListConnectionsResponse: added
  - ListFunctionsResponse.Functions: added
  - ListSystemSchemasRequest: added
  - ListSystemSchemasResponse: added
  - NewConnections: added
  - NewSystemSchemas: added
  - OptionsKvPairs: added
  - PropertiesKvPairs: added
  - StorageCredentialInfo.AzureManagedIdentity: added
  - StorageCredentialInfo.DatabricksGcpServiceAccount: added
  - SystemSchemaInfo: added
  - SystemSchemaInfoState: added
  - SystemSchemaInfoStateDisableinitialized: added
  - SystemSchemaInfoStateEnablecompleted: added
  - SystemSchemaInfoStateEnableinitialized: added
  - SystemSchemaInfoStateUnavailable: added
  - SystemSchemasAPI: added
  - SystemSchemasService: added
  - UpdateConnection: added
  - UpdateMetastore.MetastoreInfo: added
  - UpdateMetastoreAssignment.MetastoreAssignment: added
  - UpdateStorageCredential.CredentialInfo: added
  - ValidateStorageCredential.AzureManagedIdentity: added
  - ValidateStorageCredential.DatabricksGcpServiceAccount: added

github.com/databricks/databricks-sdk-go/service/ml
  - GetModelResponse.RegisteredModel: removed
  - GetModelResponse.RegisteredModelDatabricks: added

github.com/databricks/databricks-sdk-go/service/files
  - FilesAPI: added
  - FilesService: added
  - NewFiles: added

github.com/databricks/databricks-sdk-go/logger
  - SimpleLogger.Level: added

github.com/databricks/databricks-sdk-go/service/provisioning
  - (*WorkspacesAPI).Create: changed from func(context.Context, CreateWorkspaceRequest) (*Workspace, error) to func(context.Context, CreateWorkspaceRequest) (*WaitGetWorkspaceRunning[Workspace], error)
  - (*WorkspacesAPI).Update: changed from func(context.Context, UpdateWorkspaceRequest) error to func(context.Context, UpdateWorkspaceRequest) (*WaitGetWorkspaceRunning[any], error)
  - (*WorkspacesAPI).WaitGetWorkspaceRunning: added
  - WaitGetWorkspaceRunning: added

github.com/databricks/databricks-sdk-go/service/iam
  - AccountAccessControlAPI: added
  - AccountAccessControlProxyAPI: added
  - AccountAccessControlProxyService: added
  - AccountAccessControlService: added
  - GetAssignableRolesForResourceRequest: added
  - GetAssignableRolesForResourceResponse: added
  - GetRuleSetRequest: added
  - GrantRule: added
  - NewAccountAccessControl: added
  - NewAccountAccessControlProxy: added
  - RuleSetResponse: added
  - RuleSetUpdateRequest: added
  - UpdateRuleSetRequest: added

Dependency updates:

 * Bump github.com/stretchr/testify from 1.8.3 to 1.8.4 ([#406](https://github.com/databricks/databricks-sdk-go/pull/406)).
 * Bump golang.org/x/mod from 0.10.0 to 0.11.0 ([#515](https://github.com/databricks/databricks-sdk-go/pull/515)).
 * Bump golang.org/x/oauth2 from 0.8.0 to 0.9.0 ([#498](https://github.com/databricks/databricks-sdk-go/pull/498)).
 * Bump google.golang.org/api from 0.123.0 to 0.127.0 ([#405](https://github.com/databricks/databricks-sdk-go/pull/405), [#425](https://github.com/databricks/databricks-sdk-go/pull/425), [#429](https://github.com/databricks/databricks-sdk-go/pull/429)).

## 0.9.0

 * Added more usage examples for `go doc` and Go Packages ([#389](https://github.com/databricks/databricks-sdk-go/pull/389)).
 * Make u2m authentication work with new CLI ([#394](https://github.com/databricks/databricks-sdk-go/pull/394)).
 * Update from OpenAPI spec (19 may) ([#390](https://github.com/databricks/databricks-sdk-go/pull/390)).

Dependency updates:

 * Bump github.com/stretchr/testify from 1.8.2 to 1.8.3 ([#393](https://github.com/databricks/databricks-sdk-go/pull/393)).
 * Bump google.golang.org/api from 0.122.0 to 0.123.0 ([#392](https://github.com/databricks/databricks-sdk-go/pull/392)).

## 0.8.1

 * Added `in` codegen function ([#387](https://github.com/databricks/databricks-sdk-go/pull/387)).
 * Fixed mlflow acceptance tests ([#378](https://github.com/databricks/databricks-sdk-go/pull/378)).
 * Fixed MLflow integration test and removed workaround for `DELETE` query parameter ([#380](https://github.com/databricks/databricks-sdk-go/pull/380)).
 * Make clusters acceptance tests robust to duplicate cluster names ([#381](https://github.com/databricks/databricks-sdk-go/pull/381)).
 * Remove dead code from apierr/errors.go ([#376](https://github.com/databricks/databricks-sdk-go/pull/376)).
 * Serialize params to request body on delete ([#383](https://github.com/databricks/databricks-sdk-go/pull/383)).

Dependency updates:

 * Bump golang.org/x/oauth2 from 0.7.0 to 0.8.0 ([#385](https://github.com/databricks/databricks-sdk-go/pull/385)).
 * Bump google.golang.org/api from 0.118.0 to 0.122.0 ([#382](https://github.com/databricks/databricks-sdk-go/pull/382), [#386](https://github.com/databricks/databricks-sdk-go/pull/386)).

## 0.8.0

 * Added more code generation utilities ([#369](https://github.com/databricks/databricks-sdk-go/pull/369)).
 * Body logger for non-JSON payload as well ([#365](https://github.com/databricks/databricks-sdk-go/pull/365)).
 * Cleanup ephemeral groups in integration tests ([#368](https://github.com/databricks/databricks-sdk-go/pull/368)).
 * Fixed external entity generation ([#372](https://github.com/databricks/databricks-sdk-go/pull/372)).
 * Skip loading default profile if host is already configured ([#363](https://github.com/databricks/databricks-sdk-go/pull/363)).
 * Update debug messages in config loader to display correct path ([#362](https://github.com/databricks/databricks-sdk-go/pull/362)).

Dependency updates:

 * Bump golang.org/x/oauth2 from 0.6.0 to 0.7.0 ([#364](https://github.com/databricks/databricks-sdk-go/pull/364)).
 * Bump google.golang.org/api from 0.115.0 to 0.116.0 ([#361](https://github.com/databricks/databricks-sdk-go/pull/361)).
 * Bump google.golang.org/api from 0.116.0 to 0.118.0 ([#367](https://github.com/databricks/databricks-sdk-go/pull/367)).
 
API changes:

 * Moved `clusterpolicies` APIs to `compute` package.
 * Moved `clusters` APIs to `compute` package.
 * Moved `commands` APIs to `compute` package.
 * Moved `globalinitscripts` APIs to `compute` package.
 * Moved `instancepools` APIs to `compute` package.
 * Moved `scim` APIs to `iam` package.
 * Moved `permissions` APIs to `iam` package.
 * Moved `ipaccesslists` APIs to `settings` package.
 * Moved `tokenmanagement` APIs to `settings` package.
 * Moved `tokens` APIs to `settings` package.
 * Moved `workspaceconf` APIs to `settings` package.
 * Moved `gitcredentials` APIs to `workspace` package.
 * Moved `repos` APIs to `workspace` package.
 * Moved `secrets` APIs to `workspace` package.
 * Split `unitcatalog` package to `catalog` and `sharing`.
 * Renamed `mlflow` package to `ml`.
 * Renamed `dbfs` package to `files`.
 * Renamed `deployment` package to `provisioning`.
 * Renamed `endpoints` package to `serving`.
 * Renamed `clusters.List` type to `compute.ListClustersRequest`.
 * Renamed `jobs.ListRuns` type to `jobs.ListRunsRequest`.
 * Renamed `jobs.ExportRun` type to `jobs.ExportRunRequest`.
 * Renamed `clusterpolicies.List` type to `compute.ListClusterPoliciesRequest`.
 * Renamed `jobs.List` type to `jobs.ListJobsRequest`.
 * Renamed `permissions.GetPermissionLevels` type to `iam.GetPermissionLevelsRequest`.
 * Renamed `pipelines.ListPipelineEvents` type to `pipelines.ListPipelineEventsRequest`.
 * Renamed `pipelines.ListPipelines` type to `pipelines.ListPipelinesRequest`.
 * Renamed `workspaceconf.GetStatus` type to `settings.GetStatusRequest`.
 * Renamed `repos.List` type to `workspace.ListReposRequest`.
 * Renamed `tokenmanagement.List` type to `settings.ListTokenManagementRequest`.
 * Renamed `workspace.Export` type to `workspace.ExportRequest`.
 * Renamed `workspace.List` type to `workspace.ListWorkspaceRequest`.

## 0.7.0

 * Update from OpenAPI ([#359](https://github.com/databricks/databricks-sdk-go/pull/359)).
 * Experimental credentials provider via local server ([#340](https://github.com/databricks/databricks-sdk-go/pull/340)).
 * Added `isTesting` marker to support `ResourceFixture` in Terraform ([#358](https://github.com/databricks/databricks-sdk-go/pull/358)).

Dependency updates:
 
 * Bump golang.org/x/mod from 0.9.0 to 0.10.0 ([#356](https://github.com/databricks/databricks-sdk-go/pull/356)).
 * Bump google.golang.org/api from 0.114.0 to 0.115.0 ([#357](https://github.com/databricks/databricks-sdk-go/pull/357)).

## 0.6.0

 * Added type to represent a loaded configuration file ([#349](https://github.com/databricks/databricks-sdk-go/pull/349)).
 * Added named `Wait` the level of services ([#348](https://github.com/databricks/databricks-sdk-go/pull/348)).
 * Pass `azure_client_id` to Azure MSI if specified ([#354](https://github.com/databricks/databricks-sdk-go/pull/354)).

Dependency updates:
 
 * Bump google.golang.org/api from 0.112.0 to 0.114.0 ([#344](https://github.com/databricks/databricks-sdk-go/pull/344)).

## 0.5.0

 * Added Delta Live Tables events method to `pipelines.PipelinesAPI` ([#339](https://github.com/databricks/databricks-sdk-go/pull/339)).
 * Added `id` field to the `SparkVersionRequest` for easier integration with Terraform's [databricks_spark_version](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/spark_version) data source ([#335](https://github.com/databricks/databricks-sdk-go/pull/335)).
 * Improve auth fixture capturing ([#336](https://github.com/databricks/databricks-sdk-go/pull/336), [#337](https://github.com/databricks/databricks-sdk-go/pull/337)).
 * Make logger context aware and added examples of usage together with [`zerolog`](https://github.com/databricks/databricks-sdk-go/tree/main/examples/zerolog) and [`slog`](https://github.com/databricks/databricks-sdk-go/tree/main/examples/slog) ([#333](https://github.com/databricks/databricks-sdk-go/pull/333)).
 * Update from OpenAPI spec and added new APIs ([#338](https://github.com/databricks/databricks-sdk-go/pull/338)).
 
Dependency updates:
 
 * Bump google.golang.org/api from 0.111.0 to 0.112.0 ([#334](https://github.com/databricks/databricks-sdk-go/pull/334)).

## 0.4.1

 * Added an option to configure the location of the bricks CLI ([#330](https://github.com/databricks/databricks-sdk-go/pull/330)).
 * Added support for Azure CLI authentication on Azure China and Azure GovCloud ([#331](https://github.com/databricks/databricks-sdk-go/pull/331)).

## 0.4.0

 * Added `id` field to the `NodeTypeRequest` for easier integration with Terraform's [databricks_node_type](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/node_type) data source ([#325](https://github.com/databricks/databricks-sdk-go/pull/325)).
 * Don't load `~/.databrickscfg` if `azure-cli` auth is implicitly configured ([#324](https://github.com/databricks/databricks-sdk-go/pull/324)).
 * Fixed newline in codegen comments ([#326](https://github.com/databricks/databricks-sdk-go/pull/326)).
 * Update from OpenAPI ([#320](https://github.com/databricks/databricks-sdk-go/pull/320), [#328](https://github.com/databricks/databricks-sdk-go/pull/328)).

API changes:

* Renamed `deployment.AwsCredentials` to `deployment.CreateCredentialAwsCredentials`.
* Renamed `deployment.StsRole` to `deployment.CreateCredentialStsRole`.
* Removed schedules support from `sql.AlertsAPI`.

Dependency updates:

 * Bump github.com/stretchr/testify from 1.8.1 to 1.8.2 ([#318](https://github.com/databricks/databricks-sdk-go/pull/318)).
 * Bump golang.org/x/mod from 0.6.0-dev.0.20220419223038-86c51ed26bb4 to 0.8.0 ([#316](https://github.com/databricks/databricks-sdk-go/pull/316)).
 * Bump golang.org/x/mod from 0.8.0 to 0.9.0 ([#323](https://github.com/databricks/databricks-sdk-go/pull/323)).
 * Bump golang.org/x/oauth2 from 0.5.0 to 0.6.0 ([#322](https://github.com/databricks/databricks-sdk-go/pull/322)).
 * Bump golang.org/x/time from 0.0.0-20210723032227-1f47c861a9ac to 0.3.0 ([#317](https://github.com/databricks/databricks-sdk-go/pull/317)).
 * Bump google.golang.org/api from 0.110.0 to 0.111.0 ([#319](https://github.com/databricks/databricks-sdk-go/pull/319)).

## 0.3.3

 * Allow AAD SPN authentication on Databricks Account level ([#311](https://github.com/databricks/databricks-sdk-go/pull/311)).
 * Port auth tests from the JS SDK ([#313](https://github.com/databricks/databricks-sdk-go/pull/313)).
 * Skip loading `~/.databrickscfg` when not required ([#314](https://github.com/databricks/databricks-sdk-go/pull/314)).

Dependency updates:

 * Bump golang.org/x/net from 0.6.0 to 0.7.0 ([#312](https://github.com/databricks/databricks-sdk-go/pull/312)).

## 0.3.2

 * Always use the latest value for user agent key ([#309](https://github.com/databricks/databricks-sdk-go/pull/309)).

## 0.3.1

 * Change APIError to use pointer receiver ([#298](https://github.com/databricks/databricks-sdk-go/pull/298)).
 * Drop duplicate prefix in randomized email ([#299](https://github.com/databricks/databricks-sdk-go/pull/299)).
 * Expand semver pattern to be compliant with https://semver.org ([#302](https://github.com/databricks/databricks-sdk-go/pull/302)).
 * Fix `apierr.APIError` pointer receivers ([#307](https://github.com/databricks/databricks-sdk-go/pull/307)).
 * Skip loading config if auth is already explicitly configured directly ([#306](https://github.com/databricks/databricks-sdk-go/pull/306)).
 * Sync fixes for smallest node selection from Terraform ([#301](https://github.com/databricks/databricks-sdk-go/pull/301)).
 * Updated from OpenAPI spec ([#305](https://github.com/databricks/databricks-sdk-go/pull/305)).

Dependency updates:

 * Bump google.golang.org/api from 0.109.0 to 0.110.0 ([#303](https://github.com/databricks/databricks-sdk-go/pull/303)).

## 0.3.0

 * Added support for GCP deployment APIs in Public Preview.
 * Added new node type selector ([#287](https://github.com/databricks/databricks-sdk-go/pull/287)).
 * Added Azure MSI auth ([#279](https://github.com/databricks/databricks-sdk-go/pull/279)).
 * Added validation of strings in user agent package ([#291](https://github.com/databricks/databricks-sdk-go/pull/291)).
 * Added logging for error responses ([#289](https://github.com/databricks/databricks-sdk-go/pull/289)).
 * Update implementation of smallest node selection to match terraform's ([#273](https://github.com/databricks/databricks-sdk-go/pull/273)).
 * Added experimental OAuth support ([#276](https://github.com/databricks/databricks-sdk-go/pull/276), [#266](https://github.com/databricks/databricks-sdk-go/pull/266), [#292](https://github.com/databricks/databricks-sdk-go/pull/292)).
 * Fixed `google-credentials` to take precedence over `google-id` ([#283](https://github.com/databricks/databricks-sdk-go/pull/283)).
 * Test with Go 1.20 ([#295](https://github.com/databricks/databricks-sdk-go/pull/295)).
 * Regenerate from OpenAPI ([#293](https://github.com/databricks/databricks-sdk-go/pull/293), [#294](https://github.com/databricks/databricks-sdk-go/pull/294), [#296](https://github.com/databricks/databricks-sdk-go/pull/296), [#282](https://github.com/databricks/databricks-sdk-go/pull/282), [#269](https://github.com/databricks/databricks-sdk-go/pull/269), [#272](https://github.com/databricks/databricks-sdk-go/pull/272)).
 * Added `replaceAll` and `lowerFirst` template functions ([#288](https://github.com/databricks/databricks-sdk-go/pull/288), [#277](https://github.com/databricks/databricks-sdk-go/pull/277)).

Dependency updates:

 * Bump google.golang.org/api from 0.105.0 to 0.109.0 ([#271](https://github.com/databricks/databricks-sdk-go/pull/271), [#280](https://github.com/databricks/databricks-sdk-go/pull/280), [#284](https://github.com/databricks/databricks-sdk-go/pull/284), [#290](https://github.com/databricks/databricks-sdk-go/pull/290)).

API changes:

 * Renamed `clusters.CreateCluster` to `clusters.BaseClusterInfo`.
 * Renamed `jobs.Job` to `jobs.BaseJob`.
 * Renamed `jobs.Run` to `jobs.BaseRun`.
 * Fixed `commands.Results` `Schema` type from `[][]any` to `[]map[string]any`.

## 0.2.0

 * Added `DATABRICKS_AUTH_TYPE` environment variable ([#248](https://github.com/databricks/databricks-sdk-go/pull/248)).
 * Added Policy Families API ([#263](https://github.com/databricks/databricks-sdk-go/pull/263)).
 * Added experimental `ErrCannotConfigureAuth` and `ErrNotAccountClient` ([#237](https://github.com/databricks/databricks-sdk-go/pull/237), [#238](https://github.com/databricks/databricks-sdk-go/pull/238)).
 * Added DBFS file handle that supports both reading and writing ([#261](https://github.com/databricks/databricks-sdk-go/pull/261)).
 * Added `io.WriterTo` for DBFS file reader ([#249](https://github.com/databricks/databricks-sdk-go/pull/249)).
 * Added `pflag.Value` interfaces for enums ([#234](https://github.com/databricks/databricks-sdk-go/pull/234)).
 * Added support for adding custom HTTP visitors per request ([#230](https://github.com/databricks/databricks-sdk-go/pull/230)).
 * Added support for raw body as byte slice if requested ([#247](https://github.com/databricks/databricks-sdk-go/pull/247)).
 * Improved callbacks for polling the status of long-running operations ([#258](https://github.com/databricks/databricks-sdk-go/pull/258)).
 * Improved rendering of HTTP links in godoc ([#229](https://github.com/databricks/databricks-sdk-go/pull/229)).
 * Updated field types in the Jobs API from spec ([#259](https://github.com/databricks/databricks-sdk-go/pull/259)).
 * Multiple OpenAPI consistency passes ([#254](https://github.com/databricks/databricks-sdk-go/pull/254), [#241](https://github.com/databricks/databricks-sdk-go/pull/241), [#243](https://github.com/databricks/databricks-sdk-go/pull/243), [#255](https://github.com/databricks/databricks-sdk-go/pull/255), [#236](https://github.com/databricks/databricks-sdk-go/pull/236)).

API changes:
 * Renamed `IsAccountsClient` to `IsAccountClient` ([#231](https://github.com/databricks/databricks-sdk-go/pull/231)).
 * `w.ClusterPolicies.ListAll` now takes `clusterpolicies.List` as an argument.
 * `github.com/databricks/databricks-sdk-go/service/dbsql` package is renamed to `github.com/databricks/databricks-sdk-go/service/sql`.
 * `w.DataSources.ListDataSources` is renamed to `w.DataSources.List`.
 * `w.Queries.CreateQuery` is renamed to `w.Queries.CreateQuery`.
 * `w.Queries.DeleteQueryByQueryId` is renamed to `w.Queries.DeleteByQueryId`.
 * `w.Queries.GetQueryByQueryId` is renamed to `w.Queries.GetByQueryId`.
 * `w.Queries.UpdateQuery` is renamed to `w.Queries.Update`.
 * `w.Alerts.DeleteAlertByAlertId` is renamed to `w.Alerts.DeleteByAlertId`.
 * `w.Alerts.UpdateAlert` is renamed to `w.Alerts.Update`.
 * `w.Alerts.GetAlertByAlertId` is renamed to `w.Alerts.GetByAlertId`.
 * `w.Alerts.ListAlerts` is renamed to `w.Alerts.List`.
 * `w.Dashboards.CreateDashboard` is renamed to `w.Dashboards.Create`.
 * `w.Dashboards.DeleteDashboardByDashboardId` is renamed to `w.Dashboards.GetByDashboardId`.
 * `w.Dashboards.ListDashboardsAll` is renamed to `w.Dashboards.ListAll`.
 * `w.Dashboards.DeleteDashboardByDashboardId` is renamed to `w.Dashboards.DeleteByDashboardId`.
 * `w.Dashboards.RestoreDashboard` is renamed to `w.Dashboards.Restore`.
 * `deployment.CreateCustomerManagedKeyRequest` now takes `deployment.KeyUseCase` enum.
 * `w.GlobalInitScripts.CreateScript` is renamed to `w.GlobalInitScripts.Create`.
 * `w.GlobalInitScripts.DeleteScriptByScriptId` is renamed to `w.GlobalInitScripts.DeleteByScriptId`.
 * `w.GlobalInitScripts.UpdateScript` is renamed to `w.GlobalInitScripts.Update`.
 * `w.GlobalInitScripts.GetScriptByScriptId` is renamed to `w.GlobalInitScripts.GetByScriptId`.
 * `w.GlobalInitScripts.ListScriptsAll` is renamed to `w.GlobalInitScripts.ListAll`.
 * `jobs.ResetJob.NewSettings` is now required field.
 * `w.Pipelines.CreatePipeline` is renamed to `w.Pipelines.Create`.
 * `w.Pipelines.DeletePipelineByPipelineId` is renamed to `w.Pipelines.DeleteByPipelineId`.
 * `w.Pipelines.UpdatePipeline` is renamed to `w.Pipelines.Update`.
 * `w.Pipelines.GetPipelineByPipelineId` is renamed to `w.Pipelines.GetByPipelineId`.
 * `w.StorageCredentials.Update` now also returns an entity.
 * `w.ExternalLocations.Update` now also returns an entity.
 * `w.Metastores.Update` now also returns an entity.
 * `unitycatalog.CreateMetastoreAssignment.WorkspaceId` type changed from `int` to `int64`.
 * `unitycatalog.UnassignRequest.WorkspaceId` type changed from `int` to `int64`.
 * `w.Catalogs.Update` now also returns an entity.
 * `w.Schemas.Update` now also returns an entity.
 * `w.Providers.Update` now also returns an entity.
 * `w.Shares.Update` now also returns an entity.
 * `WarehousesAPI` service moved to `github.com/databricks/databricks-sdk-go/service/sql` package.
 * `w.Warehouses.CreateWarehouseAndWait` renamed to `w.Warehouses.CreateAndWait`.
 * `w.Warehouses.DeleteWarehouseByIdAndWait` renamed to `w.Warehouses.DeleteByIdAndWait`.
 * `w.Warehouses.EditWarehouse` renamed to `w.Warehouses.Edit`.
 * `w.Warehouses.GetWarehouseById` renamed to `w.Warehouses.GetById`.
 * `w.Warehouses.ListWarehousesAll` renamed to `w.Warehouses.ListAll`.
 * Removed `w.Dbfs.Overwrite` in favor of `w.Dbfs.Open("....", dbfs.FileModeOverwrite|dbfs.FileModeWrite)`.
 * Added third required argument to `w.Dbfs.Open`.

Code generation:

 * Added concept of `main` service for the package ([#239](https://github.com/databricks/databricks-sdk-go/pull/239)).
 * Added entity primitives check ([#242](https://github.com/databricks/databricks-sdk-go/pull/242)).
 * Added helpers for CRUD generation ([#246](https://github.com/databricks/databricks-sdk-go/pull/246)).
 * Added more entity-generation utils ([#257](https://github.com/databricks/databricks-sdk-go/pull/257)).
 * Dynamically generate `.gitattributes` ([#244](https://github.com/databricks/databricks-sdk-go/pull/244)).
 * Fixed required order fields ([#245](https://github.com/databricks/databricks-sdk-go/pull/245)).
 * Parse summary from descriptions ([#228](https://github.com/databricks/databricks-sdk-go/pull/228)).
 * Print error on formatter failure ([#235](https://github.com/databricks/databricks-sdk-go/pull/235)).
 * Update usage string in generator ([#260](https://github.com/databricks/databricks-sdk-go/pull/260)).
 * Fixed order of host completion ([#233](https://github.com/databricks/databricks-sdk-go/pull/233)).

Dependency updates:

 * Bump google.golang.org/api from 0.103.0 to 0.105.0 ([#232](https://github.com/databricks/databricks-sdk-go/pull/232), [#252](https://github.com/databricks/databricks-sdk-go/pull/252)).
 
## 0.1.1

* Improved usage instructions and added more documentation for services.

## 0.1.0

 * Initial release of the Databricks SDK for Go.

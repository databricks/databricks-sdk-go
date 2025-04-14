# NEXT CHANGELOG

## Release v0.63.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Added [w.EnableExportNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EnableExportNotebookAPI) workspace-level service, [w.EnableNotebookTableClipboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EnableNotebookTableClipboardAPI) workspace-level service and [w.EnableResultsDownloading](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EnableResultsDownloadingAPI) workspace-level service.
* Added `GetCredentialsForTraceDataDownload` and `GetCredentialsForTraceDataUpload` methods for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
* Added `GetDownloadFullQueryResult` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `GetPublishedDashboardTokenInfo` method for [w.LakeviewEmbedded](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewEmbeddedAPI) workspace-level service.
* Added `BindingWorkspaceIds` field for [billing.BudgetPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetPolicy).
* Added `DownloadId` field for [dashboards.GenieGenerateDownloadFullQueryResultResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieGenerateDownloadFullQueryResultResponse).
* Added `DashboardOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
* Added `DashboardTask` and `PowerBiTask` fields for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Added `DashboardTask` and `PowerBiTask` fields for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Added `DashboardTask` and `PowerBiTask` fields for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Added `IncludeFeatures` field for [ml.CreateForecastingExperimentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateForecastingExperimentRequest).
* Added `Models` field for [ml.LogInputs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#LogInputs).
* Added `DatasetDigest`, `DatasetName` and `ModelId` fields for [ml.LogMetric](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#LogMetric).
* Added `DatasetDigest`, `DatasetName`, `ModelId` and `RunId` fields for [ml.Metric](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Metric).
* Added `ModelInputs` field for [ml.RunInputs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RunInputs).
* Added `ClientApplication` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
* Added `Geography` and `Geometry` enum values for [catalog.ColumnTypeName](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ColumnTypeName).
* Added `AllocationTimeoutNoHealthyAndWarmedUpClusters`, `DockerContainerCreationException`, `DockerImageTooLargeForInstanceException` and `DockerInvalidOsException` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `Standard` enum value for [jobs.PerformanceTarget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PerformanceTarget).
* Added `CanView` enum value for [sql.WarehousePermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissionLevel).
* [Breaking] Changed `GenerateDownloadFullQueryResult` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service . Method path has changed.
* [Breaking] Changed waiter for [CommandExecutionAPI.Create](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CommandExecutionAPI.Create).
* [Breaking] Changed waiter for [CommandExecutionAPI.Execute](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CommandExecutionAPI.Execute).
* [Breaking] Removed `Error`, `Status` and `TransientStatementId` fields for [dashboards.GenieGenerateDownloadFullQueryResultResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieGenerateDownloadFullQueryResultResponse).
* [Breaking] Removed `Balanced` and `CostOptimized` enum values for [jobs.PerformanceTarget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PerformanceTarget).
* [Breaking] Removed [PipelinesAPI.WaitGetPipelineRunning](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI.WaitGetPipelineRunning) method.

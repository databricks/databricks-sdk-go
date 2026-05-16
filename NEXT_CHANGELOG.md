# NEXT CHANGELOG

## Release v0.133.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `CreateWorkspaceAssignmentDetail`, `DeleteWorkspaceAssignmentDetail`, `GetWorkspaceAssignmentDetail`, `ListWorkspaceAssignmentDetails` and `UpdateWorkspaceAssignmentDetail` methods for [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service.
* Add `CreateWorkspaceAssignmentDetailProxy`, `DeleteWorkspaceAssignmentDetailProxy`, `GetWorkspaceAssignmentDetailProxy`, `ListWorkspaceAssignmentDetailsProxy` and `UpdateWorkspaceAssignmentDetailProxy` methods for [w.WorkspaceIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API) workspace-level service.
* Add `FailoverGroupName` field for [disasterrecovery.StableUrl](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#StableUrl).
* Add `Disabled` and `PythonOperatorTask` fields for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `Disabled` and `PythonOperatorTask` fields for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Add `PythonOperatorTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Add `CatalogName`, `CreatedAt`, `CreatedBy`, `Name` and `SchemaName` fields for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* [Breaking] Add `CatalogName` and `SchemaName` fields for [ml.ListFeaturesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ListFeaturesRequest).
* Add `Rolling` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `KafkaOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `CrossWorkspaceAccess` field for [settings.CustomerFacingIngressNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingIngressNetworkPolicy).
* Add `AllowedAppsUserApiScopes` and `EffectiveAllowedAppsUserApiScopes` fields for [settingsv2.Setting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#Setting).
* Add `Hubspot`, `Github`, `Outlook` and `Smartsheet` enum values for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `GpuXlarge` enum value for [serving.ServedModelInputWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadType).
* Add `GpuXlarge` enum value for [serving.ServingModelWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingModelWorkloadType).
* [Breaking] Change `ListFeatures` method for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service with new required argument order.
* [Breaking] Remove `UnspecifiedResourceName` field for [postgres.RequestedResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RequestedResource).
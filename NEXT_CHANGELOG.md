# NEXT CHANGELOG

## Release v0.161.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `PerformanceTarget` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
* Add `DataframeSchema` and `TransformationSql` fields for [ml.StreamSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#StreamSource).
* Add `Sawtooth` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `EffectiveTablePrefix` field for [ml.UcTraceLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#UcTraceLocation).
* Add `AwsVpcEndpointInfo` and `GcpPscEndpointInfo` fields for [networking.Endpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/networking#Endpoint).
* Add `RedditAdsOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `TableNames` and `TelemetryProfileId` fields for [serving.TelemetryConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#TelemetryConfig).
* Add `GitCliEnabled` field for [workspace.GetRepoResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetRepoResponse).
* Add `KsaEccCccDcc` enum value for [settings.ComplianceStandard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ComplianceStandard).
* [Breaking] Change `CreateWorkspaceAssignmentDetail`, `DeleteWorkspaceAssignmentDetail`, `GetWorkspaceAccessDetail`, `GetWorkspaceAssignmentDetail`, `ListWorkspaceAssignmentDetails`, `ResolveGroup`, `ResolveServicePrincipal`, `ResolveUser` and `UpdateWorkspaceAssignmentDetail` methods for [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service. Method path has changed.
* [Breaking] Change `CreateWorkspaceAssignmentDetailProxy`, `DeleteWorkspaceAssignmentDetailProxy`, `GetWorkspaceAccessDetailLocal`, `GetWorkspaceAssignmentDetailProxy`, `ListWorkspaceAssignmentDetailsProxy`, `ResolveGroupProxy`, `ResolveServicePrincipalProxy`, `ResolveUserProxy` and `UpdateWorkspaceAssignmentDetailProxy` methods for [w.WorkspaceIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API) workspace-level service. Method path has changed.
* [Breaking] Change `WindowDuration` field for [ml.RollingWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RollingWindow) to no longer be required.
* Change `WindowDuration` field for [ml.RollingWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RollingWindow) to no longer be required.
* Change `WindowDuration` field for [ml.SlidingWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SlidingWindow) to no longer be required.
* [Breaking] Change `WindowDuration` field for [ml.SlidingWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SlidingWindow) to no longer be required.
* [Breaking] Remove `LongRolling` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
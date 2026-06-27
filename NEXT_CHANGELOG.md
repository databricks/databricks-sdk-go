# NEXT CHANGELOG

## Release v0.152.0

### Breaking Changes

* Query parameters listed in `ForceSendFields` are now sent on the wire when they hold a zero value. Previously `ForceSendFields` had no effect on query parameters (only JSON body fields were honored), so such a parameter was silently omitted; it is now serialized with its explicit value (for example `cascade=false`). Callers that unknowingly relied on the previous no-op behavior will now send the parameter.

### New Features and Improvements

### Bug Fixes

* Honor `ForceSendFields` for query parameters, including fields nested inside request sub-structs. A zero-valued scalar field (for example a `false` bool such as `cascade` on the pipelines Delete request) listed in `ForceSendFields` is now sent on the wire instead of being dropped by the `url` tag's `omitempty`, matching the existing behavior for JSON request bodies.
* Parse WaitTimeout and use it as poll timeout in ExecuteAndWait

### Documentation

### Internal Changes

### API Changes
* Add `Etag` field for [disasterrecovery.UpdateFailoverGroupRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#UpdateFailoverGroupRequest).
* Add `DownloadMessageAttachmentVisualization` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Add `Viz` field for [dashboards.GenieAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAttachment).
* Add `EnableVisualization` field for [dashboards.GenieCreateConversationMessageRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieCreateConversationMessageRequest).
* Add `EnableVisualization` field for [dashboards.GenieStartConversationMessageRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieStartConversationMessageRequest).
* Add `ReplaceExisting` field for [postgres.CreateDatabaseRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateDatabaseRequest).
* Add `ReplaceExisting` field for [postgres.CreateRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateRoleRequest).
* Add `AutoscalingLimitMaxCu`, `AutoscalingLimitMinCu`, `NoSuspension` and `SuspendTimeoutDuration` fields for [postgres.InitialEndpointSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#InitialEndpointSpec).
* Add `InitialBranchSpec` field for [postgres.Project](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Project).
* Add `ComputeLastActiveTime` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* Add `TelemetryConfig` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
* Add `TelemetryConfig` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
* Add `TelemetryConfig` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
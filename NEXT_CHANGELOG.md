# NEXT CHANGELOG

## Release v0.89.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `InstanceProfileArn` field for [compute.InstancePoolAwsAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAwsAttributes).
* Add `Continuous`, `Sliding` and `Tumbling` fields for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `UsagePolicyId` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Add `UsagePolicyId` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Add `UsagePolicyId` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Add `ReadFilesBytes` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
* Add `Select` enum value for [apps.AppManifestAppResourceUcSecurableSpecUcSecurablePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppManifestAppResourceUcSecurableSpecUcSecurablePermission).
* Add `Table` enum value for [apps.AppManifestAppResourceUcSecurableSpecUcSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppManifestAppResourceUcSecurableSpecUcSecurableType).
* Add `DecommissionStarted` and `DecommissionEnded` enum values for [compute.EventType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventType).
* Add `DbrImageResolutionFailure` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `DbrImageResolutionFailure` enum value for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).
* Change `OfflineStoreConfig` and `OnlineStoreConfig` fields for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature) to no longer be required.
* [Breaking] Change `OfflineStoreConfig` and `OnlineStoreConfig` fields for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature) to no longer be required.
* [Breaking] Change `LifecycleState` field for [sql.AlertV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2) to type [sql.AlertLifecycleState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertLifecycleState).
* [Breaking] Remove `Table` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings).
* [Breaking] Remove `Duration` and `Offset` fields for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
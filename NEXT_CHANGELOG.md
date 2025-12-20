# NEXT CHANGELOG

## Release v0.96.0

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `GitRepository` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `GitSource` field for [apps.AppDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppDeployment).
* Add `ExperimentSpec` field for [apps.AppManifestAppResourceSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppManifestAppResourceSpec).
* Add `Experiment` field for [apps.AppResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResource).
* Add `GitRepository` field for [apps.AppUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppUpdate).
* Add `ExcludedTableFullNames` field for [dataquality.AnomalyDetectionConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataquality#AnomalyDetectionConfig).
* Add `GroupName` field for [jobs.JobRunAs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobRunAs).
* Add `RowFilter` field for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Add `ExcludedTableFullNames` field for [qualitymonitorv2.AnomalyDetectionConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/qualitymonitorv2#AnomalyDetectionConfig).
* Add `Execute` and `UseConnection` enum values for [apps.AppManifestAppResourceUcSecurableSpecUcSecurablePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppManifestAppResourceUcSecurableSpecUcSecurablePermission).
* Add `Function` and `Connection` enum values for [apps.AppManifestAppResourceUcSecurableSpecUcSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppManifestAppResourceUcSecurableSpecUcSecurableType).
* Add `Select`, `Execute` and `UseConnection` enum values for [apps.AppResourceUcSecurableUcSecurablePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResourceUcSecurableUcSecurablePermission).
* Add `Table`, `Function` and `Connection` enum values for [apps.AppResourceUcSecurableUcSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResourceUcSecurableUcSecurableType).
* Add `Spec` and `Status` fields for [postgres.Endpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Endpoint).
* [Breaking] Remove `AutoscalingLimitMaxCu`, `AutoscalingLimitMinCu`, `CurrentState`, `Disabled`, `EffectiveAutoscalingLimitMaxCu`, `EffectiveAutoscalingLimitMinCu`, `EffectiveDisabled`, `EffectivePoolerMode`, `EffectiveSettings`, `EffectiveSuspendTimeoutDuration`, `EndpointType`, `Host`, `LastActiveTime`, `PendingState`, `PoolerMode`, `Settings`, `StartTime`, `SuspendTime` and `SuspendTimeoutDuration` fields for [postgres.Endpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Endpoint).
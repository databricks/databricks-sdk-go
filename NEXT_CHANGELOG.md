# NEXT CHANGELOG

## Release v0.104.0

### Breaking Changes

### New Features and Improvements
* Add support for single Profile for Account and Workspace operations in Unified Mode.

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `BackfillSource` field for [ml.KafkaConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaConfig).
* Add `BurstScalingEnabled` field for [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
* Add `BurstScalingEnabled` field for [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).
* Add `BurstScalingEnabled` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
* Add `BurstScalingEnabled` field for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput).
* Add `GenerateDownloadFullQueryResult` and `GetDownloadFullQueryResult` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Add `ActiveInstances` field for [apps.ComputeStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#ComputeStatus).
* [Breaking] Change `CreateRole` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service. HTTP method/verb has changed.
* Add `Compute` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `Compute` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Add `Compute` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Add `MtlsPortConnectivityFailure` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `MtlsPortConnectivityFailure` enum value for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).
* Add `BaseEnvironment` field for [compute.Environment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Environment).
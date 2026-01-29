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
* [Breaking] Change `CreateRole` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service. HTTP method/verb has changed.
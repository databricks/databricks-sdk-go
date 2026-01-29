# NEXT CHANGELOG

## Release v0.104.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Fix potential panic in `shouldRetry` function by using comma-ok idiom for type assertion ([#1438](https://github.com/databricks/databricks-sdk-go/pull/1438)).

### Documentation

### Internal Changes
* Remove deprecated `rand.Seed` call in retry backoff logic ([#1438](https://github.com/databricks/databricks-sdk-go/pull/1438)).

### API Changes
* Add `BackfillSource` field for [ml.KafkaConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaConfig).
* Add `BurstScalingEnabled` field for [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
* Add `BurstScalingEnabled` field for [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).
* Add `BurstScalingEnabled` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
* Add `BurstScalingEnabled` field for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput).
# NEXT CHANGELOG

## Release v0.147.0

### Breaking Changes

### New Features and Improvements

- Upload of big files (> 5Gb) to UC Volumes using multipart chunking ([#621](https://github.com/databricks/databricks-sdk-go/pull/1621)).

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `ServerlessComputeId` field for [pipelines.ClonePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ClonePipelineRequest).
* Add `ServerlessComputeId` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Add `ServerlessComputeId` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Add `ServerlessComputeId` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).

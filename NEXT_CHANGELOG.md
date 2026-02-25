# NEXT CHANGELOG

## Release v0.114.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add `EffectivePublishingMode` field for [pipelines.GetPipelineResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelineResponse).
* Add `DbrAutoscale` enum value for [compute.EventDetailsCause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventDetailsCause).
* Change `OutputCatalog` field for [cleanrooms.CreateCleanRoomOutputCatalogResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CreateCleanRoomOutputCatalogResponse) to be required.
* [Breaking] Remove `InternalAttributes` field for [sharing.Table](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#Table).
* [Breaking] Remove `InternalAttributes` field for [sharing.Volume](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#Volume).
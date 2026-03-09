# NEXT CHANGELOG

## Release v0.119.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Internal Changes

### API Changes
* Add [dataclassification](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataclassification) and [knowledgeassistants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants) packages.
* Add [w.DataClassification](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataclassification#DataClassificationAPI) workspace-level service.
* Add [w.KnowledgeAssistants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#KnowledgeAssistantsAPI) workspace-level service.
* Add `GenieCreateEvalRun`, `GenieGetEvalResultDetails`, `GenieGetEvalRun`, `GenieListEvalResults` and `GenieListEvalRuns` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Add `TelemetryExportDestinations` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
# NEXT CHANGELOG

## Release v0.111.0

### Breaking Changes

### New Features and Improvements

* Remove cloud type restrictions from Azure/GCP credential providers. Azure and GCP authentication now works with any Databricks host when credentials are properly configured, enabling authentication against cloud-agnostic endpoints such as aliased hosts.

### Bug Fixes

### Documentation

### Internal Changes

* Make clusters creation in AWS not depend on cloud parsed from host. Changed default AWS availability for auto-created utility clusters from SPOT to SPOT_WITH_FALLBACK (API default). 

### API Changes

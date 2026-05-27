# NEXT CHANGELOG

## Release v0.139.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Extract `?o=`/`?workspace_id=` and `?a=`/`?account_id=` from `Config.Host` into `Config.WorkspaceID`/`Config.AccountID` during host sanitization. Pasting a SPOG URL from the Databricks UI (e.g. `https://acme.databricks.net/?o=12345`) previously dropped the workspace ID, causing API calls to hit the SPOG without an `X-Databricks-Org-Id` header and return HTML.

### Documentation

### Internal Changes

### API Changes

# NEXT CHANGELOG

## Release v0.138.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* `config`: The config-file loader now sets `cfg.Profile` to `"DEFAULT"` on the legacy fallback (no profile requested, no `[__settings__].default_profile` set, `[DEFAULT]` section loaded). Previously the loader silently used `[DEFAULT]`'s values but left `cfg.Profile` empty. Consumers that derive a per-profile identifier from `cfg.Profile` (e.g. the Databricks CLI's U2M OAuth cache key) need the resolved name to match across the login and read flows; without this fix a write under `"DEFAULT"` could not be found by a no-flag read.

### Documentation

### Internal Changes

### API Changes

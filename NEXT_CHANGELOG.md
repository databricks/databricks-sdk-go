# NEXT CHANGELOG

## Release v0.81.0

### New Features and Improvements

### Bug Fixes

- Update the `U2M` port selection mechanism to try port `8020` first and fall
  back incrementally (to port `8021`, `8022`, and so on...) if a port is not 
  available. This fixes an issue with the Databricks CLI which is only 
  allowlisted on port `8020`. 

### Documentation

### Internal Changes

### API Changes

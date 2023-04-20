package compute

// make common/resource.go#ToResource read behavior consistent with "normal" resources
// func wrapMissingClusterError(err error, id string) error {
// 	if err == nil {
// 		return nil
// 	}
// 	apiErr, ok := err.(apierr.APIError)
// 	if !ok {
// 		return err
// 	}
// 	if apiErr.IsMissing() {
// 		return err
// 	}
// 	// https://github.com/databricks/terraform-provider-databricks/issues/1177
// 	// Aligned with Clusters Core team to keep behavior of these workarounds
// 	// as is in the longer term, so that this keeps working.
// 	if apiErr.ErrorCode == "INVALID_STATE" {
// 		logger.Warnf(ctx, "assuming that cluster is removed on backend: %s", apiErr)
// 		apiErr.StatusCode = 404
// 		return apiErr
// 	}
// 	// fix non-compliant error code
// 	if strings.Contains(apiErr.Message,
// 		fmt.Sprintf("Cluster %s does not exist", id)) {
// 		apiErr.StatusCode = 404
// 		return apiErr
// 	}
// 	return err
// }

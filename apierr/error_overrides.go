// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apierr

import (
	"regexp"
)

var allOverrides = []errorOverride{
	{
		debugName:         "Clusters InvalidParameterValue=>ResourceDoesNotExist",
		pathRegex:         regexp.MustCompile(`^/api/2\.\d/clusters/get`),
		verb:              "GET",
		statusCodeMatcher: regexp.MustCompile(`^400$`),
		errorCodeMatcher:  regexp.MustCompile(`INVALID_PARAMETER_VALUE`),
		messageMatcher:    regexp.MustCompile(`Cluster .* does not exist`),
		customError:       ErrResourceDoesNotExist,
	}, {
		debugName:         "Jobs InvalidParameterValue=>ResourceDoesNotExist",
		pathRegex:         regexp.MustCompile(`^/api/2\.\d/jobs/get`),
		verb:              "GET",
		statusCodeMatcher: regexp.MustCompile(`^400$`),
		errorCodeMatcher:  regexp.MustCompile(`INVALID_PARAMETER_VALUE`),
		messageMatcher:    regexp.MustCompile(`Job .* does not exist`),
		customError:       ErrResourceDoesNotExist,
	},
}

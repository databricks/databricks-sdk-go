package openapi

import "regexp"

type ErrorMappingRule struct {
	StatusCode  int    `json:"status_code"`
	ErrorCode   string `json:"error_code"`
	Description string `json:"description"`
}

// API error rules for the Databricks REST API based on HTTP status codes.
//
// The API responds with an error by setting the HTTP status code to 400 or
// higher and returning a JSON object with an error code and message. Error
// responses are hierarchical in nature, with the HTTP status code providing
// the highest level of categorization. The error_code field in the response
// provides a more granular categorization of the error.
//
// The SDK attempts to map these error responses to error objects based on this
// mapping by applying the following logic:
//
//  1. If the error matches an override defined in ErrorOverrides, return the
//     override.
//  2. If the error_code field of the response matches a key in
//     ErrorCodeMapping, return that error.
//  3. If the HTTP status code matches a rule in ErrorStatusCodeMapping, return
//     that error.
var ErrorStatusCodeMapping = []ErrorMappingRule{
	{400, "BAD_REQUEST", "the request is invalid"},
	{401, "UNAUTHENTICATED", "the request does not have valid authentication (AuthN) credentials for the operation"},
	{403, "PERMISSION_DENIED", "the caller does not have permission to execute the specified operation"},
	{404, "NOT_FOUND", "the operation was performed on a resource that does not exist"},
	{409, "RESOURCE_CONFLICT", "maps to all HTTP 409 (Conflict) responses"},
	{429, "TOO_MANY_REQUESTS", "maps to HTTP code: 429 Too Many Requests"},
	{499, "CANCELLED", "the operation was explicitly canceled by the caller"},
	{500, "INTERNAL_ERROR", "some invariants expected by the underlying system have been broken"},
	{501, "NOT_IMPLEMENTED", "the operation is not implemented or is not supported/enabled in this service"},
	{503, "TEMPORARILY_UNAVAILABLE", "the service is currently unavailable"},
	{504, "DEADLINE_EXCEEDED", "the deadline expired before the operation could complete"},
}

// Specialized  API error rules based on the error_code field in the error
// response.
//
// See the documentation on ErrorStatusCodeMapping for more information.
var ErrorCodeMapping = []ErrorMappingRule{
	{400, "INVALID_STATE", "unexpected state"},
	{400, INVALID_PARAMETER_VALUE, "supplied value for a parameter was invalid"},
	{404, RESOURCE_DOES_NOT_EXIST, "operation was performed on a resource that does not exist"},
	{409, "ABORTED", "the operation was aborted, typically due to a concurrency issue such as a sequencer check failure"},
	{409, "ALREADY_EXISTS", "operation was rejected due a conflict with an existing resource"},
	{409, "RESOURCE_ALREADY_EXISTS", "operation was rejected due a conflict with an existing resource"},
	{429, "RESOURCE_EXHAUSTED", "operation is rejected due to per-user rate limiting"},
	{429, "REQUEST_LIMIT_EXCEEDED", "cluster request was rejected because it would exceed a resource limit"},
	{500, "UNKNOWN", "this error is used as a fallback if the platform-side mapping is missing some reason"},
	{500, "DATA_LOSS", "unrecoverable data loss or corruption"},
}

type ErrorOverride struct {
	Name              string
	PathRegex         *regexp.Regexp
	Verb              string
	StatusCodeMatcher *regexp.Regexp
	ErrorCodeMatcher  *regexp.Regexp
	MessageMatcher    *regexp.Regexp
	OverrideErrorCode string
}

const INVALID_PARAMETER_VALUE = "INVALID_PARAMETER_VALUE"
const RESOURCE_DOES_NOT_EXIST = "RESOURCE_DOES_NOT_EXIST"

// Overrides in the SDK to remap specific error responses to other errors.
//
// When certain APIs return a non-standard or inconsistent error response, the
// SDK can be configured to remap these errors to a more appropriate error.
// For example, some APIs return a 400 error when a resource doesn't exist. In
// this case, the SDK can be configured to remap the error to a 404 error.
var ErrorOverrides = []ErrorOverride{
	{
		Name:              "Clusters InvalidParameterValue=>ResourceDoesNotExist",
		PathRegex:         regexp.MustCompile(`^/api/2\.\d/clusters/get`),
		Verb:              "GET",
		StatusCodeMatcher: regexp.MustCompile(`^400$`),
		MessageMatcher:    regexp.MustCompile("Cluster .* does not exist"),
		ErrorCodeMatcher:  regexp.MustCompile(INVALID_PARAMETER_VALUE),
		OverrideErrorCode: RESOURCE_DOES_NOT_EXIST,
	},
	{
		Name:              "Jobs InvalidParameterValue=>ResourceDoesNotExist",
		PathRegex:         regexp.MustCompile(`^/api/2\.\d/jobs/get`),
		Verb:              "GET",
		StatusCodeMatcher: regexp.MustCompile(`^400$`),
		MessageMatcher:    regexp.MustCompile("Job .* does not exist"),
		ErrorCodeMatcher:  regexp.MustCompile(INVALID_PARAMETER_VALUE),
		OverrideErrorCode: RESOURCE_DOES_NOT_EXIST,
	},
	{
		Name:              "Job Runs InvalidParameterValue=>ResourceDoesNotExist",
		PathRegex:         regexp.MustCompile(`^/api/2\.\d/jobs/runs/get`),
		Verb:              "GET",
		StatusCodeMatcher: regexp.MustCompile(`^400$`),
		MessageMatcher:    regexp.MustCompile("(Run .* does not exist|Run: .* in job: .* does not exist)"),
		ErrorCodeMatcher:  regexp.MustCompile(INVALID_PARAMETER_VALUE),
		OverrideErrorCode: RESOURCE_DOES_NOT_EXIST,
	},
}

var TransientErrorRegexes = []*regexp.Regexp{
	regexp.MustCompile(`com\.databricks\.backend\.manager\.util\.UnknownWorkerEnvironmentException`),
	regexp.MustCompile(`does not have any associated worker environments`),
	regexp.MustCompile(`There is no worker environment with id`),
	regexp.MustCompile(`Unknown worker environment`),
	regexp.MustCompile(`ClusterNotReadyException`),
	regexp.MustCompile(`worker env .* not found`),
	regexp.MustCompile(`Timed out after `),
	regexp.MustCompile(`deadline exceeded`),
}

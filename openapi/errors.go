package openapi

type ErrorMappingRule struct {
	StatusCode  int    `json:"status_code"`
	ErrorCode   string `json:"error_code"`
	Description string `json:"description"`
}

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

var ErrorCodeMapping = []ErrorMappingRule{
	{400, "INVALID_PARAMETER_VALUE", "supplied value for a parameter was invalid"},
	{409, "ABORTED", "the operation was aborted, typically due to a concurrency issue such as a sequencer check failure"},
	{409, "ALREADY_EXISTS", "operation was rejected due a conflict with an existing resource"},
	{409, "RESOURCE_ALREADY_EXISTS", "operation was rejected due a conflict with an existing resource"},
	{429, "RESOURCE_EXHAUSTED", "operation is rejected due to per-user rate limiting"},
	{429, "REQUEST_LIMIT_EXCEEDED", "cluster request was rejected because it would exceed a resource limit"},
	{500, "UNKNOWN", "this error is used as a fallback if the platform-side mapping is missing some reason"},
	{500, "DATA_LOSS", "unrecoverable data loss or corruption"},
}

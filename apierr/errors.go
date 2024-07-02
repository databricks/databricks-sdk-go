package apierr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/logger/httplog"
)

const (
	errorInfoType string = "type.googleapis.com/google.rpc.ErrorInfo"
)

type ErrorDetail struct {
	Type     string            `json:"@type,omitempty"`
	Reason   string            `json:"reason,omitempty"`
	Domain   string            `json:"domain,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// APIError is a generic struct for an api error on databricks
type APIError struct {
	ErrorCode  string
	Message    string
	StatusCode int
	Details    []ErrorDetail
	// If non-nil, the underlying error that should be returned by calling
	// errors.Unwrap on this error.
	unwrap error
}

// Error returns error message string instead of
func (apiError *APIError) Error() string {
	return apiError.Message
}

// IsMissing tells if error is about missing resource
func IsMissing(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// GetErrorInfo returns all entries in the list of error details of type `ErrorInfo`.
func GetErrorInfo(err error) []ErrorDetail {
	return getDetailsByType(err, errorInfoType)
}

func getDetailsByType(err error, errorDetailType string) []ErrorDetail {
	var apiError *APIError
	if !errors.As(err, &apiError) {
		return nil
	}
	filteredDetails := []ErrorDetail{}
	for _, detail := range apiError.Details {
		if errorDetailType == detail.Type {
			filteredDetails = append(filteredDetails, detail)
		}
	}
	return filteredDetails
}

// IsMissing tells if it is missing resource
func (apiError *APIError) IsMissing() bool {
	return errors.Is(apiError, ErrNotFound)
}

// IsTooManyRequests shows rate exceeded limits
func (apiError *APIError) IsTooManyRequests() bool {
	return errors.Is(apiError, ErrTooManyRequests)
}

// isRetriable returns true if error is retriable
func (apiError *APIError) IsRetriable(ctx context.Context) bool {
	if apiError.IsTooManyRequests() {
		return true
	}
	// Retry when the API is unavailable. This includes "No webapps are available
	// to handle your request. Please try again later." from the API gateway.
	// Reference: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503
	if apiError.StatusCode == http.StatusServiceUnavailable {
		return true
	}
	// Handle transient errors for retries
	for _, regex := range allTransientErrors {
		if regex.Match([]byte(apiError.Message)) {
			logger.Debugf(ctx, "Attempting retry because of %s", regex.String())
			return true
		}
	}
	// some API's recommend retries on HTTP 500, but we'll add that later
	return false
}

// NotFound returns properly formatted Not Found error
func NotFound(message string) *APIError {
	return &APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: 404,
		Message:    message,
	}
}

func ReadError(statusCode int, err error) *APIError {
	return &APIError{
		ErrorCode:  "IO_READ",
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

func TooManyRequests() *APIError {
	return &APIError{
		ErrorCode:  "TOO_MANY_REQUESTS",
		StatusCode: 429,
		Message:    "Current request has to be retried",
	}
}

func GenericIOError(ue *url.Error) *APIError {
	return &APIError{
		ErrorCode:  "IO_ERROR",
		StatusCode: 523,
		Message:    ue.Error(),
	}
}

// GetAPIError inspects HTTP errors from the Databricks API for known transient errors.
func GetAPIError(ctx context.Context, resp common.ResponseWrapper) error {
	if resp.Response.StatusCode == 429 {
		return TooManyRequests()
	}
	if resp.Response.StatusCode >= 400 {
		// read in response body as it is actually an error
		responseBodyBytes, err := io.ReadAll(resp.ReadCloser)
		if err != nil {
			return ReadError(resp.Response.StatusCode, err)
		}
		apiError := parseErrorFromResponse(resp.Response, resp.RequestBody.DebugBytes, responseBodyBytes)
		applyOverrides(ctx, apiError, resp.Response)
		return apiError
	}
	// Attempts to access private link workspaces are redirected to the login page with a specific query parameter.
	requestUrl := resp.Response.Request.URL
	if isPrivateLinkRedirect(requestUrl) {
		return privateLinkValidationError(requestUrl)
	}
	return nil
}

func parseErrorFromResponse(resp *http.Response, requestBody, responseBody []byte) *APIError {
	if len(responseBody) == 0 {
		return &APIError{StatusCode: resp.StatusCode}
	}

	// Anonymous struct used to unmarshal JSON Databricks API error responses.
	var errorBody struct {
		ErrorCode any           `json:"error_code,omitempty"` // int or string
		Message   string        `json:"message,omitempty"`
		Details   []ErrorDetail `json:"details,omitempty"`

		API12Error string `json:"error,omitempty"`

		// The following fields are for scim api only. See RFC7644 section 3.7.3
		// https://tools.ietf.org/html/rfc7644#section-3.7.3
		ScimDetail string `json:"detail,omitempty"`
		ScimStatus string `json:"status,omitempty"`
		ScimType   string `json:"scimType,omitempty"`
	}
	if err := json.Unmarshal(responseBody, &errorBody); err != nil {
		return unknownAPIError(resp, requestBody, responseBody, err)
	}

	// Convert API 1.2 error (which used a different format) to the new format.
	if errorBody.API12Error != "" {
		errorBody.Message = errorBody.API12Error
	}

	// Handle SCIM error message details.
	if errorBody.Message == "" && errorBody.ScimDetail != "" {
		if errorBody.ScimDetail == "null" {
			errorBody.Message = "SCIM API Internal Error"
		} else {
			errorBody.Message = errorBody.ScimDetail
		}
		// Add more context from SCIM responses.
		errorBody.Message = fmt.Sprintf("%s %s", errorBody.ScimType, errorBody.Message)
		errorBody.Message = strings.Trim(errorBody.Message, " ")
		errorBody.ErrorCode = fmt.Sprintf("SCIM_%s", errorBody.ScimStatus)
	}

	return &APIError{
		Message:    errorBody.Message,
		ErrorCode:  fmt.Sprintf("%v", errorBody.ErrorCode),
		StatusCode: resp.StatusCode,
		Details:    errorBody.Details,
	}
}

func unknownAPIError(resp *http.Response, requestBody, responseBody []byte, err error) *APIError {
	apiErr := &APIError{
		StatusCode: resp.StatusCode,
	}

	// This is most likely HTML... since un-marshalling JSON failed
	// Status parts first in case html message is not as expected
	statusParts := strings.SplitN(resp.Status, " ", 2)
	if len(statusParts) < 2 {
		apiErr.ErrorCode = "UNKNOWN"
	} else {
		apiErr.ErrorCode = strings.ReplaceAll(strings.ToUpper(strings.Trim(statusParts[1], " .")), " ", "_")
	}

	stringBody := string(responseBody)
	messageRE := regexp.MustCompile(`<pre>(.*)</pre>`)
	messageMatches := messageRE.FindStringSubmatch(stringBody)
	// No messages with <pre> </pre> format found so return a APIError
	if len(messageMatches) < 2 {
		apiErr.Message = MakeUnexpectedError(resp, err, requestBody, responseBody).Error()
	} else {
		apiErr.Message = strings.Trim(messageMatches[1], " .")
	}

	return apiErr
}

func MakeUnexpectedError(resp *http.Response, err error, requestBody, responseBody []byte) error {
	var req *http.Request
	if resp != nil {
		req = resp.Request
	}
	rts := httplog.RoundTripStringer{
		Request:                  req,
		Response:                 resp,
		Err:                      err,
		RequestBody:              requestBody,
		ResponseBody:             responseBody,
		DebugHeaders:             true,
		DebugTruncateBytes:       10 * 1024,
		DebugAuthorizationHeader: false,
	}
	return fmt.Errorf("unexpected error handling request: %w. This is likely a bug in the Databricks SDK for Go or the underlying REST API. Please report this issue with the following debugging information to the SDK issue tracker at https://github.com/databricks/databricks-sdk-go/issues. Request log:\n```\n%s\n```", err, rts.String())
}

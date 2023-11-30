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

	"github.com/databricks/databricks-sdk-go/logger"
)

var (
	transientErrorStringMatches = []string{
		"com.databricks.backend.manager.util.UnknownWorkerEnvironmentException",
		"does not have any associated worker environments",
		"There is no worker environment with id",
		"Unknown worker environment",
		"ClusterNotReadyException",
	}
)

const (
	errorInfoType string = "type.googleapis.com/google.rpc.ErrorInfo"
)

// APIErrorBody maps "proper" databricks rest api errors to a struct
type APIErrorBody struct {
	ErrorCode string        `json:"error_code,omitempty"`
	Message   string        `json:"message,omitempty"`
	Details   []ErrorDetail `json:"details,omitempty"`
	// The following two are for scim api only
	// for RFC 7644 Section 3.7.3 https://tools.ietf.org/html/rfc7644#section-3.7.3
	ScimDetail string `json:"detail,omitempty"`
	ScimStatus string `json:"status,omitempty"`
	ScimType   string `json:"scimType,omitempty"`
	API12Error string `json:"error,omitempty"`
}

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
}

// Error returns error message string instead of
func (apiError *APIError) Error() string {
	return apiError.Message
}

// IsMissing tells if error is about missing resource
func IsMissing(err error) bool {
	var apiError *APIError
	if errors.As(err, &apiError) {
		return apiError.IsMissing()
	}
	return false
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
	return apiError.StatusCode == http.StatusNotFound
}

// IsTooManyRequests shows rate exceeded limits
func (apiError *APIError) IsTooManyRequests() bool {
	return apiError.StatusCode == http.StatusTooManyRequests
}

// isRetriable returns true if error is retriable
func (apiError *APIError) IsRetriable(ctx context.Context) bool {
	if apiError.IsTooManyRequests() {
		return true
	}
	// Handle transient errors for retries
	for _, substring := range transientErrorStringMatches {
		if strings.Contains(apiError.Message, substring) {
			logger.Debugf(ctx, "Attempting retry because of %#v", substring)
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
func GetAPIError(ctx context.Context, resp *http.Response, body io.ReadCloser) error {
	if resp.StatusCode == 429 {
		return TooManyRequests()
	}
	if resp.StatusCode >= 400 {
		// read in response body as it is actually an error
		bodyBytes, err := io.ReadAll(body)
		if err != nil {
			return ReadError(resp.StatusCode, err)
		}
		apiError := parseErrorFromResponse(resp, bodyBytes)
		return apiError
	}
	return nil
}

func parseErrorFromResponse(resp *http.Response, body []byte) *APIError {
	// try to read in nicely formatted API error response
	var errorBody APIErrorBody
	err := json.Unmarshal(body, &errorBody)
	if err != nil {
		errorBody = parseUnknownError(resp.Status, body, err)
	}
	if errorBody.API12Error != "" {
		// API 1.2 has different response format, let's adapt
		errorBody.Message = errorBody.API12Error
	}
	// Handle SCIM error message details
	if errorBody.Message == "" && errorBody.ScimDetail != "" {
		if errorBody.ScimDetail == "null" {
			errorBody.Message = "SCIM API Internal Error"
		} else {
			errorBody.Message = errorBody.ScimDetail
		}
		// add more context from SCIM responses
		errorBody.Message = fmt.Sprintf("%s %s", errorBody.ScimType, errorBody.Message)
		errorBody.Message = strings.Trim(errorBody.Message, " ")
		errorBody.ErrorCode = fmt.Sprintf("SCIM_%s", errorBody.ScimStatus)
	}
	return &APIError{
		Message:    errorBody.Message,
		ErrorCode:  errorBody.ErrorCode,
		StatusCode: resp.StatusCode,
		Details:    errorBody.Details,
	}
}

func parseUnknownError(status string, body []byte, err error) (errorBody APIErrorBody) {
	// this is most likely HTML... since un-marshalling JSON failed
	// Status parts first in case html message is not as expected
	statusParts := strings.SplitN(status, " ", 2)
	if len(statusParts) < 2 {
		errorBody.ErrorCode = "UNKNOWN"
	} else {
		errorBody.ErrorCode = strings.ReplaceAll(
			strings.ToUpper(strings.Trim(statusParts[1], " .")),
			" ", "_")
	}
	stringBody := string(body)
	messageRE := regexp.MustCompile(`<pre>(.*)</pre>`)
	messageMatches := messageRE.FindStringSubmatch(stringBody)
	// No messages with <pre> </pre> format found so return a APIError
	if len(messageMatches) < 2 {
		errorBody.Message = fmt.Sprintf("Response from server (%s) %s: %v",
			status, stringBody, err)
		return
	}
	errorBody.Message = strings.Trim(messageMatches[1], " .")
	return
}

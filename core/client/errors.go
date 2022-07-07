package client

import (
	"log"
	"net/http"
	"strings"
)

var (
	transientErrorStringMatches = []string{
		"com.databricks.backend.manager.util.UnknownWorkerEnvironmentException",
		"does not have any associated worker environments",
		"There is no worker environment with id",
		"Unknown worker environment",
		"ClusterNotReadyException",
		"connection reset by peer",
		"TLS handshake timeout",
		"connection refused",
		"Unexpected error",
		"i/o timeout",
	}
)

// APIErrorBody maps "proper" databricks rest api errors to a struct
type APIErrorBody struct {
	ErrorCode string `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
	// The following two are for scim api only
	// for RFC 7644 Section 3.7.3 https://tools.ietf.org/html/rfc7644#section-3.7.3
	ScimDetail string `json:"detail,omitempty"`
	ScimStatus string `json:"status,omitempty"`
	ScimType   string `json:"scimType,omitempty"`
	API12Error string `json:"error,omitempty"`
}

// APIError is a generic struct for an api error on databricks
type APIError struct {
	ErrorCode  string
	Message    string
	Resource   string
	StatusCode int
}

// Error returns error message string instead of
func (apiError APIError) Error() string {
	if apiError.StatusCode != 404 {
		log.Printf("[WARN] %s:%d - %s", apiError.Resource, apiError.StatusCode, apiError.Message)
	}
	return apiError.Message
}

// IsMissing tells if error is about missing resource
func IsMissing(err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(APIError)
	return ok && e.IsMissing()
}

// IsMissing tells if it is missing resource
func (apiError APIError) IsMissing() bool {
	return apiError.StatusCode == http.StatusNotFound
}

// IsTooManyRequests shows rate exceeded limits
func (apiError APIError) IsTooManyRequests() bool {
	return apiError.StatusCode == http.StatusTooManyRequests
}

// IsRetriable returns true if error is retriable
func (apiError APIError) IsRetriable() bool {
	// Handle transient errors for retries
	for _, substring := range transientErrorStringMatches {
		if strings.Contains(apiError.Message, substring) {
			log.Printf("[INFO] Attempting retry because of %#v", substring)
			return true
		}
	}
	// some API's recommend retries on HTTP 500, but we'll add that later
	return false
}

// NotFound returns properly formatted Not Found error
func NotFound(message string) APIError {
	return APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: 404,
		Message:    message,
	}
}

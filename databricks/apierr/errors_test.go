package apierr

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/stretchr/testify/assert"
)

func TestAPIError_transientRegexMatches(t *testing.T) {
	err := APIError{
		Message: "worker env WorkerEnvId(workerenv-XXXXX) not found",
	}

	assert.True(t, err.IsRetriable(context.Background()))
}

func makeTestReponseWrapper(statusCode int, resp string) common.ResponseWrapper {
	return common.ResponseWrapper{
		Response: &http.Response{
			Status:     fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode)),
			StatusCode: statusCode,
			Request: &http.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/api/2.0/myservice",
				},
			},
		},
		DebugBytes: []byte{},
		ReadCloser: io.NopCloser(bytes.NewReader([]byte(resp))),
	}
}

func TestAPIError_GetAPIError(t *testing.T) {
	testCases := []struct {
		name           string
		resp           common.ResponseWrapper
		wantErrorIs    error
		wantErrorCode  string
		wantMessage    string
		wantStatusCode int
		wantDetails    []ErrorDetail
	}{
		{
			name:           "empty response",
			resp:           makeTestReponseWrapper(http.StatusNotFound, ""),
			wantErrorIs:    ErrNotFound,
			wantErrorCode:  "",
			wantMessage:    "Not Found",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "happy path",
			resp:           makeTestReponseWrapper(http.StatusNotFound, `{"error_code": "RESOURCE_DOES_NOT_EXIST", "message": "Cluster abc does not exist"}`),
			wantErrorIs:    ErrResourceDoesNotExist,
			wantErrorCode:  "RESOURCE_DOES_NOT_EXIST",
			wantMessage:    "Cluster abc does not exist",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "error details",
			resp:           makeTestReponseWrapper(http.StatusNotFound, `{"error_code": "RESOURCE_DOES_NOT_EXIST", "message": "Cluster abc does not exist", "details": [{"@type": "type", "reason": "reason", "domain": "domain", "metadata": {"key": "value"}}]}`),
			wantErrorIs:    ErrResourceDoesNotExist,
			wantErrorCode:  "RESOURCE_DOES_NOT_EXIST",
			wantMessage:    "Cluster abc does not exist",
			wantStatusCode: http.StatusNotFound,
			wantDetails: []ErrorDetail{
				{
					Type:     "type",
					Reason:   "reason",
					Domain:   "domain",
					Metadata: map[string]string{"key": "value"},
				},
			},
		},
		{
			name:           "string error response",
			resp:           makeTestReponseWrapper(http.StatusBadRequest, `MALFORMED_REQUEST: vpc_endpoints malformed parameters: VPC Endpoint ... with use_case ... cannot be attached in ... list`),
			wantErrorIs:    ErrBadRequest,
			wantErrorCode:  "MALFORMED_REQUEST",
			wantMessage:    "vpc_endpoints malformed parameters: VPC Endpoint ... with use_case ... cannot be attached in ... list",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "numeric error code",
			resp:           makeTestReponseWrapper(http.StatusBadRequest, `{"error_code": 500, "message": "Cluster abc does not exist"}`),
			wantErrorIs:    ErrBadRequest,
			wantErrorCode:  "500",
			wantMessage:    "Cluster abc does not exist",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "private link redirect",
			resp: common.ResponseWrapper{
				Response: &http.Response{
					Request: &http.Request{
						URL: &url.URL{
							Host:     "adb-12345678910.1.azuredatabricks.net",
							Path:     "/login.html",
							RawQuery: "error=private-link-validation-error",
						},
					},
				},
			},
			wantErrorIs:    ErrPermissionDenied,
			wantErrorCode:  "PRIVATE_LINK_VALIDATION_ERROR",
			wantMessage:    "The requested workspace has Azure Private Link enabled and is not accessible from the current network. Ensure that Azure Private Link is properly configured and that your device has access to the Azure Private Link endpoint. For more information, see https://learn.microsoft.com/en-us/azure/databricks/security/network/classic/private-link-standard#authentication-troubleshooting.",
			wantStatusCode: http.StatusForbidden,
		},
		{
			name: "applies overrides",
			resp: common.ResponseWrapper{
				Response: &http.Response{
					StatusCode: http.StatusBadRequest,
					Request: &http.Request{
						Method: "GET",
						URL: &url.URL{
							Path: "/api/2.1/clusters/get",
						},
					},
				},
				DebugBytes: []byte{},
				ReadCloser: io.NopCloser(bytes.NewReader([]byte(`{"error_code": "INVALID_PARAMETER_VALUE", "message": "Cluster abc does not exist"}`))),
			},
			wantErrorIs:    ErrResourceDoesNotExist,
			wantErrorCode:  "INVALID_PARAMETER_VALUE",
			wantMessage:    "Cluster abc does not exist",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "unexpected error",
			resp:           makeTestReponseWrapper(http.StatusInternalServerError, `unparsable error message`),
			wantErrorCode:  "INTERNAL_SERVER_ERROR",
			wantMessage:    "unable to parse response. This is likely a bug in the Databricks SDK for Go or the underlying REST API. Please report this issue with the following debugging information to the SDK issue tracker at https://github.com/databricks/databricks-sdk-go/issues. Request log:\n```\nGET /api/2.0/myservice\n> * Host: \n<  500 Internal Server Error\n< unparsable error message\n```",
			wantStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetAPIError(context.Background(), tc.resp).(*APIError)
			assert.Equal(t, tc.wantErrorCode, got.ErrorCode)
			assert.Equal(t, tc.wantMessage, got.Message)
			assert.Equal(t, tc.wantStatusCode, got.StatusCode)
			assert.Equal(t, tc.wantDetails, got.Details)
			if tc.wantErrorIs != nil {
				assert.ErrorIs(t, got, tc.wantErrorIs)
			}
		})
	}
}

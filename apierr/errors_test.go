package apierr

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestAPIError_IsRetriable_transientRegexMatches(t *testing.T) {
	err := APIError{
		Message: "worker env WorkerEnvId(workerenv-XXXXX) not found",
	}

	if !err.IsRetriable(context.Background()) {
		t.Errorf("expected error to be retriable")
	}
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

func TestGetAPIError(t *testing.T) {
	testCases := []struct {
		name        string
		resp        common.ResponseWrapper
		want        *APIError
		wantErrorIs error
	}{
		{
			name: "empty response",
			resp: makeTestReponseWrapper(http.StatusNotFound, ""),
			want: &APIError{
				ErrorCode:  "",
				Message:    "Not Found",
				StatusCode: http.StatusNotFound,
			},
			wantErrorIs: ErrNotFound,
		},
		{
			name: "happy path",
			resp: makeTestReponseWrapper(http.StatusNotFound, `{
				"error_code": "RESOURCE_DOES_NOT_EXIST",
				"message": "Cluster abc does not exist"
			}`),
			want: &APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusNotFound,
			},
			wantErrorIs: ErrResourceDoesNotExist,
		},
		{
			name: "string error response",
			resp: makeTestReponseWrapper(http.StatusBadRequest, `MALFORMED_REQUEST: vpc_endpoints malformed parameters: VPC Endpoint ... with use_case ... cannot be attached in ... list`),
			want: &APIError{
				ErrorCode:  "MALFORMED_REQUEST",
				Message:    "vpc_endpoints malformed parameters: VPC Endpoint ... with use_case ... cannot be attached in ... list",
				StatusCode: http.StatusBadRequest,
			},
			wantErrorIs: ErrBadRequest,
		},
		{
			name: "numeric error code",
			resp: makeTestReponseWrapper(http.StatusBadRequest, `{
				"error_code": 500,
				"message": "Cluster abc does not exist"}
			`),
			want: &APIError{
				ErrorCode:  "500",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusBadRequest,
			},
			wantErrorIs: ErrBadRequest,
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
			want: &APIError{
				ErrorCode:  "PRIVATE_LINK_VALIDATION_ERROR",
				Message:    "The requested workspace has Azure Private Link enabled and is not accessible from the current network. Ensure that Azure Private Link is properly configured and that your device has access to the Azure Private Link endpoint. For more information, see https://learn.microsoft.com/en-us/azure/databricks/security/network/classic/private-link-standard#authentication-troubleshooting.",
				StatusCode: http.StatusForbidden,
			},
			wantErrorIs: ErrPermissionDenied,
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
				ReadCloser: io.NopCloser(bytes.NewReader([]byte(`{
					"error_code": "INVALID_PARAMETER_VALUE",
					"message": "Cluster abc does not exist"}
				`))),
			},
			want: &APIError{
				ErrorCode:  "INVALID_PARAMETER_VALUE",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusBadRequest,
			},
			wantErrorIs: ErrResourceDoesNotExist,
		},
		{
			name: "unexpected error",
			resp: makeTestReponseWrapper(http.StatusInternalServerError, `unparsable error message`),
			want: &APIError{
				ErrorCode:  "INTERNAL_SERVER_ERROR",
				Message:    "unable to parse response. This is likely a bug in the Databricks SDK for Go or the underlying REST API. Please report this issue with the following debugging information to the SDK issue tracker at https://github.com/databricks/databricks-sdk-go/issues. Request log:\n```\nGET /api/2.0/myservice\n> * Host: \n<  500 Internal Server Error\n< unparsable error message\n```",
				StatusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "all valid error details type",
			resp: makeTestReponseWrapper(http.StatusNotFound, `{
				"error_code": "RESOURCE_DOES_NOT_EXIST", 
				"message": "Cluster abc does not exist", 
				"details": [
					{
						"@type": "type.googleapis.com/google.rpc.ErrorInfo",
						"reason": "reason",
						"domain": "domain",
						"metadata": {"k1": "v1", "k2": "v2"}
					},
					{
						"@type": "type.googleapis.com/google.rpc.RequestInfo",
						"request_id": "req42",
						"serving_data": "data"
					},
					{
						"@type": "type.googleapis.com/google.rpc.RetryInfo",
						"retry_delay": {"seconds": 1, "nanos": 1}
					},
					{
						"@type": "type.googleapis.com/google.rpc.DebugInfo",
						"stack_entries": ["entry1", "entry2"],
						"detail": "detail"
					},
					{
						"@type": "type.googleapis.com/google.rpc.QuotaFailure",
						"violations": [{"subject": "subject", "description": "description"}]
					},
					{
						"@type": "type.googleapis.com/google.rpc.PreconditionFailure",
						"violations": [{"type": "type", "subject": "subject", "description": "description"}]
					},
					{
						"@type": "type.googleapis.com/google.rpc.BadRequest",
						"field_violations": [{"field": "field", "description": "description"}]
					},
					{
						"@type": "type.googleapis.com/google.rpc.ResourceInfo",
						"resource_type": "resource_type",
						"resource_name": "resource_name",
						"owner": "owner",
						"description": "description"
					},
					{
						"@type": "type.googleapis.com/google.rpc.Help",
						"links": [{"description": "description", "url": "url"}]
					},
					{
						"@type": "foo", 
						"reason": "reason"
					}
				]
			}`),
			want: &APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusNotFound,
				Details: []ErrorDetail{
					{
						Type:   "type.googleapis.com/google.rpc.ErrorInfo",
						Reason: "reason",
						Domain: "domain",
						Metadata: map[string]string{
							"k1": "v1",
							"k2": "v2",
						},
					},
					{
						Type: "type.googleapis.com/google.rpc.RequestInfo",
					},
					{
						Type: "type.googleapis.com/google.rpc.RetryInfo",
					},
					{
						Type: "type.googleapis.com/google.rpc.DebugInfo",
					},
					{
						Type: "type.googleapis.com/google.rpc.QuotaFailure",
					},
					{
						Type: "type.googleapis.com/google.rpc.PreconditionFailure",
					},
					{
						Type: "type.googleapis.com/google.rpc.BadRequest",
					},
					{
						Type: "type.googleapis.com/google.rpc.ResourceInfo",
					},
					{
						Type: "type.googleapis.com/google.rpc.Help",
					},
					{
						Type:   "foo",
						Reason: "reason",
					},
				},
				errorDetails: ErrorDetails{
					ErrorInfo: &ErrorInfo{
						Reason:   "reason",
						Domain:   "domain",
						Metadata: map[string]string{"k1": "v1", "k2": "v2"},
					},
					RequestInfo: &RequestInfo{
						RequestID:   "req42",
						ServingData: "data",
					},
					RetryInfo: &RetryInfo{
						RetryDelay: time.Second + time.Nanosecond,
					},
					DebugInfo: &DebugInfo{
						StackEntries: []string{"entry1", "entry2"},
						Detail:       "detail",
					},
					QuotaFailure: &QuotaFailure{
						Violations: []QuotaFailureViolation{{Subject: "subject", Description: "description"}},
					},
					PreconditionFailure: &PreconditionFailure{
						Violations: []PreconditionFailureViolation{{Type: "type", Subject: "subject", Description: "description"}},
					},
					BadRequest: &BadRequest{
						FieldViolations: []BadRequestFieldViolation{{Field: "field", Description: "description"}},
					},
					ResourceInfo: &ResourceInfo{
						ResourceType: "resource_type",
						ResourceName: "resource_name",
						Owner:        "owner",
						Description:  "description",
					},
					Help: &Help{
						Links: []HelpLink{{Description: "description", URL: "url"}},
					},
					UnknownDetails: []any{map[string]interface{}{
						"@type":  "foo",
						"reason": "reason",
					}},
				},
			},
		},
		{
			name: "only keep the last error details of a type",
			resp: makeTestReponseWrapper(http.StatusNotFound, `{
				"error_code": "RESOURCE_DOES_NOT_EXIST", 
				"message": "Cluster abc does not exist", 
				"details": [
					{
						"@type": "type.googleapis.com/google.rpc.ErrorInfo",
						"reason": "first"
					},
					{
						"@type": "type.googleapis.com/google.rpc.ErrorInfo",
						"reason": "second"
					}
				]
			}`),
			want: &APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusNotFound,
				Details: []ErrorDetail{
					{
						Type:   "type.googleapis.com/google.rpc.ErrorInfo",
						Reason: "first",
					},
					{
						Type:   "type.googleapis.com/google.rpc.ErrorInfo",
						Reason: "second",
					},
				},
				errorDetails: ErrorDetails{
					ErrorInfo: &ErrorInfo{
						Reason: "second",
					},
				},
			},
		},
		{
			name: "silently drop invalid error details",
			resp: makeTestReponseWrapper(http.StatusNotFound, `{
				"error_code": "RESOURCE_DOES_NOT_EXIST", 
				"message": "Cluster abc does not exist", 
				"details": [
					42, 
					{
						"@type": "type.googleapis.com/google.rpc.ErrorInfo",
						"reason": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.RequestInfo",
						"request_id": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.RetryInfo",
						"retry_delay": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.DebugInfo",
						"stack_entries": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.QuotaFailure",
						"violations": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.PreconditionFailure",
						"violations": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.BadRequest",
						"field_violations": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.ResourceInfo",
						"resource_type": 0
					},
					{
						"@type": "type.googleapis.com/google.rpc.Help",
						"links": 0
					}
				]
			}`),
			want: &APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusNotFound,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetAPIError(context.Background(), tc.resp)

			opts := cmp.Options{
				cmp.AllowUnexported(APIError{}),            // to check ErrorDetails
				cmpopts.IgnoreFields(APIError{}, "unwrap"), // tested via wantErrorIs
			}
			if diff := cmp.Diff(tc.want, got, opts); diff != "" {
				t.Errorf("unexpected error (-want +got):\n%s", diff)
			}
			if tc.wantErrorIs != nil && !errors.Is(got, tc.wantErrorIs) {
				t.Errorf("errors.Is(%q, %q) failed", got, tc.wantErrorIs)
			}
		})
	}
}

package apierr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestAPIError_transientRegexMatches(t *testing.T) {
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

func TestAPIError_GetAPIError(t *testing.T) {
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
			name: "supported details",
			resp: makeTestReponseWrapper(http.StatusNotFound, `{
				"error_code": "RESOURCE_DOES_NOT_EXIST", 
				"message": "Cluster abc does not exist", 
				"details": [{"@type": "type", "reason": "reason", "domain": "domain", "metadata": {"key": "value"}}]
			}`),
			want: &APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusNotFound,
				Details: []ErrorDetail{{
					Type:     "type",
					Reason:   "reason",
					Domain:   "domain",
					Metadata: map[string]string{"key": "value"},
				}},
				RawDetails: []json.RawMessage{[]byte(`{"@type": "type", "reason": "reason", "domain": "domain", "metadata": {"key": "value"}}`)},
			},
			wantErrorIs: ErrResourceDoesNotExist,
		},
		{
			name: "unsupported details",
			resp: makeTestReponseWrapper(http.StatusNotFound, `{
				"error_code": "RESOURCE_DOES_NOT_EXIST", 
				"message": "Cluster abc does not exist", 
				"details": [{"@type": "type", "reason": 5}]
			}`),
			want: &APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusNotFound,
				RawDetails: []json.RawMessage{[]byte(`{"@type": "type", "reason": 5}`)},
			},
			wantErrorIs: ErrResourceDoesNotExist,
		},
		{
			name: "supported and unsupported details",
			resp: makeTestReponseWrapper(http.StatusNotFound, `{
				"error_code": "RESOURCE_DOES_NOT_EXIST", 
				"message": "Cluster abc does not exist", 
				"details": [{"@type": "type", "reason": 5}, {"@type": "type", "reason": "foo"}]
			}`),
			want: &APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				Message:    "Cluster abc does not exist",
				StatusCode: http.StatusNotFound,
				Details: []ErrorDetail{{
					Type:   "type",
					Reason: "foo",
				}},
				RawDetails: []json.RawMessage{
					[]byte(`{"@type": "type", "reason": 5}`),
					[]byte(`{"@type": "type", "reason": "foo"}`),
				},
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetAPIError(context.Background(), tc.resp)

			if diff := cmp.Diff(tc.want, got, cmpopts.IgnoreUnexported(APIError{})); diff != "" {
				t.Errorf("unexpected error (-want +got):\n%s", diff)
			}
			if tc.wantErrorIs != nil && !errors.Is(got, tc.wantErrorIs) {
				t.Errorf("errors.Is(%q, %q) failed", got, tc.wantErrorIs)
			}
		})
	}
}

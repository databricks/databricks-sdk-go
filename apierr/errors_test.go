package apierr

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/stretchr/testify/assert"
)

func TestGetAPIError_handlesEmptyResponse(t *testing.T) {
	resp := common.ResponseWrapper{
		Response: &http.Response{
			Request: &http.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/api/2.0/compute/get",
				},
			},
			StatusCode: http.StatusConflict,
		},
		DebugBytes: []byte{},
		ReadCloser: io.NopCloser(bytes.NewReader([]byte{})),
	}

	err := GetAPIError(context.Background(), resp)

	assert.Equal(t, err.(*APIError).Message, "")
}

func TestGetAPIError_appliesOverrides(t *testing.T) {
	resp := common.ResponseWrapper{
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
	}

	err := GetAPIError(context.Background(), resp)

	assert.ErrorIs(t, err, ErrResourceDoesNotExist)
}

func TestGetAPIError_parseIntErrorCode(t *testing.T) {
	resp := common.ResponseWrapper{
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
		ReadCloser: io.NopCloser(bytes.NewReader([]byte(`{"error_code": 500, "status_code": 400, "message": "Cluster abc does not exist"}`))),
	}

	err := GetAPIError(context.Background(), resp)

	assert.ErrorIs(t, err, ErrBadRequest)
	assert.Equal(t, err.(*APIError).ErrorCode, "500")
}

func TestGetAPIError_mapsPrivateLinkRedirect(t *testing.T) {
	resp := common.ResponseWrapper{
		Response: &http.Response{
			Request: &http.Request{
				URL: &url.URL{
					Host:     "adb-12345678910.1.azuredatabricks.net",
					Path:     "/login.html",
					RawQuery: "error=private-link-validation-error",
				},
			},
		},
	}

	err := GetAPIError(context.Background(), resp)

	assert.ErrorIs(t, err, ErrPermissionDenied)
	assert.Equal(t, err.(*APIError).ErrorCode, "PRIVATE_LINK_VALIDATION_ERROR")
}

func TestAPIError_transientRegexMatches(t *testing.T) {
	err := APIError{
		Message: "worker env WorkerEnvId(workerenv-XXXXX) not found",
	}

	assert.True(t, err.IsRetriable(context.Background()))
}

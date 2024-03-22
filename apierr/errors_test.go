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

func TestGetAPIErrorHandlesEmptyResponse(t *testing.T) {
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

func TestGetAPIErrorAppliesOverrides(t *testing.T) {
	resp := common.ResponseWrapper{
		Response: &http.Response{
			StatusCode: http.StatusConflict,
			Request: &http.Request{
				Method: "GET",
				URL: &url.URL{
					Path: "/api/2.0/clusters/get",
				},
			},
		},
		DebugBytes: []byte{},
		ReadCloser: io.NopCloser(bytes.NewReader([]byte(`{"error_code": "INVALID_PARAMETER_VALUE", "message": "Cluster abc does not exist"}`))),
	}
	ctx := context.Background()
	err := GetAPIError(ctx, resp)
	assert.ErrorIs(t, err, ErrResourceDoesNotExist)
}

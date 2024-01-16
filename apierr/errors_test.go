package apierr

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/stretchr/testify/assert"
)

func TestGetAPIErrorHandlesEmptyResponse(t *testing.T) {
	resp := common.ResponseWrapper{
		Response: &http.Response{
			StatusCode: http.StatusConflict,
		},
		DebugBytes: []byte{},
		ReadCloser: io.NopCloser(bytes.NewReader([]byte{})),
	}
	err := GetAPIError(context.Background(), resp)
	assert.Equal(t, err.(*APIError).Message, "")
}

func TestIsMissing_APIError(t *testing.T) {
	assert.True(t, IsMissing(&APIError{
		ErrorCode: "RESOURCE_DOES_NOT_EXIST",
	}))
	assert.False(t, IsMissing(&APIError{
		ErrorCode: "RESOURCE_ALREADY_EXISTS",
	}))
}

func TestIsMissing_BaseError(t *testing.T) {
	assert.True(t, IsMissing(ErrNotFound))
	assert.False(t, IsMissing(ErrResourceConflict))
}

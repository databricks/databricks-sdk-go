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

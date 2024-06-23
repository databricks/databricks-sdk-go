package apierr

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyOverrides(t *testing.T) {
	testCases := []struct {
		name          string
		initialError  *APIError
		path          string
		method        string
		expectedError error
	}{
		{
			initialError: &APIError{
				ErrorCode:  "INVALID_PARAMETER_VALUE",
				StatusCode: 400,
				Message:    "Cluster abc does not exist",
			},
			path:          "/api/2.0/clusters/get",
			method:        "GET",
			expectedError: ErrResourceDoesNotExist,
		},
		{
			initialError: &APIError{
				ErrorCode:  "INVALID_PARAMETER_VALUE",
				StatusCode: 400,
				Message:    "Job abc does not exist",
			},
			path:          "/api/2.0/jobs/get",
			method:        "GET",
			expectedError: ErrResourceDoesNotExist,
		},
		{
			initialError: &APIError{
				ErrorCode:  "INVALID_PARAMETER_VALUE",
				StatusCode: 400,
				Message:    "Job abc does not exist",
			},
			path:          "/api/2.1/jobs/get",
			method:        "GET",
			expectedError: ErrResourceDoesNotExist,
		},
		{
			initialError: &APIError{
				ErrorCode:  "INVALID_PARAMETER_VALUE",
				StatusCode: 400,
				Message:    "Notebook abc does not exist",
			},
			path:          "/api/2.0/notebooks/get",
			method:        "GET",
			expectedError: ErrBadRequest,
		},
		{
			initialError: &APIError{
				ErrorCode:  "INVALID_PARAMETER_VALUE",
				StatusCode: 400,
				Message:    "Invalid job attribute",
			},
			path:          "/api/2.0/jobs/get",
			method:        "GET",
			expectedError: ErrBadRequest,
		},
		{
			initialError: &APIError{
				ErrorCode:  "INVALID_PARAMETER_VALUE",
				StatusCode: 404,
				Message:    "Job abc does not exist",
			},
			path:          "/api/2.0/jobs/get",
			method:        "GET",
			expectedError: ErrBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := &http.Response{
				Request: &http.Request{
					URL: &url.URL{
						Path: tc.path,
					},
					Method: tc.method,
				},
				StatusCode: tc.initialError.StatusCode,
			}

			applyOverrides(context.Background(), tc.initialError, resp)

			assert.ErrorIs(t, tc.initialError, tc.expectedError)
		})
	}
}

package v2

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	jobsv2 "github.com/databricks/databricks-sdk-go/protos/jobs/v2"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/testing/protocmp"
)

// mockTransport implements http.RoundTripper for testing
type mockTransport func(r *http.Request) (*http.Response, error)

func (m mockTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return m(r)
}

// createTestClient creates a DatabricksClient with mocked transport
func createTestClient(t *testing.T, transport mockTransport) *client.DatabricksClient {
	cfg := &config.Config{
		Host:          "https://test.databricks.com",
		Token:         "test-token",
		ConfigFile:    "/dev/null",
		HTTPTransport: transport,
	}

	c, err := client.New(cfg)
	if err != nil {
		t.Fatalf("Failed to create test client: %v", err)
	}
	return c
}

func TestCreateJob_Success(t *testing.T) {
	// Create test request
	jobName := "test-job"
	description := "test description"
	request := &jobsv2.CreateJobRequest{
		Job: &jobsv2.Job{
			Name:        &jobName,
			Description: &description,
		},
	}

	// Expected response
	createdAt := int64(1234567890)
	expectedJob := &jobsv2.Job{
		Name:        &jobName,
		Description: &description,
		CreatedAt:   &createdAt,
	}

	// Mock transport that verifies request and returns response
	transport := mockTransport(func(r *http.Request) (*http.Response, error) {
		// Verify HTTP method and path
		if r.Method != http.MethodPost {
			t.Errorf("Expected method %s, got %s", http.MethodPost, r.Method)
		}
		if r.URL.Path != "/jobs/v2/jobs" {
			t.Errorf("Expected path /jobs/v2/jobs, got %s", r.URL.Path)
		}

		// Verify headers
		if got := r.Header.Get("Content-Type"); got != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", got)
		}
		if got := r.Header.Get("Accept"); got != "application/json" {
			t.Errorf("Expected Accept application/json, got %s", got)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-token" {
			t.Errorf("Expected Authorization 'Bearer test-token', got %s", got)
		}

		// Verify request body as raw JSON
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}

		// Validate the actual JSON structure sent over HTTP
		expectedJSON := `{"job":{"name":"test-job","description":"test description"}}`

		// Parse both as JSON to compare structure (not string comparison to avoid formatting issues)
		var actualJSON, expectedJSONParsed map[string]any
		if err := json.Unmarshal(body, &actualJSON); err != nil {
			t.Fatalf("Request body is not valid JSON: %v", err)
		}
		if err := json.Unmarshal([]byte(expectedJSON), &expectedJSONParsed); err != nil {
			t.Fatalf("Expected JSON is invalid: %v", err)
		}

		if diff := cmp.Diff(expectedJSONParsed, actualJSON); diff != "" {
			t.Errorf("JSON request mismatch (-want +got):\\n%s", diff)
		}

		// Return successful response
		responseBody, err := protojson.Marshal(expectedJob)
		if err != nil {
			t.Fatalf("Failed to marshal response: %v", err)
		}

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(string(responseBody))),
			Request:    r,
			Header:     make(http.Header),
		}, nil
	})

	// Create client and service
	testClient := createTestClient(t, transport)
	service := &jobsServiceImpl{client: testClient}

	// Execute the method
	result, err := service.CreateJob(context.Background(), request)

	// Verify results
	if err != nil {
		t.Fatalf("CreateJob failed: %v", err)
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	if diff := cmp.Diff(expectedJob, result, protocmp.Transform()); diff != "" {
		t.Errorf("Result mismatch (-want +got):\\n%s", diff)
	}
}

func TestCreateJob_HTTPError(t *testing.T) {
	// Create test request
	jobName := "test-job"
	request := &jobsv2.CreateJobRequest{
		Job: &jobsv2.Job{
			Name: &jobName,
		},
	}

	// Mock transport that returns an HTTP error
	transport := mockTransport(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 400,
			Body:       io.NopCloser(strings.NewReader(`{"error":"Bad Request","message":"Invalid job name"}`)),
			Request:    r,
			Header:     make(http.Header),
		}, nil
	})

	// Create client and service
	testClient := createTestClient(t, transport)
	service := &jobsServiceImpl{client: testClient}

	// Execute the method
	result, err := service.CreateJob(context.Background(), request)

	// Verify error handling
	if err == nil {
		t.Fatal("Expected error but got none")
	}
	if result != nil {
		t.Errorf("Expected nil result but got: %v", result)
	}
	if !strings.Contains(err.Error(), "HTTP 400") {
		t.Errorf("Expected error to contain 'HTTP 400', got: %v", err)
	}
}

func TestCreateJob_NetworkError(t *testing.T) {
	// Create test request
	jobName := "test-job"
	request := &jobsv2.CreateJobRequest{
		Job: &jobsv2.Job{
			Name: &jobName,
		},
	}

	// Mock transport that returns a network error
	transport := mockTransport(func(r *http.Request) (*http.Response, error) {
		return nil, io.ErrUnexpectedEOF
	})

	// Create client and service
	testClient := createTestClient(t, transport)
	service := &jobsServiceImpl{client: testClient}

	// Execute the method
	result, err := service.CreateJob(context.Background(), request)

	// Verify error handling
	if err == nil {
		t.Fatal("Expected error but got none")
	}
	if result != nil {
		t.Errorf("Expected nil result but got: %v", result)
	}
	if !strings.Contains(err.Error(), "HTTP request failed") {
		t.Errorf("Expected error to contain 'HTTP request failed', got: %v", err)
	}
}

func TestCreateJob_NilRequest(t *testing.T) {
	// Test with nil request - DoProto should handle this gracefully
	transport := mockTransport(func(r *http.Request) (*http.Response, error) {
		// DoProto converts nil request to empty JSON object "{}"
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}
		if got := string(body); got != "{}" {
			t.Errorf("Expected body '{}', got %s", got)
		}

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
			Request:    r,
		}, nil
	})

	testClient := createTestClient(t, transport)
	service := &jobsServiceImpl{client: testClient}

	// Test with nil request (should not cause panic)
	result, err := service.CreateJob(context.Background(), nil)
	if err != nil {
		t.Fatalf("CreateJob with nil request failed: %v", err)
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
}

func TestCreateJob_InvalidResponseUnmarshalling(t *testing.T) {
	// Create test request
	jobName := "test-job"
	request := &jobsv2.CreateJobRequest{
		Job: &jobsv2.Job{
			Name: &jobName,
		},
	}

	// Mock transport that returns invalid JSON
	transport := mockTransport(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`{invalid json}`)),
			Request:    r,
			Header:     make(http.Header),
		}, nil
	})

	// Create client and service
	testClient := createTestClient(t, transport)
	service := &jobsServiceImpl{client: testClient}

	// Execute the method
	result, err := service.CreateJob(context.Background(), request)

	// Verify error handling
	if err == nil {
		t.Fatal("Expected error but got none")
	}
	if result != nil {
		t.Errorf("Expected nil result but got: %v", result)
	}
	if !strings.Contains(err.Error(), "failed to unmarshal proto response") {
		t.Errorf("Expected error to contain 'failed to unmarshal proto response', got: %v", err)
	}
}

func TestCreateJob_EmptyResponse(t *testing.T) {
	// Create test request
	jobName := "test-job"
	request := &jobsv2.CreateJobRequest{
		Job: &jobsv2.Job{
			Name: &jobName,
		},
	}

	// Mock transport that returns empty response
	transport := mockTransport(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
			Request:    r,
			Header:     make(http.Header),
		}, nil
	})

	// Create client and service
	testClient := createTestClient(t, transport)
	service := &jobsServiceImpl{client: testClient}

	// Execute the method
	result, err := service.CreateJob(context.Background(), request)

	// Verify that empty response is handled correctly
	if err != nil {
		t.Fatalf("CreateJob with empty response failed: %v", err)
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	// Fields should be nil/empty since response was empty
	if result.Name != nil {
		t.Errorf("Expected nil Name, got: %v", result.Name)
	}
	if result.Description != nil {
		t.Errorf("Expected nil Description, got: %v", result.Description)
	}
	if result.CreatedAt != nil {
		t.Errorf("Expected nil CreatedAt, got: %v", result.CreatedAt)
	}
}

func TestCreateJob_RequestMarshalling(t *testing.T) {
	// Test that proto messages are correctly marshalled to JSON
	jobName := "test-job-marshalling"
	description := "test description for marshalling"
	request := &jobsv2.CreateJobRequest{
		Job: &jobsv2.Job{
			Name:        &jobName,
			Description: &description,
		},
	}

	// Mock transport that captures and validates the exact JSON
	transport := mockTransport(func(r *http.Request) (*http.Response, error) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}

		// Validate the actual JSON structure sent over HTTP
		expectedJSON := `{"job":{"name":"test-job-marshalling","description":"test description for marshalling"}}`

		// Parse both as JSON to compare structure
		var actualJSON, expectedJSONParsed map[string]any
		if err := json.Unmarshal(body, &actualJSON); err != nil {
			t.Fatalf("Request body is not valid JSON: %v", err)
		}
		if err := json.Unmarshal([]byte(expectedJSON), &expectedJSONParsed); err != nil {
			t.Fatalf("Expected JSON is invalid: %v", err)
		}

		if diff := cmp.Diff(expectedJSONParsed, actualJSON); diff != "" {
			t.Errorf("JSON request mismatch (-want +got):\\n%s", diff)
		}

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
			Request:    r,
		}, nil
	})

	testClient := createTestClient(t, transport)
	service := &jobsServiceImpl{client: testClient}

	_, err := service.CreateJob(context.Background(), request)
	if err != nil {
		t.Fatalf("CreateJob failed: %v", err)
	}
}

func TestCreateJob_ContextCancellation(t *testing.T) {
	// Create test request
	jobName := "test-job"
	request := &jobsv2.CreateJobRequest{
		Job: &jobsv2.Job{
			Name: &jobName,
		},
	}

	// Create cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	// Mock transport that checks context cancellation
	transport := mockTransport(func(r *http.Request) (*http.Response, error) {
		// The context should be cancelled when we get here
		select {
		case <-r.Context().Done():
			return nil, r.Context().Err()
		default:
			t.Fatal("Expected context to be cancelled")
			return nil, nil
		}
	})

	testClient := createTestClient(t, transport)
	service := &jobsServiceImpl{client: testClient}

	// Execute the method with cancelled context
	result, err := service.CreateJob(ctx, request)

	// Verify context cancellation is handled
	if err == nil {
		t.Fatal("Expected error but got none")
	}
	if result != nil {
		t.Errorf("Expected nil result but got: %v", result)
	}
	if !strings.Contains(err.Error(), "context canceled") {
		t.Errorf("Expected error to contain 'context canceled', got: %v", err)
	}
}

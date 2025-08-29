package v1

import (
	"context"

	libraryv1 "github.com/databricks/databricks-sdk-go/library/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

// JSONClient is a minimal interface for JSON-in/JSON-out requests.
// It enables plugging in a fake/stub for prototyping without a real HTTP client.
// The request and response payloads are raw JSON bytes.
// Method, path, headers, and queryParams are included for parity but can be ignored by implementations.
//
// Example stub (echo):
//
//	type Echo struct{}
//	func (Echo) Do(ctx context.Context, method, path string, headers map[string]string, queryParams map[string]any, request []byte) ([]byte, error) { return request, nil }
//
// You can pass Echo{} to NewLibraryService for a pure echo prototype.
// If nil is passed to NewLibraryService, CreateLibrary will simply echo the request's library back.
type JSONClient interface {
	Do(ctx context.Context, method, path string, headers map[string]string, queryParams map[string]any, requestbody []byte) ([]byte, error)
}

type LibraryService struct {
	jsonClient JSONClient
}

func NewLibraryService(jsonClient JSONClient) *LibraryService {
	return &LibraryService{jsonClient: jsonClient}
}

func (s *LibraryService) CreateLibrary(ctx context.Context, req *libraryv1.CreateLibraryRequest) (*libraryv1.Library, error) {
	// If no transport is provided, just echo the incoming request's library.
	if s.jsonClient == nil {
		if req == nil || req.GetLibrary() == nil {
			return &libraryv1.Library{}, nil
		}
		return req.GetLibrary(), nil
	}

	// Marshal the proto message to JSON, call the JSON client, and unmarshal the response using protojson.
	payload, err := protojson.Marshal(req.GetLibrary())
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	var queryParams map[string]any

	respJSON, err := s.jsonClient.Do(ctx, "POST", "/library/v1/library", headers, queryParams, payload)
	if err != nil {
		return nil, err
	}

	// If no response body is returned, echo back the request body.
	if len(respJSON) == 0 {
		if req == nil || req.GetLibrary() == nil {
			return &libraryv1.Library{}, nil
		}
		return req.GetLibrary(), nil
	}

	var resp libraryv1.Library
	if err := protojson.Unmarshal(respJSON, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

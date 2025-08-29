// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package v2

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	libraryv2 "github.com/databricks/databricks-sdk-go/protos/library/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

// unexported type that holds implementations of just LibraryService API methods
type library_serviceServiceImpl struct {
	client *client.DatabricksClient
}

func (a *library_serviceServiceImpl) CreateLibrary(ctx context.Context, request *libraryv2.CreateLibraryRequest) (*libraryv2.Library, error) {
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	var library libraryv2.Library
	err := a.client.DoProto(ctx, http.MethodGet, path, headers, queryParams, request, &library)
	if err != nil {
		return nil, err
	}

	return &library, nil

}

func (a *library_serviceServiceImpl) DeleteLibrary(ctx context.Context, request *libraryv2.DeleteLibraryRequest) (*emptypb.Empty, error) {
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	err := a.client.DoProto(ctx, http.MethodGet, path, headers, queryParams, request, nil)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil

}

func (a *library_serviceServiceImpl) GetLibrary(ctx context.Context, request *libraryv2.GetLibraryRequest) (*libraryv2.Library, error) {
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	var library libraryv2.Library
	err := a.client.DoProto(ctx, http.MethodGet, path, headers, queryParams, request, &library)
	if err != nil {
		return nil, err
	}

	return &library, nil

}

func (a *library_serviceServiceImpl) ListLibraries(ctx context.Context, request *libraryv2.ListLibrariesRequest) (*libraryv2.ListLibrariesResponse, error) {
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	var listLibrariesResponse libraryv2.ListLibrariesResponse
	err := a.client.DoProto(ctx, http.MethodGet, path, headers, queryParams, request, &listLibrariesResponse)
	if err != nil {
		return nil, err
	}

	return &listLibrariesResponse, nil

}

func (a *library_serviceServiceImpl) UpdateLibrary(ctx context.Context, request *libraryv2.UpdateLibraryRequest) (*libraryv2.Library, error) {
	path := ""
	queryParams := make(map[string]any)
	headers := make(map[string]string)

	var library libraryv2.Library
	err := a.client.DoProto(ctx, http.MethodGet, path, headers, queryParams, request, &library)
	if err != nil {
		return nil, err
	}

	return &library, nil

}

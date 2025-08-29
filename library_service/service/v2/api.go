// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage LibraryService Service, etc.
package v2

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
	libraryv2 "github.com/databricks/databricks-sdk-go/protos/library/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LibraryServiceInterface interface {
	CreateLibrary(ctx context.Context, request *libraryv2.CreateLibraryRequest) (*libraryv2.Library, error)

	DeleteLibrary(ctx context.Context, request *libraryv2.DeleteLibraryRequest) (*emptypb.Empty, error)

	GetLibrary(ctx context.Context, request *libraryv2.GetLibraryRequest) (*libraryv2.Library, error)

	ListLibraries(ctx context.Context, request *libraryv2.ListLibrariesRequest) (*libraryv2.ListLibrariesResponse, error)

	UpdateLibrary(ctx context.Context, request *libraryv2.UpdateLibraryRequest) (*libraryv2.Library, error)
}

func NewLibraryServiceService(client *client.DatabricksClient) *LibraryServiceAPI {
	return &LibraryServiceAPI{
		library_serviceServiceImpl: library_serviceServiceImpl{
			client: client,
		},
	}
}

type LibraryServiceAPI struct {
	library_serviceServiceImpl
}

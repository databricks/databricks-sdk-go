// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jsonmarshallv2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/internal/testspecs/service/jsonmarshallv2/jsonmarshallv2pb"
)

// unexported type that holds implementations of just JsonMarshallV2 API methods
type jsonMarshallV2Impl struct {
	client *client.DatabricksClient
}

func (a *jsonMarshallV2Impl) GetResource(ctx context.Context, request GetResourceRequest) (*Resource, error) {
	requestPb, pbErr := GetResourceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var resourcePb jsonmarshallv2pb.ResourcePb
	path := fmt.Sprintf("/api/2.0/json-marshall/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&resourcePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ResourceFromPb(&resourcePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

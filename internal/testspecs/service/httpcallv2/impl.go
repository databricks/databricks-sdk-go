// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package httpcallv2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/internal/testspecs/service/httpcallv2/httpcallv2pb"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just HttpCallV2 API methods
type httpCallV2Impl struct {
	client *client.DatabricksClient
}

func (a *httpCallV2Impl) CreateResource(ctx context.Context, request CreateResourceRequest) (*Resource, error) {
	requestPb, pbErr := CreateResourceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var resourcePb httpcallv2pb.ResourcePb
	path := fmt.Sprintf("/api/2.0/http-call/%v/%v/%v", request.PathParamString, request.PathParamInt, request.PathParamBool)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
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

func (a *httpCallV2Impl) GetResource(ctx context.Context, request GetResourceRequest) (*Resource, error) {
	requestPb, pbErr := GetResourceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var resourcePb httpcallv2pb.ResourcePb
	path := fmt.Sprintf("/api/2.0/http-call/%v/%v/%v", request.PathParamString, request.PathParamInt, request.PathParamBool)
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

func (a *httpCallV2Impl) UpdateResource(ctx context.Context, request UpdateResourceRequest) (*Resource, error) {
	requestPb, pbErr := UpdateResourceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var resourcePb httpcallv2pb.ResourcePb
	path := fmt.Sprintf("/api/2.0/http-call/%v/%v/%v", request.NestedPathParamString, request.NestedPathParamInt, request.NestedPathParamBool)
	queryParams := make(map[string]any)
	if requestPb.FieldMask != "" || slices.Contains(requestPb.ForceSendFields, "FieldMask") {
		queryParams["field_mask"] = requestPb.FieldMask
	}
	if requestPb.OptionalComplexQueryParam != nil || slices.Contains(requestPb.ForceSendFields, "OptionalComplexQueryParam") {
		queryParams["optional_complex_query_param"] = requestPb.OptionalComplexQueryParam
	}
	if requestPb.QueryParamBool != false || slices.Contains(requestPb.ForceSendFields, "QueryParamBool") {
		queryParams["query_param_bool"] = requestPb.QueryParamBool
	}
	if requestPb.QueryParamInt != 0 || slices.Contains(requestPb.ForceSendFields, "QueryParamInt") {
		queryParams["query_param_int"] = requestPb.QueryParamInt
	}
	if requestPb.QueryParamString != "" || slices.Contains(requestPb.ForceSendFields, "QueryParamString") {
		queryParams["query_param_string"] = requestPb.QueryParamString
	}
	if requestPb.RepeatedComplexQueryParam != nil || slices.Contains(requestPb.ForceSendFields, "RepeatedComplexQueryParam") {
		queryParams["repeated_complex_query_param"] = requestPb.RepeatedComplexQueryParam
	}
	if requestPb.RepeatedQueryParam != nil || slices.Contains(requestPb.ForceSendFields, "RepeatedQueryParam") {
		queryParams["repeated_query_param"] = requestPb.RepeatedQueryParam
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb).Resource,
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

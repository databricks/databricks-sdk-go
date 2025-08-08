// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package qualitymonitorv2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/qualitymonitorv2/qualitymonitorv2pb"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just QualityMonitorV2 API methods
type qualityMonitorV2Impl struct {
	client *client.DatabricksClient
}

func (a *qualityMonitorV2Impl) CreateQualityMonitor(ctx context.Context, request CreateQualityMonitorRequest) (*QualityMonitor, error) {
	requestPb, pbErr := CreateQualityMonitorRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var qualityMonitorPb qualitymonitorv2pb.QualityMonitorPb
	path := "/api/2.0/quality-monitors"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb.QualityMonitor,
		&qualityMonitorPb,
	)
	resp, err := QualityMonitorFromPb(&qualityMonitorPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *qualityMonitorV2Impl) DeleteQualityMonitor(ctx context.Context, request DeleteQualityMonitorRequest) error {
	requestPb, pbErr := DeleteQualityMonitorRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/quality-monitors/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *qualityMonitorV2Impl) GetQualityMonitor(ctx context.Context, request GetQualityMonitorRequest) (*QualityMonitor, error) {
	requestPb, pbErr := GetQualityMonitorRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var qualityMonitorPb qualitymonitorv2pb.QualityMonitorPb
	path := fmt.Sprintf("/api/2.0/quality-monitors/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&qualityMonitorPb,
	)
	resp, err := QualityMonitorFromPb(&qualityMonitorPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// (Unimplemented) List quality monitors
func (a *qualityMonitorV2Impl) ListQualityMonitor(ctx context.Context, request ListQualityMonitorRequest) listing.Iterator[QualityMonitor] {

	getNextPage := func(ctx context.Context, req ListQualityMonitorRequest) (*ListQualityMonitorResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListQualityMonitor(ctx, req)
	}
	getItems := func(resp *ListQualityMonitorResponse) []QualityMonitor {
		return resp.QualityMonitors
	}
	getNextReq := func(resp *ListQualityMonitorResponse) *ListQualityMonitorRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// (Unimplemented) List quality monitors
func (a *qualityMonitorV2Impl) ListQualityMonitorAll(ctx context.Context, request ListQualityMonitorRequest) ([]QualityMonitor, error) {
	iterator := a.ListQualityMonitor(ctx, request)
	return listing.ToSlice[QualityMonitor](ctx, iterator)
}

func (a *qualityMonitorV2Impl) internalListQualityMonitor(ctx context.Context, request ListQualityMonitorRequest) (*ListQualityMonitorResponse, error) {
	requestPb, pbErr := ListQualityMonitorRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listQualityMonitorResponsePb qualitymonitorv2pb.ListQualityMonitorResponsePb
	path := "/api/2.0/quality-monitors"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listQualityMonitorResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListQualityMonitorResponseFromPb(&listQualityMonitorResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *qualityMonitorV2Impl) UpdateQualityMonitor(ctx context.Context, request UpdateQualityMonitorRequest) (*QualityMonitor, error) {
	requestPb, pbErr := UpdateQualityMonitorRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var qualityMonitorPb qualitymonitorv2pb.QualityMonitorPb
	path := fmt.Sprintf("/api/2.0/quality-monitors/%v/%v", request.ObjectType, request.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb.QualityMonitor,
		&qualityMonitorPb,
	)
	resp, err := QualityMonitorFromPb(&qualityMonitorPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

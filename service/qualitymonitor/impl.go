// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package qualitymonitor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just QualityMonitorsV2 API methods
type qualityMonitorsV2Impl struct {
	client *client.DatabricksClient
}

func (a *qualityMonitorsV2Impl) CreateQualityMonitor(ctx context.Context, request CreateQualityMonitorRequest) (*QualityMonitor, error) {

	requestPb, pbErr := createQualityMonitorRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var qualityMonitorPb qualityMonitorPb
	path := "/api/2.0/quality-monitors"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb).QualityMonitor,
		&qualityMonitorPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := qualityMonitorFromPb(&qualityMonitorPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *qualityMonitorsV2Impl) DeleteQualityMonitor(ctx context.Context, request DeleteQualityMonitorRequest) error {

	requestPb, pbErr := deleteQualityMonitorRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteQualityMonitorResponsePb deleteQualityMonitorResponsePb
	path := fmt.Sprintf("/api/2.0/quality-monitors/%v/%v", requestPb.ObjectType, requestPb.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteQualityMonitorResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *qualityMonitorsV2Impl) GetQualityMonitor(ctx context.Context, request GetQualityMonitorRequest) (*QualityMonitor, error) {

	requestPb, pbErr := getQualityMonitorRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var qualityMonitorPb qualityMonitorPb
	path := fmt.Sprintf("/api/2.0/quality-monitors/%v/%v", requestPb.ObjectType, requestPb.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&qualityMonitorPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := qualityMonitorFromPb(&qualityMonitorPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List quality monitors.
//
// (Unimplemented) List quality monitors
func (a *qualityMonitorsV2Impl) ListQualityMonitor(ctx context.Context, request ListQualityMonitorRequest) listing.Iterator[QualityMonitor] {

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

// List quality monitors.
//
// (Unimplemented) List quality monitors
func (a *qualityMonitorsV2Impl) ListQualityMonitorAll(ctx context.Context, request ListQualityMonitorRequest) ([]QualityMonitor, error) {
	iterator := a.ListQualityMonitor(ctx, request)
	return listing.ToSlice[QualityMonitor](ctx, iterator)
}

func (a *qualityMonitorsV2Impl) internalListQualityMonitor(ctx context.Context, request ListQualityMonitorRequest) (*ListQualityMonitorResponse, error) {

	requestPb, pbErr := listQualityMonitorRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listQualityMonitorResponsePb listQualityMonitorResponsePb
	path := "/api/2.0/quality-monitors"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listQualityMonitorResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listQualityMonitorResponseFromPb(&listQualityMonitorResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *qualityMonitorsV2Impl) UpdateQualityMonitor(ctx context.Context, request UpdateQualityMonitorRequest) (*QualityMonitor, error) {

	requestPb, pbErr := updateQualityMonitorRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var qualityMonitorPb qualityMonitorPb
	path := fmt.Sprintf("/api/2.0/quality-monitors/%v/%v", requestPb.ObjectType, requestPb.ObjectId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb).QualityMonitor,
		&qualityMonitorPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := qualityMonitorFromPb(&qualityMonitorPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

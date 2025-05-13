// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type ColumnInfo struct {
	// Name of the column.
	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
}

func columnInfoToPb(st *ColumnInfo) (*columnInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &columnInfoPb{}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ColumnInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &columnInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := columnInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ColumnInfo) MarshalJSON() ([]byte, error) {
	pb, err := columnInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type columnInfoPb struct {
	// Name of the column.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func columnInfoFromPb(pb *columnInfoPb) (*ColumnInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ColumnInfo{}
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *columnInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st columnInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateEndpoint struct {
	// The budget policy id to be applied
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string
	// Type of endpoint
	// Wire name: 'endpoint_type'
	EndpointType EndpointType
	// Name of the vector search endpoint
	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
}

func createEndpointToPb(st *CreateEndpoint) (*createEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createEndpointPb{}
	pb.BudgetPolicyId = st.BudgetPolicyId

	pb.EndpointType = st.EndpointType

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateEndpoint) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createEndpointPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createEndpointFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateEndpoint) MarshalJSON() ([]byte, error) {
	pb, err := createEndpointToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createEndpointPb struct {
	// The budget policy id to be applied
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// Type of endpoint
	EndpointType EndpointType `json:"endpoint_type"`
	// Name of the vector search endpoint
	Name string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createEndpointFromPb(pb *createEndpointPb) (*CreateEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateEndpoint{}
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.EndpointType = pb.EndpointType
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createEndpointPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createEndpointPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateVectorIndexRequest struct {
	// Specification for Delta Sync Index. Required if `index_type` is
	// `DELTA_SYNC`.
	// Wire name: 'delta_sync_index_spec'
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecRequest
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	// Wire name: 'direct_access_index_spec'
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec
	// Name of the endpoint to be used for serving the index
	// Wire name: 'endpoint_name'
	EndpointName string
	// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
	// automatically syncs with a source Delta Table, automatically and
	// incrementally updating the index as the underlying data in the Delta
	// Table changes. - `DIRECT_ACCESS`: An index that supports direct read and
	// write of vectors and metadata through our REST and SDK APIs. With this
	// model, the user manages index updates.
	// Wire name: 'index_type'
	IndexType VectorIndexType
	// Name of the index
	// Wire name: 'name'
	Name string
	// Primary key of the index
	// Wire name: 'primary_key'
	PrimaryKey string
}

func createVectorIndexRequestToPb(st *CreateVectorIndexRequest) (*createVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createVectorIndexRequestPb{}
	deltaSyncIndexSpecPb, err := deltaSyncVectorIndexSpecRequestToPb(st.DeltaSyncIndexSpec)
	if err != nil {
		return nil, err
	}
	if deltaSyncIndexSpecPb != nil {
		pb.DeltaSyncIndexSpec = deltaSyncIndexSpecPb
	}

	directAccessIndexSpecPb, err := directAccessVectorIndexSpecToPb(st.DirectAccessIndexSpec)
	if err != nil {
		return nil, err
	}
	if directAccessIndexSpecPb != nil {
		pb.DirectAccessIndexSpec = directAccessIndexSpecPb
	}

	pb.EndpointName = st.EndpointName

	pb.IndexType = st.IndexType

	pb.Name = st.Name

	pb.PrimaryKey = st.PrimaryKey

	return pb, nil
}

func (st *CreateVectorIndexRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createVectorIndexRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createVectorIndexRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateVectorIndexRequest) MarshalJSON() ([]byte, error) {
	pb, err := createVectorIndexRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createVectorIndexRequestPb struct {
	// Specification for Delta Sync Index. Required if `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec *deltaSyncVectorIndexSpecRequestPb `json:"delta_sync_index_spec,omitempty"`
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec *directAccessVectorIndexSpecPb `json:"direct_access_index_spec,omitempty"`
	// Name of the endpoint to be used for serving the index
	EndpointName string `json:"endpoint_name"`
	// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
	// automatically syncs with a source Delta Table, automatically and
	// incrementally updating the index as the underlying data in the Delta
	// Table changes. - `DIRECT_ACCESS`: An index that supports direct read and
	// write of vectors and metadata through our REST and SDK APIs. With this
	// model, the user manages index updates.
	IndexType VectorIndexType `json:"index_type"`
	// Name of the index
	Name string `json:"name"`
	// Primary key of the index
	PrimaryKey string `json:"primary_key"`
}

func createVectorIndexRequestFromPb(pb *createVectorIndexRequestPb) (*CreateVectorIndexRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVectorIndexRequest{}
	deltaSyncIndexSpecField, err := deltaSyncVectorIndexSpecRequestFromPb(pb.DeltaSyncIndexSpec)
	if err != nil {
		return nil, err
	}
	if deltaSyncIndexSpecField != nil {
		st.DeltaSyncIndexSpec = deltaSyncIndexSpecField
	}
	directAccessIndexSpecField, err := directAccessVectorIndexSpecFromPb(pb.DirectAccessIndexSpec)
	if err != nil {
		return nil, err
	}
	if directAccessIndexSpecField != nil {
		st.DirectAccessIndexSpec = directAccessIndexSpecField
	}
	st.EndpointName = pb.EndpointName
	st.IndexType = pb.IndexType
	st.Name = pb.Name
	st.PrimaryKey = pb.PrimaryKey

	return st, nil
}

type CustomTag struct {
	// Key field for a vector search endpoint tag.
	// Wire name: 'key'
	Key string
	// [Optional] Value field for a vector search endpoint tag.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func customTagToPb(st *CustomTag) (*customTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customTagPb{}
	pb.Key = st.Key

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CustomTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &customTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := customTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CustomTag) MarshalJSON() ([]byte, error) {
	pb, err := customTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type customTagPb struct {
	// Key field for a vector search endpoint tag.
	Key string `json:"key"`
	// [Optional] Value field for a vector search endpoint tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func customTagFromPb(pb *customTagPb) (*CustomTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *customTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st customTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDataResult struct {
	// List of primary keys for rows that failed to process.
	// Wire name: 'failed_primary_keys'
	FailedPrimaryKeys []string
	// Count of successfully processed rows.
	// Wire name: 'success_row_count'
	SuccessRowCount int64

	ForceSendFields []string `tf:"-"`
}

func deleteDataResultToPb(st *DeleteDataResult) (*deleteDataResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDataResultPb{}
	pb.FailedPrimaryKeys = st.FailedPrimaryKeys

	pb.SuccessRowCount = st.SuccessRowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteDataResult) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDataResultPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDataResultFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDataResult) MarshalJSON() ([]byte, error) {
	pb, err := deleteDataResultToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteDataResultPb struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	// Count of successfully processed rows.
	SuccessRowCount int64 `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteDataResultFromPb(pb *deleteDataResultPb) (*DeleteDataResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDataResult{}
	st.FailedPrimaryKeys = pb.FailedPrimaryKeys
	st.SuccessRowCount = pb.SuccessRowCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteDataResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteDataResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDataStatus string
type deleteDataStatusPb string

const DeleteDataStatusFailure DeleteDataStatus = `FAILURE`

const DeleteDataStatusPartialSuccess DeleteDataStatus = `PARTIAL_SUCCESS`

const DeleteDataStatusSuccess DeleteDataStatus = `SUCCESS`

// String representation for [fmt.Print]
func (f *DeleteDataStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeleteDataStatus) Set(v string) error {
	switch v {
	case `FAILURE`, `PARTIAL_SUCCESS`, `SUCCESS`:
		*f = DeleteDataStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILURE", "PARTIAL_SUCCESS", "SUCCESS"`, v)
	}
}

// Type always returns DeleteDataStatus to satisfy [pflag.Value] interface
func (f *DeleteDataStatus) Type() string {
	return "DeleteDataStatus"
}

func deleteDataStatusToPb(st *DeleteDataStatus) (*deleteDataStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := deleteDataStatusPb(*st)
	return &pb, nil
}

func deleteDataStatusFromPb(pb *deleteDataStatusPb) (*DeleteDataStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := DeleteDataStatus(*pb)
	return &st, nil
}

// Delete data from index
type DeleteDataVectorIndexRequest struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
	// List of primary keys for the data to be deleted.
	// Wire name: 'primary_keys'
	PrimaryKeys []string `tf:"-"`
}

func deleteDataVectorIndexRequestToPb(st *DeleteDataVectorIndexRequest) (*deleteDataVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDataVectorIndexRequestPb{}
	pb.IndexName = st.IndexName

	pb.PrimaryKeys = st.PrimaryKeys

	return pb, nil
}

func (st *DeleteDataVectorIndexRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDataVectorIndexRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDataVectorIndexRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDataVectorIndexRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDataVectorIndexRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteDataVectorIndexRequestPb struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	IndexName string `json:"-" url:"-"`
	// List of primary keys for the data to be deleted.
	PrimaryKeys []string `json:"-" url:"primary_keys"`
}

func deleteDataVectorIndexRequestFromPb(pb *deleteDataVectorIndexRequestPb) (*DeleteDataVectorIndexRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDataVectorIndexRequest{}
	st.IndexName = pb.IndexName
	st.PrimaryKeys = pb.PrimaryKeys

	return st, nil
}

type DeleteDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	// Wire name: 'result'
	Result *DeleteDataResult
	// Status of the delete operation.
	// Wire name: 'status'
	Status DeleteDataStatus
}

func deleteDataVectorIndexResponseToPb(st *DeleteDataVectorIndexResponse) (*deleteDataVectorIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDataVectorIndexResponsePb{}
	resultPb, err := deleteDataResultToPb(st.Result)
	if err != nil {
		return nil, err
	}
	if resultPb != nil {
		pb.Result = resultPb
	}

	pb.Status = st.Status

	return pb, nil
}

func (st *DeleteDataVectorIndexResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDataVectorIndexResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDataVectorIndexResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDataVectorIndexResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteDataVectorIndexResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteDataVectorIndexResponsePb struct {
	// Result of the upsert or delete operation.
	Result *deleteDataResultPb `json:"result,omitempty"`
	// Status of the delete operation.
	Status DeleteDataStatus `json:"status,omitempty"`
}

func deleteDataVectorIndexResponseFromPb(pb *deleteDataVectorIndexResponsePb) (*DeleteDataVectorIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDataVectorIndexResponse{}
	resultField, err := deleteDataResultFromPb(pb.Result)
	if err != nil {
		return nil, err
	}
	if resultField != nil {
		st.Result = resultField
	}
	st.Status = pb.Status

	return st, nil
}

// Delete an endpoint
type DeleteEndpointRequest struct {
	// Name of the vector search endpoint
	// Wire name: 'endpoint_name'
	EndpointName string `tf:"-"`
}

func deleteEndpointRequestToPb(st *DeleteEndpointRequest) (*deleteEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteEndpointRequestPb{}
	pb.EndpointName = st.EndpointName

	return pb, nil
}

func (st *DeleteEndpointRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteEndpointRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteEndpointRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteEndpointRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteEndpointRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteEndpointRequestPb struct {
	// Name of the vector search endpoint
	EndpointName string `json:"-" url:"-"`
}

func deleteEndpointRequestFromPb(pb *deleteEndpointRequestPb) (*DeleteEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteEndpointRequest{}
	st.EndpointName = pb.EndpointName

	return st, nil
}

type DeleteEndpointResponse struct {
}

func deleteEndpointResponseToPb(st *DeleteEndpointResponse) (*deleteEndpointResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteEndpointResponsePb{}

	return pb, nil
}

func (st *DeleteEndpointResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteEndpointResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteEndpointResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteEndpointResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteEndpointResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteEndpointResponsePb struct {
}

func deleteEndpointResponseFromPb(pb *deleteEndpointResponsePb) (*DeleteEndpointResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteEndpointResponse{}

	return st, nil
}

// Delete an index
type DeleteIndexRequest struct {
	// Name of the index
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
}

func deleteIndexRequestToPb(st *DeleteIndexRequest) (*deleteIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteIndexRequestPb{}
	pb.IndexName = st.IndexName

	return pb, nil
}

func (st *DeleteIndexRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteIndexRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteIndexRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteIndexRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteIndexRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteIndexRequestPb struct {
	// Name of the index
	IndexName string `json:"-" url:"-"`
}

func deleteIndexRequestFromPb(pb *deleteIndexRequestPb) (*DeleteIndexRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteIndexRequest{}
	st.IndexName = pb.IndexName

	return st, nil
}

type DeleteIndexResponse struct {
}

func deleteIndexResponseToPb(st *DeleteIndexResponse) (*deleteIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteIndexResponsePb{}

	return pb, nil
}

func (st *DeleteIndexResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteIndexResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteIndexResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteIndexResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteIndexResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteIndexResponsePb struct {
}

func deleteIndexResponseFromPb(pb *deleteIndexResponsePb) (*DeleteIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteIndexResponse{}

	return st, nil
}

type DeltaSyncVectorIndexSpecRequest struct {
	// [Optional] Select the columns to sync with the vector index. If you leave
	// this field blank, all columns from the source table are synced with the
	// index. The primary key column and embedding source column or embedding
	// vector column are always synced.
	// Wire name: 'columns_to_sync'
	ColumnsToSync []string
	// The columns that contain the embedding source.
	// Wire name: 'embedding_source_columns'
	EmbeddingSourceColumns []EmbeddingSourceColumn
	// The columns that contain the embedding vectors.
	// Wire name: 'embedding_vector_columns'
	EmbeddingVectorColumns []EmbeddingVectorColumn
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	// Wire name: 'embedding_writeback_table'
	EmbeddingWritebackTable string
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	// Wire name: 'pipeline_type'
	PipelineType PipelineType
	// The name of the source table.
	// Wire name: 'source_table'
	SourceTable string

	ForceSendFields []string `tf:"-"`
}

func deltaSyncVectorIndexSpecRequestToPb(st *DeltaSyncVectorIndexSpecRequest) (*deltaSyncVectorIndexSpecRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSyncVectorIndexSpecRequestPb{}
	pb.ColumnsToSync = st.ColumnsToSync

	var embeddingSourceColumnsPb []embeddingSourceColumnPb
	for _, item := range st.EmbeddingSourceColumns {
		itemPb, err := embeddingSourceColumnToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			embeddingSourceColumnsPb = append(embeddingSourceColumnsPb, *itemPb)
		}
	}
	pb.EmbeddingSourceColumns = embeddingSourceColumnsPb

	var embeddingVectorColumnsPb []embeddingVectorColumnPb
	for _, item := range st.EmbeddingVectorColumns {
		itemPb, err := embeddingVectorColumnToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			embeddingVectorColumnsPb = append(embeddingVectorColumnsPb, *itemPb)
		}
	}
	pb.EmbeddingVectorColumns = embeddingVectorColumnsPb

	pb.EmbeddingWritebackTable = st.EmbeddingWritebackTable

	pb.PipelineType = st.PipelineType

	pb.SourceTable = st.SourceTable

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeltaSyncVectorIndexSpecRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deltaSyncVectorIndexSpecRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deltaSyncVectorIndexSpecRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeltaSyncVectorIndexSpecRequest) MarshalJSON() ([]byte, error) {
	pb, err := deltaSyncVectorIndexSpecRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deltaSyncVectorIndexSpecRequestPb struct {
	// [Optional] Select the columns to sync with the vector index. If you leave
	// this field blank, all columns from the source table are synced with the
	// index. The primary key column and embedding source column or embedding
	// vector column are always synced.
	ColumnsToSync []string `json:"columns_to_sync,omitempty"`
	// The columns that contain the embedding source.
	EmbeddingSourceColumns []embeddingSourceColumnPb `json:"embedding_source_columns,omitempty"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns []embeddingVectorColumnPb `json:"embedding_vector_columns,omitempty"`
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	EmbeddingWritebackTable string `json:"embedding_writeback_table,omitempty"`
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	PipelineType PipelineType `json:"pipeline_type,omitempty"`
	// The name of the source table.
	SourceTable string `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deltaSyncVectorIndexSpecRequestFromPb(pb *deltaSyncVectorIndexSpecRequestPb) (*DeltaSyncVectorIndexSpecRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSyncVectorIndexSpecRequest{}
	st.ColumnsToSync = pb.ColumnsToSync

	var embeddingSourceColumnsField []EmbeddingSourceColumn
	for _, itemPb := range pb.EmbeddingSourceColumns {
		item, err := embeddingSourceColumnFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			embeddingSourceColumnsField = append(embeddingSourceColumnsField, *item)
		}
	}
	st.EmbeddingSourceColumns = embeddingSourceColumnsField

	var embeddingVectorColumnsField []EmbeddingVectorColumn
	for _, itemPb := range pb.EmbeddingVectorColumns {
		item, err := embeddingVectorColumnFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			embeddingVectorColumnsField = append(embeddingVectorColumnsField, *item)
		}
	}
	st.EmbeddingVectorColumns = embeddingVectorColumnsField
	st.EmbeddingWritebackTable = pb.EmbeddingWritebackTable
	st.PipelineType = pb.PipelineType
	st.SourceTable = pb.SourceTable

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deltaSyncVectorIndexSpecRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deltaSyncVectorIndexSpecRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeltaSyncVectorIndexSpecResponse struct {
	// The columns that contain the embedding source.
	// Wire name: 'embedding_source_columns'
	EmbeddingSourceColumns []EmbeddingSourceColumn
	// The columns that contain the embedding vectors.
	// Wire name: 'embedding_vector_columns'
	EmbeddingVectorColumns []EmbeddingVectorColumn
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	// Wire name: 'embedding_writeback_table'
	EmbeddingWritebackTable string
	// The ID of the pipeline that is used to sync the index.
	// Wire name: 'pipeline_id'
	PipelineId string
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	// Wire name: 'pipeline_type'
	PipelineType PipelineType
	// The name of the source table.
	// Wire name: 'source_table'
	SourceTable string

	ForceSendFields []string `tf:"-"`
}

func deltaSyncVectorIndexSpecResponseToPb(st *DeltaSyncVectorIndexSpecResponse) (*deltaSyncVectorIndexSpecResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSyncVectorIndexSpecResponsePb{}

	var embeddingSourceColumnsPb []embeddingSourceColumnPb
	for _, item := range st.EmbeddingSourceColumns {
		itemPb, err := embeddingSourceColumnToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			embeddingSourceColumnsPb = append(embeddingSourceColumnsPb, *itemPb)
		}
	}
	pb.EmbeddingSourceColumns = embeddingSourceColumnsPb

	var embeddingVectorColumnsPb []embeddingVectorColumnPb
	for _, item := range st.EmbeddingVectorColumns {
		itemPb, err := embeddingVectorColumnToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			embeddingVectorColumnsPb = append(embeddingVectorColumnsPb, *itemPb)
		}
	}
	pb.EmbeddingVectorColumns = embeddingVectorColumnsPb

	pb.EmbeddingWritebackTable = st.EmbeddingWritebackTable

	pb.PipelineId = st.PipelineId

	pb.PipelineType = st.PipelineType

	pb.SourceTable = st.SourceTable

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeltaSyncVectorIndexSpecResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deltaSyncVectorIndexSpecResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deltaSyncVectorIndexSpecResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeltaSyncVectorIndexSpecResponse) MarshalJSON() ([]byte, error) {
	pb, err := deltaSyncVectorIndexSpecResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deltaSyncVectorIndexSpecResponsePb struct {
	// The columns that contain the embedding source.
	EmbeddingSourceColumns []embeddingSourceColumnPb `json:"embedding_source_columns,omitempty"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns []embeddingVectorColumnPb `json:"embedding_vector_columns,omitempty"`
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	EmbeddingWritebackTable string `json:"embedding_writeback_table,omitempty"`
	// The ID of the pipeline that is used to sync the index.
	PipelineId string `json:"pipeline_id,omitempty"`
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	PipelineType PipelineType `json:"pipeline_type,omitempty"`
	// The name of the source table.
	SourceTable string `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deltaSyncVectorIndexSpecResponseFromPb(pb *deltaSyncVectorIndexSpecResponsePb) (*DeltaSyncVectorIndexSpecResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSyncVectorIndexSpecResponse{}

	var embeddingSourceColumnsField []EmbeddingSourceColumn
	for _, itemPb := range pb.EmbeddingSourceColumns {
		item, err := embeddingSourceColumnFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			embeddingSourceColumnsField = append(embeddingSourceColumnsField, *item)
		}
	}
	st.EmbeddingSourceColumns = embeddingSourceColumnsField

	var embeddingVectorColumnsField []EmbeddingVectorColumn
	for _, itemPb := range pb.EmbeddingVectorColumns {
		item, err := embeddingVectorColumnFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			embeddingVectorColumnsField = append(embeddingVectorColumnsField, *item)
		}
	}
	st.EmbeddingVectorColumns = embeddingVectorColumnsField
	st.EmbeddingWritebackTable = pb.EmbeddingWritebackTable
	st.PipelineId = pb.PipelineId
	st.PipelineType = pb.PipelineType
	st.SourceTable = pb.SourceTable

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deltaSyncVectorIndexSpecResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deltaSyncVectorIndexSpecResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DirectAccessVectorIndexSpec struct {
	// The columns that contain the embedding source. The format should be
	// array[double].
	// Wire name: 'embedding_source_columns'
	EmbeddingSourceColumns []EmbeddingSourceColumn
	// The columns that contain the embedding vectors. The format should be
	// array[double].
	// Wire name: 'embedding_vector_columns'
	EmbeddingVectorColumns []EmbeddingVectorColumn
	// The schema of the index in JSON format. Supported types are `integer`,
	// `long`, `float`, `double`, `boolean`, `string`, `date`, `timestamp`.
	// Supported types for vector column: `array<float>`, `array<double>`,`.
	// Wire name: 'schema_json'
	SchemaJson string

	ForceSendFields []string `tf:"-"`
}

func directAccessVectorIndexSpecToPb(st *DirectAccessVectorIndexSpec) (*directAccessVectorIndexSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &directAccessVectorIndexSpecPb{}

	var embeddingSourceColumnsPb []embeddingSourceColumnPb
	for _, item := range st.EmbeddingSourceColumns {
		itemPb, err := embeddingSourceColumnToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			embeddingSourceColumnsPb = append(embeddingSourceColumnsPb, *itemPb)
		}
	}
	pb.EmbeddingSourceColumns = embeddingSourceColumnsPb

	var embeddingVectorColumnsPb []embeddingVectorColumnPb
	for _, item := range st.EmbeddingVectorColumns {
		itemPb, err := embeddingVectorColumnToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			embeddingVectorColumnsPb = append(embeddingVectorColumnsPb, *itemPb)
		}
	}
	pb.EmbeddingVectorColumns = embeddingVectorColumnsPb

	pb.SchemaJson = st.SchemaJson

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DirectAccessVectorIndexSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &directAccessVectorIndexSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := directAccessVectorIndexSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DirectAccessVectorIndexSpec) MarshalJSON() ([]byte, error) {
	pb, err := directAccessVectorIndexSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type directAccessVectorIndexSpecPb struct {
	// The columns that contain the embedding source. The format should be
	// array[double].
	EmbeddingSourceColumns []embeddingSourceColumnPb `json:"embedding_source_columns,omitempty"`
	// The columns that contain the embedding vectors. The format should be
	// array[double].
	EmbeddingVectorColumns []embeddingVectorColumnPb `json:"embedding_vector_columns,omitempty"`
	// The schema of the index in JSON format. Supported types are `integer`,
	// `long`, `float`, `double`, `boolean`, `string`, `date`, `timestamp`.
	// Supported types for vector column: `array<float>`, `array<double>`,`.
	SchemaJson string `json:"schema_json,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func directAccessVectorIndexSpecFromPb(pb *directAccessVectorIndexSpecPb) (*DirectAccessVectorIndexSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DirectAccessVectorIndexSpec{}

	var embeddingSourceColumnsField []EmbeddingSourceColumn
	for _, itemPb := range pb.EmbeddingSourceColumns {
		item, err := embeddingSourceColumnFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			embeddingSourceColumnsField = append(embeddingSourceColumnsField, *item)
		}
	}
	st.EmbeddingSourceColumns = embeddingSourceColumnsField

	var embeddingVectorColumnsField []EmbeddingVectorColumn
	for _, itemPb := range pb.EmbeddingVectorColumns {
		item, err := embeddingVectorColumnFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			embeddingVectorColumnsField = append(embeddingVectorColumnsField, *item)
		}
	}
	st.EmbeddingVectorColumns = embeddingVectorColumnsField
	st.SchemaJson = pb.SchemaJson

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *directAccessVectorIndexSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st directAccessVectorIndexSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EmbeddingSourceColumn struct {
	// Name of the embedding model endpoint
	// Wire name: 'embedding_model_endpoint_name'
	EmbeddingModelEndpointName string
	// Name of the column
	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
}

func embeddingSourceColumnToPb(st *EmbeddingSourceColumn) (*embeddingSourceColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &embeddingSourceColumnPb{}
	pb.EmbeddingModelEndpointName = st.EmbeddingModelEndpointName

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *EmbeddingSourceColumn) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &embeddingSourceColumnPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := embeddingSourceColumnFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EmbeddingSourceColumn) MarshalJSON() ([]byte, error) {
	pb, err := embeddingSourceColumnToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type embeddingSourceColumnPb struct {
	// Name of the embedding model endpoint
	EmbeddingModelEndpointName string `json:"embedding_model_endpoint_name,omitempty"`
	// Name of the column
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func embeddingSourceColumnFromPb(pb *embeddingSourceColumnPb) (*EmbeddingSourceColumn, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EmbeddingSourceColumn{}
	st.EmbeddingModelEndpointName = pb.EmbeddingModelEndpointName
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *embeddingSourceColumnPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st embeddingSourceColumnPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EmbeddingVectorColumn struct {
	// Dimension of the embedding vector
	// Wire name: 'embedding_dimension'
	EmbeddingDimension int
	// Name of the column
	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
}

func embeddingVectorColumnToPb(st *EmbeddingVectorColumn) (*embeddingVectorColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &embeddingVectorColumnPb{}
	pb.EmbeddingDimension = st.EmbeddingDimension

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *EmbeddingVectorColumn) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &embeddingVectorColumnPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := embeddingVectorColumnFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EmbeddingVectorColumn) MarshalJSON() ([]byte, error) {
	pb, err := embeddingVectorColumnToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type embeddingVectorColumnPb struct {
	// Dimension of the embedding vector
	EmbeddingDimension int `json:"embedding_dimension,omitempty"`
	// Name of the column
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func embeddingVectorColumnFromPb(pb *embeddingVectorColumnPb) (*EmbeddingVectorColumn, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EmbeddingVectorColumn{}
	st.EmbeddingDimension = pb.EmbeddingDimension
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *embeddingVectorColumnPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st embeddingVectorColumnPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointInfo struct {
	// Timestamp of endpoint creation
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// Creator of the endpoint
	// Wire name: 'creator'
	Creator string
	// The custom tags assigned to the endpoint
	// Wire name: 'custom_tags'
	CustomTags []CustomTag
	// The budget policy id applied to the endpoint
	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string
	// Current status of the endpoint
	// Wire name: 'endpoint_status'
	EndpointStatus *EndpointStatus
	// Type of endpoint
	// Wire name: 'endpoint_type'
	EndpointType EndpointType
	// Unique identifier of the endpoint
	// Wire name: 'id'
	Id string
	// Timestamp of last update to the endpoint
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// User who last updated the endpoint
	// Wire name: 'last_updated_user'
	LastUpdatedUser string
	// Name of the vector search endpoint
	// Wire name: 'name'
	Name string
	// Number of indexes on the endpoint
	// Wire name: 'num_indexes'
	NumIndexes int

	ForceSendFields []string `tf:"-"`
}

func endpointInfoToPb(st *EndpointInfo) (*endpointInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointInfoPb{}
	pb.CreationTimestamp = st.CreationTimestamp

	pb.Creator = st.Creator

	var customTagsPb []customTagPb
	for _, item := range st.CustomTags {
		itemPb, err := customTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb = append(customTagsPb, *itemPb)
		}
	}
	pb.CustomTags = customTagsPb

	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId

	endpointStatusPb, err := endpointStatusToPb(st.EndpointStatus)
	if err != nil {
		return nil, err
	}
	if endpointStatusPb != nil {
		pb.EndpointStatus = endpointStatusPb
	}

	pb.EndpointType = st.EndpointType

	pb.Id = st.Id

	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	pb.LastUpdatedUser = st.LastUpdatedUser

	pb.Name = st.Name

	pb.NumIndexes = st.NumIndexes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *EndpointInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointInfo) MarshalJSON() ([]byte, error) {
	pb, err := endpointInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type endpointInfoPb struct {
	// Timestamp of endpoint creation
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Creator of the endpoint
	Creator string `json:"creator,omitempty"`
	// The custom tags assigned to the endpoint
	CustomTags []customTagPb `json:"custom_tags,omitempty"`
	// The budget policy id applied to the endpoint
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// Current status of the endpoint
	EndpointStatus *endpointStatusPb `json:"endpoint_status,omitempty"`
	// Type of endpoint
	EndpointType EndpointType `json:"endpoint_type,omitempty"`
	// Unique identifier of the endpoint
	Id string `json:"id,omitempty"`
	// Timestamp of last update to the endpoint
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// User who last updated the endpoint
	LastUpdatedUser string `json:"last_updated_user,omitempty"`
	// Name of the vector search endpoint
	Name string `json:"name,omitempty"`
	// Number of indexes on the endpoint
	NumIndexes int `json:"num_indexes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointInfoFromPb(pb *endpointInfoPb) (*EndpointInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointInfo{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Creator = pb.Creator

	var customTagsField []CustomTag
	for _, itemPb := range pb.CustomTags {
		item, err := customTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField = append(customTagsField, *item)
		}
	}
	st.CustomTags = customTagsField
	st.EffectiveBudgetPolicyId = pb.EffectiveBudgetPolicyId
	endpointStatusField, err := endpointStatusFromPb(pb.EndpointStatus)
	if err != nil {
		return nil, err
	}
	if endpointStatusField != nil {
		st.EndpointStatus = endpointStatusField
	}
	st.EndpointType = pb.EndpointType
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.LastUpdatedUser = pb.LastUpdatedUser
	st.Name = pb.Name
	st.NumIndexes = pb.NumIndexes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *endpointInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Status information of an endpoint
type EndpointStatus struct {
	// Additional status message
	// Wire name: 'message'
	Message string
	// Current state of the endpoint
	// Wire name: 'state'
	State EndpointStatusState

	ForceSendFields []string `tf:"-"`
}

func endpointStatusToPb(st *EndpointStatus) (*endpointStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointStatusPb{}
	pb.Message = st.Message

	pb.State = st.State

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *EndpointStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointStatus) MarshalJSON() ([]byte, error) {
	pb, err := endpointStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type endpointStatusPb struct {
	// Additional status message
	Message string `json:"message,omitempty"`
	// Current state of the endpoint
	State EndpointStatusState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointStatusFromPb(pb *endpointStatusPb) (*EndpointStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointStatus{}
	st.Message = pb.Message
	st.State = pb.State

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *endpointStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Current state of the endpoint
type EndpointStatusState string
type endpointStatusStatePb string

const EndpointStatusStateOffline EndpointStatusState = `OFFLINE`

const EndpointStatusStateOnline EndpointStatusState = `ONLINE`

const EndpointStatusStateProvisioning EndpointStatusState = `PROVISIONING`

// String representation for [fmt.Print]
func (f *EndpointStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStatusState) Set(v string) error {
	switch v {
	case `OFFLINE`, `ONLINE`, `PROVISIONING`:
		*f = EndpointStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OFFLINE", "ONLINE", "PROVISIONING"`, v)
	}
}

// Type always returns EndpointStatusState to satisfy [pflag.Value] interface
func (f *EndpointStatusState) Type() string {
	return "EndpointStatusState"
}

func endpointStatusStateToPb(st *EndpointStatusState) (*endpointStatusStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := endpointStatusStatePb(*st)
	return &pb, nil
}

func endpointStatusStateFromPb(pb *endpointStatusStatePb) (*EndpointStatusState, error) {
	if pb == nil {
		return nil, nil
	}
	st := EndpointStatusState(*pb)
	return &st, nil
}

// Type of endpoint.
type EndpointType string
type endpointTypePb string

const EndpointTypeStandard EndpointType = `STANDARD`

// String representation for [fmt.Print]
func (f *EndpointType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointType) Set(v string) error {
	switch v {
	case `STANDARD`:
		*f = EndpointType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "STANDARD"`, v)
	}
}

// Type always returns EndpointType to satisfy [pflag.Value] interface
func (f *EndpointType) Type() string {
	return "EndpointType"
}

func endpointTypeToPb(st *EndpointType) (*endpointTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := endpointTypePb(*st)
	return &pb, nil
}

func endpointTypeFromPb(pb *endpointTypePb) (*EndpointType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EndpointType(*pb)
	return &st, nil
}

// Get an endpoint
type GetEndpointRequest struct {
	// Name of the endpoint
	// Wire name: 'endpoint_name'
	EndpointName string `tf:"-"`
}

func getEndpointRequestToPb(st *GetEndpointRequest) (*getEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEndpointRequestPb{}
	pb.EndpointName = st.EndpointName

	return pb, nil
}

func (st *GetEndpointRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getEndpointRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getEndpointRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetEndpointRequest) MarshalJSON() ([]byte, error) {
	pb, err := getEndpointRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getEndpointRequestPb struct {
	// Name of the endpoint
	EndpointName string `json:"-" url:"-"`
}

func getEndpointRequestFromPb(pb *getEndpointRequestPb) (*GetEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEndpointRequest{}
	st.EndpointName = pb.EndpointName

	return st, nil
}

// Get an index
type GetIndexRequest struct {
	// Name of the index
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
}

func getIndexRequestToPb(st *GetIndexRequest) (*getIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getIndexRequestPb{}
	pb.IndexName = st.IndexName

	return pb, nil
}

func (st *GetIndexRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getIndexRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getIndexRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetIndexRequest) MarshalJSON() ([]byte, error) {
	pb, err := getIndexRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getIndexRequestPb struct {
	// Name of the index
	IndexName string `json:"-" url:"-"`
}

func getIndexRequestFromPb(pb *getIndexRequestPb) (*GetIndexRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetIndexRequest{}
	st.IndexName = pb.IndexName

	return st, nil
}

type ListEndpointResponse struct {
	// An array of Endpoint objects
	// Wire name: 'endpoints'
	Endpoints []EndpointInfo
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listEndpointResponseToPb(st *ListEndpointResponse) (*listEndpointResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listEndpointResponsePb{}

	var endpointsPb []endpointInfoPb
	for _, item := range st.Endpoints {
		itemPb, err := endpointInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			endpointsPb = append(endpointsPb, *itemPb)
		}
	}
	pb.Endpoints = endpointsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListEndpointResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listEndpointResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listEndpointResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListEndpointResponse) MarshalJSON() ([]byte, error) {
	pb, err := listEndpointResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listEndpointResponsePb struct {
	// An array of Endpoint objects
	Endpoints []endpointInfoPb `json:"endpoints,omitempty"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listEndpointResponseFromPb(pb *listEndpointResponsePb) (*ListEndpointResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListEndpointResponse{}

	var endpointsField []EndpointInfo
	for _, itemPb := range pb.Endpoints {
		item, err := endpointInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			endpointsField = append(endpointsField, *item)
		}
	}
	st.Endpoints = endpointsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listEndpointResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listEndpointResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List all endpoints
type ListEndpointsRequest struct {
	// Token for pagination
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listEndpointsRequestToPb(st *ListEndpointsRequest) (*listEndpointsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listEndpointsRequestPb{}
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListEndpointsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listEndpointsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listEndpointsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListEndpointsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listEndpointsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listEndpointsRequestPb struct {
	// Token for pagination
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listEndpointsRequestFromPb(pb *listEndpointsRequestPb) (*ListEndpointsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListEndpointsRequest{}
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listEndpointsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listEndpointsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List indexes
type ListIndexesRequest struct {
	// Name of the endpoint
	// Wire name: 'endpoint_name'
	EndpointName string `tf:"-"`
	// Token for pagination
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listIndexesRequestToPb(st *ListIndexesRequest) (*listIndexesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listIndexesRequestPb{}
	pb.EndpointName = st.EndpointName

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListIndexesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listIndexesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listIndexesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListIndexesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listIndexesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listIndexesRequestPb struct {
	// Name of the endpoint
	EndpointName string `json:"-" url:"endpoint_name"`
	// Token for pagination
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listIndexesRequestFromPb(pb *listIndexesRequestPb) (*ListIndexesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListIndexesRequest{}
	st.EndpointName = pb.EndpointName
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listIndexesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listIndexesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// copied from proto3 / Google Well Known Types, source:
// https://github.com/protocolbuffers/protobuf/blob/450d24ca820750c5db5112a6f0b0c2efb9758021/src/google/protobuf/struct.proto
// `ListValue` is a wrapper around a repeated field of values.
//
// The JSON representation for `ListValue` is JSON array.
type ListValue struct {
	// Repeated field of dynamically typed values.
	// Wire name: 'values'
	Values []Value
}

func listValueToPb(st *ListValue) (*listValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listValuePb{}

	var valuesPb []valuePb
	for _, item := range st.Values {
		itemPb, err := valueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			valuesPb = append(valuesPb, *itemPb)
		}
	}
	pb.Values = valuesPb

	return pb, nil
}

func (st *ListValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListValue) MarshalJSON() ([]byte, error) {
	pb, err := listValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listValuePb struct {
	// Repeated field of dynamically typed values.
	Values []valuePb `json:"values,omitempty"`
}

func listValueFromPb(pb *listValuePb) (*ListValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListValue{}

	var valuesField []Value
	for _, itemPb := range pb.Values {
		item, err := valueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			valuesField = append(valuesField, *item)
		}
	}
	st.Values = valuesField

	return st, nil
}

type ListVectorIndexesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'vector_indexes'
	VectorIndexes []MiniVectorIndex

	ForceSendFields []string `tf:"-"`
}

func listVectorIndexesResponseToPb(st *ListVectorIndexesResponse) (*listVectorIndexesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVectorIndexesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var vectorIndexesPb []miniVectorIndexPb
	for _, item := range st.VectorIndexes {
		itemPb, err := miniVectorIndexToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			vectorIndexesPb = append(vectorIndexesPb, *itemPb)
		}
	}
	pb.VectorIndexes = vectorIndexesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListVectorIndexesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listVectorIndexesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listVectorIndexesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListVectorIndexesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listVectorIndexesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listVectorIndexesResponsePb struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	VectorIndexes []miniVectorIndexPb `json:"vector_indexes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listVectorIndexesResponseFromPb(pb *listVectorIndexesResponsePb) (*ListVectorIndexesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListVectorIndexesResponse{}
	st.NextPageToken = pb.NextPageToken

	var vectorIndexesField []MiniVectorIndex
	for _, itemPb := range pb.VectorIndexes {
		item, err := miniVectorIndexFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			vectorIndexesField = append(vectorIndexesField, *item)
		}
	}
	st.VectorIndexes = vectorIndexesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listVectorIndexesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listVectorIndexesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Key-value pair.
type MapStringValueEntry struct {
	// Column name.
	// Wire name: 'key'
	Key string
	// Column value, nullable.
	// Wire name: 'value'
	Value *Value

	ForceSendFields []string `tf:"-"`
}

func mapStringValueEntryToPb(st *MapStringValueEntry) (*mapStringValueEntryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mapStringValueEntryPb{}
	pb.Key = st.Key

	valuePb, err := valueToPb(st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = valuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *MapStringValueEntry) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mapStringValueEntryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := mapStringValueEntryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MapStringValueEntry) MarshalJSON() ([]byte, error) {
	pb, err := mapStringValueEntryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type mapStringValueEntryPb struct {
	// Column name.
	Key string `json:"key,omitempty"`
	// Column value, nullable.
	Value *valuePb `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func mapStringValueEntryFromPb(pb *mapStringValueEntryPb) (*MapStringValueEntry, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MapStringValueEntry{}
	st.Key = pb.Key
	valueField, err := valueFromPb(pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = valueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *mapStringValueEntryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st mapStringValueEntryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MiniVectorIndex struct {
	// The user who created the index.
	// Wire name: 'creator'
	Creator string
	// Name of the endpoint associated with the index
	// Wire name: 'endpoint_name'
	EndpointName string
	// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
	// automatically syncs with a source Delta Table, automatically and
	// incrementally updating the index as the underlying data in the Delta
	// Table changes. - `DIRECT_ACCESS`: An index that supports direct read and
	// write of vectors and metadata through our REST and SDK APIs. With this
	// model, the user manages index updates.
	// Wire name: 'index_type'
	IndexType VectorIndexType
	// Name of the index
	// Wire name: 'name'
	Name string
	// Primary key of the index
	// Wire name: 'primary_key'
	PrimaryKey string

	ForceSendFields []string `tf:"-"`
}

func miniVectorIndexToPb(st *MiniVectorIndex) (*miniVectorIndexPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &miniVectorIndexPb{}
	pb.Creator = st.Creator

	pb.EndpointName = st.EndpointName

	pb.IndexType = st.IndexType

	pb.Name = st.Name

	pb.PrimaryKey = st.PrimaryKey

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *MiniVectorIndex) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &miniVectorIndexPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := miniVectorIndexFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MiniVectorIndex) MarshalJSON() ([]byte, error) {
	pb, err := miniVectorIndexToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type miniVectorIndexPb struct {
	// The user who created the index.
	Creator string `json:"creator,omitempty"`
	// Name of the endpoint associated with the index
	EndpointName string `json:"endpoint_name,omitempty"`
	// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
	// automatically syncs with a source Delta Table, automatically and
	// incrementally updating the index as the underlying data in the Delta
	// Table changes. - `DIRECT_ACCESS`: An index that supports direct read and
	// write of vectors and metadata through our REST and SDK APIs. With this
	// model, the user manages index updates.
	IndexType VectorIndexType `json:"index_type,omitempty"`
	// Name of the index
	Name string `json:"name,omitempty"`
	// Primary key of the index
	PrimaryKey string `json:"primary_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func miniVectorIndexFromPb(pb *miniVectorIndexPb) (*MiniVectorIndex, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MiniVectorIndex{}
	st.Creator = pb.Creator
	st.EndpointName = pb.EndpointName
	st.IndexType = pb.IndexType
	st.Name = pb.Name
	st.PrimaryKey = pb.PrimaryKey

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *miniVectorIndexPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st miniVectorIndexPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PatchEndpointBudgetPolicyRequest struct {
	// The budget policy id to be applied
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string
	// Name of the vector search endpoint
	// Wire name: 'endpoint_name'
	EndpointName string `tf:"-"`
}

func patchEndpointBudgetPolicyRequestToPb(st *PatchEndpointBudgetPolicyRequest) (*patchEndpointBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &patchEndpointBudgetPolicyRequestPb{}
	pb.BudgetPolicyId = st.BudgetPolicyId

	pb.EndpointName = st.EndpointName

	return pb, nil
}

func (st *PatchEndpointBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &patchEndpointBudgetPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := patchEndpointBudgetPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PatchEndpointBudgetPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := patchEndpointBudgetPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type patchEndpointBudgetPolicyRequestPb struct {
	// The budget policy id to be applied
	BudgetPolicyId string `json:"budget_policy_id"`
	// Name of the vector search endpoint
	EndpointName string `json:"-" url:"-"`
}

func patchEndpointBudgetPolicyRequestFromPb(pb *patchEndpointBudgetPolicyRequestPb) (*PatchEndpointBudgetPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PatchEndpointBudgetPolicyRequest{}
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.EndpointName = pb.EndpointName

	return st, nil
}

type PatchEndpointBudgetPolicyResponse struct {
	// The budget policy applied to the vector search endpoint.
	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string

	ForceSendFields []string `tf:"-"`
}

func patchEndpointBudgetPolicyResponseToPb(st *PatchEndpointBudgetPolicyResponse) (*patchEndpointBudgetPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &patchEndpointBudgetPolicyResponsePb{}
	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *PatchEndpointBudgetPolicyResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &patchEndpointBudgetPolicyResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := patchEndpointBudgetPolicyResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PatchEndpointBudgetPolicyResponse) MarshalJSON() ([]byte, error) {
	pb, err := patchEndpointBudgetPolicyResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type patchEndpointBudgetPolicyResponsePb struct {
	// The budget policy applied to the vector search endpoint.
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func patchEndpointBudgetPolicyResponseFromPb(pb *patchEndpointBudgetPolicyResponsePb) (*PatchEndpointBudgetPolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PatchEndpointBudgetPolicyResponse{}
	st.EffectiveBudgetPolicyId = pb.EffectiveBudgetPolicyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *patchEndpointBudgetPolicyResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st patchEndpointBudgetPolicyResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the triggered
// execution mode, the system stops processing after successfully refreshing the
// source table in the pipeline once, ensuring the table is updated based on the
// data available when the update started. - `CONTINUOUS`: If the pipeline uses
// continuous execution, the pipeline processes new data as it arrives in the
// source table to keep vector index fresh.
type PipelineType string
type pipelineTypePb string

// If the pipeline uses continuous execution, the pipeline processes new data as
// it arrives in the source table to keep vector index fresh.
const PipelineTypeContinuous PipelineType = `CONTINUOUS`

// If the pipeline uses the triggered execution mode, the system stops
// processing after successfully refreshing the source table in the pipeline
// once, ensuring the table is updated based on the data available when the
// update started.
const PipelineTypeTriggered PipelineType = `TRIGGERED`

// String representation for [fmt.Print]
func (f *PipelineType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PipelineType) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `TRIGGERED`:
		*f = PipelineType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "TRIGGERED"`, v)
	}
}

// Type always returns PipelineType to satisfy [pflag.Value] interface
func (f *PipelineType) Type() string {
	return "PipelineType"
}

func pipelineTypeToPb(st *PipelineType) (*pipelineTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelineTypePb(*st)
	return &pb, nil
}

func pipelineTypeFromPb(pb *pipelineTypePb) (*PipelineType, error) {
	if pb == nil {
		return nil, nil
	}
	st := PipelineType(*pb)
	return &st, nil
}

// Request payload for getting next page of results.
type QueryVectorIndexNextPageRequest struct {
	// Name of the endpoint.
	// Wire name: 'endpoint_name'
	EndpointName string
	// Name of the vector index to query.
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
	// Page token returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` API.
	// Wire name: 'page_token'
	PageToken string

	ForceSendFields []string `tf:"-"`
}

func queryVectorIndexNextPageRequestToPb(st *QueryVectorIndexNextPageRequest) (*queryVectorIndexNextPageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryVectorIndexNextPageRequestPb{}
	pb.EndpointName = st.EndpointName

	pb.IndexName = st.IndexName

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *QueryVectorIndexNextPageRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryVectorIndexNextPageRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryVectorIndexNextPageRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryVectorIndexNextPageRequest) MarshalJSON() ([]byte, error) {
	pb, err := queryVectorIndexNextPageRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type queryVectorIndexNextPageRequestPb struct {
	// Name of the endpoint.
	EndpointName string `json:"endpoint_name,omitempty"`
	// Name of the vector index to query.
	IndexName string `json:"-" url:"-"`
	// Page token returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` API.
	PageToken string `json:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryVectorIndexNextPageRequestFromPb(pb *queryVectorIndexNextPageRequestPb) (*QueryVectorIndexNextPageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryVectorIndexNextPageRequest{}
	st.EndpointName = pb.EndpointName
	st.IndexName = pb.IndexName
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryVectorIndexNextPageRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryVectorIndexNextPageRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryVectorIndexRequest struct {
	// List of column names to include in the response.
	// Wire name: 'columns'
	Columns []string
	// Column names used to retrieve data to send to the reranker.
	// Wire name: 'columns_to_rerank'
	ColumnsToRerank []string
	// JSON string representing query filters.
	//
	// Example filters:
	//
	// - `{"id <": 5}`: Filter for id less than 5. - `{"id >": 5}`: Filter for
	// id greater than 5. - `{"id <=": 5}`: Filter for id less than equal to 5.
	// - `{"id >=": 5}`: Filter for id greater than equal to 5. - `{"id": 5}`:
	// Filter for id equal to 5.
	// Wire name: 'filters_json'
	FiltersJson string
	// Name of the vector index to query.
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
	// Number of results to return. Defaults to 10.
	// Wire name: 'num_results'
	NumResults int
	// Query text. Required for Delta Sync Index using model endpoint.
	// Wire name: 'query_text'
	QueryText string
	// The query type to use. Choices are `ANN` and `HYBRID`. Defaults to `ANN`.
	// Wire name: 'query_type'
	QueryType string
	// Query vector. Required for Direct Vector Access Index and Delta Sync
	// Index using self-managed vectors.
	// Wire name: 'query_vector'
	QueryVector []float64
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	// Wire name: 'score_threshold'
	ScoreThreshold float64

	ForceSendFields []string `tf:"-"`
}

func queryVectorIndexRequestToPb(st *QueryVectorIndexRequest) (*queryVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryVectorIndexRequestPb{}
	pb.Columns = st.Columns

	pb.ColumnsToRerank = st.ColumnsToRerank

	pb.FiltersJson = st.FiltersJson

	pb.IndexName = st.IndexName

	pb.NumResults = st.NumResults

	pb.QueryText = st.QueryText

	pb.QueryType = st.QueryType

	pb.QueryVector = st.QueryVector

	pb.ScoreThreshold = st.ScoreThreshold

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *QueryVectorIndexRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryVectorIndexRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryVectorIndexRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryVectorIndexRequest) MarshalJSON() ([]byte, error) {
	pb, err := queryVectorIndexRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type queryVectorIndexRequestPb struct {
	// List of column names to include in the response.
	Columns []string `json:"columns"`
	// Column names used to retrieve data to send to the reranker.
	ColumnsToRerank []string `json:"columns_to_rerank,omitempty"`
	// JSON string representing query filters.
	//
	// Example filters:
	//
	// - `{"id <": 5}`: Filter for id less than 5. - `{"id >": 5}`: Filter for
	// id greater than 5. - `{"id <=": 5}`: Filter for id less than equal to 5.
	// - `{"id >=": 5}`: Filter for id greater than equal to 5. - `{"id": 5}`:
	// Filter for id equal to 5.
	FiltersJson string `json:"filters_json,omitempty"`
	// Name of the vector index to query.
	IndexName string `json:"-" url:"-"`
	// Number of results to return. Defaults to 10.
	NumResults int `json:"num_results,omitempty"`
	// Query text. Required for Delta Sync Index using model endpoint.
	QueryText string `json:"query_text,omitempty"`
	// The query type to use. Choices are `ANN` and `HYBRID`. Defaults to `ANN`.
	QueryType string `json:"query_type,omitempty"`
	// Query vector. Required for Direct Vector Access Index and Delta Sync
	// Index using self-managed vectors.
	QueryVector []float64 `json:"query_vector,omitempty"`
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	ScoreThreshold float64 `json:"score_threshold,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryVectorIndexRequestFromPb(pb *queryVectorIndexRequestPb) (*QueryVectorIndexRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryVectorIndexRequest{}
	st.Columns = pb.Columns
	st.ColumnsToRerank = pb.ColumnsToRerank
	st.FiltersJson = pb.FiltersJson
	st.IndexName = pb.IndexName
	st.NumResults = pb.NumResults
	st.QueryText = pb.QueryText
	st.QueryType = pb.QueryType
	st.QueryVector = pb.QueryVector
	st.ScoreThreshold = pb.ScoreThreshold

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryVectorIndexRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryVectorIndexRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryVectorIndexResponse struct {
	// Metadata about the result set.
	// Wire name: 'manifest'
	Manifest *ResultManifest
	// [Optional] Token that can be used in `QueryVectorIndexNextPage` API to
	// get next page of results. If more than 1000 results satisfy the query,
	// they are returned in groups of 1000. Empty value means no more results.
	// The maximum number of results that can be returned is 10,000.
	// Wire name: 'next_page_token'
	NextPageToken string
	// Data returned in the query result.
	// Wire name: 'result'
	Result *ResultData

	ForceSendFields []string `tf:"-"`
}

func queryVectorIndexResponseToPb(st *QueryVectorIndexResponse) (*queryVectorIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryVectorIndexResponsePb{}
	manifestPb, err := resultManifestToPb(st.Manifest)
	if err != nil {
		return nil, err
	}
	if manifestPb != nil {
		pb.Manifest = manifestPb
	}

	pb.NextPageToken = st.NextPageToken

	resultPb, err := resultDataToPb(st.Result)
	if err != nil {
		return nil, err
	}
	if resultPb != nil {
		pb.Result = resultPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *QueryVectorIndexResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryVectorIndexResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryVectorIndexResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryVectorIndexResponse) MarshalJSON() ([]byte, error) {
	pb, err := queryVectorIndexResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type queryVectorIndexResponsePb struct {
	// Metadata about the result set.
	Manifest *resultManifestPb `json:"manifest,omitempty"`
	// [Optional] Token that can be used in `QueryVectorIndexNextPage` API to
	// get next page of results. If more than 1000 results satisfy the query,
	// they are returned in groups of 1000. Empty value means no more results.
	// The maximum number of results that can be returned is 10,000.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Data returned in the query result.
	Result *resultDataPb `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryVectorIndexResponseFromPb(pb *queryVectorIndexResponsePb) (*QueryVectorIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryVectorIndexResponse{}
	manifestField, err := resultManifestFromPb(pb.Manifest)
	if err != nil {
		return nil, err
	}
	if manifestField != nil {
		st.Manifest = manifestField
	}
	st.NextPageToken = pb.NextPageToken
	resultField, err := resultDataFromPb(pb.Result)
	if err != nil {
		return nil, err
	}
	if resultField != nil {
		st.Result = resultField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryVectorIndexResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryVectorIndexResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Data returned in the query result.
type ResultData struct {
	// Data rows returned in the query.
	// Wire name: 'data_array'
	DataArray []ListValue
	// Number of rows in the result set.
	// Wire name: 'row_count'
	RowCount int

	ForceSendFields []string `tf:"-"`
}

func resultDataToPb(st *ResultData) (*resultDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultDataPb{}

	var dataArrayPb []listValuePb
	for _, item := range st.DataArray {
		itemPb, err := listValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dataArrayPb = append(dataArrayPb, *itemPb)
		}
	}
	pb.DataArray = dataArrayPb

	pb.RowCount = st.RowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ResultData) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resultDataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resultDataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResultData) MarshalJSON() ([]byte, error) {
	pb, err := resultDataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type resultDataPb struct {
	// Data rows returned in the query.
	DataArray []listValuePb `json:"data_array,omitempty"`
	// Number of rows in the result set.
	RowCount int `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resultDataFromPb(pb *resultDataPb) (*ResultData, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultData{}

	var dataArrayField []ListValue
	for _, itemPb := range pb.DataArray {
		item, err := listValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dataArrayField = append(dataArrayField, *item)
		}
	}
	st.DataArray = dataArrayField
	st.RowCount = pb.RowCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resultDataPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resultDataPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Metadata about the result set.
type ResultManifest struct {
	// Number of columns in the result set.
	// Wire name: 'column_count'
	ColumnCount int
	// Information about each column in the result set.
	// Wire name: 'columns'
	Columns []ColumnInfo

	ForceSendFields []string `tf:"-"`
}

func resultManifestToPb(st *ResultManifest) (*resultManifestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultManifestPb{}
	pb.ColumnCount = st.ColumnCount

	var columnsPb []columnInfoPb
	for _, item := range st.Columns {
		itemPb, err := columnInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsPb = append(columnsPb, *itemPb)
		}
	}
	pb.Columns = columnsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ResultManifest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resultManifestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resultManifestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResultManifest) MarshalJSON() ([]byte, error) {
	pb, err := resultManifestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type resultManifestPb struct {
	// Number of columns in the result set.
	ColumnCount int `json:"column_count,omitempty"`
	// Information about each column in the result set.
	Columns []columnInfoPb `json:"columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resultManifestFromPb(pb *resultManifestPb) (*ResultManifest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultManifest{}
	st.ColumnCount = pb.ColumnCount

	var columnsField []ColumnInfo
	for _, itemPb := range pb.Columns {
		item, err := columnInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsField = append(columnsField, *item)
		}
	}
	st.Columns = columnsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resultManifestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resultManifestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ScanVectorIndexRequest struct {
	// Name of the vector index to scan.
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
	// Primary key of the last entry returned in the previous scan.
	// Wire name: 'last_primary_key'
	LastPrimaryKey string
	// Number of results to return. Defaults to 10.
	// Wire name: 'num_results'
	NumResults int

	ForceSendFields []string `tf:"-"`
}

func scanVectorIndexRequestToPb(st *ScanVectorIndexRequest) (*scanVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &scanVectorIndexRequestPb{}
	pb.IndexName = st.IndexName

	pb.LastPrimaryKey = st.LastPrimaryKey

	pb.NumResults = st.NumResults

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ScanVectorIndexRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &scanVectorIndexRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := scanVectorIndexRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ScanVectorIndexRequest) MarshalJSON() ([]byte, error) {
	pb, err := scanVectorIndexRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type scanVectorIndexRequestPb struct {
	// Name of the vector index to scan.
	IndexName string `json:"-" url:"-"`
	// Primary key of the last entry returned in the previous scan.
	LastPrimaryKey string `json:"last_primary_key,omitempty"`
	// Number of results to return. Defaults to 10.
	NumResults int `json:"num_results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func scanVectorIndexRequestFromPb(pb *scanVectorIndexRequestPb) (*ScanVectorIndexRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ScanVectorIndexRequest{}
	st.IndexName = pb.IndexName
	st.LastPrimaryKey = pb.LastPrimaryKey
	st.NumResults = pb.NumResults

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *scanVectorIndexRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st scanVectorIndexRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Response to a scan vector index request.
type ScanVectorIndexResponse struct {
	// List of data entries
	// Wire name: 'data'
	Data []Struct
	// Primary key of the last entry.
	// Wire name: 'last_primary_key'
	LastPrimaryKey string

	ForceSendFields []string `tf:"-"`
}

func scanVectorIndexResponseToPb(st *ScanVectorIndexResponse) (*scanVectorIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &scanVectorIndexResponsePb{}

	var dataPb []structPb
	for _, item := range st.Data {
		itemPb, err := structToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dataPb = append(dataPb, *itemPb)
		}
	}
	pb.Data = dataPb

	pb.LastPrimaryKey = st.LastPrimaryKey

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ScanVectorIndexResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &scanVectorIndexResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := scanVectorIndexResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ScanVectorIndexResponse) MarshalJSON() ([]byte, error) {
	pb, err := scanVectorIndexResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type scanVectorIndexResponsePb struct {
	// List of data entries
	Data []structPb `json:"data,omitempty"`
	// Primary key of the last entry.
	LastPrimaryKey string `json:"last_primary_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func scanVectorIndexResponseFromPb(pb *scanVectorIndexResponsePb) (*ScanVectorIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ScanVectorIndexResponse{}

	var dataField []Struct
	for _, itemPb := range pb.Data {
		item, err := structFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dataField = append(dataField, *item)
		}
	}
	st.Data = dataField
	st.LastPrimaryKey = pb.LastPrimaryKey

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *scanVectorIndexResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st scanVectorIndexResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// copied from proto3 / Google Well Known Types, source:
// https://github.com/protocolbuffers/protobuf/blob/450d24ca820750c5db5112a6f0b0c2efb9758021/src/google/protobuf/struct.proto
// `Struct` represents a structured data value, consisting of fields which map
// to dynamically typed values. In some languages, `Struct` might be supported
// by a native representation. For example, in scripting languages like JS a
// struct is represented as an object. The details of that representation are
// described together with the proto support for the language.
//
// The JSON representation for `Struct` is JSON object.
type Struct struct {
	// Data entry, corresponding to a row in a vector index.
	// Wire name: 'fields'
	Fields []MapStringValueEntry
}

func structToPb(st *Struct) (*structPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &structPb{}

	var fieldsPb []mapStringValueEntryPb
	for _, item := range st.Fields {
		itemPb, err := mapStringValueEntryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			fieldsPb = append(fieldsPb, *itemPb)
		}
	}
	pb.Fields = fieldsPb

	return pb, nil
}

func (st *Struct) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &structPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := structFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Struct) MarshalJSON() ([]byte, error) {
	pb, err := structToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type structPb struct {
	// Data entry, corresponding to a row in a vector index.
	Fields []mapStringValueEntryPb `json:"fields,omitempty"`
}

func structFromPb(pb *structPb) (*Struct, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Struct{}

	var fieldsField []MapStringValueEntry
	for _, itemPb := range pb.Fields {
		item, err := mapStringValueEntryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			fieldsField = append(fieldsField, *item)
		}
	}
	st.Fields = fieldsField

	return st, nil
}

// Synchronize an index
type SyncIndexRequest struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
}

func syncIndexRequestToPb(st *SyncIndexRequest) (*syncIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncIndexRequestPb{}
	pb.IndexName = st.IndexName

	return pb, nil
}

func (st *SyncIndexRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &syncIndexRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := syncIndexRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SyncIndexRequest) MarshalJSON() ([]byte, error) {
	pb, err := syncIndexRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type syncIndexRequestPb struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	IndexName string `json:"-" url:"-"`
}

func syncIndexRequestFromPb(pb *syncIndexRequestPb) (*SyncIndexRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncIndexRequest{}
	st.IndexName = pb.IndexName

	return st, nil
}

type SyncIndexResponse struct {
}

func syncIndexResponseToPb(st *SyncIndexResponse) (*syncIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncIndexResponsePb{}

	return pb, nil
}

func (st *SyncIndexResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &syncIndexResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := syncIndexResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SyncIndexResponse) MarshalJSON() ([]byte, error) {
	pb, err := syncIndexResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type syncIndexResponsePb struct {
}

func syncIndexResponseFromPb(pb *syncIndexResponsePb) (*SyncIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncIndexResponse{}

	return st, nil
}

type UpdateEndpointCustomTagsRequest struct {
	// The new custom tags for the vector search endpoint
	// Wire name: 'custom_tags'
	CustomTags []CustomTag
	// Name of the vector search endpoint
	// Wire name: 'endpoint_name'
	EndpointName string `tf:"-"`
}

func updateEndpointCustomTagsRequestToPb(st *UpdateEndpointCustomTagsRequest) (*updateEndpointCustomTagsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEndpointCustomTagsRequestPb{}

	var customTagsPb []customTagPb
	for _, item := range st.CustomTags {
		itemPb, err := customTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb = append(customTagsPb, *itemPb)
		}
	}
	pb.CustomTags = customTagsPb

	pb.EndpointName = st.EndpointName

	return pb, nil
}

func (st *UpdateEndpointCustomTagsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateEndpointCustomTagsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateEndpointCustomTagsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateEndpointCustomTagsRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateEndpointCustomTagsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateEndpointCustomTagsRequestPb struct {
	// The new custom tags for the vector search endpoint
	CustomTags []customTagPb `json:"custom_tags"`
	// Name of the vector search endpoint
	EndpointName string `json:"-" url:"-"`
}

func updateEndpointCustomTagsRequestFromPb(pb *updateEndpointCustomTagsRequestPb) (*UpdateEndpointCustomTagsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEndpointCustomTagsRequest{}

	var customTagsField []CustomTag
	for _, itemPb := range pb.CustomTags {
		item, err := customTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField = append(customTagsField, *item)
		}
	}
	st.CustomTags = customTagsField
	st.EndpointName = pb.EndpointName

	return st, nil
}

type UpdateEndpointCustomTagsResponse struct {
	// All the custom tags that are applied to the vector search endpoint.
	// Wire name: 'custom_tags'
	CustomTags []CustomTag
	// The name of the vector search endpoint whose custom tags were updated.
	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
}

func updateEndpointCustomTagsResponseToPb(st *UpdateEndpointCustomTagsResponse) (*updateEndpointCustomTagsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEndpointCustomTagsResponsePb{}

	var customTagsPb []customTagPb
	for _, item := range st.CustomTags {
		itemPb, err := customTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb = append(customTagsPb, *itemPb)
		}
	}
	pb.CustomTags = customTagsPb

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateEndpointCustomTagsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateEndpointCustomTagsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateEndpointCustomTagsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateEndpointCustomTagsResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateEndpointCustomTagsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateEndpointCustomTagsResponsePb struct {
	// All the custom tags that are applied to the vector search endpoint.
	CustomTags []customTagPb `json:"custom_tags,omitempty"`
	// The name of the vector search endpoint whose custom tags were updated.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateEndpointCustomTagsResponseFromPb(pb *updateEndpointCustomTagsResponsePb) (*UpdateEndpointCustomTagsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEndpointCustomTagsResponse{}

	var customTagsField []CustomTag
	for _, itemPb := range pb.CustomTags {
		item, err := customTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField = append(customTagsField, *item)
		}
	}
	st.CustomTags = customTagsField
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateEndpointCustomTagsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateEndpointCustomTagsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpsertDataResult struct {
	// List of primary keys for rows that failed to process.
	// Wire name: 'failed_primary_keys'
	FailedPrimaryKeys []string
	// Count of successfully processed rows.
	// Wire name: 'success_row_count'
	SuccessRowCount int64

	ForceSendFields []string `tf:"-"`
}

func upsertDataResultToPb(st *UpsertDataResult) (*upsertDataResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &upsertDataResultPb{}
	pb.FailedPrimaryKeys = st.FailedPrimaryKeys

	pb.SuccessRowCount = st.SuccessRowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpsertDataResult) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &upsertDataResultPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := upsertDataResultFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpsertDataResult) MarshalJSON() ([]byte, error) {
	pb, err := upsertDataResultToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type upsertDataResultPb struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	// Count of successfully processed rows.
	SuccessRowCount int64 `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func upsertDataResultFromPb(pb *upsertDataResultPb) (*UpsertDataResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpsertDataResult{}
	st.FailedPrimaryKeys = pb.FailedPrimaryKeys
	st.SuccessRowCount = pb.SuccessRowCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *upsertDataResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st upsertDataResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpsertDataStatus string
type upsertDataStatusPb string

const UpsertDataStatusFailure UpsertDataStatus = `FAILURE`

const UpsertDataStatusPartialSuccess UpsertDataStatus = `PARTIAL_SUCCESS`

const UpsertDataStatusSuccess UpsertDataStatus = `SUCCESS`

// String representation for [fmt.Print]
func (f *UpsertDataStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpsertDataStatus) Set(v string) error {
	switch v {
	case `FAILURE`, `PARTIAL_SUCCESS`, `SUCCESS`:
		*f = UpsertDataStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILURE", "PARTIAL_SUCCESS", "SUCCESS"`, v)
	}
}

// Type always returns UpsertDataStatus to satisfy [pflag.Value] interface
func (f *UpsertDataStatus) Type() string {
	return "UpsertDataStatus"
}

func upsertDataStatusToPb(st *UpsertDataStatus) (*upsertDataStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := upsertDataStatusPb(*st)
	return &pb, nil
}

func upsertDataStatusFromPb(pb *upsertDataStatusPb) (*UpsertDataStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := UpsertDataStatus(*pb)
	return &st, nil
}

type UpsertDataVectorIndexRequest struct {
	// Name of the vector index where data is to be upserted. Must be a Direct
	// Vector Access Index.
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
	// JSON string representing the data to be upserted.
	// Wire name: 'inputs_json'
	InputsJson string
}

func upsertDataVectorIndexRequestToPb(st *UpsertDataVectorIndexRequest) (*upsertDataVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &upsertDataVectorIndexRequestPb{}
	pb.IndexName = st.IndexName

	pb.InputsJson = st.InputsJson

	return pb, nil
}

func (st *UpsertDataVectorIndexRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &upsertDataVectorIndexRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := upsertDataVectorIndexRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpsertDataVectorIndexRequest) MarshalJSON() ([]byte, error) {
	pb, err := upsertDataVectorIndexRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type upsertDataVectorIndexRequestPb struct {
	// Name of the vector index where data is to be upserted. Must be a Direct
	// Vector Access Index.
	IndexName string `json:"-" url:"-"`
	// JSON string representing the data to be upserted.
	InputsJson string `json:"inputs_json"`
}

func upsertDataVectorIndexRequestFromPb(pb *upsertDataVectorIndexRequestPb) (*UpsertDataVectorIndexRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpsertDataVectorIndexRequest{}
	st.IndexName = pb.IndexName
	st.InputsJson = pb.InputsJson

	return st, nil
}

type UpsertDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	// Wire name: 'result'
	Result *UpsertDataResult
	// Status of the upsert operation.
	// Wire name: 'status'
	Status UpsertDataStatus
}

func upsertDataVectorIndexResponseToPb(st *UpsertDataVectorIndexResponse) (*upsertDataVectorIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &upsertDataVectorIndexResponsePb{}
	resultPb, err := upsertDataResultToPb(st.Result)
	if err != nil {
		return nil, err
	}
	if resultPb != nil {
		pb.Result = resultPb
	}

	pb.Status = st.Status

	return pb, nil
}

func (st *UpsertDataVectorIndexResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &upsertDataVectorIndexResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := upsertDataVectorIndexResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpsertDataVectorIndexResponse) MarshalJSON() ([]byte, error) {
	pb, err := upsertDataVectorIndexResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type upsertDataVectorIndexResponsePb struct {
	// Result of the upsert or delete operation.
	Result *upsertDataResultPb `json:"result,omitempty"`
	// Status of the upsert operation.
	Status UpsertDataStatus `json:"status,omitempty"`
}

func upsertDataVectorIndexResponseFromPb(pb *upsertDataVectorIndexResponsePb) (*UpsertDataVectorIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpsertDataVectorIndexResponse{}
	resultField, err := upsertDataResultFromPb(pb.Result)
	if err != nil {
		return nil, err
	}
	if resultField != nil {
		st.Result = resultField
	}
	st.Status = pb.Status

	return st, nil
}

type Value struct {

	// Wire name: 'bool_value'
	BoolValue bool
	// copied from proto3 / Google Well Known Types, source:
	// https://github.com/protocolbuffers/protobuf/blob/450d24ca820750c5db5112a6f0b0c2efb9758021/src/google/protobuf/struct.proto
	// `ListValue` is a wrapper around a repeated field of values.
	//
	// The JSON representation for `ListValue` is JSON array.
	// Wire name: 'list_value'
	ListValue *ListValue

	// Wire name: 'number_value'
	NumberValue float64

	// Wire name: 'string_value'
	StringValue string
	// copied from proto3 / Google Well Known Types, source:
	// https://github.com/protocolbuffers/protobuf/blob/450d24ca820750c5db5112a6f0b0c2efb9758021/src/google/protobuf/struct.proto
	// `Struct` represents a structured data value, consisting of fields which
	// map to dynamically typed values. In some languages, `Struct` might be
	// supported by a native representation. For example, in scripting languages
	// like JS a struct is represented as an object. The details of that
	// representation are described together with the proto support for the
	// language.
	//
	// The JSON representation for `Struct` is JSON object.
	// Wire name: 'struct_value'
	StructValue *Struct

	ForceSendFields []string `tf:"-"`
}

func valueToPb(st *Value) (*valuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &valuePb{}
	pb.BoolValue = st.BoolValue

	listValuePb, err := listValueToPb(st.ListValue)
	if err != nil {
		return nil, err
	}
	if listValuePb != nil {
		pb.ListValue = listValuePb
	}

	pb.NumberValue = st.NumberValue

	pb.StringValue = st.StringValue

	structValuePb, err := structToPb(st.StructValue)
	if err != nil {
		return nil, err
	}
	if structValuePb != nil {
		pb.StructValue = structValuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Value) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &valuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := valueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Value) MarshalJSON() ([]byte, error) {
	pb, err := valueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type valuePb struct {
	BoolValue bool `json:"bool_value,omitempty"`
	// copied from proto3 / Google Well Known Types, source:
	// https://github.com/protocolbuffers/protobuf/blob/450d24ca820750c5db5112a6f0b0c2efb9758021/src/google/protobuf/struct.proto
	// `ListValue` is a wrapper around a repeated field of values.
	//
	// The JSON representation for `ListValue` is JSON array.
	ListValue *listValuePb `json:"list_value,omitempty"`

	NumberValue float64 `json:"number_value,omitempty"`

	StringValue string `json:"string_value,omitempty"`
	// copied from proto3 / Google Well Known Types, source:
	// https://github.com/protocolbuffers/protobuf/blob/450d24ca820750c5db5112a6f0b0c2efb9758021/src/google/protobuf/struct.proto
	// `Struct` represents a structured data value, consisting of fields which
	// map to dynamically typed values. In some languages, `Struct` might be
	// supported by a native representation. For example, in scripting languages
	// like JS a struct is represented as an object. The details of that
	// representation are described together with the proto support for the
	// language.
	//
	// The JSON representation for `Struct` is JSON object.
	StructValue *structPb `json:"struct_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func valueFromPb(pb *valuePb) (*Value, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Value{}
	st.BoolValue = pb.BoolValue
	listValueField, err := listValueFromPb(pb.ListValue)
	if err != nil {
		return nil, err
	}
	if listValueField != nil {
		st.ListValue = listValueField
	}
	st.NumberValue = pb.NumberValue
	st.StringValue = pb.StringValue
	structValueField, err := structFromPb(pb.StructValue)
	if err != nil {
		return nil, err
	}
	if structValueField != nil {
		st.StructValue = structValueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *valuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st valuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type VectorIndex struct {
	// The user who created the index.
	// Wire name: 'creator'
	Creator string

	// Wire name: 'delta_sync_index_spec'
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecResponse

	// Wire name: 'direct_access_index_spec'
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec
	// Name of the endpoint associated with the index
	// Wire name: 'endpoint_name'
	EndpointName string
	// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
	// automatically syncs with a source Delta Table, automatically and
	// incrementally updating the index as the underlying data in the Delta
	// Table changes. - `DIRECT_ACCESS`: An index that supports direct read and
	// write of vectors and metadata through our REST and SDK APIs. With this
	// model, the user manages index updates.
	// Wire name: 'index_type'
	IndexType VectorIndexType
	// Name of the index
	// Wire name: 'name'
	Name string
	// Primary key of the index
	// Wire name: 'primary_key'
	PrimaryKey string

	// Wire name: 'status'
	Status *VectorIndexStatus

	ForceSendFields []string `tf:"-"`
}

func vectorIndexToPb(st *VectorIndex) (*vectorIndexPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorIndexPb{}
	pb.Creator = st.Creator

	deltaSyncIndexSpecPb, err := deltaSyncVectorIndexSpecResponseToPb(st.DeltaSyncIndexSpec)
	if err != nil {
		return nil, err
	}
	if deltaSyncIndexSpecPb != nil {
		pb.DeltaSyncIndexSpec = deltaSyncIndexSpecPb
	}

	directAccessIndexSpecPb, err := directAccessVectorIndexSpecToPb(st.DirectAccessIndexSpec)
	if err != nil {
		return nil, err
	}
	if directAccessIndexSpecPb != nil {
		pb.DirectAccessIndexSpec = directAccessIndexSpecPb
	}

	pb.EndpointName = st.EndpointName

	pb.IndexType = st.IndexType

	pb.Name = st.Name

	pb.PrimaryKey = st.PrimaryKey

	statusPb, err := vectorIndexStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *VectorIndex) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &vectorIndexPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := vectorIndexFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st VectorIndex) MarshalJSON() ([]byte, error) {
	pb, err := vectorIndexToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type vectorIndexPb struct {
	// The user who created the index.
	Creator string `json:"creator,omitempty"`

	DeltaSyncIndexSpec *deltaSyncVectorIndexSpecResponsePb `json:"delta_sync_index_spec,omitempty"`

	DirectAccessIndexSpec *directAccessVectorIndexSpecPb `json:"direct_access_index_spec,omitempty"`
	// Name of the endpoint associated with the index
	EndpointName string `json:"endpoint_name,omitempty"`
	// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
	// automatically syncs with a source Delta Table, automatically and
	// incrementally updating the index as the underlying data in the Delta
	// Table changes. - `DIRECT_ACCESS`: An index that supports direct read and
	// write of vectors and metadata through our REST and SDK APIs. With this
	// model, the user manages index updates.
	IndexType VectorIndexType `json:"index_type,omitempty"`
	// Name of the index
	Name string `json:"name,omitempty"`
	// Primary key of the index
	PrimaryKey string `json:"primary_key,omitempty"`

	Status *vectorIndexStatusPb `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func vectorIndexFromPb(pb *vectorIndexPb) (*VectorIndex, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VectorIndex{}
	st.Creator = pb.Creator
	deltaSyncIndexSpecField, err := deltaSyncVectorIndexSpecResponseFromPb(pb.DeltaSyncIndexSpec)
	if err != nil {
		return nil, err
	}
	if deltaSyncIndexSpecField != nil {
		st.DeltaSyncIndexSpec = deltaSyncIndexSpecField
	}
	directAccessIndexSpecField, err := directAccessVectorIndexSpecFromPb(pb.DirectAccessIndexSpec)
	if err != nil {
		return nil, err
	}
	if directAccessIndexSpecField != nil {
		st.DirectAccessIndexSpec = directAccessIndexSpecField
	}
	st.EndpointName = pb.EndpointName
	st.IndexType = pb.IndexType
	st.Name = pb.Name
	st.PrimaryKey = pb.PrimaryKey
	statusField, err := vectorIndexStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *vectorIndexPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st vectorIndexPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type VectorIndexStatus struct {
	// Index API Url to be used to perform operations on the index
	// Wire name: 'index_url'
	IndexUrl string
	// Number of rows indexed
	// Wire name: 'indexed_row_count'
	IndexedRowCount int64
	// Message associated with the index status
	// Wire name: 'message'
	Message string
	// Whether the index is ready for search
	// Wire name: 'ready'
	Ready bool

	ForceSendFields []string `tf:"-"`
}

func vectorIndexStatusToPb(st *VectorIndexStatus) (*vectorIndexStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorIndexStatusPb{}
	pb.IndexUrl = st.IndexUrl

	pb.IndexedRowCount = st.IndexedRowCount

	pb.Message = st.Message

	pb.Ready = st.Ready

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *VectorIndexStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &vectorIndexStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := vectorIndexStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st VectorIndexStatus) MarshalJSON() ([]byte, error) {
	pb, err := vectorIndexStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type vectorIndexStatusPb struct {
	// Index API Url to be used to perform operations on the index
	IndexUrl string `json:"index_url,omitempty"`
	// Number of rows indexed
	IndexedRowCount int64 `json:"indexed_row_count,omitempty"`
	// Message associated with the index status
	Message string `json:"message,omitempty"`
	// Whether the index is ready for search
	Ready bool `json:"ready,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func vectorIndexStatusFromPb(pb *vectorIndexStatusPb) (*VectorIndexStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VectorIndexStatus{}
	st.IndexUrl = pb.IndexUrl
	st.IndexedRowCount = pb.IndexedRowCount
	st.Message = pb.Message
	st.Ready = pb.Ready

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *vectorIndexStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st vectorIndexStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
// automatically syncs with a source Delta Table, automatically and
// incrementally updating the index as the underlying data in the Delta Table
// changes. - `DIRECT_ACCESS`: An index that supports direct read and write of
// vectors and metadata through our REST and SDK APIs. With this model, the user
// manages index updates.
type VectorIndexType string
type vectorIndexTypePb string

// An index that automatically syncs with a source Delta Table, automatically
// and incrementally updating the index as the underlying data in the Delta
// Table changes.
const VectorIndexTypeDeltaSync VectorIndexType = `DELTA_SYNC`

// An index that supports direct read and write of vectors and metadata through
// our REST and SDK APIs. With this model, the user manages index updates.
const VectorIndexTypeDirectAccess VectorIndexType = `DIRECT_ACCESS`

// String representation for [fmt.Print]
func (f *VectorIndexType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VectorIndexType) Set(v string) error {
	switch v {
	case `DELTA_SYNC`, `DIRECT_ACCESS`:
		*f = VectorIndexType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTA_SYNC", "DIRECT_ACCESS"`, v)
	}
}

// Type always returns VectorIndexType to satisfy [pflag.Value] interface
func (f *VectorIndexType) Type() string {
	return "VectorIndexType"
}

func vectorIndexTypeToPb(st *VectorIndexType) (*vectorIndexTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := vectorIndexTypePb(*st)
	return &pb, nil
}

func vectorIndexTypeFromPb(pb *vectorIndexTypePb) (*VectorIndexType, error) {
	if pb == nil {
		return nil, nil
	}
	st := VectorIndexType(*pb)
	return &st, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

func durationFromPb(s *string) (*time.Duration, error) {
	if s == nil {
		return nil, nil
	}
	d, err := time.ParseDuration(*s)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func timestampToPb(t *time.Time) (*string, error) {
	if t == nil {
		return nil, nil
	}
	s := t.Format(time.RFC3339)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func fieldMaskToPb(fm *[]string) (*string, error) {
	if fm == nil {
		return nil, nil
	}
	s := strings.Join(*fm, ",")
	return &s, nil
}

func fieldMaskFromPb(s *string) (*[]string, error) {
	if s == nil {
		return nil, nil
	}
	fm := strings.Split(*s, ",")
	return &fm, nil
}

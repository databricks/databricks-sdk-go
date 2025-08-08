// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/vectorsearch/vectorsearchpb"
)

type ColumnInfo struct {
	// Name of the column.
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ColumnInfoToPb(st *ColumnInfo) (*vectorsearchpb.ColumnInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.ColumnInfoPb{}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ColumnInfoFromPb(pb *vectorsearchpb.ColumnInfoPb) (*ColumnInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ColumnInfo{}
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateEndpoint struct {
	// The budget policy id to be applied
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string ``
	// Type of endpoint
	// Wire name: 'endpoint_type'
	EndpointType EndpointType ``
	// Name of the vector search endpoint
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateEndpointToPb(st *CreateEndpoint) (*vectorsearchpb.CreateEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.CreateEndpointPb{}
	pb.BudgetPolicyId = st.BudgetPolicyId
	endpointTypePb, err := EndpointTypeToPb(&st.EndpointType)
	if err != nil {
		return nil, err
	}
	if endpointTypePb != nil {
		pb.EndpointType = *endpointTypePb
	}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateEndpointFromPb(pb *vectorsearchpb.CreateEndpointPb) (*CreateEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateEndpoint{}
	st.BudgetPolicyId = pb.BudgetPolicyId
	endpointTypeField, err := EndpointTypeFromPb(&pb.EndpointType)
	if err != nil {
		return nil, err
	}
	if endpointTypeField != nil {
		st.EndpointType = *endpointTypeField
	}
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateVectorIndexRequest struct {
	// Specification for Delta Sync Index. Required if `index_type` is
	// `DELTA_SYNC`.
	// Wire name: 'delta_sync_index_spec'
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecRequest ``
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	// Wire name: 'direct_access_index_spec'
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec ``
	// Name of the endpoint to be used for serving the index
	// Wire name: 'endpoint_name'
	EndpointName string ``

	// Wire name: 'index_type'
	IndexType VectorIndexType ``
	// Name of the index
	// Wire name: 'name'
	Name string ``
	// Primary key of the index
	// Wire name: 'primary_key'
	PrimaryKey string ``
}

func CreateVectorIndexRequestToPb(st *CreateVectorIndexRequest) (*vectorsearchpb.CreateVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.CreateVectorIndexRequestPb{}
	deltaSyncIndexSpecPb, err := DeltaSyncVectorIndexSpecRequestToPb(st.DeltaSyncIndexSpec)
	if err != nil {
		return nil, err
	}
	if deltaSyncIndexSpecPb != nil {
		pb.DeltaSyncIndexSpec = deltaSyncIndexSpecPb
	}
	directAccessIndexSpecPb, err := DirectAccessVectorIndexSpecToPb(st.DirectAccessIndexSpec)
	if err != nil {
		return nil, err
	}
	if directAccessIndexSpecPb != nil {
		pb.DirectAccessIndexSpec = directAccessIndexSpecPb
	}
	pb.EndpointName = st.EndpointName
	indexTypePb, err := VectorIndexTypeToPb(&st.IndexType)
	if err != nil {
		return nil, err
	}
	if indexTypePb != nil {
		pb.IndexType = *indexTypePb
	}
	pb.Name = st.Name
	pb.PrimaryKey = st.PrimaryKey

	return pb, nil
}

func CreateVectorIndexRequestFromPb(pb *vectorsearchpb.CreateVectorIndexRequestPb) (*CreateVectorIndexRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVectorIndexRequest{}
	deltaSyncIndexSpecField, err := DeltaSyncVectorIndexSpecRequestFromPb(pb.DeltaSyncIndexSpec)
	if err != nil {
		return nil, err
	}
	if deltaSyncIndexSpecField != nil {
		st.DeltaSyncIndexSpec = deltaSyncIndexSpecField
	}
	directAccessIndexSpecField, err := DirectAccessVectorIndexSpecFromPb(pb.DirectAccessIndexSpec)
	if err != nil {
		return nil, err
	}
	if directAccessIndexSpecField != nil {
		st.DirectAccessIndexSpec = directAccessIndexSpecField
	}
	st.EndpointName = pb.EndpointName
	indexTypeField, err := VectorIndexTypeFromPb(&pb.IndexType)
	if err != nil {
		return nil, err
	}
	if indexTypeField != nil {
		st.IndexType = *indexTypeField
	}
	st.Name = pb.Name
	st.PrimaryKey = pb.PrimaryKey

	return st, nil
}

type CustomTag struct {
	// Key field for a vector search endpoint tag.
	// Wire name: 'key'
	Key string ``
	// [Optional] Value field for a vector search endpoint tag.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CustomTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CustomTagToPb(st *CustomTag) (*vectorsearchpb.CustomTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.CustomTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CustomTagFromPb(pb *vectorsearchpb.CustomTagPb) (*CustomTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteDataResult struct {
	// List of primary keys for rows that failed to process.
	// Wire name: 'failed_primary_keys'
	FailedPrimaryKeys []string ``
	// Count of successfully processed rows.
	// Wire name: 'success_row_count'
	SuccessRowCount int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *DeleteDataResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDataResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteDataResultToPb(st *DeleteDataResult) (*vectorsearchpb.DeleteDataResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.DeleteDataResultPb{}
	pb.FailedPrimaryKeys = st.FailedPrimaryKeys
	pb.SuccessRowCount = st.SuccessRowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteDataResultFromPb(pb *vectorsearchpb.DeleteDataResultPb) (*DeleteDataResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDataResult{}
	st.FailedPrimaryKeys = pb.FailedPrimaryKeys
	st.SuccessRowCount = pb.SuccessRowCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteDataStatus string

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

// Values returns all possible values for DeleteDataStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeleteDataStatus) Values() []DeleteDataStatus {
	return []DeleteDataStatus{
		DeleteDataStatusFailure,
		DeleteDataStatusPartialSuccess,
		DeleteDataStatusSuccess,
	}
}

// Type always returns DeleteDataStatus to satisfy [pflag.Value] interface
func (f *DeleteDataStatus) Type() string {
	return "DeleteDataStatus"
}

func DeleteDataStatusToPb(st *DeleteDataStatus) (*vectorsearchpb.DeleteDataStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := vectorsearchpb.DeleteDataStatusPb(*st)
	return &pb, nil
}

func DeleteDataStatusFromPb(pb *vectorsearchpb.DeleteDataStatusPb) (*DeleteDataStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := DeleteDataStatus(*pb)
	return &st, nil
}

type DeleteDataVectorIndexRequest struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
	// List of primary keys for the data to be deleted.
	// Wire name: 'primary_keys'
	PrimaryKeys []string `tf:"-"`
}

func DeleteDataVectorIndexRequestToPb(st *DeleteDataVectorIndexRequest) (*vectorsearchpb.DeleteDataVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.DeleteDataVectorIndexRequestPb{}
	pb.IndexName = st.IndexName
	pb.PrimaryKeys = st.PrimaryKeys

	return pb, nil
}

func DeleteDataVectorIndexRequestFromPb(pb *vectorsearchpb.DeleteDataVectorIndexRequestPb) (*DeleteDataVectorIndexRequest, error) {
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
	Result *DeleteDataResult ``
	// Status of the delete operation.
	// Wire name: 'status'
	Status DeleteDataStatus ``
}

func DeleteDataVectorIndexResponseToPb(st *DeleteDataVectorIndexResponse) (*vectorsearchpb.DeleteDataVectorIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.DeleteDataVectorIndexResponsePb{}
	resultPb, err := DeleteDataResultToPb(st.Result)
	if err != nil {
		return nil, err
	}
	if resultPb != nil {
		pb.Result = resultPb
	}
	statusPb, err := DeleteDataStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	return pb, nil
}

func DeleteDataVectorIndexResponseFromPb(pb *vectorsearchpb.DeleteDataVectorIndexResponsePb) (*DeleteDataVectorIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDataVectorIndexResponse{}
	resultField, err := DeleteDataResultFromPb(pb.Result)
	if err != nil {
		return nil, err
	}
	if resultField != nil {
		st.Result = resultField
	}
	statusField, err := DeleteDataStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	return st, nil
}

type DeleteEndpointRequest struct {
	// Name of the vector search endpoint
	// Wire name: 'endpoint_name'
	EndpointName string `tf:"-"`
}

func DeleteEndpointRequestToPb(st *DeleteEndpointRequest) (*vectorsearchpb.DeleteEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.DeleteEndpointRequestPb{}
	pb.EndpointName = st.EndpointName

	return pb, nil
}

func DeleteEndpointRequestFromPb(pb *vectorsearchpb.DeleteEndpointRequestPb) (*DeleteEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteEndpointRequest{}
	st.EndpointName = pb.EndpointName

	return st, nil
}

type DeleteIndexRequest struct {
	// Name of the index
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
}

func DeleteIndexRequestToPb(st *DeleteIndexRequest) (*vectorsearchpb.DeleteIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.DeleteIndexRequestPb{}
	pb.IndexName = st.IndexName

	return pb, nil
}

func DeleteIndexRequestFromPb(pb *vectorsearchpb.DeleteIndexRequestPb) (*DeleteIndexRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteIndexRequest{}
	st.IndexName = pb.IndexName

	return st, nil
}

type DeltaSyncVectorIndexSpecRequest struct {
	// [Optional] Select the columns to sync with the vector index. If you leave
	// this field blank, all columns from the source table are synced with the
	// index. The primary key column and embedding source column or embedding
	// vector column are always synced.
	// Wire name: 'columns_to_sync'
	ColumnsToSync []string ``
	// The columns that contain the embedding source.
	// Wire name: 'embedding_source_columns'
	EmbeddingSourceColumns []EmbeddingSourceColumn ``
	// The columns that contain the embedding vectors.
	// Wire name: 'embedding_vector_columns'
	EmbeddingVectorColumns []EmbeddingVectorColumn ``
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	// Wire name: 'embedding_writeback_table'
	EmbeddingWritebackTable string ``
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	// Wire name: 'pipeline_type'
	PipelineType PipelineType ``
	// The name of the source table.
	// Wire name: 'source_table'
	SourceTable     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DeltaSyncVectorIndexSpecRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSyncVectorIndexSpecRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeltaSyncVectorIndexSpecRequestToPb(st *DeltaSyncVectorIndexSpecRequest) (*vectorsearchpb.DeltaSyncVectorIndexSpecRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.DeltaSyncVectorIndexSpecRequestPb{}
	pb.ColumnsToSync = st.ColumnsToSync

	var embeddingSourceColumnsPb []vectorsearchpb.EmbeddingSourceColumnPb
	for _, item := range st.EmbeddingSourceColumns {
		itemPb, err := EmbeddingSourceColumnToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			embeddingSourceColumnsPb = append(embeddingSourceColumnsPb, *itemPb)
		}
	}
	pb.EmbeddingSourceColumns = embeddingSourceColumnsPb

	var embeddingVectorColumnsPb []vectorsearchpb.EmbeddingVectorColumnPb
	for _, item := range st.EmbeddingVectorColumns {
		itemPb, err := EmbeddingVectorColumnToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			embeddingVectorColumnsPb = append(embeddingVectorColumnsPb, *itemPb)
		}
	}
	pb.EmbeddingVectorColumns = embeddingVectorColumnsPb
	pb.EmbeddingWritebackTable = st.EmbeddingWritebackTable
	pipelineTypePb, err := PipelineTypeToPb(&st.PipelineType)
	if err != nil {
		return nil, err
	}
	if pipelineTypePb != nil {
		pb.PipelineType = *pipelineTypePb
	}
	pb.SourceTable = st.SourceTable

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeltaSyncVectorIndexSpecRequestFromPb(pb *vectorsearchpb.DeltaSyncVectorIndexSpecRequestPb) (*DeltaSyncVectorIndexSpecRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSyncVectorIndexSpecRequest{}
	st.ColumnsToSync = pb.ColumnsToSync

	var embeddingSourceColumnsField []EmbeddingSourceColumn
	for _, itemPb := range pb.EmbeddingSourceColumns {
		item, err := EmbeddingSourceColumnFromPb(&itemPb)
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
		item, err := EmbeddingVectorColumnFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			embeddingVectorColumnsField = append(embeddingVectorColumnsField, *item)
		}
	}
	st.EmbeddingVectorColumns = embeddingVectorColumnsField
	st.EmbeddingWritebackTable = pb.EmbeddingWritebackTable
	pipelineTypeField, err := PipelineTypeFromPb(&pb.PipelineType)
	if err != nil {
		return nil, err
	}
	if pipelineTypeField != nil {
		st.PipelineType = *pipelineTypeField
	}
	st.SourceTable = pb.SourceTable

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeltaSyncVectorIndexSpecResponse struct {
	// The columns that contain the embedding source.
	// Wire name: 'embedding_source_columns'
	EmbeddingSourceColumns []EmbeddingSourceColumn ``
	// The columns that contain the embedding vectors.
	// Wire name: 'embedding_vector_columns'
	EmbeddingVectorColumns []EmbeddingVectorColumn ``
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	// Wire name: 'embedding_writeback_table'
	EmbeddingWritebackTable string ``
	// The ID of the pipeline that is used to sync the index.
	// Wire name: 'pipeline_id'
	PipelineId string ``
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	// Wire name: 'pipeline_type'
	PipelineType PipelineType ``
	// The name of the source table.
	// Wire name: 'source_table'
	SourceTable     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DeltaSyncVectorIndexSpecResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSyncVectorIndexSpecResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeltaSyncVectorIndexSpecResponseToPb(st *DeltaSyncVectorIndexSpecResponse) (*vectorsearchpb.DeltaSyncVectorIndexSpecResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.DeltaSyncVectorIndexSpecResponsePb{}

	var embeddingSourceColumnsPb []vectorsearchpb.EmbeddingSourceColumnPb
	for _, item := range st.EmbeddingSourceColumns {
		itemPb, err := EmbeddingSourceColumnToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			embeddingSourceColumnsPb = append(embeddingSourceColumnsPb, *itemPb)
		}
	}
	pb.EmbeddingSourceColumns = embeddingSourceColumnsPb

	var embeddingVectorColumnsPb []vectorsearchpb.EmbeddingVectorColumnPb
	for _, item := range st.EmbeddingVectorColumns {
		itemPb, err := EmbeddingVectorColumnToPb(&item)
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
	pipelineTypePb, err := PipelineTypeToPb(&st.PipelineType)
	if err != nil {
		return nil, err
	}
	if pipelineTypePb != nil {
		pb.PipelineType = *pipelineTypePb
	}
	pb.SourceTable = st.SourceTable

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeltaSyncVectorIndexSpecResponseFromPb(pb *vectorsearchpb.DeltaSyncVectorIndexSpecResponsePb) (*DeltaSyncVectorIndexSpecResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSyncVectorIndexSpecResponse{}

	var embeddingSourceColumnsField []EmbeddingSourceColumn
	for _, itemPb := range pb.EmbeddingSourceColumns {
		item, err := EmbeddingSourceColumnFromPb(&itemPb)
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
		item, err := EmbeddingVectorColumnFromPb(&itemPb)
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
	pipelineTypeField, err := PipelineTypeFromPb(&pb.PipelineType)
	if err != nil {
		return nil, err
	}
	if pipelineTypeField != nil {
		st.PipelineType = *pipelineTypeField
	}
	st.SourceTable = pb.SourceTable

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DirectAccessVectorIndexSpec struct {
	// The columns that contain the embedding source. The format should be
	// array[double].
	// Wire name: 'embedding_source_columns'
	EmbeddingSourceColumns []EmbeddingSourceColumn ``
	// The columns that contain the embedding vectors. The format should be
	// array[double].
	// Wire name: 'embedding_vector_columns'
	EmbeddingVectorColumns []EmbeddingVectorColumn ``
	// The schema of the index in JSON format. Supported types are `integer`,
	// `long`, `float`, `double`, `boolean`, `string`, `date`, `timestamp`.
	// Supported types for vector column: `array<float>`, `array<double>`,`.
	// Wire name: 'schema_json'
	SchemaJson      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DirectAccessVectorIndexSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DirectAccessVectorIndexSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DirectAccessVectorIndexSpecToPb(st *DirectAccessVectorIndexSpec) (*vectorsearchpb.DirectAccessVectorIndexSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.DirectAccessVectorIndexSpecPb{}

	var embeddingSourceColumnsPb []vectorsearchpb.EmbeddingSourceColumnPb
	for _, item := range st.EmbeddingSourceColumns {
		itemPb, err := EmbeddingSourceColumnToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			embeddingSourceColumnsPb = append(embeddingSourceColumnsPb, *itemPb)
		}
	}
	pb.EmbeddingSourceColumns = embeddingSourceColumnsPb

	var embeddingVectorColumnsPb []vectorsearchpb.EmbeddingVectorColumnPb
	for _, item := range st.EmbeddingVectorColumns {
		itemPb, err := EmbeddingVectorColumnToPb(&item)
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

func DirectAccessVectorIndexSpecFromPb(pb *vectorsearchpb.DirectAccessVectorIndexSpecPb) (*DirectAccessVectorIndexSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DirectAccessVectorIndexSpec{}

	var embeddingSourceColumnsField []EmbeddingSourceColumn
	for _, itemPb := range pb.EmbeddingSourceColumns {
		item, err := EmbeddingSourceColumnFromPb(&itemPb)
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
		item, err := EmbeddingVectorColumnFromPb(&itemPb)
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

type EmbeddingSourceColumn struct {
	// Name of the embedding model endpoint
	// Wire name: 'embedding_model_endpoint_name'
	EmbeddingModelEndpointName string ``
	// Name of the column
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *EmbeddingSourceColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingSourceColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EmbeddingSourceColumnToPb(st *EmbeddingSourceColumn) (*vectorsearchpb.EmbeddingSourceColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.EmbeddingSourceColumnPb{}
	pb.EmbeddingModelEndpointName = st.EmbeddingModelEndpointName
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EmbeddingSourceColumnFromPb(pb *vectorsearchpb.EmbeddingSourceColumnPb) (*EmbeddingSourceColumn, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EmbeddingSourceColumn{}
	st.EmbeddingModelEndpointName = pb.EmbeddingModelEndpointName
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type EmbeddingVectorColumn struct {
	// Dimension of the embedding vector
	// Wire name: 'embedding_dimension'
	EmbeddingDimension int ``
	// Name of the column
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *EmbeddingVectorColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingVectorColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EmbeddingVectorColumnToPb(st *EmbeddingVectorColumn) (*vectorsearchpb.EmbeddingVectorColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.EmbeddingVectorColumnPb{}
	pb.EmbeddingDimension = st.EmbeddingDimension
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EmbeddingVectorColumnFromPb(pb *vectorsearchpb.EmbeddingVectorColumnPb) (*EmbeddingVectorColumn, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EmbeddingVectorColumn{}
	st.EmbeddingDimension = pb.EmbeddingDimension
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type EndpointInfo struct {
	// Timestamp of endpoint creation
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``
	// Creator of the endpoint
	// Wire name: 'creator'
	Creator string ``
	// The custom tags assigned to the endpoint
	// Wire name: 'custom_tags'
	CustomTags []CustomTag ``
	// The budget policy id applied to the endpoint
	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string ``
	// Current status of the endpoint
	// Wire name: 'endpoint_status'
	EndpointStatus *EndpointStatus ``
	// Type of endpoint
	// Wire name: 'endpoint_type'
	EndpointType EndpointType ``
	// Unique identifier of the endpoint
	// Wire name: 'id'
	Id string ``
	// Timestamp of last update to the endpoint
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// User who last updated the endpoint
	// Wire name: 'last_updated_user'
	LastUpdatedUser string ``
	// Name of the vector search endpoint
	// Wire name: 'name'
	Name string ``
	// Number of indexes on the endpoint
	// Wire name: 'num_indexes'
	NumIndexes      int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *EndpointInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EndpointInfoToPb(st *EndpointInfo) (*vectorsearchpb.EndpointInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.EndpointInfoPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Creator = st.Creator

	var customTagsPb []vectorsearchpb.CustomTagPb
	for _, item := range st.CustomTags {
		itemPb, err := CustomTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb = append(customTagsPb, *itemPb)
		}
	}
	pb.CustomTags = customTagsPb
	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId
	endpointStatusPb, err := EndpointStatusToPb(st.EndpointStatus)
	if err != nil {
		return nil, err
	}
	if endpointStatusPb != nil {
		pb.EndpointStatus = endpointStatusPb
	}
	endpointTypePb, err := EndpointTypeToPb(&st.EndpointType)
	if err != nil {
		return nil, err
	}
	if endpointTypePb != nil {
		pb.EndpointType = *endpointTypePb
	}
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.LastUpdatedUser = st.LastUpdatedUser
	pb.Name = st.Name
	pb.NumIndexes = st.NumIndexes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EndpointInfoFromPb(pb *vectorsearchpb.EndpointInfoPb) (*EndpointInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointInfo{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Creator = pb.Creator

	var customTagsField []CustomTag
	for _, itemPb := range pb.CustomTags {
		item, err := CustomTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField = append(customTagsField, *item)
		}
	}
	st.CustomTags = customTagsField
	st.EffectiveBudgetPolicyId = pb.EffectiveBudgetPolicyId
	endpointStatusField, err := EndpointStatusFromPb(pb.EndpointStatus)
	if err != nil {
		return nil, err
	}
	if endpointStatusField != nil {
		st.EndpointStatus = endpointStatusField
	}
	endpointTypeField, err := EndpointTypeFromPb(&pb.EndpointType)
	if err != nil {
		return nil, err
	}
	if endpointTypeField != nil {
		st.EndpointType = *endpointTypeField
	}
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.LastUpdatedUser = pb.LastUpdatedUser
	st.Name = pb.Name
	st.NumIndexes = pb.NumIndexes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Status information of an endpoint
type EndpointStatus struct {
	// Additional status message
	// Wire name: 'message'
	Message string ``
	// Current state of the endpoint
	// Wire name: 'state'
	State           EndpointStatusState ``
	ForceSendFields []string            `tf:"-"`
}

func (s *EndpointStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EndpointStatusToPb(st *EndpointStatus) (*vectorsearchpb.EndpointStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.EndpointStatusPb{}
	pb.Message = st.Message
	statePb, err := EndpointStatusStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EndpointStatusFromPb(pb *vectorsearchpb.EndpointStatusPb) (*EndpointStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointStatus{}
	st.Message = pb.Message
	stateField, err := EndpointStatusStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Current state of the endpoint
type EndpointStatusState string

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

// Values returns all possible values for EndpointStatusState.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointStatusState) Values() []EndpointStatusState {
	return []EndpointStatusState{
		EndpointStatusStateOffline,
		EndpointStatusStateOnline,
		EndpointStatusStateProvisioning,
	}
}

// Type always returns EndpointStatusState to satisfy [pflag.Value] interface
func (f *EndpointStatusState) Type() string {
	return "EndpointStatusState"
}

func EndpointStatusStateToPb(st *EndpointStatusState) (*vectorsearchpb.EndpointStatusStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := vectorsearchpb.EndpointStatusStatePb(*st)
	return &pb, nil
}

func EndpointStatusStateFromPb(pb *vectorsearchpb.EndpointStatusStatePb) (*EndpointStatusState, error) {
	if pb == nil {
		return nil, nil
	}
	st := EndpointStatusState(*pb)
	return &st, nil
}

// Type of endpoint.
type EndpointType string

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

// Values returns all possible values for EndpointType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointType) Values() []EndpointType {
	return []EndpointType{
		EndpointTypeStandard,
	}
}

// Type always returns EndpointType to satisfy [pflag.Value] interface
func (f *EndpointType) Type() string {
	return "EndpointType"
}

func EndpointTypeToPb(st *EndpointType) (*vectorsearchpb.EndpointTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := vectorsearchpb.EndpointTypePb(*st)
	return &pb, nil
}

func EndpointTypeFromPb(pb *vectorsearchpb.EndpointTypePb) (*EndpointType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EndpointType(*pb)
	return &st, nil
}

type GetEndpointRequest struct {
	// Name of the endpoint
	// Wire name: 'endpoint_name'
	EndpointName string `tf:"-"`
}

func GetEndpointRequestToPb(st *GetEndpointRequest) (*vectorsearchpb.GetEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.GetEndpointRequestPb{}
	pb.EndpointName = st.EndpointName

	return pb, nil
}

func GetEndpointRequestFromPb(pb *vectorsearchpb.GetEndpointRequestPb) (*GetEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEndpointRequest{}
	st.EndpointName = pb.EndpointName

	return st, nil
}

type GetIndexRequest struct {
	// Name of the index
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
}

func GetIndexRequestToPb(st *GetIndexRequest) (*vectorsearchpb.GetIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.GetIndexRequestPb{}
	pb.IndexName = st.IndexName

	return pb, nil
}

func GetIndexRequestFromPb(pb *vectorsearchpb.GetIndexRequestPb) (*GetIndexRequest, error) {
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
	Endpoints []EndpointInfo ``
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListEndpointResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListEndpointResponseToPb(st *ListEndpointResponse) (*vectorsearchpb.ListEndpointResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.ListEndpointResponsePb{}

	var endpointsPb []vectorsearchpb.EndpointInfoPb
	for _, item := range st.Endpoints {
		itemPb, err := EndpointInfoToPb(&item)
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

func ListEndpointResponseFromPb(pb *vectorsearchpb.ListEndpointResponsePb) (*ListEndpointResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListEndpointResponse{}

	var endpointsField []EndpointInfo
	for _, itemPb := range pb.Endpoints {
		item, err := EndpointInfoFromPb(&itemPb)
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

type ListEndpointsRequest struct {
	// Token for pagination
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListEndpointsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListEndpointsRequestToPb(st *ListEndpointsRequest) (*vectorsearchpb.ListEndpointsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.ListEndpointsRequestPb{}
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListEndpointsRequestFromPb(pb *vectorsearchpb.ListEndpointsRequestPb) (*ListEndpointsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListEndpointsRequest{}
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListIndexesRequest struct {
	// Name of the endpoint
	// Wire name: 'endpoint_name'
	EndpointName string `tf:"-"`
	// Token for pagination
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListIndexesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListIndexesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListIndexesRequestToPb(st *ListIndexesRequest) (*vectorsearchpb.ListIndexesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.ListIndexesRequestPb{}
	pb.EndpointName = st.EndpointName
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListIndexesRequestFromPb(pb *vectorsearchpb.ListIndexesRequestPb) (*ListIndexesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListIndexesRequest{}
	st.EndpointName = pb.EndpointName
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListValue struct {
	// Repeated field of dynamically typed values.
	// Wire name: 'values'
	Values []Value ``
}

func ListValueToPb(st *ListValue) (*vectorsearchpb.ListValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.ListValuePb{}

	var valuesPb []vectorsearchpb.ValuePb
	for _, item := range st.Values {
		itemPb, err := ValueToPb(&item)
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

func ListValueFromPb(pb *vectorsearchpb.ListValuePb) (*ListValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListValue{}

	var valuesField []Value
	for _, itemPb := range pb.Values {
		item, err := ValueFromPb(&itemPb)
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
	NextPageToken string ``

	// Wire name: 'vector_indexes'
	VectorIndexes   []MiniVectorIndex ``
	ForceSendFields []string          `tf:"-"`
}

func (s *ListVectorIndexesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVectorIndexesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListVectorIndexesResponseToPb(st *ListVectorIndexesResponse) (*vectorsearchpb.ListVectorIndexesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.ListVectorIndexesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var vectorIndexesPb []vectorsearchpb.MiniVectorIndexPb
	for _, item := range st.VectorIndexes {
		itemPb, err := MiniVectorIndexToPb(&item)
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

func ListVectorIndexesResponseFromPb(pb *vectorsearchpb.ListVectorIndexesResponsePb) (*ListVectorIndexesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListVectorIndexesResponse{}
	st.NextPageToken = pb.NextPageToken

	var vectorIndexesField []MiniVectorIndex
	for _, itemPb := range pb.VectorIndexes {
		item, err := MiniVectorIndexFromPb(&itemPb)
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

// Key-value pair.
type MapStringValueEntry struct {
	// Column name.
	// Wire name: 'key'
	Key string ``
	// Column value, nullable.
	// Wire name: 'value'
	Value           *Value   ``
	ForceSendFields []string `tf:"-"`
}

func (s *MapStringValueEntry) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MapStringValueEntry) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func MapStringValueEntryToPb(st *MapStringValueEntry) (*vectorsearchpb.MapStringValueEntryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.MapStringValueEntryPb{}
	pb.Key = st.Key
	valuePb, err := ValueToPb(st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = valuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func MapStringValueEntryFromPb(pb *vectorsearchpb.MapStringValueEntryPb) (*MapStringValueEntry, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MapStringValueEntry{}
	st.Key = pb.Key
	valueField, err := ValueFromPb(pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = valueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type MiniVectorIndex struct {
	// The user who created the index.
	// Wire name: 'creator'
	Creator string ``
	// Name of the endpoint associated with the index
	// Wire name: 'endpoint_name'
	EndpointName string ``

	// Wire name: 'index_type'
	IndexType VectorIndexType ``
	// Name of the index
	// Wire name: 'name'
	Name string ``
	// Primary key of the index
	// Wire name: 'primary_key'
	PrimaryKey      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *MiniVectorIndex) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MiniVectorIndex) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func MiniVectorIndexToPb(st *MiniVectorIndex) (*vectorsearchpb.MiniVectorIndexPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.MiniVectorIndexPb{}
	pb.Creator = st.Creator
	pb.EndpointName = st.EndpointName
	indexTypePb, err := VectorIndexTypeToPb(&st.IndexType)
	if err != nil {
		return nil, err
	}
	if indexTypePb != nil {
		pb.IndexType = *indexTypePb
	}
	pb.Name = st.Name
	pb.PrimaryKey = st.PrimaryKey

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func MiniVectorIndexFromPb(pb *vectorsearchpb.MiniVectorIndexPb) (*MiniVectorIndex, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MiniVectorIndex{}
	st.Creator = pb.Creator
	st.EndpointName = pb.EndpointName
	indexTypeField, err := VectorIndexTypeFromPb(&pb.IndexType)
	if err != nil {
		return nil, err
	}
	if indexTypeField != nil {
		st.IndexType = *indexTypeField
	}
	st.Name = pb.Name
	st.PrimaryKey = pb.PrimaryKey

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type PatchEndpointBudgetPolicyRequest struct {
	// The budget policy id to be applied
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string ``
	// Name of the vector search endpoint
	// Wire name: 'endpoint_name'
	EndpointName string `tf:"-"`
}

func PatchEndpointBudgetPolicyRequestToPb(st *PatchEndpointBudgetPolicyRequest) (*vectorsearchpb.PatchEndpointBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.PatchEndpointBudgetPolicyRequestPb{}
	pb.BudgetPolicyId = st.BudgetPolicyId
	pb.EndpointName = st.EndpointName

	return pb, nil
}

func PatchEndpointBudgetPolicyRequestFromPb(pb *vectorsearchpb.PatchEndpointBudgetPolicyRequestPb) (*PatchEndpointBudgetPolicyRequest, error) {
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
	EffectiveBudgetPolicyId string   ``
	ForceSendFields         []string `tf:"-"`
}

func (s *PatchEndpointBudgetPolicyResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PatchEndpointBudgetPolicyResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PatchEndpointBudgetPolicyResponseToPb(st *PatchEndpointBudgetPolicyResponse) (*vectorsearchpb.PatchEndpointBudgetPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.PatchEndpointBudgetPolicyResponsePb{}
	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PatchEndpointBudgetPolicyResponseFromPb(pb *vectorsearchpb.PatchEndpointBudgetPolicyResponsePb) (*PatchEndpointBudgetPolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PatchEndpointBudgetPolicyResponse{}
	st.EffectiveBudgetPolicyId = pb.EffectiveBudgetPolicyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the triggered
// execution mode, the system stops processing after successfully refreshing the
// source table in the pipeline once, ensuring the table is updated based on the
// data available when the update started. - `CONTINUOUS`: If the pipeline uses
// continuous execution, the pipeline processes new data as it arrives in the
// source table to keep vector index fresh.
type PipelineType string

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

// Values returns all possible values for PipelineType.
//
// There is no guarantee on the order of the values in the slice.
func (f *PipelineType) Values() []PipelineType {
	return []PipelineType{
		PipelineTypeContinuous,
		PipelineTypeTriggered,
	}
}

// Type always returns PipelineType to satisfy [pflag.Value] interface
func (f *PipelineType) Type() string {
	return "PipelineType"
}

func PipelineTypeToPb(st *PipelineType) (*vectorsearchpb.PipelineTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := vectorsearchpb.PipelineTypePb(*st)
	return &pb, nil
}

func PipelineTypeFromPb(pb *vectorsearchpb.PipelineTypePb) (*PipelineType, error) {
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
	EndpointName string ``
	// Name of the vector index to query.
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
	// Page token returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` API.
	// Wire name: 'page_token'
	PageToken       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *QueryVectorIndexNextPageRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryVectorIndexNextPageRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryVectorIndexNextPageRequestToPb(st *QueryVectorIndexNextPageRequest) (*vectorsearchpb.QueryVectorIndexNextPageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.QueryVectorIndexNextPageRequestPb{}
	pb.EndpointName = st.EndpointName
	pb.IndexName = st.IndexName
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func QueryVectorIndexNextPageRequestFromPb(pb *vectorsearchpb.QueryVectorIndexNextPageRequestPb) (*QueryVectorIndexNextPageRequest, error) {
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

type QueryVectorIndexRequest struct {
	// List of column names to include in the response.
	// Wire name: 'columns'
	Columns []string ``
	// Column names used to retrieve data to send to the reranker.
	// Wire name: 'columns_to_rerank'
	ColumnsToRerank []string ``
	// JSON string representing query filters.
	//
	// Example filters:
	//
	// - `{"id <": 5}`: Filter for id less than 5. - `{"id >": 5}`: Filter for
	// id greater than 5. - `{"id <=": 5}`: Filter for id less than equal to 5.
	// - `{"id >=": 5}`: Filter for id greater than equal to 5. - `{"id": 5}`:
	// Filter for id equal to 5.
	// Wire name: 'filters_json'
	FiltersJson string ``
	// Name of the vector index to query.
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
	// Number of results to return. Defaults to 10.
	// Wire name: 'num_results'
	NumResults int ``
	// Query text. Required for Delta Sync Index using model endpoint.
	// Wire name: 'query_text'
	QueryText string ``
	// The query type to use. Choices are `ANN` and `HYBRID`. Defaults to `ANN`.
	// Wire name: 'query_type'
	QueryType string ``
	// Query vector. Required for Direct Vector Access Index and Delta Sync
	// Index using self-managed vectors.
	// Wire name: 'query_vector'
	QueryVector []float64 ``
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	// Wire name: 'score_threshold'
	ScoreThreshold  float64  ``
	ForceSendFields []string `tf:"-"`
}

func (s *QueryVectorIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryVectorIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryVectorIndexRequestToPb(st *QueryVectorIndexRequest) (*vectorsearchpb.QueryVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.QueryVectorIndexRequestPb{}
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

func QueryVectorIndexRequestFromPb(pb *vectorsearchpb.QueryVectorIndexRequestPb) (*QueryVectorIndexRequest, error) {
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

type QueryVectorIndexResponse struct {
	// Metadata about the result set.
	// Wire name: 'manifest'
	Manifest *ResultManifest ``
	// [Optional] Token that can be used in `QueryVectorIndexNextPage` API to
	// get next page of results. If more than 1000 results satisfy the query,
	// they are returned in groups of 1000. Empty value means no more results.
	// The maximum number of results that can be returned is 10,000.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// Data returned in the query result.
	// Wire name: 'result'
	Result          *ResultData ``
	ForceSendFields []string    `tf:"-"`
}

func (s *QueryVectorIndexResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryVectorIndexResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryVectorIndexResponseToPb(st *QueryVectorIndexResponse) (*vectorsearchpb.QueryVectorIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.QueryVectorIndexResponsePb{}
	manifestPb, err := ResultManifestToPb(st.Manifest)
	if err != nil {
		return nil, err
	}
	if manifestPb != nil {
		pb.Manifest = manifestPb
	}
	pb.NextPageToken = st.NextPageToken
	resultPb, err := ResultDataToPb(st.Result)
	if err != nil {
		return nil, err
	}
	if resultPb != nil {
		pb.Result = resultPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func QueryVectorIndexResponseFromPb(pb *vectorsearchpb.QueryVectorIndexResponsePb) (*QueryVectorIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryVectorIndexResponse{}
	manifestField, err := ResultManifestFromPb(pb.Manifest)
	if err != nil {
		return nil, err
	}
	if manifestField != nil {
		st.Manifest = manifestField
	}
	st.NextPageToken = pb.NextPageToken
	resultField, err := ResultDataFromPb(pb.Result)
	if err != nil {
		return nil, err
	}
	if resultField != nil {
		st.Result = resultField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Data returned in the query result.
type ResultData struct {
	// Data rows returned in the query.
	// Wire name: 'data_array'
	DataArray [][]string ``
	// Number of rows in the result set.
	// Wire name: 'row_count'
	RowCount        int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *ResultData) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultData) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ResultDataToPb(st *ResultData) (*vectorsearchpb.ResultDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.ResultDataPb{}
	pb.DataArray = st.DataArray
	pb.RowCount = st.RowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ResultDataFromPb(pb *vectorsearchpb.ResultDataPb) (*ResultData, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultData{}
	st.DataArray = pb.DataArray
	st.RowCount = pb.RowCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Metadata about the result set.
type ResultManifest struct {
	// Number of columns in the result set.
	// Wire name: 'column_count'
	ColumnCount int ``
	// Information about each column in the result set.
	// Wire name: 'columns'
	Columns         []ColumnInfo ``
	ForceSendFields []string     `tf:"-"`
}

func (s *ResultManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultManifest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ResultManifestToPb(st *ResultManifest) (*vectorsearchpb.ResultManifestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.ResultManifestPb{}
	pb.ColumnCount = st.ColumnCount

	var columnsPb []vectorsearchpb.ColumnInfoPb
	for _, item := range st.Columns {
		itemPb, err := ColumnInfoToPb(&item)
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

func ResultManifestFromPb(pb *vectorsearchpb.ResultManifestPb) (*ResultManifest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultManifest{}
	st.ColumnCount = pb.ColumnCount

	var columnsField []ColumnInfo
	for _, itemPb := range pb.Columns {
		item, err := ColumnInfoFromPb(&itemPb)
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

type ScanVectorIndexRequest struct {
	// Name of the vector index to scan.
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
	// Primary key of the last entry returned in the previous scan.
	// Wire name: 'last_primary_key'
	LastPrimaryKey string ``
	// Number of results to return. Defaults to 10.
	// Wire name: 'num_results'
	NumResults      int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *ScanVectorIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ScanVectorIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ScanVectorIndexRequestToPb(st *ScanVectorIndexRequest) (*vectorsearchpb.ScanVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.ScanVectorIndexRequestPb{}
	pb.IndexName = st.IndexName
	pb.LastPrimaryKey = st.LastPrimaryKey
	pb.NumResults = st.NumResults

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ScanVectorIndexRequestFromPb(pb *vectorsearchpb.ScanVectorIndexRequestPb) (*ScanVectorIndexRequest, error) {
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

// Response to a scan vector index request.
type ScanVectorIndexResponse struct {
	// List of data entries
	// Wire name: 'data'
	Data []Struct ``
	// Primary key of the last entry.
	// Wire name: 'last_primary_key'
	LastPrimaryKey  string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ScanVectorIndexResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ScanVectorIndexResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ScanVectorIndexResponseToPb(st *ScanVectorIndexResponse) (*vectorsearchpb.ScanVectorIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.ScanVectorIndexResponsePb{}

	var dataPb []vectorsearchpb.StructPb
	for _, item := range st.Data {
		itemPb, err := StructToPb(&item)
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

func ScanVectorIndexResponseFromPb(pb *vectorsearchpb.ScanVectorIndexResponsePb) (*ScanVectorIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ScanVectorIndexResponse{}

	var dataField []Struct
	for _, itemPb := range pb.Data {
		item, err := StructFromPb(&itemPb)
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

type Struct struct {
	// Data entry, corresponding to a row in a vector index.
	// Wire name: 'fields'
	Fields []MapStringValueEntry ``
}

func StructToPb(st *Struct) (*vectorsearchpb.StructPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.StructPb{}

	var fieldsPb []vectorsearchpb.MapStringValueEntryPb
	for _, item := range st.Fields {
		itemPb, err := MapStringValueEntryToPb(&item)
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

func StructFromPb(pb *vectorsearchpb.StructPb) (*Struct, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Struct{}

	var fieldsField []MapStringValueEntry
	for _, itemPb := range pb.Fields {
		item, err := MapStringValueEntryFromPb(&itemPb)
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

type SyncIndexRequest struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	// Wire name: 'index_name'
	IndexName string `tf:"-"`
}

func SyncIndexRequestToPb(st *SyncIndexRequest) (*vectorsearchpb.SyncIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.SyncIndexRequestPb{}
	pb.IndexName = st.IndexName

	return pb, nil
}

func SyncIndexRequestFromPb(pb *vectorsearchpb.SyncIndexRequestPb) (*SyncIndexRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SyncIndexRequest{}
	st.IndexName = pb.IndexName

	return st, nil
}

type UpdateEndpointCustomTagsRequest struct {
	// The new custom tags for the vector search endpoint
	// Wire name: 'custom_tags'
	CustomTags []CustomTag ``
	// Name of the vector search endpoint
	// Wire name: 'endpoint_name'
	EndpointName string `tf:"-"`
}

func UpdateEndpointCustomTagsRequestToPb(st *UpdateEndpointCustomTagsRequest) (*vectorsearchpb.UpdateEndpointCustomTagsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.UpdateEndpointCustomTagsRequestPb{}

	var customTagsPb []vectorsearchpb.CustomTagPb
	for _, item := range st.CustomTags {
		itemPb, err := CustomTagToPb(&item)
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

func UpdateEndpointCustomTagsRequestFromPb(pb *vectorsearchpb.UpdateEndpointCustomTagsRequestPb) (*UpdateEndpointCustomTagsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEndpointCustomTagsRequest{}

	var customTagsField []CustomTag
	for _, itemPb := range pb.CustomTags {
		item, err := CustomTagFromPb(&itemPb)
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
	CustomTags []CustomTag ``
	// The name of the vector search endpoint whose custom tags were updated.
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateEndpointCustomTagsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateEndpointCustomTagsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateEndpointCustomTagsResponseToPb(st *UpdateEndpointCustomTagsResponse) (*vectorsearchpb.UpdateEndpointCustomTagsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.UpdateEndpointCustomTagsResponsePb{}

	var customTagsPb []vectorsearchpb.CustomTagPb
	for _, item := range st.CustomTags {
		itemPb, err := CustomTagToPb(&item)
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

func UpdateEndpointCustomTagsResponseFromPb(pb *vectorsearchpb.UpdateEndpointCustomTagsResponsePb) (*UpdateEndpointCustomTagsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEndpointCustomTagsResponse{}

	var customTagsField []CustomTag
	for _, itemPb := range pb.CustomTags {
		item, err := CustomTagFromPb(&itemPb)
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

type UpsertDataResult struct {
	// List of primary keys for rows that failed to process.
	// Wire name: 'failed_primary_keys'
	FailedPrimaryKeys []string ``
	// Count of successfully processed rows.
	// Wire name: 'success_row_count'
	SuccessRowCount int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpsertDataResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpsertDataResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpsertDataResultToPb(st *UpsertDataResult) (*vectorsearchpb.UpsertDataResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.UpsertDataResultPb{}
	pb.FailedPrimaryKeys = st.FailedPrimaryKeys
	pb.SuccessRowCount = st.SuccessRowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpsertDataResultFromPb(pb *vectorsearchpb.UpsertDataResultPb) (*UpsertDataResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpsertDataResult{}
	st.FailedPrimaryKeys = pb.FailedPrimaryKeys
	st.SuccessRowCount = pb.SuccessRowCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpsertDataStatus string

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

// Values returns all possible values for UpsertDataStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *UpsertDataStatus) Values() []UpsertDataStatus {
	return []UpsertDataStatus{
		UpsertDataStatusFailure,
		UpsertDataStatusPartialSuccess,
		UpsertDataStatusSuccess,
	}
}

// Type always returns UpsertDataStatus to satisfy [pflag.Value] interface
func (f *UpsertDataStatus) Type() string {
	return "UpsertDataStatus"
}

func UpsertDataStatusToPb(st *UpsertDataStatus) (*vectorsearchpb.UpsertDataStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := vectorsearchpb.UpsertDataStatusPb(*st)
	return &pb, nil
}

func UpsertDataStatusFromPb(pb *vectorsearchpb.UpsertDataStatusPb) (*UpsertDataStatus, error) {
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
	InputsJson string ``
}

func UpsertDataVectorIndexRequestToPb(st *UpsertDataVectorIndexRequest) (*vectorsearchpb.UpsertDataVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.UpsertDataVectorIndexRequestPb{}
	pb.IndexName = st.IndexName
	pb.InputsJson = st.InputsJson

	return pb, nil
}

func UpsertDataVectorIndexRequestFromPb(pb *vectorsearchpb.UpsertDataVectorIndexRequestPb) (*UpsertDataVectorIndexRequest, error) {
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
	Result *UpsertDataResult ``
	// Status of the upsert operation.
	// Wire name: 'status'
	Status UpsertDataStatus ``
}

func UpsertDataVectorIndexResponseToPb(st *UpsertDataVectorIndexResponse) (*vectorsearchpb.UpsertDataVectorIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.UpsertDataVectorIndexResponsePb{}
	resultPb, err := UpsertDataResultToPb(st.Result)
	if err != nil {
		return nil, err
	}
	if resultPb != nil {
		pb.Result = resultPb
	}
	statusPb, err := UpsertDataStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	return pb, nil
}

func UpsertDataVectorIndexResponseFromPb(pb *vectorsearchpb.UpsertDataVectorIndexResponsePb) (*UpsertDataVectorIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpsertDataVectorIndexResponse{}
	resultField, err := UpsertDataResultFromPb(pb.Result)
	if err != nil {
		return nil, err
	}
	if resultField != nil {
		st.Result = resultField
	}
	statusField, err := UpsertDataStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	return st, nil
}

type Value struct {

	// Wire name: 'bool_value'
	BoolValue bool ``

	// Wire name: 'list_value'
	ListValue *ListValue ``

	// Wire name: 'number_value'
	NumberValue float64 ``

	// Wire name: 'string_value'
	StringValue string ``

	// Wire name: 'struct_value'
	StructValue     *Struct  ``
	ForceSendFields []string `tf:"-"`
}

func (s *Value) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Value) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ValueToPb(st *Value) (*vectorsearchpb.ValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.ValuePb{}
	pb.BoolValue = st.BoolValue
	listValuePb, err := ListValueToPb(st.ListValue)
	if err != nil {
		return nil, err
	}
	if listValuePb != nil {
		pb.ListValue = listValuePb
	}
	pb.NumberValue = st.NumberValue
	pb.StringValue = st.StringValue
	structValuePb, err := StructToPb(st.StructValue)
	if err != nil {
		return nil, err
	}
	if structValuePb != nil {
		pb.StructValue = structValuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ValueFromPb(pb *vectorsearchpb.ValuePb) (*Value, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Value{}
	st.BoolValue = pb.BoolValue
	listValueField, err := ListValueFromPb(pb.ListValue)
	if err != nil {
		return nil, err
	}
	if listValueField != nil {
		st.ListValue = listValueField
	}
	st.NumberValue = pb.NumberValue
	st.StringValue = pb.StringValue
	structValueField, err := StructFromPb(pb.StructValue)
	if err != nil {
		return nil, err
	}
	if structValueField != nil {
		st.StructValue = structValueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type VectorIndex struct {
	// The user who created the index.
	// Wire name: 'creator'
	Creator string ``

	// Wire name: 'delta_sync_index_spec'
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecResponse ``

	// Wire name: 'direct_access_index_spec'
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec ``
	// Name of the endpoint associated with the index
	// Wire name: 'endpoint_name'
	EndpointName string ``

	// Wire name: 'index_type'
	IndexType VectorIndexType ``
	// Name of the index
	// Wire name: 'name'
	Name string ``
	// Primary key of the index
	// Wire name: 'primary_key'
	PrimaryKey string ``

	// Wire name: 'status'
	Status          *VectorIndexStatus ``
	ForceSendFields []string           `tf:"-"`
}

func (s *VectorIndex) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorIndex) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func VectorIndexToPb(st *VectorIndex) (*vectorsearchpb.VectorIndexPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.VectorIndexPb{}
	pb.Creator = st.Creator
	deltaSyncIndexSpecPb, err := DeltaSyncVectorIndexSpecResponseToPb(st.DeltaSyncIndexSpec)
	if err != nil {
		return nil, err
	}
	if deltaSyncIndexSpecPb != nil {
		pb.DeltaSyncIndexSpec = deltaSyncIndexSpecPb
	}
	directAccessIndexSpecPb, err := DirectAccessVectorIndexSpecToPb(st.DirectAccessIndexSpec)
	if err != nil {
		return nil, err
	}
	if directAccessIndexSpecPb != nil {
		pb.DirectAccessIndexSpec = directAccessIndexSpecPb
	}
	pb.EndpointName = st.EndpointName
	indexTypePb, err := VectorIndexTypeToPb(&st.IndexType)
	if err != nil {
		return nil, err
	}
	if indexTypePb != nil {
		pb.IndexType = *indexTypePb
	}
	pb.Name = st.Name
	pb.PrimaryKey = st.PrimaryKey
	statusPb, err := VectorIndexStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func VectorIndexFromPb(pb *vectorsearchpb.VectorIndexPb) (*VectorIndex, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VectorIndex{}
	st.Creator = pb.Creator
	deltaSyncIndexSpecField, err := DeltaSyncVectorIndexSpecResponseFromPb(pb.DeltaSyncIndexSpec)
	if err != nil {
		return nil, err
	}
	if deltaSyncIndexSpecField != nil {
		st.DeltaSyncIndexSpec = deltaSyncIndexSpecField
	}
	directAccessIndexSpecField, err := DirectAccessVectorIndexSpecFromPb(pb.DirectAccessIndexSpec)
	if err != nil {
		return nil, err
	}
	if directAccessIndexSpecField != nil {
		st.DirectAccessIndexSpec = directAccessIndexSpecField
	}
	st.EndpointName = pb.EndpointName
	indexTypeField, err := VectorIndexTypeFromPb(&pb.IndexType)
	if err != nil {
		return nil, err
	}
	if indexTypeField != nil {
		st.IndexType = *indexTypeField
	}
	st.Name = pb.Name
	st.PrimaryKey = pb.PrimaryKey
	statusField, err := VectorIndexStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type VectorIndexStatus struct {
	// Index API Url to be used to perform operations on the index
	// Wire name: 'index_url'
	IndexUrl string ``
	// Number of rows indexed
	// Wire name: 'indexed_row_count'
	IndexedRowCount int64 ``
	// Message associated with the index status
	// Wire name: 'message'
	Message string ``
	// Whether the index is ready for search
	// Wire name: 'ready'
	Ready           bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *VectorIndexStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorIndexStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func VectorIndexStatusToPb(st *VectorIndexStatus) (*vectorsearchpb.VectorIndexStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorsearchpb.VectorIndexStatusPb{}
	pb.IndexUrl = st.IndexUrl
	pb.IndexedRowCount = st.IndexedRowCount
	pb.Message = st.Message
	pb.Ready = st.Ready

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func VectorIndexStatusFromPb(pb *vectorsearchpb.VectorIndexStatusPb) (*VectorIndexStatus, error) {
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

// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
// automatically syncs with a source Delta Table, automatically and
// incrementally updating the index as the underlying data in the Delta Table
// changes. - `DIRECT_ACCESS`: An index that supports direct read and write of
// vectors and metadata through our REST and SDK APIs. With this model, the user
// manages index updates.
type VectorIndexType string

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

// Values returns all possible values for VectorIndexType.
//
// There is no guarantee on the order of the values in the slice.
func (f *VectorIndexType) Values() []VectorIndexType {
	return []VectorIndexType{
		VectorIndexTypeDeltaSync,
		VectorIndexTypeDirectAccess,
	}
}

// Type always returns VectorIndexType to satisfy [pflag.Value] interface
func (f *VectorIndexType) Type() string {
	return "VectorIndexType"
}

func VectorIndexTypeToPb(st *VectorIndexType) (*vectorsearchpb.VectorIndexTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := vectorsearchpb.VectorIndexTypePb(*st)
	return &pb, nil
}

func VectorIndexTypeFromPb(pb *vectorsearchpb.VectorIndexTypePb) (*VectorIndexType, error) {
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
	s := fmt.Sprintf("%.9fs", d.Seconds())
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

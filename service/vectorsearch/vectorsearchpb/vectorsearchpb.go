// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearchpb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type ColumnInfoPb struct {
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ColumnInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ColumnInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateEndpointPb struct {
	BudgetPolicyId string         `json:"budget_policy_id,omitempty"`
	EndpointType   EndpointTypePb `json:"endpoint_type"`
	Name           string         `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateEndpointPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateEndpointPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateVectorIndexRequestPb struct {
	DeltaSyncIndexSpec    *DeltaSyncVectorIndexSpecRequestPb `json:"delta_sync_index_spec,omitempty"`
	DirectAccessIndexSpec *DirectAccessVectorIndexSpecPb     `json:"direct_access_index_spec,omitempty"`
	EndpointName          string                             `json:"endpoint_name"`
	IndexType             VectorIndexTypePb                  `json:"index_type"`
	Name                  string                             `json:"name"`
	PrimaryKey            string                             `json:"primary_key"`
}

type CustomTagPb struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CustomTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CustomTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDataResultPb struct {
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	SuccessRowCount   int64    `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteDataResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteDataResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDataStatusPb string

const DeleteDataStatusFailure DeleteDataStatusPb = `FAILURE`

const DeleteDataStatusPartialSuccess DeleteDataStatusPb = `PARTIAL_SUCCESS`

const DeleteDataStatusSuccess DeleteDataStatusPb = `SUCCESS`

type DeleteDataVectorIndexRequestPb struct {
	IndexName   string   `json:"-" url:"-"`
	PrimaryKeys []string `json:"-" url:"primary_keys"`
}

type DeleteDataVectorIndexResponsePb struct {
	Result *DeleteDataResultPb `json:"result,omitempty"`
	Status DeleteDataStatusPb  `json:"status,omitempty"`
}

type DeleteEndpointRequestPb struct {
	EndpointName string `json:"-" url:"-"`
}

type DeleteEndpointResponsePb struct {
}

type DeleteIndexRequestPb struct {
	IndexName string `json:"-" url:"-"`
}

type DeleteIndexResponsePb struct {
}

type DeltaSyncVectorIndexSpecRequestPb struct {
	ColumnsToSync           []string                  `json:"columns_to_sync,omitempty"`
	EmbeddingSourceColumns  []EmbeddingSourceColumnPb `json:"embedding_source_columns,omitempty"`
	EmbeddingVectorColumns  []EmbeddingVectorColumnPb `json:"embedding_vector_columns,omitempty"`
	EmbeddingWritebackTable string                    `json:"embedding_writeback_table,omitempty"`
	PipelineType            PipelineTypePb            `json:"pipeline_type,omitempty"`
	SourceTable             string                    `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeltaSyncVectorIndexSpecRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeltaSyncVectorIndexSpecRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeltaSyncVectorIndexSpecResponsePb struct {
	EmbeddingSourceColumns  []EmbeddingSourceColumnPb `json:"embedding_source_columns,omitempty"`
	EmbeddingVectorColumns  []EmbeddingVectorColumnPb `json:"embedding_vector_columns,omitempty"`
	EmbeddingWritebackTable string                    `json:"embedding_writeback_table,omitempty"`
	PipelineId              string                    `json:"pipeline_id,omitempty"`
	PipelineType            PipelineTypePb            `json:"pipeline_type,omitempty"`
	SourceTable             string                    `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeltaSyncVectorIndexSpecResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeltaSyncVectorIndexSpecResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DirectAccessVectorIndexSpecPb struct {
	EmbeddingSourceColumns []EmbeddingSourceColumnPb `json:"embedding_source_columns,omitempty"`
	EmbeddingVectorColumns []EmbeddingVectorColumnPb `json:"embedding_vector_columns,omitempty"`
	SchemaJson             string                    `json:"schema_json,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DirectAccessVectorIndexSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DirectAccessVectorIndexSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EmbeddingSourceColumnPb struct {
	EmbeddingModelEndpointName string `json:"embedding_model_endpoint_name,omitempty"`
	Name                       string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EmbeddingSourceColumnPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EmbeddingSourceColumnPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EmbeddingVectorColumnPb struct {
	EmbeddingDimension int    `json:"embedding_dimension,omitempty"`
	Name               string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EmbeddingVectorColumnPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EmbeddingVectorColumnPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointInfoPb struct {
	CreationTimestamp       int64             `json:"creation_timestamp,omitempty"`
	Creator                 string            `json:"creator,omitempty"`
	CustomTags              []CustomTagPb     `json:"custom_tags,omitempty"`
	EffectiveBudgetPolicyId string            `json:"effective_budget_policy_id,omitempty"`
	EndpointStatus          *EndpointStatusPb `json:"endpoint_status,omitempty"`
	EndpointType            EndpointTypePb    `json:"endpoint_type,omitempty"`
	Id                      string            `json:"id,omitempty"`
	LastUpdatedTimestamp    int64             `json:"last_updated_timestamp,omitempty"`
	LastUpdatedUser         string            `json:"last_updated_user,omitempty"`
	Name                    string            `json:"name,omitempty"`
	NumIndexes              int               `json:"num_indexes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EndpointInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EndpointInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointStatusPb struct {
	Message string                `json:"message,omitempty"`
	State   EndpointStatusStatePb `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EndpointStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EndpointStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointStatusStatePb string

const EndpointStatusStateOffline EndpointStatusStatePb = `OFFLINE`

const EndpointStatusStateOnline EndpointStatusStatePb = `ONLINE`

const EndpointStatusStateProvisioning EndpointStatusStatePb = `PROVISIONING`

type EndpointTypePb string

const EndpointTypeStandard EndpointTypePb = `STANDARD`

type GetEndpointRequestPb struct {
	EndpointName string `json:"-" url:"-"`
}

type GetIndexRequestPb struct {
	IndexName string `json:"-" url:"-"`
}

type ListEndpointResponsePb struct {
	Endpoints     []EndpointInfoPb `json:"endpoints,omitempty"`
	NextPageToken string           `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListEndpointResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListEndpointResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListEndpointsRequestPb struct {
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListEndpointsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListEndpointsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListIndexesRequestPb struct {
	EndpointName string `json:"-" url:"endpoint_name"`
	PageToken    string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListIndexesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListIndexesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListValuePb struct {
	Values []ValuePb `json:"values,omitempty"`
}

type ListVectorIndexesResponsePb struct {
	NextPageToken string              `json:"next_page_token,omitempty"`
	VectorIndexes []MiniVectorIndexPb `json:"vector_indexes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListVectorIndexesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListVectorIndexesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MapStringValueEntryPb struct {
	Key   string   `json:"key,omitempty"`
	Value *ValuePb `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MapStringValueEntryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MapStringValueEntryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MiniVectorIndexPb struct {
	Creator      string            `json:"creator,omitempty"`
	EndpointName string            `json:"endpoint_name,omitempty"`
	IndexType    VectorIndexTypePb `json:"index_type,omitempty"`
	Name         string            `json:"name,omitempty"`
	PrimaryKey   string            `json:"primary_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MiniVectorIndexPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MiniVectorIndexPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PatchEndpointBudgetPolicyRequestPb struct {
	BudgetPolicyId string `json:"budget_policy_id"`
	EndpointName   string `json:"-" url:"-"`
}

type PatchEndpointBudgetPolicyResponsePb struct {
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PatchEndpointBudgetPolicyResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PatchEndpointBudgetPolicyResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineTypePb string

// If the pipeline uses continuous execution, the pipeline processes new data as
// it arrives in the source table to keep vector index fresh.
const PipelineTypeContinuous PipelineTypePb = `CONTINUOUS`

// If the pipeline uses the triggered execution mode, the system stops
// processing after successfully refreshing the source table in the pipeline
// once, ensuring the table is updated based on the data available when the
// update started.
const PipelineTypeTriggered PipelineTypePb = `TRIGGERED`

type QueryVectorIndexNextPageRequestPb struct {
	EndpointName string `json:"endpoint_name,omitempty"`
	IndexName    string `json:"-" url:"-"`
	PageToken    string `json:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryVectorIndexNextPageRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryVectorIndexNextPageRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryVectorIndexRequestPb struct {
	Columns         []string  `json:"columns"`
	ColumnsToRerank []string  `json:"columns_to_rerank,omitempty"`
	FiltersJson     string    `json:"filters_json,omitempty"`
	IndexName       string    `json:"-" url:"-"`
	NumResults      int       `json:"num_results,omitempty"`
	QueryText       string    `json:"query_text,omitempty"`
	QueryType       string    `json:"query_type,omitempty"`
	QueryVector     []float64 `json:"query_vector,omitempty"`
	ScoreThreshold  float64   `json:"score_threshold,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryVectorIndexRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryVectorIndexRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryVectorIndexResponsePb struct {
	Manifest      *ResultManifestPb `json:"manifest,omitempty"`
	NextPageToken string            `json:"next_page_token,omitempty"`
	Result        *ResultDataPb     `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryVectorIndexResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryVectorIndexResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResultDataPb struct {
	DataArray [][]string `json:"data_array,omitempty"`
	RowCount  int        `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ResultDataPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ResultDataPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResultManifestPb struct {
	ColumnCount int            `json:"column_count,omitempty"`
	Columns     []ColumnInfoPb `json:"columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ResultManifestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ResultManifestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ScanVectorIndexRequestPb struct {
	IndexName      string `json:"-" url:"-"`
	LastPrimaryKey string `json:"last_primary_key,omitempty"`
	NumResults     int    `json:"num_results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ScanVectorIndexRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ScanVectorIndexRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ScanVectorIndexResponsePb struct {
	Data           []StructPb `json:"data,omitempty"`
	LastPrimaryKey string     `json:"last_primary_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ScanVectorIndexResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ScanVectorIndexResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StructPb struct {
	Fields []MapStringValueEntryPb `json:"fields,omitempty"`
}

type SyncIndexRequestPb struct {
	IndexName string `json:"-" url:"-"`
}

type SyncIndexResponsePb struct {
}

type UpdateEndpointCustomTagsRequestPb struct {
	CustomTags   []CustomTagPb `json:"custom_tags"`
	EndpointName string        `json:"-" url:"-"`
}

type UpdateEndpointCustomTagsResponsePb struct {
	CustomTags []CustomTagPb `json:"custom_tags,omitempty"`
	Name       string        `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateEndpointCustomTagsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateEndpointCustomTagsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpsertDataResultPb struct {
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	SuccessRowCount   int64    `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpsertDataResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpsertDataResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpsertDataStatusPb string

const UpsertDataStatusFailure UpsertDataStatusPb = `FAILURE`

const UpsertDataStatusPartialSuccess UpsertDataStatusPb = `PARTIAL_SUCCESS`

const UpsertDataStatusSuccess UpsertDataStatusPb = `SUCCESS`

type UpsertDataVectorIndexRequestPb struct {
	IndexName  string `json:"-" url:"-"`
	InputsJson string `json:"inputs_json"`
}

type UpsertDataVectorIndexResponsePb struct {
	Result *UpsertDataResultPb `json:"result,omitempty"`
	Status UpsertDataStatusPb  `json:"status,omitempty"`
}

type ValuePb struct {
	BoolValue   bool         `json:"bool_value,omitempty"`
	ListValue   *ListValuePb `json:"list_value,omitempty"`
	NumberValue float64      `json:"number_value,omitempty"`
	StringValue string       `json:"string_value,omitempty"`
	StructValue *StructPb    `json:"struct_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type VectorIndexPb struct {
	Creator               string                              `json:"creator,omitempty"`
	DeltaSyncIndexSpec    *DeltaSyncVectorIndexSpecResponsePb `json:"delta_sync_index_spec,omitempty"`
	DirectAccessIndexSpec *DirectAccessVectorIndexSpecPb      `json:"direct_access_index_spec,omitempty"`
	EndpointName          string                              `json:"endpoint_name,omitempty"`
	IndexType             VectorIndexTypePb                   `json:"index_type,omitempty"`
	Name                  string                              `json:"name,omitempty"`
	PrimaryKey            string                              `json:"primary_key,omitempty"`
	Status                *VectorIndexStatusPb                `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *VectorIndexPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st VectorIndexPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type VectorIndexStatusPb struct {
	IndexUrl        string `json:"index_url,omitempty"`
	IndexedRowCount int64  `json:"indexed_row_count,omitempty"`
	Message         string `json:"message,omitempty"`
	Ready           bool   `json:"ready,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *VectorIndexStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st VectorIndexStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type VectorIndexTypePb string

// An index that automatically syncs with a source Delta Table, automatically
// and incrementally updating the index as the underlying data in the Delta
// Table changes.
const VectorIndexTypeDeltaSync VectorIndexTypePb = `DELTA_SYNC`

// An index that supports direct read and write of vectors and metadata through
// our REST and SDK APIs. With this model, the user manages index updates.
const VectorIndexTypeDirectAccess VectorIndexTypePb = `DIRECT_ACCESS`

// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"encoding/json"
	"fmt"
)

type ColumnInfo struct {
	// Name of the column.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type CreateEndpoint struct {
	// The budget policy id to be applied
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// Type of endpoint
	// Wire name: 'endpoint_type'
	EndpointType EndpointType `json:"endpoint_type"`
	// Name of the vector search endpoint
	// Wire name: 'name'
	Name string `json:"name"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type CreateVectorIndexRequest struct {
	// Specification for Delta Sync Index. Required if `index_type` is
	// `DELTA_SYNC`.
	// Wire name: 'delta_sync_index_spec'
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecRequest `json:"delta_sync_index_spec,omitempty"`
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	// Wire name: 'direct_access_index_spec'
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec `json:"direct_access_index_spec,omitempty"`
	// Name of the endpoint to be used for serving the index
	// Wire name: 'endpoint_name'
	EndpointName string `json:"endpoint_name"`
	// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
	// automatically syncs with a source Delta Table, automatically and
	// incrementally updating the index as the underlying data in the Delta
	// Table changes. - `DIRECT_ACCESS`: An index that supports direct read and
	// write of vectors and metadata through our REST and SDK APIs. With this
	// model, the user manages index updates.
	// Wire name: 'index_type'
	IndexType VectorIndexType `json:"index_type"`
	// Name of the index
	// Wire name: 'name'
	Name string `json:"name"`
	// Primary key of the index
	// Wire name: 'primary_key'
	PrimaryKey string `json:"primary_key"`
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

type CustomTag struct {
	// Key field for a vector search endpoint tag.
	// Wire name: 'key'
	Key string `json:"key"`
	// [Optional] Value field for a vector search endpoint tag.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type DeleteDataResult struct {
	// List of primary keys for rows that failed to process.
	// Wire name: 'failed_primary_keys'
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	// Count of successfully processed rows.
	// Wire name: 'success_row_count'
	SuccessRowCount int64 `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// Delete data from index
type DeleteDataVectorIndexRequest struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	IndexName string `json:"-" tf:"-"`
	// List of primary keys for the data to be deleted.
	PrimaryKeys []string `json:"-" tf:"-"`
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

type DeleteDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	// Wire name: 'result'
	Result *DeleteDataResult `json:"result,omitempty"`
	// Status of the delete operation.
	// Wire name: 'status'
	Status DeleteDataStatus `json:"status,omitempty"`
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

// Delete an endpoint
type DeleteEndpointRequest struct {
	// Name of the vector search endpoint
	EndpointName string `json:"-" tf:"-"`
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

type DeleteEndpointResponse struct {
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

// Delete an index
type DeleteIndexRequest struct {
	// Name of the index
	IndexName string `json:"-" tf:"-"`
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

type DeleteIndexResponse struct {
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

type DeltaSyncVectorIndexSpecRequest struct {
	// [Optional] Select the columns to sync with the vector index. If you leave
	// this field blank, all columns from the source table are synced with the
	// index. The primary key column and embedding source column or embedding
	// vector column are always synced.
	// Wire name: 'columns_to_sync'
	ColumnsToSync []string `json:"columns_to_sync,omitempty"`
	// The columns that contain the embedding source.
	// Wire name: 'embedding_source_columns'
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// The columns that contain the embedding vectors.
	// Wire name: 'embedding_vector_columns'
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	// Wire name: 'embedding_writeback_table'
	EmbeddingWritebackTable string `json:"embedding_writeback_table,omitempty"`
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	// Wire name: 'pipeline_type'
	PipelineType PipelineType `json:"pipeline_type,omitempty"`
	// The name of the source table.
	// Wire name: 'source_table'
	SourceTable string `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type DeltaSyncVectorIndexSpecResponse struct {
	// The columns that contain the embedding source.
	// Wire name: 'embedding_source_columns'
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// The columns that contain the embedding vectors.
	// Wire name: 'embedding_vector_columns'
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	// Wire name: 'embedding_writeback_table'
	EmbeddingWritebackTable string `json:"embedding_writeback_table,omitempty"`
	// The ID of the pipeline that is used to sync the index.
	// Wire name: 'pipeline_id'
	PipelineId string `json:"pipeline_id,omitempty"`
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	// Wire name: 'pipeline_type'
	PipelineType PipelineType `json:"pipeline_type,omitempty"`
	// The name of the source table.
	// Wire name: 'source_table'
	SourceTable string `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type DirectAccessVectorIndexSpec struct {
	// The columns that contain the embedding source. The format should be
	// array[double].
	// Wire name: 'embedding_source_columns'
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// The columns that contain the embedding vectors. The format should be
	// array[double].
	// Wire name: 'embedding_vector_columns'
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	// The schema of the index in JSON format. Supported types are `integer`,
	// `long`, `float`, `double`, `boolean`, `string`, `date`, `timestamp`.
	// Supported types for vector column: `array<float>`, `array<double>`,`.
	// Wire name: 'schema_json'
	SchemaJson string `json:"schema_json,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type EmbeddingSourceColumn struct {
	// Name of the embedding model endpoint
	// Wire name: 'embedding_model_endpoint_name'
	EmbeddingModelEndpointName string `json:"embedding_model_endpoint_name,omitempty"`
	// Name of the column
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type EmbeddingVectorColumn struct {
	// Dimension of the embedding vector
	// Wire name: 'embedding_dimension'
	EmbeddingDimension int `json:"embedding_dimension,omitempty"`
	// Name of the column
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type EndpointInfo struct {
	// Timestamp of endpoint creation
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Creator of the endpoint
	// Wire name: 'creator'
	Creator string `json:"creator,omitempty"`
	// The custom tags assigned to the endpoint
	// Wire name: 'custom_tags'
	CustomTags []CustomTag `json:"custom_tags,omitempty"`
	// The budget policy id applied to the endpoint
	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// Current status of the endpoint
	// Wire name: 'endpoint_status'
	EndpointStatus *EndpointStatus `json:"endpoint_status,omitempty"`
	// Type of endpoint
	// Wire name: 'endpoint_type'
	EndpointType EndpointType `json:"endpoint_type,omitempty"`
	// Unique identifier of the endpoint
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Timestamp of last update to the endpoint
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// User who last updated the endpoint
	// Wire name: 'last_updated_user'
	LastUpdatedUser string `json:"last_updated_user,omitempty"`
	// Name of the vector search endpoint
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Number of indexes on the endpoint
	// Wire name: 'num_indexes'
	NumIndexes int `json:"num_indexes,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// Status information of an endpoint
type EndpointStatus struct {
	// Additional status message
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// Current state of the endpoint
	// Wire name: 'state'
	State EndpointStatusState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// Get an endpoint
type GetEndpointRequest struct {
	// Name of the endpoint
	EndpointName string `json:"-" tf:"-"`
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

// Get an index
type GetIndexRequest struct {
	// If true, the URL returned for the index is guaranteed to be compatible
	// with the reranker. Currently this means we return the CP URL regardless
	// of how the index is being accessed. If not set or set to false, the URL
	// may still be compatible with the reranker depending on what URL we
	// return.
	EnsureRerankerCompatible bool `json:"-" tf:"-"`
	// Name of the index
	IndexName string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type ListEndpointResponse struct {
	// An array of Endpoint objects
	// Wire name: 'endpoints'
	Endpoints []EndpointInfo `json:"endpoints,omitempty"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// List all endpoints
type ListEndpointsRequest struct {
	// Token for pagination
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// List indexes
type ListIndexesRequest struct {
	// Name of the endpoint
	EndpointName string `json:"-" tf:"-"`
	// Token for pagination
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type ListValue struct {
	// Repeated field of dynamically typed values.
	// Wire name: 'values'
	Values []Value `json:"values,omitempty"`
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

type ListVectorIndexesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'vector_indexes'
	VectorIndexes []MiniVectorIndex `json:"vector_indexes,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// Key-value pair.
type MapStringValueEntry struct {
	// Column name.
	// Wire name: 'key'
	Key string `json:"key,omitempty"`
	// Column value, nullable.
	// Wire name: 'value'
	Value *Value `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type MiniVectorIndex struct {
	// The user who created the index.
	// Wire name: 'creator'
	Creator string `json:"creator,omitempty"`
	// Name of the endpoint associated with the index
	// Wire name: 'endpoint_name'
	EndpointName string `json:"endpoint_name,omitempty"`
	// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
	// automatically syncs with a source Delta Table, automatically and
	// incrementally updating the index as the underlying data in the Delta
	// Table changes. - `DIRECT_ACCESS`: An index that supports direct read and
	// write of vectors and metadata through our REST and SDK APIs. With this
	// model, the user manages index updates.
	// Wire name: 'index_type'
	IndexType VectorIndexType `json:"index_type,omitempty"`
	// Name of the index
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Primary key of the index
	// Wire name: 'primary_key'
	PrimaryKey string `json:"primary_key,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type PatchEndpointBudgetPolicyRequest struct {
	// The budget policy id to be applied
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string `json:"budget_policy_id"`
	// Name of the vector search endpoint
	EndpointName string `json:"-" tf:"-"`
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

type PatchEndpointBudgetPolicyResponse struct {
	// The budget policy applied to the vector search endpoint.
	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// Request payload for getting next page of results.
type QueryVectorIndexNextPageRequest struct {
	// Name of the endpoint.
	// Wire name: 'endpoint_name'
	EndpointName string `json:"endpoint_name,omitempty"`
	// Name of the vector index to query.
	IndexName string `json:"-" tf:"-"`
	// Page token returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` API.
	// Wire name: 'page_token'
	PageToken string `json:"page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type QueryVectorIndexRequest struct {
	// List of column names to include in the response.
	// Wire name: 'columns'
	Columns []string `json:"columns"`
	// Column names used to retrieve data to send to the reranker.
	// Wire name: 'columns_to_rerank'
	ColumnsToRerank []string `json:"columns_to_rerank,omitempty"`
	// JSON string representing query filters.
	//
	// Example filters:
	//
	// - `{"id <": 5}`: Filter for id less than 5. - `{"id >": 5}`: Filter for
	// id greater than 5. - `{"id <=": 5}`: Filter for id less than equal to 5.
	// - `{"id >=": 5}`: Filter for id greater than equal to 5. - `{"id": 5}`:
	// Filter for id equal to 5.
	// Wire name: 'filters_json'
	FiltersJson string `json:"filters_json,omitempty"`
	// Name of the vector index to query.
	IndexName string `json:"-" tf:"-"`
	// Number of results to return. Defaults to 10.
	// Wire name: 'num_results'
	NumResults int `json:"num_results,omitempty"`
	// Query text. Required for Delta Sync Index using model endpoint.
	// Wire name: 'query_text'
	QueryText string `json:"query_text,omitempty"`
	// The query type to use. Choices are `ANN` and `HYBRID`. Defaults to `ANN`.
	// Wire name: 'query_type'
	QueryType string `json:"query_type,omitempty"`
	// Query vector. Required for Direct Vector Access Index and Delta Sync
	// Index using self-managed vectors.
	// Wire name: 'query_vector'
	QueryVector []float64 `json:"query_vector,omitempty"`

	// Wire name: 'reranker'
	Reranker *RerankerConfig `json:"reranker,omitempty"`
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	// Wire name: 'score_threshold'
	ScoreThreshold float64 `json:"score_threshold,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type QueryVectorIndexResponse struct {
	// Metadata about the result set.
	// Wire name: 'manifest'
	Manifest *ResultManifest `json:"manifest,omitempty"`
	// [Optional] Token that can be used in `QueryVectorIndexNextPage` API to
	// get next page of results. If more than 1000 results satisfy the query,
	// they are returned in groups of 1000. Empty value means no more results.
	// The maximum number of results that can be returned is 10,000.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// Data returned in the query result.
	// Wire name: 'result'
	Result *ResultData `json:"result,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type RerankerConfig struct {

	// Wire name: 'model'
	Model string `json:"model,omitempty"`

	// Wire name: 'parameters'
	Parameters *RerankerConfigRerankerParameters `json:"parameters,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RerankerConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &rerankerConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := rerankerConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RerankerConfig) MarshalJSON() ([]byte, error) {
	pb, err := rerankerConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RerankerConfigRerankerParameters struct {

	// Wire name: 'columns_to_rerank'
	ColumnsToRerank []string `json:"columns_to_rerank,omitempty"`
}

func (st *RerankerConfigRerankerParameters) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &rerankerConfigRerankerParametersPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := rerankerConfigRerankerParametersFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RerankerConfigRerankerParameters) MarshalJSON() ([]byte, error) {
	pb, err := rerankerConfigRerankerParametersToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Data returned in the query result.
type ResultData struct {
	// Data rows returned in the query.
	// Wire name: 'data_array'
	DataArray [][]string `json:"data_array,omitempty"`
	// Number of rows in the result set.
	// Wire name: 'row_count'
	RowCount int `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// Metadata about the result set.
type ResultManifest struct {
	// Number of columns in the result set.
	// Wire name: 'column_count'
	ColumnCount int `json:"column_count,omitempty"`
	// Information about each column in the result set.
	// Wire name: 'columns'
	Columns []ColumnInfo `json:"columns,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type ScanVectorIndexRequest struct {
	// Name of the vector index to scan.
	IndexName string `json:"-" tf:"-"`
	// Primary key of the last entry returned in the previous scan.
	// Wire name: 'last_primary_key'
	LastPrimaryKey string `json:"last_primary_key,omitempty"`
	// Number of results to return. Defaults to 10.
	// Wire name: 'num_results'
	NumResults int `json:"num_results,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// Response to a scan vector index request.
type ScanVectorIndexResponse struct {
	// List of data entries
	// Wire name: 'data'
	Data []Struct `json:"data,omitempty"`
	// Primary key of the last entry.
	// Wire name: 'last_primary_key'
	LastPrimaryKey string `json:"last_primary_key,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type Struct struct {
	// Data entry, corresponding to a row in a vector index.
	// Wire name: 'fields'
	Fields []MapStringValueEntry `json:"fields,omitempty"`
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

// Synchronize an index
type SyncIndexRequest struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	IndexName string `json:"-" tf:"-"`
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

type SyncIndexResponse struct {
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

type UpdateEndpointCustomTagsRequest struct {
	// The new custom tags for the vector search endpoint
	// Wire name: 'custom_tags'
	CustomTags []CustomTag `json:"custom_tags"`
	// Name of the vector search endpoint
	EndpointName string `json:"-" tf:"-"`
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

type UpdateEndpointCustomTagsResponse struct {
	// All the custom tags that are applied to the vector search endpoint.
	// Wire name: 'custom_tags'
	CustomTags []CustomTag `json:"custom_tags,omitempty"`
	// The name of the vector search endpoint whose custom tags were updated.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type UpsertDataResult struct {
	// List of primary keys for rows that failed to process.
	// Wire name: 'failed_primary_keys'
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	// Count of successfully processed rows.
	// Wire name: 'success_row_count'
	SuccessRowCount int64 `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type UpsertDataVectorIndexRequest struct {
	// Name of the vector index where data is to be upserted. Must be a Direct
	// Vector Access Index.
	IndexName string `json:"-" tf:"-"`
	// JSON string representing the data to be upserted.
	// Wire name: 'inputs_json'
	InputsJson string `json:"inputs_json"`
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

type UpsertDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	// Wire name: 'result'
	Result *UpsertDataResult `json:"result,omitempty"`
	// Status of the upsert operation.
	// Wire name: 'status'
	Status UpsertDataStatus `json:"status,omitempty"`
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

type Value struct {

	// Wire name: 'bool_value'
	BoolValue bool `json:"bool_value,omitempty"`

	// Wire name: 'list_value'
	ListValue *ListValue `json:"list_value,omitempty"`

	// Wire name: 'number_value'
	NumberValue float64 `json:"number_value,omitempty"`

	// Wire name: 'string_value'
	StringValue string `json:"string_value,omitempty"`

	// Wire name: 'struct_value'
	StructValue *Struct `json:"struct_value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type VectorIndex struct {
	// The user who created the index.
	// Wire name: 'creator'
	Creator string `json:"creator,omitempty"`

	// Wire name: 'delta_sync_index_spec'
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecResponse `json:"delta_sync_index_spec,omitempty"`

	// Wire name: 'direct_access_index_spec'
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec `json:"direct_access_index_spec,omitempty"`
	// Name of the endpoint associated with the index
	// Wire name: 'endpoint_name'
	EndpointName string `json:"endpoint_name,omitempty"`
	// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
	// automatically syncs with a source Delta Table, automatically and
	// incrementally updating the index as the underlying data in the Delta
	// Table changes. - `DIRECT_ACCESS`: An index that supports direct read and
	// write of vectors and metadata through our REST and SDK APIs. With this
	// model, the user manages index updates.
	// Wire name: 'index_type'
	IndexType VectorIndexType `json:"index_type,omitempty"`
	// Name of the index
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Primary key of the index
	// Wire name: 'primary_key'
	PrimaryKey string `json:"primary_key,omitempty"`

	// Wire name: 'status'
	Status *VectorIndexStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type VectorIndexStatus struct {
	// Index API Url to be used to perform operations on the index
	// Wire name: 'index_url'
	IndexUrl string `json:"index_url,omitempty"`
	// Number of rows indexed
	// Wire name: 'indexed_row_count'
	IndexedRowCount int64 `json:"indexed_row_count,omitempty"`
	// Message associated with the index status
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// Whether the index is ready for search
	// Wire name: 'ready'
	Ready bool `json:"ready,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

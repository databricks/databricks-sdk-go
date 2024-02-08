// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// all definitions in this file are in alphabetical order

type ColumnInfo struct {
	// Name of the column.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateEndpoint struct {
	// Type of endpoint.
	EndpointType EndpointType `json:"endpoint_type"`
	// Name of endpoint
	Name string `json:"name"`
}

type CreateVectorIndexRequest struct {
	// Specification for Delta Sync Index. Required if `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecRequest `json:"delta_sync_index_spec,omitempty"`
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec `json:"direct_access_index_spec,omitempty"`
	// Name of the endpoint to be used for serving the index
	EndpointName string `json:"endpoint_name,omitempty"`
	// There are 2 types of Vector Search indexes:
	//
	// - `DELTA_SYNC`: An index that automatically syncs with a source Delta
	// Table, automatically and incrementally updating the index as the
	// underlying data in the Delta Table changes. - `DIRECT_ACCESS`: An index
	// that supports direct read and write of vectors and metadata through our
	// REST and SDK APIs. With this model, the user manages index updates.
	IndexType VectorIndexType `json:"index_type"`
	// Name of the index
	Name string `json:"name"`
	// Primary key of the index
	PrimaryKey string `json:"primary_key"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateVectorIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateVectorIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateVectorIndexResponse struct {
	VectorIndex *VectorIndex `json:"vector_index,omitempty"`
}

// Result of the upsert or delete operation.
type DeleteDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	// Count of successfully processed rows.
	SuccessRowCount int64 `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteDataResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDataResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Status of the delete operation.
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

// Type always returns DeleteDataStatus to satisfy [pflag.Value] interface
func (f *DeleteDataStatus) Type() string {
	return "DeleteDataStatus"
}

// Request payload for deleting data from a vector index.
type DeleteDataVectorIndexRequest struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	Name string `json:"-" url:"-"`
	// List of primary keys for the data to be deleted.
	PrimaryKeys []string `json:"primary_keys"`
}

// Response to a delete data vector index request.
type DeleteDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result *DeleteDataResult `json:"result,omitempty"`
	// Status of the delete operation.
	Status DeleteDataStatus `json:"status,omitempty"`
}

// Delete an endpoint
type DeleteEndpointRequest struct {
	// Name of the endpoint
	EndpointName string `json:"-" url:"-"`
	// Name of the endpoint to delete
	Name string `json:"-"`
}

// Delete an index
type DeleteIndexRequest struct {
	// Name of the index
	IndexName string `json:"-" url:"-"`
}

type DeltaSyncVectorIndexSpecRequest struct {
	// The columns that contain the embedding source.
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	// Pipeline execution mode.
	//
	// - `TRIGGERED`: If the pipeline uses the triggered execution mode, the
	// system stops processing after successfully refreshing the source table in
	// the pipeline once, ensuring the table is updated based on the data
	// available when the update started. - `CONTINUOUS`: If the pipeline uses
	// continuous execution, the pipeline processes new data as it arrives in
	// the source table to keep vector index fresh.
	PipelineType PipelineType `json:"pipeline_type,omitempty"`
	// The name of the source table.
	SourceTable string `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DeltaSyncVectorIndexSpecRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSyncVectorIndexSpecRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeltaSyncVectorIndexSpecResponse struct {
	// The columns that contain the embedding source.
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	// The ID of the pipeline that is used to sync the index.
	PipelineId string `json:"pipeline_id,omitempty"`
	// Pipeline execution mode.
	//
	// - `TRIGGERED`: If the pipeline uses the triggered execution mode, the
	// system stops processing after successfully refreshing the source table in
	// the pipeline once, ensuring the table is updated based on the data
	// available when the update started. - `CONTINUOUS`: If the pipeline uses
	// continuous execution, the pipeline processes new data as it arrives in
	// the source table to keep vector index fresh.
	PipelineType PipelineType `json:"pipeline_type,omitempty"`
	// The name of the source table.
	SourceTable string `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DeltaSyncVectorIndexSpecResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSyncVectorIndexSpecResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DirectAccessVectorIndexSpec struct {
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	// The schema of the index in JSON format.
	//
	// Supported types are `integer`, `long`, `float`, `double`, `boolean`,
	// `string`, `date`, `timestamp`.
	//
	// Supported types for vector column: `array<float>`, `array<double>`,`.
	SchemaJson string `json:"schema_json,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DirectAccessVectorIndexSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DirectAccessVectorIndexSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EmbeddingConfig struct {
	// Name of the embedding model endpoint
	EmbeddingModelEndpointName string `json:"embedding_model_endpoint_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EmbeddingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EmbeddingSourceColumn struct {
	EmbeddingConfig *EmbeddingConfig `json:"embedding_config,omitempty"`
	// Name of the column
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EmbeddingSourceColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingSourceColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EmbeddingVectorColumn struct {
	// Dimension of the embedding vector
	EmbeddingDimension int `json:"embedding_dimension,omitempty"`
	// Name of the column
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EmbeddingVectorColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingVectorColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointInfo struct {
	// Timestamp of endpoint creation
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Creator of the endpoint
	Creator string `json:"creator,omitempty"`
	// Current status of the endpoint
	EndpointStatus *EndpointStatus `json:"endpoint_status,omitempty"`
	// Type of endpoint.
	EndpointType EndpointType `json:"endpoint_type,omitempty"`
	// Unique identifier of the endpoint
	Id string `json:"id,omitempty"`
	// Timestamp of last update to the endpoint
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// User who last updated the endpoint
	LastUpdatedUser string `json:"last_updated_user,omitempty"`
	// Name of endpoint
	Name string `json:"name,omitempty"`
	// Number of indexes on the endpoint
	NumIndexes int `json:"num_indexes,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EndpointInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Status information of an endpoint
type EndpointStatus struct {
	// Additional status message
	Message string `json:"message,omitempty"`
	// Current state of the endpoint
	State EndpointStatusState `json:"state,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EndpointStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Type always returns EndpointType to satisfy [pflag.Value] interface
func (f *EndpointType) Type() string {
	return "EndpointType"
}

// Get an endpoint
type GetEndpointRequest struct {
	// Name of the endpoint
	EndpointName string `json:"-" url:"-"`
}

// Get an index
type GetIndexRequest struct {
	// Name of the index
	IndexName string `json:"-" url:"-"`
}

type ListEndpointResponse struct {
	// An array of Endpoint objects
	Endpoints []EndpointInfo `json:"endpoints,omitempty"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListEndpointResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List all endpoints
type ListEndpointsRequest struct {
	// Token for pagination
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListEndpointsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List indexes
type ListIndexesRequest struct {
	// Name of the endpoint
	EndpointName string `json:"-" url:"endpoint_name"`
	// Token for pagination
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListIndexesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListIndexesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListVectorIndexesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	VectorIndexes []MiniVectorIndex `json:"vector_indexes,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListVectorIndexesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVectorIndexesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MiniVectorIndex struct {
	// The user who created the index.
	Creator string `json:"creator,omitempty"`
	// Name of the endpoint associated with the index
	EndpointName string `json:"endpoint_name,omitempty"`
	// There are 2 types of Vector Search indexes:
	//
	// - `DELTA_SYNC`: An index that automatically syncs with a source Delta
	// Table, automatically and incrementally updating the index as the
	// underlying data in the Delta Table changes. - `DIRECT_ACCESS`: An index
	// that supports direct read and write of vectors and metadata through our
	// REST and SDK APIs. With this model, the user manages index updates.
	IndexType VectorIndexType `json:"index_type,omitempty"`
	// Name of the index
	Name string `json:"name,omitempty"`
	// Primary key of the index
	PrimaryKey string `json:"primary_key,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *MiniVectorIndex) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MiniVectorIndex) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Pipeline execution mode.
//
// - `TRIGGERED`: If the pipeline uses the triggered execution mode, the system
// stops processing after successfully refreshing the source table in the
// pipeline once, ensuring the table is updated based on the data available when
// the update started. - `CONTINUOUS`: If the pipeline uses continuous
// execution, the pipeline processes new data as it arrives in the source table
// to keep vector index fresh.
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

// Type always returns PipelineType to satisfy [pflag.Value] interface
func (f *PipelineType) Type() string {
	return "PipelineType"
}

type QueryVectorIndexRequest struct {
	// List of column names to include in the response.
	Columns []string `json:"columns"`
	// JSON string representing query filters.
	//
	// Example filters: - `{"id <": 5}`: Filter for id less than 5. - `{"id >":
	// 5}`: Filter for id greater than 5. - `{"id <=": 5}`: Filter for id less
	// than equal to 5. - `{"id >=": 5}`: Filter for id greater than equal to 5.
	// - `{"id": 5}`: Filter for id equal to 5.
	FiltersJson string `json:"filters_json,omitempty"`
	// Name of the vector index to query.
	IndexName string `json:"-" url:"-"`
	// Number of results to return. Defaults to 10.
	NumResults int `json:"num_results,omitempty"`
	// Query text. Required for Delta Sync Index using model endpoint.
	QueryText string `json:"query_text,omitempty"`
	// Query vector. Required for Direct Vector Access Index and Delta Sync
	// Index using self-managed vectors.
	QueryVector []float64 `json:"query_vector,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *QueryVectorIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryVectorIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryVectorIndexResponse struct {
	// Metadata about the result set.
	Manifest *ResultManifest `json:"manifest,omitempty"`
	// Data returned in the query result.
	Result *ResultData `json:"result,omitempty"`
}

// Data returned in the query result.
type ResultData struct {
	// Data rows returned in the query.
	DataArray [][]string `json:"data_array,omitempty"`
	// Number of rows in the result set.
	RowCount int `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ResultData) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultData) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metadata about the result set.
type ResultManifest struct {
	// Number of columns in the result set.
	ColumnCount int `json:"column_count,omitempty"`
	// Information about each column in the result set.
	Columns []ColumnInfo `json:"columns,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ResultManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultManifest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Synchronize an index
type SyncIndexRequest struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	IndexName string `json:"-" url:"-"`
}

// Result of the upsert or delete operation.
type UpsertDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	// Count of successfully processed rows.
	SuccessRowCount int64 `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *UpsertDataResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpsertDataResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Status of the upsert operation.
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

// Type always returns UpsertDataStatus to satisfy [pflag.Value] interface
func (f *UpsertDataStatus) Type() string {
	return "UpsertDataStatus"
}

// Request payload for upserting data into a vector index.
type UpsertDataVectorIndexRequest struct {
	// JSON string representing the data to be upserted.
	InputsJson string `json:"inputs_json"`
	// Name of the vector index where data is to be upserted. Must be a Direct
	// Vector Access Index.
	Name string `json:"-" url:"-"`
}

// Response to an upsert data vector index request.
type UpsertDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result *UpsertDataResult `json:"result,omitempty"`
	// Status of the upsert operation.
	Status UpsertDataStatus `json:"status,omitempty"`
}

type VectorIndex struct {
	// The user who created the index.
	Creator string `json:"creator,omitempty"`

	DeltaSyncVectorIndexSpec *DeltaSyncVectorIndexSpecResponse `json:"delta_sync_vector_index_spec,omitempty"`

	DirectAccessVectorIndexSpec *DirectAccessVectorIndexSpec `json:"direct_access_vector_index_spec,omitempty"`
	// Name of the endpoint associated with the index
	EndpointName string `json:"endpoint_name,omitempty"`
	// There are 2 types of Vector Search indexes:
	//
	// - `DELTA_SYNC`: An index that automatically syncs with a source Delta
	// Table, automatically and incrementally updating the index as the
	// underlying data in the Delta Table changes. - `DIRECT_ACCESS`: An index
	// that supports direct read and write of vectors and metadata through our
	// REST and SDK APIs. With this model, the user manages index updates.
	IndexType VectorIndexType `json:"index_type,omitempty"`
	// Name of the index
	Name string `json:"name,omitempty"`
	// Primary key of the index
	PrimaryKey string `json:"primary_key,omitempty"`

	Status *VectorIndexStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *VectorIndex) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorIndex) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VectorIndexStatus struct {
	// Index API Url to be used to perform operations on the index
	IndexUrl string `json:"index_url,omitempty"`
	// Number of rows indexed
	IndexedRowCount int64 `json:"indexed_row_count,omitempty"`
	// Message associated with the index status
	Message string `json:"message,omitempty"`
	// Whether the index is ready for search
	Ready bool `json:"ready,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *VectorIndexStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorIndexStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// There are 2 types of Vector Search indexes:
//
// - `DELTA_SYNC`: An index that automatically syncs with a source Delta Table,
// automatically and incrementally updating the index as the underlying data in
// the Delta Table changes. - `DIRECT_ACCESS`: An index that supports direct
// read and write of vectors and metadata through our REST and SDK APIs. With
// this model, the user manages index updates.
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

// Type always returns VectorIndexType to satisfy [pflag.Value] interface
func (f *VectorIndexType) Type() string {
	return "VectorIndexType"
}

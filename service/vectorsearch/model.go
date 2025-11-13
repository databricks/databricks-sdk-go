// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// Adjusted throughput request parameters
type AdjustedThroughputRequest struct {
	// Adjusted concurrency (total CPU) for the endpoint
	Concurrency float64 `json:"concurrency,omitempty"`
	// Adjusted maximum concurrency allowed for the endpoint
	MaximumConcurrencyAllowed float64 `json:"maximum_concurrency_allowed,omitempty"`
	// Adjusted minimum concurrency allowed for the endpoint
	MinimalConcurrencyAllowed float64 `json:"minimal_concurrency_allowed,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AdjustedThroughputRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AdjustedThroughputRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ColumnInfo struct {
	// Name of the column.
	Name string `json:"name,omitempty"`
	// Data type of the column (e.g., "string", "int", "array<float>")
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateEndpoint struct {
	// The budget policy id to be applied
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// Type of endpoint
	EndpointType EndpointType `json:"endpoint_type"`
	// Name of the vector search endpoint
	Name string `json:"name"`
	// Initial number of replicas for the endpoint. If not specified, defaults
	// to 1.
	NumReplicas int `json:"num_replicas,omitempty"`
	// The usage policy id to be applied once we've migrated to usage policies
	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateVectorIndexRequest struct {
	// Specification for Delta Sync Index. Required if `index_type` is
	// `DELTA_SYNC`.
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecRequest `json:"delta_sync_index_spec,omitempty"`
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec `json:"direct_access_index_spec,omitempty"`
	// Name of the endpoint to be used for serving the index
	EndpointName string `json:"endpoint_name"`

	IndexType VectorIndexType `json:"index_type"`
	// Name of the index
	Name string `json:"name"`
	// Primary key of the index
	PrimaryKey string `json:"primary_key"`
}

type CustomTag struct {
	// Key field for a vector search endpoint tag.
	Key string `json:"key"`
	// [Optional] Value field for a vector search endpoint tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	// Count of successfully processed rows.
	SuccessRowCount int64 `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDataResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDataResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

type DeleteDataVectorIndexRequest struct {
	// Name of the vector index where data is to be deleted. Must be a Direct
	// Vector Access Index.
	IndexName string `json:"-" url:"-"`
	// List of primary keys for the data to be deleted.
	PrimaryKeys []string `json:"-" url:"primary_keys"`
}

type DeleteDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result *DeleteDataResult `json:"result,omitempty"`
	// Status of the delete operation.
	Status DeleteDataStatus `json:"status,omitempty"`
}

type DeleteEndpointRequest struct {
	// Name of the vector search endpoint
	EndpointName string `json:"-" url:"-"`
}

type DeleteIndexRequest struct {
	// Name of the index
	IndexName string `json:"-" url:"-"`
}

type DeltaSyncVectorIndexSpecRequest struct {
	// [Optional] Select the columns to sync with the vector index. If you leave
	// this field blank, all columns from the source table are synced with the
	// index. The primary key column and embedding source column or embedding
	// vector column are always synced.
	ColumnsToSync []string `json:"columns_to_sync,omitempty"`
	// The budget policy id applied to the vector search index
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`

	EffectiveUsagePolicyId string `json:"effective_usage_policy_id,omitempty"`
	// The columns that contain the embedding source.
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
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

func (s *DeltaSyncVectorIndexSpecRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSyncVectorIndexSpecRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeltaSyncVectorIndexSpecResponse struct {
	// The budget policy id applied to the vector search index
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`

	EffectiveUsagePolicyId string `json:"effective_usage_policy_id,omitempty"`
	// The columns that contain the embedding source.
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
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

func (s *DeltaSyncVectorIndexSpecResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSyncVectorIndexSpecResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DirectAccessVectorIndexSpec struct {
	// The columns that contain the embedding source. The format should be
	// array[double].
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	// The columns that contain the embedding vectors. The format should be
	// array[double].
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	// The schema of the index in JSON format. Supported types are `integer`,
	// `long`, `float`, `double`, `boolean`, `string`, `date`, `timestamp`.
	// Supported types for vector column: `array<float>`, `array<double>`,`.
	SchemaJson string `json:"schema_json,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DirectAccessVectorIndexSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DirectAccessVectorIndexSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EmbeddingSourceColumn struct {
	// Name of the embedding model endpoint, used by default for both ingestion
	// and querying.
	EmbeddingModelEndpointName string `json:"embedding_model_endpoint_name,omitempty"`
	// Name of the embedding model endpoint which, if specified, is used for
	// querying (not ingestion).
	ModelEndpointNameForQuery string `json:"model_endpoint_name_for_query,omitempty"`
	// Name of the column
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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

	ForceSendFields []string `json:"-" url:"-"`
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
	// The custom tags assigned to the endpoint
	CustomTags []CustomTag `json:"custom_tags,omitempty"`
	// The budget policy id applied to the endpoint
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// Current status of the endpoint
	EndpointStatus *EndpointStatus `json:"endpoint_status,omitempty"`
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
	// Throughput information for the endpoint
	ThroughputInfo *EndpointThroughputInfo `json:"throughput_info,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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

	ForceSendFields []string `json:"-" url:"-"`
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

const EndpointStatusStateRedState EndpointStatusState = `RED_STATE`

const EndpointStatusStateYellowState EndpointStatusState = `YELLOW_STATE`

// String representation for [fmt.Print]
func (f *EndpointStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStatusState) Set(v string) error {
	switch v {
	case `OFFLINE`, `ONLINE`, `PROVISIONING`, `RED_STATE`, `YELLOW_STATE`:
		*f = EndpointStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OFFLINE", "ONLINE", "PROVISIONING", "RED_STATE", "YELLOW_STATE"`, v)
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
		EndpointStatusStateRedState,
		EndpointStatusStateYellowState,
	}
}

// Type always returns EndpointStatusState to satisfy [pflag.Value] interface
func (f *EndpointStatusState) Type() string {
	return "EndpointStatusState"
}

// Throughput information for an endpoint
type EndpointThroughputInfo struct {
	// Additional information about the throughput change request
	ChangeRequestMessage string `json:"change_request_message,omitempty"`
	// The state of the most recent throughput change request
	ChangeRequestState ThroughputChangeRequestState `json:"change_request_state,omitempty"`
	// The current concurrency (total CPU) allocated to the endpoint
	CurrentConcurrency float64 `json:"current_concurrency,omitempty"`
	// The current utilization of concurrency as a percentage (0-100)
	CurrentConcurrencyUtilizationPercentage float64 `json:"current_concurrency_utilization_percentage,omitempty"`
	// The current number of replicas allocated to the endpoint
	CurrentNumReplicas int `json:"current_num_replicas,omitempty"`
	// The maximum concurrency allowed for this endpoint
	MaximumConcurrencyAllowed float64 `json:"maximum_concurrency_allowed,omitempty"`
	// The minimum concurrency allowed for this endpoint
	MinimalConcurrencyAllowed float64 `json:"minimal_concurrency_allowed,omitempty"`
	// The requested concurrency (total CPU) for the endpoint
	RequestedConcurrency float64 `json:"requested_concurrency,omitempty"`
	// The requested number of replicas for the endpoint
	RequestedNumReplicas int `json:"requested_num_replicas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointThroughputInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointThroughputInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

type GetEndpointRequest struct {
	// Name of the endpoint
	EndpointName string `json:"-" url:"-"`
}

type GetIndexRequest struct {
	// If true, the URL returned for the index is guaranteed to be compatible
	// with the reranker. Currently this means we return the CP URL regardless
	// of how the index is being accessed. If not set or set to false, the URL
	// may still be compatible with the reranker depending on what URL we
	// return.
	EnsureRerankerCompatible bool `json:"-" url:"ensure_reranker_compatible,omitempty"`
	// Name of the index
	IndexName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEndpointResponse struct {
	// An array of Endpoint objects
	Endpoints []EndpointInfo `json:"endpoints,omitempty"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEndpointResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEndpointsRequest struct {
	// Token for pagination
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListEndpointsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListEndpointsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListIndexesRequest struct {
	// Name of the endpoint
	EndpointName string `json:"-" url:"endpoint_name"`
	// Token for pagination
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListIndexesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListIndexesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListValue struct {
	// Repeated field of dynamically typed values.
	Values []Value `json:"values,omitempty"`
}

type ListVectorIndexesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	VectorIndexes []MiniVectorIndex `json:"vector_indexes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListVectorIndexesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVectorIndexesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Key-value pair.
type MapStringValueEntry struct {
	// Column name.
	Key string `json:"key,omitempty"`
	// Column value, nullable.
	Value *Value `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MapStringValueEntry) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MapStringValueEntry) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Metric specification
type Metric struct {
	// Metric labels
	Labels []MetricLabel `json:"labels,omitempty"`
	// Metric name
	Name string `json:"name,omitempty"`
	// Percentile for the metric
	Percentile float64 `json:"percentile,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Metric) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Metric) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Label for a metric
type MetricLabel struct {
	// Label name
	Name string `json:"name,omitempty"`
	// Label value
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MetricLabel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetricLabel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Single metric value at a specific timestamp
type MetricValue struct {
	// Timestamp of the metric value (milliseconds since epoch)
	Timestamp int64 `json:"timestamp,omitempty"`
	// Metric value
	Value float64 `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MetricValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MetricValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Collection of metric values for a specific metric
type MetricValues struct {
	// Metric specification
	Metric *Metric `json:"metric,omitempty"`
	// Time series of metric values
	Values []MetricValue `json:"values,omitempty"`
}

type MiniVectorIndex struct {
	// The user who created the index.
	Creator string `json:"creator,omitempty"`
	// Name of the endpoint associated with the index
	EndpointName string `json:"endpoint_name,omitempty"`

	IndexType VectorIndexType `json:"index_type,omitempty"`
	// Name of the index
	Name string `json:"name,omitempty"`
	// Primary key of the index
	PrimaryKey string `json:"primary_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MiniVectorIndex) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MiniVectorIndex) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PatchEndpointBudgetPolicyRequest struct {
	// The budget policy id to be applied (hima-sheth) TODO: remove this once
	// we've migrated to usage policies
	BudgetPolicyId string `json:"budget_policy_id"`
	// Name of the vector search endpoint
	EndpointName string `json:"-" url:"-"`
}

type PatchEndpointBudgetPolicyResponse struct {
	// The budget policy applied to the vector search endpoint.
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PatchEndpointBudgetPolicyResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PatchEndpointBudgetPolicyResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PatchEndpointThroughputRequest struct {
	// If true, the request will fail if the requested concurrency or limits
	// cannot be exactly met. If false, the request will be adjusted to the
	// closest possible value.
	AllOrNothing bool `json:"all_or_nothing,omitempty"`
	// Requested concurrency (total CPU) for the endpoint. If not specified, the
	// current concurrency is maintained.
	Concurrency float64 `json:"concurrency,omitempty"`
	// Name of the vector search endpoint
	EndpointName string `json:"-" url:"-"`
	// Maximum concurrency allowed for the endpoint. If not specified, the
	// current maximum is maintained.
	MaximumConcurrencyAllowed float64 `json:"maximum_concurrency_allowed,omitempty"`
	// Minimum concurrency allowed for the endpoint. If not specified, the
	// current minimum is maintained.
	MinimalConcurrencyAllowed float64 `json:"minimal_concurrency_allowed,omitempty"`
	// Requested number of data copies for the endpoint (including primary). For
	// example: num_replicas=2 means 2 total copies of the data (1 primary + 1
	// replica). If not specified, the current replication factor is maintained.
	// Valid range: 1-6 (where 1 = no replication, 6 = 1 primary + 5 replicas).
	NumReplicas int `json:"num_replicas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PatchEndpointThroughputRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PatchEndpointThroughputRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PatchEndpointThroughputResponse struct {
	// The adjusted request if the original request could not be fully
	// fulfilled. This is only populated when the request was adjusted.
	AdjustedRequest *AdjustedThroughputRequest `json:"adjusted_request,omitempty"`
	// Message explaining the status or any adjustments made
	Message string `json:"message,omitempty"`
	// The status of the throughput change request
	Status ThroughputPatchStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PatchEndpointThroughputResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PatchEndpointThroughputResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	EndpointName string `json:"endpoint_name,omitempty"`
	// Name of the vector index to query.
	IndexName string `json:"-" url:"-"`
	// Page token returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` API.
	PageToken string `json:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryVectorIndexNextPageRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryVectorIndexNextPageRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryVectorIndexRequest struct {
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
	// The query type to use. Choices are `ANN` and `HYBRID` and `FULL_TEXT`.
	// Defaults to `ANN`.
	QueryType string `json:"query_type,omitempty"`
	// Query vector. Required for Direct Vector Access Index and Delta Sync
	// Index using self-managed vectors.
	QueryVector []float64 `json:"query_vector,omitempty"`

	Reranker *RerankerConfig `json:"reranker,omitempty"`
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	ScoreThreshold float64 `json:"score_threshold,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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
	// [Optional] Token that can be used in `QueryVectorIndexNextPage` API to
	// get next page of results. If more than 1000 results satisfy the query,
	// they are returned in groups of 1000. Empty value means no more results.
	// The maximum number of results that can be returned is 10,000.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Data returned in the query result.
	Result *ResultData `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryVectorIndexResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryVectorIndexResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RerankerConfig struct {
	Model string `json:"model,omitempty"`

	Parameters *RerankerConfigRerankerParameters `json:"parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RerankerConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RerankerConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RerankerConfigRerankerParameters struct {
	ColumnsToRerank []string `json:"columns_to_rerank,omitempty"`
}

// Data returned in the query result.
type ResultData struct {
	// Data rows returned in the query.
	DataArray [][]string `json:"data_array,omitempty"`
	// Number of rows in the result set.
	RowCount int `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResultManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultManifest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request to retrieve user-visible metrics
type RetrieveUserVisibleMetricsRequest struct {
	// End time for metrics query
	EndTime string `json:"end_time,omitempty"`
	// Granularity in seconds
	GranularityInSeconds int `json:"granularity_in_seconds,omitempty"`
	// List of metrics to retrieve
	Metrics []Metric `json:"metrics,omitempty"`
	// Vector search endpoint name
	Name string `json:"-" url:"-"`
	// Token for pagination
	PageToken string `json:"page_token,omitempty"`
	// Start time for metrics query
	StartTime string `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RetrieveUserVisibleMetricsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RetrieveUserVisibleMetricsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response containing user-visible metrics
type RetrieveUserVisibleMetricsResponse struct {
	// Collection of metric values
	MetricValues []MetricValues `json:"metric_values,omitempty"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RetrieveUserVisibleMetricsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RetrieveUserVisibleMetricsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ScanVectorIndexRequest struct {
	// Name of the vector index to scan.
	IndexName string `json:"-" url:"-"`
	// Primary key of the last entry returned in the previous scan.
	LastPrimaryKey string `json:"last_primary_key,omitempty"`
	// Number of results to return. Defaults to 10.
	NumResults int `json:"num_results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ScanVectorIndexRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ScanVectorIndexRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response to a scan vector index request.
type ScanVectorIndexResponse struct {
	// List of data entries
	Data []Struct `json:"data,omitempty"`
	// Primary key of the last entry.
	LastPrimaryKey string `json:"last_primary_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ScanVectorIndexResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ScanVectorIndexResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Struct struct {
	// Data entry, corresponding to a row in a vector index.
	Fields []MapStringValueEntry `json:"fields,omitempty"`
}

type SyncIndexRequest struct {
	// Name of the vector index to synchronize. Must be a Delta Sync Index.
	IndexName string `json:"-" url:"-"`
}

// Throughput change request state
type ThroughputChangeRequestState string

const ThroughputChangeRequestStateChangeAdjusted ThroughputChangeRequestState = `CHANGE_ADJUSTED`

const ThroughputChangeRequestStateChangeFailed ThroughputChangeRequestState = `CHANGE_FAILED`

const ThroughputChangeRequestStateChangeInProgress ThroughputChangeRequestState = `CHANGE_IN_PROGRESS`

const ThroughputChangeRequestStateChangeReachedMaximum ThroughputChangeRequestState = `CHANGE_REACHED_MAXIMUM`

const ThroughputChangeRequestStateChangeReachedMinimum ThroughputChangeRequestState = `CHANGE_REACHED_MINIMUM`

const ThroughputChangeRequestStateChangeSuccess ThroughputChangeRequestState = `CHANGE_SUCCESS`

// String representation for [fmt.Print]
func (f *ThroughputChangeRequestState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ThroughputChangeRequestState) Set(v string) error {
	switch v {
	case `CHANGE_ADJUSTED`, `CHANGE_FAILED`, `CHANGE_IN_PROGRESS`, `CHANGE_REACHED_MAXIMUM`, `CHANGE_REACHED_MINIMUM`, `CHANGE_SUCCESS`:
		*f = ThroughputChangeRequestState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CHANGE_ADJUSTED", "CHANGE_FAILED", "CHANGE_IN_PROGRESS", "CHANGE_REACHED_MAXIMUM", "CHANGE_REACHED_MINIMUM", "CHANGE_SUCCESS"`, v)
	}
}

// Values returns all possible values for ThroughputChangeRequestState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ThroughputChangeRequestState) Values() []ThroughputChangeRequestState {
	return []ThroughputChangeRequestState{
		ThroughputChangeRequestStateChangeAdjusted,
		ThroughputChangeRequestStateChangeFailed,
		ThroughputChangeRequestStateChangeInProgress,
		ThroughputChangeRequestStateChangeReachedMaximum,
		ThroughputChangeRequestStateChangeReachedMinimum,
		ThroughputChangeRequestStateChangeSuccess,
	}
}

// Type always returns ThroughputChangeRequestState to satisfy [pflag.Value] interface
func (f *ThroughputChangeRequestState) Type() string {
	return "ThroughputChangeRequestState"
}

// Response status for throughput change requests
type ThroughputPatchStatus string

const ThroughputPatchStatusPatchAccepted ThroughputPatchStatus = `PATCH_ACCEPTED`

const ThroughputPatchStatusPatchFailed ThroughputPatchStatus = `PATCH_FAILED`

const ThroughputPatchStatusPatchRejected ThroughputPatchStatus = `PATCH_REJECTED`

// String representation for [fmt.Print]
func (f *ThroughputPatchStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ThroughputPatchStatus) Set(v string) error {
	switch v {
	case `PATCH_ACCEPTED`, `PATCH_FAILED`, `PATCH_REJECTED`:
		*f = ThroughputPatchStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PATCH_ACCEPTED", "PATCH_FAILED", "PATCH_REJECTED"`, v)
	}
}

// Values returns all possible values for ThroughputPatchStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *ThroughputPatchStatus) Values() []ThroughputPatchStatus {
	return []ThroughputPatchStatus{
		ThroughputPatchStatusPatchAccepted,
		ThroughputPatchStatusPatchFailed,
		ThroughputPatchStatusPatchRejected,
	}
}

// Type always returns ThroughputPatchStatus to satisfy [pflag.Value] interface
func (f *ThroughputPatchStatus) Type() string {
	return "ThroughputPatchStatus"
}

type UpdateEndpointCustomTagsRequest struct {
	// The new custom tags for the vector search endpoint
	CustomTags []CustomTag `json:"custom_tags"`
	// Name of the vector search endpoint
	EndpointName string `json:"-" url:"-"`
}

type UpdateEndpointCustomTagsResponse struct {
	// All the custom tags that are applied to the vector search endpoint.
	CustomTags []CustomTag `json:"custom_tags,omitempty"`
	// The name of the vector search endpoint whose custom tags were updated.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateEndpointCustomTagsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateEndpointCustomTagsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateVectorIndexUsagePolicyRequest struct {
	// Name of the vector search index
	IndexName string `json:"-" url:"-"`
}

type UpdateVectorIndexUsagePolicyResponse struct {
}

type UpsertDataResult struct {
	// List of primary keys for rows that failed to process.
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	// Count of successfully processed rows.
	SuccessRowCount int64 `json:"success_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpsertDataResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpsertDataResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	IndexName string `json:"-" url:"-"`
	// JSON string representing the data to be upserted.
	InputsJson string `json:"inputs_json"`
}

type UpsertDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result *UpsertDataResult `json:"result,omitempty"`
	// Status of the upsert operation.
	Status UpsertDataStatus `json:"status,omitempty"`
}

type Value struct {
	BoolValue bool `json:"bool_value,omitempty"`

	ListValue *ListValue `json:"list_value,omitempty"`

	NumberValue float64 `json:"number_value,omitempty"`

	StringValue string `json:"string_value,omitempty"`

	StructValue *Struct `json:"struct_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Value) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Value) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VectorIndex struct {
	// The user who created the index.
	Creator string `json:"creator,omitempty"`

	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecResponse `json:"delta_sync_index_spec,omitempty"`

	DirectAccessIndexSpec *DirectAccessVectorIndexSpec `json:"direct_access_index_spec,omitempty"`
	// Name of the endpoint associated with the index
	EndpointName string `json:"endpoint_name,omitempty"`

	IndexType VectorIndexType `json:"index_type,omitempty"`
	// Name of the index
	Name string `json:"name,omitempty"`
	// Primary key of the index
	PrimaryKey string `json:"primary_key,omitempty"`

	Status *VectorIndexStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *VectorIndexStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VectorIndexStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

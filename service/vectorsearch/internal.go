// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func columnInfoToPb(st *ColumnInfo) (*columnInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &columnInfoPb{}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type columnInfoPb struct {
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

type createEndpointPb struct {
	BudgetPolicyId string       `json:"budget_policy_id,omitempty"`
	EndpointType   EndpointType `json:"endpoint_type"`
	Name           string       `json:"name"`

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

func createVectorIndexRequestToPb(st *CreateVectorIndexRequest) (*createVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createVectorIndexRequestPb{}
	pb.DeltaSyncIndexSpec = st.DeltaSyncIndexSpec
	pb.DirectAccessIndexSpec = st.DirectAccessIndexSpec
	pb.EndpointName = st.EndpointName
	pb.IndexType = st.IndexType
	pb.Name = st.Name
	pb.PrimaryKey = st.PrimaryKey

	return pb, nil
}

type createVectorIndexRequestPb struct {
	DeltaSyncIndexSpec    *DeltaSyncVectorIndexSpecRequest `json:"delta_sync_index_spec,omitempty"`
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec     `json:"direct_access_index_spec,omitempty"`
	EndpointName          string                           `json:"endpoint_name"`
	IndexType             VectorIndexType                  `json:"index_type"`
	Name                  string                           `json:"name"`
	PrimaryKey            string                           `json:"primary_key"`
}

func createVectorIndexRequestFromPb(pb *createVectorIndexRequestPb) (*CreateVectorIndexRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVectorIndexRequest{}
	st.DeltaSyncIndexSpec = pb.DeltaSyncIndexSpec
	st.DirectAccessIndexSpec = pb.DirectAccessIndexSpec
	st.EndpointName = pb.EndpointName
	st.IndexType = pb.IndexType
	st.Name = pb.Name
	st.PrimaryKey = pb.PrimaryKey

	return st, nil
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

type customTagPb struct {
	Key   string `json:"key"`
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

type deleteDataResultPb struct {
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	SuccessRowCount   int64    `json:"success_row_count,omitempty"`

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

func deleteDataVectorIndexRequestToPb(st *DeleteDataVectorIndexRequest) (*deleteDataVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDataVectorIndexRequestPb{}
	pb.IndexName = st.IndexName
	pb.PrimaryKeys = st.PrimaryKeys

	return pb, nil
}

type deleteDataVectorIndexRequestPb struct {
	IndexName   string   `json:"-" url:"-"`
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

func deleteDataVectorIndexResponseToPb(st *DeleteDataVectorIndexResponse) (*deleteDataVectorIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDataVectorIndexResponsePb{}
	pb.Result = st.Result
	pb.Status = st.Status

	return pb, nil
}

type deleteDataVectorIndexResponsePb struct {
	Result *DeleteDataResult `json:"result,omitempty"`
	Status DeleteDataStatus  `json:"status,omitempty"`
}

func deleteDataVectorIndexResponseFromPb(pb *deleteDataVectorIndexResponsePb) (*DeleteDataVectorIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDataVectorIndexResponse{}
	st.Result = pb.Result
	st.Status = pb.Status

	return st, nil
}

func deleteEndpointRequestToPb(st *DeleteEndpointRequest) (*deleteEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteEndpointRequestPb{}
	pb.EndpointName = st.EndpointName

	return pb, nil
}

type deleteEndpointRequestPb struct {
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

func deleteEndpointResponseToPb(st *DeleteEndpointResponse) (*deleteEndpointResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteEndpointResponsePb{}

	return pb, nil
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

func deleteIndexRequestToPb(st *DeleteIndexRequest) (*deleteIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteIndexRequestPb{}
	pb.IndexName = st.IndexName

	return pb, nil
}

type deleteIndexRequestPb struct {
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

func deleteIndexResponseToPb(st *DeleteIndexResponse) (*deleteIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteIndexResponsePb{}

	return pb, nil
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

func deltaSyncVectorIndexSpecRequestToPb(st *DeltaSyncVectorIndexSpecRequest) (*deltaSyncVectorIndexSpecRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSyncVectorIndexSpecRequestPb{}
	pb.ColumnsToSync = st.ColumnsToSync
	pb.EmbeddingSourceColumns = st.EmbeddingSourceColumns
	pb.EmbeddingVectorColumns = st.EmbeddingVectorColumns
	pb.EmbeddingWritebackTable = st.EmbeddingWritebackTable
	pb.PipelineType = st.PipelineType
	pb.SourceTable = st.SourceTable

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deltaSyncVectorIndexSpecRequestPb struct {
	ColumnsToSync           []string                `json:"columns_to_sync,omitempty"`
	EmbeddingSourceColumns  []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	EmbeddingVectorColumns  []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	EmbeddingWritebackTable string                  `json:"embedding_writeback_table,omitempty"`
	PipelineType            PipelineType            `json:"pipeline_type,omitempty"`
	SourceTable             string                  `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deltaSyncVectorIndexSpecRequestFromPb(pb *deltaSyncVectorIndexSpecRequestPb) (*DeltaSyncVectorIndexSpecRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSyncVectorIndexSpecRequest{}
	st.ColumnsToSync = pb.ColumnsToSync
	st.EmbeddingSourceColumns = pb.EmbeddingSourceColumns
	st.EmbeddingVectorColumns = pb.EmbeddingVectorColumns
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

func deltaSyncVectorIndexSpecResponseToPb(st *DeltaSyncVectorIndexSpecResponse) (*deltaSyncVectorIndexSpecResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSyncVectorIndexSpecResponsePb{}
	pb.EmbeddingSourceColumns = st.EmbeddingSourceColumns
	pb.EmbeddingVectorColumns = st.EmbeddingVectorColumns
	pb.EmbeddingWritebackTable = st.EmbeddingWritebackTable
	pb.PipelineId = st.PipelineId
	pb.PipelineType = st.PipelineType
	pb.SourceTable = st.SourceTable

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deltaSyncVectorIndexSpecResponsePb struct {
	EmbeddingSourceColumns  []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	EmbeddingVectorColumns  []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	EmbeddingWritebackTable string                  `json:"embedding_writeback_table,omitempty"`
	PipelineId              string                  `json:"pipeline_id,omitempty"`
	PipelineType            PipelineType            `json:"pipeline_type,omitempty"`
	SourceTable             string                  `json:"source_table,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deltaSyncVectorIndexSpecResponseFromPb(pb *deltaSyncVectorIndexSpecResponsePb) (*DeltaSyncVectorIndexSpecResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSyncVectorIndexSpecResponse{}
	st.EmbeddingSourceColumns = pb.EmbeddingSourceColumns
	st.EmbeddingVectorColumns = pb.EmbeddingVectorColumns
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

func directAccessVectorIndexSpecToPb(st *DirectAccessVectorIndexSpec) (*directAccessVectorIndexSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &directAccessVectorIndexSpecPb{}
	pb.EmbeddingSourceColumns = st.EmbeddingSourceColumns
	pb.EmbeddingVectorColumns = st.EmbeddingVectorColumns
	pb.SchemaJson = st.SchemaJson

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type directAccessVectorIndexSpecPb struct {
	EmbeddingSourceColumns []EmbeddingSourceColumn `json:"embedding_source_columns,omitempty"`
	EmbeddingVectorColumns []EmbeddingVectorColumn `json:"embedding_vector_columns,omitempty"`
	SchemaJson             string                  `json:"schema_json,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func directAccessVectorIndexSpecFromPb(pb *directAccessVectorIndexSpecPb) (*DirectAccessVectorIndexSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DirectAccessVectorIndexSpec{}
	st.EmbeddingSourceColumns = pb.EmbeddingSourceColumns
	st.EmbeddingVectorColumns = pb.EmbeddingVectorColumns
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

type embeddingSourceColumnPb struct {
	EmbeddingModelEndpointName string `json:"embedding_model_endpoint_name,omitempty"`
	Name                       string `json:"name,omitempty"`

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

type embeddingVectorColumnPb struct {
	EmbeddingDimension int    `json:"embedding_dimension,omitempty"`
	Name               string `json:"name,omitempty"`

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

func endpointInfoToPb(st *EndpointInfo) (*endpointInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointInfoPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Creator = st.Creator
	pb.CustomTags = st.CustomTags
	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId
	pb.EndpointStatus = st.EndpointStatus
	pb.EndpointType = st.EndpointType
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.LastUpdatedUser = st.LastUpdatedUser
	pb.Name = st.Name
	pb.NumIndexes = st.NumIndexes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type endpointInfoPb struct {
	CreationTimestamp       int64           `json:"creation_timestamp,omitempty"`
	Creator                 string          `json:"creator,omitempty"`
	CustomTags              []CustomTag     `json:"custom_tags,omitempty"`
	EffectiveBudgetPolicyId string          `json:"effective_budget_policy_id,omitempty"`
	EndpointStatus          *EndpointStatus `json:"endpoint_status,omitempty"`
	EndpointType            EndpointType    `json:"endpoint_type,omitempty"`
	Id                      string          `json:"id,omitempty"`
	LastUpdatedTimestamp    int64           `json:"last_updated_timestamp,omitempty"`
	LastUpdatedUser         string          `json:"last_updated_user,omitempty"`
	Name                    string          `json:"name,omitempty"`
	NumIndexes              int             `json:"num_indexes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointInfoFromPb(pb *endpointInfoPb) (*EndpointInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointInfo{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Creator = pb.Creator
	st.CustomTags = pb.CustomTags
	st.EffectiveBudgetPolicyId = pb.EffectiveBudgetPolicyId
	st.EndpointStatus = pb.EndpointStatus
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

type endpointStatusPb struct {
	Message string              `json:"message,omitempty"`
	State   EndpointStatusState `json:"state,omitempty"`

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

func getEndpointRequestToPb(st *GetEndpointRequest) (*getEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEndpointRequestPb{}
	pb.EndpointName = st.EndpointName

	return pb, nil
}

type getEndpointRequestPb struct {
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

func getIndexRequestToPb(st *GetIndexRequest) (*getIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getIndexRequestPb{}
	pb.IndexName = st.IndexName

	return pb, nil
}

type getIndexRequestPb struct {
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

func listEndpointResponseToPb(st *ListEndpointResponse) (*listEndpointResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listEndpointResponsePb{}
	pb.Endpoints = st.Endpoints
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listEndpointResponsePb struct {
	Endpoints     []EndpointInfo `json:"endpoints,omitempty"`
	NextPageToken string         `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listEndpointResponseFromPb(pb *listEndpointResponsePb) (*ListEndpointResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListEndpointResponse{}
	st.Endpoints = pb.Endpoints
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

func listEndpointsRequestToPb(st *ListEndpointsRequest) (*listEndpointsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listEndpointsRequestPb{}
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listEndpointsRequestPb struct {
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

type listIndexesRequestPb struct {
	EndpointName string `json:"-" url:"endpoint_name"`
	PageToken    string `json:"-" url:"page_token,omitempty"`

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

func listValueToPb(st *ListValue) (*listValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listValuePb{}
	pb.Values = st.Values

	return pb, nil
}

type listValuePb struct {
	Values []Value `json:"values,omitempty"`
}

func listValueFromPb(pb *listValuePb) (*ListValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListValue{}
	st.Values = pb.Values

	return st, nil
}

func listVectorIndexesResponseToPb(st *ListVectorIndexesResponse) (*listVectorIndexesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVectorIndexesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.VectorIndexes = st.VectorIndexes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listVectorIndexesResponsePb struct {
	NextPageToken string            `json:"next_page_token,omitempty"`
	VectorIndexes []MiniVectorIndex `json:"vector_indexes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listVectorIndexesResponseFromPb(pb *listVectorIndexesResponsePb) (*ListVectorIndexesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListVectorIndexesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.VectorIndexes = pb.VectorIndexes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listVectorIndexesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listVectorIndexesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func mapStringValueEntryToPb(st *MapStringValueEntry) (*mapStringValueEntryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mapStringValueEntryPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type mapStringValueEntryPb struct {
	Key   string `json:"key,omitempty"`
	Value *Value `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func mapStringValueEntryFromPb(pb *mapStringValueEntryPb) (*MapStringValueEntry, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MapStringValueEntry{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *mapStringValueEntryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st mapStringValueEntryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

type miniVectorIndexPb struct {
	Creator      string          `json:"creator,omitempty"`
	EndpointName string          `json:"endpoint_name,omitempty"`
	IndexType    VectorIndexType `json:"index_type,omitempty"`
	Name         string          `json:"name,omitempty"`
	PrimaryKey   string          `json:"primary_key,omitempty"`

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

func patchEndpointBudgetPolicyRequestToPb(st *PatchEndpointBudgetPolicyRequest) (*patchEndpointBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &patchEndpointBudgetPolicyRequestPb{}
	pb.BudgetPolicyId = st.BudgetPolicyId
	pb.EndpointName = st.EndpointName

	return pb, nil
}

type patchEndpointBudgetPolicyRequestPb struct {
	BudgetPolicyId string `json:"budget_policy_id"`
	EndpointName   string `json:"-" url:"-"`
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

func patchEndpointBudgetPolicyResponseToPb(st *PatchEndpointBudgetPolicyResponse) (*patchEndpointBudgetPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &patchEndpointBudgetPolicyResponsePb{}
	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type patchEndpointBudgetPolicyResponsePb struct {
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

type queryVectorIndexNextPageRequestPb struct {
	EndpointName string `json:"endpoint_name,omitempty"`
	IndexName    string `json:"-" url:"-"`
	PageToken    string `json:"page_token,omitempty"`

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

type queryVectorIndexRequestPb struct {
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

func queryVectorIndexResponseToPb(st *QueryVectorIndexResponse) (*queryVectorIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryVectorIndexResponsePb{}
	pb.Manifest = st.Manifest
	pb.NextPageToken = st.NextPageToken
	pb.Result = st.Result

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type queryVectorIndexResponsePb struct {
	Manifest      *ResultManifest `json:"manifest,omitempty"`
	NextPageToken string          `json:"next_page_token,omitempty"`
	Result        *ResultData     `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryVectorIndexResponseFromPb(pb *queryVectorIndexResponsePb) (*QueryVectorIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryVectorIndexResponse{}
	st.Manifest = pb.Manifest
	st.NextPageToken = pb.NextPageToken
	st.Result = pb.Result

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryVectorIndexResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryVectorIndexResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func resultDataToPb(st *ResultData) (*resultDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultDataPb{}
	pb.DataArray = st.DataArray
	pb.RowCount = st.RowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type resultDataPb struct {
	DataArray [][]string `json:"data_array,omitempty"`
	RowCount  int        `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resultDataFromPb(pb *resultDataPb) (*ResultData, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultData{}
	st.DataArray = pb.DataArray
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

func resultManifestToPb(st *ResultManifest) (*resultManifestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultManifestPb{}
	pb.ColumnCount = st.ColumnCount
	pb.Columns = st.Columns

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type resultManifestPb struct {
	ColumnCount int          `json:"column_count,omitempty"`
	Columns     []ColumnInfo `json:"columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resultManifestFromPb(pb *resultManifestPb) (*ResultManifest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultManifest{}
	st.ColumnCount = pb.ColumnCount
	st.Columns = pb.Columns

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resultManifestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resultManifestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

type scanVectorIndexRequestPb struct {
	IndexName      string `json:"-" url:"-"`
	LastPrimaryKey string `json:"last_primary_key,omitempty"`
	NumResults     int    `json:"num_results,omitempty"`

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

func scanVectorIndexResponseToPb(st *ScanVectorIndexResponse) (*scanVectorIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &scanVectorIndexResponsePb{}
	pb.Data = st.Data
	pb.LastPrimaryKey = st.LastPrimaryKey

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type scanVectorIndexResponsePb struct {
	Data           []Struct `json:"data,omitempty"`
	LastPrimaryKey string   `json:"last_primary_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func scanVectorIndexResponseFromPb(pb *scanVectorIndexResponsePb) (*ScanVectorIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ScanVectorIndexResponse{}
	st.Data = pb.Data
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

func structToPb(st *Struct) (*structPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &structPb{}
	pb.Fields = st.Fields

	return pb, nil
}

type structPb struct {
	Fields []MapStringValueEntry `json:"fields,omitempty"`
}

func structFromPb(pb *structPb) (*Struct, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Struct{}
	st.Fields = pb.Fields

	return st, nil
}

func syncIndexRequestToPb(st *SyncIndexRequest) (*syncIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncIndexRequestPb{}
	pb.IndexName = st.IndexName

	return pb, nil
}

type syncIndexRequestPb struct {
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

func syncIndexResponseToPb(st *SyncIndexResponse) (*syncIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncIndexResponsePb{}

	return pb, nil
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

func updateEndpointCustomTagsRequestToPb(st *UpdateEndpointCustomTagsRequest) (*updateEndpointCustomTagsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEndpointCustomTagsRequestPb{}
	pb.CustomTags = st.CustomTags
	pb.EndpointName = st.EndpointName

	return pb, nil
}

type updateEndpointCustomTagsRequestPb struct {
	CustomTags   []CustomTag `json:"custom_tags"`
	EndpointName string      `json:"-" url:"-"`
}

func updateEndpointCustomTagsRequestFromPb(pb *updateEndpointCustomTagsRequestPb) (*UpdateEndpointCustomTagsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEndpointCustomTagsRequest{}
	st.CustomTags = pb.CustomTags
	st.EndpointName = pb.EndpointName

	return st, nil
}

func updateEndpointCustomTagsResponseToPb(st *UpdateEndpointCustomTagsResponse) (*updateEndpointCustomTagsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEndpointCustomTagsResponsePb{}
	pb.CustomTags = st.CustomTags
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateEndpointCustomTagsResponsePb struct {
	CustomTags []CustomTag `json:"custom_tags,omitempty"`
	Name       string      `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateEndpointCustomTagsResponseFromPb(pb *updateEndpointCustomTagsResponsePb) (*UpdateEndpointCustomTagsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEndpointCustomTagsResponse{}
	st.CustomTags = pb.CustomTags
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

type upsertDataResultPb struct {
	FailedPrimaryKeys []string `json:"failed_primary_keys,omitempty"`
	SuccessRowCount   int64    `json:"success_row_count,omitempty"`

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

func upsertDataVectorIndexRequestToPb(st *UpsertDataVectorIndexRequest) (*upsertDataVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &upsertDataVectorIndexRequestPb{}
	pb.IndexName = st.IndexName
	pb.InputsJson = st.InputsJson

	return pb, nil
}

type upsertDataVectorIndexRequestPb struct {
	IndexName  string `json:"-" url:"-"`
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

func upsertDataVectorIndexResponseToPb(st *UpsertDataVectorIndexResponse) (*upsertDataVectorIndexResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &upsertDataVectorIndexResponsePb{}
	pb.Result = st.Result
	pb.Status = st.Status

	return pb, nil
}

type upsertDataVectorIndexResponsePb struct {
	Result *UpsertDataResult `json:"result,omitempty"`
	Status UpsertDataStatus  `json:"status,omitempty"`
}

func upsertDataVectorIndexResponseFromPb(pb *upsertDataVectorIndexResponsePb) (*UpsertDataVectorIndexResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpsertDataVectorIndexResponse{}
	st.Result = pb.Result
	st.Status = pb.Status

	return st, nil
}

func valueToPb(st *Value) (*valuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &valuePb{}
	pb.BoolValue = st.BoolValue
	pb.ListValue = st.ListValue
	pb.NumberValue = st.NumberValue
	pb.StringValue = st.StringValue
	pb.StructValue = st.StructValue

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type valuePb struct {
	BoolValue   bool       `json:"bool_value,omitempty"`
	ListValue   *ListValue `json:"list_value,omitempty"`
	NumberValue float64    `json:"number_value,omitempty"`
	StringValue string     `json:"string_value,omitempty"`
	StructValue *Struct    `json:"struct_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func valueFromPb(pb *valuePb) (*Value, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Value{}
	st.BoolValue = pb.BoolValue
	st.ListValue = pb.ListValue
	st.NumberValue = pb.NumberValue
	st.StringValue = pb.StringValue
	st.StructValue = pb.StructValue

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *valuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st valuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func vectorIndexToPb(st *VectorIndex) (*vectorIndexPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorIndexPb{}
	pb.Creator = st.Creator
	pb.DeltaSyncIndexSpec = st.DeltaSyncIndexSpec
	pb.DirectAccessIndexSpec = st.DirectAccessIndexSpec
	pb.EndpointName = st.EndpointName
	pb.IndexType = st.IndexType
	pb.Name = st.Name
	pb.PrimaryKey = st.PrimaryKey
	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type vectorIndexPb struct {
	Creator               string                            `json:"creator,omitempty"`
	DeltaSyncIndexSpec    *DeltaSyncVectorIndexSpecResponse `json:"delta_sync_index_spec,omitempty"`
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec      `json:"direct_access_index_spec,omitempty"`
	EndpointName          string                            `json:"endpoint_name,omitempty"`
	IndexType             VectorIndexType                   `json:"index_type,omitempty"`
	Name                  string                            `json:"name,omitempty"`
	PrimaryKey            string                            `json:"primary_key,omitempty"`
	Status                *VectorIndexStatus                `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func vectorIndexFromPb(pb *vectorIndexPb) (*VectorIndex, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VectorIndex{}
	st.Creator = pb.Creator
	st.DeltaSyncIndexSpec = pb.DeltaSyncIndexSpec
	st.DirectAccessIndexSpec = pb.DirectAccessIndexSpec
	st.EndpointName = pb.EndpointName
	st.IndexType = pb.IndexType
	st.Name = pb.Name
	st.PrimaryKey = pb.PrimaryKey
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *vectorIndexPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st vectorIndexPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

type vectorIndexStatusPb struct {
	IndexUrl        string `json:"index_url,omitempty"`
	IndexedRowCount int64  `json:"indexed_row_count,omitempty"`
	Message         string `json:"message,omitempty"`
	Ready           bool   `json:"ready,omitempty"`

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

// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func identity[T any](obj *T) (*T, error) {
	return obj, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

// Helper to strip trailing zeros in fractional part
func rstripZeros(s string) string {
	for len(s) > 0 && s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '.' {
		s = s[:len(s)-1]
	}
	return s
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

type ColumnInfo struct {
	// Name of the column.
	Name string

	ForceSendFields []string
}

func columnInfoToPb(st *ColumnInfo) (*columnInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &columnInfoPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	BudgetPolicyId string
	// Type of endpoint
	EndpointType EndpointType
	// Name of the vector search endpoint
	Name string

	ForceSendFields []string
}

func createEndpointToPb(st *CreateEndpoint) (*createEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createEndpointPb{}
	budgetPolicyIdPb, err := identity(&st.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdPb != nil {
		pb.BudgetPolicyId = *budgetPolicyIdPb
	}

	endpointTypePb, err := identity(&st.EndpointType)
	if err != nil {
		return nil, err
	}
	if endpointTypePb != nil {
		pb.EndpointType = *endpointTypePb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	budgetPolicyIdField, err := identity(&pb.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdField != nil {
		st.BudgetPolicyId = *budgetPolicyIdField
	}
	endpointTypeField, err := identity(&pb.EndpointType)
	if err != nil {
		return nil, err
	}
	if endpointTypeField != nil {
		st.EndpointType = *endpointTypeField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecRequest
	// Specification for Direct Vector Access Index. Required if `index_type` is
	// `DIRECT_ACCESS`.
	DirectAccessIndexSpec *DirectAccessVectorIndexSpec
	// Name of the endpoint to be used for serving the index
	EndpointName string
	// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
	// automatically syncs with a source Delta Table, automatically and
	// incrementally updating the index as the underlying data in the Delta
	// Table changes. - `DIRECT_ACCESS`: An index that supports direct read and
	// write of vectors and metadata through our REST and SDK APIs. With this
	// model, the user manages index updates.
	IndexType VectorIndexType
	// Name of the index
	Name string
	// Primary key of the index
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

	endpointNamePb, err := identity(&st.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNamePb != nil {
		pb.EndpointName = *endpointNamePb
	}

	indexTypePb, err := identity(&st.IndexType)
	if err != nil {
		return nil, err
	}
	if indexTypePb != nil {
		pb.IndexType = *indexTypePb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	primaryKeyPb, err := identity(&st.PrimaryKey)
	if err != nil {
		return nil, err
	}
	if primaryKeyPb != nil {
		pb.PrimaryKey = *primaryKeyPb
	}

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
	endpointNameField, err := identity(&pb.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNameField != nil {
		st.EndpointName = *endpointNameField
	}
	indexTypeField, err := identity(&pb.IndexType)
	if err != nil {
		return nil, err
	}
	if indexTypeField != nil {
		st.IndexType = *indexTypeField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	primaryKeyField, err := identity(&pb.PrimaryKey)
	if err != nil {
		return nil, err
	}
	if primaryKeyField != nil {
		st.PrimaryKey = *primaryKeyField
	}

	return st, nil
}

type CustomTag struct {
	// Key field for a vector search endpoint tag.
	Key string
	// [Optional] Value field for a vector search endpoint tag.
	Value string

	ForceSendFields []string
}

func customTagToPb(st *CustomTag) (*customTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customTagPb{}
	keyPb, err := identity(&st.Key)
	if err != nil {
		return nil, err
	}
	if keyPb != nil {
		pb.Key = *keyPb
	}

	valuePb, err := identity(&st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = *valuePb
	}

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
	keyField, err := identity(&pb.Key)
	if err != nil {
		return nil, err
	}
	if keyField != nil {
		st.Key = *keyField
	}
	valueField, err := identity(&pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = *valueField
	}

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
	FailedPrimaryKeys []string
	// Count of successfully processed rows.
	SuccessRowCount int64

	ForceSendFields []string
}

func deleteDataResultToPb(st *DeleteDataResult) (*deleteDataResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDataResultPb{}

	var failedPrimaryKeysPb []string
	for _, item := range st.FailedPrimaryKeys {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			failedPrimaryKeysPb = append(failedPrimaryKeysPb, *itemPb)
		}
	}
	pb.FailedPrimaryKeys = failedPrimaryKeysPb

	successRowCountPb, err := identity(&st.SuccessRowCount)
	if err != nil {
		return nil, err
	}
	if successRowCountPb != nil {
		pb.SuccessRowCount = *successRowCountPb
	}

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

	var failedPrimaryKeysField []string
	for _, itemPb := range pb.FailedPrimaryKeys {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			failedPrimaryKeysField = append(failedPrimaryKeysField, *item)
		}
	}
	st.FailedPrimaryKeys = failedPrimaryKeysField
	successRowCountField, err := identity(&pb.SuccessRowCount)
	if err != nil {
		return nil, err
	}
	if successRowCountField != nil {
		st.SuccessRowCount = *successRowCountField
	}

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
	IndexName string
	// List of primary keys for the data to be deleted.
	PrimaryKeys []string
}

func deleteDataVectorIndexRequestToPb(st *DeleteDataVectorIndexRequest) (*deleteDataVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDataVectorIndexRequestPb{}
	indexNamePb, err := identity(&st.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNamePb != nil {
		pb.IndexName = *indexNamePb
	}

	var primaryKeysPb []string
	for _, item := range st.PrimaryKeys {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			primaryKeysPb = append(primaryKeysPb, *itemPb)
		}
	}
	pb.PrimaryKeys = primaryKeysPb

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
	indexNameField, err := identity(&pb.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNameField != nil {
		st.IndexName = *indexNameField
	}

	var primaryKeysField []string
	for _, itemPb := range pb.PrimaryKeys {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			primaryKeysField = append(primaryKeysField, *item)
		}
	}
	st.PrimaryKeys = primaryKeysField

	return st, nil
}

type DeleteDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result *DeleteDataResult
	// Status of the delete operation.
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

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

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
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	return st, nil
}

// Delete an endpoint
type DeleteEndpointRequest struct {
	// Name of the vector search endpoint
	EndpointName string
}

func deleteEndpointRequestToPb(st *DeleteEndpointRequest) (*deleteEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteEndpointRequestPb{}
	endpointNamePb, err := identity(&st.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNamePb != nil {
		pb.EndpointName = *endpointNamePb
	}

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
	endpointNameField, err := identity(&pb.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNameField != nil {
		st.EndpointName = *endpointNameField
	}

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
	IndexName string
}

func deleteIndexRequestToPb(st *DeleteIndexRequest) (*deleteIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteIndexRequestPb{}
	indexNamePb, err := identity(&st.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNamePb != nil {
		pb.IndexName = *indexNamePb
	}

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
	indexNameField, err := identity(&pb.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNameField != nil {
		st.IndexName = *indexNameField
	}

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
	ColumnsToSync []string
	// The columns that contain the embedding source.
	EmbeddingSourceColumns []EmbeddingSourceColumn
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns []EmbeddingVectorColumn
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	EmbeddingWritebackTable string
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	PipelineType PipelineType
	// The name of the source table.
	SourceTable string

	ForceSendFields []string
}

func deltaSyncVectorIndexSpecRequestToPb(st *DeltaSyncVectorIndexSpecRequest) (*deltaSyncVectorIndexSpecRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSyncVectorIndexSpecRequestPb{}

	var columnsToSyncPb []string
	for _, item := range st.ColumnsToSync {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsToSyncPb = append(columnsToSyncPb, *itemPb)
		}
	}
	pb.ColumnsToSync = columnsToSyncPb

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

	embeddingWritebackTablePb, err := identity(&st.EmbeddingWritebackTable)
	if err != nil {
		return nil, err
	}
	if embeddingWritebackTablePb != nil {
		pb.EmbeddingWritebackTable = *embeddingWritebackTablePb
	}

	pipelineTypePb, err := identity(&st.PipelineType)
	if err != nil {
		return nil, err
	}
	if pipelineTypePb != nil {
		pb.PipelineType = *pipelineTypePb
	}

	sourceTablePb, err := identity(&st.SourceTable)
	if err != nil {
		return nil, err
	}
	if sourceTablePb != nil {
		pb.SourceTable = *sourceTablePb
	}

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

	var columnsToSyncField []string
	for _, itemPb := range pb.ColumnsToSync {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsToSyncField = append(columnsToSyncField, *item)
		}
	}
	st.ColumnsToSync = columnsToSyncField

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
	embeddingWritebackTableField, err := identity(&pb.EmbeddingWritebackTable)
	if err != nil {
		return nil, err
	}
	if embeddingWritebackTableField != nil {
		st.EmbeddingWritebackTable = *embeddingWritebackTableField
	}
	pipelineTypeField, err := identity(&pb.PipelineType)
	if err != nil {
		return nil, err
	}
	if pipelineTypeField != nil {
		st.PipelineType = *pipelineTypeField
	}
	sourceTableField, err := identity(&pb.SourceTable)
	if err != nil {
		return nil, err
	}
	if sourceTableField != nil {
		st.SourceTable = *sourceTableField
	}

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
	EmbeddingSourceColumns []EmbeddingSourceColumn
	// The columns that contain the embedding vectors.
	EmbeddingVectorColumns []EmbeddingVectorColumn
	// [Optional] Name of the Delta table to sync the vector index contents and
	// computed embeddings to.
	EmbeddingWritebackTable string
	// The ID of the pipeline that is used to sync the index.
	PipelineId string
	// Pipeline execution mode. - `TRIGGERED`: If the pipeline uses the
	// triggered execution mode, the system stops processing after successfully
	// refreshing the source table in the pipeline once, ensuring the table is
	// updated based on the data available when the update started. -
	// `CONTINUOUS`: If the pipeline uses continuous execution, the pipeline
	// processes new data as it arrives in the source table to keep vector index
	// fresh.
	PipelineType PipelineType
	// The name of the source table.
	SourceTable string

	ForceSendFields []string
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

	embeddingWritebackTablePb, err := identity(&st.EmbeddingWritebackTable)
	if err != nil {
		return nil, err
	}
	if embeddingWritebackTablePb != nil {
		pb.EmbeddingWritebackTable = *embeddingWritebackTablePb
	}

	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

	pipelineTypePb, err := identity(&st.PipelineType)
	if err != nil {
		return nil, err
	}
	if pipelineTypePb != nil {
		pb.PipelineType = *pipelineTypePb
	}

	sourceTablePb, err := identity(&st.SourceTable)
	if err != nil {
		return nil, err
	}
	if sourceTablePb != nil {
		pb.SourceTable = *sourceTablePb
	}

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
	embeddingWritebackTableField, err := identity(&pb.EmbeddingWritebackTable)
	if err != nil {
		return nil, err
	}
	if embeddingWritebackTableField != nil {
		st.EmbeddingWritebackTable = *embeddingWritebackTableField
	}
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}
	pipelineTypeField, err := identity(&pb.PipelineType)
	if err != nil {
		return nil, err
	}
	if pipelineTypeField != nil {
		st.PipelineType = *pipelineTypeField
	}
	sourceTableField, err := identity(&pb.SourceTable)
	if err != nil {
		return nil, err
	}
	if sourceTableField != nil {
		st.SourceTable = *sourceTableField
	}

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
	EmbeddingSourceColumns []EmbeddingSourceColumn
	// The columns that contain the embedding vectors. The format should be
	// array[double].
	EmbeddingVectorColumns []EmbeddingVectorColumn
	// The schema of the index in JSON format. Supported types are `integer`,
	// `long`, `float`, `double`, `boolean`, `string`, `date`, `timestamp`.
	// Supported types for vector column: `array<float>`, `array<double>`,`.
	SchemaJson string

	ForceSendFields []string
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

	schemaJsonPb, err := identity(&st.SchemaJson)
	if err != nil {
		return nil, err
	}
	if schemaJsonPb != nil {
		pb.SchemaJson = *schemaJsonPb
	}

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
	schemaJsonField, err := identity(&pb.SchemaJson)
	if err != nil {
		return nil, err
	}
	if schemaJsonField != nil {
		st.SchemaJson = *schemaJsonField
	}

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
	EmbeddingModelEndpointName string
	// Name of the column
	Name string

	ForceSendFields []string
}

func embeddingSourceColumnToPb(st *EmbeddingSourceColumn) (*embeddingSourceColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &embeddingSourceColumnPb{}
	embeddingModelEndpointNamePb, err := identity(&st.EmbeddingModelEndpointName)
	if err != nil {
		return nil, err
	}
	if embeddingModelEndpointNamePb != nil {
		pb.EmbeddingModelEndpointName = *embeddingModelEndpointNamePb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	embeddingModelEndpointNameField, err := identity(&pb.EmbeddingModelEndpointName)
	if err != nil {
		return nil, err
	}
	if embeddingModelEndpointNameField != nil {
		st.EmbeddingModelEndpointName = *embeddingModelEndpointNameField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	EmbeddingDimension int
	// Name of the column
	Name string

	ForceSendFields []string
}

func embeddingVectorColumnToPb(st *EmbeddingVectorColumn) (*embeddingVectorColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &embeddingVectorColumnPb{}
	embeddingDimensionPb, err := identity(&st.EmbeddingDimension)
	if err != nil {
		return nil, err
	}
	if embeddingDimensionPb != nil {
		pb.EmbeddingDimension = *embeddingDimensionPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	embeddingDimensionField, err := identity(&pb.EmbeddingDimension)
	if err != nil {
		return nil, err
	}
	if embeddingDimensionField != nil {
		st.EmbeddingDimension = *embeddingDimensionField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	CreationTimestamp int64
	// Creator of the endpoint
	Creator string
	// The custom tags assigned to the endpoint
	CustomTags []CustomTag
	// The budget policy id applied to the endpoint
	EffectiveBudgetPolicyId string
	// Current status of the endpoint
	EndpointStatus *EndpointStatus
	// Type of endpoint
	EndpointType EndpointType
	// Unique identifier of the endpoint
	Id string
	// Timestamp of last update to the endpoint
	LastUpdatedTimestamp int64
	// User who last updated the endpoint
	LastUpdatedUser string
	// Name of the vector search endpoint
	Name string
	// Number of indexes on the endpoint
	NumIndexes int

	ForceSendFields []string
}

func endpointInfoToPb(st *EndpointInfo) (*endpointInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointInfoPb{}
	creationTimestampPb, err := identity(&st.CreationTimestamp)
	if err != nil {
		return nil, err
	}
	if creationTimestampPb != nil {
		pb.CreationTimestamp = *creationTimestampPb
	}

	creatorPb, err := identity(&st.Creator)
	if err != nil {
		return nil, err
	}
	if creatorPb != nil {
		pb.Creator = *creatorPb
	}

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

	effectiveBudgetPolicyIdPb, err := identity(&st.EffectiveBudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if effectiveBudgetPolicyIdPb != nil {
		pb.EffectiveBudgetPolicyId = *effectiveBudgetPolicyIdPb
	}

	endpointStatusPb, err := endpointStatusToPb(st.EndpointStatus)
	if err != nil {
		return nil, err
	}
	if endpointStatusPb != nil {
		pb.EndpointStatus = endpointStatusPb
	}

	endpointTypePb, err := identity(&st.EndpointType)
	if err != nil {
		return nil, err
	}
	if endpointTypePb != nil {
		pb.EndpointType = *endpointTypePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	lastUpdatedTimestampPb, err := identity(&st.LastUpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	if lastUpdatedTimestampPb != nil {
		pb.LastUpdatedTimestamp = *lastUpdatedTimestampPb
	}

	lastUpdatedUserPb, err := identity(&st.LastUpdatedUser)
	if err != nil {
		return nil, err
	}
	if lastUpdatedUserPb != nil {
		pb.LastUpdatedUser = *lastUpdatedUserPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	numIndexesPb, err := identity(&st.NumIndexes)
	if err != nil {
		return nil, err
	}
	if numIndexesPb != nil {
		pb.NumIndexes = *numIndexesPb
	}

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
	creationTimestampField, err := identity(&pb.CreationTimestamp)
	if err != nil {
		return nil, err
	}
	if creationTimestampField != nil {
		st.CreationTimestamp = *creationTimestampField
	}
	creatorField, err := identity(&pb.Creator)
	if err != nil {
		return nil, err
	}
	if creatorField != nil {
		st.Creator = *creatorField
	}

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
	effectiveBudgetPolicyIdField, err := identity(&pb.EffectiveBudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if effectiveBudgetPolicyIdField != nil {
		st.EffectiveBudgetPolicyId = *effectiveBudgetPolicyIdField
	}
	endpointStatusField, err := endpointStatusFromPb(pb.EndpointStatus)
	if err != nil {
		return nil, err
	}
	if endpointStatusField != nil {
		st.EndpointStatus = endpointStatusField
	}
	endpointTypeField, err := identity(&pb.EndpointType)
	if err != nil {
		return nil, err
	}
	if endpointTypeField != nil {
		st.EndpointType = *endpointTypeField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	lastUpdatedTimestampField, err := identity(&pb.LastUpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	if lastUpdatedTimestampField != nil {
		st.LastUpdatedTimestamp = *lastUpdatedTimestampField
	}
	lastUpdatedUserField, err := identity(&pb.LastUpdatedUser)
	if err != nil {
		return nil, err
	}
	if lastUpdatedUserField != nil {
		st.LastUpdatedUser = *lastUpdatedUserField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	numIndexesField, err := identity(&pb.NumIndexes)
	if err != nil {
		return nil, err
	}
	if numIndexesField != nil {
		st.NumIndexes = *numIndexesField
	}

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
	Message string
	// Current state of the endpoint
	State EndpointStatusState

	ForceSendFields []string
}

func endpointStatusToPb(st *EndpointStatus) (*endpointStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointStatusPb{}
	messagePb, err := identity(&st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = *messagePb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

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
	messageField, err := identity(&pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = *messageField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

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
	EndpointName string
}

func getEndpointRequestToPb(st *GetEndpointRequest) (*getEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEndpointRequestPb{}
	endpointNamePb, err := identity(&st.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNamePb != nil {
		pb.EndpointName = *endpointNamePb
	}

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
	endpointNameField, err := identity(&pb.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNameField != nil {
		st.EndpointName = *endpointNameField
	}

	return st, nil
}

// Get an index
type GetIndexRequest struct {
	// Name of the index
	IndexName string
}

func getIndexRequestToPb(st *GetIndexRequest) (*getIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getIndexRequestPb{}
	indexNamePb, err := identity(&st.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNamePb != nil {
		pb.IndexName = *indexNamePb
	}

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
	indexNameField, err := identity(&pb.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNameField != nil {
		st.IndexName = *indexNameField
	}

	return st, nil
}

type ListEndpointResponse struct {
	// An array of Endpoint objects
	Endpoints []EndpointInfo
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string

	ForceSendFields []string
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

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	PageToken string

	ForceSendFields []string
}

func listEndpointsRequestToPb(st *ListEndpointsRequest) (*listEndpointsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listEndpointsRequestPb{}
	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	EndpointName string
	// Token for pagination
	PageToken string

	ForceSendFields []string
}

func listIndexesRequestToPb(st *ListIndexesRequest) (*listIndexesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listIndexesRequestPb{}
	endpointNamePb, err := identity(&st.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNamePb != nil {
		pb.EndpointName = *endpointNamePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	endpointNameField, err := identity(&pb.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNameField != nil {
		st.EndpointName = *endpointNameField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	NextPageToken string

	VectorIndexes []MiniVectorIndex

	ForceSendFields []string
}

func listVectorIndexesResponseToPb(st *ListVectorIndexesResponse) (*listVectorIndexesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVectorIndexesResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	Key string
	// Column value, nullable.
	Value *Value

	ForceSendFields []string
}

func mapStringValueEntryToPb(st *MapStringValueEntry) (*mapStringValueEntryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mapStringValueEntryPb{}
	keyPb, err := identity(&st.Key)
	if err != nil {
		return nil, err
	}
	if keyPb != nil {
		pb.Key = *keyPb
	}

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
	keyField, err := identity(&pb.Key)
	if err != nil {
		return nil, err
	}
	if keyField != nil {
		st.Key = *keyField
	}
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
	Creator string
	// Name of the endpoint associated with the index
	EndpointName string
	// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
	// automatically syncs with a source Delta Table, automatically and
	// incrementally updating the index as the underlying data in the Delta
	// Table changes. - `DIRECT_ACCESS`: An index that supports direct read and
	// write of vectors and metadata through our REST and SDK APIs. With this
	// model, the user manages index updates.
	IndexType VectorIndexType
	// Name of the index
	Name string
	// Primary key of the index
	PrimaryKey string

	ForceSendFields []string
}

func miniVectorIndexToPb(st *MiniVectorIndex) (*miniVectorIndexPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &miniVectorIndexPb{}
	creatorPb, err := identity(&st.Creator)
	if err != nil {
		return nil, err
	}
	if creatorPb != nil {
		pb.Creator = *creatorPb
	}

	endpointNamePb, err := identity(&st.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNamePb != nil {
		pb.EndpointName = *endpointNamePb
	}

	indexTypePb, err := identity(&st.IndexType)
	if err != nil {
		return nil, err
	}
	if indexTypePb != nil {
		pb.IndexType = *indexTypePb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	primaryKeyPb, err := identity(&st.PrimaryKey)
	if err != nil {
		return nil, err
	}
	if primaryKeyPb != nil {
		pb.PrimaryKey = *primaryKeyPb
	}

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
	creatorField, err := identity(&pb.Creator)
	if err != nil {
		return nil, err
	}
	if creatorField != nil {
		st.Creator = *creatorField
	}
	endpointNameField, err := identity(&pb.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNameField != nil {
		st.EndpointName = *endpointNameField
	}
	indexTypeField, err := identity(&pb.IndexType)
	if err != nil {
		return nil, err
	}
	if indexTypeField != nil {
		st.IndexType = *indexTypeField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	primaryKeyField, err := identity(&pb.PrimaryKey)
	if err != nil {
		return nil, err
	}
	if primaryKeyField != nil {
		st.PrimaryKey = *primaryKeyField
	}

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
	BudgetPolicyId string
	// Name of the vector search endpoint
	EndpointName string
}

func patchEndpointBudgetPolicyRequestToPb(st *PatchEndpointBudgetPolicyRequest) (*patchEndpointBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &patchEndpointBudgetPolicyRequestPb{}
	budgetPolicyIdPb, err := identity(&st.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdPb != nil {
		pb.BudgetPolicyId = *budgetPolicyIdPb
	}

	endpointNamePb, err := identity(&st.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNamePb != nil {
		pb.EndpointName = *endpointNamePb
	}

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
	budgetPolicyIdField, err := identity(&pb.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdField != nil {
		st.BudgetPolicyId = *budgetPolicyIdField
	}
	endpointNameField, err := identity(&pb.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNameField != nil {
		st.EndpointName = *endpointNameField
	}

	return st, nil
}

type PatchEndpointBudgetPolicyResponse struct {
	// The budget policy applied to the vector search endpoint.
	EffectiveBudgetPolicyId string

	ForceSendFields []string
}

func patchEndpointBudgetPolicyResponseToPb(st *PatchEndpointBudgetPolicyResponse) (*patchEndpointBudgetPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &patchEndpointBudgetPolicyResponsePb{}
	effectiveBudgetPolicyIdPb, err := identity(&st.EffectiveBudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if effectiveBudgetPolicyIdPb != nil {
		pb.EffectiveBudgetPolicyId = *effectiveBudgetPolicyIdPb
	}

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
	effectiveBudgetPolicyIdField, err := identity(&pb.EffectiveBudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if effectiveBudgetPolicyIdField != nil {
		st.EffectiveBudgetPolicyId = *effectiveBudgetPolicyIdField
	}

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
	EndpointName string
	// Name of the vector index to query.
	IndexName string
	// Page token returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` API.
	PageToken string

	ForceSendFields []string
}

func queryVectorIndexNextPageRequestToPb(st *QueryVectorIndexNextPageRequest) (*queryVectorIndexNextPageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryVectorIndexNextPageRequestPb{}
	endpointNamePb, err := identity(&st.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNamePb != nil {
		pb.EndpointName = *endpointNamePb
	}

	indexNamePb, err := identity(&st.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNamePb != nil {
		pb.IndexName = *indexNamePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	endpointNameField, err := identity(&pb.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNameField != nil {
		st.EndpointName = *endpointNameField
	}
	indexNameField, err := identity(&pb.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNameField != nil {
		st.IndexName = *indexNameField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	Columns []string
	// Column names used to retrieve data to send to the reranker.
	ColumnsToRerank []string
	// JSON string representing query filters.
	//
	// Example filters:
	//
	// - `{"id <": 5}`: Filter for id less than 5. - `{"id >": 5}`: Filter for
	// id greater than 5. - `{"id <=": 5}`: Filter for id less than equal to 5.
	// - `{"id >=": 5}`: Filter for id greater than equal to 5. - `{"id": 5}`:
	// Filter for id equal to 5.
	FiltersJson string
	// Name of the vector index to query.
	IndexName string
	// Number of results to return. Defaults to 10.
	NumResults int
	// Query text. Required for Delta Sync Index using model endpoint.
	QueryText string
	// The query type to use. Choices are `ANN` and `HYBRID`. Defaults to `ANN`.
	QueryType string
	// Query vector. Required for Direct Vector Access Index and Delta Sync
	// Index using self-managed vectors.
	QueryVector []float64
	// Threshold for the approximate nearest neighbor search. Defaults to 0.0.
	ScoreThreshold float64

	ForceSendFields []string
}

func queryVectorIndexRequestToPb(st *QueryVectorIndexRequest) (*queryVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryVectorIndexRequestPb{}

	var columnsPb []string
	for _, item := range st.Columns {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsPb = append(columnsPb, *itemPb)
		}
	}
	pb.Columns = columnsPb

	var columnsToRerankPb []string
	for _, item := range st.ColumnsToRerank {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsToRerankPb = append(columnsToRerankPb, *itemPb)
		}
	}
	pb.ColumnsToRerank = columnsToRerankPb

	filtersJsonPb, err := identity(&st.FiltersJson)
	if err != nil {
		return nil, err
	}
	if filtersJsonPb != nil {
		pb.FiltersJson = *filtersJsonPb
	}

	indexNamePb, err := identity(&st.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNamePb != nil {
		pb.IndexName = *indexNamePb
	}

	numResultsPb, err := identity(&st.NumResults)
	if err != nil {
		return nil, err
	}
	if numResultsPb != nil {
		pb.NumResults = *numResultsPb
	}

	queryTextPb, err := identity(&st.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextPb != nil {
		pb.QueryText = *queryTextPb
	}

	queryTypePb, err := identity(&st.QueryType)
	if err != nil {
		return nil, err
	}
	if queryTypePb != nil {
		pb.QueryType = *queryTypePb
	}

	var queryVectorPb []float64
	for _, item := range st.QueryVector {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			queryVectorPb = append(queryVectorPb, *itemPb)
		}
	}
	pb.QueryVector = queryVectorPb

	scoreThresholdPb, err := identity(&st.ScoreThreshold)
	if err != nil {
		return nil, err
	}
	if scoreThresholdPb != nil {
		pb.ScoreThreshold = *scoreThresholdPb
	}

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

	var columnsField []string
	for _, itemPb := range pb.Columns {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsField = append(columnsField, *item)
		}
	}
	st.Columns = columnsField

	var columnsToRerankField []string
	for _, itemPb := range pb.ColumnsToRerank {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsToRerankField = append(columnsToRerankField, *item)
		}
	}
	st.ColumnsToRerank = columnsToRerankField
	filtersJsonField, err := identity(&pb.FiltersJson)
	if err != nil {
		return nil, err
	}
	if filtersJsonField != nil {
		st.FiltersJson = *filtersJsonField
	}
	indexNameField, err := identity(&pb.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNameField != nil {
		st.IndexName = *indexNameField
	}
	numResultsField, err := identity(&pb.NumResults)
	if err != nil {
		return nil, err
	}
	if numResultsField != nil {
		st.NumResults = *numResultsField
	}
	queryTextField, err := identity(&pb.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextField != nil {
		st.QueryText = *queryTextField
	}
	queryTypeField, err := identity(&pb.QueryType)
	if err != nil {
		return nil, err
	}
	if queryTypeField != nil {
		st.QueryType = *queryTypeField
	}

	var queryVectorField []float64
	for _, itemPb := range pb.QueryVector {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			queryVectorField = append(queryVectorField, *item)
		}
	}
	st.QueryVector = queryVectorField
	scoreThresholdField, err := identity(&pb.ScoreThreshold)
	if err != nil {
		return nil, err
	}
	if scoreThresholdField != nil {
		st.ScoreThreshold = *scoreThresholdField
	}

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
	Manifest *ResultManifest
	// [Optional] Token that can be used in `QueryVectorIndexNextPage` API to
	// get next page of results. If more than 1000 results satisfy the query,
	// they are returned in groups of 1000. Empty value means no more results.
	// The maximum number of results that can be returned is 10,000.
	NextPageToken string
	// Data returned in the query result.
	Result *ResultData

	ForceSendFields []string
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

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}
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
	DataArray []ListValue
	// Number of rows in the result set.
	RowCount int

	ForceSendFields []string
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

	rowCountPb, err := identity(&st.RowCount)
	if err != nil {
		return nil, err
	}
	if rowCountPb != nil {
		pb.RowCount = *rowCountPb
	}

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
	rowCountField, err := identity(&pb.RowCount)
	if err != nil {
		return nil, err
	}
	if rowCountField != nil {
		st.RowCount = *rowCountField
	}

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
	ColumnCount int
	// Information about each column in the result set.
	Columns []ColumnInfo

	ForceSendFields []string
}

func resultManifestToPb(st *ResultManifest) (*resultManifestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultManifestPb{}
	columnCountPb, err := identity(&st.ColumnCount)
	if err != nil {
		return nil, err
	}
	if columnCountPb != nil {
		pb.ColumnCount = *columnCountPb
	}

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
	columnCountField, err := identity(&pb.ColumnCount)
	if err != nil {
		return nil, err
	}
	if columnCountField != nil {
		st.ColumnCount = *columnCountField
	}

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
	IndexName string
	// Primary key of the last entry returned in the previous scan.
	LastPrimaryKey string
	// Number of results to return. Defaults to 10.
	NumResults int

	ForceSendFields []string
}

func scanVectorIndexRequestToPb(st *ScanVectorIndexRequest) (*scanVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &scanVectorIndexRequestPb{}
	indexNamePb, err := identity(&st.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNamePb != nil {
		pb.IndexName = *indexNamePb
	}

	lastPrimaryKeyPb, err := identity(&st.LastPrimaryKey)
	if err != nil {
		return nil, err
	}
	if lastPrimaryKeyPb != nil {
		pb.LastPrimaryKey = *lastPrimaryKeyPb
	}

	numResultsPb, err := identity(&st.NumResults)
	if err != nil {
		return nil, err
	}
	if numResultsPb != nil {
		pb.NumResults = *numResultsPb
	}

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
	indexNameField, err := identity(&pb.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNameField != nil {
		st.IndexName = *indexNameField
	}
	lastPrimaryKeyField, err := identity(&pb.LastPrimaryKey)
	if err != nil {
		return nil, err
	}
	if lastPrimaryKeyField != nil {
		st.LastPrimaryKey = *lastPrimaryKeyField
	}
	numResultsField, err := identity(&pb.NumResults)
	if err != nil {
		return nil, err
	}
	if numResultsField != nil {
		st.NumResults = *numResultsField
	}

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
	Data []Struct
	// Primary key of the last entry.
	LastPrimaryKey string

	ForceSendFields []string
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

	lastPrimaryKeyPb, err := identity(&st.LastPrimaryKey)
	if err != nil {
		return nil, err
	}
	if lastPrimaryKeyPb != nil {
		pb.LastPrimaryKey = *lastPrimaryKeyPb
	}

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
	lastPrimaryKeyField, err := identity(&pb.LastPrimaryKey)
	if err != nil {
		return nil, err
	}
	if lastPrimaryKeyField != nil {
		st.LastPrimaryKey = *lastPrimaryKeyField
	}

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
	IndexName string
}

func syncIndexRequestToPb(st *SyncIndexRequest) (*syncIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &syncIndexRequestPb{}
	indexNamePb, err := identity(&st.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNamePb != nil {
		pb.IndexName = *indexNamePb
	}

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
	indexNameField, err := identity(&pb.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNameField != nil {
		st.IndexName = *indexNameField
	}

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
	CustomTags []CustomTag
	// Name of the vector search endpoint
	EndpointName string
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

	endpointNamePb, err := identity(&st.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNamePb != nil {
		pb.EndpointName = *endpointNamePb
	}

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
	endpointNameField, err := identity(&pb.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNameField != nil {
		st.EndpointName = *endpointNameField
	}

	return st, nil
}

type UpdateEndpointCustomTagsResponse struct {
	// All the custom tags that are applied to the vector search endpoint.
	CustomTags []CustomTag
	// The name of the vector search endpoint whose custom tags were updated.
	Name string

	ForceSendFields []string
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

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	FailedPrimaryKeys []string
	// Count of successfully processed rows.
	SuccessRowCount int64

	ForceSendFields []string
}

func upsertDataResultToPb(st *UpsertDataResult) (*upsertDataResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &upsertDataResultPb{}

	var failedPrimaryKeysPb []string
	for _, item := range st.FailedPrimaryKeys {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			failedPrimaryKeysPb = append(failedPrimaryKeysPb, *itemPb)
		}
	}
	pb.FailedPrimaryKeys = failedPrimaryKeysPb

	successRowCountPb, err := identity(&st.SuccessRowCount)
	if err != nil {
		return nil, err
	}
	if successRowCountPb != nil {
		pb.SuccessRowCount = *successRowCountPb
	}

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

	var failedPrimaryKeysField []string
	for _, itemPb := range pb.FailedPrimaryKeys {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			failedPrimaryKeysField = append(failedPrimaryKeysField, *item)
		}
	}
	st.FailedPrimaryKeys = failedPrimaryKeysField
	successRowCountField, err := identity(&pb.SuccessRowCount)
	if err != nil {
		return nil, err
	}
	if successRowCountField != nil {
		st.SuccessRowCount = *successRowCountField
	}

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
	IndexName string
	// JSON string representing the data to be upserted.
	InputsJson string
}

func upsertDataVectorIndexRequestToPb(st *UpsertDataVectorIndexRequest) (*upsertDataVectorIndexRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &upsertDataVectorIndexRequestPb{}
	indexNamePb, err := identity(&st.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNamePb != nil {
		pb.IndexName = *indexNamePb
	}

	inputsJsonPb, err := identity(&st.InputsJson)
	if err != nil {
		return nil, err
	}
	if inputsJsonPb != nil {
		pb.InputsJson = *inputsJsonPb
	}

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
	indexNameField, err := identity(&pb.IndexName)
	if err != nil {
		return nil, err
	}
	if indexNameField != nil {
		st.IndexName = *indexNameField
	}
	inputsJsonField, err := identity(&pb.InputsJson)
	if err != nil {
		return nil, err
	}
	if inputsJsonField != nil {
		st.InputsJson = *inputsJsonField
	}

	return st, nil
}

type UpsertDataVectorIndexResponse struct {
	// Result of the upsert or delete operation.
	Result *UpsertDataResult
	// Status of the upsert operation.
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

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

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
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	return st, nil
}

type Value struct {
	BoolValue bool
	// copied from proto3 / Google Well Known Types, source:
	// https://github.com/protocolbuffers/protobuf/blob/450d24ca820750c5db5112a6f0b0c2efb9758021/src/google/protobuf/struct.proto
	// `ListValue` is a wrapper around a repeated field of values.
	//
	// The JSON representation for `ListValue` is JSON array.
	ListValue *ListValue

	NumberValue float64

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
	StructValue *Struct

	ForceSendFields []string
}

func valueToPb(st *Value) (*valuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &valuePb{}
	boolValuePb, err := identity(&st.BoolValue)
	if err != nil {
		return nil, err
	}
	if boolValuePb != nil {
		pb.BoolValue = *boolValuePb
	}

	listValuePb, err := listValueToPb(st.ListValue)
	if err != nil {
		return nil, err
	}
	if listValuePb != nil {
		pb.ListValue = listValuePb
	}

	numberValuePb, err := identity(&st.NumberValue)
	if err != nil {
		return nil, err
	}
	if numberValuePb != nil {
		pb.NumberValue = *numberValuePb
	}

	stringValuePb, err := identity(&st.StringValue)
	if err != nil {
		return nil, err
	}
	if stringValuePb != nil {
		pb.StringValue = *stringValuePb
	}

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
	boolValueField, err := identity(&pb.BoolValue)
	if err != nil {
		return nil, err
	}
	if boolValueField != nil {
		st.BoolValue = *boolValueField
	}
	listValueField, err := listValueFromPb(pb.ListValue)
	if err != nil {
		return nil, err
	}
	if listValueField != nil {
		st.ListValue = listValueField
	}
	numberValueField, err := identity(&pb.NumberValue)
	if err != nil {
		return nil, err
	}
	if numberValueField != nil {
		st.NumberValue = *numberValueField
	}
	stringValueField, err := identity(&pb.StringValue)
	if err != nil {
		return nil, err
	}
	if stringValueField != nil {
		st.StringValue = *stringValueField
	}
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
	Creator string

	DeltaSyncIndexSpec *DeltaSyncVectorIndexSpecResponse

	DirectAccessIndexSpec *DirectAccessVectorIndexSpec
	// Name of the endpoint associated with the index
	EndpointName string
	// There are 2 types of Vector Search indexes: - `DELTA_SYNC`: An index that
	// automatically syncs with a source Delta Table, automatically and
	// incrementally updating the index as the underlying data in the Delta
	// Table changes. - `DIRECT_ACCESS`: An index that supports direct read and
	// write of vectors and metadata through our REST and SDK APIs. With this
	// model, the user manages index updates.
	IndexType VectorIndexType
	// Name of the index
	Name string
	// Primary key of the index
	PrimaryKey string

	Status *VectorIndexStatus

	ForceSendFields []string
}

func vectorIndexToPb(st *VectorIndex) (*vectorIndexPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorIndexPb{}
	creatorPb, err := identity(&st.Creator)
	if err != nil {
		return nil, err
	}
	if creatorPb != nil {
		pb.Creator = *creatorPb
	}

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

	endpointNamePb, err := identity(&st.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNamePb != nil {
		pb.EndpointName = *endpointNamePb
	}

	indexTypePb, err := identity(&st.IndexType)
	if err != nil {
		return nil, err
	}
	if indexTypePb != nil {
		pb.IndexType = *indexTypePb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	primaryKeyPb, err := identity(&st.PrimaryKey)
	if err != nil {
		return nil, err
	}
	if primaryKeyPb != nil {
		pb.PrimaryKey = *primaryKeyPb
	}

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
	creatorField, err := identity(&pb.Creator)
	if err != nil {
		return nil, err
	}
	if creatorField != nil {
		st.Creator = *creatorField
	}
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
	endpointNameField, err := identity(&pb.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNameField != nil {
		st.EndpointName = *endpointNameField
	}
	indexTypeField, err := identity(&pb.IndexType)
	if err != nil {
		return nil, err
	}
	if indexTypeField != nil {
		st.IndexType = *indexTypeField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	primaryKeyField, err := identity(&pb.PrimaryKey)
	if err != nil {
		return nil, err
	}
	if primaryKeyField != nil {
		st.PrimaryKey = *primaryKeyField
	}
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
	IndexUrl string
	// Number of rows indexed
	IndexedRowCount int64
	// Message associated with the index status
	Message string
	// Whether the index is ready for search
	Ready bool

	ForceSendFields []string
}

func vectorIndexStatusToPb(st *VectorIndexStatus) (*vectorIndexStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vectorIndexStatusPb{}
	indexUrlPb, err := identity(&st.IndexUrl)
	if err != nil {
		return nil, err
	}
	if indexUrlPb != nil {
		pb.IndexUrl = *indexUrlPb
	}

	indexedRowCountPb, err := identity(&st.IndexedRowCount)
	if err != nil {
		return nil, err
	}
	if indexedRowCountPb != nil {
		pb.IndexedRowCount = *indexedRowCountPb
	}

	messagePb, err := identity(&st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = *messagePb
	}

	readyPb, err := identity(&st.Ready)
	if err != nil {
		return nil, err
	}
	if readyPb != nil {
		pb.Ready = *readyPb
	}

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
	indexUrlField, err := identity(&pb.IndexUrl)
	if err != nil {
		return nil, err
	}
	if indexUrlField != nil {
		st.IndexUrl = *indexUrlField
	}
	indexedRowCountField, err := identity(&pb.IndexedRowCount)
	if err != nil {
		return nil, err
	}
	if indexedRowCountField != nil {
		st.IndexedRowCount = *indexedRowCountField
	}
	messageField, err := identity(&pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = *messageField
	}
	readyField, err := identity(&pb.Ready)
	if err != nil {
		return nil, err
	}
	if readyField != nil {
		st.Ready = *readyField
	}

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

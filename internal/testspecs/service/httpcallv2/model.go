// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package httpcallv2

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/internal/testspecs/service/httpcallv2/httpcallv2pb"
)

type ComplexQueryParam struct {

	// Wire name: 'nested_optional_query_param'
	NestedOptionalQueryParam string ``

	// Wire name: 'nested_repeated_query_param'
	NestedRepeatedQueryParam []string ``
	ForceSendFields          []string `tf:"-"`
}

func (st ComplexQueryParam) MarshalJSON() ([]byte, error) {
	pb, err := ComplexQueryParamToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ComplexQueryParam) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &httpcallv2pb.ComplexQueryParamPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ComplexQueryParamFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ComplexQueryParamToPb(st *ComplexQueryParam) (*httpcallv2pb.ComplexQueryParamPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &httpcallv2pb.ComplexQueryParamPb{}
	pb.NestedOptionalQueryParam = st.NestedOptionalQueryParam
	pb.NestedRepeatedQueryParam = st.NestedRepeatedQueryParam

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ComplexQueryParamFromPb(pb *httpcallv2pb.ComplexQueryParamPb) (*ComplexQueryParam, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComplexQueryParam{}
	st.NestedOptionalQueryParam = pb.NestedOptionalQueryParam
	st.NestedRepeatedQueryParam = pb.NestedRepeatedQueryParam

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// This mimics "old" style post requests which have the resource inlined.
type CreateResourceRequest struct {
	// Body element
	// Wire name: 'body_field'
	BodyField string ``

	// Wire name: 'path_param_bool'
	PathParamBool bool `tf:"-"`

	// Wire name: 'path_param_int'
	PathParamInt int `tf:"-"`

	// Wire name: 'path_param_string'
	PathParamString string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st CreateResourceRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateResourceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateResourceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &httpcallv2pb.CreateResourceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateResourceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateResourceRequestToPb(st *CreateResourceRequest) (*httpcallv2pb.CreateResourceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &httpcallv2pb.CreateResourceRequestPb{}
	pb.BodyField = st.BodyField
	pb.PathParamBool = st.PathParamBool
	pb.PathParamInt = st.PathParamInt
	pb.PathParamString = st.PathParamString

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateResourceRequestFromPb(pb *httpcallv2pb.CreateResourceRequestPb) (*CreateResourceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateResourceRequest{}
	st.BodyField = pb.BodyField
	st.PathParamBool = pb.PathParamBool
	st.PathParamInt = pb.PathParamInt
	st.PathParamString = pb.PathParamString

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetResourceRequest struct {
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	// Wire name: 'field_mask'
	FieldMask []string `tf:"-"`

	// Wire name: 'optional_complex_query_param'
	OptionalComplexQueryParam *ComplexQueryParam `tf:"-"`

	// Wire name: 'path_param_bool'
	PathParamBool bool `tf:"-"`

	// Wire name: 'path_param_int'
	PathParamInt int `tf:"-"`

	// Wire name: 'path_param_string'
	PathParamString string `tf:"-"`

	// Wire name: 'query_param_bool'
	QueryParamBool bool `tf:"-"`

	// Wire name: 'query_param_int'
	QueryParamInt int `tf:"-"`

	// Wire name: 'query_param_string'
	QueryParamString string `tf:"-"`

	// Wire name: 'repeated_complex_query_param'
	RepeatedComplexQueryParam []ComplexQueryParam `tf:"-"`

	// Wire name: 'repeated_query_param'
	RepeatedQueryParam []string `tf:"-"`
	ForceSendFields    []string `tf:"-"`
}

func (st GetResourceRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetResourceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetResourceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &httpcallv2pb.GetResourceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetResourceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetResourceRequestToPb(st *GetResourceRequest) (*httpcallv2pb.GetResourceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &httpcallv2pb.GetResourceRequestPb{}
	fieldMaskPb, err := fieldMaskToPb(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}
	optionalComplexQueryParamPb, err := ComplexQueryParamToPb(st.OptionalComplexQueryParam)
	if err != nil {
		return nil, err
	}
	if optionalComplexQueryParamPb != nil {
		pb.OptionalComplexQueryParam = optionalComplexQueryParamPb
	}
	pb.PathParamBool = st.PathParamBool
	pb.PathParamInt = st.PathParamInt
	pb.PathParamString = st.PathParamString
	pb.QueryParamBool = st.QueryParamBool
	pb.QueryParamInt = st.QueryParamInt
	pb.QueryParamString = st.QueryParamString

	var repeatedComplexQueryParamPb []httpcallv2pb.ComplexQueryParamPb
	for _, item := range st.RepeatedComplexQueryParam {
		itemPb, err := ComplexQueryParamToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repeatedComplexQueryParamPb = append(repeatedComplexQueryParamPb, *itemPb)
		}
	}
	pb.RepeatedComplexQueryParam = repeatedComplexQueryParamPb
	pb.RepeatedQueryParam = st.RepeatedQueryParam

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetResourceRequestFromPb(pb *httpcallv2pb.GetResourceRequestPb) (*GetResourceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetResourceRequest{}
	fieldMaskField, err := fieldMaskFromPb(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	optionalComplexQueryParamField, err := ComplexQueryParamFromPb(pb.OptionalComplexQueryParam)
	if err != nil {
		return nil, err
	}
	if optionalComplexQueryParamField != nil {
		st.OptionalComplexQueryParam = optionalComplexQueryParamField
	}
	st.PathParamBool = pb.PathParamBool
	st.PathParamInt = pb.PathParamInt
	st.PathParamString = pb.PathParamString
	st.QueryParamBool = pb.QueryParamBool
	st.QueryParamInt = pb.QueryParamInt
	st.QueryParamString = pb.QueryParamString

	var repeatedComplexQueryParamField []ComplexQueryParam
	for _, itemPb := range pb.RepeatedComplexQueryParam {
		item, err := ComplexQueryParamFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repeatedComplexQueryParamField = append(repeatedComplexQueryParamField, *item)
		}
	}
	st.RepeatedComplexQueryParam = repeatedComplexQueryParamField
	st.RepeatedQueryParam = pb.RepeatedQueryParam

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type Resource struct {

	// Wire name: 'body_field'
	BodyField string ``

	// Wire name: 'nested_path_param_bool'
	NestedPathParamBool bool ``

	// Wire name: 'nested_path_param_int'
	NestedPathParamInt int ``

	// Wire name: 'nested_path_param_string'
	NestedPathParamString string   ``
	ForceSendFields       []string `tf:"-"`
}

func (st Resource) MarshalJSON() ([]byte, error) {
	pb, err := ResourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Resource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &httpcallv2pb.ResourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResourceToPb(st *Resource) (*httpcallv2pb.ResourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &httpcallv2pb.ResourcePb{}
	pb.BodyField = st.BodyField
	pb.NestedPathParamBool = st.NestedPathParamBool
	pb.NestedPathParamInt = st.NestedPathParamInt
	pb.NestedPathParamString = st.NestedPathParamString

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ResourceFromPb(pb *httpcallv2pb.ResourcePb) (*Resource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Resource{}
	st.BodyField = pb.BodyField
	st.NestedPathParamBool = pb.NestedPathParamBool
	st.NestedPathParamInt = pb.NestedPathParamInt
	st.NestedPathParamString = pb.NestedPathParamString

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateResourceRequest struct {
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	// Wire name: 'field_mask'
	FieldMask []string `tf:"-"`

	// Wire name: 'nested_path_param_bool'
	NestedPathParamBool bool `tf:"-"`

	// Wire name: 'nested_path_param_int'
	NestedPathParamInt int `tf:"-"`

	// Wire name: 'nested_path_param_string'
	NestedPathParamString string `tf:"-"`

	// Wire name: 'optional_complex_query_param'
	OptionalComplexQueryParam *ComplexQueryParam `tf:"-"`

	// Wire name: 'query_param_bool'
	QueryParamBool bool `tf:"-"`

	// Wire name: 'query_param_int'
	QueryParamInt int `tf:"-"`

	// Wire name: 'query_param_string'
	QueryParamString string `tf:"-"`

	// Wire name: 'repeated_complex_query_param'
	RepeatedComplexQueryParam []ComplexQueryParam `tf:"-"`

	// Wire name: 'repeated_query_param'
	RepeatedQueryParam []string `tf:"-"`
	// Body element
	// Wire name: 'resource'
	Resource        Resource ``
	ForceSendFields []string `tf:"-"`
}

func (st UpdateResourceRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateResourceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateResourceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &httpcallv2pb.UpdateResourceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateResourceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateResourceRequestToPb(st *UpdateResourceRequest) (*httpcallv2pb.UpdateResourceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &httpcallv2pb.UpdateResourceRequestPb{}
	fieldMaskPb, err := fieldMaskToPb(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}
	pb.NestedPathParamBool = st.NestedPathParamBool
	pb.NestedPathParamInt = st.NestedPathParamInt
	pb.NestedPathParamString = st.NestedPathParamString
	optionalComplexQueryParamPb, err := ComplexQueryParamToPb(st.OptionalComplexQueryParam)
	if err != nil {
		return nil, err
	}
	if optionalComplexQueryParamPb != nil {
		pb.OptionalComplexQueryParam = optionalComplexQueryParamPb
	}
	pb.QueryParamBool = st.QueryParamBool
	pb.QueryParamInt = st.QueryParamInt
	pb.QueryParamString = st.QueryParamString

	var repeatedComplexQueryParamPb []httpcallv2pb.ComplexQueryParamPb
	for _, item := range st.RepeatedComplexQueryParam {
		itemPb, err := ComplexQueryParamToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repeatedComplexQueryParamPb = append(repeatedComplexQueryParamPb, *itemPb)
		}
	}
	pb.RepeatedComplexQueryParam = repeatedComplexQueryParamPb
	pb.RepeatedQueryParam = st.RepeatedQueryParam
	resourcePb, err := ResourceToPb(&st.Resource)
	if err != nil {
		return nil, err
	}
	if resourcePb != nil {
		pb.Resource = *resourcePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateResourceRequestFromPb(pb *httpcallv2pb.UpdateResourceRequestPb) (*UpdateResourceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateResourceRequest{}
	fieldMaskField, err := fieldMaskFromPb(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	st.NestedPathParamBool = pb.NestedPathParamBool
	st.NestedPathParamInt = pb.NestedPathParamInt
	st.NestedPathParamString = pb.NestedPathParamString
	optionalComplexQueryParamField, err := ComplexQueryParamFromPb(pb.OptionalComplexQueryParam)
	if err != nil {
		return nil, err
	}
	if optionalComplexQueryParamField != nil {
		st.OptionalComplexQueryParam = optionalComplexQueryParamField
	}
	st.QueryParamBool = pb.QueryParamBool
	st.QueryParamInt = pb.QueryParamInt
	st.QueryParamString = pb.QueryParamString

	var repeatedComplexQueryParamField []ComplexQueryParam
	for _, itemPb := range pb.RepeatedComplexQueryParam {
		item, err := ComplexQueryParamFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repeatedComplexQueryParamField = append(repeatedComplexQueryParamField, *item)
		}
	}
	st.RepeatedComplexQueryParam = repeatedComplexQueryParamField
	st.RepeatedQueryParam = pb.RepeatedQueryParam
	resourceField, err := ResourceFromPb(&pb.Resource)
	if err != nil {
		return nil, err
	}
	if resourceField != nil {
		st.Resource = *resourceField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

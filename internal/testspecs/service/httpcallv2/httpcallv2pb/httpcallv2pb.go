// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package httpcallv2pb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type ComplexQueryParamPb struct {
	NestedOptionalQueryParam string   `json:"nested_optional_query_param,omitempty" url:"nested_optional_query_param,omitempty"`
	NestedRepeatedQueryParam []string `json:"nested_repeated_query_param,omitempty" url:"nested_repeated_query_param,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ComplexQueryParamPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ComplexQueryParamPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateResourceRequestPb struct {
	BodyField       string `json:"body_field,omitempty"`
	PathParamBool   bool   `json:"-" url:"-"`
	PathParamInt    int    `json:"-" url:"-"`
	PathParamString string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateResourceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateResourceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetResourceRequestPb struct {
	FieldMask                 string                `json:"-" url:"field_mask,omitempty"`
	OptionalComplexQueryParam *ComplexQueryParamPb  `json:"-" url:"optional_complex_query_param,omitempty"`
	PathParamBool             bool                  `json:"-" url:"-"`
	PathParamInt              int                   `json:"-" url:"-"`
	PathParamString           string                `json:"-" url:"-"`
	QueryParamBool            bool                  `json:"-" url:"query_param_bool,omitempty"`
	QueryParamInt             int                   `json:"-" url:"query_param_int,omitempty"`
	QueryParamString          string                `json:"-" url:"query_param_string,omitempty"`
	RepeatedComplexQueryParam []ComplexQueryParamPb `json:"-" url:"repeated_complex_query_param,omitempty"`
	RepeatedQueryParam        []string              `json:"-" url:"repeated_query_param,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetResourceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetResourceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResourcePb struct {
	BodyField             string `json:"body_field,omitempty"`
	NestedPathParamBool   bool   `json:"nested_path_param_bool,omitempty"`
	NestedPathParamInt    int    `json:"nested_path_param_int,omitempty"`
	NestedPathParamString string `json:"nested_path_param_string,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ResourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ResourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateResourceRequestPb struct {
	FieldMask                 string                `json:"-" url:"field_mask,omitempty"`
	NestedPathParamBool       bool                  `json:"-" url:"-"`
	NestedPathParamInt        int                   `json:"-" url:"-"`
	NestedPathParamString     string                `json:"-" url:"-"`
	OptionalComplexQueryParam *ComplexQueryParamPb  `json:"-" url:"optional_complex_query_param,omitempty"`
	QueryParamBool            bool                  `json:"-" url:"query_param_bool,omitempty"`
	QueryParamInt             int                   `json:"-" url:"query_param_int,omitempty"`
	QueryParamString          string                `json:"-" url:"query_param_string,omitempty"`
	RepeatedComplexQueryParam []ComplexQueryParamPb `json:"-" url:"repeated_complex_query_param,omitempty"`
	RepeatedQueryParam        []string              `json:"-" url:"repeated_query_param,omitempty"`
	Resource                  ResourcePb            `json:"resource"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateResourceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateResourceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

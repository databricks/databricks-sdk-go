// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jsonmarshallv2pb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type GetResourceRequestPb struct {
	Name     string     `json:"-" url:"-"`
	Resource ResourcePb `json:"-" url:"resource"`
}

type NestedMessagePb struct {
	OptionalDuration  string `json:"optional_duration,omitempty" url:"optional_duration,omitempty"`
	OptionalString    string `json:"optional_string,omitempty" url:"optional_string,omitempty"`
	OptionalTimestamp string `json:"optional_timestamp,omitempty" url:"optional_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NestedMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NestedMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type OptionalFieldsPb struct {
	Duration        string            `json:"duration,omitempty" url:"duration,omitempty"`
	FieldMask       string            `json:"field_mask,omitempty" url:"field_mask,omitempty"`
	Map             map[string]string `json:"map,omitempty" url:"map,omitempty"`
	OptionalBool    bool              `json:"optional_bool,omitempty" url:"optional_bool,omitempty"`
	OptionalInt32   int               `json:"optional_int32,omitempty" url:"optional_int32,omitempty"`
	OptionalInt64   int64             `json:"optional_int64,omitempty" url:"optional_int64,omitempty"`
	OptionalMessage *NestedMessagePb  `json:"optional_message,omitempty" url:"optional_message,omitempty"`
	OptionalString  string            `json:"optional_string,omitempty" url:"optional_string,omitempty"`
	TestEnum        TestEnumPb        `json:"test_enum,omitempty" url:"test_enum,omitempty"`
	Timestamp       string            `json:"timestamp,omitempty" url:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *OptionalFieldsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st OptionalFieldsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepeatedFieldsPb struct {
	RepeatedBool      []bool            `json:"repeated_bool,omitempty" url:"repeated_bool,omitempty"`
	RepeatedDuration  []string          `json:"repeated_duration,omitempty" url:"repeated_duration,omitempty"`
	RepeatedFieldMask []string          `json:"repeated_field_mask,omitempty" url:"repeated_field_mask,omitempty"`
	RepeatedInt32     []int             `json:"repeated_int32,omitempty" url:"repeated_int32,omitempty"`
	RepeatedInt64     []int64           `json:"repeated_int64,omitempty" url:"repeated_int64,omitempty"`
	RepeatedMessage   []NestedMessagePb `json:"repeated_message,omitempty" url:"repeated_message,omitempty"`
	RepeatedString    []string          `json:"repeated_string,omitempty" url:"repeated_string,omitempty"`
	RepeatedTimestamp []string          `json:"repeated_timestamp,omitempty" url:"repeated_timestamp,omitempty"`
	TestRepeatedEnum  []TestEnumPb      `json:"test_repeated_enum,omitempty" url:"test_repeated_enum,omitempty"`
}

type RequiredFieldsPb struct {
	RequiredBool      bool            `json:"required_bool" url:"required_bool"`
	RequiredDuration  string          `json:"required_duration" url:"required_duration"`
	RequiredFieldMask string          `json:"required_field_mask" url:"required_field_mask"`
	RequiredInt32     int             `json:"required_int32" url:"required_int32"`
	RequiredInt64     int64           `json:"required_int64" url:"required_int64"`
	RequiredMessage   NestedMessagePb `json:"required_message" url:"required_message"`
	RequiredString    string          `json:"required_string" url:"required_string"`
	RequiredTimestamp string          `json:"required_timestamp" url:"required_timestamp"`
	TestRequiredEnum  TestEnumPb      `json:"test_required_enum" url:"test_required_enum"`
}

type ResourcePb struct {
	OptionalFields *OptionalFieldsPb `json:"optional_fields,omitempty" url:"optional_fields,omitempty"`
	RepeatedFields *RepeatedFieldsPb `json:"repeated_fields,omitempty" url:"repeated_fields,omitempty"`
	RequiredFields *RequiredFieldsPb `json:"required_fields,omitempty" url:"required_fields,omitempty"`
}

type TestEnumPb string

const TestEnumTestEnumOne TestEnumPb = `TEST_ENUM_ONE`

const TestEnumTestEnumTwo TestEnumPb = `TEST_ENUM_TWO`

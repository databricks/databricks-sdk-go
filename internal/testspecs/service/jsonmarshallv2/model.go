// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jsonmarshallv2

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/internal/testspecs/service/jsonmarshallv2/jsonmarshallv2pb"
)

type GetResourceRequest struct {
	Name string `json:"-" tf:"-"`
	// Description.
	Resource Resource `json:"-" tf:"-"`
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
	pb := &jsonmarshallv2pb.GetResourceRequestPb{}
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

func GetResourceRequestToPb(st *GetResourceRequest) (*jsonmarshallv2pb.GetResourceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jsonmarshallv2pb.GetResourceRequestPb{}
	pb.Name = st.Name
	resourcePb, err := ResourceToPb(&st.Resource)
	if err != nil {
		return nil, err
	}
	if resourcePb != nil {
		pb.Resource = *resourcePb
	}

	return pb, nil
}

func GetResourceRequestFromPb(pb *jsonmarshallv2pb.GetResourceRequestPb) (*GetResourceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetResourceRequest{}
	st.Name = pb.Name
	resourceField, err := ResourceFromPb(&pb.Resource)
	if err != nil {
		return nil, err
	}
	if resourceField != nil {
		st.Resource = *resourceField
	}

	return st, nil
}

type NestedMessage struct {

	// Wire name: 'optional_duration'
	OptionalDuration string `json:"optional_duration,omitempty"` //legacy

	// Wire name: 'optional_string'
	OptionalString string `json:"optional_string,omitempty"`

	// Wire name: 'optional_timestamp'
	OptionalTimestamp string   `json:"optional_timestamp,omitempty"` //legacy
	ForceSendFields   []string `json:"-" tf:"-"`
}

func (st NestedMessage) MarshalJSON() ([]byte, error) {
	pb, err := NestedMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NestedMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jsonmarshallv2pb.NestedMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NestedMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NestedMessageToPb(st *NestedMessage) (*jsonmarshallv2pb.NestedMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jsonmarshallv2pb.NestedMessagePb{}
	pb.OptionalDuration = st.OptionalDuration
	pb.OptionalString = st.OptionalString
	pb.OptionalTimestamp = st.OptionalTimestamp

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NestedMessageFromPb(pb *jsonmarshallv2pb.NestedMessagePb) (*NestedMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NestedMessage{}
	st.OptionalDuration = pb.OptionalDuration
	st.OptionalString = pb.OptionalString
	st.OptionalTimestamp = pb.OptionalTimestamp

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type OptionalFields struct {

	// Wire name: 'duration'
	Duration string `json:"duration,omitempty"` //legacy
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask,omitempty"` //legacy
	// Lint disable reason: This is a dummy field used to test SDK Generation
	// logic.
	// Wire name: 'map'
	Map map[string]string `json:"map,omitempty"`

	// Wire name: 'optional_bool'
	OptionalBool bool `json:"optional_bool,omitempty"`

	// Wire name: 'optional_int32'
	OptionalInt32 int `json:"optional_int32,omitempty"`

	// Wire name: 'optional_int64'
	OptionalInt64 int64 `json:"optional_int64,omitempty"`

	// Wire name: 'optional_message'
	OptionalMessage *NestedMessage `json:"optional_message,omitempty"`

	// Wire name: 'optional_string'
	OptionalString string `json:"optional_string,omitempty"`

	// Wire name: 'test_enum'
	TestEnum TestEnum `json:"test_enum,omitempty"`

	// Wire name: 'timestamp'
	Timestamp       string   `json:"timestamp,omitempty"` //legacy
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st OptionalFields) MarshalJSON() ([]byte, error) {
	pb, err := OptionalFieldsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *OptionalFields) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jsonmarshallv2pb.OptionalFieldsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := OptionalFieldsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func OptionalFieldsToPb(st *OptionalFields) (*jsonmarshallv2pb.OptionalFieldsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jsonmarshallv2pb.OptionalFieldsPb{}
	pb.Duration = st.Duration
	pb.FieldMask = st.FieldMask
	pb.Map = st.Map
	pb.OptionalBool = st.OptionalBool
	pb.OptionalInt32 = st.OptionalInt32
	pb.OptionalInt64 = st.OptionalInt64
	optionalMessagePb, err := NestedMessageToPb(st.OptionalMessage)
	if err != nil {
		return nil, err
	}
	if optionalMessagePb != nil {
		pb.OptionalMessage = optionalMessagePb
	}
	pb.OptionalString = st.OptionalString
	testEnumPb, err := TestEnumToPb(&st.TestEnum)
	if err != nil {
		return nil, err
	}
	if testEnumPb != nil {
		pb.TestEnum = *testEnumPb
	}
	pb.Timestamp = st.Timestamp

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func OptionalFieldsFromPb(pb *jsonmarshallv2pb.OptionalFieldsPb) (*OptionalFields, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OptionalFields{}
	st.Duration = pb.Duration
	st.FieldMask = pb.FieldMask
	st.Map = pb.Map
	st.OptionalBool = pb.OptionalBool
	st.OptionalInt32 = pb.OptionalInt32
	st.OptionalInt64 = pb.OptionalInt64
	optionalMessageField, err := NestedMessageFromPb(pb.OptionalMessage)
	if err != nil {
		return nil, err
	}
	if optionalMessageField != nil {
		st.OptionalMessage = optionalMessageField
	}
	st.OptionalString = pb.OptionalString
	testEnumField, err := TestEnumFromPb(&pb.TestEnum)
	if err != nil {
		return nil, err
	}
	if testEnumField != nil {
		st.TestEnum = *testEnumField
	}
	st.Timestamp = pb.Timestamp

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RepeatedFields struct {

	// Wire name: 'repeated_bool'
	RepeatedBool []bool `json:"repeated_bool,omitempty"`

	// Wire name: 'repeated_duration'
	RepeatedDuration []time.Duration `json:"repeated_duration,omitempty"`

	// Wire name: 'repeated_field_mask'
	RepeatedFieldMask [][]string `json:"repeated_field_mask,omitempty"`

	// Wire name: 'repeated_int32'
	RepeatedInt32 []int `json:"repeated_int32,omitempty"`

	// Wire name: 'repeated_int64'
	RepeatedInt64 []int64 `json:"repeated_int64,omitempty"`

	// Wire name: 'repeated_message'
	RepeatedMessage []NestedMessage `json:"repeated_message,omitempty"`

	// Wire name: 'repeated_string'
	RepeatedString []string `json:"repeated_string,omitempty"`

	// Wire name: 'repeated_timestamp'
	RepeatedTimestamp []time.Time `json:"repeated_timestamp,omitempty"`

	// Wire name: 'test_repeated_enum'
	TestRepeatedEnum []TestEnum `json:"test_repeated_enum,omitempty"`
}

func (st RepeatedFields) MarshalJSON() ([]byte, error) {
	pb, err := RepeatedFieldsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RepeatedFields) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jsonmarshallv2pb.RepeatedFieldsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RepeatedFieldsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RepeatedFieldsToPb(st *RepeatedFields) (*jsonmarshallv2pb.RepeatedFieldsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jsonmarshallv2pb.RepeatedFieldsPb{}
	pb.RepeatedBool = st.RepeatedBool

	var repeatedDurationPb []string
	for _, item := range st.RepeatedDuration {
		itemPb, err := durationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repeatedDurationPb = append(repeatedDurationPb, *itemPb)
		}
	}
	pb.RepeatedDuration = repeatedDurationPb

	var repeatedFieldMaskPb []string
	for _, item := range st.RepeatedFieldMask {
		itemPb, err := fieldMaskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repeatedFieldMaskPb = append(repeatedFieldMaskPb, *itemPb)
		}
	}
	pb.RepeatedFieldMask = repeatedFieldMaskPb
	pb.RepeatedInt32 = st.RepeatedInt32
	pb.RepeatedInt64 = st.RepeatedInt64

	var repeatedMessagePb []jsonmarshallv2pb.NestedMessagePb
	for _, item := range st.RepeatedMessage {
		itemPb, err := NestedMessageToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repeatedMessagePb = append(repeatedMessagePb, *itemPb)
		}
	}
	pb.RepeatedMessage = repeatedMessagePb
	pb.RepeatedString = st.RepeatedString

	var repeatedTimestampPb []string
	for _, item := range st.RepeatedTimestamp {
		itemPb, err := timestampToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repeatedTimestampPb = append(repeatedTimestampPb, *itemPb)
		}
	}
	pb.RepeatedTimestamp = repeatedTimestampPb

	var testRepeatedEnumPb []jsonmarshallv2pb.TestEnumPb
	for _, item := range st.TestRepeatedEnum {
		itemPb, err := TestEnumToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			testRepeatedEnumPb = append(testRepeatedEnumPb, *itemPb)
		}
	}
	pb.TestRepeatedEnum = testRepeatedEnumPb

	return pb, nil
}

func RepeatedFieldsFromPb(pb *jsonmarshallv2pb.RepeatedFieldsPb) (*RepeatedFields, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepeatedFields{}
	st.RepeatedBool = pb.RepeatedBool

	var repeatedDurationField []time.Duration
	for _, itemPb := range pb.RepeatedDuration {
		item, err := durationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repeatedDurationField = append(repeatedDurationField, *item)
		}
	}
	st.RepeatedDuration = repeatedDurationField

	var repeatedFieldMaskField [][]string
	for _, itemPb := range pb.RepeatedFieldMask {
		item, err := fieldMaskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repeatedFieldMaskField = append(repeatedFieldMaskField, *item)
		}
	}
	st.RepeatedFieldMask = repeatedFieldMaskField
	st.RepeatedInt32 = pb.RepeatedInt32
	st.RepeatedInt64 = pb.RepeatedInt64

	var repeatedMessageField []NestedMessage
	for _, itemPb := range pb.RepeatedMessage {
		item, err := NestedMessageFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repeatedMessageField = append(repeatedMessageField, *item)
		}
	}
	st.RepeatedMessage = repeatedMessageField
	st.RepeatedString = pb.RepeatedString

	var repeatedTimestampField []time.Time
	for _, itemPb := range pb.RepeatedTimestamp {
		item, err := timestampFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repeatedTimestampField = append(repeatedTimestampField, *item)
		}
	}
	st.RepeatedTimestamp = repeatedTimestampField

	var testRepeatedEnumField []TestEnum
	for _, itemPb := range pb.TestRepeatedEnum {
		item, err := TestEnumFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			testRepeatedEnumField = append(testRepeatedEnumField, *item)
		}
	}
	st.TestRepeatedEnum = testRepeatedEnumField

	return st, nil
}

type RequiredFields struct {

	// Wire name: 'required_bool'
	RequiredBool bool `json:"required_bool"`

	// Wire name: 'required_duration'
	RequiredDuration string `json:"required_duration"` //legacy
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	// Wire name: 'required_field_mask'
	RequiredFieldMask string `json:"required_field_mask"` //legacy

	// Wire name: 'required_int32'
	RequiredInt32 int `json:"required_int32"`

	// Wire name: 'required_int64'
	RequiredInt64 int64 `json:"required_int64"`

	// Wire name: 'required_message'
	RequiredMessage NestedMessage `json:"required_message"`

	// Wire name: 'required_string'
	RequiredString string `json:"required_string"`

	// Wire name: 'required_timestamp'
	RequiredTimestamp string `json:"required_timestamp"` //legacy

	// Wire name: 'test_required_enum'
	TestRequiredEnum TestEnum `json:"test_required_enum"`
}

func (st RequiredFields) MarshalJSON() ([]byte, error) {
	pb, err := RequiredFieldsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RequiredFields) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jsonmarshallv2pb.RequiredFieldsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RequiredFieldsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RequiredFieldsToPb(st *RequiredFields) (*jsonmarshallv2pb.RequiredFieldsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jsonmarshallv2pb.RequiredFieldsPb{}
	pb.RequiredBool = st.RequiredBool
	pb.RequiredDuration = st.RequiredDuration
	pb.RequiredFieldMask = st.RequiredFieldMask
	pb.RequiredInt32 = st.RequiredInt32
	pb.RequiredInt64 = st.RequiredInt64
	requiredMessagePb, err := NestedMessageToPb(&st.RequiredMessage)
	if err != nil {
		return nil, err
	}
	if requiredMessagePb != nil {
		pb.RequiredMessage = *requiredMessagePb
	}
	pb.RequiredString = st.RequiredString
	pb.RequiredTimestamp = st.RequiredTimestamp
	testRequiredEnumPb, err := TestEnumToPb(&st.TestRequiredEnum)
	if err != nil {
		return nil, err
	}
	if testRequiredEnumPb != nil {
		pb.TestRequiredEnum = *testRequiredEnumPb
	}

	return pb, nil
}

func RequiredFieldsFromPb(pb *jsonmarshallv2pb.RequiredFieldsPb) (*RequiredFields, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RequiredFields{}
	st.RequiredBool = pb.RequiredBool
	st.RequiredDuration = pb.RequiredDuration
	st.RequiredFieldMask = pb.RequiredFieldMask
	st.RequiredInt32 = pb.RequiredInt32
	st.RequiredInt64 = pb.RequiredInt64
	requiredMessageField, err := NestedMessageFromPb(&pb.RequiredMessage)
	if err != nil {
		return nil, err
	}
	if requiredMessageField != nil {
		st.RequiredMessage = *requiredMessageField
	}
	st.RequiredString = pb.RequiredString
	st.RequiredTimestamp = pb.RequiredTimestamp
	testRequiredEnumField, err := TestEnumFromPb(&pb.TestRequiredEnum)
	if err != nil {
		return nil, err
	}
	if testRequiredEnumField != nil {
		st.TestRequiredEnum = *testRequiredEnumField
	}

	return st, nil
}

// We separate this into 3 submessages to simplify test cases. E.g., any
// required top level field needs to be included in the expected json for each
// test case.
type Resource struct {

	// Wire name: 'optional_fields'
	OptionalFields *OptionalFields `json:"optional_fields,omitempty"`

	// Wire name: 'repeated_fields'
	RepeatedFields *RepeatedFields `json:"repeated_fields,omitempty"`

	// Wire name: 'required_fields'
	RequiredFields *RequiredFields `json:"required_fields,omitempty"`
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
	pb := &jsonmarshallv2pb.ResourcePb{}
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

func ResourceToPb(st *Resource) (*jsonmarshallv2pb.ResourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jsonmarshallv2pb.ResourcePb{}
	optionalFieldsPb, err := OptionalFieldsToPb(st.OptionalFields)
	if err != nil {
		return nil, err
	}
	if optionalFieldsPb != nil {
		pb.OptionalFields = optionalFieldsPb
	}
	repeatedFieldsPb, err := RepeatedFieldsToPb(st.RepeatedFields)
	if err != nil {
		return nil, err
	}
	if repeatedFieldsPb != nil {
		pb.RepeatedFields = repeatedFieldsPb
	}
	requiredFieldsPb, err := RequiredFieldsToPb(st.RequiredFields)
	if err != nil {
		return nil, err
	}
	if requiredFieldsPb != nil {
		pb.RequiredFields = requiredFieldsPb
	}

	return pb, nil
}

func ResourceFromPb(pb *jsonmarshallv2pb.ResourcePb) (*Resource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Resource{}
	optionalFieldsField, err := OptionalFieldsFromPb(pb.OptionalFields)
	if err != nil {
		return nil, err
	}
	if optionalFieldsField != nil {
		st.OptionalFields = optionalFieldsField
	}
	repeatedFieldsField, err := RepeatedFieldsFromPb(pb.RepeatedFields)
	if err != nil {
		return nil, err
	}
	if repeatedFieldsField != nil {
		st.RepeatedFields = repeatedFieldsField
	}
	requiredFieldsField, err := RequiredFieldsFromPb(pb.RequiredFields)
	if err != nil {
		return nil, err
	}
	if requiredFieldsField != nil {
		st.RequiredFields = requiredFieldsField
	}

	return st, nil
}

type TestEnum string

const TestEnumTestEnumOne TestEnum = `TEST_ENUM_ONE`

const TestEnumTestEnumTwo TestEnum = `TEST_ENUM_TWO`

// String representation for [fmt.Print]
func (f *TestEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TestEnum) Set(v string) error {
	switch v {
	case `TEST_ENUM_ONE`, `TEST_ENUM_TWO`:
		*f = TestEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TEST_ENUM_ONE", "TEST_ENUM_TWO"`, v)
	}
}

// Values returns all possible values for TestEnum.
//
// There is no guarantee on the order of the values in the slice.
func (f *TestEnum) Values() []TestEnum {
	return []TestEnum{
		TestEnumTestEnumOne,
		TestEnumTestEnumTwo,
	}
}

// Type always returns TestEnum to satisfy [pflag.Value] interface
func (f *TestEnum) Type() string {
	return "TestEnum"
}

func TestEnumToPb(st *TestEnum) (*jsonmarshallv2pb.TestEnumPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jsonmarshallv2pb.TestEnumPb(*st)
	return &pb, nil
}

func TestEnumFromPb(pb *jsonmarshallv2pb.TestEnumPb) (*TestEnum, error) {
	if pb == nil {
		return nil, nil
	}
	st := TestEnum(*pb)
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

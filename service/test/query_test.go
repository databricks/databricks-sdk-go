package test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func TestQueryValues(t *testing.T) {
	d := time.Duration(1000000000)
	time := time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC)
	st := RequestPb{
		AllWellKnown: AllWellKnown{
			Duration:  &d,
			Timestamp: &time,
			FieldMask: &[]string{"test", "test2"},
		},
		AllWellKnownP: &AllWellKnown{
			Duration:  &d,
			Timestamp: &time,
			FieldMask: &[]string{"test3", "test4"},
		},
	}
	vals, err := query.Values(&st)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, vals.Get("all_well_known[duration_pb]"), "1.000000000s")
	assert.Equal(t, vals.Get("all_well_known[timestamp_pb]"), "2021-01-01T01:01:01Z")
	assert.Equal(t, vals.Get("all_well_known[field_mask_pb]"), "test,test2")
	assert.Equal(t, vals.Get("all_well_known_p[duration_pb]"), "1.000000000s")
	assert.Equal(t, vals.Get("all_well_known_p[timestamp_pb]"), "2021-01-01T01:01:01Z")
	assert.Equal(t, vals.Get("all_well_known_p[field_mask_pb]"), "test3,test4")
}

type RequestPb struct {
	AllWellKnown  AllWellKnown  `url:"all_well_known"`
	AllWellKnownP *AllWellKnown `url:"all_well_known_p"`
}

// Copied from jobs/model.go
// and added the url tag to the fields

type AllWellKnown struct {

	// Wire name: 'duration'
	Duration *time.Duration `json:"duration,omitempty"`
	// optional google.protobuf.Struct struct = 3;
	//
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	// Wire name: 'field_mask'
	FieldMask *[]string `json:"field_mask,omitempty"`

	// Wire name: 'list_value'
	ListValue []json.RawMessage `json:"list_value,omitempty"`

	// Wire name: 'repeated_duration'
	RepeatedDuration []time.Duration `json:"repeated_duration,omitempty"`
	// repeated google.protobuf.Struct repeated_struct = 13;
	// Wire name: 'repeated_field_mask'
	RepeatedFieldMask [][]string `json:"repeated_field_mask,omitempty"`

	// Wire name: 'repeated_list_value'
	RepeatedListValue [][]json.RawMessage `json:"repeated_list_value,omitempty"`

	// Wire name: 'repeated_timestamp'
	RepeatedTimestamp []time.Time `json:"repeated_timestamp,omitempty"`

	// Wire name: 'repeated_value'
	RepeatedValue []json.RawMessage `json:"repeated_value,omitempty"`

	// Wire name: 'required_duration'
	RequiredDuration time.Duration `json:"required_duration"`
	// optional google.protobuf.Struct required_struct = 8 [ (validate_required)
	// = true ];
	//
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	// Wire name: 'required_field_mask'
	RequiredFieldMask []string `json:"required_field_mask"`

	// Wire name: 'required_list_value'
	RequiredListValue []json.RawMessage `json:"required_list_value"`

	// Wire name: 'required_timestamp'
	RequiredTimestamp time.Time `json:"required_timestamp"`

	// Wire name: 'required_value'
	RequiredValue json.RawMessage `json:"required_value"`

	// Wire name: 'timestamp'
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// Wire name: 'value'
	Value json.RawMessage `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AllWellKnown) EncodeValues(key string, v *url.Values) error {
	pb, err := allWellKnownToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AllWellKnown) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &allWellKnownPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := allWellKnownFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AllWellKnown) MarshalJSON() ([]byte, error) {
	pb, err := allWellKnownToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func allWellKnownToPb(st *AllWellKnown) (*allWellKnownPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &allWellKnownPb{}
	durationPb, err := durationToPb(st.Duration)
	if err != nil {
		return nil, err
	}
	if durationPb != nil {
		pb.Duration = *durationPb
	}
	fieldMaskPb, err := fieldMaskToPb(st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}
	pb.ListValue = st.ListValue

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
	pb.RepeatedListValue = st.RepeatedListValue

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
	pb.RepeatedValue = st.RepeatedValue
	requiredDurationPb, err := durationToPb(&st.RequiredDuration)
	if err != nil {
		return nil, err
	}
	if requiredDurationPb != nil {
		pb.RequiredDuration = *requiredDurationPb
	}
	requiredFieldMaskPb, err := fieldMaskToPb(&st.RequiredFieldMask)
	if err != nil {
		return nil, err
	}
	if requiredFieldMaskPb != nil {
		pb.RequiredFieldMask = *requiredFieldMaskPb
	}
	pb.RequiredListValue = st.RequiredListValue
	requiredTimestampPb, err := timestampToPb(&st.RequiredTimestamp)
	if err != nil {
		return nil, err
	}
	if requiredTimestampPb != nil {
		pb.RequiredTimestamp = *requiredTimestampPb
	}
	pb.RequiredValue = st.RequiredValue
	timestampPb, err := timestampToPb(st.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampPb != nil {
		pb.Timestamp = *timestampPb
	}
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type allWellKnownPb struct {
	Duration          string              `json:"duration,omitempty" url:"duration_pb,omitempty"`
	FieldMask         string              `json:"field_mask,omitempty" url:"field_mask_pb,omitempty"`
	ListValue         []json.RawMessage   `json:"list_value,omitempty"`
	RepeatedDuration  []string            `json:"repeated_duration,omitempty"`
	RepeatedFieldMask []string            `json:"repeated_field_mask,omitempty"`
	RepeatedListValue [][]json.RawMessage `json:"repeated_list_value,omitempty"`
	RepeatedTimestamp []string            `json:"repeated_timestamp,omitempty"`
	RepeatedValue     []json.RawMessage   `json:"repeated_value,omitempty"`
	RequiredDuration  string              `json:"required_duration"`
	RequiredFieldMask string              `json:"required_field_mask"`
	RequiredListValue []json.RawMessage   `json:"required_list_value"`
	RequiredTimestamp string              `json:"required_timestamp"`
	RequiredValue     json.RawMessage     `json:"required_value"`
	Timestamp         string              `json:"timestamp,omitempty" url:"timestamp_pb,omitempty"`
	Value             json.RawMessage     `json:"value,omitempty" url:"value_pb,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func allWellKnownFromPb(pb *allWellKnownPb) (*AllWellKnown, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AllWellKnown{}
	durationField, err := durationFromPb(&pb.Duration)
	if err != nil {
		return nil, err
	}
	if durationField != nil {
		st.Duration = durationField
	}
	fieldMaskField, err := fieldMaskFromPb(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = fieldMaskField
	}
	st.ListValue = pb.ListValue

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
	st.RepeatedListValue = pb.RepeatedListValue

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
	st.RepeatedValue = pb.RepeatedValue
	requiredDurationField, err := durationFromPb(&pb.RequiredDuration)
	if err != nil {
		return nil, err
	}
	if requiredDurationField != nil {
		st.RequiredDuration = *requiredDurationField
	}
	requiredFieldMaskField, err := fieldMaskFromPb(&pb.RequiredFieldMask)
	if err != nil {
		return nil, err
	}
	if requiredFieldMaskField != nil {
		st.RequiredFieldMask = *requiredFieldMaskField
	}
	st.RequiredListValue = pb.RequiredListValue
	requiredTimestampField, err := timestampFromPb(&pb.RequiredTimestamp)
	if err != nil {
		return nil, err
	}
	if requiredTimestampField != nil {
		st.RequiredTimestamp = *requiredTimestampField
	}
	st.RequiredValue = pb.RequiredValue
	timestampField, err := timestampFromPb(&pb.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampField != nil {
		st.Timestamp = timestampField
	}
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *allWellKnownPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st allWellKnownPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

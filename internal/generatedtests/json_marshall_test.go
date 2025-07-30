// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package generated_tests

import (
	"encoding/json"
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/testspecs/service/jsonmarshallv2"
	"github.com/google/go-cmp/cmp"
)

func TestJsonMarshall_OptionalString(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{
		OptionalString: "test",
	}
	expected := "{\"optional_string\":\"test\"}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_OptionalInt32(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{
		OptionalInt32: 42,
	}
	expected := "{\"optional_int32\":42}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_OptionalInt64(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{
		OptionalInt64: 9223372036854775807,
	}
	expected := "{\"optional_int64\":9223372036854775807}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_OptionalBool(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{
		OptionalBool: true,
	}
	expected := "{\"optional_bool\":true}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_OptionalEnum(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{
		TestEnum: jsonmarshallv2.TestEnumTestEnumOne,
	}
	expected := "{\"test_enum\":\"TEST_ENUM_ONE\"}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_OptionalNestedMessage(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{
		OptionalMessage: &jsonmarshallv2.NestedMessage{
			OptionalString: "nested_value",
		},
	}
	expected := "{\"optional_message\":{\"optional_string\":\"nested_value\"}}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_OptionalMap(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{
		Map: map[string]string{"key": "test_key", "value": "test_value"},
	}
	expected := "{\"map\":{\"key\":\"test_key\",\"value\":\"test_value\"}}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_OptionalDuration(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{
		Duration: "3600s",
	}
	expected := "{\"duration\":\"3600s\"}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_OptionalFieldMask(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{
		FieldMask: "optional_string,optional_int32",
	}
	expected := "{\"field_mask\":\"optional_string,optional_int32\"}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_OptionalTimestamp(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{
		Timestamp: "2023-01-01T00:00:00Z",
	}
	expected := "{\"timestamp\":\"2023-01-01T00:00:00Z\"}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_MultipleOptionalFields(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{
		OptionalBool:   true,
		OptionalInt32:  42,
		OptionalString: "test",
	}
	expected := "{\"optional_string\":\"test\",\"optional_int32\":42,\"optional_bool\":true}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RequiredFieldsNoInput(t *testing.T) {
	input := jsonmarshallv2.RequiredFields{}
	expected := "{\"required_string\":\"\",\"required_int32\":0,\"required_int64\":0,\"required_bool\":false,\"required_message\":{},\"test_required_enum\":\"\",\"required_duration\":\"\",\"required_field_mask\":\"\",\"required_timestamp\":\"\"}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RequiredFieldsExplicitDefaults(t *testing.T) {
	input := jsonmarshallv2.RequiredFields{
		RequiredBool:     false,
		RequiredInt32:    0,
		RequiredInt64:    0,
		RequiredString:   "",
		TestRequiredEnum: jsonmarshallv2.TestEnumTestEnumOne,
	}
	expected := "{\"required_string\":\"\",\"required_int32\":0,\"required_int64\":0,\"required_bool\":false,\"required_message\":{},\"test_required_enum\":\"TEST_ENUM_ONE\",\"required_duration\":\"\",\"required_field_mask\":\"\",\"required_timestamp\":\"\"}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RequiredFieldsNonDefaults(t *testing.T) {
	input := jsonmarshallv2.RequiredFields{
		RequiredBool:      true,
		RequiredDuration:  "7200s",
		RequiredFieldMask: "required_string,required_int32",
		RequiredInt32:     42,
		RequiredInt64:     1234567890123456789,
		RequiredString:    "non_default_string",
		RequiredTimestamp: "2023-12-31T23:59:59Z",
		TestRequiredEnum:  jsonmarshallv2.TestEnumTestEnumTwo,
	}
	expected := "{\"required_string\":\"non_default_string\",\"required_int32\":42,\"required_int64\":1234567890123456789,\"required_bool\":true,\"required_message\":{},\"test_required_enum\":\"TEST_ENUM_TWO\",\"required_duration\":\"7200s\",\"required_field_mask\":\"required_string,required_int32\",\"required_timestamp\":\"2023-12-31T23:59:59Z\"}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RequiredFieldsWithNestedMessage(t *testing.T) {
	input := jsonmarshallv2.RequiredFields{
		RequiredMessage: jsonmarshallv2.NestedMessage{
			OptionalString: "nested_value",
		},
	}
	expected := "{\"required_string\":\"\",\"required_int32\":0,\"required_int64\":0,\"required_bool\":false,\"required_message\":{\"optional_string\":\"nested_value\"},\"test_required_enum\":\"\",\"required_duration\":\"\",\"required_field_mask\":\"\",\"required_timestamp\":\"\"}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RepeatedString(t *testing.T) {
	input := jsonmarshallv2.RepeatedFields{
		RepeatedString: []string{"item1", "item2", "item3"},
	}
	expected := "{\"repeated_string\":[\"item1\",\"item2\",\"item3\"]}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RepeatedInt32(t *testing.T) {
	input := jsonmarshallv2.RepeatedFields{
		RepeatedInt32: []int{1, 2, 3, 4, 5},
	}
	expected := "{\"repeated_int32\":[1,2,3,4,5]}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RepeatedInt64(t *testing.T) {
	input := jsonmarshallv2.RepeatedFields{
		RepeatedInt64: []int64{1000000000000000000, 2000000000000000000},
	}
	expected := "{\"repeated_int64\":[1000000000000000000,2000000000000000000]}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RepeatedBool(t *testing.T) {
	input := jsonmarshallv2.RepeatedFields{
		RepeatedBool: []bool{true, false, true},
	}
	expected := "{\"repeated_bool\":[true,false,true]}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RepeatedEnum(t *testing.T) {
	input := jsonmarshallv2.RepeatedFields{
		TestRepeatedEnum: []jsonmarshallv2.TestEnum{jsonmarshallv2.TestEnumTestEnumOne, jsonmarshallv2.TestEnumTestEnumTwo},
	}
	expected := "{\"test_repeated_enum\":[\"TEST_ENUM_ONE\",\"TEST_ENUM_TWO\"]}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RepeatedNestedMessage(t *testing.T) {
	input := jsonmarshallv2.RepeatedFields{
		RepeatedMessage: []jsonmarshallv2.NestedMessage{jsonmarshallv2.NestedMessage{
			OptionalString: "nested1",
		}, jsonmarshallv2.NestedMessage{
			OptionalString: "nested2",
		}},
	}
	expected := "{\"repeated_message\":[{\"optional_string\":\"nested1\"},{\"optional_string\":\"nested2\"}]}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RepeatedDuration(t *testing.T) {
	input := jsonmarshallv2.RepeatedFields{
		RepeatedDuration: []string{"60s", "120s", "180s"},
	}
	expected := "{\"repeated_duration\":[\"60s\",\"120s\",\"180s\"]}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RepeatedFieldMask(t *testing.T) {
	input := jsonmarshallv2.RepeatedFields{
		RepeatedFieldMask: []string{"field1", "field2,field3"},
	}
	expected := "{\"repeated_field_mask\":[\"field1\",\"field2,field3\"]}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_RepeatedTimestamp(t *testing.T) {
	input := jsonmarshallv2.RepeatedFields{
		RepeatedTimestamp: []string{"2023-01-01T00:00:00Z", "2023-01-02T00:00:00Z"},
	}
	expected := "{\"repeated_timestamp\":[\"2023-01-01T00:00:00Z\",\"2023-01-02T00:00:00Z\"]}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_MultipleRepeatedFields(t *testing.T) {
	input := jsonmarshallv2.RepeatedFields{
		RepeatedBool:   []bool{true, false},
		RepeatedInt32:  []int{10, 20, 30},
		RepeatedString: []string{"a", "b", "c"},
	}
	expected := "{\"repeated_string\":[\"a\",\"b\",\"c\"],\"repeated_int32\":[10,20,30],\"repeated_bool\":[true,false]}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_EmptyRepeatedFields(t *testing.T) {
	input := jsonmarshallv2.RepeatedFields{
		RepeatedString: []string{},
	}
	expected := "{}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_OptionalFieldsNoInput(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{}
	expected := "{}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

func TestJsonMarshall_OptionalFieldsZeroValues(t *testing.T) {
	input := jsonmarshallv2.OptionalFields{
		OptionalBool:   false,
		OptionalInt32:  0,
		OptionalInt64:  0,
		OptionalString: "",
	}
	expected := "{}"
	actual, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("error marshalling input: %v", err)
	}
	// Parse both expected and actual as JSON objects
	var expectedObj map[string]interface{}
	var actualObj map[string]interface{}

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Fatalf("error unmarshalling expected: %v", err)
	}
	if err := json.Unmarshal(actual, &actualObj); err != nil {
		t.Fatalf("error unmarshalling actual: %v", err)
	}

	// Compare the JSON objects
	if diff := cmp.Diff(expectedObj, actualObj); diff != "" {
		t.Fatalf("Unexpected diff (-want, +got): %s", diff)
	}
}

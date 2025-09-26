package test

import (
	"encoding/json"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/test/new"
	"github.com/databricks/databricks-sdk-go/service/test/old"
	"github.com/stretchr/testify/assert"
)

func TestMarshallMessageJSONSerialization(t *testing.T) {
	// Create test data that will be used for both old and new structs
	testMap := map[string]string{"key1": "value1", "key2": "value2"}
	testRepeatedBool := []bool{true, false, true}
	testRepeatedString := []string{"one", "two", "three"}
	testRepeatedInt := []int{1, 2, 3}
	testRepeatedEnum := []old.TestEnum{old.TestEnumTestEnumOne, old.TestEnumTestEnumTwo}

	// Create old struct
	oldMsg := old.TestMarshallMessage{
		Map:      testMap,
		TestEnum: old.TestEnumTestEnumOne,
		TestNestedMessage: &old.TestNestedMessage{
			Map:                  testMap,
			NestedBool:           true,
			NestedEnum:           old.TestEnumTestEnumOne,
			NestedInt:            42,
			NestedRepeatedEnum:   testRepeatedEnum,
			NestedRepeatedString: testRepeatedString,
			NestedRequiredBool:   true,
			NestedRequiredInt:    100,
			NestedString:         "nested string",
		},
		TestOptionalBool:   true,
		TestOptionalInt:    42,
		TestOptionalInt64:  64,
		TestOptionalString: "optional string",
		TestRepeatedBool:   testRepeatedBool,
		TestRepeatedEnum:   testRepeatedEnum,
		TestRepeatedInt:    testRepeatedInt,
		TestRepeatedNestedMessage: []old.TestNestedMessage{{
			Map:                  testMap,
			NestedBool:           true,
			NestedEnum:           old.TestEnumTestEnumOne,
			NestedInt:            42,
			NestedRepeatedEnum:   testRepeatedEnum,
			NestedRepeatedString: testRepeatedString,
			NestedRequiredBool:   true,
			NestedRequiredInt:    100,
			NestedString:         "nested string",
		}},
		TestRepeatedString: testRepeatedString,
		TestRequiredBool:   true,
		TestRequiredEnum:   old.TestEnumTestEnumOne,
		TestRequiredInt:    100,
		TestRequiredInt64:  200,
		TestRequiredString: "required string",
	}

	// Create new struct with same data
	newMsg := new.TestMarshallMessage{
		Map:      testMap,
		TestEnum: new.TestEnum(old.TestEnumTestEnumOne),
		TestNestedMessage: &new.TestNestedMessage{
			Map:                  testMap,
			NestedBool:           true,
			NestedEnum:           new.TestEnum(old.TestEnumTestEnumOne),
			NestedInt:            42,
			NestedRepeatedEnum:   []new.TestEnum{new.TestEnum(old.TestEnumTestEnumOne), new.TestEnum(old.TestEnumTestEnumTwo)},
			NestedRepeatedString: testRepeatedString,
			NestedRequiredBool:   true,
			NestedRequiredInt:    100,
			NestedString:         "nested string",
		},
		TestOptionalBool:   true,
		TestOptionalInt:    42,
		TestOptionalInt64:  64,
		TestOptionalString: "optional string",
		TestRepeatedBool:   testRepeatedBool,
		TestRepeatedEnum:   []new.TestEnum{new.TestEnum(old.TestEnumTestEnumOne), new.TestEnum(old.TestEnumTestEnumTwo)},
		TestRepeatedInt:    testRepeatedInt,
		TestRepeatedNestedMessage: []new.TestNestedMessage{{
			Map:                  testMap,
			NestedBool:           true,
			NestedEnum:           new.TestEnum(old.TestEnumTestEnumOne),
			NestedInt:            42,
			NestedRepeatedEnum:   []new.TestEnum{new.TestEnum(old.TestEnumTestEnumOne), new.TestEnum(old.TestEnumTestEnumTwo)},
			NestedRepeatedString: testRepeatedString,
			NestedRequiredBool:   true,
			NestedRequiredInt:    100,
			NestedString:         "nested string",
		}},
		TestRepeatedString: testRepeatedString,
		TestRequiredBool:   true,
		TestRequiredEnum:   new.TestEnum(old.TestEnumTestEnumOne),
		TestRequiredInt:    100,
		TestRequiredInt64:  200,
		TestRequiredString: "required string",
	}

	// Marshal both to JSON
	oldJSON, err := json.Marshal(oldMsg)
	assert.NoError(t, err)

	newJSON, err := json.Marshal(newMsg)
	assert.NoError(t, err)

	// Compare the JSON strings
	assert.JSONEq(t, string(oldJSON), string(newJSON), "JSON serialization should be identical between old and new structs")

	// Also test the "ToPb" and "FromPb" methods
	pbNew, err := new.TestMarshallMessageToPb(&newMsg)
	assert.NoError(t, err)

	recoverNew, err := new.TestMarshallMessageFromPb(pbNew)
	assert.NoError(t, err)
	assert.Equal(t, newMsg, *recoverNew)

}

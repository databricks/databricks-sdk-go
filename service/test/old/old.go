package old

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type TestEnum string

const TestEnumTestEnumFive TestEnum = `TEST_ENUM_FIVE`

const TestEnumTestEnumFour TestEnum = `TEST_ENUM_FOUR`

const TestEnumTestEnumOne TestEnum = `TEST_ENUM_ONE`

const TestEnumTestEnumThree TestEnum = `TEST_ENUM_THREE`

const TestEnumTestEnumTwo TestEnum = `TEST_ENUM_TWO`

// String representation for [fmt.Print]
func (f *TestEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TestEnum) Set(v string) error {
	switch v {
	case `TEST_ENUM_FIVE`, `TEST_ENUM_FOUR`, `TEST_ENUM_ONE`, `TEST_ENUM_THREE`, `TEST_ENUM_TWO`:
		*f = TestEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TEST_ENUM_FIVE", "TEST_ENUM_FOUR", "TEST_ENUM_ONE", "TEST_ENUM_THREE", "TEST_ENUM_TWO"`, v)
	}
}

// Type always returns TestEnum to satisfy [pflag.Value] interface
func (f *TestEnum) Type() string {
	return "TestEnum"
}

type TestMarshallMessage struct {
	Map map[string]string `json:"map,omitempty"`

	TestEnum TestEnum `json:"test_enum,omitempty"`

	TestNestedMessage *TestNestedMessage `json:"test_nested_message,omitempty"`

	TestOptionalBool bool `json:"test_optional_bool,omitempty"`

	TestOptionalInt int `json:"test_optional_int,omitempty"`

	TestOptionalInt64 int64 `json:"test_optional_int64,omitempty"`

	TestOptionalString string `json:"test_optional_string,omitempty"`

	TestRepeatedBool []bool `json:"test_repeated_bool,omitempty"`

	TestRepeatedEnum []TestEnum `json:"test_repeated_enum,omitempty"`

	TestRepeatedInt []int `json:"test_repeated_int,omitempty"`

	TestRepeatedNestedMessage []TestNestedMessage `json:"test_repeated_nested_message,omitempty"`

	TestRepeatedString []string `json:"test_repeated_string,omitempty"`

	TestRequiredBool bool `json:"test_required_bool"`

	TestRequiredEnum TestEnum `json:"test_required_enum"`

	TestRequiredInt int `json:"test_required_int"`

	TestRequiredInt64 int64 `json:"test_required_int64"`

	TestRequiredString string `json:"test_required_string"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TestMarshallMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TestMarshallMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TestNestedMessage struct {
	Map map[string]string `json:"map,omitempty"`

	NestedBool bool `json:"nested_bool,omitempty"`

	NestedEnum TestEnum `json:"nested_enum,omitempty"`

	NestedInt int `json:"nested_int,omitempty"`

	NestedRepeatedEnum []TestEnum `json:"nested_repeated_enum,omitempty"`

	NestedRepeatedString []string `json:"nested_repeated_string,omitempty"`

	NestedRequiredBool bool `json:"nested_required_bool"`

	NestedRequiredInt int `json:"nested_required_int"`

	NestedString string `json:"nested_string,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TestNestedMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TestNestedMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

package new

import (
	"encoding/json"
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

	// Wire name: 'map'
	Map map[string]string

	// Wire name: 'test_enum'
	TestEnum TestEnum

	// Wire name: 'test_nested_message'
	TestNestedMessage *TestNestedMessage

	// Wire name: 'test_optional_bool'
	TestOptionalBool bool

	// Wire name: 'test_optional_int'
	TestOptionalInt int

	// Wire name: 'test_optional_int64'
	TestOptionalInt64 int64

	// Wire name: 'test_optional_string'
	TestOptionalString string

	// Wire name: 'test_repeated_bool'
	TestRepeatedBool []bool

	// Wire name: 'test_repeated_enum'
	TestRepeatedEnum []TestEnum

	// Wire name: 'test_repeated_int'
	TestRepeatedInt []int

	// Wire name: 'test_repeated_nested_message'
	TestRepeatedNestedMessage []TestNestedMessage

	// Wire name: 'test_repeated_string'
	TestRepeatedString []string

	// Wire name: 'test_required_bool'
	TestRequiredBool bool

	// Wire name: 'test_required_enum'
	TestRequiredEnum TestEnum

	// Wire name: 'test_required_int'
	TestRequiredInt int

	// Wire name: 'test_required_int64'
	TestRequiredInt64 int64

	// Wire name: 'test_required_string'
	TestRequiredString string

	ForceSendFields []string `tf:"-"`
}

func (st *TestMarshallMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &testMarshallMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TestMarshallMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TestMarshallMessage) MarshalJSON() ([]byte, error) {
	pb, err := TestMarshallMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TestNestedMessage struct {

	// Wire name: 'map'
	Map map[string]string

	// Wire name: 'nested_bool'
	NestedBool bool

	// Wire name: 'nested_enum'
	NestedEnum TestEnum

	// Wire name: 'nested_int'
	NestedInt int

	// Wire name: 'nested_repeated_enum'
	NestedRepeatedEnum []TestEnum

	// Wire name: 'nested_repeated_string'
	NestedRepeatedString []string

	// Wire name: 'nested_required_bool'
	NestedRequiredBool bool

	// Wire name: 'nested_required_int'
	NestedRequiredInt int

	// Wire name: 'nested_string'
	NestedString string

	ForceSendFields []string `tf:"-"`
}

func (st *TestNestedMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &testNestedMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := testNestedMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TestNestedMessage) MarshalJSON() ([]byte, error) {
	pb, err := testNestedMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func TestMarshallMessageToPb(st *TestMarshallMessage) (*testMarshallMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &testMarshallMessagePb{}
	pb.Map = st.Map
	pb.TestEnum = st.TestEnum
	pb.TestNestedMessage = st.TestNestedMessage
	pb.TestOptionalBool = st.TestOptionalBool
	pb.TestOptionalInt = st.TestOptionalInt
	pb.TestOptionalInt64 = st.TestOptionalInt64
	pb.TestOptionalString = st.TestOptionalString
	pb.TestRepeatedBool = st.TestRepeatedBool
	pb.TestRepeatedEnum = st.TestRepeatedEnum
	pb.TestRepeatedInt = st.TestRepeatedInt
	pb.TestRepeatedNestedMessage = st.TestRepeatedNestedMessage
	pb.TestRepeatedString = st.TestRepeatedString
	pb.TestRequiredBool = st.TestRequiredBool
	pb.TestRequiredEnum = st.TestRequiredEnum
	pb.TestRequiredInt = st.TestRequiredInt
	pb.TestRequiredInt64 = st.TestRequiredInt64
	pb.TestRequiredString = st.TestRequiredString

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type testMarshallMessagePb struct {
	Map                       map[string]string   `json:"map,omitempty"`
	TestEnum                  TestEnum            `json:"test_enum,omitempty"`
	TestNestedMessage         *TestNestedMessage  `json:"test_nested_message,omitempty"`
	TestOptionalBool          bool                `json:"test_optional_bool,omitempty"`
	TestOptionalInt           int                 `json:"test_optional_int,omitempty"`
	TestOptionalInt64         int64               `json:"test_optional_int64,omitempty"`
	TestOptionalString        string              `json:"test_optional_string,omitempty"`
	TestRepeatedBool          []bool              `json:"test_repeated_bool,omitempty"`
	TestRepeatedEnum          []TestEnum          `json:"test_repeated_enum,omitempty"`
	TestRepeatedInt           []int               `json:"test_repeated_int,omitempty"`
	TestRepeatedNestedMessage []TestNestedMessage `json:"test_repeated_nested_message,omitempty"`
	TestRepeatedString        []string            `json:"test_repeated_string,omitempty"`
	TestRequiredBool          bool                `json:"test_required_bool"`
	TestRequiredEnum          TestEnum            `json:"test_required_enum"`
	TestRequiredInt           int                 `json:"test_required_int"`
	TestRequiredInt64         int64               `json:"test_required_int64"`
	TestRequiredString        string              `json:"test_required_string"`

	ForceSendFields []string `json:"-" url:"-"`
}

func TestMarshallMessageFromPb(pb *testMarshallMessagePb) (*TestMarshallMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TestMarshallMessage{}
	st.Map = pb.Map
	st.TestEnum = pb.TestEnum
	st.TestNestedMessage = pb.TestNestedMessage
	st.TestOptionalBool = pb.TestOptionalBool
	st.TestOptionalInt = pb.TestOptionalInt
	st.TestOptionalInt64 = pb.TestOptionalInt64
	st.TestOptionalString = pb.TestOptionalString
	st.TestRepeatedBool = pb.TestRepeatedBool
	st.TestRepeatedEnum = pb.TestRepeatedEnum
	st.TestRepeatedInt = pb.TestRepeatedInt
	st.TestRepeatedNestedMessage = pb.TestRepeatedNestedMessage
	st.TestRepeatedString = pb.TestRepeatedString
	st.TestRequiredBool = pb.TestRequiredBool
	st.TestRequiredEnum = pb.TestRequiredEnum
	st.TestRequiredInt = pb.TestRequiredInt
	st.TestRequiredInt64 = pb.TestRequiredInt64
	st.TestRequiredString = pb.TestRequiredString

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *testMarshallMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st testMarshallMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func testNestedMessageToPb(st *TestNestedMessage) (*testNestedMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &testNestedMessagePb{}
	pb.Map = st.Map
	pb.NestedBool = st.NestedBool
	pb.NestedEnum = st.NestedEnum
	pb.NestedInt = st.NestedInt
	pb.NestedRepeatedEnum = st.NestedRepeatedEnum
	pb.NestedRepeatedString = st.NestedRepeatedString
	pb.NestedRequiredBool = st.NestedRequiredBool
	pb.NestedRequiredInt = st.NestedRequiredInt
	pb.NestedString = st.NestedString

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type testNestedMessagePb struct {
	Map                  map[string]string `json:"map,omitempty"`
	NestedBool           bool              `json:"nested_bool,omitempty"`
	NestedEnum           TestEnum          `json:"nested_enum,omitempty"`
	NestedInt            int               `json:"nested_int,omitempty"`
	NestedRepeatedEnum   []TestEnum        `json:"nested_repeated_enum,omitempty"`
	NestedRepeatedString []string          `json:"nested_repeated_string,omitempty"`
	NestedRequiredBool   bool              `json:"nested_required_bool"`
	NestedRequiredInt    int               `json:"nested_required_int"`
	NestedString         string            `json:"nested_string,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func testNestedMessageFromPb(pb *testNestedMessagePb) (*TestNestedMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TestNestedMessage{}
	st.Map = pb.Map
	st.NestedBool = pb.NestedBool
	st.NestedEnum = pb.NestedEnum
	st.NestedInt = pb.NestedInt
	st.NestedRepeatedEnum = pb.NestedRepeatedEnum
	st.NestedRepeatedString = pb.NestedRepeatedString
	st.NestedRequiredBool = pb.NestedRequiredBool
	st.NestedRequiredInt = pb.NestedRequiredInt
	st.NestedString = pb.NestedString

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *testNestedMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st testNestedMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

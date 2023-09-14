package marshal

import (
	"testing"
)

func TestStringNotSet(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st:             customStruct{},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestStringDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Str: "",
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestStringValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Str: "val",
			},
			jsonString:              `{"string":"val","childfs":{},"childnofs":{}}`,
			matchClassic:            true,
			matchUnmarshal:          false,
			unmarshalForceSendField: []string{"Str"},
		},
	)
}

func TestStringForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Str:             "",
				ForceSendFields: []string{"Str"},
			},
			jsonString:              `{"string":"","childfs":{},"childnofs":{}}`,
			matchClassic:            false,
			matchUnmarshal:          true,
			unmarshalForceSendField: []string{"Str"},
		},
	)
}

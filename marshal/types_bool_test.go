package marshal

import (
	"testing"
)

func TestBoolNotSet(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st:             customStruct{},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestBoolDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Bool: false,
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestBoolValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Bool: true,
			},
			jsonString:              `{"bool":true,"childfs":{},"childnofs":{}}`,
			matchClassic:            true,
			matchUnmarshal:          false,
			unmarshalForceSendField: []string{"Bool"},
		},
	)
}

func TestBoolForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Bool:            false,
				ForceSendFields: []string{"Bool"},
			},
			jsonString:              `{"bool":false,"childfs":{},"childnofs":{}}`,
			matchClassic:            false,
			matchUnmarshal:          true,
			unmarshalForceSendField: []string{"Bool"},
		},
	)
}

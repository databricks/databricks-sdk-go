package marshal

import (
	"testing"
)

func TestFloatNotSet(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st:             customStruct{},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestFloatDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Float: 0.0,
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestFloatValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Float: 2.3,
			},
			jsonString:              `{"float":2.3,"childfs":{},"childnofs":{}}`,
			matchClassic:            true,
			matchUnmarshal:          false,
			unmarshalForceSendField: []string{"Float"},
		},
	)
}

func TestFloatForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Float:           0,
				ForceSendFields: []string{"Float"},
			},
			jsonString:              `{"float":0,"childfs":{},"childnofs":{}}`,
			matchClassic:            false,
			matchUnmarshal:          true,
			unmarshalForceSendField: []string{"Float"},
		},
	)
}

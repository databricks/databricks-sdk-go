package marshal

import (
	"testing"
)

func TestPFloatNotSet(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st:             customStruct{},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestPFloatDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PFloat: nil,
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestPFloatValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PFloat: Ptr(0.0),
			},
			jsonString:     `{"pfloat":0,"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: false,
		},
	)
}

func TestPFloatForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PFloat:          nil,
				ForceSendFields: []string{"PFloat"},
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: false,
		},
	)
}

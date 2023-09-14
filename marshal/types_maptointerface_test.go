package marshal

import (
	"testing"
)

func TestMapToInterfaceNotSet(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st:             customStruct{},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestMapToInterfaceDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToInt: nil,
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestMapToInterfaceValue(t *testing.T) {
	element := map[string]interface{}{
		"key": "3",
	}
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToInt: element,
			},
			jsonString:     `{"maptoint":{"key":"3"},"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: false,
		},
	)
	//assert.Fail(t, "S")
}

func TestMapToInterfaceForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToInt:        nil,
				ForceSendFields: []string{"MapToInt"},
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: false, // Unmarshal won't include the ForceSendFields
		},
	)
}

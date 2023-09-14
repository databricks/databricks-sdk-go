package marshal

import (
	"testing"
)

func TestInterfaceNotSet(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st:             customStruct{},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestInterfaceDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Interface: nil,
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestInterfaceValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Interface: structNoFS{
					Int: 3,
				},
			},
			jsonString:     `{"interface":{"childint":3},"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: false,
		},
	)
}

func TestInterfaceForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Interface:       nil,
				ForceSendFields: []string{"Interface"},
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: false, // Unmarshal won't include the ForceSendFields
		},
	)
}

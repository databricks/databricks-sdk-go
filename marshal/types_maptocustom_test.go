package marshal

import (
	"testing"
)

func TestMapToCustomTypeFSDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToCustomTypeFS: nil,
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestMapToCustomTypeFSValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToCustomTypeFS: map[string]structFS{
					"key": {
						Int: 2,
					},
				},
			},
			jsonString:     `{"maptocustomtypefs":{"key":{"childint":2}},"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: false, // Unmarshal will include ForceSendFields for child
		},
	)
}

func TestMapToCustomTypeFSValueForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToCustomTypeFS: map[string]structFS{
					"key": {
						Int:             0,
						ForceSendFields: []string{"Int"},
					},
				},
			},
			jsonString:     `{"maptocustomtypefs":{"key":{"childint":0}},"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestMapToCustomTypeFSForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToCustomTypeFS: nil,
				ForceSendFields:   []string{"MapToCustomTypeFS"},
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: false, // Unmarshal won't include the ForceSendFields
		},
	)
}

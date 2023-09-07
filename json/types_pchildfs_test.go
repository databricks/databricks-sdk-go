package marshal

import (
	"testing"
)

func TestPChildFSNotSet(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st:             customStruct{},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestPChildFSDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildFS: nil,
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)
}

func TestPChildFSValueEmpty(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildFS: &structFS{},
			},
			jsonString:     `{"pchildfs":{},"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: true, // PChildFS will have ForceSendFields set
		},
	)
}

func TestPChildFSValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildFS: &structFS{
					Int: 3,
				},
			},
			jsonString:     `{"pchildfs":{"childint":3},"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: false, // PChildFS will have ForceSendFields set
		},
	)
}

func TestPChildFSForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildFS:        nil,
				ForceSendFields: []string{"PChildFS"},
			},
			jsonString:     `{"childfs":{},"childnofs":{}}`,
			matchClassic:   true,
			matchUnmarshal: false, // Unmarshal won't include the ForceSendFields
		},
	)
}

func TestPChildFSChuildForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildFS: &structFS{
					ForceSendFields: []string{"Int"},
				},
				ForceSendFields: []string{"Int"},
			},
			jsonString:              `{"childfs":{},"pchildfs":{"childint":0},"childnofs":{}, "int":0}`,
			matchClassic:            false,
			matchUnmarshal:          true, // Unmarshal won't include the ForceSendFields
			unmarshalForceSendField: []string{"Int"},
		},
	)
}

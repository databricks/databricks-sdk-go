package marshal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaults(t *testing.T) {
	executeBasicMarshalTest(t, basicMarshalTest{
		st: customStruct{
			Int:   0,
			Float: 0.0,
			Str:   "",
			Bool:  false,

			PInt:   nil,
			PFloat: nil,
			PStr:   nil,
			PBool:  nil,

			Slice:      nil,
			Interface:  nil,
			BasicMap:   nil,
			MapToArray: nil,

			ChildFS:   structFS{},
			ChildNoFS: structNoFS{},

			PChildFS:   nil,
			PChildNoFS: nil,

			MapToCustomTypeFS: nil,
		},
		jsonString:     `{"childfs":{},"childnofs":{}}`,
		matchClassic:   true,
		matchUnmarshal: true,
	})
}

func TestForce(t *testing.T) {
	executeBasicMarshalTest(t, basicMarshalTest{
		st: customStruct{
			Int:   0,
			Float: 0.0,
			Str:   "",
			Bool:  false,

			PInt:   nil,
			PFloat: nil,
			PStr:   nil,
			PBool:  nil,

			Slice:      nil,
			Interface:  nil,
			BasicMap:   nil,
			MapToArray: nil,

			ChildFS:   structFS{},
			ChildNoFS: structNoFS{},

			PChildFS:   nil,
			PChildNoFS: nil,

			MapToCustomTypeFS: nil,

			ForceSendFields: []string{"Int", "Float", "Str", "Bool"},
		},
		jsonString:              `{"int":0, "float":0, "string":"", "bool":false, "childfs":{},"childnofs":{}}`,
		matchClassic:            false,
		matchUnmarshal:          true,
		unmarshalForceSendField: []string{"Int", "Float", "Str", "Bool"},
	})
}

func TestNotDefault(t *testing.T) {
	st := customStruct{
		Int:   1,
		Float: 1.0,
		Str:   "1",
		Bool:  true,

		PInt:   Ptr(0),
		PFloat: Ptr(0.0),
		PStr:   Ptr(""),
		PBool:  Ptr(false),

		Slice: []int{9},
		BasicMap: map[string]string{
			"key": "value",
		},
		MapToArray: map[string][]interface{}{
			"key": {2.0, "field"},
		},

		ChildFS: structFS{
			Int:             1,
			ForceSendFields: []string{"Int"},
		},
		ChildNoFS: structNoFS{
			Int: 2,
		},

		PChildFS: &structFS{
			Int:             3,
			ForceSendFields: []string{"Int"},
		},
		PChildNoFS: &structNoFS{
			Int: 4,
		},

		MapToCustomTypeFS: map[string]structFS{
			"struct1": {
				Int:             1,
				ForceSendFields: []string{"Int"},
			},
		},
	}
	rc := executeBasicMarshalTest(t, basicMarshalTest{
		st: st,
		jsonString: `{"int":1, "float":1, "string":"1", "bool":true,
		"pint":0, "pfloat":0,"pbool":false,"pstring":"",
		"slice":[9], "childfs":{"childint":1},"childnofs":{"childint":2},
		"basicmap":{"key":"value"},"maptoarray":{"key":[2,"field"]},
		"pchildfs":{"childint":3},"pchildnofs":{"childint":4},
		"maptocustomtypefs":{"struct1":{"childint":1}}}`,
		matchClassic:            true,
		matchUnmarshal:          false,
		unmarshalForceSendField: []string{"Int", "Bool", "Str", "Float"},
	})
	rc.ForceSendFields = nil
	assert.Equal(t, rc, st)

}

type SimpleStruct struct {
	ID string `json:"id"`
}

func TestIntToString(t *testing.T) {
	jsonString := `{"id":1}`
	var res SimpleStruct
	err := Unmarshal([]byte(jsonString), &res)
	assert.NoError(t, err)
	assert.Equal(t, "1", res.ID)
}

package marshal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

type customStruct struct {
	// Basic types
	Int   int     `json:"int,omitempty"`
	Float float64 `json:"float,omitempty"`
	Str   string  `json:"string,omitempty"`
	Bool  bool    `json:"bool,omitempty"`

	// Pointers to basic types
	PInt   *int     `json:"pint,omitempty"`
	PFloat *float64 `json:"pfloat,omitempty"`
	PStr   *string  `json:"pstring,omitempty"`
	PBool  *bool    `json:"pbool,omitempty"`

	// Other types and structs
	Slice      []int                    `json:"slice,omitempty"`
	Interface  interface{}              `json:"interface,omitempty"`
	BasicMap   map[string]string        `json:"basicmap,omitempty"`
	MapToArray map[string][]interface{} `json:"maptoarray,omitempty"`
	MapToInt   map[string]interface{}   `json:"maptoint,omitempty"`

	ChildFS   structFS   `json:"childfs,omitempty"`
	ChildNoFS structNoFS `json:"childnofs,omitempty"`

	PChildFS   *structFS   `json:"pchildfs,omitempty"`
	PChildNoFS *structNoFS `json:"pchildnofs,omitempty"`

	MapToCustomTypeFS map[string]structFS `json:"maptocustomtypefs,omitempty"`

	ForceSendFields []string `json:"-"`
}

type structNoFS struct {
	Int int `json:"childint,omitempty"`
}

type structFS struct {
	Int int `json:"childint,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *structFS) UnmarshalJSON(b []byte) error {
	type C structFS
	return Unmarshal(b, (*C)(s))
}

func (s structFS) MarshalJSON() ([]byte, error) {
	type C structFS
	return Marshal((C)(s))
}

func (s *customStruct) UnmarshalJSON(b []byte) error {
	type C customStruct
	return Unmarshal(b, (*C)(s))
}

func (s customStruct) MarshalJSON() ([]byte, error) {
	type C customStruct
	return Marshal((C)(s))
}

func Ptr[T any](v T) *T {
	return &v
}

type basicMarshalTest struct {
	st         customStruct
	jsonString string
	// Compare marshal results with normal marshal
	matchClassic bool
	// Unmarshal may not match, since ForceSendFields will be populated during
	// custom unmarshal process
	matchUnmarshal          bool
	unmarshalForceSendField []string
}

func executeBasicMarshalTest(t *testing.T, tc basicMarshalTest) customStruct {
	// Convert to JSON
	res, err := json.Marshal(tc.st)
	assert.NoError(t, err, "error while executing custom marshal")
	compareJSON(t, tc.jsonString, string(res))

	// Compare with original json.Marshal
	if tc.matchClassic {
		type C customStruct
		original, err := json.Marshal((C)(tc.st))
		assert.NoError(t, err, "error while executing classic marshal")
		compareJSON(t, string(original), string(res))
	}

	var reconstruct customStruct
	err = json.Unmarshal(res, &reconstruct)
	assert.NoError(t, err, "error while unmarshaling")
	if tc.matchUnmarshal {
		assert.Equal(t, tc.st, reconstruct)
	}

	compareSlices(t, tc.unmarshalForceSendField, reconstruct.ForceSendFields)

	return reconstruct
}

func compareSlices(t *testing.T, s1 []string, s2 []string) {
	slices.Sort(s1)
	slices.Sort(s2)

	assert.Equal(t, s1, s2, "unexpected elements in ForceSendFields array")
}

func compareJSON(t *testing.T, json1 string, json2 string) {
	var i1 interface{}
	var i2 interface{}
	err := json.Unmarshal([]byte(json1), &i1)
	assert.NoError(t, err, "error while unmarshalling")
	err = json.Unmarshal([]byte(json2), &i2)
	assert.NoError(t, err, "error while unmarshalling")
	if !reflect.DeepEqual(i1, i2) {
		message := fmt.Sprintf("result doesn't match expected string:\nwant :%v\ngot: %v", i1, i2)
		assert.Fail(t, message)
	}
}

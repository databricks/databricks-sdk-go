package marshal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CompositeParent struct {
	Field1 int `json:"field1,omitempty"`

	ForceSendFields []string `json:"-"`
}

type CompositeSecondParent struct {
	Field3 int `json:"field3,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CompositeSecondParent) UnmarshalJSON(b []byte) error {
	return Unmarshal(b, s)
}

func (s CompositeSecondParent) MarshalJSON() ([]byte, error) {
	return Marshal(s)
}

func (s *CompositeParent) UnmarshalJSON(b []byte) error {
	return Unmarshal(b, s)
}

func (s CompositeParent) MarshalJSON() ([]byte, error) {
	return Marshal(s)
}

type compositeChild struct {
	CompositeParent
	*CompositeSecondParent
	Field2 int `json:"field2,omitempty"`
}

func (s compositeChild) MarshalJSON() ([]byte, error) {
	return Marshal(s)
}

func (s *compositeChild) UnmarshalJSON(b []byte) error {
	return Unmarshal(b, s)
}

func TestComposite(t *testing.T) {
	executeCompositeMarshalTest(
		t, compositeTest{
			st: compositeChild{
				Field2: 2,
				CompositeParent: CompositeParent{
					Field1:          1,
					ForceSendFields: []string{"Field1"},
				},
				CompositeSecondParent: &CompositeSecondParent{
					Field3:          3,
					ForceSendFields: []string{"Field3"},
				},
			},
			jsonString:     `{"field1":1, "field2":2, "field3":3}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)

}

func TestCompositeNil(t *testing.T) {
	element := compositeChild{
		Field2: 2,
		CompositeParent: CompositeParent{
			Field1:          1,
			ForceSendFields: []string{"Field1"},
		},
	}
	result := executeCompositeMarshalTest(
		t, compositeTest{
			st:           element,
			jsonString:   `{"field1":1, "field2":2}`,
			matchClassic: true,
			//CompositeSecondParent will be present due to limitations of the unmarshalling
			matchUnmarshal: false,
		},
	)
	element.CompositeSecondParent = &CompositeSecondParent{}
	assert.Equal(t, element, result)

}

func TestCompositeDefault(t *testing.T) {
	executeCompositeMarshalTest(
		t, compositeTest{
			st: compositeChild{
				Field2: 0,
				CompositeParent: CompositeParent{
					Field1:          0,
					ForceSendFields: []string{"Field1"},
				},
				CompositeSecondParent: &CompositeSecondParent{
					Field3:          0,
					ForceSendFields: []string{"Field3"},
				},
			},
			jsonString:     `{"field1":0, "field3":0}`,
			matchClassic:   true,
			matchUnmarshal: true,
		},
	)

}

type compositeTest struct {
	st         compositeChild
	jsonString string
	// Compare marshal results with normal marshal
	matchClassic bool
	// Unmarshal may not match, since ForceSendFields will be populated during
	// custom unmarshal process
	matchUnmarshal bool
}

func executeCompositeMarshalTest(t *testing.T, tc compositeTest) compositeChild {
	// Convert to JSON
	res, err := json.Marshal(tc.st)
	assert.NoError(t, err, "error while executing custom marshal")
	compareJSON(t, tc.jsonString, string(res))

	var reconstruct compositeChild
	err = json.Unmarshal(res, &reconstruct)
	assert.NoError(t, err, "error while unmarshaling")
	if tc.matchUnmarshal {
		assert.Equal(t, tc.st, reconstruct)
	}

	return reconstruct
}

type noSendFieldChild struct {
	*CompositeSecondParent
	Field2 int `json:"field2"`
}

func (s *noSendFieldChild) UnmarshalJSON(b []byte) error {
	return Unmarshal(b, s)
}

func (s noSendFieldChild) MarshalJSON() ([]byte, error) {
	return Marshal(s)
}

func TestNoSendFieldChild(t *testing.T) {
	st := noSendFieldChild{
		Field2: 2,
	}

	res, err := json.Marshal(st)
	assert.NoError(t, err, "error while executing custom marshal")
	compareJSON(t, `{"field2":2}`, string(res))

	var reconstruct noSendFieldChild
	err = json.Unmarshal(res, &reconstruct)
	assert.NoError(t, err, "error while unmarshaling")
	st.CompositeSecondParent = &CompositeSecondParent{}
	assert.Equal(t, st, reconstruct)

}

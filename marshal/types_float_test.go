package marshal

import (
	"math"
	"testing"
)

func TestFloatDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Float: 0.0,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
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
			unmarshalForceSendField: []string{"Float"},
		},
	)
}

// TestFloatNaN verifies that fields whose API value is the JSON string "NaN"
// are correctly decoded to math.NaN().  MLflow and other Databricks APIs
// encode special float values as quoted strings (see issue #1498).
func TestFloatNaN(t *testing.T) {
	var s customStruct
	err := Unmarshal([]byte(`{"float":"NaN"}`), &s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !math.IsNaN(s.Float) {
		t.Errorf("expected NaN, got %v", s.Float)
	}
}

func TestFloatInf(t *testing.T) {
	var s customStruct
	err := Unmarshal([]byte(`{"float":"Inf"}`), &s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !math.IsInf(s.Float, 1) {
		t.Errorf("expected +Inf, got %v", s.Float)
	}
}

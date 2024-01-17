package marshal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ignoredField struct {
	A string `json:"a"`
	B string `json:"-"`
}

func TestUnmarshalCustom_Default(t *testing.T) {
	raw := []byte(`{"a": "foo", "b": "bar"}`)

	var request ignoredField
	UnmarshalCustom(raw, &request, UnmarshalOptions{})

	assert.Equal(t, ignoredField{A: "foo", B: ""}, request)
}

func TestUnmarshalCustom_UnmarshalIgnoredFields(t *testing.T) {
	raw := []byte(`{"a": "foo", "b": "bar"}`)

	var request ignoredField
	UnmarshalCustom(raw, &request, UnmarshalOptions{
		UnmarshalTopLevelIgnoredFields: true,
	})

	assert.Equal(t, ignoredField{A: "foo", B: "bar"}, request)
}

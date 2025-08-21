package types

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// FieldMask represents a field mask as defined in Google's Well Known Types.
// It is used to specify which fields of a resource should be included or excluded
// in a request or response.
type FieldMask struct {
	Paths []string
}

// MarshalJSON implements the json.Marshaler interface by formatting the
// field mask as a string according to Google Well Known Type
func (f *FieldMask) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, f.ToWireFormat())), nil
}

// ToWireFormat returns the field mask as a string according to Google Well Known Type
// which is a comma-separated list of field names.
func (f FieldMask) ToWireFormat() string {
	return strings.Join(f.Paths, ",")
}

// UnmarshalJSON implements the json.Unmarshaler interface by parsing the
// field mask from a string according to Google Well Known Type
func (f *FieldMask) UnmarshalJSON(data []byte) error {
	if f == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s != "" {
		f.Paths = strings.Split(s, ",")
	} else {
		f.Paths = []string{}
	}

	return nil
}

// EncodeValues encodes the FieldMask into a url.Values object.
// If the FieldMask is nil or empty, it returns nil.
// If the url.Values is nil, it returns an error.
func (f *FieldMask) EncodeValues(key string, v *url.Values) error {
	if f == nil || len(f.Paths) == 0 {
		return nil
	}
	if v == nil {
		return fmt.Errorf("url.Values is nil")
	}

	v.Set(key, f.ToWireFormat())
	return nil
}

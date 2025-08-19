package types

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Time is a wrapper for time.Time to provide custom marshaling
// for JSON and URL query strings.
//
// It embeds time.Time, so all standard methods (Format, Parse, etc.)
// are directly accessible. The underlying time.Time value can be
// accessed via the .Time field.
//
// Example:
//
//	customTime := types.NewTime(time.Now())
//	goTime := customTime.Time // Access the underlying time.Time
type Time struct {
	time.Time
}

// NewTime creates a custom Time from a standard time.Time.
func NewTime(t time.Time) Time {
	return Time{Time: t}
}

// MarshalJSON implements the json.Marshaler interface by formatting the
// time as a string according to RFC3339Nano
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToWireFormat())), nil
}

// ToWireFormat returns the time as a string according to RFC3339Nano
func (t Time) ToWireFormat() string {
	return t.Time.Format(time.RFC3339Nano)
}

// UnmarshalJSON implements the json.Unmarshaler interface by parsing the
// time from a string according to RFC3339Nano
func (t *Time) UnmarshalJSON(b []byte) error {
	if t == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if s == "" {
		*t = Time{Time: time.Time{}}
		return nil
	}

	timeValue, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return err
	}

	*t = Time{Time: timeValue}
	return nil
}

// EncodeValues implements the query.Encoder interface by formatting the
// time as a string according to RFC3339Nano
func (t Time) EncodeValues(key string, v *url.Values) error {
	v.Set(key, t.ToWireFormat())
	return nil
}

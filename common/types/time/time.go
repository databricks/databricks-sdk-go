package time

import (
	"encoding/json"
	"fmt"
	"net/url"
	stdtime "time"
)

// Time is a wrapper for stdtime.Time to provide custom marshaling
// for JSON and URL query strings.
//
// It embeds stdtime.Time, so all standard methods (Format, Parse, etc.)
// are directly accessible. The underlying stdtime.Time value can be
// accessed via the .Time field.
//
// Example:
//
//	customTime := time.New(stdtime.Now())
//	goTime := customTime.AsTime() // Access the underlying stdtime.Time
type Time struct {
	stdtime.Time
}

// New creates a custom Time from a standard stdtime.Time.
func New(t stdtime.Time) *Time {
	return &Time{Time: t}
}

// AsTime returns the underlying stdtime.Time value.
func (x *Time) AsTime() stdtime.Time {
	if x == nil {
		return stdtime.Time{}
	}
	return x.Time
}

// MarshalJSON implements the json.Marshaler interface by formatting the
// time as a string according to RFC3339Nano
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.ToWireFormat())), nil
}

// ToWireFormat returns the time as a string according to RFC3339Nano
func (t Time) ToWireFormat() string {
	return t.Time.Format(stdtime.RFC3339Nano)
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
		*t = Time{Time: stdtime.Time{}}
		return nil
	}

	timeValue, err := stdtime.Parse(stdtime.RFC3339Nano, s)
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

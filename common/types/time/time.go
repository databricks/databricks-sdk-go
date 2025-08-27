package time

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
)

// Time is a wrapper for time.Time to provide custom marshaling
// for JSON and URL query strings.
//
// It embeds time.Time, so all standard methods (Format, Parse, etc.)
// are directly accessible. The underlying time.Time value can be
// accessed via the .AsTime() method.
//
// Example:
//
//	customTime := time.New(stdtime.Now())
//	goTime := customTime.AsTime() // Access the underlying stdtime.Time
type Time struct {
	time.Time
}

// New creates a custom Time from a standard stdtime.Time.
func New(t time.Time) *Time {
	return &Time{Time: t}
}

// AsTime returns the underlying stdtime.Time value.
func (x *Time) AsTime() time.Time {
	if x == nil {
		return time.Time{}
	}
	return x.Time
}

// MarshalJSON implements the [json.Marshaler] interface by formatting the
// time as a string according to RFC3339Nano
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format(time.RFC3339Nano))
}

// UnmarshalJSON implements the [json.Unmarshaler] interface by parsing the
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
		return errors.New("time is empty. It should be in RFC3339Nano format")
	}

	timeValue, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return err
	}

	*t = Time{Time: timeValue}
	return nil
}

// EncodeValues implements the [query.Encoder] interface by encoding the
// time as a string according to RFC3339Nano.
func (t Time) EncodeValues(key string, v *url.Values) error {
	v.Set(key, t.Time.Format(time.RFC3339Nano))
	return nil
}

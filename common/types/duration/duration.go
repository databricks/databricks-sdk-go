package duration

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Duration is a wrapper for time.Duration to provide custom marshaling
// for JSON and URL query strings.
//
// It embeds time.Duration, so all standard methods (Seconds, String, etc.)
// are directly accessible. The underlying time.Duration value can be
// accessed via the .AsDuration() method.
//
// Example:
//
//	customDur := durationpb.New(30 * time.Second)
//	goDur := customDur.AsDuration() // Access the underlying time.Duration
type Duration struct {
	time.Duration
}

// New creates a custom Duration from a standard time.Duration.
func New(d time.Duration) *Duration {
	return &Duration{Duration: d}
}

// AsDuration returns the underlying time.Duration value.
func (x *Duration) AsDuration() time.Duration {
	if x == nil {
		return 0
	}
	return x.Duration
}

// MarshalJSON implements the [json.Marshaler] interface by formatting the
// duration as a string according to Google Well Known Type.
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.toWireFormat())
}

// toWireFormat returns a string representation of Duration
// which follows the wire format from Google Well Known Type.
// a String that ends in s to indicate seconds and is preceded by
// the number of seconds, with nanoseconds expressed as fractional seconds.
//
// https://protobuf.dev/reference/protobuf/google.protobuf/#duration
func (d Duration) toWireFormat() string {
	// We do not use the standard time.Duration.String() and d.Duration.Seconds()
	// method because they use float64 which loses precision.

	// Get the total nanoseconds as a precise integer.
	sign := ""
	if d.Duration < 0 {
		sign = "-"
		d.Duration = -d.Duration
	}

	sec := d.Duration / time.Second
	nsec := d.Duration % time.Second

	if nsec == 0 {
		return fmt.Sprintf("%s%ds", sign, sec)
	}

	frac := strings.TrimRight(fmt.Sprintf("%09d", nsec), "0")
	return fmt.Sprintf("%s%d.%ss", sign, sec, frac)
}

// EncodeValues implements the [query.Encoder] interface by encoding the
// duration as a string, like "3.3s".
func (d Duration) EncodeValues(key string, v *url.Values) error {
	v.Set(key, d.toWireFormat())
	return nil
}

// UnmarshalJSON implements the [json.Unmarshaler] interface. It can parse a
// duration from the Google well-known type format (e.g., "3.123s").
func (d *Duration) UnmarshalJSON(b []byte) error {
	if d == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	// Remove the quotes from the string.
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	dur, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*d = *New(dur)
	return nil
}

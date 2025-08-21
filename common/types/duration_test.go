package types

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"
)

func TestDuration_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		duration Duration
		expected string
		wantErr  bool
	}{
		{
			name:     "zero duration",
			duration: NewDuration(0),
			expected: "0.000000000s",
		},
		{
			name:     "positive duration",
			duration: NewDuration(5 * time.Second),
			expected: "5.000000000s",
		},
		{
			name:     "negative duration",
			duration: NewDuration(-2 * time.Minute),
			expected: "-120.000000000s",
		},
		{
			name:     "negative duration with fractional seconds",
			duration: NewDuration(-2*time.Minute + 100*time.Millisecond),
			expected: "-119.900000000s",
		},
		{
			name:     "fractional seconds",
			duration: NewDuration(1500 * time.Millisecond),
			expected: "1.500000000s",
		},
		{
			name:     "large duration",
			duration: NewDuration(9223372036*time.Second + 854775000*time.Nanosecond),
			expected: "9223372036.854775000s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.duration.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Duration.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(result) != `"`+tt.expected+`"` {
				t.Errorf("Duration.MarshalJSON() = %v, want %v", string(result), `"`+tt.expected+`"`)
			}
		})
	}
}

func TestDuration_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Duration
		wantErr bool
	}{
		{
			name:    "zero duration",
			input:   `"0s"`,
			want:    NewDuration(0),
			wantErr: false,
		},
		{
			name:    "positive duration",
			input:   `"5s"`,
			want:    NewDuration(5 * time.Second),
			wantErr: false,
		},
		{
			name:    "negative duration",
			input:   `"-2s"`,
			want:    NewDuration(-2 * time.Second),
			wantErr: false,
		},
		{
			name:    "negative duration with fractional seconds",
			input:   `"-2.1s"`,
			want:    NewDuration(-2*time.Second - 100*time.Millisecond),
			wantErr: false,
		},
		{
			name:    "fractional seconds",
			input:   `"1.5s"`,
			want:    NewDuration(1500 * time.Millisecond),
			wantErr: false,
		},
		{
			name:    "large duration",
			input:   `"9223372036.854775000s"`,
			want:    NewDuration(9223372036*time.Second + 854775000*time.Nanosecond),
			wantErr: false,
		},
		{
			name:    "invalid duration format",
			input:   `"invalid"`,
			want:    NewDuration(0),
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   `""`,
			want:    NewDuration(0),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d Duration
			err := d.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("Duration.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && d != tt.want {
				t.Errorf("Duration.UnmarshalJSON() = %v, want %v", d, tt.want)
			}
		})
	}
}

func TestDuration_UnmarshalJSON_NilPointer(t *testing.T) {
	var d *Duration
	err := d.UnmarshalJSON([]byte(`"5s"`))
	if err == nil {
		t.Error("Duration.UnmarshalJSON() on nil pointer should return error")
	}
	expectedErr := "json.Unmarshal on nil pointer"
	if err.Error() != expectedErr {
		t.Errorf("Duration.UnmarshalJSON() error = %v, want %v", err.Error(), expectedErr)
	}
}

func TestDuration_EncodeValues(t *testing.T) {
	tests := []struct {
		name     string
		duration Duration
		key      string
		expected string
	}{
		{
			name:     "zero duration",
			duration: NewDuration(0),
			key:      "duration",
			expected: "0.000000000s",
		},
		{
			name:     "positive duration",
			duration: NewDuration(5 * time.Second),
			key:      "timeout",
			expected: "5.000000000s",
		},
		{
			name:     "negative duration",
			duration: NewDuration(-2 * time.Minute),
			key:      "delay",
			expected: "-120.000000000s",
		},
		{
			name:     "fractional seconds",
			duration: NewDuration(1500 * time.Millisecond),
			key:      "interval",
			expected: "1.500000000s",
		},
		{
			name:     "large duration",
			duration: NewDuration(24 * time.Hour),
			key:      "period",
			expected: "86400.000000000s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := url.Values{}
			err := tt.duration.EncodeValues(tt.key, &values)
			if err != nil {
				t.Errorf("Duration.EncodeValues() error = %v", err)
				return
			}
			result := values.Get(tt.key)
			if result != tt.expected {
				t.Errorf("Duration.EncodeValues() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDuration_JSONRoundTrip(t *testing.T) {
	tests := []struct {
		name     string
		duration Duration
	}{
		{
			name:     "zero duration",
			duration: NewDuration(0),
		},
		{
			name:     "positive duration",
			duration: NewDuration(5 * time.Second),
		},
		{
			name:     "negative duration",
			duration: NewDuration(-2 * time.Minute),
		},
		{
			name:     "fractional seconds",
			duration: NewDuration(1500 * time.Millisecond),
		},
		{
			name:     "complex duration",
			duration: NewDuration(1*time.Hour + 2*time.Minute + 3*time.Second),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal to JSON
			jsonData, err := json.Marshal(tt.duration)
			if err != nil {
				t.Errorf("json.Marshal() error = %v", err)
				return
			}

			// Unmarshal from JSON
			var result Duration
			err = json.Unmarshal(jsonData, &result)
			if err != nil {
				t.Errorf("json.Unmarshal() error = %v", err)
				return
			}

			// Check that the round trip preserved the value
			if result != tt.duration {
				t.Errorf("JSON round trip failed: original = %v, result = %v", tt.duration, result)
			}
		})
	}
}

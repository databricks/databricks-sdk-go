package time

import (
	"encoding/json"
	"net/url"
	"testing"
	stdtime "time"
)

func TestNew(t *testing.T) {
	now := stdtime.Now()
	timeVal := New(now)
	if !timeVal.Time.Equal(now) {
		t.Errorf("New() = %v, want %v", timeVal.Time, now)
	}
}

func TestAsTime(t *testing.T) {
	now := stdtime.Now()
	timeVal := New(now)
	result := timeVal.AsTime()
	if !result.Equal(now) {
		t.Errorf("AsTime() = %v, want %v", result, now)
	}
}

func TestAsTime_NilPointer(t *testing.T) {
	var timeVal *Time
	result := timeVal.AsTime()
	expected := stdtime.Time{}
	if !result.Equal(expected) {
		t.Errorf("AsTime() on nil = %v, want %v", result, expected)
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		time     Time
		expected string
		wantErr  bool
	}{
		{
			name:     "zero time",
			time:     *New(stdtime.Time{}),
			expected: `"0001-01-01T00:00:00Z"`,
		},
		{
			name:     "specific time",
			time:     *New(stdtime.Date(2023, 12, 25, 10, 30, 0, 0, stdtime.UTC)),
			expected: `"2023-12-25T10:30:00Z"`,
		},
		{
			name:     "time with nanoseconds",
			time:     *New(stdtime.Date(2023, 12, 25, 10, 30, 0, 123456789, stdtime.UTC)),
			expected: `"2023-12-25T10:30:00.123456789Z"`,
		},
		{
			name:     "time with timezone",
			time:     *New(stdtime.Date(2023, 12, 25, 10, 30, 0, 0, stdtime.FixedZone("EST", -5*3600))),
			expected: `"2023-12-25T10:30:00-05:00"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.time.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(result) != tt.expected {
				t.Errorf("Time.MarshalJSON() = %v, want %v", string(result), tt.expected)
			}
		})
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Time
		wantErr bool
	}{
		{
			name:    "zero time",
			input:   `"0001-01-01T00:00:00Z"`,
			want:    *New(stdtime.Date(1, 1, 1, 0, 0, 0, 0, stdtime.UTC)),
			wantErr: false,
		},
		{
			name:    "specific time",
			input:   `"2023-12-25T10:30:00Z"`,
			want:    *New(stdtime.Date(2023, 12, 25, 10, 30, 0, 0, stdtime.UTC)),
			wantErr: false,
		},
		{
			name:    "time with nanoseconds",
			input:   `"2023-12-25T10:30:00.123456789Z"`,
			want:    *New(stdtime.Date(2023, 12, 25, 10, 30, 0, 123456789, stdtime.UTC)),
			wantErr: false,
		},
		{
			name:    "time with timezone",
			input:   `"2023-12-25T10:30:00-05:00"`,
			want:    *New(stdtime.Date(2023, 12, 25, 10, 30, 0, 0, stdtime.FixedZone("", -5*3600))),
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   `""`,
			want:    *New(stdtime.Time{}),
			wantErr: false,
		},
		{
			name:    "invalid time format",
			input:   `"invalid-time"`,
			want:    *New(stdtime.Time{}),
			wantErr: true,
		},
		{
			name:    "invalid JSON",
			input:   `"invalid-json"`,
			want:    *New(stdtime.Time{}),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var timeVal Time
			err := timeVal.UnmarshalJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !timeVal.Time.Equal(tt.want.Time) {
				t.Errorf("Time.UnmarshalJSON() = %v, want %v", timeVal, tt.want)
			}
		})
	}
}

func TestTime_UnmarshalJSON_NilPointer(t *testing.T) {
	var timeVal *Time
	err := timeVal.UnmarshalJSON([]byte(`"2023-12-25T10:30:00Z"`))
	if err == nil {
		t.Error("Time.UnmarshalJSON() on nil pointer should return error")
	}
	expectedErr := "json.Unmarshal on nil pointer"
	if err.Error() != expectedErr {
		t.Errorf("Time.UnmarshalJSON() error = %v, want %v", err.Error(), expectedErr)
	}
}

func TestTime_EncodeValues(t *testing.T) {
	tests := []struct {
		name     string
		time     Time
		key      string
		expected string
	}{
		{
			name:     "zero time",
			time:     *New(stdtime.Time{}),
			key:      "created_at",
			expected: "0001-01-01T00:00:00Z",
		},
		{
			name:     "specific time",
			time:     *New(stdtime.Date(2023, 12, 25, 10, 30, 0, 0, stdtime.UTC)),
			key:      "updated_at",
			expected: "2023-12-25T10:30:00Z",
		},
		{
			name:     "time with nanoseconds",
			time:     *New(stdtime.Date(2023, 12, 25, 10, 30, 0, 123456789, stdtime.UTC)),
			key:      "timestamp",
			expected: "2023-12-25T10:30:00.123456789Z",
		},
		{
			name:     "time with timezone",
			time:     *New(stdtime.Date(2023, 12, 25, 10, 30, 0, 0, stdtime.FixedZone("EST", -5*3600))),
			key:      "local_time",
			expected: "2023-12-25T10:30:00-05:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := url.Values{}
			err := tt.time.EncodeValues(tt.key, &values)
			if err != nil {
				t.Errorf("Time.EncodeValues() error = %v", err)
				return
			}
			result := values.Get(tt.key)
			if result != tt.expected {
				t.Errorf("Time.EncodeValues() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTime_JSONRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		time Time
	}{
		{
			name: "zero time",
			time: *New(stdtime.Time{}),
		},
		{
			name: "specific time",
			time: *New(stdtime.Date(2023, 12, 25, 10, 30, 0, 0, stdtime.UTC)),
		},
		{
			name: "time with nanoseconds",
			time: *New(stdtime.Date(2023, 12, 25, 10, 30, 0, 123456789, stdtime.UTC)),
		},
		{
			name: "time with timezone",
			time: *New(stdtime.Date(2023, 12, 25, 10, 30, 0, 0, stdtime.FixedZone("EST", -5*3600))),
		},
		{
			name: "current time",
			time: *New(stdtime.Now()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal to JSON
			jsonData, err := json.Marshal(tt.time)
			if err != nil {
				t.Errorf("json.Marshal() error = %v", err)
				return
			}

			// Unmarshal from JSON
			var result Time
			err = json.Unmarshal(jsonData, &result)
			if err != nil {
				t.Errorf("json.Unmarshal() error = %v", err)
				return
			}

			// Check that the round trip preserved the value
			if !result.Time.Equal(tt.time.Time) {
				t.Errorf("JSON round trip failed: original = %v, result = %v", tt.time, result)
			}
		})
	}
}

func TestTime_EdgeCases(t *testing.T) {
	t.Run("very old time", func(t *testing.T) {
		oldTime := stdtime.Date(1000, 1, 1, 0, 0, 0, 0, stdtime.UTC)
		timeVal := *New(oldTime)

		jsonData, err := json.Marshal(timeVal)
		if err != nil {
			t.Errorf("json.Marshal() error = %v", err)
			return
		}

		var result Time
		err = json.Unmarshal(jsonData, &result)
		if err != nil {
			t.Errorf("json.Unmarshal() error = %v", err)
			return
		}

		if !result.Time.Equal(oldTime) {
			t.Errorf("Time round trip failed: original = %v, result = %v", oldTime, result)
		}
	})

	t.Run("very future time", func(t *testing.T) {
		futureTime := stdtime.Date(9999, 12, 31, 23, 59, 59, 999999999, stdtime.UTC)
		timeVal := *New(futureTime)

		jsonData, err := json.Marshal(timeVal)
		if err != nil {
			t.Errorf("json.Marshal() error = %v", err)
			return
		}

		var result Time
		err = json.Unmarshal(jsonData, &result)
		if err != nil {
			t.Errorf("json.Unmarshal() error = %v", err)
			return
		}

		if !result.Time.Equal(futureTime) {
			t.Errorf("Time round trip failed: original = %v, result = %v", futureTime, result)
		}
	})

	t.Run("leap year time", func(t *testing.T) {
		leapTime := stdtime.Date(2024, 2, 29, 12, 0, 0, 0, stdtime.UTC)
		timeVal := *New(leapTime)

		jsonData, err := json.Marshal(timeVal)
		if err != nil {
			t.Errorf("json.Marshal() error = %v", err)
			return
		}

		var result Time
		err = json.Unmarshal(jsonData, &result)
		if err != nil {
			t.Errorf("json.Unmarshal() error = %v", err)
			return
		}

		if !result.Time.Equal(leapTime) {
			t.Errorf("Time round trip failed: original = %v, result = %v", leapTime, result)
		}
	})
}

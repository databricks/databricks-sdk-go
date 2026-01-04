package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestConfigFile_Configure_ListParsing tests that comma-separated list values
// in configuration files are correctly parsed into slices.
func TestConfigFile_Configure_ListParsing(t *testing.T) {
	testCases := []struct {
		name    string
		profile string
		want    []string
	}{
		{
			name:    "single item",
			profile: "single-item",
			want:    []string{"clusters"},
		},
		{
			name:    "multiple items",
			profile: "multiple-items",
			want:    []string{"alpha", "beta", "gamma"},
		},
		{
			name:    "whitespace around items is trimmed",
			profile: "whitespace-around-items",
			want:    []string{"alpha", "beta", "gamma"},
		},
		{
			name:    "empty items are filtered out",
			profile: "empty-items-filtered",
			want:    []string{"alpha", "beta"},
		},
		{
			name:    "whitespace-only items are filtered out",
			profile: "whitespace-only-items-filtered",
			want:    []string{"alpha", "beta"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			withMockEnv(t, map[string]string{})

			cfg := &Config{
				Profile:    tc.profile,
				ConfigFile: "testdata/list-parsing/.databrickscfg",
			}
			err := ConfigFile.Configure(cfg)
			if err != nil {
				t.Fatalf("Configure failed: %v", err)
			}
			if diff := cmp.Diff(tc.want, cfg.Scopes); diff != "" {
				t.Errorf("list mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

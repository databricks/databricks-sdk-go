package config

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigFileLoad(t *testing.T) {
	f, err := LoadFile("testdata/.databrickscfg")
	require.NoError(t, err)
	assert.NotNil(t, f)

	for _, name := range []string{
		"password-with-double-quotes",
		"password-with-single-quotes",
		"password-without-quotes",
	} {
		section := f.Section(name)
		require.NotNil(t, section)
		assert.Equal(t, "%Y#X$Z", section.Key("password").String())
	}
}

func TestConfigFile_Scopes(t *testing.T) {
	tests := []struct {
		name     string
		profile  string
		expected []string
	}{
		{
			name:     "empty defaults to all-apis",
			profile:  "scope-empty",
			expected: []string{"all-apis"},
		},
		{
			name:     "single scope",
			profile:  "scope-single",
			expected: []string{"clusters"},
		},
		{
			name:     "multiple scopes sorted",
			profile:  "scope-multiple",
			expected: []string{"clusters", "files:read", "iam:read", "jobs", "mlflow", "model-serving", "pipelines"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env.CleanupEnvironment(t)
			t.Setenv("HOME", "testdata")

			cfg := &Config{Profile: tt.profile}
			err := cfg.EnsureResolved()
			if err != nil {
				t.Fatalf("EnsureResolved failed: %v", err)
			}
			if diff := cmp.Diff(tt.expected, cfg.GetScopes()); diff != "" {
				t.Errorf("GetScopes mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// withMockEnv replaces the getenv function with one that returns values from
// the provided map. The original function is restored when the test completes.
func withMockEnv(t *testing.T, env map[string]string) {
	original := getenv
	t.Cleanup(func() { getenv = original })
	getenv = func(key string) string {
		return env[key]
	}
}

func TestConfigFileLoad(t *testing.T) {
	f, err := LoadFile("testdata/.databrickscfg")
	if err != nil {
		t.Fatalf("LoadFile failed: %v", err)
	}
	if f == nil {
		t.Fatal("expected file to be non-nil")
	}

	for _, name := range []string{
		"password-with-double-quotes",
		"password-with-single-quotes",
		"password-without-quotes",
	} {
		section := f.Section(name)
		if section == nil {
			t.Fatalf("expected section %q to be non-nil", name)
		}
		if got, want := section.Key("password").String(), "%Y#X$Z"; got != want {
			t.Errorf("password mismatch for %q: got %q, want %q", name, got, want)
		}
	}
}

func TestConfigFile_Scopes(t *testing.T) {
	tests := []struct {
		name    string
		profile string
		want    []string
	}{
		{
			name:    "empty defaults to all-apis",
			profile: "scope-empty",
			want:    []string{"all-apis"},
		},
		{
			name:    "single scope",
			profile: "scope-single",
			want:    []string{"clusters"},
		},
		{
			name:    "multiple scopes sorted",
			profile: "scope-multiple",
			want:    []string{"clusters", "files:read", "iam:read", "jobs", "mlflow", "model-serving", "pipelines"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withMockEnv(t, map[string]string{
				"HOME": "testdata/scopes",
			})

			cfg := &Config{Profile: tt.profile}
			err := cfg.EnsureResolved()
			if err != nil {
				t.Fatalf("EnsureResolved failed: %v", err)
			}
			if diff := cmp.Diff(tt.want, cfg.GetScopes()); diff != "" {
				t.Errorf("GetScopes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

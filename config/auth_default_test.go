package config

import (
	"context"
	"strings"
	"testing"
)

func TestDefaultCredentials_Configure(t *testing.T) {
	testCases := []struct {
		desc     string
		authType string
		wantErr  string
	}{
		{
			desc:     "unknown auth type",
			authType: "unknown-auth-type-1337",
			wantErr:  "auth type \"unknown-auth-type-1337\" not found",
		},
		{
			desc:     "not valid auth",
			authType: "",
			wantErr:  "cannot configure default credentials",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := context.Background()
			cfg := &Config{
				AuthType: tc.authType,
				resolved: true, // avoid calling EnsureResolved
			}

			dc := DefaultCredentials{}
			got, gotErr := dc.Configure(ctx, cfg)

			if got != nil {
				t.Errorf("DefaultCredentials.Configure: got %v, want nil", got)
			}
			if gotErr == nil {
				t.Errorf("DefaultCredentials.Configure: got error %v, want non-nil", gotErr)
			}
			if !strings.Contains(gotErr.Error(), tc.wantErr) {
				t.Errorf("DefaultCredentials.Configure: got error %v, want error containing %q", gotErr, tc.wantErr)
			}
		})
	}
}

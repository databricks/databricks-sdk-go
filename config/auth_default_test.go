package config

import (
	"context"
	"strings"
	"testing"
)

func TestDefaultCredentials_Configure_unknownAuthType(t *testing.T) {
	ctx := context.Background()
	cfg := &Config{
		AuthType: "unknown-mode-for-test",
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
	if !strings.Contains(gotErr.Error(), "auth type \"unknown-mode-for-test\" not found") {
		t.Errorf("DefaultCredentials.Configure: got error %v, want error containing \"auth type \"unknown-mode-for-test\" not found\"", gotErr)
	}
}

func TestDefaultCredentials_Configure_noValidAuth(t *testing.T) {
	ctx := context.Background()
	cfg := &Config{
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
	if !strings.Contains(gotErr.Error(), "cannot configure default credentials") {
		t.Errorf("DefaultCredentials.Configure: got error %v, want error containing \"cannot configure default credentials\"", gotErr)
	}
}
